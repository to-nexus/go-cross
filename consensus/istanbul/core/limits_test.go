package core

import (
	"math"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/consensus/istanbul/protocols"
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
