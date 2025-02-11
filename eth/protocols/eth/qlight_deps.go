package eth

import (
	"github.com/ethereum/go-ethereum/core"
)

func CurrentENREntry(chain *core.BlockChain) *enrEntry {
	return currentENREntry(chain)
}

func NodeInfoFunc(chain *core.BlockChain, network uint64) *NodeInfo {
	return nodeInfo(chain, network)
}
