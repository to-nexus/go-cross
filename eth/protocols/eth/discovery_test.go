// Copyright 2026 The go-ethereum Authors
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

// ##CROSS: discovery filter

import (
	"testing"

	"github.com/ethereum/go-ethereum/core/forkid"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/p2p/enr"
	"github.com/stretchr/testify/require"
)

// ##CROSS: discovery filter
func TestDiscovery_NewTableFilter(t *testing.T) {
	backend := newTestBackend(0)
	defer backend.close()
	filter := NewTableFilter(backend.chain)

	newNode := func(entries ...enr.Entry) *enode.Node {
		var r enr.Record
		for _, e := range entries {
			r.Set(e)
		}
		return enode.SignNull(&r, enode.ID{1})
	}

	t.Run("keeps node without eth entry", func(t *testing.T) {
		require.True(t, filter(newNode()))
	})
	t.Run("keeps node with compatible forkid", func(t *testing.T) {
		require.True(t, filter(newNode(currentENREntry(backend.chain))))
	})
	t.Run("rejects node with incompatible forkid", func(t *testing.T) {
		other := &enrEntry{ForkID: forkid.ID{Hash: [4]byte{0x22, 0xd5, 0x23, 0xb2}}}
		require.False(t, filter(newNode(other)))
	})
	t.Run("rejects node with broken eth entry", func(t *testing.T) {
		require.False(t, filter(newNode(enr.WithEntry("eth", "broken"))))
	})
}

// ##
