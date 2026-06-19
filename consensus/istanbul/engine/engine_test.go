package engine

import (
	"bytes"
	"math"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/consensus/istanbul/validator"
	"github.com/ethereum/go-ethereum/contracts"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/triedb"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var validators = []common.Address{
	common.BytesToAddress(hexutil.MustDecode("0x44add0ec310f115a0e603b2d7db9f067778eaf8a")),
	common.BytesToAddress(hexutil.MustDecode("0x294fc7e8f22b3bcdcf955dd7ff3ba2ed833f8212")),
	common.BytesToAddress(hexutil.MustDecode("0x6beaaed781d2d2ab6350f5c4566a2c6eaac407a6")),
	common.BytesToAddress(hexutil.MustDecode("0x8be76812f765c24641ec63dc2852b378aba2b440")),
}

func TestPrepareExtra(t *testing.T) {
	expectedResult := hexutil.MustDecode("0xf87ba00000000000000000000000000000000000000000000000000000000000000000f8549444add0ec310f115a0e603b2d7db9f067778eaf8a94294fc7e8f22b3bcdcf955dd7ff3ba2ed833f8212946beaaed781d2d2ab6350f5c4566a2c6eaac407a6948be76812f765c24641ec63dc2852b378aba2b440c080c080")

	h := &types.Header{}
	err := ApplyHeaderIstanbulExtra(
		h,
		WriteValidators(validators),
	)
	require.NoError(t, err)
	assert.Equal(t, h.Extra, expectedResult)
}

// ##CROSS: consensus system contract
func TestIsSystemTransaction(t *testing.T) {
	key, err := crypto.GenerateKey()
	require.NoError(t, err)

	to := contracts.ValidatorSetAddr
	header := &types.Header{Coinbase: crypto.PubkeyToAddress(key.PublicKey)}
	signer := types.LatestSignerForChainID(params.TestChainConfig.ChainID)

	legacyTx := types.MustSignNewTx(key, signer, &types.LegacyTx{
		Nonce:    0,
		GasPrice: big.NewInt(0),
		Gas:      21000,
		To:       &to,
	})
	isSystemTx, err := IsSystemTransaction(legacyTx, header)
	require.NoError(t, err)
	assert.True(t, isSystemTx)

	dynamicFeeTx := types.MustSignNewTx(key, signer, &types.DynamicFeeTx{
		ChainID:   params.TestChainConfig.ChainID,
		Nonce:     1,
		GasTipCap: big.NewInt(0),
		GasFeeCap: big.NewInt(0),
		Gas:       21000,
		To:        &to,
	})
	isSystemTx, err = IsSystemTransaction(dynamicFeeTx, header)
	require.NoError(t, err)
	assert.False(t, isSystemTx)
}

// newTestStateDB creates an empty in-memory state database for testing.
func newTestStateDB(t *testing.T) *state.StateDB {
	t.Helper()
	memdb := rawdb.NewMemoryDatabase()
	tdb := triedb.NewDatabase(memdb, nil)
	sdb := state.NewDatabase(tdb, nil)
	statedb, err := state.New(types.EmptyRootHash, sdb)
	require.NoError(t, err)
	return statedb
}

