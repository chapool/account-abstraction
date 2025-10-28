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

// StakingReaderStakedNFTInfo is an auto generated low-level Go binding around an user-defined struct.
type StakingReaderStakedNFTInfo struct {
	TokenId             *big.Int
	Level               uint8
	StakingDuration     *big.Int
	PendingRewards      *big.Int
	TotalRewards        *big.Int
	EffectiveMultiplier *big.Int
}

// StakingReaderUserComboSummary is an auto generated low-level Go binding around an user-defined struct.
type StakingReaderUserComboSummary struct {
	CurrentComboCounts [6]*big.Int
	ComboBonus         [6]*big.Int
	NextComboThreshold [6]*big.Int
	HasPendingCombo    [6]bool
}

// StakingReaderUserDailyRewards is an auto generated low-level Go binding around an user-defined struct.
type StakingReaderUserDailyRewards struct {
	TotalBaseReward        *big.Int
	TotalDecayedReward     *big.Int
	TotalComboBonus        *big.Int
	TotalDynamicMultiplier *big.Int
	TotalFinalReward       *big.Int
	TotalBonus             *big.Int
	BaseRewardPerLevel     [6]*big.Int
	FinalRewardPerLevel    [6]*big.Int
	NftCountPerLevel       [6]*big.Int
}

// StakingReaderUserRewardStats is an auto generated low-level Go binding around an user-defined struct.
type StakingReaderUserRewardStats struct {
	TotalHistoricalRewards *big.Int
	TotalPendingRewards    *big.Int
	RewardsPerLevel        [6]*big.Int
	Last24HoursRewards     *big.Int
	AverageDailyRewards    *big.Int
}

// StakingReaderUserStakingSummary is an auto generated low-level Go binding around an user-defined struct.
type StakingReaderUserStakingSummary struct {
	TotalStakedCount         *big.Int
	TotalClaimedRewards      *big.Int
	TotalPendingRewards      *big.Int
	LevelStakingCounts       [6]*big.Int
	LongestStakingDuration   *big.Int
	TotalEffectiveMultiplier *big.Int
}

// StakingStakeInfo is an auto generated low-level Go binding around an user-defined struct.
type StakingStakeInfo struct {
	Owner                  common.Address
	TokenId                *big.Int
	Level                  uint8
	StakeTime              *big.Int
	LastClaimTime          *big.Int
	IsActive               bool
	TotalRewards           *big.Int
	PendingRewards         *big.Int
	ContinuousBonusClaimed bool
}

// StakingReaderMetaData contains all meta data concerning the StakingReader contract.
var StakingReaderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configContract\",\"outputs\":[{\"internalType\":\"contractStakingConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cpnftContract\",\"outputs\":[{\"internalType\":\"contractCPNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBasicConfig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minStakeDays\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"earlyWithdrawPenalty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quarterlyMultiplier\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getComboConfig\",\"outputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"thresholds\",\"type\":\"uint256[3]\"},{\"internalType\":\"uint256[3]\",\"name\":\"bonuses\",\"type\":\"uint256[3]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"}],\"name\":\"getComboInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentBonus\",\"type\":\"uint256\"},{\"internalType\":\"uint256[3]\",\"name\":\"thresholds\",\"type\":\"uint256[3]\"},{\"internalType\":\"uint256[3]\",\"name\":\"bonuses\",\"type\":\"uint256[3]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"}],\"name\":\"getDailyReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLevelStats\",\"outputs\":[{\"internalType\":\"uint256[7]\",\"name\":\"stakedCounts\",\"type\":\"uint256[7]\"},{\"internalType\":\"uint256[7]\",\"name\":\"supplies\",\"type\":\"uint256[7]\"},{\"internalType\":\"uint256[7]\",\"name\":\"stakingRatios\",\"type\":\"uint256[7]\"},{\"internalType\":\"uint256[7]\",\"name\":\"dynamicMultipliers\",\"type\":\"uint256[7]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getNFTDailyRewardBreakdown\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"baseReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decayedReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"comboBonus\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dynamicMultiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalReward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getPendingRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPlatformStats\",\"outputs\":[{\"internalType\":\"uint256[7]\",\"name\":\"staked\",\"type\":\"uint256[7]\"},{\"internalType\":\"uint256[7]\",\"name\":\"supply\",\"type\":\"uint256[7]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getStakeDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"stakeTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastClaimTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"totalRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingRewards\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"continuousBonusClaimed\",\"type\":\"bool\"}],\"internalType\":\"structStaking.StakeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserComboSummary\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256[6]\",\"name\":\"currentComboCounts\",\"type\":\"uint256[6]\"},{\"internalType\":\"uint256[6]\",\"name\":\"comboBonus\",\"type\":\"uint256[6]\"},{\"internalType\":\"uint256[6]\",\"name\":\"nextComboThreshold\",\"type\":\"uint256[6]\"},{\"internalType\":\"bool[6]\",\"name\":\"hasPendingCombo\",\"type\":\"bool[6]\"}],\"internalType\":\"structStakingReader.UserComboSummary\",\"name\":\"summary\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserDailyRewards\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalBaseReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDecayedReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalComboBonus\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDynamicMultiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalFinalReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBonus\",\"type\":\"uint256\"},{\"internalType\":\"uint256[6]\",\"name\":\"baseRewardPerLevel\",\"type\":\"uint256[6]\"},{\"internalType\":\"uint256[6]\",\"name\":\"finalRewardPerLevel\",\"type\":\"uint256[6]\"},{\"internalType\":\"uint256[6]\",\"name\":\"nftCountPerLevel\",\"type\":\"uint256[6]\"}],\"internalType\":\"structStakingReader.UserDailyRewards\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"}],\"name\":\"getUserDailyRewardsByLevel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"baseReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nftCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserRewardStats\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalHistoricalRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPendingRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256[6]\",\"name\":\"rewardsPerLevel\",\"type\":\"uint256[6]\"},{\"internalType\":\"uint256\",\"name\":\"last24HoursRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"averageDailyRewards\",\"type\":\"uint256\"}],\"internalType\":\"structStakingReader.UserRewardStats\",\"name\":\"stats\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getUserStakedNFTs\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"stakingDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"effectiveMultiplier\",\"type\":\"uint256\"}],\"internalType\":\"structStakingReader.StakedNFTInfo[]\",\"name\":\"nfts\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserStakingSummary\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalStakedCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalClaimedRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPendingRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256[6]\",\"name\":\"levelStakingCounts\",\"type\":\"uint256[6]\"},{\"internalType\":\"uint256\",\"name\":\"longestStakingDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalEffectiveMultiplier\",\"type\":\"uint256\"}],\"internalType\":\"structStakingReader.UserStakingSummary\",\"name\":\"summary\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVersions\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"stakingVersion\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"configVersion\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_configContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cpnftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingContract\",\"outputs\":[{\"internalType\":\"contractStaking\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_configContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cpnftContract\",\"type\":\"address\"}],\"name\":\"updateContracts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// StakingReaderABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingReaderMetaData.ABI instead.
var StakingReaderABI = StakingReaderMetaData.ABI

