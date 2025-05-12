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

package core

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/consensus/istanbul/protocols"
	"github.com/ethereum/go-ethereum/rlp"
)

// broadcastPrepare is called after receiving PRE-PREPARE from proposer node

// It
// - creates a PREPARE message
// - broadcast PREPARE message to other validators
func (c *Core) broadcastPrepare() {
	logger := c.currentLogger(true, nil)

	// Create PREPARE message from the current proposal
	sub := c.current.Subject()
	prepare := protocols.NewPrepare(sub.View.Sequence, sub.View.Round, sub.Digest)
	prepare.SetSource(c.Address())

	// Sign Message
	encodedPayload, err := prepare.EncodePayloadForSigning()
	if err != nil {
		withMsg(logger, prepare).Error("Istanbul: failed to encode payload of PREPARE message", "err", err)
		return
	}
	signature, err := c.backend.Sign(encodedPayload)
	if err != nil {
		withMsg(logger, prepare).Error("Istanbul: failed to sign PREPARE message", "err", err)
		return
	}
	prepare.SetSignature(signature)

	// RLP-encode message
	payload, err := rlp.EncodeToBytes(&prepare)
	if err != nil {
		withMsg(logger, prepare).Error("Istanbul: failed to encode PREPARE message", "err", err)
		return
	}

	withMsg(logger, prepare).Debug("Istanbul: broadcast PREPARE message", "payload", hexutil.Encode(payload))

	// Broadcast RLP-encoded message
	if err = c.backend.Broadcast(c.valSet, prepare.Code(), payload); err != nil {
		withMsg(logger, prepare).Error("Istanbul: failed to broadcast PREPARE message", "err", err)
		return
	}
}

// handlePrepare is called when receiving a PREPARE message

// It
// - validates PREPARE message digest matches the current block proposal
// - accumulates valid PREPARE message until reaching quorum
// - when quorum is reached update states to "Prepared" and broadcast COMMIT
func (c *Core) handlePrepare(prepare *protocols.Prepare) error {
	logger := c.currentLogger(true, prepare).New()

	logger.Debug("Istanbul: handle PREPARE message", "prepares.count", c.current.Prepares.Size(), "quorum", c.valSet.QuorumSize(), "prepare.digest", prepare.Digest.Hex())

	// Check digest
	if prepare.Digest != c.current.Proposal().Hash() {
		logger.Error("Istanbul: invalid PREPARE message digest", "prepare", prepare.Digest.Hex(), "current", c.current.Proposal().Hash())
		return errInvalidMessage
	}

	// Save PREPARE messages
	if err := c.current.Prepares.Add(prepare); err != nil {
		logger.Error("Istanbul: failed to save PREPARE message", "err", err)
		return err
	}

	logger = logger.New("prepares.count", c.current.Prepares.Size(), "quorum", c.valSet.QuorumSize())

	// Change to "Prepared" state if we've received quorum of PREPARE messages
	// and we are in earlier state than "Prepared"
	if (c.current.Prepares.Size() >= c.valSet.QuorumSize()) && c.state.Cmp(StatePrepared) < 0 {
		logger.Debug("Istanbul: received quorum of PREPARE messages")

		// Accumulates PREPARE messages
		c.current.preparedRound = c.currentView().Round
		c.PreparedPrepares = make([]*protocols.Prepare, 0)
		for _, m := range c.current.Prepares.Values() {
			c.PreparedPrepares = append(
				c.PreparedPrepares,
				protocols.NewPrepareWithSigAndSource(
					m.View().Sequence, m.View().Round, m.(*protocols.Prepare).Digest, m.Signature(), m.Source()))
		}

		if c.current.Proposal() != nil && c.current.Proposal().Hash() == prepare.Digest {
			logger.Debug("Istanbul: PREPARE message matches proposal", "proposal", c.current.Proposal().Hash(), "prepare", prepare.Digest)
			c.current.preparedBlock = c.current.Proposal()
		}

		c.setState(StatePrepared)
		c.broadcastCommit()
	} else {
		logger.Debug("Istanbul: accepted PREPARE messages")
	}

	return nil
}
