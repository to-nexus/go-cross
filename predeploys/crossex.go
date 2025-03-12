// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package predeploys

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// CrossExMetaData contains all meta data concerning the CrossEx contract.
var CrossExMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"LogTransfer\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"input1\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"input2\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"output1\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"output2\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// CrossExABI is the input ABI used to generate the binding from.
// Deprecated: Use CrossExMetaData.ABI instead.
var CrossExABI = CrossExMetaData.ABI

// CrossExBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const CrossExBinRuntime = "608060405234801561000f575f80fd5b5060043610610055575f3560e01c806306fdde031461005957806318160ddd14610099578063313ce567146100af57806370a08231146100be57806395d89b41146100d9575b5f80fd5b60408051808201909152601281527121b937b9b9902730ba34bb32902a37b5b2b760711b60208201525b604051610090919061011c565b60405180910390f35b6100a16100fb565b604051908152602001610090565b60405160128152602001610090565b6100a16100cc366004610168565b6001600160a01b03163190565b604080518082019091526006815265086a49ea6a6f60d31b6020820152610083565b5f6101086012600a61028b565b6101179064174876e800610296565b905090565b5f602080835283518060208501525f5b818110156101485785810183015185820160400152820161012c565b505f604082860101526040601f19601f8301168501019250505092915050565b5f60208284031215610178575f80fd5b81356001600160a01b038116811461018e575f80fd5b9392505050565b634e487b7160e01b5f52601160045260245ffd5b600181815b808511156101e357815f19048211156101c9576101c9610195565b808516156101d657918102915b93841c93908002906101ae565b509250929050565b5f826101f957506001610285565b8161020557505f610285565b816001811461021b576002811461022557610241565b6001915050610285565b60ff84111561023657610236610195565b50506001821b610285565b5060208310610133831016604e8410600b8410161715610264575081810a610285565b61026e83836101a9565b805f190482111561028157610281610195565b0290505b92915050565b5f61018e83836101eb565b80820281158282048414176102855761028561019556fea2646970667358221220ecee7b036878642a44dcf001cf156f84d6054cd98d9bb1223f4b90c2817180f864736f6c63430008180033"

// CrossEx is an auto generated Go binding around an Ethereum contract.
type CrossEx struct {
	CrossExCaller     // Read-only binding to the contract
	CrossExTransactor // Write-only binding to the contract
	CrossExFilterer   // Log filterer for contract events
}

// CrossExCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrossExCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossExTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossExTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossExFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossExFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossExSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossExSession struct {
	Contract     *CrossEx          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrossExCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossExCallerSession struct {
	Contract *CrossExCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CrossExTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossExTransactorSession struct {
	Contract     *CrossExTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CrossExRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrossExRaw struct {
	Contract *CrossEx // Generic contract binding to access the raw methods on
}

// CrossExCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossExCallerRaw struct {
	Contract *CrossExCaller // Generic read-only contract binding to access the raw methods on
}

// CrossExTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossExTransactorRaw struct {
	Contract *CrossExTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrossEx creates a new instance of CrossEx, bound to a specific deployed contract.
func NewCrossEx(address common.Address, backend bind.ContractBackend) (*CrossEx, error) {
	contract, err := bindCrossEx(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossEx{CrossExCaller: CrossExCaller{contract: contract}, CrossExTransactor: CrossExTransactor{contract: contract}, CrossExFilterer: CrossExFilterer{contract: contract}}, nil
}

// NewCrossExCaller creates a new read-only instance of CrossEx, bound to a specific deployed contract.
func NewCrossExCaller(address common.Address, caller bind.ContractCaller) (*CrossExCaller, error) {
	contract, err := bindCrossEx(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossExCaller{contract: contract}, nil
}

// NewCrossExTransactor creates a new write-only instance of CrossEx, bound to a specific deployed contract.
func NewCrossExTransactor(address common.Address, transactor bind.ContractTransactor) (*CrossExTransactor, error) {
	contract, err := bindCrossEx(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossExTransactor{contract: contract}, nil
}

// NewCrossExFilterer creates a new log filterer instance of CrossEx, bound to a specific deployed contract.
func NewCrossExFilterer(address common.Address, filterer bind.ContractFilterer) (*CrossExFilterer, error) {
	contract, err := bindCrossEx(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossExFilterer{contract: contract}, nil
}

// bindCrossEx binds a generic wrapper to an already deployed contract.
func bindCrossEx(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrossExMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossEx *CrossExRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossEx.Contract.CrossExCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossEx *CrossExRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossEx.Contract.CrossExTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossEx *CrossExRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossEx.Contract.CrossExTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossEx *CrossExCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossEx.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossEx *CrossExTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossEx.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossEx *CrossExTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossEx.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CrossEx *CrossExCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrossEx.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CrossEx *CrossExSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _CrossEx.Contract.BalanceOf(&_CrossEx.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CrossEx *CrossExCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _CrossEx.Contract.BalanceOf(&_CrossEx.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_CrossEx *CrossExCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CrossEx.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_CrossEx *CrossExSession) Decimals() (uint8, error) {
	return _CrossEx.Contract.Decimals(&_CrossEx.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_CrossEx *CrossExCallerSession) Decimals() (uint8, error) {
	return _CrossEx.Contract.Decimals(&_CrossEx.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_CrossEx *CrossExCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CrossEx.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_CrossEx *CrossExSession) Name() (string, error) {
	return _CrossEx.Contract.Name(&_CrossEx.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_CrossEx *CrossExCallerSession) Name() (string, error) {
	return _CrossEx.Contract.Name(&_CrossEx.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_CrossEx *CrossExCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CrossEx.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_CrossEx *CrossExSession) Symbol() (string, error) {
	return _CrossEx.Contract.Symbol(&_CrossEx.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_CrossEx *CrossExCallerSession) Symbol() (string, error) {
	return _CrossEx.Contract.Symbol(&_CrossEx.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() pure returns(uint256)
func (_CrossEx *CrossExCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossEx.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() pure returns(uint256)
func (_CrossEx *CrossExSession) TotalSupply() (*big.Int, error) {
	return _CrossEx.Contract.TotalSupply(&_CrossEx.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() pure returns(uint256)
func (_CrossEx *CrossExCallerSession) TotalSupply() (*big.Int, error) {
	return _CrossEx.Contract.TotalSupply(&_CrossEx.CallOpts)
}

// CrossExLogTransferIterator is returned from FilterLogTransfer and is used to iterate over the raw logs and unpacked data for LogTransfer events raised by the CrossEx contract.
type CrossExLogTransferIterator struct {
	Event *CrossExLogTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrossExLogTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossExLogTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrossExLogTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrossExLogTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossExLogTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossExLogTransfer represents a LogTransfer event raised by the CrossEx contract.
type CrossExLogTransfer struct {
	Token   common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Input1  *big.Int
	Input2  *big.Int
	Output1 *big.Int
	Output2 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogTransfer is a free log retrieval operation binding the contract event 0xe6497e3ee548a3372136af2fcb0696db31fc6cf20260707645068bd3fe97f3c4.
//
// Solidity: event LogTransfer(address indexed token, address indexed from, address indexed to, uint256 amount, uint256 input1, uint256 input2, uint256 output1, uint256 output2)
func (_CrossEx *CrossExFilterer) FilterLogTransfer(opts *bind.FilterOpts, token []common.Address, from []common.Address, to []common.Address) (*CrossExLogTransferIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CrossEx.contract.FilterLogs(opts, "LogTransfer", tokenRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CrossExLogTransferIterator{contract: _CrossEx.contract, event: "LogTransfer", logs: logs, sub: sub}, nil
}

// WatchLogTransfer is a free log subscription operation binding the contract event 0xe6497e3ee548a3372136af2fcb0696db31fc6cf20260707645068bd3fe97f3c4.
//
// Solidity: event LogTransfer(address indexed token, address indexed from, address indexed to, uint256 amount, uint256 input1, uint256 input2, uint256 output1, uint256 output2)
func (_CrossEx *CrossExFilterer) WatchLogTransfer(opts *bind.WatchOpts, sink chan<- *CrossExLogTransfer, token []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CrossEx.contract.WatchLogs(opts, "LogTransfer", tokenRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossExLogTransfer)
				if err := _CrossEx.contract.UnpackLog(event, "LogTransfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogTransfer is a log parse operation binding the contract event 0xe6497e3ee548a3372136af2fcb0696db31fc6cf20260707645068bd3fe97f3c4.
//
// Solidity: event LogTransfer(address indexed token, address indexed from, address indexed to, uint256 amount, uint256 input1, uint256 input2, uint256 output1, uint256 output2)
func (_CrossEx *CrossExFilterer) ParseLogTransfer(log types.Log) (*CrossExLogTransfer, error) {
	event := new(CrossExLogTransfer)
	if err := _CrossEx.contract.UnpackLog(event, "LogTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
