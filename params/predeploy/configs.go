package predeploy

// ##CROSS: predeploys

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

var (
	ERC1967ProxyImplementationSlot = common.HexToHash("0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc")

	CrossExAddr    = common.HexToAddress("0xFEED00000000000000000000000000000000C0DE")
	BridgeAddr     = common.HexToAddress("0xB81D6E000000000000000000000000000000C0DE")
	BridgeImplAddr = common.HexToAddress("0xB81D6E000000000000000000000000000000AAAA")
)

var (
	GenesisAllocCross = types.GenesisAlloc{
		params.FoundationCross: {
			Balance: new(big.Int).Mul(big.NewInt(50_000_000), big.NewInt(1e18)),
		},
		CrossExAddr: {
			Code: common.Hex2Bytes(CrossExBinRuntime),
		},
		BridgeAddr: {
			Balance: new(big.Int).Mul(big.NewInt(950_000_000), big.NewInt(1e18)),
			Code:    common.Hex2Bytes(ERC1967ProxyBinRuntime),
			Storage: map[common.Hash]common.Hash{
				ERC1967ProxyImplementationSlot: common.BytesToHash(BridgeImplAddr.Bytes()),
			},
		},
		BridgeImplAddr: {
			Code: common.Hex2Bytes(CrossBridgeBinRuntime),
		},
	}

	GenesisAllocCrossTest = types.GenesisAlloc{
		params.FoundationCrossTest: {
			Balance: new(big.Int).Mul(big.NewInt(1_000_000_000_000), big.NewInt(1e18)),
		},
		CrossExAddr:    GenesisAllocCross[CrossExAddr],
		BridgeImplAddr: GenesisAllocCross[BridgeImplAddr],
		BridgeAddr:     GenesisAllocCross[BridgeAddr],
	}

	GenesisAllocCrossDev3 = types.GenesisAlloc{
		params.FoundationCrossDev3: {
			Balance: new(big.Int).Mul(big.NewInt(100_000_000), big.NewInt(1e18)),
		},
		CrossExAddr:    GenesisAllocCross[CrossExAddr],
		BridgeImplAddr: GenesisAllocCross[BridgeImplAddr],
		BridgeAddr:     GenesisAllocCross[BridgeAddr],
	}

	GenesisAllocCrossDev = types.GenesisAlloc{
		params.FoundationCrossDev: {
			Balance: new(big.Int).Mul(big.NewInt(100_000_000), big.NewInt(1e18)),
		},
		CrossExAddr:    GenesisAllocCross[CrossExAddr],
		BridgeImplAddr: GenesisAllocCross[BridgeImplAddr],
		BridgeAddr:     GenesisAllocCross[BridgeAddr],
	}
)
