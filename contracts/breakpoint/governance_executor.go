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

// GovernanceExecutorMetaData contains all meta data concerning the GovernanceExecutor contract.
var GovernanceExecutorMetaData = bind.MetaData{
	ABI:        "[{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidValue\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotSystemContract\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyCoinbase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySystemContract\",\"inputs\":[{\"name\":\"contractAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyZeroGasPrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
	ID:         "GovernanceExecutor",
	BinRuntime: "0x60806040526004361015610011575f80fd5b5f3560e01c631cff79cd14610024575f80fd5b6040366003190112610225576004356001600160a01b038116908181036102255760243567ffffffffffffffff811161022557366023820112156102255780600401359267ffffffffffffffff8411610225573660248584010111610225576110123303610210577f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005c6102015760017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005d611000811480156101f6575b80156101eb575b80156101e0575b80156101d5575b80156101ca575b80156101bf575b156101ad57505f602061011f61011a86610263565b610229565b948086528060248388019501853785010152344710610196575f809161016894519034855af13d1561018e573d9161015961011a84610263565b9283523d5f602085013e61027f565b505f7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005d005b60609161027f565b4763cf47918160e01b5f526004523460245260445ffd5b6312abf22b60e11b5f5260045260245ffd5b506110128114610105565b5061101081146100fe565b5061100581146100f7565b5061100481146100f0565b5061100281146100e9565b5061100181146100e2565b633ee5aeb560e01b5f5260045ffd5b630f22c43960e41b5f5261101260045260245ffd5b5f80fd5b6040519190601f01601f1916820167ffffffffffffffff81118382101761024f57604052565b634e487b7160e01b5f52604160045260245ffd5b67ffffffffffffffff811161024f57601f01601f191660200190565b906102a3575080511561029457602081519101fd5b63d6bda27560e01b5f5260045ffd5b815115806102d4575b6102b4575090565b639996b31560e01b5f9081526001600160a01b0391909116600452602490fd5b50803b156102ac56fea26469706673582212200b01527a9052d916450975f71c52f7e485c5bf9a7f7d02df687515b22894766564736f6c634300081c0033",
}

// GovernanceExecutor is an auto generated Go binding around an Ethereum contract.
type GovernanceExecutor struct {
	abi abi.ABI
}

