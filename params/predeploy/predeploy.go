package predeploy

import (
	_ "embed"
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts"
	"github.com/ethereum/go-ethereum/contracts/breakpoint"
	"github.com/ethereum/go-ethereum/contracts/predeploy"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

// ##CROSS: predeploys

var (
	// breakpointDevValidatorSetAllocJSON contains the onedev ValidatorSet storage override.
	//go:embed breakpoint_dev_alloc.json
	breakpointDevValidatorSetAllocJSON []byte
	// breakpointDev3AllocJSON is the shared Breakpoint allocation generated for onedev3.
	//go:embed breakpoint_dev3_alloc.json
	breakpointDev3AllocJSON []byte

	alloc = types.GenesisAlloc{
		contracts.CrossExAddr: {Code: common.Hex2Bytes(predeploy.CrossExCode)},
		contracts.BridgeAddr: {
			Balance: new(big.Int).Mul(big.NewInt(980_000_000), big.NewInt(1e18)),
			Code:    common.Hex2Bytes(predeploy.CrossBridgeCode),
			Storage: map[common.Hash]common.Hash{
				contracts.BridgeImplSlot: common.BytesToHash(contracts.BridgeImplAddr.Bytes()),
			},
		},
		contracts.BridgeImplAddr: {Code: common.Hex2Bytes(predeploy.CrossBridgeImplCode)},
		contracts.Multicall3Addr: {Code: common.Hex2Bytes(predeploy.Multicall3Code)},
	}

	// ##CROSS: predeploys singleton factory
	sfAlloc = types.GenesisAlloc{
		contracts.SingletonFactoryAddr:  {Nonce: 1, Code: common.Hex2Bytes(predeploy.SingletonFactoryCode)},
		contracts.SingletonFactory2Addr: {Nonce: 1, Code: common.Hex2Bytes(predeploy.SingletonFactoryCode)},
	}
	// ##

	ethAlloc = types.GenesisAlloc{
		params.BeaconRootsAddress:        {Nonce: 1, Code: params.BeaconRootsCode, Balance: common.Big0},
		params.HistoryStorageAddress:     {Nonce: 1, Code: params.HistoryStorageCode, Balance: common.Big0},
		params.WithdrawalQueueAddress:    {Nonce: 1, Code: params.WithdrawalQueueCode, Balance: common.Big0},
		params.ConsolidationQueueAddress: {Nonce: 1, Code: params.ConsolidationQueueCode, Balance: common.Big0},
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

	GenesisAllocCrossDev3 = mergeAlloc(types.GenesisAlloc{
		params.FoundationCrossDev3: {
			Balance: new(big.Int).Mul(big.NewInt(1_000_000_000_000), big.NewInt(1e18)),
		},
		contracts.CrossExAddr:    {Nonce: 1, Code: common.FromHex(breakpoint.CrossExCode)},
		contracts.BridgeImplAddr: alloc[contracts.BridgeImplAddr],
		contracts.BridgeAddr:     alloc[contracts.BridgeAddr],
		contracts.Multicall3Addr: alloc[contracts.Multicall3Addr],
	}, sfAlloc, ethAlloc, mustDecodeAlloc(breakpointDev3AllocJSON))

	GenesisAllocCrossDev = overrideAlloc(mergeAlloc(types.GenesisAlloc{
		params.FoundationCrossDev: {
			Balance: new(big.Int).Mul(big.NewInt(1_000_000_000_000), big.NewInt(1e18)),
		},
		contracts.CrossExAddr:    {Nonce: 1, Code: common.FromHex(breakpoint.CrossExCode)},
		contracts.BridgeImplAddr: alloc[contracts.BridgeImplAddr],
		contracts.BridgeAddr:     alloc[contracts.BridgeAddr],
		contracts.Multicall3Addr: alloc[contracts.Multicall3Addr],
	}, sfAlloc, ethAlloc, mustDecodeAlloc(breakpointDev3AllocJSON)), mustDecodeAlloc(breakpointDevValidatorSetAllocJSON))
)

func mustDecodeAlloc(data []byte) types.GenesisAlloc {
	var alloc types.GenesisAlloc
	if err := json.Unmarshal(data, &alloc); err != nil {
		panic(err)
	}
	return alloc
}

func mergeAlloc(base types.GenesisAlloc, extras ...types.GenesisAlloc) types.GenesisAlloc {
	for _, extra := range extras {
		for addr, account := range extra {
			if _, exists := base[addr]; exists {
				panic("duplicate genesis allocation: " + addr.Hex())
			}
			base[addr] = account
		}
	}
	return base
}

func overrideAlloc(base, overrides types.GenesisAlloc) types.GenesisAlloc {
	for addr, account := range overrides {
		if _, exists := base[addr]; !exists {
			panic("missing genesis allocation to override: " + addr.Hex())
		}
		base[addr] = account
	}
	return base
}
