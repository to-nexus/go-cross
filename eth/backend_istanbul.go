package eth

import (
	"github.com/ethereum/go-ethereum/eth/protocols/eth"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/enode"
)

// quorum_protocol enables the eth service to return two different protocols, one for the eth mainnet "eth" service,
// and one for the quorum specific consensus algo, obtained from engine.consensus
// 2021 Jan in the future consensus (istanbul) may run from its own service and use a single subprotocol there,
// instead of overloading the eth service.

const (
	// istanbul consensus Protocol variables are optionally set in addition to the "eth" protocol variables (eth/protocol.go).
	istanbulConsensusProtocolName = "istanbul"
	Istanbul100                   = 100
)

var (
	// ProtocolVersions are the supported versions of the quorum consensus protocol (first is primary), e.g. []uint{Istanbul64, Istanbul99, Istanbul100}.
	istanbulConsensusProtocolVersions = []uint{Istanbul100}
	// protocol Length describe the number of messages support by the protocol/version map[uint]uint64{Istanbul64: 18, Istanbul99: 18, Istanbul100: 18}
	istanbulConsensusProtocolLengths = map[uint]uint64{Istanbul100: 22}
)

func (s *Ethereum) istanbulConsensusProtocols(backend eth.Backend, network uint64, dnsdisc enode.Iterator) []p2p.Protocol {
	// Set protocol Name/Version
	// keep `var protocolName = "eth"` as is, and only update the quorum consensus specific protocol
	// This is used to enable the eth service to return multiple devp2p subprotocols.
	// Previously, for istanbul/64 istnbul/99 and clique (v2.6) `protocolName` would be overridden and
	// set to the consensus subprotocol name instead of "eth", meaning the node would no longer
	// communicate over the "eth" subprotocol, e.g. "eth" or "istanbul/99" but not eth" and "istanbul/99".
	// With this change, support is added so that the "eth" subprotocol remains and optionally a consensus subprotocol
	// can be added allowing the node to communicate over "eth" and an optional consensus subprotocol, e.g. "eth" and "istanbul/100"

	protos := make([]p2p.Protocol, len(istanbulConsensusProtocolVersions))
	for i, vsn := range istanbulConsensusProtocolVersions {
		if length, ok := istanbulConsensusProtocolLengths[vsn]; !ok {
			panic("makeIstanbulConsensusProtocol for unknown version")
		} else {
			protos[i] = (*istanbulHandler)(s.handler).makeIstanbulConsensusProtocol(istanbulConsensusProtocolName, vsn, length, backend, network, dnsdisc)
		}
	}
	return protos
}
