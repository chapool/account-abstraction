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

// GasPriceOracleMetaData contains all meta data concerning the GasPriceOracle contract.
var GasPriceOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"feedType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fallbackPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"FallbackPriceUsed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxPriceAge\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"priceDeviationThreshold\",\"type\":\"uint256\"}],\"name\":\"OracleConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"feedType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"PriceFeedUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"feedType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"initialPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fallbackPrice\",\"type\":\"uint256\"}],\"name\":\"addPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"authorizeOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorizedOracles\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"name\":\"estimateGasCostInToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenCost\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"getETHForToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"weiAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getETHPriceUSD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxPriceAge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxAge\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPriceDeviationThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"feedType\",\"type\":\"string\"}],\"name\":\"getPriceWithMetadata\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"weiAmount\",\"type\":\"uint256\"}],\"name\":\"getTokenForETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenPriceUSD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxPriceAge\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_priceDeviationThreshold\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"feedType\",\"type\":\"string\"}],\"name\":\"isPriceValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPriceAge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceDeviationThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"feedType\",\"type\":\"string\"}],\"name\":\"removePriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"revokeOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"feedType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"fallbackPrice\",\"type\":\"uint256\"}],\"name\":\"setFallbackPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxPriceAge\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_priceDeviationThreshold\",\"type\":\"uint256\"}],\"name\":\"updateOracleConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"feedType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"updatePriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// GasPriceOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use GasPriceOracleMetaData.ABI instead.
var GasPriceOracleABI = GasPriceOracleMetaData.ABI

// GasPriceOracle is an auto generated Go binding around an Ethereum contract.
type GasPriceOracle struct {
	GasPriceOracleCaller     // Read-only binding to the contract
	GasPriceOracleTransactor // Write-only binding to the contract
	GasPriceOracleFilterer   // Log filterer for contract events
}

// GasPriceOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type GasPriceOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasPriceOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GasPriceOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasPriceOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GasPriceOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasPriceOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GasPriceOracleSession struct {
	Contract     *GasPriceOracle   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GasPriceOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GasPriceOracleCallerSession struct {
	Contract *GasPriceOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// GasPriceOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GasPriceOracleTransactorSession struct {
	Contract     *GasPriceOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// GasPriceOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type GasPriceOracleRaw struct {
	Contract *GasPriceOracle // Generic contract binding to access the raw methods on
}

// GasPriceOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GasPriceOracleCallerRaw struct {
	Contract *GasPriceOracleCaller // Generic read-only contract binding to access the raw methods on
}

// GasPriceOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GasPriceOracleTransactorRaw struct {
	Contract *GasPriceOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGasPriceOracle creates a new instance of GasPriceOracle, bound to a specific deployed contract.
func NewGasPriceOracle(address common.Address, backend bind.ContractBackend) (*GasPriceOracle, error) {
	contract, err := bindGasPriceOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GasPriceOracle{GasPriceOracleCaller: GasPriceOracleCaller{contract: contract}, GasPriceOracleTransactor: GasPriceOracleTransactor{contract: contract}, GasPriceOracleFilterer: GasPriceOracleFilterer{contract: contract}}, nil
}

// NewGasPriceOracleCaller creates a new read-only instance of GasPriceOracle, bound to a specific deployed contract.
func NewGasPriceOracleCaller(address common.Address, caller bind.ContractCaller) (*GasPriceOracleCaller, error) {
	contract, err := bindGasPriceOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GasPriceOracleCaller{contract: contract}, nil
}

// NewGasPriceOracleTransactor creates a new write-only instance of GasPriceOracle, bound to a specific deployed contract.
func NewGasPriceOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*GasPriceOracleTransactor, error) {
	contract, err := bindGasPriceOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GasPriceOracleTransactor{contract: contract}, nil
}

// NewGasPriceOracleFilterer creates a new log filterer instance of GasPriceOracle, bound to a specific deployed contract.
func NewGasPriceOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*GasPriceOracleFilterer, error) {
	contract, err := bindGasPriceOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GasPriceOracleFilterer{contract: contract}, nil
}

// bindGasPriceOracle binds a generic wrapper to an already deployed contract.
func bindGasPriceOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GasPriceOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasPriceOracle *GasPriceOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GasPriceOracle.Contract.GasPriceOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasPriceOracle *GasPriceOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.GasPriceOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasPriceOracle *GasPriceOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.GasPriceOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasPriceOracle *GasPriceOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GasPriceOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasPriceOracle *GasPriceOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasPriceOracle *GasPriceOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GasPriceOracle *GasPriceOracleCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GasPriceOracle.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GasPriceOracle *GasPriceOracleSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _GasPriceOracle.Contract.UPGRADEINTERFACEVERSION(&_GasPriceOracle.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GasPriceOracle *GasPriceOracleCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _GasPriceOracle.Contract.UPGRADEINTERFACEVERSION(&_GasPriceOracle.CallOpts)
}

