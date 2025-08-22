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

// BaseAccountCall is an auto generated low-level Go binding around an user-defined struct.
type BaseAccountCall struct {
	Target common.Address
	Value  *big.Int
	Data   []byte
}

// PackedUserOperation is an auto generated low-level Go binding around an user-defined struct.
type PackedUserOperation struct {
	Sender             common.Address
	Nonce              *big.Int
	InitCode           []byte
	CallData           []byte
	AccountGasLimits   [32]byte
	PreVerificationGas *big.Int
	GasFees            [32]byte
	PaymasterAndData   []byte
	Signature          []byte
}

// AAWalletMetaData contains all meta data concerning the AAWallet contract.
var AAWalletMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ExecuteError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"masterSigner\",\"type\":\"address\"}],\"name\":\"AAWalletInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAggregator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAggregator\",\"type\":\"address\"}],\"name\":\"AggregatorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldMasterSigner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newMasterSigner\",\"type\":\"address\"}],\"name\":\"MasterSignerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"masterSigner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"}],\"name\":\"MasterSignerValidation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"validAfter\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"validUntil\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"permissions\",\"type\":\"bytes32\"}],\"name\":\"SessionKeyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"}],\"name\":\"SessionKeyRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"}],\"name\":\"SessionKeyValidation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"validAfter\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"validUntil\",\"type\":\"uint48\"},{\"internalType\":\"bytes32\",\"name\":\"permissions\",\"type\":\"bytes32\"}],\"name\":\"addSessionKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggregatorAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"authorizedSessionKeyManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"canSessionKeyExecute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"canExecute\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPointAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structBaseAccount.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"executeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMasterSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"}],\"name\":\"getSessionKeyInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"},{\"internalType\":\"uint48\",\"name\":\"validAfter\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"validUntil\",\"type\":\"uint48\"},{\"internalType\":\"bytes32\",\"name\":\"permissions\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_entryPoint\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_masterSigner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMasterSignerEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"masterSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"}],\"name\":\"revokeSessionKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sessionKeys\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"validAfter\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"validUntil\",\"type\":\"uint48\"},{\"internalType\":\"bytes32\",\"name\":\"permissions\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"}],\"name\":\"setAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sessionKeyManager\",\"type\":\"address\"}],\"name\":\"setAuthorizedSessionKeyManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newMasterSigner\",\"type\":\"address\"}],\"name\":\"setMasterSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"missingAccountFunds\",\"type\":\"uint256\"}],\"name\":\"validateUserOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawDepositTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// AAWalletABI is the input ABI used to generate the binding from.
// Deprecated: Use AAWalletMetaData.ABI instead.
var AAWalletABI = AAWalletMetaData.ABI

// AAWallet is an auto generated Go binding around an Ethereum contract.
type AAWallet struct {
	AAWalletCaller     // Read-only binding to the contract
	AAWalletTransactor // Write-only binding to the contract
	AAWalletFilterer   // Log filterer for contract events
}

// AAWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type AAWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AAWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AAWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AAWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AAWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AAWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AAWalletSession struct {
	Contract     *AAWallet         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AAWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AAWalletCallerSession struct {
	Contract *AAWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// AAWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AAWalletTransactorSession struct {
	Contract     *AAWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AAWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type AAWalletRaw struct {
	Contract *AAWallet // Generic contract binding to access the raw methods on
}

// AAWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AAWalletCallerRaw struct {
	Contract *AAWalletCaller // Generic read-only contract binding to access the raw methods on
}

// AAWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AAWalletTransactorRaw struct {
	Contract *AAWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAAWallet creates a new instance of AAWallet, bound to a specific deployed contract.
func NewAAWallet(address common.Address, backend bind.ContractBackend) (*AAWallet, error) {
	contract, err := bindAAWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AAWallet{AAWalletCaller: AAWalletCaller{contract: contract}, AAWalletTransactor: AAWalletTransactor{contract: contract}, AAWalletFilterer: AAWalletFilterer{contract: contract}}, nil
}

// NewAAWalletCaller creates a new read-only instance of AAWallet, bound to a specific deployed contract.
func NewAAWalletCaller(address common.Address, caller bind.ContractCaller) (*AAWalletCaller, error) {
	contract, err := bindAAWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AAWalletCaller{contract: contract}, nil
}

// NewAAWalletTransactor creates a new write-only instance of AAWallet, bound to a specific deployed contract.
func NewAAWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*AAWalletTransactor, error) {
	contract, err := bindAAWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AAWalletTransactor{contract: contract}, nil
}

// NewAAWalletFilterer creates a new log filterer instance of AAWallet, bound to a specific deployed contract.
func NewAAWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*AAWalletFilterer, error) {
	contract, err := bindAAWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AAWalletFilterer{contract: contract}, nil
}

// bindAAWallet binds a generic wrapper to an already deployed contract.
func bindAAWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AAWalletMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AAWallet *AAWalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AAWallet.Contract.AAWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AAWallet *AAWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AAWallet.Contract.AAWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AAWallet *AAWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AAWallet.Contract.AAWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AAWallet *AAWalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AAWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AAWallet *AAWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AAWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AAWallet *AAWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AAWallet.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_AAWallet *AAWalletCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AAWallet.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_AAWallet *AAWalletSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _AAWallet.Contract.UPGRADEINTERFACEVERSION(&_AAWallet.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_AAWallet *AAWalletCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _AAWallet.Contract.UPGRADEINTERFACEVERSION(&_AAWallet.CallOpts)
}

