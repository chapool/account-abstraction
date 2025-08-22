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

// GasPaymasterMetaData contains all meta data concerning the GasPaymaster contract.
var GasPaymasterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"_entryPoint\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fallbackExchangeRate\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_burnTokens\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldBeneficiary\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newBeneficiary\",\"type\":\"address\"}],\"name\":\"BeneficiaryUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"DailyLimitUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fallbackRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"FallbackRateUsed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasAmount\",\"type\":\"uint256\"}],\"name\":\"GasPayment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOracle\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOracle\",\"type\":\"address\"}],\"name\":\"OracleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"burnTokens\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"TokenHandlingModeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_DAILY_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ORACLE_PRICE_TOLERANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beneficiary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasAmount\",\"type\":\"uint256\"}],\"name\":\"canPayForGas\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasAmount\",\"type\":\"uint256\"}],\"name\":\"estimateCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fallbackExchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getDailyLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getDailyUsage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"name\":\"getDetailedGasCostEstimate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"oracleCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fallbackCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recommendedCost\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"useOracle\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFallbackExchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOracleHealthStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isHealthy\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenHandlingMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractIGasPriceOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIPaymaster.PostOpMode\",\"name\":\"mode\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualUserOpFeePerGas\",\"type\":\"uint256\"}],\"name\":\"postOp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newBeneficiary\",\"type\":\"address\"}],\"name\":\"setBeneficiary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"setDailyLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newRate\",\"type\":\"uint256\"}],\"name\":\"setFallbackExchangeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_burnTokens\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"setTokenHandlingMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20Burnable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maxCost\",\"type\":\"uint256\"}],\"name\":\"validatePaymasterUserOp\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// GasPaymasterABI is the input ABI used to generate the binding from.
// Deprecated: Use GasPaymasterMetaData.ABI instead.
var GasPaymasterABI = GasPaymasterMetaData.ABI

// GasPaymaster is an auto generated Go binding around an Ethereum contract.
type GasPaymaster struct {
	GasPaymasterCaller     // Read-only binding to the contract
	GasPaymasterTransactor // Write-only binding to the contract
	GasPaymasterFilterer   // Log filterer for contract events
}

