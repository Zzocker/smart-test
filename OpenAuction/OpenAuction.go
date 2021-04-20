// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package openauction

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// OpenAuctionABI is the input ABI used to generate the binding from.
const OpenAuctionABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_beneficiary\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"HighestBidIncresed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AuctionEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Beneficiary\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Bid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HighestBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HighestBidder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OpenAuctionFuncSigs maps the 4-byte function signature to its string representation.
var OpenAuctionFuncSigs = map[string]string{
	"f3c88510": "AuctionEndTime()",
	"73d55379": "Beneficiary()",
	"6e6452cb": "Bid()",
	"167fb50e": "HighestBid()",
	"996657af": "HighestBidder()",
	"3ccfd60b": "withdraw()",
}

// OpenAuctionBin is the compiled bytecode used for deploying new contracts.
var OpenAuctionBin = "0x608060405234801561001057600080fd5b506040516103793803806103798339818101604052604081101561003357600080fd5b5080516020909101516000918255600180546001600160a01b0319166001600160a01b0390921691909117905561030990819061007090396000f3fe6080604052600436106100555760003560e01c8063167fb50e1461005a5780633ccfd60b146100815780636e6452cb146100aa57806373d55379146100b4578063996657af146100e5578063f3c88510146100fa575b600080fd5b34801561006657600080fd5b5061006f61010f565b60408051918252519081900360200190f35b34801561008d57600080fd5b50610096610115565b604080519115158252519081900360200190f35b6100b261017f565b005b3480156100c057600080fd5b506100c96102af565b604080516001600160a01b039092168252519081900360200190f35b3480156100f157600080fd5b506100c96102be565b34801561010657600080fd5b5061006f6102cd565b60025481565b33600090815260046020526040812054801561017657336000818152600460205260408082208290555183156108fc0291849190818181858888f115159350610176925050505733600090815260046020526040812091909155905061017c565b60019150505b90565b6000544211156101ca576040805162461bcd60e51b8152602060048201526011602482015270105d58dd1a5bdb881a185cc8195b991959607a1b604482015290519081900360640190fd5b6002543411610220576040805162461bcd60e51b815260206004820152601c60248201527f4869676865722062696420697320616c726561647920706c6163656400000000604482015290519081900360640190fd5b6003546001600160a01b031615610256576002546003546001600160a01b03166000908152600460205260409020805490910190555b346002819055600380546001600160a01b0319163390811790915560408051918252602082019290925281517fcffd2f397289f80763255196661e1934d26e4d7f4b8d00e32c66030c3080a9aa929181900390910190a1565b6001546001600160a01b031681565b6003546001600160a01b031681565b6000548156fea2646970667358221220a532bee59a584cdf7f9254b59ec849225d338ceffe7131d3ef94aac6a3211de664736f6c63430007030033"

// DeployOpenAuction deploys a new Ethereum contract, binding an instance of OpenAuction to it.
func DeployOpenAuction(auth *bind.TransactOpts, backend bind.ContractBackend, _endTime *big.Int, _beneficiary common.Address) (common.Address, *types.Transaction, *OpenAuction, error) {
	parsed, err := abi.JSON(strings.NewReader(OpenAuctionABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OpenAuctionBin), backend, _endTime, _beneficiary)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OpenAuction{OpenAuctionCaller: OpenAuctionCaller{contract: contract}, OpenAuctionTransactor: OpenAuctionTransactor{contract: contract}, OpenAuctionFilterer: OpenAuctionFilterer{contract: contract}}, nil
}

// OpenAuction is an auto generated Go binding around an Ethereum contract.
type OpenAuction struct {
	OpenAuctionCaller     // Read-only binding to the contract
	OpenAuctionTransactor // Write-only binding to the contract
	OpenAuctionFilterer   // Log filterer for contract events
}

// OpenAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type OpenAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpenAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OpenAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpenAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OpenAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpenAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OpenAuctionSession struct {
	Contract     *OpenAuction      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OpenAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OpenAuctionCallerSession struct {
	Contract *OpenAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// OpenAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OpenAuctionTransactorSession struct {
	Contract     *OpenAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// OpenAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type OpenAuctionRaw struct {
	Contract *OpenAuction // Generic contract binding to access the raw methods on
}

// OpenAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OpenAuctionCallerRaw struct {
	Contract *OpenAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// OpenAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OpenAuctionTransactorRaw struct {
	Contract *OpenAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOpenAuction creates a new instance of OpenAuction, bound to a specific deployed contract.
