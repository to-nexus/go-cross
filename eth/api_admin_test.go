// Copyright 2026 The go-ethereum Authors
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
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/consensus/beacon"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/contracts"
	"github.com/ethereum/go-ethereum/contracts/latest"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ##CROSS: istanbul posa

func TestAdminAPI_SuspendIstanbulValidator(t *testing.T) {
	// newService starts a node with nodeKey on a Breakpoint genesis whose first validators are validators, each funded.
	newService := func(t *testing.T, nodeKey *ecdsa.PrivateKey, validators ...common.Address) (*Ethereum, func(txs ...*types.Transaction)) {
		t.Helper()
		config := *params.CrossDev3ChainConfig
		istanbul := *config.Istanbul
		posa := *istanbul.PoSA
		config.Istanbul, istanbul.PoSA, config.BreakpointTime = &istanbul, &posa, nil
		posa.Validators = append([]params.PoSAValidator(nil), posa.Validators[:2]...)
		alloc := types.GenesisAlloc{crypto.PubkeyToAddress(nodeKey.PublicKey): {Balance: big.NewInt(params.Ether)}}
		for i, validator := range validators {
			posa.Validators[i].Validator = validator
			alloc[validator] = types.Account{Balance: big.NewInt(params.Ether)}
		}
		genesis := &core.Genesis{
			Config:     &config,
			GasLimit:   30_000_000,
			Difficulty: big.NewInt(1),
			Alloc:      alloc,
		}
		require.NoError(t, core.MakeGenesis(genesis, core.GenesisContracts{Breakpoint: true}))
		config.Istanbul, config.TerminalTotalDifficulty, config.AdventureTime = nil, common.Big0, nil

		stack, err := node.New(&node.Config{P2P: p2p.Config{PrivateKey: nodeKey}})
		require.NoError(t, err)
		t.Cleanup(func() { stack.Close() })
		service, err := New(stack, &ethconfig.Config{Genesis: genesis, RPCGasCap: 50_000_000})
		require.NoError(t, err)
		require.NoError(t, stack.Start())

		var blockTxs [][]*types.Transaction
		insert := func(txs ...*types.Transaction) {
			blockTxs = append(blockTxs, txs)
			// Make the whole chain again; the blocks inserted before come out the same and are skipped.
			_, blocks, _ := core.GenerateChainWithGenesis(genesis, beacon.New(ethash.NewFaker()), len(blockTxs), func(i int, b *core.BlockGen) {
				b.SetParentBeaconRoot(common.Hash{}) // the import runs the EIP-4788 call, so the maker must run it too
				for _, tx := range blockTxs[i] {
					b.AddTx(tx)
				}
			})
			_, err := service.BlockChain().InsertChain(blocks)
			require.NoError(t, err)
		}
		return service, insert
	}
	// pendingTx waits until the pool has a pending tx of addr with nonce.
	pendingTx := func(t *testing.T, service *Ethereum, addr common.Address, nonce uint64) *types.Transaction {
		t.Helper()
		var found *types.Transaction
		require.Eventually(t, func() bool {
			pending, _ := service.TxPool().ContentFrom(addr)
			for _, tx := range pending {
				if tx.Nonce() == nonce {
					found = tx
				}
			}
			return found != nil
		}, 5*time.Second, 10*time.Millisecond)
		return found
	}
	type result struct {
		hash common.Hash
		err  error
	}
	endTime := uint64(time.Now().Add(time.Hour).Unix())
	// suspend calls SuspendIstanbulValidator in the background, as it waits for the tx to be included.
	suspend := func(service *Ethereum) <-chan result {
		done := make(chan result, 1)
		go func() {
			hash, err := NewAdminAPI(service).SuspendIstanbulValidator(context.Background(), endTime)
			done <- result{hash, err}
		}()
		return done
	}

	t.Run("sends a dynamic fee tx signed by the validator and waits for the block", func(t *testing.T) {
		key, _ := crypto.GenerateKey()
		addr := crypto.PubkeyToAddress(key.PublicKey)
		service, insert := newService(t, key, addr)

		done := suspend(service)
		tx := pendingTx(t, service, addr, 0)
		assert.Equal(t, uint8(types.DynamicFeeTxType), tx.Type())
		assert.Equal(t, contracts.StakeHubAddr, *tx.To())
		assert.Equal(t, latest.NewStakeHub().PackSuspendByValidator(new(big.Int).SetUint64(endTime)), tx.Data())
		sender, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
		require.NoError(t, err)
		assert.Equal(t, addr, sender)
		select {
		case <-done:
			t.Fatal("returned before the tx is included")
		default:
		}

		insert(tx)
		res := <-done
		require.NoError(t, res.err)
		assert.Equal(t, tx.Hash(), res.hash)
	})

	t.Run("sends again when another tx takes the nonce", func(t *testing.T) {
		key, _ := crypto.GenerateKey()
		addr := crypto.PubkeyToAddress(key.PublicKey)
		service, insert := newService(t, key, addr)

		done := suspend(service)
		dropped := pendingTx(t, service, addr, 0)
		// On the real chain, a system tx of the proposer takes the nonce; a plain transfer does the same here.
		insert(types.MustSignNewTx(key, types.LatestSignerForChainID(dropped.ChainId()), &types.DynamicFeeTx{
			ChainID:   dropped.ChainId(),
			Nonce:     dropped.Nonce(),
			GasTipCap: dropped.GasTipCap(),
			GasFeeCap: dropped.GasFeeCap(),
			Gas:       params.TxGas,
			To:        &addr,
		}))
		resent := pendingTx(t, service, addr, 1)
		assert.Equal(t, dropped.Data(), resent.Data())

		insert(resent)
		res := <-done
		require.NoError(t, res.err)
		assert.Equal(t, resent.Hash(), res.hash)
	})

	t.Run("takes the tx as dropped when another tx takes the nonce before the pool adds it", func(t *testing.T) {
		key, _ := crypto.GenerateKey()
		addr := crypto.PubkeyToAddress(key.PublicKey)
		service, insert := newService(t, key, addr)
		api := NewAdminAPI(service)

		// Nonce 0 is read for the suspend tx, then another tx takes it and the pool is updated before the suspend tx is added.
		insert(types.MustSignNewTx(key, types.LatestSignerForChainID(service.BlockChain().Config().ChainID), &types.DynamicFeeTx{
			ChainID:   service.BlockChain().Config().ChainID,
			GasTipCap: big.NewInt(params.GWei),
			GasFeeCap: big.NewInt(10 * params.GWei),
			Gas:       params.TxGas,
			To:        &addr,
		}))
		require.Eventually(t, func() bool { return service.TxPool().Nonce(addr) == 1 }, 5*time.Second, 10*time.Millisecond)

		nonce, to := hexutil.Uint64(0), contracts.StakeHubAddr
		input := hexutil.Bytes(latest.NewStakeHub().PackSuspendByValidator(new(big.Int).SetUint64(endTime)))
		tx, err := api.sendSuspendTx(context.Background(), key, ethapi.TransactionArgs{From: &addr, To: &to, Input: &input, Nonce: &nonce})
		require.NoError(t, err)

		// It returns without a new head, so the retry loop sends the tx again right away.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		included, err := api.waitSuspendTx(ctx, tx, addr, nil)
		require.NoError(t, err)
		assert.False(t, included)
	})

	t.Run("fails when the included tx reverts", func(t *testing.T) {
		key, _ := crypto.GenerateKey()
		otherKey, _ := crypto.GenerateKey()
		addr, other := crypto.PubkeyToAddress(key.PublicKey), crypto.PubkeyToAddress(otherKey.PublicKey)
		service, insert := newService(t, key, addr, other)

		done := suspend(service)
		tx := pendingTx(t, service, addr, 0)
		// The other validator suspends first in the same block, so the tx of the node reverts with NotEnoughValidators.
		insert(types.MustSignNewTx(otherKey, types.LatestSignerForChainID(tx.ChainId()), &types.DynamicFeeTx{
			ChainID:   tx.ChainId(),
			GasTipCap: tx.GasTipCap(),
			GasFeeCap: tx.GasFeeCap(),
			Gas:       1_000_000,
			To:        tx.To(),
			Data:      tx.Data(),
		}), tx)
		res := <-done
		require.EqualError(t, res.err, "suspend tx "+tx.Hash().Hex()+" failed: execution reverted: NotEnoughValidators()")
		receipts := service.BlockChain().GetReceiptsByHash(service.BlockChain().CurrentBlock().Hash())
		assert.Equal(t, types.ReceiptStatusSuccessful, receipts[0].Status)
	})

	t.Run("fails without sending when the node is not a validator", func(t *testing.T) {
		key, _ := crypto.GenerateKey()
		service, _ := newService(t, key, common.HexToAddress("0x1234"))

		// Call through the RPC server to check the error code and data in the response.
		server := rpc.NewServer()
		require.NoError(t, server.RegisterName("admin", NewAdminAPI(service)))

		err := rpc.DialInProc(server).Call(new(common.Hash), "admin_suspendIstanbulValidator", endTime)
		require.ErrorContains(t, err, "execution reverted: ValidatorNotExists()")

		var rpcErr rpc.Error
		require.ErrorAs(t, err, &rpcErr)
		assert.Equal(t, 3, rpcErr.ErrorCode())

		var dataErr rpc.DataError
		require.ErrorAs(t, err, &dataErr)
		assert.Equal(t, hexutil.Encode(crypto.Keccak256([]byte("ValidatorNotExists()"))[:4]), dataErr.ErrorData())
		pending, queued := service.TxPool().Stats()
		assert.Zero(t, pending+queued)
	})
}

// ##
