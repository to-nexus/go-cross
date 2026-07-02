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

package legacypool

// ##CROSS: fee delegation
//
// These tests cover the pool-level fee-payer cost index that prevents a single
// fee payer from being committed to more than its balance across multiple pooled
// fee-delegated transactions. A single tx alone may be payable while the running
// sum of every tx the payer sponsors is not.

import (
	"crypto/ecdsa"
	"errors"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/holiman/uint256"
)

// feeDelegatedDynamicFeeTx builds a signed fee-delegated dynamic-fee transaction.
// The inner dynamic-fee transaction is signed by senderKey, and the resulting
// wrapper (which carries the separate fee payer) is signed by payerKey. Because
// the fee payer differs from the sender, the fee payer covers the gas cost
// (gasFee * gaslimit) while the sender only covers the transferred value.
func feeDelegatedDynamicFeeTx(nonce uint64, gaslimit uint64, gasFee, tip, value *big.Int, senderKey, payerKey *ecdsa.PrivateKey) *types.Transaction {
	chainID := params.TestChainConfig.ChainID
	londonSigner := types.NewLondonSigner(chainID)
	feeSigner := types.NewFeeDelegationSigner(chainID)
	payer := crypto.PubkeyToAddress(payerKey.PublicKey)

	// Sign the inner (sender) dynamic-fee transaction first to obtain the sender
	// signature values.
	senderTx, err := types.SignTx(types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: tip,
		GasFeeCap: gasFee,
		Gas:       gaslimit,
		To:        &common.Address{},
		Value:     value,
	}), londonSigner, senderKey)
	if err != nil {
		panic(err)
	}
	sv, sr, ss := senderTx.RawSignatureValues()

	// Re-create the inner data with the sender signature attached, wrap it in a
	// fee-delegated transaction and sign that with the fee payer's key.
	inner := types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: tip,
		GasFeeCap: gasFee,
		Gas:       gaslimit,
		To:        &common.Address{},
		Value:     value,
		V:         sv,
		R:         sr,
		S:         ss,
	}
	feeTx, err := types.SignTx(types.NewTx(types.NewFeeDelegatedDynamicFeeTx(&payer, inner)), feeSigner, payerKey)
	if err != nil {
		panic(err)
	}
	return feeTx
}

// setupAdventurePool returns a pool whose chain config enables the Adventure fork.
// Fee-delegated (type 0x07) transactions are only accepted once that fork is
// active (see ValidateTransaction), so the plain TestChainConfig would reject them.
func setupAdventurePool() (*LegacyPool, *ecdsa.PrivateKey) {
	cfg := *params.TestChainConfig
	adventure := uint64(0)
	cfg.AdventureTime = &adventure
	return setupPoolWithConfig(&cfg)
}

// payerCost returns the cumulative fee-payer cost currently tracked by the pool
// for the given payer (zero if the payer has no pooled obligations).
func payerCost(pool *LegacyPool, payer common.Address) *big.Int {
	pool.mu.RLock()
	defer pool.mu.RUnlock()
	if v := pool.feePayerCost[payer]; v != nil {
		return v.ToBig()
	}
	return new(big.Int)
}

// nilFeePayerTx builds a malformed fee-delegated transaction whose FeePayer is
// nil (the field is RLP-nilable) while still carrying fee-payer signature values.
// This mimics a transaction a malicious peer could craft and relay over the P2P
// path, which bypasses the RPC-side nil guard.
func nilFeePayerTx(nonce uint64, gaslimit uint64, gasFee, tip, value *big.Int, senderKey, payerKey *ecdsa.PrivateKey) *types.Transaction {
	chainID := params.TestChainConfig.ChainID
	londonSigner := types.NewLondonSigner(chainID)
	feeSigner := types.NewFeeDelegationSigner(chainID)

	senderTx, err := types.SignTx(types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: tip,
		GasFeeCap: gasFee,
		Gas:       gaslimit,
		To:        &common.Address{},
		Value:     value,
	}), londonSigner, senderKey)
	if err != nil {
		panic(err)
	}
	sv, sr, ss := senderTx.RawSignatureValues()
	inner := types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: tip,
		GasFeeCap: gasFee,
		Gas:       gaslimit,
		To:        &common.Address{},
		Value:     value,
		V:         sv,
		R:         sr,
		S:         ss,
	}
	// FeePayer is deliberately left nil.
	feeTx, err := types.SignTx(types.NewTx(types.NewFeeDelegatedDynamicFeeTx(nil, inner)), feeSigner, payerKey)
	if err != nil {
		panic(err)
	}
	return feeTx
}

