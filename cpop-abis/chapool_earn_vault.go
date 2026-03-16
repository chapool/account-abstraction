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

// ChapoolEarnVaultMetaData contains all meta data concerning the ChapoolEarnVault contract.
var ChapoolEarnVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"accountManager\",\"type\":\"address\"}],\"name\":\"AccountManagerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldWeighted\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newWeighted\",\"type\":\"uint256\"}],\"name\":\"BoostSynced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aaAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CPPRewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"EmergencyModeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdrawCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdrawExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executeAfter\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdrawInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"NftBoostControllerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cppPerSecond\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"RewardRateSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"locker\",\"type\":\"address\"}],\"name\":\"VecpotLockerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BPS_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EMERGENCY_TIMELOCK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accCPPPerWeightedUSDT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accountManagerContract\",\"outputs\":[{\"internalType\":\"contractIAccountManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelEmergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimCPP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cppToken\",\"outputs\":[{\"internalType\":\"contractICPPToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"estimatedDailyCPP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"executeEmergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getPendingCPP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserBoostBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cppToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_accountManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"initiateEmergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nftBoostController\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pendingCPP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingEmergencyWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"scheduledAt\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pending\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"positions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"depositedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastActionAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_accountManager\",\"type\":\"address\"}],\"name\":\"setAccountManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setEmergencyMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"setNftBoostController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cppPerSecond\",\"type\":\"uint256\"}],\"name\":\"setRewardRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"locker\",\"type\":\"address\"}],\"name\":\"setVecpotLocker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"syncBoost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalUsers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalVaultAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalWeightedUSDT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"usdtBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vecpotLocker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"weightedUSDT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ChapoolEarnVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use ChapoolEarnVaultMetaData.ABI instead.
var ChapoolEarnVaultABI = ChapoolEarnVaultMetaData.ABI

// ChapoolEarnVault is an auto generated Go binding around an Ethereum contract.
type ChapoolEarnVault struct {
	ChapoolEarnVaultCaller     // Read-only binding to the contract
	ChapoolEarnVaultTransactor // Write-only binding to the contract
	ChapoolEarnVaultFilterer   // Log filterer for contract events
}

// ChapoolEarnVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChapoolEarnVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChapoolEarnVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChapoolEarnVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChapoolEarnVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChapoolEarnVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChapoolEarnVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChapoolEarnVaultSession struct {
	Contract     *ChapoolEarnVault // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChapoolEarnVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChapoolEarnVaultCallerSession struct {
	Contract *ChapoolEarnVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ChapoolEarnVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChapoolEarnVaultTransactorSession struct {
	Contract     *ChapoolEarnVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ChapoolEarnVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChapoolEarnVaultRaw struct {
	Contract *ChapoolEarnVault // Generic contract binding to access the raw methods on
}

// ChapoolEarnVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChapoolEarnVaultCallerRaw struct {
	Contract *ChapoolEarnVaultCaller // Generic read-only contract binding to access the raw methods on
}

// ChapoolEarnVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChapoolEarnVaultTransactorRaw struct {
	Contract *ChapoolEarnVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChapoolEarnVault creates a new instance of ChapoolEarnVault, bound to a specific deployed contract.
func NewChapoolEarnVault(address common.Address, backend bind.ContractBackend) (*ChapoolEarnVault, error) {
	contract, err := bindChapoolEarnVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChapoolEarnVault{ChapoolEarnVaultCaller: ChapoolEarnVaultCaller{contract: contract}, ChapoolEarnVaultTransactor: ChapoolEarnVaultTransactor{contract: contract}, ChapoolEarnVaultFilterer: ChapoolEarnVaultFilterer{contract: contract}}, nil
}

// NewChapoolEarnVaultCaller creates a new read-only instance of ChapoolEarnVault, bound to a specific deployed contract.
func NewChapoolEarnVaultCaller(address common.Address, caller bind.ContractCaller) (*ChapoolEarnVaultCaller, error) {
	contract, err := bindChapoolEarnVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChapoolEarnVaultCaller{contract: contract}, nil
}

// NewChapoolEarnVaultTransactor creates a new write-only instance of ChapoolEarnVault, bound to a specific deployed contract.
func NewChapoolEarnVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*ChapoolEarnVaultTransactor, error) {
	contract, err := bindChapoolEarnVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChapoolEarnVaultTransactor{contract: contract}, nil
}

// NewChapoolEarnVaultFilterer creates a new log filterer instance of ChapoolEarnVault, bound to a specific deployed contract.
func NewChapoolEarnVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*ChapoolEarnVaultFilterer, error) {
	contract, err := bindChapoolEarnVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChapoolEarnVaultFilterer{contract: contract}, nil
}

// bindChapoolEarnVault binds a generic wrapper to an already deployed contract.
func bindChapoolEarnVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ChapoolEarnVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChapoolEarnVault *ChapoolEarnVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChapoolEarnVault.Contract.ChapoolEarnVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChapoolEarnVault *ChapoolEarnVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.ChapoolEarnVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChapoolEarnVault *ChapoolEarnVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.ChapoolEarnVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChapoolEarnVault *ChapoolEarnVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChapoolEarnVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChapoolEarnVault *ChapoolEarnVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChapoolEarnVault *ChapoolEarnVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.contract.Transact(opts, method, params...)
}

// BPSDENOMINATOR is a free data retrieval call binding the contract method 0xe1a45218.
//
// Solidity: function BPS_DENOMINATOR() view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultCaller) BPSDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChapoolEarnVault.contract.Call(opts, &out, "BPS_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BPSDENOMINATOR is a free data retrieval call binding the contract method 0xe1a45218.
//
// Solidity: function BPS_DENOMINATOR() view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultSession) BPSDENOMINATOR() (*big.Int, error) {
	return _ChapoolEarnVault.Contract.BPSDENOMINATOR(&_ChapoolEarnVault.CallOpts)
}

// BPSDENOMINATOR is a free data retrieval call binding the contract method 0xe1a45218.
//
// Solidity: function BPS_DENOMINATOR() view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultCallerSession) BPSDENOMINATOR() (*big.Int, error) {
	return _ChapoolEarnVault.Contract.BPSDENOMINATOR(&_ChapoolEarnVault.CallOpts)
}

// EMERGENCYTIMELOCK is a free data retrieval call binding the contract method 0x60d7442b.
//
// Solidity: function EMERGENCY_TIMELOCK() view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultCaller) EMERGENCYTIMELOCK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChapoolEarnVault.contract.Call(opts, &out, "EMERGENCY_TIMELOCK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EMERGENCYTIMELOCK is a free data retrieval call binding the contract method 0x60d7442b.
//
// Solidity: function EMERGENCY_TIMELOCK() view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultSession) EMERGENCYTIMELOCK() (*big.Int, error) {
	return _ChapoolEarnVault.Contract.EMERGENCYTIMELOCK(&_ChapoolEarnVault.CallOpts)
}

