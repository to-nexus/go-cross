package erc8004

import _ "embed"

// ##CROSS: contract upgrade

var (
	// ##CROSS: erc8004
	//go:embed identity_registry
	IdentityRegistryCode string
	//go:embed reputation_registry
	ReputationRegistryCode string
	// ##
)
