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

// ISessionKeyManagerBatchSessionKeyOp is an auto generated low-level Go binding around an user-defined struct.
type ISessionKeyManagerBatchSessionKeyOp struct {
	Account     common.Address
	SessionKey  common.Address
	ValidAfter  *big.Int
	ValidUntil  *big.Int
	Permissions [32]byte
}

// ISessionKeyManagerSessionKeyTemplate is an auto generated low-level Go binding around an user-defined struct.
type ISessionKeyManagerSessionKeyTemplate struct {
	Name            string
	DefaultDuration *big.Int
	Permissions     [32]byte
	IsActive        bool
}

// SessionKeyManagerMetaData contains all meta data concerning the SessionKeyManager contract.
var SessionKeyManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"masterSigner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AccountRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"masterSigner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AccountUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"masterSigner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"operationCount\",\"type\":\"uint256\"}],\"name\":\"BatchSessionKeysAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"masterSigner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"operationCount\",\"type\":\"uint256\"}],\"name\":\"BatchSessionKeysRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"masterSigner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"authorized\",\"type\":\"bool\"}],\"name\":\"MasterSignerAuthorized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"duration\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"permissions\",\"type\":\"bytes32\"}],\"name\":\"SessionKeyTemplateCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"masterSigner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"templateName\",\"type\":\"string\"},{\"internalType\":\"uint48\",\"name\":\"customDuration\",\"type\":\"uint48\"}],\"name\":\"addSessionKeyWithTemplate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"masterSigner\",\"type\":\"address\"}],\"name\":\"authorizeMasterSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorizedMasters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"validAfter\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"validUntil\",\"type\":\"uint48\"},{\"internalType\":\"bytes32\",\"name\":\"permissions\",\"type\":\"bytes32\"}],\"internalType\":\"structISessionKeyManager.BatchSessionKeyOp[]\",\"name\":\"operations\",\"type\":\"tuple[]\"}],\"name\":\"batchAddSessionKeys\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"}],\"name\":\"batchRevokeSessionKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint48\",\"name\":\"defaultDuration\",\"type\":\"uint48\"},{\"internalType\":\"bytes32\",\"name\":\"permissions\",\"type\":\"bytes32\"}],\"name\":\"createSessionKeyTemplate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"masterSigner\",\"type\":\"address\"}],\"name\":\"getRegisteredAccounts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"templateName\",\"type\":\"string\"}],\"name\":\"getTemplate\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint48\",\"name\":\"defaultDuration\",\"type\":\"uint48\"},{\"internalType\":\"bytes32\",\"name\":\"permissions\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structISessionKeyManager.SessionKeyTemplate\",\"name\":\"template\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTemplateNames\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"names\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"masterSignerAccounts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"masterSigner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"registerAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"masterSigner\",\"type\":\"address\"}],\"name\":\"revokeMasterSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"sessionKeyTemplates\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint48\",\"name\":\"defaultDuration\",\"type\":\"uint48\"},{\"internalType\":\"bytes32\",\"name\":\"permissions\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"templateNames\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"masterSigner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"unregisterAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// SessionKeyManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use SessionKeyManagerMetaData.ABI instead.
var SessionKeyManagerABI = SessionKeyManagerMetaData.ABI

// SessionKeyManager is an auto generated Go binding around an Ethereum contract.
type SessionKeyManager struct {
	SessionKeyManagerCaller     // Read-only binding to the contract
	SessionKeyManagerTransactor // Write-only binding to the contract
	SessionKeyManagerFilterer   // Log filterer for contract events
}

// SessionKeyManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type SessionKeyManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SessionKeyManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SessionKeyManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SessionKeyManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SessionKeyManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SessionKeyManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SessionKeyManagerSession struct {
	Contract     *SessionKeyManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SessionKeyManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SessionKeyManagerCallerSession struct {
	Contract *SessionKeyManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// SessionKeyManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SessionKeyManagerTransactorSession struct {
	Contract     *SessionKeyManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// SessionKeyManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type SessionKeyManagerRaw struct {
	Contract *SessionKeyManager // Generic contract binding to access the raw methods on
}

// SessionKeyManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SessionKeyManagerCallerRaw struct {
	Contract *SessionKeyManagerCaller // Generic read-only contract binding to access the raw methods on
}

// SessionKeyManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SessionKeyManagerTransactorRaw struct {
	Contract *SessionKeyManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSessionKeyManager creates a new instance of SessionKeyManager, bound to a specific deployed contract.
func NewSessionKeyManager(address common.Address, backend bind.ContractBackend) (*SessionKeyManager, error) {
	contract, err := bindSessionKeyManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SessionKeyManager{SessionKeyManagerCaller: SessionKeyManagerCaller{contract: contract}, SessionKeyManagerTransactor: SessionKeyManagerTransactor{contract: contract}, SessionKeyManagerFilterer: SessionKeyManagerFilterer{contract: contract}}, nil
}

// NewSessionKeyManagerCaller creates a new read-only instance of SessionKeyManager, bound to a specific deployed contract.
func NewSessionKeyManagerCaller(address common.Address, caller bind.ContractCaller) (*SessionKeyManagerCaller, error) {
	contract, err := bindSessionKeyManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SessionKeyManagerCaller{contract: contract}, nil
}

// NewSessionKeyManagerTransactor creates a new write-only instance of SessionKeyManager, bound to a specific deployed contract.
func NewSessionKeyManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*SessionKeyManagerTransactor, error) {
	contract, err := bindSessionKeyManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SessionKeyManagerTransactor{contract: contract}, nil
}

// NewSessionKeyManagerFilterer creates a new log filterer instance of SessionKeyManager, bound to a specific deployed contract.
func NewSessionKeyManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*SessionKeyManagerFilterer, error) {
	contract, err := bindSessionKeyManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SessionKeyManagerFilterer{contract: contract}, nil
}

// bindSessionKeyManager binds a generic wrapper to an already deployed contract.
func bindSessionKeyManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SessionKeyManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SessionKeyManager *SessionKeyManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SessionKeyManager.Contract.SessionKeyManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SessionKeyManager *SessionKeyManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SessionKeyManager.Contract.SessionKeyManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SessionKeyManager *SessionKeyManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SessionKeyManager.Contract.SessionKeyManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SessionKeyManager *SessionKeyManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SessionKeyManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SessionKeyManager *SessionKeyManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SessionKeyManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SessionKeyManager *SessionKeyManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SessionKeyManager.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_SessionKeyManager *SessionKeyManagerCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SessionKeyManager.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_SessionKeyManager *SessionKeyManagerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _SessionKeyManager.Contract.UPGRADEINTERFACEVERSION(&_SessionKeyManager.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_SessionKeyManager *SessionKeyManagerCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _SessionKeyManager.Contract.UPGRADEINTERFACEVERSION(&_SessionKeyManager.CallOpts)
}

