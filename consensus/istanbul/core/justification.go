package core

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/consensus/istanbul/protocols"
	"github.com/ethereum/go-ethereum/log"
)

// Returns true if the `proposal` is justified by the set `roundChangeMessages` of ROUND-CHANGE messages
// and by the set `prepareMessages` of PREPARE messages.
// For this we must either have:
//   - a quorum of ROUND-CHANGE messages with preparedRound and preparedBlockDigest equal to nil; or
//   - a ROUND-CHANGE message (1) whose preparedRound is not nil and is equal or higher than the
//     preparedRound of `quorumSize` ROUND-CHANGE messages and (2) whose preparedRound and
//     preparedBlockDigest match the round and block of `quorumSize` PREPARE messages.
func isJustified(
	proposal istanbul.Proposal,
	roundChangeMessages []*protocols.SignedRoundChangePayload,
	prepareMessages []*protocols.Prepare,
	quorumSize int,
	hasBadProposal func(common.Hash) bool) error {
	// Check the size of the set of ROUND-CHANGE messages
	if len(roundChangeMessages) < quorumSize {
		return errors.New("number of roundchange messages is less than required quorum of messages")
	}

	// Check the size of the set of PREPARE messages
	if len(prepareMessages) != 0 && len(prepareMessages) < quorumSize {
		return errors.New("number of prepared messages is less than required quorum of messages")
	}

	// If there are PREPARE messages, they all need to have the same round and match `proposal`
	var preparedRound *big.Int
	if len(prepareMessages) > 0 {
		preparedRound = prepareMessages[0].Round
		for _, spp := range prepareMessages {
			if preparedRound.Cmp(spp.Round) != 0 || proposal.Hash() != spp.Digest {
				return errors.New("prepared messages do not have same round or do not match proposal")
			}
		}
	}

	if preparedRound == nil {
		return hasQuorumOfRoundChangeMessagesForNil(roundChangeMessages, quorumSize, hasBadProposal)
	} else {
		return hasQuorumOfRoundChangeMessagesForPreparedRoundAndBlock(roundChangeMessages, preparedRound, proposal, quorumSize, hasBadProposal)
	}
}

// Checks whether a set of ROUND-CHANGE messages has `quorumSize` messages with nil prepared round and
// prepared block.
func hasQuorumOfRoundChangeMessagesForNil(roundChangeMessages []*protocols.SignedRoundChangePayload, quorumSize int, hasBadProposal func(common.Hash) bool) error {
	nilCount := 0
	for _, m := range roundChangeMessages {
		log.Trace("Istanbul: hasQuorumOfRoundChangeMessagesForNil", "rc", m)
		// ##CROSS: bad block mitigation
		// If the round-change message has a prepared block but is bad, we treat it as nil and count it towards the quorum
		if isNilPrepared(m) || isBadPrepared(m, hasBadProposal) {
			nilCount++
			if nilCount == quorumSize {
				return nil
			}
		}
	}
	return errors.New("quorum of roundchange messages with nil prepared round not found")
}

// Checks whether a set of ROUND-CHANGE messages has some message with `preparedRound` and `preparedBlockDigest`,
// and has `quorumSize` messages with prepared round equal to nil or equal or lower than `preparedRound`.
func hasQuorumOfRoundChangeMessagesForPreparedRoundAndBlock(roundChangeMessages []*protocols.SignedRoundChangePayload, preparedRound *big.Int, preparedBlock istanbul.Proposal, quorumSize int, hasBadProposal func(common.Hash) bool) error {
	lowerOrEqualRoundCount := 0
	hasMatchingMessage := false
	for _, m := range roundChangeMessages {
		log.Trace("Istanbul: hasQuorumOfRoundChangeMessagesForPreparedRoundAndBlock", "rc", m)
		// ##CROSS: bad block mitigation
		// If the round-change message has a prepared block but is bad, we treat it as nil and count it towards the quorum
		if m.PreparedRound == nil || isBadPrepared(m, hasBadProposal) || m.PreparedRound.Cmp(preparedRound) <= 0 {
			lowerOrEqualRoundCount++
			if !isBadPrepared(m, hasBadProposal) && m.PreparedRound != nil && m.PreparedRound.Cmp(preparedRound) == 0 && m.PreparedDigest == preparedBlock.Hash() {
				// Even though the bad block counts toward the quorum, it cannot be justified as the prepared block
				hasMatchingMessage = true
			}
			if lowerOrEqualRoundCount >= quorumSize && hasMatchingMessage {
				return nil
			}
		}
	}

	return errors.New("quorum of roundchange messages with prepared round and proposal not found")
}

