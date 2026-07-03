package contracts

import (
	"fmt"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/breakpoint"
	"github.com/ethereum/go-ethereum/contracts/predeploy"
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
	// Initializable[INITIALIZABLE_STORAGE]._initialized = type(uint64).max
	initializableSlot := common.HexToHash("0xf0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00")
	initializedData := common.HexToHash("0x000000000000000000000000000000000000000000000000ffffffffffffffff")
	// ERC1967Proxy[IMPLEMENTATION] = <impl address>
	implementationSlot := common.HexToHash("0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc")

	upgrades[forks.Prague] = make(map[uint64]*Upgrade)
	upgrades[forks.Prague][0] = &Upgrade{
		UpgradeName: forks.Prague.String(),
		Configs: []*UpgradeConfig{
			{
				Name:         "HistoryStorage",
				ContractAddr: params.HistoryStorageAddress,
				Code:         params.HistoryStorageCode,
			},
		},
	}
	upgrades[forks.Breakpoint] = make(map[uint64]*Upgrade)
	upgrades[forks.Breakpoint][0] = &Upgrade{
		UpgradeName: forks.Breakpoint.String(),
		Configs: []*UpgradeConfig{
			{
				Name:         "CrossEx",
				Commit:       "4220afc0d36775ff52177f501dd1c165444d7f4c",
				ContractAddr: CrossExAddr,
				Code:         breakpoint.CrossExCode,
			},
			// ##CROSS: consensus system contract
			{
				Name:         "ValidatorSetImpl",
				Commit:       "5cbc608bd77ef186120308ec8219a016197efba2",
				ContractAddr: ValidatorSetImplAddr,
				Code:         breakpoint.ValidatorSetMetaData.BinRuntime,
				Storage: map[common.Hash]common.Hash{
					initializableSlot: initializedData,
				},
			},
			{
				Name:         "ValidatorSetProxy",
				Commit:       "5cbc608bd77ef186120308ec8219a016197efba2",
				ContractAddr: ValidatorSetAddr,
				Code:         predeploy.ERC1967ProxyCode,
				Storage: map[common.Hash]common.Hash{
					implementationSlot: common.BytesToHash(ValidatorSetImplAddr.Bytes()),
				},
				Init: func(config *params.ChainConfig, header *types.Header) ([]byte, error) {
					// PoSA configuration must be set in the Istanbul config
					if config.Istanbul == nil || config.Istanbul.PoSA == nil {
						return nil, fmt.Errorf("PoSA configuration is not set")
					}

					admin := config.Istanbul.PoSA.Admin
					if admin == (common.Address{}) {
						return nil, fmt.Errorf("PoSA admin is not set")
					}

					var (
						validators []common.Address
						signers    [][]byte // ##CROSS: bls seal
					)
					for _, validator := range config.Istanbul.PoSA.Validators {
						validators = append(validators, validator.Validator)
						signers = append(signers, types.BytesToBLSPublicKey(validator.Signer).Bytes()) // ##CROSS: bls seal
					}
					return breakpoint.NewValidatorSet().PackInitialize(admin, validators, signers), nil
				},
			},
			{
				Name:         "StakeHubImpl",
				Commit:       "5cbc608bd77ef186120308ec8219a016197efba2",
				ContractAddr: StakeHubImplAddr,
				Code:         breakpoint.StakeHubMetaData.BinRuntime,
				Storage: map[common.Hash]common.Hash{
					initializableSlot: initializedData,
				},
			},
			{
				Name:         "StakeHubProxy",
				Commit:       "5cbc608bd77ef186120308ec8219a016197efba2",
				ContractAddr: StakeHubAddr,
				Code:         predeploy.ERC1967ProxyCode,
				Storage: map[common.Hash]common.Hash{
					implementationSlot: common.BytesToHash(StakeHubImplAddr.Bytes()),
				},
				Init: func(config *params.ChainConfig, header *types.Header) ([]byte, error) {
					// PoSA configuration must be set in the Istanbul config
					if config.Istanbul == nil || config.Istanbul.PoSA == nil {
						return nil, fmt.Errorf("PoSA configuration is not set")
					}

					pool := config.Istanbul.PoSA.DelegationPool
					admin := config.Istanbul.PoSA.Admin
					if pool == (common.Address{}) || admin == (common.Address{}) {
						return nil, fmt.Errorf("delegation pool or PoSA admin is not set")
					}

					var (
						operators  []common.Address
						validators []common.Address
						signers    [][]byte // ##CROSS: bls seal
						ids        []string
					)
					for _, validator := range config.Istanbul.PoSA.Validators {
						operators = append(operators, validator.Operator)
						validators = append(validators, validator.Validator)
						ids = append(ids, validator.ID)
						// ##CROSS: bls seal
						signers = append(signers, validator.Signer)
					}
					return breakpoint.NewStakeHub().PackInitialize(pool, admin, operators, validators, signers, ids), nil
				},
			},
			{
				Name:         "RewardHubImpl",
				Commit:       "373dfdd9ec4b53f7f3efc0516a7b7cfaa7cab0f6",
				ContractAddr: RewardHubImplAddr,
				Code:         breakpoint.RewardHubMetaData.BinRuntime,
				Storage: map[common.Hash]common.Hash{
					initializableSlot: initializedData,
				},
			},
			{
				Name:         "RewardHubProxy",
				Commit:       "373dfdd9ec4b53f7f3efc0516a7b7cfaa7cab0f6",
				ContractAddr: RewardHubAddr,
				Code:         predeploy.ERC1967ProxyCode,
				Storage: map[common.Hash]common.Hash{
					implementationSlot: common.BytesToHash(RewardHubImplAddr.Bytes()),
				},
				Init: func(config *params.ChainConfig, header *types.Header) ([]byte, error) {
					// PoSA configuration must be set in the Istanbul config
					if config.Istanbul == nil || config.Istanbul.PoSA == nil {
						return nil, fmt.Errorf("PoSA configuration is not set")
					}

					pool := config.Istanbul.PoSA.DelegationPool
					admin := config.Istanbul.PoSA.Admin
					if pool == (common.Address{}) || admin == (common.Address{}) {
						return nil, fmt.Errorf("delegation pool or PoSA admin is not set")
					}
					startBlock := config.Istanbul.PoSA.RewardStartBlock
					if startBlock == nil {
						return nil, fmt.Errorf("reward start block is not set")
					}
					return breakpoint.NewRewardHub().PackInitialize(pool, admin, startBlock), nil
				},
			},
			{
				Name:         "ValidatorSlashImpl",
				Commit:       "5cbc608bd77ef186120308ec8219a016197efba2",
				ContractAddr: ValidatorSlashImplAddr,
				Code:         breakpoint.ValidatorSlashMetaData.BinRuntime,
				Storage: map[common.Hash]common.Hash{
					initializableSlot: initializedData,
				},
			},
			{
				Name:         "ValidatorSlashProxy",
				Commit:       "5cbc608bd77ef186120308ec8219a016197efba2",
				ContractAddr: ValidatorSlashAddr,
				Code:         predeploy.ERC1967ProxyCode,
				Storage: map[common.Hash]common.Hash{
					implementationSlot: common.BytesToHash(ValidatorSlashImplAddr.Bytes()),
				},
				Init: func(config *params.ChainConfig, header *types.Header) ([]byte, error) {
					// PoSA configuration must be set in the Istanbul config
					if config.Istanbul == nil || config.Istanbul.PoSA == nil {
						return nil, fmt.Errorf("PoSA configuration is not set")
					}

					admin := config.Istanbul.PoSA.Admin
					if admin == (common.Address{}) {
						return nil, fmt.Errorf("PoSA admin is not set")
					}
					return breakpoint.NewValidatorSlash().PackInitialize(admin), nil
				},
			},
			// ##
		},
	}
	// Contract patches for testnet
	upgrades[forks.BreakpointAlpha] = make(map[uint64]*Upgrade)
	upgrades[forks.BreakpointAlpha][params.ZoneZeroChainConfig.ChainID.Uint64()] = &Upgrade{
		UpgradeName: forks.BreakpointAlpha.String(),
		Configs: []*UpgradeConfig{
			// ##CROSS: consensus system contract
			{
				Name:         "ValidatorSetImpl",
				Commit:       "5cbc608bd77ef186120308ec8219a016197efba2",
				ContractAddr: ValidatorSetImplAddr,
				Code:         breakpoint.ValidatorSetMetaData.BinRuntime,
				Storage: map[common.Hash]common.Hash{
					initializableSlot: initializedData,
				},
			},
			{
				Name:         "ValidatorSetProxy",
				Commit:       "5cbc608bd77ef186120308ec8219a016197efba2",
				ContractAddr: ValidatorSetAddr,
				Code:         predeploy.ERC1967ProxyCode,
				Storage: map[common.Hash]common.Hash{
					implementationSlot: common.BytesToHash(ValidatorSetImplAddr.Bytes()),
					common.HexToHash("0x86f0e39a3c4d21deb4b9249c3b1ec8d10355c59608a22092a826f5aa57331300"): common.BytesToHash(params.ZoneZeroPoSAConfig.Admin.Bytes()),
				},
			},
			{
				Name:         "ValidatorSlashImpl",
				Commit:       "5cbc608bd77ef186120308ec8219a016197efba2",
				ContractAddr: ValidatorSlashImplAddr,
				Code:         breakpoint.ValidatorSlashMetaData.BinRuntime,
				Storage: map[common.Hash]common.Hash{
					initializableSlot: initializedData,
				},
			},
			{
				Name:         "ValidatorSlashProxy",
				Commit:       "5cbc608bd77ef186120308ec8219a016197efba2",
				ContractAddr: ValidatorSlashAddr,
				Code:         predeploy.ERC1967ProxyCode,
				Storage: map[common.Hash]common.Hash{
					implementationSlot:      common.BytesToHash(ValidatorSlashImplAddr.Bytes()),
					common.HexToHash("0x4"): common.BytesToHash(new(big.Int).SetUint64(4).Bytes()),
					common.HexToHash("0x5"): common.BytesToHash(new(big.Int).SetUint64(50).Bytes()),
					common.HexToHash("0x86f0e39a3c4d21deb4b9249c3b1ec8d10355c59608a22092a826f5aa57331300"): common.BytesToHash(params.ZoneZeroPoSAConfig.Admin.Bytes()),
				},
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
func TryUpdateSystemContract(config *params.ChainConfig, header *types.Header, lastBlockTime uint64, statedb vm.StateDB, writeLog bool) (upgraded bool) {
	if config == nil || header == nil || statedb == nil || reflect.ValueOf(statedb).IsNil() {
		return
	}

	chainID := config.ChainID.Uint64()
	if config.IsOnPrague(header.Number, lastBlockTime, header.Time) {
		applySystemContractUpgrade(getUpgrade(forks.Prague, chainID), header, statedb, writeLog)
		upgraded = true
	}
	if config.IsOnBreakpoint(header.Number, lastBlockTime, header.Time) && isPoSAConfigured(config) {
		applySystemContractUpgrade(getUpgrade(forks.Breakpoint, chainID), header, statedb, writeLog)
		upgraded = true
	}
	if config.IsOnBreakpointAlpha(header.Number, lastBlockTime, header.Time) {
		applySystemContractUpgrade(getUpgrade(forks.BreakpointAlpha, chainID), header, statedb, writeLog)
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
	if config.IsOnBreakpoint(header.Number, lastBlockTime, header.Time) && isPoSAConfigured(config) {
		initData = append(initData, getSystemContractInitialization(getUpgrade(forks.Breakpoint, chainID), config, header)...)
	}
	return
}

// isPoSAConfigured reports whether the Istanbul PoSA config is present. When it is
// not, the Breakpoint PoSA system-contract upgrade is skipped and the chain runs in
// pre-PoSA mode. ##CROSS: istanbul posa
func isPoSAConfigured(config *params.ChainConfig) bool {
	return config.Istanbul != nil && config.Istanbul.PoSA != nil
}

func applySystemContractUpgrade(upgrade *Upgrade, header *types.Header, statedb vm.StateDB, writeLog bool) {
	var logger log.Logger
	if writeLog {
		logger = log.Root()
	} else {
		logger = log.NewLogger(log.DiscardHandler())
	}

	if upgrade == nil {
		logger.Info("Empty upgrade config", "blockNumber", header.Number.Uint64())
		return
	}

	logger.Info("Upgrading built-in contracts", "name", upgrade.UpgradeName, "blockNumber", header.Number.Uint64())
	for _, cfg := range upgrade.Configs {
		deploy := !statedb.Exist(cfg.ContractAddr)
		if deploy {
			logger.Info("Deploy contract", "name", cfg.Name, "address", cfg.ContractAddr.String(), "commit", cfg.Commit)
		} else {
			logger.Info("Upgrade contract", "name", cfg.Name, "address", cfg.ContractAddr.String(), "commit", cfg.Commit)
		}

		// write code
		prevCode := statedb.GetCode(cfg.ContractAddr)
		newCode := parseCode(cfg)
		if len(newCode) > 0 {
			if len(prevCode) > 0 {
				logger.Warn("Overwriting existing code", "name", cfg.Name, "address", cfg.ContractAddr.String())
			}
			if deploy {
				// If it is the first deployment, set the nonce to 1
				statedb.SetNonce(cfg.ContractAddr, 1, tracing.NonceChangeNewContract)
			}
			statedb.SetCode(cfg.ContractAddr, newCode)
		}

		// write storage
		if len(cfg.Storage) > 0 {
			for k, v := range cfg.Storage {
				logger.Info("Writing storage slot", "name", cfg.Name, "address", cfg.ContractAddr.String(), "slot", k.String(), "value", v.String())
				statedb.SetState(cfg.ContractAddr, k, v)
			}
		}
	}
}

func parseCode(cfg *UpgradeConfig) (code []byte) {
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