// StakingReader is an auto generated Go binding around an Ethereum contract.
type StakingReader struct {
	StakingReaderCaller     // Read-only binding to the contract
	StakingReaderTransactor // Write-only binding to the contract
	StakingReaderFilterer   // Log filterer for contract events
}

// StakingReaderCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingReaderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingReaderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingReaderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingReaderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingReaderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingReaderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingReaderSession struct {
	Contract     *StakingReader    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingReaderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingReaderCallerSession struct {
	Contract *StakingReaderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StakingReaderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingReaderTransactorSession struct {
	Contract     *StakingReaderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StakingReaderRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingReaderRaw struct {
	Contract *StakingReader // Generic contract binding to access the raw methods on
}

// StakingReaderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingReaderCallerRaw struct {
	Contract *StakingReaderCaller // Generic read-only contract binding to access the raw methods on
}

// StakingReaderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingReaderTransactorRaw struct {
	Contract *StakingReaderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingReader creates a new instance of StakingReader, bound to a specific deployed contract.
func NewStakingReader(address common.Address, backend bind.ContractBackend) (*StakingReader, error) {
	contract, err := bindStakingReader(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingReader{StakingReaderCaller: StakingReaderCaller{contract: contract}, StakingReaderTransactor: StakingReaderTransactor{contract: contract}, StakingReaderFilterer: StakingReaderFilterer{contract: contract}}, nil
}

// NewStakingReaderCaller creates a new read-only instance of StakingReader, bound to a specific deployed contract.
func NewStakingReaderCaller(address common.Address, caller bind.ContractCaller) (*StakingReaderCaller, error) {
	contract, err := bindStakingReader(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingReaderCaller{contract: contract}, nil
}

// NewStakingReaderTransactor creates a new write-only instance of StakingReader, bound to a specific deployed contract.
func NewStakingReaderTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingReaderTransactor, error) {
	contract, err := bindStakingReader(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingReaderTransactor{contract: contract}, nil
}

// NewStakingReaderFilterer creates a new log filterer instance of StakingReader, bound to a specific deployed contract.
func NewStakingReaderFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingReaderFilterer, error) {
	contract, err := bindStakingReader(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingReaderFilterer{contract: contract}, nil
}

// bindStakingReader binds a generic wrapper to an already deployed contract.
func bindStakingReader(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingReaderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingReader *StakingReaderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingReader.Contract.StakingReaderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingReader *StakingReaderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingReader.Contract.StakingReaderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingReader *StakingReaderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingReader.Contract.StakingReaderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingReader *StakingReaderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingReader.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingReader *StakingReaderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingReader.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingReader *StakingReaderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingReader.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_StakingReader *StakingReaderCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StakingReader.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_StakingReader *StakingReaderSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _StakingReader.Contract.UPGRADEINTERFACEVERSION(&_StakingReader.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_StakingReader *StakingReaderCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _StakingReader.Contract.UPGRADEINTERFACEVERSION(&_StakingReader.CallOpts)
}

// ConfigContract is a free data retrieval call binding the contract method 0xbf66a182.
//
// Solidity: function configContract() view returns(address)
func (_StakingReader *StakingReaderCaller) ConfigContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingReader.contract.Call(opts, &out, "configContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConfigContract is a free data retrieval call binding the contract method 0xbf66a182.
//
// Solidity: function configContract() view returns(address)
func (_StakingReader *StakingReaderSession) ConfigContract() (common.Address, error) {
	return _StakingReader.Contract.ConfigContract(&_StakingReader.CallOpts)
}

// ConfigContract is a free data retrieval call binding the contract method 0xbf66a182.
//
// Solidity: function configContract() view returns(address)
func (_StakingReader *StakingReaderCallerSession) ConfigContract() (common.Address, error) {
	return _StakingReader.Contract.ConfigContract(&_StakingReader.CallOpts)
}

// CpnftContract is a free data retrieval call binding the contract method 0xabc99ee4.
//
// Solidity: function cpnftContract() view returns(address)
func (_StakingReader *StakingReaderCaller) CpnftContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingReader.contract.Call(opts, &out, "cpnftContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CpnftContract is a free data retrieval call binding the contract method 0xabc99ee4.
//
// Solidity: function cpnftContract() view returns(address)
func (_StakingReader *StakingReaderSession) CpnftContract() (common.Address, error) {
	return _StakingReader.Contract.CpnftContract(&_StakingReader.CallOpts)
}

// CpnftContract is a free data retrieval call binding the contract method 0xabc99ee4.
//
// Solidity: function cpnftContract() view returns(address)
func (_StakingReader *StakingReaderCallerSession) CpnftContract() (common.Address, error) {
	return _StakingReader.Contract.CpnftContract(&_StakingReader.CallOpts)
}

// GetBasicConfig is a free data retrieval call binding the contract method 0x4fe57a6e.
//
// Solidity: function getBasicConfig() view returns(uint256 minStakeDays, uint256 earlyWithdrawPenalty, uint256 quarterlyMultiplier)
func (_StakingReader *StakingReaderCaller) GetBasicConfig(opts *bind.CallOpts) (struct {
	MinStakeDays         *big.Int
	EarlyWithdrawPenalty *big.Int
	QuarterlyMultiplier  *big.Int
}, error) {
	var out []interface{}
	err := _StakingReader.contract.Call(opts, &out, "getBasicConfig")

	outstruct := new(struct {
		MinStakeDays         *big.Int
		EarlyWithdrawPenalty *big.Int
		QuarterlyMultiplier  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinStakeDays = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EarlyWithdrawPenalty = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.QuarterlyMultiplier = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBasicConfig is a free data retrieval call binding the contract method 0x4fe57a6e.
//
// Solidity: function getBasicConfig() view returns(uint256 minStakeDays, uint256 earlyWithdrawPenalty, uint256 quarterlyMultiplier)
func (_StakingReader *StakingReaderSession) GetBasicConfig() (struct {
	MinStakeDays         *big.Int
	EarlyWithdrawPenalty *big.Int
	QuarterlyMultiplier  *big.Int
}, error) {
	return _StakingReader.Contract.GetBasicConfig(&_StakingReader.CallOpts)
}

// GetBasicConfig is a free data retrieval call binding the contract method 0x4fe57a6e.
//
// Solidity: function getBasicConfig() view returns(uint256 minStakeDays, uint256 earlyWithdrawPenalty, uint256 quarterlyMultiplier)
func (_StakingReader *StakingReaderCallerSession) GetBasicConfig() (struct {
	MinStakeDays         *big.Int
	EarlyWithdrawPenalty *big.Int
	QuarterlyMultiplier  *big.Int
}, error) {
	return _StakingReader.Contract.GetBasicConfig(&_StakingReader.CallOpts)
}

// GetComboConfig is a free data retrieval call binding the contract method 0x8da7a79f.
//
// Solidity: function getComboConfig() view returns(uint256[3] thresholds, uint256[3] bonuses)
func (_StakingReader *StakingReaderCaller) GetComboConfig(opts *bind.CallOpts) (struct {
	Thresholds [3]*big.Int
	Bonuses    [3]*big.Int
}, error) {
	var out []interface{}
	err := _StakingReader.contract.Call(opts, &out, "getComboConfig")

	outstruct := new(struct {
		Thresholds [3]*big.Int
		Bonuses    [3]*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Thresholds = *abi.ConvertType(out[0], new([3]*big.Int)).(*[3]*big.Int)
	outstruct.Bonuses = *abi.ConvertType(out[1], new([3]*big.Int)).(*[3]*big.Int)

	return *outstruct, err

}

// GetComboConfig is a free data retrieval call binding the contract method 0x8da7a79f.
//
// Solidity: function getComboConfig() view returns(uint256[3] thresholds, uint256[3] bonuses)
func (_StakingReader *StakingReaderSession) GetComboConfig() (struct {
	Thresholds [3]*big.Int
	Bonuses    [3]*big.Int
}, error) {
	return _StakingReader.Contract.GetComboConfig(&_StakingReader.CallOpts)
}

// GetComboConfig is a free data retrieval call binding the contract method 0x8da7a79f.
//
// Solidity: function getComboConfig() view returns(uint256[3] thresholds, uint256[3] bonuses)
func (_StakingReader *StakingReaderCallerSession) GetComboConfig() (struct {
	Thresholds [3]*big.Int
	Bonuses    [3]*big.Int
}, error) {
	return _StakingReader.Contract.GetComboConfig(&_StakingReader.CallOpts)
}

// GetComboInfo is a free data retrieval call binding the contract method 0x240dec05.
//
// Solidity: function getComboInfo(address user, uint8 level) view returns(uint256 count, uint256 currentBonus, uint256[3] thresholds, uint256[3] bonuses)
func (_StakingReader *StakingReaderCaller) GetComboInfo(opts *bind.CallOpts, user common.Address, level uint8) (struct {
	Count        *big.Int
	CurrentBonus *big.Int
	Thresholds   [3]*big.Int
	Bonuses      [3]*big.Int
}, error) {
	var out []interface{}
	err := _StakingReader.contract.Call(opts, &out, "getComboInfo", user, level)

	outstruct := new(struct {
		Count        *big.Int
		CurrentBonus *big.Int
		Thresholds   [3]*big.Int
		Bonuses      [3]*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Count = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CurrentBonus = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Thresholds = *abi.ConvertType(out[2], new([3]*big.Int)).(*[3]*big.Int)
	outstruct.Bonuses = *abi.ConvertType(out[3], new([3]*big.Int)).(*[3]*big.Int)

	return *outstruct, err

}

// GetComboInfo is a free data retrieval call binding the contract method 0x240dec05.
//
// Solidity: function getComboInfo(address user, uint8 level) view returns(uint256 count, uint256 currentBonus, uint256[3] thresholds, uint256[3] bonuses)
func (_StakingReader *StakingReaderSession) GetComboInfo(user common.Address, level uint8) (struct {
	Count        *big.Int
	CurrentBonus *big.Int
	Thresholds   [3]*big.Int
	Bonuses      [3]*big.Int
}, error) {
	return _StakingReader.Contract.GetComboInfo(&_StakingReader.CallOpts, user, level)
}

// GetComboInfo is a free data retrieval call binding the contract method 0x240dec05.
//
// Solidity: function getComboInfo(address user, uint8 level) view returns(uint256 count, uint256 currentBonus, uint256[3] thresholds, uint256[3] bonuses)
func (_StakingReader *StakingReaderCallerSession) GetComboInfo(user common.Address, level uint8) (struct {
	Count        *big.Int
	CurrentBonus *big.Int
	Thresholds   [3]*big.Int
	Bonuses      [3]*big.Int
}, error) {
	return _StakingReader.Contract.GetComboInfo(&_StakingReader.CallOpts, user, level)
}

// GetDailyReward is a free data retrieval call binding the contract method 0x13c57ba1.
//
// Solidity: function getDailyReward(uint8 level) view returns(uint256)
func (_StakingReader *StakingReaderCaller) GetDailyReward(opts *bind.CallOpts, level uint8) (*big.Int, error) {
	var out []interface{}
	err := _StakingReader.contract.Call(opts, &out, "getDailyReward", level)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDailyReward is a free data retrieval call binding the contract method 0x13c57ba1.
//
// Solidity: function getDailyReward(uint8 level) view returns(uint256)
func (_StakingReader *StakingReaderSession) GetDailyReward(level uint8) (*big.Int, error) {
	return _StakingReader.Contract.GetDailyReward(&_StakingReader.CallOpts, level)
}

// GetDailyReward is a free data retrieval call binding the contract method 0x13c57ba1.
//
// Solidity: function getDailyReward(uint8 level) view returns(uint256)
func (_StakingReader *StakingReaderCallerSession) GetDailyReward(level uint8) (*big.Int, error) {
	return _StakingReader.Contract.GetDailyReward(&_StakingReader.CallOpts, level)
}

// GetLevelStats is a free data retrieval call binding the contract method 0x83537b46.
//
// Solidity: function getLevelStats() view returns(uint256[7] stakedCounts, uint256[7] supplies, uint256[7] stakingRatios, uint256[7] dynamicMultipliers)
func (_StakingReader *StakingReaderCaller) GetLevelStats(opts *bind.CallOpts) (struct {
	StakedCounts       [7]*big.Int
	Supplies           [7]*big.Int
	StakingRatios      [7]*big.Int
	DynamicMultipliers [7]*big.Int
}, error) {
	var out []interface{}
	err := _StakingReader.contract.Call(opts, &out, "getLevelStats")

	outstruct := new(struct {
		StakedCounts       [7]*big.Int
		Supplies           [7]*big.Int
		StakingRatios      [7]*big.Int
		DynamicMultipliers [7]*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StakedCounts = *abi.ConvertType(out[0], new([7]*big.Int)).(*[7]*big.Int)
	outstruct.Supplies = *abi.ConvertType(out[1], new([7]*big.Int)).(*[7]*big.Int)
	outstruct.StakingRatios = *abi.ConvertType(out[2], new([7]*big.Int)).(*[7]*big.Int)
	outstruct.DynamicMultipliers = *abi.ConvertType(out[3], new([7]*big.Int)).(*[7]*big.Int)

	return *outstruct, err

}

// GetLevelStats is a free data retrieval call binding the contract method 0x83537b46.
//
// Solidity: function getLevelStats() view returns(uint256[7] stakedCounts, uint256[7] supplies, uint256[7] stakingRatios, uint256[7] dynamicMultipliers)
func (_StakingReader *StakingReaderSession) GetLevelStats() (struct {
	StakedCounts       [7]*big.Int
	Supplies           [7]*big.Int
	StakingRatios      [7]*big.Int
	DynamicMultipliers [7]*big.Int
}, error) {
	return _StakingReader.Contract.GetLevelStats(&_StakingReader.CallOpts)
}

// GetLevelStats is a free data retrieval call binding the contract method 0x83537b46.
//
// Solidity: function getLevelStats() view returns(uint256[7] stakedCounts, uint256[7] supplies, uint256[7] stakingRatios, uint256[7] dynamicMultipliers)
func (_StakingReader *StakingReaderCallerSession) GetLevelStats() (struct {
	StakedCounts       [7]*big.Int
	Supplies           [7]*big.Int
	StakingRatios      [7]*big.Int
	DynamicMultipliers [7]*big.Int
}, error) {
	return _StakingReader.Contract.GetLevelStats(&_StakingReader.CallOpts)
}

// GetNFTDailyRewardBreakdown is a free data retrieval call binding the contract method 0x346771b1.
//
// Solidity: function getNFTDailyRewardBreakdown(uint256 tokenId) view returns(uint256 baseReward, uint256 decayedReward, uint256 comboBonus, uint256 dynamicMultiplier, uint256 finalReward)
func (_StakingReader *StakingReaderCaller) GetNFTDailyRewardBreakdown(opts *bind.CallOpts, tokenId *big.Int) (struct {
	BaseReward        *big.Int
	DecayedReward     *big.Int
	ComboBonus        *big.Int
	DynamicMultiplier *big.Int
	FinalReward       *big.Int
}, error) {
	var out []interface{}
	err := _StakingReader.contract.Call(opts, &out, "getNFTDailyRewardBreakdown", tokenId)

	outstruct := new(struct {
		BaseReward        *big.Int
		DecayedReward     *big.Int
		ComboBonus        *big.Int
		DynamicMultiplier *big.Int
		FinalReward       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BaseReward = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DecayedReward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ComboBonus = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.DynamicMultiplier = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.FinalReward = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNFTDailyRewardBreakdown is a free data retrieval call binding the contract method 0x346771b1.
//
// Solidity: function getNFTDailyRewardBreakdown(uint256 tokenId) view returns(uint256 baseReward, uint256 decayedReward, uint256 comboBonus, uint256 dynamicMultiplier, uint256 finalReward)
func (_StakingReader *StakingReaderSession) GetNFTDailyRewardBreakdown(tokenId *big.Int) (struct {
	BaseReward        *big.Int
	DecayedReward     *big.Int
	ComboBonus        *big.Int
	DynamicMultiplier *big.Int
	FinalReward       *big.Int
}, error) {
	return _StakingReader.Contract.GetNFTDailyRewardBreakdown(&_StakingReader.CallOpts, tokenId)
}

// GetNFTDailyRewardBreakdown is a free data retrieval call binding the contract method 0x346771b1.
//
// Solidity: function getNFTDailyRewardBreakdown(uint256 tokenId) view returns(uint256 baseReward, uint256 decayedReward, uint256 comboBonus, uint256 dynamicMultiplier, uint256 finalReward)
func (_StakingReader *StakingReaderCallerSession) GetNFTDailyRewardBreakdown(tokenId *big.Int) (struct {
	BaseReward        *big.Int
	DecayedReward     *big.Int
	ComboBonus        *big.Int
	DynamicMultiplier *big.Int
	FinalReward       *big.Int
}, error) {
	return _StakingReader.Contract.GetNFTDailyRewardBreakdown(&_StakingReader.CallOpts, tokenId)
}

// GetPendingRewards is a free data retrieval call binding the contract method 0xf9f87c18.
//
// Solidity: function getPendingRewards(uint256 tokenId) view returns(uint256)
func (_StakingReader *StakingReaderCaller) GetPendingRewards(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingReader.contract.Call(opts, &out, "getPendingRewards", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingRewards is a free data retrieval call binding the contract method 0xf9f87c18.
//
// Solidity: function getPendingRewards(uint256 tokenId) view returns(uint256)
func (_StakingReader *StakingReaderSession) GetPendingRewards(tokenId *big.Int) (*big.Int, error) {
	return _StakingReader.Contract.GetPendingRewards(&_StakingReader.CallOpts, tokenId)
}

// GetPendingRewards is a free data retrieval call binding the contract method 0xf9f87c18.
//
// Solidity: function getPendingRewards(uint256 tokenId) view returns(uint256)
func (_StakingReader *StakingReaderCallerSession) GetPendingRewards(tokenId *big.Int) (*big.Int, error) {
	return _StakingReader.Contract.GetPendingRewards(&_StakingReader.CallOpts, tokenId)
}

// GetPlatformStats is a free data retrieval call binding the contract method 0x136d8883.
//
// Solidity: function getPlatformStats() view returns(uint256[7] staked, uint256[7] supply)
func (_StakingReader *StakingReaderCaller) GetPlatformStats(opts *bind.CallOpts) (struct {
	Staked [7]*big.Int
	Supply [7]*big.Int
}, error) {
	var out []interface{}
	err := _StakingReader.contract.Call(opts, &out, "getPlatformStats")

	outstruct := new(struct {
		Staked [7]*big.Int
		Supply [7]*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Staked = *abi.ConvertType(out[0], new([7]*big.Int)).(*[7]*big.Int)
	outstruct.Supply = *abi.ConvertType(out[1], new([7]*big.Int)).(*[7]*big.Int)

	return *outstruct, err

}

// GetPlatformStats is a free data retrieval call binding the contract method 0x136d8883.
//
// Solidity: function getPlatformStats() view returns(uint256[7] staked, uint256[7] supply)
func (_StakingReader *StakingReaderSession) GetPlatformStats() (struct {
	Staked [7]*big.Int
	Supply [7]*big.Int
}, error) {
	return _StakingReader.Contract.GetPlatformStats(&_StakingReader.CallOpts)
}

// GetPlatformStats is a free data retrieval call binding the contract method 0x136d8883.
//
// Solidity: function getPlatformStats() view returns(uint256[7] staked, uint256[7] supply)
func (_StakingReader *StakingReaderCallerSession) GetPlatformStats() (struct {
	Staked [7]*big.Int
	Supply [7]*big.Int
}, error) {
	return _StakingReader.Contract.GetPlatformStats(&_StakingReader.CallOpts)
}

// GetStakeDetails is a free data retrieval call binding the contract method 0x963c9dd3.
//
// Solidity: function getStakeDetails(uint256 tokenId) view returns((address,uint256,uint8,uint256,uint256,bool,uint256,uint256,bool))
func (_StakingReader *StakingReaderCaller) GetStakeDetails(opts *bind.CallOpts, tokenId *big.Int) (StakingStakeInfo, error) {
	var out []interface{}
	err := _StakingReader.contract.Call(opts, &out, "getStakeDetails", tokenId)

	if err != nil {
		return *new(StakingStakeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(StakingStakeInfo)).(*StakingStakeInfo)

	return out0, err

}

// GetStakeDetails is a free data retrieval call binding the contract method 0x963c9dd3.
//
// Solidity: function getStakeDetails(uint256 tokenId) view returns((address,uint256,uint8,uint256,uint256,bool,uint256,uint256,bool))
func (_StakingReader *StakingReaderSession) GetStakeDetails(tokenId *big.Int) (StakingStakeInfo, error) {
	return _StakingReader.Contract.GetStakeDetails(&_StakingReader.CallOpts, tokenId)
}

// GetStakeDetails is a free data retrieval call binding the contract method 0x963c9dd3.
//
// Solidity: function getStakeDetails(uint256 tokenId) view returns((address,uint256,uint8,uint256,uint256,bool,uint256,uint256,bool))
func (_StakingReader *StakingReaderCallerSession) GetStakeDetails(tokenId *big.Int) (StakingStakeInfo, error) {
	return _StakingReader.Contract.GetStakeDetails(&_StakingReader.CallOpts, tokenId)
}

// GetUserComboSummary is a free data retrieval call binding the contract method 0x11428ead.
//
// Solidity: function getUserComboSummary(address user) view returns((uint256[6],uint256[6],uint256[6],bool[6]) summary)
func (_StakingReader *StakingReaderCaller) GetUserComboSummary(opts *bind.CallOpts, user common.Address) (StakingReaderUserComboSummary, error) {
	var out []interface{}
	err := _StakingReader.contract.Call(opts, &out, "getUserComboSummary", user)

	if err != nil {
		return *new(StakingReaderUserComboSummary), err
	}

	out0 := *abi.ConvertType(out[0], new(StakingReaderUserComboSummary)).(*StakingReaderUserComboSummary)

	return out0, err

}

// GetUserComboSummary is a free data retrieval call binding the contract method 0x11428ead.
//
// Solidity: function getUserComboSummary(address user) view returns((uint256[6],uint256[6],uint256[6],bool[6]) summary)
func (_StakingReader *StakingReaderSession) GetUserComboSummary(user common.Address) (StakingReaderUserComboSummary, error) {
	return _StakingReader.Contract.GetUserComboSummary(&_StakingReader.CallOpts, user)
}

// GetUserComboSummary is a free data retrieval call binding the contract method 0x11428ead.
//
// Solidity: function getUserComboSummary(address user) view returns((uint256[6],uint256[6],uint256[6],bool[6]) summary)
func (_StakingReader *StakingReaderCallerSession) GetUserComboSummary(user common.Address) (StakingReaderUserComboSummary, error) {
	return _StakingReader.Contract.GetUserComboSummary(&_StakingReader.CallOpts, user)
}

// GetUserDailyRewards is a free data retrieval call binding the contract method 0x68e4c3d1.
//
// Solidity: function getUserDailyRewards(address user) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256[6],uint256[6],uint256[6]))
func (_StakingReader *StakingReaderCaller) GetUserDailyRewards(opts *bind.CallOpts, user common.Address) (StakingReaderUserDailyRewards, error) {
	var out []interface{}
	err := _StakingReader.contract.Call(opts, &out, "getUserDailyRewards", user)

	if err != nil {
		return *new(StakingReaderUserDailyRewards), err
	}

	out0 := *abi.ConvertType(out[0], new(StakingReaderUserDailyRewards)).(*StakingReaderUserDailyRewards)

	return out0, err

}

// GetUserDailyRewards is a free data retrieval call binding the contract method 0x68e4c3d1.
//
// Solidity: function getUserDailyRewards(address user) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256[6],uint256[6],uint256[6]))
func (_StakingReader *StakingReaderSession) GetUserDailyRewards(user common.Address) (StakingReaderUserDailyRewards, error) {
	return _StakingReader.Contract.GetUserDailyRewards(&_StakingReader.CallOpts, user)
}

// GetUserDailyRewards is a free data retrieval call binding the contract method 0x68e4c3d1.
//
// Solidity: function getUserDailyRewards(address user) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256[6],uint256[6],uint256[6]))
func (_StakingReader *StakingReaderCallerSession) GetUserDailyRewards(user common.Address) (StakingReaderUserDailyRewards, error) {
	return _StakingReader.Contract.GetUserDailyRewards(&_StakingReader.CallOpts, user)
}

// GetUserDailyRewardsByLevel is a free data retrieval call binding the contract method 0x907e2632.
//
// Solidity: function getUserDailyRewardsByLevel(address user, uint8 level) view returns(uint256 baseReward, uint256 finalReward, uint256 nftCount)
func (_StakingReader *StakingReaderCaller) GetUserDailyRewardsByLevel(opts *bind.CallOpts, user common.Address, level uint8) (struct {
	BaseReward  *big.Int
	FinalReward *big.Int
	NftCount    *big.Int
}, error) {
	var out []interface{}
	err := _StakingReader.contract.Call(opts, &out, "getUserDailyRewardsByLevel", user, level)

	outstruct := new(struct {
		BaseReward  *big.Int
		FinalReward *big.Int
		NftCount    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BaseReward = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FinalReward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NftCount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetUserDailyRewardsByLevel is a free data retrieval call binding the contract method 0x907e2632.
//
// Solidity: function getUserDailyRewardsByLevel(address user, uint8 level) view returns(uint256 baseReward, uint256 finalReward, uint256 nftCount)
func (_StakingReader *StakingReaderSession) GetUserDailyRewardsByLevel(user common.Address, level uint8) (struct {
	BaseReward  *big.Int
	FinalReward *big.Int
	NftCount    *big.Int
}, error) {
	return _StakingReader.Contract.GetUserDailyRewardsByLevel(&_StakingReader.CallOpts, user, level)
}

// GetUserDailyRewardsByLevel is a free data retrieval call binding the contract method 0x907e2632.
//
// Solidity: function getUserDailyRewardsByLevel(address user, uint8 level) view returns(uint256 baseReward, uint256 finalReward, uint256 nftCount)
func (_StakingReader *StakingReaderCallerSession) GetUserDailyRewardsByLevel(user common.Address, level uint8) (struct {
	BaseReward  *big.Int
	FinalReward *big.Int
	NftCount    *big.Int
}, error) {
	return _StakingReader.Contract.GetUserDailyRewardsByLevel(&_StakingReader.CallOpts, user, level)
}

// GetUserRewardStats is a free data retrieval call binding the contract method 0x9b13027d.
//
// Solidity: function getUserRewardStats(address user) view returns((uint256,uint256,uint256[6],uint256,uint256) stats)
func (_StakingReader *StakingReaderCaller) GetUserRewardStats(opts *bind.CallOpts, user common.Address) (StakingReaderUserRewardStats, error) {
	var out []interface{}
	err := _StakingReader.contract.Call(opts, &out, "getUserRewardStats", user)

	if err != nil {
		return *new(StakingReaderUserRewardStats), err
	}

	out0 := *abi.ConvertType(out[0], new(StakingReaderUserRewardStats)).(*StakingReaderUserRewardStats)

	return out0, err

}

// GetUserRewardStats is a free data retrieval call binding the contract method 0x9b13027d.
//
// Solidity: function getUserRewardStats(address user) view returns((uint256,uint256,uint256[6],uint256,uint256) stats)
func (_StakingReader *StakingReaderSession) GetUserRewardStats(user common.Address) (StakingReaderUserRewardStats, error) {
	return _StakingReader.Contract.GetUserRewardStats(&_StakingReader.CallOpts, user)
}

// GetUserRewardStats is a free data retrieval call binding the contract method 0x9b13027d.
//
// Solidity: function getUserRewardStats(address user) view returns((uint256,uint256,uint256[6],uint256,uint256) stats)
func (_StakingReader *StakingReaderCallerSession) GetUserRewardStats(user common.Address) (StakingReaderUserRewardStats, error) {
	return _StakingReader.Contract.GetUserRewardStats(&_StakingReader.CallOpts, user)
}

// GetUserStakedNFTs is a free data retrieval call binding the contract method 0x8f833a55.
//
// Solidity: function getUserStakedNFTs(address user, uint256 offset, uint256 limit) view returns((uint256,uint8,uint256,uint256,uint256,uint256)[] nfts, uint256 total)
func (_StakingReader *StakingReaderCaller) GetUserStakedNFTs(opts *bind.CallOpts, user common.Address, offset *big.Int, limit *big.Int) (struct {
	Nfts  []StakingReaderStakedNFTInfo
	Total *big.Int
}, error) {
	var out []interface{}
	err := _StakingReader.contract.Call(opts, &out, "getUserStakedNFTs", user, offset, limit)

	outstruct := new(struct {
		Nfts  []StakingReaderStakedNFTInfo
		Total *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Nfts = *abi.ConvertType(out[0], new([]StakingReaderStakedNFTInfo)).(*[]StakingReaderStakedNFTInfo)
	outstruct.Total = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetUserStakedNFTs is a free data retrieval call binding the contract method 0x8f833a55.
//
// Solidity: function getUserStakedNFTs(address user, uint256 offset, uint256 limit) view returns((uint256,uint8,uint256,uint256,uint256,uint256)[] nfts, uint256 total)
func (_StakingReader *StakingReaderSession) GetUserStakedNFTs(user common.Address, offset *big.Int, limit *big.Int) (struct {
	Nfts  []StakingReaderStakedNFTInfo
	Total *big.Int
}, error) {
	return _StakingReader.Contract.GetUserStakedNFTs(&_StakingReader.CallOpts, user, offset, limit)
}

// GetUserStakedNFTs is a free data retrieval call binding the contract method 0x8f833a55.
//
// Solidity: function getUserStakedNFTs(address user, uint256 offset, uint256 limit) view returns((uint256,uint8,uint256,uint256,uint256,uint256)[] nfts, uint256 total)
func (_StakingReader *StakingReaderCallerSession) GetUserStakedNFTs(user common.Address, offset *big.Int, limit *big.Int) (struct {
	Nfts  []StakingReaderStakedNFTInfo
	Total *big.Int
}, error) {
	return _StakingReader.Contract.GetUserStakedNFTs(&_StakingReader.CallOpts, user, offset, limit)
}

// GetUserStakingSummary is a free data retrieval call binding the contract method 0xc691ef9b.
//
// Solidity: function getUserStakingSummary(address user) view returns((uint256,uint256,uint256,uint256[6],uint256,uint256) summary)
func (_StakingReader *StakingReaderCaller) GetUserStakingSummary(opts *bind.CallOpts, user common.Address) (StakingReaderUserStakingSummary, error) {
	var out []interface{}
	err := _StakingReader.contract.Call(opts, &out, "getUserStakingSummary", user)

	if err != nil {
		return *new(StakingReaderUserStakingSummary), err
	}

	out0 := *abi.ConvertType(out[0], new(StakingReaderUserStakingSummary)).(*StakingReaderUserStakingSummary)

	return out0, err

}

// GetUserStakingSummary is a free data retrieval call binding the contract method 0xc691ef9b.
//
// Solidity: function getUserStakingSummary(address user) view returns((uint256,uint256,uint256,uint256[6],uint256,uint256) summary)
func (_StakingReader *StakingReaderSession) GetUserStakingSummary(user common.Address) (StakingReaderUserStakingSummary, error) {
	return _StakingReader.Contract.GetUserStakingSummary(&_StakingReader.CallOpts, user)
}

// GetUserStakingSummary is a free data retrieval call binding the contract method 0xc691ef9b.
//
// Solidity: function getUserStakingSummary(address user) view returns((uint256,uint256,uint256,uint256[6],uint256,uint256) summary)
func (_StakingReader *StakingReaderCallerSession) GetUserStakingSummary(user common.Address) (StakingReaderUserStakingSummary, error) {
	return _StakingReader.Contract.GetUserStakingSummary(&_StakingReader.CallOpts, user)
}

// GetVersions is a free data retrieval call binding the contract method 0x6d0cc895.
//
// Solidity: function getVersions() view returns(string stakingVersion, string configVersion)
func (_StakingReader *StakingReaderCaller) GetVersions(opts *bind.CallOpts) (struct {
	StakingVersion string
	ConfigVersion  string
}, error) {
	var out []interface{}
	err := _StakingReader.contract.Call(opts, &out, "getVersions")

	outstruct := new(struct {
		StakingVersion string
		ConfigVersion  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StakingVersion = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.ConfigVersion = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// GetVersions is a free data retrieval call binding the contract method 0x6d0cc895.
//
// Solidity: function getVersions() view returns(string stakingVersion, string configVersion)
func (_StakingReader *StakingReaderSession) GetVersions() (struct {
	StakingVersion string
	ConfigVersion  string
}, error) {
	return _StakingReader.Contract.GetVersions(&_StakingReader.CallOpts)
}

// GetVersions is a free data retrieval call binding the contract method 0x6d0cc895.
//
// Solidity: function getVersions() view returns(string stakingVersion, string configVersion)
func (_StakingReader *StakingReaderCallerSession) GetVersions() (struct {
	StakingVersion string
	ConfigVersion  string
}, error) {
	return _StakingReader.Contract.GetVersions(&_StakingReader.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingReader *StakingReaderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingReader.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingReader *StakingReaderSession) Owner() (common.Address, error) {
	return _StakingReader.Contract.Owner(&_StakingReader.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingReader *StakingReaderCallerSession) Owner() (common.Address, error) {
	return _StakingReader.Contract.Owner(&_StakingReader.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_StakingReader *StakingReaderCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingReader.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_StakingReader *StakingReaderSession) ProxiableUUID() ([32]byte, error) {
	return _StakingReader.Contract.ProxiableUUID(&_StakingReader.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_StakingReader *StakingReaderCallerSession) ProxiableUUID() ([32]byte, error) {
	return _StakingReader.Contract.ProxiableUUID(&_StakingReader.CallOpts)
}

// StakingContract is a free data retrieval call binding the contract method 0xee99205c.
//
// Solidity: function stakingContract() view returns(address)
func (_StakingReader *StakingReaderCaller) StakingContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingReader.contract.Call(opts, &out, "stakingContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingContract is a free data retrieval call binding the contract method 0xee99205c.
//
// Solidity: function stakingContract() view returns(address)
func (_StakingReader *StakingReaderSession) StakingContract() (common.Address, error) {
	return _StakingReader.Contract.StakingContract(&_StakingReader.CallOpts)
}

// StakingContract is a free data retrieval call binding the contract method 0xee99205c.
//
// Solidity: function stakingContract() view returns(address)
func (_StakingReader *StakingReaderCallerSession) StakingContract() (common.Address, error) {
	return _StakingReader.Contract.StakingContract(&_StakingReader.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _stakingContract, address _configContract, address _cpnftContract, address _owner) returns()
func (_StakingReader *StakingReaderTransactor) Initialize(opts *bind.TransactOpts, _stakingContract common.Address, _configContract common.Address, _cpnftContract common.Address, _owner common.Address) (*types.Transaction, error) {
	return _StakingReader.contract.Transact(opts, "initialize", _stakingContract, _configContract, _cpnftContract, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _stakingContract, address _configContract, address _cpnftContract, address _owner) returns()
func (_StakingReader *StakingReaderSession) Initialize(_stakingContract common.Address, _configContract common.Address, _cpnftContract common.Address, _owner common.Address) (*types.Transaction, error) {
	return _StakingReader.Contract.Initialize(&_StakingReader.TransactOpts, _stakingContract, _configContract, _cpnftContract, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _stakingContract, address _configContract, address _cpnftContract, address _owner) returns()
func (_StakingReader *StakingReaderTransactorSession) Initialize(_stakingContract common.Address, _configContract common.Address, _cpnftContract common.Address, _owner common.Address) (*types.Transaction, error) {
	return _StakingReader.Contract.Initialize(&_StakingReader.TransactOpts, _stakingContract, _configContract, _cpnftContract, _owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingReader *StakingReaderTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingReader.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingReader *StakingReaderSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingReader.Contract.RenounceOwnership(&_StakingReader.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingReader *StakingReaderTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingReader.Contract.RenounceOwnership(&_StakingReader.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingReader *StakingReaderTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StakingReader.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingReader *StakingReaderSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingReader.Contract.TransferOwnership(&_StakingReader.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingReader *StakingReaderTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingReader.Contract.TransferOwnership(&_StakingReader.TransactOpts, newOwner)
}

// UpdateContracts is a paid mutator transaction binding the contract method 0x6dca4d44.
//
// Solidity: function updateContracts(address _stakingContract, address _configContract, address _cpnftContract) returns()
func (_StakingReader *StakingReaderTransactor) UpdateContracts(opts *bind.TransactOpts, _stakingContract common.Address, _configContract common.Address, _cpnftContract common.Address) (*types.Transaction, error) {
	return _StakingReader.contract.Transact(opts, "updateContracts", _stakingContract, _configContract, _cpnftContract)
}

// UpdateContracts is a paid mutator transaction binding the contract method 0x6dca4d44.
//
// Solidity: function updateContracts(address _stakingContract, address _configContract, address _cpnftContract) returns()
func (_StakingReader *StakingReaderSession) UpdateContracts(_stakingContract common.Address, _configContract common.Address, _cpnftContract common.Address) (*types.Transaction, error) {
	return _StakingReader.Contract.UpdateContracts(&_StakingReader.TransactOpts, _stakingContract, _configContract, _cpnftContract)
}

// UpdateContracts is a paid mutator transaction binding the contract method 0x6dca4d44.
//
// Solidity: function updateContracts(address _stakingContract, address _configContract, address _cpnftContract) returns()
func (_StakingReader *StakingReaderTransactorSession) UpdateContracts(_stakingContract common.Address, _configContract common.Address, _cpnftContract common.Address) (*types.Transaction, error) {
	return _StakingReader.Contract.UpdateContracts(&_StakingReader.TransactOpts, _stakingContract, _configContract, _cpnftContract)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StakingReader *StakingReaderTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StakingReader.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StakingReader *StakingReaderSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StakingReader.Contract.UpgradeToAndCall(&_StakingReader.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StakingReader *StakingReaderTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StakingReader.Contract.UpgradeToAndCall(&_StakingReader.TransactOpts, newImplementation, data)
}

// StakingReaderInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StakingReader contract.
type StakingReaderInitializedIterator struct {
	Event *StakingReaderInitialized // Event containing the contract specifics and raw log

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
func (it *StakingReaderInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingReaderInitialized)
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
		it.Event = new(StakingReaderInitialized)
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
func (it *StakingReaderInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingReaderInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingReaderInitialized represents a Initialized event raised by the StakingReader contract.
type StakingReaderInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_StakingReader *StakingReaderFilterer) FilterInitialized(opts *bind.FilterOpts) (*StakingReaderInitializedIterator, error) {

	logs, sub, err := _StakingReader.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StakingReaderInitializedIterator{contract: _StakingReader.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_StakingReader *StakingReaderFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StakingReaderInitialized) (event.Subscription, error) {

	logs, sub, err := _StakingReader.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingReaderInitialized)
				if err := _StakingReader.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_StakingReader *StakingReaderFilterer) ParseInitialized(log types.Log) (*StakingReaderInitialized, error) {
	event := new(StakingReaderInitialized)
	if err := _StakingReader.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingReaderOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StakingReader contract.
type StakingReaderOwnershipTransferredIterator struct {
	Event *StakingReaderOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakingReaderOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingReaderOwnershipTransferred)
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
		it.Event = new(StakingReaderOwnershipTransferred)
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
func (it *StakingReaderOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingReaderOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingReaderOwnershipTransferred represents a OwnershipTransferred event raised by the StakingReader contract.
type StakingReaderOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingReader *StakingReaderFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingReaderOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingReader.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingReaderOwnershipTransferredIterator{contract: _StakingReader.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingReader *StakingReaderFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingReaderOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingReader.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingReaderOwnershipTransferred)
				if err := _StakingReader.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StakingReader *StakingReaderFilterer) ParseOwnershipTransferred(log types.Log) (*StakingReaderOwnershipTransferred, error) {
	event := new(StakingReaderOwnershipTransferred)
	if err := _StakingReader.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingReaderUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the StakingReader contract.
type StakingReaderUpgradedIterator struct {
	Event *StakingReaderUpgraded // Event containing the contract specifics and raw log

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
func (it *StakingReaderUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingReaderUpgraded)
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
		it.Event = new(StakingReaderUpgraded)
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
func (it *StakingReaderUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingReaderUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingReaderUpgraded represents a Upgraded event raised by the StakingReader contract.
type StakingReaderUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_StakingReader *StakingReaderFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*StakingReaderUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _StakingReader.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &StakingReaderUpgradedIterator{contract: _StakingReader.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_StakingReader *StakingReaderFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *StakingReaderUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _StakingReader.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingReaderUpgraded)
				if err := _StakingReader.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_StakingReader *StakingReaderFilterer) ParseUpgraded(log types.Log) (*StakingReaderUpgraded, error) {
	event := new(StakingReaderUpgraded)
	if err := _StakingReader.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
