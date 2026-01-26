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
	ABI:        "[{\"type\":\"function\",\"name\":\"getActiveValidators\",\"inputs\":[],\"outputs\":[{\"name\":\"validatorAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"signerAddrs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isActiveValidator\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isValidator\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"jail\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unstake\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateValidators\",\"inputs\":[{\"name\":\"validatorAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"signerAddrs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ValidatorsUpdated\",\"inputs\":[{\"name\":\"updater\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"validatorAddrs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"signerAddrs\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidValue\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"LengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoValidators\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCoinbase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySystemContract\",\"inputs\":[{\"name\":\"contractAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyZeroGasPrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintToInt\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	ID:         "ValidatorSet",
	BinRuntime: "0x60806040526004361015610011575f80fd5b5f3560e01c806340550a1c1461089b5780637071688a1461087e5780638ee214d31461035a5780639bcbea52146103005780639de7025814610119578063f2888dbb146100ac5763facd743b14610066575f80fd5b346100a85760203660031901126100a857602061009e6001600160a01b0361008c6108ff565b165f52600260205260405f2054151590565b6040519015158152f35b5f80fd5b346100a85760203660031901126100a8576100c56108ff565b6110023303610104576001600160a01b03165f818152600260205260409020546100eb57005b5f525f60205260405f20600160ff198254161790555f80f35b630f22c43960e41b5f5261100260045260245ffd5b346100a8575f3660031901126100a8576001545f61013682610937565b906101446040519283610915565b828252601f1961015384610937565b0136602084013761016383610937565b926101716040519485610915565b808452601f1961018082610937565b015f5b8181106102ef5750505f5b8181106101af57505080825282526101ab6040519283928361094f565b0390f35b6101b881610c58565b60018060a01b0391549060031b1c165f525f6020526101d960405f20610a7e565b6101e6575b60010161018e565b916101f083610c58565b905460039190911b1c6001600160a01b031661020c8286610a1e565b5261021683610c58565b60018060a01b0391549060031b1c165f525f602052600160405f2001604051905f9080549061024482610a46565b80855291600181169081156102c8575060011461028f575b505091816102706001949385940382610915565b61027a8289610a1e565b526102858188610a1e565b50019290506101de565b5f908152602081209092505b8183106102b257505081016020018161027061025c565b600181602092548386880101520192019161029b565b60ff191660208087019190915292151560051b85019092019250839150610270905061025c565b806060602080938901015201610183565b346100a85760203660031901126100a8576103196108ff565b6110023303610104576001600160a01b03165f8181526002602052604090205461033f57005b5f525f60205260405f2061010061ff00198254161790555f80f35b346100a85760403660031901126100a85760043567ffffffffffffffff81116100a857366023820112156100a857806004013561039681610937565b916103a46040519384610915565b8183526024602084019260051b820101903682116100a857602401915b81831061085e576024358467ffffffffffffffff82116100a857366023830112156100a85781600401356103f481610937565b926104026040519485610915565b8184526024602085019260051b820101903682116100a85760248101925b8284106107e75785854133036107d8573a6107c957805180156107ba57825181036107ab576001600160ff1b038111610799575f19810190811360011661050c5761046d905f8484610ad6565b600190815b8151811015610520576001600160a01b0361048d8284610a1e565b51165f19840184811161050c576001600160a01b03906104ad9085610a1e565b5116036104bd575b600101610472565b9160019081906001600160a01b036104d58686610a1e565b51166104e18286610a1e565b526104ec8587610a1e565b516104f78288610a1e565b526105028187610a1e565b50019290506104b5565b634e487b7160e01b5f52601160045260245ffd5b50919080835281526001545f5b8181106107165750506001545f5b8181106106f05782845f6001555f915b81518310156106b8576105706001600160a01b036105698585610a1e565b5116610c70565b5061057b8382610a1e565b51926001600160a01b0361058f8285610a1e565b51165f525f602052600160405f2001845167ffffffffffffffff81116106a4576105b98254610a46565b601f8111610669575b506020601f821160011461060457819060019596975f926105f9575b50505f19600383901b1c191690841b1790555b01919061054b565b0151905087806105de565b601f19821696835f52815f20975f5b8181106106515750916001969798918488959410610639575b505050811b0190556105f1565b01515f1960f88460031b161c1916905587808061062c565b92986020600181928c8601518155019a019301610613565b61069490835f5260205f20601f840160051c8101916020851061069a575b601f0160051c0190610c42565b866105c2565b9091508190610687565b634e487b7160e01b5f52604160045260245ffd5b6106eb7f835b57ecee103e9bbc8739ee36b503452a571775c6e409f666c04644140418059160405191829133958361094f565b0390a2005b806106fc600192610c58565b90549060031b1c5f5260026020525f60408120550161053b565b80610722600192610c58565b838060a01b0391549060031b1c165f525f6020528160405f205f8155016107498154610a46565b9081610758575b50500161052d565b81601f5f9311851461076e5750555b8580610750565b8183526020832061078991601f0160051c8101908601610c42565b8082528160208120915555610767565b63123baf0360e11b5f5260045260245ffd5b631fec674760e31b5f5260045ffd5b6313f8a3a760e31b5f5260045ffd5b6383f1b1d360e01b5f5260045ffd5b63022d8c9560e31b5f5260045ffd5b833567ffffffffffffffff81116100a8578201366043820112156100a85760248101359167ffffffffffffffff83116106a457604051610831601f8501601f191660200182610915565b83815236604484860101116100a8575f602085819660448397018386013783010152815201930192610420565b82356001600160a01b03811681036100a8578152602092830192016103c1565b346100a8575f3660031901126100a8576020600154604051908152f35b346100a85760203660031901126100a8576001600160a01b036108bc6108ff565b166108d2815f52600260205260405f2054151590565b806108e5575b6020906040519015158152f35b505f525f60205260206108fa60405f20610a7e565b6108d8565b600435906001600160a01b03821682036100a857565b90601f8019910116810190811067ffffffffffffffff8211176106a457604052565b67ffffffffffffffff81116106a45760051b60200190565b604081016040825282518091526020606083019301905f5b8181106109e7575050506020818303910152815180825260208201916020808360051b8301019401925f915b8383106109a257505050505090565b909192939460208080600193601f19868203018752818a518051918291828552018484015e5f828201840152601f01601f191601019701959491909101920190610993565b82516001600160a01b0316855260209485019490920191600101610967565b81810392915f13801582851316918412161761050c57565b8051821015610a325760209160051b010190565b634e487b7160e01b5f52603260045260245ffd5b90600182811c92168015610a74575b6020831014610a6057565b634e487b7160e01b5f52602260045260245ffd5b91607f1691610a55565b80549060ff8216159182610aac575b5081610a97575090565b603091506001610aa8910154610a46565b1490565b60081c60ff161591505f610a8d565b9190915f838201938412911290801582169115161761050c57565b929190828214610b42578183610b27610b1b610b0e610b08610b02610afb878b610a06565b6002900590565b86610abb565b89610a1e565b516001600160a01b031690565b6001600160a01b031690565b81851315610b6c5750808212610b5a575b5050828212610b48575b50505050565b610b5193610ad6565b5f808080610b42565b610b65918387610ad6565b5f80610b38565b959490969192935b86610b85610b1b610b0e848a610a1e565b1015610b9357600101610b74565b9495965b87610ba8610b1b610b0e848b610a1e565b1115610bb6575f1901610b97565b909493929196818513610b27579093600190610c0b610bd8610b0e888b610a1e565b610c01610be8610b0e858d610a1e565b610bf28a8d610a1e565b6001600160a01b039091169052565b610bf2838b610a1e565b610c158686610a1e565b51610c208287610a1e565b51610c2b8888610a1e565b52610c368287610a1e565b5201935f190190610b27565b818110610c4d575050565b5f8155600101610c42565b600154811015610a325760015f5260205f2001905f90565b805f52600260205260405f2054155f14610cca57600154680100000000000000008110156106a45760018101600155600154811015610a3257819060015f5260205f200155600154905f52600260205260405f2055600190565b505f9056fea2646970667358221220135c57382bce838dec28c8f3e6c3a2ce8f2e342b2dfd1a8d15bfcf47c6e5752e64736f6c634300081e0033",
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