// TestApplySystemTransaction verifies that replaying a system transaction keeps
// consensus execution semantics and results in the same state as the consensus path.
func TestApplySystemTransaction(t *testing.T) {
	key, err := crypto.GenerateKey()
	require.NoError(t, err)

	var (
		sender      = crypto.PubkeyToAddress(key.PublicKey)
		to          = contracts.ValidatorSetAddr
		value       = big.NewInt(100)
		data        = []byte("system call data")
		chainConfig = params.TestChainConfig
		signer      = types.LatestSignerForChainID(chainConfig.ChainID)
		funds       = uint256.NewInt(1_000_000)
	)
	header := &types.Header{
		Number:     big.NewInt(100),
		Coinbase:   sender,
		BaseFee:    big.NewInt(1_000_000_000), // non-zero to prove no fee is charged
		GasLimit:   10_000_000,
		Difficulty: big.NewInt(0),
		Time:       1_000_000_000,
	}
	tx := types.MustSignNewTx(key, signer, &types.LegacyTx{
		Nonce:    0,
		GasPrice: big.NewInt(0),
		Gas:      math.MaxUint64 / 2,
		To:       &to,
		Value:    value,
		Data:     data,
	})
	isSystemTx, err := IsSystemTransaction(tx, header)
	require.NoError(t, err)
	require.True(t, isSystemTx)

	// Replay the transaction with consensus execution semantics
	replayDB := newTestStateDB(t)
	replayDB.AddBalance(sender, funds, tracing.BalanceChangeUnspecified)
	blockCtx := core.NewEVMBlockContext(header, &chainMock{config: chainConfig}, &header.Coinbase)
	evm := vm.NewEVM(blockCtx, replayDB, chainConfig, vm.Config{})

	engine := &Engine{}
	require.NoError(t, engine.ApplySystemTransaction(evm, header, tx, 0))

	// No intrinsic gas or fee is charged; only the value is transferred
	assert.Equal(t, uint64(1), replayDB.GetNonce(sender))
	assert.Equal(t, uint256.NewInt(100), replayDB.GetBalance(to))
	assert.Equal(t, new(uint256.Int).Sub(funds, uint256.NewInt(100)), replayDB.GetBalance(sender))

	// Apply the same transaction through the consensus path and compare the state
	consensusDB := newTestStateDB(t)
	consensusDB.AddBalance(sender, funds, tracing.BalanceChangeUnspecified)
	var (
		txs      []*types.Transaction
		receipts []*types.Receipt
		usedGas  uint64
	)
	msg := newSystemMessage(sender, to, data, value)
	err = engine.applySystemTransaction(msg, tx, consensusDB, header, &chainMock{config: chainConfig}, header.Coinbase, &txs, &receipts, &usedGas, nil)
	require.NoError(t, err)

	// Both state should be equal
	deleteEmptyObjects := chainConfig.IsEIP158(header.Number)
	assert.Equal(t, consensusDB.IntermediateRoot(deleteEmptyObjects), replayDB.IntermediateRoot(deleteEmptyObjects))
}

// ##

// ##CROSS: bls seal
type signedSeal struct {
	index     uint
	signature []byte
}

func (c signedSeal) Index() uint            { return c.index }
func (c signedSeal) Signer() common.Address { return common.Address{} }
func (c signedSeal) Signature() []byte      { return c.signature }

type chainMock struct {
	config *params.ChainConfig
}

func (c *chainMock) Engine() consensus.Engine                                { return nil }
func (c *chainMock) Config() *params.ChainConfig                             { return c.config }
func (c *chainMock) CurrentHeader() *types.Header                            { return nil }
func (c *chainMock) GetHeader(hash common.Hash, number uint64) *types.Header { return nil }
func (c *chainMock) GetHeaderByNumber(number uint64) *types.Header           { return nil }
func (c *chainMock) GetHeaderByHash(hash common.Hash) *types.Header          { return nil }
func (c *chainMock) GetTd(hash common.Hash, number uint64) *big.Int          { return nil }

// ##