// NewGovernanceExecutor creates a new instance of GovernanceExecutor.
func NewGovernanceExecutor() *GovernanceExecutor {
	parsed, err := GovernanceExecutorMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &GovernanceExecutor{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *GovernanceExecutor) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackExecute is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1cff79cd.
//
// Solidity: function execute(address target, bytes data) payable returns()
func (governanceExecutor *GovernanceExecutor) PackExecute(target common.Address, data []byte) []byte {
	enc, err := governanceExecutor.abi.Pack("execute", target, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (governanceExecutor *GovernanceExecutor) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], governanceExecutor.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return governanceExecutor.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceExecutor.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return governanceExecutor.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceExecutor.abi.Errors["InsufficientBalance"].ID.Bytes()[:4]) {
		return governanceExecutor.UnpackInsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceExecutor.abi.Errors["InvalidValue"].ID.Bytes()[:4]) {
		return governanceExecutor.UnpackInvalidValueError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceExecutor.abi.Errors["NotSystemContract"].ID.Bytes()[:4]) {
		return governanceExecutor.UnpackNotSystemContractError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceExecutor.abi.Errors["OnlyCoinbase"].ID.Bytes()[:4]) {
		return governanceExecutor.UnpackOnlyCoinbaseError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceExecutor.abi.Errors["OnlySystemContract"].ID.Bytes()[:4]) {
		return governanceExecutor.UnpackOnlySystemContractError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceExecutor.abi.Errors["OnlyZeroGasPrice"].ID.Bytes()[:4]) {
		return governanceExecutor.UnpackOnlyZeroGasPriceError(raw[4:])
	}
	if bytes.Equal(raw[:4], governanceExecutor.abi.Errors["ReentrancyGuardReentrantCall"].ID.Bytes()[:4]) {
		return governanceExecutor.UnpackReentrancyGuardReentrantCallError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// GovernanceExecutorAddressEmptyCode represents a AddressEmptyCode error raised by the GovernanceExecutor contract.
type GovernanceExecutorAddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func GovernanceExecutorAddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (governanceExecutor *GovernanceExecutor) UnpackAddressEmptyCodeError(raw []byte) (*GovernanceExecutorAddressEmptyCode, error) {
	out := new(GovernanceExecutorAddressEmptyCode)
	if err := governanceExecutor.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceExecutorFailedCall represents a FailedCall error raised by the GovernanceExecutor contract.
type GovernanceExecutorFailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func GovernanceExecutorFailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (governanceExecutor *GovernanceExecutor) UnpackFailedCallError(raw []byte) (*GovernanceExecutorFailedCall, error) {
	out := new(GovernanceExecutorFailedCall)
	if err := governanceExecutor.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceExecutorInsufficientBalance represents a InsufficientBalance error raised by the GovernanceExecutor contract.
type GovernanceExecutorInsufficientBalance struct {
	Balance *big.Int
	Needed  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InsufficientBalance(uint256 balance, uint256 needed)
func GovernanceExecutorInsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xcf4791818fba6e019216eb4864093b4947f674afada5d305e57d598b641dad1d")
}

// UnpackInsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InsufficientBalance(uint256 balance, uint256 needed)
func (governanceExecutor *GovernanceExecutor) UnpackInsufficientBalanceError(raw []byte) (*GovernanceExecutorInsufficientBalance, error) {
	out := new(GovernanceExecutorInsufficientBalance)
	if err := governanceExecutor.abi.UnpackIntoInterface(out, "InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceExecutorInvalidValue represents a InvalidValue error raised by the GovernanceExecutor contract.
type GovernanceExecutorInvalidValue struct {
	Key   string
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidValue(string key, uint256 value)
func GovernanceExecutorInvalidValueErrorID() common.Hash {
	return common.HexToHash("0x2c648cf1bbb773fa5fbcfc6541df5c1891af9b6969d9a555653bcec289d77218")
}

// UnpackInvalidValueError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidValue(string key, uint256 value)
func (governanceExecutor *GovernanceExecutor) UnpackInvalidValueError(raw []byte) (*GovernanceExecutorInvalidValue, error) {
	out := new(GovernanceExecutorInvalidValue)
	if err := governanceExecutor.abi.UnpackIntoInterface(out, "InvalidValue", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceExecutorNotSystemContract represents a NotSystemContract error raised by the GovernanceExecutor contract.
type GovernanceExecutorNotSystemContract struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotSystemContract(address target)
func GovernanceExecutorNotSystemContractErrorID() common.Hash {
	return common.HexToHash("0x2557e45618676af9b2deadb772eb4a20075cc668011c612c61e09aef2f8b4b0a")
}

// UnpackNotSystemContractError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotSystemContract(address target)
func (governanceExecutor *GovernanceExecutor) UnpackNotSystemContractError(raw []byte) (*GovernanceExecutorNotSystemContract, error) {
	out := new(GovernanceExecutorNotSystemContract)
	if err := governanceExecutor.abi.UnpackIntoInterface(out, "NotSystemContract", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceExecutorOnlyCoinbase represents a OnlyCoinbase error raised by the GovernanceExecutor contract.
type GovernanceExecutorOnlyCoinbase struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyCoinbase()
func GovernanceExecutorOnlyCoinbaseErrorID() common.Hash {
	return common.HexToHash("0x116c64a8de02ce00303a109e06f26c31a3cfed64656fb9d228157fad57dff616")
}

// UnpackOnlyCoinbaseError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyCoinbase()
func (governanceExecutor *GovernanceExecutor) UnpackOnlyCoinbaseError(raw []byte) (*GovernanceExecutorOnlyCoinbase, error) {
	out := new(GovernanceExecutorOnlyCoinbase)
	if err := governanceExecutor.abi.UnpackIntoInterface(out, "OnlyCoinbase", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceExecutorOnlySystemContract represents a OnlySystemContract error raised by the GovernanceExecutor contract.
type GovernanceExecutorOnlySystemContract struct {
	ContractAddr common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlySystemContract(address contractAddr)
func GovernanceExecutorOnlySystemContractErrorID() common.Hash {
	return common.HexToHash("0xf22c4390ebf387afc834c245f4ef6f38d59454737d03e451e0d182ac61608277")
}

// UnpackOnlySystemContractError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlySystemContract(address contractAddr)
func (governanceExecutor *GovernanceExecutor) UnpackOnlySystemContractError(raw []byte) (*GovernanceExecutorOnlySystemContract, error) {
	out := new(GovernanceExecutorOnlySystemContract)
	if err := governanceExecutor.abi.UnpackIntoInterface(out, "OnlySystemContract", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceExecutorOnlyZeroGasPrice represents a OnlyZeroGasPrice error raised by the GovernanceExecutor contract.
type GovernanceExecutorOnlyZeroGasPrice struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyZeroGasPrice()
func GovernanceExecutorOnlyZeroGasPriceErrorID() common.Hash {
	return common.HexToHash("0x83f1b1d3f8cc3470adc53b59d7024697cf0374e9ddadd2111159d00fe76f2c06")
}

// UnpackOnlyZeroGasPriceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyZeroGasPrice()
func (governanceExecutor *GovernanceExecutor) UnpackOnlyZeroGasPriceError(raw []byte) (*GovernanceExecutorOnlyZeroGasPrice, error) {
	out := new(GovernanceExecutorOnlyZeroGasPrice)
	if err := governanceExecutor.abi.UnpackIntoInterface(out, "OnlyZeroGasPrice", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GovernanceExecutorReentrancyGuardReentrantCall represents a ReentrancyGuardReentrantCall error raised by the GovernanceExecutor contract.
type GovernanceExecutorReentrancyGuardReentrantCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ReentrancyGuardReentrantCall()
func GovernanceExecutorReentrancyGuardReentrantCallErrorID() common.Hash {
	return common.HexToHash("0x3ee5aeb571de7fc460830b4d0017439a1ca56fb0bc39062227ade4fe4a24c1ca")
}

// UnpackReentrancyGuardReentrantCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ReentrancyGuardReentrantCall()
func (governanceExecutor *GovernanceExecutor) UnpackReentrancyGuardReentrantCallError(raw []byte) (*GovernanceExecutorReentrancyGuardReentrantCall, error) {
	out := new(GovernanceExecutorReentrancyGuardReentrantCall)
	if err := governanceExecutor.abi.UnpackIntoInterface(out, "ReentrancyGuardReentrantCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}
