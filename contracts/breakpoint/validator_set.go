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
	ABI:        "[{\"type\":\"function\",\"name\":\"getValidatorCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isActiveValidator\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isValidator\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"jail\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unstake\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateValidators\",\"inputs\":[{\"name\":\"validatorAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ValidatorsUpdated\",\"inputs\":[{\"name\":\"updater\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"validators\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidValue\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NoValidators\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCoinbase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySystemContract\",\"inputs\":[{\"name\":\"contractAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyZeroGasPrice\",\"inputs\":[]}]",
	ID:         "ValidatorSet",
	BinRuntime: "0x60806040526004361015610011575f80fd5b5f3560e01c806340550a1c146100845780637071688a1461007f5780639bcbea521461007a578063b7ab4db514610075578063e71731e414610070578063f2888dbb1461006b5763facd743b14610066575f80fd5b610369565b610311565b6102b2565b6101d0565b61011f565b610102565b346100e85760203660031901126100e8576001600160a01b036100a56100ec565b166100bb815f52600260205260405f2054151590565b806100ce575b6020906040519015158152f35b505f525f60205260206100e360405f206105b9565b6100c1565b5f80fd5b600435906001600160a01b03821682036100e857565b346100e8575f3660031901126100e8576020600154604051908152f35b346100e85760203660031901126100e8576101386100ec565b6110023303610179576001600160a01b03165f8181526002602052604090205461015e57005b5f525f60205260405f2061010061ff00198254161790555f80f35b630f22c43960e41b5f5261100260045260245ffd5b60206040818301928281528451809452019201905f5b8181106101b15750505090565b82516001600160a01b03168452602093840193909201916001016101a4565b346100e8575f3660031901126100e8575f6001546101f56101f0826103ea565b6103bf565b81815290601f19610205826103ea565b013660208401375f5b81811061022a5783835260405180610226858261018e565b0390f35b61026061025b610239836106b0565b905460039190911b1c6001600160a01b03165f90815260208190526040902090565b6105b9565b61026d575b60010161020e565b92600180916102a9610290610284610284896106c8565b6001600160a01b031690565b61029a8388610416565b6001600160a01b039091169052565b01939050610265565b346100e85760203660031901126100e85760043567ffffffffffffffff81116100e857366023820112156100e857806004013567ffffffffffffffff81116100e8573660248260051b840101116100e857602461030f920161042f565b005b346100e85760203660031901126100e85761032a6100ec565b6110023303610179576001600160a01b03165f8181526002602052604090205461035057005b5f525f60205260405f20600160ff198254161790555f80f35b346100e85760203660031901126100e85760206103a16001600160a01b0361038f6100ec565b165f52600260205260405f2054151590565b6040519015158152f35b634e487b7160e01b5f52604160045260245ffd5b6040519190601f01601f1916820167ffffffffffffffff8111838210176103e557604052565b6103ab565b67ffffffffffffffff81116103e55760051b60200190565b634e487b7160e01b5f52603260045260245ffd5b805182101561042a5760209160051b010190565b610402565b90413303610537573a6105285780156105195761045691610451913691610546565b6105d5565b600190815b81518110156104d75761047e6104718284610416565b516001600160a01b031690565b61049661028461047161049087610597565b86610416565b6001600160a01b03909116036104af575b60010161045b565b91600180916104ce6104c46104718787610416565b61029a8387610416565b019290506104a7565b509081526104e4816105f0565b7fbb68ac2a5765ef54e5086dfb2b6dd55c8890145a9257a5e5f57c4c387b9e0aea6040518061051433948261018e565b0390a2565b6313f8a3a760e31b5f5260045ffd5b6383f1b1d360e01b5f5260045ffd5b63022d8c9560e31b5f5260045ffd5b92916105546101f0836103ea565b93828552602085019260051b81019182116100e857915b81831061057757505050565b82356001600160a01b03811681036100e85781526020928301920161056b565b5f198101919082116105a557565b634e487b7160e01b5f52601160045260245ffd5b5460ff81161590816105c9575090565b60ff915060081c161590565b6105ed60016020835160051b84010160208401610710565b90565b906001545f5b8181106106845750506001545f5b8181106106405750505f6001555f5b825181101561063b578061063461062f61047160019487610416565b6106fc565b5001610613565b509050565b60015481101561042a577fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf68101545f90815260026020526040812055600101610604565b806106906001926106b0565b838060a01b0391549060031b1c165f525f6020525f6040812055016105f6565b60015481101561042a5760015f5260205f2001905f90565b60015481101561042a5760015f527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6015490565b6105ed906001600160a01b0316600161079d565b91906040838203106107985782519282818095602084015b858110610750575050825181518452815261074292610710565b602061074e9301610710565b565b9150915080518560011461077257634e487b7160e01b5f52605160045260245ffd5b8211610785575b60200184918691610728565b6020909501805186518252865294610779565b505050565b6001810190825f528160205260405f2054155f146107f3578054680100000000000000008110156103e5576001810180835581101561042a575f828152602090200183905554915f5260205260405f2055600190565b5050505f9056fea26469706673582212202ee090f50859b442ac612b9b082ad8f5565c9374ff6b8ae5c89f25c9fdc7b80564736f6c634300081c0033",
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
	Updater    common.Address
	Validators []common.Address
	Raw        *types.Log // Blockchain specific contextual infos
}

const ValidatorSetValidatorsUpdatedEventName = "ValidatorsUpdated"

// ContractEventName returns the user-defined event name.
func (ValidatorSetValidatorsUpdated) ContractEventName() string {
	return ValidatorSetValidatorsUpdatedEventName
}

// UnpackValidatorsUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ValidatorsUpdated(address indexed updater, address[] validators)
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