// TestFeeDelegationNilFeePayerRejectedNotPanic verifies that a malformed
// fee-delegated transaction with a nil FeePayer is rejected by txpool admission
// with ErrInvalidFeePayer rather than panicking the node. This is the remote
// (peer) path, which is the one that lacks the RPC-side structural guard.
func TestFeeDelegationNilFeePayerRejectedNotPanic(t *testing.T) {
	t.Parallel()

	pool, _ := setupAdventurePool()
	defer pool.Close()

	senderKey, _ := crypto.GenerateKey()
	sender := crypto.PubkeyToAddress(senderKey.PublicKey)
	payerKey, _ := crypto.GenerateKey()

	testAddBalance(pool, sender, new(big.Int).SetUint64(params.Ether))

	tx := nilFeePayerTx(0, 100000, big.NewInt(2), big.NewInt(1), big.NewInt(100), senderKey, payerKey)
	if tx.FeePayer() != nil {
		t.Fatal("test setup: expected a nil fee payer")
	}

	// Must not panic, and must be rejected as an invalid fee payer.
	err := pool.addRemoteSync(tx)
	if !errors.Is(err, core.ErrInvalidFeePayer) {
		t.Fatalf("nil fee-payer tx: have err %v, want %v", err, core.ErrInvalidFeePayer)
	}
	if pending, queued := pool.Stats(); pending != 0 || queued != 0 {
		t.Fatalf("malformed tx leaked into pool: pending %d, queued %d", pending, queued)
	}
}

// TestFeeDelegationCumulativeOverdraftRejected verifies that a second
// fee-delegated transaction sharing a fee payer is rejected when, although each
// transaction is individually payable, their combined fee-payer cost exceeds the
// fee payer's balance.
func TestFeeDelegationCumulativeOverdraftRejected(t *testing.T) {
	t.Parallel()

	pool, _ := setupAdventurePool()
	defer pool.Close()

	payerKey, _ := crypto.GenerateKey()
	payer := crypto.PubkeyToAddress(payerKey.PublicKey)
	sender1Key, _ := crypto.GenerateKey()
	sender1 := crypto.PubkeyToAddress(sender1Key.PublicKey)
	sender2Key, _ := crypto.GenerateKey()
	sender2 := crypto.PubkeyToAddress(sender2Key.PublicKey)

	// Each tx commits the payer to gasFee*gas = 2*100000 = 200000.
	const perTxFeePayerCost = 200_000

	testAddBalance(pool, sender1, new(big.Int).SetUint64(params.Ether))
	testAddBalance(pool, sender2, new(big.Int).SetUint64(params.Ether))
	// The payer can cover one tx (200000) but not two (400000).
	testAddBalance(pool, payer, big.NewInt(300_000))

	// First tx is accepted and promoted into the pending pool, recording its cost
	// in the shared fee-payer index.
	tx1 := feeDelegatedDynamicFeeTx(0, 100000, big.NewInt(2), big.NewInt(1), big.NewInt(100), sender1Key, payerKey)
	if err := pool.addRemoteSync(tx1); err != nil {
		t.Fatalf("first fee-delegated tx rejected: %v", err)
	}
	if got, want := payerCost(pool, payer), big.NewInt(perTxFeePayerCost); got.Cmp(want) != 0 {
		t.Fatalf("payer index mismatch after tx1: have %v, want %v", got, want)
	}

	// Second tx is individually payable (200000 <= 300000) but pushes the payer's
	// cumulative obligation to 400000 > 300000, so it must be rejected.
	tx2 := feeDelegatedDynamicFeeTx(0, 100000, big.NewInt(2), big.NewInt(1), big.NewInt(100), sender2Key, payerKey)
	if err := pool.addRemoteSync(tx2); !errors.Is(err, core.ErrFeePayerInsufficientFunds) {
		t.Fatalf("second fee-delegated tx: have err %v, want %v", err, core.ErrFeePayerInsufficientFunds)
	}

	// The rejected tx must not have leaked into the index.
	if got, want := payerCost(pool, payer), big.NewInt(perTxFeePayerCost); got.Cmp(want) != 0 {
		t.Fatalf("payer index mismatch after rejected tx2: have %v, want %v", got, want)
	}
	if pending, _ := pool.Stats(); pending != 1 {
		t.Fatalf("pending count mismatch: have %d, want 1", pending)
	}
	if err := validatePoolInternals(pool); err != nil {
		t.Fatalf("pool internal state corrupted: %v", err)
	}
}

