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

// NFTBoostControllerMetaData contains all meta data concerning the NFTBoostController contract.
var NFTBoostControllerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"boostBps\",\"type\":\"uint256\"}],\"name\":\"BoostActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"BoostDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BOOST_A_BPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BOOST_B_BPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BOOST_C_BPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BOOST_NORMAL_BPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BOOST_SSS_BPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BOOST_SS_BPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BOOST_S_BPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"activateBoost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"activeTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cpnft\",\"outputs\":[{\"internalType\":\"contractICPNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateBoost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getActiveNFT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boostBps\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"}],\"name\":\"getLevelBoostBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getNFTBoostBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getNFTLevel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasActiveToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_cpnft\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_cpnft\",\"type\":\"address\"}],\"name\":\"setCPNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"boostBps\",\"type\":\"uint256\"}],\"name\":\"setLevelBoostBps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// NFTBoostControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use NFTBoostControllerMetaData.ABI instead.
var NFTBoostControllerABI = NFTBoostControllerMetaData.ABI

// NFTBoostController is an auto generated Go binding around an Ethereum contract.
type NFTBoostController struct {
	NFTBoostControllerCaller     // Read-only binding to the contract
	NFTBoostControllerTransactor // Write-only binding to the contract
	NFTBoostControllerFilterer   // Log filterer for contract events
}

// NFTBoostControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type NFTBoostControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTBoostControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NFTBoostControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTBoostControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NFTBoostControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTBoostControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NFTBoostControllerSession struct {
	Contract     *NFTBoostController // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// NFTBoostControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NFTBoostControllerCallerSession struct {
	Contract *NFTBoostControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// NFTBoostControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NFTBoostControllerTransactorSession struct {
	Contract     *NFTBoostControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// NFTBoostControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type NFTBoostControllerRaw struct {
	Contract *NFTBoostController // Generic contract binding to access the raw methods on
}

// NFTBoostControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NFTBoostControllerCallerRaw struct {
	Contract *NFTBoostControllerCaller // Generic read-only contract binding to access the raw methods on
}

// NFTBoostControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NFTBoostControllerTransactorRaw struct {
	Contract *NFTBoostControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNFTBoostController creates a new instance of NFTBoostController, bound to a specific deployed contract.
func NewNFTBoostController(address common.Address, backend bind.ContractBackend) (*NFTBoostController, error) {
	contract, err := bindNFTBoostController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NFTBoostController{NFTBoostControllerCaller: NFTBoostControllerCaller{contract: contract}, NFTBoostControllerTransactor: NFTBoostControllerTransactor{contract: contract}, NFTBoostControllerFilterer: NFTBoostControllerFilterer{contract: contract}}, nil
}

// NewNFTBoostControllerCaller creates a new read-only instance of NFTBoostController, bound to a specific deployed contract.
func NewNFTBoostControllerCaller(address common.Address, caller bind.ContractCaller) (*NFTBoostControllerCaller, error) {
	contract, err := bindNFTBoostController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NFTBoostControllerCaller{contract: contract}, nil
}

// NewNFTBoostControllerTransactor creates a new write-only instance of NFTBoostController, bound to a specific deployed contract.
func NewNFTBoostControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*NFTBoostControllerTransactor, error) {
	contract, err := bindNFTBoostController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NFTBoostControllerTransactor{contract: contract}, nil
}

// NewNFTBoostControllerFilterer creates a new log filterer instance of NFTBoostController, bound to a specific deployed contract.
func NewNFTBoostControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*NFTBoostControllerFilterer, error) {
	contract, err := bindNFTBoostController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NFTBoostControllerFilterer{contract: contract}, nil
}

// bindNFTBoostController binds a generic wrapper to an already deployed contract.
func bindNFTBoostController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NFTBoostControllerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFTBoostController *NFTBoostControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFTBoostController.Contract.NFTBoostControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFTBoostController *NFTBoostControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTBoostController.Contract.NFTBoostControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFTBoostController *NFTBoostControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFTBoostController.Contract.NFTBoostControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFTBoostController *NFTBoostControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFTBoostController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFTBoostController *NFTBoostControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTBoostController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFTBoostController *NFTBoostControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFTBoostController.Contract.contract.Transact(opts, method, params...)
}

