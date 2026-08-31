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

import (
	"errors"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/params"
)

// Node IDs for testing. The distinguishing byte is placed in the last 20 bytes so each ID
// derives a distinct nodekey address (idAddr).
var (
	idSelfValidator = enode.ID{31: 0x11} // self (validator)
	idPeerValidator = enode.ID{31: 0x22} // peer validator
	idNonValidator  = enode.ID{31: 0x99} // not a validator
)

// idAddr derives the nodekey address (last 20 bytes) from a node ID, mirroring how
// PermissionPeers converts an ID before consulting the ValidatorChecker.
func idAddr(id enode.ID) common.Address { return common.BytesToAddress(id[12:]) }

// mapValidatorChecker is a ValidatorChecker implementation for testing (address-set based).
type mapValidatorChecker map[common.Address]struct{}

func newMapChecker(addrs ...common.Address) mapValidatorChecker {
	m := make(mapValidatorChecker, len(addrs))
	for _, a := range addrs {
		m[a] = struct{}{}
	}
	return m
}

func (m mapValidatorChecker) IsValidator(addr common.Address) (bool, error) {
	_, ok := m[addr]
	return ok, nil
}

// TestVerifyPermission verifies the permission decision rules in a table-driven way.
func TestVerifyPermission(t *testing.T) {
	// Register self and the peer validator as validators (by their derived addresses).
	checker := newMapChecker(idAddr(idSelfValidator), idAddr(idPeerValidator))

	// Construct directly to verify only the decision logic.
	newPP := func(self enode.ID) *PermissionPeers {
		return &PermissionPeers{self: self, validators: checker}
	}

	tests := []struct {
		name    string
		self    enode.ID
		id      enode.ID
		trusted bool
		static  bool
		wantErr bool
	}{
		{name: "allow anyone if self is not a validator", self: idNonValidator, id: idNonValidator, wantErr: false},
		{name: "validator <-> validator is allowed", self: idSelfValidator, id: idPeerValidator, wantErr: false}, // core case
		{name: "validator rejects a non-validator", self: idSelfValidator, id: idNonValidator, wantErr: true},
		{name: "validator allows a non-validator if trusted", self: idSelfValidator, id: idNonValidator, trusted: true, wantErr: false},
		{name: "validator allows a non-validator if static", self: idSelfValidator, id: idNonValidator, static: true, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pp := newPP(tt.self)
			err := pp.verifyPermission(tt.id, tt.trusted, tt.static)
			if (err != nil) != tt.wantErr {
				t.Fatalf("verifyPermission() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestPermissionPeersDecisions verifies the decisions end-to-end through a PermissionPeers
// built by NewPermissionPeers: validator <-> validator allowed, others rejected, trusted
// exempted. Verification is now a direct synchronous call (no loop goroutine).
func TestPermissionPeersDecisions(t *testing.T) {
	checker := newMapChecker(idAddr(idSelfValidator), idAddr(idPeerValidator))

	// A node that is itself a validator.
	pp := NewPermissionPeers(idSelfValidator, checker, nil)
	defer pp.Close()

	// validator peer → allowed
	if !pp.PeerAllowed(idPeerValidator, false, false) {
		t.Errorf("validator peer should be permitted")
	}
	// non-validator peer → rejected
	if pp.PeerAllowed(idNonValidator, false, false) {
		t.Errorf("non-validator peer should be rejected")
	}
	// non-validator but trusted → allowed
	if !pp.PeerAllowed(idNonValidator, true, false) {
		t.Errorf("trusted peer should be permitted")
	}
}

// TestPermissionPeersOpenWhenNotValidator verifies that when this node is not a
// validator, all peers are allowed (open).
func TestPermissionPeersOpenWhenNotValidator(t *testing.T) {
	checker := newMapChecker(idAddr(idPeerValidator)) // self is not in the list

	pp := NewPermissionPeers(idNonValidator, checker, nil)
	defer pp.Close()

	if !pp.PeerAllowed(idNonValidator, false, false) {
		t.Errorf("non-validator self should accept any peer")
	}
}

// errChecker is a ValidatorChecker whose lookups fail for a chosen address, and otherwise
// report a fixed validator set. It models a StakeHub backend that cannot answer.
type errChecker struct {
	validators map[common.Address]struct{}
	failOn     map[common.Address]struct{} // addresses whose lookup returns an error
	err        error
}

func (c *errChecker) IsValidator(addr common.Address) (bool, error) {
	if _, fail := c.failOn[addr]; fail {
		return false, c.err
	}
	_, ok := c.validators[addr]
	return ok, nil
}

// TestVerifyPermissionFailOpen verifies A-2: when an eligibility lookup cannot be resolved
// (returns an error), the peer is allowed rather than rejected — for both the self check and
// the peer check — while a definitive "not a validator" (nil error, false) is still rejected.
func TestVerifyPermissionFailOpen(t *testing.T) {
	lookupErr := errors.New("stakehub unavailable")

	// (1) Peer lookup fails while self is a validator → fail open (allow).
	c1 := &errChecker{
		validators: map[common.Address]struct{}{idAddr(idSelfValidator): {}},
		failOn:     map[common.Address]struct{}{idAddr(idNonValidator): {}},
		err:        lookupErr,
	}
	pp1 := &PermissionPeers{self: idSelfValidator, validators: c1}
	if err := pp1.verifyPermission(idNonValidator, false, false); err != nil {
		t.Errorf("peer lookup failure should fail open (allow), got %v", err)
	}

	// (2) Self lookup fails → fail open (allow), regardless of the peer.
	c2 := &errChecker{
		validators: map[common.Address]struct{}{idAddr(idSelfValidator): {}},
		failOn:     map[common.Address]struct{}{idAddr(idSelfValidator): {}},
		err:        lookupErr,
	}
	pp2 := &PermissionPeers{self: idSelfValidator, validators: c2}
	if err := pp2.verifyPermission(idNonValidator, false, false); err != nil {
		t.Errorf("self lookup failure should fail open (allow), got %v", err)
	}

	// (3) Definitive "peer is not a validator" (no error) is still rejected — fail-open must
	// not soften a resolved negative.
	c3 := &errChecker{
		validators: map[common.Address]struct{}{idAddr(idSelfValidator): {}},
	}
	pp3 := &PermissionPeers{self: idSelfValidator, validators: c3}
	if err := pp3.verifyPermission(idNonValidator, false, false); err == nil {
		t.Errorf("a resolved non-validator peer must still be rejected")
	}
}

// TestCheckPeerPermissionNilOpen verifies that a nil PermissionPeers (permissioning
// not configured) always passes.
func TestCheckPeerPermissionNilOpen(t *testing.T) {
	var pp *PermissionPeers
	if err := pp.checkPeerPermission(nil); err != nil {
		t.Errorf("nil PermissionPeers should permit (open), got %v", err)
	}
}

// TestPermissionPeersCloseNil verifies that Close is safe on a nil receiver.
func TestPermissionPeersCloseNil(t *testing.T) {
	var pp *PermissionPeers
	pp.Close() // must not panic
}

// TestPermissionBootnodeStillRejected verifies that configuring a bootnode does not
// change the allow/reject decision (it only suppresses the reject warn): a bootnode that
// is not an eligible validator/static/trusted is still rejected, while isBootnode
// identifies it for the quieter log path.
func TestPermissionBootnodeStillRejected(t *testing.T) {
	bootnodeID := enode.ID{0xbb}
	otherID := enode.ID{0xcc}
	checker := newMapChecker(idAddr(idSelfValidator), idAddr(idPeerValidator))

	pp := NewPermissionPeers(idSelfValidator, checker, map[enode.ID]struct{}{bootnodeID: {}})
	defer pp.Close()

	// bootnode ID is recognized (so its rejection is logged quietly)...
	if !pp.isBootnode(bootnodeID) {
		t.Error("configured bootnode ID should be recognized by isBootnode")
	}
	if pp.isBootnode(otherID) {
		t.Error("non-bootnode ID should not be recognized as a bootnode")
	}
	// ...but permissioning still rejects a non-validator/non-static/non-trusted peer
	// (the bootnode set only affects logging, not the allow/reject decision).
	if pp.PeerAllowed(idNonValidator, false, false) {
		t.Error("non-validator peer must still be rejected")
	}
	if !pp.PeerAllowed(idPeerValidator, false, false) {
		t.Error("eligible validator peer should be permitted")
	}
}

// TestPeerAllowed verifies the sweep-time decision (PeerAllowed), which mirrors
// verifyPermission but is called directly (not through the loop goroutine).
func TestPeerAllowed(t *testing.T) {
	// nil receiver: permissioning not configured, everything allowed.
	var nilPP *PermissionPeers
	if !nilPP.PeerAllowed(idNonValidator, false, false) {
		t.Fatal("nil PermissionPeers should allow any peer")
	}

	checker := newMapChecker(idAddr(idSelfValidator), idAddr(idPeerValidator))

	// self is not a validator → open, everything allowed.
	openPP := &PermissionPeers{self: idNonValidator, validators: checker}
	if !openPP.PeerAllowed(idPeerValidator, false, false) {
		t.Error("non-validator self should allow any peer")
	}

	// self is a validator → apply permissioning rules.
	pp := &PermissionPeers{self: idSelfValidator, validators: checker}
	tests := []struct {
		name    string
		id      enode.ID
		trusted bool
		static  bool
		want    bool
	}{
		{"eligible validator peer allowed", idPeerValidator, false, false, true},
		{"non-validator rejected", idNonValidator, false, false, false},
		{"non-validator trusted allowed", idNonValidator, true, false, true},
		{"non-validator static allowed", idNonValidator, false, true, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pp.PeerAllowed(tt.id, tt.trusted, tt.static); got != tt.want {
				t.Errorf("PeerAllowed(%v, trusted=%v, static=%v) = %v, want %v",
					tt.id, tt.trusted, tt.static, got, tt.want)
			}
		})
	}
}

// mockValidatorSource is a test istanbulValidatorSource that records its call and
// returns a fixed result.
type mockValidatorSource struct {
	result    bool
	err       error
	called    bool
	gotAddr   common.Address
	gotHeader *types.Header
}

func (m *mockValidatorSource) IsValidatorAt(chain consensus.ChainHeaderReader, header *types.Header, addr common.Address) (bool, error) {
	m.called = true
	m.gotAddr = addr
	m.gotHeader = header
	return m.result, m.err
}

// stubChainHeaderReader is a minimal consensus.ChainHeaderReader for tests; only
// CurrentHeader returns a meaningful value.
type stubChainHeaderReader struct {
	header *types.Header
}

func (s *stubChainHeaderReader) Config() *params.ChainConfig                 { return nil }
func (s *stubChainHeaderReader) CurrentHeader() *types.Header                { return s.header }
func (s *stubChainHeaderReader) GetHeader(common.Hash, uint64) *types.Header { return nil }
func (s *stubChainHeaderReader) GetHeaderByNumber(uint64) *types.Header      { return nil }
func (s *stubChainHeaderReader) GetHeaderByHash(common.Hash) *types.Header   { return nil }
func (s *stubChainHeaderReader) GetTd(common.Hash, uint64) *big.Int          { return nil }

// TestNewIstanbulValidatorCheckerNonIstanbul verifies that a non-Istanbul (here nil)
// engine yields no checker, i.e. permissioning stays disabled.
func TestNewIstanbulValidatorCheckerNonIstanbul(t *testing.T) {
	if c := NewIstanbulValidatorChecker(nil, nil); c != nil {
		t.Errorf("expected nil checker for non-Istanbul engine, got %v", c)
	}
}

// TestIstanbulValidatorCheckerIsValidator verifies the Istanbul adapter: it returns
// false when there is no current header and otherwise delegates to the engine.
func TestIstanbulValidatorCheckerIsValidator(t *testing.T) {
	// No current header → false, and the source must not be consulted.
	src := &mockValidatorSource{result: true}
	c := &istanbulValidatorChecker{source: src, chain: &stubChainHeaderReader{header: nil}}
	if ok, err := c.IsValidator(idAddr(idPeerValidator)); ok || err != nil {
		t.Errorf("IsValidator should be (false, nil) when the current header is nil, got (%v, %v)", ok, err)
	}
	if src.called {
		t.Error("source should not be consulted when the header is nil")
	}

	// Current header present → delegate to the source and pass through the result.
	header := &types.Header{Number: big.NewInt(10)}
	for _, want := range []bool{true, false} {
		src := &mockValidatorSource{result: want}
		c := &istanbulValidatorChecker{source: src, chain: &stubChainHeaderReader{header: header}}
		if got, err := c.IsValidator(idAddr(idPeerValidator)); got != want || err != nil {
			t.Errorf("IsValidator() = (%v, %v), want (%v, nil)", got, err, want)
		}
		if !src.called {
			t.Error("source should be consulted when the header is present")
		}
		if src.gotAddr != idAddr(idPeerValidator) {
			t.Errorf("source received addr %v, want %v", src.gotAddr, idAddr(idPeerValidator))
		}
		if src.gotHeader != header {
			t.Error("source should receive the current header")
		}
	}
}
