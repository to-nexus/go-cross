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

package core

// ##CROSS: fee delegation
//
// End-to-end execution test for fee-delegated transactions on an Istanbul-consensus
// (pre-merge) chain. Unlike the engine-API/post-merge simulated backend, blocks here
// are produced with GenerateChain + an ethash faker, which stays pre-merge and is
// compatible with the Adventure fork (see consensus/beacon.isPostMerge, which forces
// pre-merge whenever Adventure is active).
//
// The test mines real fee-delegated transactions and asserts on-chain balances:
//   - the fee payer's balance drops by the gas cost (fee delegation actually charges
//     the payer, not the sender),
//   - under Istanbul consensus the entire gas fee is credited to the configured
//     beneficiary (no base-fee burn), so the payer's drop equals the beneficiary's gain,
//   - the sender only loses the transferred value and is never charged gas.

import (
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/triedb"
)

// feeDelegationIstanbulConfig returns a pre-merge chain config with the Adventure
// fork active (so fee-delegated txs are accepted) and Istanbul consensus enabled with
// the given beneficiary (so all gas fees are routed to it with no base-fee burn).
func feeDelegationIstanbulConfig(beneficiary common.Address) *params.ChainConfig {
	// Base off the dev-chain config, which already has a valid fork ordering (all
	// timestamp forks enabled at 0). Adventure sits after Shanghai in the fork order,
	// so it cannot be enabled on a config that lacks the earlier timestamp forks.
	cfg := *params.AllDevChainProtocolChanges
	adventure := uint64(0)
	cfg.AdventureTime = &adventure
	b := beneficiary
	cfg.Istanbul = &params.IstanbulConfig{
		EpochLength:        30000,
		BlockPeriodSeconds: 1,
		ProposerPolicy:     0,
		Beneficiary:        &b,
	}
	return &cfg
}

// signFeeDelegatedTx builds a signed fee-delegated dynamic-fee transaction. The inner
// dynamic-fee tx is signed by senderKey; the fee-delegated wrapper, which names a
// separate fee payer, is signed by payerKey. Because the fee payer differs from the
// sender, the payer covers the gas (gasFeeCap*gas) and the sender only the value.
func signFeeDelegatedTx(chainID *big.Int, nonce uint64, to common.Address, value, gasFeeCap, tip *big.Int, gas uint64, senderKey, payerKey *ecdsa.PrivateKey) *types.Transaction {
	londonSigner := types.NewLondonSigner(chainID)
	feeSigner := types.NewFeeDelegationSigner(chainID)
	payer := crypto.PubkeyToAddress(payerKey.PublicKey)

	senderTx, err := types.SignTx(types.NewTx(&types.DynamicFeeTx{
		ChainID: chainID, Nonce: nonce, GasTipCap: tip, GasFeeCap: gasFeeCap,
		Gas: gas, To: &to, Value: value,
	}), londonSigner, senderKey)
	if err != nil {
		panic(err)
	}
	v, r, s := senderTx.RawSignatureValues()
	inner := types.DynamicFeeTx{
		ChainID: chainID, Nonce: nonce, GasTipCap: tip, GasFeeCap: gasFeeCap,
		Gas: gas, To: &to, Value: value, V: v, R: r, S: s,
	}
	feeTx, err := types.SignTx(types.NewTx(types.NewFeeDelegatedDynamicFeeTx(&payer, inner)), feeSigner, payerKey)
	if err != nil {
		panic(err)
	}
	return feeTx
}

func TestFeeDelegationMinedChargesPayer(t *testing.T) {
	senderKey, _ := crypto.GenerateKey()
	sender := crypto.PubkeyToAddress(senderKey.PublicKey)
	payerKey, _ := crypto.GenerateKey()
	payer := crypto.PubkeyToAddress(payerKey.PublicKey)
	beneficiary := common.HexToAddress("0x00000000000000000000000000000000be0eF1c1")
	miner := common.HexToAddress("0x00000000000000000000000000000000000c01cb") // block coinbase (gets only the ethash reward)
	to := common.HexToAddress("0x00000000000000000000000000000000000000aa")

	const (
		gas   = uint64(21000)
		numTx = 3
	)
	tip := big.NewInt(params.GWei)
	gasFeeCap := big.NewInt(2 * params.GWei)
	value := big.NewInt(1)

	senderStart := new(big.Int).SetUint64(1e18)
	payerStart := new(big.Int).SetUint64(1e18)

	cfg := feeDelegationIstanbulConfig(beneficiary)
	gspec := &Genesis{
		Config:   cfg,
		GasLimit: 30_000_000,
		Alloc: types.GenesisAlloc{
			sender: {Balance: senderStart},
			payer:  {Balance: payerStart},
		},
	}

	// Produce the block with GenerateChain rather than a full miner/InsertChain: the
	// Adventure fork forces pre-merge state, and the only pre-merge engines (ethash /
	// pre-merge beacon) reject the Shanghai fork that Adventure requires — so block
	// *verification* needs the Istanbul engine. GenerateChain applies the transactions
	// through the real state-transition path (ApplyTransaction) without header
	// verification, which is exactly the execution behaviour under test.
	db := rawdb.NewMemoryDatabase()
	tdb := triedb.NewDatabase(db, triedb.HashDefaults)
	defer tdb.Close()
	genesisBlock, err := gspec.Commit(db, tdb)
	if err != nil {
		t.Fatalf("commit genesis: %v", err)
	}
	blocks, receipts := GenerateChain(cfg, genesisBlock, ethash.NewFaker(), db, 1, func(i int, b *BlockGen) {
		b.SetCoinbase(miner)
		for n := uint64(0); n < numTx; n++ {
			b.AddTx(signFeeDelegatedTx(cfg.ChainID, n, to, value, gasFeeCap, tip, gas, senderKey, payerKey))
		}
	})

	// Sanity: every fee-delegated tx actually made it into the block.
	if got := len(receipts[0]); got != numTx {
		t.Fatalf("mined %d txs, want %d", got, numTx)
	}

	st, err := state.New(blocks[0].Root(), state.NewDatabase(tdb, nil))
	if err != nil {
		t.Fatalf("open state: %v", err)
	}
	payerEnd := st.GetBalance(payer).ToBig()
	senderEnd := st.GetBalance(sender).ToBig()
	beneficiaryEnd := st.GetBalance(beneficiary).ToBig()

	payerDrop := new(big.Int).Sub(payerStart, payerEnd)
	senderDrop := new(big.Int).Sub(senderStart, senderEnd)

	// 1. The fee payer was actually charged (the original symptom under test).
	if payerDrop.Sign() <= 0 {
		t.Fatalf("fee payer balance did not drop: start %v end %v", payerStart, payerEnd)
	}

	// 2. Istanbul consensus routes the whole gas fee to the beneficiary with no burn,
	//    so the payer's drop must equal the beneficiary's gain exactly.
	if payerDrop.Cmp(beneficiaryEnd) != 0 {
		t.Fatalf("payer drop %v != beneficiary gain %v (fee not fully routed to beneficiary)", payerDrop, beneficiaryEnd)
	}

	// 3. The sender is never charged gas: it only loses the transferred value.
	wantSenderDrop := new(big.Int).Mul(value, big.NewInt(numTx))
	if senderDrop.Cmp(wantSenderDrop) != 0 {
		t.Fatalf("sender drop %v, want %v (sender must not pay gas)", senderDrop, wantSenderDrop)
	}
}
