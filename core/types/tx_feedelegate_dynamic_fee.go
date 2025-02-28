// Copyright 2021 The go-ethereum Authors
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

package types

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type FeeDelegatedDynamicFeeTx struct { // ##CROSS: fee delegation
	SenderTx DynamicFeeTx
	FeePayer *common.Address `rlp:"nil"`
	// Signature values
	Vf *big.Int `json:"vf" gencodec:"required"` // feePayer V
	Rf *big.Int `json:"rf" gencodec:"required"` // feePayer R
	Sf *big.Int `json:"sf" gencodec:"required"` // feePayer S
}

func (tx *FeeDelegatedDynamicFeeTx) SetSenderTx(senderTx DynamicFeeTx) {
	tx.SenderTx.ChainID = senderTx.ChainID
	tx.SenderTx.Nonce = senderTx.Nonce
	tx.SenderTx.GasFeeCap = senderTx.GasFeeCap
	tx.SenderTx.GasTipCap = senderTx.GasTipCap
	tx.SenderTx.Gas = senderTx.Gas
	tx.SenderTx.To = senderTx.To
	tx.SenderTx.Value = senderTx.Value
	tx.SenderTx.Data = senderTx.Data
	copy(tx.SenderTx.AccessList, senderTx.AccessList)

	v, r, s := senderTx.rawSignatureValues()
	tx.SenderTx.V = v
	tx.SenderTx.R = r
	tx.SenderTx.S = s
}

// copy creates a deep copy of the transaction data and initializes all fields.
func (tx *FeeDelegatedDynamicFeeTx) copy() TxData {
	cpy := &FeeDelegatedDynamicFeeTx{
		SenderTx: tx.copySenderTx(),
		FeePayer: copyAddressPtr(tx.FeePayer),
		Vf:       new(big.Int),
		Rf:       new(big.Int),
		Sf:       new(big.Int),
	}
	if tx.Vf != nil {
		cpy.Vf.Set(tx.Vf)
	}
	if tx.Rf != nil {
		cpy.Rf.Set(tx.Rf)
	}
	if tx.Sf != nil {
		cpy.Sf.Set(tx.Sf)
	}
	return cpy
}

func (tx *FeeDelegatedDynamicFeeTx) copySenderTx() DynamicFeeTx {
	cpy := DynamicFeeTx{
		Nonce: tx.SenderTx.Nonce,
		To:    copyAddressPtr(tx.SenderTx.To),
		Data:  common.CopyBytes(tx.SenderTx.Data),
		Gas:   tx.SenderTx.Gas,
		// These are copied below.
		AccessList: make(AccessList, len(tx.SenderTx.AccessList)),
		Value:      new(big.Int),
		ChainID:    new(big.Int),
		GasTipCap:  new(big.Int),
		GasFeeCap:  new(big.Int),
		V:          new(big.Int),
		R:          new(big.Int),
		S:          new(big.Int),
	}
	copy(cpy.AccessList, tx.SenderTx.accessList())
	if tx.SenderTx.Value != nil {
		cpy.Value.Set(tx.SenderTx.Value)
	}
	if tx.SenderTx.ChainID != nil {
		cpy.ChainID.Set(tx.SenderTx.ChainID)
	}
	if tx.SenderTx.GasTipCap != nil {
		cpy.GasTipCap.Set(tx.SenderTx.GasTipCap)
	}
	if tx.SenderTx.GasFeeCap != nil {
		cpy.GasFeeCap.Set(tx.SenderTx.GasFeeCap)
	}
	if tx.SenderTx.V != nil {
		cpy.V.Set(tx.SenderTx.V)
	}
	if tx.SenderTx.R != nil {
		cpy.R.Set(tx.SenderTx.R)
	}
	if tx.SenderTx.S != nil {
		cpy.S.Set(tx.SenderTx.S)
	}
	return cpy
}

// accessors for innerTx.
func (tx *FeeDelegatedDynamicFeeTx) txType() byte              { return FeeDelegatedDynamicFeeTxType }
func (tx *FeeDelegatedDynamicFeeTx) chainID() *big.Int         { return tx.SenderTx.ChainID }
func (tx *FeeDelegatedDynamicFeeTx) accessList() AccessList    { return tx.SenderTx.AccessList }
func (tx *FeeDelegatedDynamicFeeTx) data() []byte              { return tx.SenderTx.Data }
func (tx *FeeDelegatedDynamicFeeTx) gas() uint64               { return tx.SenderTx.Gas }
func (tx *FeeDelegatedDynamicFeeTx) gasFeeCap() *big.Int       { return tx.SenderTx.GasFeeCap }
func (tx *FeeDelegatedDynamicFeeTx) gasTipCap() *big.Int       { return tx.SenderTx.GasTipCap }
func (tx *FeeDelegatedDynamicFeeTx) gasPrice() *big.Int        { return tx.SenderTx.GasFeeCap }
func (tx *FeeDelegatedDynamicFeeTx) value() *big.Int           { return tx.SenderTx.Value }
func (tx *FeeDelegatedDynamicFeeTx) nonce() uint64             { return tx.SenderTx.Nonce }
func (tx *FeeDelegatedDynamicFeeTx) to() *common.Address       { return tx.SenderTx.To }
func (tx *FeeDelegatedDynamicFeeTx) feePayer() *common.Address { return tx.FeePayer }
func (tx *FeeDelegatedDynamicFeeTx) rawFeePayerSignatureValues() (v, r, s *big.Int) {
	return tx.Vf, tx.Rf, tx.Sf
}

func (tx *FeeDelegatedDynamicFeeTx) rawSignatureValues() (v, r, s *big.Int) {
	return tx.SenderTx.rawSignatureValues()
}

func (tx *FeeDelegatedDynamicFeeTx) setSignatureValues(chainID, v, r, s *big.Int) {
	tx.Vf, tx.Rf, tx.Sf = v, r, s
}

func (tx *FeeDelegatedDynamicFeeTx) encode(b *bytes.Buffer) error {
	return rlp.Encode(b, tx)
}

func (tx *FeeDelegatedDynamicFeeTx) decode(input []byte) error {
	return rlp.DecodeBytes(input, tx)
}

func (tx *FeeDelegatedDynamicFeeTx) effectiveGasPrice(dst *big.Int, baseFee *big.Int) *big.Int {
	return tx.SenderTx.effectiveGasPrice(dst, baseFee)
}