// BOOSTABPS is a free data retrieval call binding the contract method 0x49c61ce1.
//
// Solidity: function BOOST_A_BPS() view returns(uint256)
func (_NFTBoostController *NFTBoostControllerCaller) BOOSTABPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTBoostController.contract.Call(opts, &out, "BOOST_A_BPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BOOSTABPS is a free data retrieval call binding the contract method 0x49c61ce1.
//
// Solidity: function BOOST_A_BPS() view returns(uint256)
func (_NFTBoostController *NFTBoostControllerSession) BOOSTABPS() (*big.Int, error) {
	return _NFTBoostController.Contract.BOOSTABPS(&_NFTBoostController.CallOpts)
}

// BOOSTABPS is a free data retrieval call binding the contract method 0x49c61ce1.
//
// Solidity: function BOOST_A_BPS() view returns(uint256)
func (_NFTBoostController *NFTBoostControllerCallerSession) BOOSTABPS() (*big.Int, error) {
	return _NFTBoostController.Contract.BOOSTABPS(&_NFTBoostController.CallOpts)
}

// BOOSTBBPS is a free data retrieval call binding the contract method 0x4b34067b.
//
// Solidity: function BOOST_B_BPS() view returns(uint256)
func (_NFTBoostController *NFTBoostControllerCaller) BOOSTBBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTBoostController.contract.Call(opts, &out, "BOOST_B_BPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BOOSTBBPS is a free data retrieval call binding the contract method 0x4b34067b.
//
// Solidity: function BOOST_B_BPS() view returns(uint256)
func (_NFTBoostController *NFTBoostControllerSession) BOOSTBBPS() (*big.Int, error) {
	return _NFTBoostController.Contract.BOOSTBBPS(&_NFTBoostController.CallOpts)
}

// BOOSTBBPS is a free data retrieval call binding the contract method 0x4b34067b.
//
// Solidity: function BOOST_B_BPS() view returns(uint256)
func (_NFTBoostController *NFTBoostControllerCallerSession) BOOSTBBPS() (*big.Int, error) {
	return _NFTBoostController.Contract.BOOSTBBPS(&_NFTBoostController.CallOpts)
}

// BOOSTCBPS is a free data retrieval call binding the contract method 0x8ea1f274.
//
// Solidity: function BOOST_C_BPS() view returns(uint256)
func (_NFTBoostController *NFTBoostControllerCaller) BOOSTCBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTBoostController.contract.Call(opts, &out, "BOOST_C_BPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BOOSTCBPS is a free data retrieval call binding the contract method 0x8ea1f274.
//
// Solidity: function BOOST_C_BPS() view returns(uint256)
func (_NFTBoostController *NFTBoostControllerSession) BOOSTCBPS() (*big.Int, error) {
	return _NFTBoostController.Contract.BOOSTCBPS(&_NFTBoostController.CallOpts)
}

// BOOSTCBPS is a free data retrieval call binding the contract method 0x8ea1f274.
//
// Solidity: function BOOST_C_BPS() view returns(uint256)
func (_NFTBoostController *NFTBoostControllerCallerSession) BOOSTCBPS() (*big.Int, error) {
	return _NFTBoostController.Contract.BOOSTCBPS(&_NFTBoostController.CallOpts)
}

// BOOSTNORMALBPS is a free data retrieval call binding the contract method 0xe22c0364.
//
// Solidity: function BOOST_NORMAL_BPS() view returns(uint256)
func (_NFTBoostController *NFTBoostControllerCaller) BOOSTNORMALBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTBoostController.contract.Call(opts, &out, "BOOST_NORMAL_BPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BOOSTNORMALBPS is a free data retrieval call binding the contract method 0xe22c0364.
//
// Solidity: function BOOST_NORMAL_BPS() view returns(uint256)
func (_NFTBoostController *NFTBoostControllerSession) BOOSTNORMALBPS() (*big.Int, error) {
	return _NFTBoostController.Contract.BOOSTNORMALBPS(&_NFTBoostController.CallOpts)
}

// BOOSTNORMALBPS is a free data retrieval call binding the contract method 0xe22c0364.
//
// Solidity: function BOOST_NORMAL_BPS() view returns(uint256)
func (_NFTBoostController *NFTBoostControllerCallerSession) BOOSTNORMALBPS() (*big.Int, error) {
	return _NFTBoostController.Contract.BOOSTNORMALBPS(&_NFTBoostController.CallOpts)
}

// BOOSTSSSBPS is a free data retrieval call binding the contract method 0xa4e43fb1.
//
// Solidity: function BOOST_SSS_BPS() view returns(uint256)
func (_NFTBoostController *NFTBoostControllerCaller) BOOSTSSSBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTBoostController.contract.Call(opts, &out, "BOOST_SSS_BPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BOOSTSSSBPS is a free data retrieval call binding the contract method 0xa4e43fb1.
//
// Solidity: function BOOST_SSS_BPS() view returns(uint256)
func (_NFTBoostController *NFTBoostControllerSession) BOOSTSSSBPS() (*big.Int, error) {
	return _NFTBoostController.Contract.BOOSTSSSBPS(&_NFTBoostController.CallOpts)
}

// BOOSTSSSBPS is a free data retrieval call binding the contract method 0xa4e43fb1.
//
// Solidity: function BOOST_SSS_BPS() view returns(uint256)
func (_NFTBoostController *NFTBoostControllerCallerSession) BOOSTSSSBPS() (*big.Int, error) {
	return _NFTBoostController.Contract.BOOSTSSSBPS(&_NFTBoostController.CallOpts)
}

// BOOSTSSBPS is a free data retrieval call binding the contract method 0x884df2e3.
//
// Solidity: function BOOST_SS_BPS() view returns(uint256)
func (_NFTBoostController *NFTBoostControllerCaller) BOOSTSSBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTBoostController.contract.Call(opts, &out, "BOOST_SS_BPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BOOSTSSBPS is a free data retrieval call binding the contract method 0x884df2e3.
//
// Solidity: function BOOST_SS_BPS() view returns(uint256)
func (_NFTBoostController *NFTBoostControllerSession) BOOSTSSBPS() (*big.Int, error) {
	return _NFTBoostController.Contract.BOOSTSSBPS(&_NFTBoostController.CallOpts)
}

// BOOSTSSBPS is a free data retrieval call binding the contract method 0x884df2e3.
//
// Solidity: function BOOST_SS_BPS() view returns(uint256)
func (_NFTBoostController *NFTBoostControllerCallerSession) BOOSTSSBPS() (*big.Int, error) {
	return _NFTBoostController.Contract.BOOSTSSBPS(&_NFTBoostController.CallOpts)
}

// BOOSTSBPS is a free data retrieval call binding the contract method 0xbc675615.
//
// Solidity: function BOOST_S_BPS() view returns(uint256)
func (_NFTBoostController *NFTBoostControllerCaller) BOOSTSBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTBoostController.contract.Call(opts, &out, "BOOST_S_BPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BOOSTSBPS is a free data retrieval call binding the contract method 0xbc675615.
//
// Solidity: function BOOST_S_BPS() view returns(uint256)
func (_NFTBoostController *NFTBoostControllerSession) BOOSTSBPS() (*big.Int, error) {
	return _NFTBoostController.Contract.BOOSTSBPS(&_NFTBoostController.CallOpts)
}

// BOOSTSBPS is a free data retrieval call binding the contract method 0xbc675615.
//
// Solidity: function BOOST_S_BPS() view returns(uint256)
func (_NFTBoostController *NFTBoostControllerCallerSession) BOOSTSBPS() (*big.Int, error) {
	return _NFTBoostController.Contract.BOOSTSBPS(&_NFTBoostController.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_NFTBoostController *NFTBoostControllerCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NFTBoostController.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_NFTBoostController *NFTBoostControllerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _NFTBoostController.Contract.UPGRADEINTERFACEVERSION(&_NFTBoostController.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_NFTBoostController *NFTBoostControllerCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _NFTBoostController.Contract.UPGRADEINTERFACEVERSION(&_NFTBoostController.CallOpts)
}

// ActiveTokenId is a free data retrieval call binding the contract method 0x79f6d337.
//
// Solidity: function activeTokenId(address ) view returns(uint256)
func (_NFTBoostController *NFTBoostControllerCaller) ActiveTokenId(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NFTBoostController.contract.Call(opts, &out, "activeTokenId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveTokenId is a free data retrieval call binding the contract method 0x79f6d337.
//
// Solidity: function activeTokenId(address ) view returns(uint256)
func (_NFTBoostController *NFTBoostControllerSession) ActiveTokenId(arg0 common.Address) (*big.Int, error) {
	return _NFTBoostController.Contract.ActiveTokenId(&_NFTBoostController.CallOpts, arg0)
}

// ActiveTokenId is a free data retrieval call binding the contract method 0x79f6d337.
//
// Solidity: function activeTokenId(address ) view returns(uint256)
func (_NFTBoostController *NFTBoostControllerCallerSession) ActiveTokenId(arg0 common.Address) (*big.Int, error) {
	return _NFTBoostController.Contract.ActiveTokenId(&_NFTBoostController.CallOpts, arg0)
}

// Cpnft is a free data retrieval call binding the contract method 0x0c7b84bd.
//
// Solidity: function cpnft() view returns(address)
func (_NFTBoostController *NFTBoostControllerCaller) Cpnft(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NFTBoostController.contract.Call(opts, &out, "cpnft")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Cpnft is a free data retrieval call binding the contract method 0x0c7b84bd.
//
// Solidity: function cpnft() view returns(address)
func (_NFTBoostController *NFTBoostControllerSession) Cpnft() (common.Address, error) {
	return _NFTBoostController.Contract.Cpnft(&_NFTBoostController.CallOpts)
}

// Cpnft is a free data retrieval call binding the contract method 0x0c7b84bd.
//
// Solidity: function cpnft() view returns(address)
func (_NFTBoostController *NFTBoostControllerCallerSession) Cpnft() (common.Address, error) {
	return _NFTBoostController.Contract.Cpnft(&_NFTBoostController.CallOpts)
}

// GetActiveNFT is a free data retrieval call binding the contract method 0x20ba4541.
//
// Solidity: function getActiveNFT(address user) view returns(uint256 tokenId, uint256 level, uint256 boostBps)
func (_NFTBoostController *NFTBoostControllerCaller) GetActiveNFT(opts *bind.CallOpts, user common.Address) (struct {
	TokenId  *big.Int
	Level    *big.Int
	BoostBps *big.Int
}, error) {
	var out []interface{}
	err := _NFTBoostController.contract.Call(opts, &out, "getActiveNFT", user)

	outstruct := new(struct {
		TokenId  *big.Int
		Level    *big.Int
		BoostBps *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Level = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BoostBps = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetActiveNFT is a free data retrieval call binding the contract method 0x20ba4541.
//
// Solidity: function getActiveNFT(address user) view returns(uint256 tokenId, uint256 level, uint256 boostBps)
func (_NFTBoostController *NFTBoostControllerSession) GetActiveNFT(user common.Address) (struct {
	TokenId  *big.Int
	Level    *big.Int
	BoostBps *big.Int
}, error) {
	return _NFTBoostController.Contract.GetActiveNFT(&_NFTBoostController.CallOpts, user)
}

// GetActiveNFT is a free data retrieval call binding the contract method 0x20ba4541.
//
// Solidity: function getActiveNFT(address user) view returns(uint256 tokenId, uint256 level, uint256 boostBps)
func (_NFTBoostController *NFTBoostControllerCallerSession) GetActiveNFT(user common.Address) (struct {
	TokenId  *big.Int
	Level    *big.Int
	BoostBps *big.Int
}, error) {
	return _NFTBoostController.Contract.GetActiveNFT(&_NFTBoostController.CallOpts, user)
}

// GetLevelBoostBps is a free data retrieval call binding the contract method 0x085302bc.
//
// Solidity: function getLevelBoostBps(uint8 level) view returns(uint256)
func (_NFTBoostController *NFTBoostControllerCaller) GetLevelBoostBps(opts *bind.CallOpts, level uint8) (*big.Int, error) {
	var out []interface{}
	err := _NFTBoostController.contract.Call(opts, &out, "getLevelBoostBps", level)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLevelBoostBps is a free data retrieval call binding the contract method 0x085302bc.
//
// Solidity: function getLevelBoostBps(uint8 level) view returns(uint256)
func (_NFTBoostController *NFTBoostControllerSession) GetLevelBoostBps(level uint8) (*big.Int, error) {
	return _NFTBoostController.Contract.GetLevelBoostBps(&_NFTBoostController.CallOpts, level)
}

// GetLevelBoostBps is a free data retrieval call binding the contract method 0x085302bc.
//
// Solidity: function getLevelBoostBps(uint8 level) view returns(uint256)
func (_NFTBoostController *NFTBoostControllerCallerSession) GetLevelBoostBps(level uint8) (*big.Int, error) {
	return _NFTBoostController.Contract.GetLevelBoostBps(&_NFTBoostController.CallOpts, level)
}

// GetNFTBoostBps is a free data retrieval call binding the contract method 0x54ce4619.
//
// Solidity: function getNFTBoostBps(address user) view returns(uint256)
func (_NFTBoostController *NFTBoostControllerCaller) GetNFTBoostBps(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NFTBoostController.contract.Call(opts, &out, "getNFTBoostBps", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNFTBoostBps is a free data retrieval call binding the contract method 0x54ce4619.
//
// Solidity: function getNFTBoostBps(address user) view returns(uint256)
func (_NFTBoostController *NFTBoostControllerSession) GetNFTBoostBps(user common.Address) (*big.Int, error) {
	return _NFTBoostController.Contract.GetNFTBoostBps(&_NFTBoostController.CallOpts, user)
}

// GetNFTBoostBps is a free data retrieval call binding the contract method 0x54ce4619.
//
// Solidity: function getNFTBoostBps(address user) view returns(uint256)
func (_NFTBoostController *NFTBoostControllerCallerSession) GetNFTBoostBps(user common.Address) (*big.Int, error) {
	return _NFTBoostController.Contract.GetNFTBoostBps(&_NFTBoostController.CallOpts, user)
}

// GetNFTLevel is a free data retrieval call binding the contract method 0x080e035d.
//
// Solidity: function getNFTLevel(uint256 tokenId) view returns(uint256)
func (_NFTBoostController *NFTBoostControllerCaller) GetNFTLevel(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NFTBoostController.contract.Call(opts, &out, "getNFTLevel", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNFTLevel is a free data retrieval call binding the contract method 0x080e035d.
//
// Solidity: function getNFTLevel(uint256 tokenId) view returns(uint256)
func (_NFTBoostController *NFTBoostControllerSession) GetNFTLevel(tokenId *big.Int) (*big.Int, error) {
	return _NFTBoostController.Contract.GetNFTLevel(&_NFTBoostController.CallOpts, tokenId)
}

// GetNFTLevel is a free data retrieval call binding the contract method 0x080e035d.
//
// Solidity: function getNFTLevel(uint256 tokenId) view returns(uint256)
func (_NFTBoostController *NFTBoostControllerCallerSession) GetNFTLevel(tokenId *big.Int) (*big.Int, error) {
	return _NFTBoostController.Contract.GetNFTLevel(&_NFTBoostController.CallOpts, tokenId)
}

// HasActiveToken is a free data retrieval call binding the contract method 0x16145a54.
//
// Solidity: function hasActiveToken(address ) view returns(bool)
func (_NFTBoostController *NFTBoostControllerCaller) HasActiveToken(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _NFTBoostController.contract.Call(opts, &out, "hasActiveToken", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasActiveToken is a free data retrieval call binding the contract method 0x16145a54.
//
// Solidity: function hasActiveToken(address ) view returns(bool)
func (_NFTBoostController *NFTBoostControllerSession) HasActiveToken(arg0 common.Address) (bool, error) {
	return _NFTBoostController.Contract.HasActiveToken(&_NFTBoostController.CallOpts, arg0)
}

// HasActiveToken is a free data retrieval call binding the contract method 0x16145a54.
//
// Solidity: function hasActiveToken(address ) view returns(bool)
func (_NFTBoostController *NFTBoostControllerCallerSession) HasActiveToken(arg0 common.Address) (bool, error) {
	return _NFTBoostController.Contract.HasActiveToken(&_NFTBoostController.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NFTBoostController *NFTBoostControllerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NFTBoostController.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NFTBoostController *NFTBoostControllerSession) Owner() (common.Address, error) {
	return _NFTBoostController.Contract.Owner(&_NFTBoostController.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NFTBoostController *NFTBoostControllerCallerSession) Owner() (common.Address, error) {
	return _NFTBoostController.Contract.Owner(&_NFTBoostController.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NFTBoostController *NFTBoostControllerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NFTBoostController.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NFTBoostController *NFTBoostControllerSession) ProxiableUUID() ([32]byte, error) {
	return _NFTBoostController.Contract.ProxiableUUID(&_NFTBoostController.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NFTBoostController *NFTBoostControllerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _NFTBoostController.Contract.ProxiableUUID(&_NFTBoostController.CallOpts)
}

// ActivateBoost is a paid mutator transaction binding the contract method 0xd1e3e42a.
//
// Solidity: function activateBoost(uint256 tokenId) returns()
func (_NFTBoostController *NFTBoostControllerTransactor) ActivateBoost(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTBoostController.contract.Transact(opts, "activateBoost", tokenId)
}

// ActivateBoost is a paid mutator transaction binding the contract method 0xd1e3e42a.
//
// Solidity: function activateBoost(uint256 tokenId) returns()
func (_NFTBoostController *NFTBoostControllerSession) ActivateBoost(tokenId *big.Int) (*types.Transaction, error) {
	return _NFTBoostController.Contract.ActivateBoost(&_NFTBoostController.TransactOpts, tokenId)
}

// ActivateBoost is a paid mutator transaction binding the contract method 0xd1e3e42a.
//
// Solidity: function activateBoost(uint256 tokenId) returns()
func (_NFTBoostController *NFTBoostControllerTransactorSession) ActivateBoost(tokenId *big.Int) (*types.Transaction, error) {
	return _NFTBoostController.Contract.ActivateBoost(&_NFTBoostController.TransactOpts, tokenId)
}

// DeactivateBoost is a paid mutator transaction binding the contract method 0x2bfd3831.
//
// Solidity: function deactivateBoost() returns()
func (_NFTBoostController *NFTBoostControllerTransactor) DeactivateBoost(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTBoostController.contract.Transact(opts, "deactivateBoost")
}

// DeactivateBoost is a paid mutator transaction binding the contract method 0x2bfd3831.
//
// Solidity: function deactivateBoost() returns()
func (_NFTBoostController *NFTBoostControllerSession) DeactivateBoost() (*types.Transaction, error) {
	return _NFTBoostController.Contract.DeactivateBoost(&_NFTBoostController.TransactOpts)
}

// DeactivateBoost is a paid mutator transaction binding the contract method 0x2bfd3831.
//
// Solidity: function deactivateBoost() returns()
func (_NFTBoostController *NFTBoostControllerTransactorSession) DeactivateBoost() (*types.Transaction, error) {
	return _NFTBoostController.Contract.DeactivateBoost(&_NFTBoostController.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _cpnft, address _owner) returns()
func (_NFTBoostController *NFTBoostControllerTransactor) Initialize(opts *bind.TransactOpts, _cpnft common.Address, _owner common.Address) (*types.Transaction, error) {
	return _NFTBoostController.contract.Transact(opts, "initialize", _cpnft, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _cpnft, address _owner) returns()
func (_NFTBoostController *NFTBoostControllerSession) Initialize(_cpnft common.Address, _owner common.Address) (*types.Transaction, error) {
	return _NFTBoostController.Contract.Initialize(&_NFTBoostController.TransactOpts, _cpnft, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _cpnft, address _owner) returns()
func (_NFTBoostController *NFTBoostControllerTransactorSession) Initialize(_cpnft common.Address, _owner common.Address) (*types.Transaction, error) {
	return _NFTBoostController.Contract.Initialize(&_NFTBoostController.TransactOpts, _cpnft, _owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NFTBoostController *NFTBoostControllerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTBoostController.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NFTBoostController *NFTBoostControllerSession) RenounceOwnership() (*types.Transaction, error) {
	return _NFTBoostController.Contract.RenounceOwnership(&_NFTBoostController.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NFTBoostController *NFTBoostControllerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NFTBoostController.Contract.RenounceOwnership(&_NFTBoostController.TransactOpts)
}

// SetCPNFT is a paid mutator transaction binding the contract method 0xaf6df5a4.
//
// Solidity: function setCPNFT(address _cpnft) returns()
func (_NFTBoostController *NFTBoostControllerTransactor) SetCPNFT(opts *bind.TransactOpts, _cpnft common.Address) (*types.Transaction, error) {
	return _NFTBoostController.contract.Transact(opts, "setCPNFT", _cpnft)
}

// SetCPNFT is a paid mutator transaction binding the contract method 0xaf6df5a4.
//
// Solidity: function setCPNFT(address _cpnft) returns()
func (_NFTBoostController *NFTBoostControllerSession) SetCPNFT(_cpnft common.Address) (*types.Transaction, error) {
	return _NFTBoostController.Contract.SetCPNFT(&_NFTBoostController.TransactOpts, _cpnft)
}

// SetCPNFT is a paid mutator transaction binding the contract method 0xaf6df5a4.
//
// Solidity: function setCPNFT(address _cpnft) returns()
func (_NFTBoostController *NFTBoostControllerTransactorSession) SetCPNFT(_cpnft common.Address) (*types.Transaction, error) {
	return _NFTBoostController.Contract.SetCPNFT(&_NFTBoostController.TransactOpts, _cpnft)
}

// SetLevelBoostBps is a paid mutator transaction binding the contract method 0x96e9f00e.
//
// Solidity: function setLevelBoostBps(uint8 level, uint256 boostBps) returns()
func (_NFTBoostController *NFTBoostControllerTransactor) SetLevelBoostBps(opts *bind.TransactOpts, level uint8, boostBps *big.Int) (*types.Transaction, error) {
	return _NFTBoostController.contract.Transact(opts, "setLevelBoostBps", level, boostBps)
}

// SetLevelBoostBps is a paid mutator transaction binding the contract method 0x96e9f00e.
//
// Solidity: function setLevelBoostBps(uint8 level, uint256 boostBps) returns()
func (_NFTBoostController *NFTBoostControllerSession) SetLevelBoostBps(level uint8, boostBps *big.Int) (*types.Transaction, error) {
	return _NFTBoostController.Contract.SetLevelBoostBps(&_NFTBoostController.TransactOpts, level, boostBps)
}

// SetLevelBoostBps is a paid mutator transaction binding the contract method 0x96e9f00e.
//
// Solidity: function setLevelBoostBps(uint8 level, uint256 boostBps) returns()
func (_NFTBoostController *NFTBoostControllerTransactorSession) SetLevelBoostBps(level uint8, boostBps *big.Int) (*types.Transaction, error) {
	return _NFTBoostController.Contract.SetLevelBoostBps(&_NFTBoostController.TransactOpts, level, boostBps)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NFTBoostController *NFTBoostControllerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NFTBoostController.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NFTBoostController *NFTBoostControllerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NFTBoostController.Contract.TransferOwnership(&_NFTBoostController.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NFTBoostController *NFTBoostControllerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NFTBoostController.Contract.TransferOwnership(&_NFTBoostController.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NFTBoostController *NFTBoostControllerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NFTBoostController.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NFTBoostController *NFTBoostControllerSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NFTBoostController.Contract.UpgradeToAndCall(&_NFTBoostController.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NFTBoostController *NFTBoostControllerTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NFTBoostController.Contract.UpgradeToAndCall(&_NFTBoostController.TransactOpts, newImplementation, data)
}

// NFTBoostControllerBoostActivatedIterator is returned from FilterBoostActivated and is used to iterate over the raw logs and unpacked data for BoostActivated events raised by the NFTBoostController contract.
type NFTBoostControllerBoostActivatedIterator struct {
	Event *NFTBoostControllerBoostActivated // Event containing the contract specifics and raw log

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
func (it *NFTBoostControllerBoostActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTBoostControllerBoostActivated)
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
		it.Event = new(NFTBoostControllerBoostActivated)
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
func (it *NFTBoostControllerBoostActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTBoostControllerBoostActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTBoostControllerBoostActivated represents a BoostActivated event raised by the NFTBoostController contract.
type NFTBoostControllerBoostActivated struct {
	User     common.Address
	TokenId  *big.Int
	Level    uint8
	BoostBps *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBoostActivated is a free log retrieval operation binding the contract event 0xd20e8b20a7a1fd48b484ffd46695f6fb30397770b5afe0b8a98f0b52924839d1.
//
// Solidity: event BoostActivated(address indexed user, uint256 indexed tokenId, uint8 level, uint256 boostBps)
func (_NFTBoostController *NFTBoostControllerFilterer) FilterBoostActivated(opts *bind.FilterOpts, user []common.Address, tokenId []*big.Int) (*NFTBoostControllerBoostActivatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NFTBoostController.contract.FilterLogs(opts, "BoostActivated", userRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NFTBoostControllerBoostActivatedIterator{contract: _NFTBoostController.contract, event: "BoostActivated", logs: logs, sub: sub}, nil
}

// WatchBoostActivated is a free log subscription operation binding the contract event 0xd20e8b20a7a1fd48b484ffd46695f6fb30397770b5afe0b8a98f0b52924839d1.
//
// Solidity: event BoostActivated(address indexed user, uint256 indexed tokenId, uint8 level, uint256 boostBps)
func (_NFTBoostController *NFTBoostControllerFilterer) WatchBoostActivated(opts *bind.WatchOpts, sink chan<- *NFTBoostControllerBoostActivated, user []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NFTBoostController.contract.WatchLogs(opts, "BoostActivated", userRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTBoostControllerBoostActivated)
				if err := _NFTBoostController.contract.UnpackLog(event, "BoostActivated", log); err != nil {
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

// ParseBoostActivated is a log parse operation binding the contract event 0xd20e8b20a7a1fd48b484ffd46695f6fb30397770b5afe0b8a98f0b52924839d1.
//
// Solidity: event BoostActivated(address indexed user, uint256 indexed tokenId, uint8 level, uint256 boostBps)
func (_NFTBoostController *NFTBoostControllerFilterer) ParseBoostActivated(log types.Log) (*NFTBoostControllerBoostActivated, error) {
	event := new(NFTBoostControllerBoostActivated)
	if err := _NFTBoostController.contract.UnpackLog(event, "BoostActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTBoostControllerBoostDeactivatedIterator is returned from FilterBoostDeactivated and is used to iterate over the raw logs and unpacked data for BoostDeactivated events raised by the NFTBoostController contract.
type NFTBoostControllerBoostDeactivatedIterator struct {
	Event *NFTBoostControllerBoostDeactivated // Event containing the contract specifics and raw log

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
func (it *NFTBoostControllerBoostDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTBoostControllerBoostDeactivated)
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
		it.Event = new(NFTBoostControllerBoostDeactivated)
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
func (it *NFTBoostControllerBoostDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTBoostControllerBoostDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTBoostControllerBoostDeactivated represents a BoostDeactivated event raised by the NFTBoostController contract.
type NFTBoostControllerBoostDeactivated struct {
	User    common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBoostDeactivated is a free log retrieval operation binding the contract event 0x406624fe448b0b19d56ac94b376568fa0c648e31b35f3d3e2bc86deafb73c8c3.
//
// Solidity: event BoostDeactivated(address indexed user, uint256 indexed tokenId)
func (_NFTBoostController *NFTBoostControllerFilterer) FilterBoostDeactivated(opts *bind.FilterOpts, user []common.Address, tokenId []*big.Int) (*NFTBoostControllerBoostDeactivatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NFTBoostController.contract.FilterLogs(opts, "BoostDeactivated", userRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NFTBoostControllerBoostDeactivatedIterator{contract: _NFTBoostController.contract, event: "BoostDeactivated", logs: logs, sub: sub}, nil
}

// WatchBoostDeactivated is a free log subscription operation binding the contract event 0x406624fe448b0b19d56ac94b376568fa0c648e31b35f3d3e2bc86deafb73c8c3.
//
// Solidity: event BoostDeactivated(address indexed user, uint256 indexed tokenId)
func (_NFTBoostController *NFTBoostControllerFilterer) WatchBoostDeactivated(opts *bind.WatchOpts, sink chan<- *NFTBoostControllerBoostDeactivated, user []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NFTBoostController.contract.WatchLogs(opts, "BoostDeactivated", userRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTBoostControllerBoostDeactivated)
				if err := _NFTBoostController.contract.UnpackLog(event, "BoostDeactivated", log); err != nil {
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

// ParseBoostDeactivated is a log parse operation binding the contract event 0x406624fe448b0b19d56ac94b376568fa0c648e31b35f3d3e2bc86deafb73c8c3.
//
// Solidity: event BoostDeactivated(address indexed user, uint256 indexed tokenId)
func (_NFTBoostController *NFTBoostControllerFilterer) ParseBoostDeactivated(log types.Log) (*NFTBoostControllerBoostDeactivated, error) {
	event := new(NFTBoostControllerBoostDeactivated)
	if err := _NFTBoostController.contract.UnpackLog(event, "BoostDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTBoostControllerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NFTBoostController contract.
type NFTBoostControllerInitializedIterator struct {
	Event *NFTBoostControllerInitialized // Event containing the contract specifics and raw log

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
func (it *NFTBoostControllerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTBoostControllerInitialized)
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
		it.Event = new(NFTBoostControllerInitialized)
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
func (it *NFTBoostControllerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTBoostControllerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTBoostControllerInitialized represents a Initialized event raised by the NFTBoostController contract.
type NFTBoostControllerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NFTBoostController *NFTBoostControllerFilterer) FilterInitialized(opts *bind.FilterOpts) (*NFTBoostControllerInitializedIterator, error) {

	logs, sub, err := _NFTBoostController.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NFTBoostControllerInitializedIterator{contract: _NFTBoostController.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NFTBoostController *NFTBoostControllerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NFTBoostControllerInitialized) (event.Subscription, error) {

	logs, sub, err := _NFTBoostController.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTBoostControllerInitialized)
				if err := _NFTBoostController.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_NFTBoostController *NFTBoostControllerFilterer) ParseInitialized(log types.Log) (*NFTBoostControllerInitialized, error) {
	event := new(NFTBoostControllerInitialized)
	if err := _NFTBoostController.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTBoostControllerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NFTBoostController contract.
type NFTBoostControllerOwnershipTransferredIterator struct {
	Event *NFTBoostControllerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NFTBoostControllerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTBoostControllerOwnershipTransferred)
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
		it.Event = new(NFTBoostControllerOwnershipTransferred)
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
func (it *NFTBoostControllerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTBoostControllerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTBoostControllerOwnershipTransferred represents a OwnershipTransferred event raised by the NFTBoostController contract.
type NFTBoostControllerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NFTBoostController *NFTBoostControllerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NFTBoostControllerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NFTBoostController.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NFTBoostControllerOwnershipTransferredIterator{contract: _NFTBoostController.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NFTBoostController *NFTBoostControllerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NFTBoostControllerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NFTBoostController.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTBoostControllerOwnershipTransferred)
				if err := _NFTBoostController.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NFTBoostController *NFTBoostControllerFilterer) ParseOwnershipTransferred(log types.Log) (*NFTBoostControllerOwnershipTransferred, error) {
	event := new(NFTBoostControllerOwnershipTransferred)
	if err := _NFTBoostController.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTBoostControllerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the NFTBoostController contract.
type NFTBoostControllerUpgradedIterator struct {
	Event *NFTBoostControllerUpgraded // Event containing the contract specifics and raw log

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
func (it *NFTBoostControllerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTBoostControllerUpgraded)
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
		it.Event = new(NFTBoostControllerUpgraded)
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
func (it *NFTBoostControllerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTBoostControllerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTBoostControllerUpgraded represents a Upgraded event raised by the NFTBoostController contract.
type NFTBoostControllerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NFTBoostController *NFTBoostControllerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*NFTBoostControllerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _NFTBoostController.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &NFTBoostControllerUpgradedIterator{contract: _NFTBoostController.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NFTBoostController *NFTBoostControllerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *NFTBoostControllerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _NFTBoostController.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTBoostControllerUpgraded)
				if err := _NFTBoostController.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_NFTBoostController *NFTBoostControllerFilterer) ParseUpgraded(log types.Log) (*NFTBoostControllerUpgraded, error) {
	event := new(NFTBoostControllerUpgraded)
	if err := _NFTBoostController.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
