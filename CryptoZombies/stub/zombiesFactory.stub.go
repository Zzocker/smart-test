// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stub

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

// ZombieFactoryABI is the input ABI used to generate the binding from.
const ZombieFactoryABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_dna\",\"type\":\"uint256\"}],\"name\":\"NewZombie\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"createRandomZombie\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"zombies\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"dna\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ZombieFactoryFuncSigs maps the 4-byte function signature to its string representation.
var ZombieFactoryFuncSigs = map[string]string{
	"7bff0a01": "createRandomZombie(string)",
	"2052465e": "zombies(uint256)",
}

// ZombieFactoryBin is the compiled bytecode used for deploying new contracts.
var ZombieFactoryBin = "0x608060405234801561001057600080fd5b506104bb806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80632052465e1461003b5780637bff0a01146100d7575b600080fd5b6100586004803603602081101561005157600080fd5b503561017f565b6040518080602001838152602001828103825284818151815260200191508051906020019080838360005b8381101561009b578181015183820152602001610083565b50505050905090810190601f1680156100c85780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b61017d600480360360208110156100ed57600080fd5b81019060208101813564010000000081111561010857600080fd5b82018360208201111561011a57600080fd5b8035906020019184600183028401116401000000008311171561013c57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610232945050505050565b005b6000818154811061018c57fe5b60009182526020918290206002918202018054604080516001831615610100026000190190921693909304601f8101859004850282018501909352828152909350918391908301828280156102225780601f106101f757610100808354040283529160200191610222565b820191906000526020600020905b81548152906001019060200180831161020557829003601f168201915b5050505050908060010154905082565b600061023d8261024d565b905061024982826102d8565b5050565b600080826040516020018082805190602001908083835b602083106102835780518252601f199092019160209182019101610264565b51815160209384036101000a60001901801990921691161790526040805192909401828103601f190183529093528051920191909120935067016345785d89ffff92508391506102d09050565b069392505050565b600080546040805180820190915284815260208082018590526001830184559280528051805192939192600285027f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56301926103379284929101906103f2565b506020820151816001015550507f88f026aacbbecc90c18411df4b1185fd8d9be2470f1962f192bf84a27d0704b78184846040518084815260200180602001838152602001828103825284818151815260200191508051906020019080838360005b838110156103b1578181015183820152602001610399565b50505050905090810190601f1680156103de5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a1505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061043357805160ff1916838001178555610460565b82800160010185558215610460579182015b82811115610460578251825591602001919060010190610445565b5061046c929150610470565b5090565b5b8082111561046c576000815560010161047156fea26469706673582212209c5160e31dff18b456e500c65d9882bbf2f6c6a47b66d1eec9e47c27a776e50364736f6c63430007030033"

// DeployZombieFactory deploys a new Ethereum contract, binding an instance of ZombieFactory to it.
func DeployZombieFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZombieFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(ZombieFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZombieFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZombieFactory{ZombieFactoryCaller: ZombieFactoryCaller{contract: contract}, ZombieFactoryTransactor: ZombieFactoryTransactor{contract: contract}, ZombieFactoryFilterer: ZombieFactoryFilterer{contract: contract}}, nil
}

// ZombieFactory is an auto generated Go binding around an Ethereum contract.
type ZombieFactory struct {
	ZombieFactoryCaller     // Read-only binding to the contract
	ZombieFactoryTransactor // Write-only binding to the contract
	ZombieFactoryFilterer   // Log filterer for contract events
}

// ZombieFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZombieFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZombieFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZombieFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZombieFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZombieFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZombieFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZombieFactorySession struct {
	Contract     *ZombieFactory    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZombieFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZombieFactoryCallerSession struct {
	Contract *ZombieFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ZombieFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZombieFactoryTransactorSession struct {
	Contract     *ZombieFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ZombieFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZombieFactoryRaw struct {
	Contract *ZombieFactory // Generic contract binding to access the raw methods on
}

// ZombieFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZombieFactoryCallerRaw struct {
	Contract *ZombieFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ZombieFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZombieFactoryTransactorRaw struct {
	Contract *ZombieFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZombieFactory creates a new instance of ZombieFactory, bound to a specific deployed contract.
func NewZombieFactory(address common.Address, backend bind.ContractBackend) (*ZombieFactory, error) {
	contract, err := bindZombieFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZombieFactory{ZombieFactoryCaller: ZombieFactoryCaller{contract: contract}, ZombieFactoryTransactor: ZombieFactoryTransactor{contract: contract}, ZombieFactoryFilterer: ZombieFactoryFilterer{contract: contract}}, nil
}

// NewZombieFactoryCaller creates a new read-only instance of ZombieFactory, bound to a specific deployed contract.
func NewZombieFactoryCaller(address common.Address, caller bind.ContractCaller) (*ZombieFactoryCaller, error) {
	contract, err := bindZombieFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZombieFactoryCaller{contract: contract}, nil
}

// NewZombieFactoryTransactor creates a new write-only instance of ZombieFactory, bound to a specific deployed contract.
func NewZombieFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ZombieFactoryTransactor, error) {
	contract, err := bindZombieFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZombieFactoryTransactor{contract: contract}, nil
}

// NewZombieFactoryFilterer creates a new log filterer instance of ZombieFactory, bound to a specific deployed contract.
func NewZombieFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ZombieFactoryFilterer, error) {
	contract, err := bindZombieFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZombieFactoryFilterer{contract: contract}, nil
}

// bindZombieFactory binds a generic wrapper to an already deployed contract.
func bindZombieFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZombieFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZombieFactory *ZombieFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZombieFactory.Contract.ZombieFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZombieFactory *ZombieFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZombieFactory.Contract.ZombieFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZombieFactory *ZombieFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZombieFactory.Contract.ZombieFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZombieFactory *ZombieFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZombieFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZombieFactory *ZombieFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZombieFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZombieFactory *ZombieFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZombieFactory.Contract.contract.Transact(opts, method, params...)
}