// AggregatorAddress is a free data retrieval call binding the contract method 0x380bbe53.
//
// Solidity: function aggregatorAddress() view returns(address)
func (_AAWallet *AAWalletCaller) AggregatorAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AAWallet.contract.Call(opts, &out, "aggregatorAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AggregatorAddress is a free data retrieval call binding the contract method 0x380bbe53.
//
// Solidity: function aggregatorAddress() view returns(address)
func (_AAWallet *AAWalletSession) AggregatorAddress() (common.Address, error) {
	return _AAWallet.Contract.AggregatorAddress(&_AAWallet.CallOpts)
}

// AggregatorAddress is a free data retrieval call binding the contract method 0x380bbe53.
//
// Solidity: function aggregatorAddress() view returns(address)
func (_AAWallet *AAWalletCallerSession) AggregatorAddress() (common.Address, error) {
	return _AAWallet.Contract.AggregatorAddress(&_AAWallet.CallOpts)
}

// AuthorizedSessionKeyManager is a free data retrieval call binding the contract method 0x4b6c4f0b.
//
// Solidity: function authorizedSessionKeyManager() view returns(address)
func (_AAWallet *AAWalletCaller) AuthorizedSessionKeyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AAWallet.contract.Call(opts, &out, "authorizedSessionKeyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AuthorizedSessionKeyManager is a free data retrieval call binding the contract method 0x4b6c4f0b.
//
// Solidity: function authorizedSessionKeyManager() view returns(address)
func (_AAWallet *AAWalletSession) AuthorizedSessionKeyManager() (common.Address, error) {
	return _AAWallet.Contract.AuthorizedSessionKeyManager(&_AAWallet.CallOpts)
}

// AuthorizedSessionKeyManager is a free data retrieval call binding the contract method 0x4b6c4f0b.
//
// Solidity: function authorizedSessionKeyManager() view returns(address)
func (_AAWallet *AAWalletCallerSession) AuthorizedSessionKeyManager() (common.Address, error) {
	return _AAWallet.Contract.AuthorizedSessionKeyManager(&_AAWallet.CallOpts)
}

// CanSessionKeyExecute is a free data retrieval call binding the contract method 0x2c4d8fcf.
//
// Solidity: function canSessionKeyExecute(address sessionKey, address target, bytes4 selector) view returns(bool canExecute)
func (_AAWallet *AAWalletCaller) CanSessionKeyExecute(opts *bind.CallOpts, sessionKey common.Address, target common.Address, selector [4]byte) (bool, error) {
	var out []interface{}
	err := _AAWallet.contract.Call(opts, &out, "canSessionKeyExecute", sessionKey, target, selector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanSessionKeyExecute is a free data retrieval call binding the contract method 0x2c4d8fcf.
//
// Solidity: function canSessionKeyExecute(address sessionKey, address target, bytes4 selector) view returns(bool canExecute)
func (_AAWallet *AAWalletSession) CanSessionKeyExecute(sessionKey common.Address, target common.Address, selector [4]byte) (bool, error) {
	return _AAWallet.Contract.CanSessionKeyExecute(&_AAWallet.CallOpts, sessionKey, target, selector)
}

// CanSessionKeyExecute is a free data retrieval call binding the contract method 0x2c4d8fcf.
//
// Solidity: function canSessionKeyExecute(address sessionKey, address target, bytes4 selector) view returns(bool canExecute)
func (_AAWallet *AAWalletCallerSession) CanSessionKeyExecute(sessionKey common.Address, target common.Address, selector [4]byte) (bool, error) {
	return _AAWallet.Contract.CanSessionKeyExecute(&_AAWallet.CallOpts, sessionKey, target, selector)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_AAWallet *AAWalletCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AAWallet.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_AAWallet *AAWalletSession) EntryPoint() (common.Address, error) {
	return _AAWallet.Contract.EntryPoint(&_AAWallet.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_AAWallet *AAWalletCallerSession) EntryPoint() (common.Address, error) {
	return _AAWallet.Contract.EntryPoint(&_AAWallet.CallOpts)
}

// EntryPointAddress is a free data retrieval call binding the contract method 0x06dc245c.
//
// Solidity: function entryPointAddress() view returns(address)
func (_AAWallet *AAWalletCaller) EntryPointAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AAWallet.contract.Call(opts, &out, "entryPointAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPointAddress is a free data retrieval call binding the contract method 0x06dc245c.
//
// Solidity: function entryPointAddress() view returns(address)
func (_AAWallet *AAWalletSession) EntryPointAddress() (common.Address, error) {
	return _AAWallet.Contract.EntryPointAddress(&_AAWallet.CallOpts)
}

// EntryPointAddress is a free data retrieval call binding the contract method 0x06dc245c.
//
// Solidity: function entryPointAddress() view returns(address)
func (_AAWallet *AAWalletCallerSession) EntryPointAddress() (common.Address, error) {
	return _AAWallet.Contract.EntryPointAddress(&_AAWallet.CallOpts)
}

// GetAggregator is a free data retrieval call binding the contract method 0x3ad59dbc.
//
// Solidity: function getAggregator() view returns(address aggregator)
func (_AAWallet *AAWalletCaller) GetAggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AAWallet.contract.Call(opts, &out, "getAggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAggregator is a free data retrieval call binding the contract method 0x3ad59dbc.
//
// Solidity: function getAggregator() view returns(address aggregator)
func (_AAWallet *AAWalletSession) GetAggregator() (common.Address, error) {
	return _AAWallet.Contract.GetAggregator(&_AAWallet.CallOpts)
}

// GetAggregator is a free data retrieval call binding the contract method 0x3ad59dbc.
//
// Solidity: function getAggregator() view returns(address aggregator)
func (_AAWallet *AAWalletCallerSession) GetAggregator() (common.Address, error) {
	return _AAWallet.Contract.GetAggregator(&_AAWallet.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_AAWallet *AAWalletCaller) GetDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AAWallet.contract.Call(opts, &out, "getDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_AAWallet *AAWalletSession) GetDeposit() (*big.Int, error) {
	return _AAWallet.Contract.GetDeposit(&_AAWallet.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_AAWallet *AAWalletCallerSession) GetDeposit() (*big.Int, error) {
	return _AAWallet.Contract.GetDeposit(&_AAWallet.CallOpts)
}

// GetMasterSigner is a free data retrieval call binding the contract method 0x0b2f4d9e.
//
// Solidity: function getMasterSigner() view returns(address)
func (_AAWallet *AAWalletCaller) GetMasterSigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AAWallet.contract.Call(opts, &out, "getMasterSigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMasterSigner is a free data retrieval call binding the contract method 0x0b2f4d9e.
//
// Solidity: function getMasterSigner() view returns(address)
func (_AAWallet *AAWalletSession) GetMasterSigner() (common.Address, error) {
	return _AAWallet.Contract.GetMasterSigner(&_AAWallet.CallOpts)
}

// GetMasterSigner is a free data retrieval call binding the contract method 0x0b2f4d9e.
//
// Solidity: function getMasterSigner() view returns(address)
func (_AAWallet *AAWalletCallerSession) GetMasterSigner() (common.Address, error) {
	return _AAWallet.Contract.GetMasterSigner(&_AAWallet.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_AAWallet *AAWalletCaller) GetNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AAWallet.contract.Call(opts, &out, "getNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_AAWallet *AAWalletSession) GetNonce() (*big.Int, error) {
	return _AAWallet.Contract.GetNonce(&_AAWallet.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_AAWallet *AAWalletCallerSession) GetNonce() (*big.Int, error) {
	return _AAWallet.Contract.GetNonce(&_AAWallet.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_AAWallet *AAWalletCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AAWallet.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_AAWallet *AAWalletSession) GetOwner() (common.Address, error) {
	return _AAWallet.Contract.GetOwner(&_AAWallet.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_AAWallet *AAWalletCallerSession) GetOwner() (common.Address, error) {
	return _AAWallet.Contract.GetOwner(&_AAWallet.CallOpts)
}

// GetSessionKeyInfo is a free data retrieval call binding the contract method 0x5ad2657f.
//
// Solidity: function getSessionKeyInfo(address sessionKey) view returns(bool isValid, uint48 validAfter, uint48 validUntil, bytes32 permissions)
func (_AAWallet *AAWalletCaller) GetSessionKeyInfo(opts *bind.CallOpts, sessionKey common.Address) (struct {
	IsValid     bool
	ValidAfter  *big.Int
	ValidUntil  *big.Int
	Permissions [32]byte
}, error) {
	var out []interface{}
	err := _AAWallet.contract.Call(opts, &out, "getSessionKeyInfo", sessionKey)

	outstruct := new(struct {
		IsValid     bool
		ValidAfter  *big.Int
		ValidUntil  *big.Int
		Permissions [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsValid = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ValidAfter = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ValidUntil = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Permissions = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// GetSessionKeyInfo is a free data retrieval call binding the contract method 0x5ad2657f.
//
// Solidity: function getSessionKeyInfo(address sessionKey) view returns(bool isValid, uint48 validAfter, uint48 validUntil, bytes32 permissions)
func (_AAWallet *AAWalletSession) GetSessionKeyInfo(sessionKey common.Address) (struct {
	IsValid     bool
	ValidAfter  *big.Int
	ValidUntil  *big.Int
	Permissions [32]byte
}, error) {
	return _AAWallet.Contract.GetSessionKeyInfo(&_AAWallet.CallOpts, sessionKey)
}

// GetSessionKeyInfo is a free data retrieval call binding the contract method 0x5ad2657f.
//
// Solidity: function getSessionKeyInfo(address sessionKey) view returns(bool isValid, uint48 validAfter, uint48 validUntil, bytes32 permissions)
func (_AAWallet *AAWalletCallerSession) GetSessionKeyInfo(sessionKey common.Address) (struct {
	IsValid     bool
	ValidAfter  *big.Int
	ValidUntil  *big.Int
	Permissions [32]byte
}, error) {
	return _AAWallet.Contract.GetSessionKeyInfo(&_AAWallet.CallOpts, sessionKey)
}

// IsMasterSignerEnabled is a free data retrieval call binding the contract method 0xed536fdd.
//
// Solidity: function isMasterSignerEnabled() view returns(bool)
func (_AAWallet *AAWalletCaller) IsMasterSignerEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AAWallet.contract.Call(opts, &out, "isMasterSignerEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMasterSignerEnabled is a free data retrieval call binding the contract method 0xed536fdd.
//
// Solidity: function isMasterSignerEnabled() view returns(bool)
func (_AAWallet *AAWalletSession) IsMasterSignerEnabled() (bool, error) {
	return _AAWallet.Contract.IsMasterSignerEnabled(&_AAWallet.CallOpts)
}

// IsMasterSignerEnabled is a free data retrieval call binding the contract method 0xed536fdd.
//
// Solidity: function isMasterSignerEnabled() view returns(bool)
func (_AAWallet *AAWalletCallerSession) IsMasterSignerEnabled() (bool, error) {
	return _AAWallet.Contract.IsMasterSignerEnabled(&_AAWallet.CallOpts)
}

// MasterSigner is a free data retrieval call binding the contract method 0x20a3cd01.
//
// Solidity: function masterSigner() view returns(address)
func (_AAWallet *AAWalletCaller) MasterSigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AAWallet.contract.Call(opts, &out, "masterSigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MasterSigner is a free data retrieval call binding the contract method 0x20a3cd01.
//
// Solidity: function masterSigner() view returns(address)
func (_AAWallet *AAWalletSession) MasterSigner() (common.Address, error) {
	return _AAWallet.Contract.MasterSigner(&_AAWallet.CallOpts)
}

// MasterSigner is a free data retrieval call binding the contract method 0x20a3cd01.
//
// Solidity: function masterSigner() view returns(address)
func (_AAWallet *AAWalletCallerSession) MasterSigner() (common.Address, error) {
	return _AAWallet.Contract.MasterSigner(&_AAWallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AAWallet *AAWalletCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AAWallet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AAWallet *AAWalletSession) Owner() (common.Address, error) {
	return _AAWallet.Contract.Owner(&_AAWallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AAWallet *AAWalletCallerSession) Owner() (common.Address, error) {
	return _AAWallet.Contract.Owner(&_AAWallet.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AAWallet *AAWalletCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AAWallet.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AAWallet *AAWalletSession) ProxiableUUID() ([32]byte, error) {
	return _AAWallet.Contract.ProxiableUUID(&_AAWallet.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AAWallet *AAWalletCallerSession) ProxiableUUID() ([32]byte, error) {
	return _AAWallet.Contract.ProxiableUUID(&_AAWallet.CallOpts)
}

// SessionKeys is a free data retrieval call binding the contract method 0xb7b8d604.
//
// Solidity: function sessionKeys(address ) view returns(uint48 validAfter, uint48 validUntil, bytes32 permissions, bool isActive)
func (_AAWallet *AAWalletCaller) SessionKeys(opts *bind.CallOpts, arg0 common.Address) (struct {
	ValidAfter  *big.Int
	ValidUntil  *big.Int
	Permissions [32]byte
	IsActive    bool
}, error) {
	var out []interface{}
	err := _AAWallet.contract.Call(opts, &out, "sessionKeys", arg0)

	outstruct := new(struct {
		ValidAfter  *big.Int
		ValidUntil  *big.Int
		Permissions [32]byte
		IsActive    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ValidAfter = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ValidUntil = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Permissions = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.IsActive = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// SessionKeys is a free data retrieval call binding the contract method 0xb7b8d604.
//
// Solidity: function sessionKeys(address ) view returns(uint48 validAfter, uint48 validUntil, bytes32 permissions, bool isActive)
func (_AAWallet *AAWalletSession) SessionKeys(arg0 common.Address) (struct {
	ValidAfter  *big.Int
	ValidUntil  *big.Int
	Permissions [32]byte
	IsActive    bool
}, error) {
	return _AAWallet.Contract.SessionKeys(&_AAWallet.CallOpts, arg0)
}

// SessionKeys is a free data retrieval call binding the contract method 0xb7b8d604.
//
// Solidity: function sessionKeys(address ) view returns(uint48 validAfter, uint48 validUntil, bytes32 permissions, bool isActive)
func (_AAWallet *AAWalletCallerSession) SessionKeys(arg0 common.Address) (struct {
	ValidAfter  *big.Int
	ValidUntil  *big.Int
	Permissions [32]byte
	IsActive    bool
}, error) {
	return _AAWallet.Contract.SessionKeys(&_AAWallet.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AAWallet *AAWalletCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AAWallet.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AAWallet *AAWalletSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AAWallet.Contract.SupportsInterface(&_AAWallet.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AAWallet *AAWalletCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AAWallet.Contract.SupportsInterface(&_AAWallet.CallOpts, interfaceId)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_AAWallet *AAWalletTransactor) AddDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AAWallet.contract.Transact(opts, "addDeposit")
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_AAWallet *AAWalletSession) AddDeposit() (*types.Transaction, error) {
	return _AAWallet.Contract.AddDeposit(&_AAWallet.TransactOpts)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_AAWallet *AAWalletTransactorSession) AddDeposit() (*types.Transaction, error) {
	return _AAWallet.Contract.AddDeposit(&_AAWallet.TransactOpts)
}

// AddSessionKey is a paid mutator transaction binding the contract method 0x290f6a3f.
//
// Solidity: function addSessionKey(address sessionKey, uint48 validAfter, uint48 validUntil, bytes32 permissions) returns()
func (_AAWallet *AAWalletTransactor) AddSessionKey(opts *bind.TransactOpts, sessionKey common.Address, validAfter *big.Int, validUntil *big.Int, permissions [32]byte) (*types.Transaction, error) {
	return _AAWallet.contract.Transact(opts, "addSessionKey", sessionKey, validAfter, validUntil, permissions)
}

// AddSessionKey is a paid mutator transaction binding the contract method 0x290f6a3f.
//
// Solidity: function addSessionKey(address sessionKey, uint48 validAfter, uint48 validUntil, bytes32 permissions) returns()
func (_AAWallet *AAWalletSession) AddSessionKey(sessionKey common.Address, validAfter *big.Int, validUntil *big.Int, permissions [32]byte) (*types.Transaction, error) {
	return _AAWallet.Contract.AddSessionKey(&_AAWallet.TransactOpts, sessionKey, validAfter, validUntil, permissions)
}

// AddSessionKey is a paid mutator transaction binding the contract method 0x290f6a3f.
//
// Solidity: function addSessionKey(address sessionKey, uint48 validAfter, uint48 validUntil, bytes32 permissions) returns()
func (_AAWallet *AAWalletTransactorSession) AddSessionKey(sessionKey common.Address, validAfter *big.Int, validUntil *big.Int, permissions [32]byte) (*types.Transaction, error) {
	return _AAWallet.Contract.AddSessionKey(&_AAWallet.TransactOpts, sessionKey, validAfter, validUntil, permissions)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address target, uint256 value, bytes data) returns()
func (_AAWallet *AAWalletTransactor) Execute(opts *bind.TransactOpts, target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _AAWallet.contract.Transact(opts, "execute", target, value, data)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address target, uint256 value, bytes data) returns()
func (_AAWallet *AAWalletSession) Execute(target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _AAWallet.Contract.Execute(&_AAWallet.TransactOpts, target, value, data)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address target, uint256 value, bytes data) returns()
func (_AAWallet *AAWalletTransactorSession) Execute(target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _AAWallet.Contract.Execute(&_AAWallet.TransactOpts, target, value, data)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x34fcd5be.
//
// Solidity: function executeBatch((address,uint256,bytes)[] calls) returns()
func (_AAWallet *AAWalletTransactor) ExecuteBatch(opts *bind.TransactOpts, calls []BaseAccountCall) (*types.Transaction, error) {
	return _AAWallet.contract.Transact(opts, "executeBatch", calls)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x34fcd5be.
//
// Solidity: function executeBatch((address,uint256,bytes)[] calls) returns()
func (_AAWallet *AAWalletSession) ExecuteBatch(calls []BaseAccountCall) (*types.Transaction, error) {
	return _AAWallet.Contract.ExecuteBatch(&_AAWallet.TransactOpts, calls)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x34fcd5be.
//
// Solidity: function executeBatch((address,uint256,bytes)[] calls) returns()
func (_AAWallet *AAWalletTransactorSession) ExecuteBatch(calls []BaseAccountCall) (*types.Transaction, error) {
	return _AAWallet.Contract.ExecuteBatch(&_AAWallet.TransactOpts, calls)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _entryPoint, address _owner, address _masterSigner, address _aggregator) returns()
func (_AAWallet *AAWalletTransactor) Initialize(opts *bind.TransactOpts, _entryPoint common.Address, _owner common.Address, _masterSigner common.Address, _aggregator common.Address) (*types.Transaction, error) {
	return _AAWallet.contract.Transact(opts, "initialize", _entryPoint, _owner, _masterSigner, _aggregator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _entryPoint, address _owner, address _masterSigner, address _aggregator) returns()
func (_AAWallet *AAWalletSession) Initialize(_entryPoint common.Address, _owner common.Address, _masterSigner common.Address, _aggregator common.Address) (*types.Transaction, error) {
	return _AAWallet.Contract.Initialize(&_AAWallet.TransactOpts, _entryPoint, _owner, _masterSigner, _aggregator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _entryPoint, address _owner, address _masterSigner, address _aggregator) returns()
func (_AAWallet *AAWalletTransactorSession) Initialize(_entryPoint common.Address, _owner common.Address, _masterSigner common.Address, _aggregator common.Address) (*types.Transaction, error) {
	return _AAWallet.Contract.Initialize(&_AAWallet.TransactOpts, _entryPoint, _owner, _masterSigner, _aggregator)
}

// RevokeSessionKey is a paid mutator transaction binding the contract method 0x84f4fc6a.
//
// Solidity: function revokeSessionKey(address sessionKey) returns()
func (_AAWallet *AAWalletTransactor) RevokeSessionKey(opts *bind.TransactOpts, sessionKey common.Address) (*types.Transaction, error) {
	return _AAWallet.contract.Transact(opts, "revokeSessionKey", sessionKey)
}

// RevokeSessionKey is a paid mutator transaction binding the contract method 0x84f4fc6a.
//
// Solidity: function revokeSessionKey(address sessionKey) returns()
func (_AAWallet *AAWalletSession) RevokeSessionKey(sessionKey common.Address) (*types.Transaction, error) {
	return _AAWallet.Contract.RevokeSessionKey(&_AAWallet.TransactOpts, sessionKey)
}

// RevokeSessionKey is a paid mutator transaction binding the contract method 0x84f4fc6a.
//
// Solidity: function revokeSessionKey(address sessionKey) returns()
func (_AAWallet *AAWalletTransactorSession) RevokeSessionKey(sessionKey common.Address) (*types.Transaction, error) {
	return _AAWallet.Contract.RevokeSessionKey(&_AAWallet.TransactOpts, sessionKey)
}

// SetAggregator is a paid mutator transaction binding the contract method 0xf9120af6.
//
// Solidity: function setAggregator(address _aggregator) returns()
func (_AAWallet *AAWalletTransactor) SetAggregator(opts *bind.TransactOpts, _aggregator common.Address) (*types.Transaction, error) {
	return _AAWallet.contract.Transact(opts, "setAggregator", _aggregator)
}

// SetAggregator is a paid mutator transaction binding the contract method 0xf9120af6.
//
// Solidity: function setAggregator(address _aggregator) returns()
func (_AAWallet *AAWalletSession) SetAggregator(_aggregator common.Address) (*types.Transaction, error) {
	return _AAWallet.Contract.SetAggregator(&_AAWallet.TransactOpts, _aggregator)
}

// SetAggregator is a paid mutator transaction binding the contract method 0xf9120af6.
//
// Solidity: function setAggregator(address _aggregator) returns()
func (_AAWallet *AAWalletTransactorSession) SetAggregator(_aggregator common.Address) (*types.Transaction, error) {
	return _AAWallet.Contract.SetAggregator(&_AAWallet.TransactOpts, _aggregator)
}

// SetAuthorizedSessionKeyManager is a paid mutator transaction binding the contract method 0xe824561a.
//
// Solidity: function setAuthorizedSessionKeyManager(address _sessionKeyManager) returns()
func (_AAWallet *AAWalletTransactor) SetAuthorizedSessionKeyManager(opts *bind.TransactOpts, _sessionKeyManager common.Address) (*types.Transaction, error) {
	return _AAWallet.contract.Transact(opts, "setAuthorizedSessionKeyManager", _sessionKeyManager)
}

// SetAuthorizedSessionKeyManager is a paid mutator transaction binding the contract method 0xe824561a.
//
// Solidity: function setAuthorizedSessionKeyManager(address _sessionKeyManager) returns()
func (_AAWallet *AAWalletSession) SetAuthorizedSessionKeyManager(_sessionKeyManager common.Address) (*types.Transaction, error) {
	return _AAWallet.Contract.SetAuthorizedSessionKeyManager(&_AAWallet.TransactOpts, _sessionKeyManager)
}

// SetAuthorizedSessionKeyManager is a paid mutator transaction binding the contract method 0xe824561a.
//
// Solidity: function setAuthorizedSessionKeyManager(address _sessionKeyManager) returns()
func (_AAWallet *AAWalletTransactorSession) SetAuthorizedSessionKeyManager(_sessionKeyManager common.Address) (*types.Transaction, error) {
	return _AAWallet.Contract.SetAuthorizedSessionKeyManager(&_AAWallet.TransactOpts, _sessionKeyManager)
}

// SetMasterSigner is a paid mutator transaction binding the contract method 0x7c004623.
//
// Solidity: function setMasterSigner(address newMasterSigner) returns()
func (_AAWallet *AAWalletTransactor) SetMasterSigner(opts *bind.TransactOpts, newMasterSigner common.Address) (*types.Transaction, error) {
	return _AAWallet.contract.Transact(opts, "setMasterSigner", newMasterSigner)
}

// SetMasterSigner is a paid mutator transaction binding the contract method 0x7c004623.
//
// Solidity: function setMasterSigner(address newMasterSigner) returns()
func (_AAWallet *AAWalletSession) SetMasterSigner(newMasterSigner common.Address) (*types.Transaction, error) {
	return _AAWallet.Contract.SetMasterSigner(&_AAWallet.TransactOpts, newMasterSigner)
}

// SetMasterSigner is a paid mutator transaction binding the contract method 0x7c004623.
//
// Solidity: function setMasterSigner(address newMasterSigner) returns()
func (_AAWallet *AAWalletTransactorSession) SetMasterSigner(newMasterSigner common.Address) (*types.Transaction, error) {
	return _AAWallet.Contract.SetMasterSigner(&_AAWallet.TransactOpts, newMasterSigner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AAWallet *AAWalletTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AAWallet.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AAWallet *AAWalletSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AAWallet.Contract.UpgradeToAndCall(&_AAWallet.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AAWallet *AAWalletTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AAWallet.Contract.UpgradeToAndCall(&_AAWallet.TransactOpts, newImplementation, data)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_AAWallet *AAWalletTransactor) ValidateUserOp(opts *bind.TransactOpts, userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _AAWallet.contract.Transact(opts, "validateUserOp", userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_AAWallet *AAWalletSession) ValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _AAWallet.Contract.ValidateUserOp(&_AAWallet.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_AAWallet *AAWalletTransactorSession) ValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _AAWallet.Contract.ValidateUserOp(&_AAWallet.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_AAWallet *AAWalletTransactor) WithdrawDepositTo(opts *bind.TransactOpts, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AAWallet.contract.Transact(opts, "withdrawDepositTo", withdrawAddress, amount)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_AAWallet *AAWalletSession) WithdrawDepositTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AAWallet.Contract.WithdrawDepositTo(&_AAWallet.TransactOpts, withdrawAddress, amount)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_AAWallet *AAWalletTransactorSession) WithdrawDepositTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AAWallet.Contract.WithdrawDepositTo(&_AAWallet.TransactOpts, withdrawAddress, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AAWallet *AAWalletTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AAWallet.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AAWallet *AAWalletSession) Receive() (*types.Transaction, error) {
	return _AAWallet.Contract.Receive(&_AAWallet.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AAWallet *AAWalletTransactorSession) Receive() (*types.Transaction, error) {
	return _AAWallet.Contract.Receive(&_AAWallet.TransactOpts)
}

// AAWalletAAWalletInitializedIterator is returned from FilterAAWalletInitialized and is used to iterate over the raw logs and unpacked data for AAWalletInitialized events raised by the AAWallet contract.
type AAWalletAAWalletInitializedIterator struct {
	Event *AAWalletAAWalletInitialized // Event containing the contract specifics and raw log

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
func (it *AAWalletAAWalletInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAWalletAAWalletInitialized)
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
		it.Event = new(AAWalletAAWalletInitialized)
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
func (it *AAWalletAAWalletInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAWalletAAWalletInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAWalletAAWalletInitialized represents a AAWalletInitialized event raised by the AAWallet contract.
type AAWalletAAWalletInitialized struct {
	Owner        common.Address
	MasterSigner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAAWalletInitialized is a free log retrieval operation binding the contract event 0x94942efd10e1cddc29b9cd75631730f803a828d7b4e4b1acfe0ab15fc8687092.
//
// Solidity: event AAWalletInitialized(address indexed owner, address indexed masterSigner)
func (_AAWallet *AAWalletFilterer) FilterAAWalletInitialized(opts *bind.FilterOpts, owner []common.Address, masterSigner []common.Address) (*AAWalletAAWalletInitializedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var masterSignerRule []interface{}
	for _, masterSignerItem := range masterSigner {
		masterSignerRule = append(masterSignerRule, masterSignerItem)
	}

	logs, sub, err := _AAWallet.contract.FilterLogs(opts, "AAWalletInitialized", ownerRule, masterSignerRule)
	if err != nil {
		return nil, err
	}
	return &AAWalletAAWalletInitializedIterator{contract: _AAWallet.contract, event: "AAWalletInitialized", logs: logs, sub: sub}, nil
}

// WatchAAWalletInitialized is a free log subscription operation binding the contract event 0x94942efd10e1cddc29b9cd75631730f803a828d7b4e4b1acfe0ab15fc8687092.
//
// Solidity: event AAWalletInitialized(address indexed owner, address indexed masterSigner)
func (_AAWallet *AAWalletFilterer) WatchAAWalletInitialized(opts *bind.WatchOpts, sink chan<- *AAWalletAAWalletInitialized, owner []common.Address, masterSigner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var masterSignerRule []interface{}
	for _, masterSignerItem := range masterSigner {
		masterSignerRule = append(masterSignerRule, masterSignerItem)
	}

	logs, sub, err := _AAWallet.contract.WatchLogs(opts, "AAWalletInitialized", ownerRule, masterSignerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAWalletAAWalletInitialized)
				if err := _AAWallet.contract.UnpackLog(event, "AAWalletInitialized", log); err != nil {
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

// ParseAAWalletInitialized is a log parse operation binding the contract event 0x94942efd10e1cddc29b9cd75631730f803a828d7b4e4b1acfe0ab15fc8687092.
//
// Solidity: event AAWalletInitialized(address indexed owner, address indexed masterSigner)
func (_AAWallet *AAWalletFilterer) ParseAAWalletInitialized(log types.Log) (*AAWalletAAWalletInitialized, error) {
	event := new(AAWalletAAWalletInitialized)
	if err := _AAWallet.contract.UnpackLog(event, "AAWalletInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AAWalletAggregatorUpdatedIterator is returned from FilterAggregatorUpdated and is used to iterate over the raw logs and unpacked data for AggregatorUpdated events raised by the AAWallet contract.
type AAWalletAggregatorUpdatedIterator struct {
	Event *AAWalletAggregatorUpdated // Event containing the contract specifics and raw log

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
func (it *AAWalletAggregatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAWalletAggregatorUpdated)
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
		it.Event = new(AAWalletAggregatorUpdated)
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
func (it *AAWalletAggregatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAWalletAggregatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAWalletAggregatorUpdated represents a AggregatorUpdated event raised by the AAWallet contract.
type AAWalletAggregatorUpdated struct {
	OldAggregator common.Address
	NewAggregator common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAggregatorUpdated is a free log retrieval operation binding the contract event 0x89baabef7dfd0683c0ac16fd2a8431c51b49fbe654c3f7b5ef19763e2ccd88f2.
//
// Solidity: event AggregatorUpdated(address indexed oldAggregator, address indexed newAggregator)
func (_AAWallet *AAWalletFilterer) FilterAggregatorUpdated(opts *bind.FilterOpts, oldAggregator []common.Address, newAggregator []common.Address) (*AAWalletAggregatorUpdatedIterator, error) {

	var oldAggregatorRule []interface{}
	for _, oldAggregatorItem := range oldAggregator {
		oldAggregatorRule = append(oldAggregatorRule, oldAggregatorItem)
	}
	var newAggregatorRule []interface{}
	for _, newAggregatorItem := range newAggregator {
		newAggregatorRule = append(newAggregatorRule, newAggregatorItem)
	}

	logs, sub, err := _AAWallet.contract.FilterLogs(opts, "AggregatorUpdated", oldAggregatorRule, newAggregatorRule)
	if err != nil {
		return nil, err
	}
	return &AAWalletAggregatorUpdatedIterator{contract: _AAWallet.contract, event: "AggregatorUpdated", logs: logs, sub: sub}, nil
}

// WatchAggregatorUpdated is a free log subscription operation binding the contract event 0x89baabef7dfd0683c0ac16fd2a8431c51b49fbe654c3f7b5ef19763e2ccd88f2.
//
// Solidity: event AggregatorUpdated(address indexed oldAggregator, address indexed newAggregator)
func (_AAWallet *AAWalletFilterer) WatchAggregatorUpdated(opts *bind.WatchOpts, sink chan<- *AAWalletAggregatorUpdated, oldAggregator []common.Address, newAggregator []common.Address) (event.Subscription, error) {

	var oldAggregatorRule []interface{}
	for _, oldAggregatorItem := range oldAggregator {
		oldAggregatorRule = append(oldAggregatorRule, oldAggregatorItem)
	}
	var newAggregatorRule []interface{}
	for _, newAggregatorItem := range newAggregator {
		newAggregatorRule = append(newAggregatorRule, newAggregatorItem)
	}

	logs, sub, err := _AAWallet.contract.WatchLogs(opts, "AggregatorUpdated", oldAggregatorRule, newAggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAWalletAggregatorUpdated)
				if err := _AAWallet.contract.UnpackLog(event, "AggregatorUpdated", log); err != nil {
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

// ParseAggregatorUpdated is a log parse operation binding the contract event 0x89baabef7dfd0683c0ac16fd2a8431c51b49fbe654c3f7b5ef19763e2ccd88f2.
//
// Solidity: event AggregatorUpdated(address indexed oldAggregator, address indexed newAggregator)
func (_AAWallet *AAWalletFilterer) ParseAggregatorUpdated(log types.Log) (*AAWalletAggregatorUpdated, error) {
	event := new(AAWalletAggregatorUpdated)
	if err := _AAWallet.contract.UnpackLog(event, "AggregatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AAWalletInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AAWallet contract.
type AAWalletInitializedIterator struct {
	Event *AAWalletInitialized // Event containing the contract specifics and raw log

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
func (it *AAWalletInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAWalletInitialized)
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
		it.Event = new(AAWalletInitialized)
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
func (it *AAWalletInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAWalletInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAWalletInitialized represents a Initialized event raised by the AAWallet contract.
type AAWalletInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AAWallet *AAWalletFilterer) FilterInitialized(opts *bind.FilterOpts) (*AAWalletInitializedIterator, error) {

	logs, sub, err := _AAWallet.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AAWalletInitializedIterator{contract: _AAWallet.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AAWallet *AAWalletFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AAWalletInitialized) (event.Subscription, error) {

	logs, sub, err := _AAWallet.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAWalletInitialized)
				if err := _AAWallet.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_AAWallet *AAWalletFilterer) ParseInitialized(log types.Log) (*AAWalletInitialized, error) {
	event := new(AAWalletInitialized)
	if err := _AAWallet.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AAWalletMasterSignerUpdatedIterator is returned from FilterMasterSignerUpdated and is used to iterate over the raw logs and unpacked data for MasterSignerUpdated events raised by the AAWallet contract.
type AAWalletMasterSignerUpdatedIterator struct {
	Event *AAWalletMasterSignerUpdated // Event containing the contract specifics and raw log

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
func (it *AAWalletMasterSignerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAWalletMasterSignerUpdated)
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
		it.Event = new(AAWalletMasterSignerUpdated)
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
func (it *AAWalletMasterSignerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAWalletMasterSignerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAWalletMasterSignerUpdated represents a MasterSignerUpdated event raised by the AAWallet contract.
type AAWalletMasterSignerUpdated struct {
	OldMasterSigner common.Address
	NewMasterSigner common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMasterSignerUpdated is a free log retrieval operation binding the contract event 0x93cf965b971c37bb042b3c24d2278260742c6f2021d1c3be06984b6d03ab60ab.
//
// Solidity: event MasterSignerUpdated(address indexed oldMasterSigner, address indexed newMasterSigner)
func (_AAWallet *AAWalletFilterer) FilterMasterSignerUpdated(opts *bind.FilterOpts, oldMasterSigner []common.Address, newMasterSigner []common.Address) (*AAWalletMasterSignerUpdatedIterator, error) {

	var oldMasterSignerRule []interface{}
	for _, oldMasterSignerItem := range oldMasterSigner {
		oldMasterSignerRule = append(oldMasterSignerRule, oldMasterSignerItem)
	}
	var newMasterSignerRule []interface{}
	for _, newMasterSignerItem := range newMasterSigner {
		newMasterSignerRule = append(newMasterSignerRule, newMasterSignerItem)
	}

	logs, sub, err := _AAWallet.contract.FilterLogs(opts, "MasterSignerUpdated", oldMasterSignerRule, newMasterSignerRule)
	if err != nil {
		return nil, err
	}
	return &AAWalletMasterSignerUpdatedIterator{contract: _AAWallet.contract, event: "MasterSignerUpdated", logs: logs, sub: sub}, nil
}

// WatchMasterSignerUpdated is a free log subscription operation binding the contract event 0x93cf965b971c37bb042b3c24d2278260742c6f2021d1c3be06984b6d03ab60ab.
//
// Solidity: event MasterSignerUpdated(address indexed oldMasterSigner, address indexed newMasterSigner)
func (_AAWallet *AAWalletFilterer) WatchMasterSignerUpdated(opts *bind.WatchOpts, sink chan<- *AAWalletMasterSignerUpdated, oldMasterSigner []common.Address, newMasterSigner []common.Address) (event.Subscription, error) {

	var oldMasterSignerRule []interface{}
	for _, oldMasterSignerItem := range oldMasterSigner {
		oldMasterSignerRule = append(oldMasterSignerRule, oldMasterSignerItem)
	}
	var newMasterSignerRule []interface{}
	for _, newMasterSignerItem := range newMasterSigner {
		newMasterSignerRule = append(newMasterSignerRule, newMasterSignerItem)
	}

	logs, sub, err := _AAWallet.contract.WatchLogs(opts, "MasterSignerUpdated", oldMasterSignerRule, newMasterSignerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAWalletMasterSignerUpdated)
				if err := _AAWallet.contract.UnpackLog(event, "MasterSignerUpdated", log); err != nil {
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

// ParseMasterSignerUpdated is a log parse operation binding the contract event 0x93cf965b971c37bb042b3c24d2278260742c6f2021d1c3be06984b6d03ab60ab.
//
// Solidity: event MasterSignerUpdated(address indexed oldMasterSigner, address indexed newMasterSigner)
func (_AAWallet *AAWalletFilterer) ParseMasterSignerUpdated(log types.Log) (*AAWalletMasterSignerUpdated, error) {
	event := new(AAWalletMasterSignerUpdated)
	if err := _AAWallet.contract.UnpackLog(event, "MasterSignerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AAWalletMasterSignerValidationIterator is returned from FilterMasterSignerValidation and is used to iterate over the raw logs and unpacked data for MasterSignerValidation events raised by the AAWallet contract.
type AAWalletMasterSignerValidationIterator struct {
	Event *AAWalletMasterSignerValidation // Event containing the contract specifics and raw log

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
func (it *AAWalletMasterSignerValidationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAWalletMasterSignerValidation)
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
		it.Event = new(AAWalletMasterSignerValidation)
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
func (it *AAWalletMasterSignerValidationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAWalletMasterSignerValidationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAWalletMasterSignerValidation represents a MasterSignerValidation event raised by the AAWallet contract.
type AAWalletMasterSignerValidation struct {
	MasterSigner common.Address
	UserOpHash   [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMasterSignerValidation is a free log retrieval operation binding the contract event 0xb1ff5f74802a94e0e6c412076207eb7b19aad8b251a5304145e0bf8f15ca1ef3.
//
// Solidity: event MasterSignerValidation(address indexed masterSigner, bytes32 userOpHash)
func (_AAWallet *AAWalletFilterer) FilterMasterSignerValidation(opts *bind.FilterOpts, masterSigner []common.Address) (*AAWalletMasterSignerValidationIterator, error) {

	var masterSignerRule []interface{}
	for _, masterSignerItem := range masterSigner {
		masterSignerRule = append(masterSignerRule, masterSignerItem)
	}

	logs, sub, err := _AAWallet.contract.FilterLogs(opts, "MasterSignerValidation", masterSignerRule)
	if err != nil {
		return nil, err
	}
	return &AAWalletMasterSignerValidationIterator{contract: _AAWallet.contract, event: "MasterSignerValidation", logs: logs, sub: sub}, nil
}

// WatchMasterSignerValidation is a free log subscription operation binding the contract event 0xb1ff5f74802a94e0e6c412076207eb7b19aad8b251a5304145e0bf8f15ca1ef3.
//
// Solidity: event MasterSignerValidation(address indexed masterSigner, bytes32 userOpHash)
func (_AAWallet *AAWalletFilterer) WatchMasterSignerValidation(opts *bind.WatchOpts, sink chan<- *AAWalletMasterSignerValidation, masterSigner []common.Address) (event.Subscription, error) {

	var masterSignerRule []interface{}
	for _, masterSignerItem := range masterSigner {
		masterSignerRule = append(masterSignerRule, masterSignerItem)
	}

	logs, sub, err := _AAWallet.contract.WatchLogs(opts, "MasterSignerValidation", masterSignerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAWalletMasterSignerValidation)
				if err := _AAWallet.contract.UnpackLog(event, "MasterSignerValidation", log); err != nil {
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

// ParseMasterSignerValidation is a log parse operation binding the contract event 0xb1ff5f74802a94e0e6c412076207eb7b19aad8b251a5304145e0bf8f15ca1ef3.
//
// Solidity: event MasterSignerValidation(address indexed masterSigner, bytes32 userOpHash)
func (_AAWallet *AAWalletFilterer) ParseMasterSignerValidation(log types.Log) (*AAWalletMasterSignerValidation, error) {
	event := new(AAWalletMasterSignerValidation)
	if err := _AAWallet.contract.UnpackLog(event, "MasterSignerValidation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AAWalletSessionKeyAddedIterator is returned from FilterSessionKeyAdded and is used to iterate over the raw logs and unpacked data for SessionKeyAdded events raised by the AAWallet contract.
type AAWalletSessionKeyAddedIterator struct {
	Event *AAWalletSessionKeyAdded // Event containing the contract specifics and raw log

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
func (it *AAWalletSessionKeyAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAWalletSessionKeyAdded)
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
		it.Event = new(AAWalletSessionKeyAdded)
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
func (it *AAWalletSessionKeyAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAWalletSessionKeyAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAWalletSessionKeyAdded represents a SessionKeyAdded event raised by the AAWallet contract.
type AAWalletSessionKeyAdded struct {
	SessionKey  common.Address
	ValidAfter  *big.Int
	ValidUntil  *big.Int
	Permissions [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSessionKeyAdded is a free log retrieval operation binding the contract event 0xa00bca54e0f01b47f85e9dd47c72f5921c1e7a1541b84f3888acf901a424e94b.
//
// Solidity: event SessionKeyAdded(address indexed sessionKey, uint48 validAfter, uint48 validUntil, bytes32 permissions)
func (_AAWallet *AAWalletFilterer) FilterSessionKeyAdded(opts *bind.FilterOpts, sessionKey []common.Address) (*AAWalletSessionKeyAddedIterator, error) {

	var sessionKeyRule []interface{}
	for _, sessionKeyItem := range sessionKey {
		sessionKeyRule = append(sessionKeyRule, sessionKeyItem)
	}

	logs, sub, err := _AAWallet.contract.FilterLogs(opts, "SessionKeyAdded", sessionKeyRule)
	if err != nil {
		return nil, err
	}
	return &AAWalletSessionKeyAddedIterator{contract: _AAWallet.contract, event: "SessionKeyAdded", logs: logs, sub: sub}, nil
}

// WatchSessionKeyAdded is a free log subscription operation binding the contract event 0xa00bca54e0f01b47f85e9dd47c72f5921c1e7a1541b84f3888acf901a424e94b.
//
// Solidity: event SessionKeyAdded(address indexed sessionKey, uint48 validAfter, uint48 validUntil, bytes32 permissions)
func (_AAWallet *AAWalletFilterer) WatchSessionKeyAdded(opts *bind.WatchOpts, sink chan<- *AAWalletSessionKeyAdded, sessionKey []common.Address) (event.Subscription, error) {

	var sessionKeyRule []interface{}
	for _, sessionKeyItem := range sessionKey {
		sessionKeyRule = append(sessionKeyRule, sessionKeyItem)
	}

	logs, sub, err := _AAWallet.contract.WatchLogs(opts, "SessionKeyAdded", sessionKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAWalletSessionKeyAdded)
				if err := _AAWallet.contract.UnpackLog(event, "SessionKeyAdded", log); err != nil {
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

// ParseSessionKeyAdded is a log parse operation binding the contract event 0xa00bca54e0f01b47f85e9dd47c72f5921c1e7a1541b84f3888acf901a424e94b.
//
// Solidity: event SessionKeyAdded(address indexed sessionKey, uint48 validAfter, uint48 validUntil, bytes32 permissions)
func (_AAWallet *AAWalletFilterer) ParseSessionKeyAdded(log types.Log) (*AAWalletSessionKeyAdded, error) {
	event := new(AAWalletSessionKeyAdded)
	if err := _AAWallet.contract.UnpackLog(event, "SessionKeyAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AAWalletSessionKeyRevokedIterator is returned from FilterSessionKeyRevoked and is used to iterate over the raw logs and unpacked data for SessionKeyRevoked events raised by the AAWallet contract.
type AAWalletSessionKeyRevokedIterator struct {
	Event *AAWalletSessionKeyRevoked // Event containing the contract specifics and raw log

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
func (it *AAWalletSessionKeyRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAWalletSessionKeyRevoked)
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
		it.Event = new(AAWalletSessionKeyRevoked)
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
func (it *AAWalletSessionKeyRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAWalletSessionKeyRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAWalletSessionKeyRevoked represents a SessionKeyRevoked event raised by the AAWallet contract.
type AAWalletSessionKeyRevoked struct {
	SessionKey common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSessionKeyRevoked is a free log retrieval operation binding the contract event 0x17c796fb82086b3c9effaec517342e5ca9ed8fd78c339137ec082f748ab60cbe.
//
// Solidity: event SessionKeyRevoked(address indexed sessionKey)
func (_AAWallet *AAWalletFilterer) FilterSessionKeyRevoked(opts *bind.FilterOpts, sessionKey []common.Address) (*AAWalletSessionKeyRevokedIterator, error) {

	var sessionKeyRule []interface{}
	for _, sessionKeyItem := range sessionKey {
		sessionKeyRule = append(sessionKeyRule, sessionKeyItem)
	}

	logs, sub, err := _AAWallet.contract.FilterLogs(opts, "SessionKeyRevoked", sessionKeyRule)
	if err != nil {
		return nil, err
	}
	return &AAWalletSessionKeyRevokedIterator{contract: _AAWallet.contract, event: "SessionKeyRevoked", logs: logs, sub: sub}, nil
}

// WatchSessionKeyRevoked is a free log subscription operation binding the contract event 0x17c796fb82086b3c9effaec517342e5ca9ed8fd78c339137ec082f748ab60cbe.
//
// Solidity: event SessionKeyRevoked(address indexed sessionKey)
func (_AAWallet *AAWalletFilterer) WatchSessionKeyRevoked(opts *bind.WatchOpts, sink chan<- *AAWalletSessionKeyRevoked, sessionKey []common.Address) (event.Subscription, error) {

	var sessionKeyRule []interface{}
	for _, sessionKeyItem := range sessionKey {
		sessionKeyRule = append(sessionKeyRule, sessionKeyItem)
	}

	logs, sub, err := _AAWallet.contract.WatchLogs(opts, "SessionKeyRevoked", sessionKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAWalletSessionKeyRevoked)
				if err := _AAWallet.contract.UnpackLog(event, "SessionKeyRevoked", log); err != nil {
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

// ParseSessionKeyRevoked is a log parse operation binding the contract event 0x17c796fb82086b3c9effaec517342e5ca9ed8fd78c339137ec082f748ab60cbe.
//
// Solidity: event SessionKeyRevoked(address indexed sessionKey)
func (_AAWallet *AAWalletFilterer) ParseSessionKeyRevoked(log types.Log) (*AAWalletSessionKeyRevoked, error) {
	event := new(AAWalletSessionKeyRevoked)
	if err := _AAWallet.contract.UnpackLog(event, "SessionKeyRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AAWalletSessionKeyValidationIterator is returned from FilterSessionKeyValidation and is used to iterate over the raw logs and unpacked data for SessionKeyValidation events raised by the AAWallet contract.
type AAWalletSessionKeyValidationIterator struct {
	Event *AAWalletSessionKeyValidation // Event containing the contract specifics and raw log

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
func (it *AAWalletSessionKeyValidationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAWalletSessionKeyValidation)
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
		it.Event = new(AAWalletSessionKeyValidation)
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
func (it *AAWalletSessionKeyValidationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAWalletSessionKeyValidationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAWalletSessionKeyValidation represents a SessionKeyValidation event raised by the AAWallet contract.
type AAWalletSessionKeyValidation struct {
	SessionKey common.Address
	UserOpHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSessionKeyValidation is a free log retrieval operation binding the contract event 0x0c8028b60a408027126de0b307f68042b0a366d216ac9ad715fed951c9ad1d09.
//
// Solidity: event SessionKeyValidation(address indexed sessionKey, bytes32 userOpHash)
func (_AAWallet *AAWalletFilterer) FilterSessionKeyValidation(opts *bind.FilterOpts, sessionKey []common.Address) (*AAWalletSessionKeyValidationIterator, error) {

	var sessionKeyRule []interface{}
	for _, sessionKeyItem := range sessionKey {
		sessionKeyRule = append(sessionKeyRule, sessionKeyItem)
	}

	logs, sub, err := _AAWallet.contract.FilterLogs(opts, "SessionKeyValidation", sessionKeyRule)
	if err != nil {
		return nil, err
	}
	return &AAWalletSessionKeyValidationIterator{contract: _AAWallet.contract, event: "SessionKeyValidation", logs: logs, sub: sub}, nil
}

// WatchSessionKeyValidation is a free log subscription operation binding the contract event 0x0c8028b60a408027126de0b307f68042b0a366d216ac9ad715fed951c9ad1d09.
//
// Solidity: event SessionKeyValidation(address indexed sessionKey, bytes32 userOpHash)
func (_AAWallet *AAWalletFilterer) WatchSessionKeyValidation(opts *bind.WatchOpts, sink chan<- *AAWalletSessionKeyValidation, sessionKey []common.Address) (event.Subscription, error) {

	var sessionKeyRule []interface{}
	for _, sessionKeyItem := range sessionKey {
		sessionKeyRule = append(sessionKeyRule, sessionKeyItem)
	}

	logs, sub, err := _AAWallet.contract.WatchLogs(opts, "SessionKeyValidation", sessionKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAWalletSessionKeyValidation)
				if err := _AAWallet.contract.UnpackLog(event, "SessionKeyValidation", log); err != nil {
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

// ParseSessionKeyValidation is a log parse operation binding the contract event 0x0c8028b60a408027126de0b307f68042b0a366d216ac9ad715fed951c9ad1d09.
//
// Solidity: event SessionKeyValidation(address indexed sessionKey, bytes32 userOpHash)
func (_AAWallet *AAWalletFilterer) ParseSessionKeyValidation(log types.Log) (*AAWalletSessionKeyValidation, error) {
	event := new(AAWalletSessionKeyValidation)
	if err := _AAWallet.contract.UnpackLog(event, "SessionKeyValidation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AAWalletUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the AAWallet contract.
type AAWalletUpgradedIterator struct {
	Event *AAWalletUpgraded // Event containing the contract specifics and raw log

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
func (it *AAWalletUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAWalletUpgraded)
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
		it.Event = new(AAWalletUpgraded)
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
func (it *AAWalletUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAWalletUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAWalletUpgraded represents a Upgraded event raised by the AAWallet contract.
type AAWalletUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AAWallet *AAWalletFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*AAWalletUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AAWallet.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &AAWalletUpgradedIterator{contract: _AAWallet.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AAWallet *AAWalletFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *AAWalletUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AAWallet.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAWalletUpgraded)
				if err := _AAWallet.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_AAWallet *AAWalletFilterer) ParseUpgraded(log types.Log) (*AAWalletUpgraded, error) {
	event := new(AAWalletUpgraded)
	if err := _AAWallet.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
