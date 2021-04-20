// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ballot

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

// BallotABI is the input ABI used to generate the binding from.
const BallotABI = "[{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_proposalsName\",\"type\":\"string[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Chairperson\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Proposals\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"Name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"VoteCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"Voters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"Weight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"Voted\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"Delegate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"Vote\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"givenRightToVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposal\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"winningProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BallotFuncSigs maps the 4-byte function signature to its string representation.
var BallotFuncSigs = map[string]string{
	"ae6c1912": "Chairperson()",
	"d1dbe7d8": "Proposals(uint256)",
	"fdedaa8a": "Voters(address)",
	"5c19a95c": "delegate(address)",
	"665ceea2": "givenRightToVote(address)",
	"0121b93f": "vote(uint256)",
	"609ff1bd": "winningProposal()",
}

// BallotBin is the compiled bytecode used for deploying new contracts.
var BallotBin = "0x608060405234801561001057600080fd5b5060405161097838038061097883398101604081905261002f9161016e565b600280546001600160a01b0319163317908190556001600160a01b03166000908152602081905260408120600190555b81518110156100d4576001604051806040016040528084848151811061008157fe5b60209081029190910181015182526000918101829052835460018101855593825290819020825180519394600202909101926100c092849201906100db565b50602091909101516001918201550161005f565b5050610293565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061011c57805160ff1916838001178555610149565b82800160010185558215610149579182015b8281111561014957825182559160200191906001019061012e565b50610155929150610159565b5090565b5b80821115610155576000815560010161015a565b60006020808385031215610180578182fd5b82516001600160401b0380821115610196578384fd5b8185019150601f86818401126101aa578485fd5b8251828111156101b657fe5b6101c38586830201610270565b81815285810190858701885b8481101561026057815188018c603f8201126101e9578a8bfd5b89810151888111156101f757fe5b610208818901601f19168c01610270565b81815260408f8184860101111561021d578d8efd5b8d5b8381101561023a578481018201518382018f01528d0161021f565b8381111561024a578e8e85850101525b50508652505092880192908801906001016101cf565b50909a9950505050505050505050565b6040518181016001600160401b038111828210171561028b57fe5b604052919050565b6106d6806102a26000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063665ceea21161005b578063665ceea2146100c8578063ae6c1912146100db578063d1dbe7d8146100f0578063fdedaa8a146101115761007d565b80630121b93f146100825780635c19a95c14610097578063609ff1bd146100aa575b600080fd5b61009561009036600461052a565b610134565b005b6100956100a53660046104fc565b61019f565b6100b2610321565b6040516100bf9190610673565b60405180910390f35b6100956100d63660046104fc565b610386565b6100e3610405565b6040516100bf9190610542565b6101036100fe36600461052a565b610414565b6040516100bf929190610556565b61012461011f3660046104fc565b6104c7565b6040516100bf949392919061067c565b3360009081526020819052604090208054158015906101585750600181015460ff16155b61016157600080fd5b6001818101805460ff191682179055815481549091908490811061018157fe5b60009182526020909120600160029092020101805490910190555050565b33600090815260208190526040902080546101d55760405162461bcd60e51b81526004016101cc90610605565b60405180910390fd5b600181015460ff16156101fa5760405162461bcd60e51b81526004016101cc906105b1565b336001600160a01b03831614156102235760405162461bcd60e51b81526004016101cc9061063c565b6001600160a01b0382811660009081526020819052604090206001015461010090041615610296576001600160a01b03918216600090815260208190526040902060010154610100900490911690338214156102915760405162461bcd60e51b81526004016101cc906105e1565b610223565b6001818101805460ff19168217610100600160a81b0319166101006001600160a01b0386169081029190911790915560009081526020819052604090209081015460ff161561031457816000015460018260020154815481106102f557fe5b600091825260209091206001600290920201018054909101905561031c565b815481540181555b505050565b600080805b600154811015610380576001828154811061033d57fe5b9060005260206000209060020201600101546001828154811061035c57fe5b9060005260206000209060020201600101541115610378578091505b600101610326565b50905090565b6002546001600160a01b0316331461039d57600080fd5b6001600160a01b0381166000908152602081905260409020541580156103df57506001600160a01b03811660009081526020819052604090206001015460ff16155b6103e857600080fd5b6001600160a01b0316600090815260208190526040902060019055565b6002546001600160a01b031681565b6001818154811061042157fe5b60009182526020918290206002918202018054604080516001831615610100026000190190921693909304601f8101859004850282018501909352828152909350918391908301828280156104b75780601f1061048c576101008083540402835291602001916104b7565b820191906000526020600020905b81548152906001019060200180831161049a57829003601f168201915b5050505050908060010154905082565b600060208190529081526040902080546001820154600290920154909160ff8116916101009091046001600160a01b03169084565b60006020828403121561050d578081fd5b81356001600160a01b0381168114610523578182fd5b9392505050565b60006020828403121561053b578081fd5b5035919050565b6001600160a01b0391909116815260200190565b6000604082528351806040840152815b818110156105835760208187018101516060868401015201610566565b818111156105945782606083860101525b50602083019390935250601f91909101601f191601606001919050565b6020808252601690820152751e5bdd481a185d9948185b1c9958591e481d9bdd195960521b604082015260600190565b6020808252600a90820152690666f756e64206c6f6f760b41b604082015260600190565b6020808252601c908201527f796f7520646f6e742068617665206120766f74696e6720726967687400000000604082015260600190565b6020808252601e908201527f53656c662d64656c65676174696f6e20697320646973616c6c6f7765642e0000604082015260600190565b90815260200190565b93845291151560208401526001600160a01b0316604083015260608201526080019056fea2646970667358221220143c9f3b504b720024830121fa50fc17bf04f49f1b3ee5c8a379c8553165290364736f6c63430007030033"