func TestWriteCommittedSeals(t *testing.T) {
	istRawData := hexutil.MustDecode("0xf89f80f8549444add0ec310f115a0e603b2d7db9f067778eaf8a94294fc7e8f22b3bcdcf955dd7ff3ba2ed833f8212946beaaed781d2d2ab6350f5c4566a2c6eaac407a6948be76812f765c24641ec63dc2852b378aba2b440c080f843b841010203000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000080")
	expectedCommittedSeal := append([]byte{1, 2, 3}, bytes.Repeat([]byte{0x00}, types.IstanbulExtraSeal-3)...)
	expectedIstExtra := &types.IstanbulExtra{
		VanityData:    []byte{},
		Validators:    validators,
		CommittedSeal: [][]byte{expectedCommittedSeal},
		Round:         0,
		Vote:          nil,
		RandomReveal:  []byte{},
	}

	h := &types.Header{
		Extra: istRawData,
	}

	// normal case
	err := ApplyHeaderIstanbulExtra(
		h,
		writeCommittedSeals(nil, h, []istanbul.SignedSeal{signedSeal{0, expectedCommittedSeal}}),
	)
	require.NoError(t, err)

	// verify istanbul extra-data
	istExtra, err := getExtra(h)
	require.NoError(t, err)
	assert.Equal(t, istExtra, expectedIstExtra)

	// invalid seal
	unexpectedCommittedSeal := append(expectedCommittedSeal, make([]byte, 1)...)
	err = ApplyHeaderIstanbulExtra(
		h,
		writeCommittedSeals(nil, h, []istanbul.SignedSeal{signedSeal{0, unexpectedCommittedSeal}}),
	)
	assert.ErrorIs(t, err, istanbul.ErrInvalidCommittedSeals)
}

// ##CROSS: bls seal
func TestWriteCommittedSealsBLS(t *testing.T) {
	committedSeals := []istanbul.SignedSeal{
		signedSeal{0, hexutil.MustDecode("0x86e181c931b609a0e718a53f92bee1952bf8a4c59a5df4779d1d03917f68c16ca7a231011920c6462a5f5fc9e4a1674e0b09f31805c1a3dc2f2584624c69d96394b1946bf7ff4f78be9577aee4c7a8ae35138aada623fb17cc19d5ce81442e37")},
		signedSeal{1, hexutil.MustDecode("0xb07aa34a480b2ebdad484311dc71eb75e26fb8bea8aa572946faa98521a2890922396e78655f6c77b9c63b55f5867ed003683f4a4ecd76fe331f5902fc67cf3536a84313b4e7841085bffa1e632b739a582ee46fe5d357ea3d4f38d724dfca02")},
		signedSeal{2, hexutil.MustDecode("0x82e4a75a84ead6b7af96d24816b3d75cd531117fcee36a430d992a3118df521bcfcbd0e6e0d22d2911548e4a5d1d89a308ddbfe71db361e3dfcd9d95dc0b575b2d024d9f28b7712f46c7b6cdc4b8ff62f48859df943e55263e727384c81bf8f7")},
	}
	expectedCommittedSeal := hexutil.MustDecode("0xad370a3cb337a6f2ed0f2084cf643cb066a9db4f5b77b6c72457f1f4f0e8f82f98130d66a686571b5088ea832080f89d1567c18c7a0cd697ae619ae20dc58a950524e7ff89563fc8f2ee761d9e34fabe1fb7260d75ddb3f85449c7aa043a8002")
	expectedIstExtra := &types.IstanbulExtra{
		VanityData:    bytes.Repeat([]byte{0x00}, types.IstanbulExtraVanity),
		Validators:    []common.Address{},
		CommittedSeal: [][]byte{expectedCommittedSeal},
		RandomReveal:  []byte{},
		SignersBitset: []uint64{7},
		Signers:       []types.BLSPublicKey{},
	}

	h := &types.Header{
		Extra:  []byte{},
		Number: big.NewInt(10),
		Time:   1000,
	}

	ts := uint64(0)
	chain := &chainMock{
		config: &params.ChainConfig{
			LondonBlock:    big.NewInt(0),
			BreakpointTime: &ts,
			Istanbul: &params.IstanbulConfig{
				PoSA: &params.PoSAConfig{},
			},
		},
	}

	err := ApplyHeaderIstanbulExtra(
		h,
		writeCommittedSeals(chain, h, committedSeals),
	)
	require.NoError(t, err)

	// verify istanbul extra-data
	istExtra, err := getExtra(h)
	require.NoError(t, err)
	assert.Equal(t, expectedIstExtra, istExtra)
}