// TestFeeDelegationCumulativeWithinBalance verifies that multiple fee-delegated
// transactions sharing a fee payer are all accepted as long as the payer's
// balance covers their combined fee-payer cost.
func TestFeeDelegationCumulativeWithinBalance(t *testing.T) {
	t.Parallel()

	pool, _ := setupAdventurePool()
	defer pool.Close()

	payerKey, _ := crypto.GenerateKey()
	payer := crypto.PubkeyToAddress(payerKey.PublicKey)
	sender1Key, _ := crypto.GenerateKey()
	sender1 := crypto.PubkeyToAddress(sender1Key.PublicKey)
	sender2Key, _ := crypto.GenerateKey()
	sender2 := crypto.PubkeyToAddress(sender2Key.PublicKey)

	testAddBalance(pool, sender1, new(big.Int).SetUint64(params.Ether))
	testAddBalance(pool, sender2, new(big.Int).SetUint64(params.Ether))
	// The payer can comfortably cover both txs (2 * 200000 = 400000).
	testAddBalance(pool, payer, big.NewInt(500_000))

	tx1 := feeDelegatedDynamicFeeTx(0, 100000, big.NewInt(2), big.NewInt(1), big.NewInt(100), sender1Key, payerKey)
	if err := pool.addRemoteSync(tx1); err != nil {
		t.Fatalf("first fee-delegated tx rejected: %v", err)
	}
	tx2 := feeDelegatedDynamicFeeTx(0, 100000, big.NewInt(2), big.NewInt(1), big.NewInt(100), sender2Key, payerKey)
	if err := pool.addRemoteSync(tx2); err != nil {
		t.Fatalf("second fee-delegated tx rejected: %v", err)
	}

	if got, want := payerCost(pool, payer), big.NewInt(400_000); got.Cmp(want) != 0 {
		t.Fatalf("payer index mismatch: have %v, want %v", got, want)
	}
	if pending, _ := pool.Stats(); pending != 2 {
		t.Fatalf("pending count mismatch: have %d, want 2", pending)
	}
	if err := validatePoolInternals(pool); err != nil {
		t.Fatalf("pool internal state corrupted: %v", err)
	}
}

