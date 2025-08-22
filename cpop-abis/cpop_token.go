// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cpop

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

// CPOPTokenMetaData contains all meta data concerning the CPOPToken contract.
var CPOPTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_initialSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessDenied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ArrayLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyArray\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRole\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BURNER_ROLE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DECIMALS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SYMBOL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"adminBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"batchBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"batchMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"batchTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"from\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"to\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"batchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"roles\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CPOPTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use CPOPTokenMetaData.ABI instead.
var CPOPTokenABI = CPOPTokenMetaData.ABI

// CPOPToken is an auto generated Go binding around an Ethereum contract.
type CPOPToken struct {
	CPOPTokenCaller     // Read-only binding to the contract
	CPOPTokenTransactor // Write-only binding to the contract
	CPOPTokenFilterer   // Log filterer for contract events
}

// CPOPTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type CPOPTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CPOPTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CPOPTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CPOPTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CPOPTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CPOPTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CPOPTokenSession struct {
	Contract     *CPOPToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CPOPTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CPOPTokenCallerSession struct {
	Contract *CPOPTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CPOPTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CPOPTokenTransactorSession struct {
	Contract     *CPOPTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CPOPTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type CPOPTokenRaw struct {
	Contract *CPOPToken // Generic contract binding to access the raw methods on
}

// CPOPTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CPOPTokenCallerRaw struct {
	Contract *CPOPTokenCaller // Generic read-only contract binding to access the raw methods on
}

// CPOPTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CPOPTokenTransactorRaw struct {
	Contract *CPOPTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCPOPToken creates a new instance of CPOPToken, bound to a specific deployed contract.
func NewCPOPToken(address common.Address, backend bind.ContractBackend) (*CPOPToken, error) {
	contract, err := bindCPOPToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CPOPToken{CPOPTokenCaller: CPOPTokenCaller{contract: contract}, CPOPTokenTransactor: CPOPTokenTransactor{contract: contract}, CPOPTokenFilterer: CPOPTokenFilterer{contract: contract}}, nil
}

// NewCPOPTokenCaller creates a new read-only instance of CPOPToken, bound to a specific deployed contract.
func NewCPOPTokenCaller(address common.Address, caller bind.ContractCaller) (*CPOPTokenCaller, error) {
	contract, err := bindCPOPToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CPOPTokenCaller{contract: contract}, nil
}

// NewCPOPTokenTransactor creates a new write-only instance of CPOPToken, bound to a specific deployed contract.
func NewCPOPTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*CPOPTokenTransactor, error) {
	contract, err := bindCPOPToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CPOPTokenTransactor{contract: contract}, nil
}

// NewCPOPTokenFilterer creates a new log filterer instance of CPOPToken, bound to a specific deployed contract.
func NewCPOPTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*CPOPTokenFilterer, error) {
	contract, err := bindCPOPToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CPOPTokenFilterer{contract: contract}, nil
}