func TestWriteSigners(t *testing.T) {
	signers := []types.BLSPublicKey{
		types.BytesToBLSPublicKey(signer1),
		types.BytesToBLSPublicKey(signer2),
		types.BytesToBLSPublicKey(signer3),
		types.BytesToBLSPublicKey(signer4),
	}
	expectedIstExtra := &types.IstanbulExtra{
		VanityData:    bytes.Repeat([]byte{0x00}, types.IstanbulExtraVanity),
		Validators:    validators,
		CommittedSeal: [][]byte{},
		RandomReveal:  []byte{},
		SignersBitset: []uint64{},
		Signers:       signers,
	}

	h := &types.Header{
		Extra:  []byte{},
		Number: big.NewInt(10),
		Time:   1000,
	}

	err := ApplyHeaderIstanbulExtra(
		h,
		WriteValidators(validators),
		WriteSigners(signers),
	)
	require.NoError(t, err)

	// verify istanbul extra-data
	istExtra, err := getExtra(h)
	require.NoError(t, err)
	assert.Equal(t, expectedIstExtra, istExtra)
}

func TestBLSSigners(t *testing.T) {
	t.Run("uses byte sorted validator order", func(t *testing.T) {
		addrs := []common.Address{
			common.HexToAddress("0xc53f2189bf6d7bf56722731787127f90d319e112"),
			common.HexToAddress("0xed2d479591fe2c5626ce09bca4ed2a62e00e5bc2"),
			common.HexToAddress("0xc8417f834995aaeb35f342a67a4961e19cd4735c"),
			common.HexToAddress("0x784ae51f5013b51c8360afdf91c6bc5a16f586ea"),
			common.HexToAddress("0xecf0974e6f0630fd91ea4da8399cdb3f59e5220f"),
			common.HexToAddress("0x411c4d11acd714b82a5242667e36de14b9e1d10b"),
		}
		signers := make([]types.BLSPublicKey, len(addrs))
		for i := range signers {
			signers[i] = types.BytesToBLSPublicKey([]byte{byte(i + 1)})
		}

		stringPolicy := istanbul.NewProposerPolicyByIdAndSortFunc(istanbul.RoundRobin, istanbul.ValidatorSortByString())
		valSet := validator.NewSet(addrs, signers, stringPolicy)
		require.Equal(t, addrs[1], valSet.List()[3].Address(), "test fixture must differ between string and byte order")
		h := &types.Header{
			Extra:  []byte{},
			Number: big.NewInt(10),
			Time:   1000,
		}
		err := ApplyHeaderIstanbulExtra(
			h,
			WriteValidators(addrs),
			WriteSigners(signers),
			func(extra *types.IstanbulExtra) error {
				extra.SignersBitset = []uint64{1 << 3}
				return nil
			},
		)
		require.NoError(t, err)

		gotAddrs, gotPubkeys, err := (Engine{}).BLSSigners(h, valSet)
		require.NoError(t, err)
		require.Len(t, gotAddrs, 1)
		require.Len(t, gotPubkeys, 1)
		assert.Equal(t, addrs[2], gotAddrs[0])
		assert.Equal(t, signers[2], gotPubkeys[0])
	})

	t.Run("invalid signers bitset", func(t *testing.T) {
		signers := []types.BLSPublicKey{
			types.BytesToBLSPublicKey(signer1),
			types.BytesToBLSPublicKey(signer2),
			types.BytesToBLSPublicKey(signer3),
			types.BytesToBLSPublicKey(signer4),
		}
		valSet := validator.NewSet(validators, signers, istanbul.NewRoundRobinProposerPolicy())

		tests := []struct {
			name          string
			signersBitset []uint64
		}{
			{
				name:          "out of range set bit",
				signersBitset: []uint64{1 << uint(len(validators))},
			},
			{
				name:          "trailing zero word",
				signersBitset: []uint64{1, 0},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				h := &types.Header{
					Extra:  []byte{},
					Number: big.NewInt(10),
					Time:   1000,
				}
				err := ApplyHeaderIstanbulExtra(
					h,
					WriteValidators(validators),
					WriteSigners(signers),
					func(extra *types.IstanbulExtra) error {
						extra.SignersBitset = tt.signersBitset
						return nil
					},
				)
				require.NoError(t, err)

				_, _, err = (Engine{}).BLSSigners(h, valSet)
				assert.ErrorIs(t, err, istanbul.ErrInvalidSignersBitset)
			})
		}
	})

	t.Run("sparse signers bitset", func(t *testing.T) {
		addrs := make([]common.Address, 100)
		signers := make([]types.BLSPublicKey, 100)
		for i := range addrs {
			addrs[i] = common.BytesToAddress([]byte{byte(i + 1)})
			signers[i] = types.BytesToBLSPublicKey([]byte{byte(i + 1)})
		}

		h := &types.Header{
			Extra:  []byte{},
			Number: big.NewInt(10),
			Time:   1000,
		}
		err := ApplyHeaderIstanbulExtra(
			h,
			WriteValidators(addrs),
			WriteSigners(signers),
			func(extra *types.IstanbulExtra) error {
				extra.SignersBitset = []uint64{0, uint64(1) << 35}
				return nil
			},
		)
		require.NoError(t, err)

		valSet := validator.NewSet(addrs, signers, istanbul.NewRoundRobinProposerPolicy())
		expectedSigner := valSet.List()[99]

		gotAddrs, gotPubkeys, err := (Engine{}).BLSSigners(h, valSet)
		require.NoError(t, err)
		require.Len(t, gotAddrs, 1)
		require.Len(t, gotPubkeys, 1)
		assert.Equal(t, expectedSigner.Address(), gotAddrs[0])
		assert.Equal(t, expectedSigner.SignerAddress(), gotPubkeys[0])
	})
}