// AuthorizedOracles is a free data retrieval call binding the contract method 0x61c992a3.
//
// Solidity: function authorizedOracles(address ) view returns(bool)
func (_GasPriceOracle *GasPriceOracleCaller) AuthorizedOracles(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GasPriceOracle.contract.Call(opts, &out, "authorizedOracles", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AuthorizedOracles is a free data retrieval call binding the contract method 0x61c992a3.
//
// Solidity: function authorizedOracles(address ) view returns(bool)
func (_GasPriceOracle *GasPriceOracleSession) AuthorizedOracles(arg0 common.Address) (bool, error) {
	return _GasPriceOracle.Contract.AuthorizedOracles(&_GasPriceOracle.CallOpts, arg0)
}

// AuthorizedOracles is a free data retrieval call binding the contract method 0x61c992a3.
//
// Solidity: function authorizedOracles(address ) view returns(bool)
func (_GasPriceOracle *GasPriceOracleCallerSession) AuthorizedOracles(arg0 common.Address) (bool, error) {
	return _GasPriceOracle.Contract.AuthorizedOracles(&_GasPriceOracle.CallOpts, arg0)
}

// EstimateGasCostInToken is a free data retrieval call binding the contract method 0xd1dfa93c.
//
// Solidity: function estimateGasCostInToken(uint256 gasLimit, uint256 gasPrice) view returns(uint256 tokenCost)
func (_GasPriceOracle *GasPriceOracleCaller) EstimateGasCostInToken(opts *bind.CallOpts, gasLimit *big.Int, gasPrice *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GasPriceOracle.contract.Call(opts, &out, "estimateGasCostInToken", gasLimit, gasPrice)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateGasCostInToken is a free data retrieval call binding the contract method 0xd1dfa93c.
//
// Solidity: function estimateGasCostInToken(uint256 gasLimit, uint256 gasPrice) view returns(uint256 tokenCost)
func (_GasPriceOracle *GasPriceOracleSession) EstimateGasCostInToken(gasLimit *big.Int, gasPrice *big.Int) (*big.Int, error) {
	return _GasPriceOracle.Contract.EstimateGasCostInToken(&_GasPriceOracle.CallOpts, gasLimit, gasPrice)
}

// EstimateGasCostInToken is a free data retrieval call binding the contract method 0xd1dfa93c.
//
// Solidity: function estimateGasCostInToken(uint256 gasLimit, uint256 gasPrice) view returns(uint256 tokenCost)
func (_GasPriceOracle *GasPriceOracleCallerSession) EstimateGasCostInToken(gasLimit *big.Int, gasPrice *big.Int) (*big.Int, error) {
	return _GasPriceOracle.Contract.EstimateGasCostInToken(&_GasPriceOracle.CallOpts, gasLimit, gasPrice)
}

// GetETHForToken is a free data retrieval call binding the contract method 0x229a28eb.
//
// Solidity: function getETHForToken(uint256 tokenAmount) view returns(uint256 weiAmount)
func (_GasPriceOracle *GasPriceOracleCaller) GetETHForToken(opts *bind.CallOpts, tokenAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GasPriceOracle.contract.Call(opts, &out, "getETHForToken", tokenAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetETHForToken is a free data retrieval call binding the contract method 0x229a28eb.
//
// Solidity: function getETHForToken(uint256 tokenAmount) view returns(uint256 weiAmount)
func (_GasPriceOracle *GasPriceOracleSession) GetETHForToken(tokenAmount *big.Int) (*big.Int, error) {
	return _GasPriceOracle.Contract.GetETHForToken(&_GasPriceOracle.CallOpts, tokenAmount)
}

// GetETHForToken is a free data retrieval call binding the contract method 0x229a28eb.
//
// Solidity: function getETHForToken(uint256 tokenAmount) view returns(uint256 weiAmount)
func (_GasPriceOracle *GasPriceOracleCallerSession) GetETHForToken(tokenAmount *big.Int) (*big.Int, error) {
	return _GasPriceOracle.Contract.GetETHForToken(&_GasPriceOracle.CallOpts, tokenAmount)
}

// GetETHPriceUSD is a free data retrieval call binding the contract method 0xee69f708.
//
// Solidity: function getETHPriceUSD() view returns(uint256 price, uint256 timestamp)
func (_GasPriceOracle *GasPriceOracleCaller) GetETHPriceUSD(opts *bind.CallOpts) (struct {
	Price     *big.Int
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _GasPriceOracle.contract.Call(opts, &out, "getETHPriceUSD")

	outstruct := new(struct {
		Price     *big.Int
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetETHPriceUSD is a free data retrieval call binding the contract method 0xee69f708.
//
// Solidity: function getETHPriceUSD() view returns(uint256 price, uint256 timestamp)
func (_GasPriceOracle *GasPriceOracleSession) GetETHPriceUSD() (struct {
	Price     *big.Int
	Timestamp *big.Int
}, error) {
	return _GasPriceOracle.Contract.GetETHPriceUSD(&_GasPriceOracle.CallOpts)
}

// GetETHPriceUSD is a free data retrieval call binding the contract method 0xee69f708.
//
// Solidity: function getETHPriceUSD() view returns(uint256 price, uint256 timestamp)
func (_GasPriceOracle *GasPriceOracleCallerSession) GetETHPriceUSD() (struct {
	Price     *big.Int
	Timestamp *big.Int
}, error) {
	return _GasPriceOracle.Contract.GetETHPriceUSD(&_GasPriceOracle.CallOpts)
}

// GetMaxPriceAge is a free data retrieval call binding the contract method 0xa432cc8c.
//
// Solidity: function getMaxPriceAge() view returns(uint256 maxAge)
func (_GasPriceOracle *GasPriceOracleCaller) GetMaxPriceAge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GasPriceOracle.contract.Call(opts, &out, "getMaxPriceAge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxPriceAge is a free data retrieval call binding the contract method 0xa432cc8c.
//
// Solidity: function getMaxPriceAge() view returns(uint256 maxAge)
func (_GasPriceOracle *GasPriceOracleSession) GetMaxPriceAge() (*big.Int, error) {
	return _GasPriceOracle.Contract.GetMaxPriceAge(&_GasPriceOracle.CallOpts)
}

// GetMaxPriceAge is a free data retrieval call binding the contract method 0xa432cc8c.
//
// Solidity: function getMaxPriceAge() view returns(uint256 maxAge)
func (_GasPriceOracle *GasPriceOracleCallerSession) GetMaxPriceAge() (*big.Int, error) {
	return _GasPriceOracle.Contract.GetMaxPriceAge(&_GasPriceOracle.CallOpts)
}

// GetPriceDeviationThreshold is a free data retrieval call binding the contract method 0x66b84bff.
//
// Solidity: function getPriceDeviationThreshold() view returns(uint256 threshold)
func (_GasPriceOracle *GasPriceOracleCaller) GetPriceDeviationThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GasPriceOracle.contract.Call(opts, &out, "getPriceDeviationThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriceDeviationThreshold is a free data retrieval call binding the contract method 0x66b84bff.
//
// Solidity: function getPriceDeviationThreshold() view returns(uint256 threshold)
func (_GasPriceOracle *GasPriceOracleSession) GetPriceDeviationThreshold() (*big.Int, error) {
	return _GasPriceOracle.Contract.GetPriceDeviationThreshold(&_GasPriceOracle.CallOpts)
}

// GetPriceDeviationThreshold is a free data retrieval call binding the contract method 0x66b84bff.
//
// Solidity: function getPriceDeviationThreshold() view returns(uint256 threshold)
func (_GasPriceOracle *GasPriceOracleCallerSession) GetPriceDeviationThreshold() (*big.Int, error) {
	return _GasPriceOracle.Contract.GetPriceDeviationThreshold(&_GasPriceOracle.CallOpts)
}

// GetPriceWithMetadata is a free data retrieval call binding the contract method 0xb6bb7d50.
//
// Solidity: function getPriceWithMetadata(string feedType) view returns(uint256 price, uint256 timestamp, bool isValid)
func (_GasPriceOracle *GasPriceOracleCaller) GetPriceWithMetadata(opts *bind.CallOpts, feedType string) (struct {
	Price     *big.Int
	Timestamp *big.Int
	IsValid   bool
}, error) {
	var out []interface{}
	err := _GasPriceOracle.contract.Call(opts, &out, "getPriceWithMetadata", feedType)

	outstruct := new(struct {
		Price     *big.Int
		Timestamp *big.Int
		IsValid   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IsValid = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// GetPriceWithMetadata is a free data retrieval call binding the contract method 0xb6bb7d50.
//
// Solidity: function getPriceWithMetadata(string feedType) view returns(uint256 price, uint256 timestamp, bool isValid)
func (_GasPriceOracle *GasPriceOracleSession) GetPriceWithMetadata(feedType string) (struct {
	Price     *big.Int
	Timestamp *big.Int
	IsValid   bool
}, error) {
	return _GasPriceOracle.Contract.GetPriceWithMetadata(&_GasPriceOracle.CallOpts, feedType)
}

// GetPriceWithMetadata is a free data retrieval call binding the contract method 0xb6bb7d50.
//
// Solidity: function getPriceWithMetadata(string feedType) view returns(uint256 price, uint256 timestamp, bool isValid)
func (_GasPriceOracle *GasPriceOracleCallerSession) GetPriceWithMetadata(feedType string) (struct {
	Price     *big.Int
	Timestamp *big.Int
	IsValid   bool
}, error) {
	return _GasPriceOracle.Contract.GetPriceWithMetadata(&_GasPriceOracle.CallOpts, feedType)
}

// GetTokenForETH is a free data retrieval call binding the contract method 0x7f2742b7.
//
// Solidity: function getTokenForETH(uint256 weiAmount) view returns(uint256 tokenAmount)
func (_GasPriceOracle *GasPriceOracleCaller) GetTokenForETH(opts *bind.CallOpts, weiAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GasPriceOracle.contract.Call(opts, &out, "getTokenForETH", weiAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenForETH is a free data retrieval call binding the contract method 0x7f2742b7.
//
// Solidity: function getTokenForETH(uint256 weiAmount) view returns(uint256 tokenAmount)
func (_GasPriceOracle *GasPriceOracleSession) GetTokenForETH(weiAmount *big.Int) (*big.Int, error) {
	return _GasPriceOracle.Contract.GetTokenForETH(&_GasPriceOracle.CallOpts, weiAmount)
}

// GetTokenForETH is a free data retrieval call binding the contract method 0x7f2742b7.
//
// Solidity: function getTokenForETH(uint256 weiAmount) view returns(uint256 tokenAmount)
func (_GasPriceOracle *GasPriceOracleCallerSession) GetTokenForETH(weiAmount *big.Int) (*big.Int, error) {
	return _GasPriceOracle.Contract.GetTokenForETH(&_GasPriceOracle.CallOpts, weiAmount)
}

// GetTokenPriceUSD is a free data retrieval call binding the contract method 0x52254397.
//
// Solidity: function getTokenPriceUSD() view returns(uint256 price, uint256 timestamp)
func (_GasPriceOracle *GasPriceOracleCaller) GetTokenPriceUSD(opts *bind.CallOpts) (struct {
	Price     *big.Int
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _GasPriceOracle.contract.Call(opts, &out, "getTokenPriceUSD")

	outstruct := new(struct {
		Price     *big.Int
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTokenPriceUSD is a free data retrieval call binding the contract method 0x52254397.
//
// Solidity: function getTokenPriceUSD() view returns(uint256 price, uint256 timestamp)
func (_GasPriceOracle *GasPriceOracleSession) GetTokenPriceUSD() (struct {
	Price     *big.Int
	Timestamp *big.Int
}, error) {
	return _GasPriceOracle.Contract.GetTokenPriceUSD(&_GasPriceOracle.CallOpts)
}

// GetTokenPriceUSD is a free data retrieval call binding the contract method 0x52254397.
//
// Solidity: function getTokenPriceUSD() view returns(uint256 price, uint256 timestamp)
func (_GasPriceOracle *GasPriceOracleCallerSession) GetTokenPriceUSD() (struct {
	Price     *big.Int
	Timestamp *big.Int
}, error) {
	return _GasPriceOracle.Contract.GetTokenPriceUSD(&_GasPriceOracle.CallOpts)
}

// IsPriceValid is a free data retrieval call binding the contract method 0x717484a5.
//
// Solidity: function isPriceValid(string feedType) view returns(bool isValid)
func (_GasPriceOracle *GasPriceOracleCaller) IsPriceValid(opts *bind.CallOpts, feedType string) (bool, error) {
	var out []interface{}
	err := _GasPriceOracle.contract.Call(opts, &out, "isPriceValid", feedType)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPriceValid is a free data retrieval call binding the contract method 0x717484a5.
//
// Solidity: function isPriceValid(string feedType) view returns(bool isValid)
func (_GasPriceOracle *GasPriceOracleSession) IsPriceValid(feedType string) (bool, error) {
	return _GasPriceOracle.Contract.IsPriceValid(&_GasPriceOracle.CallOpts, feedType)
}

// IsPriceValid is a free data retrieval call binding the contract method 0x717484a5.
//
// Solidity: function isPriceValid(string feedType) view returns(bool isValid)
func (_GasPriceOracle *GasPriceOracleCallerSession) IsPriceValid(feedType string) (bool, error) {
	return _GasPriceOracle.Contract.IsPriceValid(&_GasPriceOracle.CallOpts, feedType)
}

// MaxPriceAge is a free data retrieval call binding the contract method 0x1584410a.
//
// Solidity: function maxPriceAge() view returns(uint256)
func (_GasPriceOracle *GasPriceOracleCaller) MaxPriceAge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GasPriceOracle.contract.Call(opts, &out, "maxPriceAge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPriceAge is a free data retrieval call binding the contract method 0x1584410a.
//
// Solidity: function maxPriceAge() view returns(uint256)
func (_GasPriceOracle *GasPriceOracleSession) MaxPriceAge() (*big.Int, error) {
	return _GasPriceOracle.Contract.MaxPriceAge(&_GasPriceOracle.CallOpts)
}

// MaxPriceAge is a free data retrieval call binding the contract method 0x1584410a.
//
// Solidity: function maxPriceAge() view returns(uint256)
func (_GasPriceOracle *GasPriceOracleCallerSession) MaxPriceAge() (*big.Int, error) {
	return _GasPriceOracle.Contract.MaxPriceAge(&_GasPriceOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GasPriceOracle *GasPriceOracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GasPriceOracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GasPriceOracle *GasPriceOracleSession) Owner() (common.Address, error) {
	return _GasPriceOracle.Contract.Owner(&_GasPriceOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GasPriceOracle *GasPriceOracleCallerSession) Owner() (common.Address, error) {
	return _GasPriceOracle.Contract.Owner(&_GasPriceOracle.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GasPriceOracle *GasPriceOracleCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GasPriceOracle.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GasPriceOracle *GasPriceOracleSession) Paused() (bool, error) {
	return _GasPriceOracle.Contract.Paused(&_GasPriceOracle.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GasPriceOracle *GasPriceOracleCallerSession) Paused() (bool, error) {
	return _GasPriceOracle.Contract.Paused(&_GasPriceOracle.CallOpts)
}

// PriceDeviationThreshold is a free data retrieval call binding the contract method 0xfa596d32.
//
// Solidity: function priceDeviationThreshold() view returns(uint256)
func (_GasPriceOracle *GasPriceOracleCaller) PriceDeviationThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GasPriceOracle.contract.Call(opts, &out, "priceDeviationThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceDeviationThreshold is a free data retrieval call binding the contract method 0xfa596d32.
//
// Solidity: function priceDeviationThreshold() view returns(uint256)
func (_GasPriceOracle *GasPriceOracleSession) PriceDeviationThreshold() (*big.Int, error) {
	return _GasPriceOracle.Contract.PriceDeviationThreshold(&_GasPriceOracle.CallOpts)
}

// PriceDeviationThreshold is a free data retrieval call binding the contract method 0xfa596d32.
//
// Solidity: function priceDeviationThreshold() view returns(uint256)
func (_GasPriceOracle *GasPriceOracleCallerSession) PriceDeviationThreshold() (*big.Int, error) {
	return _GasPriceOracle.Contract.PriceDeviationThreshold(&_GasPriceOracle.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GasPriceOracle *GasPriceOracleCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GasPriceOracle.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GasPriceOracle *GasPriceOracleSession) ProxiableUUID() ([32]byte, error) {
	return _GasPriceOracle.Contract.ProxiableUUID(&_GasPriceOracle.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GasPriceOracle *GasPriceOracleCallerSession) ProxiableUUID() ([32]byte, error) {
	return _GasPriceOracle.Contract.ProxiableUUID(&_GasPriceOracle.CallOpts)
}

// AddPriceFeed is a paid mutator transaction binding the contract method 0xf5e7b3a8.
//
// Solidity: function addPriceFeed(string feedType, uint256 initialPrice, uint256 fallbackPrice) returns()
func (_GasPriceOracle *GasPriceOracleTransactor) AddPriceFeed(opts *bind.TransactOpts, feedType string, initialPrice *big.Int, fallbackPrice *big.Int) (*types.Transaction, error) {
	return _GasPriceOracle.contract.Transact(opts, "addPriceFeed", feedType, initialPrice, fallbackPrice)
}

// AddPriceFeed is a paid mutator transaction binding the contract method 0xf5e7b3a8.
//
// Solidity: function addPriceFeed(string feedType, uint256 initialPrice, uint256 fallbackPrice) returns()
func (_GasPriceOracle *GasPriceOracleSession) AddPriceFeed(feedType string, initialPrice *big.Int, fallbackPrice *big.Int) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.AddPriceFeed(&_GasPriceOracle.TransactOpts, feedType, initialPrice, fallbackPrice)
}

// AddPriceFeed is a paid mutator transaction binding the contract method 0xf5e7b3a8.
//
// Solidity: function addPriceFeed(string feedType, uint256 initialPrice, uint256 fallbackPrice) returns()
func (_GasPriceOracle *GasPriceOracleTransactorSession) AddPriceFeed(feedType string, initialPrice *big.Int, fallbackPrice *big.Int) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.AddPriceFeed(&_GasPriceOracle.TransactOpts, feedType, initialPrice, fallbackPrice)
}

// AuthorizeOracle is a paid mutator transaction binding the contract method 0x0f13b763.
//
// Solidity: function authorizeOracle(address oracle) returns()
func (_GasPriceOracle *GasPriceOracleTransactor) AuthorizeOracle(opts *bind.TransactOpts, oracle common.Address) (*types.Transaction, error) {
	return _GasPriceOracle.contract.Transact(opts, "authorizeOracle", oracle)
}

// AuthorizeOracle is a paid mutator transaction binding the contract method 0x0f13b763.
//
// Solidity: function authorizeOracle(address oracle) returns()
func (_GasPriceOracle *GasPriceOracleSession) AuthorizeOracle(oracle common.Address) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.AuthorizeOracle(&_GasPriceOracle.TransactOpts, oracle)
}

// AuthorizeOracle is a paid mutator transaction binding the contract method 0x0f13b763.
//
// Solidity: function authorizeOracle(address oracle) returns()
func (_GasPriceOracle *GasPriceOracleTransactorSession) AuthorizeOracle(oracle common.Address) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.AuthorizeOracle(&_GasPriceOracle.TransactOpts, oracle)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address _owner, uint256 _maxPriceAge, uint256 _priceDeviationThreshold) returns()
func (_GasPriceOracle *GasPriceOracleTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _maxPriceAge *big.Int, _priceDeviationThreshold *big.Int) (*types.Transaction, error) {
	return _GasPriceOracle.contract.Transact(opts, "initialize", _owner, _maxPriceAge, _priceDeviationThreshold)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address _owner, uint256 _maxPriceAge, uint256 _priceDeviationThreshold) returns()
func (_GasPriceOracle *GasPriceOracleSession) Initialize(_owner common.Address, _maxPriceAge *big.Int, _priceDeviationThreshold *big.Int) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.Initialize(&_GasPriceOracle.TransactOpts, _owner, _maxPriceAge, _priceDeviationThreshold)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address _owner, uint256 _maxPriceAge, uint256 _priceDeviationThreshold) returns()
func (_GasPriceOracle *GasPriceOracleTransactorSession) Initialize(_owner common.Address, _maxPriceAge *big.Int, _priceDeviationThreshold *big.Int) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.Initialize(&_GasPriceOracle.TransactOpts, _owner, _maxPriceAge, _priceDeviationThreshold)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GasPriceOracle *GasPriceOracleTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasPriceOracle.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GasPriceOracle *GasPriceOracleSession) Pause() (*types.Transaction, error) {
	return _GasPriceOracle.Contract.Pause(&_GasPriceOracle.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GasPriceOracle *GasPriceOracleTransactorSession) Pause() (*types.Transaction, error) {
	return _GasPriceOracle.Contract.Pause(&_GasPriceOracle.TransactOpts)
}

// RemovePriceFeed is a paid mutator transaction binding the contract method 0xdb904dfa.
//
// Solidity: function removePriceFeed(string feedType) returns()
func (_GasPriceOracle *GasPriceOracleTransactor) RemovePriceFeed(opts *bind.TransactOpts, feedType string) (*types.Transaction, error) {
	return _GasPriceOracle.contract.Transact(opts, "removePriceFeed", feedType)
}

// RemovePriceFeed is a paid mutator transaction binding the contract method 0xdb904dfa.
//
// Solidity: function removePriceFeed(string feedType) returns()
func (_GasPriceOracle *GasPriceOracleSession) RemovePriceFeed(feedType string) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.RemovePriceFeed(&_GasPriceOracle.TransactOpts, feedType)
}

// RemovePriceFeed is a paid mutator transaction binding the contract method 0xdb904dfa.
//
// Solidity: function removePriceFeed(string feedType) returns()
func (_GasPriceOracle *GasPriceOracleTransactorSession) RemovePriceFeed(feedType string) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.RemovePriceFeed(&_GasPriceOracle.TransactOpts, feedType)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GasPriceOracle *GasPriceOracleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasPriceOracle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GasPriceOracle *GasPriceOracleSession) RenounceOwnership() (*types.Transaction, error) {
	return _GasPriceOracle.Contract.RenounceOwnership(&_GasPriceOracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GasPriceOracle *GasPriceOracleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GasPriceOracle.Contract.RenounceOwnership(&_GasPriceOracle.TransactOpts)
}

// RevokeOracle is a paid mutator transaction binding the contract method 0x5983e6b0.
//
// Solidity: function revokeOracle(address oracle) returns()
func (_GasPriceOracle *GasPriceOracleTransactor) RevokeOracle(opts *bind.TransactOpts, oracle common.Address) (*types.Transaction, error) {
	return _GasPriceOracle.contract.Transact(opts, "revokeOracle", oracle)
}

// RevokeOracle is a paid mutator transaction binding the contract method 0x5983e6b0.
//
// Solidity: function revokeOracle(address oracle) returns()
func (_GasPriceOracle *GasPriceOracleSession) RevokeOracle(oracle common.Address) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.RevokeOracle(&_GasPriceOracle.TransactOpts, oracle)
}

// RevokeOracle is a paid mutator transaction binding the contract method 0x5983e6b0.
//
// Solidity: function revokeOracle(address oracle) returns()
func (_GasPriceOracle *GasPriceOracleTransactorSession) RevokeOracle(oracle common.Address) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.RevokeOracle(&_GasPriceOracle.TransactOpts, oracle)
}

// SetFallbackPrice is a paid mutator transaction binding the contract method 0xa2fc3c56.
//
// Solidity: function setFallbackPrice(string feedType, uint256 fallbackPrice) returns()
func (_GasPriceOracle *GasPriceOracleTransactor) SetFallbackPrice(opts *bind.TransactOpts, feedType string, fallbackPrice *big.Int) (*types.Transaction, error) {
	return _GasPriceOracle.contract.Transact(opts, "setFallbackPrice", feedType, fallbackPrice)
}

// SetFallbackPrice is a paid mutator transaction binding the contract method 0xa2fc3c56.
//
// Solidity: function setFallbackPrice(string feedType, uint256 fallbackPrice) returns()
func (_GasPriceOracle *GasPriceOracleSession) SetFallbackPrice(feedType string, fallbackPrice *big.Int) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.SetFallbackPrice(&_GasPriceOracle.TransactOpts, feedType, fallbackPrice)
}

// SetFallbackPrice is a paid mutator transaction binding the contract method 0xa2fc3c56.
//
// Solidity: function setFallbackPrice(string feedType, uint256 fallbackPrice) returns()
func (_GasPriceOracle *GasPriceOracleTransactorSession) SetFallbackPrice(feedType string, fallbackPrice *big.Int) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.SetFallbackPrice(&_GasPriceOracle.TransactOpts, feedType, fallbackPrice)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GasPriceOracle *GasPriceOracleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GasPriceOracle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GasPriceOracle *GasPriceOracleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.TransferOwnership(&_GasPriceOracle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GasPriceOracle *GasPriceOracleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.TransferOwnership(&_GasPriceOracle.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GasPriceOracle *GasPriceOracleTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasPriceOracle.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GasPriceOracle *GasPriceOracleSession) Unpause() (*types.Transaction, error) {
	return _GasPriceOracle.Contract.Unpause(&_GasPriceOracle.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GasPriceOracle *GasPriceOracleTransactorSession) Unpause() (*types.Transaction, error) {
	return _GasPriceOracle.Contract.Unpause(&_GasPriceOracle.TransactOpts)
}

// UpdateOracleConfig is a paid mutator transaction binding the contract method 0xa923b2a6.
//
// Solidity: function updateOracleConfig(uint256 _maxPriceAge, uint256 _priceDeviationThreshold) returns()
func (_GasPriceOracle *GasPriceOracleTransactor) UpdateOracleConfig(opts *bind.TransactOpts, _maxPriceAge *big.Int, _priceDeviationThreshold *big.Int) (*types.Transaction, error) {
	return _GasPriceOracle.contract.Transact(opts, "updateOracleConfig", _maxPriceAge, _priceDeviationThreshold)
}

// UpdateOracleConfig is a paid mutator transaction binding the contract method 0xa923b2a6.
//
// Solidity: function updateOracleConfig(uint256 _maxPriceAge, uint256 _priceDeviationThreshold) returns()
func (_GasPriceOracle *GasPriceOracleSession) UpdateOracleConfig(_maxPriceAge *big.Int, _priceDeviationThreshold *big.Int) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.UpdateOracleConfig(&_GasPriceOracle.TransactOpts, _maxPriceAge, _priceDeviationThreshold)
}

// UpdateOracleConfig is a paid mutator transaction binding the contract method 0xa923b2a6.
//
// Solidity: function updateOracleConfig(uint256 _maxPriceAge, uint256 _priceDeviationThreshold) returns()
func (_GasPriceOracle *GasPriceOracleTransactorSession) UpdateOracleConfig(_maxPriceAge *big.Int, _priceDeviationThreshold *big.Int) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.UpdateOracleConfig(&_GasPriceOracle.TransactOpts, _maxPriceAge, _priceDeviationThreshold)
}

// UpdatePriceFeed is a paid mutator transaction binding the contract method 0xd47e99a5.
//
// Solidity: function updatePriceFeed(string feedType, uint256 newPrice) returns()
func (_GasPriceOracle *GasPriceOracleTransactor) UpdatePriceFeed(opts *bind.TransactOpts, feedType string, newPrice *big.Int) (*types.Transaction, error) {
	return _GasPriceOracle.contract.Transact(opts, "updatePriceFeed", feedType, newPrice)
}

// UpdatePriceFeed is a paid mutator transaction binding the contract method 0xd47e99a5.
//
// Solidity: function updatePriceFeed(string feedType, uint256 newPrice) returns()
func (_GasPriceOracle *GasPriceOracleSession) UpdatePriceFeed(feedType string, newPrice *big.Int) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.UpdatePriceFeed(&_GasPriceOracle.TransactOpts, feedType, newPrice)
}

// UpdatePriceFeed is a paid mutator transaction binding the contract method 0xd47e99a5.
//
// Solidity: function updatePriceFeed(string feedType, uint256 newPrice) returns()
func (_GasPriceOracle *GasPriceOracleTransactorSession) UpdatePriceFeed(feedType string, newPrice *big.Int) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.UpdatePriceFeed(&_GasPriceOracle.TransactOpts, feedType, newPrice)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GasPriceOracle *GasPriceOracleTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GasPriceOracle.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GasPriceOracle *GasPriceOracleSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.UpgradeToAndCall(&_GasPriceOracle.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GasPriceOracle *GasPriceOracleTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GasPriceOracle.Contract.UpgradeToAndCall(&_GasPriceOracle.TransactOpts, newImplementation, data)
}

// GasPriceOracleFallbackPriceUsedIterator is returned from FilterFallbackPriceUsed and is used to iterate over the raw logs and unpacked data for FallbackPriceUsed events raised by the GasPriceOracle contract.
type GasPriceOracleFallbackPriceUsedIterator struct {
	Event *GasPriceOracleFallbackPriceUsed // Event containing the contract specifics and raw log

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
func (it *GasPriceOracleFallbackPriceUsedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasPriceOracleFallbackPriceUsed)
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
		it.Event = new(GasPriceOracleFallbackPriceUsed)
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
func (it *GasPriceOracleFallbackPriceUsedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasPriceOracleFallbackPriceUsedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasPriceOracleFallbackPriceUsed represents a FallbackPriceUsed event raised by the GasPriceOracle contract.
type GasPriceOracleFallbackPriceUsed struct {
	FeedType      common.Hash
	FallbackPrice *big.Int
	Reason        string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFallbackPriceUsed is a free log retrieval operation binding the contract event 0x31ffa8ec85819b1e1d08d5d4b62b35b24b8f991f1d99ce7a2318520108d1b06c.
//
// Solidity: event FallbackPriceUsed(string indexed feedType, uint256 fallbackPrice, string reason)
func (_GasPriceOracle *GasPriceOracleFilterer) FilterFallbackPriceUsed(opts *bind.FilterOpts, feedType []string) (*GasPriceOracleFallbackPriceUsedIterator, error) {

	var feedTypeRule []interface{}
	for _, feedTypeItem := range feedType {
		feedTypeRule = append(feedTypeRule, feedTypeItem)
	}

	logs, sub, err := _GasPriceOracle.contract.FilterLogs(opts, "FallbackPriceUsed", feedTypeRule)
	if err != nil {
		return nil, err
	}
	return &GasPriceOracleFallbackPriceUsedIterator{contract: _GasPriceOracle.contract, event: "FallbackPriceUsed", logs: logs, sub: sub}, nil
}

// WatchFallbackPriceUsed is a free log subscription operation binding the contract event 0x31ffa8ec85819b1e1d08d5d4b62b35b24b8f991f1d99ce7a2318520108d1b06c.
//
// Solidity: event FallbackPriceUsed(string indexed feedType, uint256 fallbackPrice, string reason)
func (_GasPriceOracle *GasPriceOracleFilterer) WatchFallbackPriceUsed(opts *bind.WatchOpts, sink chan<- *GasPriceOracleFallbackPriceUsed, feedType []string) (event.Subscription, error) {

	var feedTypeRule []interface{}
	for _, feedTypeItem := range feedType {
		feedTypeRule = append(feedTypeRule, feedTypeItem)
	}

	logs, sub, err := _GasPriceOracle.contract.WatchLogs(opts, "FallbackPriceUsed", feedTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasPriceOracleFallbackPriceUsed)
				if err := _GasPriceOracle.contract.UnpackLog(event, "FallbackPriceUsed", log); err != nil {
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

// ParseFallbackPriceUsed is a log parse operation binding the contract event 0x31ffa8ec85819b1e1d08d5d4b62b35b24b8f991f1d99ce7a2318520108d1b06c.
//
// Solidity: event FallbackPriceUsed(string indexed feedType, uint256 fallbackPrice, string reason)
func (_GasPriceOracle *GasPriceOracleFilterer) ParseFallbackPriceUsed(log types.Log) (*GasPriceOracleFallbackPriceUsed, error) {
	event := new(GasPriceOracleFallbackPriceUsed)
	if err := _GasPriceOracle.contract.UnpackLog(event, "FallbackPriceUsed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasPriceOracleInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the GasPriceOracle contract.
type GasPriceOracleInitializedIterator struct {
	Event *GasPriceOracleInitialized // Event containing the contract specifics and raw log

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
func (it *GasPriceOracleInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasPriceOracleInitialized)
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
		it.Event = new(GasPriceOracleInitialized)
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
func (it *GasPriceOracleInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasPriceOracleInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasPriceOracleInitialized represents a Initialized event raised by the GasPriceOracle contract.
type GasPriceOracleInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_GasPriceOracle *GasPriceOracleFilterer) FilterInitialized(opts *bind.FilterOpts) (*GasPriceOracleInitializedIterator, error) {

	logs, sub, err := _GasPriceOracle.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GasPriceOracleInitializedIterator{contract: _GasPriceOracle.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_GasPriceOracle *GasPriceOracleFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GasPriceOracleInitialized) (event.Subscription, error) {

	logs, sub, err := _GasPriceOracle.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasPriceOracleInitialized)
				if err := _GasPriceOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_GasPriceOracle *GasPriceOracleFilterer) ParseInitialized(log types.Log) (*GasPriceOracleInitialized, error) {
	event := new(GasPriceOracleInitialized)
	if err := _GasPriceOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasPriceOracleOracleConfigUpdatedIterator is returned from FilterOracleConfigUpdated and is used to iterate over the raw logs and unpacked data for OracleConfigUpdated events raised by the GasPriceOracle contract.
type GasPriceOracleOracleConfigUpdatedIterator struct {
	Event *GasPriceOracleOracleConfigUpdated // Event containing the contract specifics and raw log

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
func (it *GasPriceOracleOracleConfigUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasPriceOracleOracleConfigUpdated)
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
		it.Event = new(GasPriceOracleOracleConfigUpdated)
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
func (it *GasPriceOracleOracleConfigUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasPriceOracleOracleConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasPriceOracleOracleConfigUpdated represents a OracleConfigUpdated event raised by the GasPriceOracle contract.
type GasPriceOracleOracleConfigUpdated struct {
	MaxPriceAge             *big.Int
	PriceDeviationThreshold *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterOracleConfigUpdated is a free log retrieval operation binding the contract event 0xb41a7d63ac4bcd91b617433dfbcee4eabd378e3ca1baed149a6d0d38a2c6a4c4.
//
// Solidity: event OracleConfigUpdated(uint256 maxPriceAge, uint256 priceDeviationThreshold)
func (_GasPriceOracle *GasPriceOracleFilterer) FilterOracleConfigUpdated(opts *bind.FilterOpts) (*GasPriceOracleOracleConfigUpdatedIterator, error) {

	logs, sub, err := _GasPriceOracle.contract.FilterLogs(opts, "OracleConfigUpdated")
	if err != nil {
		return nil, err
	}
	return &GasPriceOracleOracleConfigUpdatedIterator{contract: _GasPriceOracle.contract, event: "OracleConfigUpdated", logs: logs, sub: sub}, nil
}

// WatchOracleConfigUpdated is a free log subscription operation binding the contract event 0xb41a7d63ac4bcd91b617433dfbcee4eabd378e3ca1baed149a6d0d38a2c6a4c4.
//
// Solidity: event OracleConfigUpdated(uint256 maxPriceAge, uint256 priceDeviationThreshold)
func (_GasPriceOracle *GasPriceOracleFilterer) WatchOracleConfigUpdated(opts *bind.WatchOpts, sink chan<- *GasPriceOracleOracleConfigUpdated) (event.Subscription, error) {

	logs, sub, err := _GasPriceOracle.contract.WatchLogs(opts, "OracleConfigUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasPriceOracleOracleConfigUpdated)
				if err := _GasPriceOracle.contract.UnpackLog(event, "OracleConfigUpdated", log); err != nil {
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

// ParseOracleConfigUpdated is a log parse operation binding the contract event 0xb41a7d63ac4bcd91b617433dfbcee4eabd378e3ca1baed149a6d0d38a2c6a4c4.
//
// Solidity: event OracleConfigUpdated(uint256 maxPriceAge, uint256 priceDeviationThreshold)
func (_GasPriceOracle *GasPriceOracleFilterer) ParseOracleConfigUpdated(log types.Log) (*GasPriceOracleOracleConfigUpdated, error) {
	event := new(GasPriceOracleOracleConfigUpdated)
	if err := _GasPriceOracle.contract.UnpackLog(event, "OracleConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasPriceOracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GasPriceOracle contract.
type GasPriceOracleOwnershipTransferredIterator struct {
	Event *GasPriceOracleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GasPriceOracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasPriceOracleOwnershipTransferred)
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
		it.Event = new(GasPriceOracleOwnershipTransferred)
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
func (it *GasPriceOracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasPriceOracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasPriceOracleOwnershipTransferred represents a OwnershipTransferred event raised by the GasPriceOracle contract.
type GasPriceOracleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GasPriceOracle *GasPriceOracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GasPriceOracleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GasPriceOracle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GasPriceOracleOwnershipTransferredIterator{contract: _GasPriceOracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GasPriceOracle *GasPriceOracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GasPriceOracleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GasPriceOracle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasPriceOracleOwnershipTransferred)
				if err := _GasPriceOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_GasPriceOracle *GasPriceOracleFilterer) ParseOwnershipTransferred(log types.Log) (*GasPriceOracleOwnershipTransferred, error) {
	event := new(GasPriceOracleOwnershipTransferred)
	if err := _GasPriceOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasPriceOraclePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the GasPriceOracle contract.
type GasPriceOraclePausedIterator struct {
	Event *GasPriceOraclePaused // Event containing the contract specifics and raw log

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
func (it *GasPriceOraclePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasPriceOraclePaused)
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
		it.Event = new(GasPriceOraclePaused)
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
func (it *GasPriceOraclePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasPriceOraclePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasPriceOraclePaused represents a Paused event raised by the GasPriceOracle contract.
type GasPriceOraclePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GasPriceOracle *GasPriceOracleFilterer) FilterPaused(opts *bind.FilterOpts) (*GasPriceOraclePausedIterator, error) {

	logs, sub, err := _GasPriceOracle.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &GasPriceOraclePausedIterator{contract: _GasPriceOracle.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GasPriceOracle *GasPriceOracleFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *GasPriceOraclePaused) (event.Subscription, error) {

	logs, sub, err := _GasPriceOracle.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasPriceOraclePaused)
				if err := _GasPriceOracle.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_GasPriceOracle *GasPriceOracleFilterer) ParsePaused(log types.Log) (*GasPriceOraclePaused, error) {
	event := new(GasPriceOraclePaused)
	if err := _GasPriceOracle.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasPriceOraclePriceFeedUpdatedIterator is returned from FilterPriceFeedUpdated and is used to iterate over the raw logs and unpacked data for PriceFeedUpdated events raised by the GasPriceOracle contract.
type GasPriceOraclePriceFeedUpdatedIterator struct {
	Event *GasPriceOraclePriceFeedUpdated // Event containing the contract specifics and raw log

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
func (it *GasPriceOraclePriceFeedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasPriceOraclePriceFeedUpdated)
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
		it.Event = new(GasPriceOraclePriceFeedUpdated)
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
func (it *GasPriceOraclePriceFeedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasPriceOraclePriceFeedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasPriceOraclePriceFeedUpdated represents a PriceFeedUpdated event raised by the GasPriceOracle contract.
type GasPriceOraclePriceFeedUpdated struct {
	FeedType  common.Hash
	OldPrice  *big.Int
	NewPrice  *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPriceFeedUpdated is a free log retrieval operation binding the contract event 0x64df10a41529acca4897b76734b8b647490d64d8924ef7dda74db7799342584f.
//
// Solidity: event PriceFeedUpdated(string indexed feedType, uint256 oldPrice, uint256 newPrice, uint256 timestamp)
func (_GasPriceOracle *GasPriceOracleFilterer) FilterPriceFeedUpdated(opts *bind.FilterOpts, feedType []string) (*GasPriceOraclePriceFeedUpdatedIterator, error) {

	var feedTypeRule []interface{}
	for _, feedTypeItem := range feedType {
		feedTypeRule = append(feedTypeRule, feedTypeItem)
	}

	logs, sub, err := _GasPriceOracle.contract.FilterLogs(opts, "PriceFeedUpdated", feedTypeRule)
	if err != nil {
		return nil, err
	}
	return &GasPriceOraclePriceFeedUpdatedIterator{contract: _GasPriceOracle.contract, event: "PriceFeedUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceFeedUpdated is a free log subscription operation binding the contract event 0x64df10a41529acca4897b76734b8b647490d64d8924ef7dda74db7799342584f.
//
// Solidity: event PriceFeedUpdated(string indexed feedType, uint256 oldPrice, uint256 newPrice, uint256 timestamp)
func (_GasPriceOracle *GasPriceOracleFilterer) WatchPriceFeedUpdated(opts *bind.WatchOpts, sink chan<- *GasPriceOraclePriceFeedUpdated, feedType []string) (event.Subscription, error) {

	var feedTypeRule []interface{}
	for _, feedTypeItem := range feedType {
		feedTypeRule = append(feedTypeRule, feedTypeItem)
	}

	logs, sub, err := _GasPriceOracle.contract.WatchLogs(opts, "PriceFeedUpdated", feedTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasPriceOraclePriceFeedUpdated)
				if err := _GasPriceOracle.contract.UnpackLog(event, "PriceFeedUpdated", log); err != nil {
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

// ParsePriceFeedUpdated is a log parse operation binding the contract event 0x64df10a41529acca4897b76734b8b647490d64d8924ef7dda74db7799342584f.
//
// Solidity: event PriceFeedUpdated(string indexed feedType, uint256 oldPrice, uint256 newPrice, uint256 timestamp)
func (_GasPriceOracle *GasPriceOracleFilterer) ParsePriceFeedUpdated(log types.Log) (*GasPriceOraclePriceFeedUpdated, error) {
	event := new(GasPriceOraclePriceFeedUpdated)
	if err := _GasPriceOracle.contract.UnpackLog(event, "PriceFeedUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasPriceOracleUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the GasPriceOracle contract.
type GasPriceOracleUnpausedIterator struct {
	Event *GasPriceOracleUnpaused // Event containing the contract specifics and raw log

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
func (it *GasPriceOracleUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasPriceOracleUnpaused)
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
		it.Event = new(GasPriceOracleUnpaused)
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
func (it *GasPriceOracleUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasPriceOracleUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasPriceOracleUnpaused represents a Unpaused event raised by the GasPriceOracle contract.
type GasPriceOracleUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GasPriceOracle *GasPriceOracleFilterer) FilterUnpaused(opts *bind.FilterOpts) (*GasPriceOracleUnpausedIterator, error) {

	logs, sub, err := _GasPriceOracle.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &GasPriceOracleUnpausedIterator{contract: _GasPriceOracle.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GasPriceOracle *GasPriceOracleFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *GasPriceOracleUnpaused) (event.Subscription, error) {

	logs, sub, err := _GasPriceOracle.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasPriceOracleUnpaused)
				if err := _GasPriceOracle.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_GasPriceOracle *GasPriceOracleFilterer) ParseUnpaused(log types.Log) (*GasPriceOracleUnpaused, error) {
	event := new(GasPriceOracleUnpaused)
	if err := _GasPriceOracle.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasPriceOracleUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the GasPriceOracle contract.
type GasPriceOracleUpgradedIterator struct {
	Event *GasPriceOracleUpgraded // Event containing the contract specifics and raw log

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
func (it *GasPriceOracleUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasPriceOracleUpgraded)
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
		it.Event = new(GasPriceOracleUpgraded)
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
func (it *GasPriceOracleUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasPriceOracleUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasPriceOracleUpgraded represents a Upgraded event raised by the GasPriceOracle contract.
type GasPriceOracleUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GasPriceOracle *GasPriceOracleFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*GasPriceOracleUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GasPriceOracle.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &GasPriceOracleUpgradedIterator{contract: _GasPriceOracle.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GasPriceOracle *GasPriceOracleFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *GasPriceOracleUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GasPriceOracle.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasPriceOracleUpgraded)
				if err := _GasPriceOracle.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_GasPriceOracle *GasPriceOracleFilterer) ParseUpgraded(log types.Log) (*GasPriceOracleUpgraded, error) {
	event := new(GasPriceOracleUpgraded)
	if err := _GasPriceOracle.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
