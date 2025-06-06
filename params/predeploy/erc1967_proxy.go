// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package predeploy

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

// ERC1967ProxyMetaData contains all meta data concerning the ERC1967Proxy contract.
var ERC1967ProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"}]",
	Bin: "0x60806040526040516103d03803806103d08339810160408190526100229161023c565b61002c8282610033565b5050610321565b61003c82610091565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561008557610080828261010c565b505050565b61008d61017f565b5050565b806001600160a01b03163b5f036100cb57604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051610128919061030b565b5f60405180830381855af49150503d805f8114610160576040519150601f19603f3d011682016040523d82523d5f602084013e610165565b606091505b5090925090506101768583836101a0565b95945050505050565b341561019e5760405163b398979f60e01b815260040160405180910390fd5b565b6060826101b5576101b0826101ff565b6101f8565b81511580156101cc57506001600160a01b0384163b155b156101f557604051639996b31560e01b81526001600160a01b03851660048201526024016100c2565b50805b9392505050565b80511561020f5780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b634e487b7160e01b5f52604160045260245ffd5b5f5f6040838503121561024d575f5ffd5b82516001600160a01b0381168114610263575f5ffd5b60208401519092506001600160401b0381111561027e575f5ffd5b8301601f8101851361028e575f5ffd5b80516001600160401b038111156102a7576102a7610228565b604051601f8201601f19908116603f011681016001600160401b03811182821017156102d5576102d5610228565b6040528181528282016020018710156102ec575f5ffd5b8160208401602083015e5f602083830101528093505050509250929050565b5f82518060208501845e5f920191825250919050565b60a38061032d5f395ff3fe6080604052600a600c565b005b60186014601a565b6050565b565b5f604b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b905090565b365f5f375f5f365f845af43d5f5f3e8080156069573d5ff35b3d5ffdfea26469706673582212206e7e0d145b12f4fa81ce7fea672667da17050f1c765d58d882513fba53d69c5364736f6c634300081c0033",
}

// ERC1967ProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC1967ProxyMetaData.ABI instead.
var ERC1967ProxyABI = ERC1967ProxyMetaData.ABI

// ERC1967ProxyBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const ERC1967ProxyBinRuntime = "6080604052600a600c565b005b60186014601a565b6050565b565b5f604b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b905090565b365f5f375f5f365f845af43d5f5f3e8080156069573d5ff35b3d5ffdfea26469706673582212206e7e0d145b12f4fa81ce7fea672667da17050f1c765d58d882513fba53d69c5364736f6c634300081c0033"

// ERC1967ProxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC1967ProxyMetaData.Bin instead.
var ERC1967ProxyBin = ERC1967ProxyMetaData.Bin

// DeployERC1967Proxy deploys a new Ethereum contract, binding an instance of ERC1967Proxy to it.
func DeployERC1967Proxy(auth *bind.TransactOpts, backend bind.ContractBackend, implementation common.Address, _data []byte) (common.Address, *types.Transaction, *ERC1967Proxy, error) {
	parsed, err := ERC1967ProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC1967ProxyBin), backend, implementation, _data)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC1967Proxy{ERC1967ProxyCaller: ERC1967ProxyCaller{contract: contract}, ERC1967ProxyTransactor: ERC1967ProxyTransactor{contract: contract}, ERC1967ProxyFilterer: ERC1967ProxyFilterer{contract: contract}}, nil
}

// ERC1967Proxy is an auto generated Go binding around an Ethereum contract.
type ERC1967Proxy struct {
	ERC1967ProxyCaller     // Read-only binding to the contract
	ERC1967ProxyTransactor // Write-only binding to the contract
	ERC1967ProxyFilterer   // Log filterer for contract events
}

