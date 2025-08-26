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

// AccountManagerMetaData contains all meta data concerning the AccountManager contract.
var AccountManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"masterSigner\",\"type\":\"address\"}],\"name\":\"AccountCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"CreatorAuthorized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"CreatorRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accountImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"authorizeCreator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"masterSigner\",\"type\":\"address\"}],\"name\":\"createAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"masterSigner\",\"type\":\"address\"}],\"name\":\"createAccountWithMasterSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"masterSigner\",\"type\":\"address\"}],\"name\":\"createUserAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultMasterSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPointAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"masterSigner\",\"type\":\"address\"}],\"name\":\"getAccountAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDefaultMasterSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"masterSigner\",\"type\":\"address\"}],\"name\":\"getInitCode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"entryPoint\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"masterSigner\",\"type\":\"address\"}],\"name\":\"isAccountDeployed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isDeployed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"isAuthorizedCreator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"masterAggregatorAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"revokeCreator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"senderCreator\",\"outputs\":[{\"internalType\":\"contractISenderCreator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"masterSigner\",\"type\":\"address\"}],\"name\":\"setDefaultMasterSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"setMasterAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"updateAccountImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// AccountManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use AccountManagerMetaData.ABI instead.
var AccountManagerABI = AccountManagerMetaData.ABI

// AccountManager is an auto generated Go binding around an Ethereum contract.
type AccountManager struct {
	AccountManagerCaller     // Read-only binding to the contract
	AccountManagerTransactor // Write-only binding to the contract
	AccountManagerFilterer   // Log filterer for contract events
}

// AccountManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountManagerSession struct {
	Contract     *AccountManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccountManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountManagerCallerSession struct {
	Contract *AccountManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AccountManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountManagerTransactorSession struct {
	Contract     *AccountManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AccountManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountManagerRaw struct {
	Contract *AccountManager // Generic contract binding to access the raw methods on
}

// AccountManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountManagerCallerRaw struct {
	Contract *AccountManagerCaller // Generic read-only contract binding to access the raw methods on
}

// AccountManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountManagerTransactorRaw struct {
	Contract *AccountManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccountManager creates a new instance of AccountManager, bound to a specific deployed contract.
func NewAccountManager(address common.Address, backend bind.ContractBackend) (*AccountManager, error) {
	contract, err := bindAccountManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccountManager{AccountManagerCaller: AccountManagerCaller{contract: contract}, AccountManagerTransactor: AccountManagerTransactor{contract: contract}, AccountManagerFilterer: AccountManagerFilterer{contract: contract}}, nil
}

// NewAccountManagerCaller creates a new read-only instance of AccountManager, bound to a specific deployed contract.
func NewAccountManagerCaller(address common.Address, caller bind.ContractCaller) (*AccountManagerCaller, error) {
	contract, err := bindAccountManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountManagerCaller{contract: contract}, nil
}

// NewAccountManagerTransactor creates a new write-only instance of AccountManager, bound to a specific deployed contract.
func NewAccountManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountManagerTransactor, error) {
	contract, err := bindAccountManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountManagerTransactor{contract: contract}, nil
}

// NewAccountManagerFilterer creates a new log filterer instance of AccountManager, bound to a specific deployed contract.
func NewAccountManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountManagerFilterer, error) {
	contract, err := bindAccountManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountManagerFilterer{contract: contract}, nil
}

// bindAccountManager binds a generic wrapper to an already deployed contract.
func bindAccountManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AccountManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountManager *AccountManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccountManager.Contract.AccountManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountManager *AccountManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountManager.Contract.AccountManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountManager *AccountManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountManager.Contract.AccountManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountManager *AccountManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccountManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountManager *AccountManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountManager *AccountManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountManager.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_AccountManager *AccountManagerCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AccountManager.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_AccountManager *AccountManagerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _AccountManager.Contract.UPGRADEINTERFACEVERSION(&_AccountManager.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_AccountManager *AccountManagerCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _AccountManager.Contract.UPGRADEINTERFACEVERSION(&_AccountManager.CallOpts)
}

