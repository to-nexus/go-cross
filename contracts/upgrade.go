package contracts

import (
	"reflect"

	"github.com/ethereum/go-ethereum/common"
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
