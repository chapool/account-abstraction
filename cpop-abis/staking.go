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

// StakingComboStatus is an auto generated low-level Go binding around an user-defined struct.
type StakingComboStatus struct {
	Level         uint8
	Count         *big.Int
	EffectiveFrom *big.Int
	Bonus         *big.Int
	IsPending     bool
}

// StakingMetaData contains all meta data concerning the Staking contract.
var StakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BatchRewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BatchStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalRewards\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BatchUnstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"NFTStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"NFTUnstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"staked\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"}],\"name\":\"PlatformStatsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"RewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CPPTokenContract\",\"outputs\":[{\"internalType\":\"contractICPPToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accountManagerContract\",\"outputs\":[{\"internalType\":\"contractIAccountManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"batchClaimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"batchStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"batchUnstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"calculatePendingRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configContract\",\"outputs\":[{\"internalType\":\"contractStakingConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cpnftContract\",\"outputs\":[{\"internalType\":\"contractCPNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableTestMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialTimestamp\",\"type\":\"uint256\"}],\"name\":\"enableTestMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"days_\",\"type\":\"uint256\"}],\"name\":\"fastForwardDays\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minutes_\",\"type\":\"uint256\"}],\"name\":\"fastForwardMinutes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"seconds_\",\"type\":\"uint256\"}],\"name\":\"fastForwardTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"}],\"name\":\"getComboStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"effectiveFrom\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bonus\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPending\",\"type\":\"bool\"}],\"internalType\":\"structStaking.ComboStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dayOffset\",\"type\":\"uint256\"}],\"name\":\"getDailyRewardBreakdown\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"baseReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decayedReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quarterlyMultiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dynamicMultiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalReward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"}],\"name\":\"getEffectiveComboBonus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getHistoricalAdjustment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quarterlyMultiplier\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getHistoricalAdjustmentCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"}],\"name\":\"getHistoricalDynamicMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"getHistoricalQuarterlyMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"historicalAdjustments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quarterlyMultiplier\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_cpnftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_CPPTokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_accountManagerContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_configContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recordHistoricalAdjustment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"setTestTimestamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"stakeTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastClaimTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"totalRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingRewards\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"continuousBonusClaimed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"testMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"testTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStakedCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"totalStakedPerLevel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_cpnftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_CPPTokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_accountManagerContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_configContract\",\"type\":\"address\"}],\"name\":\"updateContracts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"userComboStatus\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"effectiveFrom\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bonus\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPending\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"userLevelCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userStakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// StakingABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingMetaData.ABI instead.
var StakingABI = StakingMetaData.ABI

// Staking is an auto generated Go binding around an Ethereum contract.
type Staking struct {
	StakingCaller     // Read-only binding to the contract
	StakingTransactor // Write-only binding to the contract
	StakingFilterer   // Log filterer for contract events
}

// StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingSession struct {
	Contract     *Staking          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingCallerSession struct {
	Contract *StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingTransactorSession struct {
	Contract     *StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingRaw struct {
	Contract *Staking // Generic contract binding to access the raw methods on
}

// StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingCallerRaw struct {
	Contract *StakingCaller // Generic read-only contract binding to access the raw methods on
}

// StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingTransactorRaw struct {
	Contract *StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaking creates a new instance of Staking, bound to a specific deployed contract.
func NewStaking(address common.Address, backend bind.ContractBackend) (*Staking, error) {
	contract, err := bindStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// NewStakingCaller creates a new read-only instance of Staking, bound to a specific deployed contract.
func NewStakingCaller(address common.Address, caller bind.ContractCaller) (*StakingCaller, error) {
	contract, err := bindStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCaller{contract: contract}, nil
}

// NewStakingTransactor creates a new write-only instance of Staking, bound to a specific deployed contract.
func NewStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingTransactor, error) {
	contract, err := bindStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTransactor{contract: contract}, nil
}

// NewStakingFilterer creates a new log filterer instance of Staking, bound to a specific deployed contract.
func NewStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingFilterer, error) {
	contract, err := bindStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingFilterer{contract: contract}, nil
}

// bindStaking binds a generic wrapper to an already deployed contract.
func bindStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transact(opts, method, params...)
}

// CPPTokenContract is a free data retrieval call binding the contract method 0xb0956c59.
//
// Solidity: function CPPTokenContract() view returns(address)
func (_Staking *StakingCaller) CPPTokenContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "CPPTokenContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CPPTokenContract is a free data retrieval call binding the contract method 0xb0956c59.
//
// Solidity: function CPPTokenContract() view returns(address)
func (_Staking *StakingSession) CPPTokenContract() (common.Address, error) {
	return _Staking.Contract.CPPTokenContract(&_Staking.CallOpts)
}

