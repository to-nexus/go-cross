// Copyright 2017 The go-ethereum Authors
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
	"errors"
	"io"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rlp"
)

// ##CROSS: bls seal
const (
	BLSPublicKeyLength = 48
	BLSSignatureLength = 96
)

type (
	BLSPublicKey [BLSPublicKeyLength]byte
	BLSSignature [BLSSignatureLength]byte
)

var (
	blsPublicKeyT = reflect.TypeOf(BLSPublicKey{})
	blsSignatureT = reflect.TypeOf(BLSSignature{})
)

// BytesToBLSPublicKey converts a byte slice to a BLS public key.
// If b is larger than BLSPublicKeyLength, it will be cropped from the left.
func BytesToBLSPublicKey(b []byte) BLSPublicKey {
	var key BLSPublicKey
	if len(b) > len(key) {
		b = b[len(b)-BLSPublicKeyLength:]
	}
	copy(key[BLSPublicKeyLength-len(b):], b)
	return key
}

// BytesToBLSSignature converts a byte slice to a BLS signature.
// If b is larger than BLSSignatureLength, it will be cropped from the left.
func BytesToBLSSignature(b []byte) BLSSignature {
	var sig BLSSignature
	if len(b) > len(sig) {
		b = b[len(b)-BLSSignatureLength:]
	}
	copy(sig[BLSSignatureLength-len(b):], b)
	return sig
}

// Bytes returns the byte slice representation of the BLS public key.
func (b BLSPublicKey) Bytes() []byte {
	return b[:]
}

// String implements the stringer interface.
func (b BLSPublicKey) String() string {
	return hexutil.Encode(b[:])
}

// MarshalText returns the hex representation of the BLS public key.
func (b BLSPublicKey) MarshalText() ([]byte, error) {
	return hexutil.Bytes(b[:]).MarshalText()
}

// UnmarshalJSON parses a BLS public key in hex syntax.
func (b *BLSPublicKey) UnmarshalJSON(input []byte) error {
	return hexutil.UnmarshalFixedJSON(blsPublicKeyT, input, b[:])
}

// UnmarshalText parses a BLS public key in hex syntax.
func (b *BLSPublicKey) UnmarshalText(input []byte) error {
	return hexutil.UnmarshalFixedText("BLSPublicKey", input, b[:])
}

// Bytes returns the byte slice representation of the BLS signature.
func (b BLSSignature) Bytes() []byte {
	return b[:]
}

// String implements the stringer interface.
func (b BLSSignature) String() string {
	return hexutil.Encode(b[:])
}

// MarshalText returns the hex representation of the BLS signature.
func (b BLSSignature) MarshalText() ([]byte, error) {
	return hexutil.Bytes(b[:]).MarshalText()
}

// UnmarshalJSON parses a BLS signature in hex syntax.
func (b *BLSSignature) UnmarshalJSON(input []byte) error {
	return hexutil.UnmarshalFixedJSON(blsSignatureT, input, b[:])
}

// UnmarshalText parses a BLS signature in hex syntax.
func (b *BLSSignature) UnmarshalText(input []byte) error {
	return hexutil.UnmarshalFixedText("BLSSignature", input, b[:])
}

// ##

var (
	// ##CROSS: istanbul digest
	// IstanbulDigest represents a hash of "Cross Istanbul. "
	// to identify whether the block is from Istanbul consensus engine
	IstanbulDigest = common.HexToHash("0x43726f737320497374616e62756c2e2000000000000000000000000000000000")

	IstanbulExtraVanity  = 32 // Fixed number of extra-data bytes reserved for validator vanity
	IstanbulExtraSeal    = 65 // Fixed number of extra-data bytes reserved for validator seal
	IstanbulExtraSealBLS = 96 // Fixed number of extra-data bytes reserved for validator BLS seal // ##CROSS: bls seal

	IstanbulAuthVote = byte(0xFF) // Magic number to vote on adding a new validator
	IstanbulDropVote = byte(0x00) // Magic number to vote on removing a validator.

	// ErrInvalidIstanbulHeaderExtra is returned if the length of extra-data is less than 32 bytes
	ErrInvalidIstanbulHeaderExtra = errors.New("invalid istanbul header extra-data")
)

// ##CROSS: istanbul digest
func IsIstanbulDigest(digest common.Hash) bool {
	return bytes.Equal(digest[:16], IstanbulDigest[:16])
}

