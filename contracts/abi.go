package contracts

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// ##CROSS: consensus system contract

func ValidatorSetABI() abi.ABI {
	v, _ := abi.JSON(strings.NewReader(validatorSetABI))
	return v
}

func StakeHubABI() abi.ABI {
	v, _ := abi.JSON(strings.NewReader(stakeHubABI))
	return v
}

const validatorSetABI = `[{"type":"function","name":"getValidators","inputs":[],"outputs":[{"name":"","type":"address[]","internalType":"address[]"}],"stateMutability":"view"},{"type":"function","name":"updateValidators","inputs":[{"name":"validatorAddrs","type":"address[]","internalType":"address[]"}],"outputs":[],"stateMutability":"nonpayable"}]`
const stakeHubABI = `[{"type":"function","name":"getValidatorCount","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getValidatorElectionInfo","inputs":[{"name":"offset","type":"uint256","internalType":"uint256"},{"name":"limit","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"validatorAddrs","type":"address[]","internalType":"address[]"},{"name":"amounts","type":"uint256[]","internalType":"uint256[]"},{"name":"totalLength","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getValidators","inputs":[{"name":"offset","type":"uint256","internalType":"uint256"},{"name":"limit","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"operators","type":"address[]","internalType":"address[]"},{"name":"totalLength","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"initialize","inputs":[{"name":"guard","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"minStakeAmount","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"validatorThreshold","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"}]`
