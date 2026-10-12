// Copyright 2025 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package locals

import (
	"errors"
	"fmt"
	"io"
	"maps"
	"math/big"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/core/txpool/legacypool"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"
)

var (
	key, _  = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	address = crypto.PubkeyToAddress(key.PublicKey)
	funds   = big.NewInt(1000000000000000)
	gspec   = &core.Genesis{
		Config: params.TestChainConfig,
		Alloc: types.GenesisAlloc{
			address: {Balance: funds},
		},
		BaseFee: big.NewInt(params.InitialBaseFee),
	}
	signer = types.LatestSigner(gspec.Config)
)

type testEnv struct {
	chain   *core.BlockChain
	pool    *txpool.TxPool
	tracker *TxTracker
	genDb   ethdb.Database
}

func newTestEnv(t *testing.T, n int, gasTip uint64, journal string) *testEnv {
	genDb, blocks, _ := core.GenerateChainWithGenesis(gspec, ethash.NewFaker(), n, func(i int, gen *core.BlockGen) {
		tx, err := types.SignTx(types.NewTransaction(gen.TxNonce(address), common.Address{0x00}, big.NewInt(1000), params.TxGas, gen.BaseFee(), nil), signer, key)
		if err != nil {
			panic(err)
		}
		gen.AddTx(tx)
	})

	db := rawdb.NewMemoryDatabase()
	chain, _ := core.NewBlockChain(db, gspec, ethash.NewFaker(), nil)

	legacyPool := legacypool.New(legacypool.DefaultConfig, chain)
	pool, err := txpool.New(gasTip, chain, []txpool.SubPool{legacyPool})
	if err != nil {
		t.Fatalf("Failed to create tx pool: %v", err)
	}
	if n, err := chain.InsertChain(blocks); err != nil {
		t.Fatalf("Failed to process block %d: %v", n, err)
	}
	if err := pool.Sync(); err != nil {
		t.Fatalf("Failed to sync the txpool, %v", err)
	}
	return &testEnv{
		chain:   chain,
		pool:    pool,
		tracker: New(journal, time.Minute, gspec.Config, pool),
		genDb:   genDb,
	}
}

func (env *testEnv) close() {
	env.chain.Stop()
}

// nolint:unused
func (env *testEnv) setGasTip(gasTip uint64) {
	env.pool.SetGasTip(new(big.Int).SetUint64(gasTip))
}

// nolint:unused
func (env *testEnv) makeTx(nonce uint64, gasPrice *big.Int) *types.Transaction {
	if nonce == 0 {
		head := env.chain.CurrentHeader()
		state, _ := env.chain.StateAt(head.Root)
		nonce = state.GetNonce(address)
	}
	if gasPrice == nil {
		gasPrice = big.NewInt(params.GWei)
	}
	tx, _ := types.SignTx(types.NewTransaction(nonce, common.Address{0x00}, big.NewInt(1000), params.TxGas, gasPrice, nil), signer, key)
	return tx
}

func (env *testEnv) makeTxs(n int) []*types.Transaction {
	head := env.chain.CurrentHeader()
	state, _ := env.chain.StateAt(head.Root)
	nonce := state.GetNonce(address)

	var txs []*types.Transaction
	for i := 0; i < n; i++ {
		tx, _ := types.SignTx(types.NewTransaction(nonce+uint64(i), common.Address{0x00}, big.NewInt(1000), params.TxGas, big.NewInt(params.GWei), nil), signer, key)
		txs = append(txs, tx)
	}
	return txs
}

// nolint:unused
func (env *testEnv) commit() {
	head := env.chain.CurrentBlock()
	block := env.chain.GetBlock(head.Hash(), head.Number.Uint64())
	blocks, _ := core.GenerateChain(env.chain.Config(), block, ethash.NewFaker(), env.genDb, 1, func(i int, gen *core.BlockGen) {
		tx, err := types.SignTx(types.NewTransaction(gen.TxNonce(address), common.Address{0x00}, big.NewInt(1000), params.TxGas, gen.BaseFee(), nil), signer, key)
		if err != nil {
			panic(err)
		}
		gen.AddTx(tx)
	})
	env.chain.InsertChain(blocks)
	if err := env.pool.Sync(); err != nil {
		panic(err)
	}
}

