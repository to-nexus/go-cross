// Copyright 2024 The go-ethereum Authors
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

package simulated

// ##CROSS: fee delegation
//
// End-to-end scenario on a real simulated backend: a fee payer is funded with a
// small amount of ether at genesis while two senders flood the pool with
// fee-delegated transactions sponsored by that payer. Each transaction is
// individually affordable, but once the payer's cumulative obligation across the
// pooled transactions exceeds its balance, further transactions are rejected by
// the txpool's ValidateTransactionWithState with an "insufficient feepayer's
// funds" error.

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

// signFeeDelegatedTx builds a signed fee-delegated dynamic-fee transaction. The
// inner (sender) transaction is signed by senderKey; the fee-delegated wrapper,
// which names a separate fee payer, is signed by payerKey. Because the fee payer
// differs from the sender, the payer covers only the gas (gasFeeCap * gas).
func signFeeDelegatedTx(chainID *big.Int, nonce uint64, to common.Address, value, gasFeeCap, gasTipCap *big.Int, gas uint64, senderKey, payerKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	londonSigner := types.NewLondonSigner(chainID)
	feeSigner := types.NewFeeDelegationSigner(chainID)
	payer := crypto.PubkeyToAddress(payerKey.PublicKey)

	// Sign the inner dynamic-fee tx with the sender key to obtain the sender
	// signature values.
	senderTx, err := types.SignTx(types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Gas:       gas,
		To:        &to,
		Value:     value,
	}), londonSigner, senderKey)
	if err != nil {
		return nil, err
	}
	v, r, s := senderTx.RawSignatureValues()

	// Re-create the inner data with the sender signature attached, wrap it in a
	// fee-delegated transaction and sign that with the fee payer's key.
	inner := types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Gas:       gas,
		To:        &to,
		Value:     value,
		V:         v,
		R:         r,
		S:         s,
	}
	return types.SignTx(types.NewTx(types.NewFeeDelegatedDynamicFeeTx(&payer, inner)), feeSigner, payerKey)
}

func TestFeeDelegationFeePayerExhausted(t *testing.T) {
	// Fee payer and two senders.
	payerKey, _ := crypto.GenerateKey()
	payer := crypto.PubkeyToAddress(payerKey.PublicKey)
	sender1Key, _ := crypto.GenerateKey()
	sender1 := crypto.PubkeyToAddress(sender1Key.PublicKey)
	sender2Key, _ := crypto.GenerateKey()
	sender2 := crypto.PubkeyToAddress(sender2Key.PublicKey)

	// Gas parameters chosen comfortably above the dev chain's 1 gwei base fee.
	const gas = uint64(21000)
	gasTipCap := big.NewInt(params.GWei)     // 1 gwei
	gasFeeCap := big.NewInt(2 * params.GWei) // 2 gwei

	// Each fee-delegated tx commits the payer to gasFeeCap*gas. Fund the payer for
	// exactly `affordable` such transactions so the next one overshoots its balance.
	const affordable = 3
	feePerTx := new(big.Int).Mul(gasFeeCap, new(big.Int).SetUint64(gas))
	payerBalance := new(big.Int).Mul(feePerTx, big.NewInt(affordable))

	// Genesis allocation: small balance for the fee payer, generous for senders so
	// only the fee payer's balance limits how many txs the pool accepts.
	sim := NewBackend(types.GenesisAlloc{
		payer:   {Balance: payerBalance},
		sender1: {Balance: big.NewInt(1e18)},
		sender2: {Balance: big.NewInt(1e18)},
	})
	defer sim.Close()

	client := sim.Client()
	ctx := context.Background()
	chainID, err := client.ChainID(ctx)
	if err != nil {
		t.Fatalf("chain id: %v", err)
	}

	to := common.HexToAddress("0x00000000000000000000000000000000000000aa")
	value := big.NewInt(1)

	// waitPending blocks until the account's pending nonce reaches want, i.e. until
	// the just-sent tx has been promoted into the pending pool (and recorded in the
	// shared fee-payer index). SendTransaction only enqueues asynchronously.
	waitPending := func(addr common.Address, want uint64) {
		t.Helper()
		for i := 0; i < 500; i++ {
			if n, err := client.PendingNonceAt(ctx, addr); err == nil && n == want {
				return
			}
			time.Sleep(10 * time.Millisecond)
		}
		t.Fatalf("timeout waiting for pending nonce %d of %s", want, addr.Hex())
	}

	senders := []struct {
		key  *ecdsa.PrivateKey
		addr common.Address
	}{
		{sender1Key, sender1},
		{sender2Key, sender2},
	}

	// Submit the two senders' transactions interleaved, one at a time, waiting for
	// each accepted tx to be promoted before sending the next.
	const perSender = 5
	accepted := 0
	for nonce := uint64(0); nonce < perSender; nonce++ {
		for _, s := range senders {
			tx, err := signFeeDelegatedTx(chainID, nonce, to, value, gasFeeCap, gasTipCap, gas, s.key, payerKey)
			if err != nil {
				t.Fatalf("sign fee-delegated tx: %v", err)
			}
			err = client.SendTransaction(ctx, tx)

			if accepted < affordable {
				if err != nil {
					t.Fatalf("tx #%d rejected before fee payer exhausted: %v", accepted+1, err)
				}
				accepted++
				waitPending(s.addr, nonce+1)
				continue
			}

			// Beyond the fee payer's means: the pool must reject it.
			if err == nil {
				t.Fatalf("tx #%d accepted, but fee payer should be exhausted", accepted+1)
			}
			if !strings.Contains(err.Error(), "insufficient feepayer's funds") {
				t.Fatalf("tx after exhaustion: unexpected error %q", err)
			}
		}
	}

	if accepted != affordable {
		t.Fatalf("accepted %d fee-delegated txs, want %d", accepted, affordable)
	}

	// The accepted transactions are real and mineable: sealing a block drains the
	// fee payer's on-chain balance.
	sim.Commit()
	bal, err := client.BalanceAt(ctx, payer, nil)
	if err != nil {
		t.Fatalf("balance: %v", err)
	}
	if bal.Cmp(payerBalance) >= 0 {
		t.Fatalf("expected fee payer balance to drop after mining: before %v, after %v", payerBalance, bal)
	}
}
