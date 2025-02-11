// Copyright 2015 The go-ethereum Authors
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
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/beacon"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/protocols/eth"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/p2p/enr"
)

var (
	// errEthPeerNil is returned when no eth peer is found to be associated with a p2p peer.
	errEthPeerNil           = errors.New("eth peer was nil")
	errEthPeerNotRegistered = errors.New("eth peer was not registered")
)

type istanbulHandler handler // ##CROSS: istanbul

func (h *istanbulHandler) Enqueue(id string, block *types.Block) {
	h.blockFetcher.Enqueue(id, block)
}

func (h *istanbulHandler) FindPeers(targets map[common.Address]bool) map[common.Address]consensus.IstanbulPeer {
	m := make(map[common.Address]consensus.IstanbulPeer)
	h.peers.lock.RLock()
	defer h.peers.lock.RUnlock()
	for _, p := range h.peers.peers {
		pubKey := p.Node().Pubkey()
		addr := crypto.PubkeyToAddress(*pubKey)
		if targets[addr] {
			m[addr] = p
		}
	}
	return m
}

// makeIstanbulConsensusProtocol is similar to eth/handler.go -> makeProtocol. Called from eth/handler.go -> Protocols.
// returns the supported subprotocol to the p2p server.
// The Run method starts the protocol and is called by the p2p server. The quorum consensus subprotocol,
// leverages the peer created and managed by the "eth" subprotocol.
// The quorum consensus protocol requires that the "eth" protocol is running as well.
func (h *istanbulHandler) makeIstanbulConsensusProtocol(protoName string, version uint, length uint64, backend eth.Backend, network uint64, dnsdisc enode.Iterator) p2p.Protocol {
	log.Debug("registering qouorum protocol ", "protoName", protoName, "version", version)

	return p2p.Protocol{
		Name:    protoName,
		Version: version,
		Length:  length,
		// no new peer created, uses the "eth" peer, so no peer management needed.
		Run: func(p *p2p.Peer, rw p2p.MsgReadWriter) error {
			/*
			* 1. wait for the eth protocol to create and register an eth peer.
			* 2. get the associate eth peer that was registered by he "eth" protocol.
			* 3. add the rw protocol for the quorum subprotocol to the eth peer.
			* 4. start listening for incoming messages.
			* 5. the incoming message will be sent on the quorum specific subprotocol, e.g. "istanbul/100".
			* 6. send messages to the consensus engine handler.
			* 7. messages to other to other peers listening to the subprotocol can be sent using the
			*    (eth)peer.ConsensusSend() which will write to the protoRW.
			 */
			// wait for the "eth" protocol to create and register the peer (added to peerset)
			select {
			case <-p.EthPeerRegistered:
				// the ethpeer should be registered, try to retrieve it and start the consensus handler.
				p2pPeerId := fmt.Sprintf("%x", p.ID().Bytes()[:8])
				ethPeer := h.peers.peer(p2pPeerId)
				if ethPeer == nil {
					p2pPeerId = fmt.Sprintf("%x", p.ID().Bytes()) //TODO:BBO
					ethPeer = h.peers.peer(p2pPeerId)
					log.Warn("full p2p peer", "id", p2pPeerId, "ethPeer", ethPeer)
				}
				if ethPeer != nil {
					p.Log().Debug("consensus subprotocol retrieved eth peer from peerset", "ethPeer.id", p2pPeerId, "ProtoName", protoName)
					// add the rw protocol for the quorum subprotocol to the eth peer.
					ethPeer.AddConsensusProtoRW(rw)
					peer := eth.NewPeer(version, p, rw, h.txpool)
					return h.handleConsensusLoop(peer, rw, nil)
				}
				p.Log().Error("consensus subprotocol retrieved nil eth peer from peerset", "ethPeer.id", p2pPeerId)
				return errEthPeerNil
			case <-p.EthPeerDisconnected:
				return errEthPeerNotRegistered
			}
		},
		NodeInfo: func() interface{} {
			return eth.NodeInfoFunc(backend.Chain(), network)
		},
		PeerInfo: func(id enode.ID) interface{} {
			return backend.PeerInfo(id)
		},
		Attributes:     []enr.Entry{eth.CurrentENREntry(backend.Chain())},
		DialCandidates: dnsdisc,
	}
}

