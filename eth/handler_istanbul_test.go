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
			require.True(t, limiter.allow(now, 1))
		}
		require.False(t, limiter.allow(now, 1))
		require.True(t, limiter.allow(now.Add(time.Second), 1))
	})

	t.Run("byte burst", func(t *testing.T) {
		limiter := newConsensusLimiter()
		require.True(t, limiter.allow(now, protocolMaxMsgSize))
		require.True(t, limiter.allow(now, protocolMaxMsgSize))
		require.False(t, limiter.allow(now, 1))
	})

	t.Run("peer limits are independent", func(t *testing.T) {
		first := newConsensusLimiter()
		second := newConsensusLimiter()
		for i := 0; i < consensusMessageBurst; i++ {
			require.True(t, first.allow(now, 1))
		}
		require.False(t, first.allow(now, 1))
		require.True(t, second.allow(now, 1))
	})
}
