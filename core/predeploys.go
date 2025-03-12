package core

// ##CROSS: predeploys

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/predeploys"
)

var (
	ERC1967ProxyImplementationSlot = common.HexToHash("0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc")
)

const (
	CrossEx    = "0x0000000000000000000000000000000FEEDADDED"
	Bridge     = "0xB81D6E000000000000000000000000000000FACE"
	BridgeImpl = "0xB81D6E000000000000000000000000000000C0DE"
)

var (
	// addresses
	CrossExAddr    = common.HexToAddress(CrossEx)
	BridgeAddr     = common.HexToAddress(Bridge)
	BridgeImplAddr = common.HexToAddress(BridgeImpl)

	// predeploys
	Predeploys = map[common.Address]types.Account{
		CrossExAddr: {
			Code: common.Hex2Bytes(predeploys.CrossExBinRuntime),
		},
		BridgeAddr: {
			Balance: new(big.Int).Mul(big.NewInt(99_900_000_000), big.NewInt(1e18)),
			Code:    common.Hex2Bytes(predeploys.ERC1967ProxyBinRuntime),
			Storage: map[common.Hash]common.Hash{
				ERC1967ProxyImplementationSlot: common.BytesToHash(BridgeImplAddr.Bytes()),
			},
		},
		BridgeImplAddr: {
			Code: common.Hex2Bytes(predeploys.CrossBridgeBinRuntime),
		},
	}
)
