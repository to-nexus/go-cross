package core

import (
	"math"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/consensus/istanbul/protocols"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ##CROSS: istanbul message limits

func TestCore_CheckMessage(t *testing.T) {
	valSet, keys := generateValidatorSetAndKeys(t, 4)
	addr := valSet.GetByIndex(0).Address()
	c, _ := newRoundChangeTestCore(t, addr, keys[addr], valSet, makeBlockWithTime(1, 1), 1, 0)

	t.Run("rejects round above uint64", func(t *testing.T) {
		round := new(big.Int).Lsh(big.NewInt(1), 64)
		view := &istanbul.View{Sequence: big.NewInt(1), Round: round}
		assert.Equal(t, errInvalidMessage, c.checkMessage(protocols.RoundChangeCode, view))
		assert.Equal(t, errInvalidMessage, c.checkMessage(protocols.PrepareCode, view))
	})
}

func TestCore_AddToBacklog(t *testing.T) {
	valSet, keys := generateValidatorSetAndKeys(t, 4)
	addr := valSet.GetByIndex(0).Address()
	src := valSet.GetByIndex(1).Address()
	c, _ := newRoundChangeTestCore(t, addr, keys[addr], valSet, makeBlockWithTime(1, 1), 1, 0)

	t.Run("caps messages per source", func(t *testing.T) {
		for i := 0; i < maxBacklogPerSource+10; i++ {
			msg := protocols.NewPrepareWithSigAndSource(big.NewInt(2), big.NewInt(int64(i)), common.Hash{byte(i)}, nil, src)
			c.addToBacklog(msg)
		}
		assert.Equal(t, maxBacklogPerSource, c.backlogs[src].Size())
	})

	t.Run("caps encoded bytes per source", func(t *testing.T) {
		c, _ := newRoundChangeTestCore(t, addr, keys[addr], valSet, makeBlockWithTime(1, 1), 1, 0)
		block := types.NewBlockWithHeader(&types.Header{Number: big.NewInt(2), Difficulty: big.NewInt(0), Extra: make([]byte, 4*1024*1024)})
		for i := 0; i < 10; i++ {
			msg := protocols.NewPreprepare(big.NewInt(2), big.NewInt(int64(i)), block)
			msg.SetSource(src)
			c.addToBacklog(msg)
		}
		// Each PRE-PREPARE is a bit over 4 MiB, so seven of them fit in the budget.
		assert.Equal(t, 7, c.backlogs[src].Size())
		used := c.backlogBytes[src]
		assert.Greater(t, used, 7*4*1024*1024)
		assert.LessOrEqual(t, used, maxBacklogBytesPerSource)

		// Messages that are still in the future keep their bytes.
		c.processBacklog()
		assert.Equal(t, 7, c.backlogs[src].Size())
		assert.Equal(t, used, c.backlogBytes[src])

		// Old messages leave the backlog and release their bytes.
		c.current = newRoundState(&istanbul.View{Sequence: big.NewInt(3), Round: big.NewInt(0)}, valSet, nil, nil, nil, nil, func(common.Hash) bool { return false })
		c.processBacklog()
		assert.Zero(t, c.backlogs[src].Size())
		assert.Zero(t, c.backlogBytes[src])
	})
}

func TestCore_VerifySignatures(t *testing.T) {
	valSet, keys := generateValidatorSetAndKeys(t, 4)
	addr := valSet.GetByIndex(0).Address()
	src := valSet.GetByIndex(1).Address()
	c, _ := newRoundChangeTestCore(t, addr, keys[addr], valSet, makeBlockWithTime(1, 1), 1, 0)

	t.Run("rejects justification longer than validator set", func(t *testing.T) {
		block := makeBlockWithTime(1, 1)
		roundChange := createSignedRoundChangeMessage(t, keys[src], src, 1, 0, block)
		prepare := createSignedPrepareMessage(t, keys[src], src, 0, block)
		for i := 0; i <= valSet.Size(); i++ {
			roundChange.Justification = append(roundChange.Justification, prepare)
		}
		payload, err := rlp.EncodeToBytes(roundChange)
		require.NoError(t, err)
		assert.Equal(t, errInvalidMessage, c.handleEncodedMsg(protocols.RoundChangeCode, payload))
	})

	// Justification was rejected before any signature recovery and returns errInvalidMessage.
	block := makeBlockWithTime(1, 1)
	unsigned := protocols.NewPrepareWithSigAndSource(big.NewInt(1), big.NewInt(0), block.Hash(), nil, src)

	t.Run("accepts prepares with a prepared block", func(t *testing.T) {
		roundChange := createSignedRoundChangeMessage(t, keys[src], src, 1, 0, block)
		roundChange.Justification = []*protocols.Prepare{createSignedPrepareMessage(t, keys[src], src, 0, block)}
		assert.NoError(t, c.verifySignatures(roundChange))
	})

	t.Run("rejects prepares without a prepared block before signature recovery", func(t *testing.T) {
		roundChange := createSignedRoundChangeMessage(t, keys[src], src, 1, 0, nil)
		roundChange.Justification = []*protocols.Prepare{unsigned}
		assert.Equal(t, errInvalidMessage, c.verifySignatures(roundChange))
	})

	t.Run("rejects round 0 PRE-PREPARE justification before signature recovery", func(t *testing.T) {
		preprepare := protocols.NewPreprepare(big.NewInt(1), big.NewInt(0), block)
		signProtocolMessage(t, preprepare, keys[src])

		preprepare.JustificationPrepares = []*protocols.Prepare{unsigned}
		assert.Equal(t, errInvalidMessage, c.verifySignatures(preprepare))

		preprepare.JustificationPrepares = nil
		preprepare.JustificationRoundChanges = []*protocols.SignedRoundChangePayload{&protocols.NewRoundChange(big.NewInt(1), big.NewInt(1), nil, nil).SignedRoundChangePayload}
		assert.Equal(t, errInvalidMessage, c.verifySignatures(preprepare))
	})
}

func TestCore_HandleDecodedMessage(t *testing.T) {
	valSet, keys := generateValidatorSetAndKeys(t, 4)
	addr := valSet.GetByIndex(0).Address()
	src := valSet.GetByIndex(1).Address()
	c, _ := newRoundChangeTestCore(t, addr, keys[addr], valSet, makeBlockWithTime(1, 1), 1, 0)
	farFuture := protocols.NewPrepareWithSigAndSource(big.NewInt(3), big.NewInt(0), common.Hash{}, nil, src)

	t.Run("logs dropped far future messages once per interval", func(t *testing.T) {
		// The first drop is logged right away and resets the counts.
		require.Equal(t, errFarFutureMessage, c.handleDecodedMessage(farFuture))
		logged := c.farFutureLogged
		require.False(t, logged.IsZero())
		assert.Empty(t, c.farFutureDrops)

		// Later drops within the interval are only counted.
		for i := 0; i < 3; i++ {
			require.Equal(t, errFarFutureMessage, c.handleDecodedMessage(farFuture))
		}
		assert.Equal(t, logged, c.farFutureLogged)
		assert.Equal(t, uint64(3), c.farFutureDrops[src])

		// After the interval, the next drop logs the counts again.
		c.farFutureLogged = time.Now().Add(-farFutureLogInterval)
		require.Equal(t, errFarFutureMessage, c.handleDecodedMessage(farFuture))
		assert.Empty(t, c.farFutureDrops)
	})
}

// ##CROSS: istanbul far future message
func TestCore_HandleEncodedMsg(t *testing.T) {
	valSet, keys := generateValidatorSetAndKeys(t, 4)
	addr := valSet.GetByIndex(0).Address()
	src := valSet.GetByIndex(1).Address()
	c, b := newRoundChangeTestCore(t, addr, keys[addr], valSet, makeBlockWithTime(1, 1), 1, 0)

	t.Run("forgets a far future message", func(t *testing.T) {
		msg := protocols.NewPrepareWithSigAndSource(big.NewInt(3), big.NewInt(0), common.Hash{}, nil, src)
		signProtocolMessage(t, msg, keys[src])
		payload, err := rlp.EncodeToBytes(msg)
		require.NoError(t, err)

		require.Equal(t, errFarFutureMessage, c.handleEncodedMsg(protocols.PrepareCode, payload))
		assert.Equal(t, []common.Hash{istanbul.RLPHash(payload)}, b.forgotten)
	})
}

// ##

func TestCore_HandleRoundChange(t *testing.T) {
	valSet, keys := generateValidatorSetAndKeys(t, 4)
	addr := valSet.GetByIndex(0).Address()
	attacker := valSet.GetByIndex(1).Address()
	honest := valSet.GetByIndex(2).Address()
	const currentRound = 10

	send := func(t *testing.T, c *Core, from common.Address, round int64) {
		t.Helper()
		payload, err := rlp.EncodeToBytes(createSignedRoundChangeMessage(t, keys[from], from, round, 0, nil))
		require.NoError(t, err)
		require.NoError(t, c.handleEncodedMsg(protocols.RoundChangeCode, payload))
	}
	futureRounds := func(c *Core) []uint64 {
		var rounds []uint64
		for k := range c.roundChangeSet.roundChanges {
			if k > currentRound {
				rounds = append(rounds, k)
			}
		}
		return rounds
	}

	t.Run("keeps one future round per validator", func(t *testing.T) {
		c, _ := newRoundChangeTestCore(t, addr, keys[addr], valSet, makeBlockWithTime(1, 1), 1, currentRound)
		for round := int64(currentRound + 1); round <= currentRound+100; round++ {
			send(t, c, attacker, round)
		}
		assert.Equal(t, []uint64{currentRound + 100}, futureRounds(c))

		// A lower future round from the same validator is ignored.
		send(t, c, attacker, currentRound+50)
		assert.Equal(t, []uint64{currentRound + 100}, futureRounds(c))
	})

	t.Run("F+1 validators still move to the lowest future round", func(t *testing.T) {
		c, _ := newRoundChangeTestCore(t, addr, keys[addr], valSet, makeBlockWithTime(1, 1), 1, currentRound)
		send(t, c, attacker, currentRound+100)
		send(t, c, honest, currentRound+2)
		assert.Equal(t, uint64(currentRound+2), c.currentView().Round.Uint64())
	})

	t.Run("rebuilds the prepared cache after dropping a lower future round", func(t *testing.T) {
		// With 7 validators F+1 is 3, so two validators do not move the node to a future round.
		valSet, keys := generateValidatorSetAndKeys(t, 7)
		addr, a, b := valSet.GetByIndex(0).Address(), valSet.GetByIndex(1).Address(), valSet.GetByIndex(2).Address()
		block := makeBlockWithTime(1, 1)
		c, _ := newRoundChangeTestCore(t, addr, keys[addr], valSet, block, 1, currentRound)

		send := func(from common.Address, round, preparedRound int64) {
			msg := createSignedRoundChangeMessage(t, keys[from], from, round, preparedRound, block)
			for _, v := range valSet.List()[:valSet.QuorumSize()] {
				msg.Justification = append(msg.Justification, createSignedPrepareMessage(t, keys[v.Address()], v.Address(), preparedRound, block))
			}
			payload, err := rlp.EncodeToBytes(msg)
			require.NoError(t, err)
			require.NoError(t, c.handleEncodedMsg(protocols.RoundChangeCode, payload))
		}
		send(a, currentRound+1, 5)
		send(b, currentRound+1, 4)
		send(a, currentRound+2, 5) // drops a from round currentRound+1

		round := uint64(currentRound + 1)
		rcs := c.roundChangeSet
		require.Equal(t, uint64(currentRound), c.currentView().Round.Uint64())
		require.Nil(t, rcs.roundChanges[round].Get(a))
		assert.Equal(t, big.NewInt(4), rcs.highestPreparedRound[round], "the prepared round must come from b")
		require.NotEmpty(t, rcs.prepareMessages[round])
		assert.Equal(t, big.NewInt(4), rcs.prepareMessages[round][0].Round)
	})
}

func TestRoundChangeSet_GetMinRoundChange(t *testing.T) {
	valSet, _ := generateValidatorSetAndKeys(t, 4)

	t.Run("handles rounds above max int64", func(t *testing.T) {
		rcs := newRoundChangeSet(valSet)
		rcs.NewRound(new(big.Int).SetUint64(math.MaxUint64))
		rcs.NewRound(big.NewInt(5))
		assert.Equal(t, big.NewInt(5), rcs.getMinRoundChange(big.NewInt(1)))
	})
}