func (h *istanbulHandler) handleConsensusLoop(p *eth.Peer, protoRW p2p.MsgReadWriter, fallThroughBackend eth.Backend) error {
	// Handle incoming messages until the connection is torn down
	for {
		if err := h.handleConsensus(p, protoRW, fallThroughBackend); err != nil {
			// allow the P2P connection to remain active during sync (when the engine is stopped)
			if errors.Is(err, istanbul.ErrStoppedEngine) && h.downloader.Synchronising() {
				// should this be warn or debug
				p.Log().Debug("Ignoring `stopped engine` consensus error due to active sync.")
				continue
			}
			p.Log().Debug("Ethereum quorum message handling failed", "err", err)
			return err
		}
	}
}

// This is a no-op because the eth handleMsg main loop handle ibf message as well.
func (h *istanbulHandler) handleConsensus(p *eth.Peer, protoRW p2p.MsgReadWriter, _ eth.Backend) error {
	// Read the next message from the remote peer (in protoRW), and ensure it's fully consumed
	msg, err := protoRW.ReadMsg()
	if err != nil {
		return err
	}
	if msg.Size > protocolMaxMsgSize {
		return fmt.Errorf("%w: %v > %v", errMsgTooLarge, msg.Size, protocolMaxMsgSize)
	}
	defer msg.Discard()

	// See if the consensus engine protocol can handle this message, e.g. istanbul will check for message is
	// istanbulMsg = 0x11, and NewBlockMsg = 0x07.
	handled, err := h.handleConsensusMsg(p, msg)
	if handled {
		p.Log().Debug("consensus message was handled by consensus engine", "msg", msg.Code, "istanbulConsensusProtocolName", istanbulConsensusProtocolName, "err", err)
		return err
	}

	return nil
}

func (h *istanbulHandler) handleConsensusMsg(p *eth.Peer, msg p2p.Msg) (bool, error) {
	var handler consensus.IstanbulHandler
	if handler, _ = h.engine.(consensus.IstanbulHandler); handler == nil {
		if beacon, ok := h.engine.(*beacon.Beacon); ok {
			handler, _ = beacon.InnerEngine().(consensus.IstanbulHandler)
		}
	}

	if handler != nil {
		pubKey := p.Node().Pubkey()
		addr := crypto.PubkeyToAddress(*pubKey)
		handled, err := handler.HandleMsg(addr, msg)
		return handled, err
	}
	return false, nil
}

// makeLegacyProtocol is basically a copy of the eth makeProtocol, but for legacy subprotocols, e.g. "istanbul/99" "istabnul/64"
// If support legacy subprotocols is removed, remove this and associated code as well.
// If quorum is using a legacy protocol then the "eth" subprotocol should not be available.
func (h *istanbulHandler) makeLegacyProtocol(protoName string, version uint, length uint64, backend eth.Backend, network uint64, dnsdisc enode.Iterator) p2p.Protocol {
	log.Debug("registering a legacy protocol ", "protoName", protoName, "version", version)
	return p2p.Protocol{
		Name:    protoName,
		Version: version,
		Length:  length,
		Run: func(p *p2p.Peer, rw p2p.MsgReadWriter) error {
			peer := eth.NewPeer(version, p, rw, h.txpool)
			return ((*handler)(h)).runEthPeer(peer, func(peer *eth.Peer) error {
				// We pass through the backend so that we can 'handle' messages that we can't handle
				return h.handleConsensusLoop(peer, rw, backend)
			})
		},
		NodeInfo: func() interface{} {
			return eth.NodeInfoFunc(backend.Chain(), network)
		},
		PeerInfo: func(id enode.ID) interface{} {
			return backend.PeerInfo(id)
		},
		Attributes:     []enr.Entry{eth.CurrentENREntry(backend.Chain())},
		DialCandidates: dnsdisc,
	}
}
