// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abis

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
)

// MetaTxForwarderForwardRequest is an auto generated low-level Go binding around an user-defined struct.
type MetaTxForwarderForwardRequest struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Gas   *big.Int
	Nonce *big.Int
	Data  []byte
}

// MetaTxForwarderMetaData contains all meta data concerning the MetaTxForwarder contract.
var MetaTxForwarderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structMetaTxForwarder.ForwardRequest\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ProxyLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"log\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structMetaTxForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structMetaTxForwarder.ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MetaTxForwarderABI is the input ABI used to generate the binding from.
// Deprecated: Use MetaTxForwarderMetaData.ABI instead.
var MetaTxForwarderABI = MetaTxForwarderMetaData.ABI

// MetaTxForwarder is an auto generated Go binding around an Ethereum contract.
type MetaTxForwarder struct {
	MetaTxForwarderCaller     // Read-only binding to the contract
	MetaTxForwarderTransactor // Write-only binding to the contract
	MetaTxForwarderFilterer   // Log filterer for contract events
}

// MetaTxForwarderCaller is an auto generated read-only Go binding around an Ethereum contract.
type MetaTxForwarderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaTxForwarderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MetaTxForwarderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaTxForwarderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MetaTxForwarderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaTxForwarderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MetaTxForwarderSession struct {
	Contract     *MetaTxForwarder  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MetaTxForwarderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MetaTxForwarderCallerSession struct {
	Contract *MetaTxForwarderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// MetaTxForwarderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MetaTxForwarderTransactorSession struct {
	Contract     *MetaTxForwarderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// MetaTxForwarderRaw is an auto generated low-level Go binding around an Ethereum contract.
type MetaTxForwarderRaw struct {
	Contract *MetaTxForwarder // Generic contract binding to access the raw methods on
}

// MetaTxForwarderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MetaTxForwarderCallerRaw struct {
	Contract *MetaTxForwarderCaller // Generic read-only contract binding to access the raw methods on
}

// MetaTxForwarderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MetaTxForwarderTransactorRaw struct {
	Contract *MetaTxForwarderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMetaTxForwarder creates a new instance of MetaTxForwarder, bound to a specific deployed contract.
func NewMetaTxForwarder(address common.Address, backend bind.ContractBackend) (*MetaTxForwarder, error) {
	contract, err := bindMetaTxForwarder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MetaTxForwarder{MetaTxForwarderCaller: MetaTxForwarderCaller{contract: contract}, MetaTxForwarderTransactor: MetaTxForwarderTransactor{contract: contract}, MetaTxForwarderFilterer: MetaTxForwarderFilterer{contract: contract}}, nil
}

// NewMetaTxForwarderCaller creates a new read-only instance of MetaTxForwarder, bound to a specific deployed contract.
func NewMetaTxForwarderCaller(address common.Address, caller bind.ContractCaller) (*MetaTxForwarderCaller, error) {
	contract, err := bindMetaTxForwarder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MetaTxForwarderCaller{contract: contract}, nil
}

// NewMetaTxForwarderTransactor creates a new write-only instance of MetaTxForwarder, bound to a specific deployed contract.
func NewMetaTxForwarderTransactor(address common.Address, transactor bind.ContractTransactor) (*MetaTxForwarderTransactor, error) {
	contract, err := bindMetaTxForwarder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MetaTxForwarderTransactor{contract: contract}, nil
}

// NewMetaTxForwarderFilterer creates a new log filterer instance of MetaTxForwarder, bound to a specific deployed contract.
func NewMetaTxForwarderFilterer(address common.Address, filterer bind.ContractFilterer) (*MetaTxForwarderFilterer, error) {
	contract, err := bindMetaTxForwarder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MetaTxForwarderFilterer{contract: contract}, nil
}

// bindMetaTxForwarder binds a generic wrapper to an already deployed contract.
func bindMetaTxForwarder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MetaTxForwarderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetaTxForwarder *MetaTxForwarderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MetaTxForwarder.Contract.MetaTxForwarderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetaTxForwarder *MetaTxForwarderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaTxForwarder.Contract.MetaTxForwarderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetaTxForwarder *MetaTxForwarderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetaTxForwarder.Contract.MetaTxForwarderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetaTxForwarder *MetaTxForwarderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MetaTxForwarder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetaTxForwarder *MetaTxForwarderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaTxForwarder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetaTxForwarder *MetaTxForwarderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetaTxForwarder.Contract.contract.Transact(opts, method, params...)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address from) view returns(uint256)
func (_MetaTxForwarder *MetaTxForwarderCaller) GetNonce(opts *bind.CallOpts, from common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MetaTxForwarder.contract.Call(opts, &out, "getNonce", from)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address from) view returns(uint256)
func (_MetaTxForwarder *MetaTxForwarderSession) GetNonce(from common.Address) (*big.Int, error) {
	return _MetaTxForwarder.Contract.GetNonce(&_MetaTxForwarder.CallOpts, from)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address from) view returns(uint256)
