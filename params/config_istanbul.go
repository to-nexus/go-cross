package params

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
)

// ##CROSS: istanbul

type IstanbulConfig struct {
	EpochLength              uint64           `json:"epochlength"`                       // Number of blocks that should pass before pending validator votes are reset
	BlockPeriodSeconds       uint64           `json:"blockperiodseconds"`                // Minimum time between two consecutive QBFT blocks’ timestamps in seconds
	EmptyBlockPeriodSeconds  uint64           `json:"emptyblockperiodseconds,omitempty"` // Minimum time between two consecutive QBFT a block and empty block’ timestamps in seconds
	RequestTimeoutSeconds    uint64           `json:"requesttimeoutseconds"`             // Minimum request timeout for each QBFT round in milliseconds
	ProposerPolicy           uint64           `json:"policy"`                            // The policy for proposer selection
	Validators               []common.Address `json:"validators"`                        // Validators list
	MaxRequestTimeoutSeconds *uint64          `json:"maxRequestTimeoutSeconds"`          // The max round time

	// ##CROSS: fee collection
	Beneficiary *common.Address `json:"beneficiary"`        // Cross Beneficiary address
	GasLimit    *uint64         `json:"gaslimit,omitempty"` // Block gas limit

	// ##CROSS: basefee
	ElasticityMultiplier     *uint64               `json:"elasticitymultiplier"`     // Bounds the maximum gas limit an EIP-1559 block may have.
	BaseFeeChangeDenominator *uint64               `json:"basefeechangedenominator"` // Bounds the amount the base fee can change between blocks.
	MaxBaseFee               *math.HexOrDecimal256 `json:"maxbasefee,omitempty"`     // MaxBaseFee
	MinBaseFee               *math.HexOrDecimal256 `json:"minbasefee,omitempty"`     // MinBaseFee
}

func (c IstanbulConfig) String() string {
	return "istanbul"
}

type Transition struct {
	Block                   *big.Int `json:"block"`
	EpochLength             uint64   `json:"epochlength,omitempty"`             // Number of blocks that should pass before pending validator votes are reset
	BlockPeriodSeconds      uint64   `json:"blockperiodseconds,omitempty"`      // Minimum time between two consecutive QBFT blocks’ timestamps in seconds
	EmptyBlockPeriodSeconds *uint64  `json:"emptyblockperiodseconds,omitempty"` // Minimum time between two consecutive QBFT a block and empty block’ timestamps in seconds
	RequestTimeoutSeconds   uint64   `json:"requesttimeoutseconds,omitempty"`   // Minimum request timeout for each QBFT round in milliseconds
	//	Validators               []common.Address `json:"validators,omitempty"`               // List of validators
	MaxRequestTimeoutSeconds *uint64 `json:"maxRequestTimeoutSeconds,omitempty"` // The max a timeout should be for a round change

	// ##CROSS: fee collection
	Beneficiary *common.Address `json:"Beneficiary,omitempty"` // Cross Beneficiary address
	GasLimit    *uint64         `json:"gaslimit,omitempty"`    // gas limit

	// ##CROSS: basefee
	ElasticityMultiplier     *uint64               `json:"elasticitymultiplier,omitempty"`     // Bounds the maximum gas limit an EIP-1559 block may have.
	BaseFeeChangeDenominator *uint64               `json:"basefeechangedenominator,omitempty"` // Bounds the amount the base fee can change between blocks.
	MaxBaseFee               *math.HexOrDecimal256 `json:"maxbasefee,omitempty"`               // MaxBaseFee
	MinBaseFee               *math.HexOrDecimal256 `json:"minbasefee,omitempty"`               // MinBaseFee
}

func (c *ChainConfig) IsIstanbulConsensus() bool {
	return c.Istanbul != nil
}

// GetTransitionValue iterates over the transitions and calls the callback for each transition
// that is less than or equal to the given block number.
func (c *ChainConfig) GetTransitionValue(num *big.Int, callback func(transition Transition)) {
	if c != nil && num != nil && c.Transitions != nil {
		for i := 0; i < len(c.Transitions) && c.Transitions[i].Block.Cmp(num) <= 0; i++ {
			callback(c.Transitions[i])
		}
	}
}

func (c *ChainConfig) GetEpochLength(num *big.Int) (epochLength uint64) {
	// ##CROSS: istanbul param
	if cfg := IstanbulConfigAt(num.Uint64()); cfg != nil && cfg.EpochLength != 0 {
		return cfg.EpochLength
	}
	// ##

	if c.Istanbul != nil {
		epochLength = c.Istanbul.EpochLength
	}
	c.GetTransitionValue(num, func(transition Transition) {
		if transition.EpochLength != 0 {
			epochLength = transition.EpochLength
		}
	})
	if epochLength == 0 {
		epochLength = 30000
	}
	return
}

