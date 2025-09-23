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

// CPNFTMetaData contains all meta data concerning the CPNFT contract.
var CPNFTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"baseTokenURI_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumCPNFT.NFTLevel\",\"name\":\"level\",\"type\":\"uint8\"}],\"name\":\"TokenLevelSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isStaked\",\"type\":\"bool\"}],\"name\":\"TokenStakeStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"batchBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"to\",\"type\":\"address[]\"},{\"internalType\":\"enumCPNFT.NFTLevel[]\",\"name\":\"levels\",\"type\":\"uint8[]\"}],\"name\":\"batchMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"from\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"to\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"batchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getTokenLevel\",\"outputs\":[{\"internalType\":\"enumCPNFT.NFTLevel\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isStaked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"enumCPNFT.NFTLevel\",\"name\":\"level\",\"type\":\"uint8\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"staked\",\"type\":\"bool\"}],\"name\":\"setStakeStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"enumCPNFT.NFTLevel\",\"name\":\"level\",\"type\":\"uint8\"}],\"name\":\"setTokenLevel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CPNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use CPNFTMetaData.ABI instead.
var CPNFTABI = CPNFTMetaData.ABI

// CPNFT is an auto generated Go binding around an Ethereum contract.
type CPNFT struct {
	CPNFTCaller     // Read-only binding to the contract
	CPNFTTransactor // Write-only binding to the contract
	CPNFTFilterer   // Log filterer for contract events
}

// CPNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type CPNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CPNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CPNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CPNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CPNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CPNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CPNFTSession struct {
	Contract     *CPNFT            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CPNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CPNFTCallerSession struct {
	Contract *CPNFTCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CPNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CPNFTTransactorSession struct {
	Contract     *CPNFTTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CPNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type CPNFTRaw struct {
	Contract *CPNFT // Generic contract binding to access the raw methods on
}

// CPNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CPNFTCallerRaw struct {
	Contract *CPNFTCaller // Generic read-only contract binding to access the raw methods on
}

// CPNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CPNFTTransactorRaw struct {
	Contract *CPNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCPNFT creates a new instance of CPNFT, bound to a specific deployed contract.
func NewCPNFT(address common.Address, backend bind.ContractBackend) (*CPNFT, error) {
	contract, err := bindCPNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CPNFT{CPNFTCaller: CPNFTCaller{contract: contract}, CPNFTTransactor: CPNFTTransactor{contract: contract}, CPNFTFilterer: CPNFTFilterer{contract: contract}}, nil
}

// NewCPNFTCaller creates a new read-only instance of CPNFT, bound to a specific deployed contract.
func NewCPNFTCaller(address common.Address, caller bind.ContractCaller) (*CPNFTCaller, error) {
	contract, err := bindCPNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CPNFTCaller{contract: contract}, nil
}

// NewCPNFTTransactor creates a new write-only instance of CPNFT, bound to a specific deployed contract.
func NewCPNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*CPNFTTransactor, error) {
	contract, err := bindCPNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CPNFTTransactor{contract: contract}, nil
}

// NewCPNFTFilterer creates a new log filterer instance of CPNFT, bound to a specific deployed contract.
func NewCPNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*CPNFTFilterer, error) {
	contract, err := bindCPNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CPNFTFilterer{contract: contract}, nil
}

