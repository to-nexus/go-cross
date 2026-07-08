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
//
// PermissionPeers verifies the remote peer during the handshake. If this node is
// an eligible validator:
//   - the peer is an eligible validator, or
//   - the peer is registered as a static / trusted node,
// the connection is allowed; otherwise the handshake is aborted and the connection
// is dropped. If this node is not an eligible validator, no restriction is applied (open).
//
// "Eligible validator" means the address passes the ValidatorChecker: it is not just
// registered but also meets the consensus requirements (e.g. not blacklisted and holds
// the minimum stake). A registered-but-ineligible validator is treated the same as a
// non-validator. Whether an address is eligible is abstracted behind the ValidatorChecker
// interface. The real implementation is backed by the Istanbul consensus engine
// (StakeHub when PoSA is active, otherwise config/snapshot) and is injected by the
// parent package (eth). Verification is serialized in the loop goroutine (channel handler).

import (
	"errors"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

// ValidatorChecker abstracts the logic that decides whether an address is an
// eligible validator (registered and meeting the consensus requirements, e.g. not
// blacklisted and holding the minimum stake). Both this node and the remote peer are
// looked up through this interface. A hardcoded static implementation and a StakeHub
// contract lookup implementation both satisfy this interface.
type ValidatorChecker interface {
	IsValidator(addr common.Address) bool
}

// errPeerNotPermitted is returned for a peer that fails verification.
var errPeerNotPermitted = errors.New("peer rejected by consensus permissioning (not an eligible validator/static/trusted)")

// permissionTimeout is the maximum time to wait for a verification request/response.
const permissionTimeout = 5 * time.Second

// permissionRequest is a single verification request. The result is returned over the result channel.
type permissionRequest struct {
	addr    common.Address // nodekey address of the remote peer
	trusted bool           // whether the peer is a trusted node
	static  bool           // whether the peer is a static node
	result  chan error     // verification result (nil = allowed, err = rejected)
}

// PermissionPeers is the verifier responsible for consensus peer permissioning.
// The handler creates and owns it using its own nodekey address and passes it to Handshake.
type PermissionPeers struct {
	self       common.Address   // nodekey address of this node itself
	validators ValidatorChecker // validator checker (static / StakeHub, etc.)

	reqCh chan *permissionRequest
	quit  chan struct{}
}

// NewPermissionPeers creates a PermissionPeers and starts the verification loop.
// self is this node's nodekey address, checker is the validator checker.
func NewPermissionPeers(self common.Address, checker ValidatorChecker) *PermissionPeers {
	pp := &PermissionPeers{
		self:       self,
		validators: checker,
		reqCh:      make(chan *permissionRequest),
		quit:       make(chan struct{}),
	}
	go pp.loop()
	log.Info("Consensus permissioning enabled", "self", self, "isValidator", pp.selfIsValidator())
	return pp
}

// Close terminates the verification loop. It is safe to call on a nil receiver.
func (pp *PermissionPeers) Close() {
	if pp == nil {
		return
	}
	close(pp.quit)
}

// isValidator delegates to the checker and returns whether the address is an eligible validator.
// If there is no checker (nil), the address is treated as not eligible (= open).
func (pp *PermissionPeers) isValidator(addr common.Address) bool {
	return pp.validators != nil && pp.validators.IsValidator(addr)
}

// selfIsValidator returns whether this node itself is an eligible validator.
func (pp *PermissionPeers) selfIsValidator() bool {
	return pp.isValidator(pp.self)
}

// loop is the channel handler that processes verification requests serially.
func (pp *PermissionPeers) loop() {
	for {
		select {
		case req := <-pp.reqCh:
			req.result <- pp.verifyPermission(req)
		case <-pp.quit:
			return
		}
	}
}

// verifyPermission makes the actual allow/reject decision.
func (pp *PermissionPeers) verifyPermission(req *permissionRequest) error {
	// Static / trusted peers are always allowed, regardless of validator status.
	// Check this first: it is a cheap flag test and short-circuits a possibly-cold
	// StakeHub eligibility lookup (below) for such peers.
	if req.trusted || req.static {
		return nil
	}
	// If this node is not an eligible validator, no restriction is applied.
	if !pp.selfIsValidator() {
		return nil
	}
	// This node is an eligible validator: allow only if the peer is also eligible.
	if pp.isValidator(req.addr) {
		return nil
	}
	// otherwise reject
	return errPeerNotPermitted
}

// PeerAllowed returns whether an already-connected peer (addr / trusted / static)
// may stay connected under the current validator set. It matches the handshake-time
// decision (verifyPermission) and is used by the periodic sweep. If the receiver is
// nil (permissioning not configured), it always allows. verifyPermission only reads
// immutable fields (self / validators), so it is safe to call concurrently with loop.
func (pp *PermissionPeers) PeerAllowed(addr common.Address, trusted, static bool) bool {
	if pp == nil {
		return true
	}
	return pp.verifyPermission(&permissionRequest{addr: addr, trusted: trusted, static: static}) == nil
}

// checkPeerPermission is called from Handshake. It sends a verification request to
// loop and waits until the result arrives. It returns nil on pass and an error on
// rejection; when an error is returned, the caller (Handshake) aborts the handshake
// and drops the connection. If the receiver is nil (permissioning not configured),
// it always passes.
func (pp *PermissionPeers) checkPeerPermission(p *Peer) error {
	if pp == nil {
		return nil
	}
	pub := p.Node().Pubkey()
	if pub == nil {
		return errPeerNotPermitted
	}
	addr := crypto.PubkeyToAddress(*pub)

	req := &permissionRequest{
		addr:    addr,
		trusted: p.Trusted(),
		static:  p.StaticDialed(),
		result:  make(chan error, 1),
	}

	// Send the verification request (with a timeout in case the handler is busy or stopped).
	select {
	case pp.reqCh <- req:
	case <-pp.quit:
		return nil
	case <-time.After(permissionTimeout):
		log.Warn("Consensus permissioning timed out on submit", "addr", addr)
		return errPeerNotPermitted
	}

	// Wait for the verification result.
	select {
	case err := <-req.result:
		if err != nil {
			log.Warn("Rejecting peer by consensus permissioning",
				"addr", addr, "trusted", req.trusted, "static", req.static, "err", err)
		} else {
			log.Trace("Peer permitted by consensus permissioning", "addr", addr)
		}
		return err
	case <-time.After(permissionTimeout):
		log.Warn("Consensus permissioning timed out on result", "addr", addr)
		return errPeerNotPermitted
	}
}

// istanbulValidatorSource is the minimal interface of the Istanbul engine needed for
// validator eligibility determination. *backend.Backend implements it.
type istanbulValidatorSource interface {
	IsValidatorAt(chain consensus.ChainHeaderReader, header *types.Header, addr common.Address) bool
}

// istanbulValidatorChecker implements ValidatorChecker using the Istanbul consensus
// engine's eligibility determination. Eligibility is queried directly from the
// consensus engine (StakeHub when PoSA is active, otherwise config/snapshot — the
// engine handles the branching). Queries are made against the current head.
type istanbulValidatorChecker struct {
	source istanbulValidatorSource
	chain  consensus.ChainHeaderReader
}

var _ ValidatorChecker = (*istanbulValidatorChecker)(nil)

// NewIstanbulValidatorChecker returns a checker if engine is an Istanbul consensus
// engine that exposes IsValidatorAt, otherwise it returns nil (= permissioning not applied).
func NewIstanbulValidatorChecker(engine consensus.Engine, chain consensus.ChainHeaderReader) ValidatorChecker {
	e := consensus.ToIstanbulEngine(engine)
	if e == nil {
		return nil
	}
	source, ok := e.(istanbulValidatorSource)
	if !ok {
		return nil
	}
	return &istanbulValidatorChecker{source: source, chain: chain}
}

// IsValidator returns whether the address is an eligible validator against the current head.
func (c *istanbulValidatorChecker) IsValidator(addr common.Address) bool {
	header := c.chain.CurrentHeader()
	if header == nil {
		return false
	}
	return c.source.IsValidatorAt(c.chain, header, addr)
}
