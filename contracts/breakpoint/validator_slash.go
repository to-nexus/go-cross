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
	ABI:        "[{\"type\":\"function\",\"name\":\"MITIGATE_RATE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlashInfo\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlashedValidators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mitigate\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"offlineSlashThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setOfflineSlashThreshold\",\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashOffline\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OfflineSlashThresholdUpdated\",\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorMitigated\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"count\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorSlashed\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"count\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValue\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OncePerBlock\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCoinbase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySystemContract\",\"inputs\":[{\"name\":\"contractAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyZeroGasPrice\",\"inputs\":[]}]",
	ID:         "ValidatorSlash",
	BinRuntime: "0x6080806040526004361015610012575f80fd5b5f3560e01c9081632ea22d0b146103d75750806344690e5c1461039e5780634ac9302e146102df5780635b19c262146102c25780637ac6010d146102135780638129fc1c146100d75780638d02d4be146100935763e2b33d5b14610074575f80fd5b3461008f575f36600319011261008f57602060405160048152f35b5f80fd5b3461008f575f36600319011261008f574133036100c8573a6100b9576100b761061d565b005b6383f1b1d360e01b5f5260045ffd5b63022d8c9560e31b5f5260045ffd5b3461008f575f36600319011261008f575f5160206108885f395f51905f525460ff8160401c16159067ffffffffffffffff81168015908161020b575b6001149081610201575b1590816101f8575b506101e95767ffffffffffffffff1981166001175f5160206108885f395f51905f5255816101bd575b504133036100c8573a6100b957603260035561016657005b68ff0000000000000000195f5160206108885f395f51905f5254165f5160206108885f395f51905f52557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a1005b68ffffffffffffffffff191668010000000000000001175f5160206108885f395f51905f52558161014e565b63f92ee8a960e01b5f5260045ffd5b90501583610125565b303b15915061011d565b839150610113565b3461008f57602036600319011261008f5760043561101333036102ad5760048110610269576020817f44cf37fc8d2185243e5927aa52de715bc8ff0d81f4785d5d391bc8178f1a2cdc92600355604051908152a1005b60849060405190632c648cf160e01b82526040600483015260156044830152741bd9999b1a5b9954db185cda151a1c995cda1bdb19605a1b60648301526024820152fd5b630f22c43960e41b5f5261101360045260245ffd5b3461008f575f36600319011261008f576020600354604051908152f35b3461008f575f36600319011261008f576040518060206001549283815201809260015f527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6905f5b818110610388575050508161033d910382610455565b604051918291602083019060208452518091526040830191905f5b818110610366575050500390f35b82516001600160a01b0316845285945060209384019390920191600101610358565b8254845260209093019260019283019201610327565b3461008f57602036600319011261008f576004356001600160a01b038116810361008f574133036100c8573a6100b9576100b790610477565b3461008f57602036600319011261008f576004356001600160a01b0381169081900361008f575f525f60205260405f206040820182811067ffffffffffffffff82111761044157604092602091845260018354938483520154918291015282519182526020820152f35b634e487b7160e01b5f52604160045260245ffd5b90601f8019910116810190811067ffffffffffffffff82111761044157604052565b604051631015428760e21b81526001600160a01b0390911660048201819052906020816024816110015afa908115610595575f916105e2575b50156105df57805f525f60205260405f2080544381146105d0576105a0576001906104da83610752565b5081808201555b438155018054600354111561051f575b60207f17bddadfd7ec8898c3b9eadd0cf5ae77ba8d5df3a50e96ab86ec2dd711aa8fbb9154604051908152a2565b5f81556110023b1561008f5760405163111a439760e21b815260048101839052905f82602481836110025af1908115610595577f17bddadfd7ec8898c3b9eadd0cf5ae77ba8d5df3a50e96ab86ec2dd711aa8fbb92602092610585575b509150506104f1565b5f61058f91610455565b5f61057c565b6040513d5f823e3d90fd5b60018101805491905f1983146105bc57600180930190556104e1565b634e487b7160e01b5f52601160045260245ffd5b638aaaa79d60e01b5f5260045ffd5b50565b90506020813d602011610615575b816105fd60209383610455565b8101031261008f5751801515810361008f575f6104b0565b3d91506105f0565b60015480156105df5760035460021c90805b610637575050565b5f1981018181116105bc576001541115610729577fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf58101546001600160a01b03165f81815260208190526040902060010180548410156106dc578054908482039182116105bc57817f5bd07e8a6722365853b9a8e0af5f9e09a944cbdfc7e86bb21c93cc71486036bd9260209255604051908152a25b80156105bc575f19018061062f565b50805f525f6020525f60016040822082815501556106f9816107c2565b507f5bd07e8a6722365853b9a8e0af5f9e09a944cbdfc7e86bb21c93cc71486036bd60206040515f8152a26106cd565b634e487b7160e01b5f52603260045260245ffd5b8054821015610729575f5260205f2001905f90565b805f52600260205260405f2054155f146107bd5760015468010000000000000000811015610441576107a6610790826001859401600155600161073d565b819391549060031b91821b915f19901b19161790565b9055600154905f52600260205260405f2055600190565b505f90565b5f818152600260205260409020548015610881575f1981018181116105bc576001545f198101919082116105bc57818103610849575b5050506001548015610835575f190161081281600161073d565b8154905f199060031b1b191690556001555f5260026020525f6040812055600190565b634e487b7160e01b5f52603160045260245ffd5b61086b61085a61079093600161073d565b90549060031b1c928392600161073d565b90555f52600260205260405f20555f80806107f8565b50505f9056fef0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a2646970667358221220596e9efddbe589f8a8d99f8e5c49e4d0fc2994395829e7bbafdac5ccf12cf81964736f6c634300081e0033",
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

