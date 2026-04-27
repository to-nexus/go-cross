package contracts

import (
	"fmt"
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
	// fork -> chainID -> upgrade; chainID = 0 for default
	upgrades = make(map[forks.Fork]map[uint64]*Upgrade)
)

func init() {
	initialize := common.FromHex("8129fc1c")

	upgrades[forks.Prague] = make(map[uint64]*Upgrade)
	upgrades[forks.Prague][0] = &Upgrade{
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
	upgrades[forks.Breakpoint] = make(map[uint64]*Upgrade)
	upgrades[forks.Breakpoint][0] = &Upgrade{
		UpgradeName: forks.Breakpoint.String(),
		Configs: []*UpgradeConfig{
			{
				Name:         "CrossEx",
				ContractAddr: CrossExAddr,
				Code:         breakpoint.CrossExCode,
			},
			// ##CROSS: consensus system contract
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
					// ##CROSS: bls seal
					// browse Istanbul config to build signers list
					signers := make([][]byte, 0, len(extra.Validators))
					for i, validator := range extra.Validators {
						if config.Istanbul != nil {
							for idx := range config.Istanbul.Validators {
								if validator == config.Istanbul.Validators[idx] && idx < len(config.Istanbul.Signers) {
									signers = append(signers, config.Istanbul.Signers[idx])
									break
								}
							}
						}
						if len(signers) < i+1 {
							// this validator has no signer, write empty signer
							signers = append(signers, types.BLSPublicKey{}.Bytes())
						}
					}
					// ##
					return breakpoint.NewValidatorSet().PackUpdateValidators(extra.Validators, signers), nil
				},
			},
			{
				Name:         "StakeHub",
				ContractAddr: StakeHubAddr,
				Code:         breakpoint.StakeHubMetaData.BinRuntime,
				Deploy:       true,
				Init: func(config *params.ChainConfig, header *types.Header) ([]byte, error) {
					pool := config.Istanbul.DelegationPool
					admin := config.Istanbul.PoSAAdmin
					if pool == nil || admin == nil {
						return nil, fmt.Errorf("delegation pool or PoSA admin is not set")
					}
					return breakpoint.NewStakeHub().PackInitialize(*pool, *admin), nil
				},
			},
			{
				Name:         "RewardHub",
				ContractAddr: RewardHubAddr,
				Code:         breakpoint.RewardHubMetaData.BinRuntime,
				Deploy:       true,
				Init: func(config *params.ChainConfig, header *types.Header) ([]byte, error) {
					pool := config.Istanbul.DelegationPool
					admin := config.Istanbul.PoSAAdmin
					if pool == nil || admin == nil {
						return nil, fmt.Errorf("delegation pool or PoSA admin is not set")
					}
					startBlock := config.Istanbul.RewardStartBlock
					if startBlock == nil {
						return nil, fmt.Errorf("reward start block is not set")
					}
					return breakpoint.NewRewardHub().PackInitialize(*pool, *admin, startBlock), nil
				},
			},
			{
				Name:         "ValidatorSlash",
				ContractAddr: ValidatorSlashAddr,
				Code:         breakpoint.ValidatorSlashMetaData.BinRuntime,
				Deploy:       true,
				Init: func(config *params.ChainConfig, header *types.Header) ([]byte, error) {
					return initialize, nil
				},
			},
			{
				Name:         "CrossGovernor",
				ContractAddr: GovernorAddr,
				Code:         breakpoint.CrossGovernorMetaData.BinRuntime,
				Deploy:       true,
				Init: func(config *params.ChainConfig, header *types.Header) ([]byte, error) {
					admin := config.Istanbul.PoSAAdmin
					if admin == nil {
						return nil, fmt.Errorf("PoSA admin is not set")
					}
					return breakpoint.NewCrossGovernor().PackInitialize(*admin), nil
				},
			},
			{
				Name:         "GovernanceToken",
				ContractAddr: GovernanceTokenAddr,
				Code:         breakpoint.GovernanceTokenCode,
				Deploy:       true,
				Init: func(config *params.ChainConfig, header *types.Header) ([]byte, error) {
					return initialize, nil
				},
			},
			{
				Name:         "GovernanceTimelock",
				ContractAddr: GovernanceTimelockAddr,
				Code:         breakpoint.GovernanceTimelockCode,
				Deploy:       true,
				Init: func(config *params.ChainConfig, header *types.Header) ([]byte, error) {
					return initialize, nil
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
}

func getUpgrade(fork forks.Fork, chainID uint64) *Upgrade {
	upgrade, ok := upgrades[fork][chainID]
	if !ok {
		upgrade = upgrades[fork][0]
	}
	return upgrade
}

// TryUpdateSystemContract checks if the block is exactly on the fork and upgrades the system contracts if it is.
func TryUpdateSystemContract(config *params.ChainConfig, header *types.Header, lastBlockTime uint64, statedb vm.StateDB) (upgraded bool) {
	if config == nil || header == nil || statedb == nil || reflect.ValueOf(statedb).IsNil() {
		return
	}

	chainID := config.ChainID.Uint64()
	if config.IsOnPrague(header.Number, lastBlockTime, header.Time) {
		applySystemContractUpgrade(getUpgrade(forks.Prague, chainID), header, statedb)
		upgraded = true
	}
	if config.IsOnBreakpoint(header.Number, lastBlockTime, header.Time) {
		applySystemContractUpgrade(getUpgrade(forks.Breakpoint, chainID), header, statedb)
		upgraded = true
	}
	return
}

// InitSystemContract checks if the block is exactly on the fork and returns the initialization data for the system contracts if it is.
func InitSystemContract(config *params.ChainConfig, header *types.Header, lastBlockTime uint64) (initData []ContractInitData) {
	if config == nil || header == nil {
		return
	}

	chainID := config.ChainID.Uint64()
	if config.IsOnBreakpoint(header.Number, lastBlockTime, header.Time) {
		initData = append(initData, getSystemContractInitialization(getUpgrade(forks.Breakpoint, chainID), config, header)...)
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
		if len(code) == 0 {
			panic(fmt.Errorf("%s: failed to parse code: %s", upgrade.UpgradeName, cfg.Name))
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