// bindCPOPToken binds a generic wrapper to an already deployed contract.
func bindCPOPToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CPOPTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CPOPToken *CPOPTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CPOPToken.Contract.CPOPTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CPOPToken *CPOPTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CPOPToken.Contract.CPOPTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CPOPToken *CPOPTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CPOPToken.Contract.CPOPTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CPOPToken *CPOPTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CPOPToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CPOPToken *CPOPTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CPOPToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CPOPToken *CPOPTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CPOPToken.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(uint8)
func (_CPOPToken *CPOPTokenCaller) ADMINROLE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CPOPToken.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(uint8)
func (_CPOPToken *CPOPTokenSession) ADMINROLE() (uint8, error) {
	return _CPOPToken.Contract.ADMINROLE(&_CPOPToken.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(uint8)
func (_CPOPToken *CPOPTokenCallerSession) ADMINROLE() (uint8, error) {
	return _CPOPToken.Contract.ADMINROLE(&_CPOPToken.CallOpts)
}

// BURNERROLE is a free data retrieval call binding the contract method 0x282c51f3.
//
// Solidity: function BURNER_ROLE() view returns(uint8)
func (_CPOPToken *CPOPTokenCaller) BURNERROLE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CPOPToken.contract.Call(opts, &out, "BURNER_ROLE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BURNERROLE is a free data retrieval call binding the contract method 0x282c51f3.
//
// Solidity: function BURNER_ROLE() view returns(uint8)
func (_CPOPToken *CPOPTokenSession) BURNERROLE() (uint8, error) {
	return _CPOPToken.Contract.BURNERROLE(&_CPOPToken.CallOpts)
}

// BURNERROLE is a free data retrieval call binding the contract method 0x282c51f3.
//
// Solidity: function BURNER_ROLE() view returns(uint8)
func (_CPOPToken *CPOPTokenCallerSession) BURNERROLE() (uint8, error) {
	return _CPOPToken.Contract.BURNERROLE(&_CPOPToken.CallOpts)
}

// DECIMALS is a free data retrieval call binding the contract method 0x2e0f2625.
//
// Solidity: function DECIMALS() view returns(uint8)
func (_CPOPToken *CPOPTokenCaller) DECIMALS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CPOPToken.contract.Call(opts, &out, "DECIMALS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DECIMALS is a free data retrieval call binding the contract method 0x2e0f2625.
//
// Solidity: function DECIMALS() view returns(uint8)
func (_CPOPToken *CPOPTokenSession) DECIMALS() (uint8, error) {
	return _CPOPToken.Contract.DECIMALS(&_CPOPToken.CallOpts)
}

// DECIMALS is a free data retrieval call binding the contract method 0x2e0f2625.
//
// Solidity: function DECIMALS() view returns(uint8)
func (_CPOPToken *CPOPTokenCallerSession) DECIMALS() (uint8, error) {
	return _CPOPToken.Contract.DECIMALS(&_CPOPToken.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(uint8)
func (_CPOPToken *CPOPTokenCaller) MINTERROLE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CPOPToken.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(uint8)
func (_CPOPToken *CPOPTokenSession) MINTERROLE() (uint8, error) {
	return _CPOPToken.Contract.MINTERROLE(&_CPOPToken.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(uint8)
func (_CPOPToken *CPOPTokenCallerSession) MINTERROLE() (uint8, error) {
	return _CPOPToken.Contract.MINTERROLE(&_CPOPToken.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_CPOPToken *CPOPTokenCaller) NAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CPOPToken.contract.Call(opts, &out, "NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_CPOPToken *CPOPTokenSession) NAME() (string, error) {
	return _CPOPToken.Contract.NAME(&_CPOPToken.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_CPOPToken *CPOPTokenCallerSession) NAME() (string, error) {
	return _CPOPToken.Contract.NAME(&_CPOPToken.CallOpts)
}

// SYMBOL is a free data retrieval call binding the contract method 0xf76f8d78.
//
// Solidity: function SYMBOL() view returns(string)
func (_CPOPToken *CPOPTokenCaller) SYMBOL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CPOPToken.contract.Call(opts, &out, "SYMBOL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SYMBOL is a free data retrieval call binding the contract method 0xf76f8d78.
//
// Solidity: function SYMBOL() view returns(string)
func (_CPOPToken *CPOPTokenSession) SYMBOL() (string, error) {
	return _CPOPToken.Contract.SYMBOL(&_CPOPToken.CallOpts)
}

// SYMBOL is a free data retrieval call binding the contract method 0xf76f8d78.
//
// Solidity: function SYMBOL() view returns(string)
func (_CPOPToken *CPOPTokenCallerSession) SYMBOL() (string, error) {
	return _CPOPToken.Contract.SYMBOL(&_CPOPToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_CPOPToken *CPOPTokenCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CPOPToken.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_CPOPToken *CPOPTokenSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CPOPToken.Contract.Allowance(&_CPOPToken.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_CPOPToken *CPOPTokenCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CPOPToken.Contract.Allowance(&_CPOPToken.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_CPOPToken *CPOPTokenCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CPOPToken.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_CPOPToken *CPOPTokenSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CPOPToken.Contract.BalanceOf(&_CPOPToken.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_CPOPToken *CPOPTokenCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CPOPToken.Contract.BalanceOf(&_CPOPToken.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_CPOPToken *CPOPTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CPOPToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_CPOPToken *CPOPTokenSession) Decimals() (uint8, error) {
	return _CPOPToken.Contract.Decimals(&_CPOPToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_CPOPToken *CPOPTokenCallerSession) Decimals() (uint8, error) {
	return _CPOPToken.Contract.Decimals(&_CPOPToken.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x95a8c58d.
//
// Solidity: function hasRole(address account, uint8 role) view returns(bool)
func (_CPOPToken *CPOPTokenCaller) HasRole(opts *bind.CallOpts, account common.Address, role uint8) (bool, error) {
	var out []interface{}
	err := _CPOPToken.contract.Call(opts, &out, "hasRole", account, role)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x95a8c58d.
//
// Solidity: function hasRole(address account, uint8 role) view returns(bool)
func (_CPOPToken *CPOPTokenSession) HasRole(account common.Address, role uint8) (bool, error) {
	return _CPOPToken.Contract.HasRole(&_CPOPToken.CallOpts, account, role)
}

// HasRole is a free data retrieval call binding the contract method 0x95a8c58d.
//
// Solidity: function hasRole(address account, uint8 role) view returns(bool)
func (_CPOPToken *CPOPTokenCallerSession) HasRole(account common.Address, role uint8) (bool, error) {
	return _CPOPToken.Contract.HasRole(&_CPOPToken.CallOpts, account, role)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_CPOPToken *CPOPTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CPOPToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_CPOPToken *CPOPTokenSession) Name() (string, error) {
	return _CPOPToken.Contract.Name(&_CPOPToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_CPOPToken *CPOPTokenCallerSession) Name() (string, error) {
	return _CPOPToken.Contract.Name(&_CPOPToken.CallOpts)
}

// Roles is a free data retrieval call binding the contract method 0x99374642.
//
// Solidity: function roles(address ) view returns(uint8)
func (_CPOPToken *CPOPTokenCaller) Roles(opts *bind.CallOpts, arg0 common.Address) (uint8, error) {
	var out []interface{}
	err := _CPOPToken.contract.Call(opts, &out, "roles", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Roles is a free data retrieval call binding the contract method 0x99374642.
//
// Solidity: function roles(address ) view returns(uint8)
func (_CPOPToken *CPOPTokenSession) Roles(arg0 common.Address) (uint8, error) {
	return _CPOPToken.Contract.Roles(&_CPOPToken.CallOpts, arg0)
}

// Roles is a free data retrieval call binding the contract method 0x99374642.
//
// Solidity: function roles(address ) view returns(uint8)
func (_CPOPToken *CPOPTokenCallerSession) Roles(arg0 common.Address) (uint8, error) {
	return _CPOPToken.Contract.Roles(&_CPOPToken.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_CPOPToken *CPOPTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CPOPToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_CPOPToken *CPOPTokenSession) Symbol() (string, error) {
	return _CPOPToken.Contract.Symbol(&_CPOPToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_CPOPToken *CPOPTokenCallerSession) Symbol() (string, error) {
	return _CPOPToken.Contract.Symbol(&_CPOPToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CPOPToken *CPOPTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CPOPToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CPOPToken *CPOPTokenSession) TotalSupply() (*big.Int, error) {
	return _CPOPToken.Contract.TotalSupply(&_CPOPToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CPOPToken *CPOPTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _CPOPToken.Contract.TotalSupply(&_CPOPToken.CallOpts)
}

// AdminBurn is a paid mutator transaction binding the contract method 0x06dd0419.
//
// Solidity: function adminBurn(address from, uint256 amount) returns()
func (_CPOPToken *CPOPTokenTransactor) AdminBurn(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CPOPToken.contract.Transact(opts, "adminBurn", from, amount)
}

// AdminBurn is a paid mutator transaction binding the contract method 0x06dd0419.
//
// Solidity: function adminBurn(address from, uint256 amount) returns()
func (_CPOPToken *CPOPTokenSession) AdminBurn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CPOPToken.Contract.AdminBurn(&_CPOPToken.TransactOpts, from, amount)
}

// AdminBurn is a paid mutator transaction binding the contract method 0x06dd0419.
//
// Solidity: function adminBurn(address from, uint256 amount) returns()
func (_CPOPToken *CPOPTokenTransactorSession) AdminBurn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CPOPToken.Contract.AdminBurn(&_CPOPToken.TransactOpts, from, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CPOPToken *CPOPTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CPOPToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CPOPToken *CPOPTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CPOPToken.Contract.Approve(&_CPOPToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CPOPToken *CPOPTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CPOPToken.Contract.Approve(&_CPOPToken.TransactOpts, spender, amount)
}

// BatchBurn is a paid mutator transaction binding the contract method 0x4a6cc677.
//
// Solidity: function batchBurn(address[] accounts, uint256[] amounts) returns()
func (_CPOPToken *CPOPTokenTransactor) BatchBurn(opts *bind.TransactOpts, accounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _CPOPToken.contract.Transact(opts, "batchBurn", accounts, amounts)
}

// BatchBurn is a paid mutator transaction binding the contract method 0x4a6cc677.
//
// Solidity: function batchBurn(address[] accounts, uint256[] amounts) returns()
func (_CPOPToken *CPOPTokenSession) BatchBurn(accounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _CPOPToken.Contract.BatchBurn(&_CPOPToken.TransactOpts, accounts, amounts)
}

// BatchBurn is a paid mutator transaction binding the contract method 0x4a6cc677.
//
// Solidity: function batchBurn(address[] accounts, uint256[] amounts) returns()
func (_CPOPToken *CPOPTokenTransactorSession) BatchBurn(accounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _CPOPToken.Contract.BatchBurn(&_CPOPToken.TransactOpts, accounts, amounts)
}

// BatchMint is a paid mutator transaction binding the contract method 0x68573107.
//
// Solidity: function batchMint(address[] recipients, uint256[] amounts) returns()
func (_CPOPToken *CPOPTokenTransactor) BatchMint(opts *bind.TransactOpts, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _CPOPToken.contract.Transact(opts, "batchMint", recipients, amounts)
}

// BatchMint is a paid mutator transaction binding the contract method 0x68573107.
//
// Solidity: function batchMint(address[] recipients, uint256[] amounts) returns()
func (_CPOPToken *CPOPTokenSession) BatchMint(recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _CPOPToken.Contract.BatchMint(&_CPOPToken.TransactOpts, recipients, amounts)
}

// BatchMint is a paid mutator transaction binding the contract method 0x68573107.
//
// Solidity: function batchMint(address[] recipients, uint256[] amounts) returns()
func (_CPOPToken *CPOPTokenTransactorSession) BatchMint(recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _CPOPToken.Contract.BatchMint(&_CPOPToken.TransactOpts, recipients, amounts)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0x88d695b2.
//
// Solidity: function batchTransfer(address[] recipients, uint256[] amounts) returns()
func (_CPOPToken *CPOPTokenTransactor) BatchTransfer(opts *bind.TransactOpts, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _CPOPToken.contract.Transact(opts, "batchTransfer", recipients, amounts)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0x88d695b2.
//
// Solidity: function batchTransfer(address[] recipients, uint256[] amounts) returns()
func (_CPOPToken *CPOPTokenSession) BatchTransfer(recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _CPOPToken.Contract.BatchTransfer(&_CPOPToken.TransactOpts, recipients, amounts)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0x88d695b2.
//
// Solidity: function batchTransfer(address[] recipients, uint256[] amounts) returns()
func (_CPOPToken *CPOPTokenTransactorSession) BatchTransfer(recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _CPOPToken.Contract.BatchTransfer(&_CPOPToken.TransactOpts, recipients, amounts)
}

// BatchTransferFrom is a paid mutator transaction binding the contract method 0xb818f9e4.
//
// Solidity: function batchTransferFrom(address[] from, address[] to, uint256[] amounts) returns()
func (_CPOPToken *CPOPTokenTransactor) BatchTransferFrom(opts *bind.TransactOpts, from []common.Address, to []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _CPOPToken.contract.Transact(opts, "batchTransferFrom", from, to, amounts)
}

// BatchTransferFrom is a paid mutator transaction binding the contract method 0xb818f9e4.
//
// Solidity: function batchTransferFrom(address[] from, address[] to, uint256[] amounts) returns()
func (_CPOPToken *CPOPTokenSession) BatchTransferFrom(from []common.Address, to []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _CPOPToken.Contract.BatchTransferFrom(&_CPOPToken.TransactOpts, from, to, amounts)
}

// BatchTransferFrom is a paid mutator transaction binding the contract method 0xb818f9e4.
//
// Solidity: function batchTransferFrom(address[] from, address[] to, uint256[] amounts) returns()
func (_CPOPToken *CPOPTokenTransactorSession) BatchTransferFrom(from []common.Address, to []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _CPOPToken.Contract.BatchTransferFrom(&_CPOPToken.TransactOpts, from, to, amounts)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_CPOPToken *CPOPTokenTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _CPOPToken.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_CPOPToken *CPOPTokenSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _CPOPToken.Contract.Burn(&_CPOPToken.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_CPOPToken *CPOPTokenTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _CPOPToken.Contract.Burn(&_CPOPToken.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 amount) returns()
func (_CPOPToken *CPOPTokenTransactor) BurnFrom(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CPOPToken.contract.Transact(opts, "burnFrom", from, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 amount) returns()
func (_CPOPToken *CPOPTokenSession) BurnFrom(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CPOPToken.Contract.BurnFrom(&_CPOPToken.TransactOpts, from, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 amount) returns()
func (_CPOPToken *CPOPTokenTransactorSession) BurnFrom(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CPOPToken.Contract.BurnFrom(&_CPOPToken.TransactOpts, from, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x3e840236.
//
// Solidity: function grantRole(address account, uint8 role) returns()
func (_CPOPToken *CPOPTokenTransactor) GrantRole(opts *bind.TransactOpts, account common.Address, role uint8) (*types.Transaction, error) {
	return _CPOPToken.contract.Transact(opts, "grantRole", account, role)
}

// GrantRole is a paid mutator transaction binding the contract method 0x3e840236.
//
// Solidity: function grantRole(address account, uint8 role) returns()
func (_CPOPToken *CPOPTokenSession) GrantRole(account common.Address, role uint8) (*types.Transaction, error) {
	return _CPOPToken.Contract.GrantRole(&_CPOPToken.TransactOpts, account, role)
}

// GrantRole is a paid mutator transaction binding the contract method 0x3e840236.
//
// Solidity: function grantRole(address account, uint8 role) returns()
func (_CPOPToken *CPOPTokenTransactorSession) GrantRole(account common.Address, role uint8) (*types.Transaction, error) {
	return _CPOPToken.Contract.GrantRole(&_CPOPToken.TransactOpts, account, role)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_CPOPToken *CPOPTokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CPOPToken.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_CPOPToken *CPOPTokenSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CPOPToken.Contract.Mint(&_CPOPToken.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_CPOPToken *CPOPTokenTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CPOPToken.Contract.Mint(&_CPOPToken.TransactOpts, to, amount)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x077d3c03.
//
// Solidity: function revokeRole(address account, uint8 role) returns()
func (_CPOPToken *CPOPTokenTransactor) RevokeRole(opts *bind.TransactOpts, account common.Address, role uint8) (*types.Transaction, error) {
	return _CPOPToken.contract.Transact(opts, "revokeRole", account, role)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x077d3c03.
//
// Solidity: function revokeRole(address account, uint8 role) returns()
func (_CPOPToken *CPOPTokenSession) RevokeRole(account common.Address, role uint8) (*types.Transaction, error) {
	return _CPOPToken.Contract.RevokeRole(&_CPOPToken.TransactOpts, account, role)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x077d3c03.
//
// Solidity: function revokeRole(address account, uint8 role) returns()
func (_CPOPToken *CPOPTokenTransactorSession) RevokeRole(account common.Address, role uint8) (*types.Transaction, error) {
	return _CPOPToken.Contract.RevokeRole(&_CPOPToken.TransactOpts, account, role)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_CPOPToken *CPOPTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CPOPToken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_CPOPToken *CPOPTokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CPOPToken.Contract.Transfer(&_CPOPToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_CPOPToken *CPOPTokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CPOPToken.Contract.Transfer(&_CPOPToken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_CPOPToken *CPOPTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CPOPToken.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_CPOPToken *CPOPTokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CPOPToken.Contract.TransferFrom(&_CPOPToken.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_CPOPToken *CPOPTokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CPOPToken.Contract.TransferFrom(&_CPOPToken.TransactOpts, from, to, amount)
}

// CPOPTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CPOPToken contract.
type CPOPTokenApprovalIterator struct {
	Event *CPOPTokenApproval // Event containing the contract specifics and raw log

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
func (it *CPOPTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CPOPTokenApproval)
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
		it.Event = new(CPOPTokenApproval)
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
func (it *CPOPTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CPOPTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CPOPTokenApproval represents a Approval event raised by the CPOPToken contract.
type CPOPTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CPOPToken *CPOPTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CPOPTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CPOPToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CPOPTokenApprovalIterator{contract: _CPOPToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CPOPToken *CPOPTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CPOPTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CPOPToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CPOPTokenApproval)
				if err := _CPOPToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_CPOPToken *CPOPTokenFilterer) ParseApproval(log types.Log) (*CPOPTokenApproval, error) {
	event := new(CPOPTokenApproval)
	if err := _CPOPToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CPOPTokenRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the CPOPToken contract.
type CPOPTokenRoleGrantedIterator struct {
	Event *CPOPTokenRoleGranted // Event containing the contract specifics and raw log

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
func (it *CPOPTokenRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CPOPTokenRoleGranted)
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
		it.Event = new(CPOPTokenRoleGranted)
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
func (it *CPOPTokenRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CPOPTokenRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CPOPTokenRoleGranted represents a RoleGranted event raised by the CPOPToken contract.
type CPOPTokenRoleGranted struct {
	Account common.Address
	Role    uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0xaa259565575c834bc07e74dca784b4071676133ac78513b431afb6ee7edae121.
//
// Solidity: event RoleGranted(address indexed account, uint8 role)
func (_CPOPToken *CPOPTokenFilterer) FilterRoleGranted(opts *bind.FilterOpts, account []common.Address) (*CPOPTokenRoleGrantedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CPOPToken.contract.FilterLogs(opts, "RoleGranted", accountRule)
	if err != nil {
		return nil, err
	}
	return &CPOPTokenRoleGrantedIterator{contract: _CPOPToken.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0xaa259565575c834bc07e74dca784b4071676133ac78513b431afb6ee7edae121.
//
// Solidity: event RoleGranted(address indexed account, uint8 role)
func (_CPOPToken *CPOPTokenFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CPOPTokenRoleGranted, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CPOPToken.contract.WatchLogs(opts, "RoleGranted", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CPOPTokenRoleGranted)
				if err := _CPOPToken.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0xaa259565575c834bc07e74dca784b4071676133ac78513b431afb6ee7edae121.
//
// Solidity: event RoleGranted(address indexed account, uint8 role)
func (_CPOPToken *CPOPTokenFilterer) ParseRoleGranted(log types.Log) (*CPOPTokenRoleGranted, error) {
	event := new(CPOPTokenRoleGranted)
	if err := _CPOPToken.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CPOPTokenRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the CPOPToken contract.
type CPOPTokenRoleRevokedIterator struct {
	Event *CPOPTokenRoleRevoked // Event containing the contract specifics and raw log

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
func (it *CPOPTokenRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CPOPTokenRoleRevoked)
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
		it.Event = new(CPOPTokenRoleRevoked)
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
func (it *CPOPTokenRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CPOPTokenRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CPOPTokenRoleRevoked represents a RoleRevoked event raised by the CPOPToken contract.
type CPOPTokenRoleRevoked struct {
	Account common.Address
	Role    uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0x34a1009b84e077aee5bc8197faf7105e54f9ba6879e2a51a716e4156ea9ad769.
//
// Solidity: event RoleRevoked(address indexed account, uint8 role)
func (_CPOPToken *CPOPTokenFilterer) FilterRoleRevoked(opts *bind.FilterOpts, account []common.Address) (*CPOPTokenRoleRevokedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CPOPToken.contract.FilterLogs(opts, "RoleRevoked", accountRule)
	if err != nil {
		return nil, err
	}
	return &CPOPTokenRoleRevokedIterator{contract: _CPOPToken.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0x34a1009b84e077aee5bc8197faf7105e54f9ba6879e2a51a716e4156ea9ad769.
//
// Solidity: event RoleRevoked(address indexed account, uint8 role)
func (_CPOPToken *CPOPTokenFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CPOPTokenRoleRevoked, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CPOPToken.contract.WatchLogs(opts, "RoleRevoked", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CPOPTokenRoleRevoked)
				if err := _CPOPToken.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0x34a1009b84e077aee5bc8197faf7105e54f9ba6879e2a51a716e4156ea9ad769.
//
// Solidity: event RoleRevoked(address indexed account, uint8 role)
func (_CPOPToken *CPOPTokenFilterer) ParseRoleRevoked(log types.Log) (*CPOPTokenRoleRevoked, error) {
	event := new(CPOPTokenRoleRevoked)
	if err := _CPOPToken.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CPOPTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CPOPToken contract.
type CPOPTokenTransferIterator struct {
	Event *CPOPTokenTransfer // Event containing the contract specifics and raw log

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
func (it *CPOPTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CPOPTokenTransfer)
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
		it.Event = new(CPOPTokenTransfer)
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
func (it *CPOPTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CPOPTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CPOPTokenTransfer represents a Transfer event raised by the CPOPToken contract.
type CPOPTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CPOPToken *CPOPTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CPOPTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CPOPToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CPOPTokenTransferIterator{contract: _CPOPToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CPOPToken *CPOPTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CPOPTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CPOPToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CPOPTokenTransfer)
				if err := _CPOPToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_CPOPToken *CPOPTokenFilterer) ParseTransfer(log types.Log) (*CPOPTokenTransfer, error) {
	event := new(CPOPTokenTransfer)
	if err := _CPOPToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
