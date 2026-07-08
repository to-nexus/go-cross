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
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

// Addresses for testing.
var (
	addrSelfValidator = common.HexToAddress("0x0000000000000000000000000000000000000011") // self (validator)
	addrPeerValidator = common.HexToAddress("0x0000000000000000000000000000000000000022") // peer validator
	addrNonValidator  = common.HexToAddress("0x0000000000000000000000000000000000000099") // not a validator
)

// mapValidatorChecker is a ValidatorChecker implementation for testing (address-set based).
type mapValidatorChecker map[common.Address]struct{}

func newMapChecker(addrs ...common.Address) mapValidatorChecker {
	m := make(mapValidatorChecker, len(addrs))
	for _, a := range addrs {
		m[a] = struct{}{}
	}
	return m
}

func (m mapValidatorChecker) IsValidator(addr common.Address) bool {
	_, ok := m[addr]
	return ok
}

// TestVerifyPermission verifies the permission decision rules in a table-driven way.
func TestVerifyPermission(t *testing.T) {
	// Register self and the peer validator as validators.
	checker := newMapChecker(addrSelfValidator, addrPeerValidator)

	// Construct directly to verify only the decision logic without the loop goroutine.
	newPP := func(self common.Address) *PermissionPeers {
		return &PermissionPeers{self: self, validators: checker}
	}

	tests := []struct {
		name    string
		self    common.Address
		req     *permissionRequest
		wantErr bool
	}{
		{
			name:    "allow anyone if self is not a validator",
			self:    addrNonValidator,
			req:     &permissionRequest{addr: addrNonValidator},
			wantErr: false,
		},
		{
			name:    "validator <-> validator is allowed", // core case
			self:    addrSelfValidator,
			req:     &permissionRequest{addr: addrPeerValidator},
			wantErr: false,
		},
		{
			name:    "validator rejects a non-validator",
			self:    addrSelfValidator,
			req:     &permissionRequest{addr: addrNonValidator},
			wantErr: true,
		},
		{
			name:    "validator allows a non-validator if trusted",
			self:    addrSelfValidator,
			req:     &permissionRequest{addr: addrNonValidator, trusted: true},
			wantErr: false,
		},
		{
			name:    "validator allows a non-validator if static",
			self:    addrSelfValidator,
			req:     &permissionRequest{addr: addrNonValidator, static: true},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pp := newPP(tt.self)
			err := pp.verifyPermission(tt.req)
			if (err != nil) != tt.wantErr {
				t.Fatalf("verifyPermission() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestPermissionPeersChannel verifies the verification path through the channel handler
// (loop goroutine) end-to-end: connections between validators are allowed, others rejected.
func TestPermissionPeersChannel(t *testing.T) {
	checker := newMapChecker(addrSelfValidator, addrPeerValidator)

	// A node that is itself a validator.
	pp := NewPermissionPeers(addrSelfValidator, checker)
	defer pp.Close()

	// Helper that sends a request over the channel and receives the result.
	ask := func(req *permissionRequest) error {
		req.result = make(chan error, 1)
		select {
		case pp.reqCh <- req:
		case <-time.After(time.Second):
			t.Fatal("failed to submit permission request")
		}
		select {
		case err := <-req.result:
			return err
		case <-time.After(time.Second):
			t.Fatal("timed out waiting for permission result")
			return nil
		}
	}

	// validator peer → allowed
	if err := ask(&permissionRequest{addr: addrPeerValidator}); err != nil {
		t.Errorf("validator peer should be permitted, got %v", err)
	}
	// non-validator peer → rejected
	if err := ask(&permissionRequest{addr: addrNonValidator}); err == nil {
		t.Errorf("non-validator peer should be rejected")
	}
	// non-validator but trusted → allowed
	if err := ask(&permissionRequest{addr: addrNonValidator, trusted: true}); err != nil {
		t.Errorf("trusted peer should be permitted, got %v", err)
	}
}

// TestPermissionPeersOpenWhenNotValidator verifies that when this node is not a
// validator, all peers are allowed (open).
func TestPermissionPeersOpenWhenNotValidator(t *testing.T) {
	checker := newMapChecker(addrPeerValidator) // self is not in the list

	pp := NewPermissionPeers(addrNonValidator, checker)
	defer pp.Close()

	req := &permissionRequest{addr: addrNonValidator, result: make(chan error, 1)}
	pp.reqCh <- req
	if err := <-req.result; err != nil {
		t.Errorf("non-validator self should accept any peer, got %v", err)
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

// TestPeerAllowed verifies the sweep-time decision (PeerAllowed), which mirrors
// verifyPermission but is called directly (not through the loop goroutine).
func TestPeerAllowed(t *testing.T) {
	// nil receiver: permissioning not configured, everything allowed.
	var nilPP *PermissionPeers
	if !nilPP.PeerAllowed(addrNonValidator, false, false) {
		t.Fatal("nil PermissionPeers should allow any peer")
	}

	checker := newMapChecker(addrSelfValidator, addrPeerValidator)

	// self is not a validator → open, everything allowed.
	openPP := &PermissionPeers{self: addrNonValidator, validators: checker}
	if !openPP.PeerAllowed(addrPeerValidator, false, false) {
		t.Error("non-validator self should allow any peer")
	}

	// self is a validator → apply permissioning rules.
	pp := &PermissionPeers{self: addrSelfValidator, validators: checker}
	tests := []struct {
		name    string
		addr    common.Address
		trusted bool
		static  bool
		want    bool
	}{
		{"eligible validator peer allowed", addrPeerValidator, false, false, true},
		{"non-validator rejected", addrNonValidator, false, false, false},
		{"non-validator trusted allowed", addrNonValidator, true, false, true},
		{"non-validator static allowed", addrNonValidator, false, true, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pp.PeerAllowed(tt.addr, tt.trusted, tt.static); got != tt.want {
				t.Errorf("PeerAllowed(%v, trusted=%v, static=%v) = %v, want %v",
					tt.addr, tt.trusted, tt.static, got, tt.want)
			}
		})
	}
}

// mockValidatorSource is a test istanbulValidatorSource that records its call and
// returns a fixed result.
type mockValidatorSource struct {
	result    bool
	called    bool
	gotAddr   common.Address
	gotHeader *types.Header
}

func (m *mockValidatorSource) IsValidatorAt(chain consensus.ChainHeaderReader, header *types.Header, addr common.Address) bool {
	m.called = true
	m.gotAddr = addr
	m.gotHeader = header
	return m.result
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
	if c.IsValidator(addrPeerValidator) {
		t.Error("IsValidator should be false when the current header is nil")
	}
	if src.called {
		t.Error("source should not be consulted when the header is nil")
	}

	// Current header present → delegate to the source and pass through the result.
	header := &types.Header{Number: big.NewInt(10)}
	for _, want := range []bool{true, false} {
		src := &mockValidatorSource{result: want}
		c := &istanbulValidatorChecker{source: src, chain: &stubChainHeaderReader{header: header}}
		if got := c.IsValidator(addrPeerValidator); got != want {
			t.Errorf("IsValidator() = %v, want %v", got, want)
		}
		if !src.called {
			t.Error("source should be consulted when the header is present")
		}
		if src.gotAddr != addrPeerValidator {
			t.Errorf("source received addr %v, want %v", src.gotAddr, addrPeerValidator)
		}
		if src.gotHeader != header {
			t.Error("source should receive the current header")
		}
	}
}