// TestFeeDelegationReplacementNotDoubleCounted verifies that replacing a pooled
// fee-delegated transaction with a price-bumped one does not double count the
// fee payer's obligation. The cost of the transaction being replaced is excluded
// from the cumulative check, so the replacement is accepted even though naively
// summing both would exceed the payer's balance.
func TestFeeDelegationReplacementNotDoubleCounted(t *testing.T) {
	t.Parallel()

	pool, _ := setupAdventurePool()
	defer pool.Close()

	payerKey, _ := crypto.GenerateKey()
	payer := crypto.PubkeyToAddress(payerKey.PublicKey)
	senderKey, _ := crypto.GenerateKey()
	sender := crypto.PubkeyToAddress(senderKey.PublicKey)

	testAddBalance(pool, sender, new(big.Int).SetUint64(params.Ether))
	// Enough for a single 300000 obligation, but not for 200000 + 300000 = 500000
	// if the replaced tx were double counted.
	testAddBalance(pool, payer, big.NewInt(300_000))

	// Original tx: fee-payer cost = 2 * 100000 = 200000.
	tx := feeDelegatedDynamicFeeTx(0, 100000, big.NewInt(2), big.NewInt(1), big.NewInt(100), senderKey, payerKey)
	if err := pool.addRemoteSync(tx); err != nil {
		t.Fatalf("original fee-delegated tx rejected: %v", err)
	}
	if got, want := payerCost(pool, payer), big.NewInt(200_000); got.Cmp(want) != 0 {
		t.Fatalf("payer index mismatch after original: have %v, want %v", got, want)
	}

	// Price-bumped replacement at the same (sender, nonce): fee-payer cost =
	// 3 * 100000 = 300000. The old 200000 must be excluded from the cumulative
	// check (300000 <= 300000), otherwise the replacement would be wrongly rejected.
	repl := feeDelegatedDynamicFeeTx(0, 100000, big.NewInt(3), big.NewInt(2), big.NewInt(100), senderKey, payerKey)
	if err := pool.addRemoteSync(repl); err != nil {
		t.Fatalf("price-bumped replacement rejected: %v", err)
	}

	// Only the replacement remains, so the index reflects its cost alone.
	if got, want := payerCost(pool, payer), big.NewInt(300_000); got.Cmp(want) != 0 {
		t.Fatalf("payer index mismatch after replacement: have %v, want %v", got, want)
	}
	if pending, _ := pool.Stats(); pending != 1 {
		t.Fatalf("pending count mismatch: have %d, want 1", pending)
	}
	if err := validatePoolInternals(pool); err != nil {
		t.Fatalf("pool internal state corrupted: %v", err)
	}
}