func (_MetaTxForwarder *MetaTxForwarderCallerSession) GetNonce(from common.Address) (*big.Int, error) {
	return _MetaTxForwarder.Contract.GetNonce(&_MetaTxForwarder.CallOpts, from)
}

// Verify is a free data retrieval call binding the contract method 0xbf5d3bdb.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,bytes) req, bytes signature) view returns(bool)
func (_MetaTxForwarder *MetaTxForwarderCaller) Verify(opts *bind.CallOpts, req MetaTxForwarderForwardRequest, signature []byte) (bool, error) {
	var out []interface{}
	err := _MetaTxForwarder.contract.Call(opts, &out, "verify", req, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0xbf5d3bdb.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,bytes) req, bytes signature) view returns(bool)
func (_MetaTxForwarder *MetaTxForwarderSession) Verify(req MetaTxForwarderForwardRequest, signature []byte) (bool, error) {
	return _MetaTxForwarder.Contract.Verify(&_MetaTxForwarder.CallOpts, req, signature)
}

// Verify is a free data retrieval call binding the contract method 0xbf5d3bdb.
//
// Solidity: function verify((address,address,uint256,uint256,uint256,bytes) req, bytes signature) view returns(bool)
func (_MetaTxForwarder *MetaTxForwarderCallerSession) Verify(req MetaTxForwarderForwardRequest, signature []byte) (bool, error) {
	return _MetaTxForwarder.Contract.Verify(&_MetaTxForwarder.CallOpts, req, signature)
}

// Execute is a paid mutator transaction binding the contract method 0x47153f82.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,bytes) req, bytes signature) payable returns(bool, bytes)
func (_MetaTxForwarder *MetaTxForwarderTransactor) Execute(opts *bind.TransactOpts, req MetaTxForwarderForwardRequest, signature []byte) (*types.Transaction, error) {
	return _MetaTxForwarder.contract.Transact(opts, "execute", req, signature)
}

// Execute is a paid mutator transaction binding the contract method 0x47153f82.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,bytes) req, bytes signature) payable returns(bool, bytes)
func (_MetaTxForwarder *MetaTxForwarderSession) Execute(req MetaTxForwarderForwardRequest, signature []byte) (*types.Transaction, error) {
	return _MetaTxForwarder.Contract.Execute(&_MetaTxForwarder.TransactOpts, req, signature)
}

// Execute is a paid mutator transaction binding the contract method 0x47153f82.
//
// Solidity: function execute((address,address,uint256,uint256,uint256,bytes) req, bytes signature) payable returns(bool, bytes)
func (_MetaTxForwarder *MetaTxForwarderTransactorSession) Execute(req MetaTxForwarderForwardRequest, signature []byte) (*types.Transaction, error) {
	return _MetaTxForwarder.Contract.Execute(&_MetaTxForwarder.TransactOpts, req, signature)
}

