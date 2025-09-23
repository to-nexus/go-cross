package contracts

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/breakpoint"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
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
		Prepare      func(header *types.Header, cfg *UpgradeConfig) error
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
	upgrades[forks.Prague] = &Upgrade{
		UpgradeName: forks.Prague.String(),
		Configs: []*UpgradeConfig{
			{
				Name:         "HistoryStorage",
				ContractAddr: params.HistoryStorageAddress,
				Deploy:       true,
				Prepare: func(header *types.Header, cfg *UpgradeConfig) error {
					cfg.Code = hex.EncodeToString(params.HistoryStorageCode)
					return nil
				},
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
				Code:         breakpoint.ValidatorSetCode,
				Deploy:       true,
				Prepare: func(header *types.Header, cfg *UpgradeConfig) error {
					// Prepare storage of the contract with current validators
					extra, err := types.ExtractIstanbulExtra(header)
					if err != nil {
						return err
					}
					if cfg.Storage == nil {
						cfg.Storage = make(map[common.Hash]common.Hash)
					}
					arrSlot, _ := new(big.Int).SetString("290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563", 16)
					mapBase := common.LeftPadBytes(common.Big1.Bytes(), 32)
					for i, addr := range extra.Validators {
						// validators[i] = addr
						cfg.Storage[common.BigToHash(arrSlot)] = common.BytesToHash(common.LeftPadBytes(addr.Bytes(), 32))
						arrSlot = arrSlot.Add(arrSlot, big.NewInt(20))

						// validatorsMap[addr] = i+1
						mapSlot := crypto.Keccak256(common.LeftPadBytes(addr.Bytes(), 32), mapBase)
						cfg.Storage[common.BytesToHash(mapSlot)] = common.BytesToHash(common.LeftPadBytes(big.NewInt(int64(i+1)).Bytes(), 32))
					}
					cfg.Storage[common.BigToHash(common.Big0)] = common.BigToHash(big.NewInt(int64(len(extra.Validators)))) // set length of 'validators'
					return nil
				},
			},
			// ##
		},
	}
}

// TryUpdateSystemContract checks if the block is exactly on the fork and upgrades the system contracts if it is.
func TryUpdateSystemContract(config *params.ChainConfig, header *types.Header, lastBlockTime uint64, statedb vm.StateDB) {
	if config == nil || header == nil || statedb == nil || reflect.ValueOf(statedb).IsNil() {
		return
	}

	if config.IsOnPrague(header.Number, lastBlockTime, header.Time) {
		applySystemContractUpgrade(upgrades[forks.Prague], header, statedb)
	}
	if config.IsOnBreakpoint(header.Number, lastBlockTime, header.Time) {
		applySystemContractUpgrade(upgrades[forks.Breakpoint], header, statedb)
	}
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

		if cfg.Prepare != nil {
			err := cfg.Prepare(header, cfg)
			if err != nil {
				panic(fmt.Errorf("failed to prepare: %s", err.Error()))
			}
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