// Zombies is a free data retrieval call binding the contract method 0x2052465e.
//
// Solidity: function zombies(uint256 ) view returns(string name, uint256 dna)
func (_ZombieFactory *ZombieFactoryCaller) Zombies(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name string
	Dna  *big.Int
}, error) {
	var out []interface{}
	err := _ZombieFactory.contract.Call(opts, &out, "zombies", arg0)

	outstruct := new(struct {
		Name string
		Dna  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Dna = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Zombies is a free data retrieval call binding the contract method 0x2052465e.
//
// Solidity: function zombies(uint256 ) view returns(string name, uint256 dna)
func (_ZombieFactory *ZombieFactorySession) Zombies(arg0 *big.Int) (struct {
	Name string
	Dna  *big.Int
}, error) {
	return _ZombieFactory.Contract.Zombies(&_ZombieFactory.CallOpts, arg0)
}

// Zombies is a free data retrieval call binding the contract method 0x2052465e.
//
// Solidity: function zombies(uint256 ) view returns(string name, uint256 dna)
func (_ZombieFactory *ZombieFactoryCallerSession) Zombies(arg0 *big.Int) (struct {
	Name string
	Dna  *big.Int
}, error) {
	return _ZombieFactory.Contract.Zombies(&_ZombieFactory.CallOpts, arg0)
}

// CreateRandomZombie is a paid mutator transaction binding the contract method 0x7bff0a01.
//
// Solidity: function createRandomZombie(string _name) returns()
func (_ZombieFactory *ZombieFactoryTransactor) CreateRandomZombie(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _ZombieFactory.contract.Transact(opts, "createRandomZombie", _name)
}

// CreateRandomZombie is a paid mutator transaction binding the contract method 0x7bff0a01.
//
// Solidity: function createRandomZombie(string _name) returns()
func (_ZombieFactory *ZombieFactorySession) CreateRandomZombie(_name string) (*types.Transaction, error) {
	return _ZombieFactory.Contract.CreateRandomZombie(&_ZombieFactory.TransactOpts, _name)
}

// CreateRandomZombie is a paid mutator transaction binding the contract method 0x7bff0a01.
//
// Solidity: function createRandomZombie(string _name) returns()
func (_ZombieFactory *ZombieFactoryTransactorSession) CreateRandomZombie(_name string) (*types.Transaction, error) {
	return _ZombieFactory.Contract.CreateRandomZombie(&_ZombieFactory.TransactOpts, _name)
}

// ZombieFactoryNewZombieIterator is returned from FilterNewZombie and is used to iterate over the raw logs and unpacked data for NewZombie events raised by the ZombieFactory contract.
type ZombieFactoryNewZombieIterator struct {
	Event *ZombieFactoryNewZombie // Event containing the contract specifics and raw log

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
func (it *ZombieFactoryNewZombieIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZombieFactoryNewZombie)
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
		it.Event = new(ZombieFactoryNewZombie)
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
func (it *ZombieFactoryNewZombieIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZombieFactoryNewZombieIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZombieFactoryNewZombie represents a NewZombie event raised by the ZombieFactory contract.
type ZombieFactoryNewZombie struct {
	Id   *big.Int
	Name string
	Dna  *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNewZombie is a free log retrieval operation binding the contract event 0x88f026aacbbecc90c18411df4b1185fd8d9be2470f1962f192bf84a27d0704b7.
//
// Solidity: event NewZombie(uint256 _id, string _name, uint256 _dna)
func (_ZombieFactory *ZombieFactoryFilterer) FilterNewZombie(opts *bind.FilterOpts) (*ZombieFactoryNewZombieIterator, error) {

	logs, sub, err := _ZombieFactory.contract.FilterLogs(opts, "NewZombie")
	if err != nil {
		return nil, err
	}
	return &ZombieFactoryNewZombieIterator{contract: _ZombieFactory.contract, event: "NewZombie", logs: logs, sub: sub}, nil
}

// WatchNewZombie is a free log subscription operation binding the contract event 0x88f026aacbbecc90c18411df4b1185fd8d9be2470f1962f192bf84a27d0704b7.
//
// Solidity: event NewZombie(uint256 _id, string _name, uint256 _dna)
func (_ZombieFactory *ZombieFactoryFilterer) WatchNewZombie(opts *bind.WatchOpts, sink chan<- *ZombieFactoryNewZombie) (event.Subscription, error) {

	logs, sub, err := _ZombieFactory.contract.WatchLogs(opts, "NewZombie")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZombieFactoryNewZombie)
				if err := _ZombieFactory.contract.UnpackLog(event, "NewZombie", log); err != nil {
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

// ParseNewZombie is a log parse operation binding the contract event 0x88f026aacbbecc90c18411df4b1185fd8d9be2470f1962f192bf84a27d0704b7.
//
// Solidity: event NewZombie(uint256 _id, string _name, uint256 _dna)
func (_ZombieFactory *ZombieFactoryFilterer) ParseNewZombie(log types.Log) (*ZombieFactoryNewZombie, error) {
	event := new(ZombieFactoryNewZombie)
	if err := _ZombieFactory.contract.UnpackLog(event, "NewZombie", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