func NewOpenAuction(address common.Address, backend bind.ContractBackend) (*OpenAuction, error) {
	contract, err := bindOpenAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OpenAuction{OpenAuctionCaller: OpenAuctionCaller{contract: contract}, OpenAuctionTransactor: OpenAuctionTransactor{contract: contract}, OpenAuctionFilterer: OpenAuctionFilterer{contract: contract}}, nil
}

// NewOpenAuctionCaller creates a new read-only instance of OpenAuction, bound to a specific deployed contract.
func NewOpenAuctionCaller(address common.Address, caller bind.ContractCaller) (*OpenAuctionCaller, error) {
	contract, err := bindOpenAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OpenAuctionCaller{contract: contract}, nil
}

// NewOpenAuctionTransactor creates a new write-only instance of OpenAuction, bound to a specific deployed contract.
func NewOpenAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*OpenAuctionTransactor, error) {
	contract, err := bindOpenAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OpenAuctionTransactor{contract: contract}, nil
}

// NewOpenAuctionFilterer creates a new log filterer instance of OpenAuction, bound to a specific deployed contract.
func NewOpenAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*OpenAuctionFilterer, error) {
	contract, err := bindOpenAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OpenAuctionFilterer{contract: contract}, nil
}

// bindOpenAuction binds a generic wrapper to an already deployed contract.
func bindOpenAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OpenAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OpenAuction *OpenAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OpenAuction.Contract.OpenAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OpenAuction *OpenAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpenAuction.Contract.OpenAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OpenAuction *OpenAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OpenAuction.Contract.OpenAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OpenAuction *OpenAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OpenAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OpenAuction *OpenAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpenAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OpenAuction *OpenAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OpenAuction.Contract.contract.Transact(opts, method, params...)
}

