package contracts

import "github.com/ethereum/go-ethereum/common"

// ##CROSS: contract upgrade

var (
	// predeploy
	BridgeImplSlot = common.HexToHash("0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc")

	CrossExAddr    = common.HexToAddress("0xfeed00000000000000000000000000000000C0DE")
	BridgeAddr     = common.HexToAddress("0xb81d6e000000000000000000000000000000C0de")
	BridgeImplAddr = common.HexToAddress("0xB81D6e000000000000000000000000000000AAaA")
	Multicall3Addr = common.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11")

	// breakpoint
	// ##CROSS: consensus system contract
	IstanbulParamAddr      = common.HexToAddress("0x0000000000000000000000000000000000001000")
	ValidatorSetAddr       = common.HexToAddress("0x0000000000000000000000000000000000001001")
	StakeHubAddr           = common.HexToAddress("0x0000000000000000000000000000000000001002")
	ValidatorShareAddr     = common.HexToAddress("0x0000000000000000000000000000000000001003")
	ValidatorSlashAddr     = common.HexToAddress("0x0000000000000000000000000000000000001004")
	SystemRewardAddr       = common.HexToAddress("0x0000000000000000000000000000000000001005")
	GovernorAddr           = common.HexToAddress("0x0000000000000000000000000000000000001010")
	GovernanceTokenAddr    = common.HexToAddress("0x0000000000000000000000000000000000001011")
	GovernanceTimelockAddr = common.HexToAddress("0x0000000000000000000000000000000000001012")
	GovernanceExecutorAddr = common.HexToAddress("0x0000000000000000000000000000000000001013")
	// ##
	// ##CROSS: erc8004
	ERC8004IdentityRegistryAddr          = common.HexToAddress("0x8004A169FB4a3325136EB29fA0ceB6D2e539a432")
	ERC8004ReputationRegistryAddr        = common.HexToAddress("0x8004BAa17C55a88189AE136b182e5fdA19dE9b63")
	ERC8004IdentityRegistryAddrTestnet   = common.HexToAddress("0x8004A818BFB912233c491871b3d84c89A494BD9e")
	ERC8004ReputationRegistryAddrTestnet = common.HexToAddress("0x8004B663056A597Dffe9eCcC1965A193B7388713")
	// ##
)

var systemContracts = map[common.Address]bool{
	IstanbulParamAddr:      true,
	ValidatorSetAddr:       true,
	StakeHubAddr:           true,
	ValidatorSlashAddr:     true,
	SystemRewardAddr:       true,
	GovernorAddr:           true,
	GovernanceTokenAddr:    true,
	GovernanceTimelockAddr: true,
	GovernanceExecutorAddr: true,
}

// IsSystemContract checks if the address is a system contract.
func IsSystemContract(addr common.Address) bool {
	return systemContracts[addr]
}
