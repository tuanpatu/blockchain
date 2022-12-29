// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tptoken

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

// TptokenMetaData contains all meta data concerning the Tptoken contract.
var TptokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040518060400160405280600781526020017f5450546f6b656e000000000000000000000000000000000000000000000000008152506040518060400160405280600381526020017f545054000000000000000000000000000000000000000000000000000000000081525081600390816200008f9190620004c1565b508060049081620000a19190620004c1565b505050620000bf33683635c9adc5dea00000620000c560201b60201c565b620006c3565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160362000137576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200012e9062000609565b60405180910390fd5b6200014b600083836200023d60201b60201c565b80600260008282546200015f91906200065a565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254620001b691906200065a565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516200021d9190620006a6565b60405180910390a362000239600083836200024260201b60201c565b5050565b505050565b505050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620002c957607f821691505b602082108103620002df57620002de62000281565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620003497fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826200030a565b6200035586836200030a565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000620003a26200039c62000396846200036d565b62000377565b6200036d565b9050919050565b6000819050919050565b620003be8362000381565b620003d6620003cd82620003a9565b84845462000317565b825550505050565b600090565b620003ed620003de565b620003fa818484620003b3565b505050565b5b81811015620004225762000416600082620003e3565b60018101905062000400565b5050565b601f82111562000471576200043b81620002e5565b6200044684620002fa565b8101602085101562000456578190505b6200046e6200046585620002fa565b830182620003ff565b50505b505050565b600082821c905092915050565b6000620004966000198460080262000476565b1980831691505092915050565b6000620004b1838362000483565b9150826002028217905092915050565b620004cc8262000247565b67ffffffffffffffff811115620004e857620004e762000252565b5b620004f48254620002b0565b6200050182828562000426565b600060209050601f83116001811462000539576000841562000524578287015190505b620005308582620004a3565b865550620005a0565b601f1984166200054986620002e5565b60005b8281101562000573578489015182556001820191506020850194506020810190506200054c565b868310156200059357848901516200058f601f89168262000483565b8355505b6001600288020188555050505b505050505050565b600082825260208201905092915050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b6000620005f1601f83620005a8565b9150620005fe82620005b9565b602082019050919050565b600060208201905081810360008301526200062481620005e2565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600062000667826200036d565b915062000674836200036d565b92508282019050808211156200068f576200068e6200062b565b5b92915050565b620006a0816200036d565b82525050565b6000602082019050620006bd600083018462000695565b92915050565b61123880620006d36000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80633950935111610071578063395093511461016857806370a082311461019857806395d89b41146101c8578063a457c2d7146101e6578063a9059cbb14610216578063dd62ed3e14610246576100a9565b806306fdde03146100ae578063095ea7b3146100cc57806318160ddd146100fc57806323b872dd1461011a578063313ce5671461014a575b600080fd5b6100b6610276565b6040516100c39190610b15565b60405180910390f35b6100e660048036038101906100e19190610bd0565b610308565b6040516100f39190610c2b565b60405180910390f35b61010461032b565b6040516101119190610c55565b60405180910390f35b610134600480360381019061012f9190610c70565b610335565b6040516101419190610c2b565b60405180910390f35b610152610364565b60405161015f9190610cdf565b60405180910390f35b610182600480360381019061017d9190610bd0565b61036d565b60405161018f9190610c2b565b60405180910390f35b6101b260048036038101906101ad9190610cfa565b6103a4565b6040516101bf9190610c55565b60405180910390f35b6101d06103ec565b6040516101dd9190610b15565b60405180910390f35b61020060048036038101906101fb9190610bd0565b61047e565b60405161020d9190610c2b565b60405180910390f35b610230600480360381019061022b9190610bd0565b6104f5565b60405161023d9190610c2b565b60405180910390f35b610260600480360381019061025b9190610d27565b610518565b60405161026d9190610c55565b60405180910390f35b60606003805461028590610d96565b80601f01602080910402602001604051908101604052809291908181526020018280546102b190610d96565b80156102fe5780601f106102d3576101008083540402835291602001916102fe565b820191906000526020600020905b8154815290600101906020018083116102e157829003601f168201915b5050505050905090565b60008061031361059f565b90506103208185856105a7565b600191505092915050565b6000600254905090565b60008061034061059f565b905061034d858285610770565b6103588585856107fc565b60019150509392505050565b60006012905090565b60008061037861059f565b905061039981858561038a8589610518565b6103949190610df6565b6105a7565b600191505092915050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6060600480546103fb90610d96565b80601f016020809104026020016040519081016040528092919081815260200182805461042790610d96565b80156104745780601f1061044957610100808354040283529160200191610474565b820191906000526020600020905b81548152906001019060200180831161045757829003601f168201915b5050505050905090565b60008061048961059f565b905060006104978286610518565b9050838110156104dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104d390610e9c565b60405180910390fd5b6104e982868684036105a7565b60019250505092915050565b60008061050061059f565b905061050d8185856107fc565b600191505092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610616576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161060d90610f2e565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610685576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161067c90610fc0565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040516107639190610c55565b60405180910390a3505050565b600061077c8484610518565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146107f657818110156107e8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107df9061102c565b60405180910390fd5b6107f584848484036105a7565b5b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361086b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610862906110be565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036108da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108d190611150565b60405180910390fd5b6108e5838383610a7b565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508181101561096b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610962906111e2565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546109fe9190610df6565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610a629190610c55565b60405180910390a3610a75848484610a80565b50505050565b505050565b505050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610abf578082015181840152602081019050610aa4565b60008484015250505050565b6000601f19601f8301169050919050565b6000610ae782610a85565b610af18185610a90565b9350610b01818560208601610aa1565b610b0a81610acb565b840191505092915050565b60006020820190508181036000830152610b2f8184610adc565b905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610b6782610b3c565b9050919050565b610b7781610b5c565b8114610b8257600080fd5b50565b600081359050610b9481610b6e565b92915050565b6000819050919050565b610bad81610b9a565b8114610bb857600080fd5b50565b600081359050610bca81610ba4565b92915050565b60008060408385031215610be757610be6610b37565b5b6000610bf585828601610b85565b9250506020610c0685828601610bbb565b9150509250929050565b60008115159050919050565b610c2581610c10565b82525050565b6000602082019050610c406000830184610c1c565b92915050565b610c4f81610b9a565b82525050565b6000602082019050610c6a6000830184610c46565b92915050565b600080600060608486031215610c8957610c88610b37565b5b6000610c9786828701610b85565b9350506020610ca886828701610b85565b9250506040610cb986828701610bbb565b9150509250925092565b600060ff82169050919050565b610cd981610cc3565b82525050565b6000602082019050610cf46000830184610cd0565b92915050565b600060208284031215610d1057610d0f610b37565b5b6000610d1e84828501610b85565b91505092915050565b60008060408385031215610d3e57610d3d610b37565b5b6000610d4c85828601610b85565b9250506020610d5d85828601610b85565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610dae57607f821691505b602082108103610dc157610dc0610d67565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610e0182610b9a565b9150610e0c83610b9a565b9250828201905080821115610e2457610e23610dc7565b5b92915050565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b6000610e86602583610a90565b9150610e9182610e2a565b604082019050919050565b60006020820190508181036000830152610eb581610e79565b9050919050565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b6000610f18602483610a90565b9150610f2382610ebc565b604082019050919050565b60006020820190508181036000830152610f4781610f0b565b9050919050565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b6000610faa602283610a90565b9150610fb582610f4e565b604082019050919050565b60006020820190508181036000830152610fd981610f9d565b9050919050565b7f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000600082015250565b6000611016601d83610a90565b915061102182610fe0565b602082019050919050565b6000602082019050818103600083015261104581611009565b9050919050565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b60006110a8602583610a90565b91506110b38261104c565b604082019050919050565b600060208201905081810360008301526110d78161109b565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b600061113a602383610a90565b9150611145826110de565b604082019050919050565b600060208201905081810360008301526111698161112d565b9050919050565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b60006111cc602683610a90565b91506111d782611170565b604082019050919050565b600060208201905081810360008301526111fb816111bf565b905091905056fea26469706673582212203a7ef60b7aba92b6a349b8ce956577e8ffddfdbef9e565833d0ce9d39418dd3364736f6c63430008110033",
}

