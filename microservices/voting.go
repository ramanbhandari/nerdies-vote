// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package voting

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

// VotingMetaData contains all meta data concerning the Voting contract.
var VotingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"candidateId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"CandidateAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"candidateId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"voterHash\",\"type\":\"bytes32\"}],\"name\":\"VoteCasted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"voterHash\",\"type\":\"bytes32\"}],\"name\":\"VoterRegistered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"candidates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"candidatesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"hasVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isEligible\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"voterHash\",\"type\":\"bytes32\"}],\"name\":\"registerVoterHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"voterHash\",\"type\":\"bytes32\"}],\"name\":\"checkVoterStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"candidateId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"voterHash\",\"type\":\"bytes32\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"voterHash\",\"type\":\"bytes32\"}],\"name\":\"getVoter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCount\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"addCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5033600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611282806100616000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806368bb8bb61161007157806368bb8bb61461017a578063a87d942c14610196578063b44cf604146101b4578063d8e63542146101e4578063e2a18da014610214578063f851a44014610230576100a9565b80631b4613cb146100ae5780632d35a8a2146100de5780633477ee2e146100fc5780633c74e0d91461012e578063462e91ec1461015e575b600080fd5b6100c860048036038101906100c391906109b2565b61024e565b6040516100d591906109fa565b60405180910390f35b6100e661026e565b6040516100f39190610a2e565b60405180910390f35b61011660048036038101906101119190610a75565b610274565b60405161012593929190610b3b565b60405180910390f35b610148600480360381019061014391906109b2565b610326565b60405161015591906109fa565b60405180910390f35b61017860048036038101906101739190610cae565b610346565b005b610194600480360381019061018f9190610cf7565b610492565b005b61019e610638565b6040516101ab9190610df5565b60405180910390f35b6101ce60048036038101906101c991906109b2565b6106f5565b6040516101db91906109fa565b60405180910390f35b6101fe60048036038101906101f991906109b2565b61071f565b60405161020b91906109fa565b60405180910390f35b61022e600480360381019061022991906109b2565b610749565b005b61023861089f565b6040516102459190610e58565b60405180910390f35b60026020528060005260406000206000915054906101000a900460ff1681565b60035481565b600060205280600052604060002060009150905080600001549080600101805461029d90610ea2565b80601f01602080910402602001604051908101604052809291908181526020018280546102c990610ea2565b80156103165780601f106102eb57610100808354040283529160200191610316565b820191906000526020600020905b8154815290600101906020018083116102f957829003601f168201915b5050505050908060020154905083565b60016020528060005260406000206000915054906101000a900460ff1681565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103d6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103cd90610f1f565b60405180910390fd5b600360008154809291906103e990610f6e565b919050555060405180606001604052806003548152602001828152602001600081525060008060035481526020019081526020016000206000820151816000015560208201518160010190805190602001906104469291906108c5565b50604082015181600201559050507fe83b2a43e7e82d975c8a0a6d2f045153c869e111136a34d1889ab7b598e396a360035482604051610487929190610fb6565b60405180910390a150565b6001600082815260200190815260200160002060009054906101000a900460ff166104f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104e990611032565b60405180910390fd5b6002600082815260200190815260200160002060009054906101000a900460ff1615610553576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161054a9061109e565b60405180910390fd5b60008211801561056557506003548211155b6105a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161059b9061110a565b60405180910390fd5b60008083815260200190815260200160002060020160008154809291906105ca90610f6e565b919050555060016002600083815260200190815260200160002060006101000a81548160ff0219169083151502179055507fc8058fd305bbbcba2a5a8df5471ff25dd0f41cc61f4e2a9f0ed519b44f848847828260405161062c929190611139565b60405180910390a15050565b6060600060035467ffffffffffffffff81111561065857610657610b83565b5b6040519080825280602002602001820160405280156106865781602001602082028036833780820191505090505b5090506000600190505b60035481116106ed5760008082815260200190815260200160002060020154826001836106bd9190611162565b815181106106ce576106cd611196565b5b60200260200101818152505080806106e590610f6e565b915050610690565b508091505090565b60006002600083815260200190815260200160002060009054906101000a900460ff169050919050565b60006001600083815260200190815260200160002060009054906101000a900460ff169050919050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107d9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107d090610f1f565b60405180910390fd5b6001600082815260200190815260200160002060009054906101000a900460ff161561083a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161083190611211565b60405180910390fd5b600180600083815260200190815260200160002060006101000a81548160ff0219169083151502179055507fe1e93aab4d6cae623ae197f9c4d63ac5eede4c69f9d0fbc92e4dbb0766abd115816040516108949190611231565b60405180910390a150565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b8280546108d190610ea2565b90600052602060002090601f0160209004810192826108f3576000855561093a565b82601f1061090c57805160ff191683800117855561093a565b8280016001018555821561093a579182015b8281111561093957825182559160200191906001019061091e565b5b509050610947919061094b565b5090565b5b8082111561096457600081600090555060010161094c565b5090565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b61098f8161097c565b811461099a57600080fd5b50565b6000813590506109ac81610986565b92915050565b6000602082840312156109c8576109c7610972565b5b60006109d68482850161099d565b91505092915050565b60008115159050919050565b6109f4816109df565b82525050565b6000602082019050610a0f60008301846109eb565b92915050565b6000819050919050565b610a2881610a15565b82525050565b6000602082019050610a436000830184610a1f565b92915050565b610a5281610a15565b8114610a5d57600080fd5b50565b600081359050610a6f81610a49565b92915050565b600060208284031215610a8b57610a8a610972565b5b6000610a9984828501610a60565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610adc578082015181840152602081019050610ac1565b83811115610aeb576000848401525b50505050565b6000601f19601f8301169050919050565b6000610b0d82610aa2565b610b178185610aad565b9350610b27818560208601610abe565b610b3081610af1565b840191505092915050565b6000606082019050610b506000830186610a1f565b8181036020830152610b628185610b02565b9050610b716040830184610a1f565b949350505050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610bbb82610af1565b810181811067ffffffffffffffff82111715610bda57610bd9610b83565b5b80604052505050565b6000610bed610968565b9050610bf98282610bb2565b919050565b600067ffffffffffffffff821115610c1957610c18610b83565b5b610c2282610af1565b9050602081019050919050565b82818337600083830152505050565b6000610c51610c4c84610bfe565b610be3565b905082815260208101848484011115610c6d57610c6c610b7e565b5b610c78848285610c2f565b509392505050565b600082601f830112610c9557610c94610b79565b5b8135610ca5848260208601610c3e565b91505092915050565b600060208284031215610cc457610cc3610972565b5b600082013567ffffffffffffffff811115610ce257610ce1610977565b5b610cee84828501610c80565b91505092915050565b60008060408385031215610d0e57610d0d610972565b5b6000610d1c85828601610a60565b9250506020610d2d8582860161099d565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610d6c81610a15565b82525050565b6000610d7e8383610d63565b60208301905092915050565b6000602082019050919050565b6000610da282610d37565b610dac8185610d42565b9350610db783610d53565b8060005b83811015610de8578151610dcf8882610d72565b9750610dda83610d8a565b925050600181019050610dbb565b5085935050505092915050565b60006020820190508181036000830152610e0f8184610d97565b905092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610e4282610e17565b9050919050565b610e5281610e37565b82525050565b6000602082019050610e6d6000830184610e49565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610eba57607f821691505b602082108103610ecd57610ecc610e73565b5b50919050565b7f4e6f7420617574686f72697a6564000000000000000000000000000000000000600082015250565b6000610f09600e83610aad565b9150610f1482610ed3565b602082019050919050565b60006020820190508181036000830152610f3881610efc565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610f7982610a15565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610fab57610faa610f3f565b5b600182019050919050565b6000604082019050610fcb6000830185610a1f565b8181036020830152610fdd8184610b02565b90509392505050565b7f566f746572204e6f742052656769737465726564000000000000000000000000600082015250565b600061101c601483610aad565b915061102782610fe6565b602082019050919050565b6000602082019050818103600083015261104b8161100f565b9050919050565b7f416c726561647920566f74656400000000000000000000000000000000000000600082015250565b6000611088600d83610aad565b915061109382611052565b602082019050919050565b600060208201905081810360008301526110b78161107b565b9050919050565b7f496e76616c69642063616e646964617465204944000000000000000000000000600082015250565b60006110f4601483610aad565b91506110ff826110be565b602082019050919050565b60006020820190508181036000830152611123816110e7565b9050919050565b6111338161097c565b82525050565b600060408201905061114e6000830185610a1f565b61115b602083018461112a565b9392505050565b600061116d82610a15565b915061117883610a15565b92508282101561118b5761118a610f3f565b5b828203905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f416c726561647920526567697374657265640000000000000000000000000000600082015250565b60006111fb601283610aad565b9150611206826111c5565b602082019050919050565b6000602082019050818103600083015261122a816111ee565b9050919050565b6000602082019050611246600083018461112a565b9291505056fea2646970667358221220f5b071e6aac18bd5a7993d2a2e91f41b4c6d115b1112b0991491f7dfa22e429e64736f6c634300080d0033",
}

