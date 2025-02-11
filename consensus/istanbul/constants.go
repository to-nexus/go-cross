package istanbul

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

var (
	DefaultDifficulty = big.NewInt(1)
	EmptyBlockNonce   = types.BlockNonce{}
)

// ElasticityMultiplier bounds the maximum gas limit an EIP-1559 block may have.
func ElasticityMultiplier(num *big.Int) uint64 {
	return params.DefaultElasticityMultiplier
}
