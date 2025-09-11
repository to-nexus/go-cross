package predeploy

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts"
	"github.com/ethereum/go-ethereum/contracts/predeploy"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

// ##CROSS: predeploys

var (
	alloc = types.GenesisAlloc{
		contracts.CrossExAddr: {
			Code: common.Hex2Bytes(predeploy.CrossExCode),
		},
		contracts.BridgeAddr: {
			Balance: new(big.Int).Mul(big.NewInt(980_000_000), big.NewInt(1e18)),
			Code:    common.Hex2Bytes(predeploy.CrossBridgeCode),
			Storage: map[common.Hash]common.Hash{
				contracts.BridgeImplSlot: common.BytesToHash(contracts.BridgeImplAddr.Bytes()),
			},
		},
		contracts.BridgeImplAddr: {
			Code: common.Hex2Bytes(predeploy.CrossBridgeImplCode),
		},
		contracts.Multicall3Addr: {
			Code: common.Hex2Bytes(predeploy.Multicall3Code),
		},
	}

	GenesisAllocCross = types.GenesisAlloc{
		params.FoundationCross: {
			Balance: new(big.Int).Mul(big.NewInt(10_000_000), big.NewInt(1e18)),
		},
		params.EcoCross: {
			Balance: new(big.Int).Mul(big.NewInt(10_000_000), big.NewInt(1e18)),
		},
		contracts.CrossExAddr:    alloc[contracts.CrossExAddr],
		contracts.BridgeAddr:     alloc[contracts.BridgeAddr],
		contracts.BridgeImplAddr: alloc[contracts.BridgeImplAddr],
		contracts.Multicall3Addr: alloc[contracts.Multicall3Addr],
	}

	GenesisAllocZoneZero = types.GenesisAlloc{
		params.FoundationZoneZero: {
			Balance: new(big.Int).Mul(big.NewInt(1_000_000_000_000), big.NewInt(1e18)),
		},
		contracts.CrossExAddr:    alloc[contracts.CrossExAddr],
		contracts.BridgeImplAddr: alloc[contracts.BridgeImplAddr],
		contracts.BridgeAddr:     alloc[contracts.BridgeAddr],
		contracts.Multicall3Addr: alloc[contracts.Multicall3Addr],
	}

	GenesisAllocCrossDev3 = types.GenesisAlloc{
		params.FoundationCrossDev3: {
			Balance: new(big.Int).Mul(big.NewInt(100_000_000), big.NewInt(1e18)),
		},
		contracts.CrossExAddr:    alloc[contracts.CrossExAddr],
		contracts.BridgeImplAddr: alloc[contracts.BridgeImplAddr],
		contracts.BridgeAddr:     alloc[contracts.BridgeAddr],
		contracts.Multicall3Addr: alloc[contracts.Multicall3Addr],
	}

	GenesisAllocCrossDev = types.GenesisAlloc{
		params.FoundationCrossDev: {
			Balance: new(big.Int).Mul(big.NewInt(100_000_000), big.NewInt(1e18)),
		},
		contracts.CrossExAddr:    alloc[contracts.CrossExAddr],
		contracts.BridgeImplAddr: alloc[contracts.BridgeImplAddr],
		contracts.BridgeAddr:     alloc[contracts.BridgeAddr],
		contracts.Multicall3Addr: alloc[contracts.Multicall3Addr],
	}
)
