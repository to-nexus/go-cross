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
	Foundation *common.Address `json:"foundation"` // Cross Foundation address
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
	Foundation *common.Address `json:"Foundation,omitempty"` // foundation  address
	GasLimit   *uint64         `json:"gaslimit,omitempty"`   // gas limit
	// ##CROSS: basefee
	ElasticityMultiplier     *uint64               `json:"elasticitymultiplier,omitempty"`     // Bounds the maximum gas limit an EIP-1559 block may have.
	BaseFeeChangeDenominator *uint64               `json:"basefeechangedenominator,omitempty"` // Bounds the amount the base fee can change between blocks.
	MaxBaseFee               *math.HexOrDecimal256 `json:"maxbasefee,omitempty"`               // MaxBaseFee
	MinBaseFee               *math.HexOrDecimal256 `json:"minbasefee,omitempty"`               // MinBaseFee
}

// gets value at or after a transition
func (c *ChainConfig) GetTransitionValue(num *big.Int, callback func(transition Transition)) {
	if c != nil && num != nil && c.Transitions != nil {
		for i := 0; i < len(c.Transitions) && c.Transitions[i].Block.Cmp(num) <= 0; i++ {
			callback(c.Transitions[i])
		}
	}
}

func (c *ChainConfig) GetGasLimit(num *big.Int) (gasLimit *uint64) {
	c.GetTransitionValue(num, func(transition Transition) {
		if transition.GasLimit != nil {
			gasLimit = transition.GasLimit
		}
	})
	return
}

// ##CROSS: fee collection
func (c *ChainConfig) GetFoundationAddress(num *big.Int) (foundation *common.Address) {
	if c.Istanbul != nil && c.Istanbul.Foundation != nil {
		foundation = c.Istanbul.Foundation
	}

	c.GetTransitionValue(num, func(transition Transition) {
		if transition.Foundation != nil {
			foundation = transition.Foundation
		}
	})
	return
}

// ##CROSS: basefee
func (c *ChainConfig) GetElasticityMultiplier(num *big.Int) (multiplier uint64) {
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