func (c *ChainConfig) GetBlockPeriodSeconds(num *big.Int) (blockPeriod uint64) {
	// ##CROSS: istanbul param
	if cfg := IstanbulConfigAt(num.Uint64()); cfg != nil && cfg.BlockPeriodSeconds != 0 {
		return cfg.BlockPeriodSeconds
	}
	// ##

	if c.Istanbul != nil {
		blockPeriod = c.Istanbul.BlockPeriodSeconds
	}
	c.GetTransitionValue(num, func(transition Transition) {
		if transition.BlockPeriodSeconds != 0 {
			blockPeriod = transition.BlockPeriodSeconds
		}
	})
	if blockPeriod == 0 {
		blockPeriod = 1
	}
	return
}

func (c *ChainConfig) GetEmptyBlockPeriodSeconds(num *big.Int) (emptyBlockPeriods uint64) {
	// ##CROSS: istanbul param
	if cfg := IstanbulConfigAt(num.Uint64()); cfg != nil {
		return cfg.EmptyBlockPeriodSeconds
	}
	// ##

	if c.Istanbul != nil {
		emptyBlockPeriods = c.Istanbul.EmptyBlockPeriodSeconds
	}
	c.GetTransitionValue(num, func(transition Transition) {
		if transition.EmptyBlockPeriodSeconds != nil {
			emptyBlockPeriods = *transition.EmptyBlockPeriodSeconds
		}
	})
	return
}

func (c *ChainConfig) GetRequestTimeoutSeconds(num *big.Int) (requestTimeout uint64) {
	// ##CROSS: istanbul param
	if cfg := IstanbulConfigAt(num.Uint64()); cfg != nil && cfg.RequestTimeoutSeconds != 0 {
		return cfg.RequestTimeoutSeconds
	}
	// ##

	if c.Istanbul != nil {
		requestTimeout = c.Istanbul.RequestTimeoutSeconds
	}
	c.GetTransitionValue(num, func(transition Transition) {
		if transition.RequestTimeoutSeconds != 0 {
			requestTimeout = transition.RequestTimeoutSeconds
		}
	})
	if requestTimeout == 0 {
		requestTimeout = 10
	}
	return
}

func (c *ChainConfig) GetMaxRequestTimeoutSeconds(num *big.Int) (maxRequestTimeout uint64) {
	// ##CROSS: istanbul param
	if cfg := IstanbulConfigAt(num.Uint64()); cfg != nil && cfg.MaxRequestTimeoutSeconds != nil {
		return *cfg.MaxRequestTimeoutSeconds
	}
	// ##

	if c.Istanbul != nil && c.Istanbul.MaxRequestTimeoutSeconds != nil {
		maxRequestTimeout = *c.Istanbul.MaxRequestTimeoutSeconds
	}
	c.GetTransitionValue(num, func(transition Transition) {
		if transition.MaxRequestTimeoutSeconds != nil {
			maxRequestTimeout = *transition.MaxRequestTimeoutSeconds
		}
	})
	return
}

func (c *ChainConfig) GetProposerPolicy(num *big.Int) (policy uint64) {
	// ##CROSS: istanbul param
	if cfg := IstanbulConfigAt(num.Uint64()); cfg != nil {
		return cfg.ProposerPolicy
	}
	// ##

	if c.Istanbul != nil {
		policy = c.Istanbul.ProposerPolicy
	}
	return
}

// ##CROSS: fee collection
func (c *ChainConfig) GetBeneficiaryAddress(num *big.Int) (beneficiary *common.Address) {
	if c.Istanbul != nil && c.Istanbul.Beneficiary != nil {
		beneficiary = c.Istanbul.Beneficiary
	}
	c.GetTransitionValue(num, func(transition Transition) {
		if transition.Beneficiary != nil {
			beneficiary = transition.Beneficiary
		}
	})
	return
}

func (c *ChainConfig) GetGasLimit(num *big.Int) (gasLimit *uint64) {
	// ##CROSS: istanbul param
	if cfg := IstanbulConfigAt(num.Uint64()); cfg != nil && cfg.GasLimit != nil {
		return cfg.GasLimit
	}
	// ##

	if c.Istanbul != nil && c.Istanbul.GasLimit != nil {
		gasLimit = c.Istanbul.GasLimit
	}
	c.GetTransitionValue(num, func(transition Transition) {
		if transition.GasLimit != nil {
			gasLimit = transition.GasLimit
		}
	})
	return
}

// ##CROSS: basefee
func (c *ChainConfig) GetElasticityMultiplier(num *big.Int) (multiplier uint64) {
	// ##CROSS: istanbul param
	if cfg := IstanbulConfigAt(num.Uint64()); cfg != nil && cfg.ElasticityMultiplier != nil {
		return *cfg.ElasticityMultiplier
	}
	// ##

	if c.Istanbul != nil && c.Istanbul.ElasticityMultiplier != nil {
		multiplier = *c.Istanbul.ElasticityMultiplier
	}
	c.GetTransitionValue(num, func(transition Transition) {
		if transition.ElasticityMultiplier != nil {
			multiplier = *transition.ElasticityMultiplier
		}
	})
	if multiplier == 0 {
		multiplier = c.ElasticityMultiplier()
	}
	return
}

