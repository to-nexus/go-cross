// Copyright 2024 The go-ethereum Authors
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

package eth

// ##CROSS: peer permission
import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// eachPeerBounded is the bounded-concurrency fan-out used by sweepPermissionedPeers.
// These tests exercise that orchestration directly (the per-peer decision, PeerAllowed,
// is covered in the eth/protocols/eth package). The peer values are only counted, never
// dereferenced, so empty ethPeer values are sufficient.

// TestEachPeerBoundedRunsAllOnce verifies every peer is processed exactly once.
func TestEachPeerBoundedRunsAllOnce(t *testing.T) {
	const n = 50
	peers := make([]*ethPeer, n)
	idx := make(map[*ethPeer]int, n)
	for i := range peers {
		peers[i] = &ethPeer{}
		idx[peers[i]] = i
	}

	var total int64
	seen := make([]int32, n)
	eachPeerBounded(peers, 8, func(p *ethPeer) {
		atomic.AddInt64(&total, 1)
		atomic.AddInt32(&seen[idx[p]], 1) // idx is read-only here, safe for concurrent reads
	})

	if total != n {
		t.Fatalf("fn called %d times, want %d", total, n)
	}
	for i, c := range seen {
		if c != 1 {
			t.Errorf("peer %d processed %d times, want 1", i, c)
		}
	}
}

// TestEachPeerBoundedConcurrencyLimit verifies concurrent invocations never exceed the
// limit, while still running more than one at a time.
func TestEachPeerBoundedConcurrencyLimit(t *testing.T) {
	const (
		n     = 40
		limit = 5
	)
	peers := make([]*ethPeer, n)
	for i := range peers {
		peers[i] = &ethPeer{}
	}

	var (
		cur int64
		mu  sync.Mutex
		max int64
	)
	eachPeerBounded(peers, limit, func(p *ethPeer) {
		c := atomic.AddInt64(&cur, 1)
		mu.Lock()
		if c > max {
			max = c
		}
		mu.Unlock()
		time.Sleep(2 * time.Millisecond) // hold the slot to force overlap
		atomic.AddInt64(&cur, -1)
	})

	if max > limit {
		t.Errorf("observed concurrency %d exceeded limit %d", max, limit)
	}
	if max < 2 {
		t.Errorf("expected concurrent execution, observed max %d", max)
	}
}

// TestEachPeerBoundedEmpty verifies an empty peer list returns promptly without calling fn.
func TestEachPeerBoundedEmpty(t *testing.T) {
	done := make(chan struct{})
	go func() {
		eachPeerBounded(nil, 8, func(*ethPeer) {
			t.Error("fn should not be called for an empty peer list")
		})
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("eachPeerBounded did not return for an empty peer list")
	}
}
