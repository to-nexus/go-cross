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
	"math"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/prque"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/consensus/istanbul/protocols"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"
)

var (
	roundMeter     = metrics.NewRegisteredMeter("consensus/istanbul/core/round", nil)
	sequenceMeter  = metrics.NewRegisteredMeter("consensus/istanbul/core/sequence", nil)
	consensusTimer = metrics.NewRegisteredTimer("consensus/istanbul/core/consensus", nil)
)

// New creates an Istanbul consensus core
func New(backend istanbul.Backend, config *istanbul.Config) *Core {
	c := &Core{
		config:             config,
		address:            backend.Address(),
		state:              StateAcceptRequest,
		handlerWg:          new(sync.WaitGroup),
		logger:             log.New("address", backend.Address()),
		backend:            backend,
		backlogs:           make(map[common.Address]*prque.Prque[int64, protocols.Message]),
		backlogsMu:         new(sync.Mutex),
		pendingRequests:    prque.New[int64, *Request](nil),
		pendingRequestsMu:  new(sync.Mutex),
		consensusTimestamp: time.Time{},
	}

	c.validateFn = c.checkValidatorSignature
	return c
}

// ----------------------------------------------------------------------------

type Core struct {
	config  *istanbul.Config
	address common.Address
	state   State
	logger  log.Logger

	backend               istanbul.Backend
	events                *event.TypeMuxSubscription
	finalCommittedSub     *event.TypeMuxSubscription
	timeoutSub            *event.TypeMuxSubscription
	futurePreprepareTimer *time.Timer

	valSet     istanbul.ValidatorSet
	validateFn func([]byte, []byte) (common.Address, error)

	backlogs   map[common.Address]*prque.Prque[int64, protocols.Message]
	backlogsMu *sync.Mutex

	current      *roundState
	currentMutex sync.Mutex
	handlerWg    *sync.WaitGroup

	roundChangeSet   *roundChangeSet
	roundChangeTimer *time.Timer

	PreparedPrepares []*protocols.Prepare

	pendingRequests   *prque.Prque[int64, *Request]
	pendingRequestsMu *sync.Mutex

	consensusTimestamp time.Time

	newRoundMutex sync.Mutex
	newRoundTimer *time.Timer
}

// ##CROSS: istanbul stats

// CurrentStat returns the current validator list and view for stat logging.
func (c *Core) CurrentStat() (istanbul.ValidatorSet, *istanbul.View) {
	c.currentMutex.Lock()
	defer c.currentMutex.Unlock()
	return c.valSet, c.currentView()
}

// ##

func (c *Core) currentView() *istanbul.View {
	return &istanbul.View{
		Sequence: new(big.Int).Set(c.current.Sequence()),
		Round:    new(big.Int).Set(c.current.Round()),
	}
}

func (c *Core) IsProposer() bool {
	v := c.valSet
	if v == nil {
		return false
	}
	return v.IsProposer(c.backend.Address())
}

func (c *Core) IsCurrentProposal(blockHash common.Hash) bool {
	return c.current != nil && c.current.pendingRequest != nil && c.current.pendingRequest.Proposal.Hash() == blockHash
}

// startNewRound starts a new round. if round equals to 0, it means to starts a new sequence
func (c *Core) startNewRound(round *big.Int) {
	c.currentMutex.Lock()
	defer c.currentMutex.Unlock()

	var logger log.Logger
	if c.current == nil {
		logger = c.logger.New("old.round", -1, "old.seq", 0)
	} else {
		logger = c.currentLogger(false, nil)
	}
	logger = logger.New("target.round", round)

	roundChange := false

	// Try to get last proposal
	lastProposal, lastProposer := c.backend.LastProposal()
	if lastProposal != nil {
		logger = logger.New("lastProposal.number", lastProposal.Number().Uint64(), "lastProposal.hash", lastProposal.Hash())
	}

	logger.Debug("Istanbul: initialize new round")

	if c.current == nil {
		logger.Debug("Istanbul: start at the initial round")
	} else if lastProposal.Number().Cmp(c.current.Sequence()) >= 0 {
		diff := new(big.Int).Sub(lastProposal.Number(), c.current.Sequence())
		sequenceMeter.Mark(new(big.Int).Add(diff, common.Big1).Int64())

		if !c.consensusTimestamp.IsZero() {
			consensusTimer.UpdateSince(c.consensusTimestamp)
			c.consensusTimestamp = time.Time{}
		}
		logger.Debug("Istanbul: catch up last block proposal")
	} else if lastProposal.Number().Cmp(big.NewInt(c.current.Sequence().Int64()-1)) == 0 {
		if round.Cmp(common.Big0) == 0 {
			// same seq and round, don't need to start new round
			logger.Debug("Istanbul: same round, no need to start new round")
			return
		} else if round.Cmp(c.current.Round()) < 0 {
			logger.Warn("Istanbul: next round is inferior to current round")
			return
		}
		roundChange = true
	} else {
		logger.Warn("Istanbul: next sequence is before last block proposal")
		return
	}

	var oldLogger log.Logger
	if c.current == nil {
		oldLogger = c.logger.New("old.round", -1, "old.seq", 0)
	} else {
		oldLogger = c.logger.New("old.round", c.current.Round().Uint64(), "old.sequence", c.current.Sequence().Uint64(), "old.state", c.state.String(), "old.proposer", c.valSet.GetProposer())
	}

	// Create next view
	var newView *istanbul.View
	if roundChange {
		newView = &istanbul.View{
			Sequence: new(big.Int).Set(c.current.Sequence()),
			Round:    new(big.Int).Set(round),
		}
	} else {
		newView = &istanbul.View{
			Sequence: new(big.Int).Add(lastProposal.Number(), common.Big1),
			Round:    new(big.Int),
		}
		c.valSet = c.backend.Validators(lastProposal)
	}

	// New snapshot for new round
	c.updateRoundState(newView, c.valSet, roundChange)

	// Calculate new proposer
	c.valSet.CalcProposer(lastProposer, newView.Round.Uint64())
	c.setState(StateAcceptRequest)

	if c.current != nil && round.Cmp(c.current.Round()) > 0 {
		roundMeter.Mark(new(big.Int).Sub(round, c.current.Round()).Int64())
	}

	// Update RoundChangeSet by deleting older round messages
	if round.Uint64() == 0 {
		c.PreparedPrepares = nil
		c.roundChangeSet = newRoundChangeSet(c.valSet)
	} else {
		// Clear earlier round messages
		c.roundChangeSet.ClearLowerThan(round)
	}
	c.roundChangeSet.NewRound(round)

	// ##CROSS: istanbul stats
	go c.sendEvent(istanbul.NewRoundEvent{
		Validators: c.valSet,
		Sequence:   newView.Sequence.Uint64(),
		Round:      newView.Round.Uint64(),
	})
	// ##

	if round.Uint64() > 0 {
		c.newRoundChangeTimer()
	}

	oldLogger.Info("Istanbul: start new round",
		"next.round", newView.Round,
		"next.seq", newView.Sequence,
		"next.proposer", c.valSet.GetProposer(),
		"next.valSet", c.valSet.List(),
		"next.size", c.valSet.Size(),
		"next.IsProposer", c.IsProposer(),
	)
}