// GasPaymasterCaller is an auto generated read-only Go binding around an Ethereum contract.
type GasPaymasterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasPaymasterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GasPaymasterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasPaymasterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GasPaymasterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasPaymasterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GasPaymasterSession struct {
	Contract     *GasPaymaster     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GasPaymasterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GasPaymasterCallerSession struct {
	Contract *GasPaymasterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// GasPaymasterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GasPaymasterTransactorSession struct {
	Contract     *GasPaymasterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// GasPaymasterRaw is an auto generated low-level Go binding around an Ethereum contract.
type GasPaymasterRaw struct {
	Contract *GasPaymaster // Generic contract binding to access the raw methods on
}

// GasPaymasterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GasPaymasterCallerRaw struct {
	Contract *GasPaymasterCaller // Generic read-only contract binding to access the raw methods on
}

// GasPaymasterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GasPaymasterTransactorRaw struct {
	Contract *GasPaymasterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGasPaymaster creates a new instance of GasPaymaster, bound to a specific deployed contract.
func NewGasPaymaster(address common.Address, backend bind.ContractBackend) (*GasPaymaster, error) {
	contract, err := bindGasPaymaster(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GasPaymaster{GasPaymasterCaller: GasPaymasterCaller{contract: contract}, GasPaymasterTransactor: GasPaymasterTransactor{contract: contract}, GasPaymasterFilterer: GasPaymasterFilterer{contract: contract}}, nil
}

// NewGasPaymasterCaller creates a new read-only instance of GasPaymaster, bound to a specific deployed contract.
func NewGasPaymasterCaller(address common.Address, caller bind.ContractCaller) (*GasPaymasterCaller, error) {
	contract, err := bindGasPaymaster(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GasPaymasterCaller{contract: contract}, nil
}

// NewGasPaymasterTransactor creates a new write-only instance of GasPaymaster, bound to a specific deployed contract.
func NewGasPaymasterTransactor(address common.Address, transactor bind.ContractTransactor) (*GasPaymasterTransactor, error) {
	contract, err := bindGasPaymaster(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GasPaymasterTransactor{contract: contract}, nil
}

// NewGasPaymasterFilterer creates a new log filterer instance of GasPaymaster, bound to a specific deployed contract.
func NewGasPaymasterFilterer(address common.Address, filterer bind.ContractFilterer) (*GasPaymasterFilterer, error) {
	contract, err := bindGasPaymaster(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GasPaymasterFilterer{contract: contract}, nil
}

// bindGasPaymaster binds a generic wrapper to an already deployed contract.
func bindGasPaymaster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GasPaymasterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasPaymaster *GasPaymasterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GasPaymaster.Contract.GasPaymasterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasPaymaster *GasPaymasterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasPaymaster.Contract.GasPaymasterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasPaymaster *GasPaymasterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasPaymaster.Contract.GasPaymasterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasPaymaster *GasPaymasterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GasPaymaster.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasPaymaster *GasPaymasterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasPaymaster.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasPaymaster *GasPaymasterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasPaymaster.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTDAILYLIMIT is a free data retrieval call binding the contract method 0xaecbdfd7.
//
// Solidity: function DEFAULT_DAILY_LIMIT() view returns(uint256)
func (_GasPaymaster *GasPaymasterCaller) DEFAULTDAILYLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GasPaymaster.contract.Call(opts, &out, "DEFAULT_DAILY_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTDAILYLIMIT is a free data retrieval call binding the contract method 0xaecbdfd7.
//
// Solidity: function DEFAULT_DAILY_LIMIT() view returns(uint256)
func (_GasPaymaster *GasPaymasterSession) DEFAULTDAILYLIMIT() (*big.Int, error) {
	return _GasPaymaster.Contract.DEFAULTDAILYLIMIT(&_GasPaymaster.CallOpts)
}

// DEFAULTDAILYLIMIT is a free data retrieval call binding the contract method 0xaecbdfd7.
//
// Solidity: function DEFAULT_DAILY_LIMIT() view returns(uint256)
func (_GasPaymaster *GasPaymasterCallerSession) DEFAULTDAILYLIMIT() (*big.Int, error) {
	return _GasPaymaster.Contract.DEFAULTDAILYLIMIT(&_GasPaymaster.CallOpts)
}

// ORACLEPRICETOLERANCE is a free data retrieval call binding the contract method 0x7408cdfb.
//
// Solidity: function ORACLE_PRICE_TOLERANCE() view returns(uint256)
func (_GasPaymaster *GasPaymasterCaller) ORACLEPRICETOLERANCE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GasPaymaster.contract.Call(opts, &out, "ORACLE_PRICE_TOLERANCE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ORACLEPRICETOLERANCE is a free data retrieval call binding the contract method 0x7408cdfb.
//
// Solidity: function ORACLE_PRICE_TOLERANCE() view returns(uint256)
func (_GasPaymaster *GasPaymasterSession) ORACLEPRICETOLERANCE() (*big.Int, error) {
	return _GasPaymaster.Contract.ORACLEPRICETOLERANCE(&_GasPaymaster.CallOpts)
}

// ORACLEPRICETOLERANCE is a free data retrieval call binding the contract method 0x7408cdfb.
//
// Solidity: function ORACLE_PRICE_TOLERANCE() view returns(uint256)
func (_GasPaymaster *GasPaymasterCallerSession) ORACLEPRICETOLERANCE() (*big.Int, error) {
	return _GasPaymaster.Contract.ORACLEPRICETOLERANCE(&_GasPaymaster.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_GasPaymaster *GasPaymasterCaller) Beneficiary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GasPaymaster.contract.Call(opts, &out, "beneficiary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_GasPaymaster *GasPaymasterSession) Beneficiary() (common.Address, error) {
	return _GasPaymaster.Contract.Beneficiary(&_GasPaymaster.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_GasPaymaster *GasPaymasterCallerSession) Beneficiary() (common.Address, error) {
	return _GasPaymaster.Contract.Beneficiary(&_GasPaymaster.CallOpts)
}

// BurnTokens is a free data retrieval call binding the contract method 0x08003f78.
//
// Solidity: function burnTokens() view returns(bool)
func (_GasPaymaster *GasPaymasterCaller) BurnTokens(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GasPaymaster.contract.Call(opts, &out, "burnTokens")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BurnTokens is a free data retrieval call binding the contract method 0x08003f78.
//
// Solidity: function burnTokens() view returns(bool)
func (_GasPaymaster *GasPaymasterSession) BurnTokens() (bool, error) {
	return _GasPaymaster.Contract.BurnTokens(&_GasPaymaster.CallOpts)
}

// BurnTokens is a free data retrieval call binding the contract method 0x08003f78.
//
// Solidity: function burnTokens() view returns(bool)
func (_GasPaymaster *GasPaymasterCallerSession) BurnTokens() (bool, error) {
	return _GasPaymaster.Contract.BurnTokens(&_GasPaymaster.CallOpts)
}

// CanPayForGas is a free data retrieval call binding the contract method 0x952dd2b2.
//
// Solidity: function canPayForGas(address user, uint256 gasAmount) view returns(bool)
func (_GasPaymaster *GasPaymasterCaller) CanPayForGas(opts *bind.CallOpts, user common.Address, gasAmount *big.Int) (bool, error) {
	var out []interface{}
	err := _GasPaymaster.contract.Call(opts, &out, "canPayForGas", user, gasAmount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanPayForGas is a free data retrieval call binding the contract method 0x952dd2b2.
//
// Solidity: function canPayForGas(address user, uint256 gasAmount) view returns(bool)
func (_GasPaymaster *GasPaymasterSession) CanPayForGas(user common.Address, gasAmount *big.Int) (bool, error) {
	return _GasPaymaster.Contract.CanPayForGas(&_GasPaymaster.CallOpts, user, gasAmount)
}

// CanPayForGas is a free data retrieval call binding the contract method 0x952dd2b2.
//
// Solidity: function canPayForGas(address user, uint256 gasAmount) view returns(bool)
func (_GasPaymaster *GasPaymasterCallerSession) CanPayForGas(user common.Address, gasAmount *big.Int) (bool, error) {
	return _GasPaymaster.Contract.CanPayForGas(&_GasPaymaster.CallOpts, user, gasAmount)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_GasPaymaster *GasPaymasterCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GasPaymaster.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_GasPaymaster *GasPaymasterSession) EntryPoint() (common.Address, error) {
	return _GasPaymaster.Contract.EntryPoint(&_GasPaymaster.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_GasPaymaster *GasPaymasterCallerSession) EntryPoint() (common.Address, error) {
	return _GasPaymaster.Contract.EntryPoint(&_GasPaymaster.CallOpts)
}

// EstimateCost is a free data retrieval call binding the contract method 0x41807a54.
//
// Solidity: function estimateCost(uint256 gasAmount) view returns(uint256)
func (_GasPaymaster *GasPaymasterCaller) EstimateCost(opts *bind.CallOpts, gasAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GasPaymaster.contract.Call(opts, &out, "estimateCost", gasAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateCost is a free data retrieval call binding the contract method 0x41807a54.
//
// Solidity: function estimateCost(uint256 gasAmount) view returns(uint256)
func (_GasPaymaster *GasPaymasterSession) EstimateCost(gasAmount *big.Int) (*big.Int, error) {
	return _GasPaymaster.Contract.EstimateCost(&_GasPaymaster.CallOpts, gasAmount)
}

// EstimateCost is a free data retrieval call binding the contract method 0x41807a54.
//
// Solidity: function estimateCost(uint256 gasAmount) view returns(uint256)
func (_GasPaymaster *GasPaymasterCallerSession) EstimateCost(gasAmount *big.Int) (*big.Int, error) {
	return _GasPaymaster.Contract.EstimateCost(&_GasPaymaster.CallOpts, gasAmount)
}

// FallbackExchangeRate is a free data retrieval call binding the contract method 0x533f3a1b.
//
// Solidity: function fallbackExchangeRate() view returns(uint256)
func (_GasPaymaster *GasPaymasterCaller) FallbackExchangeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GasPaymaster.contract.Call(opts, &out, "fallbackExchangeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FallbackExchangeRate is a free data retrieval call binding the contract method 0x533f3a1b.
//
// Solidity: function fallbackExchangeRate() view returns(uint256)
func (_GasPaymaster *GasPaymasterSession) FallbackExchangeRate() (*big.Int, error) {
	return _GasPaymaster.Contract.FallbackExchangeRate(&_GasPaymaster.CallOpts)
}

// FallbackExchangeRate is a free data retrieval call binding the contract method 0x533f3a1b.
//
// Solidity: function fallbackExchangeRate() view returns(uint256)
func (_GasPaymaster *GasPaymasterCallerSession) FallbackExchangeRate() (*big.Int, error) {
	return _GasPaymaster.Contract.FallbackExchangeRate(&_GasPaymaster.CallOpts)
}

// GetDailyLimit is a free data retrieval call binding the contract method 0x58f2ef17.
//
// Solidity: function getDailyLimit(address user) view returns(uint256)
func (_GasPaymaster *GasPaymasterCaller) GetDailyLimit(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GasPaymaster.contract.Call(opts, &out, "getDailyLimit", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDailyLimit is a free data retrieval call binding the contract method 0x58f2ef17.
//
// Solidity: function getDailyLimit(address user) view returns(uint256)
func (_GasPaymaster *GasPaymasterSession) GetDailyLimit(user common.Address) (*big.Int, error) {
	return _GasPaymaster.Contract.GetDailyLimit(&_GasPaymaster.CallOpts, user)
}

// GetDailyLimit is a free data retrieval call binding the contract method 0x58f2ef17.
//
// Solidity: function getDailyLimit(address user) view returns(uint256)
func (_GasPaymaster *GasPaymasterCallerSession) GetDailyLimit(user common.Address) (*big.Int, error) {
	return _GasPaymaster.Contract.GetDailyLimit(&_GasPaymaster.CallOpts, user)
}

// GetDailyUsage is a free data retrieval call binding the contract method 0xf11d5ea9.
//
// Solidity: function getDailyUsage(address user) view returns(uint256)
func (_GasPaymaster *GasPaymasterCaller) GetDailyUsage(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GasPaymaster.contract.Call(opts, &out, "getDailyUsage", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDailyUsage is a free data retrieval call binding the contract method 0xf11d5ea9.
//
// Solidity: function getDailyUsage(address user) view returns(uint256)
func (_GasPaymaster *GasPaymasterSession) GetDailyUsage(user common.Address) (*big.Int, error) {
	return _GasPaymaster.Contract.GetDailyUsage(&_GasPaymaster.CallOpts, user)
}

// GetDailyUsage is a free data retrieval call binding the contract method 0xf11d5ea9.
//
// Solidity: function getDailyUsage(address user) view returns(uint256)
func (_GasPaymaster *GasPaymasterCallerSession) GetDailyUsage(user common.Address) (*big.Int, error) {
	return _GasPaymaster.Contract.GetDailyUsage(&_GasPaymaster.CallOpts, user)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_GasPaymaster *GasPaymasterCaller) GetDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GasPaymaster.contract.Call(opts, &out, "getDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_GasPaymaster *GasPaymasterSession) GetDeposit() (*big.Int, error) {
	return _GasPaymaster.Contract.GetDeposit(&_GasPaymaster.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_GasPaymaster *GasPaymasterCallerSession) GetDeposit() (*big.Int, error) {
	return _GasPaymaster.Contract.GetDeposit(&_GasPaymaster.CallOpts)
}

// GetDetailedGasCostEstimate is a free data retrieval call binding the contract method 0x6b8e0892.
//
// Solidity: function getDetailedGasCostEstimate(uint256 gasLimit, uint256 gasPrice) view returns(uint256 oracleCost, uint256 fallbackCost, uint256 recommendedCost, bool useOracle)
func (_GasPaymaster *GasPaymasterCaller) GetDetailedGasCostEstimate(opts *bind.CallOpts, gasLimit *big.Int, gasPrice *big.Int) (struct {
	OracleCost      *big.Int
	FallbackCost    *big.Int
	RecommendedCost *big.Int
	UseOracle       bool
}, error) {
	var out []interface{}
	err := _GasPaymaster.contract.Call(opts, &out, "getDetailedGasCostEstimate", gasLimit, gasPrice)

	outstruct := new(struct {
		OracleCost      *big.Int
		FallbackCost    *big.Int
		RecommendedCost *big.Int
		UseOracle       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OracleCost = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FallbackCost = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RecommendedCost = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UseOracle = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// GetDetailedGasCostEstimate is a free data retrieval call binding the contract method 0x6b8e0892.
//
// Solidity: function getDetailedGasCostEstimate(uint256 gasLimit, uint256 gasPrice) view returns(uint256 oracleCost, uint256 fallbackCost, uint256 recommendedCost, bool useOracle)
func (_GasPaymaster *GasPaymasterSession) GetDetailedGasCostEstimate(gasLimit *big.Int, gasPrice *big.Int) (struct {
	OracleCost      *big.Int
	FallbackCost    *big.Int
	RecommendedCost *big.Int
	UseOracle       bool
}, error) {
	return _GasPaymaster.Contract.GetDetailedGasCostEstimate(&_GasPaymaster.CallOpts, gasLimit, gasPrice)
}

// GetDetailedGasCostEstimate is a free data retrieval call binding the contract method 0x6b8e0892.
//
// Solidity: function getDetailedGasCostEstimate(uint256 gasLimit, uint256 gasPrice) view returns(uint256 oracleCost, uint256 fallbackCost, uint256 recommendedCost, bool useOracle)
func (_GasPaymaster *GasPaymasterCallerSession) GetDetailedGasCostEstimate(gasLimit *big.Int, gasPrice *big.Int) (struct {
	OracleCost      *big.Int
	FallbackCost    *big.Int
	RecommendedCost *big.Int
	UseOracle       bool
}, error) {
	return _GasPaymaster.Contract.GetDetailedGasCostEstimate(&_GasPaymaster.CallOpts, gasLimit, gasPrice)
}

// GetFallbackExchangeRate is a free data retrieval call binding the contract method 0x0caa2c20.
//
// Solidity: function getFallbackExchangeRate() view returns(uint256)
func (_GasPaymaster *GasPaymasterCaller) GetFallbackExchangeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GasPaymaster.contract.Call(opts, &out, "getFallbackExchangeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFallbackExchangeRate is a free data retrieval call binding the contract method 0x0caa2c20.
//
// Solidity: function getFallbackExchangeRate() view returns(uint256)
func (_GasPaymaster *GasPaymasterSession) GetFallbackExchangeRate() (*big.Int, error) {
	return _GasPaymaster.Contract.GetFallbackExchangeRate(&_GasPaymaster.CallOpts)
}

// GetFallbackExchangeRate is a free data retrieval call binding the contract method 0x0caa2c20.
//
// Solidity: function getFallbackExchangeRate() view returns(uint256)
func (_GasPaymaster *GasPaymasterCallerSession) GetFallbackExchangeRate() (*big.Int, error) {
	return _GasPaymaster.Contract.GetFallbackExchangeRate(&_GasPaymaster.CallOpts)
}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_GasPaymaster *GasPaymasterCaller) GetOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GasPaymaster.contract.Call(opts, &out, "getOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_GasPaymaster *GasPaymasterSession) GetOracle() (common.Address, error) {
	return _GasPaymaster.Contract.GetOracle(&_GasPaymaster.CallOpts)
}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_GasPaymaster *GasPaymasterCallerSession) GetOracle() (common.Address, error) {
	return _GasPaymaster.Contract.GetOracle(&_GasPaymaster.CallOpts)
}

// GetOracleHealthStatus is a free data retrieval call binding the contract method 0x9a10542b.
//
// Solidity: function getOracleHealthStatus() view returns(bool isHealthy, uint256 lastUpdate)
func (_GasPaymaster *GasPaymasterCaller) GetOracleHealthStatus(opts *bind.CallOpts) (struct {
	IsHealthy  bool
	LastUpdate *big.Int
}, error) {
	var out []interface{}
	err := _GasPaymaster.contract.Call(opts, &out, "getOracleHealthStatus")

	outstruct := new(struct {
		IsHealthy  bool
		LastUpdate *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsHealthy = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.LastUpdate = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetOracleHealthStatus is a free data retrieval call binding the contract method 0x9a10542b.
//
// Solidity: function getOracleHealthStatus() view returns(bool isHealthy, uint256 lastUpdate)
func (_GasPaymaster *GasPaymasterSession) GetOracleHealthStatus() (struct {
	IsHealthy  bool
	LastUpdate *big.Int
}, error) {
	return _GasPaymaster.Contract.GetOracleHealthStatus(&_GasPaymaster.CallOpts)
}

// GetOracleHealthStatus is a free data retrieval call binding the contract method 0x9a10542b.
//
// Solidity: function getOracleHealthStatus() view returns(bool isHealthy, uint256 lastUpdate)
func (_GasPaymaster *GasPaymasterCallerSession) GetOracleHealthStatus() (struct {
	IsHealthy  bool
	LastUpdate *big.Int
}, error) {
	return _GasPaymaster.Contract.GetOracleHealthStatus(&_GasPaymaster.CallOpts)
}

// GetToken is a free data retrieval call binding the contract method 0x21df0da7.
//
// Solidity: function getToken() view returns(address)
func (_GasPaymaster *GasPaymasterCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GasPaymaster.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetToken is a free data retrieval call binding the contract method 0x21df0da7.
//
// Solidity: function getToken() view returns(address)
func (_GasPaymaster *GasPaymasterSession) GetToken() (common.Address, error) {
	return _GasPaymaster.Contract.GetToken(&_GasPaymaster.CallOpts)
}

// GetToken is a free data retrieval call binding the contract method 0x21df0da7.
//
// Solidity: function getToken() view returns(address)
func (_GasPaymaster *GasPaymasterCallerSession) GetToken() (common.Address, error) {
	return _GasPaymaster.Contract.GetToken(&_GasPaymaster.CallOpts)
}

// GetTokenHandlingMode is a free data retrieval call binding the contract method 0xbce8d4d9.
//
// Solidity: function getTokenHandlingMode() view returns(bool, address)
func (_GasPaymaster *GasPaymasterCaller) GetTokenHandlingMode(opts *bind.CallOpts) (bool, common.Address, error) {
	var out []interface{}
	err := _GasPaymaster.contract.Call(opts, &out, "getTokenHandlingMode")

	if err != nil {
		return *new(bool), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// GetTokenHandlingMode is a free data retrieval call binding the contract method 0xbce8d4d9.
//
// Solidity: function getTokenHandlingMode() view returns(bool, address)
func (_GasPaymaster *GasPaymasterSession) GetTokenHandlingMode() (bool, common.Address, error) {
	return _GasPaymaster.Contract.GetTokenHandlingMode(&_GasPaymaster.CallOpts)
}

// GetTokenHandlingMode is a free data retrieval call binding the contract method 0xbce8d4d9.
//
// Solidity: function getTokenHandlingMode() view returns(bool, address)
func (_GasPaymaster *GasPaymasterCallerSession) GetTokenHandlingMode() (bool, common.Address, error) {
	return _GasPaymaster.Contract.GetTokenHandlingMode(&_GasPaymaster.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_GasPaymaster *GasPaymasterCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GasPaymaster.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_GasPaymaster *GasPaymasterSession) Oracle() (common.Address, error) {
	return _GasPaymaster.Contract.Oracle(&_GasPaymaster.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_GasPaymaster *GasPaymasterCallerSession) Oracle() (common.Address, error) {
	return _GasPaymaster.Contract.Oracle(&_GasPaymaster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GasPaymaster *GasPaymasterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GasPaymaster.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GasPaymaster *GasPaymasterSession) Owner() (common.Address, error) {
	return _GasPaymaster.Contract.Owner(&_GasPaymaster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GasPaymaster *GasPaymasterCallerSession) Owner() (common.Address, error) {
	return _GasPaymaster.Contract.Owner(&_GasPaymaster.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GasPaymaster *GasPaymasterCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GasPaymaster.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GasPaymaster *GasPaymasterSession) Paused() (bool, error) {
	return _GasPaymaster.Contract.Paused(&_GasPaymaster.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GasPaymaster *GasPaymasterCallerSession) Paused() (bool, error) {
	return _GasPaymaster.Contract.Paused(&_GasPaymaster.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_GasPaymaster *GasPaymasterCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GasPaymaster.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_GasPaymaster *GasPaymasterSession) PendingOwner() (common.Address, error) {
	return _GasPaymaster.Contract.PendingOwner(&_GasPaymaster.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_GasPaymaster *GasPaymasterCallerSession) PendingOwner() (common.Address, error) {
	return _GasPaymaster.Contract.PendingOwner(&_GasPaymaster.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_GasPaymaster *GasPaymasterCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GasPaymaster.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_GasPaymaster *GasPaymasterSession) Token() (common.Address, error) {
	return _GasPaymaster.Contract.Token(&_GasPaymaster.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_GasPaymaster *GasPaymasterCallerSession) Token() (common.Address, error) {
	return _GasPaymaster.Contract.Token(&_GasPaymaster.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_GasPaymaster *GasPaymasterTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasPaymaster.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_GasPaymaster *GasPaymasterSession) AcceptOwnership() (*types.Transaction, error) {
	return _GasPaymaster.Contract.AcceptOwnership(&_GasPaymaster.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_GasPaymaster *GasPaymasterTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _GasPaymaster.Contract.AcceptOwnership(&_GasPaymaster.TransactOpts)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_GasPaymaster *GasPaymasterTransactor) AddStake(opts *bind.TransactOpts, unstakeDelaySec uint32) (*types.Transaction, error) {
	return _GasPaymaster.contract.Transact(opts, "addStake", unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_GasPaymaster *GasPaymasterSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _GasPaymaster.Contract.AddStake(&_GasPaymaster.TransactOpts, unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_GasPaymaster *GasPaymasterTransactorSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _GasPaymaster.Contract.AddStake(&_GasPaymaster.TransactOpts, unstakeDelaySec)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_GasPaymaster *GasPaymasterTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasPaymaster.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_GasPaymaster *GasPaymasterSession) Deposit() (*types.Transaction, error) {
	return _GasPaymaster.Contract.Deposit(&_GasPaymaster.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_GasPaymaster *GasPaymasterTransactorSession) Deposit() (*types.Transaction, error) {
	return _GasPaymaster.Contract.Deposit(&_GasPaymaster.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GasPaymaster *GasPaymasterTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasPaymaster.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GasPaymaster *GasPaymasterSession) Pause() (*types.Transaction, error) {
	return _GasPaymaster.Contract.Pause(&_GasPaymaster.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GasPaymaster *GasPaymasterTransactorSession) Pause() (*types.Transaction, error) {
	return _GasPaymaster.Contract.Pause(&_GasPaymaster.TransactOpts)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_GasPaymaster *GasPaymasterTransactor) PostOp(opts *bind.TransactOpts, mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _GasPaymaster.contract.Transact(opts, "postOp", mode, context, actualGasCost, actualUserOpFeePerGas)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_GasPaymaster *GasPaymasterSession) PostOp(mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _GasPaymaster.Contract.PostOp(&_GasPaymaster.TransactOpts, mode, context, actualGasCost, actualUserOpFeePerGas)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_GasPaymaster *GasPaymasterTransactorSession) PostOp(mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _GasPaymaster.Contract.PostOp(&_GasPaymaster.TransactOpts, mode, context, actualGasCost, actualUserOpFeePerGas)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GasPaymaster *GasPaymasterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasPaymaster.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GasPaymaster *GasPaymasterSession) RenounceOwnership() (*types.Transaction, error) {
	return _GasPaymaster.Contract.RenounceOwnership(&_GasPaymaster.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GasPaymaster *GasPaymasterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GasPaymaster.Contract.RenounceOwnership(&_GasPaymaster.TransactOpts)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x1c31f710.
//
// Solidity: function setBeneficiary(address _newBeneficiary) returns()
func (_GasPaymaster *GasPaymasterTransactor) SetBeneficiary(opts *bind.TransactOpts, _newBeneficiary common.Address) (*types.Transaction, error) {
	return _GasPaymaster.contract.Transact(opts, "setBeneficiary", _newBeneficiary)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x1c31f710.
//
// Solidity: function setBeneficiary(address _newBeneficiary) returns()
func (_GasPaymaster *GasPaymasterSession) SetBeneficiary(_newBeneficiary common.Address) (*types.Transaction, error) {
	return _GasPaymaster.Contract.SetBeneficiary(&_GasPaymaster.TransactOpts, _newBeneficiary)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x1c31f710.
//
// Solidity: function setBeneficiary(address _newBeneficiary) returns()
func (_GasPaymaster *GasPaymasterTransactorSession) SetBeneficiary(_newBeneficiary common.Address) (*types.Transaction, error) {
	return _GasPaymaster.Contract.SetBeneficiary(&_GasPaymaster.TransactOpts, _newBeneficiary)
}

// SetDailyLimit is a paid mutator transaction binding the contract method 0x2803212f.
//
// Solidity: function setDailyLimit(address user, uint256 limit) returns()
func (_GasPaymaster *GasPaymasterTransactor) SetDailyLimit(opts *bind.TransactOpts, user common.Address, limit *big.Int) (*types.Transaction, error) {
	return _GasPaymaster.contract.Transact(opts, "setDailyLimit", user, limit)
}

// SetDailyLimit is a paid mutator transaction binding the contract method 0x2803212f.
//
// Solidity: function setDailyLimit(address user, uint256 limit) returns()
func (_GasPaymaster *GasPaymasterSession) SetDailyLimit(user common.Address, limit *big.Int) (*types.Transaction, error) {
	return _GasPaymaster.Contract.SetDailyLimit(&_GasPaymaster.TransactOpts, user, limit)
}

// SetDailyLimit is a paid mutator transaction binding the contract method 0x2803212f.
//
// Solidity: function setDailyLimit(address user, uint256 limit) returns()
func (_GasPaymaster *GasPaymasterTransactorSession) SetDailyLimit(user common.Address, limit *big.Int) (*types.Transaction, error) {
	return _GasPaymaster.Contract.SetDailyLimit(&_GasPaymaster.TransactOpts, user, limit)
}

// SetFallbackExchangeRate is a paid mutator transaction binding the contract method 0x0422f0ad.
//
// Solidity: function setFallbackExchangeRate(uint256 newRate) returns()
func (_GasPaymaster *GasPaymasterTransactor) SetFallbackExchangeRate(opts *bind.TransactOpts, newRate *big.Int) (*types.Transaction, error) {
	return _GasPaymaster.contract.Transact(opts, "setFallbackExchangeRate", newRate)
}

// SetFallbackExchangeRate is a paid mutator transaction binding the contract method 0x0422f0ad.
//
// Solidity: function setFallbackExchangeRate(uint256 newRate) returns()
func (_GasPaymaster *GasPaymasterSession) SetFallbackExchangeRate(newRate *big.Int) (*types.Transaction, error) {
	return _GasPaymaster.Contract.SetFallbackExchangeRate(&_GasPaymaster.TransactOpts, newRate)
}

// SetFallbackExchangeRate is a paid mutator transaction binding the contract method 0x0422f0ad.
//
// Solidity: function setFallbackExchangeRate(uint256 newRate) returns()
func (_GasPaymaster *GasPaymasterTransactorSession) SetFallbackExchangeRate(newRate *big.Int) (*types.Transaction, error) {
	return _GasPaymaster.Contract.SetFallbackExchangeRate(&_GasPaymaster.TransactOpts, newRate)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address newOracle) returns()
func (_GasPaymaster *GasPaymasterTransactor) SetOracle(opts *bind.TransactOpts, newOracle common.Address) (*types.Transaction, error) {
	return _GasPaymaster.contract.Transact(opts, "setOracle", newOracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address newOracle) returns()
func (_GasPaymaster *GasPaymasterSession) SetOracle(newOracle common.Address) (*types.Transaction, error) {
	return _GasPaymaster.Contract.SetOracle(&_GasPaymaster.TransactOpts, newOracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address newOracle) returns()
func (_GasPaymaster *GasPaymasterTransactorSession) SetOracle(newOracle common.Address) (*types.Transaction, error) {
	return _GasPaymaster.Contract.SetOracle(&_GasPaymaster.TransactOpts, newOracle)
}

// SetTokenHandlingMode is a paid mutator transaction binding the contract method 0x4f7c989b.
//
// Solidity: function setTokenHandlingMode(bool _burnTokens, address _beneficiary) returns()
func (_GasPaymaster *GasPaymasterTransactor) SetTokenHandlingMode(opts *bind.TransactOpts, _burnTokens bool, _beneficiary common.Address) (*types.Transaction, error) {
	return _GasPaymaster.contract.Transact(opts, "setTokenHandlingMode", _burnTokens, _beneficiary)
}

// SetTokenHandlingMode is a paid mutator transaction binding the contract method 0x4f7c989b.
//
// Solidity: function setTokenHandlingMode(bool _burnTokens, address _beneficiary) returns()
func (_GasPaymaster *GasPaymasterSession) SetTokenHandlingMode(_burnTokens bool, _beneficiary common.Address) (*types.Transaction, error) {
	return _GasPaymaster.Contract.SetTokenHandlingMode(&_GasPaymaster.TransactOpts, _burnTokens, _beneficiary)
}

// SetTokenHandlingMode is a paid mutator transaction binding the contract method 0x4f7c989b.
//
// Solidity: function setTokenHandlingMode(bool _burnTokens, address _beneficiary) returns()
func (_GasPaymaster *GasPaymasterTransactorSession) SetTokenHandlingMode(_burnTokens bool, _beneficiary common.Address) (*types.Transaction, error) {
	return _GasPaymaster.Contract.SetTokenHandlingMode(&_GasPaymaster.TransactOpts, _burnTokens, _beneficiary)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GasPaymaster *GasPaymasterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GasPaymaster.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GasPaymaster *GasPaymasterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GasPaymaster.Contract.TransferOwnership(&_GasPaymaster.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GasPaymaster *GasPaymasterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GasPaymaster.Contract.TransferOwnership(&_GasPaymaster.TransactOpts, newOwner)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_GasPaymaster *GasPaymasterTransactor) UnlockStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasPaymaster.contract.Transact(opts, "unlockStake")
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_GasPaymaster *GasPaymasterSession) UnlockStake() (*types.Transaction, error) {
	return _GasPaymaster.Contract.UnlockStake(&_GasPaymaster.TransactOpts)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_GasPaymaster *GasPaymasterTransactorSession) UnlockStake() (*types.Transaction, error) {
	return _GasPaymaster.Contract.UnlockStake(&_GasPaymaster.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GasPaymaster *GasPaymasterTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasPaymaster.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GasPaymaster *GasPaymasterSession) Unpause() (*types.Transaction, error) {
	return _GasPaymaster.Contract.Unpause(&_GasPaymaster.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GasPaymaster *GasPaymasterTransactorSession) Unpause() (*types.Transaction, error) {
	return _GasPaymaster.Contract.Unpause(&_GasPaymaster.TransactOpts)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_GasPaymaster *GasPaymasterTransactor) ValidatePaymasterUserOp(opts *bind.TransactOpts, userOp PackedUserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _GasPaymaster.contract.Transact(opts, "validatePaymasterUserOp", userOp, userOpHash, maxCost)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_GasPaymaster *GasPaymasterSession) ValidatePaymasterUserOp(userOp PackedUserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _GasPaymaster.Contract.ValidatePaymasterUserOp(&_GasPaymaster.TransactOpts, userOp, userOpHash, maxCost)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_GasPaymaster *GasPaymasterTransactorSession) ValidatePaymasterUserOp(userOp PackedUserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _GasPaymaster.Contract.ValidatePaymasterUserOp(&_GasPaymaster.TransactOpts, userOp, userOpHash, maxCost)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_GasPaymaster *GasPaymasterTransactor) WithdrawStake(opts *bind.TransactOpts, withdrawAddress common.Address) (*types.Transaction, error) {
	return _GasPaymaster.contract.Transact(opts, "withdrawStake", withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_GasPaymaster *GasPaymasterSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _GasPaymaster.Contract.WithdrawStake(&_GasPaymaster.TransactOpts, withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_GasPaymaster *GasPaymasterTransactorSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _GasPaymaster.Contract.WithdrawStake(&_GasPaymaster.TransactOpts, withdrawAddress)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_GasPaymaster *GasPaymasterTransactor) WithdrawTo(opts *bind.TransactOpts, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GasPaymaster.contract.Transact(opts, "withdrawTo", withdrawAddress, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_GasPaymaster *GasPaymasterSession) WithdrawTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GasPaymaster.Contract.WithdrawTo(&_GasPaymaster.TransactOpts, withdrawAddress, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_GasPaymaster *GasPaymasterTransactorSession) WithdrawTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GasPaymaster.Contract.WithdrawTo(&_GasPaymaster.TransactOpts, withdrawAddress, amount)
}

// GasPaymasterBeneficiaryUpdatedIterator is returned from FilterBeneficiaryUpdated and is used to iterate over the raw logs and unpacked data for BeneficiaryUpdated events raised by the GasPaymaster contract.
type GasPaymasterBeneficiaryUpdatedIterator struct {
	Event *GasPaymasterBeneficiaryUpdated // Event containing the contract specifics and raw log

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
func (it *GasPaymasterBeneficiaryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasPaymasterBeneficiaryUpdated)
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
		it.Event = new(GasPaymasterBeneficiaryUpdated)
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
func (it *GasPaymasterBeneficiaryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasPaymasterBeneficiaryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasPaymasterBeneficiaryUpdated represents a BeneficiaryUpdated event raised by the GasPaymaster contract.
type GasPaymasterBeneficiaryUpdated struct {
	OldBeneficiary common.Address
	NewBeneficiary common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBeneficiaryUpdated is a free log retrieval operation binding the contract event 0xe72eaf6addaa195f3c83095031dd08f3a96808dcf047babed1fe4e4f69d6c622.
//
// Solidity: event BeneficiaryUpdated(address indexed oldBeneficiary, address indexed newBeneficiary)
func (_GasPaymaster *GasPaymasterFilterer) FilterBeneficiaryUpdated(opts *bind.FilterOpts, oldBeneficiary []common.Address, newBeneficiary []common.Address) (*GasPaymasterBeneficiaryUpdatedIterator, error) {

	var oldBeneficiaryRule []interface{}
	for _, oldBeneficiaryItem := range oldBeneficiary {
		oldBeneficiaryRule = append(oldBeneficiaryRule, oldBeneficiaryItem)
	}
	var newBeneficiaryRule []interface{}
	for _, newBeneficiaryItem := range newBeneficiary {
		newBeneficiaryRule = append(newBeneficiaryRule, newBeneficiaryItem)
	}

	logs, sub, err := _GasPaymaster.contract.FilterLogs(opts, "BeneficiaryUpdated", oldBeneficiaryRule, newBeneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &GasPaymasterBeneficiaryUpdatedIterator{contract: _GasPaymaster.contract, event: "BeneficiaryUpdated", logs: logs, sub: sub}, nil
}

// WatchBeneficiaryUpdated is a free log subscription operation binding the contract event 0xe72eaf6addaa195f3c83095031dd08f3a96808dcf047babed1fe4e4f69d6c622.
//
// Solidity: event BeneficiaryUpdated(address indexed oldBeneficiary, address indexed newBeneficiary)
func (_GasPaymaster *GasPaymasterFilterer) WatchBeneficiaryUpdated(opts *bind.WatchOpts, sink chan<- *GasPaymasterBeneficiaryUpdated, oldBeneficiary []common.Address, newBeneficiary []common.Address) (event.Subscription, error) {

	var oldBeneficiaryRule []interface{}
	for _, oldBeneficiaryItem := range oldBeneficiary {
		oldBeneficiaryRule = append(oldBeneficiaryRule, oldBeneficiaryItem)
	}
	var newBeneficiaryRule []interface{}
	for _, newBeneficiaryItem := range newBeneficiary {
		newBeneficiaryRule = append(newBeneficiaryRule, newBeneficiaryItem)
	}

	logs, sub, err := _GasPaymaster.contract.WatchLogs(opts, "BeneficiaryUpdated", oldBeneficiaryRule, newBeneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasPaymasterBeneficiaryUpdated)
				if err := _GasPaymaster.contract.UnpackLog(event, "BeneficiaryUpdated", log); err != nil {
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

// ParseBeneficiaryUpdated is a log parse operation binding the contract event 0xe72eaf6addaa195f3c83095031dd08f3a96808dcf047babed1fe4e4f69d6c622.
//
// Solidity: event BeneficiaryUpdated(address indexed oldBeneficiary, address indexed newBeneficiary)
func (_GasPaymaster *GasPaymasterFilterer) ParseBeneficiaryUpdated(log types.Log) (*GasPaymasterBeneficiaryUpdated, error) {
	event := new(GasPaymasterBeneficiaryUpdated)
	if err := _GasPaymaster.contract.UnpackLog(event, "BeneficiaryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasPaymasterDailyLimitUpdatedIterator is returned from FilterDailyLimitUpdated and is used to iterate over the raw logs and unpacked data for DailyLimitUpdated events raised by the GasPaymaster contract.
type GasPaymasterDailyLimitUpdatedIterator struct {
	Event *GasPaymasterDailyLimitUpdated // Event containing the contract specifics and raw log

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
func (it *GasPaymasterDailyLimitUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasPaymasterDailyLimitUpdated)
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
		it.Event = new(GasPaymasterDailyLimitUpdated)
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
func (it *GasPaymasterDailyLimitUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasPaymasterDailyLimitUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasPaymasterDailyLimitUpdated represents a DailyLimitUpdated event raised by the GasPaymaster contract.
type GasPaymasterDailyLimitUpdated struct {
	User     common.Address
	NewLimit *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDailyLimitUpdated is a free log retrieval operation binding the contract event 0x074424510750f675220b94bac388bff82d6407b38f341b0226af2ab45af33dc2.
//
// Solidity: event DailyLimitUpdated(address indexed user, uint256 newLimit)
func (_GasPaymaster *GasPaymasterFilterer) FilterDailyLimitUpdated(opts *bind.FilterOpts, user []common.Address) (*GasPaymasterDailyLimitUpdatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _GasPaymaster.contract.FilterLogs(opts, "DailyLimitUpdated", userRule)
	if err != nil {
		return nil, err
	}
	return &GasPaymasterDailyLimitUpdatedIterator{contract: _GasPaymaster.contract, event: "DailyLimitUpdated", logs: logs, sub: sub}, nil
}

// WatchDailyLimitUpdated is a free log subscription operation binding the contract event 0x074424510750f675220b94bac388bff82d6407b38f341b0226af2ab45af33dc2.
//
// Solidity: event DailyLimitUpdated(address indexed user, uint256 newLimit)
func (_GasPaymaster *GasPaymasterFilterer) WatchDailyLimitUpdated(opts *bind.WatchOpts, sink chan<- *GasPaymasterDailyLimitUpdated, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _GasPaymaster.contract.WatchLogs(opts, "DailyLimitUpdated", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasPaymasterDailyLimitUpdated)
				if err := _GasPaymaster.contract.UnpackLog(event, "DailyLimitUpdated", log); err != nil {
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

// ParseDailyLimitUpdated is a log parse operation binding the contract event 0x074424510750f675220b94bac388bff82d6407b38f341b0226af2ab45af33dc2.
//
// Solidity: event DailyLimitUpdated(address indexed user, uint256 newLimit)
func (_GasPaymaster *GasPaymasterFilterer) ParseDailyLimitUpdated(log types.Log) (*GasPaymasterDailyLimitUpdated, error) {
	event := new(GasPaymasterDailyLimitUpdated)
	if err := _GasPaymaster.contract.UnpackLog(event, "DailyLimitUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasPaymasterFallbackRateUsedIterator is returned from FilterFallbackRateUsed and is used to iterate over the raw logs and unpacked data for FallbackRateUsed events raised by the GasPaymaster contract.
type GasPaymasterFallbackRateUsedIterator struct {
	Event *GasPaymasterFallbackRateUsed // Event containing the contract specifics and raw log

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
func (it *GasPaymasterFallbackRateUsedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasPaymasterFallbackRateUsed)
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
		it.Event = new(GasPaymasterFallbackRateUsed)
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
func (it *GasPaymasterFallbackRateUsedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasPaymasterFallbackRateUsedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasPaymasterFallbackRateUsed represents a FallbackRateUsed event raised by the GasPaymaster contract.
type GasPaymasterFallbackRateUsed struct {
	User         common.Address
	FallbackRate *big.Int
	Reason       string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFallbackRateUsed is a free log retrieval operation binding the contract event 0xe52ec37b7b27186785250eb52ca942ad41cc14b2460efbf23a2cd4c17d43f3b6.
//
// Solidity: event FallbackRateUsed(address indexed user, uint256 fallbackRate, string reason)
func (_GasPaymaster *GasPaymasterFilterer) FilterFallbackRateUsed(opts *bind.FilterOpts, user []common.Address) (*GasPaymasterFallbackRateUsedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _GasPaymaster.contract.FilterLogs(opts, "FallbackRateUsed", userRule)
	if err != nil {
		return nil, err
	}
	return &GasPaymasterFallbackRateUsedIterator{contract: _GasPaymaster.contract, event: "FallbackRateUsed", logs: logs, sub: sub}, nil
}

// WatchFallbackRateUsed is a free log subscription operation binding the contract event 0xe52ec37b7b27186785250eb52ca942ad41cc14b2460efbf23a2cd4c17d43f3b6.
//
// Solidity: event FallbackRateUsed(address indexed user, uint256 fallbackRate, string reason)
func (_GasPaymaster *GasPaymasterFilterer) WatchFallbackRateUsed(opts *bind.WatchOpts, sink chan<- *GasPaymasterFallbackRateUsed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _GasPaymaster.contract.WatchLogs(opts, "FallbackRateUsed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasPaymasterFallbackRateUsed)
				if err := _GasPaymaster.contract.UnpackLog(event, "FallbackRateUsed", log); err != nil {
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

// ParseFallbackRateUsed is a log parse operation binding the contract event 0xe52ec37b7b27186785250eb52ca942ad41cc14b2460efbf23a2cd4c17d43f3b6.
//
// Solidity: event FallbackRateUsed(address indexed user, uint256 fallbackRate, string reason)
func (_GasPaymaster *GasPaymasterFilterer) ParseFallbackRateUsed(log types.Log) (*GasPaymasterFallbackRateUsed, error) {
	event := new(GasPaymasterFallbackRateUsed)
	if err := _GasPaymaster.contract.UnpackLog(event, "FallbackRateUsed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasPaymasterGasPaymentIterator is returned from FilterGasPayment and is used to iterate over the raw logs and unpacked data for GasPayment events raised by the GasPaymaster contract.
type GasPaymasterGasPaymentIterator struct {
	Event *GasPaymasterGasPayment // Event containing the contract specifics and raw log

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
func (it *GasPaymasterGasPaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasPaymasterGasPayment)
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
		it.Event = new(GasPaymasterGasPayment)
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
func (it *GasPaymasterGasPaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasPaymasterGasPaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasPaymasterGasPayment represents a GasPayment event raised by the GasPaymaster contract.
type GasPaymasterGasPayment struct {
	User      common.Address
	Amount    *big.Int
	GasAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGasPayment is a free log retrieval operation binding the contract event 0xf1badaa3c213d6926962a5f8e5567a267454e2d561b7faeaafc8e71e2e216b93.
//
// Solidity: event GasPayment(address indexed user, uint256 amount, uint256 gasAmount)
func (_GasPaymaster *GasPaymasterFilterer) FilterGasPayment(opts *bind.FilterOpts, user []common.Address) (*GasPaymasterGasPaymentIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _GasPaymaster.contract.FilterLogs(opts, "GasPayment", userRule)
	if err != nil {
		return nil, err
	}
	return &GasPaymasterGasPaymentIterator{contract: _GasPaymaster.contract, event: "GasPayment", logs: logs, sub: sub}, nil
}

// WatchGasPayment is a free log subscription operation binding the contract event 0xf1badaa3c213d6926962a5f8e5567a267454e2d561b7faeaafc8e71e2e216b93.
//
// Solidity: event GasPayment(address indexed user, uint256 amount, uint256 gasAmount)
func (_GasPaymaster *GasPaymasterFilterer) WatchGasPayment(opts *bind.WatchOpts, sink chan<- *GasPaymasterGasPayment, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _GasPaymaster.contract.WatchLogs(opts, "GasPayment", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasPaymasterGasPayment)
				if err := _GasPaymaster.contract.UnpackLog(event, "GasPayment", log); err != nil {
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

// ParseGasPayment is a log parse operation binding the contract event 0xf1badaa3c213d6926962a5f8e5567a267454e2d561b7faeaafc8e71e2e216b93.
//
// Solidity: event GasPayment(address indexed user, uint256 amount, uint256 gasAmount)
func (_GasPaymaster *GasPaymasterFilterer) ParseGasPayment(log types.Log) (*GasPaymasterGasPayment, error) {
	event := new(GasPaymasterGasPayment)
	if err := _GasPaymaster.contract.UnpackLog(event, "GasPayment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasPaymasterOracleUpdatedIterator is returned from FilterOracleUpdated and is used to iterate over the raw logs and unpacked data for OracleUpdated events raised by the GasPaymaster contract.
type GasPaymasterOracleUpdatedIterator struct {
	Event *GasPaymasterOracleUpdated // Event containing the contract specifics and raw log

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
func (it *GasPaymasterOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasPaymasterOracleUpdated)
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
		it.Event = new(GasPaymasterOracleUpdated)
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
func (it *GasPaymasterOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasPaymasterOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasPaymasterOracleUpdated represents a OracleUpdated event raised by the GasPaymaster contract.
type GasPaymasterOracleUpdated struct {
	OldOracle common.Address
	NewOracle common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOracleUpdated is a free log retrieval operation binding the contract event 0x078c3b417dadf69374a59793b829c52001247130433427049317bde56607b1b7.
//
// Solidity: event OracleUpdated(address indexed oldOracle, address indexed newOracle)
func (_GasPaymaster *GasPaymasterFilterer) FilterOracleUpdated(opts *bind.FilterOpts, oldOracle []common.Address, newOracle []common.Address) (*GasPaymasterOracleUpdatedIterator, error) {

	var oldOracleRule []interface{}
	for _, oldOracleItem := range oldOracle {
		oldOracleRule = append(oldOracleRule, oldOracleItem)
	}
	var newOracleRule []interface{}
	for _, newOracleItem := range newOracle {
		newOracleRule = append(newOracleRule, newOracleItem)
	}

	logs, sub, err := _GasPaymaster.contract.FilterLogs(opts, "OracleUpdated", oldOracleRule, newOracleRule)
	if err != nil {
		return nil, err
	}
	return &GasPaymasterOracleUpdatedIterator{contract: _GasPaymaster.contract, event: "OracleUpdated", logs: logs, sub: sub}, nil
}

// WatchOracleUpdated is a free log subscription operation binding the contract event 0x078c3b417dadf69374a59793b829c52001247130433427049317bde56607b1b7.
//
// Solidity: event OracleUpdated(address indexed oldOracle, address indexed newOracle)
func (_GasPaymaster *GasPaymasterFilterer) WatchOracleUpdated(opts *bind.WatchOpts, sink chan<- *GasPaymasterOracleUpdated, oldOracle []common.Address, newOracle []common.Address) (event.Subscription, error) {

	var oldOracleRule []interface{}
	for _, oldOracleItem := range oldOracle {
		oldOracleRule = append(oldOracleRule, oldOracleItem)
	}
	var newOracleRule []interface{}
	for _, newOracleItem := range newOracle {
		newOracleRule = append(newOracleRule, newOracleItem)
	}

	logs, sub, err := _GasPaymaster.contract.WatchLogs(opts, "OracleUpdated", oldOracleRule, newOracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasPaymasterOracleUpdated)
				if err := _GasPaymaster.contract.UnpackLog(event, "OracleUpdated", log); err != nil {
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

// ParseOracleUpdated is a log parse operation binding the contract event 0x078c3b417dadf69374a59793b829c52001247130433427049317bde56607b1b7.
//
// Solidity: event OracleUpdated(address indexed oldOracle, address indexed newOracle)
func (_GasPaymaster *GasPaymasterFilterer) ParseOracleUpdated(log types.Log) (*GasPaymasterOracleUpdated, error) {
	event := new(GasPaymasterOracleUpdated)
	if err := _GasPaymaster.contract.UnpackLog(event, "OracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasPaymasterOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the GasPaymaster contract.
type GasPaymasterOwnershipTransferStartedIterator struct {
	Event *GasPaymasterOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *GasPaymasterOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasPaymasterOwnershipTransferStarted)
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
		it.Event = new(GasPaymasterOwnershipTransferStarted)
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
func (it *GasPaymasterOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasPaymasterOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasPaymasterOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the GasPaymaster contract.
type GasPaymasterOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_GasPaymaster *GasPaymasterFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GasPaymasterOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GasPaymaster.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GasPaymasterOwnershipTransferStartedIterator{contract: _GasPaymaster.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_GasPaymaster *GasPaymasterFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *GasPaymasterOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GasPaymaster.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasPaymasterOwnershipTransferStarted)
				if err := _GasPaymaster.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_GasPaymaster *GasPaymasterFilterer) ParseOwnershipTransferStarted(log types.Log) (*GasPaymasterOwnershipTransferStarted, error) {
	event := new(GasPaymasterOwnershipTransferStarted)
	if err := _GasPaymaster.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasPaymasterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GasPaymaster contract.
type GasPaymasterOwnershipTransferredIterator struct {
	Event *GasPaymasterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GasPaymasterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasPaymasterOwnershipTransferred)
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
		it.Event = new(GasPaymasterOwnershipTransferred)
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
func (it *GasPaymasterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasPaymasterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasPaymasterOwnershipTransferred represents a OwnershipTransferred event raised by the GasPaymaster contract.
type GasPaymasterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GasPaymaster *GasPaymasterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GasPaymasterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GasPaymaster.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GasPaymasterOwnershipTransferredIterator{contract: _GasPaymaster.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GasPaymaster *GasPaymasterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GasPaymasterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GasPaymaster.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasPaymasterOwnershipTransferred)
				if err := _GasPaymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_GasPaymaster *GasPaymasterFilterer) ParseOwnershipTransferred(log types.Log) (*GasPaymasterOwnershipTransferred, error) {
	event := new(GasPaymasterOwnershipTransferred)
	if err := _GasPaymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasPaymasterPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the GasPaymaster contract.
type GasPaymasterPausedIterator struct {
	Event *GasPaymasterPaused // Event containing the contract specifics and raw log

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
func (it *GasPaymasterPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasPaymasterPaused)
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
		it.Event = new(GasPaymasterPaused)
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
func (it *GasPaymasterPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasPaymasterPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasPaymasterPaused represents a Paused event raised by the GasPaymaster contract.
type GasPaymasterPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GasPaymaster *GasPaymasterFilterer) FilterPaused(opts *bind.FilterOpts) (*GasPaymasterPausedIterator, error) {

	logs, sub, err := _GasPaymaster.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &GasPaymasterPausedIterator{contract: _GasPaymaster.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GasPaymaster *GasPaymasterFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *GasPaymasterPaused) (event.Subscription, error) {

	logs, sub, err := _GasPaymaster.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasPaymasterPaused)
				if err := _GasPaymaster.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GasPaymaster *GasPaymasterFilterer) ParsePaused(log types.Log) (*GasPaymasterPaused, error) {
	event := new(GasPaymasterPaused)
	if err := _GasPaymaster.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasPaymasterTokenHandlingModeUpdatedIterator is returned from FilterTokenHandlingModeUpdated and is used to iterate over the raw logs and unpacked data for TokenHandlingModeUpdated events raised by the GasPaymaster contract.
type GasPaymasterTokenHandlingModeUpdatedIterator struct {
	Event *GasPaymasterTokenHandlingModeUpdated // Event containing the contract specifics and raw log

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
func (it *GasPaymasterTokenHandlingModeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasPaymasterTokenHandlingModeUpdated)
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
		it.Event = new(GasPaymasterTokenHandlingModeUpdated)
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
func (it *GasPaymasterTokenHandlingModeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasPaymasterTokenHandlingModeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasPaymasterTokenHandlingModeUpdated represents a TokenHandlingModeUpdated event raised by the GasPaymaster contract.
type GasPaymasterTokenHandlingModeUpdated struct {
	BurnTokens  bool
	Beneficiary common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenHandlingModeUpdated is a free log retrieval operation binding the contract event 0x51e4ad57fe93123e507a804b447de9654b22837c72e05db06485242038df708f.
//
// Solidity: event TokenHandlingModeUpdated(bool burnTokens, address beneficiary)
func (_GasPaymaster *GasPaymasterFilterer) FilterTokenHandlingModeUpdated(opts *bind.FilterOpts) (*GasPaymasterTokenHandlingModeUpdatedIterator, error) {

	logs, sub, err := _GasPaymaster.contract.FilterLogs(opts, "TokenHandlingModeUpdated")
	if err != nil {
		return nil, err
	}
	return &GasPaymasterTokenHandlingModeUpdatedIterator{contract: _GasPaymaster.contract, event: "TokenHandlingModeUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenHandlingModeUpdated is a free log subscription operation binding the contract event 0x51e4ad57fe93123e507a804b447de9654b22837c72e05db06485242038df708f.
//
// Solidity: event TokenHandlingModeUpdated(bool burnTokens, address beneficiary)
func (_GasPaymaster *GasPaymasterFilterer) WatchTokenHandlingModeUpdated(opts *bind.WatchOpts, sink chan<- *GasPaymasterTokenHandlingModeUpdated) (event.Subscription, error) {

	logs, sub, err := _GasPaymaster.contract.WatchLogs(opts, "TokenHandlingModeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasPaymasterTokenHandlingModeUpdated)
				if err := _GasPaymaster.contract.UnpackLog(event, "TokenHandlingModeUpdated", log); err != nil {
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

// ParseTokenHandlingModeUpdated is a log parse operation binding the contract event 0x51e4ad57fe93123e507a804b447de9654b22837c72e05db06485242038df708f.
//
// Solidity: event TokenHandlingModeUpdated(bool burnTokens, address beneficiary)
func (_GasPaymaster *GasPaymasterFilterer) ParseTokenHandlingModeUpdated(log types.Log) (*GasPaymasterTokenHandlingModeUpdated, error) {
	event := new(GasPaymasterTokenHandlingModeUpdated)
	if err := _GasPaymaster.contract.UnpackLog(event, "TokenHandlingModeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasPaymasterTokensBurnedIterator is returned from FilterTokensBurned and is used to iterate over the raw logs and unpacked data for TokensBurned events raised by the GasPaymaster contract.
type GasPaymasterTokensBurnedIterator struct {
	Event *GasPaymasterTokensBurned // Event containing the contract specifics and raw log

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
func (it *GasPaymasterTokensBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasPaymasterTokensBurned)
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
		it.Event = new(GasPaymasterTokensBurned)
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
func (it *GasPaymasterTokensBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasPaymasterTokensBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasPaymasterTokensBurned represents a TokensBurned event raised by the GasPaymaster contract.
type GasPaymasterTokensBurned struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokensBurned is a free log retrieval operation binding the contract event 0xfd38818f5291bf0bb3a2a48aadc06ba8757865d1dabd804585338aab3009dcb6.
//
// Solidity: event TokensBurned(address indexed user, uint256 amount)
func (_GasPaymaster *GasPaymasterFilterer) FilterTokensBurned(opts *bind.FilterOpts, user []common.Address) (*GasPaymasterTokensBurnedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _GasPaymaster.contract.FilterLogs(opts, "TokensBurned", userRule)
	if err != nil {
		return nil, err
	}
	return &GasPaymasterTokensBurnedIterator{contract: _GasPaymaster.contract, event: "TokensBurned", logs: logs, sub: sub}, nil
}

// WatchTokensBurned is a free log subscription operation binding the contract event 0xfd38818f5291bf0bb3a2a48aadc06ba8757865d1dabd804585338aab3009dcb6.
//
// Solidity: event TokensBurned(address indexed user, uint256 amount)
func (_GasPaymaster *GasPaymasterFilterer) WatchTokensBurned(opts *bind.WatchOpts, sink chan<- *GasPaymasterTokensBurned, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _GasPaymaster.contract.WatchLogs(opts, "TokensBurned", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasPaymasterTokensBurned)
				if err := _GasPaymaster.contract.UnpackLog(event, "TokensBurned", log); err != nil {
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

// ParseTokensBurned is a log parse operation binding the contract event 0xfd38818f5291bf0bb3a2a48aadc06ba8757865d1dabd804585338aab3009dcb6.
//
// Solidity: event TokensBurned(address indexed user, uint256 amount)
func (_GasPaymaster *GasPaymasterFilterer) ParseTokensBurned(log types.Log) (*GasPaymasterTokensBurned, error) {
	event := new(GasPaymasterTokensBurned)
	if err := _GasPaymaster.contract.UnpackLog(event, "TokensBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasPaymasterTokensTransferredIterator is returned from FilterTokensTransferred and is used to iterate over the raw logs and unpacked data for TokensTransferred events raised by the GasPaymaster contract.
type GasPaymasterTokensTransferredIterator struct {
	Event *GasPaymasterTokensTransferred // Event containing the contract specifics and raw log

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
func (it *GasPaymasterTokensTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasPaymasterTokensTransferred)
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
		it.Event = new(GasPaymasterTokensTransferred)
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
func (it *GasPaymasterTokensTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasPaymasterTokensTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasPaymasterTokensTransferred represents a TokensTransferred event raised by the GasPaymaster contract.
type GasPaymasterTokensTransferred struct {
	User        common.Address
	Beneficiary common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokensTransferred is a free log retrieval operation binding the contract event 0x1b89874203ff7f0bba87c969ada3f32fda22ed38a6706d35199d21280c7811b1.
//
// Solidity: event TokensTransferred(address indexed user, address indexed beneficiary, uint256 amount)
func (_GasPaymaster *GasPaymasterFilterer) FilterTokensTransferred(opts *bind.FilterOpts, user []common.Address, beneficiary []common.Address) (*GasPaymasterTokensTransferredIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _GasPaymaster.contract.FilterLogs(opts, "TokensTransferred", userRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &GasPaymasterTokensTransferredIterator{contract: _GasPaymaster.contract, event: "TokensTransferred", logs: logs, sub: sub}, nil
}

// WatchTokensTransferred is a free log subscription operation binding the contract event 0x1b89874203ff7f0bba87c969ada3f32fda22ed38a6706d35199d21280c7811b1.
//
// Solidity: event TokensTransferred(address indexed user, address indexed beneficiary, uint256 amount)
func (_GasPaymaster *GasPaymasterFilterer) WatchTokensTransferred(opts *bind.WatchOpts, sink chan<- *GasPaymasterTokensTransferred, user []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _GasPaymaster.contract.WatchLogs(opts, "TokensTransferred", userRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasPaymasterTokensTransferred)
				if err := _GasPaymaster.contract.UnpackLog(event, "TokensTransferred", log); err != nil {
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

// ParseTokensTransferred is a log parse operation binding the contract event 0x1b89874203ff7f0bba87c969ada3f32fda22ed38a6706d35199d21280c7811b1.
//
// Solidity: event TokensTransferred(address indexed user, address indexed beneficiary, uint256 amount)
func (_GasPaymaster *GasPaymasterFilterer) ParseTokensTransferred(log types.Log) (*GasPaymasterTokensTransferred, error) {
	event := new(GasPaymasterTokensTransferred)
	if err := _GasPaymaster.contract.UnpackLog(event, "TokensTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasPaymasterUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the GasPaymaster contract.
type GasPaymasterUnpausedIterator struct {
	Event *GasPaymasterUnpaused // Event containing the contract specifics and raw log

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
func (it *GasPaymasterUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasPaymasterUnpaused)
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
		it.Event = new(GasPaymasterUnpaused)
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
func (it *GasPaymasterUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasPaymasterUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasPaymasterUnpaused represents a Unpaused event raised by the GasPaymaster contract.
type GasPaymasterUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GasPaymaster *GasPaymasterFilterer) FilterUnpaused(opts *bind.FilterOpts) (*GasPaymasterUnpausedIterator, error) {

	logs, sub, err := _GasPaymaster.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &GasPaymasterUnpausedIterator{contract: _GasPaymaster.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GasPaymaster *GasPaymasterFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *GasPaymasterUnpaused) (event.Subscription, error) {

	logs, sub, err := _GasPaymaster.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasPaymasterUnpaused)
				if err := _GasPaymaster.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GasPaymaster *GasPaymasterFilterer) ParseUnpaused(log types.Log) (*GasPaymasterUnpaused, error) {
	event := new(GasPaymasterUnpaused)
	if err := _GasPaymaster.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
