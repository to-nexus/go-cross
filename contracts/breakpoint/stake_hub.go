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
	ABI:        "[{\"type\":\"function\",\"name\":\"MAX_COMMISSION_RATE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPDATE_FREQUENCY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addToBlackList\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyValidator\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"commissionRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"id\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"claimUndelegation\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimUndelegation\",\"inputs\":[{\"name\":\"id\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimUnstaked\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegate\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"delegate\",\"inputs\":[{\"name\":\"id\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"distributeReward\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getIds\",\"inputs\":[{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"ids\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"totalLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getKeeper\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperators\",\"inputs\":[{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"operatorAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"totalLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorAddress\",\"inputs\":[{\"name\":\"id\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorAddress\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorDescription\",\"inputs\":[{\"name\":\"id_\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"commissionRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"id\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorDescription\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"commissionRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"id\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorInfo\",\"inputs\":[{\"name\":\"id\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shareContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isJailed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isStaked\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorInfo\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shareContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isJailed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isStaked\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorJailInfo\",\"inputs\":[{\"name\":\"id\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"jailUntil\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isJailed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorJailInfo\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"jailUntil\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isJailed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorTotalStake\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorTotalStake\",\"inputs\":[{\"name\":\"id\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorUnstakingInfo\",\"inputs\":[{\"name\":\"id\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"unstakeTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isUnstaking\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isUnstaked\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorUnstakingInfo\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"unstakeTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isUnstaking\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isUnstaked\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidators\",\"inputs\":[{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"validatorAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"totalLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"idToOperator\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isActiveValidator\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isActiveValidator\",\"inputs\":[{\"name\":\"id\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minDelegateAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minStakeAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"offlineJailDuration\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"offlineSlashAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeFromBlackList\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setKeeper\",\"inputs\":[{\"name\":\"keeper\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinDelegateAmount\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinStakeAmount\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOfflineJailDuration\",\"inputs\":[{\"name\":\"duration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOfflineSlashAmount\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUnbondingPeriod\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setValidatorThreshold\",\"inputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashOffline\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalSlashedAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unbondingPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"undelegate\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"undelegate\",\"inputs\":[{\"name\":\"id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unjail\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unstakeValidator\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateValidator\",\"inputs\":[{\"name\":\"newValidatorAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newCommissionRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validatorThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatorToOperator\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"BlackListed\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Claimed\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Delegated\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"KeeperChanged\",\"inputs\":[{\"name\":\"previousKeeper\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newKeeper\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"KeeperUpdated\",\"inputs\":[{\"name\":\"keeper\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinDelegateAmountUpdated\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinStakeAmountUpdated\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OfflineJailDurationUpdated\",\"inputs\":[{\"name\":\"duration\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OfflineSlashAmountUpdated\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardDistributed\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnBlackListed\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnbondingPeriodUpdated\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Undelegated\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorApplied\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"shareContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorJailed\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"jailUntil\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorSlashed\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"slashAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorThresholdUpdated\",\"inputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorUnjailed\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorUnstaked\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"unstakeTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorUpdated\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newValidatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newCommissionRate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"DuplicateId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DuplicateOperatorAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DuplicateValidatorAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedDeployment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InBlackList\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidCommissionRate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValidatorAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValidatorStatus\",\"inputs\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumStakeHub.ValidatorStatus\"}]},{\"type\":\"error\",\"name\":\"InvalidValue\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"JailedValidator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughShares\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughValidators\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotJailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCoinbase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyKeeper\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySystemContract\",\"inputs\":[{\"name\":\"contractAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyZeroGasPrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"StillInJailPeriod\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TooFrequentUpdate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValidatorNotExists\",\"inputs\":[]}]",
	ID:         "StakeHub",
	BinRuntime: "0x60806040526004361015610011575f80fd5b5f3560e01c8063092193ab146103945780630e222b361461038f578063114eaf551461038a578063176763aa14610385578063207239c0146103805780632aa8e06d1461037b57806335efc73414610376578063391b6f4e146103715780633e732eba1461036c5780633f4ba83a1461036757806340550a1c14610362578063416985731461035d578063417c73a71461035857806342fce3cd1461035357806343a769d91461034e57806344690e5c146103495780634a49ac4c146103445780634d99dd161461033f5780634fd101d71461033a5780635bc0fbc7146103355780635c19a95c146103305780635c975abb1461032b578063604ababc1461032657806363973e4b146103215780636cf6d6751461031c5780636f7aa7ec146103175780637071688a1461031257806372e9c9341461030d578063748747e6146103085780638129fc1c146103035780638456cb59146102fe5780638661213a146102f95780638a11d7c9146102f45780638c7d0dbc146102ef5780638dfc8897146102ea578063935e08d6146102e55780639679904a146102e05780639ac83d85146102db5780639ddb511a146102d6578063a43569b3146102d1578063a4cac964146102cc578063a9cc60f3146102c7578063ba723eda146102c2578063bbe5e345146102bd578063bff02e20146102b8578063c470ed65146102b3578063d1379332146102ae578063d276290d146102a9578063d86c7fe8146102a4578063db6c4c061461029f578063e3f95ed91461029a578063e6635f5d14610295578063ea4dd2b914610290578063eb4af0451461028b578063f1887684146102865763f679d30514610281575f80fd5b612160565b612143565b6120b2565b612072565b611e4a565b611df0565b611dd3565b611ba9565b611a00565b6119c0565b61197b565b611900565b611878565b61185b565b61183e565b61181a565b6117f9565b6117a4565b611772565b6116dd565b611657565b6115d7565b61159f565b61157b565b61153a565b6114ab565b611345565b611286565b611269565b61124d565b611209565b6111ec565b611196565b610dbf565b610d91565b610d52565b610cbd565b610ca0565b610c46565b610bcf565b610a03565b61096d565b610950565b6108d6565b610862565b610821565b610786565b6106f2565b6106be565b61067e565b61061b565b6105ff565b6105b1565b6104cc565b610457565b6103b3565b600435906001600160a01b03821682036103af57565b5f80fd5b60203660031901126103af576103c7610399565b4133036103ed573a6103de576103dc90612251565b005b6383f1b1d360e01b5f5260045ffd5b63022d8c9560e31b5f5260045ffd5b9181601f840112156103af578235916001600160401b0383116103af57602083818601950101116103af57565b60206003198201126103af57600435906001600160401b0382116103af57610453916004016103fc565b9091565b346103af5761046f61046836610429565b369161240c565b80516020918201205f908152600482526040808220546001600160a01b03168252600292839052902060038101549101546001600160401b03919091169060e81c60ff165b604080519283529015156020830152819081015b0390f35b346103af5760203660031901126103af57600435611013330361055e578015610520576020817f38a1644f59872db6fd17fdced395495fcaa3ca7d2825a0704db5a3acbd1dd06492600955604051908152a1005b60849060405190632c648cf160e01b825260406004830152600f60448301526e1d5b989bdb991a5b99d4195c9a5bd9608a1b60648301526024820152fd5b630f22c43960e41b5f5261101360045260245ffd5b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b6040906105ae939281528160208201520190610573565b90565b346103af576105c261046836610429565b602081519101205f5260046020526105e560018060a01b0360405f205416612867565b906104c860405192839283610597565b5f9103126103af57565b346103af575f3660031901126103af5760206040516113888152f35b346103af5760203660031901126103af57602061063e610639610399565b612451565b604051908152f35b6001600160a01b03165f9081527f46944d7356e409d0d8c174d262d4ab0ddb2d4597e3e424151a463d8a4af4e5016020526040902090565b346103af5760203660031901126103af576001600160a01b0361069f610399565b165f526003602052602060018060a01b0360405f205416604051908152f35b346103af575f3660031901126103af575f5160206135475f395f51905f52546040516001600160a01b039091168152602090f35b346103af5760203660031901126103af57600435611013330361055e578015610746576020817f02cd8ef316564ca78b75bf239c0a630008374c1fb1d26d941a6e9b19e42b2aa592600755604051908152a1005b60849060405190632c648cf160e01b82526040600483015260116044830152701b5a5b91195b1959d85d19505b5bdd5b9d607a1b60648301526024820152fd5b346103af575f3660031901126103af575f5160206135475f395f51905f52546001600160a01b03163303610812576107bc612bff565b6107c4612bff565b60ff195f5160206135675f395f51905f5254165f5160206135675f395f51905f52557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b63c60eb33560e01b5f5260045ffd5b346103af5760203660031901126103af576001600160a01b03610842610399565b165f526002602052602061085860405f20612c27565b6040519015158152f35b346103af5761087361046836610429565b602081519101205f5260046020526104c861089960018060a01b0360405f2054166125bb565b604080516001600160a01b0396871681529590941660208601529115159284019290925290151560608301521515608082015290819060a0820190565b346103af5760203660031901126103af576108ef610399565b5f5160206135475f395f51905f52546001600160a01b031633036108125761091681610646565b805460ff191660011790556001600160a01b03167f7fd26be6fc92aff63f1f4409b2b2ddeb272a888031d7f55ec830485ec61941865f80a2005b346103af575f3660031901126103af576020600b54604051908152f35b346103af5760203660031901126103af57600435611013330361055e5780156109c1576020817fbdf592a42adc0296ad8a75a5e7b8a0912c37793a1a275d78275bc335911c6bd392600b55604051908152a1005b60849060405190632c648cf160e01b825260406004830152601360448301527237b3333634b732a530b4b6223ab930ba34b7b760691b60648301526024820152fd5b346103af5760203660031901126103af57610a1c610399565b6110043303610bba576001600160a01b03165f908152600360205260409020610a4d905b546001600160a01b031690565b610a5681612c71565b6001600160a01b0381165f9081526002602052604081206002810154602090610abd90610a9990610a8d906001600160a01b031681565b6001600160a01b031690565b600a5460405195868094819363045bc4d160e41b8352600483019190602083019252565b03925af18015610b8657610b30925f91610b8b575b50610ae7610ae2826005546124f8565b600555565b6040519081526001600160a01b038416907f17bddadfd7ec8898c3b9eadd0cf5ae77ba8d5df3a50e96ab86ec2dd711aa8fbb90602090a2610b2a600b54426124f8565b90612cb7565b6110113b156103af57604051632961046560e21b81526001600160a01b0390911660048201525f8180602481015b0381836110115af18015610b8657610b7257005b80610b805f6103dc93612220565b806105f5565b612246565b610bad915060203d602011610bb3575b610ba58183612220565b810190612442565b5f610ad2565b503d610b9b565b630f22c43960e41b5f5261100460045260245ffd5b346103af5760203660031901126103af57610be8610399565b5f5160206135475f395f51905f52546001600160a01b0316330361081257610c0f81610646565b805460ff191690556001600160a01b03167fe0db3499b7fdc3da4cddff5f45d694549c19835e7f719fb5606d3ad1a5de40115f80a2005b346103af5760403660031901126103af57610c5f610399565b602435610c6a612dbc565b60ff610c7533610646565b5416610c8d5781610c886103dc93612c71565b612de3565b63033788ff60e01b5f523360045260245ffd5b346103af575f3660031901126103af576020600854604051908152f35b346103af5760203660031901126103af57600435611013330361055e578015610d11576020817f324143af489ab0ba77b0d3580e2486ee612b673b87db2179c589f61ac532fd5992600855604051908152a1005b60849060405190632c648cf160e01b82526040600483015260126044830152711d985b1a59185d1bdc951a1c995cda1bdb1960721b60648301526024820152fd5b60203660031901126103af57610d66610399565b610d6e612dbc565b60ff610d7933610646565b5416610c8d5780610d8c6103dc92612c71565b612f3b565b346103af575f3660031901126103af57602060ff5f5160206135675f395f51905f5254166040519015158152f35b60603660031901126103af57610dd3610399565b6024356044356001600160401b0381116103af57610df59036906004016103fc565b91610dfe612dbc565b60ff610e0933610646565b5416610c8d576001600160a01b03841692831561118757600654341061117857610e3233610a8d565b94610e48865f52600160205260405f2054151590565b611169576001600160a01b0381165f908152600360205260409020610e7090610a8d90610a40565b61115a57611388831161114b57610e90610e8b36848761240c565b613089565b610ea8610e9e36848761240c565b6020815191012090565b93610ec1610a8d610a40875f52600460205260405f2090565b61113c5761105f94610eed610ee7610ee1610f3e9461105197369161240c565b3361315e565b986134c3565b506001600160a01b0383165f908152600360205260409020610f2b9033905b80546001600160a01b0319166001600160a01b03909216919091179055565b610f0c33915f52600460205260405f2090565b335f908152600260205260409020610f6f9080546001600160a01b0319166001600160a01b03909316929092178255565b6001810180546001600160a01b0319163317815561102e906002830180546001600160a01b0319166001600160a01b038a161781559461101a90610fb79061321a565b61321a565b93610fe96003820195869067ffffffffffffffff60401b82549160401b169067ffffffffffffffff60401b1916179055565b610ff24261321a565b815467ffffffffffffffff60a01b191660a09190911b67ffffffffffffffff60a01b16179055565b805467ffffffffffffffff60a01b19169055565b825467ffffffffffffffff60a01b19168355805467ffffffffffffffff19169055565b805461ffff60e01b19169055565b604051916001600160a01b03169033907f02a48e2789d19f84e9cfb18b02359e63c1e9779fa51ff1ae65bba5ab0a2739a35f80a4348082526020820152339081907f24d7bda8602b916d64417f0dbfe2e2e88ec9b1157bd9f596dfdb91ba26624e0490604090a36110113b156103af57604051632961046560e21b81523360048201525f81602481836110115af18015610b8657611128575b506110113b156103af57604051632949582d60e01b8152336004820181905260248201525f818060448101610b5e565b80610b805f61113693612220565b5f6110f8565b630b1fb6bd60e41b5f5260045ffd5b63047f2a9760e51b5f5260045ffd5b63813f326760e01b5f5260045ffd5b639cc20ced60e01b5f5260045ffd5b63e008b5f960e01b5f5260045ffd5b63713ce51160e01b5f5260045ffd5b346103af576111a761046836610429565b602081519101205f5260046020526104c86111cd60018060a01b0360405f205416612915565b6040805193845291151560208401521515908201529081906060820190565b346103af575f3660031901126103af576020600954604051908152f35b346103af5760203660031901126103af57611222610399565b61122a612dbc565b60ff61123533610646565b5416610c8d57806112486103dc92612c71565b61324b565b346103af575f3660031901126103af5760205f54604051908152f35b346103af575f3660031901126103af576020600754604051908152f35b346103af5760203660031901126103af5761129f610399565b611013330361055e576001600160a01b03168015611311575f5160206135475f395f51905f5254816001600160a01b0382167f068b48a2fe7f498b57ff6da64f075ae658fde8d77124b092e62b3dc58d91ce355f80a36001600160a01b031916175f5160206135475f395f51905f5255005b6084604051632c648cf160e01b815260406004820152600660448201526535b2b2b832b960d11b60648201525f6024820152fd5b346103af575f3660031901126103af575f5160206135875f395f51905f52546001600160401b0361138e60ff604084901c1615611381565b1590565b926001600160401b031690565b16801590816114a3575b6001149081611499575b159081611490575b5061148157806113e060016001600160401b03195f5160206135875f395f51905f525416175f5160206135875f395f51905f5255565b61144c575b6113ed612523565b6113f357005b61141d60ff60401b195f5160206135875f395f51905f5254165f5160206135875f395f51905f5255565b604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b61147c600160401b60ff60401b195f5160206135875f395f51905f525416175f5160206135875f395f51905f5255565b6113e5565b63f92ee8a960e01b5f5260045ffd5b9050155f6113aa565b303b1591506113a2565b829150611398565b346103af575f3660031901126103af575f5160206135475f395f51905f52546001600160a01b03163303610812576114e1612dbc565b6114e9612dbc565b600160ff195f5160206135675f395f51905f525416175f5160206135675f395f51905f52557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b346103af5761154b61046836610429565b80516020918201205f908152600482526040808220546001600160a01b0316825260028352902061085890612c27565b346103af5760203660031901126103af576104c861089961159a610399565b6125bb565b346103af576115b061046836610429565b80516020918201205f90815260048252604090205461063e906001600160a01b0316612451565b346103af5760403660031901126103af576004356001600160401b0381116103af576116079036906004016103fc565b60243590611613612dbc565b60ff61161e33610646565b5416610c8d576103dc9261163391369161240c565b602081519101205f52600460205260018060a01b0360405f205416610c8881612c71565b346103af5760403660031901126103af57611676602435600435612751565b9060405190604082016040835281518091526060830190602060608260051b8601019301915f905b8282106116b2578580868960208301520390f35b909192936020806116cf600193605f198a82030186528851610573565b96019201920190929161169e565b346103af5760203660031901126103af57600435611013330361055e578015611731576020817fff6e1245b60ea5b68ddca24a93cb8707677289cbdbdec9ecb7ccfc25541a774992600a55604051908152a1005b60849060405190632c648cf160e01b82526040600483015260126044830152711bd9999b1a5b9954db185cda105b5bdd5b9d60721b60648301526024820152fd5b346103af5760203660031901126103af576004355f526004602052602060018060a01b0360405f205416604051908152f35b6117ad36610429565b906117b6612dbc565b60ff6117c133610646565b5416610c8d576117d291369161240c565b602081519101205f5260046020526103dc60018060a01b0360405f205416610d8c81612c71565b346103af5760203660031901126103af576105e5611815610399565b612867565b346103af5760203660031901126103af576104c86111cd611839610399565b612915565b346103af575f3660031901126103af576020600a54604051908152f35b346103af575f3660031901126103af576020600554604051908152f35b346103af5760203660031901126103af576104b4611894610399565b60018060a01b03165f52600260205260405f209060ff60026001600160401b0360038501541693015460e81c1690565b90602080835192838152019201905f5b8181106118e15750505090565b82516001600160a01b03168452602093840193909201916001016118d4565b346103af5760403660031901126103af5761191f6024356004356129ac565b9190611936604051926060845260608401906118c4565b9282840360208401526020808351958681520192015f945b80861061196357505082935060408301520390f35b9092602080600192865181520194019501949061194e565b346103af5761198c61046836610429565b80516020918201205f908152600482526040808220546001600160a01b039081168352600284529181902054905191168152f35b346103af5760203660031901126103af576001600160a01b036119e1610399565b165f526002602052602060018060a01b0360405f205416604051908152f35b346103af575f3660031901126103af57611a18612dbc565b60ff611a2333610646565b5416610c8d57611a3233612c71565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005c611b9a5760017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005d335f90815260026020819052604090912001805460e081901c60ff1690611aa282612505565b60018203611b8757611ac79060a01c6001600160401b03165b6001600160401b031690565b4210611b7157815460ff60e01b1916600160e11b1782555f6020611af6610a8d8086546001600160a01b031690565b600460405180948193630f8f004960e31b83525af1908115610b86575f91611b52575b50604051908152339081907ff7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd399268390602090a36103dc613334565b611b6b915060203d602011610bb357610ba58183612220565b81611b19565b63170cc93360e21b5f52611b8490612ae7565b5ffd5b63170cc93360e21b5f52611b8482612ae7565b633ee5aeb560e01b5f5260045ffd5b346103af575f3660031901126103af57611bc1612dbc565b60ff611bcc33610646565b5416610c8d57611bdb33612c71565b335f90815260026020526040902060028101805460e081901c60ff16611c0081612505565b80611b71575060e81c60ff16611dc457611c1b61137d613359565b611db557805460ff60e01b1916600160e01b178155611c6a611c42610fb2600954426124f8565b825467ffffffffffffffff60a01b191660a09190911b67ffffffffffffffff60a01b16178255565b8154611c8a906001600160a01b0316915460a01c6001600160401b031690565b6040516001600160401b0391909116815233916001600160a01b0316907fc88e6a643abb00b7da809fc581edda961c3ea06e7b32c1bce7e7671f37a77eed90602090a36110113b156103af57604051632961046560e21b81523360048201525f81602481836110115af18015610b8657611da1575b506110113b156103af57604051632949582d60e01b81523360048201525f6024820181905290919082604481836110115af1918215610b8657611d4f92611d8d575b50546001600160a01b031690565b6110013b156103af5760405163f2888dbb60e01b81526001600160a01b039190911660048201525f81602481836110015af18015610b8657610b7257005b80610b805f611d9b93612220565b5f611d41565b80610b805f611daf93612220565b5f611cff565b6315caeb5160e31b5f5260045ffd5b633c59519b60e11b5f5260045ffd5b346103af575f3660031901126103af576020604051620151808152f35b346103af57611dfe36610429565b90611e07612dbc565b60ff611e1233610646565b5416610c8d57611e2391369161240c565b602081519101205f5260046020526103dc60018060a01b0360405f20541661124881612c71565b346103af5760403660031901126103af57611e63610399565b602435611e6e612dbc565b60ff611e7933610646565b5416610c8d57611e8833612c71565b335f908152600260205260409020916001830190611ec26001600160401b03611ebc84546001600160401b039060a01c1690565b166124e3565b4210612063576001600160a01b0381168015159081612044575b50611fd8575b5081151580611fb5575b611f6c575b611eff9150610ff24261321a565b8054611f24906003906001600160a01b03165b92015460401c6001600160401b031690565b6040516001600160401b039190911681526001600160a01b03919091169033907f1a4d5ce1be705973f55eb762d6ca5ed08d73caab380fa407cfaf0c1fa64a2ae490602090a3005b611388821161114b57611fb0611f84611eff9361321a565b600385019067ffffffffffffffff60401b82549160401b169067ffffffffffffffff60401b1916179055565b611ef1565b506003830154611fd09060401c6001600160401b0316611abb565b821415611eec565b6001600160a01b0381165f908152600360205260409020611ffc90610a8d90610a40565b61115a5761203e9061202233610f0c8360018060a01b03165f52600360205260405f2090565b84546001600160a01b0319166001600160a01b03909116178455565b5f611ee2565b855490915061205b906001600160a01b0316610a8d565b14155f611edc565b63190bb86b60e31b5f5260045ffd5b346103af5760403660031901126103af576120a8612094602435600435612af9565b6040519283926040845260408401906118c4565b9060208301520390f35b346103af5760203660031901126103af57600435611013330361055e578015612106576020817f8448c02797b448f4946bc25b3bf925e5556d1df822c944da701c54bab8a3162f92600655604051908152a1005b60849060405190632c648cf160e01b825260406004830152600e60448301526d1b5a5b94dd185ad9505b5bdd5b9d60921b60648301526024820152fd5b346103af575f3660031901126103af576020600654604051908152f35b346103af575f3660031901126103af57612178612dbc565b60ff61218333610646565b5416610c8d5761219233612c71565b335f52600260205260405f206121b361137d600283015460ff9060e81c1690565b6121fd576121ce611abb60038301546001600160401b031690565b42106121ee576121dd816133e2565b60065411611178576103dc9061346c565b63065f425760e21b5f5260045ffd5b63f9200d4f60e01b5f5260045ffd5b634e487b7160e01b5f52604160045260245ffd5b90601f801991011681019081106001600160401b0382111761224157604052565b61220c565b6040513d5f823e3d90fd5b6001600160a01b03165f90815260036020526040902061227090610a40565b6001600160a01b0381165f9081526002602052604090206002810180546001600160a01b0380821616159081156123e3575b506123a057546122c390600390611f1290610a8d906001600160a01b031681565b90803b156103af57604051632f303ebb60e11b81526001600160401b039290921660048301525f908290602490829034905af18015610b865761238c575b506040513481526001600160a01b038216907fe34918ff1c7084970068b53fd71ad6d8b04e9f15d3886cbf006443e6cdc52ea690602090a26110113b156103af57604051632961046560e21b81526001600160a01b0390911660048201525f8180602481015b0381836110115af18015610b865761237c5750565b80610b805f61238a93612220565b565b80610b805f61239a93612220565b5f612301565b5050506123ac34612b81565b604051348152611005907fe34918ff1c7084970068b53fd71ad6d8b04e9f15d3886cbf006443e6cdc52ea69080602081015b0390a2565b60e81c60ff1690505f6122a2565b6001600160401b03811161224157601f01601f191660200190565b929192612418826123f1565b916124266040519384612220565b8294818452818301116103af578281602093845f960137010152565b908160209103126103af575190565b6001600160a01b039081165f90815260026020819052604090912001541680156124ca5760049060209061248d906001600160a01b0316610a8d565b6040516278744560e21b815292839182905afa908115610b86575f916124b1575090565b6105ae915060203d602011610bb357610ba58183612220565b505f90565b634e487b7160e01b5f52601160045260245ffd5b906201518082018092116124f357565b6124cf565b919082018092116124f357565b6003111561250f57565b634e487b7160e01b5f52602160045260245ffd5b4133036103ed573a6103de5769021e19e0c9bab2400000600655678ac7230489e8000060075560156008556201518060095568056bc75e2d63100000600a5562015180600b5561257161351b565b61257961351b565b5f5160206135475f395f51905f5280546001600160a01b03191673e4dd1087233043dca7532c4fb9d8a31fe99b731e1790556125b361351b565b61238a61351b565b6001600160a01b039081165f90815260026020819052604090912080549181015482841694938116939092916125f090612c27565b926001600160401b0360ff8260e81c169360a01c1615159081612611575090565b60ff915060e01c1661137d81612505565b6001600160401b0381116122415760051b60200190565b60405190612648602083612220565b5f80835282815b82811061265b57505050565b80606060208093850101520161264f565b9061267682612622565b6126836040519182612220565b8281528092612694601f1991612622565b01905f5b8281106126a457505050565b806060602080938501015201612698565b919082039182116124f357565b6020818303126103af578051906001600160401b0382116103af570181601f820112156103af578051906126f5826123f1565b926127036040519485612220565b828452602083830101116103af57815f9260208093018386015e8301015290565b634e487b7160e01b5f52603260045260245ffd5b805182101561274c5760209160051b010190565b612724565b5f549291838210156128595780612850575061276e835b826124f8565b81848211612848575b612780916126b5565b6127898161266c565b915f5b82811061279a575050509190565b805f6127e5610a8d610a8d60026127d76127be6127b96004998b6124f8565b6132f9565b6001600160a01b03165f90815260026020526040902090565b01546001600160a01b031690565b6040516395d89b4160e01b815293849182905afa8015610b86576001925f91612826575b506128148287612738565b5261281f8186612738565b500161278c565b61284291503d805f833e61283a8183612220565b8101906126c2565b5f612809565b849150612777565b61276e90612768565b5050612863612639565b9190565b6001600160a01b03165f90815260026020526040902060028101546001600160a01b0316919082156128fd575f6128c4610a8d6128b7611abb600360049601546001600160401b039060401c1690565b956001600160a01b031690565b6040516395d89b4160e01b815292839182905afa908115610b86575f916128e9575090565b6105ae91503d805f833e61283a8183612220565b505f915060405161290f602082612220565b5f815290565b60018060a01b03165f526002602052600260405f2001549060ff6001600160401b038360a01c169260e01c1690600382101561250f576002600183149261295b81612505565b1490565b6040519061296e602083612220565b5f808352366020840137565b9061298482612622565b6129916040519182612220565b82815280926129a2601f1991612622565b0190602036910137565b905f549081831015612ad35780612aca57506129c9815b836124f8565b82828211612ac2575b6129db916126b5565b926129e58461297a565b936129ef8161297a565b935f5b8281106129fe57505050565b612a0e6127be6127b983856124f8565b8054909190612a39906001600160a01b0316612a2a838b612738565b6001600160a01b039091169052565b612a4282612c27565b15612ab8576020612a64610a8d610a8d6002600496015460018060a01b031690565b6040516278744560e21b815293849182905afa8015610b86576001925f91612a9a575b505b612a938289612738565b52016129f2565b612ab2915060203d8111610bb357610ba58183612220565b5f612a87565b600191505f612a89565b8291506129d2565b6129c9906129c3565b509050612ade61295f565b9161286361295f565b90602491600381101561250f57600452565b5f54929183821015612b775780612b725750825b81018082116124f357838111612b6b575b8181039081116124f357612b318161297a565b915f5b828110612b42575050509190565b808201908183116124f357612b65612b5b6001936132f9565b612a2a8388612738565b01612b34565b5082612b1e565b612b0d565b505061286361295f565b804710612be9575f808080936110055af13d15612be1573d90612ba3826123f1565b91612bb16040519384612220565b82523d5f602084013e5b15612bc35750565b805115612bd257602081519101fd5b63d6bda27560e01b5f5260045ffd5b606090612bbb565b4763cf47918160e01b5f5260045260245260445ffd5b60ff5f5160206135675f395f51905f52541615612c1857565b638dfc202b60e01b5f5260045ffd5b600281015460ff8160e01c16600381101561250f57159182612c5a575b5081612c4e575090565b60ff915060e81c161590565b612c659192506133e2565b6006541115905f612c44565b6001600160a01b03168015908115612c9a575b50612c8b57565b635926e0c360e01b5f5260045ffd5b612cb091505f52600160205260405f2054151590565b155f612c84565b60038101612ccf611abb82546001600160401b031690565b8311612d8f575b5060028101805460ff60e81b1916600160e81b17905580546001600160a01b0316906110013b156103af57604051634de5f52960e11b81526001600160a01b039290921660048301525f82602481836110015af1908115610b86577feb7d7a49847ec491969db21a0e31b234565a9923145a2d1b56a75c9e95825802926123de92612d7b575b50600101546040519384526001600160a01b0316929081906020820190565b80610b805f612d8993612220565b5f612d5c565b612db690612d9c8461321a565b6001600160401b03166001600160401b0319825416179055565b5f612cd6565b60ff5f5160206135675f395f51905f525416612dd457565b63d93c066560e01b5f5260045ffd5b908015612f2c576001600160a01b0382165f90815260026020526040902060028101805491929091612e1f90610a8d906001600160a01b031681565b604051635d043b2960e11b815260048101839052336024820181905260448201529190602090839060649082905f905af1908115610b8657612eaa925f92612f0b575b506040805192835260208301919091526001600160a01b03861692339284927f3aace7340547de7b9156593a7652dc07ee900cea3fd8f82cb6c9d38b40829802928291820190565b0390a33314612ee8575b50506110113b156103af57604051632961046560e21b81526001600160a01b0390911660048201525f818060248101612367565b5460e81c60ff16611dc457612efc906133e2565b60065411611178575f80612eb4565b612f2591925060203d602011610bb357610ba58183612220565b905f612e62565b633c57b48560e21b5f5260045ffd5b6007543410611178576001600160a01b0381165f908152600260208190526040909120015460e081901c60ff16612f7181612505565b80611b71575060e881901c60ff1680613065575b611dc457612f9d90610a8d906001600160a01b031681565b604051636e553f6560e01b815234600482018190523360248301529091602091839160449183915af1908115610b86575f91613046575b5060408051348152602081019290925233916001600160a01b038416917f24d7bda8602b916d64417f0dbfe2e2e88ec9b1157bd9f596dfdb91ba26624e0491a36110113b156103af57604051632961046560e21b81526001600160a01b0390911660048201525f818060248101612367565b61305f915060203d602011610bb357610ba58183612220565b5f612fd4565b50336001600160a01b0383161415612f85565b90815181101561274c570160200190565b805160038110908115613153575b506130f9575f5b815181101561314f576130d46130ce6130c86130ba8486613078565b516001600160f81b03191690565b60f81c90565b60ff1690565b60308110908115613144575b81613127575b81613108575b506130f95760010161309e565b631bf4348160e31b5f5260045ffd5b606181109150811561311c575b505f6130ec565b607a9150115f613115565b905060418110801561313a575b906130e6565b50605a8111613134565b6039811191506130e0565b5050565b60099150115f613097565b90763d602d80600a3d3981f3363d3d373d3d3d363d730000005f527010035af43d82803e903d91602b57fd5bf3602052603760095ff0916001600160a01b03831691821561320b57823b156103af57604080516379ccf11760e11b81526001600160a01b039093166004840152602483015290915f91839190829081906131e9906044830190610573565b039134905af18015610b86576131fd575090565b80610b805f6105ae93612220565b63b06ebf3d60e01b5f5260045ffd5b6001600160401b038111613234576001600160401b031690565b6306dfcc6560e41b5f52604060045260245260445ffd5b6001600160a01b039081165f8181526002602081905260409091200154909161327691610a8d911681565b604051630f41a04d60e11b81523360048201529190602090839060249082905f905af1918215610b86575f926132d8575b5060405191825233917ff7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd399268390602090a3565b6132f291925060203d602011610bb357610ba58183612220565b905f6132a7565b5f5481101561274c575f80527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56301546001600160a01b031690565b5f7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005d565b5f80549081815b83811061336f57505050505f90565b8181101561274c577f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5638101546001600160a01b03165f9081526002602052604090206133ba90612c27565b6133d8575b600183116133cf57600101613360565b50505050600190565b91600101916133bf565b60028101546001600160a01b03168015613466576134519160209161342590600190613416906001600160a01b0316610a8d565b9201546001600160a01b031690565b60405163ce96cb7760e01b81526001600160a01b03909116600482015292839190829081906024820190565b03915afa908115610b86575f916124b1575090565b50505f90565b60028101805460ff60e81b19169055600101546001600160a01b03167f9390b453426557da5ebdc31f19a37753ca04addf656d32f35232211bb2af3f195f80a2565b805482101561274c575f5260205f2001905f90565b5f818152600160205260409020546124ca575f54600160401b811015612241578060016134f392015f555f6134ae565b81549060031b9083821b915f19901b19161790555f54905f52600160205260405f2055600190565b60ff5f5160206135875f395f51905f525460401c161561353757565b631afcd79f60e31b5f5260045ffdfe46944d7356e409d0d8c174d262d4ab0ddb2d4597e3e424151a463d8a4af4e500cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a264697066735822122018424e41043d7ad1178bb50e371c44fedddc45d8f81974fbaa7fd25e317fc67e64736f6c634300081c0033",
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

