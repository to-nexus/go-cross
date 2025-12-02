package breakpoint

import _ "embed"

// ##CROSS: contract upgrade

var (
	//go:embed cross_ex
	CrossExCode string
	//go:embed validator_share
	ValidatorShareCode string
	//go:embed system_reward
	SystemRewardCode string
	//go:embed cross_governor
	CrossGovernorCode string
	//go:embed governance_token
	GovernanceTokenCode string
	//go:embed governance_timelock
	GovernanceTimelockCode string
	//go:embed governance_executor
	GovernanceExecutorCode string
)