// ERC1967ProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC1967ProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1967ProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC1967ProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1967ProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC1967ProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1967ProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC1967ProxySession struct {
	Contract     *ERC1967Proxy     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC1967ProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC1967ProxyCallerSession struct {
	Contract *ERC1967ProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ERC1967ProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC1967ProxyTransactorSession struct {
	Contract     *ERC1967ProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ERC1967ProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC1967ProxyRaw struct {
	Contract *ERC1967Proxy // Generic contract binding to access the raw methods on
}

// ERC1967ProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC1967ProxyCallerRaw struct {
	Contract *ERC1967ProxyCaller // Generic read-only contract binding to access the raw methods on
}

// ERC1967ProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC1967ProxyTransactorRaw struct {
	Contract *ERC1967ProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC1967Proxy creates a new instance of ERC1967Proxy, bound to a specific deployed contract.
func NewERC1967Proxy(address common.Address, backend bind.ContractBackend) (*ERC1967Proxy, error) {
	contract, err := bindERC1967Proxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC1967Proxy{ERC1967ProxyCaller: ERC1967ProxyCaller{contract: contract}, ERC1967ProxyTransactor: ERC1967ProxyTransactor{contract: contract}, ERC1967ProxyFilterer: ERC1967ProxyFilterer{contract: contract}}, nil
}

// NewERC1967ProxyCaller creates a new read-only instance of ERC1967Proxy, bound to a specific deployed contract.
func NewERC1967ProxyCaller(address common.Address, caller bind.ContractCaller) (*ERC1967ProxyCaller, error) {
	contract, err := bindERC1967Proxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1967ProxyCaller{contract: contract}, nil
}

// NewERC1967ProxyTransactor creates a new write-only instance of ERC1967Proxy, bound to a specific deployed contract.
func NewERC1967ProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC1967ProxyTransactor, error) {
	contract, err := bindERC1967Proxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1967ProxyTransactor{contract: contract}, nil
}

// NewERC1967ProxyFilterer creates a new log filterer instance of ERC1967Proxy, bound to a specific deployed contract.
func NewERC1967ProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC1967ProxyFilterer, error) {
	contract, err := bindERC1967Proxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC1967ProxyFilterer{contract: contract}, nil
}

// bindERC1967Proxy binds a generic wrapper to an already deployed contract.
func bindERC1967Proxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC1967ProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1967Proxy *ERC1967ProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1967Proxy.Contract.ERC1967ProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1967Proxy *ERC1967ProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1967Proxy.Contract.ERC1967ProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1967Proxy *ERC1967ProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1967Proxy.Contract.ERC1967ProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1967Proxy *ERC1967ProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1967Proxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1967Proxy *ERC1967ProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1967Proxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1967Proxy *ERC1967ProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1967Proxy.Contract.contract.Transact(opts, method, params...)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ERC1967Proxy *ERC1967ProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ERC1967Proxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ERC1967Proxy *ERC1967ProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ERC1967Proxy.Contract.Fallback(&_ERC1967Proxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ERC1967Proxy *ERC1967ProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ERC1967Proxy.Contract.Fallback(&_ERC1967Proxy.TransactOpts, calldata)
}

// ERC1967ProxyUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ERC1967Proxy contract.
type ERC1967ProxyUpgradedIterator struct {
	Event *ERC1967ProxyUpgraded // Event containing the contract specifics and raw log

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
func (it *ERC1967ProxyUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1967ProxyUpgraded)
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
		it.Event = new(ERC1967ProxyUpgraded)
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
func (it *ERC1967ProxyUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1967ProxyUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1967ProxyUpgraded represents a Upgraded event raised by the ERC1967Proxy contract.
type ERC1967ProxyUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC1967Proxy *ERC1967ProxyFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ERC1967ProxyUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ERC1967Proxy.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ERC1967ProxyUpgradedIterator{contract: _ERC1967Proxy.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC1967Proxy *ERC1967ProxyFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ERC1967ProxyUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ERC1967Proxy.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1967ProxyUpgraded)
				if err := _ERC1967Proxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC1967Proxy *ERC1967ProxyFilterer) ParseUpgraded(log types.Log) (*ERC1967ProxyUpgraded, error) {
	event := new(ERC1967ProxyUpgraded)
	if err := _ERC1967Proxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
