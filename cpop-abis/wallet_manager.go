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

// WalletManagerMetaData contains all meta data concerning the WalletManager contract.
var WalletManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"masterSigner\",\"type\":\"address\"}],\"name\":\"AccountCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"CreatorAuthorized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"CreatorRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accountImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"authorizeCreator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"masterSigner\",\"type\":\"address\"}],\"name\":\"createAccountWithMasterSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"masterSigner\",\"type\":\"address\"}],\"name\":\"createUserAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"masterSigner\",\"type\":\"address\"}],\"name\":\"createWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultMasterSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPointAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"masterSigner\",\"type\":\"address\"}],\"name\":\"getAccountAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDefaultMasterSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"masterSigner\",\"type\":\"address\"}],\"name\":\"getInitCode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"entryPoint\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"masterSigner\",\"type\":\"address\"}],\"name\":\"isAccountDeployed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isDeployed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"isAuthorizedCreator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"masterAggregatorAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"revokeCreator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"senderCreator\",\"outputs\":[{\"internalType\":\"contractISenderCreator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"masterSigner\",\"type\":\"address\"}],\"name\":\"setDefaultMasterSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"setMasterAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"updateAccountImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// WalletManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletManagerMetaData.ABI instead.
var WalletManagerABI = WalletManagerMetaData.ABI

// WalletManager is an auto generated Go binding around an Ethereum contract.
type WalletManager struct {
	WalletManagerCaller     // Read-only binding to the contract
	WalletManagerTransactor // Write-only binding to the contract
	WalletManagerFilterer   // Log filterer for contract events
}

// WalletManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletManagerSession struct {
	Contract     *WalletManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletManagerCallerSession struct {
	Contract *WalletManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// WalletManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletManagerTransactorSession struct {
	Contract     *WalletManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// WalletManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletManagerRaw struct {
	Contract *WalletManager // Generic contract binding to access the raw methods on
}

// WalletManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletManagerCallerRaw struct {
	Contract *WalletManagerCaller // Generic read-only contract binding to access the raw methods on
}

// WalletManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletManagerTransactorRaw struct {
	Contract *WalletManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletManager creates a new instance of WalletManager, bound to a specific deployed contract.
func NewWalletManager(address common.Address, backend bind.ContractBackend) (*WalletManager, error) {
	contract, err := bindWalletManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WalletManager{WalletManagerCaller: WalletManagerCaller{contract: contract}, WalletManagerTransactor: WalletManagerTransactor{contract: contract}, WalletManagerFilterer: WalletManagerFilterer{contract: contract}}, nil
}

// NewWalletManagerCaller creates a new read-only instance of WalletManager, bound to a specific deployed contract.
func NewWalletManagerCaller(address common.Address, caller bind.ContractCaller) (*WalletManagerCaller, error) {
	contract, err := bindWalletManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletManagerCaller{contract: contract}, nil
}

// NewWalletManagerTransactor creates a new write-only instance of WalletManager, bound to a specific deployed contract.
func NewWalletManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletManagerTransactor, error) {
	contract, err := bindWalletManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletManagerTransactor{contract: contract}, nil
}

// NewWalletManagerFilterer creates a new log filterer instance of WalletManager, bound to a specific deployed contract.
func NewWalletManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletManagerFilterer, error) {
	contract, err := bindWalletManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletManagerFilterer{contract: contract}, nil
}

// bindWalletManager binds a generic wrapper to an already deployed contract.
func bindWalletManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WalletManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletManager *WalletManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletManager.Contract.WalletManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletManager *WalletManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletManager.Contract.WalletManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletManager *WalletManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletManager.Contract.WalletManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletManager *WalletManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletManager *WalletManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletManager *WalletManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletManager.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_WalletManager *WalletManagerCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WalletManager.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_WalletManager *WalletManagerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _WalletManager.Contract.UPGRADEINTERFACEVERSION(&_WalletManager.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_WalletManager *WalletManagerCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _WalletManager.Contract.UPGRADEINTERFACEVERSION(&_WalletManager.CallOpts)
}