// PackMITIGATERATE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe2b33d5b.
//
// Solidity: function MITIGATE_RATE() view returns(uint256)
func (validatorSlash *ValidatorSlash) PackMITIGATERATE() []byte {
	enc, err := validatorSlash.abi.Pack("MITIGATE_RATE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMITIGATERATE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe2b33d5b.
//
// Solidity: function MITIGATE_RATE() view returns(uint256)
func (validatorSlash *ValidatorSlash) UnpackMITIGATERATE(data []byte) (*big.Int, error) {
	out, err := validatorSlash.abi.Unpack("MITIGATE_RATE", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
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

// PackMitigate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8d02d4be.
//
// Solidity: function mitigate() returns()
func (validatorSlash *ValidatorSlash) PackMitigate() []byte {
	enc, err := validatorSlash.abi.Pack("mitigate")
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

// ValidatorSlashValidatorMitigated represents a ValidatorMitigated event raised by the ValidatorSlash contract.
type ValidatorSlashValidatorMitigated struct {
	ValidatorAddr common.Address
	Count         *big.Int
	Raw           *types.Log // Blockchain specific contextual infos
}

const ValidatorSlashValidatorMitigatedEventName = "ValidatorMitigated"

// ContractEventName returns the user-defined event name.
func (ValidatorSlashValidatorMitigated) ContractEventName() string {
	return ValidatorSlashValidatorMitigatedEventName
}

// UnpackValidatorMitigatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ValidatorMitigated(address indexed validatorAddr, uint256 count)
func (validatorSlash *ValidatorSlash) UnpackValidatorMitigatedEvent(log *types.Log) (*ValidatorSlashValidatorMitigated, error) {
	event := "ValidatorMitigated"
	if log.Topics[0] != validatorSlash.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ValidatorSlashValidatorMitigated)
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

// ValidatorSlashValidatorSlashed represents a ValidatorSlashed event raised by the ValidatorSlash contract.
type ValidatorSlashValidatorSlashed struct {
	ValidatorAddr common.Address
	Count         *big.Int
	Raw           *types.Log // Blockchain specific contextual infos
}

const ValidatorSlashValidatorSlashedEventName = "ValidatorSlashed"

// ContractEventName returns the user-defined event name.
func (ValidatorSlashValidatorSlashed) ContractEventName() string {
	return ValidatorSlashValidatorSlashedEventName
}

// UnpackValidatorSlashedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ValidatorSlashed(address indexed validatorAddr, uint256 count)
func (validatorSlash *ValidatorSlash) UnpackValidatorSlashedEvent(log *types.Log) (*ValidatorSlashValidatorSlashed, error) {
	event := "ValidatorSlashed"
	if log.Topics[0] != validatorSlash.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ValidatorSlashValidatorSlashed)
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
	if bytes.Equal(raw[:4], validatorSlash.abi.Errors["OncePerBlock"].ID.Bytes()[:4]) {
		return validatorSlash.UnpackOncePerBlockError(raw[4:])
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

// ValidatorSlashOncePerBlock represents a OncePerBlock error raised by the ValidatorSlash contract.
type ValidatorSlashOncePerBlock struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OncePerBlock()
func ValidatorSlashOncePerBlockErrorID() common.Hash {
	return common.HexToHash("0x8aaaa79d8f19a3e4f2e6009e25cdfe18acbcdc0260b835dc213c1a242e2c3af6")
}

// UnpackOncePerBlockError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OncePerBlock()
func (validatorSlash *ValidatorSlash) UnpackOncePerBlockError(raw []byte) (*ValidatorSlashOncePerBlock, error) {
	out := new(ValidatorSlashOncePerBlock)
	if err := validatorSlash.abi.UnpackIntoInterface(out, "OncePerBlock", raw); err != nil {
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