// DeployBallot deploys a new Ethereum contract, binding an instance of Ballot to it.
func DeployBallot(auth *bind.TransactOpts, backend bind.ContractBackend, _proposalsName []string) (common.Address, *types.Transaction, *Ballot, error) {
	parsed, err := abi.JSON(strings.NewReader(BallotABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BallotBin), backend, _proposalsName)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ballot{BallotCaller: BallotCaller{contract: contract}, BallotTransactor: BallotTransactor{contract: contract}, BallotFilterer: BallotFilterer{contract: contract}}, nil
}

// Ballot is an auto generated Go binding around an Ethereum contract.
type Ballot struct {
	BallotCaller     // Read-only binding to the contract
	BallotTransactor // Write-only binding to the contract
	BallotFilterer   // Log filterer for contract events
}

// BallotCaller is an auto generated read-only Go binding around an Ethereum contract.
type BallotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BallotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BallotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BallotFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BallotFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BallotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BallotSession struct {
	Contract     *Ballot           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BallotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BallotCallerSession struct {
	Contract *BallotCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BallotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BallotTransactorSession struct {
	Contract     *BallotTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BallotRaw is an auto generated low-level Go binding around an Ethereum contract.
type BallotRaw struct {
	Contract *Ballot // Generic contract binding to access the raw methods on
}

// BallotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BallotCallerRaw struct {
	Contract *BallotCaller // Generic read-only contract binding to access the raw methods on
}

// BallotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BallotTransactorRaw struct {
	Contract *BallotTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBallot creates a new instance of Ballot, bound to a specific deployed contract.
func NewBallot(address common.Address, backend bind.ContractBackend) (*Ballot, error) {
	contract, err := bindBallot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ballot{BallotCaller: BallotCaller{contract: contract}, BallotTransactor: BallotTransactor{contract: contract}, BallotFilterer: BallotFilterer{contract: contract}}, nil
}

// NewBallotCaller creates a new read-only instance of Ballot, bound to a specific deployed contract.
func NewBallotCaller(address common.Address, caller bind.ContractCaller) (*BallotCaller, error) {
	contract, err := bindBallot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BallotCaller{contract: contract}, nil
}

// NewBallotTransactor creates a new write-only instance of Ballot, bound to a specific deployed contract.
func NewBallotTransactor(address common.Address, transactor bind.ContractTransactor) (*BallotTransactor, error) {
	contract, err := bindBallot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BallotTransactor{contract: contract}, nil
}

// NewBallotFilterer creates a new log filterer instance of Ballot, bound to a specific deployed contract.
func NewBallotFilterer(address common.Address, filterer bind.ContractFilterer) (*BallotFilterer, error) {
	contract, err := bindBallot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BallotFilterer{contract: contract}, nil
}

// bindBallot binds a generic wrapper to an already deployed contract.
func bindBallot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BallotABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ballot *BallotRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ballot.Contract.BallotCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ballot *BallotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ballot.Contract.BallotTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ballot *BallotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ballot.Contract.BallotTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ballot *BallotCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ballot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ballot *BallotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ballot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ballot *BallotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ballot.Contract.contract.Transact(opts, method, params...)
}

// Chairperson is a free data retrieval call binding the contract method 0xae6c1912.
//
// Solidity: function Chairperson() view returns(address)
func (_Ballot *BallotCaller) Chairperson(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ballot.contract.Call(opts, &out, "Chairperson")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Chairperson is a free data retrieval call binding the contract method 0xae6c1912.
//
// Solidity: function Chairperson() view returns(address)
func (_Ballot *BallotSession) Chairperson() (common.Address, error) {
	return _Ballot.Contract.Chairperson(&_Ballot.CallOpts)
}

// Chairperson is a free data retrieval call binding the contract method 0xae6c1912.
//
// Solidity: function Chairperson() view returns(address)
func (_Ballot *BallotCallerSession) Chairperson() (common.Address, error) {
	return _Ballot.Contract.Chairperson(&_Ballot.CallOpts)
}

// Proposals is a free data retrieval call binding the contract method 0xd1dbe7d8.
//
// Solidity: function Proposals(uint256 ) view returns(string Name, uint256 VoteCount)
func (_Ballot *BallotCaller) Proposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name      string
	VoteCount *big.Int
}, error) {
	var out []interface{}
	err := _Ballot.contract.Call(opts, &out, "Proposals", arg0)

	outstruct := new(struct {
		Name      string
		VoteCount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.VoteCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0xd1dbe7d8.
//
// Solidity: function Proposals(uint256 ) view returns(string Name, uint256 VoteCount)
func (_Ballot *BallotSession) Proposals(arg0 *big.Int) (struct {
	Name      string
	VoteCount *big.Int
}, error) {
	return _Ballot.Contract.Proposals(&_Ballot.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0xd1dbe7d8.
//
// Solidity: function Proposals(uint256 ) view returns(string Name, uint256 VoteCount)
func (_Ballot *BallotCallerSession) Proposals(arg0 *big.Int) (struct {
	Name      string
	VoteCount *big.Int
}, error) {
	return _Ballot.Contract.Proposals(&_Ballot.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xfdedaa8a.
//
// Solidity: function Voters(address ) view returns(uint256 Weight, bool Voted, address Delegate, uint256 Vote)
func (_Ballot *BallotCaller) Voters(opts *bind.CallOpts, arg0 common.Address) (struct {
	Weight   *big.Int
	Voted    bool
	Delegate common.Address
	Vote     *big.Int
}, error) {
	var out []interface{}
	err := _Ballot.contract.Call(opts, &out, "Voters", arg0)

	outstruct := new(struct {
		Weight   *big.Int
		Voted    bool
		Delegate common.Address
		Vote     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Weight = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Voted = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Delegate = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Vote = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Voters is a free data retrieval call binding the contract method 0xfdedaa8a.
//
// Solidity: function Voters(address ) view returns(uint256 Weight, bool Voted, address Delegate, uint256 Vote)
func (_Ballot *BallotSession) Voters(arg0 common.Address) (struct {
	Weight   *big.Int
	Voted    bool
	Delegate common.Address
	Vote     *big.Int
}, error) {
	return _Ballot.Contract.Voters(&_Ballot.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xfdedaa8a.
//
// Solidity: function Voters(address ) view returns(uint256 Weight, bool Voted, address Delegate, uint256 Vote)
func (_Ballot *BallotCallerSession) Voters(arg0 common.Address) (struct {
	Weight   *big.Int
	Voted    bool
	Delegate common.Address
	Vote     *big.Int
}, error) {
	return _Ballot.Contract.Voters(&_Ballot.CallOpts, arg0)
}

// WinningProposal is a free data retrieval call binding the contract method 0x609ff1bd.
//
// Solidity: function winningProposal() view returns(uint256)
func (_Ballot *BallotCaller) WinningProposal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ballot.contract.Call(opts, &out, "winningProposal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WinningProposal is a free data retrieval call binding the contract method 0x609ff1bd.
//
// Solidity: function winningProposal() view returns(uint256)
func (_Ballot *BallotSession) WinningProposal() (*big.Int, error) {
	return _Ballot.Contract.WinningProposal(&_Ballot.CallOpts)
}

// WinningProposal is a free data retrieval call binding the contract method 0x609ff1bd.
//
// Solidity: function winningProposal() view returns(uint256)
func (_Ballot *BallotCallerSession) WinningProposal() (*big.Int, error) {
	return _Ballot.Contract.WinningProposal(&_Ballot.CallOpts)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address _to) returns()
func (_Ballot *BallotTransactor) Delegate(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Ballot.contract.Transact(opts, "delegate", _to)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address _to) returns()
func (_Ballot *BallotSession) Delegate(_to common.Address) (*types.Transaction, error) {
	return _Ballot.Contract.Delegate(&_Ballot.TransactOpts, _to)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address _to) returns()
func (_Ballot *BallotTransactorSession) Delegate(_to common.Address) (*types.Transaction, error) {
	return _Ballot.Contract.Delegate(&_Ballot.TransactOpts, _to)
}

// GivenRightToVote is a paid mutator transaction binding the contract method 0x665ceea2.
//
// Solidity: function givenRightToVote(address voter) returns()
func (_Ballot *BallotTransactor) GivenRightToVote(opts *bind.TransactOpts, voter common.Address) (*types.Transaction, error) {
	return _Ballot.contract.Transact(opts, "givenRightToVote", voter)
}

// GivenRightToVote is a paid mutator transaction binding the contract method 0x665ceea2.
//
// Solidity: function givenRightToVote(address voter) returns()
func (_Ballot *BallotSession) GivenRightToVote(voter common.Address) (*types.Transaction, error) {
	return _Ballot.Contract.GivenRightToVote(&_Ballot.TransactOpts, voter)
}

// GivenRightToVote is a paid mutator transaction binding the contract method 0x665ceea2.
//
// Solidity: function givenRightToVote(address voter) returns()
func (_Ballot *BallotTransactorSession) GivenRightToVote(voter common.Address) (*types.Transaction, error) {
	return _Ballot.Contract.GivenRightToVote(&_Ballot.TransactOpts, voter)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 _proposal) returns()
func (_Ballot *BallotTransactor) Vote(opts *bind.TransactOpts, _proposal *big.Int) (*types.Transaction, error) {
	return _Ballot.contract.Transact(opts, "vote", _proposal)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 _proposal) returns()
func (_Ballot *BallotSession) Vote(_proposal *big.Int) (*types.Transaction, error) {
	return _Ballot.Contract.Vote(&_Ballot.TransactOpts, _proposal)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 _proposal) returns()
func (_Ballot *BallotTransactorSession) Vote(_proposal *big.Int) (*types.Transaction, error) {
	return _Ballot.Contract.Vote(&_Ballot.TransactOpts, _proposal)
}