// CPPTokenContract is a free data retrieval call binding the contract method 0xb0956c59.
//
// Solidity: function CPPTokenContract() view returns(address)
func (_Staking *StakingCallerSession) CPPTokenContract() (common.Address, error) {
	return _Staking.Contract.CPPTokenContract(&_Staking.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Staking *StakingCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Staking *StakingSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Staking.Contract.UPGRADEINTERFACEVERSION(&_Staking.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Staking *StakingCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Staking.Contract.UPGRADEINTERFACEVERSION(&_Staking.CallOpts)
}

// AccountManagerContract is a free data retrieval call binding the contract method 0x23676a71.
//
// Solidity: function accountManagerContract() view returns(address)
func (_Staking *StakingCaller) AccountManagerContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "accountManagerContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountManagerContract is a free data retrieval call binding the contract method 0x23676a71.
//
// Solidity: function accountManagerContract() view returns(address)
func (_Staking *StakingSession) AccountManagerContract() (common.Address, error) {
	return _Staking.Contract.AccountManagerContract(&_Staking.CallOpts)
}

// AccountManagerContract is a free data retrieval call binding the contract method 0x23676a71.
//
// Solidity: function accountManagerContract() view returns(address)
func (_Staking *StakingCallerSession) AccountManagerContract() (common.Address, error) {
	return _Staking.Contract.AccountManagerContract(&_Staking.CallOpts)
}

// CalculatePendingRewards is a free data retrieval call binding the contract method 0xb6ed7d62.
//
// Solidity: function calculatePendingRewards(uint256 tokenId) view returns(uint256)
func (_Staking *StakingCaller) CalculatePendingRewards(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "calculatePendingRewards", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculatePendingRewards is a free data retrieval call binding the contract method 0xb6ed7d62.
//
// Solidity: function calculatePendingRewards(uint256 tokenId) view returns(uint256)
func (_Staking *StakingSession) CalculatePendingRewards(tokenId *big.Int) (*big.Int, error) {
	return _Staking.Contract.CalculatePendingRewards(&_Staking.CallOpts, tokenId)
}

// CalculatePendingRewards is a free data retrieval call binding the contract method 0xb6ed7d62.
//
// Solidity: function calculatePendingRewards(uint256 tokenId) view returns(uint256)
func (_Staking *StakingCallerSession) CalculatePendingRewards(tokenId *big.Int) (*big.Int, error) {
	return _Staking.Contract.CalculatePendingRewards(&_Staking.CallOpts, tokenId)
}

// ConfigContract is a free data retrieval call binding the contract method 0xbf66a182.
//
// Solidity: function configContract() view returns(address)
func (_Staking *StakingCaller) ConfigContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "configContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConfigContract is a free data retrieval call binding the contract method 0xbf66a182.
//
// Solidity: function configContract() view returns(address)
func (_Staking *StakingSession) ConfigContract() (common.Address, error) {
	return _Staking.Contract.ConfigContract(&_Staking.CallOpts)
}

// ConfigContract is a free data retrieval call binding the contract method 0xbf66a182.
//
// Solidity: function configContract() view returns(address)
func (_Staking *StakingCallerSession) ConfigContract() (common.Address, error) {
	return _Staking.Contract.ConfigContract(&_Staking.CallOpts)
}

// CpnftContract is a free data retrieval call binding the contract method 0xabc99ee4.
//
// Solidity: function cpnftContract() view returns(address)
func (_Staking *StakingCaller) CpnftContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "cpnftContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CpnftContract is a free data retrieval call binding the contract method 0xabc99ee4.
//
// Solidity: function cpnftContract() view returns(address)
func (_Staking *StakingSession) CpnftContract() (common.Address, error) {
	return _Staking.Contract.CpnftContract(&_Staking.CallOpts)
}

// CpnftContract is a free data retrieval call binding the contract method 0xabc99ee4.
//
// Solidity: function cpnftContract() view returns(address)
func (_Staking *StakingCallerSession) CpnftContract() (common.Address, error) {
	return _Staking.Contract.CpnftContract(&_Staking.CallOpts)
}

// GetComboStatus is a free data retrieval call binding the contract method 0x2a20af8e.
//
// Solidity: function getComboStatus(address user, uint8 level) view returns((uint8,uint256,uint256,uint256,bool))
func (_Staking *StakingCaller) GetComboStatus(opts *bind.CallOpts, user common.Address, level uint8) (StakingComboStatus, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getComboStatus", user, level)

	if err != nil {
		return *new(StakingComboStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(StakingComboStatus)).(*StakingComboStatus)

	return out0, err

}

// GetComboStatus is a free data retrieval call binding the contract method 0x2a20af8e.
//
// Solidity: function getComboStatus(address user, uint8 level) view returns((uint8,uint256,uint256,uint256,bool))
func (_Staking *StakingSession) GetComboStatus(user common.Address, level uint8) (StakingComboStatus, error) {
	return _Staking.Contract.GetComboStatus(&_Staking.CallOpts, user, level)
}

// GetComboStatus is a free data retrieval call binding the contract method 0x2a20af8e.
//
// Solidity: function getComboStatus(address user, uint8 level) view returns((uint8,uint256,uint256,uint256,bool))
func (_Staking *StakingCallerSession) GetComboStatus(user common.Address, level uint8) (StakingComboStatus, error) {
	return _Staking.Contract.GetComboStatus(&_Staking.CallOpts, user, level)
}

// GetCurrentTimestamp is a free data retrieval call binding the contract method 0x6c9230db.
//
// Solidity: function getCurrentTimestamp() view returns(uint256)
func (_Staking *StakingCaller) GetCurrentTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getCurrentTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentTimestamp is a free data retrieval call binding the contract method 0x6c9230db.
//
// Solidity: function getCurrentTimestamp() view returns(uint256)
func (_Staking *StakingSession) GetCurrentTimestamp() (*big.Int, error) {
	return _Staking.Contract.GetCurrentTimestamp(&_Staking.CallOpts)
}

// GetCurrentTimestamp is a free data retrieval call binding the contract method 0x6c9230db.
//
// Solidity: function getCurrentTimestamp() view returns(uint256)
func (_Staking *StakingCallerSession) GetCurrentTimestamp() (*big.Int, error) {
	return _Staking.Contract.GetCurrentTimestamp(&_Staking.CallOpts)
}

// GetDailyRewardBreakdown is a free data retrieval call binding the contract method 0x515a10f9.
//
// Solidity: function getDailyRewardBreakdown(uint256 tokenId, uint256 dayOffset) view returns(uint256 baseReward, uint256 decayedReward, uint256 quarterlyMultiplier, uint256 dynamicMultiplier, uint256 finalReward)
func (_Staking *StakingCaller) GetDailyRewardBreakdown(opts *bind.CallOpts, tokenId *big.Int, dayOffset *big.Int) (struct {
	BaseReward          *big.Int
	DecayedReward       *big.Int
	QuarterlyMultiplier *big.Int
	DynamicMultiplier   *big.Int
	FinalReward         *big.Int
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getDailyRewardBreakdown", tokenId, dayOffset)

	outstruct := new(struct {
		BaseReward          *big.Int
		DecayedReward       *big.Int
		QuarterlyMultiplier *big.Int
		DynamicMultiplier   *big.Int
		FinalReward         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BaseReward = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DecayedReward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.QuarterlyMultiplier = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.DynamicMultiplier = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.FinalReward = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetDailyRewardBreakdown is a free data retrieval call binding the contract method 0x515a10f9.
//
// Solidity: function getDailyRewardBreakdown(uint256 tokenId, uint256 dayOffset) view returns(uint256 baseReward, uint256 decayedReward, uint256 quarterlyMultiplier, uint256 dynamicMultiplier, uint256 finalReward)
func (_Staking *StakingSession) GetDailyRewardBreakdown(tokenId *big.Int, dayOffset *big.Int) (struct {
	BaseReward          *big.Int
	DecayedReward       *big.Int
	QuarterlyMultiplier *big.Int
	DynamicMultiplier   *big.Int
	FinalReward         *big.Int
}, error) {
	return _Staking.Contract.GetDailyRewardBreakdown(&_Staking.CallOpts, tokenId, dayOffset)
}

// GetDailyRewardBreakdown is a free data retrieval call binding the contract method 0x515a10f9.
//
// Solidity: function getDailyRewardBreakdown(uint256 tokenId, uint256 dayOffset) view returns(uint256 baseReward, uint256 decayedReward, uint256 quarterlyMultiplier, uint256 dynamicMultiplier, uint256 finalReward)
func (_Staking *StakingCallerSession) GetDailyRewardBreakdown(tokenId *big.Int, dayOffset *big.Int) (struct {
	BaseReward          *big.Int
	DecayedReward       *big.Int
	QuarterlyMultiplier *big.Int
	DynamicMultiplier   *big.Int
	FinalReward         *big.Int
}, error) {
	return _Staking.Contract.GetDailyRewardBreakdown(&_Staking.CallOpts, tokenId, dayOffset)
}

// GetEffectiveComboBonus is a free data retrieval call binding the contract method 0x572501b5.
//
// Solidity: function getEffectiveComboBonus(address user, uint8 level) view returns(uint256)
func (_Staking *StakingCaller) GetEffectiveComboBonus(opts *bind.CallOpts, user common.Address, level uint8) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getEffectiveComboBonus", user, level)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEffectiveComboBonus is a free data retrieval call binding the contract method 0x572501b5.
//
// Solidity: function getEffectiveComboBonus(address user, uint8 level) view returns(uint256)
func (_Staking *StakingSession) GetEffectiveComboBonus(user common.Address, level uint8) (*big.Int, error) {
	return _Staking.Contract.GetEffectiveComboBonus(&_Staking.CallOpts, user, level)
}

// GetEffectiveComboBonus is a free data retrieval call binding the contract method 0x572501b5.
//
// Solidity: function getEffectiveComboBonus(address user, uint8 level) view returns(uint256)
func (_Staking *StakingCallerSession) GetEffectiveComboBonus(user common.Address, level uint8) (*big.Int, error) {
	return _Staking.Contract.GetEffectiveComboBonus(&_Staking.CallOpts, user, level)
}

// GetHistoricalAdjustment is a free data retrieval call binding the contract method 0xe5cbfa2e.
//
// Solidity: function getHistoricalAdjustment(uint256 index) view returns(uint256 timestamp, uint256 quarterlyMultiplier)
func (_Staking *StakingCaller) GetHistoricalAdjustment(opts *bind.CallOpts, index *big.Int) (struct {
	Timestamp           *big.Int
	QuarterlyMultiplier *big.Int
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getHistoricalAdjustment", index)

	outstruct := new(struct {
		Timestamp           *big.Int
		QuarterlyMultiplier *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.QuarterlyMultiplier = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetHistoricalAdjustment is a free data retrieval call binding the contract method 0xe5cbfa2e.
//
// Solidity: function getHistoricalAdjustment(uint256 index) view returns(uint256 timestamp, uint256 quarterlyMultiplier)
func (_Staking *StakingSession) GetHistoricalAdjustment(index *big.Int) (struct {
	Timestamp           *big.Int
	QuarterlyMultiplier *big.Int
}, error) {
	return _Staking.Contract.GetHistoricalAdjustment(&_Staking.CallOpts, index)
}

// GetHistoricalAdjustment is a free data retrieval call binding the contract method 0xe5cbfa2e.
//
// Solidity: function getHistoricalAdjustment(uint256 index) view returns(uint256 timestamp, uint256 quarterlyMultiplier)
func (_Staking *StakingCallerSession) GetHistoricalAdjustment(index *big.Int) (struct {
	Timestamp           *big.Int
	QuarterlyMultiplier *big.Int
}, error) {
	return _Staking.Contract.GetHistoricalAdjustment(&_Staking.CallOpts, index)
}

// GetHistoricalAdjustmentCount is a free data retrieval call binding the contract method 0x60cdf4e2.
//
// Solidity: function getHistoricalAdjustmentCount() view returns(uint256)
func (_Staking *StakingCaller) GetHistoricalAdjustmentCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getHistoricalAdjustmentCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetHistoricalAdjustmentCount is a free data retrieval call binding the contract method 0x60cdf4e2.
//
// Solidity: function getHistoricalAdjustmentCount() view returns(uint256)
func (_Staking *StakingSession) GetHistoricalAdjustmentCount() (*big.Int, error) {
	return _Staking.Contract.GetHistoricalAdjustmentCount(&_Staking.CallOpts)
}

// GetHistoricalAdjustmentCount is a free data retrieval call binding the contract method 0x60cdf4e2.
//
// Solidity: function getHistoricalAdjustmentCount() view returns(uint256)
func (_Staking *StakingCallerSession) GetHistoricalAdjustmentCount() (*big.Int, error) {
	return _Staking.Contract.GetHistoricalAdjustmentCount(&_Staking.CallOpts)
}

// GetHistoricalDynamicMultiplier is a free data retrieval call binding the contract method 0x79fe2082.
//
// Solidity: function getHistoricalDynamicMultiplier(uint256 index, uint8 level) view returns(uint256)
func (_Staking *StakingCaller) GetHistoricalDynamicMultiplier(opts *bind.CallOpts, index *big.Int, level uint8) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getHistoricalDynamicMultiplier", index, level)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetHistoricalDynamicMultiplier is a free data retrieval call binding the contract method 0x79fe2082.
//
// Solidity: function getHistoricalDynamicMultiplier(uint256 index, uint8 level) view returns(uint256)
func (_Staking *StakingSession) GetHistoricalDynamicMultiplier(index *big.Int, level uint8) (*big.Int, error) {
	return _Staking.Contract.GetHistoricalDynamicMultiplier(&_Staking.CallOpts, index, level)
}

// GetHistoricalDynamicMultiplier is a free data retrieval call binding the contract method 0x79fe2082.
//
// Solidity: function getHistoricalDynamicMultiplier(uint256 index, uint8 level) view returns(uint256)
func (_Staking *StakingCallerSession) GetHistoricalDynamicMultiplier(index *big.Int, level uint8) (*big.Int, error) {
	return _Staking.Contract.GetHistoricalDynamicMultiplier(&_Staking.CallOpts, index, level)
}

// GetHistoricalQuarterlyMultiplier is a free data retrieval call binding the contract method 0x1782ed78.
//
// Solidity: function getHistoricalQuarterlyMultiplier(uint256 timestamp) view returns(uint256)
func (_Staking *StakingCaller) GetHistoricalQuarterlyMultiplier(opts *bind.CallOpts, timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getHistoricalQuarterlyMultiplier", timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetHistoricalQuarterlyMultiplier is a free data retrieval call binding the contract method 0x1782ed78.
//
// Solidity: function getHistoricalQuarterlyMultiplier(uint256 timestamp) view returns(uint256)
func (_Staking *StakingSession) GetHistoricalQuarterlyMultiplier(timestamp *big.Int) (*big.Int, error) {
	return _Staking.Contract.GetHistoricalQuarterlyMultiplier(&_Staking.CallOpts, timestamp)
}

// GetHistoricalQuarterlyMultiplier is a free data retrieval call binding the contract method 0x1782ed78.
//
// Solidity: function getHistoricalQuarterlyMultiplier(uint256 timestamp) view returns(uint256)
func (_Staking *StakingCallerSession) GetHistoricalQuarterlyMultiplier(timestamp *big.Int) (*big.Int, error) {
	return _Staking.Contract.GetHistoricalQuarterlyMultiplier(&_Staking.CallOpts, timestamp)
}

// HistoricalAdjustments is a free data retrieval call binding the contract method 0x8af6582c.
//
// Solidity: function historicalAdjustments(uint256 ) view returns(uint256 timestamp, uint256 quarterlyMultiplier)
func (_Staking *StakingCaller) HistoricalAdjustments(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Timestamp           *big.Int
	QuarterlyMultiplier *big.Int
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "historicalAdjustments", arg0)

	outstruct := new(struct {
		Timestamp           *big.Int
		QuarterlyMultiplier *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.QuarterlyMultiplier = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// HistoricalAdjustments is a free data retrieval call binding the contract method 0x8af6582c.
//
// Solidity: function historicalAdjustments(uint256 ) view returns(uint256 timestamp, uint256 quarterlyMultiplier)
func (_Staking *StakingSession) HistoricalAdjustments(arg0 *big.Int) (struct {
	Timestamp           *big.Int
	QuarterlyMultiplier *big.Int
}, error) {
	return _Staking.Contract.HistoricalAdjustments(&_Staking.CallOpts, arg0)
}

// HistoricalAdjustments is a free data retrieval call binding the contract method 0x8af6582c.
//
// Solidity: function historicalAdjustments(uint256 ) view returns(uint256 timestamp, uint256 quarterlyMultiplier)
func (_Staking *StakingCallerSession) HistoricalAdjustments(arg0 *big.Int) (struct {
	Timestamp           *big.Int
	QuarterlyMultiplier *big.Int
}, error) {
	return _Staking.Contract.HistoricalAdjustments(&_Staking.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingSession) Owner() (common.Address, error) {
	return _Staking.Contract.Owner(&_Staking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingCallerSession) Owner() (common.Address, error) {
	return _Staking.Contract.Owner(&_Staking.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Staking *StakingCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Staking *StakingSession) Paused() (bool, error) {
	return _Staking.Contract.Paused(&_Staking.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Staking *StakingCallerSession) Paused() (bool, error) {
	return _Staking.Contract.Paused(&_Staking.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Staking *StakingCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Staking *StakingSession) ProxiableUUID() ([32]byte, error) {
	return _Staking.Contract.ProxiableUUID(&_Staking.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Staking *StakingCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Staking.Contract.ProxiableUUID(&_Staking.CallOpts)
}

// Stakes is a free data retrieval call binding the contract method 0xd5a44f86.
//
// Solidity: function stakes(uint256 ) view returns(address owner, uint256 tokenId, uint8 level, uint256 stakeTime, uint256 lastClaimTime, bool isActive, uint256 totalRewards, uint256 pendingRewards, bool continuousBonusClaimed)
func (_Staking *StakingCaller) Stakes(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner                  common.Address
	TokenId                *big.Int
	Level                  uint8
	StakeTime              *big.Int
	LastClaimTime          *big.Int
	IsActive               bool
	TotalRewards           *big.Int
	PendingRewards         *big.Int
	ContinuousBonusClaimed bool
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "stakes", arg0)

	outstruct := new(struct {
		Owner                  common.Address
		TokenId                *big.Int
		Level                  uint8
		StakeTime              *big.Int
		LastClaimTime          *big.Int
		IsActive               bool
		TotalRewards           *big.Int
		PendingRewards         *big.Int
		ContinuousBonusClaimed bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Level = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.StakeTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LastClaimTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.IsActive = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.TotalRewards = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.PendingRewards = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.ContinuousBonusClaimed = *abi.ConvertType(out[8], new(bool)).(*bool)

	return *outstruct, err

}

// Stakes is a free data retrieval call binding the contract method 0xd5a44f86.
//
// Solidity: function stakes(uint256 ) view returns(address owner, uint256 tokenId, uint8 level, uint256 stakeTime, uint256 lastClaimTime, bool isActive, uint256 totalRewards, uint256 pendingRewards, bool continuousBonusClaimed)
func (_Staking *StakingSession) Stakes(arg0 *big.Int) (struct {
	Owner                  common.Address
	TokenId                *big.Int
	Level                  uint8
	StakeTime              *big.Int
	LastClaimTime          *big.Int
	IsActive               bool
	TotalRewards           *big.Int
	PendingRewards         *big.Int
	ContinuousBonusClaimed bool
}, error) {
	return _Staking.Contract.Stakes(&_Staking.CallOpts, arg0)
}

// Stakes is a free data retrieval call binding the contract method 0xd5a44f86.
//
// Solidity: function stakes(uint256 ) view returns(address owner, uint256 tokenId, uint8 level, uint256 stakeTime, uint256 lastClaimTime, bool isActive, uint256 totalRewards, uint256 pendingRewards, bool continuousBonusClaimed)
func (_Staking *StakingCallerSession) Stakes(arg0 *big.Int) (struct {
	Owner                  common.Address
	TokenId                *big.Int
	Level                  uint8
	StakeTime              *big.Int
	LastClaimTime          *big.Int
	IsActive               bool
	TotalRewards           *big.Int
	PendingRewards         *big.Int
	ContinuousBonusClaimed bool
}, error) {
	return _Staking.Contract.Stakes(&_Staking.CallOpts, arg0)
}

// TestMode is a free data retrieval call binding the contract method 0xcd9ea342.
//
// Solidity: function testMode() view returns(bool)
func (_Staking *StakingCaller) TestMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "testMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TestMode is a free data retrieval call binding the contract method 0xcd9ea342.
//
// Solidity: function testMode() view returns(bool)
func (_Staking *StakingSession) TestMode() (bool, error) {
	return _Staking.Contract.TestMode(&_Staking.CallOpts)
}

// TestMode is a free data retrieval call binding the contract method 0xcd9ea342.
//
// Solidity: function testMode() view returns(bool)
func (_Staking *StakingCallerSession) TestMode() (bool, error) {
	return _Staking.Contract.TestMode(&_Staking.CallOpts)
}

// TestTimestamp is a free data retrieval call binding the contract method 0x7e773d9d.
//
// Solidity: function testTimestamp() view returns(uint256)
func (_Staking *StakingCaller) TestTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "testTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TestTimestamp is a free data retrieval call binding the contract method 0x7e773d9d.
//
// Solidity: function testTimestamp() view returns(uint256)
func (_Staking *StakingSession) TestTimestamp() (*big.Int, error) {
	return _Staking.Contract.TestTimestamp(&_Staking.CallOpts)
}

// TestTimestamp is a free data retrieval call binding the contract method 0x7e773d9d.
//
// Solidity: function testTimestamp() view returns(uint256)
func (_Staking *StakingCallerSession) TestTimestamp() (*big.Int, error) {
	return _Staking.Contract.TestTimestamp(&_Staking.CallOpts)
}

// TotalStakedCount is a free data retrieval call binding the contract method 0x920387b4.
//
// Solidity: function totalStakedCount() view returns(uint256)
func (_Staking *StakingCaller) TotalStakedCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "totalStakedCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStakedCount is a free data retrieval call binding the contract method 0x920387b4.
//
// Solidity: function totalStakedCount() view returns(uint256)
func (_Staking *StakingSession) TotalStakedCount() (*big.Int, error) {
	return _Staking.Contract.TotalStakedCount(&_Staking.CallOpts)
}

// TotalStakedCount is a free data retrieval call binding the contract method 0x920387b4.
//
// Solidity: function totalStakedCount() view returns(uint256)
func (_Staking *StakingCallerSession) TotalStakedCount() (*big.Int, error) {
	return _Staking.Contract.TotalStakedCount(&_Staking.CallOpts)
}

// TotalStakedPerLevel is a free data retrieval call binding the contract method 0x56445218.
//
// Solidity: function totalStakedPerLevel(uint8 ) view returns(uint256)
func (_Staking *StakingCaller) TotalStakedPerLevel(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "totalStakedPerLevel", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStakedPerLevel is a free data retrieval call binding the contract method 0x56445218.
//
// Solidity: function totalStakedPerLevel(uint8 ) view returns(uint256)
func (_Staking *StakingSession) TotalStakedPerLevel(arg0 uint8) (*big.Int, error) {
	return _Staking.Contract.TotalStakedPerLevel(&_Staking.CallOpts, arg0)
}

// TotalStakedPerLevel is a free data retrieval call binding the contract method 0x56445218.
//
// Solidity: function totalStakedPerLevel(uint8 ) view returns(uint256)
func (_Staking *StakingCallerSession) TotalStakedPerLevel(arg0 uint8) (*big.Int, error) {
	return _Staking.Contract.TotalStakedPerLevel(&_Staking.CallOpts, arg0)
}

// UserComboStatus is a free data retrieval call binding the contract method 0xfc450f37.
//
// Solidity: function userComboStatus(address , uint8 ) view returns(uint8 level, uint256 count, uint256 effectiveFrom, uint256 bonus, bool isPending)
func (_Staking *StakingCaller) UserComboStatus(opts *bind.CallOpts, arg0 common.Address, arg1 uint8) (struct {
	Level         uint8
	Count         *big.Int
	EffectiveFrom *big.Int
	Bonus         *big.Int
	IsPending     bool
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "userComboStatus", arg0, arg1)

	outstruct := new(struct {
		Level         uint8
		Count         *big.Int
		EffectiveFrom *big.Int
		Bonus         *big.Int
		IsPending     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Level = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Count = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EffectiveFrom = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Bonus = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.IsPending = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// UserComboStatus is a free data retrieval call binding the contract method 0xfc450f37.
//
// Solidity: function userComboStatus(address , uint8 ) view returns(uint8 level, uint256 count, uint256 effectiveFrom, uint256 bonus, bool isPending)
func (_Staking *StakingSession) UserComboStatus(arg0 common.Address, arg1 uint8) (struct {
	Level         uint8
	Count         *big.Int
	EffectiveFrom *big.Int
	Bonus         *big.Int
	IsPending     bool
}, error) {
	return _Staking.Contract.UserComboStatus(&_Staking.CallOpts, arg0, arg1)
}

// UserComboStatus is a free data retrieval call binding the contract method 0xfc450f37.
//
// Solidity: function userComboStatus(address , uint8 ) view returns(uint8 level, uint256 count, uint256 effectiveFrom, uint256 bonus, bool isPending)
func (_Staking *StakingCallerSession) UserComboStatus(arg0 common.Address, arg1 uint8) (struct {
	Level         uint8
	Count         *big.Int
	EffectiveFrom *big.Int
	Bonus         *big.Int
	IsPending     bool
}, error) {
	return _Staking.Contract.UserComboStatus(&_Staking.CallOpts, arg0, arg1)
}

// UserLevelCounts is a free data retrieval call binding the contract method 0x71ce85ad.
//
// Solidity: function userLevelCounts(address , uint8 ) view returns(uint256)
func (_Staking *StakingCaller) UserLevelCounts(opts *bind.CallOpts, arg0 common.Address, arg1 uint8) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "userLevelCounts", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserLevelCounts is a free data retrieval call binding the contract method 0x71ce85ad.
//
// Solidity: function userLevelCounts(address , uint8 ) view returns(uint256)
func (_Staking *StakingSession) UserLevelCounts(arg0 common.Address, arg1 uint8) (*big.Int, error) {
	return _Staking.Contract.UserLevelCounts(&_Staking.CallOpts, arg0, arg1)
}

// UserLevelCounts is a free data retrieval call binding the contract method 0x71ce85ad.
//
// Solidity: function userLevelCounts(address , uint8 ) view returns(uint256)
func (_Staking *StakingCallerSession) UserLevelCounts(arg0 common.Address, arg1 uint8) (*big.Int, error) {
	return _Staking.Contract.UserLevelCounts(&_Staking.CallOpts, arg0, arg1)
}

// UserStakes is a free data retrieval call binding the contract method 0xb5d5b5fa.
//
// Solidity: function userStakes(address , uint256 ) view returns(uint256)
func (_Staking *StakingCaller) UserStakes(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "userStakes", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserStakes is a free data retrieval call binding the contract method 0xb5d5b5fa.
//
// Solidity: function userStakes(address , uint256 ) view returns(uint256)
func (_Staking *StakingSession) UserStakes(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Staking.Contract.UserStakes(&_Staking.CallOpts, arg0, arg1)
}

// UserStakes is a free data retrieval call binding the contract method 0xb5d5b5fa.
//
// Solidity: function userStakes(address , uint256 ) view returns(uint256)
func (_Staking *StakingCallerSession) UserStakes(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Staking.Contract.UserStakes(&_Staking.CallOpts, arg0, arg1)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Staking *StakingCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Staking *StakingSession) Version() (string, error) {
	return _Staking.Contract.Version(&_Staking.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Staking *StakingCallerSession) Version() (string, error) {
	return _Staking.Contract.Version(&_Staking.CallOpts)
}

// BatchClaimRewards is a paid mutator transaction binding the contract method 0x2ac8eaff.
//
// Solidity: function batchClaimRewards(address userAddress, uint256[] tokenIds) returns()
func (_Staking *StakingTransactor) BatchClaimRewards(opts *bind.TransactOpts, userAddress common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "batchClaimRewards", userAddress, tokenIds)
}

// BatchClaimRewards is a paid mutator transaction binding the contract method 0x2ac8eaff.
//
// Solidity: function batchClaimRewards(address userAddress, uint256[] tokenIds) returns()
func (_Staking *StakingSession) BatchClaimRewards(userAddress common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Staking.Contract.BatchClaimRewards(&_Staking.TransactOpts, userAddress, tokenIds)
}

// BatchClaimRewards is a paid mutator transaction binding the contract method 0x2ac8eaff.
//
// Solidity: function batchClaimRewards(address userAddress, uint256[] tokenIds) returns()
func (_Staking *StakingTransactorSession) BatchClaimRewards(userAddress common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Staking.Contract.BatchClaimRewards(&_Staking.TransactOpts, userAddress, tokenIds)
}

// BatchStake is a paid mutator transaction binding the contract method 0x717cee0b.
//
// Solidity: function batchStake(address userAddress, uint256[] tokenIds) returns()
func (_Staking *StakingTransactor) BatchStake(opts *bind.TransactOpts, userAddress common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "batchStake", userAddress, tokenIds)
}

// BatchStake is a paid mutator transaction binding the contract method 0x717cee0b.
//
// Solidity: function batchStake(address userAddress, uint256[] tokenIds) returns()
func (_Staking *StakingSession) BatchStake(userAddress common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Staking.Contract.BatchStake(&_Staking.TransactOpts, userAddress, tokenIds)
}

// BatchStake is a paid mutator transaction binding the contract method 0x717cee0b.
//
// Solidity: function batchStake(address userAddress, uint256[] tokenIds) returns()
func (_Staking *StakingTransactorSession) BatchStake(userAddress common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Staking.Contract.BatchStake(&_Staking.TransactOpts, userAddress, tokenIds)
}

// BatchUnstake is a paid mutator transaction binding the contract method 0x961de0d7.
//
// Solidity: function batchUnstake(address userAddress, uint256[] tokenIds) returns()
func (_Staking *StakingTransactor) BatchUnstake(opts *bind.TransactOpts, userAddress common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "batchUnstake", userAddress, tokenIds)
}

// BatchUnstake is a paid mutator transaction binding the contract method 0x961de0d7.
//
// Solidity: function batchUnstake(address userAddress, uint256[] tokenIds) returns()
func (_Staking *StakingSession) BatchUnstake(userAddress common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Staking.Contract.BatchUnstake(&_Staking.TransactOpts, userAddress, tokenIds)
}

// BatchUnstake is a paid mutator transaction binding the contract method 0x961de0d7.
//
// Solidity: function batchUnstake(address userAddress, uint256[] tokenIds) returns()
func (_Staking *StakingTransactorSession) BatchUnstake(userAddress common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Staking.Contract.BatchUnstake(&_Staking.TransactOpts, userAddress, tokenIds)
}

// DisableTestMode is a paid mutator transaction binding the contract method 0x9ddca0c8.
//
// Solidity: function disableTestMode() returns()
func (_Staking *StakingTransactor) DisableTestMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "disableTestMode")
}

// DisableTestMode is a paid mutator transaction binding the contract method 0x9ddca0c8.
//
// Solidity: function disableTestMode() returns()
func (_Staking *StakingSession) DisableTestMode() (*types.Transaction, error) {
	return _Staking.Contract.DisableTestMode(&_Staking.TransactOpts)
}

// DisableTestMode is a paid mutator transaction binding the contract method 0x9ddca0c8.
//
// Solidity: function disableTestMode() returns()
func (_Staking *StakingTransactorSession) DisableTestMode() (*types.Transaction, error) {
	return _Staking.Contract.DisableTestMode(&_Staking.TransactOpts)
}

// EnableTestMode is a paid mutator transaction binding the contract method 0x7dc21e65.
//
// Solidity: function enableTestMode(uint256 initialTimestamp) returns()
func (_Staking *StakingTransactor) EnableTestMode(opts *bind.TransactOpts, initialTimestamp *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "enableTestMode", initialTimestamp)
}

// EnableTestMode is a paid mutator transaction binding the contract method 0x7dc21e65.
//
// Solidity: function enableTestMode(uint256 initialTimestamp) returns()
func (_Staking *StakingSession) EnableTestMode(initialTimestamp *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.EnableTestMode(&_Staking.TransactOpts, initialTimestamp)
}

// EnableTestMode is a paid mutator transaction binding the contract method 0x7dc21e65.
//
// Solidity: function enableTestMode(uint256 initialTimestamp) returns()
func (_Staking *StakingTransactorSession) EnableTestMode(initialTimestamp *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.EnableTestMode(&_Staking.TransactOpts, initialTimestamp)
}

// FastForwardDays is a paid mutator transaction binding the contract method 0xb3f1cfed.
//
// Solidity: function fastForwardDays(uint256 days_) returns()
func (_Staking *StakingTransactor) FastForwardDays(opts *bind.TransactOpts, days_ *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "fastForwardDays", days_)
}

// FastForwardDays is a paid mutator transaction binding the contract method 0xb3f1cfed.
//
// Solidity: function fastForwardDays(uint256 days_) returns()
func (_Staking *StakingSession) FastForwardDays(days_ *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.FastForwardDays(&_Staking.TransactOpts, days_)
}

// FastForwardDays is a paid mutator transaction binding the contract method 0xb3f1cfed.
//
// Solidity: function fastForwardDays(uint256 days_) returns()
func (_Staking *StakingTransactorSession) FastForwardDays(days_ *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.FastForwardDays(&_Staking.TransactOpts, days_)
}

// FastForwardMinutes is a paid mutator transaction binding the contract method 0x7783ea07.
//
// Solidity: function fastForwardMinutes(uint256 minutes_) returns()
func (_Staking *StakingTransactor) FastForwardMinutes(opts *bind.TransactOpts, minutes_ *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "fastForwardMinutes", minutes_)
}

// FastForwardMinutes is a paid mutator transaction binding the contract method 0x7783ea07.
//
// Solidity: function fastForwardMinutes(uint256 minutes_) returns()
func (_Staking *StakingSession) FastForwardMinutes(minutes_ *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.FastForwardMinutes(&_Staking.TransactOpts, minutes_)
}

// FastForwardMinutes is a paid mutator transaction binding the contract method 0x7783ea07.
//
// Solidity: function fastForwardMinutes(uint256 minutes_) returns()
func (_Staking *StakingTransactorSession) FastForwardMinutes(minutes_ *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.FastForwardMinutes(&_Staking.TransactOpts, minutes_)
}

// FastForwardTime is a paid mutator transaction binding the contract method 0xcd0e0e80.
//
// Solidity: function fastForwardTime(uint256 seconds_) returns()
func (_Staking *StakingTransactor) FastForwardTime(opts *bind.TransactOpts, seconds_ *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "fastForwardTime", seconds_)
}

// FastForwardTime is a paid mutator transaction binding the contract method 0xcd0e0e80.
//
// Solidity: function fastForwardTime(uint256 seconds_) returns()
func (_Staking *StakingSession) FastForwardTime(seconds_ *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.FastForwardTime(&_Staking.TransactOpts, seconds_)
}

// FastForwardTime is a paid mutator transaction binding the contract method 0xcd0e0e80.
//
// Solidity: function fastForwardTime(uint256 seconds_) returns()
func (_Staking *StakingTransactorSession) FastForwardTime(seconds_ *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.FastForwardTime(&_Staking.TransactOpts, seconds_)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _cpnftContract, address _CPPTokenContract, address _accountManagerContract, address _configContract, address _owner) returns()
func (_Staking *StakingTransactor) Initialize(opts *bind.TransactOpts, _cpnftContract common.Address, _CPPTokenContract common.Address, _accountManagerContract common.Address, _configContract common.Address, _owner common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "initialize", _cpnftContract, _CPPTokenContract, _accountManagerContract, _configContract, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _cpnftContract, address _CPPTokenContract, address _accountManagerContract, address _configContract, address _owner) returns()
func (_Staking *StakingSession) Initialize(_cpnftContract common.Address, _CPPTokenContract common.Address, _accountManagerContract common.Address, _configContract common.Address, _owner common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Initialize(&_Staking.TransactOpts, _cpnftContract, _CPPTokenContract, _accountManagerContract, _configContract, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _cpnftContract, address _CPPTokenContract, address _accountManagerContract, address _configContract, address _owner) returns()
func (_Staking *StakingTransactorSession) Initialize(_cpnftContract common.Address, _CPPTokenContract common.Address, _accountManagerContract common.Address, _configContract common.Address, _owner common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Initialize(&_Staking.TransactOpts, _cpnftContract, _CPPTokenContract, _accountManagerContract, _configContract, _owner)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Staking *StakingTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Staking *StakingSession) Pause() (*types.Transaction, error) {
	return _Staking.Contract.Pause(&_Staking.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Staking *StakingTransactorSession) Pause() (*types.Transaction, error) {
	return _Staking.Contract.Pause(&_Staking.TransactOpts)
}

// RecordHistoricalAdjustment is a paid mutator transaction binding the contract method 0x4278efde.
//
// Solidity: function recordHistoricalAdjustment() returns()
func (_Staking *StakingTransactor) RecordHistoricalAdjustment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "recordHistoricalAdjustment")
}

// RecordHistoricalAdjustment is a paid mutator transaction binding the contract method 0x4278efde.
//
// Solidity: function recordHistoricalAdjustment() returns()
func (_Staking *StakingSession) RecordHistoricalAdjustment() (*types.Transaction, error) {
	return _Staking.Contract.RecordHistoricalAdjustment(&_Staking.TransactOpts)
}

// RecordHistoricalAdjustment is a paid mutator transaction binding the contract method 0x4278efde.
//
// Solidity: function recordHistoricalAdjustment() returns()
func (_Staking *StakingTransactorSession) RecordHistoricalAdjustment() (*types.Transaction, error) {
	return _Staking.Contract.RecordHistoricalAdjustment(&_Staking.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingSession) RenounceOwnership() (*types.Transaction, error) {
	return _Staking.Contract.RenounceOwnership(&_Staking.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Staking.Contract.RenounceOwnership(&_Staking.TransactOpts)
}

// SetTestTimestamp is a paid mutator transaction binding the contract method 0x8b0a49a7.
//
// Solidity: function setTestTimestamp(uint256 timestamp) returns()
func (_Staking *StakingTransactor) SetTestTimestamp(opts *bind.TransactOpts, timestamp *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setTestTimestamp", timestamp)
}

// SetTestTimestamp is a paid mutator transaction binding the contract method 0x8b0a49a7.
//
// Solidity: function setTestTimestamp(uint256 timestamp) returns()
func (_Staking *StakingSession) SetTestTimestamp(timestamp *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetTestTimestamp(&_Staking.TransactOpts, timestamp)
}

// SetTestTimestamp is a paid mutator transaction binding the contract method 0x8b0a49a7.
//
// Solidity: function setTestTimestamp(uint256 timestamp) returns()
func (_Staking *StakingTransactorSession) SetTestTimestamp(timestamp *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetTestTimestamp(&_Staking.TransactOpts, timestamp)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Staking.Contract.TransferOwnership(&_Staking.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Staking.Contract.TransferOwnership(&_Staking.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Staking *StakingTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Staking *StakingSession) Unpause() (*types.Transaction, error) {
	return _Staking.Contract.Unpause(&_Staking.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Staking *StakingTransactorSession) Unpause() (*types.Transaction, error) {
	return _Staking.Contract.Unpause(&_Staking.TransactOpts)
}

// UpdateContracts is a paid mutator transaction binding the contract method 0x0b751279.
//
// Solidity: function updateContracts(address _cpnftContract, address _CPPTokenContract, address _accountManagerContract, address _configContract) returns()
func (_Staking *StakingTransactor) UpdateContracts(opts *bind.TransactOpts, _cpnftContract common.Address, _CPPTokenContract common.Address, _accountManagerContract common.Address, _configContract common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "updateContracts", _cpnftContract, _CPPTokenContract, _accountManagerContract, _configContract)
}

// UpdateContracts is a paid mutator transaction binding the contract method 0x0b751279.
//
// Solidity: function updateContracts(address _cpnftContract, address _CPPTokenContract, address _accountManagerContract, address _configContract) returns()
func (_Staking *StakingSession) UpdateContracts(_cpnftContract common.Address, _CPPTokenContract common.Address, _accountManagerContract common.Address, _configContract common.Address) (*types.Transaction, error) {
	return _Staking.Contract.UpdateContracts(&_Staking.TransactOpts, _cpnftContract, _CPPTokenContract, _accountManagerContract, _configContract)
}

// UpdateContracts is a paid mutator transaction binding the contract method 0x0b751279.
//
// Solidity: function updateContracts(address _cpnftContract, address _CPPTokenContract, address _accountManagerContract, address _configContract) returns()
func (_Staking *StakingTransactorSession) UpdateContracts(_cpnftContract common.Address, _CPPTokenContract common.Address, _accountManagerContract common.Address, _configContract common.Address) (*types.Transaction, error) {
	return _Staking.Contract.UpdateContracts(&_Staking.TransactOpts, _cpnftContract, _CPPTokenContract, _accountManagerContract, _configContract)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Staking *StakingTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Staking *StakingSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Staking.Contract.UpgradeToAndCall(&_Staking.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Staking *StakingTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Staking.Contract.UpgradeToAndCall(&_Staking.TransactOpts, newImplementation, data)
}

// StakingBatchRewardsClaimedIterator is returned from FilterBatchRewardsClaimed and is used to iterate over the raw logs and unpacked data for BatchRewardsClaimed events raised by the Staking contract.
type StakingBatchRewardsClaimedIterator struct {
	Event *StakingBatchRewardsClaimed // Event containing the contract specifics and raw log

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
func (it *StakingBatchRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingBatchRewardsClaimed)
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
		it.Event = new(StakingBatchRewardsClaimed)
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
func (it *StakingBatchRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingBatchRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingBatchRewardsClaimed represents a BatchRewardsClaimed event raised by the Staking contract.
type StakingBatchRewardsClaimed struct {
	User        common.Address
	TokenCount  *big.Int
	TotalAmount *big.Int
	Timestamp   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchRewardsClaimed is a free log retrieval operation binding the contract event 0xf1f6dd2e19164e6e054315354946fd50dccfbcff365bc9ebfaf1f7a0dffa492f.
//
// Solidity: event BatchRewardsClaimed(address indexed user, uint256 tokenCount, uint256 totalAmount, uint256 timestamp)
func (_Staking *StakingFilterer) FilterBatchRewardsClaimed(opts *bind.FilterOpts, user []common.Address) (*StakingBatchRewardsClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "BatchRewardsClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return &StakingBatchRewardsClaimedIterator{contract: _Staking.contract, event: "BatchRewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchBatchRewardsClaimed is a free log subscription operation binding the contract event 0xf1f6dd2e19164e6e054315354946fd50dccfbcff365bc9ebfaf1f7a0dffa492f.
//
// Solidity: event BatchRewardsClaimed(address indexed user, uint256 tokenCount, uint256 totalAmount, uint256 timestamp)
func (_Staking *StakingFilterer) WatchBatchRewardsClaimed(opts *bind.WatchOpts, sink chan<- *StakingBatchRewardsClaimed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "BatchRewardsClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingBatchRewardsClaimed)
				if err := _Staking.contract.UnpackLog(event, "BatchRewardsClaimed", log); err != nil {
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

// ParseBatchRewardsClaimed is a log parse operation binding the contract event 0xf1f6dd2e19164e6e054315354946fd50dccfbcff365bc9ebfaf1f7a0dffa492f.
//
// Solidity: event BatchRewardsClaimed(address indexed user, uint256 tokenCount, uint256 totalAmount, uint256 timestamp)
func (_Staking *StakingFilterer) ParseBatchRewardsClaimed(log types.Log) (*StakingBatchRewardsClaimed, error) {
	event := new(StakingBatchRewardsClaimed)
	if err := _Staking.contract.UnpackLog(event, "BatchRewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingBatchStakedIterator is returned from FilterBatchStaked and is used to iterate over the raw logs and unpacked data for BatchStaked events raised by the Staking contract.
type StakingBatchStakedIterator struct {
	Event *StakingBatchStaked // Event containing the contract specifics and raw log

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
func (it *StakingBatchStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingBatchStaked)
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
		it.Event = new(StakingBatchStaked)
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
func (it *StakingBatchStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingBatchStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingBatchStaked represents a BatchStaked event raised by the Staking contract.
type StakingBatchStaked struct {
	User      common.Address
	TokenIds  []*big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBatchStaked is a free log retrieval operation binding the contract event 0xb4de6d90ecb1f810530cdbe96c79478f24e8a22c1dd3090bc856c9ef6b7a1f42.
//
// Solidity: event BatchStaked(address indexed user, uint256[] tokenIds, uint256 timestamp)
func (_Staking *StakingFilterer) FilterBatchStaked(opts *bind.FilterOpts, user []common.Address) (*StakingBatchStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "BatchStaked", userRule)
	if err != nil {
		return nil, err
	}
	return &StakingBatchStakedIterator{contract: _Staking.contract, event: "BatchStaked", logs: logs, sub: sub}, nil
}

// WatchBatchStaked is a free log subscription operation binding the contract event 0xb4de6d90ecb1f810530cdbe96c79478f24e8a22c1dd3090bc856c9ef6b7a1f42.
//
// Solidity: event BatchStaked(address indexed user, uint256[] tokenIds, uint256 timestamp)
func (_Staking *StakingFilterer) WatchBatchStaked(opts *bind.WatchOpts, sink chan<- *StakingBatchStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "BatchStaked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingBatchStaked)
				if err := _Staking.contract.UnpackLog(event, "BatchStaked", log); err != nil {
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

// ParseBatchStaked is a log parse operation binding the contract event 0xb4de6d90ecb1f810530cdbe96c79478f24e8a22c1dd3090bc856c9ef6b7a1f42.
//
// Solidity: event BatchStaked(address indexed user, uint256[] tokenIds, uint256 timestamp)
func (_Staking *StakingFilterer) ParseBatchStaked(log types.Log) (*StakingBatchStaked, error) {
	event := new(StakingBatchStaked)
	if err := _Staking.contract.UnpackLog(event, "BatchStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingBatchUnstakedIterator is returned from FilterBatchUnstaked and is used to iterate over the raw logs and unpacked data for BatchUnstaked events raised by the Staking contract.
type StakingBatchUnstakedIterator struct {
	Event *StakingBatchUnstaked // Event containing the contract specifics and raw log

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
func (it *StakingBatchUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingBatchUnstaked)
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
		it.Event = new(StakingBatchUnstaked)
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
func (it *StakingBatchUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingBatchUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingBatchUnstaked represents a BatchUnstaked event raised by the Staking contract.
type StakingBatchUnstaked struct {
	User         common.Address
	TokenIds     []*big.Int
	TotalRewards *big.Int
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBatchUnstaked is a free log retrieval operation binding the contract event 0xa685df6db7e38bbbe7890e285450dcdc4f0a30a30a051c9a04f4c46cd5a03641.
//
// Solidity: event BatchUnstaked(address indexed user, uint256[] tokenIds, uint256 totalRewards, uint256 timestamp)
func (_Staking *StakingFilterer) FilterBatchUnstaked(opts *bind.FilterOpts, user []common.Address) (*StakingBatchUnstakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "BatchUnstaked", userRule)
	if err != nil {
		return nil, err
	}
	return &StakingBatchUnstakedIterator{contract: _Staking.contract, event: "BatchUnstaked", logs: logs, sub: sub}, nil
}

// WatchBatchUnstaked is a free log subscription operation binding the contract event 0xa685df6db7e38bbbe7890e285450dcdc4f0a30a30a051c9a04f4c46cd5a03641.
//
// Solidity: event BatchUnstaked(address indexed user, uint256[] tokenIds, uint256 totalRewards, uint256 timestamp)
func (_Staking *StakingFilterer) WatchBatchUnstaked(opts *bind.WatchOpts, sink chan<- *StakingBatchUnstaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "BatchUnstaked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingBatchUnstaked)
				if err := _Staking.contract.UnpackLog(event, "BatchUnstaked", log); err != nil {
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

// ParseBatchUnstaked is a log parse operation binding the contract event 0xa685df6db7e38bbbe7890e285450dcdc4f0a30a30a051c9a04f4c46cd5a03641.
//
// Solidity: event BatchUnstaked(address indexed user, uint256[] tokenIds, uint256 totalRewards, uint256 timestamp)
func (_Staking *StakingFilterer) ParseBatchUnstaked(log types.Log) (*StakingBatchUnstaked, error) {
	event := new(StakingBatchUnstaked)
	if err := _Staking.contract.UnpackLog(event, "BatchUnstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Staking contract.
type StakingInitializedIterator struct {
	Event *StakingInitialized // Event containing the contract specifics and raw log

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
func (it *StakingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingInitialized)
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
		it.Event = new(StakingInitialized)
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
func (it *StakingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingInitialized represents a Initialized event raised by the Staking contract.
type StakingInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Staking *StakingFilterer) FilterInitialized(opts *bind.FilterOpts) (*StakingInitializedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StakingInitializedIterator{contract: _Staking.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Staking *StakingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StakingInitialized) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingInitialized)
				if err := _Staking.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Staking *StakingFilterer) ParseInitialized(log types.Log) (*StakingInitialized, error) {
	event := new(StakingInitialized)
	if err := _Staking.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingNFTStakedIterator is returned from FilterNFTStaked and is used to iterate over the raw logs and unpacked data for NFTStaked events raised by the Staking contract.
type StakingNFTStakedIterator struct {
	Event *StakingNFTStaked // Event containing the contract specifics and raw log

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
func (it *StakingNFTStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingNFTStaked)
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
		it.Event = new(StakingNFTStaked)
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
func (it *StakingNFTStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingNFTStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingNFTStaked represents a NFTStaked event raised by the Staking contract.
type StakingNFTStaked struct {
	User      common.Address
	TokenId   *big.Int
	Level     uint8
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNFTStaked is a free log retrieval operation binding the contract event 0xb7f42a117a7de13499e08cdb12729b20c510b7f623fc79fec9e8bfbe1a024487.
//
// Solidity: event NFTStaked(address indexed user, uint256 indexed tokenId, uint8 level, uint256 timestamp)
func (_Staking *StakingFilterer) FilterNFTStaked(opts *bind.FilterOpts, user []common.Address, tokenId []*big.Int) (*StakingNFTStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "NFTStaked", userRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &StakingNFTStakedIterator{contract: _Staking.contract, event: "NFTStaked", logs: logs, sub: sub}, nil
}

// WatchNFTStaked is a free log subscription operation binding the contract event 0xb7f42a117a7de13499e08cdb12729b20c510b7f623fc79fec9e8bfbe1a024487.
//
// Solidity: event NFTStaked(address indexed user, uint256 indexed tokenId, uint8 level, uint256 timestamp)
func (_Staking *StakingFilterer) WatchNFTStaked(opts *bind.WatchOpts, sink chan<- *StakingNFTStaked, user []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "NFTStaked", userRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingNFTStaked)
				if err := _Staking.contract.UnpackLog(event, "NFTStaked", log); err != nil {
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

// ParseNFTStaked is a log parse operation binding the contract event 0xb7f42a117a7de13499e08cdb12729b20c510b7f623fc79fec9e8bfbe1a024487.
//
// Solidity: event NFTStaked(address indexed user, uint256 indexed tokenId, uint8 level, uint256 timestamp)
func (_Staking *StakingFilterer) ParseNFTStaked(log types.Log) (*StakingNFTStaked, error) {
	event := new(StakingNFTStaked)
	if err := _Staking.contract.UnpackLog(event, "NFTStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingNFTUnstakedIterator is returned from FilterNFTUnstaked and is used to iterate over the raw logs and unpacked data for NFTUnstaked events raised by the Staking contract.
type StakingNFTUnstakedIterator struct {
	Event *StakingNFTUnstaked // Event containing the contract specifics and raw log

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
func (it *StakingNFTUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingNFTUnstaked)
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
		it.Event = new(StakingNFTUnstaked)
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
func (it *StakingNFTUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingNFTUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingNFTUnstaked represents a NFTUnstaked event raised by the Staking contract.
type StakingNFTUnstaked struct {
	User      common.Address
	TokenId   *big.Int
	Rewards   *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNFTUnstaked is a free log retrieval operation binding the contract event 0x84bcc89bc1b3cb66a2b22a282fa6e1bb013db80ebce0e9c70f9beba416ac2b70.
//
// Solidity: event NFTUnstaked(address indexed user, uint256 indexed tokenId, uint256 rewards, uint256 timestamp)
func (_Staking *StakingFilterer) FilterNFTUnstaked(opts *bind.FilterOpts, user []common.Address, tokenId []*big.Int) (*StakingNFTUnstakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "NFTUnstaked", userRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &StakingNFTUnstakedIterator{contract: _Staking.contract, event: "NFTUnstaked", logs: logs, sub: sub}, nil
}

// WatchNFTUnstaked is a free log subscription operation binding the contract event 0x84bcc89bc1b3cb66a2b22a282fa6e1bb013db80ebce0e9c70f9beba416ac2b70.
//
// Solidity: event NFTUnstaked(address indexed user, uint256 indexed tokenId, uint256 rewards, uint256 timestamp)
func (_Staking *StakingFilterer) WatchNFTUnstaked(opts *bind.WatchOpts, sink chan<- *StakingNFTUnstaked, user []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "NFTUnstaked", userRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingNFTUnstaked)
				if err := _Staking.contract.UnpackLog(event, "NFTUnstaked", log); err != nil {
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

// ParseNFTUnstaked is a log parse operation binding the contract event 0x84bcc89bc1b3cb66a2b22a282fa6e1bb013db80ebce0e9c70f9beba416ac2b70.
//
// Solidity: event NFTUnstaked(address indexed user, uint256 indexed tokenId, uint256 rewards, uint256 timestamp)
func (_Staking *StakingFilterer) ParseNFTUnstaked(log types.Log) (*StakingNFTUnstaked, error) {
	event := new(StakingNFTUnstaked)
	if err := _Staking.contract.UnpackLog(event, "NFTUnstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Staking contract.
type StakingOwnershipTransferredIterator struct {
	Event *StakingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingOwnershipTransferred)
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
		it.Event = new(StakingOwnershipTransferred)
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
func (it *StakingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingOwnershipTransferred represents a OwnershipTransferred event raised by the Staking contract.
type StakingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Staking *StakingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingOwnershipTransferredIterator{contract: _Staking.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Staking *StakingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingOwnershipTransferred)
				if err := _Staking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Staking *StakingFilterer) ParseOwnershipTransferred(log types.Log) (*StakingOwnershipTransferred, error) {
	event := new(StakingOwnershipTransferred)
	if err := _Staking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Staking contract.
type StakingPausedIterator struct {
	Event *StakingPaused // Event containing the contract specifics and raw log

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
func (it *StakingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPaused)
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
		it.Event = new(StakingPaused)
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
func (it *StakingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPaused represents a Paused event raised by the Staking contract.
type StakingPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Staking *StakingFilterer) FilterPaused(opts *bind.FilterOpts) (*StakingPausedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &StakingPausedIterator{contract: _Staking.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Staking *StakingFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StakingPaused) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPaused)
				if err := _Staking.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Staking *StakingFilterer) ParsePaused(log types.Log) (*StakingPaused, error) {
	event := new(StakingPaused)
	if err := _Staking.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPlatformStatsUpdatedIterator is returned from FilterPlatformStatsUpdated and is used to iterate over the raw logs and unpacked data for PlatformStatsUpdated events raised by the Staking contract.
type StakingPlatformStatsUpdatedIterator struct {
	Event *StakingPlatformStatsUpdated // Event containing the contract specifics and raw log

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
func (it *StakingPlatformStatsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPlatformStatsUpdated)
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
		it.Event = new(StakingPlatformStatsUpdated)
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
func (it *StakingPlatformStatsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPlatformStatsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPlatformStatsUpdated represents a PlatformStatsUpdated event raised by the Staking contract.
type StakingPlatformStatsUpdated struct {
	Level  uint8
	Staked *big.Int
	Supply *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPlatformStatsUpdated is a free log retrieval operation binding the contract event 0xc331e8057295fc53acc3dbefbf32a694fd6c7a9ead970e1763608d7b5c5545bf.
//
// Solidity: event PlatformStatsUpdated(uint8 level, uint256 staked, uint256 supply)
func (_Staking *StakingFilterer) FilterPlatformStatsUpdated(opts *bind.FilterOpts) (*StakingPlatformStatsUpdatedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "PlatformStatsUpdated")
	if err != nil {
		return nil, err
	}
	return &StakingPlatformStatsUpdatedIterator{contract: _Staking.contract, event: "PlatformStatsUpdated", logs: logs, sub: sub}, nil
}

// WatchPlatformStatsUpdated is a free log subscription operation binding the contract event 0xc331e8057295fc53acc3dbefbf32a694fd6c7a9ead970e1763608d7b5c5545bf.
//
// Solidity: event PlatformStatsUpdated(uint8 level, uint256 staked, uint256 supply)
func (_Staking *StakingFilterer) WatchPlatformStatsUpdated(opts *bind.WatchOpts, sink chan<- *StakingPlatformStatsUpdated) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "PlatformStatsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPlatformStatsUpdated)
				if err := _Staking.contract.UnpackLog(event, "PlatformStatsUpdated", log); err != nil {
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

// ParsePlatformStatsUpdated is a log parse operation binding the contract event 0xc331e8057295fc53acc3dbefbf32a694fd6c7a9ead970e1763608d7b5c5545bf.
//
// Solidity: event PlatformStatsUpdated(uint8 level, uint256 staked, uint256 supply)
func (_Staking *StakingFilterer) ParsePlatformStatsUpdated(log types.Log) (*StakingPlatformStatsUpdated, error) {
	event := new(StakingPlatformStatsUpdated)
	if err := _Staking.contract.UnpackLog(event, "PlatformStatsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRewardsClaimedIterator is returned from FilterRewardsClaimed and is used to iterate over the raw logs and unpacked data for RewardsClaimed events raised by the Staking contract.
type StakingRewardsClaimedIterator struct {
	Event *StakingRewardsClaimed // Event containing the contract specifics and raw log

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
func (it *StakingRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRewardsClaimed)
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
		it.Event = new(StakingRewardsClaimed)
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
func (it *StakingRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRewardsClaimed represents a RewardsClaimed event raised by the Staking contract.
type StakingRewardsClaimed struct {
	User      common.Address
	TokenId   *big.Int
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewardsClaimed is a free log retrieval operation binding the contract event 0x56253d287efacdb2c4cd76dd03624a4821c1ce721d1152e8f5f5718f6087c9bf.
//
// Solidity: event RewardsClaimed(address indexed user, uint256 indexed tokenId, uint256 amount, uint256 timestamp)
func (_Staking *StakingFilterer) FilterRewardsClaimed(opts *bind.FilterOpts, user []common.Address, tokenId []*big.Int) (*StakingRewardsClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "RewardsClaimed", userRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &StakingRewardsClaimedIterator{contract: _Staking.contract, event: "RewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardsClaimed is a free log subscription operation binding the contract event 0x56253d287efacdb2c4cd76dd03624a4821c1ce721d1152e8f5f5718f6087c9bf.
//
// Solidity: event RewardsClaimed(address indexed user, uint256 indexed tokenId, uint256 amount, uint256 timestamp)
func (_Staking *StakingFilterer) WatchRewardsClaimed(opts *bind.WatchOpts, sink chan<- *StakingRewardsClaimed, user []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "RewardsClaimed", userRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRewardsClaimed)
				if err := _Staking.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
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

// ParseRewardsClaimed is a log parse operation binding the contract event 0x56253d287efacdb2c4cd76dd03624a4821c1ce721d1152e8f5f5718f6087c9bf.
//
// Solidity: event RewardsClaimed(address indexed user, uint256 indexed tokenId, uint256 amount, uint256 timestamp)
func (_Staking *StakingFilterer) ParseRewardsClaimed(log types.Log) (*StakingRewardsClaimed, error) {
	event := new(StakingRewardsClaimed)
	if err := _Staking.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Staking contract.
type StakingUnpausedIterator struct {
	Event *StakingUnpaused // Event containing the contract specifics and raw log

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
func (it *StakingUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingUnpaused)
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
		it.Event = new(StakingUnpaused)
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
func (it *StakingUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingUnpaused represents a Unpaused event raised by the Staking contract.
type StakingUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Staking *StakingFilterer) FilterUnpaused(opts *bind.FilterOpts) (*StakingUnpausedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &StakingUnpausedIterator{contract: _Staking.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Staking *StakingFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StakingUnpaused) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingUnpaused)
				if err := _Staking.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Staking *StakingFilterer) ParseUnpaused(log types.Log) (*StakingUnpaused, error) {
	event := new(StakingUnpaused)
	if err := _Staking.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Staking contract.
type StakingUpgradedIterator struct {
	Event *StakingUpgraded // Event containing the contract specifics and raw log

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
func (it *StakingUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingUpgraded)
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
		it.Event = new(StakingUpgraded)
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
func (it *StakingUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingUpgraded represents a Upgraded event raised by the Staking contract.
type StakingUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Staking *StakingFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*StakingUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &StakingUpgradedIterator{contract: _Staking.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Staking *StakingFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *StakingUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingUpgraded)
				if err := _Staking.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Staking *StakingFilterer) ParseUpgraded(log types.Log) (*StakingUpgraded, error) {
	event := new(StakingUpgraded)
	if err := _Staking.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
