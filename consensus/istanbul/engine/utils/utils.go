package utils

import (
	"bytes"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
	"math/big"
)

var (
	NilUncleHash = types.CalcUncleHash(nil) // Always Keccak256(RLP([])) as uncles are meaningless outside of PoW.
)

// WriteCommittedSeals writes the extra-data field of a block header with given committed seals.
func WriteCommittedSeals(committedSeals [][]byte) ApplyExtra {
	return func(extra *types.IstanbulExtra) error {
		if len(committedSeals) == 0 {
			return istanbul.ErrInvalidCommittedSeals
		}

		for _, seal := range committedSeals {
			if len(seal) != types.IstanbulExtraSeal {
				return istanbul.ErrInvalidCommittedSeals
			}
		}

		extra.CommittedSeal = make([][]byte, len(committedSeals))
		copy(extra.CommittedSeal, committedSeals)

		return nil
	}
}

// WriteRoundNumber writes the extra-data field of a block header with given round.
func WriteRoundNumber(round *big.Int) ApplyExtra {
	return func(extra *types.IstanbulExtra) error {
		extra.Round = uint32(round.Uint64())
		return nil
	}
}

func WriteValidators(validators []common.Address) ApplyExtra {
	return func(extra *types.IstanbulExtra) error {
		extra.Validators = validators
		return nil
	}
}

// FIXME: Need to update this for Istanbul
// SigHash returns the hash which is used as input for the Istanbul
// signing. It is the hash of the entire header apart from the 65 byte signature
// contained at the end of the extra data.
//
// Note, the method requires the extra data to be at least 65 bytes, otherwise it
// panics. This is done to avoid accidentally using both forms (signature present
// or not), which could be abused to produce different hashes for the same header.
func SigHash(header *types.Header) (hash common.Hash) {
	hasher := sha3.NewLegacyKeccak256()
	rlp.Encode(hasher, types.IstanbulFilteredHeader(header))
	hasher.Sum(hash[:0])
	return hash
}

// PrepareCommittedSeal returns a committed seal for the given hash
func PrepareCommittedSeal(header *types.Header, round uint32) []byte {
	h := types.CopyHeader(header)
	return h.IstanbulHashWithRoundNumber(round).Bytes()
}

func WriteVote(candidate common.Address, authorize bool) ApplyExtra {
	return func(extra *types.IstanbulExtra) error {
		voteType := types.IstanbulDropVote
		if authorize {
			voteType = types.IstanbulAuthVote
		}

		vote := &types.ValidatorVote{RecipientAddress: candidate, VoteType: voteType}
		extra.Vote = vote
		return nil
	}
}

func GetExtra(header *types.Header) (*types.IstanbulExtra, error) {
	if len(header.Extra) < types.IstanbulExtraVanity {
		// In this scenario, the header extradata only contains client specific information, hence create a new extra and set vanity
		vanity := append(header.Extra, bytes.Repeat([]byte{0x00}, types.IstanbulExtraVanity-len(header.Extra))...)
		return &types.IstanbulExtra{
			VanityData:    vanity,
			Validators:    []common.Address{},
			CommittedSeal: [][]byte{},
			Round:         0,
			Vote:          nil,
		}, nil
	}

	// This is the case when Extra has already been set
	return types.ExtractIstanbulExtra(header)
}

func SetExtra(h *types.Header, extra *types.IstanbulExtra) error {
	payload, err := rlp.EncodeToBytes(extra)
	if err != nil {
		return err
	}

	h.Extra = payload
	return nil
}
