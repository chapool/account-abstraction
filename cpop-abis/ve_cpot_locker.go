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

// VeCPOTLockerLockPosition is an auto generated low-level Go binding around an user-defined struct.
type VeCPOTLockerLockPosition struct {
	LockId       *big.Int
	Amount       *big.Int
	VeAmount     *big.Int
	StartTime    *big.Int
	UnlockTime   *big.Int
	DurationDays *big.Int
	Active       bool
}

// VeCPOTLockerMetaData contains all meta data concerning the VeCPOTLocker contract.
var VeCPOTLockerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"boostPerVeUnit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxVecpotBoostBps\",\"type\":\"uint256\"}],\"name\":\"BoostParamsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"veAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"durationDays\",\"type\":\"uint256\"}],\"name\":\"LockedCPOT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"additionalAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVeAmount\",\"type\":\"uint256\"}],\"name\":\"MoreCPOTLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"UnlockedCPOT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"VaultSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BOOST_VE_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DURATION_180\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DURATION_30\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DURATION_360\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DURATION_90\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"boostPerVeUnit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cpot\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getBoostBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lockId\",\"type\":\"uint256\"}],\"name\":\"getLockPosition\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lockId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"veAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"durationDays\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"internalType\":\"structVeCPOTLocker.LockPosition\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getLockPositions\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lockId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"veAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"durationDays\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"internalType\":\"structVeCPOTLocker.LockPosition[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getTotalVeCPOT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_cpot\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"durationDays\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lockId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lockId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"additionalAmount\",\"type\":\"uint256\"}],\"name\":\"lockMore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lockPositions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lockId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"veAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"durationDays\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxVecpotBoostBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nextLockId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"additionalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"durationDays\",\"type\":\"uint256\"}],\"name\":\"previewBoostBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"durationDays\",\"type\":\"uint256\"}],\"name\":\"previewVeCPOT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_boostPerVeUnit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxVecpotBoostBps\",\"type\":\"uint256\"}],\"name\":\"setBoostParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"setVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lockId\",\"type\":\"uint256\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VeCPOTLockerABI is the input ABI used to generate the binding from.
// Deprecated: Use VeCPOTLockerMetaData.ABI instead.
var VeCPOTLockerABI = VeCPOTLockerMetaData.ABI

// VeCPOTLocker is an auto generated Go binding around an Ethereum contract.
type VeCPOTLocker struct {
	VeCPOTLockerCaller     // Read-only binding to the contract
	VeCPOTLockerTransactor // Write-only binding to the contract
	VeCPOTLockerFilterer   // Log filterer for contract events
}

// VeCPOTLockerCaller is an auto generated read-only Go binding around an Ethereum contract.
type VeCPOTLockerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VeCPOTLockerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VeCPOTLockerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VeCPOTLockerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VeCPOTLockerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VeCPOTLockerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VeCPOTLockerSession struct {
	Contract     *VeCPOTLocker     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VeCPOTLockerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VeCPOTLockerCallerSession struct {
	Contract *VeCPOTLockerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// VeCPOTLockerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VeCPOTLockerTransactorSession struct {
	Contract     *VeCPOTLockerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// VeCPOTLockerRaw is an auto generated low-level Go binding around an Ethereum contract.
type VeCPOTLockerRaw struct {
	Contract *VeCPOTLocker // Generic contract binding to access the raw methods on
}

// VeCPOTLockerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VeCPOTLockerCallerRaw struct {
	Contract *VeCPOTLockerCaller // Generic read-only contract binding to access the raw methods on
}

// VeCPOTLockerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VeCPOTLockerTransactorRaw struct {
	Contract *VeCPOTLockerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVeCPOTLocker creates a new instance of VeCPOTLocker, bound to a specific deployed contract.
func NewVeCPOTLocker(address common.Address, backend bind.ContractBackend) (*VeCPOTLocker, error) {
	contract, err := bindVeCPOTLocker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VeCPOTLocker{VeCPOTLockerCaller: VeCPOTLockerCaller{contract: contract}, VeCPOTLockerTransactor: VeCPOTLockerTransactor{contract: contract}, VeCPOTLockerFilterer: VeCPOTLockerFilterer{contract: contract}}, nil
}

// NewVeCPOTLockerCaller creates a new read-only instance of VeCPOTLocker, bound to a specific deployed contract.
func NewVeCPOTLockerCaller(address common.Address, caller bind.ContractCaller) (*VeCPOTLockerCaller, error) {
	contract, err := bindVeCPOTLocker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VeCPOTLockerCaller{contract: contract}, nil
}

// NewVeCPOTLockerTransactor creates a new write-only instance of VeCPOTLocker, bound to a specific deployed contract.
func NewVeCPOTLockerTransactor(address common.Address, transactor bind.ContractTransactor) (*VeCPOTLockerTransactor, error) {
	contract, err := bindVeCPOTLocker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VeCPOTLockerTransactor{contract: contract}, nil
}

// NewVeCPOTLockerFilterer creates a new log filterer instance of VeCPOTLocker, bound to a specific deployed contract.
func NewVeCPOTLockerFilterer(address common.Address, filterer bind.ContractFilterer) (*VeCPOTLockerFilterer, error) {
	contract, err := bindVeCPOTLocker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VeCPOTLockerFilterer{contract: contract}, nil
}

// bindVeCPOTLocker binds a generic wrapper to an already deployed contract.
func bindVeCPOTLocker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VeCPOTLockerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VeCPOTLocker *VeCPOTLockerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VeCPOTLocker.Contract.VeCPOTLockerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VeCPOTLocker *VeCPOTLockerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VeCPOTLocker.Contract.VeCPOTLockerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VeCPOTLocker *VeCPOTLockerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VeCPOTLocker.Contract.VeCPOTLockerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VeCPOTLocker *VeCPOTLockerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VeCPOTLocker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VeCPOTLocker *VeCPOTLockerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VeCPOTLocker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VeCPOTLocker *VeCPOTLockerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VeCPOTLocker.Contract.contract.Transact(opts, method, params...)
}

