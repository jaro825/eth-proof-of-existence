// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"key\",\"type\":\"bytes16\"},{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"created\",\"type\":\"uint32\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"key\",\"type\":\"bytes16\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50335f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610ce68061005c5f395ff3fe608060405234801561000f575f80fd5b5060043610610055575f3560e01c80638da5cb5b146100595780638db70012146100775780638ffdb076146100a9578063a6f9dae1146100c5578063de8fa431146100e1575b5f80fd5b6100616100ff565b60405161006e91906104a9565b60405180910390f35b610091600480360381019061008c9190610528565b610126565b6040516100a093929190610613565b60405180910390f35b6100c360048036038101906100be91906107cf565b610249565b005b6100df60048036038101906100da9190610879565b610392565b005b6100e9610461565b6040516100f691906108bc565b60405180910390f35b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60605f805f60015f866fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff191681526020019081526020015f206040518060600160405290815f8201805461017c90610902565b80601f01602080910402602001604051908101604052809291908181526020018280546101a890610902565b80156101f35780601f106101ca576101008083540402835291602001916101f3565b820191905f5260205f20905b8154815290600101906020018083116101d657829003601f168201915b5050505050815260200160018201548152602001600282015f9054906101000a900463ffffffff1663ffffffff1663ffffffff16815250509050805f015181602001518260400151935093509350509193909250565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102d6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102cd906109b2565b60405180910390fd5b60405180606001604052808481526020018381526020018263ffffffff1681525060015f866fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff191681526020019081526020015f205f820151815f0190816103419190610b6d565b50602082015181600101556040820151816002015f6101000a81548163ffffffff021916908363ffffffff16021790555090505060025f81548092919061038790610c69565b919050555050505050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461041f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610416906109b2565b60405180910390fd5b805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f600254905090565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6104938261046a565b9050919050565b6104a381610489565b82525050565b5f6020820190506104bc5f83018461049a565b92915050565b5f604051905090565b5f80fd5b5f80fd5b5f7fffffffffffffffffffffffffffffffff0000000000000000000000000000000082169050919050565b610507816104d3565b8114610511575f80fd5b50565b5f81359050610522816104fe565b92915050565b5f6020828403121561053d5761053c6104cb565b5b5f61054a84828501610514565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561058a57808201518184015260208101905061056f565b5f8484015250505050565b5f601f19601f8301169050919050565b5f6105af82610553565b6105b9818561055d565b93506105c981856020860161056d565b6105d281610595565b840191505092915050565b5f819050919050565b6105ef816105dd565b82525050565b5f63ffffffff82169050919050565b61060d816105f5565b82525050565b5f6060820190508181035f83015261062b81866105a5565b905061063a60208301856105e6565b6106476040830184610604565b949350505050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61068d82610595565b810181811067ffffffffffffffff821117156106ac576106ab610657565b5b80604052505050565b5f6106be6104c2565b90506106ca8282610684565b919050565b5f67ffffffffffffffff8211156106e9576106e8610657565b5b6106f282610595565b9050602081019050919050565b828183375f83830152505050565b5f61071f61071a846106cf565b6106b5565b90508281526020810184848401111561073b5761073a610653565b5b6107468482856106ff565b509392505050565b5f82601f8301126107625761076161064f565b5b813561077284826020860161070d565b91505092915050565b610784816105dd565b811461078e575f80fd5b50565b5f8135905061079f8161077b565b92915050565b6107ae816105f5565b81146107b8575f80fd5b50565b5f813590506107c9816107a5565b92915050565b5f805f80608085870312156107e7576107e66104cb565b5b5f6107f487828801610514565b945050602085013567ffffffffffffffff811115610815576108146104cf565b5b6108218782880161074e565b935050604061083287828801610791565b9250506060610843878288016107bb565b91505092959194509250565b61085881610489565b8114610862575f80fd5b50565b5f813590506108738161084f565b92915050565b5f6020828403121561088e5761088d6104cb565b5b5f61089b84828501610865565b91505092915050565b5f819050919050565b6108b6816108a4565b82525050565b5f6020820190506108cf5f8301846108ad565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061091957607f821691505b60208210810361092c5761092b6108d5565b5b50919050565b5f82825260208201905092915050565b7f6f6e6c7920746865206f776e65722063616e2063616c6c20746869732066756e5f8201527f6374696f6e000000000000000000000000000000000000000000000000000000602082015250565b5f61099c602583610932565b91506109a782610942565b604082019050919050565b5f6020820190508181035f8301526109c981610990565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610a2c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826109f1565b610a3686836109f1565b95508019841693508086168417925050509392505050565b5f819050919050565b5f610a71610a6c610a67846108a4565b610a4e565b6108a4565b9050919050565b5f819050919050565b610a8a83610a57565b610a9e610a9682610a78565b8484546109fd565b825550505050565b5f90565b610ab2610aa6565b610abd818484610a81565b505050565b5b81811015610ae057610ad55f82610aaa565b600181019050610ac3565b5050565b601f821115610b2557610af6816109d0565b610aff846109e2565b81016020851015610b0e578190505b610b22610b1a856109e2565b830182610ac2565b50505b505050565b5f82821c905092915050565b5f610b455f1984600802610b2a565b1980831691505092915050565b5f610b5d8383610b36565b9150826002028217905092915050565b610b7682610553565b67ffffffffffffffff811115610b8f57610b8e610657565b5b610b998254610902565b610ba4828285610ae4565b5f60209050601f831160018114610bd5575f8415610bc3578287015190505b610bcd8582610b52565b865550610c34565b601f198416610be3866109d0565b5f5b82811015610c0a57848901518255600182019150602085019450602081019050610be5565b86831015610c275784890151610c23601f891682610b36565b8355505b6001600288020188555050505b505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610c73826108a4565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610ca557610ca4610c3c565b5b60018201905091905056fea26469706673582212203f491405302f9bbed9535bd03fd07e7588fe6ff1baabecb90362f5a4ab13be0f64736f6c63430008150033",
}

