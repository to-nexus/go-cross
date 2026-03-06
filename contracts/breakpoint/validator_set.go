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
	ABI:        "[{\"type\":\"function\",\"name\":\"getActiveValidators\",\"inputs\":[],\"outputs\":[{\"name\":\"validatorAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"signerAddrs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isActiveValidator\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isValidator\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"jail\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"suspend\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unstake\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateValidators\",\"inputs\":[{\"name\":\"validatorAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"signerAddrs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ValidatorsUpdated\",\"inputs\":[{\"name\":\"updater\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"validatorAddrs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"signerAddrs\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidSigner\",\"inputs\":[{\"name\":\"signerAddr\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidValidator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValue\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"LengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoValidators\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCoinbase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySystemContract\",\"inputs\":[{\"name\":\"contractAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyZeroGasPrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintToInt\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	ID:         "ValidatorSet",
	BinRuntime: "0x60806040526004361015610011575f80fd5b5f3560e01c8063286781c71461098757806340550a1c146109235780637071688a146109065780638ee214d3146103655780639bcbea521461030b5780639de7025814610124578063f2888dbb146100b75763facd743b14610071575f80fd5b346100b35760203660031901126100b35760206100a96001600160a01b036100976109e3565b165f52600260205260405f2054151590565b6040519015158152f35b5f80fd5b346100b35760203660031901126100b3576100d06109e3565b611002330361010f576001600160a01b03165f818152600260205260409020546100f657005b5f525f60205260405f20600160ff198254161790555f80f35b630f22c43960e41b5f5261100260045260245ffd5b346100b3575f3660031901126100b3576001545f61014182610a1b565b9061014f60405192836109f9565b828252601f1961015e84610a1b565b0136602084013761016e83610a1b565b9261017c60405194856109f9565b808452601f1961018b82610a1b565b015f5b8181106102fa5750505f5b8181106101ba57505080825282526101b660405192839283610a57565b0390f35b6101c381610d5e565b60018060a01b0391549060031b1c165f525f6020526101e460405f20610b6e565b6101f1575b600101610199565b916101fb83610d5e565b905460039190911b1c6001600160a01b03166102178286610af6565b5261022183610d5e565b60018060a01b0391549060031b1c165f525f602052600160405f2001604051905f9080549061024f82610b36565b80855291600181169081156102d3575060011461029a575b5050918161027b60019493859403826109f9565b6102858289610af6565b526102908188610af6565b50019290506101e9565b5f908152602081209092505b8183106102bd57505081016020018161027b610267565b60018160209254838688010152019201916102a6565b60ff191660208087019190915292151560051b8501909201925083915061027b9050610267565b80606060208093890101520161018e565b346100b35760203660031901126100b3576103246109e3565b611002330361010f576001600160a01b03165f8181526002602052604090205461034a57005b5f525f60205260405f2061010061ff00198254161790555f80f35b346100b35760403660031901126100b35760043567ffffffffffffffff81116100b357366023820112156100b35780600401356103a181610a1b565b916103af60405193846109f9565b8183526024602084019260051b820101903682116100b357602401915b8183106108e6578360243567ffffffffffffffff81116100b357366023820112156100b35780600401356103ff81610a1b565b9161040d60405193846109f9565b8183526024602084019260051b820101903682116100b35760248101925b82841061086f578585413303610860573a61085157815180156108425781518103610833575f5b8181106107bd57506001600160ff1b0381116107ab575f19810160018282131661079757610482905f8486610bdc565b600190815b81811061070657505080835281526001545f5b8181106106835750506001545f5b81811061065d5782845f6001555f915b8151831015610625576104dd6001600160a01b036104d68585610af6565b5116610d76565b506104e88382610af6565b51926001600160a01b036104fc8285610af6565b51165f525f602052600160405f2001845167ffffffffffffffff8111610611576105268254610b36565b601f81116105d6575b506020601f821160011461057157819060019596975f92610566575b50505f19600383901b1c191690841b1790555b0191906104b8565b01519050878061054b565b601f19821696835f52815f20975f5b8181106105be57509160019697989184889594106105a6575b505050811b01905561055e565b01515f1960f88460031b161c19169055878080610599565b92986020600181928c8601518155019a019301610580565b61060190835f5260205f20601f840160051c81019160208510610607575b601f0160051c0190610d48565b8661052f565b90915081906105f4565b634e487b7160e01b5f52604160045260245ffd5b6106587f835b57ecee103e9bbc8739ee36b503452a571775c6e409f666c046441404180591604051918291339583610a57565b0390a2005b80610669600192610d5e565b90549060031b1c5f5260026020525f6040812055016104a8565b8061068f600192610d5e565b838060a01b0391549060031b1c165f525f6020528160405f205f8155016106b68154610b36565b90816106c5575b50500161049a565b81601f5f931185146106db5750555b85806106bd565b818352602083206106f691601f0160051c8101908601610d48565b80825281602081209155556106d4565b6001600160a01b036107188287610af6565b51165f198401848111610797576001600160a01b03906107389088610af6565b511603610748575b600101610487565b9160019081906001600160a01b036107608689610af6565b511661076c8289610af6565b526107778587610af6565b516107828288610af6565b5261078d8187610af6565b5001929050610740565b634e487b7160e01b5f52601160045260245ffd5b63123baf0360e11b5f5260045260245ffd5b6001600160a01b036107cf8286610af6565b5116156108245760306107e28285610af6565b5151036107f157600101610452565b6107fe6108209184610af6565b5160405163158e2ce560e01b8152602060048201529182916024830190610a33565b0390fd5b631a0a9b9f60e21b5f5260045ffd5b631fec674760e31b5f5260045ffd5b6313f8a3a760e31b5f5260045ffd5b6383f1b1d360e01b5f5260045ffd5b63022d8c9560e31b5f5260045ffd5b833567ffffffffffffffff81116100b3578201366043820112156100b35760248101359167ffffffffffffffff8311610611576040516108b9601f8501601f1916602001826109f9565b83815236604484860101116100b3575f60208581966044839701838601378301015281520193019261042b565b82356001600160a01b03811681036100b3578152602092830192016103cc565b346100b3575f3660031901126100b3576020600154604051908152f35b346100b35760203660031901126100b3576001600160a01b036109446109e3565b1661095a815f52600260205260405f2054151590565b8061096d575b6020906040519015158152f35b505f525f602052602061098260405f20610b6e565b610960565b346100b35760203660031901126100b3576109a06109e3565b611002330361010f576001600160a01b03165f818152600260205260409020546109c657005b5f525f60205260405f206201000062ff0000198254161790555f80f35b600435906001600160a01b03821682036100b357565b90601f8019910116810190811067ffffffffffffffff82111761061157604052565b67ffffffffffffffff81116106115760051b60200190565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b604081016040825282518091526020606083019301905f5b818110610ad7575050506020818303910152815180825260208201916020808360051b8301019401925f915b838310610aaa57505050505090565b9091929394602080610ac8600193601f198682030187528951610a33565b97019301930191939290610a9b565b82516001600160a01b0316855260209485019490920191600101610a6f565b8051821015610b0a5760209160051b010190565b634e487b7160e01b5f52603260045260245ffd5b81810392915f13801582851316918412161761079757565b90600182811c92168015610b64575b6020831014610b5057565b634e487b7160e01b5f52602260045260245ffd5b91607f1691610b45565b80549060ff8216159182610bb1575b82610ba2575b5081610b8d575090565b603091506001610b9e910154610b36565b1490565b60101c60ff161591505f610b83565b915060ff8260081c161591610b7d565b9190915f838201938412911290801582169115161761079757565b929190828214610c48578183610c2d610c21610c14610c0e610c08610c01878b610b1e565b6002900590565b86610bc1565b89610af6565b516001600160a01b031690565b6001600160a01b031690565b81851315610c725750808212610c60575b5050828212610c4e575b50505050565b610c5793610bdc565b5f808080610c48565b610c6b918387610bdc565b5f80610c3e565b959490969192935b86610c8b610c21610c14848a610af6565b1015610c9957600101610c7a565b9495965b87610cae610c21610c14848b610af6565b1115610cbc575f1901610c9d565b909493929196818513610c2d579093600190610d11610cde610c14888b610af6565b610d07610cee610c14858d610af6565b610cf88a8d610af6565b6001600160a01b039091169052565b610cf8838b610af6565b610d1b8686610af6565b51610d268287610af6565b51610d318888610af6565b52610d3c8287610af6565b5201935f190190610c2d565b818110610d53575050565b5f8155600101610d48565b600154811015610b0a5760015f5260205f2001905f90565b805f52600260205260405f2054155f14610dd057600154680100000000000000008110156106115760018101600155600154811015610b0a57819060015f5260205f200155600154905f52600260205260405f2055600190565b505f9056fea26469706673582212206a4eba792955d34591f006840f697acb454faffc47d335017a81425c2bb3f74164736f6c634300081e0033",
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

