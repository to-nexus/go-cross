package predeploy

import (
	_ "embed"
	"encoding/json"
	"maps"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts"
	"github.com/ethereum/go-ethereum/contracts/predeploy"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

// ##CROSS: predeploys

var (
	// breakpointAllocJSON contains uninitialized Breakpoint contract code and proxy links.
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
			Balance: new(big.Int).Mul(big.NewInt(1_000_000_000_000), big.NewInt(1e18)),
		},
		contracts.CrossReserveAddr: {
			Balance: new(big.Int).Mul(big.NewInt(100_000_000), big.NewInt(1e18)),
		},
	}, crossAlloc, sfAlloc)

	GenesisAllocCrossDev = mergeAlloc(types.GenesisAlloc{
		params.FoundationCrossDev: {
			Balance: new(big.Int).Mul(big.NewInt(1_000_000_000_000), big.NewInt(1e18)),
		},
		contracts.CrossReserveAddr: {
			Balance: new(big.Int).Mul(big.NewInt(100_000_000), big.NewInt(1e18)),
		},
	}, crossAlloc, sfAlloc)
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
func BreakpointGenesisAlloc() types.GenesisAlloc {
	source := mustDecodeAlloc(breakpointAllocJSON)
	impls := []common.Address{
		contracts.ValidatorSetImplAddr,
		contracts.StakeHubImplAddr,
		contracts.RewardHubImplAddr,
		contracts.ValidatorSlashImplAddr,
		contracts.CrossReserveImplAddr,
		contracts.ProtocolReserveImplAddr,
		contracts.DelegationPoolImplAddr,
	}
	proxies := map[common.Address]common.Address{
		contracts.ValidatorSetAddr:    contracts.ValidatorSetImplAddr,
		contracts.StakeHubAddr:        contracts.StakeHubImplAddr,
		contracts.RewardHubAddr:       contracts.RewardHubImplAddr,
		contracts.ValidatorSlashAddr:  contracts.ValidatorSlashImplAddr,
		contracts.CrossReserveAddr:    contracts.CrossReserveImplAddr,
		contracts.ProtocolReserveAddr: contracts.ProtocolReserveImplAddr,
		contracts.DelegationPoolAddr:  contracts.DelegationPoolImplAddr,
	}

	result := make(types.GenesisAlloc, len(impls)+len(proxies))
	for _, addr := range impls {
		account := source[addr]
		result[addr] = types.Account{
			Balance: new(big.Int),
			Nonce:   account.Nonce,
			Code:    common.CopyBytes(account.Code),
			Storage: maps.Clone(account.Storage),
		}
	}
	for addr, impl := range proxies {
		account := source[addr]
		result[addr] = types.Account{
			Balance: new(big.Int),
			Nonce:   account.Nonce,
			Code:    common.CopyBytes(account.Code),
			Storage: map[common.Hash]common.Hash{
				contracts.ProxyImplSlot: common.BytesToHash(impl.Bytes()),
			},
		}
	}
	return result
}