// StoreABI is the input ABI used to generate the binding from.
// Deprecated: Use StoreMetaData.ABI instead.
var StoreABI = StoreMetaData.ABI

// StoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StoreMetaData.Bin instead.
var StoreBin = StoreMetaData.Bin

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := StoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x8db70012.
//
// Solidity: function get(bytes16 key) view returns(bytes, bytes32, uint32)
func (_Store *StoreCaller) Get(opts *bind.CallOpts, key [16]byte) ([]byte, [32]byte, uint32, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "get", key)

	if err != nil {
		return *new([]byte), *new([32]byte), *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	out2 := *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return out0, out1, out2, err

}

// Get is a free data retrieval call binding the contract method 0x8db70012.
//
// Solidity: function get(bytes16 key) view returns(bytes, bytes32, uint32)
func (_Store *StoreSession) Get(key [16]byte) ([]byte, [32]byte, uint32, error) {
	return _Store.Contract.Get(&_Store.CallOpts, key)
}

// Get is a free data retrieval call binding the contract method 0x8db70012.
//
// Solidity: function get(bytes16 key) view returns(bytes, bytes32, uint32)
func (_Store *StoreCallerSession) Get(key [16]byte) ([]byte, [32]byte, uint32, error) {
	return _Store.Contract.Get(&_Store.CallOpts, key)
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() view returns(uint256)
func (_Store *StoreCaller) GetSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() view returns(uint256)
func (_Store *StoreSession) GetSize() (*big.Int, error) {
	return _Store.Contract.GetSize(&_Store.CallOpts)
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() view returns(uint256)
func (_Store *StoreCallerSession) GetSize() (*big.Int, error) {
	return _Store.Contract.GetSize(&_Store.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Store *StoreCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Store *StoreSession) Owner() (common.Address, error) {
	return _Store.Contract.Owner(&_Store.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Store *StoreCallerSession) Owner() (common.Address, error) {
	return _Store.Contract.Owner(&_Store.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0x8ffdb076.
//
// Solidity: function add(bytes16 key, bytes name, bytes32 hash, uint32 created) returns()
func (_Store *StoreTransactor) Add(opts *bind.TransactOpts, key [16]byte, name []byte, hash [32]byte, created uint32) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "add", key, name, hash, created)
}

// Add is a paid mutator transaction binding the contract method 0x8ffdb076.
//
// Solidity: function add(bytes16 key, bytes name, bytes32 hash, uint32 created) returns()
func (_Store *StoreSession) Add(key [16]byte, name []byte, hash [32]byte, created uint32) (*types.Transaction, error) {
	return _Store.Contract.Add(&_Store.TransactOpts, key, name, hash, created)
}

// Add is a paid mutator transaction binding the contract method 0x8ffdb076.
//
// Solidity: function add(bytes16 key, bytes name, bytes32 hash, uint32 created) returns()
func (_Store *StoreTransactorSession) Add(key [16]byte, name []byte, hash [32]byte, created uint32) (*types.Transaction, error) {
	return _Store.Contract.Add(&_Store.TransactOpts, key, name, hash, created)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Store *StoreTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Store *StoreSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Store.Contract.ChangeOwner(&_Store.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Store *StoreTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Store.Contract.ChangeOwner(&_Store.TransactOpts, newOwner)
}
