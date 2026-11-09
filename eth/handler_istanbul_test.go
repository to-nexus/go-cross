// Copyright 2026 The go-cross Authors
// This file is part of the go-cross library.
//
// The go-cross library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

package eth

import (
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/eth/protocols/eth"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/enode"
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

// ##CROSS: fix upstream
// TestConsensusProtocolNoPeerLeak checks that an istanbul session does not start eth peer goroutines.
func TestConsensusProtocolNoPeerLeak(t *testing.T) {
	h := newTestHandler()
	defer h.close()

	p := p2p.NewPeer(enode.ID{1}, "", nil)
	ethApp, ethNet := p2p.MsgPipe()
	defer ethApp.Close()
	defer ethNet.Close()
	ethPeer := eth.NewPeer(eth.ETH68, p, ethApp, h.txpool)
	defer ethPeer.Close()
	require.NoError(t, h.handler.peers.registerPeer(ethPeer, nil))
	p.EthPeerRegistered <- struct{}{}

	proto := (*istanbulHandler)(h.handler).makeIstanbulConsensusProtocol(istanbulProtocolName, IstanbulProtocolVersion, istanbulProtocolLengths, (*ethHandler)(h.handler), 1)
	app, net := p2p.MsgPipe()
	defer app.Close()

	// Count one of the loops eth.NewPeer starts; the whole goroutine count is noisy.
	// Only ethPeer's own loop should be running, so wait until exactly one is left.
	peerLoops := func() (n int) {
		for range 100 {
			buf := make([]byte, 1<<20)
			if n = strings.Count(string(buf[:runtime.Stack(buf, true)]), "(*Peer).broadcastTransactions("); n == 1 {
				break
			}
			time.Sleep(10 * time.Millisecond)
		}
		return n
	}
	require.Equal(t, 1, peerLoops())
	errc := make(chan error, 1)
	go func() { errc <- proto.Run(p, app) }()
	time.Sleep(50 * time.Millisecond) // let Run reach the read loop
	net.Close()
	require.ErrorIs(t, <-errc, p2p.ErrPipeClosed)
	require.Equal(t, 1, peerLoops(), "istanbul session left eth peer goroutines running")
}

// ##