// TestFeeDelegationScenarioPayerExhausted is an end-to-end scenario: a fee payer
// holds only a small amount of ether while two senders flood the pool with long
// runs of sequential-nonce fee-delegated transactions, all sponsored by that same
// payer. Each transaction is individually affordable, but as accepted txs pile up
// in the pending pool the payer's cumulative fee obligation grows. Once it would
// exceed the payer's balance, every further transaction is rejected by
// ValidateTransactionWithState with ErrFeePayerInsufficientFunds.
func TestFeeDelegationScenarioPayerExhausted(t *testing.T) {
	t.Parallel()

	pool, _ := setupAdventurePool()
	defer pool.Close()

	payerKey, _ := crypto.GenerateKey()
	payer := crypto.PubkeyToAddress(payerKey.PublicKey)

	// Two senders, each generously funded so the sender-side checks never bind and
	// only the fee payer's balance limits how many txs the pool accepts.
	senderKeys := make([]*ecdsa.PrivateKey, 2)
	for i := range senderKeys {
		senderKeys[i], _ = crypto.GenerateKey()
		testAddBalance(pool, crypto.PubkeyToAddress(senderKeys[i].PublicKey), new(big.Int).SetUint64(params.Ether))
	}

	// Each fee-delegated tx commits the payer to gasFee*gas = 2*100000 = 200000.
	// The payer is funded for exactly 5 such transactions; the 6th overshoots.
	const (
		feePerTx   = 200_000
		affordable = 5
		perSender  = 8 // each sender tries more than its fair share
	)
	testAddBalance(pool, payer, big.NewInt(feePerTx*affordable))

	// Submit the two senders' transactions interleaved and one at a time, so each
	// accepted tx is promoted into the pending pool (and recorded in the shared
	// fee-payer index) before the next one is validated.
	accepted := 0
	for nonce := uint64(0); nonce < perSender; nonce++ {
		for _, key := range senderKeys {
			tx := feeDelegatedDynamicFeeTx(nonce, 100000, big.NewInt(2), big.NewInt(1), big.NewInt(100), key, payerKey)
			err := pool.addRemoteSync(tx)
			if accepted < affordable {
				if err != nil {
					t.Fatalf("tx #%d rejected before payer exhausted: %v", accepted+1, err)
				}
				accepted++
				continue
			}
			// Beyond the payer's means: must fail in ValidateTransactionWithState.
			if !errors.Is(err, core.ErrFeePayerInsufficientFunds) {
				t.Fatalf("tx after payer exhausted: have err %v, want %v", err, core.ErrFeePayerInsufficientFunds)
			}
		}
	}

	if accepted != affordable {
		t.Fatalf("accepted %d fee-delegated txs, want %d", accepted, affordable)
	}
	// The payer ends fully committed, exactly at its balance.
	if got, want := payerCost(pool, payer), big.NewInt(feePerTx*affordable); got.Cmp(want) != 0 {
		t.Fatalf("payer index mismatch: have %v, want %v", got, want)
	}
	if pending, _ := pool.Stats(); pending != affordable {
		t.Fatalf("pending count mismatch: have %d, want %d", pending, affordable)
	}
	if err := validatePoolInternals(pool); err != nil {
		t.Fatalf("pool internal state corrupted: %v", err)
	}
}

// TestListFeePayerCostIndex exercises the list-level add/sub bookkeeping that
// keeps the shared pool fee-payer index in sync as fee-delegated transactions
// enter and leave a pending list.
func TestListFeePayerCostIndex(t *testing.T) {
	t.Parallel()

	payerKey, _ := crypto.GenerateKey()
	payer := crypto.PubkeyToAddress(payerKey.PublicKey)
	senderKey, _ := crypto.GenerateKey()

	shared := make(map[common.Address]*uint256.Int)
	list := newList(true)
	list.feePayerCost = shared // wire the list to the shared index, as promoteTx does

	// Each tx's fee-payer cost = gasFee * gas = 2 * 100000 = 200000.
	tx0 := feeDelegatedDynamicFeeTx(0, 100000, big.NewInt(2), big.NewInt(1), big.NewInt(100), senderKey, payerKey)
	tx1 := feeDelegatedDynamicFeeTx(1, 100000, big.NewInt(2), big.NewInt(1), big.NewInt(100), senderKey, payerKey)

	if ok, _ := list.Add(tx0, DefaultConfig.PriceBump); !ok {
		t.Fatal("failed to add tx0")
	}
	if got, want := shared[payer].ToBig(), big.NewInt(200_000); got.Cmp(want) != 0 {
		t.Fatalf("index after tx0: have %v, want %v", got, want)
	}
	if ok, _ := list.Add(tx1, DefaultConfig.PriceBump); !ok {
		t.Fatal("failed to add tx1")
	}
	if got, want := shared[payer].ToBig(), big.NewInt(400_000); got.Cmp(want) != 0 {
		t.Fatalf("index after tx1: have %v, want %v", got, want)
	}

	// Remove the highest nonce first so strict-mode invalidation doesn't also drop
	// the lower nonce.
	if ok, _ := list.Remove(tx1); !ok {
		t.Fatal("failed to remove tx1")
	}
	if got, want := shared[payer].ToBig(), big.NewInt(200_000); got.Cmp(want) != 0 {
		t.Fatalf("index after removing tx1: have %v, want %v", got, want)
	}

	// Removing the last tx for this payer should delete the entry entirely.
	if ok, _ := list.Remove(tx0); !ok {
		t.Fatal("failed to remove tx0")
	}
	if _, ok := shared[payer]; ok {
		t.Fatalf("index entry for payer should be deleted, have %v", shared[payer])
	}
}

