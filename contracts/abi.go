package contracts

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func ValidatorSetABI() abi.ABI {
	v, _ := abi.JSON(strings.NewReader(validatorSetABI))
	return v
}

const validatorSetABI = `[{"type":"function","name":"getValidators","inputs":[],"outputs":[{"name":"","type":"address[]","internalType":"address[]"}],"stateMutability":"view"},{"type":"function","name":"updateValidators","inputs":[{"name":"validatorAddrs","type":"address[]","internalType":"address[]"}],"outputs":[],"stateMutability":"nonpayable"}]`
