// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package breakpoint

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// StakeHubMetaData contains all meta data concerning the StakeHub contract.
var StakeHubMetaData = bind.MetaData{
	ABI:        "[{\"type\":\"function\",\"name\":\"MAX_COMMISSION_RATE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SIGNER_ADDRESS_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPDATE_INTERVAL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addToBlackList\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyValidator\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"signerAddr\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"commissionRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"id\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"claimUndelegation\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimUndelegation\",\"inputs\":[{\"name\":\"id\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimUnstaked\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegate\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"delegate\",\"inputs\":[{\"name\":\"id\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"distributeReward\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getIds\",\"inputs\":[{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"ids\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"totalLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getKeeper\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperators\",\"inputs\":[{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"operatorAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"totalLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSignerAddress\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSignerAddress\",\"inputs\":[{\"name\":\"id\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorAddress\",\"inputs\":[{\"name\":\"id\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorAddress\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorDescription\",\"inputs\":[{\"name\":\"id_\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"commissionRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"id\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorDescription\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"commissionRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"id\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorInfo\",\"inputs\":[{\"name\":\"id\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"signerAddr\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"shareContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isJailed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isStaked\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorInfo\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"signerAddr\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"shareContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isJailed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isStaked\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorJailInfo\",\"inputs\":[{\"name\":\"id\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"jailUntil\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isJailed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorJailInfo\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"jailUntil\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isJailed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorTotalStake\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorTotalStake\",\"inputs\":[{\"name\":\"id\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorUnstakingInfo\",\"inputs\":[{\"name\":\"id\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"unstakeTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isUnstaking\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isUnstaked\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorUnstakingInfo\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"unstakeTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isUnstaking\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isUnstaked\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidators\",\"inputs\":[{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"validatorAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"signerAddrs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"totalLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"idToOperator\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isActiveValidator\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isActiveValidator\",\"inputs\":[{\"name\":\"id\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minDelegateAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minStakeAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"offlineJailDuration\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"offlineSlashAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeFromBlackList\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setKeeper\",\"inputs\":[{\"name\":\"keeper\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinDelegateAmount\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinStakeAmount\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOfflineJailDuration\",\"inputs\":[{\"name\":\"duration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOfflineSlashAmount\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUnbondingPeriod\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setValidatorThreshold\",\"inputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"signerToOperator\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"slashOffline\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalSlashedAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unbondingPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"undelegate\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"undelegate\",\"inputs\":[{\"name\":\"id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unjail\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unstakeValidator\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateCommissionRate\",\"inputs\":[{\"name\":\"newCommissionRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateSigner\",\"inputs\":[{\"name\":\"signerAddr\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateValidator\",\"inputs\":[{\"name\":\"newValidatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validatorThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatorToOperator\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"BlackListed\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Claimed\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CommissionRateUpdated\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newCommissionRate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Delegated\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"KeeperChanged\",\"inputs\":[{\"name\":\"previousKeeper\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newKeeper\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"KeeperUpdated\",\"inputs\":[{\"name\":\"keeper\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinDelegateAmountUpdated\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinStakeAmountUpdated\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OfflineJailDurationUpdated\",\"inputs\":[{\"name\":\"duration\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OfflineSlashAmountUpdated\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardDistributed\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignerUpdated\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"signerAddr\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnBlackListed\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnbondingPeriodUpdated\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Undelegated\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorApplied\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"shareContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"signerAddr\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorJailed\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"jailUntil\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorSlashed\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"slashAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorThresholdUpdated\",\"inputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorUnjailed\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorUnstaked\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"unstakeTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorUpdated\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newValidatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"DuplicateId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DuplicateOperatorAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DuplicateSignerAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DuplicateValidatorAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedDeployment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InBlackList\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidCommissionRate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignerAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValidatorAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValidatorStatus\",\"inputs\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumStakeHub.ValidatorStatus\"}]},{\"type\":\"error\",\"name\":\"InvalidValue\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"JailedValidator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughShares\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughValidators\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotJailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCoinbase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyKeeper\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySystemContract\",\"inputs\":[{\"name\":\"contractAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyZeroGasPrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"StillInJailPeriod\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TooFrequentUpdate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValidatorNotExists\",\"inputs\":[]}]",
	ID:         "StakeHub",
	BinRuntime: "0x60806040526004361015610011575f80fd5b5f3560e01c8062fa3d50146103f3578063092193ab146103ee5780630e222b36146103e9578063114eaf55146103e4578063176763aa146103df578063207239c0146103da5780632aa8e06d146103d557806335efc734146103d0578063391b6f4e146103cb5780633e732eba146103c65780633f4ba83a146103c157806340550a1c146103bc57806341698573146103b7578063417c73a7146103b257806342fce3cd146103ad57806343a769d9146103a857806344690e5c146103a35780634a49ac4c1461039e5780634c23c8fa146103995780634d99dd16146103945780634fd101d71461038f5780635bc0fbc71461038a5780635c19a95c146103855780635c975abb14610380578063635b27701461037b57806363973e4b146103765780636cf6d675146103715780636f7aa7ec1461036c5780637071688a14610367578063727ec27c1461036257806372e9c9341461035d578063748747e6146103585780638129fc1c146103535780638456cb591461034e5780638661213a14610349578063891239af1461034457806389ac41471461033f5780638a11d7c91461033a5780638c7d0dbc146103355780638dfc889714610330578063935e08d61461032b5780639679904a146103265780639ac83d85146103215780639ddb511a1461031c5780639e76f95814610317578063a43569b314610312578063a4cac9641461030d578063a9cc60f314610308578063ba723eda14610303578063bbe5e345146102fe578063bff02e20146102f9578063c470ed65146102f4578063d1379332146102ef578063d276290d146102ea578063d86c7fe8146102e5578063d8d50a2a146102e0578063e3f95ed9146102db578063ea4dd2b9146102d6578063eb4af045146102d1578063f1887684146102cc578063f26d717e146102c75763f679d305146102c2575f80fd5b612547565b612502565b6124e5565b612454565b612414565b6123ba565b6122dc565b6120b2565b611f12565b611ed2565b611e8d565b611de0565b611d25565b611d08565b611ceb565b611cc7565b611ca6565b611c8b565b611c36565b611c04565b611b6f565b611ae9565b611a69565b611a31565b611a0d565b6119f0565b6119a1565b611960565b6118d1565b61176b565b6116ac565b61168f565b6112b5565b611299565b611255565b611238565b6111e2565b61108a565b61105c565b61101d565b610f88565b610f6b565b610f24565b610ea6565b610da4565b610bd8565b610b42565b610b25565b610aab565b610a61565b6109d6565b61093b565b6108a7565b610873565b610833565b6107d0565b6107b4565b610766565b610681565b610609565b610565565b3461054b57602036600319011261054b5760043561040f6131d1565b60ff61041a336107fb565b541661053857610429336131f8565b6113888111610529576104ea90335f5260026020526104da60405f20916104a9600461047d60018601936104786104736001600160401b03875460a01c166001600160401b031690565b61323e565b613262565b940193849067ffffffffffffffff60401b82549160401b169067ffffffffffffffff60401b1916179055565b6104b242613262565b815467ffffffffffffffff60a01b191660a09190911b67ffffffffffffffff60a01b16179055565b5460401c6001600160401b031690565b6040516001600160401b03909116815233907f86d576c20e383fc2413ef692209cc48ddad5e52f25db5b32f8f7ec5076461ae99080602081015b0390a2005b63047f2a9760e51b5f5260045ffd5b63033788ff60e01b5f523360045260245ffd5b5f80fd5b600435906001600160a01b038216820361054b57565b602036600319011261054b5761057961054f565b41330361059f573a6105905761058e906125f7565b005b6383f1b1d360e01b5f5260045ffd5b63022d8c9560e31b5f5260045ffd5b9181601f8401121561054b578235916001600160401b03831161054b576020838186019501011161054b57565b602060031982011261054b57600435906001600160401b03821161054b57610605916004016105ae565b9091565b3461054b5761062161061a366105db565b3691610e70565b80516020918201205f90815260048083526040808320546001600160a01b031683526002909352919020908101546003909101546001600160401b03919091169060e81c60ff165b604080519283529015156020830152819081015b0390f35b3461054b57602036600319011261054b5760043561101333036107135780156106d5576020817f38a1644f59872db6fd17fdced395495fcaa3ca7d2825a0704db5a3acbd1dd06492600a55604051908152a1005b60849060405190632c648cf160e01b825260406004830152600f60448301526e1d5b989bdb991a5b99d4195c9a5bd9608a1b60648301526024820152fd5b630f22c43960e41b5f5261101360045260245ffd5b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b604090610763939281528160208201520190610728565b90565b3461054b5761077761061a366105db565b602081519101205f52600460205261079a60018060a01b0360405f205416612def565b9061067d6040519283928361074c565b5f91031261054b57565b3461054b575f36600319011261054b5760206040516113888152f35b3461054b57602036600319011261054b5760206107f36107ee61054f565b6127b8565b604051908152f35b6001600160a01b03165f9081527f46944d7356e409d0d8c174d262d4ab0ddb2d4597e3e424151a463d8a4af4e5016020526040902090565b3461054b57602036600319011261054b576001600160a01b0361085461054f565b165f526003602052602060018060a01b0360405f205416604051908152f35b3461054b575f36600319011261054b575f516020613bcd5f395f51905f52546040516001600160a01b039091168152602090f35b3461054b57602036600319011261054b5760043561101333036107135780156108fb576020817f02cd8ef316564ca78b75bf239c0a630008374c1fb1d26d941a6e9b19e42b2aa592600855604051908152a1005b60849060405190632c648cf160e01b82526040600483015260116044830152701b5a5b91195b1959d85d19505b5bdd5b9d607a1b60648301526024820152fd5b3461054b575f36600319011261054b575f516020613bcd5f395f51905f52546001600160a01b031633036109c757610971613311565b610979613311565b60ff195f516020613bed5f395f51905f5254165f516020613bed5f395f51905f52557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b63c60eb33560e01b5f5260045ffd5b3461054b57602036600319011261054b576001600160a01b036109f761054f565b165f5260026020526020610a0d60405f20613339565b6040519015158152f35b9491610a3f9060a0959298979498600180881b0316875260c0602088015260c0870190610728565b96600180861b0316604086015215156060850152151560808401521515910152565b3461054b57610a7261061a366105db565b602081519101205f52600460205261067d610a9860018060a01b0360405f205416612b3e565b9260409694969291925196879687610a17565b3461054b57602036600319011261054b57610ac461054f565b5f516020613bcd5f395f51905f52546001600160a01b031633036109c757610aeb816107fb565b805460ff191660011790556001600160a01b03167f7fd26be6fc92aff63f1f4409b2b2ddeb272a888031d7f55ec830485ec61941865f80a2005b3461054b575f36600319011261054b576020600c54604051908152f35b3461054b57602036600319011261054b576004356110133303610713578015610b96576020817fbdf592a42adc0296ad8a75a5e7b8a0912c37793a1a275d78275bc335911c6bd392600c55604051908152a1005b60849060405190632c648cf160e01b825260406004830152601360448301527237b3333634b732a530b4b6223ab930ba34b7b760691b60648301526024820152fd5b3461054b57602036600319011261054b57610bf161054f565b6110043303610d8f576001600160a01b03165f908152600360205260409020610c22905b546001600160a01b031690565b610c2b816131f8565b6001600160a01b0381165f9081526002602052604081206003810154602090610c9290610c6e90610c62906001600160a01b031681565b6001600160a01b031690565b600b5460405195868094819363045bc4d160e41b8352600483019190602083019252565b03925af18015610d5b57610d05925f91610d60575b50610cbc610cb782600654612849565b600655565b6040519081526001600160a01b038416907f17bddadfd7ec8898c3b9eadd0cf5ae77ba8d5df3a50e96ab86ec2dd711aa8fbb90602090a2610cff600c5442612849565b90613383565b6110113b1561054b57604051632961046560e21b81526001600160a01b0390911660048201525f8180602481015b0381836110115af18015610d5b57610d4757005b80610d555f61058e93610e2f565b806107aa565b6125ec565b610d82915060203d602011610d88575b610d7a8183610e2f565b8101906127a9565b5f610ca7565b503d610d70565b630f22c43960e41b5f5261100460045260245ffd5b3461054b57602036600319011261054b57610dbd61054f565b5f516020613bcd5f395f51905f52546001600160a01b031633036109c757610de4816107fb565b805460ff191690556001600160a01b03167fe0db3499b7fdc3da4cddff5f45d694549c19835e7f719fb5606d3ad1a5de40115f80a2005b634e487b7160e01b5f52604160045260245ffd5b90601f801991011681019081106001600160401b03821117610e5057604052565b610e1b565b6001600160401b038111610e5057601f01601f191660200190565b929192610e7c82610e55565b91610e8a6040519384610e2f565b82948184528183011161054b578281602093845f960137010152565b3461054b57602036600319011261054b576004356001600160401b03811161054b573660238201121561054b57610eea61067d913690602481600401359101610e70565b602060405191805191829101835e60059082019081526020908290038101909120546040516001600160a01b039091168152918291820190565b3461054b57604036600319011261054b57610f3d61054f565b602435610f486131d1565b60ff610f53336107fb565b54166105385781610f6661058e936131f8565b613488565b3461054b575f36600319011261054b576020600954604051908152f35b3461054b57602036600319011261054b576004356110133303610713578015610fdc576020817f324143af489ab0ba77b0d3580e2486ee612b673b87db2179c589f61ac532fd5992600955604051908152a1005b60849060405190632c648cf160e01b82526040600483015260126044830152711d985b1a59185d1bdc951a1c995cda1bdb1960721b60648301526024820152fd5b602036600319011261054b5761103161054f565b6110396131d1565b60ff611044336107fb565b5416610538578061105761058e926131f8565b6135e0565b3461054b575f36600319011261054b57602060ff5f516020613bed5f395f51905f5254166040519015158152f35b3461054b57602036600319011261054b576110a361054f565b6110ab6131d1565b60ff6110b6336107fb565b5416610538576110c5336131f8565b6001600160a01b038116156111d3576001600160a01b0381165f9081526003602052604090206110f890610c6290610c15565b6111c457335f908152600260205260409020611193909161115b6001840161113d61047361113183546001600160401b039060a01c1690565b6001600160401b031690565b84546001600160a01b0319166001600160a01b0384161785556104a9565b6001600160a01b03165f90815260036020526040902033905b80546001600160a01b0319166001600160a01b03909216919091179055565b546001600160a01b0316337fcfac5dc75b8d9a7e074162f59d9adcd33da59f0fe8dfb21580db298fc0fdad0d5f80a3005b63813f326760e01b5f5260045ffd5b63713ce51160e01b5f5260045ffd5b3461054b576111f361061a366105db565b602081519101205f52600460205261067d61121960018060a01b0360405f205416612e9c565b6040805193845291151560208401521515908201529081906060820190565b3461054b575f36600319011261054b576020600a54604051908152f35b3461054b57602036600319011261054b5761126e61054f565b6112766131d1565b60ff611281336107fb565b5416610538578061129461058e926131f8565b61371c565b3461054b575f36600319011261054b5760205f54604051908152f35b608036600319011261054b576112c961054f565b6024356001600160401b03811161054b576112e89036906004016105ae565b91604435906064356001600160401b03811161054b5761130c9036906004016105ae565b926113156131d1565b60ff611320336107fb565b5416610538576001600160a01b0383169384156111d35760075434106116805761134933610c62565b9161135f835f52600160205260405f2054151590565b611671576001600160a01b0385165f90815260036020526040902061138790610c6290610c15565b6111c45761139e611399368a8a610e70565b6137c3565b6113ae610c62610c158a8a61285b565b611662576113c56113c0368487610e70565b6137ee565b6113dd6113d3368487610e70565b6020815191012090565b916113f6610c62610c15855f52600460205260405f2090565b611653576113888211610529577f3f08ef9dcc0ca304c93d66df7d03be5aa3ff16f8991cb126e49161be0e19ae5d9561149561158e946114828b6111748e61145861145261144c6115a59e6115809b3691610e70565b336138c3565b9b613b34565b506001600160a01b0387165f90815260036020526040902061147b903390611174565b339261285b565b61117433915f52600460205260405f2090565b335f9081526002602052604090206114c69080546001600160a01b0319166001600160a01b03909316929092178255565b6001810180546001600160a01b0319163317815561155d906114ec8c8c600286016128fb565b6003830180546001600160a01b0319166001600160a01b038916178155946115499061151790613262565b936104a96004820195869067ffffffffffffffff60401b82549160401b169067ffffffffffffffff60401b1916179055565b805467ffffffffffffffff60a01b19169055565b825467ffffffffffffffff60a01b19168355805467ffffffffffffffff19169055565b805461ffff60e01b19169055565b60405191829160018060a01b0316963396836129e1565b0390a460408051348082526020820152339182917f24d7bda8602b916d64417f0dbfe2e2e88ec9b1157bd9f596dfdb91ba26624e049190a36110113b1561054b57604051632961046560e21b81523360048201525f81602481836110115af18015610d5b5761163f575b506110113b1561054b57604051632949582d60e01b8152336004820181905260248201525f818060448101610d33565b80610d555f61164d93610e2f565b5f61160f565b630b1fb6bd60e41b5f5260045ffd5b63966a8ad360e01b5f5260045ffd5b639cc20ced60e01b5f5260045ffd5b63e008b5f960e01b5f5260045ffd5b3461054b575f36600319011261054b576020600854604051908152f35b3461054b57602036600319011261054b576116c561054f565b6110133303610713576001600160a01b03168015611737575f516020613bcd5f395f51905f5254816001600160a01b0382167f068b48a2fe7f498b57ff6da64f075ae658fde8d77124b092e62b3dc58d91ce355f80a36001600160a01b031916175f516020613bcd5f395f51905f5255005b6084604051632c648cf160e01b815260406004820152600660448201526535b2b2b832b960d11b60648201525f6024820152fd5b3461054b575f36600319011261054b575f516020613c0d5f395f51905f52546001600160401b036117b460ff604084901c16156117a7565b1590565b926001600160401b031690565b16801590816118c9575b60011490816118bf575b1590816118b6575b506118a7578061180660016001600160401b03195f516020613c0d5f395f51905f525416175f516020613c0d5f395f51905f5255565b611872575b611813612a08565b61181957005b61184360ff60401b195f516020613c0d5f395f51905f5254165f516020613c0d5f395f51905f5255565b604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b6118a2600160401b60ff60401b195f516020613c0d5f395f51905f525416175f516020613c0d5f395f51905f5255565b61180b565b63f92ee8a960e01b5f5260045ffd5b9050155f6117d0565b303b1591506117c8565b8291506117be565b3461054b575f36600319011261054b575f516020613bcd5f395f51905f52546001600160a01b031633036109c7576119076131d1565b61190f6131d1565b600160ff195f516020613bed5f395f51905f525416175f516020613bed5f395f51905f52557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b3461054b5761197161061a366105db565b80516020918201205f908152600482526040808220546001600160a01b03168252600283529020610a0d90613339565b3461054b57602036600319011261054b576001600160a01b036119c261054f565b165f52600260205261067d6119dc600260405f2001612aa0565b604051918291602083526020830190610728565b3461054b575f36600319011261054b576020604051620151808152f35b3461054b57602036600319011261054b5761067d610a98611a2c61054f565b612b3e565b3461054b57611a4261061a366105db565b80516020918201205f9081526004825260409020546107f3906001600160a01b03166127b8565b3461054b57604036600319011261054b576004356001600160401b03811161054b57611a999036906004016105ae565b60243590611aa56131d1565b60ff611ab0336107fb565b54166105385761058e92611ac5913691610e70565b602081519101205f52600460205260018060a01b0360405f205416610f66816131f8565b3461054b57604036600319011261054b57611b08602435600435612cd9565b9060405190604082016040835281518091526060830190602060608260051b8601019301915f905b828210611b44578580868960208301520390f35b90919293602080611b61600193605f198a82030186528851610728565b960192019201909291611b30565b3461054b57602036600319011261054b576004356110133303610713578015611bc3576020817fff6e1245b60ea5b68ddca24a93cb8707677289cbdbdec9ecb7ccfc25541a774992600b55604051908152a1005b60849060405190632c648cf160e01b82526040600483015260126044830152711bd9999b1a5b9954db185cda105b5bdd5b9d60721b60648301526024820152fd5b3461054b57602036600319011261054b576004355f526004602052602060018060a01b0360405f205416604051908152f35b611c3f366105db565b90611c486131d1565b60ff611c53336107fb565b541661053857611c64913691610e70565b602081519101205f52600460205261058e60018060a01b0360405f205416611057816131f8565b3461054b575f36600319011261054b57602060405160308152f35b3461054b57602036600319011261054b5761079a611cc261054f565b612def565b3461054b57602036600319011261054b5761067d611219611ce661054f565b612e9c565b3461054b575f36600319011261054b576020600b54604051908152f35b3461054b575f36600319011261054b576020600654604051908152f35b3461054b57602036600319011261054b57610669611d4161054f565b60018060a01b03165f52600260205260405f209060ff60036001600160401b0360048501541693015460e81c1690565b90602080835192838152019201905f5b818110611d8e5750505090565b82516001600160a01b0316845260209384019390920191600101611d81565b90602080835192838152019201905f5b818110611dca5750505090565b8251845260209384019390920191600101611dbd565b3461054b57604036600319011261054b57611dff602435600435612f33565b90611e1860409493945193608085526080850190611d71565b8381036020850152845180825260208201916020808360051b8301019701925f915b838310611e6057878088611e568c8a8482036040860152611dad565b9060608301520390f35b9091929397602080611e7e600193601f198682030187528c51610728565b9a019301930191939290611e3a565b3461054b57611e9e61061a366105db565b80516020918201205f908152600482526040808220546001600160a01b039081168352600284529181902054905191168152f35b3461054b57602036600319011261054b576001600160a01b03611ef361054f565b165f526002602052602060018060a01b0360405f205416604051908152f35b3461054b575f36600319011261054b57611f2a6131d1565b60ff611f35336107fb565b541661053857611f44336131f8565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005c6120a35760017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005d335f908152600260205260409020600301805460e081901c60ff1690611fb3826129c3565b6001820361209057611fd09060a01c6001600160401b0316611131565b421061207a57815460ff60e01b1916600160e11b1782555f6020611fff610c628086546001600160a01b031690565b600460405180948193630f8f004960e31b83525af1908115610d5b575f9161205b575b50604051908152339081907ff7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd399268390602090a361058e6139ba565b612074915060203d602011610d8857610d7a8183610e2f565b81612022565b63170cc93360e21b5f5261208d906130a7565b5ffd5b63170cc93360e21b5f5261208d826130a7565b633ee5aeb560e01b5f5260045ffd5b3461054b575f36600319011261054b576120ca6131d1565b60ff6120d5336107fb565b5416610538576120e4336131f8565b335f90815260026020526040902060038101805460e081901c60ff16612109816129c3565b8061207a575060e81c60ff166122cd576121246117a36139df565b6122be57805460ff60e01b1916600160e01b17815561217361214b610478600a5442612849565b825467ffffffffffffffff60a01b191660a09190911b67ffffffffffffffff60a01b16178255565b8154612193906001600160a01b0316915460a01c6001600160401b031690565b6040516001600160401b0391909116815233916001600160a01b0316907fc88e6a643abb00b7da809fc581edda961c3ea06e7b32c1bce7e7671f37a77eed90602090a36110113b1561054b57604051632961046560e21b81523360048201525f81602481836110115af18015610d5b576122aa575b506110113b1561054b57604051632949582d60e01b81523360048201525f6024820181905290919082604481836110115af1918215610d5b5761225892612296575b50546001600160a01b031690565b6110013b1561054b5760405163f2888dbb60e01b81526001600160a01b039190911660048201525f81602481836110015af18015610d5b57610d4757005b80610d555f6122a493610e2f565b5f61224a565b80610d555f6122b893610e2f565b5f612208565b6315caeb5160e31b5f5260045ffd5b633c59519b60e11b5f5260045ffd5b3461054b576122ea366105db565b906122f36131d1565b60ff6122fe336107fb565b54166105385761230d336131f8565b61231b611399368484610e70565b6040518282823760058184019081528190036020019020546001600160a01b031661166257335f90815260026020526040902061238a9190611174909361147b6002600187019661237d6104736111318a546001600160401b039060a01c1690565b01956104a98386896128fb565b7f3b8c9c0f9c86a8c4815de4d82b5220138d828ca5c72db704d3809e43118c0c4d604051806105243394826130b9565b3461054b576123c8366105db565b906123d16131d1565b60ff6123dc336107fb565b5416610538576123ed913691610e70565b602081519101205f52600460205261058e60018060a01b0360405f205416611294816131f8565b3461054b57604036600319011261054b5761244a612436602435600435613149565b604051928392604084526040840190611d71565b9060208301520390f35b3461054b57602036600319011261054b5760043561101333036107135780156124a8576020817f8448c02797b448f4946bc25b3bf925e5556d1df822c944da701c54bab8a3162f92600755604051908152a1005b60849060405190632c648cf160e01b825260406004830152600e60448301526d1b5a5b94dd185ad9505b5bdd5b9d60921b60648301526024820152fd5b3461054b575f36600319011261054b576020600754604051908152f35b3461054b5761251361061a366105db565b602081519101205f52600460205260018060a01b0360405f2054165f52600260205261067d6119dc600260405f2001612aa0565b3461054b575f36600319011261054b5761255f6131d1565b60ff61256a336107fb565b541661053857612579336131f8565b335f908152600260205260409020600381015460e81c60ff16156125dd576125ae61113160048301546001600160401b031690565b42106125ce576125bd81613a68565b600754116116805761058e90613af2565b63065f425760e21b5f5260045ffd5b63f9200d4f60e01b5f5260045ffd5b6040513d5f823e3d90fd5b6001600160a01b03165f90815260036020526040902061261690610c15565b6001600160a01b0381165f9081526002602052604090206003810180546001600160a01b03808216161590811561279b575b50612758575461267b9060049061266990610c62906001600160a01b031681565b92015460401c6001600160401b031690565b90803b1561054b57604051632f303ebb60e11b81526001600160401b039290921660048301525f908290602490829034905af18015610d5b57612744575b506040513481526001600160a01b038216907fe34918ff1c7084970068b53fd71ad6d8b04e9f15d3886cbf006443e6cdc52ea690602090a26110113b1561054b57604051632961046560e21b81526001600160a01b0390911660048201525f8180602481015b0381836110115af18015610d5b576127345750565b80610d555f61274293610e2f565b565b80610d555f61275293610e2f565b5f6126b9565b50505061276434613293565b604051348152611005907fe34918ff1c7084970068b53fd71ad6d8b04e9f15d3886cbf006443e6cdc52ea69080602081015b0390a2565b60e81c60ff1690505f612648565b9081602091031261054b575190565b6001600160a01b039081165f90815260026020526040902060030154168015612830576004906020906127f3906001600160a01b0316610c62565b6040516278744560e21b815292839182905afa908115610d5b575f91612817575090565b610763915060203d602011610d8857610d7a8183610e2f565b505f90565b634e487b7160e01b5f52601160045260245ffd5b9190820180921161285657565b612835565b6020908260405193849283378101600581520301902090565b90600182811c921680156128a2575b602083101461288e57565b634e487b7160e01b5f52602260045260245ffd5b91607f1691612883565b601f82116128b957505050565b5f5260205f20906020601f840160051c830193106128f1575b601f0160051c01905b8181106128e6575050565b5f81556001016128db565b90915081906128d2565b9092916001600160401b038111610e50576129208161291a8454612874565b846128ac565b5f601f821160011461295e57819061294f9394955f92612953575b50508160011b915f199060031b1c19161790565b9055565b013590505f8061293b565b601f19821694612971845f5260205f2090565b915f5b8781106129ab575083600195969710612992575b505050811b019055565b01355f19600384901b60f8161c191690555f8080612988565b90926020600181928686013581550194019101612974565b600311156129cd57565b634e487b7160e01b5f52602160045260245ffd5b90918060409360208452816020850152848401375f828201840152601f01601f1916010190565b41330361059f573a6105905769021e19e0c9bab2400000600755678ac7230489e80000600855601560095562015180600a5568056bc75e2d63100000600b5562015180600c55612a56613ba1565b612a5e613ba1565b5f516020613bcd5f395f51905f5280546001600160a01b03191673e4dd1087233043dca7532c4fb9d8a31fe99b731e179055612a98613ba1565b612742613ba1565b9060405191825f825492612ab384612874565b8084529360018116908115612b1c5750600114612ad8575b5061274292500383610e2f565b90505f9291925260205f20905f915b818310612b00575050906020612742928201015f612acb565b6020919350806001915483858901015201910190918492612ae7565b90506020925061274294915060ff191682840152151560051b8201015f612acb565b60018060a01b03165f52600260205260405f2080549060018060a01b03821692612b6a60028301612aa0565b92600383015491612b8360018060a01b03841694613339565b926001600160401b0360ff8260e81c169360a01c1615159081612ba4575090565b60ff915060e01c166117a3816129c3565b6001600160401b038111610e505760051b60200190565b5f5b828110612bda57505050565b606082820152602001612bce565b60405190612bf7602083610e2f565b5f8083526127429060208401612bcc565b90612742612c1583612bb5565b612c226040519182610e2f565b83815260208194612c35601f1991612bb5565b019101612bcc565b9190820391821161285657565b60208183031261054b578051906001600160401b03821161054b570181601f8201121561054b57805190612c7d82610e55565b92612c8b6040519485610e2f565b8284526020838301011161054b57815f9260208093018386015e8301015290565b634e487b7160e01b5f52603260045260245ffd5b8051821015612cd45760209160051b010190565b612cac565b5f54929183821015612de15780612dd85750612cf6835b82612849565b81848211612dd0575b612d0891612c3d565b612d1181612c08565b915f5b828110612d22575050509190565b805f612d6d610c62610c626003612d5f612d46612d416004998b612849565b61397f565b6001600160a01b03165f90815260026020526040902090565b01546001600160a01b031690565b6040516395d89b4160e01b815293849182905afa8015610d5b576001925f91612dae575b50612d9c8287612cc0565b52612da78186612cc0565b5001612d14565b612dca91503d805f833e612dc28183610e2f565b810190612c4a565b5f612d91565b849150612cff565b612cf690612cf0565b5050612deb612be8565b9190565b6001600160a01b03165f90815260026020526040902060038101546001600160a01b031691908215612e84575f612e4b610c62612e3e6111316004809601546001600160401b039060401c1690565b956001600160a01b031690565b6040516395d89b4160e01b815292839182905afa908115610d5b575f91612e70575090565b61076391503d805f833e612dc28183610e2f565b505f9150604051612e96602082610e2f565b5f815290565b60018060a01b03165f526002602052600360405f2001549060ff6001600160401b038360a01c169260e01c1690612ed2826129c3565b60026001831492612ee2816129c3565b1490565b60405190612ef5602083610e2f565b5f808352366020840137565b90612f0b82612bb5565b612f186040519182610e2f565b8281528092612f29601f1991612bb5565b0190602036910137565b5f549182821015613084578061307b5750612f4e8282612849565b81838211613073575b612f6091612c3d565b91612f6a83612f01565b93612f7484612c08565b93612f7e81612f01565b935f5b828110612f8d57505050565b612f9d612d46612d418385612849565b8054909190612fc8906001600160a01b0316612fb9838c612cc0565b6001600160a01b039091169052565b612fd460028301612aa0565b612fde828a612cc0565b52612fe98189612cc0565b50612ff382613339565b15613069576020613015610c62610c626003600496015460018060a01b031690565b6040516278744560e21b815293849182905afa8015610d5b576001925f9161304b575b505b6130448289612cc0565b5201612f81565b613063915060203d8111610d8857610d7a8183610e2f565b5f613038565b600191505f61303a565b839150612f57565b612f4e90612cf0565b505061308e612ee6565b90613097612be8565b926130a0612ee6565b9293929190565b9060249160038110156129cd57600452565b90602082525f918154916130cc83612874565b928360208401526001811690815f1461312857506001146130ee575b50505090565b5f90815260208120929350915b83831061311157506040925001015f80806130e8565b8054604084840101526020909201916001016130fb565b9150506040935060ff929192191683830152151560051b01015f80806130e8565b5f549291838210156131c757806131c25750825b8101808211612856578381116131bb575b8181039081116128565761318181612f01565b915f5b828110613192575050509190565b80820190818311612856576131b56131ab60019361397f565b612fb98388612cc0565b01613184565b508261316e565b61315d565b5050612deb612ee6565b60ff5f516020613bed5f395f51905f5254166131e957565b63d93c066560e01b5f5260045ffd5b6001600160a01b03168015908115613221575b5061321257565b635926e0c360e01b5f5260045ffd5b61323791505f52600160205260405f2054151590565b155f61320b565b62015180810180911161285657421061325357565b63190bb86b60e31b5f5260045ffd5b6001600160401b03811161327c576001600160401b031690565b6306dfcc6560e41b5f52604060045260245260445ffd5b8047106132fb575f808080936110055af13d156132f3573d906132b582610e55565b916132c36040519384610e2f565b82523d5f602084013e5b156132d55750565b8051156132e457602081519101fd5b63d6bda27560e01b5f5260045ffd5b6060906132cd565b4763cf47918160e01b5f5260045260245260445ffd5b60ff5f516020613bed5f395f51905f5254161561332a57565b638dfc202b60e01b5f5260045ffd5b600381015460ff8160e01c1660038110156129cd5715918261336c575b5081613360575090565b60ff915060e81c161590565b613377919250613a68565b6007541115905f613356565b6004810161339b61113182546001600160401b031690565b831161345b575b5060038101805460ff60e81b1916600160e81b17905580546001600160a01b0316906110013b1561054b57604051634de5f52960e11b81526001600160a01b039290921660048301525f82602481836110015af1908115610d5b577feb7d7a49847ec491969db21a0e31b234565a9923145a2d1b56a75c9e958258029261279692613447575b50600101546040519384526001600160a01b0316929081906020820190565b80610d555f61345593610e2f565b5f613428565b6134829061346884613262565b6001600160401b03166001600160401b0319825416179055565b5f6133a2565b9080156135d1576001600160a01b0382165f908152600260205260409020600381018054919290916134c490610c62906001600160a01b031681565b604051635d043b2960e11b815260048101839052336024820181905260448201529190602090839060649082905f905af1908115610d5b5761354f925f926135b0575b506040805192835260208301919091526001600160a01b03861692339284927f3aace7340547de7b9156593a7652dc07ee900cea3fd8f82cb6c9d38b40829802928291820190565b0390a3331461358d575b50506110113b1561054b57604051632961046560e21b81526001600160a01b0390911660048201525f81806024810161271f565b5460e81c60ff166122cd576135a190613a68565b60075411611680575f80613559565b6135ca91925060203d602011610d8857610d7a8183610e2f565b905f613507565b633c57b48560e21b5f5260045ffd5b6008543410611680576001600160a01b0381165f9081526002602052604090206003015460e081901c60ff16613615816129c3565b8061207a575060e881901c60ff1680613709575b6122cd5761364190610c62906001600160a01b031681565b604051636e553f6560e01b815234600482018190523360248301529091602091839160449183915af1908115610d5b575f916136ea575b5060408051348152602081019290925233916001600160a01b038416917f24d7bda8602b916d64417f0dbfe2e2e88ec9b1157bd9f596dfdb91ba26624e0491a36110113b1561054b57604051632961046560e21b81526001600160a01b0390911660048201525f81806024810161271f565b613703915060203d602011610d8857610d7a8183610e2f565b5f613678565b50336001600160a01b0383161415613629565b6001600160a01b038181165f9081526002602052604090206003015416604051630f41a04d60e11b815233600482015290602090829060249082905f905af1908115610d5b575f916137a4575b5060405190815233916001600160a01b0316907ff7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd399268390602090a3565b6137bd915060203d602011610d8857610d7a8183610e2f565b5f613769565b51602f19016137ce57565b634501a91960e01b5f5260045ffd5b908151811015612cd4570160200190565b8051600381109081156138b8575b5061385e575f5b81518110156138b45761383961383361382d61381f84866137dd565b516001600160f81b03191690565b60f81c90565b60ff1690565b603081109081156138a9575b8161388c575b8161386d575b5061385e57600101613803565b631bf4348160e31b5f5260045ffd5b6061811091508115613881575b505f613851565b607a9150115f61387a565b905060418110801561389f575b9061384b565b50605a8111613899565b603981119150613845565b5050565b60099150115f6137fc565b90763d602d80600a3d3981f3363d3d373d3d3d363d730000005f527010035af43d82803e903d91602b57fd5bf3602052603760095ff0916001600160a01b03831691821561397057823b1561054b57604080516379ccf11760e11b81526001600160a01b039093166004840152602483015290915f918391908290819061394e906044830190610728565b039134905af18015610d5b57613962575090565b80610d555f61076393610e2f565b63b06ebf3d60e01b5f5260045ffd5b5f54811015612cd4575f80527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56301546001600160a01b031690565b5f7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005d565b5f80549081815b8381106139f557505050505f90565b81811015612cd4577f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5638101546001600160a01b03165f908152600260205260409020613a4090613339565b613a5e575b60018311613a55576001016139e6565b50505050600190565b9160010191613a45565b60038101546001600160a01b03168015613aec57613ad791602091613aab90600190613a9c906001600160a01b0316610c62565b9201546001600160a01b031690565b60405163ce96cb7760e01b81526001600160a01b03909116600482015292839190829081906024820190565b03915afa908115610d5b575f91612817575090565b50505f90565b60038101805460ff60e81b19169055600101546001600160a01b03167f9390b453426557da5ebdc31f19a37753ca04addf656d32f35232211bb2af3f195f80a2565b805f52600160205260405f2054155f14612830575f54600160401b811015610e5057600181015f555f54811015612cd4575f80527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563018190555f54905f52600160205260405f2055600190565b60ff5f516020613c0d5f395f51905f525460401c1615613bbd57565b631afcd79f60e31b5f5260045ffdfe46944d7356e409d0d8c174d262d4ab0ddb2d4597e3e424151a463d8a4af4e500cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a26469706673582212206d282c3d0d0cc6289127c77df06d2c24bad98b5f438521c09f146eb9884df19364736f6c634300081e0033",
}