func TestResubmit(t *testing.T) {
	env := newTestEnv(t, 10, 0, "")
	defer env.close()

	txs := env.makeTxs(10)
	txsA := txs[:len(txs)/2]
	txsB := txs[len(txs)/2:]
	env.pool.Add(txsA, true)

	pending, queued := env.pool.ContentFrom(address)
	if len(pending) != len(txsA) || len(queued) != 0 {
		t.Fatalf("Unexpected txpool content: %d, %d", len(pending), len(queued))
	}
	env.tracker.TrackAll(txs)

	resubmit := env.tracker.recheck(true)
	if len(resubmit) != len(txsB) {
		t.Fatalf("Unexpected transactions to resubmit, got: %d, want: %d", len(resubmit), len(txsB))
	}
	env.tracker.mu.Lock()
	allCopy := maps.Clone(env.tracker.all)
	env.tracker.mu.Unlock()

	if len(allCopy) != len(txs) {
		t.Fatalf("Unexpected transactions being tracked, got: %d, want: %d", len(allCopy), len(txs))
	}
}

func TestJournal(t *testing.T) {
	journalPath := filepath.Join(t.TempDir(), fmt.Sprintf("%d", rand.Int63()))
	env := newTestEnv(t, 10, 0, journalPath)
	defer env.close()

	env.tracker.Start()
	defer env.tracker.Stop()

	txs := env.makeTxs(10)
	txsA := txs[:len(txs)/2]
	txsB := txs[len(txs)/2:]
	env.pool.Add(txsA, true)

	pending, queued := env.pool.ContentFrom(address)
	if len(pending) != len(txsA) || len(queued) != 0 {
		t.Fatalf("Unexpected txpool content: %d, %d", len(pending), len(queued))
	}
	env.tracker.TrackAll(txsA)
	env.tracker.TrackAll(txsB)
	env.tracker.recheck(true) // manually rejournal the tracker

	// Make sure all the transactions are properly journalled
	// ##CROSS: journal readiness
	trackerB := New("", time.Minute, gspec.Config, env.pool)
	newTxJournal(journalPath).load(func(transactions []*types.Transaction) []error {
		trackerB.TrackAll(transactions)
		return nil
	})

	// ##

	trackerB.mu.Lock()
	allCopy := maps.Clone(trackerB.all)
	trackerB.mu.Unlock()

	if len(allCopy) != len(txs) {
		t.Fatalf("Unexpected transactions being tracked, got: %d, want: %d", len(allCopy), len(txs))
	}
}

