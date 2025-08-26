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

// AAccountMetaData contains all meta data concerning the AAccount contract.
var AAccountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ExecuteError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"masterSigner\",\"type\":\"address\"}],\"name\":\"AAccountInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAggregator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAggregator\",\"type\":\"address\"}],\"name\":\"AggregatorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldMasterSigner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newMasterSigner\",\"type\":\"address\"}],\"name\":\"MasterSignerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"masterSigner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"}],\"name\":\"MasterSignerValidation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"validAfter\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"validUntil\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"permissions\",\"type\":\"bytes32\"}],\"name\":\"SessionKeyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"}],\"name\":\"SessionKeyRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"}],\"name\":\"SessionKeyValidation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"validAfter\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"validUntil\",\"type\":\"uint48\"},{\"internalType\":\"bytes32\",\"name\":\"permissions\",\"type\":\"bytes32\"}],\"name\":\"addSessionKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggregatorAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"authorizedSessionKeyManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"canSessionKeyExecute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"canExecute\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPointAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structBaseAccount.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"executeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMasterSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"}],\"name\":\"getSessionKeyInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"},{\"internalType\":\"uint48\",\"name\":\"validAfter\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"validUntil\",\"type\":\"uint48\"},{\"internalType\":\"bytes32\",\"name\":\"permissions\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_entryPoint\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_masterSigner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMasterSignerEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"masterSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"}],\"name\":\"revokeSessionKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sessionKeys\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"validAfter\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"validUntil\",\"type\":\"uint48\"},{\"internalType\":\"bytes32\",\"name\":\"permissions\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"}],\"name\":\"setAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sessionKeyManager\",\"type\":\"address\"}],\"name\":\"setAuthorizedSessionKeyManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newMasterSigner\",\"type\":\"address\"}],\"name\":\"setMasterSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"missingAccountFunds\",\"type\":\"uint256\"}],\"name\":\"validateUserOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawDepositTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// AAccountABI is the input ABI used to generate the binding from.
// Deprecated: Use AAccountMetaData.ABI instead.
var AAccountABI = AAccountMetaData.ABI

// AAccount is an auto generated Go binding around an Ethereum contract.
type AAccount struct {
	AAccountCaller     // Read-only binding to the contract
	AAccountTransactor // Write-only binding to the contract
	AAccountFilterer   // Log filterer for contract events
}

// AAccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type AAccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AAccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AAccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AAccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AAccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AAccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AAccountSession struct {
	Contract     *AAccount         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AAccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AAccountCallerSession struct {
	Contract *AAccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// AAccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AAccountTransactorSession struct {
	Contract     *AAccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AAccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type AAccountRaw struct {
	Contract *AAccount // Generic contract binding to access the raw methods on
}

// AAccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AAccountCallerRaw struct {
	Contract *AAccountCaller // Generic read-only contract binding to access the raw methods on
}

// AAccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AAccountTransactorRaw struct {
	Contract *AAccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAAccount creates a new instance of AAccount, bound to a specific deployed contract.
func NewAAccount(address common.Address, backend bind.ContractBackend) (*AAccount, error) {
	contract, err := bindAAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AAccount{AAccountCaller: AAccountCaller{contract: contract}, AAccountTransactor: AAccountTransactor{contract: contract}, AAccountFilterer: AAccountFilterer{contract: contract}}, nil
}

// NewAAccountCaller creates a new read-only instance of AAccount, bound to a specific deployed contract.
func NewAAccountCaller(address common.Address, caller bind.ContractCaller) (*AAccountCaller, error) {
	contract, err := bindAAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AAccountCaller{contract: contract}, nil
}

// NewAAccountTransactor creates a new write-only instance of AAccount, bound to a specific deployed contract.
func NewAAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*AAccountTransactor, error) {
	contract, err := bindAAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AAccountTransactor{contract: contract}, nil
}

// NewAAccountFilterer creates a new log filterer instance of AAccount, bound to a specific deployed contract.
func NewAAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*AAccountFilterer, error) {
	contract, err := bindAAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AAccountFilterer{contract: contract}, nil
}

// bindAAccount binds a generic wrapper to an already deployed contract.
func bindAAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AAccountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AAccount *AAccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AAccount.Contract.AAccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AAccount *AAccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AAccount.Contract.AAccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AAccount *AAccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AAccount.Contract.AAccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AAccount *AAccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AAccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AAccount *AAccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AAccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AAccount *AAccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AAccount.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_AAccount *AAccountCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AAccount.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_AAccount *AAccountSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _AAccount.Contract.UPGRADEINTERFACEVERSION(&_AAccount.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_AAccount *AAccountCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _AAccount.Contract.UPGRADEINTERFACEVERSION(&_AAccount.CallOpts)
}

