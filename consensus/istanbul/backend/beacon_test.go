package backend

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/beacon"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/consensus/misc/eip1559"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ##CROSS: istanbul

// newBeaconTestChain creates an isolated Istanbul chain and a signed, unimported block.
func newBeaconTestChain(t *testing.T) (*core.BlockChain, *types.Block) {
	t.Helper()

	key, err := crypto.GenerateKey()
	require.NoError(t, err)

	address := crypto.PubkeyToAddress(key.PublicKey)
	config := *params.TestChainConfig
	config.Ethash = nil
	config.Istanbul = &params.IstanbulConfig{EpochLength: 30000, BlockPeriodSeconds: 1, Validators: []common.Address{address}}
	config.ShanghaiTime = new(uint64)
	extra, err := rlp.EncodeToBytes(&types.IstanbulExtra{VanityData: make([]byte, types.IstanbulExtraVanity), Validators: []common.Address{address}})
	require.NoError(t, err)

	genesis := &core.Genesis{
		Config: &config, ExtraData: extra, Mixhash: types.IstanbulDigest,
		Difficulty: big.NewInt(1), GasLimit: 10_000_000, BaseFee: big.NewInt(params.InitialBaseFee),
	}
	db := rawdb.NewMemoryDatabase()
	t.Cleanup(func() { db.Close() })
	backend := New(&config, key, nil, db, nil)
	chain, err := core.NewBlockChain(db, genesis, beacon.New(backend), nil)
	require.NoError(t, err)
	t.Cleanup(chain.Stop)

	header := makeHeader(chain.Genesis(), backend.config, chain.Config())
	require.NoError(t, backend.Prepare(chain, header))

	// Prepare uses wall-clock time; leave room for a child without making it a future block.
	header.Time = chain.Genesis().Time() + 1
	header.Coinbase = address
	statedb, err := chain.State()
	require.NoError(t, err)

	block, _, err := backend.FinalizeAndAssemble(chain, header, statedb, &types.Body{}, nil, nil)
	require.NoError(t, err)

	header = block.Header()
	require.NoError(t, backend.Engine().CommitHeader(chain, header, []istanbul.SignedSeal{commitSeal(t, key, block)}, big.NewInt(0)))
	return chain, block.WithSeal(header)
}

type beaconTestChainReader struct {
	consensus.ChainHeaderReader
	config *params.ChainConfig
}

// Config returns the configuration under test without changing the underlying chain.
func (c beaconTestChainReader) Config() *params.ChainConfig { return c.config }

// postMergeTestHeader keeps the header valid under PoS rules without any Istanbul seal.
func postMergeTestHeader(header *types.Header) *types.Header {
	header = types.CopyHeader(header)
	header.Difficulty = new(big.Int)
	header.Extra = nil
	header.MixDigest = common.Hash{}
	header.WithdrawalsHash = &types.EmptyWithdrawalsHash
	return header
}

func TestBeacon_VerifyHeader(t *testing.T) {
	chain, block := newBeaconTestChain(t)
	engine := chain.Engine()
	t.Run("valid Istanbul header", func(t *testing.T) {
		require.NoError(t, engine.VerifyHeader(chain, block.Header()))
	})
	t.Run("reject PoS header on Istanbul", func(t *testing.T) {
		header := postMergeTestHeader(block.Header())
		config := *chain.Config()
		config.Istanbul = nil
		// The same header must pass PoS checks, so an unrelated format error cannot hide a routing regression.
		require.NoError(t, engine.VerifyHeader(beaconTestChainReader{chain, &config}, header))
		require.ErrorIs(t, engine.VerifyHeader(chain, header), istanbul.ErrInvalidExtraDataFormat)
	})
}

func TestBeacon_VerifyHeaders(t *testing.T) {
	chain, block := newBeaconTestChain(t)
	child := postMergeTestHeader(block.Header())
	child.ParentHash = block.Hash()
	child.Number = big.NewInt(2)
	child.Time++
	child.BaseFee = eip1559.CalcBaseFee(chain.Config(), block.Header())
	for _, test := range []struct {
		name    string
		headers []*types.Header
		bad     int
	}{
		{"valid Istanbul", []*types.Header{block.Header()}, -1},
		{"PoS only", []*types.Header{postMergeTestHeader(block.Header())}, 0},
		{"mixed batch", []*types.Header{block.Header(), child}, 1},
	} {
		t.Run(test.name, func(t *testing.T) {
			abort, results := chain.Engine().VerifyHeaders(chain, test.headers)
			defer close(abort)
			for i := range test.headers {
				select {
				case err := <-results:
					if i == test.bad {
						require.ErrorIs(t, err, istanbul.ErrInvalidExtraDataFormat)
					} else {
						require.NoError(t, err)
					}
				case <-time.After(2 * time.Second):
					t.Fatalf("missing verification result for header %d", i)
				}
			}
		})
	}
}