// AccountImplementation is a free data retrieval call binding the contract method 0x11464fbe.
//
// Solidity: function accountImplementation() view returns(address)
func (_AccountManager *AccountManagerCaller) AccountImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccountManager.contract.Call(opts, &out, "accountImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountImplementation is a free data retrieval call binding the contract method 0x11464fbe.
//
// Solidity: function accountImplementation() view returns(address)
func (_AccountManager *AccountManagerSession) AccountImplementation() (common.Address, error) {
	return _AccountManager.Contract.AccountImplementation(&_AccountManager.CallOpts)
}

// AccountImplementation is a free data retrieval call binding the contract method 0x11464fbe.
//
// Solidity: function accountImplementation() view returns(address)
func (_AccountManager *AccountManagerCallerSession) AccountImplementation() (common.Address, error) {
	return _AccountManager.Contract.AccountImplementation(&_AccountManager.CallOpts)
}

// DefaultMasterSigner is a free data retrieval call binding the contract method 0x1137a196.
//
// Solidity: function defaultMasterSigner() view returns(address)
func (_AccountManager *AccountManagerCaller) DefaultMasterSigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccountManager.contract.Call(opts, &out, "defaultMasterSigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultMasterSigner is a free data retrieval call binding the contract method 0x1137a196.
//
// Solidity: function defaultMasterSigner() view returns(address)
func (_AccountManager *AccountManagerSession) DefaultMasterSigner() (common.Address, error) {
	return _AccountManager.Contract.DefaultMasterSigner(&_AccountManager.CallOpts)
}

// DefaultMasterSigner is a free data retrieval call binding the contract method 0x1137a196.
//
// Solidity: function defaultMasterSigner() view returns(address)
func (_AccountManager *AccountManagerCallerSession) DefaultMasterSigner() (common.Address, error) {
	return _AccountManager.Contract.DefaultMasterSigner(&_AccountManager.CallOpts)
}

// EntryPointAddress is a free data retrieval call binding the contract method 0x06dc245c.
//
// Solidity: function entryPointAddress() view returns(address)
func (_AccountManager *AccountManagerCaller) EntryPointAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccountManager.contract.Call(opts, &out, "entryPointAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPointAddress is a free data retrieval call binding the contract method 0x06dc245c.
//
// Solidity: function entryPointAddress() view returns(address)
func (_AccountManager *AccountManagerSession) EntryPointAddress() (common.Address, error) {
	return _AccountManager.Contract.EntryPointAddress(&_AccountManager.CallOpts)
}

// EntryPointAddress is a free data retrieval call binding the contract method 0x06dc245c.
//
// Solidity: function entryPointAddress() view returns(address)
func (_AccountManager *AccountManagerCallerSession) EntryPointAddress() (common.Address, error) {
	return _AccountManager.Contract.EntryPointAddress(&_AccountManager.CallOpts)
}

// GetAccountAddress is a free data retrieval call binding the contract method 0x3c505342.
//
// Solidity: function getAccountAddress(address owner, address masterSigner) view returns(address account)
func (_AccountManager *AccountManagerCaller) GetAccountAddress(opts *bind.CallOpts, owner common.Address, masterSigner common.Address) (common.Address, error) {
	var out []interface{}
	err := _AccountManager.contract.Call(opts, &out, "getAccountAddress", owner, masterSigner)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAccountAddress is a free data retrieval call binding the contract method 0x3c505342.
//
// Solidity: function getAccountAddress(address owner, address masterSigner) view returns(address account)
func (_AccountManager *AccountManagerSession) GetAccountAddress(owner common.Address, masterSigner common.Address) (common.Address, error) {
	return _AccountManager.Contract.GetAccountAddress(&_AccountManager.CallOpts, owner, masterSigner)
}

// GetAccountAddress is a free data retrieval call binding the contract method 0x3c505342.
//
// Solidity: function getAccountAddress(address owner, address masterSigner) view returns(address account)
func (_AccountManager *AccountManagerCallerSession) GetAccountAddress(owner common.Address, masterSigner common.Address) (common.Address, error) {
	return _AccountManager.Contract.GetAccountAddress(&_AccountManager.CallOpts, owner, masterSigner)
}

// GetDefaultMasterSigner is a free data retrieval call binding the contract method 0xa4cd40a2.
//
// Solidity: function getDefaultMasterSigner() view returns(address)
func (_AccountManager *AccountManagerCaller) GetDefaultMasterSigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccountManager.contract.Call(opts, &out, "getDefaultMasterSigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDefaultMasterSigner is a free data retrieval call binding the contract method 0xa4cd40a2.
//
// Solidity: function getDefaultMasterSigner() view returns(address)
func (_AccountManager *AccountManagerSession) GetDefaultMasterSigner() (common.Address, error) {
	return _AccountManager.Contract.GetDefaultMasterSigner(&_AccountManager.CallOpts)
}

// GetDefaultMasterSigner is a free data retrieval call binding the contract method 0xa4cd40a2.
//
// Solidity: function getDefaultMasterSigner() view returns(address)
func (_AccountManager *AccountManagerCallerSession) GetDefaultMasterSigner() (common.Address, error) {
	return _AccountManager.Contract.GetDefaultMasterSigner(&_AccountManager.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_AccountManager *AccountManagerCaller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccountManager.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_AccountManager *AccountManagerSession) GetImplementation() (common.Address, error) {
	return _AccountManager.Contract.GetImplementation(&_AccountManager.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_AccountManager *AccountManagerCallerSession) GetImplementation() (common.Address, error) {
	return _AccountManager.Contract.GetImplementation(&_AccountManager.CallOpts)
}

// GetInitCode is a free data retrieval call binding the contract method 0xfa5c8528.
//
// Solidity: function getInitCode(address owner, address masterSigner) view returns(bytes initCode)
func (_AccountManager *AccountManagerCaller) GetInitCode(opts *bind.CallOpts, owner common.Address, masterSigner common.Address) ([]byte, error) {
	var out []interface{}
	err := _AccountManager.contract.Call(opts, &out, "getInitCode", owner, masterSigner)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInitCode is a free data retrieval call binding the contract method 0xfa5c8528.
//
// Solidity: function getInitCode(address owner, address masterSigner) view returns(bytes initCode)
func (_AccountManager *AccountManagerSession) GetInitCode(owner common.Address, masterSigner common.Address) ([]byte, error) {
	return _AccountManager.Contract.GetInitCode(&_AccountManager.CallOpts, owner, masterSigner)
}

// GetInitCode is a free data retrieval call binding the contract method 0xfa5c8528.
//
// Solidity: function getInitCode(address owner, address masterSigner) view returns(bytes initCode)
func (_AccountManager *AccountManagerCallerSession) GetInitCode(owner common.Address, masterSigner common.Address) ([]byte, error) {
	return _AccountManager.Contract.GetInitCode(&_AccountManager.CallOpts, owner, masterSigner)
}

// IsAccountDeployed is a free data retrieval call binding the contract method 0x4797aa4c.
//
// Solidity: function isAccountDeployed(address owner, address masterSigner) view returns(bool isDeployed)
func (_AccountManager *AccountManagerCaller) IsAccountDeployed(opts *bind.CallOpts, owner common.Address, masterSigner common.Address) (bool, error) {
	var out []interface{}
	err := _AccountManager.contract.Call(opts, &out, "isAccountDeployed", owner, masterSigner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAccountDeployed is a free data retrieval call binding the contract method 0x4797aa4c.
//
// Solidity: function isAccountDeployed(address owner, address masterSigner) view returns(bool isDeployed)
func (_AccountManager *AccountManagerSession) IsAccountDeployed(owner common.Address, masterSigner common.Address) (bool, error) {
	return _AccountManager.Contract.IsAccountDeployed(&_AccountManager.CallOpts, owner, masterSigner)
}

// IsAccountDeployed is a free data retrieval call binding the contract method 0x4797aa4c.
//
// Solidity: function isAccountDeployed(address owner, address masterSigner) view returns(bool isDeployed)
func (_AccountManager *AccountManagerCallerSession) IsAccountDeployed(owner common.Address, masterSigner common.Address) (bool, error) {
	return _AccountManager.Contract.IsAccountDeployed(&_AccountManager.CallOpts, owner, masterSigner)
}

// IsAuthorizedCreator is a free data retrieval call binding the contract method 0x79b53ebe.
//
// Solidity: function isAuthorizedCreator(address creator) view returns(bool)
func (_AccountManager *AccountManagerCaller) IsAuthorizedCreator(opts *bind.CallOpts, creator common.Address) (bool, error) {
	var out []interface{}
	err := _AccountManager.contract.Call(opts, &out, "isAuthorizedCreator", creator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuthorizedCreator is a free data retrieval call binding the contract method 0x79b53ebe.
//
// Solidity: function isAuthorizedCreator(address creator) view returns(bool)
func (_AccountManager *AccountManagerSession) IsAuthorizedCreator(creator common.Address) (bool, error) {
	return _AccountManager.Contract.IsAuthorizedCreator(&_AccountManager.CallOpts, creator)
}

// IsAuthorizedCreator is a free data retrieval call binding the contract method 0x79b53ebe.
//
// Solidity: function isAuthorizedCreator(address creator) view returns(bool)
func (_AccountManager *AccountManagerCallerSession) IsAuthorizedCreator(creator common.Address) (bool, error) {
	return _AccountManager.Contract.IsAuthorizedCreator(&_AccountManager.CallOpts, creator)
}

// MasterAggregatorAddress is a free data retrieval call binding the contract method 0x482aab42.
//
// Solidity: function masterAggregatorAddress() view returns(address)
func (_AccountManager *AccountManagerCaller) MasterAggregatorAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccountManager.contract.Call(opts, &out, "masterAggregatorAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MasterAggregatorAddress is a free data retrieval call binding the contract method 0x482aab42.
//
// Solidity: function masterAggregatorAddress() view returns(address)
func (_AccountManager *AccountManagerSession) MasterAggregatorAddress() (common.Address, error) {
	return _AccountManager.Contract.MasterAggregatorAddress(&_AccountManager.CallOpts)
}

// MasterAggregatorAddress is a free data retrieval call binding the contract method 0x482aab42.
//
// Solidity: function masterAggregatorAddress() view returns(address)
func (_AccountManager *AccountManagerCallerSession) MasterAggregatorAddress() (common.Address, error) {
	return _AccountManager.Contract.MasterAggregatorAddress(&_AccountManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AccountManager *AccountManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccountManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AccountManager *AccountManagerSession) Owner() (common.Address, error) {
	return _AccountManager.Contract.Owner(&_AccountManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AccountManager *AccountManagerCallerSession) Owner() (common.Address, error) {
	return _AccountManager.Contract.Owner(&_AccountManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AccountManager *AccountManagerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccountManager.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AccountManager *AccountManagerSession) ProxiableUUID() ([32]byte, error) {
	return _AccountManager.Contract.ProxiableUUID(&_AccountManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AccountManager *AccountManagerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _AccountManager.Contract.ProxiableUUID(&_AccountManager.CallOpts)
}

// SenderCreator is a free data retrieval call binding the contract method 0x09ccb880.
//
// Solidity: function senderCreator() view returns(address)
func (_AccountManager *AccountManagerCaller) SenderCreator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccountManager.contract.Call(opts, &out, "senderCreator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SenderCreator is a free data retrieval call binding the contract method 0x09ccb880.
//
// Solidity: function senderCreator() view returns(address)
func (_AccountManager *AccountManagerSession) SenderCreator() (common.Address, error) {
	return _AccountManager.Contract.SenderCreator(&_AccountManager.CallOpts)
}

// SenderCreator is a free data retrieval call binding the contract method 0x09ccb880.
//
// Solidity: function senderCreator() view returns(address)
func (_AccountManager *AccountManagerCallerSession) SenderCreator() (common.Address, error) {
	return _AccountManager.Contract.SenderCreator(&_AccountManager.CallOpts)
}

// AuthorizeCreator is a paid mutator transaction binding the contract method 0xd8625b07.
//
// Solidity: function authorizeCreator(address creator) returns()
func (_AccountManager *AccountManagerTransactor) AuthorizeCreator(opts *bind.TransactOpts, creator common.Address) (*types.Transaction, error) {
	return _AccountManager.contract.Transact(opts, "authorizeCreator", creator)
}

// AuthorizeCreator is a paid mutator transaction binding the contract method 0xd8625b07.
//
// Solidity: function authorizeCreator(address creator) returns()
func (_AccountManager *AccountManagerSession) AuthorizeCreator(creator common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.AuthorizeCreator(&_AccountManager.TransactOpts, creator)
}

// AuthorizeCreator is a paid mutator transaction binding the contract method 0xd8625b07.
//
// Solidity: function authorizeCreator(address creator) returns()
func (_AccountManager *AccountManagerTransactorSession) AuthorizeCreator(creator common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.AuthorizeCreator(&_AccountManager.TransactOpts, creator)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x133572ba.
//
// Solidity: function createAccount(address owner, address masterSigner) returns(address account)
func (_AccountManager *AccountManagerTransactor) CreateAccount(opts *bind.TransactOpts, owner common.Address, masterSigner common.Address) (*types.Transaction, error) {
	return _AccountManager.contract.Transact(opts, "createAccount", owner, masterSigner)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x133572ba.
//
// Solidity: function createAccount(address owner, address masterSigner) returns(address account)
func (_AccountManager *AccountManagerSession) CreateAccount(owner common.Address, masterSigner common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.CreateAccount(&_AccountManager.TransactOpts, owner, masterSigner)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x133572ba.
//
// Solidity: function createAccount(address owner, address masterSigner) returns(address account)
func (_AccountManager *AccountManagerTransactorSession) CreateAccount(owner common.Address, masterSigner common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.CreateAccount(&_AccountManager.TransactOpts, owner, masterSigner)
}

// CreateAccountWithMasterSigner is a paid mutator transaction binding the contract method 0x4242339f.
//
// Solidity: function createAccountWithMasterSigner(address owner, address masterSigner) returns(address account)
func (_AccountManager *AccountManagerTransactor) CreateAccountWithMasterSigner(opts *bind.TransactOpts, owner common.Address, masterSigner common.Address) (*types.Transaction, error) {
	return _AccountManager.contract.Transact(opts, "createAccountWithMasterSigner", owner, masterSigner)
}

// CreateAccountWithMasterSigner is a paid mutator transaction binding the contract method 0x4242339f.
//
// Solidity: function createAccountWithMasterSigner(address owner, address masterSigner) returns(address account)
func (_AccountManager *AccountManagerSession) CreateAccountWithMasterSigner(owner common.Address, masterSigner common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.CreateAccountWithMasterSigner(&_AccountManager.TransactOpts, owner, masterSigner)
}

// CreateAccountWithMasterSigner is a paid mutator transaction binding the contract method 0x4242339f.
//
// Solidity: function createAccountWithMasterSigner(address owner, address masterSigner) returns(address account)
func (_AccountManager *AccountManagerTransactorSession) CreateAccountWithMasterSigner(owner common.Address, masterSigner common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.CreateAccountWithMasterSigner(&_AccountManager.TransactOpts, owner, masterSigner)
}

// CreateUserAccount is a paid mutator transaction binding the contract method 0x64857d26.
//
// Solidity: function createUserAccount(address owner, address masterSigner) returns(address account)
func (_AccountManager *AccountManagerTransactor) CreateUserAccount(opts *bind.TransactOpts, owner common.Address, masterSigner common.Address) (*types.Transaction, error) {
	return _AccountManager.contract.Transact(opts, "createUserAccount", owner, masterSigner)
}

// CreateUserAccount is a paid mutator transaction binding the contract method 0x64857d26.
//
// Solidity: function createUserAccount(address owner, address masterSigner) returns(address account)
func (_AccountManager *AccountManagerSession) CreateUserAccount(owner common.Address, masterSigner common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.CreateUserAccount(&_AccountManager.TransactOpts, owner, masterSigner)
}

// CreateUserAccount is a paid mutator transaction binding the contract method 0x64857d26.
//
// Solidity: function createUserAccount(address owner, address masterSigner) returns(address account)
func (_AccountManager *AccountManagerTransactorSession) CreateUserAccount(owner common.Address, masterSigner common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.CreateUserAccount(&_AccountManager.TransactOpts, owner, masterSigner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address entryPoint, address owner) returns()
func (_AccountManager *AccountManagerTransactor) Initialize(opts *bind.TransactOpts, entryPoint common.Address, owner common.Address) (*types.Transaction, error) {
	return _AccountManager.contract.Transact(opts, "initialize", entryPoint, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address entryPoint, address owner) returns()
func (_AccountManager *AccountManagerSession) Initialize(entryPoint common.Address, owner common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.Initialize(&_AccountManager.TransactOpts, entryPoint, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address entryPoint, address owner) returns()
func (_AccountManager *AccountManagerTransactorSession) Initialize(entryPoint common.Address, owner common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.Initialize(&_AccountManager.TransactOpts, entryPoint, owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AccountManager *AccountManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AccountManager *AccountManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _AccountManager.Contract.RenounceOwnership(&_AccountManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AccountManager *AccountManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AccountManager.Contract.RenounceOwnership(&_AccountManager.TransactOpts)
}

// RevokeCreator is a paid mutator transaction binding the contract method 0xc0e69af6.
//
// Solidity: function revokeCreator(address creator) returns()
func (_AccountManager *AccountManagerTransactor) RevokeCreator(opts *bind.TransactOpts, creator common.Address) (*types.Transaction, error) {
	return _AccountManager.contract.Transact(opts, "revokeCreator", creator)
}

// RevokeCreator is a paid mutator transaction binding the contract method 0xc0e69af6.
//
// Solidity: function revokeCreator(address creator) returns()
func (_AccountManager *AccountManagerSession) RevokeCreator(creator common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.RevokeCreator(&_AccountManager.TransactOpts, creator)
}

// RevokeCreator is a paid mutator transaction binding the contract method 0xc0e69af6.
//
// Solidity: function revokeCreator(address creator) returns()
func (_AccountManager *AccountManagerTransactorSession) RevokeCreator(creator common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.RevokeCreator(&_AccountManager.TransactOpts, creator)
}

// SetDefaultMasterSigner is a paid mutator transaction binding the contract method 0xa3a2bb76.
//
// Solidity: function setDefaultMasterSigner(address masterSigner) returns()
func (_AccountManager *AccountManagerTransactor) SetDefaultMasterSigner(opts *bind.TransactOpts, masterSigner common.Address) (*types.Transaction, error) {
	return _AccountManager.contract.Transact(opts, "setDefaultMasterSigner", masterSigner)
}

// SetDefaultMasterSigner is a paid mutator transaction binding the contract method 0xa3a2bb76.
//
// Solidity: function setDefaultMasterSigner(address masterSigner) returns()
func (_AccountManager *AccountManagerSession) SetDefaultMasterSigner(masterSigner common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.SetDefaultMasterSigner(&_AccountManager.TransactOpts, masterSigner)
}

// SetDefaultMasterSigner is a paid mutator transaction binding the contract method 0xa3a2bb76.
//
// Solidity: function setDefaultMasterSigner(address masterSigner) returns()
func (_AccountManager *AccountManagerTransactorSession) SetDefaultMasterSigner(masterSigner common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.SetDefaultMasterSigner(&_AccountManager.TransactOpts, masterSigner)
}

// SetMasterAggregator is a paid mutator transaction binding the contract method 0x67d944fd.
//
// Solidity: function setMasterAggregator(address aggregator) returns()
func (_AccountManager *AccountManagerTransactor) SetMasterAggregator(opts *bind.TransactOpts, aggregator common.Address) (*types.Transaction, error) {
	return _AccountManager.contract.Transact(opts, "setMasterAggregator", aggregator)
}

// SetMasterAggregator is a paid mutator transaction binding the contract method 0x67d944fd.
//
// Solidity: function setMasterAggregator(address aggregator) returns()
func (_AccountManager *AccountManagerSession) SetMasterAggregator(aggregator common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.SetMasterAggregator(&_AccountManager.TransactOpts, aggregator)
}

// SetMasterAggregator is a paid mutator transaction binding the contract method 0x67d944fd.
//
// Solidity: function setMasterAggregator(address aggregator) returns()
func (_AccountManager *AccountManagerTransactorSession) SetMasterAggregator(aggregator common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.SetMasterAggregator(&_AccountManager.TransactOpts, aggregator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AccountManager *AccountManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AccountManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AccountManager *AccountManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.TransferOwnership(&_AccountManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AccountManager *AccountManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.TransferOwnership(&_AccountManager.TransactOpts, newOwner)
}

// UpdateAccountImplementation is a paid mutator transaction binding the contract method 0x0a5b9e6e.
//
// Solidity: function updateAccountImplementation(address newImplementation) returns()
func (_AccountManager *AccountManagerTransactor) UpdateAccountImplementation(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _AccountManager.contract.Transact(opts, "updateAccountImplementation", newImplementation)
}

// UpdateAccountImplementation is a paid mutator transaction binding the contract method 0x0a5b9e6e.
//
// Solidity: function updateAccountImplementation(address newImplementation) returns()
func (_AccountManager *AccountManagerSession) UpdateAccountImplementation(newImplementation common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.UpdateAccountImplementation(&_AccountManager.TransactOpts, newImplementation)
}

// UpdateAccountImplementation is a paid mutator transaction binding the contract method 0x0a5b9e6e.
//
// Solidity: function updateAccountImplementation(address newImplementation) returns()
func (_AccountManager *AccountManagerTransactorSession) UpdateAccountImplementation(newImplementation common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.UpdateAccountImplementation(&_AccountManager.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AccountManager *AccountManagerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AccountManager.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AccountManager *AccountManagerSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AccountManager.Contract.UpgradeToAndCall(&_AccountManager.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AccountManager *AccountManagerTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AccountManager.Contract.UpgradeToAndCall(&_AccountManager.TransactOpts, newImplementation, data)
}

// AccountManagerAccountCreatedIterator is returned from FilterAccountCreated and is used to iterate over the raw logs and unpacked data for AccountCreated events raised by the AccountManager contract.
type AccountManagerAccountCreatedIterator struct {
	Event *AccountManagerAccountCreated // Event containing the contract specifics and raw log

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
func (it *AccountManagerAccountCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountManagerAccountCreated)
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
		it.Event = new(AccountManagerAccountCreated)
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
func (it *AccountManagerAccountCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountManagerAccountCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountManagerAccountCreated represents a AccountCreated event raised by the AccountManager contract.
type AccountManagerAccountCreated struct {
	Account      common.Address
	Owner        common.Address
	MasterSigner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAccountCreated is a free log retrieval operation binding the contract event 0x9db37ffcb3d1eadb0306e35c0453e35a675f2e80757b09251d848185ac5cc22a.
//
// Solidity: event AccountCreated(address indexed account, address indexed owner, address indexed masterSigner)
func (_AccountManager *AccountManagerFilterer) FilterAccountCreated(opts *bind.FilterOpts, account []common.Address, owner []common.Address, masterSigner []common.Address) (*AccountManagerAccountCreatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var masterSignerRule []interface{}
	for _, masterSignerItem := range masterSigner {
		masterSignerRule = append(masterSignerRule, masterSignerItem)
	}

	logs, sub, err := _AccountManager.contract.FilterLogs(opts, "AccountCreated", accountRule, ownerRule, masterSignerRule)
	if err != nil {
		return nil, err
	}
	return &AccountManagerAccountCreatedIterator{contract: _AccountManager.contract, event: "AccountCreated", logs: logs, sub: sub}, nil
}

// WatchAccountCreated is a free log subscription operation binding the contract event 0x9db37ffcb3d1eadb0306e35c0453e35a675f2e80757b09251d848185ac5cc22a.
//
// Solidity: event AccountCreated(address indexed account, address indexed owner, address indexed masterSigner)
func (_AccountManager *AccountManagerFilterer) WatchAccountCreated(opts *bind.WatchOpts, sink chan<- *AccountManagerAccountCreated, account []common.Address, owner []common.Address, masterSigner []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var masterSignerRule []interface{}
	for _, masterSignerItem := range masterSigner {
		masterSignerRule = append(masterSignerRule, masterSignerItem)
	}

	logs, sub, err := _AccountManager.contract.WatchLogs(opts, "AccountCreated", accountRule, ownerRule, masterSignerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountManagerAccountCreated)
				if err := _AccountManager.contract.UnpackLog(event, "AccountCreated", log); err != nil {
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

// ParseAccountCreated is a log parse operation binding the contract event 0x9db37ffcb3d1eadb0306e35c0453e35a675f2e80757b09251d848185ac5cc22a.
//
// Solidity: event AccountCreated(address indexed account, address indexed owner, address indexed masterSigner)
func (_AccountManager *AccountManagerFilterer) ParseAccountCreated(log types.Log) (*AccountManagerAccountCreated, error) {
	event := new(AccountManagerAccountCreated)
	if err := _AccountManager.contract.UnpackLog(event, "AccountCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountManagerCreatorAuthorizedIterator is returned from FilterCreatorAuthorized and is used to iterate over the raw logs and unpacked data for CreatorAuthorized events raised by the AccountManager contract.
type AccountManagerCreatorAuthorizedIterator struct {
	Event *AccountManagerCreatorAuthorized // Event containing the contract specifics and raw log

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
func (it *AccountManagerCreatorAuthorizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountManagerCreatorAuthorized)
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
		it.Event = new(AccountManagerCreatorAuthorized)
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
func (it *AccountManagerCreatorAuthorizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountManagerCreatorAuthorizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountManagerCreatorAuthorized represents a CreatorAuthorized event raised by the AccountManager contract.
type AccountManagerCreatorAuthorized struct {
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCreatorAuthorized is a free log retrieval operation binding the contract event 0x00b6d3c5059e85c667aa67cfa572e42162916972ed621a4618e55dfcb4a17463.
//
// Solidity: event CreatorAuthorized(address indexed creator)
func (_AccountManager *AccountManagerFilterer) FilterCreatorAuthorized(opts *bind.FilterOpts, creator []common.Address) (*AccountManagerCreatorAuthorizedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _AccountManager.contract.FilterLogs(opts, "CreatorAuthorized", creatorRule)
	if err != nil {
		return nil, err
	}
	return &AccountManagerCreatorAuthorizedIterator{contract: _AccountManager.contract, event: "CreatorAuthorized", logs: logs, sub: sub}, nil
}

// WatchCreatorAuthorized is a free log subscription operation binding the contract event 0x00b6d3c5059e85c667aa67cfa572e42162916972ed621a4618e55dfcb4a17463.
//
// Solidity: event CreatorAuthorized(address indexed creator)
func (_AccountManager *AccountManagerFilterer) WatchCreatorAuthorized(opts *bind.WatchOpts, sink chan<- *AccountManagerCreatorAuthorized, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _AccountManager.contract.WatchLogs(opts, "CreatorAuthorized", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountManagerCreatorAuthorized)
				if err := _AccountManager.contract.UnpackLog(event, "CreatorAuthorized", log); err != nil {
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

// ParseCreatorAuthorized is a log parse operation binding the contract event 0x00b6d3c5059e85c667aa67cfa572e42162916972ed621a4618e55dfcb4a17463.
//
// Solidity: event CreatorAuthorized(address indexed creator)
func (_AccountManager *AccountManagerFilterer) ParseCreatorAuthorized(log types.Log) (*AccountManagerCreatorAuthorized, error) {
	event := new(AccountManagerCreatorAuthorized)
	if err := _AccountManager.contract.UnpackLog(event, "CreatorAuthorized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountManagerCreatorRevokedIterator is returned from FilterCreatorRevoked and is used to iterate over the raw logs and unpacked data for CreatorRevoked events raised by the AccountManager contract.
type AccountManagerCreatorRevokedIterator struct {
	Event *AccountManagerCreatorRevoked // Event containing the contract specifics and raw log

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
func (it *AccountManagerCreatorRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountManagerCreatorRevoked)
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
		it.Event = new(AccountManagerCreatorRevoked)
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
func (it *AccountManagerCreatorRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountManagerCreatorRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountManagerCreatorRevoked represents a CreatorRevoked event raised by the AccountManager contract.
type AccountManagerCreatorRevoked struct {
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCreatorRevoked is a free log retrieval operation binding the contract event 0xc77d51d04162a5d61b26217b8a8c78870a02749a7a198336685c57d6f128fa4b.
//
// Solidity: event CreatorRevoked(address indexed creator)
func (_AccountManager *AccountManagerFilterer) FilterCreatorRevoked(opts *bind.FilterOpts, creator []common.Address) (*AccountManagerCreatorRevokedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _AccountManager.contract.FilterLogs(opts, "CreatorRevoked", creatorRule)
	if err != nil {
		return nil, err
	}
	return &AccountManagerCreatorRevokedIterator{contract: _AccountManager.contract, event: "CreatorRevoked", logs: logs, sub: sub}, nil
}

// WatchCreatorRevoked is a free log subscription operation binding the contract event 0xc77d51d04162a5d61b26217b8a8c78870a02749a7a198336685c57d6f128fa4b.
//
// Solidity: event CreatorRevoked(address indexed creator)
func (_AccountManager *AccountManagerFilterer) WatchCreatorRevoked(opts *bind.WatchOpts, sink chan<- *AccountManagerCreatorRevoked, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _AccountManager.contract.WatchLogs(opts, "CreatorRevoked", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountManagerCreatorRevoked)
				if err := _AccountManager.contract.UnpackLog(event, "CreatorRevoked", log); err != nil {
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

// ParseCreatorRevoked is a log parse operation binding the contract event 0xc77d51d04162a5d61b26217b8a8c78870a02749a7a198336685c57d6f128fa4b.
//
// Solidity: event CreatorRevoked(address indexed creator)
func (_AccountManager *AccountManagerFilterer) ParseCreatorRevoked(log types.Log) (*AccountManagerCreatorRevoked, error) {
	event := new(AccountManagerCreatorRevoked)
	if err := _AccountManager.contract.UnpackLog(event, "CreatorRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AccountManager contract.
type AccountManagerInitializedIterator struct {
	Event *AccountManagerInitialized // Event containing the contract specifics and raw log

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
func (it *AccountManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountManagerInitialized)
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
		it.Event = new(AccountManagerInitialized)
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
func (it *AccountManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountManagerInitialized represents a Initialized event raised by the AccountManager contract.
type AccountManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AccountManager *AccountManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*AccountManagerInitializedIterator, error) {

	logs, sub, err := _AccountManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AccountManagerInitializedIterator{contract: _AccountManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AccountManager *AccountManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AccountManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _AccountManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountManagerInitialized)
				if err := _AccountManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AccountManager *AccountManagerFilterer) ParseInitialized(log types.Log) (*AccountManagerInitialized, error) {
	event := new(AccountManagerInitialized)
	if err := _AccountManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AccountManager contract.
type AccountManagerOwnershipTransferredIterator struct {
	Event *AccountManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AccountManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountManagerOwnershipTransferred)
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
		it.Event = new(AccountManagerOwnershipTransferred)
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
func (it *AccountManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountManagerOwnershipTransferred represents a OwnershipTransferred event raised by the AccountManager contract.
type AccountManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AccountManager *AccountManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AccountManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AccountManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AccountManagerOwnershipTransferredIterator{contract: _AccountManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AccountManager *AccountManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AccountManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AccountManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountManagerOwnershipTransferred)
				if err := _AccountManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AccountManager *AccountManagerFilterer) ParseOwnershipTransferred(log types.Log) (*AccountManagerOwnershipTransferred, error) {
	event := new(AccountManagerOwnershipTransferred)
	if err := _AccountManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountManagerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the AccountManager contract.
type AccountManagerUpgradedIterator struct {
	Event *AccountManagerUpgraded // Event containing the contract specifics and raw log

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
func (it *AccountManagerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountManagerUpgraded)
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
		it.Event = new(AccountManagerUpgraded)
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
func (it *AccountManagerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountManagerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountManagerUpgraded represents a Upgraded event raised by the AccountManager contract.
type AccountManagerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AccountManager *AccountManagerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*AccountManagerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AccountManager.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &AccountManagerUpgradedIterator{contract: _AccountManager.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AccountManager *AccountManagerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *AccountManagerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AccountManager.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountManagerUpgraded)
				if err := _AccountManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AccountManager *AccountManagerFilterer) ParseUpgraded(log types.Log) (*AccountManagerUpgraded, error) {
	event := new(AccountManagerUpgraded)
	if err := _AccountManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