func isNilPrepared(m *protocols.SignedRoundChangePayload) bool {
	return (m.PreparedRound == nil || m.PreparedRound.Cmp(common.Big0) == 0) && m.PreparedDigest == (common.Hash{})
}

func isBadPrepared(m *protocols.SignedRoundChangePayload, hasBadProposal func(common.Hash) bool) bool {
	return hasBadProposal != nil && m.PreparedRound != nil && m.PreparedDigest != (common.Hash{}) && hasBadProposal(m.PreparedDigest)
}

// Checks whether the round and block of a set of PREPARE messages of at least quorumSize match the
// preparedRound and preparedBlockDigest of a ROUND-CHANGE messages.
func hasMatchingRoundChangeAndPrepares(
	roundChange *protocols.RoundChange, prepareMessages []*protocols.Prepare, quorumSize int) error {
	if len(prepareMessages) < quorumSize {
		return errors.New("number of prepare messages is less than quorum of messages")
	}

	for _, spp := range prepareMessages {
		if spp.Digest != roundChange.PreparedDigest {
			return errors.New("prepared message digest does not match roundchange prepared digest")
		}
		if spp.Round.Cmp(roundChange.PreparedRound) != 0 {
			return errors.New("round number in prepared message does not match prepared round in roundchange")
		}
	}
	return nil
}

// ##CROSS: istanbul validation
// validateProposalSequence checks that a proposal belongs to the given consensus sequence.
func validateProposalSequence(proposal istanbul.Proposal, sequence *big.Int) error {
	if proposal == nil || proposal.Number() == nil || sequence == nil {
		return errors.New("proposal sequence is invalid")
	}
	if proposal.Number().Cmp(sequence) != 0 {
		return errors.New("proposal number does not match sequence")
	}
	return nil
}

type protocolMessage interface {
	protocols.Message
	comparable
}

// validatePrepareSequence checks that the message matches the outer sequence.
func validateMessageSequence[T protocolMessage](messages []T, sequence *big.Int) error {
	if sequence == nil {
		return errors.New("sequence is invalid")
	}
	var zero T
	for _, msg := range messages {
		if msg == zero {
			return errors.New("message is invalid")
		}
		view := msg.View()
		if view.Sequence == nil || view.Round == nil {
			return errors.New("message view is invalid")
		}
		if view.Sequence.Cmp(sequence) != 0 {
			return errors.New("message sequence does not match")
		}
	}
	return nil
}

// validateMessageJustification checks that all source of the given messages are member of validator set and distinct.
func validateMessageJustification[T protocolMessage](messages []T, valSet istanbul.ValidatorSet, quorumSize int) error {
	if len(messages) < quorumSize {
		return errors.New("number of messages is less than required quorum")
	}
	var zero T
	seen := make(map[common.Address]struct{})
	for _, msg := range messages {
		if msg == zero {
			return errors.New("message is invalid")
		}
		source := msg.Source()
		if source == (common.Address{}) {
			return errors.New("message source is empty")
		}
		if valSet != nil {
			if _, validator := valSet.GetByAddress(source); validator == nil {
				return errors.New("message source is not a validator")
			}
		}
		if _, ok := seen[source]; ok {
			return errors.New("duplicate message source")
		}
		seen[source] = struct{}{}
	}
	return nil
}

// ##
