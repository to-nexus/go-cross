package breakpoint

import _ "embed"

// ##CROSS: contract upgrade

var (
	// ##CROSS: consensus system contract
	//go:embed validator_set
	ValidatorSetCode string
	//go:embed stake_hub
	StakeHubCode string
	// ##
)