func (c *ChainConfig) GetBaseFeeChangeDenominator(num *big.Int) (denominator uint64) {
	// ##CROSS: istanbul param
	if cfg := IstanbulConfigAt(num.Uint64()); cfg != nil && cfg.BaseFeeChangeDenominator != nil {
		return *cfg.BaseFeeChangeDenominator
	}
	// ##

	if c.Istanbul != nil && c.Istanbul.BaseFeeChangeDenominator != nil {
		denominator = *c.Istanbul.BaseFeeChangeDenominator
	}
	c.GetTransitionValue(num, func(transition Transition) {
		if transition.BaseFeeChangeDenominator != nil {
			denominator = *transition.BaseFeeChangeDenominator
		}
	})
	if denominator == 0 {
		denominator = c.BaseFeeChangeDenominator()
	}
	return
}

func (c *ChainConfig) GetMaxBaseFee(num *big.Int) (max *big.Int) {
	// ##CROSS: istanbul param
	if cfg := IstanbulConfigAt(num.Uint64()); cfg != nil && cfg.MaxBaseFee != nil {
		return (*big.Int)(cfg.MaxBaseFee)
	}
	// ##

	if c.Istanbul != nil && c.Istanbul.MaxBaseFee != nil {
		max = (*big.Int)(c.Istanbul.MaxBaseFee)
	}

	c.GetTransitionValue(num, func(transition Transition) {
		if transition.MaxBaseFee != nil {
			max = (*big.Int)(transition.MaxBaseFee)
		}
	})
	return
}

func (c *ChainConfig) GetMinBaseFee(num *big.Int) (min *big.Int) {
	// ##CROSS: istanbul param
	if cfg := IstanbulConfigAt(num.Uint64()); cfg != nil && cfg.MinBaseFee != nil {
		return (*big.Int)(cfg.MinBaseFee)
	}
	// ##

	if c.Istanbul != nil && c.Istanbul.MinBaseFee != nil {
		min = (*big.Int)(c.Istanbul.MinBaseFee)
	}

	c.GetTransitionValue(num, func(transition Transition) {
		if transition.MinBaseFee != nil {
			min = (*big.Int)(transition.MinBaseFee)
		}
	})
	return
}

// ##

// ##CROSS: istanbul

type transitionIncompatibleErr struct {
	field string
}

func newTransitionIncompatibleErr(field string) *transitionIncompatibleErr {
	return &transitionIncompatibleErr{field: field}
}

func (e *transitionIncompatibleErr) Error() string {
	return fmt.Sprintf("transitions.%s data incompatible. %s historical data does not match", e.field, e.field)
}

// checks if changes to transitions proposed are compatible
// with already existing genesis data
func isTransitionsConfigCompatible(c1, c2 *ChainConfig, head *big.Int) (*big.Int, *big.Int, error) {
	if len(c1.Transitions) == 0 && len(c2.Transitions) == 0 {
		// maxCodeSizeConfig not used. return
		return big.NewInt(0), big.NewInt(0), nil
	}

	// existing config had Transitions and new one does not have the same return error
	if len(c1.Transitions) > 0 && len(c2.Transitions) == 0 {
		return head, head, fmt.Errorf("genesis file missing transitions information")
	}

	if len(c2.Transitions) > 0 && len(c1.Transitions) == 0 {
		return big.NewInt(0), big.NewInt(0), nil
	}

	// check the number of records below current head in both configs
	// if they do not match throw an error
	c1RecsBelowHead := 0
	for _, data := range c1.Transitions {
		if data.Block.Cmp(head) <= 0 {
			c1RecsBelowHead++
		} else {
			break
		}
	}

	c2RecsBelowHead := 0
	for _, data := range c2.Transitions {
		if data.Block.Cmp(head) <= 0 {
			c2RecsBelowHead++
		} else {
			break
		}
	}

	// if the count of past records is not matching return error
	if c1RecsBelowHead != c2RecsBelowHead {
		return head, head, errors.New("transitions data incompatible. updating transitions for past")
	}

	// validate that each past record is matching exactly. if not return error
	for i := 0; i < c1RecsBelowHead; i++ {
		same := c1.Transitions[i].Block.Cmp(c2.Transitions[i].Block) == 0
		if !same || c1.Transitions[i].BlockPeriodSeconds != c2.Transitions[i].BlockPeriodSeconds {
			return head, head, newTransitionIncompatibleErr("BlockPeriodSeconds")
		}
		if !same || c1.Transitions[i].RequestTimeoutSeconds != c2.Transitions[i].RequestTimeoutSeconds {
			return head, head, newTransitionIncompatibleErr("RequestTimeoutSeconds")
		}
		if !same || c1.Transitions[i].EpochLength != c2.Transitions[i].EpochLength {
			return head, head, newTransitionIncompatibleErr("EpochLength")
		}
	}

	return big.NewInt(0), big.NewInt(0), nil
}