// BOOSTVEPRECISION is a free data retrieval call binding the contract method 0x7c701b53.
//
// Solidity: function BOOST_VE_PRECISION() view returns(uint256)
func (_VeCPOTLocker *VeCPOTLockerCaller) BOOSTVEPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VeCPOTLocker.contract.Call(opts, &out, "BOOST_VE_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BOOSTVEPRECISION is a free data retrieval call binding the contract method 0x7c701b53.
//
// Solidity: function BOOST_VE_PRECISION() view returns(uint256)
func (_VeCPOTLocker *VeCPOTLockerSession) BOOSTVEPRECISION() (*big.Int, error) {
	return _VeCPOTLocker.Contract.BOOSTVEPRECISION(&_VeCPOTLocker.CallOpts)
}

// BOOSTVEPRECISION is a free data retrieval call binding the contract method 0x7c701b53.
//
// Solidity: function BOOST_VE_PRECISION() view returns(uint256)
func (_VeCPOTLocker *VeCPOTLockerCallerSession) BOOSTVEPRECISION() (*big.Int, error) {
	return _VeCPOTLocker.Contract.BOOSTVEPRECISION(&_VeCPOTLocker.CallOpts)
}

// DURATION180 is a free data retrieval call binding the contract method 0x25c0fb67.
//
// Solidity: function DURATION_180() view returns(uint256)
func (_VeCPOTLocker *VeCPOTLockerCaller) DURATION180(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VeCPOTLocker.contract.Call(opts, &out, "DURATION_180")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DURATION180 is a free data retrieval call binding the contract method 0x25c0fb67.
//
// Solidity: function DURATION_180() view returns(uint256)
func (_VeCPOTLocker *VeCPOTLockerSession) DURATION180() (*big.Int, error) {
	return _VeCPOTLocker.Contract.DURATION180(&_VeCPOTLocker.CallOpts)
}

// DURATION180 is a free data retrieval call binding the contract method 0x25c0fb67.
//
// Solidity: function DURATION_180() view returns(uint256)
func (_VeCPOTLocker *VeCPOTLockerCallerSession) DURATION180() (*big.Int, error) {
	return _VeCPOTLocker.Contract.DURATION180(&_VeCPOTLocker.CallOpts)
}

// DURATION30 is a free data retrieval call binding the contract method 0x5cbfb795.
//
// Solidity: function DURATION_30() view returns(uint256)
func (_VeCPOTLocker *VeCPOTLockerCaller) DURATION30(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VeCPOTLocker.contract.Call(opts, &out, "DURATION_30")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DURATION30 is a free data retrieval call binding the contract method 0x5cbfb795.
//
// Solidity: function DURATION_30() view returns(uint256)
func (_VeCPOTLocker *VeCPOTLockerSession) DURATION30() (*big.Int, error) {
	return _VeCPOTLocker.Contract.DURATION30(&_VeCPOTLocker.CallOpts)
}

// DURATION30 is a free data retrieval call binding the contract method 0x5cbfb795.
//
// Solidity: function DURATION_30() view returns(uint256)
func (_VeCPOTLocker *VeCPOTLockerCallerSession) DURATION30() (*big.Int, error) {
	return _VeCPOTLocker.Contract.DURATION30(&_VeCPOTLocker.CallOpts)
}

// DURATION360 is a free data retrieval call binding the contract method 0xa6ce00b1.
//
// Solidity: function DURATION_360() view returns(uint256)
func (_VeCPOTLocker *VeCPOTLockerCaller) DURATION360(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VeCPOTLocker.contract.Call(opts, &out, "DURATION_360")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DURATION360 is a free data retrieval call binding the contract method 0xa6ce00b1.
//
// Solidity: function DURATION_360() view returns(uint256)
func (_VeCPOTLocker *VeCPOTLockerSession) DURATION360() (*big.Int, error) {
	return _VeCPOTLocker.Contract.DURATION360(&_VeCPOTLocker.CallOpts)
}

// DURATION360 is a free data retrieval call binding the contract method 0xa6ce00b1.
//
// Solidity: function DURATION_360() view returns(uint256)
func (_VeCPOTLocker *VeCPOTLockerCallerSession) DURATION360() (*big.Int, error) {
	return _VeCPOTLocker.Contract.DURATION360(&_VeCPOTLocker.CallOpts)
}

// DURATION90 is a free data retrieval call binding the contract method 0x729d3599.
//
// Solidity: function DURATION_90() view returns(uint256)
func (_VeCPOTLocker *VeCPOTLockerCaller) DURATION90(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VeCPOTLocker.contract.Call(opts, &out, "DURATION_90")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DURATION90 is a free data retrieval call binding the contract method 0x729d3599.
//
// Solidity: function DURATION_90() view returns(uint256)
func (_VeCPOTLocker *VeCPOTLockerSession) DURATION90() (*big.Int, error) {
	return _VeCPOTLocker.Contract.DURATION90(&_VeCPOTLocker.CallOpts)
}

// DURATION90 is a free data retrieval call binding the contract method 0x729d3599.
//
// Solidity: function DURATION_90() view returns(uint256)
func (_VeCPOTLocker *VeCPOTLockerCallerSession) DURATION90() (*big.Int, error) {
	return _VeCPOTLocker.Contract.DURATION90(&_VeCPOTLocker.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_VeCPOTLocker *VeCPOTLockerCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VeCPOTLocker.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_VeCPOTLocker *VeCPOTLockerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _VeCPOTLocker.Contract.UPGRADEINTERFACEVERSION(&_VeCPOTLocker.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_VeCPOTLocker *VeCPOTLockerCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _VeCPOTLocker.Contract.UPGRADEINTERFACEVERSION(&_VeCPOTLocker.CallOpts)
}

// BoostPerVeUnit is a free data retrieval call binding the contract method 0xd7ccd714.
//
// Solidity: function boostPerVeUnit() view returns(uint256)
func (_VeCPOTLocker *VeCPOTLockerCaller) BoostPerVeUnit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VeCPOTLocker.contract.Call(opts, &out, "boostPerVeUnit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BoostPerVeUnit is a free data retrieval call binding the contract method 0xd7ccd714.
//
// Solidity: function boostPerVeUnit() view returns(uint256)
func (_VeCPOTLocker *VeCPOTLockerSession) BoostPerVeUnit() (*big.Int, error) {
	return _VeCPOTLocker.Contract.BoostPerVeUnit(&_VeCPOTLocker.CallOpts)
}

// BoostPerVeUnit is a free data retrieval call binding the contract method 0xd7ccd714.
//
// Solidity: function boostPerVeUnit() view returns(uint256)
func (_VeCPOTLocker *VeCPOTLockerCallerSession) BoostPerVeUnit() (*big.Int, error) {
	return _VeCPOTLocker.Contract.BoostPerVeUnit(&_VeCPOTLocker.CallOpts)
}

// Cpot is a free data retrieval call binding the contract method 0x9ccf800f.
//
// Solidity: function cpot() view returns(address)
func (_VeCPOTLocker *VeCPOTLockerCaller) Cpot(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VeCPOTLocker.contract.Call(opts, &out, "cpot")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Cpot is a free data retrieval call binding the contract method 0x9ccf800f.
//
// Solidity: function cpot() view returns(address)
func (_VeCPOTLocker *VeCPOTLockerSession) Cpot() (common.Address, error) {
	return _VeCPOTLocker.Contract.Cpot(&_VeCPOTLocker.CallOpts)
}

// Cpot is a free data retrieval call binding the contract method 0x9ccf800f.
//
// Solidity: function cpot() view returns(address)
func (_VeCPOTLocker *VeCPOTLockerCallerSession) Cpot() (common.Address, error) {
	return _VeCPOTLocker.Contract.Cpot(&_VeCPOTLocker.CallOpts)
}

// GetBoostBps is a free data retrieval call binding the contract method 0x37688fb8.
//
// Solidity: function getBoostBps(address user) view returns(uint256)
func (_VeCPOTLocker *VeCPOTLockerCaller) GetBoostBps(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VeCPOTLocker.contract.Call(opts, &out, "getBoostBps", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBoostBps is a free data retrieval call binding the contract method 0x37688fb8.
//
// Solidity: function getBoostBps(address user) view returns(uint256)
func (_VeCPOTLocker *VeCPOTLockerSession) GetBoostBps(user common.Address) (*big.Int, error) {
	return _VeCPOTLocker.Contract.GetBoostBps(&_VeCPOTLocker.CallOpts, user)
}

// GetBoostBps is a free data retrieval call binding the contract method 0x37688fb8.
//
// Solidity: function getBoostBps(address user) view returns(uint256)
func (_VeCPOTLocker *VeCPOTLockerCallerSession) GetBoostBps(user common.Address) (*big.Int, error) {
	return _VeCPOTLocker.Contract.GetBoostBps(&_VeCPOTLocker.CallOpts, user)
}

// GetLockPosition is a free data retrieval call binding the contract method 0x8bc26dd2.
//
// Solidity: function getLockPosition(address user, uint256 lockId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,bool))
func (_VeCPOTLocker *VeCPOTLockerCaller) GetLockPosition(opts *bind.CallOpts, user common.Address, lockId *big.Int) (VeCPOTLockerLockPosition, error) {
	var out []interface{}
	err := _VeCPOTLocker.contract.Call(opts, &out, "getLockPosition", user, lockId)

	if err != nil {
		return *new(VeCPOTLockerLockPosition), err
	}

	out0 := *abi.ConvertType(out[0], new(VeCPOTLockerLockPosition)).(*VeCPOTLockerLockPosition)

	return out0, err

}

// GetLockPosition is a free data retrieval call binding the contract method 0x8bc26dd2.
//
// Solidity: function getLockPosition(address user, uint256 lockId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,bool))
func (_VeCPOTLocker *VeCPOTLockerSession) GetLockPosition(user common.Address, lockId *big.Int) (VeCPOTLockerLockPosition, error) {
	return _VeCPOTLocker.Contract.GetLockPosition(&_VeCPOTLocker.CallOpts, user, lockId)
}

// GetLockPosition is a free data retrieval call binding the contract method 0x8bc26dd2.
//
// Solidity: function getLockPosition(address user, uint256 lockId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,bool))
func (_VeCPOTLocker *VeCPOTLockerCallerSession) GetLockPosition(user common.Address, lockId *big.Int) (VeCPOTLockerLockPosition, error) {
	return _VeCPOTLocker.Contract.GetLockPosition(&_VeCPOTLocker.CallOpts, user, lockId)
}

// GetLockPositions is a free data retrieval call binding the contract method 0xa793f6c4.
//
// Solidity: function getLockPositions(address user) view returns((uint256,uint256,uint256,uint256,uint256,uint256,bool)[])
func (_VeCPOTLocker *VeCPOTLockerCaller) GetLockPositions(opts *bind.CallOpts, user common.Address) ([]VeCPOTLockerLockPosition, error) {
	var out []interface{}
	err := _VeCPOTLocker.contract.Call(opts, &out, "getLockPositions", user)

	if err != nil {
		return *new([]VeCPOTLockerLockPosition), err
	}

	out0 := *abi.ConvertType(out[0], new([]VeCPOTLockerLockPosition)).(*[]VeCPOTLockerLockPosition)

	return out0, err

}

// GetLockPositions is a free data retrieval call binding the contract method 0xa793f6c4.
//
// Solidity: function getLockPositions(address user) view returns((uint256,uint256,uint256,uint256,uint256,uint256,bool)[])
func (_VeCPOTLocker *VeCPOTLockerSession) GetLockPositions(user common.Address) ([]VeCPOTLockerLockPosition, error) {
	return _VeCPOTLocker.Contract.GetLockPositions(&_VeCPOTLocker.CallOpts, user)
}

// GetLockPositions is a free data retrieval call binding the contract method 0xa793f6c4.
//
// Solidity: function getLockPositions(address user) view returns((uint256,uint256,uint256,uint256,uint256,uint256,bool)[])
func (_VeCPOTLocker *VeCPOTLockerCallerSession) GetLockPositions(user common.Address) ([]VeCPOTLockerLockPosition, error) {
	return _VeCPOTLocker.Contract.GetLockPositions(&_VeCPOTLocker.CallOpts, user)
}

// GetTotalVeCPOT is a free data retrieval call binding the contract method 0xd2c6894f.
//
// Solidity: function getTotalVeCPOT(address user) view returns(uint256)
func (_VeCPOTLocker *VeCPOTLockerCaller) GetTotalVeCPOT(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VeCPOTLocker.contract.Call(opts, &out, "getTotalVeCPOT", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalVeCPOT is a free data retrieval call binding the contract method 0xd2c6894f.
//
// Solidity: function getTotalVeCPOT(address user) view returns(uint256)
func (_VeCPOTLocker *VeCPOTLockerSession) GetTotalVeCPOT(user common.Address) (*big.Int, error) {
	return _VeCPOTLocker.Contract.GetTotalVeCPOT(&_VeCPOTLocker.CallOpts, user)
}

// GetTotalVeCPOT is a free data retrieval call binding the contract method 0xd2c6894f.
//
// Solidity: function getTotalVeCPOT(address user) view returns(uint256)
func (_VeCPOTLocker *VeCPOTLockerCallerSession) GetTotalVeCPOT(user common.Address) (*big.Int, error) {
	return _VeCPOTLocker.Contract.GetTotalVeCPOT(&_VeCPOTLocker.CallOpts, user)
}

// LockPositions is a free data retrieval call binding the contract method 0x0787acf7.
//
// Solidity: function lockPositions(address , uint256 ) view returns(uint256 lockId, uint256 amount, uint256 veAmount, uint256 startTime, uint256 unlockTime, uint256 durationDays, bool active)
func (_VeCPOTLocker *VeCPOTLockerCaller) LockPositions(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	LockId       *big.Int
	Amount       *big.Int
	VeAmount     *big.Int
	StartTime    *big.Int
	UnlockTime   *big.Int
	DurationDays *big.Int
	Active       bool
}, error) {
	var out []interface{}
	err := _VeCPOTLocker.contract.Call(opts, &out, "lockPositions", arg0, arg1)

	outstruct := new(struct {
		LockId       *big.Int
		Amount       *big.Int
		VeAmount     *big.Int
		StartTime    *big.Int
		UnlockTime   *big.Int
		DurationDays *big.Int
		Active       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LockId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.VeAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.UnlockTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.DurationDays = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Active = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// LockPositions is a free data retrieval call binding the contract method 0x0787acf7.
//
// Solidity: function lockPositions(address , uint256 ) view returns(uint256 lockId, uint256 amount, uint256 veAmount, uint256 startTime, uint256 unlockTime, uint256 durationDays, bool active)
func (_VeCPOTLocker *VeCPOTLockerSession) LockPositions(arg0 common.Address, arg1 *big.Int) (struct {
	LockId       *big.Int
	Amount       *big.Int
	VeAmount     *big.Int
	StartTime    *big.Int
	UnlockTime   *big.Int
	DurationDays *big.Int
	Active       bool
}, error) {
	return _VeCPOTLocker.Contract.LockPositions(&_VeCPOTLocker.CallOpts, arg0, arg1)
}

// LockPositions is a free data retrieval call binding the contract method 0x0787acf7.
//
// Solidity: function lockPositions(address , uint256 ) view returns(uint256 lockId, uint256 amount, uint256 veAmount, uint256 startTime, uint256 unlockTime, uint256 durationDays, bool active)
func (_VeCPOTLocker *VeCPOTLockerCallerSession) LockPositions(arg0 common.Address, arg1 *big.Int) (struct {
	LockId       *big.Int
	Amount       *big.Int
	VeAmount     *big.Int
	StartTime    *big.Int
	UnlockTime   *big.Int
	DurationDays *big.Int
	Active       bool
}, error) {
	return _VeCPOTLocker.Contract.LockPositions(&_VeCPOTLocker.CallOpts, arg0, arg1)
}

// MaxVecpotBoostBps is a free data retrieval call binding the contract method 0x17dab899.
//
// Solidity: function maxVecpotBoostBps() view returns(uint256)
func (_VeCPOTLocker *VeCPOTLockerCaller) MaxVecpotBoostBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VeCPOTLocker.contract.Call(opts, &out, "maxVecpotBoostBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxVecpotBoostBps is a free data retrieval call binding the contract method 0x17dab899.
//
// Solidity: function maxVecpotBoostBps() view returns(uint256)
func (_VeCPOTLocker *VeCPOTLockerSession) MaxVecpotBoostBps() (*big.Int, error) {
	return _VeCPOTLocker.Contract.MaxVecpotBoostBps(&_VeCPOTLocker.CallOpts)
}

// MaxVecpotBoostBps is a free data retrieval call binding the contract method 0x17dab899.
//
// Solidity: function maxVecpotBoostBps() view returns(uint256)
func (_VeCPOTLocker *VeCPOTLockerCallerSession) MaxVecpotBoostBps() (*big.Int, error) {
	return _VeCPOTLocker.Contract.MaxVecpotBoostBps(&_VeCPOTLocker.CallOpts)
}

// NextLockId is a free data retrieval call binding the contract method 0x80e9c989.
//
// Solidity: function nextLockId(address ) view returns(uint256)
func (_VeCPOTLocker *VeCPOTLockerCaller) NextLockId(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VeCPOTLocker.contract.Call(opts, &out, "nextLockId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextLockId is a free data retrieval call binding the contract method 0x80e9c989.
//
// Solidity: function nextLockId(address ) view returns(uint256)
func (_VeCPOTLocker *VeCPOTLockerSession) NextLockId(arg0 common.Address) (*big.Int, error) {
	return _VeCPOTLocker.Contract.NextLockId(&_VeCPOTLocker.CallOpts, arg0)
}

// NextLockId is a free data retrieval call binding the contract method 0x80e9c989.
//
// Solidity: function nextLockId(address ) view returns(uint256)
func (_VeCPOTLocker *VeCPOTLockerCallerSession) NextLockId(arg0 common.Address) (*big.Int, error) {
	return _VeCPOTLocker.Contract.NextLockId(&_VeCPOTLocker.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VeCPOTLocker *VeCPOTLockerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VeCPOTLocker.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VeCPOTLocker *VeCPOTLockerSession) Owner() (common.Address, error) {
	return _VeCPOTLocker.Contract.Owner(&_VeCPOTLocker.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VeCPOTLocker *VeCPOTLockerCallerSession) Owner() (common.Address, error) {
	return _VeCPOTLocker.Contract.Owner(&_VeCPOTLocker.CallOpts)
}

// PreviewBoostBps is a free data retrieval call binding the contract method 0x8b5b703c.
//
// Solidity: function previewBoostBps(address user, uint256 additionalAmount, uint256 durationDays) view returns(uint256)
func (_VeCPOTLocker *VeCPOTLockerCaller) PreviewBoostBps(opts *bind.CallOpts, user common.Address, additionalAmount *big.Int, durationDays *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VeCPOTLocker.contract.Call(opts, &out, "previewBoostBps", user, additionalAmount, durationDays)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewBoostBps is a free data retrieval call binding the contract method 0x8b5b703c.
//
// Solidity: function previewBoostBps(address user, uint256 additionalAmount, uint256 durationDays) view returns(uint256)
func (_VeCPOTLocker *VeCPOTLockerSession) PreviewBoostBps(user common.Address, additionalAmount *big.Int, durationDays *big.Int) (*big.Int, error) {
	return _VeCPOTLocker.Contract.PreviewBoostBps(&_VeCPOTLocker.CallOpts, user, additionalAmount, durationDays)
}

// PreviewBoostBps is a free data retrieval call binding the contract method 0x8b5b703c.
//
// Solidity: function previewBoostBps(address user, uint256 additionalAmount, uint256 durationDays) view returns(uint256)
func (_VeCPOTLocker *VeCPOTLockerCallerSession) PreviewBoostBps(user common.Address, additionalAmount *big.Int, durationDays *big.Int) (*big.Int, error) {
	return _VeCPOTLocker.Contract.PreviewBoostBps(&_VeCPOTLocker.CallOpts, user, additionalAmount, durationDays)
}

// PreviewVeCPOT is a free data retrieval call binding the contract method 0x2fe21365.
//
// Solidity: function previewVeCPOT(uint256 amount, uint256 durationDays) pure returns(uint256)
func (_VeCPOTLocker *VeCPOTLockerCaller) PreviewVeCPOT(opts *bind.CallOpts, amount *big.Int, durationDays *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VeCPOTLocker.contract.Call(opts, &out, "previewVeCPOT", amount, durationDays)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewVeCPOT is a free data retrieval call binding the contract method 0x2fe21365.
//
// Solidity: function previewVeCPOT(uint256 amount, uint256 durationDays) pure returns(uint256)
func (_VeCPOTLocker *VeCPOTLockerSession) PreviewVeCPOT(amount *big.Int, durationDays *big.Int) (*big.Int, error) {
	return _VeCPOTLocker.Contract.PreviewVeCPOT(&_VeCPOTLocker.CallOpts, amount, durationDays)
}

// PreviewVeCPOT is a free data retrieval call binding the contract method 0x2fe21365.
//
// Solidity: function previewVeCPOT(uint256 amount, uint256 durationDays) pure returns(uint256)
func (_VeCPOTLocker *VeCPOTLockerCallerSession) PreviewVeCPOT(amount *big.Int, durationDays *big.Int) (*big.Int, error) {
	return _VeCPOTLocker.Contract.PreviewVeCPOT(&_VeCPOTLocker.CallOpts, amount, durationDays)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_VeCPOTLocker *VeCPOTLockerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VeCPOTLocker.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_VeCPOTLocker *VeCPOTLockerSession) ProxiableUUID() ([32]byte, error) {
	return _VeCPOTLocker.Contract.ProxiableUUID(&_VeCPOTLocker.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_VeCPOTLocker *VeCPOTLockerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _VeCPOTLocker.Contract.ProxiableUUID(&_VeCPOTLocker.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_VeCPOTLocker *VeCPOTLockerCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VeCPOTLocker.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_VeCPOTLocker *VeCPOTLockerSession) Vault() (common.Address, error) {
	return _VeCPOTLocker.Contract.Vault(&_VeCPOTLocker.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_VeCPOTLocker *VeCPOTLockerCallerSession) Vault() (common.Address, error) {
	return _VeCPOTLocker.Contract.Vault(&_VeCPOTLocker.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _cpot, address _owner) returns()
func (_VeCPOTLocker *VeCPOTLockerTransactor) Initialize(opts *bind.TransactOpts, _cpot common.Address, _owner common.Address) (*types.Transaction, error) {
	return _VeCPOTLocker.contract.Transact(opts, "initialize", _cpot, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _cpot, address _owner) returns()
func (_VeCPOTLocker *VeCPOTLockerSession) Initialize(_cpot common.Address, _owner common.Address) (*types.Transaction, error) {
	return _VeCPOTLocker.Contract.Initialize(&_VeCPOTLocker.TransactOpts, _cpot, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _cpot, address _owner) returns()
func (_VeCPOTLocker *VeCPOTLockerTransactorSession) Initialize(_cpot common.Address, _owner common.Address) (*types.Transaction, error) {
	return _VeCPOTLocker.Contract.Initialize(&_VeCPOTLocker.TransactOpts, _cpot, _owner)
}

// Lock is a paid mutator transaction binding the contract method 0x1338736f.
//
// Solidity: function lock(uint256 amount, uint256 durationDays) returns(uint256 lockId)
func (_VeCPOTLocker *VeCPOTLockerTransactor) Lock(opts *bind.TransactOpts, amount *big.Int, durationDays *big.Int) (*types.Transaction, error) {
	return _VeCPOTLocker.contract.Transact(opts, "lock", amount, durationDays)
}

// Lock is a paid mutator transaction binding the contract method 0x1338736f.
//
// Solidity: function lock(uint256 amount, uint256 durationDays) returns(uint256 lockId)
func (_VeCPOTLocker *VeCPOTLockerSession) Lock(amount *big.Int, durationDays *big.Int) (*types.Transaction, error) {
	return _VeCPOTLocker.Contract.Lock(&_VeCPOTLocker.TransactOpts, amount, durationDays)
}

// Lock is a paid mutator transaction binding the contract method 0x1338736f.
//
// Solidity: function lock(uint256 amount, uint256 durationDays) returns(uint256 lockId)
func (_VeCPOTLocker *VeCPOTLockerTransactorSession) Lock(amount *big.Int, durationDays *big.Int) (*types.Transaction, error) {
	return _VeCPOTLocker.Contract.Lock(&_VeCPOTLocker.TransactOpts, amount, durationDays)
}

// LockMore is a paid mutator transaction binding the contract method 0xe6ee2217.
//
// Solidity: function lockMore(uint256 lockId, uint256 additionalAmount) returns()
func (_VeCPOTLocker *VeCPOTLockerTransactor) LockMore(opts *bind.TransactOpts, lockId *big.Int, additionalAmount *big.Int) (*types.Transaction, error) {
	return _VeCPOTLocker.contract.Transact(opts, "lockMore", lockId, additionalAmount)
}

// LockMore is a paid mutator transaction binding the contract method 0xe6ee2217.
//
// Solidity: function lockMore(uint256 lockId, uint256 additionalAmount) returns()
func (_VeCPOTLocker *VeCPOTLockerSession) LockMore(lockId *big.Int, additionalAmount *big.Int) (*types.Transaction, error) {
	return _VeCPOTLocker.Contract.LockMore(&_VeCPOTLocker.TransactOpts, lockId, additionalAmount)
}

// LockMore is a paid mutator transaction binding the contract method 0xe6ee2217.
//
// Solidity: function lockMore(uint256 lockId, uint256 additionalAmount) returns()
func (_VeCPOTLocker *VeCPOTLockerTransactorSession) LockMore(lockId *big.Int, additionalAmount *big.Int) (*types.Transaction, error) {
	return _VeCPOTLocker.Contract.LockMore(&_VeCPOTLocker.TransactOpts, lockId, additionalAmount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VeCPOTLocker *VeCPOTLockerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VeCPOTLocker.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VeCPOTLocker *VeCPOTLockerSession) RenounceOwnership() (*types.Transaction, error) {
	return _VeCPOTLocker.Contract.RenounceOwnership(&_VeCPOTLocker.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VeCPOTLocker *VeCPOTLockerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VeCPOTLocker.Contract.RenounceOwnership(&_VeCPOTLocker.TransactOpts)
}

// SetBoostParams is a paid mutator transaction binding the contract method 0x62a42442.
//
// Solidity: function setBoostParams(uint256 _boostPerVeUnit, uint256 _maxVecpotBoostBps) returns()
func (_VeCPOTLocker *VeCPOTLockerTransactor) SetBoostParams(opts *bind.TransactOpts, _boostPerVeUnit *big.Int, _maxVecpotBoostBps *big.Int) (*types.Transaction, error) {
	return _VeCPOTLocker.contract.Transact(opts, "setBoostParams", _boostPerVeUnit, _maxVecpotBoostBps)
}

// SetBoostParams is a paid mutator transaction binding the contract method 0x62a42442.
//
// Solidity: function setBoostParams(uint256 _boostPerVeUnit, uint256 _maxVecpotBoostBps) returns()
func (_VeCPOTLocker *VeCPOTLockerSession) SetBoostParams(_boostPerVeUnit *big.Int, _maxVecpotBoostBps *big.Int) (*types.Transaction, error) {
	return _VeCPOTLocker.Contract.SetBoostParams(&_VeCPOTLocker.TransactOpts, _boostPerVeUnit, _maxVecpotBoostBps)
}

// SetBoostParams is a paid mutator transaction binding the contract method 0x62a42442.
//
// Solidity: function setBoostParams(uint256 _boostPerVeUnit, uint256 _maxVecpotBoostBps) returns()
func (_VeCPOTLocker *VeCPOTLockerTransactorSession) SetBoostParams(_boostPerVeUnit *big.Int, _maxVecpotBoostBps *big.Int) (*types.Transaction, error) {
	return _VeCPOTLocker.Contract.SetBoostParams(&_VeCPOTLocker.TransactOpts, _boostPerVeUnit, _maxVecpotBoostBps)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address _vault) returns()
func (_VeCPOTLocker *VeCPOTLockerTransactor) SetVault(opts *bind.TransactOpts, _vault common.Address) (*types.Transaction, error) {
	return _VeCPOTLocker.contract.Transact(opts, "setVault", _vault)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address _vault) returns()
func (_VeCPOTLocker *VeCPOTLockerSession) SetVault(_vault common.Address) (*types.Transaction, error) {
	return _VeCPOTLocker.Contract.SetVault(&_VeCPOTLocker.TransactOpts, _vault)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address _vault) returns()
func (_VeCPOTLocker *VeCPOTLockerTransactorSession) SetVault(_vault common.Address) (*types.Transaction, error) {
	return _VeCPOTLocker.Contract.SetVault(&_VeCPOTLocker.TransactOpts, _vault)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VeCPOTLocker *VeCPOTLockerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VeCPOTLocker.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VeCPOTLocker *VeCPOTLockerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VeCPOTLocker.Contract.TransferOwnership(&_VeCPOTLocker.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VeCPOTLocker *VeCPOTLockerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VeCPOTLocker.Contract.TransferOwnership(&_VeCPOTLocker.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 lockId) returns()
func (_VeCPOTLocker *VeCPOTLockerTransactor) Unlock(opts *bind.TransactOpts, lockId *big.Int) (*types.Transaction, error) {
	return _VeCPOTLocker.contract.Transact(opts, "unlock", lockId)
}

// Unlock is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 lockId) returns()
func (_VeCPOTLocker *VeCPOTLockerSession) Unlock(lockId *big.Int) (*types.Transaction, error) {
	return _VeCPOTLocker.Contract.Unlock(&_VeCPOTLocker.TransactOpts, lockId)
}

// Unlock is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 lockId) returns()
func (_VeCPOTLocker *VeCPOTLockerTransactorSession) Unlock(lockId *big.Int) (*types.Transaction, error) {
	return _VeCPOTLocker.Contract.Unlock(&_VeCPOTLocker.TransactOpts, lockId)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_VeCPOTLocker *VeCPOTLockerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _VeCPOTLocker.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_VeCPOTLocker *VeCPOTLockerSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _VeCPOTLocker.Contract.UpgradeToAndCall(&_VeCPOTLocker.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_VeCPOTLocker *VeCPOTLockerTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _VeCPOTLocker.Contract.UpgradeToAndCall(&_VeCPOTLocker.TransactOpts, newImplementation, data)
}

// VeCPOTLockerBoostParamsSetIterator is returned from FilterBoostParamsSet and is used to iterate over the raw logs and unpacked data for BoostParamsSet events raised by the VeCPOTLocker contract.
type VeCPOTLockerBoostParamsSetIterator struct {
	Event *VeCPOTLockerBoostParamsSet // Event containing the contract specifics and raw log

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
func (it *VeCPOTLockerBoostParamsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeCPOTLockerBoostParamsSet)
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
		it.Event = new(VeCPOTLockerBoostParamsSet)
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
func (it *VeCPOTLockerBoostParamsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeCPOTLockerBoostParamsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeCPOTLockerBoostParamsSet represents a BoostParamsSet event raised by the VeCPOTLocker contract.
type VeCPOTLockerBoostParamsSet struct {
	BoostPerVeUnit    *big.Int
	MaxVecpotBoostBps *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterBoostParamsSet is a free log retrieval operation binding the contract event 0xe9163a86c9383a3712a9c137f0d584b02f37d0447328cd773c899d93dd75b840.
//
// Solidity: event BoostParamsSet(uint256 boostPerVeUnit, uint256 maxVecpotBoostBps)
func (_VeCPOTLocker *VeCPOTLockerFilterer) FilterBoostParamsSet(opts *bind.FilterOpts) (*VeCPOTLockerBoostParamsSetIterator, error) {

	logs, sub, err := _VeCPOTLocker.contract.FilterLogs(opts, "BoostParamsSet")
	if err != nil {
		return nil, err
	}
	return &VeCPOTLockerBoostParamsSetIterator{contract: _VeCPOTLocker.contract, event: "BoostParamsSet", logs: logs, sub: sub}, nil
}

// WatchBoostParamsSet is a free log subscription operation binding the contract event 0xe9163a86c9383a3712a9c137f0d584b02f37d0447328cd773c899d93dd75b840.
//
// Solidity: event BoostParamsSet(uint256 boostPerVeUnit, uint256 maxVecpotBoostBps)
func (_VeCPOTLocker *VeCPOTLockerFilterer) WatchBoostParamsSet(opts *bind.WatchOpts, sink chan<- *VeCPOTLockerBoostParamsSet) (event.Subscription, error) {

	logs, sub, err := _VeCPOTLocker.contract.WatchLogs(opts, "BoostParamsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeCPOTLockerBoostParamsSet)
				if err := _VeCPOTLocker.contract.UnpackLog(event, "BoostParamsSet", log); err != nil {
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

// ParseBoostParamsSet is a log parse operation binding the contract event 0xe9163a86c9383a3712a9c137f0d584b02f37d0447328cd773c899d93dd75b840.
//
// Solidity: event BoostParamsSet(uint256 boostPerVeUnit, uint256 maxVecpotBoostBps)
func (_VeCPOTLocker *VeCPOTLockerFilterer) ParseBoostParamsSet(log types.Log) (*VeCPOTLockerBoostParamsSet, error) {
	event := new(VeCPOTLockerBoostParamsSet)
	if err := _VeCPOTLocker.contract.UnpackLog(event, "BoostParamsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeCPOTLockerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the VeCPOTLocker contract.
type VeCPOTLockerInitializedIterator struct {
	Event *VeCPOTLockerInitialized // Event containing the contract specifics and raw log

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
func (it *VeCPOTLockerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeCPOTLockerInitialized)
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
		it.Event = new(VeCPOTLockerInitialized)
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
func (it *VeCPOTLockerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeCPOTLockerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeCPOTLockerInitialized represents a Initialized event raised by the VeCPOTLocker contract.
type VeCPOTLockerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_VeCPOTLocker *VeCPOTLockerFilterer) FilterInitialized(opts *bind.FilterOpts) (*VeCPOTLockerInitializedIterator, error) {

	logs, sub, err := _VeCPOTLocker.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &VeCPOTLockerInitializedIterator{contract: _VeCPOTLocker.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_VeCPOTLocker *VeCPOTLockerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *VeCPOTLockerInitialized) (event.Subscription, error) {

	logs, sub, err := _VeCPOTLocker.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeCPOTLockerInitialized)
				if err := _VeCPOTLocker.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_VeCPOTLocker *VeCPOTLockerFilterer) ParseInitialized(log types.Log) (*VeCPOTLockerInitialized, error) {
	event := new(VeCPOTLockerInitialized)
	if err := _VeCPOTLocker.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeCPOTLockerLockedCPOTIterator is returned from FilterLockedCPOT and is used to iterate over the raw logs and unpacked data for LockedCPOT events raised by the VeCPOTLocker contract.
type VeCPOTLockerLockedCPOTIterator struct {
	Event *VeCPOTLockerLockedCPOT // Event containing the contract specifics and raw log

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
func (it *VeCPOTLockerLockedCPOTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeCPOTLockerLockedCPOT)
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
		it.Event = new(VeCPOTLockerLockedCPOT)
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
func (it *VeCPOTLockerLockedCPOTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeCPOTLockerLockedCPOTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeCPOTLockerLockedCPOT represents a LockedCPOT event raised by the VeCPOTLocker contract.
type VeCPOTLockerLockedCPOT struct {
	User         common.Address
	LockId       *big.Int
	Amount       *big.Int
	UnlockTime   *big.Int
	VeAmount     *big.Int
	DurationDays *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLockedCPOT is a free log retrieval operation binding the contract event 0xaecaebcccc0d65d57a05ea25ceabf33184b4d012230640a8ef0d3fb127e94a91.
//
// Solidity: event LockedCPOT(address indexed user, uint256 lockId, uint256 amount, uint256 unlockTime, uint256 veAmount, uint256 durationDays)
func (_VeCPOTLocker *VeCPOTLockerFilterer) FilterLockedCPOT(opts *bind.FilterOpts, user []common.Address) (*VeCPOTLockerLockedCPOTIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _VeCPOTLocker.contract.FilterLogs(opts, "LockedCPOT", userRule)
	if err != nil {
		return nil, err
	}
	return &VeCPOTLockerLockedCPOTIterator{contract: _VeCPOTLocker.contract, event: "LockedCPOT", logs: logs, sub: sub}, nil
}

// WatchLockedCPOT is a free log subscription operation binding the contract event 0xaecaebcccc0d65d57a05ea25ceabf33184b4d012230640a8ef0d3fb127e94a91.
//
// Solidity: event LockedCPOT(address indexed user, uint256 lockId, uint256 amount, uint256 unlockTime, uint256 veAmount, uint256 durationDays)
func (_VeCPOTLocker *VeCPOTLockerFilterer) WatchLockedCPOT(opts *bind.WatchOpts, sink chan<- *VeCPOTLockerLockedCPOT, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _VeCPOTLocker.contract.WatchLogs(opts, "LockedCPOT", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeCPOTLockerLockedCPOT)
				if err := _VeCPOTLocker.contract.UnpackLog(event, "LockedCPOT", log); err != nil {
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

// ParseLockedCPOT is a log parse operation binding the contract event 0xaecaebcccc0d65d57a05ea25ceabf33184b4d012230640a8ef0d3fb127e94a91.
//
// Solidity: event LockedCPOT(address indexed user, uint256 lockId, uint256 amount, uint256 unlockTime, uint256 veAmount, uint256 durationDays)
func (_VeCPOTLocker *VeCPOTLockerFilterer) ParseLockedCPOT(log types.Log) (*VeCPOTLockerLockedCPOT, error) {
	event := new(VeCPOTLockerLockedCPOT)
	if err := _VeCPOTLocker.contract.UnpackLog(event, "LockedCPOT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeCPOTLockerMoreCPOTLockedIterator is returned from FilterMoreCPOTLocked and is used to iterate over the raw logs and unpacked data for MoreCPOTLocked events raised by the VeCPOTLocker contract.
type VeCPOTLockerMoreCPOTLockedIterator struct {
	Event *VeCPOTLockerMoreCPOTLocked // Event containing the contract specifics and raw log

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
func (it *VeCPOTLockerMoreCPOTLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeCPOTLockerMoreCPOTLocked)
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
		it.Event = new(VeCPOTLockerMoreCPOTLocked)
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
func (it *VeCPOTLockerMoreCPOTLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeCPOTLockerMoreCPOTLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeCPOTLockerMoreCPOTLocked represents a MoreCPOTLocked event raised by the VeCPOTLocker contract.
type VeCPOTLockerMoreCPOTLocked struct {
	User             common.Address
	LockId           *big.Int
	AdditionalAmount *big.Int
	NewVeAmount      *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMoreCPOTLocked is a free log retrieval operation binding the contract event 0xb08b72fd9435a140d4b63215204687ff8bdd8a15afb12d94c0c205d367ddba66.
//
// Solidity: event MoreCPOTLocked(address indexed user, uint256 lockId, uint256 additionalAmount, uint256 newVeAmount)
func (_VeCPOTLocker *VeCPOTLockerFilterer) FilterMoreCPOTLocked(opts *bind.FilterOpts, user []common.Address) (*VeCPOTLockerMoreCPOTLockedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _VeCPOTLocker.contract.FilterLogs(opts, "MoreCPOTLocked", userRule)
	if err != nil {
		return nil, err
	}
	return &VeCPOTLockerMoreCPOTLockedIterator{contract: _VeCPOTLocker.contract, event: "MoreCPOTLocked", logs: logs, sub: sub}, nil
}

// WatchMoreCPOTLocked is a free log subscription operation binding the contract event 0xb08b72fd9435a140d4b63215204687ff8bdd8a15afb12d94c0c205d367ddba66.
//
// Solidity: event MoreCPOTLocked(address indexed user, uint256 lockId, uint256 additionalAmount, uint256 newVeAmount)
func (_VeCPOTLocker *VeCPOTLockerFilterer) WatchMoreCPOTLocked(opts *bind.WatchOpts, sink chan<- *VeCPOTLockerMoreCPOTLocked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _VeCPOTLocker.contract.WatchLogs(opts, "MoreCPOTLocked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeCPOTLockerMoreCPOTLocked)
				if err := _VeCPOTLocker.contract.UnpackLog(event, "MoreCPOTLocked", log); err != nil {
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

// ParseMoreCPOTLocked is a log parse operation binding the contract event 0xb08b72fd9435a140d4b63215204687ff8bdd8a15afb12d94c0c205d367ddba66.
//
// Solidity: event MoreCPOTLocked(address indexed user, uint256 lockId, uint256 additionalAmount, uint256 newVeAmount)
func (_VeCPOTLocker *VeCPOTLockerFilterer) ParseMoreCPOTLocked(log types.Log) (*VeCPOTLockerMoreCPOTLocked, error) {
	event := new(VeCPOTLockerMoreCPOTLocked)
	if err := _VeCPOTLocker.contract.UnpackLog(event, "MoreCPOTLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeCPOTLockerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VeCPOTLocker contract.
type VeCPOTLockerOwnershipTransferredIterator struct {
	Event *VeCPOTLockerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VeCPOTLockerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeCPOTLockerOwnershipTransferred)
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
		it.Event = new(VeCPOTLockerOwnershipTransferred)
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
func (it *VeCPOTLockerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeCPOTLockerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeCPOTLockerOwnershipTransferred represents a OwnershipTransferred event raised by the VeCPOTLocker contract.
type VeCPOTLockerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VeCPOTLocker *VeCPOTLockerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VeCPOTLockerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VeCPOTLocker.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VeCPOTLockerOwnershipTransferredIterator{contract: _VeCPOTLocker.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VeCPOTLocker *VeCPOTLockerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VeCPOTLockerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VeCPOTLocker.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeCPOTLockerOwnershipTransferred)
				if err := _VeCPOTLocker.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VeCPOTLocker *VeCPOTLockerFilterer) ParseOwnershipTransferred(log types.Log) (*VeCPOTLockerOwnershipTransferred, error) {
	event := new(VeCPOTLockerOwnershipTransferred)
	if err := _VeCPOTLocker.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeCPOTLockerUnlockedCPOTIterator is returned from FilterUnlockedCPOT and is used to iterate over the raw logs and unpacked data for UnlockedCPOT events raised by the VeCPOTLocker contract.
type VeCPOTLockerUnlockedCPOTIterator struct {
	Event *VeCPOTLockerUnlockedCPOT // Event containing the contract specifics and raw log

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
func (it *VeCPOTLockerUnlockedCPOTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeCPOTLockerUnlockedCPOT)
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
		it.Event = new(VeCPOTLockerUnlockedCPOT)
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
func (it *VeCPOTLockerUnlockedCPOTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeCPOTLockerUnlockedCPOTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeCPOTLockerUnlockedCPOT represents a UnlockedCPOT event raised by the VeCPOTLocker contract.
type VeCPOTLockerUnlockedCPOT struct {
	User      common.Address
	LockId    *big.Int
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnlockedCPOT is a free log retrieval operation binding the contract event 0x6b70cac7d437336e3f4af649c888592f8503004a6962f5953390ca34d36cbe0b.
//
// Solidity: event UnlockedCPOT(address indexed user, uint256 lockId, uint256 amount, uint256 timestamp)
func (_VeCPOTLocker *VeCPOTLockerFilterer) FilterUnlockedCPOT(opts *bind.FilterOpts, user []common.Address) (*VeCPOTLockerUnlockedCPOTIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _VeCPOTLocker.contract.FilterLogs(opts, "UnlockedCPOT", userRule)
	if err != nil {
		return nil, err
	}
	return &VeCPOTLockerUnlockedCPOTIterator{contract: _VeCPOTLocker.contract, event: "UnlockedCPOT", logs: logs, sub: sub}, nil
}

// WatchUnlockedCPOT is a free log subscription operation binding the contract event 0x6b70cac7d437336e3f4af649c888592f8503004a6962f5953390ca34d36cbe0b.
//
// Solidity: event UnlockedCPOT(address indexed user, uint256 lockId, uint256 amount, uint256 timestamp)
func (_VeCPOTLocker *VeCPOTLockerFilterer) WatchUnlockedCPOT(opts *bind.WatchOpts, sink chan<- *VeCPOTLockerUnlockedCPOT, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _VeCPOTLocker.contract.WatchLogs(opts, "UnlockedCPOT", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeCPOTLockerUnlockedCPOT)
				if err := _VeCPOTLocker.contract.UnpackLog(event, "UnlockedCPOT", log); err != nil {
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

// ParseUnlockedCPOT is a log parse operation binding the contract event 0x6b70cac7d437336e3f4af649c888592f8503004a6962f5953390ca34d36cbe0b.
//
// Solidity: event UnlockedCPOT(address indexed user, uint256 lockId, uint256 amount, uint256 timestamp)
func (_VeCPOTLocker *VeCPOTLockerFilterer) ParseUnlockedCPOT(log types.Log) (*VeCPOTLockerUnlockedCPOT, error) {
	event := new(VeCPOTLockerUnlockedCPOT)
	if err := _VeCPOTLocker.contract.UnpackLog(event, "UnlockedCPOT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeCPOTLockerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the VeCPOTLocker contract.
type VeCPOTLockerUpgradedIterator struct {
	Event *VeCPOTLockerUpgraded // Event containing the contract specifics and raw log

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
func (it *VeCPOTLockerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeCPOTLockerUpgraded)
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
		it.Event = new(VeCPOTLockerUpgraded)
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
func (it *VeCPOTLockerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeCPOTLockerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeCPOTLockerUpgraded represents a Upgraded event raised by the VeCPOTLocker contract.
type VeCPOTLockerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_VeCPOTLocker *VeCPOTLockerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*VeCPOTLockerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _VeCPOTLocker.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &VeCPOTLockerUpgradedIterator{contract: _VeCPOTLocker.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_VeCPOTLocker *VeCPOTLockerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *VeCPOTLockerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _VeCPOTLocker.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeCPOTLockerUpgraded)
				if err := _VeCPOTLocker.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_VeCPOTLocker *VeCPOTLockerFilterer) ParseUpgraded(log types.Log) (*VeCPOTLockerUpgraded, error) {
	event := new(VeCPOTLockerUpgraded)
	if err := _VeCPOTLocker.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeCPOTLockerVaultSetIterator is returned from FilterVaultSet and is used to iterate over the raw logs and unpacked data for VaultSet events raised by the VeCPOTLocker contract.
type VeCPOTLockerVaultSetIterator struct {
	Event *VeCPOTLockerVaultSet // Event containing the contract specifics and raw log

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
func (it *VeCPOTLockerVaultSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeCPOTLockerVaultSet)
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
		it.Event = new(VeCPOTLockerVaultSet)
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
func (it *VeCPOTLockerVaultSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeCPOTLockerVaultSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeCPOTLockerVaultSet represents a VaultSet event raised by the VeCPOTLocker contract.
type VeCPOTLockerVaultSet struct {
	Vault common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVaultSet is a free log retrieval operation binding the contract event 0xe7ae49f883c825b05681b3e00e8be6fdea9ed2a8a45e4c6ecb9390fc44cce615.
//
// Solidity: event VaultSet(address indexed vault)
func (_VeCPOTLocker *VeCPOTLockerFilterer) FilterVaultSet(opts *bind.FilterOpts, vault []common.Address) (*VeCPOTLockerVaultSetIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _VeCPOTLocker.contract.FilterLogs(opts, "VaultSet", vaultRule)
	if err != nil {
		return nil, err
	}
	return &VeCPOTLockerVaultSetIterator{contract: _VeCPOTLocker.contract, event: "VaultSet", logs: logs, sub: sub}, nil
}

// WatchVaultSet is a free log subscription operation binding the contract event 0xe7ae49f883c825b05681b3e00e8be6fdea9ed2a8a45e4c6ecb9390fc44cce615.
//
// Solidity: event VaultSet(address indexed vault)
func (_VeCPOTLocker *VeCPOTLockerFilterer) WatchVaultSet(opts *bind.WatchOpts, sink chan<- *VeCPOTLockerVaultSet, vault []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _VeCPOTLocker.contract.WatchLogs(opts, "VaultSet", vaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeCPOTLockerVaultSet)
				if err := _VeCPOTLocker.contract.UnpackLog(event, "VaultSet", log); err != nil {
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

// ParseVaultSet is a log parse operation binding the contract event 0xe7ae49f883c825b05681b3e00e8be6fdea9ed2a8a45e4c6ecb9390fc44cce615.
//
// Solidity: event VaultSet(address indexed vault)
func (_VeCPOTLocker *VeCPOTLockerFilterer) ParseVaultSet(log types.Log) (*VeCPOTLockerVaultSet, error) {
	event := new(VeCPOTLockerVaultSet)
	if err := _VeCPOTLocker.contract.UnpackLog(event, "VaultSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
