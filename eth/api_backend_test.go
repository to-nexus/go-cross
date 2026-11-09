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

package eth

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math"
	"math/big"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/beacon"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/core/txpool/blobpool"
	"github.com/ethereum/go-ethereum/core/txpool/legacypool"
	"github.com/ethereum/go-ethereum/core/txpool/locals"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	key, _  = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	address = crypto.PubkeyToAddress(key.PublicKey)
	funds   = big.NewInt(1000_000_000_000_000)
	gspec   = &core.Genesis{
		Config: params.MergedTestChainConfig,
		Alloc: types.GenesisAlloc{
			address: {Balance: funds},
		},
		Difficulty: common.Big0,
		BaseFee:    big.NewInt(params.InitialBaseFee),
	}
	signer = types.LatestSignerForChainID(gspec.Config.ChainID)
)

func initBackend(withLocal bool) *EthAPIBackend {
	var (
		// Create a database pre-initialize with a genesis block
		db     = rawdb.NewMemoryDatabase()
		engine = beacon.New(ethash.NewFaker())
	)
	chain, _ := core.NewBlockChain(db, gspec, engine, nil)

	txconfig := legacypool.DefaultConfig
	txconfig.Journal = "" // Don't litter the disk with test journals

	blobPool := blobpool.New(blobpool.Config{Datadir: ""}, chain, nil)
	legacyPool := legacypool.New(txconfig, chain)
	txpool, _ := txpool.New(txconfig.PriceLimit, chain, []txpool.SubPool{legacyPool, blobPool})

	eth := &Ethereum{
		blockchain: chain,
		txPool:     txpool,
	}
	if withLocal {
		eth.localTxTracker = locals.New("", time.Minute, gspec.Config, txpool)
	}
	return &EthAPIBackend{
		eth: eth,
	}
}

func makeTx(nonce uint64, gasPrice *big.Int, amount *big.Int, key *ecdsa.PrivateKey) *types.Transaction {
	if gasPrice == nil {
		gasPrice = big.NewInt(params.GWei)
	}
	if amount == nil {
		amount = big.NewInt(1000)
	}
	tx, _ := types.SignTx(types.NewTransaction(nonce, common.Address{0x00}, amount, params.TxGas, gasPrice, nil), signer, key)
	return tx
}

type unsignedAuth struct {
	nonce uint64
	key   *ecdsa.PrivateKey
}

func pricedSetCodeTx(nonce uint64, gaslimit uint64, gasFee, tip *uint256.Int, key *ecdsa.PrivateKey, unsigned []unsignedAuth) *types.Transaction {
	var authList []types.SetCodeAuthorization
	for _, u := range unsigned {
		auth, _ := types.SignSetCode(u.key, types.SetCodeAuthorization{
			ChainID: *uint256.MustFromBig(gspec.Config.ChainID),
			Address: common.Address{0x42},
			Nonce:   u.nonce,
		})
		authList = append(authList, auth)
	}
	return pricedSetCodeTxWithAuth(nonce, gaslimit, gasFee, tip, key, authList)
}

func pricedSetCodeTxWithAuth(nonce uint64, gaslimit uint64, gasFee, tip *uint256.Int, key *ecdsa.PrivateKey, authList []types.SetCodeAuthorization) *types.Transaction {
	return types.MustSignNewTx(key, signer, &types.SetCodeTx{
		ChainID:    uint256.MustFromBig(gspec.Config.ChainID),
		Nonce:      nonce,
		GasTipCap:  tip,
		GasFeeCap:  gasFee,
		Gas:        gaslimit,
		To:         common.Address{},
		Value:      uint256.NewInt(100),
		Data:       nil,
		AccessList: nil,
		AuthList:   authList,
	})
}

func TestSendTx(t *testing.T) {
	testSendTx(t, false)
	testSendTx(t, true)
}

func TestSendTxEIP2681(t *testing.T) {
	b := initBackend(false)

	// Test EIP-2681: nonce overflow should be rejected
	tx := makeTx(uint64(math.MaxUint64), nil, nil, key) // max uint64 nonce
	err := b.SendTx(context.Background(), tx)
	if err == nil {
		t.Fatal("Expected EIP-2681 nonce overflow error, but transaction was accepted")
	}
	if !errors.Is(err, core.ErrNonceMax) {
		t.Errorf("Expected core.ErrNonceMax, got: %v", err)
	}

	// Test normal case: should succeed
	normalTx := makeTx(0, nil, nil, key)
	err = b.SendTx(context.Background(), normalTx)
	if err != nil {
		t.Errorf("Normal transaction should succeed, got error: %v", err)
	}
}

