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
var BallotBin = "0x60806040523480156200001157600080fd5b5060405162000b2538038062000b258339810160408190526200003491620001a6565b600280546001600160a01b031916339081179091556000908152602081905260408120600190555b8151811015620000f857600160405180604001604052808484815181106200009457634e487b7160e01b600052603260045260246000fd5b6020908102919091018101518252600091810182905283546001810185559382529081902082518051939460020290910192620000d5928492019062000100565b506020820151816001015550508080620000ef9062000339565b9150506200005c565b505062000377565b8280546200010e90620002fc565b90600052602060002090601f0160209004810192826200013257600085556200017d565b82601f106200014d57805160ff19168380011785556200017d565b828001600101855582156200017d579182015b828111156200017d57825182559160200191906001019062000160565b506200018b9291506200018f565b5090565b5b808211156200018b576000815560010162000190565b60006020808385031215620001b9578182fd5b82516001600160401b0380821115620001d0578384fd5b8185019150601f8681840112620001e5578485fd5b825182811115620001fa57620001fa62000361565b6200020a858260051b01620002c9565b81815285810190858701885b84811015620002b957815188018c603f82011262000232578a8bfd5b898101518881111562000249576200024962000361565b6200025c818901601f19168c01620002c9565b81815260408f8184860101111562000272578d8efd5b8d5b8381101562000291578481018201518382018f01528d0162000274565b83811115620002a2578e8e85850101525b505086525050928801929088019060010162000216565b50909a9950505050505050505050565b604051601f8201601f191681016001600160401b0381118282101715620002f457620002f462000361565b604052919050565b600181811c908216806200031157607f821691505b602082108114156200033357634e487b7160e01b600052602260045260246000fd5b50919050565b60006000198214156200035a57634e487b7160e01b81526011600452602481fd5b5060010190565b634e487b7160e01b600052604160045260246000fd5b61079e80620003876000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063665ceea21161005b578063665ceea2146100c5578063ae6c1912146100d8578063d1dbe7d814610103578063fdedaa8a146101245761007d565b80630121b93f146100825780635c19a95c14610097578063609ff1bd146100aa575b600080fd5b610095610090366004610671565b610196565b005b6100956100a5366004610643565b610220565b6100b2610471565b6040519081526020015b60405180910390f35b6100956100d3366004610643565b610508565b6002546100eb906001600160a01b031681565b6040516001600160a01b0390911681526020016100bc565b610116610111366004610671565b610587565b6040516100bc929190610689565b610167610132366004610643565b600060208190529081526040902080546001820154600290920154909160ff8116916101009091046001600160a01b03169084565b6040516100bc949392919093845291151560208401526001600160a01b03166040830152606082015260800190565b3360009081526020819052604090208054158015906101ba5750600181015460ff16155b6101c357600080fd5b6001818101805460ff19168217905581548154909190849081106101f757634e487b7160e01b600052603260045260246000fd5b9060005260206000209060020201600101600082825461021791906106e4565b90915550505050565b33600090815260208190526040902080546102825760405162461bcd60e51b815260206004820152601c60248201527f796f7520646f6e742068617665206120766f74696e672072696768740000000060448201526064015b60405180910390fd5b600181015460ff16156102d05760405162461bcd60e51b81526020600482015260166024820152751e5bdd481a185d9948185b1c9958591e481d9bdd195960521b6044820152606401610279565b336001600160a01b03831614156103295760405162461bcd60e51b815260206004820152601e60248201527f53656c662d64656c65676174696f6e20697320646973616c6c6f7765642e00006044820152606401610279565b6001600160a01b03828116600090815260208190526040902060010154610100900416156103b9576001600160a01b03918216600090815260208190526040902060010154610100900490911690338214156103b45760405162461bcd60e51b815260206004820152600a6024820152690666f756e64206c6f6f760b41b6044820152606401610279565b610329565b600181810180546001600160a81b0319166101006001600160a01b03861690810291909117831790915560009081526020819052604090209081015460ff1615610452578160000154600182600201548154811061042757634e487b7160e01b600052603260045260246000fd5b9060005260206000209060020201600101600082825461044791906106e4565b9091555061046c9050565b8154815482906000906104669084906106e4565b90915550505b505050565b600080805b60015481101561050257600182815481106104a157634e487b7160e01b600052603260045260246000fd5b906000526020600020906002020160010154600182815481106104d457634e487b7160e01b600052603260045260246000fd5b90600052602060002090600202016001015411156104f0578091505b806104fa81610737565b915050610476565b50905090565b6002546001600160a01b0316331461051f57600080fd5b6001600160a01b03811660009081526020819052604090205415801561056157506001600160a01b03811660009081526020819052604090206001015460ff16155b61056a57600080fd5b6001600160a01b0316600090815260208190526040902060019055565b6001818154811061059757600080fd5b90600052602060002090600202016000915090508060000180546105ba906106fc565b80601f01602080910402602001604051908101604052809291908181526020018280546105e6906106fc565b80156106335780601f1061060857610100808354040283529160200191610633565b820191906000526020600020905b81548152906001019060200180831161061657829003601f168201915b5050505050908060010154905082565b600060208284031215610654578081fd5b81356001600160a01b038116811461066a578182fd5b9392505050565b600060208284031215610682578081fd5b5035919050565b6000604082528351806040840152815b818110156106b65760208187018101516060868401015201610699565b818111156106c75782606083860101525b50602083019390935250601f91909101601f191601606001919050565b600082198211156106f7576106f7610752565b500190565b600181811c9082168061071057607f821691505b6020821081141561073157634e487b7160e01b600052602260045260246000fd5b50919050565b600060001982141561074b5761074b610752565b5060010190565b634e487b7160e01b600052601160045260246000fdfea26469706673582212206c9c1b3c5bd9d89326eadebabd36fad980e73f2d30e24a54a16a91ae0820831964736f6c63430008030033"

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