// StakeHub is an auto generated Go binding around an Ethereum contract.
type StakeHub struct {
	abi abi.ABI
}

// NewStakeHub creates a new instance of StakeHub.
func NewStakeHub() *StakeHub {
	parsed, err := StakeHubMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &StakeHub{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *StakeHub) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackMAXCOMMISSIONRATE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x207239c0.
//
// Solidity: function MAX_COMMISSION_RATE() view returns(uint256)
func (stakeHub *StakeHub) PackMAXCOMMISSIONRATE() []byte {
	enc, err := stakeHub.abi.Pack("MAX_COMMISSION_RATE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMAXCOMMISSIONRATE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x207239c0.
//
// Solidity: function MAX_COMMISSION_RATE() view returns(uint256)
func (stakeHub *StakeHub) UnpackMAXCOMMISSIONRATE(data []byte) (*big.Int, error) {
	out, err := stakeHub.abi.Unpack("MAX_COMMISSION_RATE", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackSIGNERADDRESSLENGTH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9e76f958.
//
// Solidity: function SIGNER_ADDRESS_LENGTH() view returns(uint256)
func (stakeHub *StakeHub) PackSIGNERADDRESSLENGTH() []byte {
	enc, err := stakeHub.abi.Pack("SIGNER_ADDRESS_LENGTH")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSIGNERADDRESSLENGTH is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9e76f958.
//
// Solidity: function SIGNER_ADDRESS_LENGTH() view returns(uint256)
func (stakeHub *StakeHub) UnpackSIGNERADDRESSLENGTH(data []byte) (*big.Int, error) {
	out, err := stakeHub.abi.Unpack("SIGNER_ADDRESS_LENGTH", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackUPDATEINTERVAL is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x89ac4147.
//
// Solidity: function UPDATE_INTERVAL() view returns(uint256)
func (stakeHub *StakeHub) PackUPDATEINTERVAL() []byte {
	enc, err := stakeHub.abi.Pack("UPDATE_INTERVAL")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUPDATEINTERVAL is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x89ac4147.
//
// Solidity: function UPDATE_INTERVAL() view returns(uint256)
func (stakeHub *StakeHub) UnpackUPDATEINTERVAL(data []byte) (*big.Int, error) {
	out, err := stakeHub.abi.Unpack("UPDATE_INTERVAL", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackAddToBlackList is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x417c73a7.
//
// Solidity: function addToBlackList(address account) returns()
func (stakeHub *StakeHub) PackAddToBlackList(account common.Address) []byte {
	enc, err := stakeHub.abi.Pack("addToBlackList", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackApplyValidator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x727ec27c.
//
// Solidity: function applyValidator(address validatorAddr, bytes signerAddr, uint256 commissionRate, string id) payable returns()
func (stakeHub *StakeHub) PackApplyValidator(validatorAddr common.Address, signerAddr []byte, commissionRate *big.Int, id string) []byte {
	enc, err := stakeHub.abi.Pack("applyValidator", validatorAddr, signerAddr, commissionRate, id)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackClaimUndelegation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6f7aa7ec.
//
// Solidity: function claimUndelegation(address operatorAddr) returns()
func (stakeHub *StakeHub) PackClaimUndelegation(operatorAddr common.Address) []byte {
	enc, err := stakeHub.abi.Pack("claimUndelegation", operatorAddr)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackClaimUndelegation0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe3f95ed9.
//
// Solidity: function claimUndelegation(string id) returns()
func (stakeHub *StakeHub) PackClaimUndelegation0(id string) []byte {
	enc, err := stakeHub.abi.Pack("claimUndelegation0", id)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackClaimUnstaked is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd276290d.
//
// Solidity: function claimUnstaked() returns()
func (stakeHub *StakeHub) PackClaimUnstaked() []byte {
	enc, err := stakeHub.abi.Pack("claimUnstaked")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDelegate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c19a95c.
//
// Solidity: function delegate(address operatorAddr) payable returns()
func (stakeHub *StakeHub) PackDelegate(operatorAddr common.Address) []byte {
	enc, err := stakeHub.abi.Pack("delegate", operatorAddr)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDelegate0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9ddb511a.
//
// Solidity: function delegate(string id) payable returns()
func (stakeHub *StakeHub) PackDelegate0(id string) []byte {
	enc, err := stakeHub.abi.Pack("delegate0", id)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDistributeReward is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x092193ab.
//
// Solidity: function distributeReward(address validatorAddr) payable returns()
func (stakeHub *StakeHub) PackDistributeReward(validatorAddr common.Address) []byte {
	enc, err := stakeHub.abi.Pack("distributeReward", validatorAddr)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackGetIds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x935e08d6.
//
// Solidity: function getIds(uint256 offset, uint256 limit) view returns(string[] ids, uint256 totalLength)
func (stakeHub *StakeHub) PackGetIds(offset *big.Int, limit *big.Int) []byte {
	enc, err := stakeHub.abi.Pack("getIds", offset, limit)
	if err != nil {
		panic(err)
	}
	return enc
}

// GetIdsOutput serves as a container for the return parameters of contract
// method GetIds.
type GetIdsOutput struct {
	Ids         []string
	TotalLength *big.Int
}

// UnpackGetIds is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x935e08d6.
//
// Solidity: function getIds(uint256 offset, uint256 limit) view returns(string[] ids, uint256 totalLength)
func (stakeHub *StakeHub) UnpackGetIds(data []byte) (GetIdsOutput, error) {
	out, err := stakeHub.abi.Unpack("getIds", data)
	outstruct := new(GetIdsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Ids = *abi.ConvertType(out[0], new([]string)).(*[]string)
	outstruct.TotalLength = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, err

}

// PackGetKeeper is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x391b6f4e.
//
// Solidity: function getKeeper() view returns(address)
func (stakeHub *StakeHub) PackGetKeeper() []byte {
	enc, err := stakeHub.abi.Pack("getKeeper")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetKeeper is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x391b6f4e.
//
// Solidity: function getKeeper() view returns(address)
func (stakeHub *StakeHub) UnpackGetKeeper(data []byte) (common.Address, error) {
	out, err := stakeHub.abi.Unpack("getKeeper", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetOperators is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xea4dd2b9.
//
// Solidity: function getOperators(uint256 offset, uint256 limit) view returns(address[] operatorAddrs, uint256 totalLength)
func (stakeHub *StakeHub) PackGetOperators(offset *big.Int, limit *big.Int) []byte {
	enc, err := stakeHub.abi.Pack("getOperators", offset, limit)
	if err != nil {
		panic(err)
	}
	return enc
}

// GetOperatorsOutput serves as a container for the return parameters of contract
// method GetOperators.
type GetOperatorsOutput struct {
	OperatorAddrs []common.Address
	TotalLength   *big.Int
}

// UnpackGetOperators is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xea4dd2b9.
//
// Solidity: function getOperators(uint256 offset, uint256 limit) view returns(address[] operatorAddrs, uint256 totalLength)
func (stakeHub *StakeHub) UnpackGetOperators(data []byte) (GetOperatorsOutput, error) {
	out, err := stakeHub.abi.Unpack("getOperators", data)
	outstruct := new(GetOperatorsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.OperatorAddrs = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.TotalLength = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, err

}

// PackGetSignerAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x891239af.
//
// Solidity: function getSignerAddress(address operatorAddr) view returns(bytes)
func (stakeHub *StakeHub) PackGetSignerAddress(operatorAddr common.Address) []byte {
	enc, err := stakeHub.abi.Pack("getSignerAddress", operatorAddr)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetSignerAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x891239af.
//
// Solidity: function getSignerAddress(address operatorAddr) view returns(bytes)
func (stakeHub *StakeHub) UnpackGetSignerAddress(data []byte) ([]byte, error) {
	out, err := stakeHub.abi.Unpack("getSignerAddress", data)
	if err != nil {
		return *new([]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	return out0, err
}

// PackGetSignerAddress0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf26d717e.
//
// Solidity: function getSignerAddress(string id) view returns(bytes)
func (stakeHub *StakeHub) PackGetSignerAddress0(id string) []byte {
	enc, err := stakeHub.abi.Pack("getSignerAddress0", id)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetSignerAddress0 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf26d717e.
//
// Solidity: function getSignerAddress(string id) view returns(bytes)
func (stakeHub *StakeHub) UnpackGetSignerAddress0(data []byte) ([]byte, error) {
	out, err := stakeHub.abi.Unpack("getSignerAddress0", data)
	if err != nil {
		return *new([]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	return out0, err
}

// PackGetValidatorAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc470ed65.
//
// Solidity: function getValidatorAddress(string id) view returns(address)
func (stakeHub *StakeHub) PackGetValidatorAddress(id string) []byte {
	enc, err := stakeHub.abi.Pack("getValidatorAddress", id)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetValidatorAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc470ed65.
//
// Solidity: function getValidatorAddress(string id) view returns(address)
func (stakeHub *StakeHub) UnpackGetValidatorAddress(data []byte) (common.Address, error) {
	out, err := stakeHub.abi.Unpack("getValidatorAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetValidatorAddress0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd1379332.
//
// Solidity: function getValidatorAddress(address operatorAddr) view returns(address)
func (stakeHub *StakeHub) PackGetValidatorAddress0(operatorAddr common.Address) []byte {
	enc, err := stakeHub.abi.Pack("getValidatorAddress0", operatorAddr)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetValidatorAddress0 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd1379332.
//
// Solidity: function getValidatorAddress(address operatorAddr) view returns(address)
func (stakeHub *StakeHub) UnpackGetValidatorAddress0(data []byte) (common.Address, error) {
	out, err := stakeHub.abi.Unpack("getValidatorAddress0", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetValidatorCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7071688a.
//
// Solidity: function getValidatorCount() view returns(uint256)
func (stakeHub *StakeHub) PackGetValidatorCount() []byte {
	enc, err := stakeHub.abi.Pack("getValidatorCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetValidatorCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7071688a.
//
// Solidity: function getValidatorCount() view returns(uint256)
func (stakeHub *StakeHub) UnpackGetValidatorCount(data []byte) (*big.Int, error) {
	out, err := stakeHub.abi.Unpack("getValidatorCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetValidatorDescription is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x176763aa.
//
// Solidity: function getValidatorDescription(string id_) view returns(uint256 commissionRate, string id)
func (stakeHub *StakeHub) PackGetValidatorDescription(id string) []byte {
	enc, err := stakeHub.abi.Pack("getValidatorDescription", id)
	if err != nil {
		panic(err)
	}
	return enc
}

// GetValidatorDescriptionOutput serves as a container for the return parameters of contract
// method GetValidatorDescription.
type GetValidatorDescriptionOutput struct {
	CommissionRate *big.Int
	Id             string
}

// UnpackGetValidatorDescription is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x176763aa.
//
// Solidity: function getValidatorDescription(string id_) view returns(uint256 commissionRate, string id)
func (stakeHub *StakeHub) UnpackGetValidatorDescription(data []byte) (GetValidatorDescriptionOutput, error) {
	out, err := stakeHub.abi.Unpack("getValidatorDescription", data)
	outstruct := new(GetValidatorDescriptionOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.CommissionRate = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.Id = *abi.ConvertType(out[1], new(string)).(*string)
	return *outstruct, err

}

// PackGetValidatorDescription0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa43569b3.
//
// Solidity: function getValidatorDescription(address operatorAddr) view returns(uint256 commissionRate, string id)
func (stakeHub *StakeHub) PackGetValidatorDescription0(operatorAddr common.Address) []byte {
	enc, err := stakeHub.abi.Pack("getValidatorDescription0", operatorAddr)
	if err != nil {
		panic(err)
	}
	return enc
}

// GetValidatorDescription0Output serves as a container for the return parameters of contract
// method GetValidatorDescription0.
type GetValidatorDescription0Output struct {
	CommissionRate *big.Int
	Id             string
}

// UnpackGetValidatorDescription0 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa43569b3.
//
// Solidity: function getValidatorDescription(address operatorAddr) view returns(uint256 commissionRate, string id)
func (stakeHub *StakeHub) UnpackGetValidatorDescription0(data []byte) (GetValidatorDescription0Output, error) {
	out, err := stakeHub.abi.Unpack("getValidatorDescription0", data)
	outstruct := new(GetValidatorDescription0Output)
	if err != nil {
		return *outstruct, err
	}
	outstruct.CommissionRate = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.Id = *abi.ConvertType(out[1], new(string)).(*string)
	return *outstruct, err

}

// PackGetValidatorInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x41698573.
//
// Solidity: function getValidatorInfo(string id) view returns(address validatorAddr, bytes signerAddr, address shareContract, bool isActive, bool isJailed, bool isStaked)
func (stakeHub *StakeHub) PackGetValidatorInfo(id string) []byte {
	enc, err := stakeHub.abi.Pack("getValidatorInfo", id)
	if err != nil {
		panic(err)
	}
	return enc
}

// GetValidatorInfoOutput serves as a container for the return parameters of contract
// method GetValidatorInfo.
type GetValidatorInfoOutput struct {
	ValidatorAddr common.Address
	SignerAddr    []byte
	ShareContract common.Address
	IsActive      bool
	IsJailed      bool
	IsStaked      bool
}

// UnpackGetValidatorInfo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x41698573.
//
// Solidity: function getValidatorInfo(string id) view returns(address validatorAddr, bytes signerAddr, address shareContract, bool isActive, bool isJailed, bool isStaked)
func (stakeHub *StakeHub) UnpackGetValidatorInfo(data []byte) (GetValidatorInfoOutput, error) {
	out, err := stakeHub.abi.Unpack("getValidatorInfo", data)
	outstruct := new(GetValidatorInfoOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.ValidatorAddr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.SignerAddr = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.ShareContract = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.IsActive = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.IsJailed = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.IsStaked = *abi.ConvertType(out[5], new(bool)).(*bool)
	return *outstruct, err

}

// PackGetValidatorInfo0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8a11d7c9.
//
// Solidity: function getValidatorInfo(address operatorAddr) view returns(address validatorAddr, bytes signerAddr, address shareContract, bool isActive, bool isJailed, bool isStaked)
func (stakeHub *StakeHub) PackGetValidatorInfo0(operatorAddr common.Address) []byte {
	enc, err := stakeHub.abi.Pack("getValidatorInfo0", operatorAddr)
	if err != nil {
		panic(err)
	}
	return enc
}

// GetValidatorInfo0Output serves as a container for the return parameters of contract
// method GetValidatorInfo0.
type GetValidatorInfo0Output struct {
	ValidatorAddr common.Address
	SignerAddr    []byte
	ShareContract common.Address
	IsActive      bool
	IsJailed      bool
	IsStaked      bool
}

// UnpackGetValidatorInfo0 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8a11d7c9.
//
// Solidity: function getValidatorInfo(address operatorAddr) view returns(address validatorAddr, bytes signerAddr, address shareContract, bool isActive, bool isJailed, bool isStaked)
func (stakeHub *StakeHub) UnpackGetValidatorInfo0(data []byte) (GetValidatorInfo0Output, error) {
	out, err := stakeHub.abi.Unpack("getValidatorInfo0", data)
	outstruct := new(GetValidatorInfo0Output)
	if err != nil {
		return *outstruct, err
	}
	outstruct.ValidatorAddr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.SignerAddr = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.ShareContract = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.IsActive = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.IsJailed = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.IsStaked = *abi.ConvertType(out[5], new(bool)).(*bool)
	return *outstruct, err

}

// PackGetValidatorJailInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0e222b36.
//
// Solidity: function getValidatorJailInfo(string id) view returns(uint256 jailUntil, bool isJailed)
func (stakeHub *StakeHub) PackGetValidatorJailInfo(id string) []byte {
	enc, err := stakeHub.abi.Pack("getValidatorJailInfo", id)
	if err != nil {
		panic(err)
	}
	return enc
}

// GetValidatorJailInfoOutput serves as a container for the return parameters of contract
// method GetValidatorJailInfo.
type GetValidatorJailInfoOutput struct {
	JailUntil *big.Int
	IsJailed  bool
}

// UnpackGetValidatorJailInfo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0e222b36.
//
// Solidity: function getValidatorJailInfo(string id) view returns(uint256 jailUntil, bool isJailed)
func (stakeHub *StakeHub) UnpackGetValidatorJailInfo(data []byte) (GetValidatorJailInfoOutput, error) {
	out, err := stakeHub.abi.Unpack("getValidatorJailInfo", data)
	outstruct := new(GetValidatorJailInfoOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.JailUntil = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.IsJailed = *abi.ConvertType(out[1], new(bool)).(*bool)
	return *outstruct, err

}

// PackGetValidatorJailInfo0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbbe5e345.
//
// Solidity: function getValidatorJailInfo(address operatorAddr) view returns(uint256 jailUntil, bool isJailed)
func (stakeHub *StakeHub) PackGetValidatorJailInfo0(operatorAddr common.Address) []byte {
	enc, err := stakeHub.abi.Pack("getValidatorJailInfo0", operatorAddr)
	if err != nil {
		panic(err)
	}
	return enc
}

// GetValidatorJailInfo0Output serves as a container for the return parameters of contract
// method GetValidatorJailInfo0.
type GetValidatorJailInfo0Output struct {
	JailUntil *big.Int
	IsJailed  bool
}

// UnpackGetValidatorJailInfo0 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbbe5e345.
//
// Solidity: function getValidatorJailInfo(address operatorAddr) view returns(uint256 jailUntil, bool isJailed)
func (stakeHub *StakeHub) UnpackGetValidatorJailInfo0(data []byte) (GetValidatorJailInfo0Output, error) {
	out, err := stakeHub.abi.Unpack("getValidatorJailInfo0", data)
	outstruct := new(GetValidatorJailInfo0Output)
	if err != nil {
		return *outstruct, err
	}
	outstruct.JailUntil = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.IsJailed = *abi.ConvertType(out[1], new(bool)).(*bool)
	return *outstruct, err

}

// PackGetValidatorTotalStake is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2aa8e06d.
//
// Solidity: function getValidatorTotalStake(address operatorAddr) view returns(uint256)
func (stakeHub *StakeHub) PackGetValidatorTotalStake(operatorAddr common.Address) []byte {
	enc, err := stakeHub.abi.Pack("getValidatorTotalStake", operatorAddr)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetValidatorTotalStake is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2aa8e06d.
//
// Solidity: function getValidatorTotalStake(address operatorAddr) view returns(uint256)
func (stakeHub *StakeHub) UnpackGetValidatorTotalStake(data []byte) (*big.Int, error) {
	out, err := stakeHub.abi.Unpack("getValidatorTotalStake", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetValidatorTotalStake0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8c7d0dbc.
//
// Solidity: function getValidatorTotalStake(string id) view returns(uint256)
func (stakeHub *StakeHub) PackGetValidatorTotalStake0(id string) []byte {
	enc, err := stakeHub.abi.Pack("getValidatorTotalStake0", id)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetValidatorTotalStake0 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8c7d0dbc.
//
// Solidity: function getValidatorTotalStake(string id) view returns(uint256)
func (stakeHub *StakeHub) UnpackGetValidatorTotalStake0(data []byte) (*big.Int, error) {
	out, err := stakeHub.abi.Unpack("getValidatorTotalStake0", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetValidatorUnstakingInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x63973e4b.
//
// Solidity: function getValidatorUnstakingInfo(string id) view returns(uint256 unstakeTime, bool isUnstaking, bool isUnstaked)
func (stakeHub *StakeHub) PackGetValidatorUnstakingInfo(id string) []byte {
	enc, err := stakeHub.abi.Pack("getValidatorUnstakingInfo", id)
	if err != nil {
		panic(err)
	}
	return enc
}

// GetValidatorUnstakingInfoOutput serves as a container for the return parameters of contract
// method GetValidatorUnstakingInfo.
type GetValidatorUnstakingInfoOutput struct {
	UnstakeTime *big.Int
	IsUnstaking bool
	IsUnstaked  bool
}

// UnpackGetValidatorUnstakingInfo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x63973e4b.
//
// Solidity: function getValidatorUnstakingInfo(string id) view returns(uint256 unstakeTime, bool isUnstaking, bool isUnstaked)
func (stakeHub *StakeHub) UnpackGetValidatorUnstakingInfo(data []byte) (GetValidatorUnstakingInfoOutput, error) {
	out, err := stakeHub.abi.Unpack("getValidatorUnstakingInfo", data)
	outstruct := new(GetValidatorUnstakingInfoOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.UnstakeTime = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.IsUnstaking = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.IsUnstaked = *abi.ConvertType(out[2], new(bool)).(*bool)
	return *outstruct, err

}

// PackGetValidatorUnstakingInfo0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa4cac964.
//
// Solidity: function getValidatorUnstakingInfo(address operatorAddr) view returns(uint256 unstakeTime, bool isUnstaking, bool isUnstaked)
func (stakeHub *StakeHub) PackGetValidatorUnstakingInfo0(operatorAddr common.Address) []byte {
	enc, err := stakeHub.abi.Pack("getValidatorUnstakingInfo0", operatorAddr)
	if err != nil {
		panic(err)
	}
	return enc
}

// GetValidatorUnstakingInfo0Output serves as a container for the return parameters of contract
// method GetValidatorUnstakingInfo0.
type GetValidatorUnstakingInfo0Output struct {
	UnstakeTime *big.Int
	IsUnstaking bool
	IsUnstaked  bool
}

// UnpackGetValidatorUnstakingInfo0 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa4cac964.
//
// Solidity: function getValidatorUnstakingInfo(address operatorAddr) view returns(uint256 unstakeTime, bool isUnstaking, bool isUnstaked)
func (stakeHub *StakeHub) UnpackGetValidatorUnstakingInfo0(data []byte) (GetValidatorUnstakingInfo0Output, error) {
	out, err := stakeHub.abi.Unpack("getValidatorUnstakingInfo0", data)
	outstruct := new(GetValidatorUnstakingInfo0Output)
	if err != nil {
		return *outstruct, err
	}
	outstruct.UnstakeTime = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.IsUnstaking = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.IsUnstaked = *abi.ConvertType(out[2], new(bool)).(*bool)
	return *outstruct, err

}

// PackGetValidators is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbff02e20.
//
// Solidity: function getValidators(uint256 offset, uint256 limit) view returns(address[] validatorAddrs, bytes[] signerAddrs, uint256[] amounts, uint256 totalLength)
func (stakeHub *StakeHub) PackGetValidators(offset *big.Int, limit *big.Int) []byte {
	enc, err := stakeHub.abi.Pack("getValidators", offset, limit)
	if err != nil {
		panic(err)
	}
	return enc
}

// GetValidatorsOutput serves as a container for the return parameters of contract
// method GetValidators.
type GetValidatorsOutput struct {
	ValidatorAddrs []common.Address
	SignerAddrs    [][]byte
	Amounts        []*big.Int
	TotalLength    *big.Int
}

// UnpackGetValidators is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbff02e20.
//
// Solidity: function getValidators(uint256 offset, uint256 limit) view returns(address[] validatorAddrs, bytes[] signerAddrs, uint256[] amounts, uint256 totalLength)
func (stakeHub *StakeHub) UnpackGetValidators(data []byte) (GetValidatorsOutput, error) {
	out, err := stakeHub.abi.Unpack("getValidators", data)
	outstruct := new(GetValidatorsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.ValidatorAddrs = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.SignerAddrs = *abi.ConvertType(out[1], new([][]byte)).(*[][]byte)
	outstruct.Amounts = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	outstruct.TotalLength = abi.ConvertType(out[3], new(big.Int)).(*big.Int)
	return *outstruct, err

}

// PackIdToOperator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9ac83d85.
//
// Solidity: function idToOperator(bytes32 ) view returns(address)
func (stakeHub *StakeHub) PackIdToOperator(arg0 [32]byte) []byte {
	enc, err := stakeHub.abi.Pack("idToOperator", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIdToOperator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9ac83d85.
//
// Solidity: function idToOperator(bytes32 ) view returns(address)
func (stakeHub *StakeHub) UnpackIdToOperator(data []byte) (common.Address, error) {
	out, err := stakeHub.abi.Unpack("idToOperator", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (stakeHub *StakeHub) PackInitialize() []byte {
	enc, err := stakeHub.abi.Pack("initialize")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackIsActiveValidator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x40550a1c.
//
// Solidity: function isActiveValidator(address operatorAddr) view returns(bool)
func (stakeHub *StakeHub) PackIsActiveValidator(operatorAddr common.Address) []byte {
	enc, err := stakeHub.abi.Pack("isActiveValidator", operatorAddr)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsActiveValidator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x40550a1c.
//
// Solidity: function isActiveValidator(address operatorAddr) view returns(bool)
func (stakeHub *StakeHub) UnpackIsActiveValidator(data []byte) (bool, error) {
	out, err := stakeHub.abi.Unpack("isActiveValidator", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackIsActiveValidator0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8661213a.
//
// Solidity: function isActiveValidator(string id) view returns(bool)
func (stakeHub *StakeHub) PackIsActiveValidator0(id string) []byte {
	enc, err := stakeHub.abi.Pack("isActiveValidator0", id)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsActiveValidator0 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8661213a.
//
// Solidity: function isActiveValidator(string id) view returns(bool)
func (stakeHub *StakeHub) UnpackIsActiveValidator0(data []byte) (bool, error) {
	out, err := stakeHub.abi.Unpack("isActiveValidator0", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackMinDelegateAmount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x72e9c934.
//
// Solidity: function minDelegateAmount() view returns(uint256)
func (stakeHub *StakeHub) PackMinDelegateAmount() []byte {
	enc, err := stakeHub.abi.Pack("minDelegateAmount")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMinDelegateAmount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x72e9c934.
//
// Solidity: function minDelegateAmount() view returns(uint256)
func (stakeHub *StakeHub) UnpackMinDelegateAmount(data []byte) (*big.Int, error) {
	out, err := stakeHub.abi.Unpack("minDelegateAmount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackMinStakeAmount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf1887684.
//
// Solidity: function minStakeAmount() view returns(uint256)
func (stakeHub *StakeHub) PackMinStakeAmount() []byte {
	enc, err := stakeHub.abi.Pack("minStakeAmount")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMinStakeAmount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf1887684.
//
// Solidity: function minStakeAmount() view returns(uint256)
func (stakeHub *StakeHub) UnpackMinStakeAmount(data []byte) (*big.Int, error) {
	out, err := stakeHub.abi.Unpack("minStakeAmount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackOfflineJailDuration is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42fce3cd.
//
// Solidity: function offlineJailDuration() view returns(uint256)
func (stakeHub *StakeHub) PackOfflineJailDuration() []byte {
	enc, err := stakeHub.abi.Pack("offlineJailDuration")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOfflineJailDuration is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x42fce3cd.
//
// Solidity: function offlineJailDuration() view returns(uint256)
func (stakeHub *StakeHub) UnpackOfflineJailDuration(data []byte) (*big.Int, error) {
	out, err := stakeHub.abi.Unpack("offlineJailDuration", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackOfflineSlashAmount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9cc60f3.
//
// Solidity: function offlineSlashAmount() view returns(uint256)
func (stakeHub *StakeHub) PackOfflineSlashAmount() []byte {
	enc, err := stakeHub.abi.Pack("offlineSlashAmount")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOfflineSlashAmount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa9cc60f3.
//
// Solidity: function offlineSlashAmount() view returns(uint256)
func (stakeHub *StakeHub) UnpackOfflineSlashAmount(data []byte) (*big.Int, error) {
	out, err := stakeHub.abi.Unpack("offlineSlashAmount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackPause is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8456cb59.
//
// Solidity: function pause() returns()
func (stakeHub *StakeHub) PackPause() []byte {
	enc, err := stakeHub.abi.Pack("pause")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackPaused is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (stakeHub *StakeHub) PackPaused() []byte {
	enc, err := stakeHub.abi.Pack("paused")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPaused is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (stakeHub *StakeHub) UnpackPaused(data []byte) (bool, error) {
	out, err := stakeHub.abi.Unpack("paused", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackRemoveFromBlackList is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4a49ac4c.
//
// Solidity: function removeFromBlackList(address account) returns()
func (stakeHub *StakeHub) PackRemoveFromBlackList(account common.Address) []byte {
	enc, err := stakeHub.abi.Pack("removeFromBlackList", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetKeeper is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x748747e6.
//
// Solidity: function setKeeper(address keeper) returns()
func (stakeHub *StakeHub) PackSetKeeper(keeper common.Address) []byte {
	enc, err := stakeHub.abi.Pack("setKeeper", keeper)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetMinDelegateAmount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3e732eba.
//
// Solidity: function setMinDelegateAmount(uint256 amount) returns()
func (stakeHub *StakeHub) PackSetMinDelegateAmount(amount *big.Int) []byte {
	enc, err := stakeHub.abi.Pack("setMinDelegateAmount", amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetMinStakeAmount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xeb4af045.
//
// Solidity: function setMinStakeAmount(uint256 amount) returns()
func (stakeHub *StakeHub) PackSetMinStakeAmount(amount *big.Int) []byte {
	enc, err := stakeHub.abi.Pack("setMinStakeAmount", amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetOfflineJailDuration is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x43a769d9.
//
// Solidity: function setOfflineJailDuration(uint256 duration) returns()
func (stakeHub *StakeHub) PackSetOfflineJailDuration(duration *big.Int) []byte {
	enc, err := stakeHub.abi.Pack("setOfflineJailDuration", duration)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetOfflineSlashAmount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9679904a.
//
// Solidity: function setOfflineSlashAmount(uint256 amount) returns()
func (stakeHub *StakeHub) PackSetOfflineSlashAmount(amount *big.Int) []byte {
	enc, err := stakeHub.abi.Pack("setOfflineSlashAmount", amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetUnbondingPeriod is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x114eaf55.
//
// Solidity: function setUnbondingPeriod(uint256 period) returns()
func (stakeHub *StakeHub) PackSetUnbondingPeriod(period *big.Int) []byte {
	enc, err := stakeHub.abi.Pack("setUnbondingPeriod", period)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetValidatorThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5bc0fbc7.
//
// Solidity: function setValidatorThreshold(uint256 count) returns()
func (stakeHub *StakeHub) PackSetValidatorThreshold(count *big.Int) []byte {
	enc, err := stakeHub.abi.Pack("setValidatorThreshold", count)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSignerToOperator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4c23c8fa.
//
// Solidity: function signerToOperator(bytes ) view returns(address)
func (stakeHub *StakeHub) PackSignerToOperator(arg0 []byte) []byte {
	enc, err := stakeHub.abi.Pack("signerToOperator", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSignerToOperator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4c23c8fa.
//
// Solidity: function signerToOperator(bytes ) view returns(address)
func (stakeHub *StakeHub) UnpackSignerToOperator(data []byte) (common.Address, error) {
	out, err := stakeHub.abi.Unpack("signerToOperator", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackSlashOffline is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x44690e5c.
//
// Solidity: function slashOffline(address validatorAddr) returns()
func (stakeHub *StakeHub) PackSlashOffline(validatorAddr common.Address) []byte {
	enc, err := stakeHub.abi.Pack("slashOffline", validatorAddr)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackTotalSlashedAmount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xba723eda.
//
// Solidity: function totalSlashedAmount() view returns(uint256)
func (stakeHub *StakeHub) PackTotalSlashedAmount() []byte {
	enc, err := stakeHub.abi.Pack("totalSlashedAmount")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTotalSlashedAmount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xba723eda.
//
// Solidity: function totalSlashedAmount() view returns(uint256)
func (stakeHub *StakeHub) UnpackTotalSlashedAmount(data []byte) (*big.Int, error) {
	out, err := stakeHub.abi.Unpack("totalSlashedAmount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackUnbondingPeriod is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6cf6d675.
//
// Solidity: function unbondingPeriod() view returns(uint256)
func (stakeHub *StakeHub) PackUnbondingPeriod() []byte {
	enc, err := stakeHub.abi.Pack("unbondingPeriod")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUnbondingPeriod is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6cf6d675.
//
// Solidity: function unbondingPeriod() view returns(uint256)
func (stakeHub *StakeHub) UnpackUnbondingPeriod(data []byte) (*big.Int, error) {
	out, err := stakeHub.abi.Unpack("unbondingPeriod", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackUndelegate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4d99dd16.
//
// Solidity: function undelegate(address operatorAddr, uint256 shares) returns()
func (stakeHub *StakeHub) PackUndelegate(operatorAddr common.Address, shares *big.Int) []byte {
	enc, err := stakeHub.abi.Pack("undelegate", operatorAddr, shares)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUndelegate0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8dfc8897.
//
// Solidity: function undelegate(string id, uint256 shares) returns()
func (stakeHub *StakeHub) PackUndelegate0(id string, shares *big.Int) []byte {
	enc, err := stakeHub.abi.Pack("undelegate0", id, shares)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUnjail is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf679d305.
//
// Solidity: function unjail() returns()
func (stakeHub *StakeHub) PackUnjail() []byte {
	enc, err := stakeHub.abi.Pack("unjail")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUnpause is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (stakeHub *StakeHub) PackUnpause() []byte {
	enc, err := stakeHub.abi.Pack("unpause")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUnstakeValidator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd86c7fe8.
//
// Solidity: function unstakeValidator() returns()
func (stakeHub *StakeHub) PackUnstakeValidator() []byte {
	enc, err := stakeHub.abi.Pack("unstakeValidator")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUpdateCommissionRate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x00fa3d50.
//
// Solidity: function updateCommissionRate(uint256 newCommissionRate) returns()
func (stakeHub *StakeHub) PackUpdateCommissionRate(newCommissionRate *big.Int) []byte {
	enc, err := stakeHub.abi.Pack("updateCommissionRate", newCommissionRate)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUpdateSigner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd8d50a2a.
//
// Solidity: function updateSigner(bytes signerAddr) returns()
func (stakeHub *StakeHub) PackUpdateSigner(signerAddr []byte) []byte {
	enc, err := stakeHub.abi.Pack("updateSigner", signerAddr)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUpdateValidator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x635b2770.
//
// Solidity: function updateValidator(address newValidatorAddr) returns()
func (stakeHub *StakeHub) PackUpdateValidator(newValidatorAddr common.Address) []byte {
	enc, err := stakeHub.abi.Pack("updateValidator", newValidatorAddr)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackValidatorThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4fd101d7.
//
// Solidity: function validatorThreshold() view returns(uint256)
func (stakeHub *StakeHub) PackValidatorThreshold() []byte {
	enc, err := stakeHub.abi.Pack("validatorThreshold")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackValidatorThreshold is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4fd101d7.
//
// Solidity: function validatorThreshold() view returns(uint256)
func (stakeHub *StakeHub) UnpackValidatorThreshold(data []byte) (*big.Int, error) {
	out, err := stakeHub.abi.Unpack("validatorThreshold", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackValidatorToOperator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x35efc734.
//
// Solidity: function validatorToOperator(address ) view returns(address)
func (stakeHub *StakeHub) PackValidatorToOperator(arg0 common.Address) []byte {
	enc, err := stakeHub.abi.Pack("validatorToOperator", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackValidatorToOperator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x35efc734.
//
// Solidity: function validatorToOperator(address ) view returns(address)
func (stakeHub *StakeHub) UnpackValidatorToOperator(data []byte) (common.Address, error) {
	out, err := stakeHub.abi.Unpack("validatorToOperator", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// StakeHubBlackListed represents a BlackListed event raised by the StakeHub contract.
type StakeHubBlackListed struct {
	Account common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const StakeHubBlackListedEventName = "BlackListed"

// ContractEventName returns the user-defined event name.
func (StakeHubBlackListed) ContractEventName() string {
	return StakeHubBlackListedEventName
}

// UnpackBlackListedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event BlackListed(address indexed account)
func (stakeHub *StakeHub) UnpackBlackListedEvent(log *types.Log) (*StakeHubBlackListed, error) {
	event := "BlackListed"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubBlackListed)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// StakeHubClaimed represents a Claimed event raised by the StakeHub contract.
type StakeHubClaimed struct {
	OperatorAddr common.Address
	Delegator    common.Address
	Amount       *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const StakeHubClaimedEventName = "Claimed"

// ContractEventName returns the user-defined event name.
func (StakeHubClaimed) ContractEventName() string {
	return StakeHubClaimedEventName
}

// UnpackClaimedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Claimed(address indexed operatorAddr, address indexed delegator, uint256 amount)
func (stakeHub *StakeHub) UnpackClaimedEvent(log *types.Log) (*StakeHubClaimed, error) {
	event := "Claimed"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubClaimed)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// StakeHubCommissionRateUpdated represents a CommissionRateUpdated event raised by the StakeHub contract.
type StakeHubCommissionRateUpdated struct {
	OperatorAddr      common.Address
	NewCommissionRate *big.Int
	Raw               *types.Log // Blockchain specific contextual infos
}

const StakeHubCommissionRateUpdatedEventName = "CommissionRateUpdated"

// ContractEventName returns the user-defined event name.
func (StakeHubCommissionRateUpdated) ContractEventName() string {
	return StakeHubCommissionRateUpdatedEventName
}

// UnpackCommissionRateUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event CommissionRateUpdated(address indexed operatorAddr, uint256 newCommissionRate)
func (stakeHub *StakeHub) UnpackCommissionRateUpdatedEvent(log *types.Log) (*StakeHubCommissionRateUpdated, error) {
	event := "CommissionRateUpdated"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubCommissionRateUpdated)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// StakeHubDelegated represents a Delegated event raised by the StakeHub contract.
type StakeHubDelegated struct {
	OperatorAddr common.Address
	Delegator    common.Address
	Amount       *big.Int
	Shares       *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const StakeHubDelegatedEventName = "Delegated"

// ContractEventName returns the user-defined event name.
func (StakeHubDelegated) ContractEventName() string {
	return StakeHubDelegatedEventName
}

// UnpackDelegatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Delegated(address indexed operatorAddr, address indexed delegator, uint256 amount, uint256 shares)
func (stakeHub *StakeHub) UnpackDelegatedEvent(log *types.Log) (*StakeHubDelegated, error) {
	event := "Delegated"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubDelegated)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// StakeHubInitialized represents a Initialized event raised by the StakeHub contract.
type StakeHubInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const StakeHubInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (StakeHubInitialized) ContractEventName() string {
	return StakeHubInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (stakeHub *StakeHub) UnpackInitializedEvent(log *types.Log) (*StakeHubInitialized, error) {
	event := "Initialized"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubInitialized)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// StakeHubKeeperChanged represents a KeeperChanged event raised by the StakeHub contract.
type StakeHubKeeperChanged struct {
	PreviousKeeper common.Address
	NewKeeper      common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const StakeHubKeeperChangedEventName = "KeeperChanged"

// ContractEventName returns the user-defined event name.
func (StakeHubKeeperChanged) ContractEventName() string {
	return StakeHubKeeperChangedEventName
}

// UnpackKeeperChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event KeeperChanged(address indexed previousKeeper, address indexed newKeeper)
func (stakeHub *StakeHub) UnpackKeeperChangedEvent(log *types.Log) (*StakeHubKeeperChanged, error) {
	event := "KeeperChanged"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubKeeperChanged)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// StakeHubKeeperUpdated represents a KeeperUpdated event raised by the StakeHub contract.
type StakeHubKeeperUpdated struct {
	Keeper common.Address
	Raw    *types.Log // Blockchain specific contextual infos
}

const StakeHubKeeperUpdatedEventName = "KeeperUpdated"

// ContractEventName returns the user-defined event name.
func (StakeHubKeeperUpdated) ContractEventName() string {
	return StakeHubKeeperUpdatedEventName
}

// UnpackKeeperUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event KeeperUpdated(address keeper)
func (stakeHub *StakeHub) UnpackKeeperUpdatedEvent(log *types.Log) (*StakeHubKeeperUpdated, error) {
	event := "KeeperUpdated"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubKeeperUpdated)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// StakeHubMinDelegateAmountUpdated represents a MinDelegateAmountUpdated event raised by the StakeHub contract.
type StakeHubMinDelegateAmountUpdated struct {
	Amount *big.Int
	Raw    *types.Log // Blockchain specific contextual infos
}

const StakeHubMinDelegateAmountUpdatedEventName = "MinDelegateAmountUpdated"

// ContractEventName returns the user-defined event name.
func (StakeHubMinDelegateAmountUpdated) ContractEventName() string {
	return StakeHubMinDelegateAmountUpdatedEventName
}

// UnpackMinDelegateAmountUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MinDelegateAmountUpdated(uint256 amount)
func (stakeHub *StakeHub) UnpackMinDelegateAmountUpdatedEvent(log *types.Log) (*StakeHubMinDelegateAmountUpdated, error) {
	event := "MinDelegateAmountUpdated"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubMinDelegateAmountUpdated)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// StakeHubMinStakeAmountUpdated represents a MinStakeAmountUpdated event raised by the StakeHub contract.
type StakeHubMinStakeAmountUpdated struct {
	Amount *big.Int
	Raw    *types.Log // Blockchain specific contextual infos
}

const StakeHubMinStakeAmountUpdatedEventName = "MinStakeAmountUpdated"

// ContractEventName returns the user-defined event name.
func (StakeHubMinStakeAmountUpdated) ContractEventName() string {
	return StakeHubMinStakeAmountUpdatedEventName
}

// UnpackMinStakeAmountUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MinStakeAmountUpdated(uint256 amount)
func (stakeHub *StakeHub) UnpackMinStakeAmountUpdatedEvent(log *types.Log) (*StakeHubMinStakeAmountUpdated, error) {
	event := "MinStakeAmountUpdated"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubMinStakeAmountUpdated)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// StakeHubOfflineJailDurationUpdated represents a OfflineJailDurationUpdated event raised by the StakeHub contract.
type StakeHubOfflineJailDurationUpdated struct {
	Duration *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const StakeHubOfflineJailDurationUpdatedEventName = "OfflineJailDurationUpdated"

// ContractEventName returns the user-defined event name.
func (StakeHubOfflineJailDurationUpdated) ContractEventName() string {
	return StakeHubOfflineJailDurationUpdatedEventName
}

// UnpackOfflineJailDurationUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OfflineJailDurationUpdated(uint256 duration)
func (stakeHub *StakeHub) UnpackOfflineJailDurationUpdatedEvent(log *types.Log) (*StakeHubOfflineJailDurationUpdated, error) {
	event := "OfflineJailDurationUpdated"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubOfflineJailDurationUpdated)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// StakeHubOfflineSlashAmountUpdated represents a OfflineSlashAmountUpdated event raised by the StakeHub contract.
type StakeHubOfflineSlashAmountUpdated struct {
	Amount *big.Int
	Raw    *types.Log // Blockchain specific contextual infos
}

const StakeHubOfflineSlashAmountUpdatedEventName = "OfflineSlashAmountUpdated"

// ContractEventName returns the user-defined event name.
func (StakeHubOfflineSlashAmountUpdated) ContractEventName() string {
	return StakeHubOfflineSlashAmountUpdatedEventName
}

// UnpackOfflineSlashAmountUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OfflineSlashAmountUpdated(uint256 amount)
func (stakeHub *StakeHub) UnpackOfflineSlashAmountUpdatedEvent(log *types.Log) (*StakeHubOfflineSlashAmountUpdated, error) {
	event := "OfflineSlashAmountUpdated"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubOfflineSlashAmountUpdated)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// StakeHubPaused represents a Paused event raised by the StakeHub contract.
type StakeHubPaused struct {
	Account common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const StakeHubPausedEventName = "Paused"

// ContractEventName returns the user-defined event name.
func (StakeHubPaused) ContractEventName() string {
	return StakeHubPausedEventName
}

// UnpackPausedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Paused(address account)
func (stakeHub *StakeHub) UnpackPausedEvent(log *types.Log) (*StakeHubPaused, error) {
	event := "Paused"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubPaused)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// StakeHubRewardDistributed represents a RewardDistributed event raised by the StakeHub contract.
type StakeHubRewardDistributed struct {
	OperatorAddr common.Address
	Amount       *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const StakeHubRewardDistributedEventName = "RewardDistributed"

// ContractEventName returns the user-defined event name.
func (StakeHubRewardDistributed) ContractEventName() string {
	return StakeHubRewardDistributedEventName
}

// UnpackRewardDistributedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RewardDistributed(address indexed operatorAddr, uint256 amount)
func (stakeHub *StakeHub) UnpackRewardDistributedEvent(log *types.Log) (*StakeHubRewardDistributed, error) {
	event := "RewardDistributed"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubRewardDistributed)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// StakeHubSignerUpdated represents a SignerUpdated event raised by the StakeHub contract.
type StakeHubSignerUpdated struct {
	OperatorAddr common.Address
	SignerAddr   []byte
	Raw          *types.Log // Blockchain specific contextual infos
}

const StakeHubSignerUpdatedEventName = "SignerUpdated"

// ContractEventName returns the user-defined event name.
func (StakeHubSignerUpdated) ContractEventName() string {
	return StakeHubSignerUpdatedEventName
}

// UnpackSignerUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SignerUpdated(address indexed operatorAddr, bytes signerAddr)
func (stakeHub *StakeHub) UnpackSignerUpdatedEvent(log *types.Log) (*StakeHubSignerUpdated, error) {
	event := "SignerUpdated"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubSignerUpdated)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// StakeHubUnBlackListed represents a UnBlackListed event raised by the StakeHub contract.
type StakeHubUnBlackListed struct {
	Account common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const StakeHubUnBlackListedEventName = "UnBlackListed"

// ContractEventName returns the user-defined event name.
func (StakeHubUnBlackListed) ContractEventName() string {
	return StakeHubUnBlackListedEventName
}

// UnpackUnBlackListedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event UnBlackListed(address indexed account)
func (stakeHub *StakeHub) UnpackUnBlackListedEvent(log *types.Log) (*StakeHubUnBlackListed, error) {
	event := "UnBlackListed"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubUnBlackListed)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// StakeHubUnbondingPeriodUpdated represents a UnbondingPeriodUpdated event raised by the StakeHub contract.
type StakeHubUnbondingPeriodUpdated struct {
	Period *big.Int
	Raw    *types.Log // Blockchain specific contextual infos
}

const StakeHubUnbondingPeriodUpdatedEventName = "UnbondingPeriodUpdated"

// ContractEventName returns the user-defined event name.
func (StakeHubUnbondingPeriodUpdated) ContractEventName() string {
	return StakeHubUnbondingPeriodUpdatedEventName
}

// UnpackUnbondingPeriodUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event UnbondingPeriodUpdated(uint256 period)
func (stakeHub *StakeHub) UnpackUnbondingPeriodUpdatedEvent(log *types.Log) (*StakeHubUnbondingPeriodUpdated, error) {
	event := "UnbondingPeriodUpdated"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubUnbondingPeriodUpdated)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// StakeHubUndelegated represents a Undelegated event raised by the StakeHub contract.
type StakeHubUndelegated struct {
	OperatorAddr common.Address
	Delegator    common.Address
	Amount       *big.Int
	Shares       *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const StakeHubUndelegatedEventName = "Undelegated"

// ContractEventName returns the user-defined event name.
func (StakeHubUndelegated) ContractEventName() string {
	return StakeHubUndelegatedEventName
}

// UnpackUndelegatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Undelegated(address indexed operatorAddr, address indexed delegator, uint256 amount, uint256 shares)
func (stakeHub *StakeHub) UnpackUndelegatedEvent(log *types.Log) (*StakeHubUndelegated, error) {
	event := "Undelegated"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubUndelegated)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// StakeHubUnpaused represents a Unpaused event raised by the StakeHub contract.
type StakeHubUnpaused struct {
	Account common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const StakeHubUnpausedEventName = "Unpaused"

// ContractEventName returns the user-defined event name.
func (StakeHubUnpaused) ContractEventName() string {
	return StakeHubUnpausedEventName
}

// UnpackUnpausedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Unpaused(address account)
func (stakeHub *StakeHub) UnpackUnpausedEvent(log *types.Log) (*StakeHubUnpaused, error) {
	event := "Unpaused"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubUnpaused)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// StakeHubValidatorApplied represents a ValidatorApplied event raised by the StakeHub contract.
type StakeHubValidatorApplied struct {
	ValidatorAddr common.Address
	OperatorAddr  common.Address
	ShareContract common.Address
	SignerAddr    []byte
	Raw           *types.Log // Blockchain specific contextual infos
}

const StakeHubValidatorAppliedEventName = "ValidatorApplied"

// ContractEventName returns the user-defined event name.
func (StakeHubValidatorApplied) ContractEventName() string {
	return StakeHubValidatorAppliedEventName
}

// UnpackValidatorAppliedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ValidatorApplied(address indexed validatorAddr, address indexed operatorAddr, address indexed shareContract, bytes signerAddr)
func (stakeHub *StakeHub) UnpackValidatorAppliedEvent(log *types.Log) (*StakeHubValidatorApplied, error) {
	event := "ValidatorApplied"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubValidatorApplied)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// StakeHubValidatorJailed represents a ValidatorJailed event raised by the StakeHub contract.
type StakeHubValidatorJailed struct {
	OperatorAddr common.Address
	JailUntil    *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const StakeHubValidatorJailedEventName = "ValidatorJailed"

// ContractEventName returns the user-defined event name.
func (StakeHubValidatorJailed) ContractEventName() string {
	return StakeHubValidatorJailedEventName
}

// UnpackValidatorJailedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ValidatorJailed(address indexed operatorAddr, uint256 jailUntil)
func (stakeHub *StakeHub) UnpackValidatorJailedEvent(log *types.Log) (*StakeHubValidatorJailed, error) {
	event := "ValidatorJailed"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubValidatorJailed)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// StakeHubValidatorSlashed represents a ValidatorSlashed event raised by the StakeHub contract.
type StakeHubValidatorSlashed struct {
	OperatorAddr common.Address
	SlashAmount  *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const StakeHubValidatorSlashedEventName = "ValidatorSlashed"

// ContractEventName returns the user-defined event name.
func (StakeHubValidatorSlashed) ContractEventName() string {
	return StakeHubValidatorSlashedEventName
}

// UnpackValidatorSlashedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ValidatorSlashed(address indexed operatorAddr, uint256 slashAmount)
func (stakeHub *StakeHub) UnpackValidatorSlashedEvent(log *types.Log) (*StakeHubValidatorSlashed, error) {
	event := "ValidatorSlashed"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubValidatorSlashed)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// StakeHubValidatorThresholdUpdated represents a ValidatorThresholdUpdated event raised by the StakeHub contract.
type StakeHubValidatorThresholdUpdated struct {
	Count *big.Int
	Raw   *types.Log // Blockchain specific contextual infos
}

const StakeHubValidatorThresholdUpdatedEventName = "ValidatorThresholdUpdated"

// ContractEventName returns the user-defined event name.
func (StakeHubValidatorThresholdUpdated) ContractEventName() string {
	return StakeHubValidatorThresholdUpdatedEventName
}

// UnpackValidatorThresholdUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ValidatorThresholdUpdated(uint256 count)
func (stakeHub *StakeHub) UnpackValidatorThresholdUpdatedEvent(log *types.Log) (*StakeHubValidatorThresholdUpdated, error) {
	event := "ValidatorThresholdUpdated"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubValidatorThresholdUpdated)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// StakeHubValidatorUnjailed represents a ValidatorUnjailed event raised by the StakeHub contract.
type StakeHubValidatorUnjailed struct {
	OperatorAddr common.Address
	Raw          *types.Log // Blockchain specific contextual infos
}

const StakeHubValidatorUnjailedEventName = "ValidatorUnjailed"

// ContractEventName returns the user-defined event name.
func (StakeHubValidatorUnjailed) ContractEventName() string {
	return StakeHubValidatorUnjailedEventName
}

// UnpackValidatorUnjailedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ValidatorUnjailed(address indexed operatorAddr)
func (stakeHub *StakeHub) UnpackValidatorUnjailedEvent(log *types.Log) (*StakeHubValidatorUnjailed, error) {
	event := "ValidatorUnjailed"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubValidatorUnjailed)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// StakeHubValidatorUnstaked represents a ValidatorUnstaked event raised by the StakeHub contract.
type StakeHubValidatorUnstaked struct {
	ValidatorAddr common.Address
	OperatorAddr  common.Address
	UnstakeTime   *big.Int
	Raw           *types.Log // Blockchain specific contextual infos
}

const StakeHubValidatorUnstakedEventName = "ValidatorUnstaked"

// ContractEventName returns the user-defined event name.
func (StakeHubValidatorUnstaked) ContractEventName() string {
	return StakeHubValidatorUnstakedEventName
}

// UnpackValidatorUnstakedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ValidatorUnstaked(address indexed validatorAddr, address indexed operatorAddr, uint256 unstakeTime)
func (stakeHub *StakeHub) UnpackValidatorUnstakedEvent(log *types.Log) (*StakeHubValidatorUnstaked, error) {
	event := "ValidatorUnstaked"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubValidatorUnstaked)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// StakeHubValidatorUpdated represents a ValidatorUpdated event raised by the StakeHub contract.
type StakeHubValidatorUpdated struct {
	OperatorAddr     common.Address
	NewValidatorAddr common.Address
	Raw              *types.Log // Blockchain specific contextual infos
}

const StakeHubValidatorUpdatedEventName = "ValidatorUpdated"

// ContractEventName returns the user-defined event name.
func (StakeHubValidatorUpdated) ContractEventName() string {
	return StakeHubValidatorUpdatedEventName
}

// UnpackValidatorUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ValidatorUpdated(address indexed operatorAddr, address indexed newValidatorAddr)
func (stakeHub *StakeHub) UnpackValidatorUpdatedEvent(log *types.Log) (*StakeHubValidatorUpdated, error) {
	event := "ValidatorUpdated"
	if log.Topics[0] != stakeHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(StakeHubValidatorUpdated)
	if len(log.Data) > 0 {
		if err := stakeHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range stakeHub.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (stakeHub *StakeHub) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["DuplicateId"].ID.Bytes()[:4]) {
		return stakeHub.UnpackDuplicateIdError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["DuplicateOperatorAddress"].ID.Bytes()[:4]) {
		return stakeHub.UnpackDuplicateOperatorAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["DuplicateSignerAddress"].ID.Bytes()[:4]) {
		return stakeHub.UnpackDuplicateSignerAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["DuplicateValidatorAddress"].ID.Bytes()[:4]) {
		return stakeHub.UnpackDuplicateValidatorAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["EnforcedPause"].ID.Bytes()[:4]) {
		return stakeHub.UnpackEnforcedPauseError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["ExpectedPause"].ID.Bytes()[:4]) {
		return stakeHub.UnpackExpectedPauseError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return stakeHub.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["FailedDeployment"].ID.Bytes()[:4]) {
		return stakeHub.UnpackFailedDeploymentError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["InBlackList"].ID.Bytes()[:4]) {
		return stakeHub.UnpackInBlackListError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["InsufficientBalance"].ID.Bytes()[:4]) {
		return stakeHub.UnpackInsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["InvalidCommissionRate"].ID.Bytes()[:4]) {
		return stakeHub.UnpackInvalidCommissionRateError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["InvalidId"].ID.Bytes()[:4]) {
		return stakeHub.UnpackInvalidIdError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return stakeHub.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["InvalidSignerAddress"].ID.Bytes()[:4]) {
		return stakeHub.UnpackInvalidSignerAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["InvalidValidatorAddress"].ID.Bytes()[:4]) {
		return stakeHub.UnpackInvalidValidatorAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["InvalidValidatorStatus"].ID.Bytes()[:4]) {
		return stakeHub.UnpackInvalidValidatorStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["InvalidValue"].ID.Bytes()[:4]) {
		return stakeHub.UnpackInvalidValueError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["JailedValidator"].ID.Bytes()[:4]) {
		return stakeHub.UnpackJailedValidatorError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["NotEnoughAmount"].ID.Bytes()[:4]) {
		return stakeHub.UnpackNotEnoughAmountError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["NotEnoughShares"].ID.Bytes()[:4]) {
		return stakeHub.UnpackNotEnoughSharesError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["NotEnoughValidators"].ID.Bytes()[:4]) {
		return stakeHub.UnpackNotEnoughValidatorsError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return stakeHub.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["NotJailed"].ID.Bytes()[:4]) {
		return stakeHub.UnpackNotJailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["OnlyCoinbase"].ID.Bytes()[:4]) {
		return stakeHub.UnpackOnlyCoinbaseError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["OnlyKeeper"].ID.Bytes()[:4]) {
		return stakeHub.UnpackOnlyKeeperError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["OnlySystemContract"].ID.Bytes()[:4]) {
		return stakeHub.UnpackOnlySystemContractError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["OnlyZeroGasPrice"].ID.Bytes()[:4]) {
		return stakeHub.UnpackOnlyZeroGasPriceError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["ReentrancyGuardReentrantCall"].ID.Bytes()[:4]) {
		return stakeHub.UnpackReentrancyGuardReentrantCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["SafeCastOverflowedUintDowncast"].ID.Bytes()[:4]) {
		return stakeHub.UnpackSafeCastOverflowedUintDowncastError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["StillInJailPeriod"].ID.Bytes()[:4]) {
		return stakeHub.UnpackStillInJailPeriodError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["TooFrequentUpdate"].ID.Bytes()[:4]) {
		return stakeHub.UnpackTooFrequentUpdateError(raw[4:])
	}
	if bytes.Equal(raw[:4], stakeHub.abi.Errors["ValidatorNotExists"].ID.Bytes()[:4]) {
		return stakeHub.UnpackValidatorNotExistsError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// StakeHubDuplicateId represents a DuplicateId error raised by the StakeHub contract.
type StakeHubDuplicateId struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DuplicateId()
func StakeHubDuplicateIdErrorID() common.Hash {
	return common.HexToHash("0xb1fb6bd0569a843fae5509e7270458906bc5e0359ad0fe264fabe5345bacf83d")
}

// UnpackDuplicateIdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DuplicateId()
func (stakeHub *StakeHub) UnpackDuplicateIdError(raw []byte) (*StakeHubDuplicateId, error) {
	out := new(StakeHubDuplicateId)
	if err := stakeHub.abi.UnpackIntoInterface(out, "DuplicateId", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubDuplicateOperatorAddress represents a DuplicateOperatorAddress error raised by the StakeHub contract.
type StakeHubDuplicateOperatorAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DuplicateOperatorAddress()
func StakeHubDuplicateOperatorAddressErrorID() common.Hash {
	return common.HexToHash("0x9cc20cedeb767979223b806ace03a4cd1b95ca667f7eb9f2be73281565920a41")
}

// UnpackDuplicateOperatorAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DuplicateOperatorAddress()
func (stakeHub *StakeHub) UnpackDuplicateOperatorAddressError(raw []byte) (*StakeHubDuplicateOperatorAddress, error) {
	out := new(StakeHubDuplicateOperatorAddress)
	if err := stakeHub.abi.UnpackIntoInterface(out, "DuplicateOperatorAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubDuplicateSignerAddress represents a DuplicateSignerAddress error raised by the StakeHub contract.
type StakeHubDuplicateSignerAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DuplicateSignerAddress()
func StakeHubDuplicateSignerAddressErrorID() common.Hash {
	return common.HexToHash("0x966a8ad3cdc93e00523c8216b0c27e60847a5a239de630a2a187f4632b639ca7")
}

// UnpackDuplicateSignerAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DuplicateSignerAddress()
func (stakeHub *StakeHub) UnpackDuplicateSignerAddressError(raw []byte) (*StakeHubDuplicateSignerAddress, error) {
	out := new(StakeHubDuplicateSignerAddress)
	if err := stakeHub.abi.UnpackIntoInterface(out, "DuplicateSignerAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubDuplicateValidatorAddress represents a DuplicateValidatorAddress error raised by the StakeHub contract.
type StakeHubDuplicateValidatorAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DuplicateValidatorAddress()
func StakeHubDuplicateValidatorAddressErrorID() common.Hash {
	return common.HexToHash("0x813f326766a6a6c709fc0f61143d5c4586a07d19cd7652a1fcbe25d4f7562c0e")
}

// UnpackDuplicateValidatorAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DuplicateValidatorAddress()
func (stakeHub *StakeHub) UnpackDuplicateValidatorAddressError(raw []byte) (*StakeHubDuplicateValidatorAddress, error) {
	out := new(StakeHubDuplicateValidatorAddress)
	if err := stakeHub.abi.UnpackIntoInterface(out, "DuplicateValidatorAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubEnforcedPause represents a EnforcedPause error raised by the StakeHub contract.
type StakeHubEnforcedPause struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EnforcedPause()
func StakeHubEnforcedPauseErrorID() common.Hash {
	return common.HexToHash("0xd93c0665d6c96d04a8f174024fc4ddd66c250604aff22bbec808de86dd3637e3")
}

// UnpackEnforcedPauseError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EnforcedPause()
func (stakeHub *StakeHub) UnpackEnforcedPauseError(raw []byte) (*StakeHubEnforcedPause, error) {
	out := new(StakeHubEnforcedPause)
	if err := stakeHub.abi.UnpackIntoInterface(out, "EnforcedPause", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubExpectedPause represents a ExpectedPause error raised by the StakeHub contract.
type StakeHubExpectedPause struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ExpectedPause()
func StakeHubExpectedPauseErrorID() common.Hash {
	return common.HexToHash("0x8dfc202bcfe9a735b559bee70674422512bc5c30f687046ae8778315fb81da44")
}

// UnpackExpectedPauseError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ExpectedPause()
func (stakeHub *StakeHub) UnpackExpectedPauseError(raw []byte) (*StakeHubExpectedPause, error) {
	out := new(StakeHubExpectedPause)
	if err := stakeHub.abi.UnpackIntoInterface(out, "ExpectedPause", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubFailedCall represents a FailedCall error raised by the StakeHub contract.
type StakeHubFailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func StakeHubFailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (stakeHub *StakeHub) UnpackFailedCallError(raw []byte) (*StakeHubFailedCall, error) {
	out := new(StakeHubFailedCall)
	if err := stakeHub.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubFailedDeployment represents a FailedDeployment error raised by the StakeHub contract.
type StakeHubFailedDeployment struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedDeployment()
func StakeHubFailedDeploymentErrorID() common.Hash {
	return common.HexToHash("0xb06ebf3d5067824a3fe5d5ba19471e035a7de6c88dac362c77b162830a5b9093")
}

// UnpackFailedDeploymentError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedDeployment()
func (stakeHub *StakeHub) UnpackFailedDeploymentError(raw []byte) (*StakeHubFailedDeployment, error) {
	out := new(StakeHubFailedDeployment)
	if err := stakeHub.abi.UnpackIntoInterface(out, "FailedDeployment", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubInBlackList represents a InBlackList error raised by the StakeHub contract.
type StakeHubInBlackList struct {
	Account common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InBlackList(address account)
func StakeHubInBlackListErrorID() common.Hash {
	return common.HexToHash("0x033788ffabe13708428b49b6cedeba51975d173696096da94c2097fc14da662e")
}

// UnpackInBlackListError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InBlackList(address account)
func (stakeHub *StakeHub) UnpackInBlackListError(raw []byte) (*StakeHubInBlackList, error) {
	out := new(StakeHubInBlackList)
	if err := stakeHub.abi.UnpackIntoInterface(out, "InBlackList", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubInsufficientBalance represents a InsufficientBalance error raised by the StakeHub contract.
type StakeHubInsufficientBalance struct {
	Balance *big.Int
	Needed  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InsufficientBalance(uint256 balance, uint256 needed)
func StakeHubInsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xcf4791818fba6e019216eb4864093b4947f674afada5d305e57d598b641dad1d")
}

// UnpackInsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InsufficientBalance(uint256 balance, uint256 needed)
func (stakeHub *StakeHub) UnpackInsufficientBalanceError(raw []byte) (*StakeHubInsufficientBalance, error) {
	out := new(StakeHubInsufficientBalance)
	if err := stakeHub.abi.UnpackIntoInterface(out, "InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubInvalidCommissionRate represents a InvalidCommissionRate error raised by the StakeHub contract.
type StakeHubInvalidCommissionRate struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidCommissionRate()
func StakeHubInvalidCommissionRateErrorID() common.Hash {
	return common.HexToHash("0x8fe552e0fc2e1a1a1656c6c0141e0073be5b3b424244bc17f9ee08d7f29ab97e")
}

// UnpackInvalidCommissionRateError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidCommissionRate()
func (stakeHub *StakeHub) UnpackInvalidCommissionRateError(raw []byte) (*StakeHubInvalidCommissionRate, error) {
	out := new(StakeHubInvalidCommissionRate)
	if err := stakeHub.abi.UnpackIntoInterface(out, "InvalidCommissionRate", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubInvalidId represents a InvalidId error raised by the StakeHub contract.
type StakeHubInvalidId struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidId()
func StakeHubInvalidIdErrorID() common.Hash {
	return common.HexToHash("0xdfa1a408860d09b3a48293a777aa052aa693dea19a077bbc934298b3ce4acd29")
}

// UnpackInvalidIdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidId()
func (stakeHub *StakeHub) UnpackInvalidIdError(raw []byte) (*StakeHubInvalidId, error) {
	out := new(StakeHubInvalidId)
	if err := stakeHub.abi.UnpackIntoInterface(out, "InvalidId", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubInvalidInitialization represents a InvalidInitialization error raised by the StakeHub contract.
type StakeHubInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func StakeHubInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (stakeHub *StakeHub) UnpackInvalidInitializationError(raw []byte) (*StakeHubInvalidInitialization, error) {
	out := new(StakeHubInvalidInitialization)
	if err := stakeHub.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubInvalidSignerAddress represents a InvalidSignerAddress error raised by the StakeHub contract.
type StakeHubInvalidSignerAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidSignerAddress()
func StakeHubInvalidSignerAddressErrorID() common.Hash {
	return common.HexToHash("0x4501a919f1e46d6aa4c5d40a66f88774f85a2c937171bc784733d47027bb9da7")
}

// UnpackInvalidSignerAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidSignerAddress()
func (stakeHub *StakeHub) UnpackInvalidSignerAddressError(raw []byte) (*StakeHubInvalidSignerAddress, error) {
	out := new(StakeHubInvalidSignerAddress)
	if err := stakeHub.abi.UnpackIntoInterface(out, "InvalidSignerAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubInvalidValidatorAddress represents a InvalidValidatorAddress error raised by the StakeHub contract.
type StakeHubInvalidValidatorAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidValidatorAddress()
func StakeHubInvalidValidatorAddressErrorID() common.Hash {
	return common.HexToHash("0x713ce511aa50868d0c2d8b07e0b3982819c17809cacda19b9e7a5e3b4d4bb677")
}

// UnpackInvalidValidatorAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidValidatorAddress()
func (stakeHub *StakeHub) UnpackInvalidValidatorAddressError(raw []byte) (*StakeHubInvalidValidatorAddress, error) {
	out := new(StakeHubInvalidValidatorAddress)
	if err := stakeHub.abi.UnpackIntoInterface(out, "InvalidValidatorAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubInvalidValidatorStatus represents a InvalidValidatorStatus error raised by the StakeHub contract.
type StakeHubInvalidValidatorStatus struct {
	Status uint8
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidValidatorStatus(uint8 status)
func StakeHubInvalidValidatorStatusErrorID() common.Hash {
	return common.HexToHash("0x5c3324cccfdef5487c5ba458c99fd8eac87de089ab6d0759af0d114eaad96a3e")
}

// UnpackInvalidValidatorStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidValidatorStatus(uint8 status)
func (stakeHub *StakeHub) UnpackInvalidValidatorStatusError(raw []byte) (*StakeHubInvalidValidatorStatus, error) {
	out := new(StakeHubInvalidValidatorStatus)
	if err := stakeHub.abi.UnpackIntoInterface(out, "InvalidValidatorStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubInvalidValue represents a InvalidValue error raised by the StakeHub contract.
type StakeHubInvalidValue struct {
	Key   string
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidValue(string key, uint256 value)
func StakeHubInvalidValueErrorID() common.Hash {
	return common.HexToHash("0x2c648cf1bbb773fa5fbcfc6541df5c1891af9b6969d9a555653bcec289d77218")
}

// UnpackInvalidValueError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidValue(string key, uint256 value)
func (stakeHub *StakeHub) UnpackInvalidValueError(raw []byte) (*StakeHubInvalidValue, error) {
	out := new(StakeHubInvalidValue)
	if err := stakeHub.abi.UnpackIntoInterface(out, "InvalidValue", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubJailedValidator represents a JailedValidator error raised by the StakeHub contract.
type StakeHubJailedValidator struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error JailedValidator()
func StakeHubJailedValidatorErrorID() common.Hash {
	return common.HexToHash("0x78b2a3366692ee51cd0e427429ed8d526c9090e05446ec218ab6a9eaa846431b")
}

// UnpackJailedValidatorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error JailedValidator()
func (stakeHub *StakeHub) UnpackJailedValidatorError(raw []byte) (*StakeHubJailedValidator, error) {
	out := new(StakeHubJailedValidator)
	if err := stakeHub.abi.UnpackIntoInterface(out, "JailedValidator", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubNotEnoughAmount represents a NotEnoughAmount error raised by the StakeHub contract.
type StakeHubNotEnoughAmount struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotEnoughAmount()
func StakeHubNotEnoughAmountErrorID() common.Hash {
	return common.HexToHash("0xe008b5f9775a2bbe458e4c5378d46e4d64183e024a4785f91837700c840acfee")
}

// UnpackNotEnoughAmountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotEnoughAmount()
func (stakeHub *StakeHub) UnpackNotEnoughAmountError(raw []byte) (*StakeHubNotEnoughAmount, error) {
	out := new(StakeHubNotEnoughAmount)
	if err := stakeHub.abi.UnpackIntoInterface(out, "NotEnoughAmount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubNotEnoughShares represents a NotEnoughShares error raised by the StakeHub contract.
type StakeHubNotEnoughShares struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotEnoughShares()
func StakeHubNotEnoughSharesErrorID() common.Hash {
	return common.HexToHash("0xf15ed2146de0776d0c3238708a746d883c49ffc423a1269fe2022bfaf59c7446")
}

// UnpackNotEnoughSharesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotEnoughShares()
func (stakeHub *StakeHub) UnpackNotEnoughSharesError(raw []byte) (*StakeHubNotEnoughShares, error) {
	out := new(StakeHubNotEnoughShares)
	if err := stakeHub.abi.UnpackIntoInterface(out, "NotEnoughShares", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubNotEnoughValidators represents a NotEnoughValidators error raised by the StakeHub contract.
type StakeHubNotEnoughValidators struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotEnoughValidators()
func StakeHubNotEnoughValidatorsErrorID() common.Hash {
	return common.HexToHash("0xae575a8893fe8a557a7152418b327ad764f6981a872ee0ffd0046c6e1fcc52c4")
}

// UnpackNotEnoughValidatorsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotEnoughValidators()
func (stakeHub *StakeHub) UnpackNotEnoughValidatorsError(raw []byte) (*StakeHubNotEnoughValidators, error) {
	out := new(StakeHubNotEnoughValidators)
	if err := stakeHub.abi.UnpackIntoInterface(out, "NotEnoughValidators", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubNotInitializing represents a NotInitializing error raised by the StakeHub contract.
type StakeHubNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func StakeHubNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (stakeHub *StakeHub) UnpackNotInitializingError(raw []byte) (*StakeHubNotInitializing, error) {
	out := new(StakeHubNotInitializing)
	if err := stakeHub.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubNotJailed represents a NotJailed error raised by the StakeHub contract.
type StakeHubNotJailed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotJailed()
func StakeHubNotJailedErrorID() common.Hash {
	return common.HexToHash("0xf9200d4f9403acb31301fdd63749a1f0c233891cee0618b3f06f5f9768d0c018")
}

// UnpackNotJailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotJailed()
func (stakeHub *StakeHub) UnpackNotJailedError(raw []byte) (*StakeHubNotJailed, error) {
	out := new(StakeHubNotJailed)
	if err := stakeHub.abi.UnpackIntoInterface(out, "NotJailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubOnlyCoinbase represents a OnlyCoinbase error raised by the StakeHub contract.
type StakeHubOnlyCoinbase struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyCoinbase()
func StakeHubOnlyCoinbaseErrorID() common.Hash {
	return common.HexToHash("0x116c64a8de02ce00303a109e06f26c31a3cfed64656fb9d228157fad57dff616")
}

// UnpackOnlyCoinbaseError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyCoinbase()
func (stakeHub *StakeHub) UnpackOnlyCoinbaseError(raw []byte) (*StakeHubOnlyCoinbase, error) {
	out := new(StakeHubOnlyCoinbase)
	if err := stakeHub.abi.UnpackIntoInterface(out, "OnlyCoinbase", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubOnlyKeeper represents a OnlyKeeper error raised by the StakeHub contract.
type StakeHubOnlyKeeper struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyKeeper()
func StakeHubOnlyKeeperErrorID() common.Hash {
	return common.HexToHash("0xc60eb3352805112c61799a78f842543bbf71f1a5e9cd075fb2e23066f7db6914")
}

// UnpackOnlyKeeperError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyKeeper()
func (stakeHub *StakeHub) UnpackOnlyKeeperError(raw []byte) (*StakeHubOnlyKeeper, error) {
	out := new(StakeHubOnlyKeeper)
	if err := stakeHub.abi.UnpackIntoInterface(out, "OnlyKeeper", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubOnlySystemContract represents a OnlySystemContract error raised by the StakeHub contract.
type StakeHubOnlySystemContract struct {
	ContractAddr common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlySystemContract(address contractAddr)
func StakeHubOnlySystemContractErrorID() common.Hash {
	return common.HexToHash("0xf22c4390ebf387afc834c245f4ef6f38d59454737d03e451e0d182ac61608277")
}

// UnpackOnlySystemContractError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlySystemContract(address contractAddr)
func (stakeHub *StakeHub) UnpackOnlySystemContractError(raw []byte) (*StakeHubOnlySystemContract, error) {
	out := new(StakeHubOnlySystemContract)
	if err := stakeHub.abi.UnpackIntoInterface(out, "OnlySystemContract", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubOnlyZeroGasPrice represents a OnlyZeroGasPrice error raised by the StakeHub contract.
type StakeHubOnlyZeroGasPrice struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyZeroGasPrice()
func StakeHubOnlyZeroGasPriceErrorID() common.Hash {
	return common.HexToHash("0x83f1b1d3f8cc3470adc53b59d7024697cf0374e9ddadd2111159d00fe76f2c06")
}

// UnpackOnlyZeroGasPriceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyZeroGasPrice()
func (stakeHub *StakeHub) UnpackOnlyZeroGasPriceError(raw []byte) (*StakeHubOnlyZeroGasPrice, error) {
	out := new(StakeHubOnlyZeroGasPrice)
	if err := stakeHub.abi.UnpackIntoInterface(out, "OnlyZeroGasPrice", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubReentrancyGuardReentrantCall represents a ReentrancyGuardReentrantCall error raised by the StakeHub contract.
type StakeHubReentrancyGuardReentrantCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ReentrancyGuardReentrantCall()
func StakeHubReentrancyGuardReentrantCallErrorID() common.Hash {
	return common.HexToHash("0x3ee5aeb571de7fc460830b4d0017439a1ca56fb0bc39062227ade4fe4a24c1ca")
}

// UnpackReentrancyGuardReentrantCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ReentrancyGuardReentrantCall()
func (stakeHub *StakeHub) UnpackReentrancyGuardReentrantCallError(raw []byte) (*StakeHubReentrancyGuardReentrantCall, error) {
	out := new(StakeHubReentrancyGuardReentrantCall)
	if err := stakeHub.abi.UnpackIntoInterface(out, "ReentrancyGuardReentrantCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubSafeCastOverflowedUintDowncast represents a SafeCastOverflowedUintDowncast error raised by the StakeHub contract.
type StakeHubSafeCastOverflowedUintDowncast struct {
	Bits  uint8
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func StakeHubSafeCastOverflowedUintDowncastErrorID() common.Hash {
	return common.HexToHash("0x6dfcc6503a32754ce7a89698e18201fc5294fd4aad43edefee786f88423b1a12")
}

// UnpackSafeCastOverflowedUintDowncastError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func (stakeHub *StakeHub) UnpackSafeCastOverflowedUintDowncastError(raw []byte) (*StakeHubSafeCastOverflowedUintDowncast, error) {
	out := new(StakeHubSafeCastOverflowedUintDowncast)
	if err := stakeHub.abi.UnpackIntoInterface(out, "SafeCastOverflowedUintDowncast", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubStillInJailPeriod represents a StillInJailPeriod error raised by the StakeHub contract.
type StakeHubStillInJailPeriod struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StillInJailPeriod()
func StakeHubStillInJailPeriodErrorID() common.Hash {
	return common.HexToHash("0x197d095ce12fcbcecb760823b2e9199a8b252bc1aae24c14e24b5670f2738dff")
}

// UnpackStillInJailPeriodError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StillInJailPeriod()
func (stakeHub *StakeHub) UnpackStillInJailPeriodError(raw []byte) (*StakeHubStillInJailPeriod, error) {
	out := new(StakeHubStillInJailPeriod)
	if err := stakeHub.abi.UnpackIntoInterface(out, "StillInJailPeriod", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubTooFrequentUpdate represents a TooFrequentUpdate error raised by the StakeHub contract.
type StakeHubTooFrequentUpdate struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TooFrequentUpdate()
func StakeHubTooFrequentUpdateErrorID() common.Hash {
	return common.HexToHash("0xc85dc358bb7c88429397fc64ed42d3cef735fef6b38bb9aae96e5c1d3e864224")
}

// UnpackTooFrequentUpdateError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TooFrequentUpdate()
func (stakeHub *StakeHub) UnpackTooFrequentUpdateError(raw []byte) (*StakeHubTooFrequentUpdate, error) {
	out := new(StakeHubTooFrequentUpdate)
	if err := stakeHub.abi.UnpackIntoInterface(out, "TooFrequentUpdate", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StakeHubValidatorNotExists represents a ValidatorNotExists error raised by the StakeHub contract.
type StakeHubValidatorNotExists struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ValidatorNotExists()
func StakeHubValidatorNotExistsErrorID() common.Hash {
	return common.HexToHash("0x5926e0c333f1cbbc6225b5f94c36bb66fcde4ead7b7b00e0b28ebb7bcf906e5e")
}

// UnpackValidatorNotExistsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ValidatorNotExists()
func (stakeHub *StakeHub) UnpackValidatorNotExistsError(raw []byte) (*StakeHubValidatorNotExists, error) {
	out := new(StakeHubValidatorNotExists)
	if err := stakeHub.abi.UnpackIntoInterface(out, "ValidatorNotExists", raw); err != nil {
		return nil, err
	}
	return out, nil
}