func testSendTx(t *testing.T, withLocal bool) {
	b := initBackend(withLocal)

	txA := pricedSetCodeTx(0, 250000, uint256.NewInt(params.GWei), uint256.NewInt(params.GWei), key, []unsignedAuth{{nonce: 0, key: key}})
	if err := b.SendTx(context.Background(), txA); err != nil {
		t.Fatalf("Failed to submit tx: %v", err)
	}
	for {
		pending, _ := b.TxPool().ContentFrom(address)
		if len(pending) == 1 {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}

	txB := makeTx(1, nil, nil, key)
	err := b.SendTx(context.Background(), txB)

	if withLocal {
		if err != nil {
			t.Fatalf("Unexpected error sending tx: %v", err)
		}
	} else {
		if !errors.Is(err, txpool.ErrInflightTxLimitReached) {
			t.Fatalf("Unexpected error, want: %v, got: %v", txpool.ErrInflightTxLimitReached, err)
		}
	}
}

// ##CROSS: journal readiness
func TestEthAPIBackend_SendTxJournal(t *testing.T) {
	b := initBackend(false)
	defer b.eth.blockchain.Stop()
	defer b.eth.txPool.Close()
	path := filepath.Join(t.TempDir(), "transactions.rlp")
	tracker := locals.New(path, time.Hour, gspec.Config, b.eth.txPool)
	b.eth.localTxTracker = tracker
	tx := makeTx(0, nil, nil, key)
	require.ErrorIs(t, b.SendTx(context.Background(), tx), locals.ErrNotReady)
	require.False(t, b.eth.txPool.Has(tx.Hash()))
	require.NoError(t, tracker.Start())
	defer tracker.Stop()
	require.NoError(t, b.SendTx(context.Background(), tx))
	require.True(t, b.eth.txPool.Has(tx.Hash()))
	data, err := os.ReadFile(path)
	require.NoError(t, err)
	var restored types.Transaction
	require.NoError(t, rlp.DecodeBytes(data, &restored))
	require.Equal(t, tx.Hash(), restored.Hash())
}

// ##

// ##CROSS: fix upstream 76e3dc6b
func TestEthAPIBackend_SetHead(t *testing.T) {
	// newBackend returns a path mode backend whose states below the head are not present and not recoverable,
	// because the memory database has no state freezer. The handler is left nil, so a rejected call must return
	// before it cancels the downloader.
	newBackend := func(t *testing.T) *EthAPIBackend {
		db := rawdb.NewMemoryDatabase()
		engine := beacon.New(ethash.NewFaker())
		chain, err := core.NewBlockChain(db, gspec, engine, core.DefaultConfig().WithStateScheme(rawdb.PathScheme))
		require.NoError(t, err)
		t.Cleanup(chain.Stop)

		_, blocks, _ := core.GenerateChainWithGenesis(gspec, engine, 10, func(i int, b *core.BlockGen) {
			b.AddTx(makeTx(uint64(i), nil, nil, key)) // change the state root in each block
		})
		_, err = chain.InsertChain(blocks)
		require.NoError(t, err)
		require.NoError(t, chain.TrieDB().Commit(blocks[len(blocks)-1].Root(), false))
		return &EthAPIBackend{eth: &Ethereum{blockchain: chain, chainDb: db}}
	}

	t.Run("rejects a target before the snap sync pivot", func(t *testing.T) {
		b := newBackend(t)
		rawdb.WriteLastPivotNumber(b.eth.chainDb, 8)

		require.ErrorContains(t, b.SetHead(7), "before the snap-sync pivot")
		assert.Equal(t, uint64(10), b.CurrentBlock().Number.Uint64())
	})
	t.Run("rejects a target whose state can't be recovered", func(t *testing.T) {
		b := newBackend(t)
		require.False(t, b.eth.blockchain.HasState(b.eth.blockchain.GetHeaderByNumber(5).Root))

		require.ErrorContains(t, b.SetHead(5), "not recoverable")
		assert.Equal(t, uint64(10), b.CurrentBlock().Number.Uint64())
	})
}

// ##
