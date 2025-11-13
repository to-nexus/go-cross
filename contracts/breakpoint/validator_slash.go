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

// ValidatorSlashMetaData contains all meta data concerning the ValidatorSlash contract.
var ValidatorSlashMetaData = bind.MetaData{
	ABI:        "[{\"type\":\"function\",\"name\":\"getSlashInfo\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlashedValidators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"offlineSlashThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setOfflineSlashThreshold\",\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashOffline\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OfflineSlashThresholdUpdated\",\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OfflineSlashed\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValue\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCoinbase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySystemContract\",\"inputs\":[{\"name\":\"contractAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyZeroGasPrice\",\"inputs\":[]}]",
	ID:         "ValidatorSlash",
	BinRuntime: "0x6080806040526004361015610012575f80fd5b5f3560e01c9081632ea22d0b146103805750806344690e5c146103455780634ac9302e146102865780635b19c262146102695780637ac6010d146101bc57638129fc1c1461005e575f80fd5b346101b8575f3660031901126101b8575f5160206106755f395f51905f525460ff8160401c16159067ffffffffffffffff8116801590816101b0575b60011490816101a6575b15908161019d575b5061018e5767ffffffffffffffff1981166001175f5160206106755f395f51905f525581610162575b50413303610153573a6101445760326003556100ed57005b68ff0000000000000000195f5160206106755f395f51905f5254165f5160206106755f395f51905f52557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a1005b6383f1b1d360e01b5f5260045ffd5b63022d8c9560e31b5f5260045ffd5b68ffffffffffffffffff191668010000000000000001175f5160206106755f395f51905f52555f6100d5565b63f92ee8a960e01b5f5260045ffd5b9050155f6100ac565b303b1591506100a4565b83915061009a565b5f80fd5b346101b85760203660031901126101b8576004356110083303610254578015610210576020817f44cf37fc8d2185243e5927aa52de715bc8ff0d81f4785d5d391bc8178f1a2cdc92600355604051908152a1005b60849060405190632c648cf160e01b82526040600483015260156044830152741bd9999b1a5b9954db185cda151a1c995cda1bdb19605a1b60648301526024820152fd5b630f22c43960e41b5f5261100860045260245ffd5b346101b8575f3660031901126101b8576020600354604051908152f35b346101b8575f3660031901126101b8576040518060206001549283815201809260015f527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6905f5b81811061032f57505050816102e491038261040a565b604051918291602083019060208452518091526040830191905f5b81811061030d575050500390f35b82516001600160a01b03168452859450602093840193909201916001016102ff565b82548452602090930192600192830192016102ce565b346101b85760203660031901126101b8576004356001600160a01b03811681036101b857413303610153573a6101445761037e9061042c565b005b346101b85760203660031901126101b8576004356001600160a01b03811691908290036101b8576040915f525f6020526020825f20916103bf816103da565b60018354938483520154918291015282519182526020820152f35b6040810190811067ffffffffffffffff8211176103f657604052565b634e487b7160e01b5f52604160045260245ffd5b90601f8019910116810190811067ffffffffffffffff8211176103f657604052565b60405163facd743b60e01b81526001600160a01b0390911660048201819052905f906020816024816110015afa908115610578575f916105c8575b50156105c457815f525f60205260405f2090604051610485816103da565b60018354938483520154602082019381855243146105bd578061059757506104ac84610603565b50600183525b4381528251600354908115610583570615610503575b60017fa134ecb0a8ea69d0a9d1f43e7c16431fe63e3671536b461b96c16a5f2e7860d5938584528360205260408420925183555191015580a2565b6110023b156101b85760405163111a439760e21b815260048101859052925f84602481836110025af1938415610578577fa134ecb0a8ea69d0a9d1f43e7c16431fe63e3671536b461b96c16a5f2e7860d594610562575b5092506104c8565b61056f9193505f9061040a565b5f91600161055a565b6040513d5f823e3d90fd5b634e487b7160e01b5f52601260045260245ffd5b5f1981146105a95760010183526104b2565b634e487b7160e01b5f52601160045260245ffd5b5050505050565b5050565b90506020813d6020116105fb575b816105e36020938361040a565b810103126101b8575180151581036101b8575f610467565b3d91506105d6565b805f52600260205260405f2054155f1461066f57600154680100000000000000008110156103f657600181018060015581101561065b57819060015f5260205f200155600154905f52600260205260405f2055600190565b634e487b7160e01b5f52603260045260245ffd5b505f9056fef0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a264697066735822122058471f0206e5d6b3ed0f8f36a5aae35bf0f1b02dfa6418fb0a4e079a4abc573764736f6c634300081c0033",
}

