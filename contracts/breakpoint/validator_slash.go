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
	ABI:        "[{\"type\":\"function\",\"name\":\"MITIGATE_RATE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlashInfo\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlashedValidators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mitigate\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"offlineSlashThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setOfflineSlashThreshold\",\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashOffline\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OfflineSlashThresholdUpdated\",\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorMitigated\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"count\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorSlashed\",\"inputs\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"count\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValue\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotValidator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OncePerBlock\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCoinbase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySystemContract\",\"inputs\":[{\"name\":\"contractAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyZeroGasPrice\",\"inputs\":[]}]",
	ID:         "ValidatorSlash",
	BinRuntime: "0x6080806040526004361015610012575f80fd5b5f905f3560e01c9081632ea22d0b1461059d5750806344690e5c146103c55780634ac9302e146103065780635b19c262146102e85780637ac6010d146102375780638129fc1c146100da5780638d02d4be146100945763e2b33d5b14610076575f80fd5b34610091578060031936011261009157602060405160048152f35b80fd5b50346100915780600319360112610091574133036100cb573a6100bc576100b961063d565b80f35b6383f1b1d360e01b8152600490fd5b63022d8c9560e31b8152600490fd5b50346100915780600319360112610091575f5160206108ab5f395f51905f525460ff8160401c16159067ffffffffffffffff81168015908161022f575b6001149081610225575b15908161021c575b5061020d5767ffffffffffffffff1981166001175f5160206108ab5f395f51905f5255816101e1575b504133036101d2573a6101c357603260035561016b5780f35b68ff0000000000000000195f5160206108ab5f395f51905f5254165f5160206108ab5f395f51905f52557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a180f35b6383f1b1d360e01b8252600482fd5b63022d8c9560e31b8252600482fd5b68ffffffffffffffffff191668010000000000000001175f5160206108ab5f395f51905f52555f610152565b63f92ee8a960e01b8352600483fd5b9050155f610129565b303b159150610121565b839150610117565b50346100915760203660031901126100915760043561101333036102d3576004811061028f576020817f44cf37fc8d2185243e5927aa52de715bc8ff0d81f4785d5d391bc8178f1a2cdc92600355604051908152a180f35b60849060405190632c648cf160e01b82526040600483015260156044830152741bd9999b1a5b9954db185cda151a1c995cda1bdb19605a1b60648301526024820152fd5b630f22c43960e41b8252611013600452602482fd5b50346100915780600319360112610091576020600354604051908152f35b503461009157806003193601126100915760405180602060015491828152018091600185527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf690855b8181106103af575050508261036591038361061b565b604051928392602084019060208552518091526040840192915b81811061038d575050500390f35b82516001600160a01b031684528594506020938401939092019160010161037f565b825484526020909301926001928301920161034f565b50346104f25760203660031901126104f2576004356001600160a01b038116908190036104f25741330361058e573a61057f5760405163facd743b60e01b8152600481018290526020816024816110015afa9081156104e7575f91610544575b501561053557805f525f60205260405f208054438114610526576104f65760019061044f83610775565b5081808201555b4381550180546003541115610495575b60207f17bddadfd7ec8898c3b9eadd0cf5ae77ba8d5df3a50e96ab86ec2dd711aa8fbb9154604051908152a280f35b5f81556110023b156104f25760405163111a439760e21b8152600481018390525f81602481836110025af180156104e7576104d1575b50610466565b6104de9193505f9061061b565b5f9160206104cb565b6040513d5f823e3d90fd5b5f80fd5b60018101805491905f1983146105125760018093019055610456565b634e487b7160e01b5f52601160045260245ffd5b638aaaa79d60e01b5f5260045ffd5b632ec5b44960e01b5f5260045ffd5b90506020813d602011610577575b8161055f6020938361061b565b810103126104f2575180151581036104f2575f610425565b3d9150610552565b6383f1b1d360e01b5f5260045ffd5b63022d8c9560e31b5f5260045ffd5b346104f25760203660031901126104f2576004356001600160a01b038116908190036104f2575f525f60205260405f206040820182811067ffffffffffffffff82111761060757604092602091845260018354938483520154918291015282519182526020820152f35b634e487b7160e01b5f52604160045260245ffd5b90601f8019910116810190811067ffffffffffffffff82111761060757604052565b600154801561075d5760035460021c90805b610657575050565b5f198101818111610512576001541115610749577fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf58101546001600160a01b03165f81815260208190526040902060010180548410156106fc5780549084820391821161051257817f5bd07e8a6722365853b9a8e0af5f9e09a944cbdfc7e86bb21c93cc71486036bd9260209255604051908152a25b8015610512575f19018061064f565b50805f525f6020525f6001604082208281550155610719816107e5565b507f5bd07e8a6722365853b9a8e0af5f9e09a944cbdfc7e86bb21c93cc71486036bd60206040515f8152a26106ed565b634e487b7160e01b5f52603260045260245ffd5b50565b8054821015610749575f5260205f2001905f90565b805f52600260205260405f2054155f146107e05760015468010000000000000000811015610607576107c96107b38260018594016001556001610760565b819391549060031b91821b915f19901b19161790565b9055600154905f52600260205260405f2055600190565b505f90565b5f8181526002602052604090205480156108a4575f198101818111610512576001545f198101919082116105125781810361086c575b5050506001548015610858575f1901610835816001610760565b8154905f199060031b1b191690556001555f5260026020525f6040812055600190565b634e487b7160e01b5f52603160045260245ffd5b61088e61087d6107b3936001610760565b90549060031b1c9283926001610760565b90555f52600260205260405f20555f808061081b565b50505f9056fef0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a264697066735822122009c8838bb34177ee43ac015120d1061326927af3a080ec55d5380e637a55e70f64736f6c634300081c0033",
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
	if bytes.Equal(raw[:4], validatorSlash.abi.Errors["NotValidator"].ID.Bytes()[:4]) {
		return validatorSlash.UnpackNotValidatorError(raw[4:])
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

// ValidatorSlashNotValidator represents a NotValidator error raised by the ValidatorSlash contract.
type ValidatorSlashNotValidator struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotValidator()
func ValidatorSlashNotValidatorErrorID() common.Hash {
	return common.HexToHash("0x2ec5b449e1d63fa34c160c67ce2d5826d939017f27bc0c78ef142b8357c69f9e")
}

// UnpackNotValidatorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotValidator()
func (validatorSlash *ValidatorSlash) UnpackNotValidatorError(raw []byte) (*ValidatorSlashNotValidator, error) {
	out := new(ValidatorSlashNotValidator)
	if err := validatorSlash.abi.UnpackIntoInterface(out, "NotValidator", raw); err != nil {
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