// PackGetActiveValidators is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9de70258.
//
// Solidity: function getActiveValidators() view returns(address[] validatorAddrs, bytes[] signerAddrs)
func (validatorSet *ValidatorSet) PackGetActiveValidators() []byte {
	enc, err := validatorSet.abi.Pack("getActiveValidators")
	if err != nil {
		panic(err)
	}
	return enc
}

// GetActiveValidatorsOutput serves as a container for the return parameters of contract
// method GetActiveValidators.
type GetActiveValidatorsOutput struct {
	ValidatorAddrs []common.Address
	SignerAddrs    [][]byte
}

// UnpackGetActiveValidators is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9de70258.
//
// Solidity: function getActiveValidators() view returns(address[] validatorAddrs, bytes[] signerAddrs)
func (validatorSet *ValidatorSet) UnpackGetActiveValidators(data []byte) (GetActiveValidatorsOutput, error) {
	out, err := validatorSet.abi.Unpack("getActiveValidators", data)
	outstruct := new(GetActiveValidatorsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.ValidatorAddrs = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.SignerAddrs = *abi.ConvertType(out[1], new([][]byte)).(*[][]byte)
	return *outstruct, err

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

// PackIsActiveValidator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x40550a1c.
//
// Solidity: function isActiveValidator(address validator) view returns(bool)
func (validatorSet *ValidatorSet) PackIsActiveValidator(validator common.Address) []byte {
	enc, err := validatorSet.abi.Pack("isActiveValidator", validator)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsActiveValidator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x40550a1c.
//
// Solidity: function isActiveValidator(address validator) view returns(bool)
func (validatorSet *ValidatorSet) UnpackIsActiveValidator(data []byte) (bool, error) {
	out, err := validatorSet.abi.Unpack("isActiveValidator", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
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

// PackJail is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9bcbea52.
//
// Solidity: function jail(address validator) returns()
func (validatorSet *ValidatorSet) PackJail(validator common.Address) []byte {
	enc, err := validatorSet.abi.Pack("jail", validator)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUnstake is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2888dbb.
//
// Solidity: function unstake(address validator) returns()
func (validatorSet *ValidatorSet) PackUnstake(validator common.Address) []byte {
	enc, err := validatorSet.abi.Pack("unstake", validator)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUpdateValidators is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8ee214d3.
//
// Solidity: function updateValidators(address[] validatorAddrs, bytes[] signerAddrs) returns()
func (validatorSet *ValidatorSet) PackUpdateValidators(validatorAddrs []common.Address, signerAddrs [][]byte) []byte {
	enc, err := validatorSet.abi.Pack("updateValidators", validatorAddrs, signerAddrs)
	if err != nil {
		panic(err)
	}
	return enc
}

// ValidatorSetValidatorsUpdated represents a ValidatorsUpdated event raised by the ValidatorSet contract.
type ValidatorSetValidatorsUpdated struct {
	Updater        common.Address
	ValidatorAddrs []common.Address
	SignerAddrs    [][]byte
	Raw            *types.Log // Blockchain specific contextual infos
}

const ValidatorSetValidatorsUpdatedEventName = "ValidatorsUpdated"

// ContractEventName returns the user-defined event name.
func (ValidatorSetValidatorsUpdated) ContractEventName() string {
	return ValidatorSetValidatorsUpdatedEventName
}

// UnpackValidatorsUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ValidatorsUpdated(address indexed updater, address[] validatorAddrs, bytes[] signerAddrs)
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
	if bytes.Equal(raw[:4], validatorSet.abi.Errors["LengthMismatch"].ID.Bytes()[:4]) {
		return validatorSet.UnpackLengthMismatchError(raw[4:])
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
	if bytes.Equal(raw[:4], validatorSet.abi.Errors["SafeCastOverflowedUintToInt"].ID.Bytes()[:4]) {
		return validatorSet.UnpackSafeCastOverflowedUintToIntError(raw[4:])
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

// ValidatorSetLengthMismatch represents a LengthMismatch error raised by the ValidatorSet contract.
type ValidatorSetLengthMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LengthMismatch()
func ValidatorSetLengthMismatchErrorID() common.Hash {
	return common.HexToHash("0xff633a3803c58b9bc21e58efecee59f27e033cc0b1883fccb4969c76146fe60f")
}

// UnpackLengthMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LengthMismatch()
func (validatorSet *ValidatorSet) UnpackLengthMismatchError(raw []byte) (*ValidatorSetLengthMismatch, error) {
	out := new(ValidatorSetLengthMismatch)
	if err := validatorSet.abi.UnpackIntoInterface(out, "LengthMismatch", raw); err != nil {
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

// ValidatorSetSafeCastOverflowedUintToInt represents a SafeCastOverflowedUintToInt error raised by the ValidatorSet contract.
type ValidatorSetSafeCastOverflowedUintToInt struct {
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCastOverflowedUintToInt(uint256 value)
func ValidatorSetSafeCastOverflowedUintToIntErrorID() common.Hash {
	return common.HexToHash("0x24775e0629ae69d78c11bae050651b81820407f300ff750ff2be51e4ce75c37f")
}

// UnpackSafeCastOverflowedUintToIntError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCastOverflowedUintToInt(uint256 value)
func (validatorSet *ValidatorSet) UnpackSafeCastOverflowedUintToIntError(raw []byte) (*ValidatorSetSafeCastOverflowedUintToInt, error) {
	out := new(ValidatorSetSafeCastOverflowedUintToInt)
	if err := validatorSet.abi.UnpackIntoInterface(out, "SafeCastOverflowedUintToInt", raw); err != nil {
		return nil, err
	}
	return out, nil
}