// updateRoundState updates round state by checking if locking block is necessary
func (c *Core) updateRoundState(view *istanbul.View, validatorSet istanbul.ValidatorSet, roundChange bool) {
	if roundChange && c.current != nil {
		c.current = newRoundState(view, validatorSet, c.current.Preprepare, c.current.preparedRound, c.current.preparedBlock, c.current.pendingRequest, c.backend.HasBadProposal)
	} else {
		c.current = newRoundState(view, validatorSet, nil, nil, nil, nil, c.backend.HasBadProposal)
	}
}

func (c *Core) setState(state State) {
	if c.state != state {
		oldState := c.state
		c.state = state
		c.currentLogger(false, nil).Debug("Istanbul: changed state", "old.state", oldState.String(), "new.state", state.String())
	}
	if state == StateAcceptRequest {
		c.processPendingRequests()
	}

	// each time we change state, we process backlog for possible message that are
	// now ready
	c.processBacklog()
}

func (c *Core) Address() common.Address {
	return c.address
}

func (c *Core) stopFuturePreprepareTimer() {
	if c.futurePreprepareTimer != nil {
		c.futurePreprepareTimer.Stop()
	}
}

func (c *Core) stopTimer() {
	c.stopFuturePreprepareTimer()
	if c.roundChangeTimer != nil {
		c.roundChangeTimer.Stop()
	}
}

func (c *Core) newRoundChangeTimer() {
	c.stopTimer()

	for c.current == nil { // wait because it is asynchronous in handleRequest
		time.Sleep(10 * time.Millisecond)
	}

	// set timeout based on the round number
	baseTimeout := time.Duration(c.config.GetConfig(c.current.Sequence()).RequestTimeout) * time.Millisecond
	round := c.current.Round().Uint64()
	maxRequestTimeout := time.Duration(c.config.GetConfig(c.current.Sequence()).MaxRequestTimeoutSeconds) * time.Second

	// If the upper limit of the request timeout is capped by small maxRequestTimeout, round can be a quite large number,
	// which leads to float64 overflow, making its value negative or zero forever after some point.
	// In this case we cannot simply use math.Pow and have to implement a safeguard on our own, at the cost of performance (which is not important in this case).
	var timeout time.Duration
	if maxRequestTimeout > time.Duration(0) {
		timeout = baseTimeout
		for i := uint64(0); i < round; i++ {
			timeout = timeout * 2
			if timeout > maxRequestTimeout {
				timeout = maxRequestTimeout
				break
			}
		}
		// prevent log storm when unexpected overflow happens
		if timeout < baseTimeout {
			c.currentLogger(true, nil).Error("Istanbul: Possible request timeout overflow detected, setting timeout value to maxRequestTimeout",
				"timeout", timeout.Seconds(),
				"max_request_timeout", maxRequestTimeout.Seconds(),
			)
			timeout = maxRequestTimeout
		}
	} else {
		// effectively impossible to observe overflow happen when maxRequestTimeout is disabled
		timeout = baseTimeout * time.Duration(math.Pow(2, float64(round)))
	}

	c.currentLogger(true, nil).Trace("Istanbul: start new ROUND-CHANGE timer", "timeout", timeout.Seconds())
	c.roundChangeTimer = time.AfterFunc(timeout, func() {
		c.sendEvent(timeoutEvent{})
	})
}

func (c *Core) checkValidatorSignature(data []byte, sig []byte) (common.Address, error) {
	return istanbul.CheckValidatorSignature(c.valSet, data, sig)
}

// PrepareCommittedSeal returns a committed seal for the given header and takes current round under consideration
func PrepareCommittedSeal(header *types.Header, round uint32) []byte {
	h := types.CopyHeader(header)
	return h.IstanbulHashWithRoundNumber(round).Bytes()
}