// ##CROSS: journal readiness
func TestTxTracker_Start(t *testing.T) {
	t.Run("replay and concurrent tracking", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "transactions.rlp")
		env := newTestEnv(t, 0, 0, "")
		defer env.close()
		defer env.pool.Close()
		config := *gspec.Config
		config.AdventureTime = new(uint64)
		tracker := New(path, time.Hour, &config, env.pool)

		txs := env.makeTxs(1058)
		seed := newTxJournal(path)
		require.NoError(t, seed.setupWriter())
		for _, tx := range txs[:1026] {
			require.NoError(t, seed.insert(tx))
		}
		require.NoError(t, seed.close())
		require.False(t, tracker.Ready())
		require.ErrorIs(t, tracker.Track(txs[1026]), ErrNotReady)
		require.Empty(t, tracker.all)

		// Include a fee-delegated transaction in the live submissions.
		inner := types.DynamicFeeTx{
			ChainID: config.ChainID, Nonce: uint64(len(txs)), To: &address,
			Gas: params.TxGas, GasTipCap: big.NewInt(params.GWei), GasFeeCap: big.NewInt(params.GWei),
			Value: new(big.Int),
		}
		senderTx, err := types.SignTx(types.NewTx(&inner), signer, key)
		require.NoError(t, err)
		inner.V, inner.R, inner.S = senderTx.RawSignatureValues()
		payerKey, err := crypto.GenerateKey()
		require.NoError(t, err)
		payer := crypto.PubkeyToAddress(payerKey.PublicKey)
		feeTx, err := types.SignTx(types.NewTx(types.NewFeeDelegatedDynamicFeeTx(&payer, inner)), types.NewFeeDelegationSigner(config.ChainID), payerKey)
		require.NoError(t, err)
		txs = append(txs, feeTx)

		live := txs[1026:]
		start := make(chan struct{})
		started := make(chan error, 1)
		go func() { <-start; started <- tracker.Start() }()
		errs := make([]error, len(live))
		var wg sync.WaitGroup
		for i, tx := range live {
			wg.Add(1)
			go func() {
				defer wg.Done()
				<-start
				errs[i] = tracker.Track(tx)
			}()
		}
		close(start)
		startErr := <-started
		wg.Wait()
		require.NoError(t, startErr)
		require.True(t, tracker.Ready())
		for i, err := range errs {
			if errors.Is(err, ErrNotReady) {
				err = tracker.Track(live[i])
			}
			require.NoError(t, err)
		}
		require.NoError(t, tracker.Stop())
		require.False(t, tracker.Ready())
		require.ErrorIs(t, tracker.Track(txs[0]), ErrNotReady)

		// Restart before any rotation. Both replayed and live transactions must survive.
		restored := New(path, time.Hour, &config, env.pool)
		require.NoError(t, restored.Start())
		defer restored.Stop()
		restored.mu.Lock()
		defer restored.mu.Unlock()
		require.Len(t, restored.all, len(txs))
		for _, tx := range txs {
			got := restored.all[tx.Hash()]
			require.NotNil(t, got)
			wantRaw, err := tx.MarshalBinary()
			require.NoError(t, err)
			gotRaw, err := got.MarshalBinary()
			require.NoError(t, err)
			require.Equal(t, wantRaw, gotRaw)
		}
	})
	t.Run("writer setup error", func(t *testing.T) {
		tracker := New(filepath.Join(t.TempDir(), "missing", "transactions.rlp"), time.Hour, gspec.Config, nil)
		require.ErrorIs(t, tracker.Start(), os.ErrNotExist)
		require.False(t, tracker.Ready())
		require.Nil(t, tracker.journal.writer)
	})
	t.Run("replay error", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "transactions.rlp")
		tx := types.MustSignNewTx(key, signer, &types.LegacyTx{GasPrice: big.NewInt(params.GWei)})
		seed := newTxJournal(path)
		require.NoError(t, seed.setupWriter())
		require.NoError(t, seed.insert(tx))
		_, err := seed.writer.Write([]byte{0xff})
		require.NoError(t, err)
		require.NoError(t, seed.close())
		before, err := os.ReadFile(path)
		require.NoError(t, err)
		tracker := New(path, time.Hour, gspec.Config, nil)
		require.Error(t, tracker.Start())
		require.False(t, tracker.Ready())
		require.Nil(t, tracker.journal.writer)
		require.Contains(t, tracker.all, tx.Hash())
		after, err := os.ReadFile(path)
		require.NoError(t, err)
		require.Equal(t, before, after)
	})
}

type failingJournalWriter struct {
	io.WriteCloser
	err error
}

// Write simulates a journal write failure.
func (w *failingJournalWriter) Write([]byte) (int, error) { return 0, w.err }

// Close releases the writer and simulates a journal close failure.
func (w *failingJournalWriter) Close() error {
	w.WriteCloser.Close()
	return w.err
}

func TestTxTracker_JournalErrors(t *testing.T) {
	tracker := New(filepath.Join(t.TempDir(), "transactions.rlp"), time.Hour, gspec.Config, nil)
	require.NoError(t, tracker.Start())
	failure := errors.New("journal I/O failure")
	tracker.mu.Lock()
	tracker.journal.writer = &failingJournalWriter{tracker.journal.writer, failure}
	tracker.mu.Unlock()
	tx := types.MustSignNewTx(key, signer, &types.LegacyTx{GasPrice: big.NewInt(params.GWei)})
	require.NoError(t, tracker.Track(tx))
	require.ErrorIs(t, tracker.Stop(), failure)
	require.Contains(t, tracker.all, tx.Hash())
	require.False(t, tracker.Ready())
	require.Nil(t, tracker.journal.writer)
}

// ##
