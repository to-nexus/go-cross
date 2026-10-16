package predeploy

import (
	_ "embed"
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts"
	"github.com/ethereum/go-ethereum/contracts/latest"
	"github.com/ethereum/go-ethereum/contracts/predeploy"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

// ##CROSS: predeploys

var (
	// breakpointAllocJSON contains uninitialized PoSA reserve and pool contract code and proxy links.
	//go:embed breakpoint_alloc.json
	breakpointAllocJSON []byte // ##CROSS: fork breakpoint

	crossAlloc = types.GenesisAlloc{
		contracts.CrossExAddr: {Code: common.Hex2Bytes(predeploy.CrossExCode)},
		contracts.BridgeAddr: {
			Balance: new(big.Int).Mul(big.NewInt(980_000_000), big.NewInt(1e18)),
			Code:    common.Hex2Bytes(predeploy.CrossBridgeCode),
			Storage: map[common.Hash]common.Hash{
				contracts.ProxyImplSlot: common.BytesToHash(contracts.BridgeImplAddr.Bytes()),
			},
		},
		contracts.BridgeImplAddr: {Code: common.Hex2Bytes(predeploy.CrossBridgeImplCode)},
		contracts.Multicall3Addr: {Code: common.Hex2Bytes(predeploy.Multicall3Code)},
	}

	// ##CROSS: predeploys singleton factory
	//nolint:unused
	sfAlloc = types.GenesisAlloc{
		contracts.SingletonFactoryAddr:  {Nonce: 1, Code: common.Hex2Bytes(predeploy.SingletonFactoryCode)},
		contracts.SingletonFactory2Addr: {Nonce: 1, Code: common.Hex2Bytes(predeploy.SingletonFactoryCode)},
	}
	// ##

	GenesisAllocCross = mergeAlloc(types.GenesisAlloc{
		params.FoundationCross: {
			Balance: new(big.Int).Mul(big.NewInt(10_000_000), big.NewInt(1e18)),
		},
		params.EcoCross: {
			Balance: new(big.Int).Mul(big.NewInt(10_000_000), big.NewInt(1e18)),
		},
	}, crossAlloc)

	GenesisAllocZoneZero = mergeAlloc(types.GenesisAlloc{
		params.FoundationZoneZero: {
			Balance: new(big.Int).Mul(big.NewInt(1_000_000_000_000), big.NewInt(1e18)),
		},
	}, crossAlloc)

	GenesisAllocCrossDev3 = mergeAlloc(types.GenesisAlloc{
		params.FoundationCrossDev3: {
			Balance: new(big.Int).Mul(big.NewInt(100_000_000), big.NewInt(1e18)),
		},
	}, crossAlloc)

	GenesisAllocCrossDev = mergeAlloc(types.GenesisAlloc{
		params.FoundationCrossDev: {
			Balance: new(big.Int).Mul(big.NewInt(100_000_000), big.NewInt(1e18)),
		},
	}, crossAlloc)
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

// ##CROSS: fork breakpoint

// BreakpointGenesisAlloc returns uninitialized contracts used to build a custom Breakpoint genesis.
// The PoSA reserve and pool contracts come from the JSON, and the system contracts use the code in contracts/latest.
func BreakpointGenesisAlloc() types.GenesisAlloc {
	alloc := mustDecodeAlloc(breakpointAllocJSON)
	systemContracts := []struct {
		proxy, impl common.Address
		code        string
	}{
		{contracts.ValidatorSetAddr, contracts.ValidatorSetImplAddr, latest.ValidatorSetMetaData.BinRuntime},
		{contracts.StakeHubAddr, contracts.StakeHubImplAddr, latest.StakeHubMetaData.BinRuntime},
		{contracts.RewardHubAddr, contracts.RewardHubImplAddr, latest.RewardHubMetaData.BinRuntime},
		{contracts.ValidatorSlashAddr, contracts.ValidatorSlashImplAddr, latest.ValidatorSlashMetaData.BinRuntime},
	}
	for _, c := range systemContracts {
		// Genesis code never runs a constructor, so disable the implementation initializer here.
		alloc[c.impl] = types.Account{
			Balance: new(big.Int),
			Nonce:   1,
			Code:    common.FromHex(c.code),
			Storage: map[common.Hash]common.Hash{contracts.InitializableSlot: contracts.InitializersDisabled},
		}
		alloc[c.proxy] = types.Account{
			Balance: new(big.Int),
			Nonce:   1,
			Code:    common.FromHex(predeploy.ERC1967ProxyCode),
			Storage: map[common.Hash]common.Hash{contracts.ProxyImplSlot: common.BytesToHash(c.impl.Bytes())},
		}
	}
	return alloc
}
