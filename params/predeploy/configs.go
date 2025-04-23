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

	CrossExAddr    = common.HexToAddress("0xfeed00000000000000000000000000000000C0DE")
	BridgeAddr     = common.HexToAddress("0xb81d6e000000000000000000000000000000C0de")
	BridgeImplAddr = common.HexToAddress("0xB81D6e000000000000000000000000000000AAaA")
	Multicall3Addr = common.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11")
)

var (
	GenesisAllocCross = types.GenesisAlloc{
		params.FoundationCross: {
			Balance: new(big.Int).Mul(big.NewInt(10_000_000), big.NewInt(1e18)),
		},
		params.EchoCross: {
			Balance: new(big.Int).Mul(big.NewInt(10_000_000), big.NewInt(1e18)),
		},
		CrossExAddr: {
			Code: common.Hex2Bytes(CrossExBinRuntime),
		},
		BridgeAddr: {
			Balance: new(big.Int).Mul(big.NewInt(980_000_000), big.NewInt(1e18)),
			Code:    common.Hex2Bytes(ERC1967ProxyBinRuntime),
			Storage: map[common.Hash]common.Hash{
				ERC1967ProxyImplementationSlot: common.BytesToHash(BridgeImplAddr.Bytes()),
			},
		},
		BridgeImplAddr: {
			Code: common.Hex2Bytes(CrossBridgeBinRuntime),
		},
		Multicall3Addr: {
			Code: common.Hex2Bytes(Multicall3BinRuntime),
		},
	}

	GenesisAllocZoneZero = types.GenesisAlloc{
		params.FoundationZoneZero: {
			Balance: new(big.Int).Mul(big.NewInt(1_000_000_000_000), big.NewInt(1e18)),
		},
		CrossExAddr:    GenesisAllocCross[CrossExAddr],
		BridgeImplAddr: GenesisAllocCross[BridgeImplAddr],
		BridgeAddr:     GenesisAllocCross[BridgeAddr],
		Multicall3Addr: GenesisAllocCross[Multicall3Addr],
	}

	GenesisAllocCrossDev3 = types.GenesisAlloc{
		params.FoundationCrossDev3: {
			Balance: new(big.Int).Mul(big.NewInt(100_000_000), big.NewInt(1e18)),
		},
		CrossExAddr:    GenesisAllocCross[CrossExAddr],
		BridgeImplAddr: GenesisAllocCross[BridgeImplAddr],
		BridgeAddr:     GenesisAllocCross[BridgeAddr],
		Multicall3Addr: GenesisAllocCross[Multicall3Addr],
	}

	GenesisAllocCrossDev = types.GenesisAlloc{
		params.FoundationCrossDev: {
			Balance: new(big.Int).Mul(big.NewInt(100_000_000), big.NewInt(1e18)),
		},
		CrossExAddr:    GenesisAllocCross[CrossExAddr],
		BridgeImplAddr: GenesisAllocCross[BridgeImplAddr],
		BridgeAddr:     GenesisAllocCross[BridgeAddr],
		Multicall3Addr: GenesisAllocCross[Multicall3Addr],
	}
)