// bindCPNFT binds a generic wrapper to an already deployed contract.
func bindCPNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CPNFTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CPNFT *CPNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CPNFT.Contract.CPNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CPNFT *CPNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CPNFT.Contract.CPNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CPNFT *CPNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CPNFT.Contract.CPNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CPNFT *CPNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CPNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CPNFT *CPNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CPNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CPNFT *CPNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CPNFT.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner_) view returns(uint256)
func (_CPNFT *CPNFTCaller) BalanceOf(opts *bind.CallOpts, owner_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CPNFT.contract.Call(opts, &out, "balanceOf", owner_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner_) view returns(uint256)
func (_CPNFT *CPNFTSession) BalanceOf(owner_ common.Address) (*big.Int, error) {
	return _CPNFT.Contract.BalanceOf(&_CPNFT.CallOpts, owner_)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner_) view returns(uint256)
func (_CPNFT *CPNFTCallerSession) BalanceOf(owner_ common.Address) (*big.Int, error) {
	return _CPNFT.Contract.BalanceOf(&_CPNFT.CallOpts, owner_)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CPNFT *CPNFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CPNFT.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CPNFT *CPNFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _CPNFT.Contract.GetApproved(&_CPNFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CPNFT *CPNFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _CPNFT.Contract.GetApproved(&_CPNFT.CallOpts, tokenId)
}

// GetCurrentTokenId is a free data retrieval call binding the contract method 0x56189236.
//
// Solidity: function getCurrentTokenId() view returns(uint256)
func (_CPNFT *CPNFTCaller) GetCurrentTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CPNFT.contract.Call(opts, &out, "getCurrentTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentTokenId is a free data retrieval call binding the contract method 0x56189236.
//
// Solidity: function getCurrentTokenId() view returns(uint256)
func (_CPNFT *CPNFTSession) GetCurrentTokenId() (*big.Int, error) {
	return _CPNFT.Contract.GetCurrentTokenId(&_CPNFT.CallOpts)
}

// GetCurrentTokenId is a free data retrieval call binding the contract method 0x56189236.
//
// Solidity: function getCurrentTokenId() view returns(uint256)
func (_CPNFT *CPNFTCallerSession) GetCurrentTokenId() (*big.Int, error) {
	return _CPNFT.Contract.GetCurrentTokenId(&_CPNFT.CallOpts)
}

// GetNextTokenId is a free data retrieval call binding the contract method 0xcaa0f92a.
//
// Solidity: function getNextTokenId() view returns(uint256)
func (_CPNFT *CPNFTCaller) GetNextTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CPNFT.contract.Call(opts, &out, "getNextTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextTokenId is a free data retrieval call binding the contract method 0xcaa0f92a.
//
// Solidity: function getNextTokenId() view returns(uint256)
func (_CPNFT *CPNFTSession) GetNextTokenId() (*big.Int, error) {
	return _CPNFT.Contract.GetNextTokenId(&_CPNFT.CallOpts)
}

// GetNextTokenId is a free data retrieval call binding the contract method 0xcaa0f92a.
//
// Solidity: function getNextTokenId() view returns(uint256)
func (_CPNFT *CPNFTCallerSession) GetNextTokenId() (*big.Int, error) {
	return _CPNFT.Contract.GetNextTokenId(&_CPNFT.CallOpts)
}

// GetTokenLevel is a free data retrieval call binding the contract method 0xd011645c.
//
// Solidity: function getTokenLevel(uint256 tokenId) view returns(uint8)
func (_CPNFT *CPNFTCaller) GetTokenLevel(opts *bind.CallOpts, tokenId *big.Int) (uint8, error) {
	var out []interface{}
	err := _CPNFT.contract.Call(opts, &out, "getTokenLevel", tokenId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetTokenLevel is a free data retrieval call binding the contract method 0xd011645c.
//
// Solidity: function getTokenLevel(uint256 tokenId) view returns(uint8)
func (_CPNFT *CPNFTSession) GetTokenLevel(tokenId *big.Int) (uint8, error) {
	return _CPNFT.Contract.GetTokenLevel(&_CPNFT.CallOpts, tokenId)
}

// GetTokenLevel is a free data retrieval call binding the contract method 0xd011645c.
//
// Solidity: function getTokenLevel(uint256 tokenId) view returns(uint8)
func (_CPNFT *CPNFTCallerSession) GetTokenLevel(tokenId *big.Int) (uint8, error) {
	return _CPNFT.Contract.GetTokenLevel(&_CPNFT.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner_, address operator) view returns(bool)
func (_CPNFT *CPNFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner_ common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _CPNFT.contract.Call(opts, &out, "isApprovedForAll", owner_, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner_, address operator) view returns(bool)
func (_CPNFT *CPNFTSession) IsApprovedForAll(owner_ common.Address, operator common.Address) (bool, error) {
	return _CPNFT.Contract.IsApprovedForAll(&_CPNFT.CallOpts, owner_, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner_, address operator) view returns(bool)
func (_CPNFT *CPNFTCallerSession) IsApprovedForAll(owner_ common.Address, operator common.Address) (bool, error) {
	return _CPNFT.Contract.IsApprovedForAll(&_CPNFT.CallOpts, owner_, operator)
}

// IsStaked is a free data retrieval call binding the contract method 0xbaa51f86.
//
// Solidity: function isStaked(uint256 tokenId) view returns(bool)
func (_CPNFT *CPNFTCaller) IsStaked(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _CPNFT.contract.Call(opts, &out, "isStaked", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStaked is a free data retrieval call binding the contract method 0xbaa51f86.
//
// Solidity: function isStaked(uint256 tokenId) view returns(bool)
func (_CPNFT *CPNFTSession) IsStaked(tokenId *big.Int) (bool, error) {
	return _CPNFT.Contract.IsStaked(&_CPNFT.CallOpts, tokenId)
}

// IsStaked is a free data retrieval call binding the contract method 0xbaa51f86.
//
// Solidity: function isStaked(uint256 tokenId) view returns(bool)
func (_CPNFT *CPNFTCallerSession) IsStaked(tokenId *big.Int) (bool, error) {
	return _CPNFT.Contract.IsStaked(&_CPNFT.CallOpts, tokenId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CPNFT *CPNFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CPNFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CPNFT *CPNFTSession) Name() (string, error) {
	return _CPNFT.Contract.Name(&_CPNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CPNFT *CPNFTCallerSession) Name() (string, error) {
	return _CPNFT.Contract.Name(&_CPNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CPNFT *CPNFTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CPNFT.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CPNFT *CPNFTSession) Owner() (common.Address, error) {
	return _CPNFT.Contract.Owner(&_CPNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CPNFT *CPNFTCallerSession) Owner() (common.Address, error) {
	return _CPNFT.Contract.Owner(&_CPNFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CPNFT *CPNFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CPNFT.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CPNFT *CPNFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _CPNFT.Contract.OwnerOf(&_CPNFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CPNFT *CPNFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _CPNFT.Contract.OwnerOf(&_CPNFT.CallOpts, tokenId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CPNFT *CPNFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CPNFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CPNFT *CPNFTSession) Symbol() (string, error) {
	return _CPNFT.Contract.Symbol(&_CPNFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CPNFT *CPNFTCallerSession) Symbol() (string, error) {
	return _CPNFT.Contract.Symbol(&_CPNFT.CallOpts)
}

// TokenExists is a free data retrieval call binding the contract method 0x00923f9e.
//
// Solidity: function tokenExists(uint256 tokenId) view returns(bool)
func (_CPNFT *CPNFTCaller) TokenExists(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _CPNFT.contract.Call(opts, &out, "tokenExists", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TokenExists is a free data retrieval call binding the contract method 0x00923f9e.
//
// Solidity: function tokenExists(uint256 tokenId) view returns(bool)
func (_CPNFT *CPNFTSession) TokenExists(tokenId *big.Int) (bool, error) {
	return _CPNFT.Contract.TokenExists(&_CPNFT.CallOpts, tokenId)
}

// TokenExists is a free data retrieval call binding the contract method 0x00923f9e.
//
// Solidity: function tokenExists(uint256 tokenId) view returns(bool)
func (_CPNFT *CPNFTCallerSession) TokenExists(tokenId *big.Int) (bool, error) {
	return _CPNFT.Contract.TokenExists(&_CPNFT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CPNFT *CPNFTCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _CPNFT.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CPNFT *CPNFTSession) TokenURI(tokenId *big.Int) (string, error) {
	return _CPNFT.Contract.TokenURI(&_CPNFT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CPNFT *CPNFTCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _CPNFT.Contract.TokenURI(&_CPNFT.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CPNFT *CPNFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CPNFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CPNFT *CPNFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CPNFT.Contract.Approve(&_CPNFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CPNFT *CPNFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CPNFT.Contract.Approve(&_CPNFT.TransactOpts, to, tokenId)
}

// BatchBurn is a paid mutator transaction binding the contract method 0xdc8e92ea.
//
// Solidity: function batchBurn(uint256[] tokenIds) returns()
func (_CPNFT *CPNFTTransactor) BatchBurn(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _CPNFT.contract.Transact(opts, "batchBurn", tokenIds)
}

// BatchBurn is a paid mutator transaction binding the contract method 0xdc8e92ea.
//
// Solidity: function batchBurn(uint256[] tokenIds) returns()
func (_CPNFT *CPNFTSession) BatchBurn(tokenIds []*big.Int) (*types.Transaction, error) {
	return _CPNFT.Contract.BatchBurn(&_CPNFT.TransactOpts, tokenIds)
}

// BatchBurn is a paid mutator transaction binding the contract method 0xdc8e92ea.
//
// Solidity: function batchBurn(uint256[] tokenIds) returns()
func (_CPNFT *CPNFTTransactorSession) BatchBurn(tokenIds []*big.Int) (*types.Transaction, error) {
	return _CPNFT.Contract.BatchBurn(&_CPNFT.TransactOpts, tokenIds)
}

// BatchMint is a paid mutator transaction binding the contract method 0x79c352a7.
//
// Solidity: function batchMint(address[] to, uint8[] levels) returns()
func (_CPNFT *CPNFTTransactor) BatchMint(opts *bind.TransactOpts, to []common.Address, levels []uint8) (*types.Transaction, error) {
	return _CPNFT.contract.Transact(opts, "batchMint", to, levels)
}

// BatchMint is a paid mutator transaction binding the contract method 0x79c352a7.
//
// Solidity: function batchMint(address[] to, uint8[] levels) returns()
func (_CPNFT *CPNFTSession) BatchMint(to []common.Address, levels []uint8) (*types.Transaction, error) {
	return _CPNFT.Contract.BatchMint(&_CPNFT.TransactOpts, to, levels)
}

// BatchMint is a paid mutator transaction binding the contract method 0x79c352a7.
//
// Solidity: function batchMint(address[] to, uint8[] levels) returns()
func (_CPNFT *CPNFTTransactorSession) BatchMint(to []common.Address, levels []uint8) (*types.Transaction, error) {
	return _CPNFT.Contract.BatchMint(&_CPNFT.TransactOpts, to, levels)
}

// BatchTransferFrom is a paid mutator transaction binding the contract method 0xb818f9e4.
//
// Solidity: function batchTransferFrom(address[] from, address[] to, uint256[] tokenIds) returns()
func (_CPNFT *CPNFTTransactor) BatchTransferFrom(opts *bind.TransactOpts, from []common.Address, to []common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _CPNFT.contract.Transact(opts, "batchTransferFrom", from, to, tokenIds)
}

// BatchTransferFrom is a paid mutator transaction binding the contract method 0xb818f9e4.
//
// Solidity: function batchTransferFrom(address[] from, address[] to, uint256[] tokenIds) returns()
func (_CPNFT *CPNFTSession) BatchTransferFrom(from []common.Address, to []common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _CPNFT.Contract.BatchTransferFrom(&_CPNFT.TransactOpts, from, to, tokenIds)
}

// BatchTransferFrom is a paid mutator transaction binding the contract method 0xb818f9e4.
//
// Solidity: function batchTransferFrom(address[] from, address[] to, uint256[] tokenIds) returns()
func (_CPNFT *CPNFTTransactorSession) BatchTransferFrom(from []common.Address, to []common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _CPNFT.Contract.BatchTransferFrom(&_CPNFT.TransactOpts, from, to, tokenIds)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_CPNFT *CPNFTTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _CPNFT.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_CPNFT *CPNFTSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _CPNFT.Contract.Burn(&_CPNFT.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_CPNFT *CPNFTTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _CPNFT.Contract.Burn(&_CPNFT.TransactOpts, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x691562a0.
//
// Solidity: function mint(address to, uint8 level) returns(uint256)
func (_CPNFT *CPNFTTransactor) Mint(opts *bind.TransactOpts, to common.Address, level uint8) (*types.Transaction, error) {
	return _CPNFT.contract.Transact(opts, "mint", to, level)
}

// Mint is a paid mutator transaction binding the contract method 0x691562a0.
//
// Solidity: function mint(address to, uint8 level) returns(uint256)
func (_CPNFT *CPNFTSession) Mint(to common.Address, level uint8) (*types.Transaction, error) {
	return _CPNFT.Contract.Mint(&_CPNFT.TransactOpts, to, level)
}

// Mint is a paid mutator transaction binding the contract method 0x691562a0.
//
// Solidity: function mint(address to, uint8 level) returns(uint256)
func (_CPNFT *CPNFTTransactorSession) Mint(to common.Address, level uint8) (*types.Transaction, error) {
	return _CPNFT.Contract.Mint(&_CPNFT.TransactOpts, to, level)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CPNFT *CPNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CPNFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CPNFT *CPNFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CPNFT.Contract.SafeTransferFrom(&_CPNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CPNFT *CPNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CPNFT.Contract.SafeTransferFrom(&_CPNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_CPNFT *CPNFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _CPNFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_CPNFT *CPNFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _CPNFT.Contract.SafeTransferFrom0(&_CPNFT.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_CPNFT *CPNFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _CPNFT.Contract.SafeTransferFrom0(&_CPNFT.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CPNFT *CPNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _CPNFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CPNFT *CPNFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _CPNFT.Contract.SetApprovalForAll(&_CPNFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CPNFT *CPNFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _CPNFT.Contract.SetApprovalForAll(&_CPNFT.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_CPNFT *CPNFTTransactor) SetBaseURI(opts *bind.TransactOpts, baseURI string) (*types.Transaction, error) {
	return _CPNFT.contract.Transact(opts, "setBaseURI", baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_CPNFT *CPNFTSession) SetBaseURI(baseURI string) (*types.Transaction, error) {
	return _CPNFT.Contract.SetBaseURI(&_CPNFT.TransactOpts, baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_CPNFT *CPNFTTransactorSession) SetBaseURI(baseURI string) (*types.Transaction, error) {
	return _CPNFT.Contract.SetBaseURI(&_CPNFT.TransactOpts, baseURI)
}

// SetStakeStatus is a paid mutator transaction binding the contract method 0xe7b82624.
//
// Solidity: function setStakeStatus(uint256 tokenId, bool staked) returns()
func (_CPNFT *CPNFTTransactor) SetStakeStatus(opts *bind.TransactOpts, tokenId *big.Int, staked bool) (*types.Transaction, error) {
	return _CPNFT.contract.Transact(opts, "setStakeStatus", tokenId, staked)
}

// SetStakeStatus is a paid mutator transaction binding the contract method 0xe7b82624.
//
// Solidity: function setStakeStatus(uint256 tokenId, bool staked) returns()
func (_CPNFT *CPNFTSession) SetStakeStatus(tokenId *big.Int, staked bool) (*types.Transaction, error) {
	return _CPNFT.Contract.SetStakeStatus(&_CPNFT.TransactOpts, tokenId, staked)
}

// SetStakeStatus is a paid mutator transaction binding the contract method 0xe7b82624.
//
// Solidity: function setStakeStatus(uint256 tokenId, bool staked) returns()
func (_CPNFT *CPNFTTransactorSession) SetStakeStatus(tokenId *big.Int, staked bool) (*types.Transaction, error) {
	return _CPNFT.Contract.SetStakeStatus(&_CPNFT.TransactOpts, tokenId, staked)
}

// SetTokenLevel is a paid mutator transaction binding the contract method 0x08ad480b.
//
// Solidity: function setTokenLevel(uint256 tokenId, uint8 level) returns()
func (_CPNFT *CPNFTTransactor) SetTokenLevel(opts *bind.TransactOpts, tokenId *big.Int, level uint8) (*types.Transaction, error) {
	return _CPNFT.contract.Transact(opts, "setTokenLevel", tokenId, level)
}

// SetTokenLevel is a paid mutator transaction binding the contract method 0x08ad480b.
//
// Solidity: function setTokenLevel(uint256 tokenId, uint8 level) returns()
func (_CPNFT *CPNFTSession) SetTokenLevel(tokenId *big.Int, level uint8) (*types.Transaction, error) {
	return _CPNFT.Contract.SetTokenLevel(&_CPNFT.TransactOpts, tokenId, level)
}

// SetTokenLevel is a paid mutator transaction binding the contract method 0x08ad480b.
//
// Solidity: function setTokenLevel(uint256 tokenId, uint8 level) returns()
func (_CPNFT *CPNFTTransactorSession) SetTokenLevel(tokenId *big.Int, level uint8) (*types.Transaction, error) {
	return _CPNFT.Contract.SetTokenLevel(&_CPNFT.TransactOpts, tokenId, level)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CPNFT *CPNFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CPNFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CPNFT *CPNFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CPNFT.Contract.TransferFrom(&_CPNFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CPNFT *CPNFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CPNFT.Contract.TransferFrom(&_CPNFT.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CPNFT *CPNFTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CPNFT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CPNFT *CPNFTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CPNFT.Contract.TransferOwnership(&_CPNFT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CPNFT *CPNFTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CPNFT.Contract.TransferOwnership(&_CPNFT.TransactOpts, newOwner)
}

// CPNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CPNFT contract.
type CPNFTApprovalIterator struct {
	Event *CPNFTApproval // Event containing the contract specifics and raw log

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
func (it *CPNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CPNFTApproval)
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
		it.Event = new(CPNFTApproval)
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
func (it *CPNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CPNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CPNFTApproval represents a Approval event raised by the CPNFT contract.
type CPNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CPNFT *CPNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*CPNFTApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CPNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CPNFTApprovalIterator{contract: _CPNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CPNFT *CPNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CPNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CPNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CPNFTApproval)
				if err := _CPNFT.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CPNFT *CPNFTFilterer) ParseApproval(log types.Log) (*CPNFTApproval, error) {
	event := new(CPNFTApproval)
	if err := _CPNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CPNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the CPNFT contract.
type CPNFTApprovalForAllIterator struct {
	Event *CPNFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *CPNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CPNFTApprovalForAll)
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
		it.Event = new(CPNFTApprovalForAll)
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
func (it *CPNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CPNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CPNFTApprovalForAll represents a ApprovalForAll event raised by the CPNFT contract.
type CPNFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CPNFT *CPNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*CPNFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _CPNFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &CPNFTApprovalForAllIterator{contract: _CPNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CPNFT *CPNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *CPNFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _CPNFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CPNFTApprovalForAll)
				if err := _CPNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CPNFT *CPNFTFilterer) ParseApprovalForAll(log types.Log) (*CPNFTApprovalForAll, error) {
	event := new(CPNFTApprovalForAll)
	if err := _CPNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CPNFTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CPNFT contract.
type CPNFTOwnershipTransferredIterator struct {
	Event *CPNFTOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CPNFTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CPNFTOwnershipTransferred)
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
		it.Event = new(CPNFTOwnershipTransferred)
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
func (it *CPNFTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CPNFTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CPNFTOwnershipTransferred represents a OwnershipTransferred event raised by the CPNFT contract.
type CPNFTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CPNFT *CPNFTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CPNFTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CPNFT.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CPNFTOwnershipTransferredIterator{contract: _CPNFT.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CPNFT *CPNFTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CPNFTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CPNFT.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CPNFTOwnershipTransferred)
				if err := _CPNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CPNFT *CPNFTFilterer) ParseOwnershipTransferred(log types.Log) (*CPNFTOwnershipTransferred, error) {
	event := new(CPNFTOwnershipTransferred)
	if err := _CPNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CPNFTTokenLevelSetIterator is returned from FilterTokenLevelSet and is used to iterate over the raw logs and unpacked data for TokenLevelSet events raised by the CPNFT contract.
type CPNFTTokenLevelSetIterator struct {
	Event *CPNFTTokenLevelSet // Event containing the contract specifics and raw log

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
func (it *CPNFTTokenLevelSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CPNFTTokenLevelSet)
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
		it.Event = new(CPNFTTokenLevelSet)
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
func (it *CPNFTTokenLevelSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CPNFTTokenLevelSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CPNFTTokenLevelSet represents a TokenLevelSet event raised by the CPNFT contract.
type CPNFTTokenLevelSet struct {
	TokenId *big.Int
	Level   uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenLevelSet is a free log retrieval operation binding the contract event 0xf3a12011211dc86e9c6704718a339c980ffca05c6096d33c6579a86082b63272.
//
// Solidity: event TokenLevelSet(uint256 indexed tokenId, uint8 level)
func (_CPNFT *CPNFTFilterer) FilterTokenLevelSet(opts *bind.FilterOpts, tokenId []*big.Int) (*CPNFTTokenLevelSetIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CPNFT.contract.FilterLogs(opts, "TokenLevelSet", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CPNFTTokenLevelSetIterator{contract: _CPNFT.contract, event: "TokenLevelSet", logs: logs, sub: sub}, nil
}

// WatchTokenLevelSet is a free log subscription operation binding the contract event 0xf3a12011211dc86e9c6704718a339c980ffca05c6096d33c6579a86082b63272.
//
// Solidity: event TokenLevelSet(uint256 indexed tokenId, uint8 level)
func (_CPNFT *CPNFTFilterer) WatchTokenLevelSet(opts *bind.WatchOpts, sink chan<- *CPNFTTokenLevelSet, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CPNFT.contract.WatchLogs(opts, "TokenLevelSet", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CPNFTTokenLevelSet)
				if err := _CPNFT.contract.UnpackLog(event, "TokenLevelSet", log); err != nil {
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

// ParseTokenLevelSet is a log parse operation binding the contract event 0xf3a12011211dc86e9c6704718a339c980ffca05c6096d33c6579a86082b63272.
//
// Solidity: event TokenLevelSet(uint256 indexed tokenId, uint8 level)
func (_CPNFT *CPNFTFilterer) ParseTokenLevelSet(log types.Log) (*CPNFTTokenLevelSet, error) {
	event := new(CPNFTTokenLevelSet)
	if err := _CPNFT.contract.UnpackLog(event, "TokenLevelSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CPNFTTokenStakeStatusChangedIterator is returned from FilterTokenStakeStatusChanged and is used to iterate over the raw logs and unpacked data for TokenStakeStatusChanged events raised by the CPNFT contract.
type CPNFTTokenStakeStatusChangedIterator struct {
	Event *CPNFTTokenStakeStatusChanged // Event containing the contract specifics and raw log

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
func (it *CPNFTTokenStakeStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CPNFTTokenStakeStatusChanged)
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
		it.Event = new(CPNFTTokenStakeStatusChanged)
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
func (it *CPNFTTokenStakeStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CPNFTTokenStakeStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CPNFTTokenStakeStatusChanged represents a TokenStakeStatusChanged event raised by the CPNFT contract.
type CPNFTTokenStakeStatusChanged struct {
	TokenId  *big.Int
	IsStaked bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTokenStakeStatusChanged is a free log retrieval operation binding the contract event 0x9b5f059a05a4b3deddf04e14f11520a9e0337d12b80a3f3b611a23b62b03e278.
//
// Solidity: event TokenStakeStatusChanged(uint256 indexed tokenId, bool isStaked)
func (_CPNFT *CPNFTFilterer) FilterTokenStakeStatusChanged(opts *bind.FilterOpts, tokenId []*big.Int) (*CPNFTTokenStakeStatusChangedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CPNFT.contract.FilterLogs(opts, "TokenStakeStatusChanged", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CPNFTTokenStakeStatusChangedIterator{contract: _CPNFT.contract, event: "TokenStakeStatusChanged", logs: logs, sub: sub}, nil
}

// WatchTokenStakeStatusChanged is a free log subscription operation binding the contract event 0x9b5f059a05a4b3deddf04e14f11520a9e0337d12b80a3f3b611a23b62b03e278.
//
// Solidity: event TokenStakeStatusChanged(uint256 indexed tokenId, bool isStaked)
func (_CPNFT *CPNFTFilterer) WatchTokenStakeStatusChanged(opts *bind.WatchOpts, sink chan<- *CPNFTTokenStakeStatusChanged, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CPNFT.contract.WatchLogs(opts, "TokenStakeStatusChanged", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CPNFTTokenStakeStatusChanged)
				if err := _CPNFT.contract.UnpackLog(event, "TokenStakeStatusChanged", log); err != nil {
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

// ParseTokenStakeStatusChanged is a log parse operation binding the contract event 0x9b5f059a05a4b3deddf04e14f11520a9e0337d12b80a3f3b611a23b62b03e278.
//
// Solidity: event TokenStakeStatusChanged(uint256 indexed tokenId, bool isStaked)
func (_CPNFT *CPNFTFilterer) ParseTokenStakeStatusChanged(log types.Log) (*CPNFTTokenStakeStatusChanged, error) {
	event := new(CPNFTTokenStakeStatusChanged)
	if err := _CPNFT.contract.UnpackLog(event, "TokenStakeStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CPNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CPNFT contract.
type CPNFTTransferIterator struct {
	Event *CPNFTTransfer // Event containing the contract specifics and raw log

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
func (it *CPNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CPNFTTransfer)
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
		it.Event = new(CPNFTTransfer)
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
func (it *CPNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CPNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CPNFTTransfer represents a Transfer event raised by the CPNFT contract.
type CPNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CPNFT *CPNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*CPNFTTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CPNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CPNFTTransferIterator{contract: _CPNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CPNFT *CPNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CPNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CPNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CPNFTTransfer)
				if err := _CPNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CPNFT *CPNFTFilterer) ParseTransfer(log types.Log) (*CPNFTTransfer, error) {
	event := new(CPNFTTransfer)
	if err := _CPNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
