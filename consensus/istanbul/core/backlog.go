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
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/prque"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/consensus/istanbul/protocols"
)

var (
	// msgPriority is defined for calculating processing priority to speedup consensus
	// msgPreprepare > msgCommit > msgPrepare
	msgPriority = map[uint64]int{
		protocols.PreprepareCode: 1,
		protocols.CommitCode:     2,
		protocols.PrepareCode:    3,
	}
)

// checkMessage checks that a message matches our current QBFT state
//
// In particular it ensures that
// - message has the expected round
// - message has the expected sequence
// - message type is expected given our current state

// return errInvalidMessage if the message is invalid
// return errFutureMessage if the message view is 1 round/sequence larger than current view
// return errFarFutureMessage if the message view is more than 1+ round/sequence larger than current view
// return errOldMessage if the message view is smaller than current view
func (c *Core) checkMessage(msgCode uint64, view *istanbul.View) error {
	if view == nil || view.Sequence == nil || view.Round == nil {
		return errInvalidMessage
	}

	if msgCode == protocols.RoundChangeCode {
		// if ROUND-CHANGE message
		// check that
		// - sequence matches our current sequence
		// - round is in the future
		if view.Sequence.Cmp(c.currentView().Sequence) > 0 {
			if view.Sequence.Cmp(new(big.Int).Add(c.currentView().Sequence, common.Big1)) > 0 {
				return errFarFutureMessage
			}
			return errFutureMessage
		} else if view.Cmp(c.currentView()) < 0 {
			return errOldMessage
		}
		return nil
	}

	// If not ROUND-CHANGE
	// check that round and sequence equals our current round and sequence
	if cmp := view.Cmp(c.currentView()); cmp > 0 {
		if cmp > 1 {
			return errFarFutureMessage
		}
		return errFutureMessage
	}

	if view.Cmp(c.currentView()) < 0 {
		return errOldMessage
	}

	switch c.state {
	case StateAcceptRequest:
		// StateAcceptRequest only accepts msgPreprepare and msgRoundChange
		// other messages are future messages
		if msgCode > protocols.PreprepareCode {
			return errFutureMessage
		}
		return nil
	case StatePreprepared:
		// StatePreprepared only accepts msgPrepare and msgRoundChange
		// message less than msgPrepare are invalid and greater are future messages
		if msgCode < protocols.PrepareCode {
			return errInvalidMessage
		} else if msgCode > protocols.PrepareCode {
			return errFutureMessage
		}
		return nil
	case StatePrepared:
		// StatePrepared only accepts msgCommit and msgRoundChange
		// other messages are invalid messages
		if msgCode < protocols.CommitCode {
			return errInvalidMessage
		}
		return nil
	case StateCommitted:
		// StateCommit rejects all messages other than msgRoundChange
		return errInvalidMessage
	}
	return nil
}

// addToBacklog allows to postpone the processing of future messages

// it adds the message to backlog which is read on every state change
func (c *Core) addToBacklog(msg protocols.Message) {
	logger := c.currentLogger(true, msg)

	src := msg.Source()
	if src == c.Address() {
		logger.Warn("Istanbul: backlog from self")
		return
	}

	logger.Trace("Istanbul: new backlog message", "backlogs_size", len(c.backlogs))

	c.backlogsMu.Lock()
	defer c.backlogsMu.Unlock()

	backlog := c.backlogs[src]
	if backlog == nil {
		backlog = prque.New[int64, protocols.Message](nil)
		c.backlogs[src] = backlog
	}
	view := msg.View()
	backlog.Push(msg, toPriority(msg.Code(), &view))
}

// processBacklog lookup for future messages that have been backlogged and post it on
// the event channel so main handler loop can handle it

// It is called on every state change
func (c *Core) processBacklog() {
	c.backlogsMu.Lock()
	defer c.backlogsMu.Unlock()

	for srcAddress, backlog := range c.backlogs {
		if backlog == nil {
			continue
		}
		_, src := c.valSet.GetByAddress(srcAddress)
		if src == nil {
			// validator is not available
			delete(c.backlogs, srcAddress)
			continue
		}
		logger := c.logger.New("from", src, "state", c.state)
		isFuture := false

		logger.Trace("Istanbul: process backlog")

		// We stop processing if
		//   1. backlog is empty
		//   2. The first message in queue is a future message
		for !(backlog.Empty() || isFuture) {
			m, prio := backlog.Pop()

			var code uint64
			var view istanbul.View
			var event backlogEvent

			msg := m.(protocols.Message)
			code = msg.Code()
			view = msg.View()
			event.msg = msg

			// Push back if it's a future message
			err := c.checkMessage(code, &view)
			if err != nil {
				if err == errFutureMessage {
					// this is still a future message
					logger.Trace("Istanbul: stop processing backlog", "msg", m)
					backlog.Push(m, prio)
					isFuture = true
					break
				}
				logger.Trace("Istanbul: skip backlog message", "msg", m, "err", err)
				continue
			}
			logger.Trace("Istanbul: post backlog event", "msg", m)

			event.src = src
			go c.sendEvent(event)
		}
	}
}

func toPriority(msgCode uint64, view *istanbul.View) int64 {
	if msgCode == protocols.RoundChangeCode {
		// For msgRoundChange, set the message priority based on its sequence
		return -int64(view.Sequence.Uint64() * 1000)
	}
	// FIXME: round will be reset as 0 while new sequence
	// 10 * Round limits the range of message code is from 0 to 9
	// 1000 * Sequence limits the range of round is from 0 to 99
	return -int64(view.Sequence.Uint64()*1000 + view.Round.Uint64()*10 + uint64(msgPriority[msgCode]))
}
