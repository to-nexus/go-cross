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

// ValidatorSetMetaData contains all meta data concerning the ValidatorSet contract.
var ValidatorSetMetaData = bind.MetaData{
	ABI:        "[{\"type\":\"function\",\"name\":\"getValidatorCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isValidator\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateValidators\",\"inputs\":[{\"name\":\"validatorAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ValidatorsUpdated\",\"inputs\":[{\"name\":\"updater\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"validatorsCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidValue\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NoValidators\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCoinbase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySystemContract\",\"inputs\":[{\"name\":\"contractAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyZeroGasPrice\",\"inputs\":[]}]",
	ID:         "ValidatorSet",
	BinRuntime: "0x6080806040526004361015610012575f80fd5b5f3560e01c9081637071688a1461043d57508063b7ab4db514610383578063e71731e41461008a5763facd743b14610048575f80fd5b34610086576020366003190112610086576004356001600160a01b03811690819003610086575f526001602052602060405f20541515604051908152f35b5f80fd5b346100865760203660031901126100865760043567ffffffffffffffff8111610086573660238201121561008657806004013567ffffffffffffffff8111610086573660248260051b8401011161008657413303610374573a610365578015610356576100f6816104e1565b915f5b82811015610134576024600582901b830101356001600160a01b03811691908290036100865760019161012c82876104cd565b5152016100f9565b82845f545f5b81811061032e5750505f545f8055806102e2575b50805161015a816104e1565b8051906801000000000000000082116102ce575f54825f55808310610273575b506020015f80525f5160206105b25f395f51905f525f915b83831061025557505050505f5b8181106101d557836040519081527f59789ca63ec0c402ea2e2e4224898146dfec220e577c1501583d32f90f54bb8360203392a2005b6101df81846104cd565b516101e982610494565b919091610242576101f991610569565b600181019081811161022e576001916001600160a01b0361021a83876104cd565b5151165f528260205260405f20550161019f565b634e487b7160e01b5f52601160045260245ffd5b634e487b7160e01b5f525f60045260245ffd5b60146020826102676001945186610569565b01920192019190610192565b8060140290601482040361022e578260140260148104840361022e575f80525f5160206105b25f395f51905f5291820191015b8181106102b3575061017a565b805f601492556102c882820160018301610553565b016102a6565b634e487b7160e01b5f52604160045260245ffd5b8060140290601482040361022e575f80525f5160206105b25f395f51905f52908101905b818110610313575061014e565b805f6014925561032882820160018301610553565b01610306565b8061033a600192610494565b50828060a01b039054165f52816020525f60408120550161013a565b6313f8a3a760e31b5f5260045ffd5b6383f1b1d360e01b5f5260045ffd5b63022d8c9560e31b5f5260045ffd5b34610086575f366003190112610086575f546103a66103a18261047c565b610456565b908082526103b38161047c565b602083019190601f19013683375f5b818110610416578284604051918291602083019060208452518091526040830191905f5b8181106103f4575050500390f35b82516001600160a01b03168452859450602093840193909201916001016103e6565b80610422600192610494565b50828060a01b0390541661043682876104cd565b52016103c2565b34610086575f366003190112610086576020905f548152f35b6040519190601f01601f1916820167ffffffffffffffff8111838210176102ce57604052565b67ffffffffffffffff81116102ce5760051b60200190565b5f548110156104b9575f8080526014919091025f5160206105b25f395f51905f520191565b634e487b7160e01b5f52603260045260245ffd5b80518210156104b95760209160051b010190565b906104ee6103a18361047c565b82815280926104ff601f199161047c565b015f5b81811061050e57505050565b60405190604082019180831067ffffffffffffffff8411176102ce576020926040525f815261026061053f81610456565b903682378382015282828601015201610502565b81811061055e575050565b5f8155600101610553565b815181546001600160a01b0319166001600160a01b03919091161781556020909101515f5b6013811061059b57505050565b600190602083519301928282860101550161058e56fe290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563a2646970667358221220d752196c8321f8d6caf1b32d52b44dfda920de18f039a62583360c3790c5b1c764736f6c634300081c0033",
}