// EMERGENCYTIMELOCK is a free data retrieval call binding the contract method 0x60d7442b.
//
// Solidity: function EMERGENCY_TIMELOCK() view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultCallerSession) EMERGENCYTIMELOCK() (*big.Int, error) {
	return _ChapoolEarnVault.Contract.EMERGENCYTIMELOCK(&_ChapoolEarnVault.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultCaller) PRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChapoolEarnVault.contract.Call(opts, &out, "PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultSession) PRECISION() (*big.Int, error) {
	return _ChapoolEarnVault.Contract.PRECISION(&_ChapoolEarnVault.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultCallerSession) PRECISION() (*big.Int, error) {
	return _ChapoolEarnVault.Contract.PRECISION(&_ChapoolEarnVault.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ChapoolEarnVault *ChapoolEarnVaultCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ChapoolEarnVault.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ChapoolEarnVault *ChapoolEarnVaultSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ChapoolEarnVault.Contract.UPGRADEINTERFACEVERSION(&_ChapoolEarnVault.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ChapoolEarnVault *ChapoolEarnVaultCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ChapoolEarnVault.Contract.UPGRADEINTERFACEVERSION(&_ChapoolEarnVault.CallOpts)
}

// AccCPPPerWeightedUSDT is a free data retrieval call binding the contract method 0x4f91cbc4.
//
// Solidity: function accCPPPerWeightedUSDT() view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultCaller) AccCPPPerWeightedUSDT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChapoolEarnVault.contract.Call(opts, &out, "accCPPPerWeightedUSDT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccCPPPerWeightedUSDT is a free data retrieval call binding the contract method 0x4f91cbc4.
//
// Solidity: function accCPPPerWeightedUSDT() view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultSession) AccCPPPerWeightedUSDT() (*big.Int, error) {
	return _ChapoolEarnVault.Contract.AccCPPPerWeightedUSDT(&_ChapoolEarnVault.CallOpts)
}

// AccCPPPerWeightedUSDT is a free data retrieval call binding the contract method 0x4f91cbc4.
//
// Solidity: function accCPPPerWeightedUSDT() view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultCallerSession) AccCPPPerWeightedUSDT() (*big.Int, error) {
	return _ChapoolEarnVault.Contract.AccCPPPerWeightedUSDT(&_ChapoolEarnVault.CallOpts)
}

// AccountManagerContract is a free data retrieval call binding the contract method 0x23676a71.
//
// Solidity: function accountManagerContract() view returns(address)
func (_ChapoolEarnVault *ChapoolEarnVaultCaller) AccountManagerContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChapoolEarnVault.contract.Call(opts, &out, "accountManagerContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountManagerContract is a free data retrieval call binding the contract method 0x23676a71.
//
// Solidity: function accountManagerContract() view returns(address)
func (_ChapoolEarnVault *ChapoolEarnVaultSession) AccountManagerContract() (common.Address, error) {
	return _ChapoolEarnVault.Contract.AccountManagerContract(&_ChapoolEarnVault.CallOpts)
}

// AccountManagerContract is a free data retrieval call binding the contract method 0x23676a71.
//
// Solidity: function accountManagerContract() view returns(address)
func (_ChapoolEarnVault *ChapoolEarnVaultCallerSession) AccountManagerContract() (common.Address, error) {
	return _ChapoolEarnVault.Contract.AccountManagerContract(&_ChapoolEarnVault.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_ChapoolEarnVault *ChapoolEarnVaultCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChapoolEarnVault.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_ChapoolEarnVault *ChapoolEarnVaultSession) Asset() (common.Address, error) {
	return _ChapoolEarnVault.Contract.Asset(&_ChapoolEarnVault.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_ChapoolEarnVault *ChapoolEarnVaultCallerSession) Asset() (common.Address, error) {
	return _ChapoolEarnVault.Contract.Asset(&_ChapoolEarnVault.CallOpts)
}

// CppToken is a free data retrieval call binding the contract method 0x94002996.
//
// Solidity: function cppToken() view returns(address)
func (_ChapoolEarnVault *ChapoolEarnVaultCaller) CppToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChapoolEarnVault.contract.Call(opts, &out, "cppToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CppToken is a free data retrieval call binding the contract method 0x94002996.
//
// Solidity: function cppToken() view returns(address)
func (_ChapoolEarnVault *ChapoolEarnVaultSession) CppToken() (common.Address, error) {
	return _ChapoolEarnVault.Contract.CppToken(&_ChapoolEarnVault.CallOpts)
}

// CppToken is a free data retrieval call binding the contract method 0x94002996.
//
// Solidity: function cppToken() view returns(address)
func (_ChapoolEarnVault *ChapoolEarnVaultCallerSession) CppToken() (common.Address, error) {
	return _ChapoolEarnVault.Contract.CppToken(&_ChapoolEarnVault.CallOpts)
}

// EmergencyMode is a free data retrieval call binding the contract method 0x0905f560.
//
// Solidity: function emergencyMode() view returns(bool)
func (_ChapoolEarnVault *ChapoolEarnVaultCaller) EmergencyMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ChapoolEarnVault.contract.Call(opts, &out, "emergencyMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EmergencyMode is a free data retrieval call binding the contract method 0x0905f560.
//
// Solidity: function emergencyMode() view returns(bool)
func (_ChapoolEarnVault *ChapoolEarnVaultSession) EmergencyMode() (bool, error) {
	return _ChapoolEarnVault.Contract.EmergencyMode(&_ChapoolEarnVault.CallOpts)
}

// EmergencyMode is a free data retrieval call binding the contract method 0x0905f560.
//
// Solidity: function emergencyMode() view returns(bool)
func (_ChapoolEarnVault *ChapoolEarnVaultCallerSession) EmergencyMode() (bool, error) {
	return _ChapoolEarnVault.Contract.EmergencyMode(&_ChapoolEarnVault.CallOpts)
}

// EstimatedDailyCPP is a free data retrieval call binding the contract method 0x0d1afde6.
//
// Solidity: function estimatedDailyCPP(address user) view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultCaller) EstimatedDailyCPP(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ChapoolEarnVault.contract.Call(opts, &out, "estimatedDailyCPP", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimatedDailyCPP is a free data retrieval call binding the contract method 0x0d1afde6.
//
// Solidity: function estimatedDailyCPP(address user) view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultSession) EstimatedDailyCPP(user common.Address) (*big.Int, error) {
	return _ChapoolEarnVault.Contract.EstimatedDailyCPP(&_ChapoolEarnVault.CallOpts, user)
}

// EstimatedDailyCPP is a free data retrieval call binding the contract method 0x0d1afde6.
//
// Solidity: function estimatedDailyCPP(address user) view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultCallerSession) EstimatedDailyCPP(user common.Address) (*big.Int, error) {
	return _ChapoolEarnVault.Contract.EstimatedDailyCPP(&_ChapoolEarnVault.CallOpts, user)
}

// GetPendingCPP is a free data retrieval call binding the contract method 0xd61497f6.
//
// Solidity: function getPendingCPP(address user) view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultCaller) GetPendingCPP(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ChapoolEarnVault.contract.Call(opts, &out, "getPendingCPP", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingCPP is a free data retrieval call binding the contract method 0xd61497f6.
//
// Solidity: function getPendingCPP(address user) view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultSession) GetPendingCPP(user common.Address) (*big.Int, error) {
	return _ChapoolEarnVault.Contract.GetPendingCPP(&_ChapoolEarnVault.CallOpts, user)
}

// GetPendingCPP is a free data retrieval call binding the contract method 0xd61497f6.
//
// Solidity: function getPendingCPP(address user) view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultCallerSession) GetPendingCPP(user common.Address) (*big.Int, error) {
	return _ChapoolEarnVault.Contract.GetPendingCPP(&_ChapoolEarnVault.CallOpts, user)
}

// GetUserBoostBps is a free data retrieval call binding the contract method 0x73468d42.
//
// Solidity: function getUserBoostBps(address user) view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultCaller) GetUserBoostBps(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ChapoolEarnVault.contract.Call(opts, &out, "getUserBoostBps", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserBoostBps is a free data retrieval call binding the contract method 0x73468d42.
//
// Solidity: function getUserBoostBps(address user) view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultSession) GetUserBoostBps(user common.Address) (*big.Int, error) {
	return _ChapoolEarnVault.Contract.GetUserBoostBps(&_ChapoolEarnVault.CallOpts, user)
}

// GetUserBoostBps is a free data retrieval call binding the contract method 0x73468d42.
//
// Solidity: function getUserBoostBps(address user) view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultCallerSession) GetUserBoostBps(user common.Address) (*big.Int, error) {
	return _ChapoolEarnVault.Contract.GetUserBoostBps(&_ChapoolEarnVault.CallOpts, user)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultCaller) LastUpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChapoolEarnVault.contract.Call(opts, &out, "lastUpdateTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultSession) LastUpdateTime() (*big.Int, error) {
	return _ChapoolEarnVault.Contract.LastUpdateTime(&_ChapoolEarnVault.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultCallerSession) LastUpdateTime() (*big.Int, error) {
	return _ChapoolEarnVault.Contract.LastUpdateTime(&_ChapoolEarnVault.CallOpts)
}

// NftBoostController is a free data retrieval call binding the contract method 0x2371b1bb.
//
// Solidity: function nftBoostController() view returns(address)
func (_ChapoolEarnVault *ChapoolEarnVaultCaller) NftBoostController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChapoolEarnVault.contract.Call(opts, &out, "nftBoostController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NftBoostController is a free data retrieval call binding the contract method 0x2371b1bb.
//
// Solidity: function nftBoostController() view returns(address)
func (_ChapoolEarnVault *ChapoolEarnVaultSession) NftBoostController() (common.Address, error) {
	return _ChapoolEarnVault.Contract.NftBoostController(&_ChapoolEarnVault.CallOpts)
}

// NftBoostController is a free data retrieval call binding the contract method 0x2371b1bb.
//
// Solidity: function nftBoostController() view returns(address)
func (_ChapoolEarnVault *ChapoolEarnVaultCallerSession) NftBoostController() (common.Address, error) {
	return _ChapoolEarnVault.Contract.NftBoostController(&_ChapoolEarnVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChapoolEarnVault *ChapoolEarnVaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChapoolEarnVault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChapoolEarnVault *ChapoolEarnVaultSession) Owner() (common.Address, error) {
	return _ChapoolEarnVault.Contract.Owner(&_ChapoolEarnVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChapoolEarnVault *ChapoolEarnVaultCallerSession) Owner() (common.Address, error) {
	return _ChapoolEarnVault.Contract.Owner(&_ChapoolEarnVault.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ChapoolEarnVault *ChapoolEarnVaultCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ChapoolEarnVault.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ChapoolEarnVault *ChapoolEarnVaultSession) Paused() (bool, error) {
	return _ChapoolEarnVault.Contract.Paused(&_ChapoolEarnVault.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ChapoolEarnVault *ChapoolEarnVaultCallerSession) Paused() (bool, error) {
	return _ChapoolEarnVault.Contract.Paused(&_ChapoolEarnVault.CallOpts)
}

// PendingCPP is a free data retrieval call binding the contract method 0xa284f6c8.
//
// Solidity: function pendingCPP(address ) view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultCaller) PendingCPP(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ChapoolEarnVault.contract.Call(opts, &out, "pendingCPP", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingCPP is a free data retrieval call binding the contract method 0xa284f6c8.
//
// Solidity: function pendingCPP(address ) view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultSession) PendingCPP(arg0 common.Address) (*big.Int, error) {
	return _ChapoolEarnVault.Contract.PendingCPP(&_ChapoolEarnVault.CallOpts, arg0)
}

// PendingCPP is a free data retrieval call binding the contract method 0xa284f6c8.
//
// Solidity: function pendingCPP(address ) view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultCallerSession) PendingCPP(arg0 common.Address) (*big.Int, error) {
	return _ChapoolEarnVault.Contract.PendingCPP(&_ChapoolEarnVault.CallOpts, arg0)
}

// PendingEmergencyWithdraw is a free data retrieval call binding the contract method 0xef2ec4db.
//
// Solidity: function pendingEmergencyWithdraw() view returns(uint256 amount, address to, uint256 scheduledAt, bool pending)
func (_ChapoolEarnVault *ChapoolEarnVaultCaller) PendingEmergencyWithdraw(opts *bind.CallOpts) (struct {
	Amount      *big.Int
	To          common.Address
	ScheduledAt *big.Int
	Pending     bool
}, error) {
	var out []interface{}
	err := _ChapoolEarnVault.contract.Call(opts, &out, "pendingEmergencyWithdraw")

	outstruct := new(struct {
		Amount      *big.Int
		To          common.Address
		ScheduledAt *big.Int
		Pending     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.To = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.ScheduledAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Pending = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// PendingEmergencyWithdraw is a free data retrieval call binding the contract method 0xef2ec4db.
//
// Solidity: function pendingEmergencyWithdraw() view returns(uint256 amount, address to, uint256 scheduledAt, bool pending)
func (_ChapoolEarnVault *ChapoolEarnVaultSession) PendingEmergencyWithdraw() (struct {
	Amount      *big.Int
	To          common.Address
	ScheduledAt *big.Int
	Pending     bool
}, error) {
	return _ChapoolEarnVault.Contract.PendingEmergencyWithdraw(&_ChapoolEarnVault.CallOpts)
}

// PendingEmergencyWithdraw is a free data retrieval call binding the contract method 0xef2ec4db.
//
// Solidity: function pendingEmergencyWithdraw() view returns(uint256 amount, address to, uint256 scheduledAt, bool pending)
func (_ChapoolEarnVault *ChapoolEarnVaultCallerSession) PendingEmergencyWithdraw() (struct {
	Amount      *big.Int
	To          common.Address
	ScheduledAt *big.Int
	Pending     bool
}, error) {
	return _ChapoolEarnVault.Contract.PendingEmergencyWithdraw(&_ChapoolEarnVault.CallOpts)
}

// Positions is a free data retrieval call binding the contract method 0x55f57510.
//
// Solidity: function positions(address ) view returns(uint256 depositedAt, uint256 lastActionAt)
func (_ChapoolEarnVault *ChapoolEarnVaultCaller) Positions(opts *bind.CallOpts, arg0 common.Address) (struct {
	DepositedAt  *big.Int
	LastActionAt *big.Int
}, error) {
	var out []interface{}
	err := _ChapoolEarnVault.contract.Call(opts, &out, "positions", arg0)

	outstruct := new(struct {
		DepositedAt  *big.Int
		LastActionAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DepositedAt = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastActionAt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Positions is a free data retrieval call binding the contract method 0x55f57510.
//
// Solidity: function positions(address ) view returns(uint256 depositedAt, uint256 lastActionAt)
func (_ChapoolEarnVault *ChapoolEarnVaultSession) Positions(arg0 common.Address) (struct {
	DepositedAt  *big.Int
	LastActionAt *big.Int
}, error) {
	return _ChapoolEarnVault.Contract.Positions(&_ChapoolEarnVault.CallOpts, arg0)
}

// Positions is a free data retrieval call binding the contract method 0x55f57510.
//
// Solidity: function positions(address ) view returns(uint256 depositedAt, uint256 lastActionAt)
func (_ChapoolEarnVault *ChapoolEarnVaultCallerSession) Positions(arg0 common.Address) (struct {
	DepositedAt  *big.Int
	LastActionAt *big.Int
}, error) {
	return _ChapoolEarnVault.Contract.Positions(&_ChapoolEarnVault.CallOpts, arg0)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ChapoolEarnVault *ChapoolEarnVaultCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ChapoolEarnVault.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ChapoolEarnVault *ChapoolEarnVaultSession) ProxiableUUID() ([32]byte, error) {
	return _ChapoolEarnVault.Contract.ProxiableUUID(&_ChapoolEarnVault.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ChapoolEarnVault *ChapoolEarnVaultCallerSession) ProxiableUUID() ([32]byte, error) {
	return _ChapoolEarnVault.Contract.ProxiableUUID(&_ChapoolEarnVault.CallOpts)
}

// RewardDebt is a free data retrieval call binding the contract method 0x5873eb9b.
//
// Solidity: function rewardDebt(address ) view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultCaller) RewardDebt(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ChapoolEarnVault.contract.Call(opts, &out, "rewardDebt", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardDebt is a free data retrieval call binding the contract method 0x5873eb9b.
//
// Solidity: function rewardDebt(address ) view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultSession) RewardDebt(arg0 common.Address) (*big.Int, error) {
	return _ChapoolEarnVault.Contract.RewardDebt(&_ChapoolEarnVault.CallOpts, arg0)
}

// RewardDebt is a free data retrieval call binding the contract method 0x5873eb9b.
//
// Solidity: function rewardDebt(address ) view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultCallerSession) RewardDebt(arg0 common.Address) (*big.Int, error) {
	return _ChapoolEarnVault.Contract.RewardDebt(&_ChapoolEarnVault.CallOpts, arg0)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultCaller) RewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChapoolEarnVault.contract.Call(opts, &out, "rewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultSession) RewardRate() (*big.Int, error) {
	return _ChapoolEarnVault.Contract.RewardRate(&_ChapoolEarnVault.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultCallerSession) RewardRate() (*big.Int, error) {
	return _ChapoolEarnVault.Contract.RewardRate(&_ChapoolEarnVault.CallOpts)
}

// TotalUsers is a free data retrieval call binding the contract method 0xbff1f9e1.
//
// Solidity: function totalUsers() view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultCaller) TotalUsers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChapoolEarnVault.contract.Call(opts, &out, "totalUsers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalUsers is a free data retrieval call binding the contract method 0xbff1f9e1.
//
// Solidity: function totalUsers() view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultSession) TotalUsers() (*big.Int, error) {
	return _ChapoolEarnVault.Contract.TotalUsers(&_ChapoolEarnVault.CallOpts)
}

// TotalUsers is a free data retrieval call binding the contract method 0xbff1f9e1.
//
// Solidity: function totalUsers() view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultCallerSession) TotalUsers() (*big.Int, error) {
	return _ChapoolEarnVault.Contract.TotalUsers(&_ChapoolEarnVault.CallOpts)
}

// TotalVaultAssets is a free data retrieval call binding the contract method 0xa28cb110.
//
// Solidity: function totalVaultAssets() view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultCaller) TotalVaultAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChapoolEarnVault.contract.Call(opts, &out, "totalVaultAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalVaultAssets is a free data retrieval call binding the contract method 0xa28cb110.
//
// Solidity: function totalVaultAssets() view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultSession) TotalVaultAssets() (*big.Int, error) {
	return _ChapoolEarnVault.Contract.TotalVaultAssets(&_ChapoolEarnVault.CallOpts)
}

// TotalVaultAssets is a free data retrieval call binding the contract method 0xa28cb110.
//
// Solidity: function totalVaultAssets() view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultCallerSession) TotalVaultAssets() (*big.Int, error) {
	return _ChapoolEarnVault.Contract.TotalVaultAssets(&_ChapoolEarnVault.CallOpts)
}

// TotalWeightedUSDT is a free data retrieval call binding the contract method 0x93b6dac5.
//
// Solidity: function totalWeightedUSDT() view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultCaller) TotalWeightedUSDT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChapoolEarnVault.contract.Call(opts, &out, "totalWeightedUSDT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalWeightedUSDT is a free data retrieval call binding the contract method 0x93b6dac5.
//
// Solidity: function totalWeightedUSDT() view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultSession) TotalWeightedUSDT() (*big.Int, error) {
	return _ChapoolEarnVault.Contract.TotalWeightedUSDT(&_ChapoolEarnVault.CallOpts)
}

// TotalWeightedUSDT is a free data retrieval call binding the contract method 0x93b6dac5.
//
// Solidity: function totalWeightedUSDT() view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultCallerSession) TotalWeightedUSDT() (*big.Int, error) {
	return _ChapoolEarnVault.Contract.TotalWeightedUSDT(&_ChapoolEarnVault.CallOpts)
}

// UsdtBalance is a free data retrieval call binding the contract method 0xaaa36eb4.
//
// Solidity: function usdtBalance(address ) view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultCaller) UsdtBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ChapoolEarnVault.contract.Call(opts, &out, "usdtBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdtBalance is a free data retrieval call binding the contract method 0xaaa36eb4.
//
// Solidity: function usdtBalance(address ) view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultSession) UsdtBalance(arg0 common.Address) (*big.Int, error) {
	return _ChapoolEarnVault.Contract.UsdtBalance(&_ChapoolEarnVault.CallOpts, arg0)
}

// UsdtBalance is a free data retrieval call binding the contract method 0xaaa36eb4.
//
// Solidity: function usdtBalance(address ) view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultCallerSession) UsdtBalance(arg0 common.Address) (*big.Int, error) {
	return _ChapoolEarnVault.Contract.UsdtBalance(&_ChapoolEarnVault.CallOpts, arg0)
}

// VecpotLocker is a free data retrieval call binding the contract method 0xfee61125.
//
// Solidity: function vecpotLocker() view returns(address)
func (_ChapoolEarnVault *ChapoolEarnVaultCaller) VecpotLocker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChapoolEarnVault.contract.Call(opts, &out, "vecpotLocker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VecpotLocker is a free data retrieval call binding the contract method 0xfee61125.
//
// Solidity: function vecpotLocker() view returns(address)
func (_ChapoolEarnVault *ChapoolEarnVaultSession) VecpotLocker() (common.Address, error) {
	return _ChapoolEarnVault.Contract.VecpotLocker(&_ChapoolEarnVault.CallOpts)
}

// VecpotLocker is a free data retrieval call binding the contract method 0xfee61125.
//
// Solidity: function vecpotLocker() view returns(address)
func (_ChapoolEarnVault *ChapoolEarnVaultCallerSession) VecpotLocker() (common.Address, error) {
	return _ChapoolEarnVault.Contract.VecpotLocker(&_ChapoolEarnVault.CallOpts)
}

// WeightedUSDT is a free data retrieval call binding the contract method 0x76d80199.
//
// Solidity: function weightedUSDT(address ) view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultCaller) WeightedUSDT(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ChapoolEarnVault.contract.Call(opts, &out, "weightedUSDT", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeightedUSDT is a free data retrieval call binding the contract method 0x76d80199.
//
// Solidity: function weightedUSDT(address ) view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultSession) WeightedUSDT(arg0 common.Address) (*big.Int, error) {
	return _ChapoolEarnVault.Contract.WeightedUSDT(&_ChapoolEarnVault.CallOpts, arg0)
}

// WeightedUSDT is a free data retrieval call binding the contract method 0x76d80199.
//
// Solidity: function weightedUSDT(address ) view returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultCallerSession) WeightedUSDT(arg0 common.Address) (*big.Int, error) {
	return _ChapoolEarnVault.Contract.WeightedUSDT(&_ChapoolEarnVault.CallOpts, arg0)
}

// CancelEmergencyWithdraw is a paid mutator transaction binding the contract method 0x683131c9.
//
// Solidity: function cancelEmergencyWithdraw() returns()
func (_ChapoolEarnVault *ChapoolEarnVaultTransactor) CancelEmergencyWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChapoolEarnVault.contract.Transact(opts, "cancelEmergencyWithdraw")
}

