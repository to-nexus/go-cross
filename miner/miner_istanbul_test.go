package miner

import (
	"errors"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	istanbulBackend "github.com/ethereum/go-ethereum/consensus/istanbul/backend"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/bls"
	"github.com/ethereum/go-ethereum/eth/downloader"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/eth/protocols/eth"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/require"
)

// ##CROSS: istanbul

func TestMiner_Update(t *testing.T) {
	for _, mode := range []ethconfig.SyncMode{ethconfig.FullSync, ethconfig.SnapSync} {
		t.Run(fmt.Sprintf("keeps consensus running across failed %v syncs", mode), func(t *testing.T) {
			miner, mux := newIstanbulMiner(t)
			startMiner(t, miner)
			chain := miner.worker.chain
			dl := downloader.New(chain.TrieDB().Disk(), mux, chain, nil, nil)
			t.Cleanup(dl.Terminate)

			for i := 0; i < 3; i++ {
				peer := &withheldHeadPeer{requested: make(chan struct{}, 1)}
				id := fmt.Sprintf("attacker-%d", i)
				require.NoError(t, dl.RegisterPeer(id, eth.ETH68, peer))

				head, td := peer.Head()
				done := make(chan error, 1)
				go func() { done <- dl.LegacySync(id, head, td, nil, mode) }()
				select {
				case <-peer.requested:
				case <-time.After(time.Second):
					t.Fatal("missing head request")
				}
				// The downloader posts StartEvent before it asks for the head.
				syncEvents(mux)
				requireRunning(t, miner)

				dl.Cancel() // End the withheld request without waiting for the network timeout.
				select {
				case err := <-done:
					require.Error(t, err)
				case <-time.After(time.Second):
					t.Fatal("head request was not canceled")
				}
				syncEvents(mux)
				require.Empty(t, miner.worker.startCh, "a failed sync must not restart the running engine")
				require.NoError(t, dl.UnregisterPeer(id))
			}
		})
	}

	t.Run("follows new chain heads during a sync", func(t *testing.T) {
		miner, mux := newIstanbulMiner(t)
		startMiner(t, miner)
		mux.Post(downloader.StartEvent{})
		syncEvents(mux)

		chain := miner.worker.chain
		header := types.CopyHeader(chain.CurrentBlock())
		statedb, err := chain.StateAt(header.Root)
		require.NoError(t, err)

		header.ParentHash = header.Hash()
		header.Number = new(big.Int).Add(header.Number, common.Big1)
		header.Time++

		_, err = chain.WriteBlockAndSetHead(types.NewBlockWithHeader(header), nil, nil, statedb, false, false)
		require.NoError(t, err)
		require.NoError(t, miner.worker.istanbulBackend.NewChainHead())
		require.Eventually(t, func() bool {
			_, view := miner.worker.istanbulBackend.CurrentStat()
			return view.Sequence.Uint64() == header.Number.Uint64()+1
		}, time.Second, time.Millisecond)
		requireRunning(t, miner)
	})

	t.Run("starts right away during a sync", func(t *testing.T) {
		miner, mux := newIstanbulMiner(t)
		mux.Post(downloader.StartEvent{})
		syncEvents(mux)
		startMiner(t, miner)
		requireRunning(t, miner)
	})

	t.Run("keeps a manual stop during a sync", func(t *testing.T) {
		miner, mux := newIstanbulMiner(t)
		startMiner(t, miner)
		mux.Post(downloader.StartEvent{})
		miner.Stop()
		mux.Post(downloader.FailedEvent{Err: errors.New("header timeout")})
		mux.Post(downloader.DoneEvent{Latest: miner.worker.chain.CurrentHeader()})
		miner.Stop() // The loop takes this only after it handled the first stop and the events.
		require.False(t, miner.Mining())
		require.False(t, miner.worker.istanbulBackend.Started())
		require.Empty(t, miner.worker.startCh)
	})
}

func syncEvents(mux *event.TypeMux) {
	mux.Post(downloader.StartEvent{})
}