// ValidatorSlash is an auto generated Go binding around an Ethereum contract.
type ValidatorSlash struct {
	abi abi.ABI
}

// NewValidatorSlash creates a new instance of ValidatorSlash.
func NewValidatorSlash() *ValidatorSlash {
	parsed, err := ValidatorSlashMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ValidatorSlash{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ValidatorSlash) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackGetSlashInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2ea22d0b.
//
// Solidity: function getSlashInfo(address validatorAddr) view returns(uint256 height, uint256 count)
func (validatorSlash *ValidatorSlash) PackGetSlashInfo(validatorAddr common.Address) []byte {
	enc, err := validatorSlash.abi.Pack("getSlashInfo", validatorAddr)
	if err != nil {
		panic(err)
	}
	return enc
}

// GetSlashInfoOutput serves as a container for the return parameters of contract
// method GetSlashInfo.
type GetSlashInfoOutput struct {
	Height *big.Int
	Count  *big.Int
}

// UnpackGetSlashInfo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2ea22d0b.
//
// Solidity: function getSlashInfo(address validatorAddr) view returns(uint256 height, uint256 count)
func (validatorSlash *ValidatorSlash) UnpackGetSlashInfo(data []byte) (GetSlashInfoOutput, error) {
	out, err := validatorSlash.abi.Unpack("getSlashInfo", data)
	outstruct := new(GetSlashInfoOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Height = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.Count = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, err

}

// PackGetSlashedValidators is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4ac9302e.
//
// Solidity: function getSlashedValidators() view returns(address[])
func (validatorSlash *ValidatorSlash) PackGetSlashedValidators() []byte {
	enc, err := validatorSlash.abi.Pack("getSlashedValidators")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetSlashedValidators is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4ac9302e.
//
// Solidity: function getSlashedValidators() view returns(address[])
func (validatorSlash *ValidatorSlash) UnpackGetSlashedValidators(data []byte) ([]common.Address, error) {
	out, err := validatorSlash.abi.Unpack("getSlashedValidators", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, err
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (validatorSlash *ValidatorSlash) PackInitialize() []byte {
	enc, err := validatorSlash.abi.Pack("initialize")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackOfflineSlashThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5b19c262.
//
// Solidity: function offlineSlashThreshold() view returns(uint256)
func (validatorSlash *ValidatorSlash) PackOfflineSlashThreshold() []byte {
	enc, err := validatorSlash.abi.Pack("offlineSlashThreshold")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOfflineSlashThreshold is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5b19c262.
//
// Solidity: function offlineSlashThreshold() view returns(uint256)
func (validatorSlash *ValidatorSlash) UnpackOfflineSlashThreshold(data []byte) (*big.Int, error) {
	out, err := validatorSlash.abi.Unpack("offlineSlashThreshold", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackSetOfflineSlashThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7ac6010d.
//
// Solidity: function setOfflineSlashThreshold(uint256 threshold) returns()
func (validatorSlash *ValidatorSlash) PackSetOfflineSlashThreshold(threshold *big.Int) []byte {
	enc, err := validatorSlash.abi.Pack("setOfflineSlashThreshold", threshold)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSlashOffline is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x44690e5c.
//
// Solidity: function slashOffline(address validatorAddr) returns()
func (validatorSlash *ValidatorSlash) PackSlashOffline(validatorAddr common.Address) []byte {
	enc, err := validatorSlash.abi.Pack("slashOffline", validatorAddr)
	if err != nil {
		panic(err)
	}
	return enc
}

// ValidatorSlashInitialized represents a Initialized event raised by the ValidatorSlash contract.
type ValidatorSlashInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const ValidatorSlashInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (ValidatorSlashInitialized) ContractEventName() string {
	return ValidatorSlashInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (validatorSlash *ValidatorSlash) UnpackInitializedEvent(log *types.Log) (*ValidatorSlashInitialized, error) {
	event := "Initialized"
	if log.Topics[0] != validatorSlash.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ValidatorSlashInitialized)
	if len(log.Data) > 0 {
		if err := validatorSlash.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range validatorSlash.abi.Events[event].Inputs {
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

// ValidatorSlashOfflineSlashThresholdUpdated represents a OfflineSlashThresholdUpdated event raised by the ValidatorSlash contract.
type ValidatorSlashOfflineSlashThresholdUpdated struct {
	Threshold *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const ValidatorSlashOfflineSlashThresholdUpdatedEventName = "OfflineSlashThresholdUpdated"

// ContractEventName returns the user-defined event name.
func (ValidatorSlashOfflineSlashThresholdUpdated) ContractEventName() string {
	return ValidatorSlashOfflineSlashThresholdUpdatedEventName
}

// UnpackOfflineSlashThresholdUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OfflineSlashThresholdUpdated(uint256 threshold)
func (validatorSlash *ValidatorSlash) UnpackOfflineSlashThresholdUpdatedEvent(log *types.Log) (*ValidatorSlashOfflineSlashThresholdUpdated, error) {
	event := "OfflineSlashThresholdUpdated"
	if log.Topics[0] != validatorSlash.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ValidatorSlashOfflineSlashThresholdUpdated)
	if len(log.Data) > 0 {
		if err := validatorSlash.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range validatorSlash.abi.Events[event].Inputs {
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

// ValidatorSlashOfflineSlashed represents a OfflineSlashed event raised by the ValidatorSlash contract.
type ValidatorSlashOfflineSlashed struct {
	ValidatorAddr common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const ValidatorSlashOfflineSlashedEventName = "OfflineSlashed"

// ContractEventName returns the user-defined event name.
func (ValidatorSlashOfflineSlashed) ContractEventName() string {
	return ValidatorSlashOfflineSlashedEventName
}

// UnpackOfflineSlashedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OfflineSlashed(address indexed validatorAddr)
func (validatorSlash *ValidatorSlash) UnpackOfflineSlashedEvent(log *types.Log) (*ValidatorSlashOfflineSlashed, error) {
	event := "OfflineSlashed"
	if log.Topics[0] != validatorSlash.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ValidatorSlashOfflineSlashed)
	if len(log.Data) > 0 {
		if err := validatorSlash.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range validatorSlash.abi.Events[event].Inputs {
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
func (validatorSlash *ValidatorSlash) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], validatorSlash.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return validatorSlash.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], validatorSlash.abi.Errors["InvalidValue"].ID.Bytes()[:4]) {
		return validatorSlash.UnpackInvalidValueError(raw[4:])
	}
	if bytes.Equal(raw[:4], validatorSlash.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return validatorSlash.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], validatorSlash.abi.Errors["OnlyCoinbase"].ID.Bytes()[:4]) {
		return validatorSlash.UnpackOnlyCoinbaseError(raw[4:])
	}
	if bytes.Equal(raw[:4], validatorSlash.abi.Errors["OnlySystemContract"].ID.Bytes()[:4]) {
		return validatorSlash.UnpackOnlySystemContractError(raw[4:])
	}
	if bytes.Equal(raw[:4], validatorSlash.abi.Errors["OnlyZeroGasPrice"].ID.Bytes()[:4]) {
		return validatorSlash.UnpackOnlyZeroGasPriceError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ValidatorSlashInvalidInitialization represents a InvalidInitialization error raised by the ValidatorSlash contract.
type ValidatorSlashInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func ValidatorSlashInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (validatorSlash *ValidatorSlash) UnpackInvalidInitializationError(raw []byte) (*ValidatorSlashInvalidInitialization, error) {
	out := new(ValidatorSlashInvalidInitialization)
	if err := validatorSlash.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorSlashInvalidValue represents a InvalidValue error raised by the ValidatorSlash contract.
type ValidatorSlashInvalidValue struct {
	Key   string
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidValue(string key, uint256 value)
func ValidatorSlashInvalidValueErrorID() common.Hash {
	return common.HexToHash("0x2c648cf1bbb773fa5fbcfc6541df5c1891af9b6969d9a555653bcec289d77218")
}

// UnpackInvalidValueError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidValue(string key, uint256 value)
func (validatorSlash *ValidatorSlash) UnpackInvalidValueError(raw []byte) (*ValidatorSlashInvalidValue, error) {
	out := new(ValidatorSlashInvalidValue)
	if err := validatorSlash.abi.UnpackIntoInterface(out, "InvalidValue", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorSlashNotInitializing represents a NotInitializing error raised by the ValidatorSlash contract.
type ValidatorSlashNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func ValidatorSlashNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (validatorSlash *ValidatorSlash) UnpackNotInitializingError(raw []byte) (*ValidatorSlashNotInitializing, error) {
	out := new(ValidatorSlashNotInitializing)
	if err := validatorSlash.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorSlashOnlyCoinbase represents a OnlyCoinbase error raised by the ValidatorSlash contract.
type ValidatorSlashOnlyCoinbase struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyCoinbase()
func ValidatorSlashOnlyCoinbaseErrorID() common.Hash {
	return common.HexToHash("0x116c64a8de02ce00303a109e06f26c31a3cfed64656fb9d228157fad57dff616")
}

// UnpackOnlyCoinbaseError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyCoinbase()
func (validatorSlash *ValidatorSlash) UnpackOnlyCoinbaseError(raw []byte) (*ValidatorSlashOnlyCoinbase, error) {
	out := new(ValidatorSlashOnlyCoinbase)
	if err := validatorSlash.abi.UnpackIntoInterface(out, "OnlyCoinbase", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorSlashOnlySystemContract represents a OnlySystemContract error raised by the ValidatorSlash contract.
type ValidatorSlashOnlySystemContract struct {
	ContractAddr common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlySystemContract(address contractAddr)
func ValidatorSlashOnlySystemContractErrorID() common.Hash {
	return common.HexToHash("0xf22c4390ebf387afc834c245f4ef6f38d59454737d03e451e0d182ac61608277")
}

// UnpackOnlySystemContractError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlySystemContract(address contractAddr)
func (validatorSlash *ValidatorSlash) UnpackOnlySystemContractError(raw []byte) (*ValidatorSlashOnlySystemContract, error) {
	out := new(ValidatorSlashOnlySystemContract)
	if err := validatorSlash.abi.UnpackIntoInterface(out, "OnlySystemContract", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorSlashOnlyZeroGasPrice represents a OnlyZeroGasPrice error raised by the ValidatorSlash contract.
type ValidatorSlashOnlyZeroGasPrice struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyZeroGasPrice()
func ValidatorSlashOnlyZeroGasPriceErrorID() common.Hash {
	return common.HexToHash("0x83f1b1d3f8cc3470adc53b59d7024697cf0374e9ddadd2111159d00fe76f2c06")
}

// UnpackOnlyZeroGasPriceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyZeroGasPrice()
func (validatorSlash *ValidatorSlash) UnpackOnlyZeroGasPriceError(raw []byte) (*ValidatorSlashOnlyZeroGasPrice, error) {
	out := new(ValidatorSlashOnlyZeroGasPrice)
	if err := validatorSlash.abi.UnpackIntoInterface(out, "OnlyZeroGasPrice", raw); err != nil {
		return nil, err
	}
	return out, nil
}