// TptokenABI is the input ABI used to generate the binding from.
// Deprecated: Use TptokenMetaData.ABI instead.
var TptokenABI = TptokenMetaData.ABI

// TptokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TptokenMetaData.Bin instead.
var TptokenBin = TptokenMetaData.Bin

// DeployTptoken deploys a new Ethereum contract, binding an instance of Tptoken to it.
func DeployTptoken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Tptoken, error) {
	parsed, err := TptokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TptokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Tptoken{TptokenCaller: TptokenCaller{contract: contract}, TptokenTransactor: TptokenTransactor{contract: contract}, TptokenFilterer: TptokenFilterer{contract: contract}}, nil
}

// Tptoken is an auto generated Go binding around an Ethereum contract.
type Tptoken struct {
	TptokenCaller     // Read-only binding to the contract
	TptokenTransactor // Write-only binding to the contract
	TptokenFilterer   // Log filterer for contract events
}

// TptokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TptokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TptokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TptokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TptokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TptokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TptokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TptokenSession struct {
	Contract     *Tptoken          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TptokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TptokenCallerSession struct {
	Contract *TptokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// TptokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TptokenTransactorSession struct {
	Contract     *TptokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TptokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TptokenRaw struct {
	Contract *Tptoken // Generic contract binding to access the raw methods on
}

// TptokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TptokenCallerRaw struct {
	Contract *TptokenCaller // Generic read-only contract binding to access the raw methods on
}

// TptokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TptokenTransactorRaw struct {
	Contract *TptokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTptoken creates a new instance of Tptoken, bound to a specific deployed contract.
func NewTptoken(address common.Address, backend bind.ContractBackend) (*Tptoken, error) {
	contract, err := bindTptoken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tptoken{TptokenCaller: TptokenCaller{contract: contract}, TptokenTransactor: TptokenTransactor{contract: contract}, TptokenFilterer: TptokenFilterer{contract: contract}}, nil
}

// NewTptokenCaller creates a new read-only instance of Tptoken, bound to a specific deployed contract.
func NewTptokenCaller(address common.Address, caller bind.ContractCaller) (*TptokenCaller, error) {
	contract, err := bindTptoken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TptokenCaller{contract: contract}, nil
}

// NewTptokenTransactor creates a new write-only instance of Tptoken, bound to a specific deployed contract.
func NewTptokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TptokenTransactor, error) {
	contract, err := bindTptoken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TptokenTransactor{contract: contract}, nil
}

// NewTptokenFilterer creates a new log filterer instance of Tptoken, bound to a specific deployed contract.
func NewTptokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TptokenFilterer, error) {
	contract, err := bindTptoken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TptokenFilterer{contract: contract}, nil
}

// bindTptoken binds a generic wrapper to an already deployed contract.
func bindTptoken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TptokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tptoken *TptokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tptoken.Contract.TptokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tptoken *TptokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tptoken.Contract.TptokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tptoken *TptokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tptoken.Contract.TptokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tptoken *TptokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tptoken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tptoken *TptokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tptoken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tptoken *TptokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tptoken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Tptoken *TptokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tptoken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Tptoken *TptokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Tptoken.Contract.Allowance(&_Tptoken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Tptoken *TptokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Tptoken.Contract.Allowance(&_Tptoken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Tptoken *TptokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tptoken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Tptoken *TptokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Tptoken.Contract.BalanceOf(&_Tptoken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Tptoken *TptokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Tptoken.Contract.BalanceOf(&_Tptoken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Tptoken *TptokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tptoken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Tptoken *TptokenSession) Decimals() (uint8, error) {
	return _Tptoken.Contract.Decimals(&_Tptoken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Tptoken *TptokenCallerSession) Decimals() (uint8, error) {
	return _Tptoken.Contract.Decimals(&_Tptoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Tptoken *TptokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Tptoken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Tptoken *TptokenSession) Name() (string, error) {
	return _Tptoken.Contract.Name(&_Tptoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Tptoken *TptokenCallerSession) Name() (string, error) {
	return _Tptoken.Contract.Name(&_Tptoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Tptoken *TptokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Tptoken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Tptoken *TptokenSession) Symbol() (string, error) {
	return _Tptoken.Contract.Symbol(&_Tptoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Tptoken *TptokenCallerSession) Symbol() (string, error) {
	return _Tptoken.Contract.Symbol(&_Tptoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Tptoken *TptokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tptoken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Tptoken *TptokenSession) TotalSupply() (*big.Int, error) {
	return _Tptoken.Contract.TotalSupply(&_Tptoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Tptoken *TptokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Tptoken.Contract.TotalSupply(&_Tptoken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Tptoken *TptokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tptoken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Tptoken *TptokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tptoken.Contract.Approve(&_Tptoken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Tptoken *TptokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tptoken.Contract.Approve(&_Tptoken.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Tptoken *TptokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Tptoken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Tptoken *TptokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Tptoken.Contract.DecreaseAllowance(&_Tptoken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Tptoken *TptokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Tptoken.Contract.DecreaseAllowance(&_Tptoken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Tptoken *TptokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Tptoken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Tptoken *TptokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Tptoken.Contract.IncreaseAllowance(&_Tptoken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Tptoken *TptokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Tptoken.Contract.IncreaseAllowance(&_Tptoken.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Tptoken *TptokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tptoken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Tptoken *TptokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tptoken.Contract.Transfer(&_Tptoken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Tptoken *TptokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tptoken.Contract.Transfer(&_Tptoken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Tptoken *TptokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tptoken.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Tptoken *TptokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tptoken.Contract.TransferFrom(&_Tptoken.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Tptoken *TptokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tptoken.Contract.TransferFrom(&_Tptoken.TransactOpts, from, to, amount)
}

// TptokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Tptoken contract.
type TptokenApprovalIterator struct {
	Event *TptokenApproval // Event containing the contract specifics and raw log

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
func (it *TptokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TptokenApproval)
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
		it.Event = new(TptokenApproval)
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
func (it *TptokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TptokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TptokenApproval represents a Approval event raised by the Tptoken contract.
type TptokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Tptoken *TptokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TptokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Tptoken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TptokenApprovalIterator{contract: _Tptoken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Tptoken *TptokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TptokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Tptoken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TptokenApproval)
				if err := _Tptoken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Tptoken *TptokenFilterer) ParseApproval(log types.Log) (*TptokenApproval, error) {
	event := new(TptokenApproval)
	if err := _Tptoken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TptokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Tptoken contract.
type TptokenTransferIterator struct {
	Event *TptokenTransfer // Event containing the contract specifics and raw log

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
func (it *TptokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TptokenTransfer)
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
		it.Event = new(TptokenTransfer)
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
func (it *TptokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TptokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TptokenTransfer represents a Transfer event raised by the Tptoken contract.
type TptokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Tptoken *TptokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TptokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Tptoken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TptokenTransferIterator{contract: _Tptoken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Tptoken *TptokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TptokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Tptoken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TptokenTransfer)
				if err := _Tptoken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Tptoken *TptokenFilterer) ParseTransfer(log types.Log) (*TptokenTransfer, error) {
	event := new(TptokenTransfer)
	if err := _Tptoken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
