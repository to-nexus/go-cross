package contracts

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
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
		ContractAddr common.Address
		Commit       string
		Code         string
		Storage      map[common.Hash]common.Hash
	}
	Upgrade struct {
		UpgradeName string
		Configs     []*UpgradeConfig
	}
)

var (
	upgrades = make(map[forks.Fork]*Upgrade)
)

func init() {
	upgrades[forks.Breakpoint] = &Upgrade{
		UpgradeName: forks.Breakpoint.String(),
		Configs:     []*UpgradeConfig{},
	}
}

// TryUpdateSystemContract checks if the block is exactly on the fork and upgrades the system contracts if it is.
func TryUpdateSystemContract(config *params.ChainConfig, blockNumber *big.Int, lastBlockTime uint64, blockTime uint64, statedb vm.StateDB, atBlockBegin bool) {
	if config == nil || blockNumber == nil || statedb == nil || reflect.ValueOf(statedb).IsNil() {
		return
	}

	if atBlockBegin {
		// HistoryStorage contract for Prague fork is predefined in params package
		if config.IsOnPrague(blockNumber, lastBlockTime, blockTime) {
			log.Info("Deploy contract", "name", "HistoryStorage", "address", params.HistoryStorageAddress.String(), "blockNumber", blockNumber.Uint64())
			statedb.SetCode(params.HistoryStorageAddress, params.HistoryStorageCode)
			statedb.SetNonce(params.HistoryStorageAddress, 1, tracing.NonceChangeNewContract)
		}
	} else {
		upgradeSystemContract(config, blockNumber, lastBlockTime, blockTime, statedb)
	}
}

func upgradeSystemContract(config *params.ChainConfig, blockNumber *big.Int, lastBlockTime uint64, blockTime uint64, statedb vm.StateDB) {
	if config.IsOnBreakpoint(blockNumber, lastBlockTime, blockTime) {
		applySystemContractUpgrade(upgrades[forks.Breakpoint], blockNumber, statedb)
	}
}

func applySystemContractUpgrade(upgrade *Upgrade, blockNumber *big.Int, statedb vm.StateDB) {
	if upgrade == nil {
		log.Info("Empty upgrade config", "blockNumber", blockNumber.Uint64())
		return
	}

	log.Info("Upgrading built-in contracts", "name", upgrade.UpgradeName, "blockNumber", blockNumber.Uint64())
	for _, cfg := range upgrade.Configs {
		if cfg.Deploy {
			log.Info("Deploy contract", "name", cfg.Name, "address", cfg.ContractAddr.String(), "commit", cfg.Commit)
		} else {
			log.Info("Upgrade contract", "name", cfg.Name, "address", cfg.ContractAddr.String(), "commit", cfg.Commit)
		}

		if len(cfg.Code) > 0 {
			code, err := hex.DecodeString(strings.TrimSpace(cfg.Code))
			if err != nil {
				panic(fmt.Errorf("failed to decode contract code: %s", err.Error()))
			}
			if len(code) > 0 {
				if cfg.Deploy {
					if prevCode := statedb.GetCode(cfg.ContractAddr); len(prevCode) > 0 {
						log.Warn("Contract code already exists", "address", cfg.ContractAddr.String())
					} else {
						// If it is the first deployment, we need to set the nonce to 1
						statedb.SetNonce(cfg.ContractAddr, 1, tracing.NonceChangeNewContract)
					}
				}
				statedb.SetCode(cfg.ContractAddr, code)
			}
		}
		for k, v := range cfg.Storage {
			statedb.SetState(cfg.ContractAddr, k, v)
		}
	}
}