func MakeIstanbulDigest(seed common.Hash) common.Hash {
	digest := IstanbulDigest
	copy(digest[16:], seed[:16])
	return digest
}

// ##

// IstanbulExtra represents header extradata for qbft protocol
type IstanbulExtra struct {
	VanityData    []byte
	Validators    []common.Address
	Vote          *ValidatorVote
	Round         uint32
	CommittedSeal [][]byte
	RandomReveal  []byte
	// ##CROSS: bls seal
	SignersBitset []uint64
	Signers       []BLSPublicKey
	// ##
}

// EncodeRLP serializes qist into the Ethereum RLP format.
func (qst *IstanbulExtra) EncodeRLP(w io.Writer) error {
	val := []interface{}{
		qst.VanityData,
		qst.Validators,
		qst.Vote,
		qst.Round,
		qst.CommittedSeal,
		qst.RandomReveal,
	}
	// ##CROSS: bls seal
	if qst.SignersBitset != nil || qst.Signers != nil {
		val = append(val, qst.SignersBitset, qst.Signers)
	}
	// ##
	return rlp.Encode(w, val)
}

// DecodeRLP implements rlp.Decoder, and load the IstanbulExtra fields from a RLP stream.
func (qst *IstanbulExtra) DecodeRLP(s *rlp.Stream) error {
	var extra struct {
		VanityData    []byte
		Validators    []common.Address
		Vote          *ValidatorVote `rlp:"nil"`
		Round         uint32
		CommittedSeal [][]byte
		RandomReveal  []byte
		// ##CROSS: bls seal
		SignersBitset []uint64       `rlp:"optional"`
		Signers       []BLSPublicKey `rlp:"optional"`
		// ##
	}
	if err := s.Decode(&extra); err != nil {
		return err
	}
	qst.VanityData = extra.VanityData
	qst.Validators = extra.Validators
	qst.Vote = extra.Vote
	qst.Round = extra.Round
	qst.CommittedSeal = extra.CommittedSeal
	qst.RandomReveal = extra.RandomReveal
	// ##CROSS: bls seal
	qst.SignersBitset = extra.SignersBitset
	qst.Signers = extra.Signers
	// ##
	return nil
}

type ValidatorVote struct {
	RecipientAddress common.Address
	VoteType         byte
}

// EncodeRLP serializes ValidatorVote into the Ethereum RLP format.
func (vv *ValidatorVote) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, []interface{}{
		vv.RecipientAddress,
		vv.VoteType,
	})
}

// DecodeRLP implements rlp.Decoder, and load the ValidatorVote fields from a RLP stream.
func (vv *ValidatorVote) DecodeRLP(s *rlp.Stream) error {
	var validatorVote struct {
		RecipientAddress common.Address
		VoteType         byte
	}
	if err := s.Decode(&validatorVote); err != nil {
		return err
	}
	vv.RecipientAddress, vv.VoteType = validatorVote.RecipientAddress, validatorVote.VoteType
	return nil
}

// ExtractIstanbulExtra extracts all values of the IstanbulExtra from the header. It returns an
// error if the length of the given extra-data is less than 32 bytes or the extra-data can not
// be decoded.
func ExtractIstanbulExtra(h *Header) (*IstanbulExtra, error) {
	extra := new(IstanbulExtra)
	err := rlp.DecodeBytes(h.Extra[:], extra)
	if err != nil {
		return nil, err
	}
	return extra, nil
}

// IstanbulFilteredHeader returns a filtered header which some information (like committed seals, round, validator vote)
// are clean to fulfill the Istanbul hash rules. It returns nil if the extra-data cannot be
// decoded/encoded by rlp.
func IstanbulFilteredHeader(h *Header) *Header {
	return IstanbulFilteredHeaderWithRound(h, 0)
}

// IstanbulFilteredHeaderWithRound returns the copy of the header with round number set to the given round number
// and commit seal set to its null value
func IstanbulFilteredHeaderWithRound(h *Header, round uint32) *Header {
	newHeader := CopyHeader(h)
	extra, err := ExtractIstanbulExtra(newHeader)
	if err != nil {
		return nil
	}

	extra.CommittedSeal = [][]byte{}
	extra.Round = round

	// ##CROSS: bls seal
	extra.SignersBitset = nil
	extra.Signers = nil
	// ##

	payload, err := rlp.EncodeToBytes(&extra)
	if err != nil {
		return nil
	}

	newHeader.Extra = payload
	return newHeader
}