// ValidatorSet is an auto generated Go binding around an Ethereum contract.
type ValidatorSet struct {
	abi abi.ABI
}

// NewValidatorSet creates a new instance of ValidatorSet.
func NewValidatorSet() *ValidatorSet {
	parsed, err := ValidatorSetMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ValidatorSet{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ValidatorSet) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackGetValidatorCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7071688a.
//
// Solidity: function getValidatorCount() view returns(uint256)
func (validatorSet *ValidatorSet) PackGetValidatorCount() []byte {
	enc, err := validatorSet.abi.Pack("getValidatorCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetValidatorCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7071688a.
//
// Solidity: function getValidatorCount() view returns(uint256)
func (validatorSet *ValidatorSet) UnpackGetValidatorCount(data []byte) (*big.Int, error) {
	out, err := validatorSet.abi.Unpack("getValidatorCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetValidators is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (validatorSet *ValidatorSet) PackGetValidators() []byte {
	enc, err := validatorSet.abi.Pack("getValidators")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetValidators is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (validatorSet *ValidatorSet) UnpackGetValidators(data []byte) ([]common.Address, error) {
	out, err := validatorSet.abi.Unpack("getValidators", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, err
}

// PackIsValidator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfacd743b.
//
// Solidity: function isValidator(address validator) view returns(bool)
func (validatorSet *ValidatorSet) PackIsValidator(validator common.Address) []byte {
	enc, err := validatorSet.abi.Pack("isValidator", validator)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsValidator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfacd743b.
//
// Solidity: function isValidator(address validator) view returns(bool)
func (validatorSet *ValidatorSet) UnpackIsValidator(data []byte) (bool, error) {
	out, err := validatorSet.abi.Unpack("isValidator", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackUpdateValidators is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe71731e4.
//
// Solidity: function updateValidators(address[] validatorAddrs) returns()
func (validatorSet *ValidatorSet) PackUpdateValidators(validatorAddrs []common.Address) []byte {
	enc, err := validatorSet.abi.Pack("updateValidators", validatorAddrs)
	if err != nil {
		panic(err)
	}
	return enc
}

// ValidatorSetValidatorsUpdated represents a ValidatorsUpdated event raised by the ValidatorSet contract.
type ValidatorSetValidatorsUpdated struct {
	Updater         common.Address
	ValidatorsCount *big.Int
	Raw             *types.Log // Blockchain specific contextual infos
}

const ValidatorSetValidatorsUpdatedEventName = "ValidatorsUpdated"

// ContractEventName returns the user-defined event name.
func (ValidatorSetValidatorsUpdated) ContractEventName() string {
	return ValidatorSetValidatorsUpdatedEventName
}

// UnpackValidatorsUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ValidatorsUpdated(address indexed updater, uint256 validatorsCount)
func (validatorSet *ValidatorSet) UnpackValidatorsUpdatedEvent(log *types.Log) (*ValidatorSetValidatorsUpdated, error) {
	event := "ValidatorsUpdated"
	if log.Topics[0] != validatorSet.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ValidatorSetValidatorsUpdated)
	if len(log.Data) > 0 {
		if err := validatorSet.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range validatorSet.abi.Events[event].Inputs {
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
func (validatorSet *ValidatorSet) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], validatorSet.abi.Errors["InvalidValue"].ID.Bytes()[:4]) {
		return validatorSet.UnpackInvalidValueError(raw[4:])
	}
	if bytes.Equal(raw[:4], validatorSet.abi.Errors["NoValidators"].ID.Bytes()[:4]) {
		return validatorSet.UnpackNoValidatorsError(raw[4:])
	}
	if bytes.Equal(raw[:4], validatorSet.abi.Errors["OnlyCoinbase"].ID.Bytes()[:4]) {
		return validatorSet.UnpackOnlyCoinbaseError(raw[4:])
	}
	if bytes.Equal(raw[:4], validatorSet.abi.Errors["OnlySystemContract"].ID.Bytes()[:4]) {
		return validatorSet.UnpackOnlySystemContractError(raw[4:])
	}
	if bytes.Equal(raw[:4], validatorSet.abi.Errors["OnlyZeroGasPrice"].ID.Bytes()[:4]) {
		return validatorSet.UnpackOnlyZeroGasPriceError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ValidatorSetInvalidValue represents a InvalidValue error raised by the ValidatorSet contract.
type ValidatorSetInvalidValue struct {
	Key   string
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidValue(string key, uint256 value)
func ValidatorSetInvalidValueErrorID() common.Hash {
	return common.HexToHash("0x2c648cf1bbb773fa5fbcfc6541df5c1891af9b6969d9a555653bcec289d77218")
}

// UnpackInvalidValueError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidValue(string key, uint256 value)
func (validatorSet *ValidatorSet) UnpackInvalidValueError(raw []byte) (*ValidatorSetInvalidValue, error) {
	out := new(ValidatorSetInvalidValue)
	if err := validatorSet.abi.UnpackIntoInterface(out, "InvalidValue", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorSetNoValidators represents a NoValidators error raised by the ValidatorSet contract.
type ValidatorSetNoValidators struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoValidators()
func ValidatorSetNoValidatorsErrorID() common.Hash {
	return common.HexToHash("0x9fc51d38400b51f3e674fdb3d5301f19c8eacfc8dbff7f6c33697c643c6cbfa9")
}

// UnpackNoValidatorsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoValidators()
func (validatorSet *ValidatorSet) UnpackNoValidatorsError(raw []byte) (*ValidatorSetNoValidators, error) {
	out := new(ValidatorSetNoValidators)
	if err := validatorSet.abi.UnpackIntoInterface(out, "NoValidators", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorSetOnlyCoinbase represents a OnlyCoinbase error raised by the ValidatorSet contract.
type ValidatorSetOnlyCoinbase struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyCoinbase()
func ValidatorSetOnlyCoinbaseErrorID() common.Hash {
	return common.HexToHash("0x116c64a8de02ce00303a109e06f26c31a3cfed64656fb9d228157fad57dff616")
}

// UnpackOnlyCoinbaseError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyCoinbase()
func (validatorSet *ValidatorSet) UnpackOnlyCoinbaseError(raw []byte) (*ValidatorSetOnlyCoinbase, error) {
	out := new(ValidatorSetOnlyCoinbase)
	if err := validatorSet.abi.UnpackIntoInterface(out, "OnlyCoinbase", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorSetOnlySystemContract represents a OnlySystemContract error raised by the ValidatorSet contract.
type ValidatorSetOnlySystemContract struct {
	ContractAddr common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlySystemContract(address contractAddr)
func ValidatorSetOnlySystemContractErrorID() common.Hash {
	return common.HexToHash("0xf22c4390ebf387afc834c245f4ef6f38d59454737d03e451e0d182ac61608277")
}

// UnpackOnlySystemContractError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlySystemContract(address contractAddr)
func (validatorSet *ValidatorSet) UnpackOnlySystemContractError(raw []byte) (*ValidatorSetOnlySystemContract, error) {
	out := new(ValidatorSetOnlySystemContract)
	if err := validatorSet.abi.UnpackIntoInterface(out, "OnlySystemContract", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorSetOnlyZeroGasPrice represents a OnlyZeroGasPrice error raised by the ValidatorSet contract.
type ValidatorSetOnlyZeroGasPrice struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyZeroGasPrice()
func ValidatorSetOnlyZeroGasPriceErrorID() common.Hash {
	return common.HexToHash("0x83f1b1d3f8cc3470adc53b59d7024697cf0374e9ddadd2111159d00fe76f2c06")
}

// UnpackOnlyZeroGasPriceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyZeroGasPrice()
func (validatorSet *ValidatorSet) UnpackOnlyZeroGasPriceError(raw []byte) (*ValidatorSetOnlyZeroGasPrice, error) {
	out := new(ValidatorSetOnlyZeroGasPrice)
	if err := validatorSet.abi.UnpackIntoInterface(out, "OnlyZeroGasPrice", raw); err != nil {
		return nil, err
	}
	return out, nil
}