// AuctionEndTime is a free data retrieval call binding the contract method 0xf3c88510.
//
// Solidity: function AuctionEndTime() view returns(uint256)
func (_OpenAuction *OpenAuctionCaller) AuctionEndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OpenAuction.contract.Call(opts, &out, "AuctionEndTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AuctionEndTime is a free data retrieval call binding the contract method 0xf3c88510.
//
// Solidity: function AuctionEndTime() view returns(uint256)
func (_OpenAuction *OpenAuctionSession) AuctionEndTime() (*big.Int, error) {
	return _OpenAuction.Contract.AuctionEndTime(&_OpenAuction.CallOpts)
}

// AuctionEndTime is a free data retrieval call binding the contract method 0xf3c88510.
//
// Solidity: function AuctionEndTime() view returns(uint256)
func (_OpenAuction *OpenAuctionCallerSession) AuctionEndTime() (*big.Int, error) {
	return _OpenAuction.Contract.AuctionEndTime(&_OpenAuction.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x73d55379.
//
// Solidity: function Beneficiary() view returns(address)
func (_OpenAuction *OpenAuctionCaller) Beneficiary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OpenAuction.contract.Call(opts, &out, "Beneficiary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Beneficiary is a free data retrieval call binding the contract method 0x73d55379.
//
// Solidity: function Beneficiary() view returns(address)
func (_OpenAuction *OpenAuctionSession) Beneficiary() (common.Address, error) {
	return _OpenAuction.Contract.Beneficiary(&_OpenAuction.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x73d55379.
//
// Solidity: function Beneficiary() view returns(address)
func (_OpenAuction *OpenAuctionCallerSession) Beneficiary() (common.Address, error) {
	return _OpenAuction.Contract.Beneficiary(&_OpenAuction.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0x167fb50e.
//
// Solidity: function HighestBid() view returns(uint256)
func (_OpenAuction *OpenAuctionCaller) HighestBid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OpenAuction.contract.Call(opts, &out, "HighestBid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HighestBid is a free data retrieval call binding the contract method 0x167fb50e.
//
// Solidity: function HighestBid() view returns(uint256)
func (_OpenAuction *OpenAuctionSession) HighestBid() (*big.Int, error) {
	return _OpenAuction.Contract.HighestBid(&_OpenAuction.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0x167fb50e.
//
// Solidity: function HighestBid() view returns(uint256)
func (_OpenAuction *OpenAuctionCallerSession) HighestBid() (*big.Int, error) {
	return _OpenAuction.Contract.HighestBid(&_OpenAuction.CallOpts)
}

// HighestBidder is a free data retrieval call binding the contract method 0x996657af.
//
// Solidity: function HighestBidder() view returns(address)
func (_OpenAuction *OpenAuctionCaller) HighestBidder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OpenAuction.contract.Call(opts, &out, "HighestBidder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HighestBidder is a free data retrieval call binding the contract method 0x996657af.
//
// Solidity: function HighestBidder() view returns(address)
func (_OpenAuction *OpenAuctionSession) HighestBidder() (common.Address, error) {
	return _OpenAuction.Contract.HighestBidder(&_OpenAuction.CallOpts)
}

// HighestBidder is a free data retrieval call binding the contract method 0x996657af.
//
// Solidity: function HighestBidder() view returns(address)
func (_OpenAuction *OpenAuctionCallerSession) HighestBidder() (common.Address, error) {
	return _OpenAuction.Contract.HighestBidder(&_OpenAuction.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() payable returns()
func (_OpenAuction *OpenAuctionTransactor) Bid(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpenAuction.contract.Transact(opts, "Bid")
}

// Bid is a paid mutator transaction binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() payable returns()
func (_OpenAuction *OpenAuctionSession) Bid() (*types.Transaction, error) {
	return _OpenAuction.Contract.Bid(&_OpenAuction.TransactOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() payable returns()
func (_OpenAuction *OpenAuctionTransactorSession) Bid() (*types.Transaction, error) {
	return _OpenAuction.Contract.Bid(&_OpenAuction.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_OpenAuction *OpenAuctionTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpenAuction.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_OpenAuction *OpenAuctionSession) Withdraw() (*types.Transaction, error) {
	return _OpenAuction.Contract.Withdraw(&_OpenAuction.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_OpenAuction *OpenAuctionTransactorSession) Withdraw() (*types.Transaction, error) {
	return _OpenAuction.Contract.Withdraw(&_OpenAuction.TransactOpts)
}

// OpenAuctionHighestBidIncresedIterator is returned from FilterHighestBidIncresed and is used to iterate over the raw logs and unpacked data for HighestBidIncresed events raised by the OpenAuction contract.
type OpenAuctionHighestBidIncresedIterator struct {
	Event *OpenAuctionHighestBidIncresed // Event containing the contract specifics and raw log

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
func (it *OpenAuctionHighestBidIncresedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpenAuctionHighestBidIncresed)
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
		it.Event = new(OpenAuctionHighestBidIncresed)
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
func (it *OpenAuctionHighestBidIncresedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpenAuctionHighestBidIncresedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpenAuctionHighestBidIncresed represents a HighestBidIncresed event raised by the OpenAuction contract.
type OpenAuctionHighestBidIncresed struct {
	Bidder common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterHighestBidIncresed is a free log retrieval operation binding the contract event 0xcffd2f397289f80763255196661e1934d26e4d7f4b8d00e32c66030c3080a9aa.
//
// Solidity: event HighestBidIncresed(address bidder, uint256 amount)
func (_OpenAuction *OpenAuctionFilterer) FilterHighestBidIncresed(opts *bind.FilterOpts) (*OpenAuctionHighestBidIncresedIterator, error) {

	logs, sub, err := _OpenAuction.contract.FilterLogs(opts, "HighestBidIncresed")
	if err != nil {
		return nil, err
	}
	return &OpenAuctionHighestBidIncresedIterator{contract: _OpenAuction.contract, event: "HighestBidIncresed", logs: logs, sub: sub}, nil
}

// WatchHighestBidIncresed is a free log subscription operation binding the contract event 0xcffd2f397289f80763255196661e1934d26e4d7f4b8d00e32c66030c3080a9aa.
//
// Solidity: event HighestBidIncresed(address bidder, uint256 amount)
func (_OpenAuction *OpenAuctionFilterer) WatchHighestBidIncresed(opts *bind.WatchOpts, sink chan<- *OpenAuctionHighestBidIncresed) (event.Subscription, error) {

	logs, sub, err := _OpenAuction.contract.WatchLogs(opts, "HighestBidIncresed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpenAuctionHighestBidIncresed)
				if err := _OpenAuction.contract.UnpackLog(event, "HighestBidIncresed", log); err != nil {
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

// ParseHighestBidIncresed is a log parse operation binding the contract event 0xcffd2f397289f80763255196661e1934d26e4d7f4b8d00e32c66030c3080a9aa.
//
// Solidity: event HighestBidIncresed(address bidder, uint256 amount)
func (_OpenAuction *OpenAuctionFilterer) ParseHighestBidIncresed(log types.Log) (*OpenAuctionHighestBidIncresed, error) {
	event := new(OpenAuctionHighestBidIncresed)
	if err := _OpenAuction.contract.UnpackLog(event, "HighestBidIncresed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