// TestListFeePayerCostUnwired verifies that a list which is not wired to the
// shared index (e.g. a queue list) leaves the index untouched and does not panic.
func TestListFeePayerCostUnwired(t *testing.T) {
	t.Parallel()

	payerKey, _ := crypto.GenerateKey()
	senderKey, _ := crypto.GenerateKey()

	list := newList(true) // feePayerCost left nil, as queue lists are
	tx := feeDelegatedDynamicFeeTx(0, 100000, big.NewInt(2), big.NewInt(1), big.NewInt(100), senderKey, payerKey)

	if ok, _ := list.Add(tx, DefaultConfig.PriceBump); !ok {
		t.Fatal("failed to add tx to unwired list")
	}
	if list.feePayerCost != nil {
		t.Fatal("unwired list must not allocate a fee-payer index")
	}
	// Removal must likewise be a no-op rather than panic.
	if ok, _ := list.Remove(tx); !ok {
		t.Fatal("failed to remove tx from unwired list")
	}
}

// TestFeeDelegationPromoteRespectsCumulativeBudget reproduces the queue-accounting
// gap: the shared fee-payer index only tracks pending transactions, so a
// fee-delegated tx sitting in the queue (behind a nonce gap) is not counted at
// admission time. When the gap is later filled and the queued tx becomes
// promotable, promoting it together with the tx that filled the gap would push the
// fee payer's committed cost past its balance. promoteExecutables must enforce the
// cumulative fee-payer budget and drop the excess instead of over-committing.
func TestFeeDelegationPromoteRespectsCumulativeBudget(t *testing.T) {
	t.Parallel()

	pool, _ := setupAdventurePool()
	defer pool.Close()

	payerKey, _ := crypto.GenerateKey()
	payer := crypto.PubkeyToAddress(payerKey.PublicKey)
	senderKey, _ := crypto.GenerateKey()
	sender := crypto.PubkeyToAddress(senderKey.PublicKey)

	testAddBalance(pool, sender, new(big.Int).SetUint64(params.Ether))
	// Each tx commits the payer to gasFee*gas = 2*100000 = 200000. Fund the payer
	// for exactly two such transactions, not three.
	const perTxFeePayerCost = 200_000
	testAddBalance(pool, payer, big.NewInt(2*perTxFeePayerCost))

	// nonce 0: no gap, promotes into pending and is counted in the index.
	if err := pool.addRemoteSync(feeDelegatedDynamicFeeTx(0, 100000, big.NewInt(2), big.NewInt(1), big.NewInt(100), senderKey, payerKey)); err != nil {
		t.Fatalf("nonce 0 rejected: %v", err)
	}
	// nonce 2: nonce gap (1 is missing) keeps it in the queue, where it is NOT
	// counted in the pending fee-payer index. Admission sees only nonce 0's cost
	// (200000) + this tx (200000) = 400000 <= balance, so it is accepted.
	if err := pool.addRemoteSync(feeDelegatedDynamicFeeTx(2, 100000, big.NewInt(2), big.NewInt(1), big.NewInt(100), senderKey, payerKey)); err != nil {
		t.Fatalf("nonce 2 (queued) rejected: %v", err)
	}
	if pending, queued := pool.Stats(); pending != 1 || queued != 1 {
		t.Fatalf("setup mismatch: pending %d queued %d, want pending 1 queued 1", pending, queued)
	}

	// nonce 1: fills the gap. Now nonce 1 and nonce 2 are both promotable, but
	// promoting both would commit the payer to 3*200000 = 600000 > 400000. The
	// promotion gate must promote only what fits (nonce 1) and hold nonce 2 back in
	// the queue for a later attempt.
	if err := pool.addRemoteSync(feeDelegatedDynamicFeeTx(1, 100000, big.NewInt(2), big.NewInt(1), big.NewInt(100), senderKey, payerKey)); err != nil {
		t.Fatalf("nonce 1 rejected: %v", err)
	}

	// The pending fee-payer commitment must never exceed the payer's balance.
	if got, want := payerCost(pool, payer), big.NewInt(2*perTxFeePayerCost); got.Cmp(want) != 0 {
		t.Fatalf("fee-payer index over-committed: have %v, want %v", got, want)
	}
	// nonce 0 and nonce 1 promote into pending; nonce 2 is held back in the queue
	// (not dropped) because the payer's budget is exhausted for now.
	if pending, queued := pool.Stats(); pending != 2 || queued != 1 {
		t.Fatalf("pool state mismatch: pending %d queued %d, want pending 2 queued 1", pending, queued)
	}
	if err := validatePoolInternals(pool); err != nil {
		t.Fatalf("pool internal state corrupted: %v", err)
	}
}

