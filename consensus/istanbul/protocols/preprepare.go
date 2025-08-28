package protocols

import (
	"fmt"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

type Preprepare struct {
	CommonPayload
	Proposal                  istanbul.Proposal
	JustificationRoundChanges []*SignedRoundChangePayload
	JustificationPrepares     []*Prepare
}

func NewPreprepare(sequence *big.Int, round *big.Int, proposal istanbul.Proposal) *Preprepare {
	return &Preprepare{
		CommonPayload: CommonPayload{
			code:     PreprepareCode,
			Sequence: sequence,
			Round:    round,
		},
		Proposal: proposal,
	}
}

func (m *Preprepare) EncodePayloadForSigning() ([]byte, error) {
	return rlp.EncodeToBytes(
		[]interface{}{
			m.Code(),
			[]interface{}{m.Sequence, m.Round, m.Proposal},
		})
}

func (m *Preprepare) EncodeRLP(w io.Writer) error {
	val := []interface{}{
		[]interface{}{
			[]interface{}{m.Sequence, m.Round, m.Proposal},
			m.signature,
		},
		[]interface{}{
			m.JustificationRoundChanges,
			m.JustificationPrepares,
		},
	}
	// encode extra only if there are any data
	// ##CROSS: blob sidecars
	if m.Proposal.Sidecars().Len() > 0 {
		val = append(val, []any{m.Proposal.Sidecars()})
	}
	// ##
	return rlp.Encode(w, val)
}

func (m *Preprepare) DecodeRLP(stream *rlp.Stream) error {
	var message struct {
		SignedPayload struct {
			Payload struct {
				Sequence *big.Int
				Round    *big.Int
				Proposal *types.Block
			}
			Signature []byte
		}
		Justification struct {
			RoundChanges []*SignedRoundChangePayload
			Prepares     []*Prepare
		}
		Extra struct {
			Sidecars types.BlobSidecars // ##CROSS: blob sidecars
		} `rlp:"optional"`
	}
	if err := stream.Decode(&message); err != nil {
		return err
	}
	m.code = PreprepareCode
	m.Sequence = message.SignedPayload.Payload.Sequence
	m.Round = message.SignedPayload.Payload.Round
	m.Proposal = message.SignedPayload.Payload.Proposal.WithSidecars(message.Extra.Sidecars) // ##CROSS: blob sidecars
	m.signature = message.SignedPayload.Signature
	m.JustificationPrepares = message.Justification.Prepares
	m.JustificationRoundChanges = message.Justification.RoundChanges
	return nil
}

func (m *Preprepare) String() string {
	return fmt.Sprintf("code: %d, sequence: %d, round: %d, proposal: %v", m.code, m.Sequence, m.Round, m.Proposal.Hash().Hex())
}