// PackSuspend is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x286781c7.
//
// Solidity: function suspend(address validator) returns()
func (validatorSet *ValidatorSet) PackSuspend(validator common.Address) []byte {
	enc, err := validatorSet.abi.Pack("suspend", validator)
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
	if bytes.Equal(raw[:4], validatorSet.abi.Errors["InvalidSigner"].ID.Bytes()[:4]) {
		return validatorSet.UnpackInvalidSignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], validatorSet.abi.Errors["InvalidValidator"].ID.Bytes()[:4]) {
		return validatorSet.UnpackInvalidValidatorError(raw[4:])
	}
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

// ValidatorSetInvalidSigner represents a InvalidSigner error raised by the ValidatorSet contract.
type ValidatorSetInvalidSigner struct {
	SignerAddr []byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidSigner(bytes signerAddr)
func ValidatorSetInvalidSignerErrorID() common.Hash {
	return common.HexToHash("0x158e2ce56dd0a6d142fc86ecff1750a53c20cf498df396eab667797120538c2c")
}

// UnpackInvalidSignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidSigner(bytes signerAddr)
func (validatorSet *ValidatorSet) UnpackInvalidSignerError(raw []byte) (*ValidatorSetInvalidSigner, error) {
	out := new(ValidatorSetInvalidSigner)
	if err := validatorSet.abi.UnpackIntoInterface(out, "InvalidSigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorSetInvalidValidator represents a InvalidValidator error raised by the ValidatorSet contract.
type ValidatorSetInvalidValidator struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidValidator()
func ValidatorSetInvalidValidatorErrorID() common.Hash {
	return common.HexToHash("0x682a6e7c20600f361eaefb598cbd5420061f2b13ac94bbde7c6f97af1ebeec30")
}

// UnpackInvalidValidatorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidValidator()
func (validatorSet *ValidatorSet) UnpackInvalidValidatorError(raw []byte) (*ValidatorSetInvalidValidator, error) {
	out := new(ValidatorSetInvalidValidator)
	if err := validatorSet.abi.UnpackIntoInterface(out, "InvalidValidator", raw); err != nil {
		return nil, err
	}
	return out, nil
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