// AuthorizedMasters is a free data retrieval call binding the contract method 0xfe0221fe.
//
// Solidity: function authorizedMasters(address ) view returns(bool)
func (_SessionKeyManager *SessionKeyManagerCaller) AuthorizedMasters(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _SessionKeyManager.contract.Call(opts, &out, "authorizedMasters", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AuthorizedMasters is a free data retrieval call binding the contract method 0xfe0221fe.
//
// Solidity: function authorizedMasters(address ) view returns(bool)
func (_SessionKeyManager *SessionKeyManagerSession) AuthorizedMasters(arg0 common.Address) (bool, error) {
	return _SessionKeyManager.Contract.AuthorizedMasters(&_SessionKeyManager.CallOpts, arg0)
}

// AuthorizedMasters is a free data retrieval call binding the contract method 0xfe0221fe.
//
// Solidity: function authorizedMasters(address ) view returns(bool)
func (_SessionKeyManager *SessionKeyManagerCallerSession) AuthorizedMasters(arg0 common.Address) (bool, error) {
	return _SessionKeyManager.Contract.AuthorizedMasters(&_SessionKeyManager.CallOpts, arg0)
}

// GetRegisteredAccounts is a free data retrieval call binding the contract method 0xe8f880e1.
//
// Solidity: function getRegisteredAccounts(address masterSigner) view returns(address[] accounts)
func (_SessionKeyManager *SessionKeyManagerCaller) GetRegisteredAccounts(opts *bind.CallOpts, masterSigner common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _SessionKeyManager.contract.Call(opts, &out, "getRegisteredAccounts", masterSigner)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRegisteredAccounts is a free data retrieval call binding the contract method 0xe8f880e1.
//
// Solidity: function getRegisteredAccounts(address masterSigner) view returns(address[] accounts)
func (_SessionKeyManager *SessionKeyManagerSession) GetRegisteredAccounts(masterSigner common.Address) ([]common.Address, error) {
	return _SessionKeyManager.Contract.GetRegisteredAccounts(&_SessionKeyManager.CallOpts, masterSigner)
}

// GetRegisteredAccounts is a free data retrieval call binding the contract method 0xe8f880e1.
//
// Solidity: function getRegisteredAccounts(address masterSigner) view returns(address[] accounts)
func (_SessionKeyManager *SessionKeyManagerCallerSession) GetRegisteredAccounts(masterSigner common.Address) ([]common.Address, error) {
	return _SessionKeyManager.Contract.GetRegisteredAccounts(&_SessionKeyManager.CallOpts, masterSigner)
}

// GetTemplate is a free data retrieval call binding the contract method 0x7d67cd6f.
//
// Solidity: function getTemplate(string templateName) view returns((string,uint48,bytes32,bool) template)
func (_SessionKeyManager *SessionKeyManagerCaller) GetTemplate(opts *bind.CallOpts, templateName string) (ISessionKeyManagerSessionKeyTemplate, error) {
	var out []interface{}
	err := _SessionKeyManager.contract.Call(opts, &out, "getTemplate", templateName)

	if err != nil {
		return *new(ISessionKeyManagerSessionKeyTemplate), err
	}

	out0 := *abi.ConvertType(out[0], new(ISessionKeyManagerSessionKeyTemplate)).(*ISessionKeyManagerSessionKeyTemplate)

	return out0, err

}

// GetTemplate is a free data retrieval call binding the contract method 0x7d67cd6f.
//
// Solidity: function getTemplate(string templateName) view returns((string,uint48,bytes32,bool) template)
func (_SessionKeyManager *SessionKeyManagerSession) GetTemplate(templateName string) (ISessionKeyManagerSessionKeyTemplate, error) {
	return _SessionKeyManager.Contract.GetTemplate(&_SessionKeyManager.CallOpts, templateName)
}

// GetTemplate is a free data retrieval call binding the contract method 0x7d67cd6f.
//
// Solidity: function getTemplate(string templateName) view returns((string,uint48,bytes32,bool) template)
func (_SessionKeyManager *SessionKeyManagerCallerSession) GetTemplate(templateName string) (ISessionKeyManagerSessionKeyTemplate, error) {
	return _SessionKeyManager.Contract.GetTemplate(&_SessionKeyManager.CallOpts, templateName)
}

// GetTemplateNames is a free data retrieval call binding the contract method 0xf5d5d0b0.
//
// Solidity: function getTemplateNames() view returns(string[] names)
func (_SessionKeyManager *SessionKeyManagerCaller) GetTemplateNames(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _SessionKeyManager.contract.Call(opts, &out, "getTemplateNames")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetTemplateNames is a free data retrieval call binding the contract method 0xf5d5d0b0.
//
// Solidity: function getTemplateNames() view returns(string[] names)
func (_SessionKeyManager *SessionKeyManagerSession) GetTemplateNames() ([]string, error) {
	return _SessionKeyManager.Contract.GetTemplateNames(&_SessionKeyManager.CallOpts)
}

// GetTemplateNames is a free data retrieval call binding the contract method 0xf5d5d0b0.
//
// Solidity: function getTemplateNames() view returns(string[] names)
func (_SessionKeyManager *SessionKeyManagerCallerSession) GetTemplateNames() ([]string, error) {
	return _SessionKeyManager.Contract.GetTemplateNames(&_SessionKeyManager.CallOpts)
}

// MasterSignerAccounts is a free data retrieval call binding the contract method 0x42b526e7.
//
// Solidity: function masterSignerAccounts(address , uint256 ) view returns(address)
func (_SessionKeyManager *SessionKeyManagerCaller) MasterSignerAccounts(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SessionKeyManager.contract.Call(opts, &out, "masterSignerAccounts", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MasterSignerAccounts is a free data retrieval call binding the contract method 0x42b526e7.
//
// Solidity: function masterSignerAccounts(address , uint256 ) view returns(address)
func (_SessionKeyManager *SessionKeyManagerSession) MasterSignerAccounts(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _SessionKeyManager.Contract.MasterSignerAccounts(&_SessionKeyManager.CallOpts, arg0, arg1)
}

// MasterSignerAccounts is a free data retrieval call binding the contract method 0x42b526e7.
//
// Solidity: function masterSignerAccounts(address , uint256 ) view returns(address)
func (_SessionKeyManager *SessionKeyManagerCallerSession) MasterSignerAccounts(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _SessionKeyManager.Contract.MasterSignerAccounts(&_SessionKeyManager.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SessionKeyManager *SessionKeyManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SessionKeyManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SessionKeyManager *SessionKeyManagerSession) Owner() (common.Address, error) {
	return _SessionKeyManager.Contract.Owner(&_SessionKeyManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SessionKeyManager *SessionKeyManagerCallerSession) Owner() (common.Address, error) {
	return _SessionKeyManager.Contract.Owner(&_SessionKeyManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SessionKeyManager *SessionKeyManagerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SessionKeyManager.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SessionKeyManager *SessionKeyManagerSession) ProxiableUUID() ([32]byte, error) {
	return _SessionKeyManager.Contract.ProxiableUUID(&_SessionKeyManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SessionKeyManager *SessionKeyManagerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _SessionKeyManager.Contract.ProxiableUUID(&_SessionKeyManager.CallOpts)
}

// SessionKeyTemplates is a free data retrieval call binding the contract method 0x789f4ead.
//
// Solidity: function sessionKeyTemplates(string ) view returns(string name, uint48 defaultDuration, bytes32 permissions, bool isActive)
func (_SessionKeyManager *SessionKeyManagerCaller) SessionKeyTemplates(opts *bind.CallOpts, arg0 string) (struct {
	Name            string
	DefaultDuration *big.Int
	Permissions     [32]byte
	IsActive        bool
}, error) {
	var out []interface{}
	err := _SessionKeyManager.contract.Call(opts, &out, "sessionKeyTemplates", arg0)

	outstruct := new(struct {
		Name            string
		DefaultDuration *big.Int
		Permissions     [32]byte
		IsActive        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.DefaultDuration = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Permissions = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.IsActive = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// SessionKeyTemplates is a free data retrieval call binding the contract method 0x789f4ead.
//
// Solidity: function sessionKeyTemplates(string ) view returns(string name, uint48 defaultDuration, bytes32 permissions, bool isActive)
func (_SessionKeyManager *SessionKeyManagerSession) SessionKeyTemplates(arg0 string) (struct {
	Name            string
	DefaultDuration *big.Int
	Permissions     [32]byte
	IsActive        bool
}, error) {
	return _SessionKeyManager.Contract.SessionKeyTemplates(&_SessionKeyManager.CallOpts, arg0)
}

// SessionKeyTemplates is a free data retrieval call binding the contract method 0x789f4ead.
//
// Solidity: function sessionKeyTemplates(string ) view returns(string name, uint48 defaultDuration, bytes32 permissions, bool isActive)
func (_SessionKeyManager *SessionKeyManagerCallerSession) SessionKeyTemplates(arg0 string) (struct {
	Name            string
	DefaultDuration *big.Int
	Permissions     [32]byte
	IsActive        bool
}, error) {
	return _SessionKeyManager.Contract.SessionKeyTemplates(&_SessionKeyManager.CallOpts, arg0)
}

// TemplateNames is a free data retrieval call binding the contract method 0xe842abe8.
//
// Solidity: function templateNames(uint256 ) view returns(string)
func (_SessionKeyManager *SessionKeyManagerCaller) TemplateNames(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _SessionKeyManager.contract.Call(opts, &out, "templateNames", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TemplateNames is a free data retrieval call binding the contract method 0xe842abe8.
//
// Solidity: function templateNames(uint256 ) view returns(string)
func (_SessionKeyManager *SessionKeyManagerSession) TemplateNames(arg0 *big.Int) (string, error) {
	return _SessionKeyManager.Contract.TemplateNames(&_SessionKeyManager.CallOpts, arg0)
}

// TemplateNames is a free data retrieval call binding the contract method 0xe842abe8.
//
// Solidity: function templateNames(uint256 ) view returns(string)
func (_SessionKeyManager *SessionKeyManagerCallerSession) TemplateNames(arg0 *big.Int) (string, error) {
	return _SessionKeyManager.Contract.TemplateNames(&_SessionKeyManager.CallOpts, arg0)
}

// AddSessionKeyWithTemplate is a paid mutator transaction binding the contract method 0x06381952.
//
// Solidity: function addSessionKeyWithTemplate(address masterSigner, address sessionKey, string templateName, uint48 customDuration) returns()
func (_SessionKeyManager *SessionKeyManagerTransactor) AddSessionKeyWithTemplate(opts *bind.TransactOpts, masterSigner common.Address, sessionKey common.Address, templateName string, customDuration *big.Int) (*types.Transaction, error) {
	return _SessionKeyManager.contract.Transact(opts, "addSessionKeyWithTemplate", masterSigner, sessionKey, templateName, customDuration)
}

// AddSessionKeyWithTemplate is a paid mutator transaction binding the contract method 0x06381952.
//
// Solidity: function addSessionKeyWithTemplate(address masterSigner, address sessionKey, string templateName, uint48 customDuration) returns()
func (_SessionKeyManager *SessionKeyManagerSession) AddSessionKeyWithTemplate(masterSigner common.Address, sessionKey common.Address, templateName string, customDuration *big.Int) (*types.Transaction, error) {
	return _SessionKeyManager.Contract.AddSessionKeyWithTemplate(&_SessionKeyManager.TransactOpts, masterSigner, sessionKey, templateName, customDuration)
}

// AddSessionKeyWithTemplate is a paid mutator transaction binding the contract method 0x06381952.
//
// Solidity: function addSessionKeyWithTemplate(address masterSigner, address sessionKey, string templateName, uint48 customDuration) returns()
func (_SessionKeyManager *SessionKeyManagerTransactorSession) AddSessionKeyWithTemplate(masterSigner common.Address, sessionKey common.Address, templateName string, customDuration *big.Int) (*types.Transaction, error) {
	return _SessionKeyManager.Contract.AddSessionKeyWithTemplate(&_SessionKeyManager.TransactOpts, masterSigner, sessionKey, templateName, customDuration)
}

// AuthorizeMasterSigner is a paid mutator transaction binding the contract method 0x941f058d.
//
// Solidity: function authorizeMasterSigner(address masterSigner) returns()
func (_SessionKeyManager *SessionKeyManagerTransactor) AuthorizeMasterSigner(opts *bind.TransactOpts, masterSigner common.Address) (*types.Transaction, error) {
	return _SessionKeyManager.contract.Transact(opts, "authorizeMasterSigner", masterSigner)
}

// AuthorizeMasterSigner is a paid mutator transaction binding the contract method 0x941f058d.
//
// Solidity: function authorizeMasterSigner(address masterSigner) returns()
func (_SessionKeyManager *SessionKeyManagerSession) AuthorizeMasterSigner(masterSigner common.Address) (*types.Transaction, error) {
	return _SessionKeyManager.Contract.AuthorizeMasterSigner(&_SessionKeyManager.TransactOpts, masterSigner)
}

// AuthorizeMasterSigner is a paid mutator transaction binding the contract method 0x941f058d.
//
// Solidity: function authorizeMasterSigner(address masterSigner) returns()
func (_SessionKeyManager *SessionKeyManagerTransactorSession) AuthorizeMasterSigner(masterSigner common.Address) (*types.Transaction, error) {
	return _SessionKeyManager.Contract.AuthorizeMasterSigner(&_SessionKeyManager.TransactOpts, masterSigner)
}

// BatchAddSessionKeys is a paid mutator transaction binding the contract method 0x83fc8724.
//
// Solidity: function batchAddSessionKeys((address,address,uint48,uint48,bytes32)[] operations) returns()
func (_SessionKeyManager *SessionKeyManagerTransactor) BatchAddSessionKeys(opts *bind.TransactOpts, operations []ISessionKeyManagerBatchSessionKeyOp) (*types.Transaction, error) {
	return _SessionKeyManager.contract.Transact(opts, "batchAddSessionKeys", operations)
}

// BatchAddSessionKeys is a paid mutator transaction binding the contract method 0x83fc8724.
//
// Solidity: function batchAddSessionKeys((address,address,uint48,uint48,bytes32)[] operations) returns()
func (_SessionKeyManager *SessionKeyManagerSession) BatchAddSessionKeys(operations []ISessionKeyManagerBatchSessionKeyOp) (*types.Transaction, error) {
	return _SessionKeyManager.Contract.BatchAddSessionKeys(&_SessionKeyManager.TransactOpts, operations)
}

// BatchAddSessionKeys is a paid mutator transaction binding the contract method 0x83fc8724.
//
// Solidity: function batchAddSessionKeys((address,address,uint48,uint48,bytes32)[] operations) returns()
func (_SessionKeyManager *SessionKeyManagerTransactorSession) BatchAddSessionKeys(operations []ISessionKeyManagerBatchSessionKeyOp) (*types.Transaction, error) {
	return _SessionKeyManager.Contract.BatchAddSessionKeys(&_SessionKeyManager.TransactOpts, operations)
}

// BatchRevokeSessionKey is a paid mutator transaction binding the contract method 0xbb5763dd.
//
// Solidity: function batchRevokeSessionKey(address sessionKey) returns()
func (_SessionKeyManager *SessionKeyManagerTransactor) BatchRevokeSessionKey(opts *bind.TransactOpts, sessionKey common.Address) (*types.Transaction, error) {
	return _SessionKeyManager.contract.Transact(opts, "batchRevokeSessionKey", sessionKey)
}

// BatchRevokeSessionKey is a paid mutator transaction binding the contract method 0xbb5763dd.
//
// Solidity: function batchRevokeSessionKey(address sessionKey) returns()
func (_SessionKeyManager *SessionKeyManagerSession) BatchRevokeSessionKey(sessionKey common.Address) (*types.Transaction, error) {
	return _SessionKeyManager.Contract.BatchRevokeSessionKey(&_SessionKeyManager.TransactOpts, sessionKey)
}

// BatchRevokeSessionKey is a paid mutator transaction binding the contract method 0xbb5763dd.
//
// Solidity: function batchRevokeSessionKey(address sessionKey) returns()
func (_SessionKeyManager *SessionKeyManagerTransactorSession) BatchRevokeSessionKey(sessionKey common.Address) (*types.Transaction, error) {
	return _SessionKeyManager.Contract.BatchRevokeSessionKey(&_SessionKeyManager.TransactOpts, sessionKey)
}

// CreateSessionKeyTemplate is a paid mutator transaction binding the contract method 0xb6d076e7.
//
// Solidity: function createSessionKeyTemplate(string name, uint48 defaultDuration, bytes32 permissions) returns()
func (_SessionKeyManager *SessionKeyManagerTransactor) CreateSessionKeyTemplate(opts *bind.TransactOpts, name string, defaultDuration *big.Int, permissions [32]byte) (*types.Transaction, error) {
	return _SessionKeyManager.contract.Transact(opts, "createSessionKeyTemplate", name, defaultDuration, permissions)
}

// CreateSessionKeyTemplate is a paid mutator transaction binding the contract method 0xb6d076e7.
//
// Solidity: function createSessionKeyTemplate(string name, uint48 defaultDuration, bytes32 permissions) returns()
func (_SessionKeyManager *SessionKeyManagerSession) CreateSessionKeyTemplate(name string, defaultDuration *big.Int, permissions [32]byte) (*types.Transaction, error) {
	return _SessionKeyManager.Contract.CreateSessionKeyTemplate(&_SessionKeyManager.TransactOpts, name, defaultDuration, permissions)
}

// CreateSessionKeyTemplate is a paid mutator transaction binding the contract method 0xb6d076e7.
//
// Solidity: function createSessionKeyTemplate(string name, uint48 defaultDuration, bytes32 permissions) returns()
func (_SessionKeyManager *SessionKeyManagerTransactorSession) CreateSessionKeyTemplate(name string, defaultDuration *big.Int, permissions [32]byte) (*types.Transaction, error) {
	return _SessionKeyManager.Contract.CreateSessionKeyTemplate(&_SessionKeyManager.TransactOpts, name, defaultDuration, permissions)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_SessionKeyManager *SessionKeyManagerTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _SessionKeyManager.contract.Transact(opts, "initialize", _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_SessionKeyManager *SessionKeyManagerSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _SessionKeyManager.Contract.Initialize(&_SessionKeyManager.TransactOpts, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_SessionKeyManager *SessionKeyManagerTransactorSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _SessionKeyManager.Contract.Initialize(&_SessionKeyManager.TransactOpts, _owner)
}

// RegisterAccount is a paid mutator transaction binding the contract method 0x71fe98bc.
//
// Solidity: function registerAccount(address masterSigner, address account) returns()
func (_SessionKeyManager *SessionKeyManagerTransactor) RegisterAccount(opts *bind.TransactOpts, masterSigner common.Address, account common.Address) (*types.Transaction, error) {
	return _SessionKeyManager.contract.Transact(opts, "registerAccount", masterSigner, account)
}

// RegisterAccount is a paid mutator transaction binding the contract method 0x71fe98bc.
//
// Solidity: function registerAccount(address masterSigner, address account) returns()
func (_SessionKeyManager *SessionKeyManagerSession) RegisterAccount(masterSigner common.Address, account common.Address) (*types.Transaction, error) {
	return _SessionKeyManager.Contract.RegisterAccount(&_SessionKeyManager.TransactOpts, masterSigner, account)
}

// RegisterAccount is a paid mutator transaction binding the contract method 0x71fe98bc.
//
// Solidity: function registerAccount(address masterSigner, address account) returns()
func (_SessionKeyManager *SessionKeyManagerTransactorSession) RegisterAccount(masterSigner common.Address, account common.Address) (*types.Transaction, error) {
	return _SessionKeyManager.Contract.RegisterAccount(&_SessionKeyManager.TransactOpts, masterSigner, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SessionKeyManager *SessionKeyManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SessionKeyManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SessionKeyManager *SessionKeyManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _SessionKeyManager.Contract.RenounceOwnership(&_SessionKeyManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SessionKeyManager *SessionKeyManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SessionKeyManager.Contract.RenounceOwnership(&_SessionKeyManager.TransactOpts)
}

// RevokeMasterSigner is a paid mutator transaction binding the contract method 0x136620b1.
//
// Solidity: function revokeMasterSigner(address masterSigner) returns()
func (_SessionKeyManager *SessionKeyManagerTransactor) RevokeMasterSigner(opts *bind.TransactOpts, masterSigner common.Address) (*types.Transaction, error) {
	return _SessionKeyManager.contract.Transact(opts, "revokeMasterSigner", masterSigner)
}

// RevokeMasterSigner is a paid mutator transaction binding the contract method 0x136620b1.
//
// Solidity: function revokeMasterSigner(address masterSigner) returns()
func (_SessionKeyManager *SessionKeyManagerSession) RevokeMasterSigner(masterSigner common.Address) (*types.Transaction, error) {
	return _SessionKeyManager.Contract.RevokeMasterSigner(&_SessionKeyManager.TransactOpts, masterSigner)
}

// RevokeMasterSigner is a paid mutator transaction binding the contract method 0x136620b1.
//
// Solidity: function revokeMasterSigner(address masterSigner) returns()
func (_SessionKeyManager *SessionKeyManagerTransactorSession) RevokeMasterSigner(masterSigner common.Address) (*types.Transaction, error) {
	return _SessionKeyManager.Contract.RevokeMasterSigner(&_SessionKeyManager.TransactOpts, masterSigner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SessionKeyManager *SessionKeyManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SessionKeyManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SessionKeyManager *SessionKeyManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SessionKeyManager.Contract.TransferOwnership(&_SessionKeyManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SessionKeyManager *SessionKeyManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SessionKeyManager.Contract.TransferOwnership(&_SessionKeyManager.TransactOpts, newOwner)
}

// UnregisterAccount is a paid mutator transaction binding the contract method 0xe05e2477.
//
// Solidity: function unregisterAccount(address masterSigner, address account) returns()
func (_SessionKeyManager *SessionKeyManagerTransactor) UnregisterAccount(opts *bind.TransactOpts, masterSigner common.Address, account common.Address) (*types.Transaction, error) {
	return _SessionKeyManager.contract.Transact(opts, "unregisterAccount", masterSigner, account)
}

// UnregisterAccount is a paid mutator transaction binding the contract method 0xe05e2477.
//
// Solidity: function unregisterAccount(address masterSigner, address account) returns()
func (_SessionKeyManager *SessionKeyManagerSession) UnregisterAccount(masterSigner common.Address, account common.Address) (*types.Transaction, error) {
	return _SessionKeyManager.Contract.UnregisterAccount(&_SessionKeyManager.TransactOpts, masterSigner, account)
}

// UnregisterAccount is a paid mutator transaction binding the contract method 0xe05e2477.
//
// Solidity: function unregisterAccount(address masterSigner, address account) returns()
func (_SessionKeyManager *SessionKeyManagerTransactorSession) UnregisterAccount(masterSigner common.Address, account common.Address) (*types.Transaction, error) {
	return _SessionKeyManager.Contract.UnregisterAccount(&_SessionKeyManager.TransactOpts, masterSigner, account)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SessionKeyManager *SessionKeyManagerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SessionKeyManager.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SessionKeyManager *SessionKeyManagerSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SessionKeyManager.Contract.UpgradeToAndCall(&_SessionKeyManager.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SessionKeyManager *SessionKeyManagerTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SessionKeyManager.Contract.UpgradeToAndCall(&_SessionKeyManager.TransactOpts, newImplementation, data)
}

// SessionKeyManagerAccountRegisteredIterator is returned from FilterAccountRegistered and is used to iterate over the raw logs and unpacked data for AccountRegistered events raised by the SessionKeyManager contract.
type SessionKeyManagerAccountRegisteredIterator struct {
	Event *SessionKeyManagerAccountRegistered // Event containing the contract specifics and raw log

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
func (it *SessionKeyManagerAccountRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SessionKeyManagerAccountRegistered)
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
		it.Event = new(SessionKeyManagerAccountRegistered)
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
func (it *SessionKeyManagerAccountRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SessionKeyManagerAccountRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SessionKeyManagerAccountRegistered represents a AccountRegistered event raised by the SessionKeyManager contract.
type SessionKeyManagerAccountRegistered struct {
	MasterSigner common.Address
	Account      common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAccountRegistered is a free log retrieval operation binding the contract event 0x415c9c2b518f5b14d4a1dd06e9ea39351fbcf788a8000a8726eb2ad48fca781f.
//
// Solidity: event AccountRegistered(address indexed masterSigner, address indexed account)
func (_SessionKeyManager *SessionKeyManagerFilterer) FilterAccountRegistered(opts *bind.FilterOpts, masterSigner []common.Address, account []common.Address) (*SessionKeyManagerAccountRegisteredIterator, error) {

	var masterSignerRule []interface{}
	for _, masterSignerItem := range masterSigner {
		masterSignerRule = append(masterSignerRule, masterSignerItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SessionKeyManager.contract.FilterLogs(opts, "AccountRegistered", masterSignerRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &SessionKeyManagerAccountRegisteredIterator{contract: _SessionKeyManager.contract, event: "AccountRegistered", logs: logs, sub: sub}, nil
}

// WatchAccountRegistered is a free log subscription operation binding the contract event 0x415c9c2b518f5b14d4a1dd06e9ea39351fbcf788a8000a8726eb2ad48fca781f.
//
// Solidity: event AccountRegistered(address indexed masterSigner, address indexed account)
func (_SessionKeyManager *SessionKeyManagerFilterer) WatchAccountRegistered(opts *bind.WatchOpts, sink chan<- *SessionKeyManagerAccountRegistered, masterSigner []common.Address, account []common.Address) (event.Subscription, error) {

	var masterSignerRule []interface{}
	for _, masterSignerItem := range masterSigner {
		masterSignerRule = append(masterSignerRule, masterSignerItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SessionKeyManager.contract.WatchLogs(opts, "AccountRegistered", masterSignerRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SessionKeyManagerAccountRegistered)
				if err := _SessionKeyManager.contract.UnpackLog(event, "AccountRegistered", log); err != nil {
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

// ParseAccountRegistered is a log parse operation binding the contract event 0x415c9c2b518f5b14d4a1dd06e9ea39351fbcf788a8000a8726eb2ad48fca781f.
//
// Solidity: event AccountRegistered(address indexed masterSigner, address indexed account)
func (_SessionKeyManager *SessionKeyManagerFilterer) ParseAccountRegistered(log types.Log) (*SessionKeyManagerAccountRegistered, error) {
	event := new(SessionKeyManagerAccountRegistered)
	if err := _SessionKeyManager.contract.UnpackLog(event, "AccountRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SessionKeyManagerAccountUnregisteredIterator is returned from FilterAccountUnregistered and is used to iterate over the raw logs and unpacked data for AccountUnregistered events raised by the SessionKeyManager contract.
type SessionKeyManagerAccountUnregisteredIterator struct {
	Event *SessionKeyManagerAccountUnregistered // Event containing the contract specifics and raw log

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
func (it *SessionKeyManagerAccountUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SessionKeyManagerAccountUnregistered)
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
		it.Event = new(SessionKeyManagerAccountUnregistered)
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
func (it *SessionKeyManagerAccountUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SessionKeyManagerAccountUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SessionKeyManagerAccountUnregistered represents a AccountUnregistered event raised by the SessionKeyManager contract.
type SessionKeyManagerAccountUnregistered struct {
	MasterSigner common.Address
	Account      common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAccountUnregistered is a free log retrieval operation binding the contract event 0x2c64115beb9454d3020b88d9f9abfe70572db3fbc3d4f9df0afd55e759bb2aff.
//
// Solidity: event AccountUnregistered(address indexed masterSigner, address indexed account)
func (_SessionKeyManager *SessionKeyManagerFilterer) FilterAccountUnregistered(opts *bind.FilterOpts, masterSigner []common.Address, account []common.Address) (*SessionKeyManagerAccountUnregisteredIterator, error) {

	var masterSignerRule []interface{}
	for _, masterSignerItem := range masterSigner {
		masterSignerRule = append(masterSignerRule, masterSignerItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SessionKeyManager.contract.FilterLogs(opts, "AccountUnregistered", masterSignerRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &SessionKeyManagerAccountUnregisteredIterator{contract: _SessionKeyManager.contract, event: "AccountUnregistered", logs: logs, sub: sub}, nil
}

// WatchAccountUnregistered is a free log subscription operation binding the contract event 0x2c64115beb9454d3020b88d9f9abfe70572db3fbc3d4f9df0afd55e759bb2aff.
//
// Solidity: event AccountUnregistered(address indexed masterSigner, address indexed account)
func (_SessionKeyManager *SessionKeyManagerFilterer) WatchAccountUnregistered(opts *bind.WatchOpts, sink chan<- *SessionKeyManagerAccountUnregistered, masterSigner []common.Address, account []common.Address) (event.Subscription, error) {

	var masterSignerRule []interface{}
	for _, masterSignerItem := range masterSigner {
		masterSignerRule = append(masterSignerRule, masterSignerItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SessionKeyManager.contract.WatchLogs(opts, "AccountUnregistered", masterSignerRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SessionKeyManagerAccountUnregistered)
				if err := _SessionKeyManager.contract.UnpackLog(event, "AccountUnregistered", log); err != nil {
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

// ParseAccountUnregistered is a log parse operation binding the contract event 0x2c64115beb9454d3020b88d9f9abfe70572db3fbc3d4f9df0afd55e759bb2aff.
//
// Solidity: event AccountUnregistered(address indexed masterSigner, address indexed account)
func (_SessionKeyManager *SessionKeyManagerFilterer) ParseAccountUnregistered(log types.Log) (*SessionKeyManagerAccountUnregistered, error) {
	event := new(SessionKeyManagerAccountUnregistered)
	if err := _SessionKeyManager.contract.UnpackLog(event, "AccountUnregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SessionKeyManagerBatchSessionKeysAddedIterator is returned from FilterBatchSessionKeysAdded and is used to iterate over the raw logs and unpacked data for BatchSessionKeysAdded events raised by the SessionKeyManager contract.
type SessionKeyManagerBatchSessionKeysAddedIterator struct {
	Event *SessionKeyManagerBatchSessionKeysAdded // Event containing the contract specifics and raw log

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
func (it *SessionKeyManagerBatchSessionKeysAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SessionKeyManagerBatchSessionKeysAdded)
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
		it.Event = new(SessionKeyManagerBatchSessionKeysAdded)
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
func (it *SessionKeyManagerBatchSessionKeysAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SessionKeyManagerBatchSessionKeysAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SessionKeyManagerBatchSessionKeysAdded represents a BatchSessionKeysAdded event raised by the SessionKeyManager contract.
type SessionKeyManagerBatchSessionKeysAdded struct {
	MasterSigner   common.Address
	OperationCount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBatchSessionKeysAdded is a free log retrieval operation binding the contract event 0xe186f599a84746ba9416e7f7db192ea291e818087b5e799899fce3cbdbf0b4dc.
//
// Solidity: event BatchSessionKeysAdded(address indexed masterSigner, uint256 operationCount)
func (_SessionKeyManager *SessionKeyManagerFilterer) FilterBatchSessionKeysAdded(opts *bind.FilterOpts, masterSigner []common.Address) (*SessionKeyManagerBatchSessionKeysAddedIterator, error) {

	var masterSignerRule []interface{}
	for _, masterSignerItem := range masterSigner {
		masterSignerRule = append(masterSignerRule, masterSignerItem)
	}

	logs, sub, err := _SessionKeyManager.contract.FilterLogs(opts, "BatchSessionKeysAdded", masterSignerRule)
	if err != nil {
		return nil, err
	}
	return &SessionKeyManagerBatchSessionKeysAddedIterator{contract: _SessionKeyManager.contract, event: "BatchSessionKeysAdded", logs: logs, sub: sub}, nil
}

// WatchBatchSessionKeysAdded is a free log subscription operation binding the contract event 0xe186f599a84746ba9416e7f7db192ea291e818087b5e799899fce3cbdbf0b4dc.
//
// Solidity: event BatchSessionKeysAdded(address indexed masterSigner, uint256 operationCount)
func (_SessionKeyManager *SessionKeyManagerFilterer) WatchBatchSessionKeysAdded(opts *bind.WatchOpts, sink chan<- *SessionKeyManagerBatchSessionKeysAdded, masterSigner []common.Address) (event.Subscription, error) {

	var masterSignerRule []interface{}
	for _, masterSignerItem := range masterSigner {
		masterSignerRule = append(masterSignerRule, masterSignerItem)
	}

	logs, sub, err := _SessionKeyManager.contract.WatchLogs(opts, "BatchSessionKeysAdded", masterSignerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SessionKeyManagerBatchSessionKeysAdded)
				if err := _SessionKeyManager.contract.UnpackLog(event, "BatchSessionKeysAdded", log); err != nil {
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

// ParseBatchSessionKeysAdded is a log parse operation binding the contract event 0xe186f599a84746ba9416e7f7db192ea291e818087b5e799899fce3cbdbf0b4dc.
//
// Solidity: event BatchSessionKeysAdded(address indexed masterSigner, uint256 operationCount)
func (_SessionKeyManager *SessionKeyManagerFilterer) ParseBatchSessionKeysAdded(log types.Log) (*SessionKeyManagerBatchSessionKeysAdded, error) {
	event := new(SessionKeyManagerBatchSessionKeysAdded)
	if err := _SessionKeyManager.contract.UnpackLog(event, "BatchSessionKeysAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SessionKeyManagerBatchSessionKeysRevokedIterator is returned from FilterBatchSessionKeysRevoked and is used to iterate over the raw logs and unpacked data for BatchSessionKeysRevoked events raised by the SessionKeyManager contract.
type SessionKeyManagerBatchSessionKeysRevokedIterator struct {
	Event *SessionKeyManagerBatchSessionKeysRevoked // Event containing the contract specifics and raw log

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
func (it *SessionKeyManagerBatchSessionKeysRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SessionKeyManagerBatchSessionKeysRevoked)
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
		it.Event = new(SessionKeyManagerBatchSessionKeysRevoked)
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
func (it *SessionKeyManagerBatchSessionKeysRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SessionKeyManagerBatchSessionKeysRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SessionKeyManagerBatchSessionKeysRevoked represents a BatchSessionKeysRevoked event raised by the SessionKeyManager contract.
type SessionKeyManagerBatchSessionKeysRevoked struct {
	MasterSigner   common.Address
	OperationCount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBatchSessionKeysRevoked is a free log retrieval operation binding the contract event 0x1c6bc406f9685ed6e86bd13e734c85a81bed930976063beec8ddefb94274bf4e.
//
// Solidity: event BatchSessionKeysRevoked(address indexed masterSigner, uint256 operationCount)
func (_SessionKeyManager *SessionKeyManagerFilterer) FilterBatchSessionKeysRevoked(opts *bind.FilterOpts, masterSigner []common.Address) (*SessionKeyManagerBatchSessionKeysRevokedIterator, error) {

	var masterSignerRule []interface{}
	for _, masterSignerItem := range masterSigner {
		masterSignerRule = append(masterSignerRule, masterSignerItem)
	}

	logs, sub, err := _SessionKeyManager.contract.FilterLogs(opts, "BatchSessionKeysRevoked", masterSignerRule)
	if err != nil {
		return nil, err
	}
	return &SessionKeyManagerBatchSessionKeysRevokedIterator{contract: _SessionKeyManager.contract, event: "BatchSessionKeysRevoked", logs: logs, sub: sub}, nil
}

// WatchBatchSessionKeysRevoked is a free log subscription operation binding the contract event 0x1c6bc406f9685ed6e86bd13e734c85a81bed930976063beec8ddefb94274bf4e.
//
// Solidity: event BatchSessionKeysRevoked(address indexed masterSigner, uint256 operationCount)
func (_SessionKeyManager *SessionKeyManagerFilterer) WatchBatchSessionKeysRevoked(opts *bind.WatchOpts, sink chan<- *SessionKeyManagerBatchSessionKeysRevoked, masterSigner []common.Address) (event.Subscription, error) {

	var masterSignerRule []interface{}
	for _, masterSignerItem := range masterSigner {
		masterSignerRule = append(masterSignerRule, masterSignerItem)
	}

	logs, sub, err := _SessionKeyManager.contract.WatchLogs(opts, "BatchSessionKeysRevoked", masterSignerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SessionKeyManagerBatchSessionKeysRevoked)
				if err := _SessionKeyManager.contract.UnpackLog(event, "BatchSessionKeysRevoked", log); err != nil {
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

// ParseBatchSessionKeysRevoked is a log parse operation binding the contract event 0x1c6bc406f9685ed6e86bd13e734c85a81bed930976063beec8ddefb94274bf4e.
//
// Solidity: event BatchSessionKeysRevoked(address indexed masterSigner, uint256 operationCount)
func (_SessionKeyManager *SessionKeyManagerFilterer) ParseBatchSessionKeysRevoked(log types.Log) (*SessionKeyManagerBatchSessionKeysRevoked, error) {
	event := new(SessionKeyManagerBatchSessionKeysRevoked)
	if err := _SessionKeyManager.contract.UnpackLog(event, "BatchSessionKeysRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SessionKeyManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SessionKeyManager contract.
type SessionKeyManagerInitializedIterator struct {
	Event *SessionKeyManagerInitialized // Event containing the contract specifics and raw log

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
func (it *SessionKeyManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SessionKeyManagerInitialized)
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
		it.Event = new(SessionKeyManagerInitialized)
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
func (it *SessionKeyManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SessionKeyManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SessionKeyManagerInitialized represents a Initialized event raised by the SessionKeyManager contract.
type SessionKeyManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SessionKeyManager *SessionKeyManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*SessionKeyManagerInitializedIterator, error) {

	logs, sub, err := _SessionKeyManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SessionKeyManagerInitializedIterator{contract: _SessionKeyManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SessionKeyManager *SessionKeyManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SessionKeyManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _SessionKeyManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SessionKeyManagerInitialized)
				if err := _SessionKeyManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SessionKeyManager *SessionKeyManagerFilterer) ParseInitialized(log types.Log) (*SessionKeyManagerInitialized, error) {
	event := new(SessionKeyManagerInitialized)
	if err := _SessionKeyManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SessionKeyManagerMasterSignerAuthorizedIterator is returned from FilterMasterSignerAuthorized and is used to iterate over the raw logs and unpacked data for MasterSignerAuthorized events raised by the SessionKeyManager contract.
type SessionKeyManagerMasterSignerAuthorizedIterator struct {
	Event *SessionKeyManagerMasterSignerAuthorized // Event containing the contract specifics and raw log

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
func (it *SessionKeyManagerMasterSignerAuthorizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SessionKeyManagerMasterSignerAuthorized)
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
		it.Event = new(SessionKeyManagerMasterSignerAuthorized)
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
func (it *SessionKeyManagerMasterSignerAuthorizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SessionKeyManagerMasterSignerAuthorizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SessionKeyManagerMasterSignerAuthorized represents a MasterSignerAuthorized event raised by the SessionKeyManager contract.
type SessionKeyManagerMasterSignerAuthorized struct {
	MasterSigner common.Address
	Authorized   bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMasterSignerAuthorized is a free log retrieval operation binding the contract event 0xf701060b30680725cd6015ad977083a38f59b5027e945f6c193060adda7630a8.
//
// Solidity: event MasterSignerAuthorized(address indexed masterSigner, bool authorized)
func (_SessionKeyManager *SessionKeyManagerFilterer) FilterMasterSignerAuthorized(opts *bind.FilterOpts, masterSigner []common.Address) (*SessionKeyManagerMasterSignerAuthorizedIterator, error) {

	var masterSignerRule []interface{}
	for _, masterSignerItem := range masterSigner {
		masterSignerRule = append(masterSignerRule, masterSignerItem)
	}

	logs, sub, err := _SessionKeyManager.contract.FilterLogs(opts, "MasterSignerAuthorized", masterSignerRule)
	if err != nil {
		return nil, err
	}
	return &SessionKeyManagerMasterSignerAuthorizedIterator{contract: _SessionKeyManager.contract, event: "MasterSignerAuthorized", logs: logs, sub: sub}, nil
}

// WatchMasterSignerAuthorized is a free log subscription operation binding the contract event 0xf701060b30680725cd6015ad977083a38f59b5027e945f6c193060adda7630a8.
//
// Solidity: event MasterSignerAuthorized(address indexed masterSigner, bool authorized)
func (_SessionKeyManager *SessionKeyManagerFilterer) WatchMasterSignerAuthorized(opts *bind.WatchOpts, sink chan<- *SessionKeyManagerMasterSignerAuthorized, masterSigner []common.Address) (event.Subscription, error) {

	var masterSignerRule []interface{}
	for _, masterSignerItem := range masterSigner {
		masterSignerRule = append(masterSignerRule, masterSignerItem)
	}

	logs, sub, err := _SessionKeyManager.contract.WatchLogs(opts, "MasterSignerAuthorized", masterSignerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SessionKeyManagerMasterSignerAuthorized)
				if err := _SessionKeyManager.contract.UnpackLog(event, "MasterSignerAuthorized", log); err != nil {
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

// ParseMasterSignerAuthorized is a log parse operation binding the contract event 0xf701060b30680725cd6015ad977083a38f59b5027e945f6c193060adda7630a8.
//
// Solidity: event MasterSignerAuthorized(address indexed masterSigner, bool authorized)
func (_SessionKeyManager *SessionKeyManagerFilterer) ParseMasterSignerAuthorized(log types.Log) (*SessionKeyManagerMasterSignerAuthorized, error) {
	event := new(SessionKeyManagerMasterSignerAuthorized)
	if err := _SessionKeyManager.contract.UnpackLog(event, "MasterSignerAuthorized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SessionKeyManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SessionKeyManager contract.
type SessionKeyManagerOwnershipTransferredIterator struct {
	Event *SessionKeyManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SessionKeyManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SessionKeyManagerOwnershipTransferred)
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
		it.Event = new(SessionKeyManagerOwnershipTransferred)
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
func (it *SessionKeyManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SessionKeyManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SessionKeyManagerOwnershipTransferred represents a OwnershipTransferred event raised by the SessionKeyManager contract.
type SessionKeyManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SessionKeyManager *SessionKeyManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SessionKeyManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SessionKeyManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SessionKeyManagerOwnershipTransferredIterator{contract: _SessionKeyManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SessionKeyManager *SessionKeyManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SessionKeyManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SessionKeyManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SessionKeyManagerOwnershipTransferred)
				if err := _SessionKeyManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SessionKeyManager *SessionKeyManagerFilterer) ParseOwnershipTransferred(log types.Log) (*SessionKeyManagerOwnershipTransferred, error) {
	event := new(SessionKeyManagerOwnershipTransferred)
	if err := _SessionKeyManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SessionKeyManagerSessionKeyTemplateCreatedIterator is returned from FilterSessionKeyTemplateCreated and is used to iterate over the raw logs and unpacked data for SessionKeyTemplateCreated events raised by the SessionKeyManager contract.
type SessionKeyManagerSessionKeyTemplateCreatedIterator struct {
	Event *SessionKeyManagerSessionKeyTemplateCreated // Event containing the contract specifics and raw log

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
func (it *SessionKeyManagerSessionKeyTemplateCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SessionKeyManagerSessionKeyTemplateCreated)
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
		it.Event = new(SessionKeyManagerSessionKeyTemplateCreated)
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
func (it *SessionKeyManagerSessionKeyTemplateCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SessionKeyManagerSessionKeyTemplateCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SessionKeyManagerSessionKeyTemplateCreated represents a SessionKeyTemplateCreated event raised by the SessionKeyManager contract.
type SessionKeyManagerSessionKeyTemplateCreated struct {
	Name        common.Hash
	Duration    *big.Int
	Permissions [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSessionKeyTemplateCreated is a free log retrieval operation binding the contract event 0x572b22ddc5d2a88148a7bbf6c9e22fb55f2b3e9aa0ea539f14f45b19bc12fe78.
//
// Solidity: event SessionKeyTemplateCreated(string indexed name, uint48 duration, bytes32 permissions)
func (_SessionKeyManager *SessionKeyManagerFilterer) FilterSessionKeyTemplateCreated(opts *bind.FilterOpts, name []string) (*SessionKeyManagerSessionKeyTemplateCreatedIterator, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _SessionKeyManager.contract.FilterLogs(opts, "SessionKeyTemplateCreated", nameRule)
	if err != nil {
		return nil, err
	}
	return &SessionKeyManagerSessionKeyTemplateCreatedIterator{contract: _SessionKeyManager.contract, event: "SessionKeyTemplateCreated", logs: logs, sub: sub}, nil
}

// WatchSessionKeyTemplateCreated is a free log subscription operation binding the contract event 0x572b22ddc5d2a88148a7bbf6c9e22fb55f2b3e9aa0ea539f14f45b19bc12fe78.
//
// Solidity: event SessionKeyTemplateCreated(string indexed name, uint48 duration, bytes32 permissions)
func (_SessionKeyManager *SessionKeyManagerFilterer) WatchSessionKeyTemplateCreated(opts *bind.WatchOpts, sink chan<- *SessionKeyManagerSessionKeyTemplateCreated, name []string) (event.Subscription, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _SessionKeyManager.contract.WatchLogs(opts, "SessionKeyTemplateCreated", nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SessionKeyManagerSessionKeyTemplateCreated)
				if err := _SessionKeyManager.contract.UnpackLog(event, "SessionKeyTemplateCreated", log); err != nil {
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

// ParseSessionKeyTemplateCreated is a log parse operation binding the contract event 0x572b22ddc5d2a88148a7bbf6c9e22fb55f2b3e9aa0ea539f14f45b19bc12fe78.
//
// Solidity: event SessionKeyTemplateCreated(string indexed name, uint48 duration, bytes32 permissions)
func (_SessionKeyManager *SessionKeyManagerFilterer) ParseSessionKeyTemplateCreated(log types.Log) (*SessionKeyManagerSessionKeyTemplateCreated, error) {
	event := new(SessionKeyManagerSessionKeyTemplateCreated)
	if err := _SessionKeyManager.contract.UnpackLog(event, "SessionKeyTemplateCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SessionKeyManagerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the SessionKeyManager contract.
type SessionKeyManagerUpgradedIterator struct {
	Event *SessionKeyManagerUpgraded // Event containing the contract specifics and raw log

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
func (it *SessionKeyManagerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SessionKeyManagerUpgraded)
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
		it.Event = new(SessionKeyManagerUpgraded)
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
func (it *SessionKeyManagerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SessionKeyManagerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SessionKeyManagerUpgraded represents a Upgraded event raised by the SessionKeyManager contract.
type SessionKeyManagerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SessionKeyManager *SessionKeyManagerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*SessionKeyManagerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _SessionKeyManager.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &SessionKeyManagerUpgradedIterator{contract: _SessionKeyManager.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SessionKeyManager *SessionKeyManagerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *SessionKeyManagerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _SessionKeyManager.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SessionKeyManagerUpgraded)
				if err := _SessionKeyManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_SessionKeyManager *SessionKeyManagerFilterer) ParseUpgraded(log types.Log) (*SessionKeyManagerUpgraded, error) {
	event := new(SessionKeyManagerUpgraded)
	if err := _SessionKeyManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