// CancelEmergencyWithdraw is a paid mutator transaction binding the contract method 0x683131c9.
//
// Solidity: function cancelEmergencyWithdraw() returns()
func (_ChapoolEarnVault *ChapoolEarnVaultSession) CancelEmergencyWithdraw() (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.CancelEmergencyWithdraw(&_ChapoolEarnVault.TransactOpts)
}

// CancelEmergencyWithdraw is a paid mutator transaction binding the contract method 0x683131c9.
//
// Solidity: function cancelEmergencyWithdraw() returns()
func (_ChapoolEarnVault *ChapoolEarnVaultTransactorSession) CancelEmergencyWithdraw() (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.CancelEmergencyWithdraw(&_ChapoolEarnVault.TransactOpts)
}

// ClaimCPP is a paid mutator transaction binding the contract method 0x71841d90.
//
// Solidity: function claimCPP() returns(uint256 amount)
func (_ChapoolEarnVault *ChapoolEarnVaultTransactor) ClaimCPP(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChapoolEarnVault.contract.Transact(opts, "claimCPP")
}

// ClaimCPP is a paid mutator transaction binding the contract method 0x71841d90.
//
// Solidity: function claimCPP() returns(uint256 amount)
func (_ChapoolEarnVault *ChapoolEarnVaultSession) ClaimCPP() (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.ClaimCPP(&_ChapoolEarnVault.TransactOpts)
}