// MetaTxForwarderProxyLogIterator is returned from FilterProxyLog and is used to iterate over the raw logs and unpacked data for ProxyLog events raised by the MetaTxForwarder contract.
type MetaTxForwarderProxyLogIterator struct {
	Event *MetaTxForwarderProxyLog // Event containing the contract specifics and raw log

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
func (it *MetaTxForwarderProxyLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaTxForwarderProxyLog)
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
		it.Event = new(MetaTxForwarderProxyLog)
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
func (it *MetaTxForwarderProxyLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaTxForwarderProxyLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaTxForwarderProxyLog represents a ProxyLog event raised by the MetaTxForwarder contract.
type MetaTxForwarderProxyLog struct {
	Arg0 MetaTxForwarderForwardRequest
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterProxyLog is a free log retrieval operation binding the contract event 0xd03bf592f687a60523d2ddd30db6948ef1c7934c57c181027cd0f2b9049c4adc.
//
// Solidity: event ProxyLog((address,address,uint256,uint256,uint256,bytes) arg0)
func (_MetaTxForwarder *MetaTxForwarderFilterer) FilterProxyLog(opts *bind.FilterOpts) (*MetaTxForwarderProxyLogIterator, error) {

	logs, sub, err := _MetaTxForwarder.contract.FilterLogs(opts, "ProxyLog")
	if err != nil {
		return nil, err
	}
	return &MetaTxForwarderProxyLogIterator{contract: _MetaTxForwarder.contract, event: "ProxyLog", logs: logs, sub: sub}, nil
}

// WatchProxyLog is a free log subscription operation binding the contract event 0xd03bf592f687a60523d2ddd30db6948ef1c7934c57c181027cd0f2b9049c4adc.
//
// Solidity: event ProxyLog((address,address,uint256,uint256,uint256,bytes) arg0)
func (_MetaTxForwarder *MetaTxForwarderFilterer) WatchProxyLog(opts *bind.WatchOpts, sink chan<- *MetaTxForwarderProxyLog) (event.Subscription, error) {

	logs, sub, err := _MetaTxForwarder.contract.WatchLogs(opts, "ProxyLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaTxForwarderProxyLog)
				if err := _MetaTxForwarder.contract.UnpackLog(event, "ProxyLog", log); err != nil {
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

// ParseProxyLog is a log parse operation binding the contract event 0xd03bf592f687a60523d2ddd30db6948ef1c7934c57c181027cd0f2b9049c4adc.
//
// Solidity: event ProxyLog((address,address,uint256,uint256,uint256,bytes) arg0)
func (_MetaTxForwarder *MetaTxForwarderFilterer) ParseProxyLog(log types.Log) (*MetaTxForwarderProxyLog, error) {
	event := new(MetaTxForwarderProxyLog)
	if err := _MetaTxForwarder.contract.UnpackLog(event, "ProxyLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaTxForwarderLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the MetaTxForwarder contract.
type MetaTxForwarderLogIterator struct {
	Event *MetaTxForwarderLog // Event containing the contract specifics and raw log

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
func (it *MetaTxForwarderLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaTxForwarderLog)
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
		it.Event = new(MetaTxForwarderLog)
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
func (it *MetaTxForwarderLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaTxForwarderLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaTxForwarderLog represents a Log event raised by the MetaTxForwarder contract.
type MetaTxForwarderLog struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0x2c2ecbc2212ac38c2f9ec89aa5fcef7f532a5db24dbf7cad1f48bc82843b7428.
//
// Solidity: event log(address arg0)
func (_MetaTxForwarder *MetaTxForwarderFilterer) FilterLog(opts *bind.FilterOpts) (*MetaTxForwarderLogIterator, error) {

	logs, sub, err := _MetaTxForwarder.contract.FilterLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return &MetaTxForwarderLogIterator{contract: _MetaTxForwarder.contract, event: "log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0x2c2ecbc2212ac38c2f9ec89aa5fcef7f532a5db24dbf7cad1f48bc82843b7428.
//
// Solidity: event log(address arg0)
func (_MetaTxForwarder *MetaTxForwarderFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *MetaTxForwarderLog) (event.Subscription, error) {

	logs, sub, err := _MetaTxForwarder.contract.WatchLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaTxForwarderLog)
				if err := _MetaTxForwarder.contract.UnpackLog(event, "log", log); err != nil {
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

// ParseLog is a log parse operation binding the contract event 0x2c2ecbc2212ac38c2f9ec89aa5fcef7f532a5db24dbf7cad1f48bc82843b7428.
//
// Solidity: event log(address arg0)
func (_MetaTxForwarder *MetaTxForwarderFilterer) ParseLog(log types.Log) (*MetaTxForwarderLog, error) {
	event := new(MetaTxForwarderLog)
	if err := _MetaTxForwarder.contract.UnpackLog(event, "log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