// AggregatorAddress is a free data retrieval call binding the contract method 0x380bbe53.
//
// Solidity: function aggregatorAddress() view returns(address)
func (_AAccount *AAccountCaller) AggregatorAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AAccount.contract.Call(opts, &out, "aggregatorAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AggregatorAddress is a free data retrieval call binding the contract method 0x380bbe53.
//
// Solidity: function aggregatorAddress() view returns(address)
func (_AAccount *AAccountSession) AggregatorAddress() (common.Address, error) {
	return _AAccount.Contract.AggregatorAddress(&_AAccount.CallOpts)
}

// AggregatorAddress is a free data retrieval call binding the contract method 0x380bbe53.
//
// Solidity: function aggregatorAddress() view returns(address)
func (_AAccount *AAccountCallerSession) AggregatorAddress() (common.Address, error) {
	return _AAccount.Contract.AggregatorAddress(&_AAccount.CallOpts)
}

// AuthorizedSessionKeyManager is a free data retrieval call binding the contract method 0x4b6c4f0b.
//
// Solidity: function authorizedSessionKeyManager() view returns(address)
func (_AAccount *AAccountCaller) AuthorizedSessionKeyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AAccount.contract.Call(opts, &out, "authorizedSessionKeyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AuthorizedSessionKeyManager is a free data retrieval call binding the contract method 0x4b6c4f0b.
//
// Solidity: function authorizedSessionKeyManager() view returns(address)
func (_AAccount *AAccountSession) AuthorizedSessionKeyManager() (common.Address, error) {
	return _AAccount.Contract.AuthorizedSessionKeyManager(&_AAccount.CallOpts)
}

// AuthorizedSessionKeyManager is a free data retrieval call binding the contract method 0x4b6c4f0b.
//
// Solidity: function authorizedSessionKeyManager() view returns(address)
func (_AAccount *AAccountCallerSession) AuthorizedSessionKeyManager() (common.Address, error) {
	return _AAccount.Contract.AuthorizedSessionKeyManager(&_AAccount.CallOpts)
}

// CanSessionKeyExecute is a free data retrieval call binding the contract method 0x2c4d8fcf.
//
// Solidity: function canSessionKeyExecute(address sessionKey, address target, bytes4 selector) view returns(bool canExecute)
func (_AAccount *AAccountCaller) CanSessionKeyExecute(opts *bind.CallOpts, sessionKey common.Address, target common.Address, selector [4]byte) (bool, error) {
	var out []interface{}
	err := _AAccount.contract.Call(opts, &out, "canSessionKeyExecute", sessionKey, target, selector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanSessionKeyExecute is a free data retrieval call binding the contract method 0x2c4d8fcf.
//
// Solidity: function canSessionKeyExecute(address sessionKey, address target, bytes4 selector) view returns(bool canExecute)
func (_AAccount *AAccountSession) CanSessionKeyExecute(sessionKey common.Address, target common.Address, selector [4]byte) (bool, error) {
	return _AAccount.Contract.CanSessionKeyExecute(&_AAccount.CallOpts, sessionKey, target, selector)
}

// CanSessionKeyExecute is a free data retrieval call binding the contract method 0x2c4d8fcf.
//
// Solidity: function canSessionKeyExecute(address sessionKey, address target, bytes4 selector) view returns(bool canExecute)
func (_AAccount *AAccountCallerSession) CanSessionKeyExecute(sessionKey common.Address, target common.Address, selector [4]byte) (bool, error) {
	return _AAccount.Contract.CanSessionKeyExecute(&_AAccount.CallOpts, sessionKey, target, selector)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_AAccount *AAccountCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AAccount.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_AAccount *AAccountSession) EntryPoint() (common.Address, error) {
	return _AAccount.Contract.EntryPoint(&_AAccount.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_AAccount *AAccountCallerSession) EntryPoint() (common.Address, error) {
	return _AAccount.Contract.EntryPoint(&_AAccount.CallOpts)
}

// EntryPointAddress is a free data retrieval call binding the contract method 0x06dc245c.
//
// Solidity: function entryPointAddress() view returns(address)
func (_AAccount *AAccountCaller) EntryPointAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AAccount.contract.Call(opts, &out, "entryPointAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPointAddress is a free data retrieval call binding the contract method 0x06dc245c.
//
// Solidity: function entryPointAddress() view returns(address)
func (_AAccount *AAccountSession) EntryPointAddress() (common.Address, error) {
	return _AAccount.Contract.EntryPointAddress(&_AAccount.CallOpts)
}

// EntryPointAddress is a free data retrieval call binding the contract method 0x06dc245c.
//
// Solidity: function entryPointAddress() view returns(address)
func (_AAccount *AAccountCallerSession) EntryPointAddress() (common.Address, error) {
	return _AAccount.Contract.EntryPointAddress(&_AAccount.CallOpts)
}

// GetAggregator is a free data retrieval call binding the contract method 0x3ad59dbc.
//
// Solidity: function getAggregator() view returns(address aggregator)
func (_AAccount *AAccountCaller) GetAggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AAccount.contract.Call(opts, &out, "getAggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAggregator is a free data retrieval call binding the contract method 0x3ad59dbc.
//
// Solidity: function getAggregator() view returns(address aggregator)
func (_AAccount *AAccountSession) GetAggregator() (common.Address, error) {
	return _AAccount.Contract.GetAggregator(&_AAccount.CallOpts)
}

// GetAggregator is a free data retrieval call binding the contract method 0x3ad59dbc.
//
// Solidity: function getAggregator() view returns(address aggregator)
func (_AAccount *AAccountCallerSession) GetAggregator() (common.Address, error) {
	return _AAccount.Contract.GetAggregator(&_AAccount.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_AAccount *AAccountCaller) GetDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AAccount.contract.Call(opts, &out, "getDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_AAccount *AAccountSession) GetDeposit() (*big.Int, error) {
	return _AAccount.Contract.GetDeposit(&_AAccount.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_AAccount *AAccountCallerSession) GetDeposit() (*big.Int, error) {
	return _AAccount.Contract.GetDeposit(&_AAccount.CallOpts)
}

// GetMasterSigner is a free data retrieval call binding the contract method 0x0b2f4d9e.
//
// Solidity: function getMasterSigner() view returns(address)
func (_AAccount *AAccountCaller) GetMasterSigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AAccount.contract.Call(opts, &out, "getMasterSigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMasterSigner is a free data retrieval call binding the contract method 0x0b2f4d9e.
//
// Solidity: function getMasterSigner() view returns(address)
func (_AAccount *AAccountSession) GetMasterSigner() (common.Address, error) {
	return _AAccount.Contract.GetMasterSigner(&_AAccount.CallOpts)
}

// GetMasterSigner is a free data retrieval call binding the contract method 0x0b2f4d9e.
//
// Solidity: function getMasterSigner() view returns(address)
func (_AAccount *AAccountCallerSession) GetMasterSigner() (common.Address, error) {
	return _AAccount.Contract.GetMasterSigner(&_AAccount.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_AAccount *AAccountCaller) GetNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AAccount.contract.Call(opts, &out, "getNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_AAccount *AAccountSession) GetNonce() (*big.Int, error) {
	return _AAccount.Contract.GetNonce(&_AAccount.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_AAccount *AAccountCallerSession) GetNonce() (*big.Int, error) {
	return _AAccount.Contract.GetNonce(&_AAccount.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_AAccount *AAccountCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AAccount.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_AAccount *AAccountSession) GetOwner() (common.Address, error) {
	return _AAccount.Contract.GetOwner(&_AAccount.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_AAccount *AAccountCallerSession) GetOwner() (common.Address, error) {
	return _AAccount.Contract.GetOwner(&_AAccount.CallOpts)
}

// GetSessionKeyInfo is a free data retrieval call binding the contract method 0x5ad2657f.
//
// Solidity: function getSessionKeyInfo(address sessionKey) view returns(bool isValid, uint48 validAfter, uint48 validUntil, bytes32 permissions)
func (_AAccount *AAccountCaller) GetSessionKeyInfo(opts *bind.CallOpts, sessionKey common.Address) (struct {
	IsValid     bool
	ValidAfter  *big.Int
	ValidUntil  *big.Int
	Permissions [32]byte
}, error) {
	var out []interface{}
	err := _AAccount.contract.Call(opts, &out, "getSessionKeyInfo", sessionKey)

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
func (_AAccount *AAccountSession) GetSessionKeyInfo(sessionKey common.Address) (struct {
	IsValid     bool
	ValidAfter  *big.Int
	ValidUntil  *big.Int
	Permissions [32]byte
}, error) {
	return _AAccount.Contract.GetSessionKeyInfo(&_AAccount.CallOpts, sessionKey)
}

// GetSessionKeyInfo is a free data retrieval call binding the contract method 0x5ad2657f.
//
// Solidity: function getSessionKeyInfo(address sessionKey) view returns(bool isValid, uint48 validAfter, uint48 validUntil, bytes32 permissions)
func (_AAccount *AAccountCallerSession) GetSessionKeyInfo(sessionKey common.Address) (struct {
	IsValid     bool
	ValidAfter  *big.Int
	ValidUntil  *big.Int
	Permissions [32]byte
}, error) {
	return _AAccount.Contract.GetSessionKeyInfo(&_AAccount.CallOpts, sessionKey)
}

// IsMasterSignerEnabled is a free data retrieval call binding the contract method 0xed536fdd.
//
// Solidity: function isMasterSignerEnabled() view returns(bool)
func (_AAccount *AAccountCaller) IsMasterSignerEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AAccount.contract.Call(opts, &out, "isMasterSignerEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMasterSignerEnabled is a free data retrieval call binding the contract method 0xed536fdd.
//
// Solidity: function isMasterSignerEnabled() view returns(bool)
func (_AAccount *AAccountSession) IsMasterSignerEnabled() (bool, error) {
	return _AAccount.Contract.IsMasterSignerEnabled(&_AAccount.CallOpts)
}

// IsMasterSignerEnabled is a free data retrieval call binding the contract method 0xed536fdd.
//
// Solidity: function isMasterSignerEnabled() view returns(bool)
func (_AAccount *AAccountCallerSession) IsMasterSignerEnabled() (bool, error) {
	return _AAccount.Contract.IsMasterSignerEnabled(&_AAccount.CallOpts)
}

// MasterSigner is a free data retrieval call binding the contract method 0x20a3cd01.
//
// Solidity: function masterSigner() view returns(address)
func (_AAccount *AAccountCaller) MasterSigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AAccount.contract.Call(opts, &out, "masterSigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MasterSigner is a free data retrieval call binding the contract method 0x20a3cd01.
//
// Solidity: function masterSigner() view returns(address)
func (_AAccount *AAccountSession) MasterSigner() (common.Address, error) {
	return _AAccount.Contract.MasterSigner(&_AAccount.CallOpts)
}

// MasterSigner is a free data retrieval call binding the contract method 0x20a3cd01.
//
// Solidity: function masterSigner() view returns(address)
func (_AAccount *AAccountCallerSession) MasterSigner() (common.Address, error) {
	return _AAccount.Contract.MasterSigner(&_AAccount.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AAccount *AAccountCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AAccount.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AAccount *AAccountSession) Owner() (common.Address, error) {
	return _AAccount.Contract.Owner(&_AAccount.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AAccount *AAccountCallerSession) Owner() (common.Address, error) {
	return _AAccount.Contract.Owner(&_AAccount.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AAccount *AAccountCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AAccount.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AAccount *AAccountSession) ProxiableUUID() ([32]byte, error) {
	return _AAccount.Contract.ProxiableUUID(&_AAccount.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AAccount *AAccountCallerSession) ProxiableUUID() ([32]byte, error) {
	return _AAccount.Contract.ProxiableUUID(&_AAccount.CallOpts)
}

// SessionKeys is a free data retrieval call binding the contract method 0xb7b8d604.
//
// Solidity: function sessionKeys(address ) view returns(uint48 validAfter, uint48 validUntil, bytes32 permissions, bool isActive)
func (_AAccount *AAccountCaller) SessionKeys(opts *bind.CallOpts, arg0 common.Address) (struct {
	ValidAfter  *big.Int
	ValidUntil  *big.Int
	Permissions [32]byte
	IsActive    bool
}, error) {
	var out []interface{}
	err := _AAccount.contract.Call(opts, &out, "sessionKeys", arg0)

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
func (_AAccount *AAccountSession) SessionKeys(arg0 common.Address) (struct {
	ValidAfter  *big.Int
	ValidUntil  *big.Int
	Permissions [32]byte
	IsActive    bool
}, error) {
	return _AAccount.Contract.SessionKeys(&_AAccount.CallOpts, arg0)
}

// SessionKeys is a free data retrieval call binding the contract method 0xb7b8d604.
//
// Solidity: function sessionKeys(address ) view returns(uint48 validAfter, uint48 validUntil, bytes32 permissions, bool isActive)
func (_AAccount *AAccountCallerSession) SessionKeys(arg0 common.Address) (struct {
	ValidAfter  *big.Int
	ValidUntil  *big.Int
	Permissions [32]byte
	IsActive    bool
}, error) {
	return _AAccount.Contract.SessionKeys(&_AAccount.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AAccount *AAccountCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AAccount.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AAccount *AAccountSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AAccount.Contract.SupportsInterface(&_AAccount.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AAccount *AAccountCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AAccount.Contract.SupportsInterface(&_AAccount.CallOpts, interfaceId)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_AAccount *AAccountTransactor) AddDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AAccount.contract.Transact(opts, "addDeposit")
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_AAccount *AAccountSession) AddDeposit() (*types.Transaction, error) {
	return _AAccount.Contract.AddDeposit(&_AAccount.TransactOpts)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_AAccount *AAccountTransactorSession) AddDeposit() (*types.Transaction, error) {
	return _AAccount.Contract.AddDeposit(&_AAccount.TransactOpts)
}

// AddSessionKey is a paid mutator transaction binding the contract method 0x290f6a3f.
//
// Solidity: function addSessionKey(address sessionKey, uint48 validAfter, uint48 validUntil, bytes32 permissions) returns()
func (_AAccount *AAccountTransactor) AddSessionKey(opts *bind.TransactOpts, sessionKey common.Address, validAfter *big.Int, validUntil *big.Int, permissions [32]byte) (*types.Transaction, error) {
	return _AAccount.contract.Transact(opts, "addSessionKey", sessionKey, validAfter, validUntil, permissions)
}

// AddSessionKey is a paid mutator transaction binding the contract method 0x290f6a3f.
//
// Solidity: function addSessionKey(address sessionKey, uint48 validAfter, uint48 validUntil, bytes32 permissions) returns()
func (_AAccount *AAccountSession) AddSessionKey(sessionKey common.Address, validAfter *big.Int, validUntil *big.Int, permissions [32]byte) (*types.Transaction, error) {
	return _AAccount.Contract.AddSessionKey(&_AAccount.TransactOpts, sessionKey, validAfter, validUntil, permissions)
}

// AddSessionKey is a paid mutator transaction binding the contract method 0x290f6a3f.
//
// Solidity: function addSessionKey(address sessionKey, uint48 validAfter, uint48 validUntil, bytes32 permissions) returns()
func (_AAccount *AAccountTransactorSession) AddSessionKey(sessionKey common.Address, validAfter *big.Int, validUntil *big.Int, permissions [32]byte) (*types.Transaction, error) {
	return _AAccount.Contract.AddSessionKey(&_AAccount.TransactOpts, sessionKey, validAfter, validUntil, permissions)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address target, uint256 value, bytes data) returns()
func (_AAccount *AAccountTransactor) Execute(opts *bind.TransactOpts, target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _AAccount.contract.Transact(opts, "execute", target, value, data)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address target, uint256 value, bytes data) returns()
func (_AAccount *AAccountSession) Execute(target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _AAccount.Contract.Execute(&_AAccount.TransactOpts, target, value, data)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address target, uint256 value, bytes data) returns()
func (_AAccount *AAccountTransactorSession) Execute(target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _AAccount.Contract.Execute(&_AAccount.TransactOpts, target, value, data)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x34fcd5be.
//
// Solidity: function executeBatch((address,uint256,bytes)[] calls) returns()
func (_AAccount *AAccountTransactor) ExecuteBatch(opts *bind.TransactOpts, calls []BaseAccountCall) (*types.Transaction, error) {
	return _AAccount.contract.Transact(opts, "executeBatch", calls)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x34fcd5be.
//
// Solidity: function executeBatch((address,uint256,bytes)[] calls) returns()
func (_AAccount *AAccountSession) ExecuteBatch(calls []BaseAccountCall) (*types.Transaction, error) {
	return _AAccount.Contract.ExecuteBatch(&_AAccount.TransactOpts, calls)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x34fcd5be.
//
// Solidity: function executeBatch((address,uint256,bytes)[] calls) returns()
func (_AAccount *AAccountTransactorSession) ExecuteBatch(calls []BaseAccountCall) (*types.Transaction, error) {
	return _AAccount.Contract.ExecuteBatch(&_AAccount.TransactOpts, calls)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _entryPoint, address _owner, address _masterSigner, address _aggregator) returns()
func (_AAccount *AAccountTransactor) Initialize(opts *bind.TransactOpts, _entryPoint common.Address, _owner common.Address, _masterSigner common.Address, _aggregator common.Address) (*types.Transaction, error) {
	return _AAccount.contract.Transact(opts, "initialize", _entryPoint, _owner, _masterSigner, _aggregator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _entryPoint, address _owner, address _masterSigner, address _aggregator) returns()
func (_AAccount *AAccountSession) Initialize(_entryPoint common.Address, _owner common.Address, _masterSigner common.Address, _aggregator common.Address) (*types.Transaction, error) {
	return _AAccount.Contract.Initialize(&_AAccount.TransactOpts, _entryPoint, _owner, _masterSigner, _aggregator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _entryPoint, address _owner, address _masterSigner, address _aggregator) returns()
func (_AAccount *AAccountTransactorSession) Initialize(_entryPoint common.Address, _owner common.Address, _masterSigner common.Address, _aggregator common.Address) (*types.Transaction, error) {
	return _AAccount.Contract.Initialize(&_AAccount.TransactOpts, _entryPoint, _owner, _masterSigner, _aggregator)
}

// RevokeSessionKey is a paid mutator transaction binding the contract method 0x84f4fc6a.
//
// Solidity: function revokeSessionKey(address sessionKey) returns()
func (_AAccount *AAccountTransactor) RevokeSessionKey(opts *bind.TransactOpts, sessionKey common.Address) (*types.Transaction, error) {
	return _AAccount.contract.Transact(opts, "revokeSessionKey", sessionKey)
}

// RevokeSessionKey is a paid mutator transaction binding the contract method 0x84f4fc6a.
//
// Solidity: function revokeSessionKey(address sessionKey) returns()
func (_AAccount *AAccountSession) RevokeSessionKey(sessionKey common.Address) (*types.Transaction, error) {
	return _AAccount.Contract.RevokeSessionKey(&_AAccount.TransactOpts, sessionKey)
}

// RevokeSessionKey is a paid mutator transaction binding the contract method 0x84f4fc6a.
//
// Solidity: function revokeSessionKey(address sessionKey) returns()
func (_AAccount *AAccountTransactorSession) RevokeSessionKey(sessionKey common.Address) (*types.Transaction, error) {
	return _AAccount.Contract.RevokeSessionKey(&_AAccount.TransactOpts, sessionKey)
}

// SetAggregator is a paid mutator transaction binding the contract method 0xf9120af6.
//
// Solidity: function setAggregator(address _aggregator) returns()
func (_AAccount *AAccountTransactor) SetAggregator(opts *bind.TransactOpts, _aggregator common.Address) (*types.Transaction, error) {
	return _AAccount.contract.Transact(opts, "setAggregator", _aggregator)
}

// SetAggregator is a paid mutator transaction binding the contract method 0xf9120af6.
//
// Solidity: function setAggregator(address _aggregator) returns()
func (_AAccount *AAccountSession) SetAggregator(_aggregator common.Address) (*types.Transaction, error) {
	return _AAccount.Contract.SetAggregator(&_AAccount.TransactOpts, _aggregator)
}

// SetAggregator is a paid mutator transaction binding the contract method 0xf9120af6.
//
// Solidity: function setAggregator(address _aggregator) returns()
func (_AAccount *AAccountTransactorSession) SetAggregator(_aggregator common.Address) (*types.Transaction, error) {
	return _AAccount.Contract.SetAggregator(&_AAccount.TransactOpts, _aggregator)
}

// SetAuthorizedSessionKeyManager is a paid mutator transaction binding the contract method 0xe824561a.
//
// Solidity: function setAuthorizedSessionKeyManager(address _sessionKeyManager) returns()
func (_AAccount *AAccountTransactor) SetAuthorizedSessionKeyManager(opts *bind.TransactOpts, _sessionKeyManager common.Address) (*types.Transaction, error) {
	return _AAccount.contract.Transact(opts, "setAuthorizedSessionKeyManager", _sessionKeyManager)
}

// SetAuthorizedSessionKeyManager is a paid mutator transaction binding the contract method 0xe824561a.
//
// Solidity: function setAuthorizedSessionKeyManager(address _sessionKeyManager) returns()
func (_AAccount *AAccountSession) SetAuthorizedSessionKeyManager(_sessionKeyManager common.Address) (*types.Transaction, error) {
	return _AAccount.Contract.SetAuthorizedSessionKeyManager(&_AAccount.TransactOpts, _sessionKeyManager)
}

// SetAuthorizedSessionKeyManager is a paid mutator transaction binding the contract method 0xe824561a.
//
// Solidity: function setAuthorizedSessionKeyManager(address _sessionKeyManager) returns()
func (_AAccount *AAccountTransactorSession) SetAuthorizedSessionKeyManager(_sessionKeyManager common.Address) (*types.Transaction, error) {
	return _AAccount.Contract.SetAuthorizedSessionKeyManager(&_AAccount.TransactOpts, _sessionKeyManager)
}

// SetMasterSigner is a paid mutator transaction binding the contract method 0x7c004623.
//
// Solidity: function setMasterSigner(address newMasterSigner) returns()
func (_AAccount *AAccountTransactor) SetMasterSigner(opts *bind.TransactOpts, newMasterSigner common.Address) (*types.Transaction, error) {
	return _AAccount.contract.Transact(opts, "setMasterSigner", newMasterSigner)
}

// SetMasterSigner is a paid mutator transaction binding the contract method 0x7c004623.
//
// Solidity: function setMasterSigner(address newMasterSigner) returns()
func (_AAccount *AAccountSession) SetMasterSigner(newMasterSigner common.Address) (*types.Transaction, error) {
	return _AAccount.Contract.SetMasterSigner(&_AAccount.TransactOpts, newMasterSigner)
}

// SetMasterSigner is a paid mutator transaction binding the contract method 0x7c004623.
//
// Solidity: function setMasterSigner(address newMasterSigner) returns()
func (_AAccount *AAccountTransactorSession) SetMasterSigner(newMasterSigner common.Address) (*types.Transaction, error) {
	return _AAccount.Contract.SetMasterSigner(&_AAccount.TransactOpts, newMasterSigner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AAccount *AAccountTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AAccount.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AAccount *AAccountSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AAccount.Contract.UpgradeToAndCall(&_AAccount.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AAccount *AAccountTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AAccount.Contract.UpgradeToAndCall(&_AAccount.TransactOpts, newImplementation, data)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_AAccount *AAccountTransactor) ValidateUserOp(opts *bind.TransactOpts, userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _AAccount.contract.Transact(opts, "validateUserOp", userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_AAccount *AAccountSession) ValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _AAccount.Contract.ValidateUserOp(&_AAccount.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_AAccount *AAccountTransactorSession) ValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _AAccount.Contract.ValidateUserOp(&_AAccount.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_AAccount *AAccountTransactor) WithdrawDepositTo(opts *bind.TransactOpts, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AAccount.contract.Transact(opts, "withdrawDepositTo", withdrawAddress, amount)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_AAccount *AAccountSession) WithdrawDepositTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AAccount.Contract.WithdrawDepositTo(&_AAccount.TransactOpts, withdrawAddress, amount)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_AAccount *AAccountTransactorSession) WithdrawDepositTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AAccount.Contract.WithdrawDepositTo(&_AAccount.TransactOpts, withdrawAddress, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AAccount *AAccountTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AAccount.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AAccount *AAccountSession) Receive() (*types.Transaction, error) {
	return _AAccount.Contract.Receive(&_AAccount.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AAccount *AAccountTransactorSession) Receive() (*types.Transaction, error) {
	return _AAccount.Contract.Receive(&_AAccount.TransactOpts)
}

// AAccountAAccountInitializedIterator is returned from FilterAAccountInitialized and is used to iterate over the raw logs and unpacked data for AAccountInitialized events raised by the AAccount contract.
type AAccountAAccountInitializedIterator struct {
	Event *AAccountAAccountInitialized // Event containing the contract specifics and raw log

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
func (it *AAccountAAccountInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAccountAAccountInitialized)
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
		it.Event = new(AAccountAAccountInitialized)
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
func (it *AAccountAAccountInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAccountAAccountInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAccountAAccountInitialized represents a AAccountInitialized event raised by the AAccount contract.
type AAccountAAccountInitialized struct {
	Owner        common.Address
	MasterSigner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAAccountInitialized is a free log retrieval operation binding the contract event 0x377d2e23c6fc6c40a79ccf8789bddd229d795f19a448c30414a3c2396a9334a3.
//
// Solidity: event AAccountInitialized(address indexed owner, address indexed masterSigner)
func (_AAccount *AAccountFilterer) FilterAAccountInitialized(opts *bind.FilterOpts, owner []common.Address, masterSigner []common.Address) (*AAccountAAccountInitializedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var masterSignerRule []interface{}
	for _, masterSignerItem := range masterSigner {
		masterSignerRule = append(masterSignerRule, masterSignerItem)
	}

	logs, sub, err := _AAccount.contract.FilterLogs(opts, "AAccountInitialized", ownerRule, masterSignerRule)
	if err != nil {
		return nil, err
	}
	return &AAccountAAccountInitializedIterator{contract: _AAccount.contract, event: "AAccountInitialized", logs: logs, sub: sub}, nil
}

// WatchAAccountInitialized is a free log subscription operation binding the contract event 0x377d2e23c6fc6c40a79ccf8789bddd229d795f19a448c30414a3c2396a9334a3.
//
// Solidity: event AAccountInitialized(address indexed owner, address indexed masterSigner)
func (_AAccount *AAccountFilterer) WatchAAccountInitialized(opts *bind.WatchOpts, sink chan<- *AAccountAAccountInitialized, owner []common.Address, masterSigner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var masterSignerRule []interface{}
	for _, masterSignerItem := range masterSigner {
		masterSignerRule = append(masterSignerRule, masterSignerItem)
	}

	logs, sub, err := _AAccount.contract.WatchLogs(opts, "AAccountInitialized", ownerRule, masterSignerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAccountAAccountInitialized)
				if err := _AAccount.contract.UnpackLog(event, "AAccountInitialized", log); err != nil {
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

// ParseAAccountInitialized is a log parse operation binding the contract event 0x377d2e23c6fc6c40a79ccf8789bddd229d795f19a448c30414a3c2396a9334a3.
//
// Solidity: event AAccountInitialized(address indexed owner, address indexed masterSigner)
func (_AAccount *AAccountFilterer) ParseAAccountInitialized(log types.Log) (*AAccountAAccountInitialized, error) {
	event := new(AAccountAAccountInitialized)
	if err := _AAccount.contract.UnpackLog(event, "AAccountInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AAccountAggregatorUpdatedIterator is returned from FilterAggregatorUpdated and is used to iterate over the raw logs and unpacked data for AggregatorUpdated events raised by the AAccount contract.
type AAccountAggregatorUpdatedIterator struct {
	Event *AAccountAggregatorUpdated // Event containing the contract specifics and raw log

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
func (it *AAccountAggregatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAccountAggregatorUpdated)
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
		it.Event = new(AAccountAggregatorUpdated)
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
func (it *AAccountAggregatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAccountAggregatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAccountAggregatorUpdated represents a AggregatorUpdated event raised by the AAccount contract.
type AAccountAggregatorUpdated struct {
	OldAggregator common.Address
	NewAggregator common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAggregatorUpdated is a free log retrieval operation binding the contract event 0x89baabef7dfd0683c0ac16fd2a8431c51b49fbe654c3f7b5ef19763e2ccd88f2.
//
// Solidity: event AggregatorUpdated(address indexed oldAggregator, address indexed newAggregator)
func (_AAccount *AAccountFilterer) FilterAggregatorUpdated(opts *bind.FilterOpts, oldAggregator []common.Address, newAggregator []common.Address) (*AAccountAggregatorUpdatedIterator, error) {

	var oldAggregatorRule []interface{}
	for _, oldAggregatorItem := range oldAggregator {
		oldAggregatorRule = append(oldAggregatorRule, oldAggregatorItem)
	}
	var newAggregatorRule []interface{}
	for _, newAggregatorItem := range newAggregator {
		newAggregatorRule = append(newAggregatorRule, newAggregatorItem)
	}

	logs, sub, err := _AAccount.contract.FilterLogs(opts, "AggregatorUpdated", oldAggregatorRule, newAggregatorRule)
	if err != nil {
		return nil, err
	}
	return &AAccountAggregatorUpdatedIterator{contract: _AAccount.contract, event: "AggregatorUpdated", logs: logs, sub: sub}, nil
}

// WatchAggregatorUpdated is a free log subscription operation binding the contract event 0x89baabef7dfd0683c0ac16fd2a8431c51b49fbe654c3f7b5ef19763e2ccd88f2.
//
// Solidity: event AggregatorUpdated(address indexed oldAggregator, address indexed newAggregator)
func (_AAccount *AAccountFilterer) WatchAggregatorUpdated(opts *bind.WatchOpts, sink chan<- *AAccountAggregatorUpdated, oldAggregator []common.Address, newAggregator []common.Address) (event.Subscription, error) {

	var oldAggregatorRule []interface{}
	for _, oldAggregatorItem := range oldAggregator {
		oldAggregatorRule = append(oldAggregatorRule, oldAggregatorItem)
	}
	var newAggregatorRule []interface{}
	for _, newAggregatorItem := range newAggregator {
		newAggregatorRule = append(newAggregatorRule, newAggregatorItem)
	}

	logs, sub, err := _AAccount.contract.WatchLogs(opts, "AggregatorUpdated", oldAggregatorRule, newAggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAccountAggregatorUpdated)
				if err := _AAccount.contract.UnpackLog(event, "AggregatorUpdated", log); err != nil {
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
func (_AAccount *AAccountFilterer) ParseAggregatorUpdated(log types.Log) (*AAccountAggregatorUpdated, error) {
	event := new(AAccountAggregatorUpdated)
	if err := _AAccount.contract.UnpackLog(event, "AggregatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AAccountInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AAccount contract.
type AAccountInitializedIterator struct {
	Event *AAccountInitialized // Event containing the contract specifics and raw log

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
func (it *AAccountInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAccountInitialized)
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
		it.Event = new(AAccountInitialized)
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
func (it *AAccountInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAccountInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAccountInitialized represents a Initialized event raised by the AAccount contract.
type AAccountInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AAccount *AAccountFilterer) FilterInitialized(opts *bind.FilterOpts) (*AAccountInitializedIterator, error) {

	logs, sub, err := _AAccount.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AAccountInitializedIterator{contract: _AAccount.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AAccount *AAccountFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AAccountInitialized) (event.Subscription, error) {

	logs, sub, err := _AAccount.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAccountInitialized)
				if err := _AAccount.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_AAccount *AAccountFilterer) ParseInitialized(log types.Log) (*AAccountInitialized, error) {
	event := new(AAccountInitialized)
	if err := _AAccount.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AAccountMasterSignerUpdatedIterator is returned from FilterMasterSignerUpdated and is used to iterate over the raw logs and unpacked data for MasterSignerUpdated events raised by the AAccount contract.
type AAccountMasterSignerUpdatedIterator struct {
	Event *AAccountMasterSignerUpdated // Event containing the contract specifics and raw log

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
func (it *AAccountMasterSignerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAccountMasterSignerUpdated)
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
		it.Event = new(AAccountMasterSignerUpdated)
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
func (it *AAccountMasterSignerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAccountMasterSignerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAccountMasterSignerUpdated represents a MasterSignerUpdated event raised by the AAccount contract.
type AAccountMasterSignerUpdated struct {
	OldMasterSigner common.Address
	NewMasterSigner common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMasterSignerUpdated is a free log retrieval operation binding the contract event 0x93cf965b971c37bb042b3c24d2278260742c6f2021d1c3be06984b6d03ab60ab.
//
// Solidity: event MasterSignerUpdated(address indexed oldMasterSigner, address indexed newMasterSigner)
func (_AAccount *AAccountFilterer) FilterMasterSignerUpdated(opts *bind.FilterOpts, oldMasterSigner []common.Address, newMasterSigner []common.Address) (*AAccountMasterSignerUpdatedIterator, error) {

	var oldMasterSignerRule []interface{}
	for _, oldMasterSignerItem := range oldMasterSigner {
		oldMasterSignerRule = append(oldMasterSignerRule, oldMasterSignerItem)
	}
	var newMasterSignerRule []interface{}
	for _, newMasterSignerItem := range newMasterSigner {
		newMasterSignerRule = append(newMasterSignerRule, newMasterSignerItem)
	}

	logs, sub, err := _AAccount.contract.FilterLogs(opts, "MasterSignerUpdated", oldMasterSignerRule, newMasterSignerRule)
	if err != nil {
		return nil, err
	}
	return &AAccountMasterSignerUpdatedIterator{contract: _AAccount.contract, event: "MasterSignerUpdated", logs: logs, sub: sub}, nil
}

// WatchMasterSignerUpdated is a free log subscription operation binding the contract event 0x93cf965b971c37bb042b3c24d2278260742c6f2021d1c3be06984b6d03ab60ab.
//
// Solidity: event MasterSignerUpdated(address indexed oldMasterSigner, address indexed newMasterSigner)
func (_AAccount *AAccountFilterer) WatchMasterSignerUpdated(opts *bind.WatchOpts, sink chan<- *AAccountMasterSignerUpdated, oldMasterSigner []common.Address, newMasterSigner []common.Address) (event.Subscription, error) {

	var oldMasterSignerRule []interface{}
	for _, oldMasterSignerItem := range oldMasterSigner {
		oldMasterSignerRule = append(oldMasterSignerRule, oldMasterSignerItem)
	}
	var newMasterSignerRule []interface{}
	for _, newMasterSignerItem := range newMasterSigner {
		newMasterSignerRule = append(newMasterSignerRule, newMasterSignerItem)
	}

	logs, sub, err := _AAccount.contract.WatchLogs(opts, "MasterSignerUpdated", oldMasterSignerRule, newMasterSignerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAccountMasterSignerUpdated)
				if err := _AAccount.contract.UnpackLog(event, "MasterSignerUpdated", log); err != nil {
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
func (_AAccount *AAccountFilterer) ParseMasterSignerUpdated(log types.Log) (*AAccountMasterSignerUpdated, error) {
	event := new(AAccountMasterSignerUpdated)
	if err := _AAccount.contract.UnpackLog(event, "MasterSignerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AAccountMasterSignerValidationIterator is returned from FilterMasterSignerValidation and is used to iterate over the raw logs and unpacked data for MasterSignerValidation events raised by the AAccount contract.
type AAccountMasterSignerValidationIterator struct {
	Event *AAccountMasterSignerValidation // Event containing the contract specifics and raw log

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
func (it *AAccountMasterSignerValidationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAccountMasterSignerValidation)
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
		it.Event = new(AAccountMasterSignerValidation)
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
func (it *AAccountMasterSignerValidationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAccountMasterSignerValidationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAccountMasterSignerValidation represents a MasterSignerValidation event raised by the AAccount contract.
type AAccountMasterSignerValidation struct {
	MasterSigner common.Address
	UserOpHash   [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMasterSignerValidation is a free log retrieval operation binding the contract event 0xb1ff5f74802a94e0e6c412076207eb7b19aad8b251a5304145e0bf8f15ca1ef3.
//
// Solidity: event MasterSignerValidation(address indexed masterSigner, bytes32 userOpHash)
func (_AAccount *AAccountFilterer) FilterMasterSignerValidation(opts *bind.FilterOpts, masterSigner []common.Address) (*AAccountMasterSignerValidationIterator, error) {

	var masterSignerRule []interface{}
	for _, masterSignerItem := range masterSigner {
		masterSignerRule = append(masterSignerRule, masterSignerItem)
	}

	logs, sub, err := _AAccount.contract.FilterLogs(opts, "MasterSignerValidation", masterSignerRule)
	if err != nil {
		return nil, err
	}
	return &AAccountMasterSignerValidationIterator{contract: _AAccount.contract, event: "MasterSignerValidation", logs: logs, sub: sub}, nil
}

// WatchMasterSignerValidation is a free log subscription operation binding the contract event 0xb1ff5f74802a94e0e6c412076207eb7b19aad8b251a5304145e0bf8f15ca1ef3.
//
// Solidity: event MasterSignerValidation(address indexed masterSigner, bytes32 userOpHash)
func (_AAccount *AAccountFilterer) WatchMasterSignerValidation(opts *bind.WatchOpts, sink chan<- *AAccountMasterSignerValidation, masterSigner []common.Address) (event.Subscription, error) {

	var masterSignerRule []interface{}
	for _, masterSignerItem := range masterSigner {
		masterSignerRule = append(masterSignerRule, masterSignerItem)
	}

	logs, sub, err := _AAccount.contract.WatchLogs(opts, "MasterSignerValidation", masterSignerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAccountMasterSignerValidation)
				if err := _AAccount.contract.UnpackLog(event, "MasterSignerValidation", log); err != nil {
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
func (_AAccount *AAccountFilterer) ParseMasterSignerValidation(log types.Log) (*AAccountMasterSignerValidation, error) {
	event := new(AAccountMasterSignerValidation)
	if err := _AAccount.contract.UnpackLog(event, "MasterSignerValidation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AAccountSessionKeyAddedIterator is returned from FilterSessionKeyAdded and is used to iterate over the raw logs and unpacked data for SessionKeyAdded events raised by the AAccount contract.
type AAccountSessionKeyAddedIterator struct {
	Event *AAccountSessionKeyAdded // Event containing the contract specifics and raw log

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
func (it *AAccountSessionKeyAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAccountSessionKeyAdded)
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
		it.Event = new(AAccountSessionKeyAdded)
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
func (it *AAccountSessionKeyAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAccountSessionKeyAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAccountSessionKeyAdded represents a SessionKeyAdded event raised by the AAccount contract.
type AAccountSessionKeyAdded struct {
	SessionKey  common.Address
	ValidAfter  *big.Int
	ValidUntil  *big.Int
	Permissions [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSessionKeyAdded is a free log retrieval operation binding the contract event 0xa00bca54e0f01b47f85e9dd47c72f5921c1e7a1541b84f3888acf901a424e94b.
//
// Solidity: event SessionKeyAdded(address indexed sessionKey, uint48 validAfter, uint48 validUntil, bytes32 permissions)
func (_AAccount *AAccountFilterer) FilterSessionKeyAdded(opts *bind.FilterOpts, sessionKey []common.Address) (*AAccountSessionKeyAddedIterator, error) {

	var sessionKeyRule []interface{}
	for _, sessionKeyItem := range sessionKey {
		sessionKeyRule = append(sessionKeyRule, sessionKeyItem)
	}

	logs, sub, err := _AAccount.contract.FilterLogs(opts, "SessionKeyAdded", sessionKeyRule)
	if err != nil {
		return nil, err
	}
	return &AAccountSessionKeyAddedIterator{contract: _AAccount.contract, event: "SessionKeyAdded", logs: logs, sub: sub}, nil
}

// WatchSessionKeyAdded is a free log subscription operation binding the contract event 0xa00bca54e0f01b47f85e9dd47c72f5921c1e7a1541b84f3888acf901a424e94b.
//
// Solidity: event SessionKeyAdded(address indexed sessionKey, uint48 validAfter, uint48 validUntil, bytes32 permissions)
func (_AAccount *AAccountFilterer) WatchSessionKeyAdded(opts *bind.WatchOpts, sink chan<- *AAccountSessionKeyAdded, sessionKey []common.Address) (event.Subscription, error) {

	var sessionKeyRule []interface{}
	for _, sessionKeyItem := range sessionKey {
		sessionKeyRule = append(sessionKeyRule, sessionKeyItem)
	}

	logs, sub, err := _AAccount.contract.WatchLogs(opts, "SessionKeyAdded", sessionKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAccountSessionKeyAdded)
				if err := _AAccount.contract.UnpackLog(event, "SessionKeyAdded", log); err != nil {
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
func (_AAccount *AAccountFilterer) ParseSessionKeyAdded(log types.Log) (*AAccountSessionKeyAdded, error) {
	event := new(AAccountSessionKeyAdded)
	if err := _AAccount.contract.UnpackLog(event, "SessionKeyAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AAccountSessionKeyRevokedIterator is returned from FilterSessionKeyRevoked and is used to iterate over the raw logs and unpacked data for SessionKeyRevoked events raised by the AAccount contract.
type AAccountSessionKeyRevokedIterator struct {
	Event *AAccountSessionKeyRevoked // Event containing the contract specifics and raw log

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
func (it *AAccountSessionKeyRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAccountSessionKeyRevoked)
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
		it.Event = new(AAccountSessionKeyRevoked)
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
func (it *AAccountSessionKeyRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAccountSessionKeyRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAccountSessionKeyRevoked represents a SessionKeyRevoked event raised by the AAccount contract.
type AAccountSessionKeyRevoked struct {
	SessionKey common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSessionKeyRevoked is a free log retrieval operation binding the contract event 0x17c796fb82086b3c9effaec517342e5ca9ed8fd78c339137ec082f748ab60cbe.
//
// Solidity: event SessionKeyRevoked(address indexed sessionKey)
func (_AAccount *AAccountFilterer) FilterSessionKeyRevoked(opts *bind.FilterOpts, sessionKey []common.Address) (*AAccountSessionKeyRevokedIterator, error) {

	var sessionKeyRule []interface{}
	for _, sessionKeyItem := range sessionKey {
		sessionKeyRule = append(sessionKeyRule, sessionKeyItem)
	}

	logs, sub, err := _AAccount.contract.FilterLogs(opts, "SessionKeyRevoked", sessionKeyRule)
	if err != nil {
		return nil, err
	}
	return &AAccountSessionKeyRevokedIterator{contract: _AAccount.contract, event: "SessionKeyRevoked", logs: logs, sub: sub}, nil
}

// WatchSessionKeyRevoked is a free log subscription operation binding the contract event 0x17c796fb82086b3c9effaec517342e5ca9ed8fd78c339137ec082f748ab60cbe.
//
// Solidity: event SessionKeyRevoked(address indexed sessionKey)
func (_AAccount *AAccountFilterer) WatchSessionKeyRevoked(opts *bind.WatchOpts, sink chan<- *AAccountSessionKeyRevoked, sessionKey []common.Address) (event.Subscription, error) {

	var sessionKeyRule []interface{}
	for _, sessionKeyItem := range sessionKey {
		sessionKeyRule = append(sessionKeyRule, sessionKeyItem)
	}

	logs, sub, err := _AAccount.contract.WatchLogs(opts, "SessionKeyRevoked", sessionKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAccountSessionKeyRevoked)
				if err := _AAccount.contract.UnpackLog(event, "SessionKeyRevoked", log); err != nil {
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
func (_AAccount *AAccountFilterer) ParseSessionKeyRevoked(log types.Log) (*AAccountSessionKeyRevoked, error) {
	event := new(AAccountSessionKeyRevoked)
	if err := _AAccount.contract.UnpackLog(event, "SessionKeyRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AAccountSessionKeyValidationIterator is returned from FilterSessionKeyValidation and is used to iterate over the raw logs and unpacked data for SessionKeyValidation events raised by the AAccount contract.
type AAccountSessionKeyValidationIterator struct {
	Event *AAccountSessionKeyValidation // Event containing the contract specifics and raw log

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
func (it *AAccountSessionKeyValidationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAccountSessionKeyValidation)
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
		it.Event = new(AAccountSessionKeyValidation)
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
func (it *AAccountSessionKeyValidationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAccountSessionKeyValidationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAccountSessionKeyValidation represents a SessionKeyValidation event raised by the AAccount contract.
type AAccountSessionKeyValidation struct {
	SessionKey common.Address
	UserOpHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSessionKeyValidation is a free log retrieval operation binding the contract event 0x0c8028b60a408027126de0b307f68042b0a366d216ac9ad715fed951c9ad1d09.
//
// Solidity: event SessionKeyValidation(address indexed sessionKey, bytes32 userOpHash)
func (_AAccount *AAccountFilterer) FilterSessionKeyValidation(opts *bind.FilterOpts, sessionKey []common.Address) (*AAccountSessionKeyValidationIterator, error) {

	var sessionKeyRule []interface{}
	for _, sessionKeyItem := range sessionKey {
		sessionKeyRule = append(sessionKeyRule, sessionKeyItem)
	}

	logs, sub, err := _AAccount.contract.FilterLogs(opts, "SessionKeyValidation", sessionKeyRule)
	if err != nil {
		return nil, err
	}
	return &AAccountSessionKeyValidationIterator{contract: _AAccount.contract, event: "SessionKeyValidation", logs: logs, sub: sub}, nil
}

// WatchSessionKeyValidation is a free log subscription operation binding the contract event 0x0c8028b60a408027126de0b307f68042b0a366d216ac9ad715fed951c9ad1d09.
//
// Solidity: event SessionKeyValidation(address indexed sessionKey, bytes32 userOpHash)
func (_AAccount *AAccountFilterer) WatchSessionKeyValidation(opts *bind.WatchOpts, sink chan<- *AAccountSessionKeyValidation, sessionKey []common.Address) (event.Subscription, error) {

	var sessionKeyRule []interface{}
	for _, sessionKeyItem := range sessionKey {
		sessionKeyRule = append(sessionKeyRule, sessionKeyItem)
	}

	logs, sub, err := _AAccount.contract.WatchLogs(opts, "SessionKeyValidation", sessionKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAccountSessionKeyValidation)
				if err := _AAccount.contract.UnpackLog(event, "SessionKeyValidation", log); err != nil {
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
func (_AAccount *AAccountFilterer) ParseSessionKeyValidation(log types.Log) (*AAccountSessionKeyValidation, error) {
	event := new(AAccountSessionKeyValidation)
	if err := _AAccount.contract.UnpackLog(event, "SessionKeyValidation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AAccountUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the AAccount contract.
type AAccountUpgradedIterator struct {
	Event *AAccountUpgraded // Event containing the contract specifics and raw log

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
func (it *AAccountUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAccountUpgraded)
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
		it.Event = new(AAccountUpgraded)
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
func (it *AAccountUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAccountUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAccountUpgraded represents a Upgraded event raised by the AAccount contract.
type AAccountUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AAccount *AAccountFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*AAccountUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AAccount.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &AAccountUpgradedIterator{contract: _AAccount.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AAccount *AAccountFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *AAccountUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AAccount.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAccountUpgraded)
				if err := _AAccount.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_AAccount *AAccountFilterer) ParseUpgraded(log types.Log) (*AAccountUpgraded, error) {
	event := new(AAccountUpgraded)
	if err := _AAccount.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