func TestBeacon_IstanbulFinalization(t *testing.T) {
	for _, method := range []string{"Finalize", "FinalizeAndAssemble"} {
		t.Run(method, func(t *testing.T) {
			for _, test := range []struct {
				name     string
				istanbul bool
				shanghai bool
				balance  uint64
			}{
				{"Istanbul", true, true, 0},
				{"Istanbul before Shanghai", true, false, 0},
				{"PoS withdrawals unchanged", false, true, 7 * params.GWei},
			} {
				t.Run(test.name, func(t *testing.T) {
					chain, block := newBeaconTestChain(t)
					config := *chain.Config()
					if !test.istanbul {
						config.Istanbul = nil
					}
					if !test.shanghai {
						config.ShanghaiTime = nil
					}
					reader := beaconTestChainReader{chain, &config}
					header := postMergeTestHeader(block.Header())
					recipient := common.Address{0xaa}
					body := &types.Body{Withdrawals: []*types.Withdrawal{{Address: recipient, Amount: 7}}}
					statedb, err := chain.State()
					require.NoError(t, err)
					if method == "Finalize" {
						var receipts types.Receipts
						err = chain.Engine().Finalize(reader, header, statedb, body, &body.Transactions, &receipts, nil, &header.GasUsed, nil)
					} else {
						_, _, err = chain.Engine().FinalizeAndAssemble(reader, header, statedb, body, nil, nil)
					}
					require.NoError(t, err)
					assert.Equal(t, uint256.NewInt(test.balance), statedb.GetBalance(recipient))
				})
			}
		})
	}
}

func TestBlockChain_InsertChainIstanbul(t *testing.T) {
	for _, test := range []struct {
		name string
		bad  int
	}{
		{"valid block", -1},
		{"PoS at index zero", 0},
		{"PoS at index one", 1},
	} {
		t.Run(test.name, func(t *testing.T) {
			chain, block := newBeaconTestChain(t)
			blocks := types.Blocks{block}
			if test.bad >= 0 {
				header := postMergeTestHeader(block.Header())
				if test.bad == 1 {
					header.ParentHash = block.Hash()
					header.Number = big.NewInt(2)
					header.Time++
					header.BaseFee = eip1559.CalcBaseFee(chain.Config(), block.Header())
					blocks = append(blocks, nil)
				}
				blocks[test.bad] = types.NewBlockWithHeader(header).WithBody(types.Body{Withdrawals: []*types.Withdrawal{}})
			}
			index, err := chain.InsertChain(blocks)
			if test.bad < 0 {
				require.NoError(t, err)
				assert.Equal(t, block.Hash(), chain.CurrentBlock().Hash())
				return
			}
			require.ErrorContains(t, err, "rejecting post-merge (difficulty 0)")
			assert.Equal(t, test.bad, index)
			assert.Equal(t, chain.Genesis().Hash(), chain.CurrentBlock().Hash())
			for _, rejected := range blocks {
				assert.False(t, chain.HasBlock(rejected.Hash(), rejected.NumberU64()))
			}
		})
	}
}

func TestBlockChain_WriteBlockAndSetHeadIstanbul(t *testing.T) {
	for _, test := range []struct {
		name       string
		difficulty int64
	}{
		{"finalize Istanbul block", 1},
		{"do not finalize PoS block", 0},
	} {
		t.Run(test.name, func(t *testing.T) {
			chain, block := newBeaconTestChain(t)
			chain.SetSafe(chain.Genesis().Header())
			chain.SetFinalized(chain.Genesis().Header())
			header := block.Header()
			header.Difficulty = big.NewInt(test.difficulty)
			block = block.WithSeal(header)
			statedb, err := chain.State()
			require.NoError(t, err)

			// Exercise the write path directly; InsertChain rejects difficulty zero before reaching this guard.
			status, err := chain.WriteBlockAndSetHead(block, nil, nil, statedb, false, false)
			require.NoError(t, err)
			assert.Equal(t, core.CanonStatTy, status)
			assert.Equal(t, block.Hash(), chain.CurrentBlock().Hash())

			want := chain.Genesis().Hash()
			if test.difficulty > 0 {
				want = block.Hash()
			}
			require.NotNil(t, chain.CurrentSafeBlock())
			require.NotNil(t, chain.CurrentFinalBlock())
			assert.Equal(t, want, chain.CurrentSafeBlock().Hash())
			assert.Equal(t, want, chain.CurrentFinalBlock().Hash())
		})
	}
}
