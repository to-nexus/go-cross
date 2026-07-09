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
// When this node is a validator, it only connects to eligible validators and to
// static/trusted nodes; otherwise it applies no restriction (open).

import (
	"errors"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p/enode"
)

// ValidatorChecker decides whether an address is an eligible validator.
type ValidatorChecker interface {
	IsValidator(addr common.Address) bool
}

// errPeerNotPermitted is returned for a peer that fails verification.
var errPeerNotPermitted = errors.New("peer rejected by consensus permissioning (not an eligible validator/static/trusted)")

// permissionTimeout is the maximum time to wait for a verification request/response.
const permissionTimeout = 5 * time.Second

// permissionRequest is a single verification request. The result is returned over the result channel.
type permissionRequest struct {
	id      enode.ID   // node ID of the remote peer
	trusted bool       // whether the peer is a trusted node
	static  bool       // whether the peer is a static node
	result  chan error // verification result (nil = allowed, err = rejected)
}

// PermissionPeers verifies peers for consensus permissioning. Owned by the handler and
// passed to Handshake.
type PermissionPeers struct {
	self       enode.ID              // node ID of this node itself
	validators ValidatorChecker      // validator checker (static / StakeHub, etc.)
	bootnodes  map[enode.ID]struct{} // configured bootnode IDs (still rejected, but not warned)

	reqCh chan *permissionRequest
	quit  chan struct{}
}

// NewPermissionPeers creates a PermissionPeers and starts its verification loop.
// bootnodes are still rejected, but their rejection is logged quietly (not warned).
func NewPermissionPeers(self enode.ID, checker ValidatorChecker, bootnodes map[enode.ID]struct{}) *PermissionPeers {
	pp := &PermissionPeers{
		self:       self,
		validators: checker,
		bootnodes:  bootnodes,
		reqCh:      make(chan *permissionRequest),
		quit:       make(chan struct{}),
	}
	go pp.loop()
	log.Info("Consensus permissioning enabled", "self", self, "bootnodes", len(bootnodes))
	return pp
}

// isBootnode reports whether id is one of the configured bootnodes.
func (pp *PermissionPeers) isBootnode(id enode.ID) bool {
	_, ok := pp.bootnodes[id]
	return ok
}

// Close terminates the verification loop. It is safe to call on a nil receiver.
func (pp *PermissionPeers) Close() {
	if pp == nil {
		return
	}
	close(pp.quit)
}

// isValidatorNode reports whether the node (by its nodekey address) is an eligible
// validator. With no checker it is treated as not eligible (= open).
func (pp *PermissionPeers) isValidatorNode(id enode.ID) bool {
	return pp.validators != nil && pp.validators.IsValidator(common.BytesToAddress(id[12:]))
}

// selfIsValidatorNode returns whether this node itself is an eligible validator.
func (pp *PermissionPeers) selfIsValidatorNode() bool {
	return pp.isValidatorNode(pp.self)
}

// LogSelfStatus logs this node's own validator status. Call only after the node is started
func (pp *PermissionPeers) LogSelfStatus() {
	if pp == nil {
		return
	}
	log.Warn("Consensus permissioning self status", "self", pp.self.String(), "isValidator", pp.selfIsValidatorNode())
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

// verifyPermission makes the allow/reject decision for a peer.
func (pp *PermissionPeers) verifyPermission(req *permissionRequest) error {
	// static/trusted always allowed (cheap check first, avoids the eligibility lookup).
	if req.trusted || req.static {
		return nil
	}
	// if we are not a validator, everything is allowed (open).
	if !pp.selfIsValidatorNode() {
		return nil
	}
	// as a validator, allow only eligible-validator peers.
	if pp.isValidatorNode(req.id) {
		return nil
	}
	return errPeerNotPermitted
}

// PeerAllowed reports whether a connected peer may stay connected; used by the periodic
// sweep. Safe to call concurrently (reads immutable fields only); nil receiver allows all.
func (pp *PermissionPeers) PeerAllowed(id enode.ID, trusted, static bool) bool {
	if pp == nil {
		return true
	}
	return pp.verifyPermission(&permissionRequest{id: id, trusted: trusted, static: static}) == nil
}

// checkPeerPermission is called from Handshake: it verifies the peer and returns an error
// to reject it (aborting the handshake). A nil receiver always passes.
func (pp *PermissionPeers) checkPeerPermission(p *Peer) error {
	if pp == nil {
		return nil
	}
	id := p.Node().ID()

	req := &permissionRequest{
		id:      id,
		trusted: p.Trusted(),
		static:  p.StaticDialed(),
		result:  make(chan error, 1),
	}

	// submit the request (timeout guards against a busy/stopped handler).
	select {
	case pp.reqCh <- req:
	case <-pp.quit:
		return nil
	case <-time.After(permissionTimeout):
		log.Warn("Consensus permissioning timed out on submit", "id", id)
		return errPeerNotPermitted
	}
	// wait for the result.
	select {
	case err := <-req.result:
		if err != nil {
			// bootnodes keep dialing validators, so log their rejection quietly.
			if !pp.isBootnode(id) {
				log.Warn("Rejecting peer by consensus permissioning",
					"id", id.String(), "trusted", req.trusted, "static", req.static, "err", err)
			}
		} else {
			log.Trace("Peer permitted by consensus permissioning", "id", id.String())
		}
		return err
	case <-time.After(permissionTimeout):
		log.Warn("Consensus permissioning timed out on result", "id", id.String())
		return errPeerNotPermitted
	}
}

// istanbulValidatorSource is the minimal Istanbul engine interface for eligibility checks.
type istanbulValidatorSource interface {
	IsValidatorAt(chain consensus.ChainHeaderReader, header *types.Header, addr common.Address) bool
}

// istanbulValidatorChecker implements ValidatorChecker via the Istanbul engine,
// querying eligibility against the current head.
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
