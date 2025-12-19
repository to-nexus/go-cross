package contracts

import (
	"fmt"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
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
		Init         func(*params.ChainConfig, *types.Header) ([]byte, error)
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
	// ##CROSS: bridge fix
	upgrades[forks.AdventureFix] = &Upgrade{
		UpgradeName: forks.AdventureFix.String(),
		Configs: []*UpgradeConfig{
			{
				Name:         "CrossBridge",
				ContractAddr: BridgeAddr,
				Storage: map[common.Hash]common.Hash{
					BridgeImplSlot: common.BytesToHash(BridgeImplAddr.Bytes()),
				},
			},
		},
	}
	// ##
	upgrades[forks.Prague] = &Upgrade{
		UpgradeName: forks.Prague.String(),
		Configs: []*UpgradeConfig{
			{
				Name:         "HistoryStorage",
				ContractAddr: params.HistoryStorageAddress,
				Code:         params.HistoryStorageCode,
				Deploy:       true,
			},
		},
	}
	// ##CROSS: fork breakpoint
	upgrades[forks.Breakpoint] = &Upgrade{
		UpgradeName: forks.Breakpoint.String(),
		Configs: []*UpgradeConfig{
			{
				Name:         "CrossEx",
				ContractAddr: CrossExAddr,
				Code:         breakpoint.CrossExCode,
			},
			// ##CROSS: consensus system contract
			{
				Name:         "IstanbulParam",
				ContractAddr: IstanbulParamAddr,
				Code:         breakpoint.IstanbulParamMetaData.BinRuntime,
				Deploy:       true,
				Init: func(config *params.ChainConfig, header *types.Header) ([]byte, error) {
					epochLength := config.GetEpochLength(header.Number)
					blockPeriod := config.GetBlockPeriodSeconds(header.Number)
					emptyBlockPeriod := config.GetEmptyBlockPeriodSeconds(header.Number)
					requestTimeout := config.GetRequestTimeoutSeconds(header.Number)
					maxRequestTimeout := config.GetMaxRequestTimeoutSeconds(header.Number)
					baseFeeChangeDenominator := config.GetBaseFeeChangeDenominator(header.Number)
					elasticityMultiplier := config.GetElasticityMultiplier(header.Number)
					maxBaseFee := config.GetMaxBaseFee(header.Number)
					if maxBaseFee == nil {
						maxBaseFee = big.NewInt(0)
					}
					minBaseFee := config.GetMinBaseFee(header.Number)
					if minBaseFee == nil {
						minBaseFee = big.NewInt(0)
					}
					proposerPolicy := config.GetProposerPolicy(header.Number)
					gasLimit := header.GasLimit
					if limit := config.GetGasLimit(header.Number); limit != nil {
						gasLimit = *limit
					}
					return breakpoint.NewIstanbulParam().PackInitialize(
						epochLength,
						blockPeriod,
						emptyBlockPeriod,
						requestTimeout,
						maxRequestTimeout,
						elasticityMultiplier,
						baseFeeChangeDenominator,
						maxBaseFee,
						minBaseFee,
						proposerPolicy,
						gasLimit,
					), nil
				},
			},
			{
				Name:         "ValidatorSet",
				ContractAddr: ValidatorSetAddr,
				Code:         breakpoint.ValidatorSetMetaData.BinRuntime,
				Deploy:       true,
				Init: func(config *params.ChainConfig, header *types.Header) ([]byte, error) {
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
				Init: func(config *params.ChainConfig, header *types.Header) ([]byte, error) {
					return breakpoint.NewStakeHub().PackInitialize(), nil
				},
			},
			{
				Name:         "ValidatorShare",
				ContractAddr: ValidatorShareAddr,
				Code:         breakpoint.ValidatorShareCode,
				Deploy:       true,
			},
			{
				Name:         "ValidatorSlash",
				ContractAddr: ValidatorSlashAddr,
				Code:         breakpoint.ValidatorSlashMetaData.BinRuntime,
				Deploy:       true,
				Init: func(config *params.ChainConfig, header *types.Header) ([]byte, error) {
					return breakpoint.NewValidatorSlash().PackInitialize(), nil
				},
			},
			{
				Name:         "SystemReward",
				ContractAddr: SystemRewardAddr,
				Code:         breakpoint.SystemRewardCode,
				Deploy:       true,
			},
			{
				Name:         "CrossGovernor",
				ContractAddr: GovernorAddr,
				Code:         breakpoint.CrossGovernorCode,
				Deploy:       true,
				Init: func(config *params.ChainConfig, header *types.Header) ([]byte, error) {
					return common.FromHex("8129fc1c"), nil
				},
			},
			{
				Name:         "GovernanceToken",
				ContractAddr: GovernanceTokenAddr,
				Code:         breakpoint.GovernanceTokenCode,
				Deploy:       true,
				Init: func(config *params.ChainConfig, header *types.Header) ([]byte, error) {
					return common.FromHex("8129fc1c"), nil
				},
			},
			{
				Name:         "GovernanceTimelock",
				ContractAddr: GovernanceTimelockAddr,
				Code:         breakpoint.GovernanceTimelockCode,
				Deploy:       true,
				Init: func(config *params.ChainConfig, header *types.Header) ([]byte, error) {
					return common.FromHex("8129fc1c"), nil
				},
			},
			{
				Name:         "GovernanceExecutor",
				ContractAddr: GovernanceExecutorAddr,
				Code:         breakpoint.GovernanceExecutorCode,
				Deploy:       true,
			},
			// ##
		},
	}
	// ##
}

// TryUpdateSystemContract checks if the block is exactly on the fork and upgrades the system contracts if it is.
func TryUpdateSystemContract(config *params.ChainConfig, header *types.Header, lastBlockTime uint64, statedb vm.StateDB) (upgraded bool) {
	if config == nil || header == nil || statedb == nil || reflect.ValueOf(statedb).IsNil() {
		return
	}

	// ##CROSS: bridge fix
	if config.IsOnAdventureFix(header.Number, lastBlockTime, header.Time) {
		applySystemContractUpgrade(upgrades[forks.AdventureFix], header, statedb)
		upgraded = true
	}
	// ##
	if config.IsOnPrague(header.Number, lastBlockTime, header.Time) {
		applySystemContractUpgrade(upgrades[forks.Prague], header, statedb)
		upgraded = true
	}
	// ##CROSS: fork breakpoint
	if config.IsOnBreakpoint(header.Number, lastBlockTime, header.Time) {
		applySystemContractUpgrade(upgrades[forks.Breakpoint], header, statedb)
		upgraded = true
	}
	// ##
	return
}

// InitSystemContract checks if the block is exactly on the fork and returns the initialization data for the system contracts if it is.
func InitSystemContract(config *params.ChainConfig, header *types.Header, lastBlockTime uint64) (initData []ContractInitData) {
	if config == nil || header == nil {
		return
	}

	if config.IsOnBreakpoint(header.Number, lastBlockTime, header.Time) {
		initData = append(initData, getSystemContractInitialization(upgrades[forks.Breakpoint], config, header)...)
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
		code := getCode(cfg)
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

func getCode(cfg *UpgradeConfig) (code []byte) {
	switch v := cfg.Code.(type) {
	case string:
		code = common.FromHex(v)
	case []byte:
		code = v
	}
	return
}

func getSystemContractInitialization(upgrade *Upgrade, config *params.ChainConfig, header *types.Header) (initData []ContractInitData) {
	if upgrade == nil {
		return
	}
	for _, cfg := range upgrade.Configs {
		if cfg.Init != nil {
			data, err := cfg.Init(config, header)
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