// VotingABI is the input ABI used to generate the binding from.
// Deprecated: Use VotingMetaData.ABI instead.
var VotingABI = VotingMetaData.ABI

// VotingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VotingMetaData.Bin instead.
var VotingBin = VotingMetaData.Bin

// DeployVoting deploys a new Ethereum contract, binding an instance of Voting to it.
func DeployVoting(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Voting, error) {
	parsed, err := VotingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VotingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Voting{VotingCaller: VotingCaller{contract: contract}, VotingTransactor: VotingTransactor{contract: contract}, VotingFilterer: VotingFilterer{contract: contract}}, nil
}

// Voting is an auto generated Go binding around an Ethereum contract.
type Voting struct {
	VotingCaller     // Read-only binding to the contract
	VotingTransactor // Write-only binding to the contract
	VotingFilterer   // Log filterer for contract events
}

// VotingCaller is an auto generated read-only Go binding around an Ethereum contract.
type VotingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VotingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VotingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VotingSession struct {
	Contract     *Voting           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VotingCallerSession struct {
	Contract *VotingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VotingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VotingTransactorSession struct {
	Contract     *VotingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingRaw is an auto generated low-level Go binding around an Ethereum contract.
type VotingRaw struct {
	Contract *Voting // Generic contract binding to access the raw methods on
}

// VotingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VotingCallerRaw struct {
	Contract *VotingCaller // Generic read-only contract binding to access the raw methods on
}

// VotingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VotingTransactorRaw struct {
	Contract *VotingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoting creates a new instance of Voting, bound to a specific deployed contract.
func NewVoting(address common.Address, backend bind.ContractBackend) (*Voting, error) {
	contract, err := bindVoting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Voting{VotingCaller: VotingCaller{contract: contract}, VotingTransactor: VotingTransactor{contract: contract}, VotingFilterer: VotingFilterer{contract: contract}}, nil
}

// NewVotingCaller creates a new read-only instance of Voting, bound to a specific deployed contract.
func NewVotingCaller(address common.Address, caller bind.ContractCaller) (*VotingCaller, error) {
	contract, err := bindVoting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VotingCaller{contract: contract}, nil
}

// NewVotingTransactor creates a new write-only instance of Voting, bound to a specific deployed contract.
func NewVotingTransactor(address common.Address, transactor bind.ContractTransactor) (*VotingTransactor, error) {
	contract, err := bindVoting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VotingTransactor{contract: contract}, nil
}

// NewVotingFilterer creates a new log filterer instance of Voting, bound to a specific deployed contract.
func NewVotingFilterer(address common.Address, filterer bind.ContractFilterer) (*VotingFilterer, error) {
	contract, err := bindVoting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VotingFilterer{contract: contract}, nil
}

// bindVoting binds a generic wrapper to an already deployed contract.
func bindVoting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VotingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voting *VotingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Voting.Contract.VotingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voting *VotingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.Contract.VotingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voting *VotingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voting.Contract.VotingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voting *VotingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Voting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voting *VotingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voting *VotingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voting.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Voting *VotingCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Voting *VotingSession) Admin() (common.Address, error) {
	return _Voting.Contract.Admin(&_Voting.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Voting *VotingCallerSession) Admin() (common.Address, error) {
	return _Voting.Contract.Admin(&_Voting.CallOpts)
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates(uint256 ) view returns(uint256 id, string name, uint256 voteCount)
func (_Voting *VotingCaller) Candidates(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id        *big.Int
	Name      string
	VoteCount *big.Int
}, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "candidates", arg0)

	outstruct := new(struct {
		Id        *big.Int
		Name      string
		VoteCount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.VoteCount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates(uint256 ) view returns(uint256 id, string name, uint256 voteCount)
func (_Voting *VotingSession) Candidates(arg0 *big.Int) (struct {
	Id        *big.Int
	Name      string
	VoteCount *big.Int
}, error) {
	return _Voting.Contract.Candidates(&_Voting.CallOpts, arg0)
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates(uint256 ) view returns(uint256 id, string name, uint256 voteCount)
func (_Voting *VotingCallerSession) Candidates(arg0 *big.Int) (struct {
	Id        *big.Int
	Name      string
	VoteCount *big.Int
}, error) {
	return _Voting.Contract.Candidates(&_Voting.CallOpts, arg0)
}

// CandidatesCount is a free data retrieval call binding the contract method 0x2d35a8a2.
//
// Solidity: function candidatesCount() view returns(uint256)
func (_Voting *VotingCaller) CandidatesCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "candidatesCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CandidatesCount is a free data retrieval call binding the contract method 0x2d35a8a2.
//
// Solidity: function candidatesCount() view returns(uint256)
func (_Voting *VotingSession) CandidatesCount() (*big.Int, error) {
	return _Voting.Contract.CandidatesCount(&_Voting.CallOpts)
}

// CandidatesCount is a free data retrieval call binding the contract method 0x2d35a8a2.
//
// Solidity: function candidatesCount() view returns(uint256)
func (_Voting *VotingCallerSession) CandidatesCount() (*big.Int, error) {
	return _Voting.Contract.CandidatesCount(&_Voting.CallOpts)
}

// CheckVoterStatus is a free data retrieval call binding the contract method 0xb44cf604.
//
// Solidity: function checkVoterStatus(bytes32 voterHash) view returns(bool)
func (_Voting *VotingCaller) CheckVoterStatus(opts *bind.CallOpts, voterHash [32]byte) (bool, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "checkVoterStatus", voterHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckVoterStatus is a free data retrieval call binding the contract method 0xb44cf604.
//
// Solidity: function checkVoterStatus(bytes32 voterHash) view returns(bool)
func (_Voting *VotingSession) CheckVoterStatus(voterHash [32]byte) (bool, error) {
	return _Voting.Contract.CheckVoterStatus(&_Voting.CallOpts, voterHash)
}

// CheckVoterStatus is a free data retrieval call binding the contract method 0xb44cf604.
//
// Solidity: function checkVoterStatus(bytes32 voterHash) view returns(bool)
func (_Voting *VotingCallerSession) CheckVoterStatus(voterHash [32]byte) (bool, error) {
	return _Voting.Contract.CheckVoterStatus(&_Voting.CallOpts, voterHash)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256[])
func (_Voting *VotingCaller) GetCount(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "getCount")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256[])
func (_Voting *VotingSession) GetCount() ([]*big.Int, error) {
	return _Voting.Contract.GetCount(&_Voting.CallOpts)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256[])
func (_Voting *VotingCallerSession) GetCount() ([]*big.Int, error) {
	return _Voting.Contract.GetCount(&_Voting.CallOpts)
}

// GetVoter is a free data retrieval call binding the contract method 0xd8e63542.
//
// Solidity: function getVoter(bytes32 voterHash) view returns(bool)
func (_Voting *VotingCaller) GetVoter(opts *bind.CallOpts, voterHash [32]byte) (bool, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "getVoter", voterHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetVoter is a free data retrieval call binding the contract method 0xd8e63542.
//
// Solidity: function getVoter(bytes32 voterHash) view returns(bool)
func (_Voting *VotingSession) GetVoter(voterHash [32]byte) (bool, error) {
	return _Voting.Contract.GetVoter(&_Voting.CallOpts, voterHash)
}

// GetVoter is a free data retrieval call binding the contract method 0xd8e63542.
//
// Solidity: function getVoter(bytes32 voterHash) view returns(bool)
func (_Voting *VotingCallerSession) GetVoter(voterHash [32]byte) (bool, error) {
	return _Voting.Contract.GetVoter(&_Voting.CallOpts, voterHash)
}

// HasVoted is a free data retrieval call binding the contract method 0x1b4613cb.
//
// Solidity: function hasVoted(bytes32 ) view returns(bool)
func (_Voting *VotingCaller) HasVoted(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "hasVoted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVoted is a free data retrieval call binding the contract method 0x1b4613cb.
//
// Solidity: function hasVoted(bytes32 ) view returns(bool)
func (_Voting *VotingSession) HasVoted(arg0 [32]byte) (bool, error) {
	return _Voting.Contract.HasVoted(&_Voting.CallOpts, arg0)
}

// HasVoted is a free data retrieval call binding the contract method 0x1b4613cb.
//
// Solidity: function hasVoted(bytes32 ) view returns(bool)
func (_Voting *VotingCallerSession) HasVoted(arg0 [32]byte) (bool, error) {
	return _Voting.Contract.HasVoted(&_Voting.CallOpts, arg0)
}

// IsEligible is a free data retrieval call binding the contract method 0x3c74e0d9.
//
// Solidity: function isEligible(bytes32 ) view returns(bool)
func (_Voting *VotingCaller) IsEligible(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "isEligible", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEligible is a free data retrieval call binding the contract method 0x3c74e0d9.
//
// Solidity: function isEligible(bytes32 ) view returns(bool)
func (_Voting *VotingSession) IsEligible(arg0 [32]byte) (bool, error) {
	return _Voting.Contract.IsEligible(&_Voting.CallOpts, arg0)
}

// IsEligible is a free data retrieval call binding the contract method 0x3c74e0d9.
//
// Solidity: function isEligible(bytes32 ) view returns(bool)
func (_Voting *VotingCallerSession) IsEligible(arg0 [32]byte) (bool, error) {
	return _Voting.Contract.IsEligible(&_Voting.CallOpts, arg0)
}

// AddCandidate is a paid mutator transaction binding the contract method 0x462e91ec.
//
// Solidity: function addCandidate(string _name) returns()
func (_Voting *VotingTransactor) AddCandidate(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "addCandidate", _name)
}

// AddCandidate is a paid mutator transaction binding the contract method 0x462e91ec.
//
// Solidity: function addCandidate(string _name) returns()
func (_Voting *VotingSession) AddCandidate(_name string) (*types.Transaction, error) {
	return _Voting.Contract.AddCandidate(&_Voting.TransactOpts, _name)
}

// AddCandidate is a paid mutator transaction binding the contract method 0x462e91ec.
//
// Solidity: function addCandidate(string _name) returns()
func (_Voting *VotingTransactorSession) AddCandidate(_name string) (*types.Transaction, error) {
	return _Voting.Contract.AddCandidate(&_Voting.TransactOpts, _name)
}

// RegisterVoterHash is a paid mutator transaction binding the contract method 0xe2a18da0.
//
// Solidity: function registerVoterHash(bytes32 voterHash) returns()
func (_Voting *VotingTransactor) RegisterVoterHash(opts *bind.TransactOpts, voterHash [32]byte) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "registerVoterHash", voterHash)
}

// RegisterVoterHash is a paid mutator transaction binding the contract method 0xe2a18da0.
//
// Solidity: function registerVoterHash(bytes32 voterHash) returns()
func (_Voting *VotingSession) RegisterVoterHash(voterHash [32]byte) (*types.Transaction, error) {
	return _Voting.Contract.RegisterVoterHash(&_Voting.TransactOpts, voterHash)
}

// RegisterVoterHash is a paid mutator transaction binding the contract method 0xe2a18da0.
//
// Solidity: function registerVoterHash(bytes32 voterHash) returns()
func (_Voting *VotingTransactorSession) RegisterVoterHash(voterHash [32]byte) (*types.Transaction, error) {
	return _Voting.Contract.RegisterVoterHash(&_Voting.TransactOpts, voterHash)
}

// Vote is a paid mutator transaction binding the contract method 0x68bb8bb6.
//
// Solidity: function vote(uint256 candidateId, bytes32 voterHash) returns()
func (_Voting *VotingTransactor) Vote(opts *bind.TransactOpts, candidateId *big.Int, voterHash [32]byte) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "vote", candidateId, voterHash)
}

// Vote is a paid mutator transaction binding the contract method 0x68bb8bb6.
//
// Solidity: function vote(uint256 candidateId, bytes32 voterHash) returns()
func (_Voting *VotingSession) Vote(candidateId *big.Int, voterHash [32]byte) (*types.Transaction, error) {
	return _Voting.Contract.Vote(&_Voting.TransactOpts, candidateId, voterHash)
}

// Vote is a paid mutator transaction binding the contract method 0x68bb8bb6.
//
// Solidity: function vote(uint256 candidateId, bytes32 voterHash) returns()
func (_Voting *VotingTransactorSession) Vote(candidateId *big.Int, voterHash [32]byte) (*types.Transaction, error) {
	return _Voting.Contract.Vote(&_Voting.TransactOpts, candidateId, voterHash)
}

// VotingCandidateAddedIterator is returned from FilterCandidateAdded and is used to iterate over the raw logs and unpacked data for CandidateAdded events raised by the Voting contract.
type VotingCandidateAddedIterator struct {
	Event *VotingCandidateAdded // Event containing the contract specifics and raw log

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
func (it *VotingCandidateAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingCandidateAdded)
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
		it.Event = new(VotingCandidateAdded)
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
func (it *VotingCandidateAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingCandidateAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingCandidateAdded represents a CandidateAdded event raised by the Voting contract.
type VotingCandidateAdded struct {
	CandidateId *big.Int
	Name        string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCandidateAdded is a free log retrieval operation binding the contract event 0xe83b2a43e7e82d975c8a0a6d2f045153c869e111136a34d1889ab7b598e396a3.
//
// Solidity: event CandidateAdded(uint256 candidateId, string name)
func (_Voting *VotingFilterer) FilterCandidateAdded(opts *bind.FilterOpts) (*VotingCandidateAddedIterator, error) {

	logs, sub, err := _Voting.contract.FilterLogs(opts, "CandidateAdded")
	if err != nil {
		return nil, err
	}
	return &VotingCandidateAddedIterator{contract: _Voting.contract, event: "CandidateAdded", logs: logs, sub: sub}, nil
}

// WatchCandidateAdded is a free log subscription operation binding the contract event 0xe83b2a43e7e82d975c8a0a6d2f045153c869e111136a34d1889ab7b598e396a3.
//
// Solidity: event CandidateAdded(uint256 candidateId, string name)
func (_Voting *VotingFilterer) WatchCandidateAdded(opts *bind.WatchOpts, sink chan<- *VotingCandidateAdded) (event.Subscription, error) {

	logs, sub, err := _Voting.contract.WatchLogs(opts, "CandidateAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingCandidateAdded)
				if err := _Voting.contract.UnpackLog(event, "CandidateAdded", log); err != nil {
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

// ParseCandidateAdded is a log parse operation binding the contract event 0xe83b2a43e7e82d975c8a0a6d2f045153c869e111136a34d1889ab7b598e396a3.
//
// Solidity: event CandidateAdded(uint256 candidateId, string name)
func (_Voting *VotingFilterer) ParseCandidateAdded(log types.Log) (*VotingCandidateAdded, error) {
	event := new(VotingCandidateAdded)
	if err := _Voting.contract.UnpackLog(event, "CandidateAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingVoteCastedIterator is returned from FilterVoteCasted and is used to iterate over the raw logs and unpacked data for VoteCasted events raised by the Voting contract.
type VotingVoteCastedIterator struct {
	Event *VotingVoteCasted // Event containing the contract specifics and raw log

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
func (it *VotingVoteCastedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingVoteCasted)
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
		it.Event = new(VotingVoteCasted)
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
func (it *VotingVoteCastedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingVoteCastedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingVoteCasted represents a VoteCasted event raised by the Voting contract.
type VotingVoteCasted struct {
	CandidateId *big.Int
	VoterHash   [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVoteCasted is a free log retrieval operation binding the contract event 0xc8058fd305bbbcba2a5a8df5471ff25dd0f41cc61f4e2a9f0ed519b44f848847.
//
// Solidity: event VoteCasted(uint256 candidateId, bytes32 voterHash)
func (_Voting *VotingFilterer) FilterVoteCasted(opts *bind.FilterOpts) (*VotingVoteCastedIterator, error) {

	logs, sub, err := _Voting.contract.FilterLogs(opts, "VoteCasted")
	if err != nil {
		return nil, err
	}
	return &VotingVoteCastedIterator{contract: _Voting.contract, event: "VoteCasted", logs: logs, sub: sub}, nil
}

// WatchVoteCasted is a free log subscription operation binding the contract event 0xc8058fd305bbbcba2a5a8df5471ff25dd0f41cc61f4e2a9f0ed519b44f848847.
//
// Solidity: event VoteCasted(uint256 candidateId, bytes32 voterHash)
func (_Voting *VotingFilterer) WatchVoteCasted(opts *bind.WatchOpts, sink chan<- *VotingVoteCasted) (event.Subscription, error) {

	logs, sub, err := _Voting.contract.WatchLogs(opts, "VoteCasted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingVoteCasted)
				if err := _Voting.contract.UnpackLog(event, "VoteCasted", log); err != nil {
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

// ParseVoteCasted is a log parse operation binding the contract event 0xc8058fd305bbbcba2a5a8df5471ff25dd0f41cc61f4e2a9f0ed519b44f848847.
//
// Solidity: event VoteCasted(uint256 candidateId, bytes32 voterHash)
func (_Voting *VotingFilterer) ParseVoteCasted(log types.Log) (*VotingVoteCasted, error) {
	event := new(VotingVoteCasted)
	if err := _Voting.contract.UnpackLog(event, "VoteCasted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingVoterRegisteredIterator is returned from FilterVoterRegistered and is used to iterate over the raw logs and unpacked data for VoterRegistered events raised by the Voting contract.
type VotingVoterRegisteredIterator struct {
	Event *VotingVoterRegistered // Event containing the contract specifics and raw log

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
func (it *VotingVoterRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingVoterRegistered)
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
		it.Event = new(VotingVoterRegistered)
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
func (it *VotingVoterRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingVoterRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingVoterRegistered represents a VoterRegistered event raised by the Voting contract.
type VotingVoterRegistered struct {
	VoterHash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVoterRegistered is a free log retrieval operation binding the contract event 0xe1e93aab4d6cae623ae197f9c4d63ac5eede4c69f9d0fbc92e4dbb0766abd115.
//
// Solidity: event VoterRegistered(bytes32 voterHash)
func (_Voting *VotingFilterer) FilterVoterRegistered(opts *bind.FilterOpts) (*VotingVoterRegisteredIterator, error) {

	logs, sub, err := _Voting.contract.FilterLogs(opts, "VoterRegistered")
	if err != nil {
		return nil, err
	}
	return &VotingVoterRegisteredIterator{contract: _Voting.contract, event: "VoterRegistered", logs: logs, sub: sub}, nil
}

// WatchVoterRegistered is a free log subscription operation binding the contract event 0xe1e93aab4d6cae623ae197f9c4d63ac5eede4c69f9d0fbc92e4dbb0766abd115.
//
// Solidity: event VoterRegistered(bytes32 voterHash)
func (_Voting *VotingFilterer) WatchVoterRegistered(opts *bind.WatchOpts, sink chan<- *VotingVoterRegistered) (event.Subscription, error) {

	logs, sub, err := _Voting.contract.WatchLogs(opts, "VoterRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingVoterRegistered)
				if err := _Voting.contract.UnpackLog(event, "VoterRegistered", log); err != nil {
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

// ParseVoterRegistered is a log parse operation binding the contract event 0xe1e93aab4d6cae623ae197f9c4d63ac5eede4c69f9d0fbc92e4dbb0766abd115.
//
// Solidity: event VoterRegistered(bytes32 voterHash)
func (_Voting *VotingFilterer) ParseVoterRegistered(log types.Log) (*VotingVoterRegistered, error) {
	event := new(VotingVoterRegistered)
	if err := _Voting.contract.UnpackLog(event, "VoterRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