// ClaimCPP is a paid mutator transaction binding the contract method 0x71841d90.
//
// Solidity: function claimCPP() returns(uint256 amount)
func (_ChapoolEarnVault *ChapoolEarnVaultTransactorSession) ClaimCPP() (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.ClaimCPP(&_ChapoolEarnVault.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultTransactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ChapoolEarnVault.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.Deposit(&_ChapoolEarnVault.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultTransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.Deposit(&_ChapoolEarnVault.TransactOpts, assets, receiver)
}

// ExecuteEmergencyWithdraw is a paid mutator transaction binding the contract method 0x2a8821cc.
//
// Solidity: function executeEmergencyWithdraw() returns()
func (_ChapoolEarnVault *ChapoolEarnVaultTransactor) ExecuteEmergencyWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChapoolEarnVault.contract.Transact(opts, "executeEmergencyWithdraw")
}

// ExecuteEmergencyWithdraw is a paid mutator transaction binding the contract method 0x2a8821cc.
//
// Solidity: function executeEmergencyWithdraw() returns()
func (_ChapoolEarnVault *ChapoolEarnVaultSession) ExecuteEmergencyWithdraw() (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.ExecuteEmergencyWithdraw(&_ChapoolEarnVault.TransactOpts)
}

// ExecuteEmergencyWithdraw is a paid mutator transaction binding the contract method 0x2a8821cc.
//
// Solidity: function executeEmergencyWithdraw() returns()
func (_ChapoolEarnVault *ChapoolEarnVaultTransactorSession) ExecuteEmergencyWithdraw() (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.ExecuteEmergencyWithdraw(&_ChapoolEarnVault.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _asset, address _cppToken, address _accountManager, address _owner) returns()
func (_ChapoolEarnVault *ChapoolEarnVaultTransactor) Initialize(opts *bind.TransactOpts, _asset common.Address, _cppToken common.Address, _accountManager common.Address, _owner common.Address) (*types.Transaction, error) {
	return _ChapoolEarnVault.contract.Transact(opts, "initialize", _asset, _cppToken, _accountManager, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _asset, address _cppToken, address _accountManager, address _owner) returns()
func (_ChapoolEarnVault *ChapoolEarnVaultSession) Initialize(_asset common.Address, _cppToken common.Address, _accountManager common.Address, _owner common.Address) (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.Initialize(&_ChapoolEarnVault.TransactOpts, _asset, _cppToken, _accountManager, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _asset, address _cppToken, address _accountManager, address _owner) returns()
func (_ChapoolEarnVault *ChapoolEarnVaultTransactorSession) Initialize(_asset common.Address, _cppToken common.Address, _accountManager common.Address, _owner common.Address) (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.Initialize(&_ChapoolEarnVault.TransactOpts, _asset, _cppToken, _accountManager, _owner)
}

// InitiateEmergencyWithdraw is a paid mutator transaction binding the contract method 0x8dd0f770.
//
// Solidity: function initiateEmergencyWithdraw(uint256 amount, address to) returns()
func (_ChapoolEarnVault *ChapoolEarnVaultTransactor) InitiateEmergencyWithdraw(opts *bind.TransactOpts, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _ChapoolEarnVault.contract.Transact(opts, "initiateEmergencyWithdraw", amount, to)
}

// InitiateEmergencyWithdraw is a paid mutator transaction binding the contract method 0x8dd0f770.
//
// Solidity: function initiateEmergencyWithdraw(uint256 amount, address to) returns()
func (_ChapoolEarnVault *ChapoolEarnVaultSession) InitiateEmergencyWithdraw(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.InitiateEmergencyWithdraw(&_ChapoolEarnVault.TransactOpts, amount, to)
}

// InitiateEmergencyWithdraw is a paid mutator transaction binding the contract method 0x8dd0f770.
//
// Solidity: function initiateEmergencyWithdraw(uint256 amount, address to) returns()
func (_ChapoolEarnVault *ChapoolEarnVaultTransactorSession) InitiateEmergencyWithdraw(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.InitiateEmergencyWithdraw(&_ChapoolEarnVault.TransactOpts, amount, to)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ChapoolEarnVault *ChapoolEarnVaultTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChapoolEarnVault.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ChapoolEarnVault *ChapoolEarnVaultSession) Pause() (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.Pause(&_ChapoolEarnVault.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ChapoolEarnVault *ChapoolEarnVaultTransactorSession) Pause() (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.Pause(&_ChapoolEarnVault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChapoolEarnVault *ChapoolEarnVaultTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChapoolEarnVault.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChapoolEarnVault *ChapoolEarnVaultSession) RenounceOwnership() (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.RenounceOwnership(&_ChapoolEarnVault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChapoolEarnVault *ChapoolEarnVaultTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.RenounceOwnership(&_ChapoolEarnVault.TransactOpts)
}

// SetAccountManager is a paid mutator transaction binding the contract method 0x8e4586af.
//
// Solidity: function setAccountManager(address _accountManager) returns()
func (_ChapoolEarnVault *ChapoolEarnVaultTransactor) SetAccountManager(opts *bind.TransactOpts, _accountManager common.Address) (*types.Transaction, error) {
	return _ChapoolEarnVault.contract.Transact(opts, "setAccountManager", _accountManager)
}

// SetAccountManager is a paid mutator transaction binding the contract method 0x8e4586af.
//
// Solidity: function setAccountManager(address _accountManager) returns()
func (_ChapoolEarnVault *ChapoolEarnVaultSession) SetAccountManager(_accountManager common.Address) (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.SetAccountManager(&_ChapoolEarnVault.TransactOpts, _accountManager)
}

// SetAccountManager is a paid mutator transaction binding the contract method 0x8e4586af.
//
// Solidity: function setAccountManager(address _accountManager) returns()
func (_ChapoolEarnVault *ChapoolEarnVaultTransactorSession) SetAccountManager(_accountManager common.Address) (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.SetAccountManager(&_ChapoolEarnVault.TransactOpts, _accountManager)
}

// SetEmergencyMode is a paid mutator transaction binding the contract method 0xbe32b3f8.
//
// Solidity: function setEmergencyMode(bool enabled) returns()
func (_ChapoolEarnVault *ChapoolEarnVaultTransactor) SetEmergencyMode(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _ChapoolEarnVault.contract.Transact(opts, "setEmergencyMode", enabled)
}

// SetEmergencyMode is a paid mutator transaction binding the contract method 0xbe32b3f8.
//
// Solidity: function setEmergencyMode(bool enabled) returns()
func (_ChapoolEarnVault *ChapoolEarnVaultSession) SetEmergencyMode(enabled bool) (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.SetEmergencyMode(&_ChapoolEarnVault.TransactOpts, enabled)
}

// SetEmergencyMode is a paid mutator transaction binding the contract method 0xbe32b3f8.
//
// Solidity: function setEmergencyMode(bool enabled) returns()
func (_ChapoolEarnVault *ChapoolEarnVaultTransactorSession) SetEmergencyMode(enabled bool) (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.SetEmergencyMode(&_ChapoolEarnVault.TransactOpts, enabled)
}

// SetNftBoostController is a paid mutator transaction binding the contract method 0xc0c82d84.
//
// Solidity: function setNftBoostController(address controller) returns()
func (_ChapoolEarnVault *ChapoolEarnVaultTransactor) SetNftBoostController(opts *bind.TransactOpts, controller common.Address) (*types.Transaction, error) {
	return _ChapoolEarnVault.contract.Transact(opts, "setNftBoostController", controller)
}

// SetNftBoostController is a paid mutator transaction binding the contract method 0xc0c82d84.
//
// Solidity: function setNftBoostController(address controller) returns()
func (_ChapoolEarnVault *ChapoolEarnVaultSession) SetNftBoostController(controller common.Address) (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.SetNftBoostController(&_ChapoolEarnVault.TransactOpts, controller)
}

// SetNftBoostController is a paid mutator transaction binding the contract method 0xc0c82d84.
//
// Solidity: function setNftBoostController(address controller) returns()
func (_ChapoolEarnVault *ChapoolEarnVaultTransactorSession) SetNftBoostController(controller common.Address) (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.SetNftBoostController(&_ChapoolEarnVault.TransactOpts, controller)
}

// SetRewardRate is a paid mutator transaction binding the contract method 0x9e447fc6.
//
// Solidity: function setRewardRate(uint256 cppPerSecond) returns()
func (_ChapoolEarnVault *ChapoolEarnVaultTransactor) SetRewardRate(opts *bind.TransactOpts, cppPerSecond *big.Int) (*types.Transaction, error) {
	return _ChapoolEarnVault.contract.Transact(opts, "setRewardRate", cppPerSecond)
}

// SetRewardRate is a paid mutator transaction binding the contract method 0x9e447fc6.
//
// Solidity: function setRewardRate(uint256 cppPerSecond) returns()
func (_ChapoolEarnVault *ChapoolEarnVaultSession) SetRewardRate(cppPerSecond *big.Int) (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.SetRewardRate(&_ChapoolEarnVault.TransactOpts, cppPerSecond)
}

// SetRewardRate is a paid mutator transaction binding the contract method 0x9e447fc6.
//
// Solidity: function setRewardRate(uint256 cppPerSecond) returns()
func (_ChapoolEarnVault *ChapoolEarnVaultTransactorSession) SetRewardRate(cppPerSecond *big.Int) (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.SetRewardRate(&_ChapoolEarnVault.TransactOpts, cppPerSecond)
}

// SetVecpotLocker is a paid mutator transaction binding the contract method 0xe3d48d4b.
//
// Solidity: function setVecpotLocker(address locker) returns()
func (_ChapoolEarnVault *ChapoolEarnVaultTransactor) SetVecpotLocker(opts *bind.TransactOpts, locker common.Address) (*types.Transaction, error) {
	return _ChapoolEarnVault.contract.Transact(opts, "setVecpotLocker", locker)
}

// SetVecpotLocker is a paid mutator transaction binding the contract method 0xe3d48d4b.
//
// Solidity: function setVecpotLocker(address locker) returns()
func (_ChapoolEarnVault *ChapoolEarnVaultSession) SetVecpotLocker(locker common.Address) (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.SetVecpotLocker(&_ChapoolEarnVault.TransactOpts, locker)
}

// SetVecpotLocker is a paid mutator transaction binding the contract method 0xe3d48d4b.
//
// Solidity: function setVecpotLocker(address locker) returns()
func (_ChapoolEarnVault *ChapoolEarnVaultTransactorSession) SetVecpotLocker(locker common.Address) (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.SetVecpotLocker(&_ChapoolEarnVault.TransactOpts, locker)
}

// SyncBoost is a paid mutator transaction binding the contract method 0x423297c6.
//
// Solidity: function syncBoost(address user) returns()
func (_ChapoolEarnVault *ChapoolEarnVaultTransactor) SyncBoost(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _ChapoolEarnVault.contract.Transact(opts, "syncBoost", user)
}

// SyncBoost is a paid mutator transaction binding the contract method 0x423297c6.
//
// Solidity: function syncBoost(address user) returns()
func (_ChapoolEarnVault *ChapoolEarnVaultSession) SyncBoost(user common.Address) (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.SyncBoost(&_ChapoolEarnVault.TransactOpts, user)
}

// SyncBoost is a paid mutator transaction binding the contract method 0x423297c6.
//
// Solidity: function syncBoost(address user) returns()
func (_ChapoolEarnVault *ChapoolEarnVaultTransactorSession) SyncBoost(user common.Address) (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.SyncBoost(&_ChapoolEarnVault.TransactOpts, user)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ChapoolEarnVault *ChapoolEarnVaultTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ChapoolEarnVault.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ChapoolEarnVault *ChapoolEarnVaultSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.TransferOwnership(&_ChapoolEarnVault.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ChapoolEarnVault *ChapoolEarnVaultTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.TransferOwnership(&_ChapoolEarnVault.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ChapoolEarnVault *ChapoolEarnVaultTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChapoolEarnVault.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ChapoolEarnVault *ChapoolEarnVaultSession) Unpause() (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.Unpause(&_ChapoolEarnVault.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ChapoolEarnVault *ChapoolEarnVaultTransactorSession) Unpause() (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.Unpause(&_ChapoolEarnVault.TransactOpts)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ChapoolEarnVault *ChapoolEarnVaultTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ChapoolEarnVault.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ChapoolEarnVault *ChapoolEarnVaultSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.UpgradeToAndCall(&_ChapoolEarnVault.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ChapoolEarnVault *ChapoolEarnVaultTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.UpgradeToAndCall(&_ChapoolEarnVault.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 assets, address receiver) returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultTransactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ChapoolEarnVault.contract.Transact(opts, "withdraw", assets, receiver)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 assets, address receiver) returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultSession) Withdraw(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.Withdraw(&_ChapoolEarnVault.TransactOpts, assets, receiver)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 assets, address receiver) returns(uint256)
func (_ChapoolEarnVault *ChapoolEarnVaultTransactorSession) Withdraw(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ChapoolEarnVault.Contract.Withdraw(&_ChapoolEarnVault.TransactOpts, assets, receiver)
}

// ChapoolEarnVaultAccountManagerSetIterator is returned from FilterAccountManagerSet and is used to iterate over the raw logs and unpacked data for AccountManagerSet events raised by the ChapoolEarnVault contract.
type ChapoolEarnVaultAccountManagerSetIterator struct {
	Event *ChapoolEarnVaultAccountManagerSet // Event containing the contract specifics and raw log

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
func (it *ChapoolEarnVaultAccountManagerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChapoolEarnVaultAccountManagerSet)
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
		it.Event = new(ChapoolEarnVaultAccountManagerSet)
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
func (it *ChapoolEarnVaultAccountManagerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChapoolEarnVaultAccountManagerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChapoolEarnVaultAccountManagerSet represents a AccountManagerSet event raised by the ChapoolEarnVault contract.
type ChapoolEarnVaultAccountManagerSet struct {
	AccountManager common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAccountManagerSet is a free log retrieval operation binding the contract event 0xb2ac641eff98719b19180b58c39bda050bee14e73e47358cbfdc2013c2f3819d.
//
// Solidity: event AccountManagerSet(address indexed accountManager)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) FilterAccountManagerSet(opts *bind.FilterOpts, accountManager []common.Address) (*ChapoolEarnVaultAccountManagerSetIterator, error) {

	var accountManagerRule []interface{}
	for _, accountManagerItem := range accountManager {
		accountManagerRule = append(accountManagerRule, accountManagerItem)
	}

	logs, sub, err := _ChapoolEarnVault.contract.FilterLogs(opts, "AccountManagerSet", accountManagerRule)
	if err != nil {
		return nil, err
	}
	return &ChapoolEarnVaultAccountManagerSetIterator{contract: _ChapoolEarnVault.contract, event: "AccountManagerSet", logs: logs, sub: sub}, nil
}

// WatchAccountManagerSet is a free log subscription operation binding the contract event 0xb2ac641eff98719b19180b58c39bda050bee14e73e47358cbfdc2013c2f3819d.
//
// Solidity: event AccountManagerSet(address indexed accountManager)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) WatchAccountManagerSet(opts *bind.WatchOpts, sink chan<- *ChapoolEarnVaultAccountManagerSet, accountManager []common.Address) (event.Subscription, error) {

	var accountManagerRule []interface{}
	for _, accountManagerItem := range accountManager {
		accountManagerRule = append(accountManagerRule, accountManagerItem)
	}

	logs, sub, err := _ChapoolEarnVault.contract.WatchLogs(opts, "AccountManagerSet", accountManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChapoolEarnVaultAccountManagerSet)
				if err := _ChapoolEarnVault.contract.UnpackLog(event, "AccountManagerSet", log); err != nil {
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

// ParseAccountManagerSet is a log parse operation binding the contract event 0xb2ac641eff98719b19180b58c39bda050bee14e73e47358cbfdc2013c2f3819d.
//
// Solidity: event AccountManagerSet(address indexed accountManager)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) ParseAccountManagerSet(log types.Log) (*ChapoolEarnVaultAccountManagerSet, error) {
	event := new(ChapoolEarnVaultAccountManagerSet)
	if err := _ChapoolEarnVault.contract.UnpackLog(event, "AccountManagerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChapoolEarnVaultBoostSyncedIterator is returned from FilterBoostSynced and is used to iterate over the raw logs and unpacked data for BoostSynced events raised by the ChapoolEarnVault contract.
type ChapoolEarnVaultBoostSyncedIterator struct {
	Event *ChapoolEarnVaultBoostSynced // Event containing the contract specifics and raw log

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
func (it *ChapoolEarnVaultBoostSyncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChapoolEarnVaultBoostSynced)
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
		it.Event = new(ChapoolEarnVaultBoostSynced)
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
func (it *ChapoolEarnVaultBoostSyncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChapoolEarnVaultBoostSyncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChapoolEarnVaultBoostSynced represents a BoostSynced event raised by the ChapoolEarnVault contract.
type ChapoolEarnVaultBoostSynced struct {
	User        common.Address
	OldWeighted *big.Int
	NewWeighted *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBoostSynced is a free log retrieval operation binding the contract event 0x5ea52e354c5bf30e5c4375ce0134e3e7bc444da0d737e86a24bf2e8b70259da1.
//
// Solidity: event BoostSynced(address indexed user, uint256 oldWeighted, uint256 newWeighted)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) FilterBoostSynced(opts *bind.FilterOpts, user []common.Address) (*ChapoolEarnVaultBoostSyncedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ChapoolEarnVault.contract.FilterLogs(opts, "BoostSynced", userRule)
	if err != nil {
		return nil, err
	}
	return &ChapoolEarnVaultBoostSyncedIterator{contract: _ChapoolEarnVault.contract, event: "BoostSynced", logs: logs, sub: sub}, nil
}

// WatchBoostSynced is a free log subscription operation binding the contract event 0x5ea52e354c5bf30e5c4375ce0134e3e7bc444da0d737e86a24bf2e8b70259da1.
//
// Solidity: event BoostSynced(address indexed user, uint256 oldWeighted, uint256 newWeighted)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) WatchBoostSynced(opts *bind.WatchOpts, sink chan<- *ChapoolEarnVaultBoostSynced, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ChapoolEarnVault.contract.WatchLogs(opts, "BoostSynced", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChapoolEarnVaultBoostSynced)
				if err := _ChapoolEarnVault.contract.UnpackLog(event, "BoostSynced", log); err != nil {
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

// ParseBoostSynced is a log parse operation binding the contract event 0x5ea52e354c5bf30e5c4375ce0134e3e7bc444da0d737e86a24bf2e8b70259da1.
//
// Solidity: event BoostSynced(address indexed user, uint256 oldWeighted, uint256 newWeighted)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) ParseBoostSynced(log types.Log) (*ChapoolEarnVaultBoostSynced, error) {
	event := new(ChapoolEarnVaultBoostSynced)
	if err := _ChapoolEarnVault.contract.UnpackLog(event, "BoostSynced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChapoolEarnVaultCPPRewardClaimedIterator is returned from FilterCPPRewardClaimed and is used to iterate over the raw logs and unpacked data for CPPRewardClaimed events raised by the ChapoolEarnVault contract.
type ChapoolEarnVaultCPPRewardClaimedIterator struct {
	Event *ChapoolEarnVaultCPPRewardClaimed // Event containing the contract specifics and raw log

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
func (it *ChapoolEarnVaultCPPRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChapoolEarnVaultCPPRewardClaimed)
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
		it.Event = new(ChapoolEarnVaultCPPRewardClaimed)
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
func (it *ChapoolEarnVaultCPPRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChapoolEarnVaultCPPRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChapoolEarnVaultCPPRewardClaimed represents a CPPRewardClaimed event raised by the ChapoolEarnVault contract.
type ChapoolEarnVaultCPPRewardClaimed struct {
	User      common.Address
	AaAccount common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCPPRewardClaimed is a free log retrieval operation binding the contract event 0x083be9ede0850f39b60b6aad1fa9a6b5bfb142a4cc100e4d4b4dc3424555cbfb.
//
// Solidity: event CPPRewardClaimed(address indexed user, address indexed aaAccount, uint256 amount, uint256 timestamp)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) FilterCPPRewardClaimed(opts *bind.FilterOpts, user []common.Address, aaAccount []common.Address) (*ChapoolEarnVaultCPPRewardClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var aaAccountRule []interface{}
	for _, aaAccountItem := range aaAccount {
		aaAccountRule = append(aaAccountRule, aaAccountItem)
	}

	logs, sub, err := _ChapoolEarnVault.contract.FilterLogs(opts, "CPPRewardClaimed", userRule, aaAccountRule)
	if err != nil {
		return nil, err
	}
	return &ChapoolEarnVaultCPPRewardClaimedIterator{contract: _ChapoolEarnVault.contract, event: "CPPRewardClaimed", logs: logs, sub: sub}, nil
}

// WatchCPPRewardClaimed is a free log subscription operation binding the contract event 0x083be9ede0850f39b60b6aad1fa9a6b5bfb142a4cc100e4d4b4dc3424555cbfb.
//
// Solidity: event CPPRewardClaimed(address indexed user, address indexed aaAccount, uint256 amount, uint256 timestamp)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) WatchCPPRewardClaimed(opts *bind.WatchOpts, sink chan<- *ChapoolEarnVaultCPPRewardClaimed, user []common.Address, aaAccount []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var aaAccountRule []interface{}
	for _, aaAccountItem := range aaAccount {
		aaAccountRule = append(aaAccountRule, aaAccountItem)
	}

	logs, sub, err := _ChapoolEarnVault.contract.WatchLogs(opts, "CPPRewardClaimed", userRule, aaAccountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChapoolEarnVaultCPPRewardClaimed)
				if err := _ChapoolEarnVault.contract.UnpackLog(event, "CPPRewardClaimed", log); err != nil {
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

// ParseCPPRewardClaimed is a log parse operation binding the contract event 0x083be9ede0850f39b60b6aad1fa9a6b5bfb142a4cc100e4d4b4dc3424555cbfb.
//
// Solidity: event CPPRewardClaimed(address indexed user, address indexed aaAccount, uint256 amount, uint256 timestamp)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) ParseCPPRewardClaimed(log types.Log) (*ChapoolEarnVaultCPPRewardClaimed, error) {
	event := new(ChapoolEarnVaultCPPRewardClaimed)
	if err := _ChapoolEarnVault.contract.UnpackLog(event, "CPPRewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChapoolEarnVaultDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the ChapoolEarnVault contract.
type ChapoolEarnVaultDepositedIterator struct {
	Event *ChapoolEarnVaultDeposited // Event containing the contract specifics and raw log

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
func (it *ChapoolEarnVaultDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChapoolEarnVaultDeposited)
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
		it.Event = new(ChapoolEarnVaultDeposited)
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
func (it *ChapoolEarnVaultDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChapoolEarnVaultDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChapoolEarnVaultDeposited represents a Deposited event raised by the ChapoolEarnVault contract.
type ChapoolEarnVaultDeposited struct {
	User      common.Address
	Assets    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x73a19dd210f1a7f902193214c0ee91dd35ee5b4d920cba8d519eca65a7b488ca.
//
// Solidity: event Deposited(address indexed user, uint256 assets, uint256 timestamp)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) FilterDeposited(opts *bind.FilterOpts, user []common.Address) (*ChapoolEarnVaultDepositedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ChapoolEarnVault.contract.FilterLogs(opts, "Deposited", userRule)
	if err != nil {
		return nil, err
	}
	return &ChapoolEarnVaultDepositedIterator{contract: _ChapoolEarnVault.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x73a19dd210f1a7f902193214c0ee91dd35ee5b4d920cba8d519eca65a7b488ca.
//
// Solidity: event Deposited(address indexed user, uint256 assets, uint256 timestamp)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *ChapoolEarnVaultDeposited, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ChapoolEarnVault.contract.WatchLogs(opts, "Deposited", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChapoolEarnVaultDeposited)
				if err := _ChapoolEarnVault.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x73a19dd210f1a7f902193214c0ee91dd35ee5b4d920cba8d519eca65a7b488ca.
//
// Solidity: event Deposited(address indexed user, uint256 assets, uint256 timestamp)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) ParseDeposited(log types.Log) (*ChapoolEarnVaultDeposited, error) {
	event := new(ChapoolEarnVaultDeposited)
	if err := _ChapoolEarnVault.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChapoolEarnVaultEmergencyModeSetIterator is returned from FilterEmergencyModeSet and is used to iterate over the raw logs and unpacked data for EmergencyModeSet events raised by the ChapoolEarnVault contract.
type ChapoolEarnVaultEmergencyModeSetIterator struct {
	Event *ChapoolEarnVaultEmergencyModeSet // Event containing the contract specifics and raw log

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
func (it *ChapoolEarnVaultEmergencyModeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChapoolEarnVaultEmergencyModeSet)
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
		it.Event = new(ChapoolEarnVaultEmergencyModeSet)
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
func (it *ChapoolEarnVaultEmergencyModeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChapoolEarnVaultEmergencyModeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChapoolEarnVaultEmergencyModeSet represents a EmergencyModeSet event raised by the ChapoolEarnVault contract.
type ChapoolEarnVaultEmergencyModeSet struct {
	Enabled   bool
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEmergencyModeSet is a free log retrieval operation binding the contract event 0x4dd6184000833260f54b5c2ff5681acbbe71e2e732f03270603842169cd7e4e3.
//
// Solidity: event EmergencyModeSet(bool enabled, uint256 timestamp)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) FilterEmergencyModeSet(opts *bind.FilterOpts) (*ChapoolEarnVaultEmergencyModeSetIterator, error) {

	logs, sub, err := _ChapoolEarnVault.contract.FilterLogs(opts, "EmergencyModeSet")
	if err != nil {
		return nil, err
	}
	return &ChapoolEarnVaultEmergencyModeSetIterator{contract: _ChapoolEarnVault.contract, event: "EmergencyModeSet", logs: logs, sub: sub}, nil
}

// WatchEmergencyModeSet is a free log subscription operation binding the contract event 0x4dd6184000833260f54b5c2ff5681acbbe71e2e732f03270603842169cd7e4e3.
//
// Solidity: event EmergencyModeSet(bool enabled, uint256 timestamp)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) WatchEmergencyModeSet(opts *bind.WatchOpts, sink chan<- *ChapoolEarnVaultEmergencyModeSet) (event.Subscription, error) {

	logs, sub, err := _ChapoolEarnVault.contract.WatchLogs(opts, "EmergencyModeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChapoolEarnVaultEmergencyModeSet)
				if err := _ChapoolEarnVault.contract.UnpackLog(event, "EmergencyModeSet", log); err != nil {
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

// ParseEmergencyModeSet is a log parse operation binding the contract event 0x4dd6184000833260f54b5c2ff5681acbbe71e2e732f03270603842169cd7e4e3.
//
// Solidity: event EmergencyModeSet(bool enabled, uint256 timestamp)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) ParseEmergencyModeSet(log types.Log) (*ChapoolEarnVaultEmergencyModeSet, error) {
	event := new(ChapoolEarnVaultEmergencyModeSet)
	if err := _ChapoolEarnVault.contract.UnpackLog(event, "EmergencyModeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChapoolEarnVaultEmergencyWithdrawCancelledIterator is returned from FilterEmergencyWithdrawCancelled and is used to iterate over the raw logs and unpacked data for EmergencyWithdrawCancelled events raised by the ChapoolEarnVault contract.
type ChapoolEarnVaultEmergencyWithdrawCancelledIterator struct {
	Event *ChapoolEarnVaultEmergencyWithdrawCancelled // Event containing the contract specifics and raw log

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
func (it *ChapoolEarnVaultEmergencyWithdrawCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChapoolEarnVaultEmergencyWithdrawCancelled)
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
		it.Event = new(ChapoolEarnVaultEmergencyWithdrawCancelled)
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
func (it *ChapoolEarnVaultEmergencyWithdrawCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChapoolEarnVaultEmergencyWithdrawCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChapoolEarnVaultEmergencyWithdrawCancelled represents a EmergencyWithdrawCancelled event raised by the ChapoolEarnVault contract.
type ChapoolEarnVaultEmergencyWithdrawCancelled struct {
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdrawCancelled is a free log retrieval operation binding the contract event 0x063c09fbefb35d5bb16d75da375f0e4c90a798eb8bfad5783b2cdcfbf8c89f83.
//
// Solidity: event EmergencyWithdrawCancelled(uint256 timestamp)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) FilterEmergencyWithdrawCancelled(opts *bind.FilterOpts) (*ChapoolEarnVaultEmergencyWithdrawCancelledIterator, error) {

	logs, sub, err := _ChapoolEarnVault.contract.FilterLogs(opts, "EmergencyWithdrawCancelled")
	if err != nil {
		return nil, err
	}
	return &ChapoolEarnVaultEmergencyWithdrawCancelledIterator{contract: _ChapoolEarnVault.contract, event: "EmergencyWithdrawCancelled", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdrawCancelled is a free log subscription operation binding the contract event 0x063c09fbefb35d5bb16d75da375f0e4c90a798eb8bfad5783b2cdcfbf8c89f83.
//
// Solidity: event EmergencyWithdrawCancelled(uint256 timestamp)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) WatchEmergencyWithdrawCancelled(opts *bind.WatchOpts, sink chan<- *ChapoolEarnVaultEmergencyWithdrawCancelled) (event.Subscription, error) {

	logs, sub, err := _ChapoolEarnVault.contract.WatchLogs(opts, "EmergencyWithdrawCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChapoolEarnVaultEmergencyWithdrawCancelled)
				if err := _ChapoolEarnVault.contract.UnpackLog(event, "EmergencyWithdrawCancelled", log); err != nil {
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

// ParseEmergencyWithdrawCancelled is a log parse operation binding the contract event 0x063c09fbefb35d5bb16d75da375f0e4c90a798eb8bfad5783b2cdcfbf8c89f83.
//
// Solidity: event EmergencyWithdrawCancelled(uint256 timestamp)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) ParseEmergencyWithdrawCancelled(log types.Log) (*ChapoolEarnVaultEmergencyWithdrawCancelled, error) {
	event := new(ChapoolEarnVaultEmergencyWithdrawCancelled)
	if err := _ChapoolEarnVault.contract.UnpackLog(event, "EmergencyWithdrawCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChapoolEarnVaultEmergencyWithdrawExecutedIterator is returned from FilterEmergencyWithdrawExecuted and is used to iterate over the raw logs and unpacked data for EmergencyWithdrawExecuted events raised by the ChapoolEarnVault contract.
type ChapoolEarnVaultEmergencyWithdrawExecutedIterator struct {
	Event *ChapoolEarnVaultEmergencyWithdrawExecuted // Event containing the contract specifics and raw log

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
func (it *ChapoolEarnVaultEmergencyWithdrawExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChapoolEarnVaultEmergencyWithdrawExecuted)
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
		it.Event = new(ChapoolEarnVaultEmergencyWithdrawExecuted)
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
func (it *ChapoolEarnVaultEmergencyWithdrawExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChapoolEarnVaultEmergencyWithdrawExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChapoolEarnVaultEmergencyWithdrawExecuted represents a EmergencyWithdrawExecuted event raised by the ChapoolEarnVault contract.
type ChapoolEarnVaultEmergencyWithdrawExecuted struct {
	Amount    *big.Int
	To        common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdrawExecuted is a free log retrieval operation binding the contract event 0xd8fb3598854d395a849b994e522dd276e4f0e3600d7bc444b8a04cf7e0ad72f5.
//
// Solidity: event EmergencyWithdrawExecuted(uint256 amount, address indexed to, uint256 timestamp)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) FilterEmergencyWithdrawExecuted(opts *bind.FilterOpts, to []common.Address) (*ChapoolEarnVaultEmergencyWithdrawExecutedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ChapoolEarnVault.contract.FilterLogs(opts, "EmergencyWithdrawExecuted", toRule)
	if err != nil {
		return nil, err
	}
	return &ChapoolEarnVaultEmergencyWithdrawExecutedIterator{contract: _ChapoolEarnVault.contract, event: "EmergencyWithdrawExecuted", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdrawExecuted is a free log subscription operation binding the contract event 0xd8fb3598854d395a849b994e522dd276e4f0e3600d7bc444b8a04cf7e0ad72f5.
//
// Solidity: event EmergencyWithdrawExecuted(uint256 amount, address indexed to, uint256 timestamp)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) WatchEmergencyWithdrawExecuted(opts *bind.WatchOpts, sink chan<- *ChapoolEarnVaultEmergencyWithdrawExecuted, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ChapoolEarnVault.contract.WatchLogs(opts, "EmergencyWithdrawExecuted", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChapoolEarnVaultEmergencyWithdrawExecuted)
				if err := _ChapoolEarnVault.contract.UnpackLog(event, "EmergencyWithdrawExecuted", log); err != nil {
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

// ParseEmergencyWithdrawExecuted is a log parse operation binding the contract event 0xd8fb3598854d395a849b994e522dd276e4f0e3600d7bc444b8a04cf7e0ad72f5.
//
// Solidity: event EmergencyWithdrawExecuted(uint256 amount, address indexed to, uint256 timestamp)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) ParseEmergencyWithdrawExecuted(log types.Log) (*ChapoolEarnVaultEmergencyWithdrawExecuted, error) {
	event := new(ChapoolEarnVaultEmergencyWithdrawExecuted)
	if err := _ChapoolEarnVault.contract.UnpackLog(event, "EmergencyWithdrawExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChapoolEarnVaultEmergencyWithdrawInitiatedIterator is returned from FilterEmergencyWithdrawInitiated and is used to iterate over the raw logs and unpacked data for EmergencyWithdrawInitiated events raised by the ChapoolEarnVault contract.
type ChapoolEarnVaultEmergencyWithdrawInitiatedIterator struct {
	Event *ChapoolEarnVaultEmergencyWithdrawInitiated // Event containing the contract specifics and raw log

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
func (it *ChapoolEarnVaultEmergencyWithdrawInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChapoolEarnVaultEmergencyWithdrawInitiated)
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
		it.Event = new(ChapoolEarnVaultEmergencyWithdrawInitiated)
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
func (it *ChapoolEarnVaultEmergencyWithdrawInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChapoolEarnVaultEmergencyWithdrawInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChapoolEarnVaultEmergencyWithdrawInitiated represents a EmergencyWithdrawInitiated event raised by the ChapoolEarnVault contract.
type ChapoolEarnVaultEmergencyWithdrawInitiated struct {
	Amount       *big.Int
	To           common.Address
	ExecuteAfter *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdrawInitiated is a free log retrieval operation binding the contract event 0xc1fcee2d2666ba0f2721abaf371b69f1b7ccf6fddaec164f4cc44281ae0af044.
//
// Solidity: event EmergencyWithdrawInitiated(uint256 amount, address indexed to, uint256 executeAfter)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) FilterEmergencyWithdrawInitiated(opts *bind.FilterOpts, to []common.Address) (*ChapoolEarnVaultEmergencyWithdrawInitiatedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ChapoolEarnVault.contract.FilterLogs(opts, "EmergencyWithdrawInitiated", toRule)
	if err != nil {
		return nil, err
	}
	return &ChapoolEarnVaultEmergencyWithdrawInitiatedIterator{contract: _ChapoolEarnVault.contract, event: "EmergencyWithdrawInitiated", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdrawInitiated is a free log subscription operation binding the contract event 0xc1fcee2d2666ba0f2721abaf371b69f1b7ccf6fddaec164f4cc44281ae0af044.
//
// Solidity: event EmergencyWithdrawInitiated(uint256 amount, address indexed to, uint256 executeAfter)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) WatchEmergencyWithdrawInitiated(opts *bind.WatchOpts, sink chan<- *ChapoolEarnVaultEmergencyWithdrawInitiated, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ChapoolEarnVault.contract.WatchLogs(opts, "EmergencyWithdrawInitiated", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChapoolEarnVaultEmergencyWithdrawInitiated)
				if err := _ChapoolEarnVault.contract.UnpackLog(event, "EmergencyWithdrawInitiated", log); err != nil {
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

// ParseEmergencyWithdrawInitiated is a log parse operation binding the contract event 0xc1fcee2d2666ba0f2721abaf371b69f1b7ccf6fddaec164f4cc44281ae0af044.
//
// Solidity: event EmergencyWithdrawInitiated(uint256 amount, address indexed to, uint256 executeAfter)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) ParseEmergencyWithdrawInitiated(log types.Log) (*ChapoolEarnVaultEmergencyWithdrawInitiated, error) {
	event := new(ChapoolEarnVaultEmergencyWithdrawInitiated)
	if err := _ChapoolEarnVault.contract.UnpackLog(event, "EmergencyWithdrawInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChapoolEarnVaultInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ChapoolEarnVault contract.
type ChapoolEarnVaultInitializedIterator struct {
	Event *ChapoolEarnVaultInitialized // Event containing the contract specifics and raw log

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
func (it *ChapoolEarnVaultInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChapoolEarnVaultInitialized)
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
		it.Event = new(ChapoolEarnVaultInitialized)
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
func (it *ChapoolEarnVaultInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChapoolEarnVaultInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChapoolEarnVaultInitialized represents a Initialized event raised by the ChapoolEarnVault contract.
type ChapoolEarnVaultInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) FilterInitialized(opts *bind.FilterOpts) (*ChapoolEarnVaultInitializedIterator, error) {

	logs, sub, err := _ChapoolEarnVault.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ChapoolEarnVaultInitializedIterator{contract: _ChapoolEarnVault.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ChapoolEarnVaultInitialized) (event.Subscription, error) {

	logs, sub, err := _ChapoolEarnVault.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChapoolEarnVaultInitialized)
				if err := _ChapoolEarnVault.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) ParseInitialized(log types.Log) (*ChapoolEarnVaultInitialized, error) {
	event := new(ChapoolEarnVaultInitialized)
	if err := _ChapoolEarnVault.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChapoolEarnVaultNftBoostControllerSetIterator is returned from FilterNftBoostControllerSet and is used to iterate over the raw logs and unpacked data for NftBoostControllerSet events raised by the ChapoolEarnVault contract.
type ChapoolEarnVaultNftBoostControllerSetIterator struct {
	Event *ChapoolEarnVaultNftBoostControllerSet // Event containing the contract specifics and raw log

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
func (it *ChapoolEarnVaultNftBoostControllerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChapoolEarnVaultNftBoostControllerSet)
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
		it.Event = new(ChapoolEarnVaultNftBoostControllerSet)
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
func (it *ChapoolEarnVaultNftBoostControllerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChapoolEarnVaultNftBoostControllerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChapoolEarnVaultNftBoostControllerSet represents a NftBoostControllerSet event raised by the ChapoolEarnVault contract.
type ChapoolEarnVaultNftBoostControllerSet struct {
	Controller common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNftBoostControllerSet is a free log retrieval operation binding the contract event 0xfb713de2df92903ed1db315b501b4345c43334edc4831308792e42371ad0f7b2.
//
// Solidity: event NftBoostControllerSet(address indexed controller)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) FilterNftBoostControllerSet(opts *bind.FilterOpts, controller []common.Address) (*ChapoolEarnVaultNftBoostControllerSetIterator, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _ChapoolEarnVault.contract.FilterLogs(opts, "NftBoostControllerSet", controllerRule)
	if err != nil {
		return nil, err
	}
	return &ChapoolEarnVaultNftBoostControllerSetIterator{contract: _ChapoolEarnVault.contract, event: "NftBoostControllerSet", logs: logs, sub: sub}, nil
}

// WatchNftBoostControllerSet is a free log subscription operation binding the contract event 0xfb713de2df92903ed1db315b501b4345c43334edc4831308792e42371ad0f7b2.
//
// Solidity: event NftBoostControllerSet(address indexed controller)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) WatchNftBoostControllerSet(opts *bind.WatchOpts, sink chan<- *ChapoolEarnVaultNftBoostControllerSet, controller []common.Address) (event.Subscription, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _ChapoolEarnVault.contract.WatchLogs(opts, "NftBoostControllerSet", controllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChapoolEarnVaultNftBoostControllerSet)
				if err := _ChapoolEarnVault.contract.UnpackLog(event, "NftBoostControllerSet", log); err != nil {
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

// ParseNftBoostControllerSet is a log parse operation binding the contract event 0xfb713de2df92903ed1db315b501b4345c43334edc4831308792e42371ad0f7b2.
//
// Solidity: event NftBoostControllerSet(address indexed controller)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) ParseNftBoostControllerSet(log types.Log) (*ChapoolEarnVaultNftBoostControllerSet, error) {
	event := new(ChapoolEarnVaultNftBoostControllerSet)
	if err := _ChapoolEarnVault.contract.UnpackLog(event, "NftBoostControllerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChapoolEarnVaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ChapoolEarnVault contract.
type ChapoolEarnVaultOwnershipTransferredIterator struct {
	Event *ChapoolEarnVaultOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ChapoolEarnVaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChapoolEarnVaultOwnershipTransferred)
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
		it.Event = new(ChapoolEarnVaultOwnershipTransferred)
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
func (it *ChapoolEarnVaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChapoolEarnVaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChapoolEarnVaultOwnershipTransferred represents a OwnershipTransferred event raised by the ChapoolEarnVault contract.
type ChapoolEarnVaultOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ChapoolEarnVaultOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ChapoolEarnVault.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ChapoolEarnVaultOwnershipTransferredIterator{contract: _ChapoolEarnVault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ChapoolEarnVaultOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ChapoolEarnVault.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChapoolEarnVaultOwnershipTransferred)
				if err := _ChapoolEarnVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) ParseOwnershipTransferred(log types.Log) (*ChapoolEarnVaultOwnershipTransferred, error) {
	event := new(ChapoolEarnVaultOwnershipTransferred)
	if err := _ChapoolEarnVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChapoolEarnVaultPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ChapoolEarnVault contract.
type ChapoolEarnVaultPausedIterator struct {
	Event *ChapoolEarnVaultPaused // Event containing the contract specifics and raw log

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
func (it *ChapoolEarnVaultPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChapoolEarnVaultPaused)
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
		it.Event = new(ChapoolEarnVaultPaused)
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
func (it *ChapoolEarnVaultPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChapoolEarnVaultPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChapoolEarnVaultPaused represents a Paused event raised by the ChapoolEarnVault contract.
type ChapoolEarnVaultPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) FilterPaused(opts *bind.FilterOpts) (*ChapoolEarnVaultPausedIterator, error) {

	logs, sub, err := _ChapoolEarnVault.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ChapoolEarnVaultPausedIterator{contract: _ChapoolEarnVault.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ChapoolEarnVaultPaused) (event.Subscription, error) {

	logs, sub, err := _ChapoolEarnVault.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChapoolEarnVaultPaused)
				if err := _ChapoolEarnVault.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) ParsePaused(log types.Log) (*ChapoolEarnVaultPaused, error) {
	event := new(ChapoolEarnVaultPaused)
	if err := _ChapoolEarnVault.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChapoolEarnVaultRewardRateSetIterator is returned from FilterRewardRateSet and is used to iterate over the raw logs and unpacked data for RewardRateSet events raised by the ChapoolEarnVault contract.
type ChapoolEarnVaultRewardRateSetIterator struct {
	Event *ChapoolEarnVaultRewardRateSet // Event containing the contract specifics and raw log

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
func (it *ChapoolEarnVaultRewardRateSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChapoolEarnVaultRewardRateSet)
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
		it.Event = new(ChapoolEarnVaultRewardRateSet)
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
func (it *ChapoolEarnVaultRewardRateSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChapoolEarnVaultRewardRateSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChapoolEarnVaultRewardRateSet represents a RewardRateSet event raised by the ChapoolEarnVault contract.
type ChapoolEarnVaultRewardRateSet struct {
	CppPerSecond *big.Int
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRewardRateSet is a free log retrieval operation binding the contract event 0x45e1b857394f54806156e6c3c35b79213af3de735bf9b1832622d802aa5ff918.
//
// Solidity: event RewardRateSet(uint256 cppPerSecond, uint256 timestamp)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) FilterRewardRateSet(opts *bind.FilterOpts) (*ChapoolEarnVaultRewardRateSetIterator, error) {

	logs, sub, err := _ChapoolEarnVault.contract.FilterLogs(opts, "RewardRateSet")
	if err != nil {
		return nil, err
	}
	return &ChapoolEarnVaultRewardRateSetIterator{contract: _ChapoolEarnVault.contract, event: "RewardRateSet", logs: logs, sub: sub}, nil
}

// WatchRewardRateSet is a free log subscription operation binding the contract event 0x45e1b857394f54806156e6c3c35b79213af3de735bf9b1832622d802aa5ff918.
//
// Solidity: event RewardRateSet(uint256 cppPerSecond, uint256 timestamp)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) WatchRewardRateSet(opts *bind.WatchOpts, sink chan<- *ChapoolEarnVaultRewardRateSet) (event.Subscription, error) {

	logs, sub, err := _ChapoolEarnVault.contract.WatchLogs(opts, "RewardRateSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChapoolEarnVaultRewardRateSet)
				if err := _ChapoolEarnVault.contract.UnpackLog(event, "RewardRateSet", log); err != nil {
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

// ParseRewardRateSet is a log parse operation binding the contract event 0x45e1b857394f54806156e6c3c35b79213af3de735bf9b1832622d802aa5ff918.
//
// Solidity: event RewardRateSet(uint256 cppPerSecond, uint256 timestamp)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) ParseRewardRateSet(log types.Log) (*ChapoolEarnVaultRewardRateSet, error) {
	event := new(ChapoolEarnVaultRewardRateSet)
	if err := _ChapoolEarnVault.contract.UnpackLog(event, "RewardRateSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChapoolEarnVaultUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ChapoolEarnVault contract.
type ChapoolEarnVaultUnpausedIterator struct {
	Event *ChapoolEarnVaultUnpaused // Event containing the contract specifics and raw log

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
func (it *ChapoolEarnVaultUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChapoolEarnVaultUnpaused)
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
		it.Event = new(ChapoolEarnVaultUnpaused)
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
func (it *ChapoolEarnVaultUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChapoolEarnVaultUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChapoolEarnVaultUnpaused represents a Unpaused event raised by the ChapoolEarnVault contract.
type ChapoolEarnVaultUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ChapoolEarnVaultUnpausedIterator, error) {

	logs, sub, err := _ChapoolEarnVault.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ChapoolEarnVaultUnpausedIterator{contract: _ChapoolEarnVault.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ChapoolEarnVaultUnpaused) (event.Subscription, error) {

	logs, sub, err := _ChapoolEarnVault.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChapoolEarnVaultUnpaused)
				if err := _ChapoolEarnVault.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) ParseUnpaused(log types.Log) (*ChapoolEarnVaultUnpaused, error) {
	event := new(ChapoolEarnVaultUnpaused)
	if err := _ChapoolEarnVault.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChapoolEarnVaultUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ChapoolEarnVault contract.
type ChapoolEarnVaultUpgradedIterator struct {
	Event *ChapoolEarnVaultUpgraded // Event containing the contract specifics and raw log

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
func (it *ChapoolEarnVaultUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChapoolEarnVaultUpgraded)
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
		it.Event = new(ChapoolEarnVaultUpgraded)
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
func (it *ChapoolEarnVaultUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChapoolEarnVaultUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChapoolEarnVaultUpgraded represents a Upgraded event raised by the ChapoolEarnVault contract.
type ChapoolEarnVaultUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ChapoolEarnVaultUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ChapoolEarnVault.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ChapoolEarnVaultUpgradedIterator{contract: _ChapoolEarnVault.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ChapoolEarnVaultUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ChapoolEarnVault.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChapoolEarnVaultUpgraded)
				if err := _ChapoolEarnVault.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) ParseUpgraded(log types.Log) (*ChapoolEarnVaultUpgraded, error) {
	event := new(ChapoolEarnVaultUpgraded)
	if err := _ChapoolEarnVault.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChapoolEarnVaultVecpotLockerSetIterator is returned from FilterVecpotLockerSet and is used to iterate over the raw logs and unpacked data for VecpotLockerSet events raised by the ChapoolEarnVault contract.
type ChapoolEarnVaultVecpotLockerSetIterator struct {
	Event *ChapoolEarnVaultVecpotLockerSet // Event containing the contract specifics and raw log

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
func (it *ChapoolEarnVaultVecpotLockerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChapoolEarnVaultVecpotLockerSet)
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
		it.Event = new(ChapoolEarnVaultVecpotLockerSet)
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
func (it *ChapoolEarnVaultVecpotLockerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChapoolEarnVaultVecpotLockerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChapoolEarnVaultVecpotLockerSet represents a VecpotLockerSet event raised by the ChapoolEarnVault contract.
type ChapoolEarnVaultVecpotLockerSet struct {
	Locker common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterVecpotLockerSet is a free log retrieval operation binding the contract event 0x6e6fcf56757336fc8ea967d5a6e92e794b598b6f562c0b87c362ca59605528a3.
//
// Solidity: event VecpotLockerSet(address indexed locker)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) FilterVecpotLockerSet(opts *bind.FilterOpts, locker []common.Address) (*ChapoolEarnVaultVecpotLockerSetIterator, error) {

	var lockerRule []interface{}
	for _, lockerItem := range locker {
		lockerRule = append(lockerRule, lockerItem)
	}

	logs, sub, err := _ChapoolEarnVault.contract.FilterLogs(opts, "VecpotLockerSet", lockerRule)
	if err != nil {
		return nil, err
	}
	return &ChapoolEarnVaultVecpotLockerSetIterator{contract: _ChapoolEarnVault.contract, event: "VecpotLockerSet", logs: logs, sub: sub}, nil
}

// WatchVecpotLockerSet is a free log subscription operation binding the contract event 0x6e6fcf56757336fc8ea967d5a6e92e794b598b6f562c0b87c362ca59605528a3.
//
// Solidity: event VecpotLockerSet(address indexed locker)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) WatchVecpotLockerSet(opts *bind.WatchOpts, sink chan<- *ChapoolEarnVaultVecpotLockerSet, locker []common.Address) (event.Subscription, error) {

	var lockerRule []interface{}
	for _, lockerItem := range locker {
		lockerRule = append(lockerRule, lockerItem)
	}

	logs, sub, err := _ChapoolEarnVault.contract.WatchLogs(opts, "VecpotLockerSet", lockerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChapoolEarnVaultVecpotLockerSet)
				if err := _ChapoolEarnVault.contract.UnpackLog(event, "VecpotLockerSet", log); err != nil {
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

// ParseVecpotLockerSet is a log parse operation binding the contract event 0x6e6fcf56757336fc8ea967d5a6e92e794b598b6f562c0b87c362ca59605528a3.
//
// Solidity: event VecpotLockerSet(address indexed locker)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) ParseVecpotLockerSet(log types.Log) (*ChapoolEarnVaultVecpotLockerSet, error) {
	event := new(ChapoolEarnVaultVecpotLockerSet)
	if err := _ChapoolEarnVault.contract.UnpackLog(event, "VecpotLockerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChapoolEarnVaultWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the ChapoolEarnVault contract.
type ChapoolEarnVaultWithdrawnIterator struct {
	Event *ChapoolEarnVaultWithdrawn // Event containing the contract specifics and raw log

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
func (it *ChapoolEarnVaultWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChapoolEarnVaultWithdrawn)
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
		it.Event = new(ChapoolEarnVaultWithdrawn)
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
func (it *ChapoolEarnVaultWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChapoolEarnVaultWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChapoolEarnVaultWithdrawn represents a Withdrawn event raised by the ChapoolEarnVault contract.
type ChapoolEarnVaultWithdrawn struct {
	User      common.Address
	Assets    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed user, uint256 assets, uint256 timestamp)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*ChapoolEarnVaultWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ChapoolEarnVault.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &ChapoolEarnVaultWithdrawnIterator{contract: _ChapoolEarnVault.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed user, uint256 assets, uint256 timestamp)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *ChapoolEarnVaultWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ChapoolEarnVault.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChapoolEarnVaultWithdrawn)
				if err := _ChapoolEarnVault.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed user, uint256 assets, uint256 timestamp)
func (_ChapoolEarnVault *ChapoolEarnVaultFilterer) ParseWithdrawn(log types.Log) (*ChapoolEarnVaultWithdrawn, error) {
	event := new(ChapoolEarnVaultWithdrawn)
	if err := _ChapoolEarnVault.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
