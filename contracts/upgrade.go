package contracts

import (
	"fmt"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/contracts/breakpoint"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/params/forks"
)

// ##CROSS: contract upgrade

type (
	UpgradeConfig struct {
		Deploy       bool
		Name         string
		Commit       string
		ContractAddr common.Address
		Code         any // string or []byte
		Storage      map[common.Hash]common.Hash
		Init         func(header *types.Header) ([]byte, error)
	}
	Upgrade struct {
		UpgradeName string
		Configs     []*UpgradeConfig
	}

	ContractInitData struct {
		To   common.Address
		Data []byte
	}
)

var (
	upgrades = make(map[forks.Fork]*Upgrade)
)

func init() {
	upgrades[forks.Prague] = &Upgrade{
		UpgradeName: forks.Prague.String(),
		Configs: []*UpgradeConfig{
			{
				Name:         "HistoryStorage",
				ContractAddr: params.HistoryStorageAddress,
				Deploy:       true,
				Code:         params.HistoryStorageCode,
			},
		},
	}
	upgrades[forks.Breakpoint] = &Upgrade{
		UpgradeName: forks.Breakpoint.String(),
		Configs: []*UpgradeConfig{
			// ##CROSS: consensus system contract
			{
				Name:         "ValidatorSet",
				ContractAddr: ValidatorSetAddr,
				Code:         breakpoint.ValidatorSetMetaData.BinRuntime,
				Deploy:       true,
				Init: func(header *types.Header) ([]byte, error) {
					extra, err := types.ExtractIstanbulExtra(header)
					if err != nil {
						return nil, err
					}
					return breakpoint.NewValidatorSet().PackUpdateValidators(extra.Validators), nil
				},
			},
			{
				Name:         "StakeHub",
				ContractAddr: StakeHubAddr,
				Code:         breakpoint.StakeHubMetaData.BinRuntime,
				Deploy:       true,
				Init: func(header *types.Header) ([]byte, error) {
					return breakpoint.NewStakeHub().PackInitialize(common.Address{}), nil
				},
			},
			{
				Name:         "GovernanceToken",
				ContractAddr: GovernanceTokenAddr,
				Code:         breakpoint.GovernanceTokenMetaData.BinRuntime,
				Deploy:       true,
				Init: func(header *types.Header) ([]byte, error) {
					return breakpoint.NewGovernanceToken().PackInitialize(), nil
				},
			},
			// ##
		},
	}
}

// TryUpdateSystemContract checks if the block is exactly on the fork and upgrades the system contracts if it is.
func TryUpdateSystemContract(config *params.ChainConfig, header *types.Header, lastBlockTime uint64, statedb vm.StateDB) (upgraded bool) {
	if config == nil || header == nil || statedb == nil || reflect.ValueOf(statedb).IsNil() {
		return
	}

	if config.IsOnPrague(header.Number, lastBlockTime, header.Time) {
		applySystemContractUpgrade(upgrades[forks.Prague], header, statedb)
		upgraded = true
	}
	if config.IsOnBreakpoint(header.Number, lastBlockTime, header.Time) {
		applySystemContractUpgrade(upgrades[forks.Breakpoint], header, statedb)
		upgraded = true
	}
	return
}

func InitSystemContract(config *params.ChainConfig, header *types.Header, lastBlockTime uint64) (initData []ContractInitData) {
	if config == nil || header == nil {
		return
	}

	if config.IsOnBreakpoint(header.Number, lastBlockTime, header.Time) {
		initData = append(initData, getSystemContractInitialization(upgrades[forks.Breakpoint], header)...)
	}
	return
}

func applySystemContractUpgrade(upgrade *Upgrade, header *types.Header, statedb vm.StateDB) {
	if upgrade == nil {
		log.Info("Empty upgrade config", "blockNumber", header.Number.Uint64())
		return
	}

	log.Info("Upgrading built-in contracts", "name", upgrade.UpgradeName, "blockNumber", header.Number.Uint64())
	for _, cfg := range upgrade.Configs {
		if cfg.Deploy {
			log.Info("Deploy contract", "name", cfg.Name, "address", cfg.ContractAddr.String(), "commit", cfg.Commit)
		} else {
			log.Info("Upgrade contract", "name", cfg.Name, "address", cfg.ContractAddr.String(), "commit", cfg.Commit)
		}

		exists := statedb.Exist(cfg.ContractAddr)

		// write code
		code, err := getCode(cfg)
		if err != nil {
			panic(fmt.Errorf("failed to parse code: %w", err))
		}

		if len(code) > 0 {
			if cfg.Deploy {
				if prevCode := statedb.GetCode(cfg.ContractAddr); len(prevCode) > 0 {
					log.Warn("Overwriting existing code", "name", cfg.Name, "address", cfg.ContractAddr.String())
				} else {
					// If it is the first deployment, set the nonce to 1
					statedb.SetNonce(cfg.ContractAddr, 1, tracing.NonceChangeNewContract)
				}
			} else if !exists {
				log.Warn("Writing code to non-existent account", "name", cfg.Name, "address", cfg.ContractAddr.String())
			}
			statedb.SetCode(cfg.ContractAddr, code)
		}

		// write storage
		if len(cfg.Storage) > 0 {
			if !exists {
				log.Warn("Writing storage slots of non-existent account", "name", cfg.Name, "address", cfg.ContractAddr.String())
			}
			for k, v := range cfg.Storage {
				statedb.SetState(cfg.ContractAddr, k, v)
			}
		}
	}
}

func getCode(cfg *UpgradeConfig) (code []byte, err error) {
	switch v := cfg.Code.(type) {
	case string:
		code, err = hexutil.Decode(v)
	case []byte:
		code = v
	}
	return
}

func getSystemContractInitialization(upgrade *Upgrade, header *types.Header) (initData []ContractInitData) {
	if upgrade == nil {
		return
	}
	for _, cfg := range upgrade.Configs {
		if cfg.Init != nil {
			data, err := cfg.Init(header)
			if err != nil {
				panic(fmt.Errorf("failed to get init data: %s", err.Error()))
			}
			initData = append(initData, ContractInitData{
				To:   cfg.ContractAddr,
				Data: data,
			})
		}
	}
	return
}
