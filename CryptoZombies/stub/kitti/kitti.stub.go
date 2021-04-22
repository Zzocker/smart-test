// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kitti

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

// KittiABI is the input ABI used to generate the binding from.
const KittiABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_kittiId\",\"type\":\"uint256\"}],\"name\":\"getKitti\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// KittiFuncSigs maps the 4-byte function signature to its string representation.
var KittiFuncSigs = map[string]string{
	"431143cd": "getKitti(uint256)",
}

// KittiBin is the compiled bytecode used for deploying new contracts.
var KittiBin = "0x6080604052601060005567016345785d89ffff60015534801561002157600080fd5b5060c3806100306000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063431143cd14602d575b600080fd5b604760048036036020811015604157600080fd5b50356059565b60408051918252519081900360200190f35b60015460408051602080820185905282518083038201815291830190925280519101206000919081608657fe5b069291505056fea264697066735822122078e91ddbce61537be92dff0ffb6522de0a17a5c5d44d8e13d0a26d7cb0c913c064736f6c63430007030033"

// DeployKitti deploys a new Ethereum contract, binding an instance of Kitti to it.
func DeployKitti(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Kitti, error) {
	parsed, err := abi.JSON(strings.NewReader(KittiABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KittiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Kitti{KittiCaller: KittiCaller{contract: contract}, KittiTransactor: KittiTransactor{contract: contract}, KittiFilterer: KittiFilterer{contract: contract}}, nil
}

// Kitti is an auto generated Go binding around an Ethereum contract.
type Kitti struct {
	KittiCaller     // Read-only binding to the contract
	KittiTransactor // Write-only binding to the contract
	KittiFilterer   // Log filterer for contract events
}

// KittiCaller is an auto generated read-only Go binding around an Ethereum contract.
type KittiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KittiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KittiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KittiSession struct {
	Contract     *Kitti            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KittiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KittiCallerSession struct {
	Contract *KittiCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// KittiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KittiTransactorSession struct {
	Contract     *KittiTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KittiRaw is an auto generated low-level Go binding around an Ethereum contract.
type KittiRaw struct {
	Contract *Kitti // Generic contract binding to access the raw methods on
}

// KittiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KittiCallerRaw struct {
	Contract *KittiCaller // Generic read-only contract binding to access the raw methods on
}

// KittiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KittiTransactorRaw struct {
	Contract *KittiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKitti creates a new instance of Kitti, bound to a specific deployed contract.
func NewKitti(address common.Address, backend bind.ContractBackend) (*Kitti, error) {
	contract, err := bindKitti(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Kitti{KittiCaller: KittiCaller{contract: contract}, KittiTransactor: KittiTransactor{contract: contract}, KittiFilterer: KittiFilterer{contract: contract}}, nil
}

// NewKittiCaller creates a new read-only instance of Kitti, bound to a specific deployed contract.
func NewKittiCaller(address common.Address, caller bind.ContractCaller) (*KittiCaller, error) {
	contract, err := bindKitti(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KittiCaller{contract: contract}, nil
}

// NewKittiTransactor creates a new write-only instance of Kitti, bound to a specific deployed contract.
func NewKittiTransactor(address common.Address, transactor bind.ContractTransactor) (*KittiTransactor, error) {
	contract, err := bindKitti(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KittiTransactor{contract: contract}, nil
}

// NewKittiFilterer creates a new log filterer instance of Kitti, bound to a specific deployed contract.
func NewKittiFilterer(address common.Address, filterer bind.ContractFilterer) (*KittiFilterer, error) {
	contract, err := bindKitti(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KittiFilterer{contract: contract}, nil
}

// bindKitti binds a generic wrapper to an already deployed contract.
func bindKitti(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KittiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kitti *KittiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kitti.Contract.KittiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kitti *KittiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kitti.Contract.KittiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kitti *KittiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kitti.Contract.KittiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kitti *KittiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kitti.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kitti *KittiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kitti.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kitti *KittiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kitti.Contract.contract.Transact(opts, method, params...)
}

// GetKitti is a free data retrieval call binding the contract method 0x431143cd.
//
// Solidity: function getKitti(uint256 _kittiId) view returns(uint256)
func (_Kitti *KittiCaller) GetKitti(opts *bind.CallOpts, _kittiId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Kitti.contract.Call(opts, &out, "getKitti", _kittiId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetKitti is a free data retrieval call binding the contract method 0x431143cd.
//
// Solidity: function getKitti(uint256 _kittiId) view returns(uint256)
func (_Kitti *KittiSession) GetKitti(_kittiId *big.Int) (*big.Int, error) {
	return _Kitti.Contract.GetKitti(&_Kitti.CallOpts, _kittiId)
}

// GetKitti is a free data retrieval call binding the contract method 0x431143cd.
//
// Solidity: function getKitti(uint256 _kittiId) view returns(uint256)
func (_Kitti *KittiCallerSession) GetKitti(_kittiId *big.Int) (*big.Int, error) {
	return _Kitti.Contract.GetKitti(&_Kitti.CallOpts, _kittiId)
}
