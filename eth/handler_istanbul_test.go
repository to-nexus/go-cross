// Copyright 2026 The go-cross Authors
// This file is part of the go-cross library.
//
// The go-cross library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

package eth

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestConsensusLimiter(t *testing.T) {
	now := time.Now()

	t.Run("message burst", func(t *testing.T) {
		limiter := newConsensusLimiter()
		for i := 0; i < consensusMessageBurst; i++ {
			delay, ok := limiter.reserve(now, 1)
			require.True(t, ok)
			require.Zero(t, delay)
		}
		delay, ok := limiter.reserve(now, 1)
		require.True(t, ok)
		require.Equal(t, 10*time.Millisecond, delay)

		delay, ok = limiter.reserve(now.Add(time.Second), 1)
		require.True(t, ok)
		require.Zero(t, delay)
	})

	t.Run("byte burst", func(t *testing.T) {
		limiter := newConsensusLimiter()
		for i := 0; i < 2; i++ {
			delay, ok := limiter.reserve(now, protocolMaxMsgSize)
			require.True(t, ok)
			require.Zero(t, delay)
		}
		delay, ok := limiter.reserve(now, protocolMaxMsgSize)
		require.True(t, ok)
		require.Equal(t, 625*time.Millisecond, delay)
	})

	t.Run("failed reservation restores tokens", func(t *testing.T) {
		limiter := newConsensusLimiter()
		_, ok := limiter.reserve(now, uint32(consensusByteBurst+1))
		require.False(t, ok)

		for i := 0; i < consensusMessageBurst; i++ {
			delay, ok := limiter.reserve(now, 1)
			require.True(t, ok)
			require.Zero(t, delay)
		}
	})

	t.Run("peer limits are independent", func(t *testing.T) {
		first := newConsensusLimiter()
		second := newConsensusLimiter()
		for i := 0; i < consensusMessageBurst; i++ {
			delay, ok := first.reserve(now, 1)
			require.True(t, ok)
			require.Zero(t, delay)
		}
		firstDelay, ok := first.reserve(now, 1)
		require.True(t, ok)
		require.Equal(t, 10*time.Millisecond, firstDelay)

		secondDelay, ok := second.reserve(now, 1)
		require.True(t, ok)
		require.Zero(t, secondDelay)
	})
}
