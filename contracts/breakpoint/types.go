package breakpoint

import _ "embed"

// ##CROSS: contract upgrade

var (
	//go:embed validator_set
	ValidatorSetCode string // ##CROSS: consensus system contract
)
