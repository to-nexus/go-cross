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
	"crypto/rand"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"io"
)

var (
	// ##CROSS: istanbul digest
	// IstanbulDigest represents a hash of "Cross Istanbul. "
	// to identify whether the block is from Istanbul consensus engine
	IstanbulDigest = common.HexToHash("0x43726f737320497374616e62756c2e2000000000000000000000000000000000")

	IstanbulExtraVanity = 32 // Fixed number of extra-data bytes reserved for validator vanity
	IstanbulExtraSeal   = 65 // Fixed number of extra-data bytes reserved for validator seal

	IstanbulAuthVote = byte(0xFF) // Magic number to vote on adding a new validator
	IstanbulDropVote = byte(0x00) // Magic number to vote on removing a validator.

	// ErrInvalidIstanbulHeaderExtra is returned if the length of extra-data is less than 32 bytes
	ErrInvalidIstanbulHeaderExtra = errors.New("invalid istanbul header extra-data")
)

// ##CROSS: istanbul digest
func IsIstanbulDigest(digest common.Hash) bool {
	return bytes.Equal(digest[:16], IstanbulDigest[:16])
}

func MakeIstanbulDigest() common.Hash {
	digest := IstanbulDigest

	_, err := rand.Read(digest[16:])
	if err != nil {
		log.Crit("Failed to generate random istanbul hash", "err", err)
	}
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
}

// EncodeRLP serializes qist into the Ethereum RLP format.
func (qst *IstanbulExtra) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, []interface{}{
		qst.VanityData,
		qst.Validators,
		qst.Vote,
		qst.Round,
		qst.CommittedSeal,
	})
}

// DecodeRLP implements rlp.Decoder, and load the IstanbulExtra fields from a RLP stream.
func (qst *IstanbulExtra) DecodeRLP(s *rlp.Stream) error {
	var extra struct {
		VanityData    []byte
		Validators    []common.Address
		Vote          *ValidatorVote `rlp:"nil"`
		Round         uint32
		CommittedSeal [][]byte
	}
	if err := s.Decode(&extra); err != nil {
		return err
	}
	qst.VanityData, qst.Validators, qst.Vote, qst.Round, qst.CommittedSeal = extra.VanityData, extra.Validators, extra.Vote, extra.Round, extra.CommittedSeal

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

	payload, err := rlp.EncodeToBytes(&extra)
	if err != nil {
		return nil
	}

	newHeader.Extra = payload

	return newHeader
}
