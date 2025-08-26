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

// MasterAggregatorMetaData contains all meta data concerning the MasterAggregator contract.
var MasterAggregatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"master\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"Account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"authorized\",\"type\":\"bool\"}],\"name\":\"AccountAuthorized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"master\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"operationCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"aggregatedHash\",\"type\":\"bytes32\"}],\"name\":\"AggregatedValidation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"master\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"authorized\",\"type\":\"bool\"}],\"name\":\"MasterAuthorized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"entryPoint\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"delay\",\"type\":\"uint32\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation[]\",\"name\":\"userOps\",\"type\":\"tuple[]\"}],\"name\":\"aggregateSignatures\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"aggregatedSignature\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorizedMasters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"authorizedMastersList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"Account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"master\",\"type\":\"address\"}],\"name\":\"autoAuthorizeAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"master\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"Accounts\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"authorized\",\"type\":\"bool\"}],\"name\":\"batchSetAccountAuthorization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"operationCount\",\"type\":\"uint256\"}],\"name\":\"calculateGasSavings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"savings\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation[]\",\"name\":\"userOps\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"masterSigner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"masterSignature\",\"type\":\"bytes\"}],\"name\":\"createMasterAggregatedSignature\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"aggregatedSignature\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"master\",\"type\":\"address\"}],\"name\":\"getMasterNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation[]\",\"name\":\"userOps\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"masterSigner\",\"type\":\"address\"}],\"name\":\"getMasterSigningData\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashToSign\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_initialMasters\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"Account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"master\",\"type\":\"address\"}],\"name\":\"isAccountControlledByMaster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"masterNonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"masterToAccounts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxAggregatedOps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"master\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"Account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"authorized\",\"type\":\"bool\"}],\"name\":\"setAccountAuthorization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"master\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"authorized\",\"type\":\"bool\"}],\"name\":\"setMasterAuthorization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"entryPoint\",\"type\":\"address\"}],\"name\":\"unlockStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxAggregatedOps\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_validationWindow\",\"type\":\"uint256\"}],\"name\":\"updateConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation[]\",\"name\":\"userOps\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"validateSignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"validateUserOpSignature\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validationWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"entryPoint\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MasterAggregatorABI is the input ABI used to generate the binding from.
// Deprecated: Use MasterAggregatorMetaData.ABI instead.
var MasterAggregatorABI = MasterAggregatorMetaData.ABI

// MasterAggregator is an auto generated Go binding around an Ethereum contract.
type MasterAggregator struct {
	MasterAggregatorCaller     // Read-only binding to the contract
	MasterAggregatorTransactor // Write-only binding to the contract
	MasterAggregatorFilterer   // Log filterer for contract events
}

// MasterAggregatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type MasterAggregatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterAggregatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MasterAggregatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterAggregatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MasterAggregatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterAggregatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MasterAggregatorSession struct {
	Contract     *MasterAggregator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MasterAggregatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MasterAggregatorCallerSession struct {
	Contract *MasterAggregatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MasterAggregatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MasterAggregatorTransactorSession struct {
	Contract     *MasterAggregatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MasterAggregatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type MasterAggregatorRaw struct {
	Contract *MasterAggregator // Generic contract binding to access the raw methods on
}

// MasterAggregatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MasterAggregatorCallerRaw struct {
	Contract *MasterAggregatorCaller // Generic read-only contract binding to access the raw methods on
}

// MasterAggregatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MasterAggregatorTransactorRaw struct {
	Contract *MasterAggregatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMasterAggregator creates a new instance of MasterAggregator, bound to a specific deployed contract.
func NewMasterAggregator(address common.Address, backend bind.ContractBackend) (*MasterAggregator, error) {
	contract, err := bindMasterAggregator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MasterAggregator{MasterAggregatorCaller: MasterAggregatorCaller{contract: contract}, MasterAggregatorTransactor: MasterAggregatorTransactor{contract: contract}, MasterAggregatorFilterer: MasterAggregatorFilterer{contract: contract}}, nil
}

// NewMasterAggregatorCaller creates a new read-only instance of MasterAggregator, bound to a specific deployed contract.
func NewMasterAggregatorCaller(address common.Address, caller bind.ContractCaller) (*MasterAggregatorCaller, error) {
	contract, err := bindMasterAggregator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MasterAggregatorCaller{contract: contract}, nil
}

// NewMasterAggregatorTransactor creates a new write-only instance of MasterAggregator, bound to a specific deployed contract.
func NewMasterAggregatorTransactor(address common.Address, transactor bind.ContractTransactor) (*MasterAggregatorTransactor, error) {
	contract, err := bindMasterAggregator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MasterAggregatorTransactor{contract: contract}, nil
}

// NewMasterAggregatorFilterer creates a new log filterer instance of MasterAggregator, bound to a specific deployed contract.
func NewMasterAggregatorFilterer(address common.Address, filterer bind.ContractFilterer) (*MasterAggregatorFilterer, error) {
	contract, err := bindMasterAggregator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MasterAggregatorFilterer{contract: contract}, nil
}

// bindMasterAggregator binds a generic wrapper to an already deployed contract.
func bindMasterAggregator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MasterAggregatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MasterAggregator *MasterAggregatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MasterAggregator.Contract.MasterAggregatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MasterAggregator *MasterAggregatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterAggregator.Contract.MasterAggregatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MasterAggregator *MasterAggregatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MasterAggregator.Contract.MasterAggregatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MasterAggregator *MasterAggregatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MasterAggregator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MasterAggregator *MasterAggregatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterAggregator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MasterAggregator *MasterAggregatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MasterAggregator.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_MasterAggregator *MasterAggregatorCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MasterAggregator.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_MasterAggregator *MasterAggregatorSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _MasterAggregator.Contract.UPGRADEINTERFACEVERSION(&_MasterAggregator.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_MasterAggregator *MasterAggregatorCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _MasterAggregator.Contract.UPGRADEINTERFACEVERSION(&_MasterAggregator.CallOpts)
}

// AggregateSignatures is a free data retrieval call binding the contract method 0xae574a43.
//
// Solidity: function aggregateSignatures((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] userOps) view returns(bytes aggregatedSignature)
func (_MasterAggregator *MasterAggregatorCaller) AggregateSignatures(opts *bind.CallOpts, userOps []PackedUserOperation) ([]byte, error) {
	var out []interface{}
	err := _MasterAggregator.contract.Call(opts, &out, "aggregateSignatures", userOps)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// AggregateSignatures is a free data retrieval call binding the contract method 0xae574a43.
//
// Solidity: function aggregateSignatures((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] userOps) view returns(bytes aggregatedSignature)
func (_MasterAggregator *MasterAggregatorSession) AggregateSignatures(userOps []PackedUserOperation) ([]byte, error) {
	return _MasterAggregator.Contract.AggregateSignatures(&_MasterAggregator.CallOpts, userOps)
}

// AggregateSignatures is a free data retrieval call binding the contract method 0xae574a43.
//
// Solidity: function aggregateSignatures((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] userOps) view returns(bytes aggregatedSignature)
func (_MasterAggregator *MasterAggregatorCallerSession) AggregateSignatures(userOps []PackedUserOperation) ([]byte, error) {
	return _MasterAggregator.Contract.AggregateSignatures(&_MasterAggregator.CallOpts, userOps)
}

// AuthorizedMasters is a free data retrieval call binding the contract method 0xfe0221fe.
//
// Solidity: function authorizedMasters(address ) view returns(bool)
func (_MasterAggregator *MasterAggregatorCaller) AuthorizedMasters(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MasterAggregator.contract.Call(opts, &out, "authorizedMasters", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AuthorizedMasters is a free data retrieval call binding the contract method 0xfe0221fe.
//
// Solidity: function authorizedMasters(address ) view returns(bool)
func (_MasterAggregator *MasterAggregatorSession) AuthorizedMasters(arg0 common.Address) (bool, error) {
	return _MasterAggregator.Contract.AuthorizedMasters(&_MasterAggregator.CallOpts, arg0)
}

// AuthorizedMasters is a free data retrieval call binding the contract method 0xfe0221fe.
//
// Solidity: function authorizedMasters(address ) view returns(bool)
func (_MasterAggregator *MasterAggregatorCallerSession) AuthorizedMasters(arg0 common.Address) (bool, error) {
	return _MasterAggregator.Contract.AuthorizedMasters(&_MasterAggregator.CallOpts, arg0)
}

// AuthorizedMastersList is a free data retrieval call binding the contract method 0x25df7a3e.
//
// Solidity: function authorizedMastersList(uint256 ) view returns(address)
func (_MasterAggregator *MasterAggregatorCaller) AuthorizedMastersList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MasterAggregator.contract.Call(opts, &out, "authorizedMastersList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AuthorizedMastersList is a free data retrieval call binding the contract method 0x25df7a3e.
//
// Solidity: function authorizedMastersList(uint256 ) view returns(address)
func (_MasterAggregator *MasterAggregatorSession) AuthorizedMastersList(arg0 *big.Int) (common.Address, error) {
	return _MasterAggregator.Contract.AuthorizedMastersList(&_MasterAggregator.CallOpts, arg0)
}

// AuthorizedMastersList is a free data retrieval call binding the contract method 0x25df7a3e.
//
// Solidity: function authorizedMastersList(uint256 ) view returns(address)
func (_MasterAggregator *MasterAggregatorCallerSession) AuthorizedMastersList(arg0 *big.Int) (common.Address, error) {
	return _MasterAggregator.Contract.AuthorizedMastersList(&_MasterAggregator.CallOpts, arg0)
}

// CalculateGasSavings is a free data retrieval call binding the contract method 0xa8b772e4.
//
// Solidity: function calculateGasSavings(uint256 operationCount) pure returns(uint256 savings)
func (_MasterAggregator *MasterAggregatorCaller) CalculateGasSavings(opts *bind.CallOpts, operationCount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MasterAggregator.contract.Call(opts, &out, "calculateGasSavings", operationCount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateGasSavings is a free data retrieval call binding the contract method 0xa8b772e4.
//
// Solidity: function calculateGasSavings(uint256 operationCount) pure returns(uint256 savings)
func (_MasterAggregator *MasterAggregatorSession) CalculateGasSavings(operationCount *big.Int) (*big.Int, error) {
	return _MasterAggregator.Contract.CalculateGasSavings(&_MasterAggregator.CallOpts, operationCount)
}

// CalculateGasSavings is a free data retrieval call binding the contract method 0xa8b772e4.
//
// Solidity: function calculateGasSavings(uint256 operationCount) pure returns(uint256 savings)
func (_MasterAggregator *MasterAggregatorCallerSession) CalculateGasSavings(operationCount *big.Int) (*big.Int, error) {
	return _MasterAggregator.Contract.CalculateGasSavings(&_MasterAggregator.CallOpts, operationCount)
}

// GetMasterNonce is a free data retrieval call binding the contract method 0x3671255b.
//
// Solidity: function getMasterNonce(address master) view returns(uint256 nonce)
func (_MasterAggregator *MasterAggregatorCaller) GetMasterNonce(opts *bind.CallOpts, master common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MasterAggregator.contract.Call(opts, &out, "getMasterNonce", master)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMasterNonce is a free data retrieval call binding the contract method 0x3671255b.
//
// Solidity: function getMasterNonce(address master) view returns(uint256 nonce)
func (_MasterAggregator *MasterAggregatorSession) GetMasterNonce(master common.Address) (*big.Int, error) {
	return _MasterAggregator.Contract.GetMasterNonce(&_MasterAggregator.CallOpts, master)
}

// GetMasterNonce is a free data retrieval call binding the contract method 0x3671255b.
//
// Solidity: function getMasterNonce(address master) view returns(uint256 nonce)
func (_MasterAggregator *MasterAggregatorCallerSession) GetMasterNonce(master common.Address) (*big.Int, error) {
	return _MasterAggregator.Contract.GetMasterNonce(&_MasterAggregator.CallOpts, master)
}

// GetMasterSigningData is a free data retrieval call binding the contract method 0xbe7bf224.
//
// Solidity: function getMasterSigningData((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] userOps, address masterSigner) view returns(bytes32 hashToSign, uint256 nonce)
func (_MasterAggregator *MasterAggregatorCaller) GetMasterSigningData(opts *bind.CallOpts, userOps []PackedUserOperation, masterSigner common.Address) (struct {
	HashToSign [32]byte
	Nonce      *big.Int
}, error) {
	var out []interface{}
	err := _MasterAggregator.contract.Call(opts, &out, "getMasterSigningData", userOps, masterSigner)

	outstruct := new(struct {
		HashToSign [32]byte
		Nonce      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.HashToSign = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Nonce = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetMasterSigningData is a free data retrieval call binding the contract method 0xbe7bf224.
//
// Solidity: function getMasterSigningData((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] userOps, address masterSigner) view returns(bytes32 hashToSign, uint256 nonce)
func (_MasterAggregator *MasterAggregatorSession) GetMasterSigningData(userOps []PackedUserOperation, masterSigner common.Address) (struct {
	HashToSign [32]byte
	Nonce      *big.Int
}, error) {
	return _MasterAggregator.Contract.GetMasterSigningData(&_MasterAggregator.CallOpts, userOps, masterSigner)
}

// GetMasterSigningData is a free data retrieval call binding the contract method 0xbe7bf224.
//
// Solidity: function getMasterSigningData((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] userOps, address masterSigner) view returns(bytes32 hashToSign, uint256 nonce)
func (_MasterAggregator *MasterAggregatorCallerSession) GetMasterSigningData(userOps []PackedUserOperation, masterSigner common.Address) (struct {
	HashToSign [32]byte
	Nonce      *big.Int
}, error) {
	return _MasterAggregator.Contract.GetMasterSigningData(&_MasterAggregator.CallOpts, userOps, masterSigner)
}

// IsAccountControlledByMaster is a free data retrieval call binding the contract method 0x3e90df87.
//
// Solidity: function isAccountControlledByMaster(address Account, address master) view returns(bool isValid)
func (_MasterAggregator *MasterAggregatorCaller) IsAccountControlledByMaster(opts *bind.CallOpts, Account common.Address, master common.Address) (bool, error) {
	var out []interface{}
	err := _MasterAggregator.contract.Call(opts, &out, "isAccountControlledByMaster", Account, master)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAccountControlledByMaster is a free data retrieval call binding the contract method 0x3e90df87.
//
// Solidity: function isAccountControlledByMaster(address Account, address master) view returns(bool isValid)
func (_MasterAggregator *MasterAggregatorSession) IsAccountControlledByMaster(Account common.Address, master common.Address) (bool, error) {
	return _MasterAggregator.Contract.IsAccountControlledByMaster(&_MasterAggregator.CallOpts, Account, master)
}

// IsAccountControlledByMaster is a free data retrieval call binding the contract method 0x3e90df87.
//
// Solidity: function isAccountControlledByMaster(address Account, address master) view returns(bool isValid)
func (_MasterAggregator *MasterAggregatorCallerSession) IsAccountControlledByMaster(Account common.Address, master common.Address) (bool, error) {
	return _MasterAggregator.Contract.IsAccountControlledByMaster(&_MasterAggregator.CallOpts, Account, master)
}

// MasterNonces is a free data retrieval call binding the contract method 0x45253c53.
//
// Solidity: function masterNonces(address ) view returns(uint256)
func (_MasterAggregator *MasterAggregatorCaller) MasterNonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MasterAggregator.contract.Call(opts, &out, "masterNonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MasterNonces is a free data retrieval call binding the contract method 0x45253c53.
//
// Solidity: function masterNonces(address ) view returns(uint256)
func (_MasterAggregator *MasterAggregatorSession) MasterNonces(arg0 common.Address) (*big.Int, error) {
	return _MasterAggregator.Contract.MasterNonces(&_MasterAggregator.CallOpts, arg0)
}

// MasterNonces is a free data retrieval call binding the contract method 0x45253c53.
//
// Solidity: function masterNonces(address ) view returns(uint256)
func (_MasterAggregator *MasterAggregatorCallerSession) MasterNonces(arg0 common.Address) (*big.Int, error) {
	return _MasterAggregator.Contract.MasterNonces(&_MasterAggregator.CallOpts, arg0)
}

// MasterToAccounts is a free data retrieval call binding the contract method 0x1914355c.
//
// Solidity: function masterToAccounts(address , address ) view returns(bool)
func (_MasterAggregator *MasterAggregatorCaller) MasterToAccounts(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _MasterAggregator.contract.Call(opts, &out, "masterToAccounts", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MasterToAccounts is a free data retrieval call binding the contract method 0x1914355c.
//
// Solidity: function masterToAccounts(address , address ) view returns(bool)
func (_MasterAggregator *MasterAggregatorSession) MasterToAccounts(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _MasterAggregator.Contract.MasterToAccounts(&_MasterAggregator.CallOpts, arg0, arg1)
}

// MasterToAccounts is a free data retrieval call binding the contract method 0x1914355c.
//
// Solidity: function masterToAccounts(address , address ) view returns(bool)
func (_MasterAggregator *MasterAggregatorCallerSession) MasterToAccounts(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _MasterAggregator.Contract.MasterToAccounts(&_MasterAggregator.CallOpts, arg0, arg1)
}

// MaxAggregatedOps is a free data retrieval call binding the contract method 0x0e166cfe.
//
// Solidity: function maxAggregatedOps() view returns(uint256)
func (_MasterAggregator *MasterAggregatorCaller) MaxAggregatedOps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterAggregator.contract.Call(opts, &out, "maxAggregatedOps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAggregatedOps is a free data retrieval call binding the contract method 0x0e166cfe.
//
// Solidity: function maxAggregatedOps() view returns(uint256)
func (_MasterAggregator *MasterAggregatorSession) MaxAggregatedOps() (*big.Int, error) {
	return _MasterAggregator.Contract.MaxAggregatedOps(&_MasterAggregator.CallOpts)
}

// MaxAggregatedOps is a free data retrieval call binding the contract method 0x0e166cfe.
//
// Solidity: function maxAggregatedOps() view returns(uint256)
func (_MasterAggregator *MasterAggregatorCallerSession) MaxAggregatedOps() (*big.Int, error) {
	return _MasterAggregator.Contract.MaxAggregatedOps(&_MasterAggregator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MasterAggregator *MasterAggregatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterAggregator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MasterAggregator *MasterAggregatorSession) Owner() (common.Address, error) {
	return _MasterAggregator.Contract.Owner(&_MasterAggregator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MasterAggregator *MasterAggregatorCallerSession) Owner() (common.Address, error) {
	return _MasterAggregator.Contract.Owner(&_MasterAggregator.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MasterAggregator *MasterAggregatorCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MasterAggregator.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MasterAggregator *MasterAggregatorSession) ProxiableUUID() ([32]byte, error) {
	return _MasterAggregator.Contract.ProxiableUUID(&_MasterAggregator.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MasterAggregator *MasterAggregatorCallerSession) ProxiableUUID() ([32]byte, error) {
	return _MasterAggregator.Contract.ProxiableUUID(&_MasterAggregator.CallOpts)
}

// ValidateUserOpSignature is a free data retrieval call binding the contract method 0x062a422b.
//
// Solidity: function validateUserOpSignature((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) ) pure returns(bytes)
func (_MasterAggregator *MasterAggregatorCaller) ValidateUserOpSignature(opts *bind.CallOpts, arg0 PackedUserOperation) ([]byte, error) {
	var out []interface{}
	err := _MasterAggregator.contract.Call(opts, &out, "validateUserOpSignature", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ValidateUserOpSignature is a free data retrieval call binding the contract method 0x062a422b.
//
// Solidity: function validateUserOpSignature((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) ) pure returns(bytes)
func (_MasterAggregator *MasterAggregatorSession) ValidateUserOpSignature(arg0 PackedUserOperation) ([]byte, error) {
	return _MasterAggregator.Contract.ValidateUserOpSignature(&_MasterAggregator.CallOpts, arg0)
}

// ValidateUserOpSignature is a free data retrieval call binding the contract method 0x062a422b.
//
// Solidity: function validateUserOpSignature((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) ) pure returns(bytes)
func (_MasterAggregator *MasterAggregatorCallerSession) ValidateUserOpSignature(arg0 PackedUserOperation) ([]byte, error) {
	return _MasterAggregator.Contract.ValidateUserOpSignature(&_MasterAggregator.CallOpts, arg0)
}

// ValidationWindow is a free data retrieval call binding the contract method 0x2d395e4d.
//
// Solidity: function validationWindow() view returns(uint256)
func (_MasterAggregator *MasterAggregatorCaller) ValidationWindow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterAggregator.contract.Call(opts, &out, "validationWindow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidationWindow is a free data retrieval call binding the contract method 0x2d395e4d.
//
// Solidity: function validationWindow() view returns(uint256)
func (_MasterAggregator *MasterAggregatorSession) ValidationWindow() (*big.Int, error) {
	return _MasterAggregator.Contract.ValidationWindow(&_MasterAggregator.CallOpts)
}

// ValidationWindow is a free data retrieval call binding the contract method 0x2d395e4d.
//
// Solidity: function validationWindow() view returns(uint256)
func (_MasterAggregator *MasterAggregatorCallerSession) ValidationWindow() (*big.Int, error) {
	return _MasterAggregator.Contract.ValidationWindow(&_MasterAggregator.CallOpts)
}

// AddStake is a paid mutator transaction binding the contract method 0x45171159.
//
// Solidity: function addStake(address entryPoint, uint32 delay) payable returns()
func (_MasterAggregator *MasterAggregatorTransactor) AddStake(opts *bind.TransactOpts, entryPoint common.Address, delay uint32) (*types.Transaction, error) {
	return _MasterAggregator.contract.Transact(opts, "addStake", entryPoint, delay)
}

// AddStake is a paid mutator transaction binding the contract method 0x45171159.
//
// Solidity: function addStake(address entryPoint, uint32 delay) payable returns()
func (_MasterAggregator *MasterAggregatorSession) AddStake(entryPoint common.Address, delay uint32) (*types.Transaction, error) {
	return _MasterAggregator.Contract.AddStake(&_MasterAggregator.TransactOpts, entryPoint, delay)
}

// AddStake is a paid mutator transaction binding the contract method 0x45171159.
//
// Solidity: function addStake(address entryPoint, uint32 delay) payable returns()
func (_MasterAggregator *MasterAggregatorTransactorSession) AddStake(entryPoint common.Address, delay uint32) (*types.Transaction, error) {
	return _MasterAggregator.Contract.AddStake(&_MasterAggregator.TransactOpts, entryPoint, delay)
}

// AutoAuthorizeAccount is a paid mutator transaction binding the contract method 0x34f10a69.
//
// Solidity: function autoAuthorizeAccount(address Account, address master) returns()
func (_MasterAggregator *MasterAggregatorTransactor) AutoAuthorizeAccount(opts *bind.TransactOpts, Account common.Address, master common.Address) (*types.Transaction, error) {
	return _MasterAggregator.contract.Transact(opts, "autoAuthorizeAccount", Account, master)
}

// AutoAuthorizeAccount is a paid mutator transaction binding the contract method 0x34f10a69.
//
// Solidity: function autoAuthorizeAccount(address Account, address master) returns()
func (_MasterAggregator *MasterAggregatorSession) AutoAuthorizeAccount(Account common.Address, master common.Address) (*types.Transaction, error) {
	return _MasterAggregator.Contract.AutoAuthorizeAccount(&_MasterAggregator.TransactOpts, Account, master)
}

// AutoAuthorizeAccount is a paid mutator transaction binding the contract method 0x34f10a69.
//
// Solidity: function autoAuthorizeAccount(address Account, address master) returns()
func (_MasterAggregator *MasterAggregatorTransactorSession) AutoAuthorizeAccount(Account common.Address, master common.Address) (*types.Transaction, error) {
	return _MasterAggregator.Contract.AutoAuthorizeAccount(&_MasterAggregator.TransactOpts, Account, master)
}

// BatchSetAccountAuthorization is a paid mutator transaction binding the contract method 0x1b6053f9.
//
// Solidity: function batchSetAccountAuthorization(address master, address[] Accounts, bool authorized) returns()
func (_MasterAggregator *MasterAggregatorTransactor) BatchSetAccountAuthorization(opts *bind.TransactOpts, master common.Address, Accounts []common.Address, authorized bool) (*types.Transaction, error) {
	return _MasterAggregator.contract.Transact(opts, "batchSetAccountAuthorization", master, Accounts, authorized)
}

// BatchSetAccountAuthorization is a paid mutator transaction binding the contract method 0x1b6053f9.
//
// Solidity: function batchSetAccountAuthorization(address master, address[] Accounts, bool authorized) returns()
func (_MasterAggregator *MasterAggregatorSession) BatchSetAccountAuthorization(master common.Address, Accounts []common.Address, authorized bool) (*types.Transaction, error) {
	return _MasterAggregator.Contract.BatchSetAccountAuthorization(&_MasterAggregator.TransactOpts, master, Accounts, authorized)
}

// BatchSetAccountAuthorization is a paid mutator transaction binding the contract method 0x1b6053f9.
//
// Solidity: function batchSetAccountAuthorization(address master, address[] Accounts, bool authorized) returns()
func (_MasterAggregator *MasterAggregatorTransactorSession) BatchSetAccountAuthorization(master common.Address, Accounts []common.Address, authorized bool) (*types.Transaction, error) {
	return _MasterAggregator.Contract.BatchSetAccountAuthorization(&_MasterAggregator.TransactOpts, master, Accounts, authorized)
}

// CreateMasterAggregatedSignature is a paid mutator transaction binding the contract method 0xcc8a43ea.
//
// Solidity: function createMasterAggregatedSignature((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] userOps, address masterSigner, bytes masterSignature) returns(bytes aggregatedSignature)
func (_MasterAggregator *MasterAggregatorTransactor) CreateMasterAggregatedSignature(opts *bind.TransactOpts, userOps []PackedUserOperation, masterSigner common.Address, masterSignature []byte) (*types.Transaction, error) {
	return _MasterAggregator.contract.Transact(opts, "createMasterAggregatedSignature", userOps, masterSigner, masterSignature)
}

// CreateMasterAggregatedSignature is a paid mutator transaction binding the contract method 0xcc8a43ea.
//
// Solidity: function createMasterAggregatedSignature((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] userOps, address masterSigner, bytes masterSignature) returns(bytes aggregatedSignature)
func (_MasterAggregator *MasterAggregatorSession) CreateMasterAggregatedSignature(userOps []PackedUserOperation, masterSigner common.Address, masterSignature []byte) (*types.Transaction, error) {
	return _MasterAggregator.Contract.CreateMasterAggregatedSignature(&_MasterAggregator.TransactOpts, userOps, masterSigner, masterSignature)
}

// CreateMasterAggregatedSignature is a paid mutator transaction binding the contract method 0xcc8a43ea.
//
// Solidity: function createMasterAggregatedSignature((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] userOps, address masterSigner, bytes masterSignature) returns(bytes aggregatedSignature)
func (_MasterAggregator *MasterAggregatorTransactorSession) CreateMasterAggregatedSignature(userOps []PackedUserOperation, masterSigner common.Address, masterSignature []byte) (*types.Transaction, error) {
	return _MasterAggregator.Contract.CreateMasterAggregatedSignature(&_MasterAggregator.TransactOpts, userOps, masterSigner, masterSignature)
}

// Initialize is a paid mutator transaction binding the contract method 0x946d9204.
//
// Solidity: function initialize(address _owner, address[] _initialMasters) returns()
func (_MasterAggregator *MasterAggregatorTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _initialMasters []common.Address) (*types.Transaction, error) {
	return _MasterAggregator.contract.Transact(opts, "initialize", _owner, _initialMasters)
}

// Initialize is a paid mutator transaction binding the contract method 0x946d9204.
//
// Solidity: function initialize(address _owner, address[] _initialMasters) returns()
func (_MasterAggregator *MasterAggregatorSession) Initialize(_owner common.Address, _initialMasters []common.Address) (*types.Transaction, error) {
	return _MasterAggregator.Contract.Initialize(&_MasterAggregator.TransactOpts, _owner, _initialMasters)
}

// Initialize is a paid mutator transaction binding the contract method 0x946d9204.
//
// Solidity: function initialize(address _owner, address[] _initialMasters) returns()
func (_MasterAggregator *MasterAggregatorTransactorSession) Initialize(_owner common.Address, _initialMasters []common.Address) (*types.Transaction, error) {
	return _MasterAggregator.Contract.Initialize(&_MasterAggregator.TransactOpts, _owner, _initialMasters)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MasterAggregator *MasterAggregatorTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterAggregator.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MasterAggregator *MasterAggregatorSession) Pause() (*types.Transaction, error) {
	return _MasterAggregator.Contract.Pause(&_MasterAggregator.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MasterAggregator *MasterAggregatorTransactorSession) Pause() (*types.Transaction, error) {
	return _MasterAggregator.Contract.Pause(&_MasterAggregator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MasterAggregator *MasterAggregatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterAggregator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MasterAggregator *MasterAggregatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MasterAggregator.Contract.RenounceOwnership(&_MasterAggregator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MasterAggregator *MasterAggregatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MasterAggregator.Contract.RenounceOwnership(&_MasterAggregator.TransactOpts)
}

// SetAccountAuthorization is a paid mutator transaction binding the contract method 0x14f368ce.
//
// Solidity: function setAccountAuthorization(address master, address Account, bool authorized) returns()
func (_MasterAggregator *MasterAggregatorTransactor) SetAccountAuthorization(opts *bind.TransactOpts, master common.Address, Account common.Address, authorized bool) (*types.Transaction, error) {
	return _MasterAggregator.contract.Transact(opts, "setAccountAuthorization", master, Account, authorized)
}

// SetAccountAuthorization is a paid mutator transaction binding the contract method 0x14f368ce.
//
// Solidity: function setAccountAuthorization(address master, address Account, bool authorized) returns()
func (_MasterAggregator *MasterAggregatorSession) SetAccountAuthorization(master common.Address, Account common.Address, authorized bool) (*types.Transaction, error) {
	return _MasterAggregator.Contract.SetAccountAuthorization(&_MasterAggregator.TransactOpts, master, Account, authorized)
}

// SetAccountAuthorization is a paid mutator transaction binding the contract method 0x14f368ce.
//
// Solidity: function setAccountAuthorization(address master, address Account, bool authorized) returns()
func (_MasterAggregator *MasterAggregatorTransactorSession) SetAccountAuthorization(master common.Address, Account common.Address, authorized bool) (*types.Transaction, error) {
	return _MasterAggregator.Contract.SetAccountAuthorization(&_MasterAggregator.TransactOpts, master, Account, authorized)
}

// SetMasterAuthorization is a paid mutator transaction binding the contract method 0x79d37814.
//
// Solidity: function setMasterAuthorization(address master, bool authorized) returns()
func (_MasterAggregator *MasterAggregatorTransactor) SetMasterAuthorization(opts *bind.TransactOpts, master common.Address, authorized bool) (*types.Transaction, error) {
	return _MasterAggregator.contract.Transact(opts, "setMasterAuthorization", master, authorized)
}

// SetMasterAuthorization is a paid mutator transaction binding the contract method 0x79d37814.
//
// Solidity: function setMasterAuthorization(address master, bool authorized) returns()
func (_MasterAggregator *MasterAggregatorSession) SetMasterAuthorization(master common.Address, authorized bool) (*types.Transaction, error) {
	return _MasterAggregator.Contract.SetMasterAuthorization(&_MasterAggregator.TransactOpts, master, authorized)
}

// SetMasterAuthorization is a paid mutator transaction binding the contract method 0x79d37814.
//
// Solidity: function setMasterAuthorization(address master, bool authorized) returns()
func (_MasterAggregator *MasterAggregatorTransactorSession) SetMasterAuthorization(master common.Address, authorized bool) (*types.Transaction, error) {
	return _MasterAggregator.Contract.SetMasterAuthorization(&_MasterAggregator.TransactOpts, master, authorized)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MasterAggregator *MasterAggregatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MasterAggregator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MasterAggregator *MasterAggregatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MasterAggregator.Contract.TransferOwnership(&_MasterAggregator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MasterAggregator *MasterAggregatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MasterAggregator.Contract.TransferOwnership(&_MasterAggregator.TransactOpts, newOwner)
}

// UnlockStake is a paid mutator transaction binding the contract method 0x4a1ce599.
//
// Solidity: function unlockStake(address entryPoint) returns()
func (_MasterAggregator *MasterAggregatorTransactor) UnlockStake(opts *bind.TransactOpts, entryPoint common.Address) (*types.Transaction, error) {
	return _MasterAggregator.contract.Transact(opts, "unlockStake", entryPoint)
}

// UnlockStake is a paid mutator transaction binding the contract method 0x4a1ce599.
//
// Solidity: function unlockStake(address entryPoint) returns()
func (_MasterAggregator *MasterAggregatorSession) UnlockStake(entryPoint common.Address) (*types.Transaction, error) {
	return _MasterAggregator.Contract.UnlockStake(&_MasterAggregator.TransactOpts, entryPoint)
}

// UnlockStake is a paid mutator transaction binding the contract method 0x4a1ce599.
//
// Solidity: function unlockStake(address entryPoint) returns()
func (_MasterAggregator *MasterAggregatorTransactorSession) UnlockStake(entryPoint common.Address) (*types.Transaction, error) {
	return _MasterAggregator.Contract.UnlockStake(&_MasterAggregator.TransactOpts, entryPoint)
}

// UpdateConfig is a paid mutator transaction binding the contract method 0x8248e722.
//
// Solidity: function updateConfig(uint256 _maxAggregatedOps, uint256 _validationWindow) returns()
func (_MasterAggregator *MasterAggregatorTransactor) UpdateConfig(opts *bind.TransactOpts, _maxAggregatedOps *big.Int, _validationWindow *big.Int) (*types.Transaction, error) {
	return _MasterAggregator.contract.Transact(opts, "updateConfig", _maxAggregatedOps, _validationWindow)
}

// UpdateConfig is a paid mutator transaction binding the contract method 0x8248e722.
//
// Solidity: function updateConfig(uint256 _maxAggregatedOps, uint256 _validationWindow) returns()
func (_MasterAggregator *MasterAggregatorSession) UpdateConfig(_maxAggregatedOps *big.Int, _validationWindow *big.Int) (*types.Transaction, error) {
	return _MasterAggregator.Contract.UpdateConfig(&_MasterAggregator.TransactOpts, _maxAggregatedOps, _validationWindow)
}

// UpdateConfig is a paid mutator transaction binding the contract method 0x8248e722.
//
// Solidity: function updateConfig(uint256 _maxAggregatedOps, uint256 _validationWindow) returns()
func (_MasterAggregator *MasterAggregatorTransactorSession) UpdateConfig(_maxAggregatedOps *big.Int, _validationWindow *big.Int) (*types.Transaction, error) {
	return _MasterAggregator.Contract.UpdateConfig(&_MasterAggregator.TransactOpts, _maxAggregatedOps, _validationWindow)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MasterAggregator *MasterAggregatorTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MasterAggregator.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MasterAggregator *MasterAggregatorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MasterAggregator.Contract.UpgradeToAndCall(&_MasterAggregator.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MasterAggregator *MasterAggregatorTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MasterAggregator.Contract.UpgradeToAndCall(&_MasterAggregator.TransactOpts, newImplementation, data)
}

// ValidateSignatures is a paid mutator transaction binding the contract method 0x2dd81133.
//
// Solidity: function validateSignatures((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] userOps, bytes signature) returns()
func (_MasterAggregator *MasterAggregatorTransactor) ValidateSignatures(opts *bind.TransactOpts, userOps []PackedUserOperation, signature []byte) (*types.Transaction, error) {
	return _MasterAggregator.contract.Transact(opts, "validateSignatures", userOps, signature)
}

// ValidateSignatures is a paid mutator transaction binding the contract method 0x2dd81133.
//
// Solidity: function validateSignatures((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] userOps, bytes signature) returns()
func (_MasterAggregator *MasterAggregatorSession) ValidateSignatures(userOps []PackedUserOperation, signature []byte) (*types.Transaction, error) {
	return _MasterAggregator.Contract.ValidateSignatures(&_MasterAggregator.TransactOpts, userOps, signature)
}

// ValidateSignatures is a paid mutator transaction binding the contract method 0x2dd81133.
//
// Solidity: function validateSignatures((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] userOps, bytes signature) returns()
func (_MasterAggregator *MasterAggregatorTransactorSession) ValidateSignatures(userOps []PackedUserOperation, signature []byte) (*types.Transaction, error) {
	return _MasterAggregator.Contract.ValidateSignatures(&_MasterAggregator.TransactOpts, userOps, signature)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xb36f9705.
//
// Solidity: function withdrawStake(address entryPoint, address withdrawAddress) returns()
func (_MasterAggregator *MasterAggregatorTransactor) WithdrawStake(opts *bind.TransactOpts, entryPoint common.Address, withdrawAddress common.Address) (*types.Transaction, error) {
	return _MasterAggregator.contract.Transact(opts, "withdrawStake", entryPoint, withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xb36f9705.
//
// Solidity: function withdrawStake(address entryPoint, address withdrawAddress) returns()
func (_MasterAggregator *MasterAggregatorSession) WithdrawStake(entryPoint common.Address, withdrawAddress common.Address) (*types.Transaction, error) {
	return _MasterAggregator.Contract.WithdrawStake(&_MasterAggregator.TransactOpts, entryPoint, withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xb36f9705.
//
// Solidity: function withdrawStake(address entryPoint, address withdrawAddress) returns()
func (_MasterAggregator *MasterAggregatorTransactorSession) WithdrawStake(entryPoint common.Address, withdrawAddress common.Address) (*types.Transaction, error) {
	return _MasterAggregator.Contract.WithdrawStake(&_MasterAggregator.TransactOpts, entryPoint, withdrawAddress)
}

// MasterAggregatorAccountAuthorizedIterator is returned from FilterAccountAuthorized and is used to iterate over the raw logs and unpacked data for AccountAuthorized events raised by the MasterAggregator contract.
type MasterAggregatorAccountAuthorizedIterator struct {
	Event *MasterAggregatorAccountAuthorized // Event containing the contract specifics and raw log

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
func (it *MasterAggregatorAccountAuthorizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterAggregatorAccountAuthorized)
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
		it.Event = new(MasterAggregatorAccountAuthorized)
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
func (it *MasterAggregatorAccountAuthorizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterAggregatorAccountAuthorizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterAggregatorAccountAuthorized represents a AccountAuthorized event raised by the MasterAggregator contract.
type MasterAggregatorAccountAuthorized struct {
	Master     common.Address
	Account    common.Address
	Authorized bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAccountAuthorized is a free log retrieval operation binding the contract event 0x31b0379890007b62e73ee4db4531159d8d00428b2426e3942589e1b604d82e8d.
//
// Solidity: event AccountAuthorized(address indexed master, address indexed Account, bool authorized)
func (_MasterAggregator *MasterAggregatorFilterer) FilterAccountAuthorized(opts *bind.FilterOpts, master []common.Address, Account []common.Address) (*MasterAggregatorAccountAuthorizedIterator, error) {

	var masterRule []interface{}
	for _, masterItem := range master {
		masterRule = append(masterRule, masterItem)
	}
	var AccountRule []interface{}
	for _, AccountItem := range Account {
		AccountRule = append(AccountRule, AccountItem)
	}

	logs, sub, err := _MasterAggregator.contract.FilterLogs(opts, "AccountAuthorized", masterRule, AccountRule)
	if err != nil {
		return nil, err
	}
	return &MasterAggregatorAccountAuthorizedIterator{contract: _MasterAggregator.contract, event: "AccountAuthorized", logs: logs, sub: sub}, nil
}

// WatchAccountAuthorized is a free log subscription operation binding the contract event 0x31b0379890007b62e73ee4db4531159d8d00428b2426e3942589e1b604d82e8d.
//
// Solidity: event AccountAuthorized(address indexed master, address indexed Account, bool authorized)
func (_MasterAggregator *MasterAggregatorFilterer) WatchAccountAuthorized(opts *bind.WatchOpts, sink chan<- *MasterAggregatorAccountAuthorized, master []common.Address, Account []common.Address) (event.Subscription, error) {

	var masterRule []interface{}
	for _, masterItem := range master {
		masterRule = append(masterRule, masterItem)
	}
	var AccountRule []interface{}
	for _, AccountItem := range Account {
		AccountRule = append(AccountRule, AccountItem)
	}

	logs, sub, err := _MasterAggregator.contract.WatchLogs(opts, "AccountAuthorized", masterRule, AccountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterAggregatorAccountAuthorized)
				if err := _MasterAggregator.contract.UnpackLog(event, "AccountAuthorized", log); err != nil {
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

// ParseAccountAuthorized is a log parse operation binding the contract event 0x31b0379890007b62e73ee4db4531159d8d00428b2426e3942589e1b604d82e8d.
//
// Solidity: event AccountAuthorized(address indexed master, address indexed Account, bool authorized)
func (_MasterAggregator *MasterAggregatorFilterer) ParseAccountAuthorized(log types.Log) (*MasterAggregatorAccountAuthorized, error) {
	event := new(MasterAggregatorAccountAuthorized)
	if err := _MasterAggregator.contract.UnpackLog(event, "AccountAuthorized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterAggregatorAggregatedValidationIterator is returned from FilterAggregatedValidation and is used to iterate over the raw logs and unpacked data for AggregatedValidation events raised by the MasterAggregator contract.
type MasterAggregatorAggregatedValidationIterator struct {
	Event *MasterAggregatorAggregatedValidation // Event containing the contract specifics and raw log

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
func (it *MasterAggregatorAggregatedValidationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterAggregatorAggregatedValidation)
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
		it.Event = new(MasterAggregatorAggregatedValidation)
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
func (it *MasterAggregatorAggregatedValidationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterAggregatorAggregatedValidationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterAggregatorAggregatedValidation represents a AggregatedValidation event raised by the MasterAggregator contract.
type MasterAggregatorAggregatedValidation struct {
	Master         common.Address
	OperationCount *big.Int
	AggregatedHash [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAggregatedValidation is a free log retrieval operation binding the contract event 0x2accd4101bebfe73d1f34f89ea9116bffa7389a0f30ec4a00b98ac9a17b6b2e6.
//
// Solidity: event AggregatedValidation(address indexed master, uint256 operationCount, bytes32 aggregatedHash)
func (_MasterAggregator *MasterAggregatorFilterer) FilterAggregatedValidation(opts *bind.FilterOpts, master []common.Address) (*MasterAggregatorAggregatedValidationIterator, error) {

	var masterRule []interface{}
	for _, masterItem := range master {
		masterRule = append(masterRule, masterItem)
	}

	logs, sub, err := _MasterAggregator.contract.FilterLogs(opts, "AggregatedValidation", masterRule)
	if err != nil {
		return nil, err
	}
	return &MasterAggregatorAggregatedValidationIterator{contract: _MasterAggregator.contract, event: "AggregatedValidation", logs: logs, sub: sub}, nil
}

// WatchAggregatedValidation is a free log subscription operation binding the contract event 0x2accd4101bebfe73d1f34f89ea9116bffa7389a0f30ec4a00b98ac9a17b6b2e6.
//
// Solidity: event AggregatedValidation(address indexed master, uint256 operationCount, bytes32 aggregatedHash)
func (_MasterAggregator *MasterAggregatorFilterer) WatchAggregatedValidation(opts *bind.WatchOpts, sink chan<- *MasterAggregatorAggregatedValidation, master []common.Address) (event.Subscription, error) {

	var masterRule []interface{}
	for _, masterItem := range master {
		masterRule = append(masterRule, masterItem)
	}

	logs, sub, err := _MasterAggregator.contract.WatchLogs(opts, "AggregatedValidation", masterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterAggregatorAggregatedValidation)
				if err := _MasterAggregator.contract.UnpackLog(event, "AggregatedValidation", log); err != nil {
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

// ParseAggregatedValidation is a log parse operation binding the contract event 0x2accd4101bebfe73d1f34f89ea9116bffa7389a0f30ec4a00b98ac9a17b6b2e6.
//
// Solidity: event AggregatedValidation(address indexed master, uint256 operationCount, bytes32 aggregatedHash)
func (_MasterAggregator *MasterAggregatorFilterer) ParseAggregatedValidation(log types.Log) (*MasterAggregatorAggregatedValidation, error) {
	event := new(MasterAggregatorAggregatedValidation)
	if err := _MasterAggregator.contract.UnpackLog(event, "AggregatedValidation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterAggregatorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MasterAggregator contract.
type MasterAggregatorInitializedIterator struct {
	Event *MasterAggregatorInitialized // Event containing the contract specifics and raw log

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
func (it *MasterAggregatorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterAggregatorInitialized)
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
		it.Event = new(MasterAggregatorInitialized)
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
func (it *MasterAggregatorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterAggregatorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterAggregatorInitialized represents a Initialized event raised by the MasterAggregator contract.
type MasterAggregatorInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MasterAggregator *MasterAggregatorFilterer) FilterInitialized(opts *bind.FilterOpts) (*MasterAggregatorInitializedIterator, error) {

	logs, sub, err := _MasterAggregator.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MasterAggregatorInitializedIterator{contract: _MasterAggregator.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MasterAggregator *MasterAggregatorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MasterAggregatorInitialized) (event.Subscription, error) {

	logs, sub, err := _MasterAggregator.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterAggregatorInitialized)
				if err := _MasterAggregator.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_MasterAggregator *MasterAggregatorFilterer) ParseInitialized(log types.Log) (*MasterAggregatorInitialized, error) {
	event := new(MasterAggregatorInitialized)
	if err := _MasterAggregator.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterAggregatorMasterAuthorizedIterator is returned from FilterMasterAuthorized and is used to iterate over the raw logs and unpacked data for MasterAuthorized events raised by the MasterAggregator contract.
type MasterAggregatorMasterAuthorizedIterator struct {
	Event *MasterAggregatorMasterAuthorized // Event containing the contract specifics and raw log

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
func (it *MasterAggregatorMasterAuthorizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterAggregatorMasterAuthorized)
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
		it.Event = new(MasterAggregatorMasterAuthorized)
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
func (it *MasterAggregatorMasterAuthorizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterAggregatorMasterAuthorizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterAggregatorMasterAuthorized represents a MasterAuthorized event raised by the MasterAggregator contract.
type MasterAggregatorMasterAuthorized struct {
	Master     common.Address
	Authorized bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMasterAuthorized is a free log retrieval operation binding the contract event 0xfdb8d49d4935686ba354500071275240db16ad1b420d5b24fd933b4ac6ce165b.
//
// Solidity: event MasterAuthorized(address indexed master, bool authorized)
func (_MasterAggregator *MasterAggregatorFilterer) FilterMasterAuthorized(opts *bind.FilterOpts, master []common.Address) (*MasterAggregatorMasterAuthorizedIterator, error) {

	var masterRule []interface{}
	for _, masterItem := range master {
		masterRule = append(masterRule, masterItem)
	}

	logs, sub, err := _MasterAggregator.contract.FilterLogs(opts, "MasterAuthorized", masterRule)
	if err != nil {
		return nil, err
	}
	return &MasterAggregatorMasterAuthorizedIterator{contract: _MasterAggregator.contract, event: "MasterAuthorized", logs: logs, sub: sub}, nil
}

// WatchMasterAuthorized is a free log subscription operation binding the contract event 0xfdb8d49d4935686ba354500071275240db16ad1b420d5b24fd933b4ac6ce165b.
//
// Solidity: event MasterAuthorized(address indexed master, bool authorized)
func (_MasterAggregator *MasterAggregatorFilterer) WatchMasterAuthorized(opts *bind.WatchOpts, sink chan<- *MasterAggregatorMasterAuthorized, master []common.Address) (event.Subscription, error) {

	var masterRule []interface{}
	for _, masterItem := range master {
		masterRule = append(masterRule, masterItem)
	}

	logs, sub, err := _MasterAggregator.contract.WatchLogs(opts, "MasterAuthorized", masterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterAggregatorMasterAuthorized)
				if err := _MasterAggregator.contract.UnpackLog(event, "MasterAuthorized", log); err != nil {
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

// ParseMasterAuthorized is a log parse operation binding the contract event 0xfdb8d49d4935686ba354500071275240db16ad1b420d5b24fd933b4ac6ce165b.
//
// Solidity: event MasterAuthorized(address indexed master, bool authorized)
func (_MasterAggregator *MasterAggregatorFilterer) ParseMasterAuthorized(log types.Log) (*MasterAggregatorMasterAuthorized, error) {
	event := new(MasterAggregatorMasterAuthorized)
	if err := _MasterAggregator.contract.UnpackLog(event, "MasterAuthorized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterAggregatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MasterAggregator contract.
type MasterAggregatorOwnershipTransferredIterator struct {
	Event *MasterAggregatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MasterAggregatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterAggregatorOwnershipTransferred)
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
		it.Event = new(MasterAggregatorOwnershipTransferred)
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
func (it *MasterAggregatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterAggregatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterAggregatorOwnershipTransferred represents a OwnershipTransferred event raised by the MasterAggregator contract.
type MasterAggregatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MasterAggregator *MasterAggregatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MasterAggregatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MasterAggregator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MasterAggregatorOwnershipTransferredIterator{contract: _MasterAggregator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MasterAggregator *MasterAggregatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MasterAggregatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MasterAggregator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterAggregatorOwnershipTransferred)
				if err := _MasterAggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MasterAggregator *MasterAggregatorFilterer) ParseOwnershipTransferred(log types.Log) (*MasterAggregatorOwnershipTransferred, error) {
	event := new(MasterAggregatorOwnershipTransferred)
	if err := _MasterAggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterAggregatorUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the MasterAggregator contract.
type MasterAggregatorUpgradedIterator struct {
	Event *MasterAggregatorUpgraded // Event containing the contract specifics and raw log

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
func (it *MasterAggregatorUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterAggregatorUpgraded)
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
		it.Event = new(MasterAggregatorUpgraded)
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
func (it *MasterAggregatorUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterAggregatorUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterAggregatorUpgraded represents a Upgraded event raised by the MasterAggregator contract.
type MasterAggregatorUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MasterAggregator *MasterAggregatorFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*MasterAggregatorUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _MasterAggregator.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &MasterAggregatorUpgradedIterator{contract: _MasterAggregator.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MasterAggregator *MasterAggregatorFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *MasterAggregatorUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _MasterAggregator.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterAggregatorUpgraded)
				if err := _MasterAggregator.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_MasterAggregator *MasterAggregatorFilterer) ParseUpgraded(log types.Log) (*MasterAggregatorUpgraded, error) {
	event := new(MasterAggregatorUpgraded)
	if err := _MasterAggregator.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