// PackUPDATEFREQUENCY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdb6c4c06.
//
// Solidity: function UPDATE_FREQUENCY() view returns(uint256)
func (stakeHub *StakeHub) PackUPDATEFREQUENCY() []byte {
	enc, err := stakeHub.abi.Pack("UPDATE_FREQUENCY")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUPDATEFREQUENCY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdb6c4c06.
//
// Solidity: function UPDATE_FREQUENCY() view returns(uint256)
func (stakeHub *StakeHub) UnpackUPDATEFREQUENCY(data []byte) (*big.Int, error) {
	out, err := stakeHub.abi.Unpack("UPDATE_FREQUENCY", data)
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
// the contract method with ID 0x604ababc.
//
// Solidity: function applyValidator(address validatorAddr, uint256 commissionRate, string id) payable returns()
func (stakeHub *StakeHub) PackApplyValidator(validatorAddr common.Address, commissionRate *big.Int, id string) []byte {
	enc, err := stakeHub.abi.Pack("applyValidator", validatorAddr, commissionRate, id)
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
// Solidity: function getValidatorInfo(string id) view returns(address validatorAddr, address shareContract, bool isActive, bool isJailed, bool isStaked)
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
	ShareContract common.Address
	IsActive      bool
	IsJailed      bool
	IsStaked      bool
}

// UnpackGetValidatorInfo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x41698573.
//
// Solidity: function getValidatorInfo(string id) view returns(address validatorAddr, address shareContract, bool isActive, bool isJailed, bool isStaked)
func (stakeHub *StakeHub) UnpackGetValidatorInfo(data []byte) (GetValidatorInfoOutput, error) {
	out, err := stakeHub.abi.Unpack("getValidatorInfo", data)
	outstruct := new(GetValidatorInfoOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.ValidatorAddr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ShareContract = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.IsActive = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.IsJailed = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.IsStaked = *abi.ConvertType(out[4], new(bool)).(*bool)
	return *outstruct, err

}

// PackGetValidatorInfo0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8a11d7c9.
//
// Solidity: function getValidatorInfo(address operatorAddr) view returns(address validatorAddr, address shareContract, bool isActive, bool isJailed, bool isStaked)
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
	ShareContract common.Address
	IsActive      bool
	IsJailed      bool
	IsStaked      bool
}

// UnpackGetValidatorInfo0 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8a11d7c9.
//
// Solidity: function getValidatorInfo(address operatorAddr) view returns(address validatorAddr, address shareContract, bool isActive, bool isJailed, bool isStaked)
func (stakeHub *StakeHub) UnpackGetValidatorInfo0(data []byte) (GetValidatorInfo0Output, error) {
	out, err := stakeHub.abi.Unpack("getValidatorInfo0", data)
	outstruct := new(GetValidatorInfo0Output)
	if err != nil {
		return *outstruct, err
	}
	outstruct.ValidatorAddr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ShareContract = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.IsActive = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.IsJailed = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.IsStaked = *abi.ConvertType(out[4], new(bool)).(*bool)
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
// Solidity: function getValidators(uint256 offset, uint256 limit) view returns(address[] validatorAddrs, uint256[] amounts, uint256 totalLength)
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
	Amounts        []*big.Int
	TotalLength    *big.Int
}

// UnpackGetValidators is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbff02e20.
//
// Solidity: function getValidators(uint256 offset, uint256 limit) view returns(address[] validatorAddrs, uint256[] amounts, uint256 totalLength)
func (stakeHub *StakeHub) UnpackGetValidators(data []byte) (GetValidatorsOutput, error) {
	out, err := stakeHub.abi.Unpack("getValidators", data)
	outstruct := new(GetValidatorsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.ValidatorAddrs = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Amounts = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.TotalLength = abi.ConvertType(out[2], new(big.Int)).(*big.Int)
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

// PackUpdateValidator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe6635f5d.
//
// Solidity: function updateValidator(address newValidatorAddr, uint256 newCommissionRate) returns()
func (stakeHub *StakeHub) PackUpdateValidator(newValidatorAddr common.Address, newCommissionRate *big.Int) []byte {
	enc, err := stakeHub.abi.Pack("updateValidator", newValidatorAddr, newCommissionRate)
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
// Solidity: event ValidatorApplied(address indexed validatorAddr, address indexed operatorAddr, address indexed shareContract)
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
	OperatorAddr      common.Address
	NewValidatorAddr  common.Address
	NewCommissionRate *big.Int
	Raw               *types.Log // Blockchain specific contextual infos
}

const StakeHubValidatorUpdatedEventName = "ValidatorUpdated"

// ContractEventName returns the user-defined event name.
func (StakeHubValidatorUpdated) ContractEventName() string {
	return StakeHubValidatorUpdatedEventName
}

// UnpackValidatorUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ValidatorUpdated(address indexed operatorAddr, address indexed newValidatorAddr, uint256 newCommissionRate)
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