// ##

func TestWriteRoundNumber(t *testing.T) {
	istRawData := hexutil.MustDecode("0xf85b80f8549444add0ec310f115a0e603b2d7db9f067778eaf8a94294fc7e8f22b3bcdcf955dd7ff3ba2ed833f8212946beaaed781d2d2ab6350f5c4566a2c6eaac407a6948be76812f765c24641ec63dc2852b378aba2b440c005c080")
	expectedIstExtra := &types.IstanbulExtra{
		VanityData:    []byte{},
		Validators:    validators,
		CommittedSeal: [][]byte{},
		Round:         5,
		Vote:          nil,
		RandomReveal:  []byte{},
	}

	h := &types.Header{
		Extra: istRawData,
	}

	// normal case
	err := ApplyHeaderIstanbulExtra(
		h,
		writeRoundNumber(big.NewInt(5)),
	)
	require.NoError(t, err)

	// verify istanbul extra-data
	istExtra, err := getExtra(h)
	require.NoError(t, err)
	assert.Equal(t, istExtra, expectedIstExtra)
}

func TestWriteValidatorVote(t *testing.T) {
	vanity := bytes.Repeat([]byte{0x00}, types.IstanbulExtraVanity)
	istRawData := hexutil.MustDecode("0xf83da00000000000000000000000000000000000000000000000000000000000000000c0d79444add0ec310f115a0e603b2d7db9f0677712345681ff80c080")
	vote := &types.ValidatorVote{RecipientAddress: common.BytesToAddress(hexutil.MustDecode("0x44add0ec310f115a0e603b2d7db9f06777123456")), VoteType: types.IstanbulAuthVote}
	expectedIstExtra := &types.IstanbulExtra{
		VanityData:    vanity,
		Validators:    []common.Address{},
		CommittedSeal: [][]byte{},
		Round:         0,
		Vote:          vote,
	}

	h := &types.Header{
		Extra: istRawData,
	}

	// normal case
	err := ApplyHeaderIstanbulExtra(
		h,
		WriteVote(common.BytesToAddress(hexutil.MustDecode("0x44add0ec310f115a0e603b2d7db9f06777123456")), true),
	)
	require.NoError(t, err)

	// verify istanbul extra-data
	istExtra, err := getExtra(h)
	require.NoError(t, err)
	assert.Equal(t, istExtra.Vote, expectedIstExtra.Vote)
}
