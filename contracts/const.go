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
	ValidatorSetAddr       = common.HexToAddress("0x0000000000000000000000000000000000001001")
	StakeHubAddr           = common.HexToAddress("0x0000000000000000000000000000000000001002")
	GovernorAddr           = common.HexToAddress("0x0000000000000000000000000000000000001005")
	GovernanceTokenAddr    = common.HexToAddress("0x0000000000000000000000000000000000001006")
	GovernanceTimelockAddr = common.HexToAddress("0x0000000000000000000000000000000000001007")
	GovernanceExecutorAddr = common.HexToAddress("0x0000000000000000000000000000000000001008")
	// ##
)