// TestFeeDelegationPromoteIgnoresGappedQueue verifies that the promotion-time
// fee-payer budget check only considers the contiguous, promotable nonce run and
// ignores transactions stranded behind a nonce gap. Such transactions cannot be
// promoted this round, so they must neither be dropped by the budget check nor
// reserve any of the fee payer's budget (which would wrongly penalize other
// promotable transactions sharing the same fee payer).
func TestFeeDelegationPromoteIgnoresGappedQueue(t *testing.T) {
	t.Parallel()

	pool, _ := setupAdventurePool()
	defer pool.Close()

	payerKey, _ := crypto.GenerateKey()
	payer := crypto.PubkeyToAddress(payerKey.PublicKey)
	senderKey, _ := crypto.GenerateKey()
	sender := crypto.PubkeyToAddress(senderKey.PublicKey)

	testAddBalance(pool, sender, new(big.Int).SetUint64(params.Ether))
	// Payer covers two txs (2 * 200000).
	const perTxFeePayerCost = 200_000
	testAddBalance(pool, payer, big.NewInt(2*perTxFeePayerCost))

	// nonce 0 promotes; the payer now has 200000 committed in pending.
	if err := pool.addRemoteSync(feeDelegatedDynamicFeeTx(0, 100000, big.NewInt(2), big.NewInt(1), big.NewInt(100), senderKey, payerKey)); err != nil {
		t.Fatalf("nonce 0 rejected: %v", err)
	}
	// nonce 3 and nonce 4 sit behind a gap (1 and 2 are missing), so they can never
	// be promoted this round. Each is individually within the payer's remaining
	// budget at admission, so both are accepted into the queue.
	if err := pool.addRemoteSync(feeDelegatedDynamicFeeTx(3, 100000, big.NewInt(2), big.NewInt(1), big.NewInt(100), senderKey, payerKey)); err != nil {
		t.Fatalf("nonce 3 rejected: %v", err)
	}
	if err := pool.addRemoteSync(feeDelegatedDynamicFeeTx(4, 100000, big.NewInt(2), big.NewInt(1), big.NewInt(100), senderKey, payerKey)); err != nil {
		t.Fatalf("nonce 4 rejected: %v", err)
	}

	// The gapped txs are not promotable, so the promotion budget check must leave
	// them untouched: both remain queued and only nonce 0 is counted in the index.
	if pending, queued := pool.Stats(); pending != 1 || queued != 2 {
		t.Fatalf("pool state mismatch: pending %d queued %d, want pending 1 queued 2", pending, queued)
	}
	if got, want := payerCost(pool, payer), big.NewInt(perTxFeePayerCost); got.Cmp(want) != 0 {
		t.Fatalf("fee-payer index mismatch: have %v, want %v", got, want)
	}
	if err := validatePoolInternals(pool); err != nil {
		t.Fatalf("pool internal state corrupted: %v", err)
	}
}