// AccountImplementation is a free data retrieval call binding the contract method 0x11464fbe.
//
// Solidity: function accountImplementation() view returns(address)
func (_WalletManager *WalletManagerCaller) AccountImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletManager.contract.Call(opts, &out, "accountImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountImplementation is a free data retrieval call binding the contract method 0x11464fbe.
//
// Solidity: function accountImplementation() view returns(address)
func (_WalletManager *WalletManagerSession) AccountImplementation() (common.Address, error) {
	return _WalletManager.Contract.AccountImplementation(&_WalletManager.CallOpts)
}

// AccountImplementation is a free data retrieval call binding the contract method 0x11464fbe.
//
// Solidity: function accountImplementation() view returns(address)
func (_WalletManager *WalletManagerCallerSession) AccountImplementation() (common.Address, error) {
	return _WalletManager.Contract.AccountImplementation(&_WalletManager.CallOpts)
}

// DefaultMasterSigner is a free data retrieval call binding the contract method 0x1137a196.
//
// Solidity: function defaultMasterSigner() view returns(address)
func (_WalletManager *WalletManagerCaller) DefaultMasterSigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletManager.contract.Call(opts, &out, "defaultMasterSigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultMasterSigner is a free data retrieval call binding the contract method 0x1137a196.
//
// Solidity: function defaultMasterSigner() view returns(address)
func (_WalletManager *WalletManagerSession) DefaultMasterSigner() (common.Address, error) {
	return _WalletManager.Contract.DefaultMasterSigner(&_WalletManager.CallOpts)
}

// DefaultMasterSigner is a free data retrieval call binding the contract method 0x1137a196.
//
// Solidity: function defaultMasterSigner() view returns(address)
func (_WalletManager *WalletManagerCallerSession) DefaultMasterSigner() (common.Address, error) {
	return _WalletManager.Contract.DefaultMasterSigner(&_WalletManager.CallOpts)
}

// EntryPointAddress is a free data retrieval call binding the contract method 0x06dc245c.
//
// Solidity: function entryPointAddress() view returns(address)
func (_WalletManager *WalletManagerCaller) EntryPointAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletManager.contract.Call(opts, &out, "entryPointAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPointAddress is a free data retrieval call binding the contract method 0x06dc245c.
//
// Solidity: function entryPointAddress() view returns(address)
func (_WalletManager *WalletManagerSession) EntryPointAddress() (common.Address, error) {
	return _WalletManager.Contract.EntryPointAddress(&_WalletManager.CallOpts)
}

// EntryPointAddress is a free data retrieval call binding the contract method 0x06dc245c.
//
// Solidity: function entryPointAddress() view returns(address)
func (_WalletManager *WalletManagerCallerSession) EntryPointAddress() (common.Address, error) {
	return _WalletManager.Contract.EntryPointAddress(&_WalletManager.CallOpts)
}

// GetAccountAddress is a free data retrieval call binding the contract method 0x3c505342.
//
// Solidity: function getAccountAddress(address owner, address masterSigner) view returns(address account)
func (_WalletManager *WalletManagerCaller) GetAccountAddress(opts *bind.CallOpts, owner common.Address, masterSigner common.Address) (common.Address, error) {
	var out []interface{}
	err := _WalletManager.contract.Call(opts, &out, "getAccountAddress", owner, masterSigner)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAccountAddress is a free data retrieval call binding the contract method 0x3c505342.
//
// Solidity: function getAccountAddress(address owner, address masterSigner) view returns(address account)
func (_WalletManager *WalletManagerSession) GetAccountAddress(owner common.Address, masterSigner common.Address) (common.Address, error) {
	return _WalletManager.Contract.GetAccountAddress(&_WalletManager.CallOpts, owner, masterSigner)
}

// GetAccountAddress is a free data retrieval call binding the contract method 0x3c505342.
//
// Solidity: function getAccountAddress(address owner, address masterSigner) view returns(address account)
func (_WalletManager *WalletManagerCallerSession) GetAccountAddress(owner common.Address, masterSigner common.Address) (common.Address, error) {
	return _WalletManager.Contract.GetAccountAddress(&_WalletManager.CallOpts, owner, masterSigner)
}

// GetDefaultMasterSigner is a free data retrieval call binding the contract method 0xa4cd40a2.
//
// Solidity: function getDefaultMasterSigner() view returns(address)
func (_WalletManager *WalletManagerCaller) GetDefaultMasterSigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletManager.contract.Call(opts, &out, "getDefaultMasterSigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDefaultMasterSigner is a free data retrieval call binding the contract method 0xa4cd40a2.
//
// Solidity: function getDefaultMasterSigner() view returns(address)
func (_WalletManager *WalletManagerSession) GetDefaultMasterSigner() (common.Address, error) {
	return _WalletManager.Contract.GetDefaultMasterSigner(&_WalletManager.CallOpts)
}

// GetDefaultMasterSigner is a free data retrieval call binding the contract method 0xa4cd40a2.
//
// Solidity: function getDefaultMasterSigner() view returns(address)
func (_WalletManager *WalletManagerCallerSession) GetDefaultMasterSigner() (common.Address, error) {
	return _WalletManager.Contract.GetDefaultMasterSigner(&_WalletManager.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_WalletManager *WalletManagerCaller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletManager.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_WalletManager *WalletManagerSession) GetImplementation() (common.Address, error) {
	return _WalletManager.Contract.GetImplementation(&_WalletManager.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_WalletManager *WalletManagerCallerSession) GetImplementation() (common.Address, error) {
	return _WalletManager.Contract.GetImplementation(&_WalletManager.CallOpts)
}

// GetInitCode is a free data retrieval call binding the contract method 0xfa5c8528.
//
// Solidity: function getInitCode(address owner, address masterSigner) view returns(bytes initCode)
func (_WalletManager *WalletManagerCaller) GetInitCode(opts *bind.CallOpts, owner common.Address, masterSigner common.Address) ([]byte, error) {
	var out []interface{}
	err := _WalletManager.contract.Call(opts, &out, "getInitCode", owner, masterSigner)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInitCode is a free data retrieval call binding the contract method 0xfa5c8528.
//
// Solidity: function getInitCode(address owner, address masterSigner) view returns(bytes initCode)
func (_WalletManager *WalletManagerSession) GetInitCode(owner common.Address, masterSigner common.Address) ([]byte, error) {
	return _WalletManager.Contract.GetInitCode(&_WalletManager.CallOpts, owner, masterSigner)
}

// GetInitCode is a free data retrieval call binding the contract method 0xfa5c8528.
//
// Solidity: function getInitCode(address owner, address masterSigner) view returns(bytes initCode)
func (_WalletManager *WalletManagerCallerSession) GetInitCode(owner common.Address, masterSigner common.Address) ([]byte, error) {
	return _WalletManager.Contract.GetInitCode(&_WalletManager.CallOpts, owner, masterSigner)
}

// IsAccountDeployed is a free data retrieval call binding the contract method 0x4797aa4c.
//
// Solidity: function isAccountDeployed(address owner, address masterSigner) view returns(bool isDeployed)
func (_WalletManager *WalletManagerCaller) IsAccountDeployed(opts *bind.CallOpts, owner common.Address, masterSigner common.Address) (bool, error) {
	var out []interface{}
	err := _WalletManager.contract.Call(opts, &out, "isAccountDeployed", owner, masterSigner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAccountDeployed is a free data retrieval call binding the contract method 0x4797aa4c.
//
// Solidity: function isAccountDeployed(address owner, address masterSigner) view returns(bool isDeployed)
func (_WalletManager *WalletManagerSession) IsAccountDeployed(owner common.Address, masterSigner common.Address) (bool, error) {
	return _WalletManager.Contract.IsAccountDeployed(&_WalletManager.CallOpts, owner, masterSigner)
}

// IsAccountDeployed is a free data retrieval call binding the contract method 0x4797aa4c.
//
// Solidity: function isAccountDeployed(address owner, address masterSigner) view returns(bool isDeployed)
func (_WalletManager *WalletManagerCallerSession) IsAccountDeployed(owner common.Address, masterSigner common.Address) (bool, error) {
	return _WalletManager.Contract.IsAccountDeployed(&_WalletManager.CallOpts, owner, masterSigner)
}

// IsAuthorizedCreator is a free data retrieval call binding the contract method 0x79b53ebe.
//
// Solidity: function isAuthorizedCreator(address creator) view returns(bool)
func (_WalletManager *WalletManagerCaller) IsAuthorizedCreator(opts *bind.CallOpts, creator common.Address) (bool, error) {
	var out []interface{}
	err := _WalletManager.contract.Call(opts, &out, "isAuthorizedCreator", creator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuthorizedCreator is a free data retrieval call binding the contract method 0x79b53ebe.
//
// Solidity: function isAuthorizedCreator(address creator) view returns(bool)
func (_WalletManager *WalletManagerSession) IsAuthorizedCreator(creator common.Address) (bool, error) {
	return _WalletManager.Contract.IsAuthorizedCreator(&_WalletManager.CallOpts, creator)
}

// IsAuthorizedCreator is a free data retrieval call binding the contract method 0x79b53ebe.
//
// Solidity: function isAuthorizedCreator(address creator) view returns(bool)
func (_WalletManager *WalletManagerCallerSession) IsAuthorizedCreator(creator common.Address) (bool, error) {
	return _WalletManager.Contract.IsAuthorizedCreator(&_WalletManager.CallOpts, creator)
}

// MasterAggregatorAddress is a free data retrieval call binding the contract method 0x482aab42.
//
// Solidity: function masterAggregatorAddress() view returns(address)
func (_WalletManager *WalletManagerCaller) MasterAggregatorAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletManager.contract.Call(opts, &out, "masterAggregatorAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MasterAggregatorAddress is a free data retrieval call binding the contract method 0x482aab42.
//
// Solidity: function masterAggregatorAddress() view returns(address)
func (_WalletManager *WalletManagerSession) MasterAggregatorAddress() (common.Address, error) {
	return _WalletManager.Contract.MasterAggregatorAddress(&_WalletManager.CallOpts)
}

// MasterAggregatorAddress is a free data retrieval call binding the contract method 0x482aab42.
//
// Solidity: function masterAggregatorAddress() view returns(address)
func (_WalletManager *WalletManagerCallerSession) MasterAggregatorAddress() (common.Address, error) {
	return _WalletManager.Contract.MasterAggregatorAddress(&_WalletManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WalletManager *WalletManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WalletManager *WalletManagerSession) Owner() (common.Address, error) {
	return _WalletManager.Contract.Owner(&_WalletManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WalletManager *WalletManagerCallerSession) Owner() (common.Address, error) {
	return _WalletManager.Contract.Owner(&_WalletManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_WalletManager *WalletManagerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WalletManager.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_WalletManager *WalletManagerSession) ProxiableUUID() ([32]byte, error) {
	return _WalletManager.Contract.ProxiableUUID(&_WalletManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_WalletManager *WalletManagerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _WalletManager.Contract.ProxiableUUID(&_WalletManager.CallOpts)
}

// SenderCreator is a free data retrieval call binding the contract method 0x09ccb880.
//
// Solidity: function senderCreator() view returns(address)
func (_WalletManager *WalletManagerCaller) SenderCreator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletManager.contract.Call(opts, &out, "senderCreator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SenderCreator is a free data retrieval call binding the contract method 0x09ccb880.
//
// Solidity: function senderCreator() view returns(address)
func (_WalletManager *WalletManagerSession) SenderCreator() (common.Address, error) {
	return _WalletManager.Contract.SenderCreator(&_WalletManager.CallOpts)
}

// SenderCreator is a free data retrieval call binding the contract method 0x09ccb880.
//
// Solidity: function senderCreator() view returns(address)
func (_WalletManager *WalletManagerCallerSession) SenderCreator() (common.Address, error) {
	return _WalletManager.Contract.SenderCreator(&_WalletManager.CallOpts)
}

// AuthorizeCreator is a paid mutator transaction binding the contract method 0xd8625b07.
//
// Solidity: function authorizeCreator(address creator) returns()
func (_WalletManager *WalletManagerTransactor) AuthorizeCreator(opts *bind.TransactOpts, creator common.Address) (*types.Transaction, error) {
	return _WalletManager.contract.Transact(opts, "authorizeCreator", creator)
}

// AuthorizeCreator is a paid mutator transaction binding the contract method 0xd8625b07.
//
// Solidity: function authorizeCreator(address creator) returns()
func (_WalletManager *WalletManagerSession) AuthorizeCreator(creator common.Address) (*types.Transaction, error) {
	return _WalletManager.Contract.AuthorizeCreator(&_WalletManager.TransactOpts, creator)
}

// AuthorizeCreator is a paid mutator transaction binding the contract method 0xd8625b07.
//
// Solidity: function authorizeCreator(address creator) returns()
func (_WalletManager *WalletManagerTransactorSession) AuthorizeCreator(creator common.Address) (*types.Transaction, error) {
	return _WalletManager.Contract.AuthorizeCreator(&_WalletManager.TransactOpts, creator)
}

// CreateAccountWithMasterSigner is a paid mutator transaction binding the contract method 0x4242339f.
//
// Solidity: function createAccountWithMasterSigner(address owner, address masterSigner) returns(address account)
func (_WalletManager *WalletManagerTransactor) CreateAccountWithMasterSigner(opts *bind.TransactOpts, owner common.Address, masterSigner common.Address) (*types.Transaction, error) {
	return _WalletManager.contract.Transact(opts, "createAccountWithMasterSigner", owner, masterSigner)
}

// CreateAccountWithMasterSigner is a paid mutator transaction binding the contract method 0x4242339f.
//
// Solidity: function createAccountWithMasterSigner(address owner, address masterSigner) returns(address account)
func (_WalletManager *WalletManagerSession) CreateAccountWithMasterSigner(owner common.Address, masterSigner common.Address) (*types.Transaction, error) {
	return _WalletManager.Contract.CreateAccountWithMasterSigner(&_WalletManager.TransactOpts, owner, masterSigner)
}

// CreateAccountWithMasterSigner is a paid mutator transaction binding the contract method 0x4242339f.
//
// Solidity: function createAccountWithMasterSigner(address owner, address masterSigner) returns(address account)
func (_WalletManager *WalletManagerTransactorSession) CreateAccountWithMasterSigner(owner common.Address, masterSigner common.Address) (*types.Transaction, error) {
	return _WalletManager.Contract.CreateAccountWithMasterSigner(&_WalletManager.TransactOpts, owner, masterSigner)
}

// CreateUserAccount is a paid mutator transaction binding the contract method 0x64857d26.
//
// Solidity: function createUserAccount(address owner, address masterSigner) returns(address account)
func (_WalletManager *WalletManagerTransactor) CreateUserAccount(opts *bind.TransactOpts, owner common.Address, masterSigner common.Address) (*types.Transaction, error) {
	return _WalletManager.contract.Transact(opts, "createUserAccount", owner, masterSigner)
}

// CreateUserAccount is a paid mutator transaction binding the contract method 0x64857d26.
//
// Solidity: function createUserAccount(address owner, address masterSigner) returns(address account)
func (_WalletManager *WalletManagerSession) CreateUserAccount(owner common.Address, masterSigner common.Address) (*types.Transaction, error) {
	return _WalletManager.Contract.CreateUserAccount(&_WalletManager.TransactOpts, owner, masterSigner)
}

// CreateUserAccount is a paid mutator transaction binding the contract method 0x64857d26.
//
// Solidity: function createUserAccount(address owner, address masterSigner) returns(address account)
func (_WalletManager *WalletManagerTransactorSession) CreateUserAccount(owner common.Address, masterSigner common.Address) (*types.Transaction, error) {
	return _WalletManager.Contract.CreateUserAccount(&_WalletManager.TransactOpts, owner, masterSigner)
}

// CreateWallet is a paid mutator transaction binding the contract method 0x8f860c5f.
//
// Solidity: function createWallet(address owner, address masterSigner) returns(address account)
func (_WalletManager *WalletManagerTransactor) CreateWallet(opts *bind.TransactOpts, owner common.Address, masterSigner common.Address) (*types.Transaction, error) {
	return _WalletManager.contract.Transact(opts, "createWallet", owner, masterSigner)
}

// CreateWallet is a paid mutator transaction binding the contract method 0x8f860c5f.
//
// Solidity: function createWallet(address owner, address masterSigner) returns(address account)
func (_WalletManager *WalletManagerSession) CreateWallet(owner common.Address, masterSigner common.Address) (*types.Transaction, error) {
	return _WalletManager.Contract.CreateWallet(&_WalletManager.TransactOpts, owner, masterSigner)
}

// CreateWallet is a paid mutator transaction binding the contract method 0x8f860c5f.
//
// Solidity: function createWallet(address owner, address masterSigner) returns(address account)
func (_WalletManager *WalletManagerTransactorSession) CreateWallet(owner common.Address, masterSigner common.Address) (*types.Transaction, error) {
	return _WalletManager.Contract.CreateWallet(&_WalletManager.TransactOpts, owner, masterSigner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address entryPoint, address owner) returns()
func (_WalletManager *WalletManagerTransactor) Initialize(opts *bind.TransactOpts, entryPoint common.Address, owner common.Address) (*types.Transaction, error) {
	return _WalletManager.contract.Transact(opts, "initialize", entryPoint, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address entryPoint, address owner) returns()
func (_WalletManager *WalletManagerSession) Initialize(entryPoint common.Address, owner common.Address) (*types.Transaction, error) {
	return _WalletManager.Contract.Initialize(&_WalletManager.TransactOpts, entryPoint, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address entryPoint, address owner) returns()
func (_WalletManager *WalletManagerTransactorSession) Initialize(entryPoint common.Address, owner common.Address) (*types.Transaction, error) {
	return _WalletManager.Contract.Initialize(&_WalletManager.TransactOpts, entryPoint, owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WalletManager *WalletManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WalletManager *WalletManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _WalletManager.Contract.RenounceOwnership(&_WalletManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WalletManager *WalletManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _WalletManager.Contract.RenounceOwnership(&_WalletManager.TransactOpts)
}

// RevokeCreator is a paid mutator transaction binding the contract method 0xc0e69af6.
//
// Solidity: function revokeCreator(address creator) returns()
func (_WalletManager *WalletManagerTransactor) RevokeCreator(opts *bind.TransactOpts, creator common.Address) (*types.Transaction, error) {
	return _WalletManager.contract.Transact(opts, "revokeCreator", creator)
}

// RevokeCreator is a paid mutator transaction binding the contract method 0xc0e69af6.
//
// Solidity: function revokeCreator(address creator) returns()
func (_WalletManager *WalletManagerSession) RevokeCreator(creator common.Address) (*types.Transaction, error) {
	return _WalletManager.Contract.RevokeCreator(&_WalletManager.TransactOpts, creator)
}

// RevokeCreator is a paid mutator transaction binding the contract method 0xc0e69af6.
//
// Solidity: function revokeCreator(address creator) returns()
func (_WalletManager *WalletManagerTransactorSession) RevokeCreator(creator common.Address) (*types.Transaction, error) {
	return _WalletManager.Contract.RevokeCreator(&_WalletManager.TransactOpts, creator)
}

// SetDefaultMasterSigner is a paid mutator transaction binding the contract method 0xa3a2bb76.
//
// Solidity: function setDefaultMasterSigner(address masterSigner) returns()
func (_WalletManager *WalletManagerTransactor) SetDefaultMasterSigner(opts *bind.TransactOpts, masterSigner common.Address) (*types.Transaction, error) {
	return _WalletManager.contract.Transact(opts, "setDefaultMasterSigner", masterSigner)
}

// SetDefaultMasterSigner is a paid mutator transaction binding the contract method 0xa3a2bb76.
//
// Solidity: function setDefaultMasterSigner(address masterSigner) returns()
func (_WalletManager *WalletManagerSession) SetDefaultMasterSigner(masterSigner common.Address) (*types.Transaction, error) {
	return _WalletManager.Contract.SetDefaultMasterSigner(&_WalletManager.TransactOpts, masterSigner)
}

// SetDefaultMasterSigner is a paid mutator transaction binding the contract method 0xa3a2bb76.
//
// Solidity: function setDefaultMasterSigner(address masterSigner) returns()
func (_WalletManager *WalletManagerTransactorSession) SetDefaultMasterSigner(masterSigner common.Address) (*types.Transaction, error) {
	return _WalletManager.Contract.SetDefaultMasterSigner(&_WalletManager.TransactOpts, masterSigner)
}

// SetMasterAggregator is a paid mutator transaction binding the contract method 0x67d944fd.
//
// Solidity: function setMasterAggregator(address aggregator) returns()
func (_WalletManager *WalletManagerTransactor) SetMasterAggregator(opts *bind.TransactOpts, aggregator common.Address) (*types.Transaction, error) {
	return _WalletManager.contract.Transact(opts, "setMasterAggregator", aggregator)
}

// SetMasterAggregator is a paid mutator transaction binding the contract method 0x67d944fd.
//
// Solidity: function setMasterAggregator(address aggregator) returns()
func (_WalletManager *WalletManagerSession) SetMasterAggregator(aggregator common.Address) (*types.Transaction, error) {
	return _WalletManager.Contract.SetMasterAggregator(&_WalletManager.TransactOpts, aggregator)
}

// SetMasterAggregator is a paid mutator transaction binding the contract method 0x67d944fd.
//
// Solidity: function setMasterAggregator(address aggregator) returns()
func (_WalletManager *WalletManagerTransactorSession) SetMasterAggregator(aggregator common.Address) (*types.Transaction, error) {
	return _WalletManager.Contract.SetMasterAggregator(&_WalletManager.TransactOpts, aggregator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WalletManager *WalletManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WalletManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WalletManager *WalletManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WalletManager.Contract.TransferOwnership(&_WalletManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WalletManager *WalletManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WalletManager.Contract.TransferOwnership(&_WalletManager.TransactOpts, newOwner)
}

// UpdateAccountImplementation is a paid mutator transaction binding the contract method 0x0a5b9e6e.
//
// Solidity: function updateAccountImplementation(address newImplementation) returns()
func (_WalletManager *WalletManagerTransactor) UpdateAccountImplementation(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _WalletManager.contract.Transact(opts, "updateAccountImplementation", newImplementation)
}

// UpdateAccountImplementation is a paid mutator transaction binding the contract method 0x0a5b9e6e.
//
// Solidity: function updateAccountImplementation(address newImplementation) returns()
func (_WalletManager *WalletManagerSession) UpdateAccountImplementation(newImplementation common.Address) (*types.Transaction, error) {
	return _WalletManager.Contract.UpdateAccountImplementation(&_WalletManager.TransactOpts, newImplementation)
}

// UpdateAccountImplementation is a paid mutator transaction binding the contract method 0x0a5b9e6e.
//
// Solidity: function updateAccountImplementation(address newImplementation) returns()
func (_WalletManager *WalletManagerTransactorSession) UpdateAccountImplementation(newImplementation common.Address) (*types.Transaction, error) {
	return _WalletManager.Contract.UpdateAccountImplementation(&_WalletManager.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_WalletManager *WalletManagerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _WalletManager.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_WalletManager *WalletManagerSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _WalletManager.Contract.UpgradeToAndCall(&_WalletManager.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_WalletManager *WalletManagerTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _WalletManager.Contract.UpgradeToAndCall(&_WalletManager.TransactOpts, newImplementation, data)
}

// WalletManagerAccountCreatedIterator is returned from FilterAccountCreated and is used to iterate over the raw logs and unpacked data for AccountCreated events raised by the WalletManager contract.
type WalletManagerAccountCreatedIterator struct {
	Event *WalletManagerAccountCreated // Event containing the contract specifics and raw log

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
func (it *WalletManagerAccountCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletManagerAccountCreated)
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
		it.Event = new(WalletManagerAccountCreated)
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
func (it *WalletManagerAccountCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletManagerAccountCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletManagerAccountCreated represents a AccountCreated event raised by the WalletManager contract.
type WalletManagerAccountCreated struct {
	Account      common.Address
	Owner        common.Address
	MasterSigner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAccountCreated is a free log retrieval operation binding the contract event 0x9db37ffcb3d1eadb0306e35c0453e35a675f2e80757b09251d848185ac5cc22a.
//
// Solidity: event AccountCreated(address indexed account, address indexed owner, address indexed masterSigner)
func (_WalletManager *WalletManagerFilterer) FilterAccountCreated(opts *bind.FilterOpts, account []common.Address, owner []common.Address, masterSigner []common.Address) (*WalletManagerAccountCreatedIterator, error) {

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

	logs, sub, err := _WalletManager.contract.FilterLogs(opts, "AccountCreated", accountRule, ownerRule, masterSignerRule)
	if err != nil {
		return nil, err
	}
	return &WalletManagerAccountCreatedIterator{contract: _WalletManager.contract, event: "AccountCreated", logs: logs, sub: sub}, nil
}

// WatchAccountCreated is a free log subscription operation binding the contract event 0x9db37ffcb3d1eadb0306e35c0453e35a675f2e80757b09251d848185ac5cc22a.
//
// Solidity: event AccountCreated(address indexed account, address indexed owner, address indexed masterSigner)
func (_WalletManager *WalletManagerFilterer) WatchAccountCreated(opts *bind.WatchOpts, sink chan<- *WalletManagerAccountCreated, account []common.Address, owner []common.Address, masterSigner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _WalletManager.contract.WatchLogs(opts, "AccountCreated", accountRule, ownerRule, masterSignerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletManagerAccountCreated)
				if err := _WalletManager.contract.UnpackLog(event, "AccountCreated", log); err != nil {
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
func (_WalletManager *WalletManagerFilterer) ParseAccountCreated(log types.Log) (*WalletManagerAccountCreated, error) {
	event := new(WalletManagerAccountCreated)
	if err := _WalletManager.contract.UnpackLog(event, "AccountCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletManagerCreatorAuthorizedIterator is returned from FilterCreatorAuthorized and is used to iterate over the raw logs and unpacked data for CreatorAuthorized events raised by the WalletManager contract.
type WalletManagerCreatorAuthorizedIterator struct {
	Event *WalletManagerCreatorAuthorized // Event containing the contract specifics and raw log

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
func (it *WalletManagerCreatorAuthorizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletManagerCreatorAuthorized)
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
		it.Event = new(WalletManagerCreatorAuthorized)
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
func (it *WalletManagerCreatorAuthorizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletManagerCreatorAuthorizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletManagerCreatorAuthorized represents a CreatorAuthorized event raised by the WalletManager contract.
type WalletManagerCreatorAuthorized struct {
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCreatorAuthorized is a free log retrieval operation binding the contract event 0x00b6d3c5059e85c667aa67cfa572e42162916972ed621a4618e55dfcb4a17463.
//
// Solidity: event CreatorAuthorized(address indexed creator)
func (_WalletManager *WalletManagerFilterer) FilterCreatorAuthorized(opts *bind.FilterOpts, creator []common.Address) (*WalletManagerCreatorAuthorizedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _WalletManager.contract.FilterLogs(opts, "CreatorAuthorized", creatorRule)
	if err != nil {
		return nil, err
	}
	return &WalletManagerCreatorAuthorizedIterator{contract: _WalletManager.contract, event: "CreatorAuthorized", logs: logs, sub: sub}, nil
}

// WatchCreatorAuthorized is a free log subscription operation binding the contract event 0x00b6d3c5059e85c667aa67cfa572e42162916972ed621a4618e55dfcb4a17463.
//
// Solidity: event CreatorAuthorized(address indexed creator)
func (_WalletManager *WalletManagerFilterer) WatchCreatorAuthorized(opts *bind.WatchOpts, sink chan<- *WalletManagerCreatorAuthorized, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _WalletManager.contract.WatchLogs(opts, "CreatorAuthorized", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletManagerCreatorAuthorized)
				if err := _WalletManager.contract.UnpackLog(event, "CreatorAuthorized", log); err != nil {
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
func (_WalletManager *WalletManagerFilterer) ParseCreatorAuthorized(log types.Log) (*WalletManagerCreatorAuthorized, error) {
	event := new(WalletManagerCreatorAuthorized)
	if err := _WalletManager.contract.UnpackLog(event, "CreatorAuthorized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletManagerCreatorRevokedIterator is returned from FilterCreatorRevoked and is used to iterate over the raw logs and unpacked data for CreatorRevoked events raised by the WalletManager contract.
type WalletManagerCreatorRevokedIterator struct {
	Event *WalletManagerCreatorRevoked // Event containing the contract specifics and raw log

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
func (it *WalletManagerCreatorRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletManagerCreatorRevoked)
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
		it.Event = new(WalletManagerCreatorRevoked)
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
func (it *WalletManagerCreatorRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletManagerCreatorRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletManagerCreatorRevoked represents a CreatorRevoked event raised by the WalletManager contract.
type WalletManagerCreatorRevoked struct {
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCreatorRevoked is a free log retrieval operation binding the contract event 0xc77d51d04162a5d61b26217b8a8c78870a02749a7a198336685c57d6f128fa4b.
//
// Solidity: event CreatorRevoked(address indexed creator)
func (_WalletManager *WalletManagerFilterer) FilterCreatorRevoked(opts *bind.FilterOpts, creator []common.Address) (*WalletManagerCreatorRevokedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _WalletManager.contract.FilterLogs(opts, "CreatorRevoked", creatorRule)
	if err != nil {
		return nil, err
	}
	return &WalletManagerCreatorRevokedIterator{contract: _WalletManager.contract, event: "CreatorRevoked", logs: logs, sub: sub}, nil
}

// WatchCreatorRevoked is a free log subscription operation binding the contract event 0xc77d51d04162a5d61b26217b8a8c78870a02749a7a198336685c57d6f128fa4b.
//
// Solidity: event CreatorRevoked(address indexed creator)
func (_WalletManager *WalletManagerFilterer) WatchCreatorRevoked(opts *bind.WatchOpts, sink chan<- *WalletManagerCreatorRevoked, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _WalletManager.contract.WatchLogs(opts, "CreatorRevoked", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletManagerCreatorRevoked)
				if err := _WalletManager.contract.UnpackLog(event, "CreatorRevoked", log); err != nil {
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
func (_WalletManager *WalletManagerFilterer) ParseCreatorRevoked(log types.Log) (*WalletManagerCreatorRevoked, error) {
	event := new(WalletManagerCreatorRevoked)
	if err := _WalletManager.contract.UnpackLog(event, "CreatorRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the WalletManager contract.
type WalletManagerInitializedIterator struct {
	Event *WalletManagerInitialized // Event containing the contract specifics and raw log

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
func (it *WalletManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletManagerInitialized)
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
		it.Event = new(WalletManagerInitialized)
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
func (it *WalletManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletManagerInitialized represents a Initialized event raised by the WalletManager contract.
type WalletManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_WalletManager *WalletManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*WalletManagerInitializedIterator, error) {

	logs, sub, err := _WalletManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &WalletManagerInitializedIterator{contract: _WalletManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_WalletManager *WalletManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *WalletManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _WalletManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletManagerInitialized)
				if err := _WalletManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_WalletManager *WalletManagerFilterer) ParseInitialized(log types.Log) (*WalletManagerInitialized, error) {
	event := new(WalletManagerInitialized)
	if err := _WalletManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WalletManager contract.
type WalletManagerOwnershipTransferredIterator struct {
	Event *WalletManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WalletManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletManagerOwnershipTransferred)
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
		it.Event = new(WalletManagerOwnershipTransferred)
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
func (it *WalletManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletManagerOwnershipTransferred represents a OwnershipTransferred event raised by the WalletManager contract.
type WalletManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WalletManager *WalletManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WalletManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WalletManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WalletManagerOwnershipTransferredIterator{contract: _WalletManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WalletManager *WalletManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WalletManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WalletManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletManagerOwnershipTransferred)
				if err := _WalletManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_WalletManager *WalletManagerFilterer) ParseOwnershipTransferred(log types.Log) (*WalletManagerOwnershipTransferred, error) {
	event := new(WalletManagerOwnershipTransferred)
	if err := _WalletManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletManagerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the WalletManager contract.
type WalletManagerUpgradedIterator struct {
	Event *WalletManagerUpgraded // Event containing the contract specifics and raw log

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
func (it *WalletManagerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletManagerUpgraded)
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
		it.Event = new(WalletManagerUpgraded)
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
func (it *WalletManagerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletManagerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletManagerUpgraded represents a Upgraded event raised by the WalletManager contract.
type WalletManagerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_WalletManager *WalletManagerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*WalletManagerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _WalletManager.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &WalletManagerUpgradedIterator{contract: _WalletManager.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_WalletManager *WalletManagerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *WalletManagerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _WalletManager.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletManagerUpgraded)
				if err := _WalletManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_WalletManager *WalletManagerFilterer) ParseUpgraded(log types.Log) (*WalletManagerUpgraded, error) {
	event := new(WalletManagerUpgraded)
	if err := _WalletManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
