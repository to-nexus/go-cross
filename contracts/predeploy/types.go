package predeploy

import _ "embed"

// ##CROSS: contract upgrade

var (
	//go:embed cross_ex
	CrossExCode string
	//go:embed cross_bridge
	CrossBridgeCode string
	//go:embed cross_bridge_impl
	CrossBridgeImplCode string
	//go:embed multicall3
	Multicall3Code string
)