// startMiner starts the miner and waits until the worker has started the engine.
func startMiner(t *testing.T, miner *Miner) {
	t.Helper()
	miner.Start()
	select {
	case <-miner.worker.startCh:
	case <-time.After(time.Second):
		t.Fatal("miner did not start")
	}
}

func requireRunning(t *testing.T, miner *Miner) {
	t.Helper()
	require.True(t, miner.Mining())
	require.True(t, miner.worker.istanbulBackend.Started())
	require.False(t, miner.worker.syncing.Load())
}

// withheldHeadPeer claims a high TD and never answers header requests.
type withheldHeadPeer struct {
	downloader.Peer
	requested chan struct{}
}

func (p *withheldHeadPeer) Head() (common.Hash, *big.Int) {
	return common.Hash{0xff}, new(big.Int).Lsh(big.NewInt(1), 99)
}

func (p *withheldHeadPeer) RequestHeadersByHash(common.Hash, int, int, bool, chan *eth.Response) (*eth.Request, error) {
	p.requested <- struct{}{}
	return &eth.Request{}, nil
}

// newIstanbulMiner runs the real PoSA backend and Miner event loop without the block-building worker loops.
func newIstanbulMiner(t *testing.T) (*Miner, *event.TypeMux) {
	t.Helper()

	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	blsKey, err := bls.RandKey()
	require.NoError(t, err)

	address := crypto.PubkeyToAddress(key.PublicKey)
	config := *params.CrossTestChainConfig
	config.Istanbul = &params.IstanbulConfig{
		PoSA: &params.PoSAConfig{
			Admin:                address,
			CouncilPeriod:        86400,
			ValidatorEpochLength: 300,
			Validators: []params.PoSAValidator{
				{Operator: address, Validator: address, Signer: blsKey.PublicKey().Marshal(), ID: "validator"},
			},
		},
	}
	extra, err := rlp.EncodeToBytes(&types.IstanbulExtra{
		Validators: []common.Address{address}, Signers: []types.BLSPublicKey{types.BytesToBLSPublicKey(blsKey.PublicKey().Marshal())},
	})
	require.NoError(t, err)

	genesis := &core.Genesis{
		Config:     &config,
		GasLimit:   30_000_000,
		Difficulty: big.NewInt(1),
		Mixhash:    types.IstanbulDigest,
		ExtraData:  extra,
	}
	require.NoError(t, core.MakeBreakpointGenesis(genesis))

	db := rawdb.NewMemoryDatabase()
	engine := istanbulBackend.New(&config, key, blsKey, db, nil)
	chain, err := core.NewBlockChain(db, genesis, engine, nil)
	require.NoError(t, err)
	t.Cleanup(func() { chain.Stop(); db.Close() })

	// Install a PoSA head; block verification is not part of this lifecycle test.
	statedb, err := chain.StateAt(chain.Genesis().Root())
	require.NoError(t, err)

	header := chain.Genesis().Header()
	header.ParentHash, header.Number, header.Time, header.Coinbase = header.Hash(), big.NewInt(1), 1, address
	header.Extra, err = rlp.EncodeToBytes(&types.IstanbulExtra{})
	require.NoError(t, err)

	_, err = chain.WriteBlockAndSetHead(types.NewBlockWithHeader(header), nil, nil, statedb, false, false)
	require.NoError(t, err)
	require.True(t, config.IsIstanbulPoSA(header.Number, header.Time))

	mux := new(event.TypeMux)
	miner := &Miner{
		mux:     mux,
		exitCh:  make(chan struct{}),
		startCh: make(chan struct{}),
		stopCh:  make(chan struct{}),
		worker: &worker{
			chain:           chain,
			chainConfig:     &config,
			istanbulBackend: engine,
			startCh:         make(chan struct{}, 2),
			exitCh:          make(chan struct{}),
		},
	}
	miner.wg.Add(1)
	go miner.update()
	t.Cleanup(func() {
		miner.Close()
		if engine.Started() {
			require.NoError(t, engine.Stop())
		}
		mux.Stop()
	})
	miner.Stop() // Returns once update has set up its downloader subscription.
	return miner, mux
}
