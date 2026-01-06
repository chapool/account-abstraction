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

// IMarketplaceBid is an auto generated low-level Go binding around an user-defined struct.
type IMarketplaceBid struct {
	Bidder    common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Refunded  bool
}

// IMarketplaceListing is an auto generated low-level Go binding around an user-defined struct.
type IMarketplaceListing struct {
	ListingId       *big.Int
	Seller          common.Address
	NftContract     common.Address
	TokenId         *big.Int
	ListingType     uint8
	Price           *big.Int
	Status          uint8
	StartTime       *big.Int
	EndTime         *big.Int
	MinBidIncrement *big.Int
}

// MarketplaceMetaData contains all meta data concerning the Marketplace contract.
var MarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"finalPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"platformFee\",\"type\":\"uint256\"}],\"name\":\"AuctionSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousBidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"}],\"name\":\"BidPlaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BidRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"platformFee\",\"type\":\"uint256\"}],\"name\":\"ItemSold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"ListingCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumIMarketplace.ListingType\",\"name\":\"listingType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"ListingCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bids\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"refunded\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"buyItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"}],\"name\":\"cancelListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"enumIMarketplace.ListingType\",\"name\":\"listingType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBidIncrement\",\"type\":\"uint256\"}],\"name\":\"createListing\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delistingLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delistingRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastResetTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delistingWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableTestMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"emergencyWithdrawNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialTimestamp\",\"type\":\"uint256\"}],\"name\":\"enableTestMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"}],\"name\":\"getBids\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"refunded\",\"type\":\"bool\"}],\"internalType\":\"structIMarketplace.Bid[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getDelistingRecord\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastResetTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"}],\"name\":\"getHighestBid\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"refunded\",\"type\":\"bool\"}],\"internalType\":\"structIMarketplace.Bid\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"}],\"name\":\"getListing\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"enumIMarketplace.ListingType\",\"name\":\"listingType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"enumIMarketplace.ListingStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBidIncrement\",\"type\":\"uint256\"}],\"internalType\":\"structIMarketplace.Listing\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserListings\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"highestBidIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_paymentToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_platformFeeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_platformFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_delistingLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_delistingWindow\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"listings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"enumIMarketplace.ListingType\",\"name\":\"listingType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"enumIMarketplace.ListingStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBidIncrement\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextListingId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nftContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paymentToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"name\":\"placeBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformFeeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"setTestTimestamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"settleAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"testMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"testTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_delistingLimit\",\"type\":\"uint256\"}],\"name\":\"updateDelistingLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_delistingWindow\",\"type\":\"uint256\"}],\"name\":\"updateDelistingWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftContract\",\"type\":\"address\"}],\"name\":\"updateNFTContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_paymentToken\",\"type\":\"address\"}],\"name\":\"updatePaymentToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_platformFeeRate\",\"type\":\"uint256\"}],\"name\":\"updatePlatformFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_platformFeeRecipient\",\"type\":\"address\"}],\"name\":\"updatePlatformFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userListings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// MarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketplaceMetaData.ABI instead.
var MarketplaceABI = MarketplaceMetaData.ABI

// Marketplace is an auto generated Go binding around an Ethereum contract.
type Marketplace struct {
	MarketplaceCaller     // Read-only binding to the contract
	MarketplaceTransactor // Write-only binding to the contract
	MarketplaceFilterer   // Log filterer for contract events
}

// MarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketplaceSession struct {
	Contract     *Marketplace      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketplaceCallerSession struct {
	Contract *MarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketplaceTransactorSession struct {
	Contract     *MarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketplaceRaw struct {
	Contract *Marketplace // Generic contract binding to access the raw methods on
}

// MarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketplaceCallerRaw struct {
	Contract *MarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// MarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketplaceTransactorRaw struct {
	Contract *MarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketplace creates a new instance of Marketplace, bound to a specific deployed contract.
func NewMarketplace(address common.Address, backend bind.ContractBackend) (*Marketplace, error) {
	contract, err := bindMarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Marketplace{MarketplaceCaller: MarketplaceCaller{contract: contract}, MarketplaceTransactor: MarketplaceTransactor{contract: contract}, MarketplaceFilterer: MarketplaceFilterer{contract: contract}}, nil
}

// NewMarketplaceCaller creates a new read-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceCaller(address common.Address, caller bind.ContractCaller) (*MarketplaceCaller, error) {
	contract, err := bindMarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceCaller{contract: contract}, nil
}

// NewMarketplaceTransactor creates a new write-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketplaceTransactor, error) {
	contract, err := bindMarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceTransactor{contract: contract}, nil
}

// NewMarketplaceFilterer creates a new log filterer instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketplaceFilterer, error) {
	contract, err := bindMarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketplaceFilterer{contract: contract}, nil
}

// bindMarketplace binds a generic wrapper to an already deployed contract.
func bindMarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MarketplaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.MarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Marketplace *MarketplaceCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Marketplace *MarketplaceSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Marketplace.Contract.UPGRADEINTERFACEVERSION(&_Marketplace.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Marketplace *MarketplaceCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Marketplace.Contract.UPGRADEINTERFACEVERSION(&_Marketplace.CallOpts)
}

// Bids is a free data retrieval call binding the contract method 0x7b3c4baa.
//
// Solidity: function bids(uint256 , uint256 ) view returns(address bidder, uint256 amount, uint256 timestamp, bool refunded)
func (_Marketplace *MarketplaceCaller) Bids(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Bidder    common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Refunded  bool
}, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "bids", arg0, arg1)

	outstruct := new(struct {
		Bidder    common.Address
		Amount    *big.Int
		Timestamp *big.Int
		Refunded  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Bidder = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Refunded = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Bids is a free data retrieval call binding the contract method 0x7b3c4baa.
//
// Solidity: function bids(uint256 , uint256 ) view returns(address bidder, uint256 amount, uint256 timestamp, bool refunded)
func (_Marketplace *MarketplaceSession) Bids(arg0 *big.Int, arg1 *big.Int) (struct {
	Bidder    common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Refunded  bool
}, error) {
	return _Marketplace.Contract.Bids(&_Marketplace.CallOpts, arg0, arg1)
}

// Bids is a free data retrieval call binding the contract method 0x7b3c4baa.
//
// Solidity: function bids(uint256 , uint256 ) view returns(address bidder, uint256 amount, uint256 timestamp, bool refunded)
func (_Marketplace *MarketplaceCallerSession) Bids(arg0 *big.Int, arg1 *big.Int) (struct {
	Bidder    common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Refunded  bool
}, error) {
	return _Marketplace.Contract.Bids(&_Marketplace.CallOpts, arg0, arg1)
}

// DelistingLimit is a free data retrieval call binding the contract method 0xd93d1dc8.
//
// Solidity: function delistingLimit() view returns(uint256)
func (_Marketplace *MarketplaceCaller) DelistingLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "delistingLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelistingLimit is a free data retrieval call binding the contract method 0xd93d1dc8.
//
// Solidity: function delistingLimit() view returns(uint256)
func (_Marketplace *MarketplaceSession) DelistingLimit() (*big.Int, error) {
	return _Marketplace.Contract.DelistingLimit(&_Marketplace.CallOpts)
}

// DelistingLimit is a free data retrieval call binding the contract method 0xd93d1dc8.
//
// Solidity: function delistingLimit() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) DelistingLimit() (*big.Int, error) {
	return _Marketplace.Contract.DelistingLimit(&_Marketplace.CallOpts)
}

// DelistingRecords is a free data retrieval call binding the contract method 0xb39b6e74.
//
// Solidity: function delistingRecords(address ) view returns(uint256 count, uint256 lastResetTime)
func (_Marketplace *MarketplaceCaller) DelistingRecords(opts *bind.CallOpts, arg0 common.Address) (struct {
	Count         *big.Int
	LastResetTime *big.Int
}, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "delistingRecords", arg0)

	outstruct := new(struct {
		Count         *big.Int
		LastResetTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Count = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastResetTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DelistingRecords is a free data retrieval call binding the contract method 0xb39b6e74.
//
// Solidity: function delistingRecords(address ) view returns(uint256 count, uint256 lastResetTime)
func (_Marketplace *MarketplaceSession) DelistingRecords(arg0 common.Address) (struct {
	Count         *big.Int
	LastResetTime *big.Int
}, error) {
	return _Marketplace.Contract.DelistingRecords(&_Marketplace.CallOpts, arg0)
}

// DelistingRecords is a free data retrieval call binding the contract method 0xb39b6e74.
//
// Solidity: function delistingRecords(address ) view returns(uint256 count, uint256 lastResetTime)
func (_Marketplace *MarketplaceCallerSession) DelistingRecords(arg0 common.Address) (struct {
	Count         *big.Int
	LastResetTime *big.Int
}, error) {
	return _Marketplace.Contract.DelistingRecords(&_Marketplace.CallOpts, arg0)
}

// DelistingWindow is a free data retrieval call binding the contract method 0x50d832f8.
//
// Solidity: function delistingWindow() view returns(uint256)
func (_Marketplace *MarketplaceCaller) DelistingWindow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "delistingWindow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelistingWindow is a free data retrieval call binding the contract method 0x50d832f8.
//
// Solidity: function delistingWindow() view returns(uint256)
func (_Marketplace *MarketplaceSession) DelistingWindow() (*big.Int, error) {
	return _Marketplace.Contract.DelistingWindow(&_Marketplace.CallOpts)
}

// DelistingWindow is a free data retrieval call binding the contract method 0x50d832f8.
//
// Solidity: function delistingWindow() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) DelistingWindow() (*big.Int, error) {
	return _Marketplace.Contract.DelistingWindow(&_Marketplace.CallOpts)
}

// GetBids is a free data retrieval call binding the contract method 0x131d9a27.
//
// Solidity: function getBids(uint256 listingId) view returns((address,uint256,uint256,bool)[])
func (_Marketplace *MarketplaceCaller) GetBids(opts *bind.CallOpts, listingId *big.Int) ([]IMarketplaceBid, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getBids", listingId)

	if err != nil {
		return *new([]IMarketplaceBid), err
	}

	out0 := *abi.ConvertType(out[0], new([]IMarketplaceBid)).(*[]IMarketplaceBid)

	return out0, err

}

// GetBids is a free data retrieval call binding the contract method 0x131d9a27.
//
// Solidity: function getBids(uint256 listingId) view returns((address,uint256,uint256,bool)[])
func (_Marketplace *MarketplaceSession) GetBids(listingId *big.Int) ([]IMarketplaceBid, error) {
	return _Marketplace.Contract.GetBids(&_Marketplace.CallOpts, listingId)
}

// GetBids is a free data retrieval call binding the contract method 0x131d9a27.
//
// Solidity: function getBids(uint256 listingId) view returns((address,uint256,uint256,bool)[])
func (_Marketplace *MarketplaceCallerSession) GetBids(listingId *big.Int) ([]IMarketplaceBid, error) {
	return _Marketplace.Contract.GetBids(&_Marketplace.CallOpts, listingId)
}

// GetCurrentTimestamp is a free data retrieval call binding the contract method 0x6c9230db.
//
// Solidity: function getCurrentTimestamp() view returns(uint256)
func (_Marketplace *MarketplaceCaller) GetCurrentTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getCurrentTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentTimestamp is a free data retrieval call binding the contract method 0x6c9230db.
//
// Solidity: function getCurrentTimestamp() view returns(uint256)
func (_Marketplace *MarketplaceSession) GetCurrentTimestamp() (*big.Int, error) {
	return _Marketplace.Contract.GetCurrentTimestamp(&_Marketplace.CallOpts)
}

// GetCurrentTimestamp is a free data retrieval call binding the contract method 0x6c9230db.
//
// Solidity: function getCurrentTimestamp() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) GetCurrentTimestamp() (*big.Int, error) {
	return _Marketplace.Contract.GetCurrentTimestamp(&_Marketplace.CallOpts)
}

// GetDelistingRecord is a free data retrieval call binding the contract method 0x0c5bf9be.
//
// Solidity: function getDelistingRecord(address user) view returns(uint256 count, uint256 lastResetTime, uint256 remainingCount)
func (_Marketplace *MarketplaceCaller) GetDelistingRecord(opts *bind.CallOpts, user common.Address) (struct {
	Count          *big.Int
	LastResetTime  *big.Int
	RemainingCount *big.Int
}, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getDelistingRecord", user)

	outstruct := new(struct {
		Count          *big.Int
		LastResetTime  *big.Int
		RemainingCount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Count = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastResetTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RemainingCount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetDelistingRecord is a free data retrieval call binding the contract method 0x0c5bf9be.
//
// Solidity: function getDelistingRecord(address user) view returns(uint256 count, uint256 lastResetTime, uint256 remainingCount)
func (_Marketplace *MarketplaceSession) GetDelistingRecord(user common.Address) (struct {
	Count          *big.Int
	LastResetTime  *big.Int
	RemainingCount *big.Int
}, error) {
	return _Marketplace.Contract.GetDelistingRecord(&_Marketplace.CallOpts, user)
}

// GetDelistingRecord is a free data retrieval call binding the contract method 0x0c5bf9be.
//
// Solidity: function getDelistingRecord(address user) view returns(uint256 count, uint256 lastResetTime, uint256 remainingCount)
func (_Marketplace *MarketplaceCallerSession) GetDelistingRecord(user common.Address) (struct {
	Count          *big.Int
	LastResetTime  *big.Int
	RemainingCount *big.Int
}, error) {
	return _Marketplace.Contract.GetDelistingRecord(&_Marketplace.CallOpts, user)
}

// GetHighestBid is a free data retrieval call binding the contract method 0x8f288644.
//
// Solidity: function getHighestBid(uint256 listingId) view returns((address,uint256,uint256,bool))
func (_Marketplace *MarketplaceCaller) GetHighestBid(opts *bind.CallOpts, listingId *big.Int) (IMarketplaceBid, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getHighestBid", listingId)

	if err != nil {
		return *new(IMarketplaceBid), err
	}

	out0 := *abi.ConvertType(out[0], new(IMarketplaceBid)).(*IMarketplaceBid)

	return out0, err

}

// GetHighestBid is a free data retrieval call binding the contract method 0x8f288644.
//
// Solidity: function getHighestBid(uint256 listingId) view returns((address,uint256,uint256,bool))
func (_Marketplace *MarketplaceSession) GetHighestBid(listingId *big.Int) (IMarketplaceBid, error) {
	return _Marketplace.Contract.GetHighestBid(&_Marketplace.CallOpts, listingId)
}

// GetHighestBid is a free data retrieval call binding the contract method 0x8f288644.
//
// Solidity: function getHighestBid(uint256 listingId) view returns((address,uint256,uint256,bool))
func (_Marketplace *MarketplaceCallerSession) GetHighestBid(listingId *big.Int) (IMarketplaceBid, error) {
	return _Marketplace.Contract.GetHighestBid(&_Marketplace.CallOpts, listingId)
}

// GetListing is a free data retrieval call binding the contract method 0x107a274a.
//
// Solidity: function getListing(uint256 listingId) view returns((uint256,address,address,uint256,uint8,uint256,uint8,uint256,uint256,uint256))
func (_Marketplace *MarketplaceCaller) GetListing(opts *bind.CallOpts, listingId *big.Int) (IMarketplaceListing, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getListing", listingId)

	if err != nil {
		return *new(IMarketplaceListing), err
	}

	out0 := *abi.ConvertType(out[0], new(IMarketplaceListing)).(*IMarketplaceListing)

	return out0, err

}

// GetListing is a free data retrieval call binding the contract method 0x107a274a.
//
// Solidity: function getListing(uint256 listingId) view returns((uint256,address,address,uint256,uint8,uint256,uint8,uint256,uint256,uint256))
func (_Marketplace *MarketplaceSession) GetListing(listingId *big.Int) (IMarketplaceListing, error) {
	return _Marketplace.Contract.GetListing(&_Marketplace.CallOpts, listingId)
}

// GetListing is a free data retrieval call binding the contract method 0x107a274a.
//
// Solidity: function getListing(uint256 listingId) view returns((uint256,address,address,uint256,uint8,uint256,uint8,uint256,uint256,uint256))
func (_Marketplace *MarketplaceCallerSession) GetListing(listingId *big.Int) (IMarketplaceListing, error) {
	return _Marketplace.Contract.GetListing(&_Marketplace.CallOpts, listingId)
}

// GetUserListings is a free data retrieval call binding the contract method 0xeb18a9d8.
//
// Solidity: function getUserListings(address user) view returns(uint256[])
func (_Marketplace *MarketplaceCaller) GetUserListings(opts *bind.CallOpts, user common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getUserListings", user)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetUserListings is a free data retrieval call binding the contract method 0xeb18a9d8.
//
// Solidity: function getUserListings(address user) view returns(uint256[])
func (_Marketplace *MarketplaceSession) GetUserListings(user common.Address) ([]*big.Int, error) {
	return _Marketplace.Contract.GetUserListings(&_Marketplace.CallOpts, user)
}

// GetUserListings is a free data retrieval call binding the contract method 0xeb18a9d8.
//
// Solidity: function getUserListings(address user) view returns(uint256[])
func (_Marketplace *MarketplaceCallerSession) GetUserListings(user common.Address) ([]*big.Int, error) {
	return _Marketplace.Contract.GetUserListings(&_Marketplace.CallOpts, user)
}

// HighestBidIndex is a free data retrieval call binding the contract method 0x3f6f8a02.
//
// Solidity: function highestBidIndex(uint256 ) view returns(uint256)
func (_Marketplace *MarketplaceCaller) HighestBidIndex(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "highestBidIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HighestBidIndex is a free data retrieval call binding the contract method 0x3f6f8a02.
//
// Solidity: function highestBidIndex(uint256 ) view returns(uint256)
func (_Marketplace *MarketplaceSession) HighestBidIndex(arg0 *big.Int) (*big.Int, error) {
	return _Marketplace.Contract.HighestBidIndex(&_Marketplace.CallOpts, arg0)
}

// HighestBidIndex is a free data retrieval call binding the contract method 0x3f6f8a02.
//
// Solidity: function highestBidIndex(uint256 ) view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) HighestBidIndex(arg0 *big.Int) (*big.Int, error) {
	return _Marketplace.Contract.HighestBidIndex(&_Marketplace.CallOpts, arg0)
}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(uint256 listingId, address seller, address nftContract, uint256 tokenId, uint8 listingType, uint256 price, uint8 status, uint256 startTime, uint256 endTime, uint256 minBidIncrement)
func (_Marketplace *MarketplaceCaller) Listings(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ListingId       *big.Int
	Seller          common.Address
	NftContract     common.Address
	TokenId         *big.Int
	ListingType     uint8
	Price           *big.Int
	Status          uint8
	StartTime       *big.Int
	EndTime         *big.Int
	MinBidIncrement *big.Int
}, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "listings", arg0)

	outstruct := new(struct {
		ListingId       *big.Int
		Seller          common.Address
		NftContract     common.Address
		TokenId         *big.Int
		ListingType     uint8
		Price           *big.Int
		Status          uint8
		StartTime       *big.Int
		EndTime         *big.Int
		MinBidIncrement *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ListingId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Seller = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.NftContract = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ListingType = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.Price = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[6], new(uint8)).(*uint8)
	outstruct.StartTime = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.MinBidIncrement = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(uint256 listingId, address seller, address nftContract, uint256 tokenId, uint8 listingType, uint256 price, uint8 status, uint256 startTime, uint256 endTime, uint256 minBidIncrement)
func (_Marketplace *MarketplaceSession) Listings(arg0 *big.Int) (struct {
	ListingId       *big.Int
	Seller          common.Address
	NftContract     common.Address
	TokenId         *big.Int
	ListingType     uint8
	Price           *big.Int
	Status          uint8
	StartTime       *big.Int
	EndTime         *big.Int
	MinBidIncrement *big.Int
}, error) {
	return _Marketplace.Contract.Listings(&_Marketplace.CallOpts, arg0)
}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(uint256 listingId, address seller, address nftContract, uint256 tokenId, uint8 listingType, uint256 price, uint8 status, uint256 startTime, uint256 endTime, uint256 minBidIncrement)
func (_Marketplace *MarketplaceCallerSession) Listings(arg0 *big.Int) (struct {
	ListingId       *big.Int
	Seller          common.Address
	NftContract     common.Address
	TokenId         *big.Int
	ListingType     uint8
	Price           *big.Int
	Status          uint8
	StartTime       *big.Int
	EndTime         *big.Int
	MinBidIncrement *big.Int
}, error) {
	return _Marketplace.Contract.Listings(&_Marketplace.CallOpts, arg0)
}

// NextListingId is a free data retrieval call binding the contract method 0xaaccf1ec.
//
// Solidity: function nextListingId() view returns(uint256)
func (_Marketplace *MarketplaceCaller) NextListingId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "nextListingId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextListingId is a free data retrieval call binding the contract method 0xaaccf1ec.
//
// Solidity: function nextListingId() view returns(uint256)
func (_Marketplace *MarketplaceSession) NextListingId() (*big.Int, error) {
	return _Marketplace.Contract.NextListingId(&_Marketplace.CallOpts)
}

// NextListingId is a free data retrieval call binding the contract method 0xaaccf1ec.
//
// Solidity: function nextListingId() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) NextListingId() (*big.Int, error) {
	return _Marketplace.Contract.NextListingId(&_Marketplace.CallOpts)
}

// NftContract is a free data retrieval call binding the contract method 0xd56d229d.
//
// Solidity: function nftContract() view returns(address)
func (_Marketplace *MarketplaceCaller) NftContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "nftContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NftContract is a free data retrieval call binding the contract method 0xd56d229d.
//
// Solidity: function nftContract() view returns(address)
func (_Marketplace *MarketplaceSession) NftContract() (common.Address, error) {
	return _Marketplace.Contract.NftContract(&_Marketplace.CallOpts)
}

// NftContract is a free data retrieval call binding the contract method 0xd56d229d.
//
// Solidity: function nftContract() view returns(address)
func (_Marketplace *MarketplaceCallerSession) NftContract() (common.Address, error) {
	return _Marketplace.Contract.NftContract(&_Marketplace.CallOpts)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Marketplace *MarketplaceCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Marketplace *MarketplaceSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Marketplace.Contract.OnERC721Received(&_Marketplace.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Marketplace *MarketplaceCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Marketplace.Contract.OnERC721Received(&_Marketplace.CallOpts, arg0, arg1, arg2, arg3)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Marketplace *MarketplaceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Marketplace *MarketplaceSession) Owner() (common.Address, error) {
	return _Marketplace.Contract.Owner(&_Marketplace.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Marketplace *MarketplaceCallerSession) Owner() (common.Address, error) {
	return _Marketplace.Contract.Owner(&_Marketplace.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Marketplace *MarketplaceCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Marketplace *MarketplaceSession) Paused() (bool, error) {
	return _Marketplace.Contract.Paused(&_Marketplace.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Marketplace *MarketplaceCallerSession) Paused() (bool, error) {
	return _Marketplace.Contract.Paused(&_Marketplace.CallOpts)
}

// PaymentToken is a free data retrieval call binding the contract method 0x3013ce29.
//
// Solidity: function paymentToken() view returns(address)
func (_Marketplace *MarketplaceCaller) PaymentToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "paymentToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PaymentToken is a free data retrieval call binding the contract method 0x3013ce29.
//
// Solidity: function paymentToken() view returns(address)
func (_Marketplace *MarketplaceSession) PaymentToken() (common.Address, error) {
	return _Marketplace.Contract.PaymentToken(&_Marketplace.CallOpts)
}

// PaymentToken is a free data retrieval call binding the contract method 0x3013ce29.
//
// Solidity: function paymentToken() view returns(address)
func (_Marketplace *MarketplaceCallerSession) PaymentToken() (common.Address, error) {
	return _Marketplace.Contract.PaymentToken(&_Marketplace.CallOpts)
}

// PlatformFeeRate is a free data retrieval call binding the contract method 0xeeca08f0.
//
// Solidity: function platformFeeRate() view returns(uint256)
func (_Marketplace *MarketplaceCaller) PlatformFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "platformFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformFeeRate is a free data retrieval call binding the contract method 0xeeca08f0.
//
// Solidity: function platformFeeRate() view returns(uint256)
func (_Marketplace *MarketplaceSession) PlatformFeeRate() (*big.Int, error) {
	return _Marketplace.Contract.PlatformFeeRate(&_Marketplace.CallOpts)
}

// PlatformFeeRate is a free data retrieval call binding the contract method 0xeeca08f0.
//
// Solidity: function platformFeeRate() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) PlatformFeeRate() (*big.Int, error) {
	return _Marketplace.Contract.PlatformFeeRate(&_Marketplace.CallOpts)
}

// PlatformFeeRecipient is a free data retrieval call binding the contract method 0xeb13554f.
//
// Solidity: function platformFeeRecipient() view returns(address)
func (_Marketplace *MarketplaceCaller) PlatformFeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "platformFeeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PlatformFeeRecipient is a free data retrieval call binding the contract method 0xeb13554f.
//
// Solidity: function platformFeeRecipient() view returns(address)
func (_Marketplace *MarketplaceSession) PlatformFeeRecipient() (common.Address, error) {
	return _Marketplace.Contract.PlatformFeeRecipient(&_Marketplace.CallOpts)
}

// PlatformFeeRecipient is a free data retrieval call binding the contract method 0xeb13554f.
//
// Solidity: function platformFeeRecipient() view returns(address)
func (_Marketplace *MarketplaceCallerSession) PlatformFeeRecipient() (common.Address, error) {
	return _Marketplace.Contract.PlatformFeeRecipient(&_Marketplace.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Marketplace *MarketplaceCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Marketplace *MarketplaceSession) ProxiableUUID() ([32]byte, error) {
	return _Marketplace.Contract.ProxiableUUID(&_Marketplace.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Marketplace *MarketplaceCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Marketplace.Contract.ProxiableUUID(&_Marketplace.CallOpts)
}

// TestMode is a free data retrieval call binding the contract method 0xcd9ea342.
//
// Solidity: function testMode() view returns(bool)
func (_Marketplace *MarketplaceCaller) TestMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "testMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TestMode is a free data retrieval call binding the contract method 0xcd9ea342.
//
// Solidity: function testMode() view returns(bool)
func (_Marketplace *MarketplaceSession) TestMode() (bool, error) {
	return _Marketplace.Contract.TestMode(&_Marketplace.CallOpts)
}

// TestMode is a free data retrieval call binding the contract method 0xcd9ea342.
//
// Solidity: function testMode() view returns(bool)
func (_Marketplace *MarketplaceCallerSession) TestMode() (bool, error) {
	return _Marketplace.Contract.TestMode(&_Marketplace.CallOpts)
}

// TestTimestamp is a free data retrieval call binding the contract method 0x7e773d9d.
//
// Solidity: function testTimestamp() view returns(uint256)
func (_Marketplace *MarketplaceCaller) TestTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "testTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TestTimestamp is a free data retrieval call binding the contract method 0x7e773d9d.
//
// Solidity: function testTimestamp() view returns(uint256)
func (_Marketplace *MarketplaceSession) TestTimestamp() (*big.Int, error) {
	return _Marketplace.Contract.TestTimestamp(&_Marketplace.CallOpts)
}

// TestTimestamp is a free data retrieval call binding the contract method 0x7e773d9d.
//
// Solidity: function testTimestamp() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) TestTimestamp() (*big.Int, error) {
	return _Marketplace.Contract.TestTimestamp(&_Marketplace.CallOpts)
}

// UserListings is a free data retrieval call binding the contract method 0x86cfa7d3.
//
// Solidity: function userListings(address , uint256 ) view returns(uint256)
func (_Marketplace *MarketplaceCaller) UserListings(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "userListings", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserListings is a free data retrieval call binding the contract method 0x86cfa7d3.
//
// Solidity: function userListings(address , uint256 ) view returns(uint256)
func (_Marketplace *MarketplaceSession) UserListings(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Marketplace.Contract.UserListings(&_Marketplace.CallOpts, arg0, arg1)
}

// UserListings is a free data retrieval call binding the contract method 0x86cfa7d3.
//
// Solidity: function userListings(address , uint256 ) view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) UserListings(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Marketplace.Contract.UserListings(&_Marketplace.CallOpts, arg0, arg1)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Marketplace *MarketplaceCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Marketplace *MarketplaceSession) Version() (string, error) {
	return _Marketplace.Contract.Version(&_Marketplace.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Marketplace *MarketplaceCallerSession) Version() (string, error) {
	return _Marketplace.Contract.Version(&_Marketplace.CallOpts)
}

// BuyItem is a paid mutator transaction binding the contract method 0x7383d709.
//
// Solidity: function buyItem(uint256 listingId, address buyer) returns()
func (_Marketplace *MarketplaceTransactor) BuyItem(opts *bind.TransactOpts, listingId *big.Int, buyer common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "buyItem", listingId, buyer)
}

// BuyItem is a paid mutator transaction binding the contract method 0x7383d709.
//
// Solidity: function buyItem(uint256 listingId, address buyer) returns()
func (_Marketplace *MarketplaceSession) BuyItem(listingId *big.Int, buyer common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.BuyItem(&_Marketplace.TransactOpts, listingId, buyer)
}

// BuyItem is a paid mutator transaction binding the contract method 0x7383d709.
//
// Solidity: function buyItem(uint256 listingId, address buyer) returns()
func (_Marketplace *MarketplaceTransactorSession) BuyItem(listingId *big.Int, buyer common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.BuyItem(&_Marketplace.TransactOpts, listingId, buyer)
}

// CancelListing is a paid mutator transaction binding the contract method 0x305a67a8.
//
// Solidity: function cancelListing(uint256 listingId) returns()
func (_Marketplace *MarketplaceTransactor) CancelListing(opts *bind.TransactOpts, listingId *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "cancelListing", listingId)
}

// CancelListing is a paid mutator transaction binding the contract method 0x305a67a8.
//
// Solidity: function cancelListing(uint256 listingId) returns()
func (_Marketplace *MarketplaceSession) CancelListing(listingId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CancelListing(&_Marketplace.TransactOpts, listingId)
}

// CancelListing is a paid mutator transaction binding the contract method 0x305a67a8.
//
// Solidity: function cancelListing(uint256 listingId) returns()
func (_Marketplace *MarketplaceTransactorSession) CancelListing(listingId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CancelListing(&_Marketplace.TransactOpts, listingId)
}

// CreateListing is a paid mutator transaction binding the contract method 0xc58b4cc5.
//
// Solidity: function createListing(address seller, uint256 tokenId, uint8 listingType, uint256 price, uint256 duration, uint256 minBidIncrement) returns(uint256 listingId)
func (_Marketplace *MarketplaceTransactor) CreateListing(opts *bind.TransactOpts, seller common.Address, tokenId *big.Int, listingType uint8, price *big.Int, duration *big.Int, minBidIncrement *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "createListing", seller, tokenId, listingType, price, duration, minBidIncrement)
}

// CreateListing is a paid mutator transaction binding the contract method 0xc58b4cc5.
//
// Solidity: function createListing(address seller, uint256 tokenId, uint8 listingType, uint256 price, uint256 duration, uint256 minBidIncrement) returns(uint256 listingId)
func (_Marketplace *MarketplaceSession) CreateListing(seller common.Address, tokenId *big.Int, listingType uint8, price *big.Int, duration *big.Int, minBidIncrement *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CreateListing(&_Marketplace.TransactOpts, seller, tokenId, listingType, price, duration, minBidIncrement)
}

// CreateListing is a paid mutator transaction binding the contract method 0xc58b4cc5.
//
// Solidity: function createListing(address seller, uint256 tokenId, uint8 listingType, uint256 price, uint256 duration, uint256 minBidIncrement) returns(uint256 listingId)
func (_Marketplace *MarketplaceTransactorSession) CreateListing(seller common.Address, tokenId *big.Int, listingType uint8, price *big.Int, duration *big.Int, minBidIncrement *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CreateListing(&_Marketplace.TransactOpts, seller, tokenId, listingType, price, duration, minBidIncrement)
}

// DisableTestMode is a paid mutator transaction binding the contract method 0x9ddca0c8.
//
// Solidity: function disableTestMode() returns()
func (_Marketplace *MarketplaceTransactor) DisableTestMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "disableTestMode")
}

// DisableTestMode is a paid mutator transaction binding the contract method 0x9ddca0c8.
//
// Solidity: function disableTestMode() returns()
func (_Marketplace *MarketplaceSession) DisableTestMode() (*types.Transaction, error) {
	return _Marketplace.Contract.DisableTestMode(&_Marketplace.TransactOpts)
}

// DisableTestMode is a paid mutator transaction binding the contract method 0x9ddca0c8.
//
// Solidity: function disableTestMode() returns()
func (_Marketplace *MarketplaceTransactorSession) DisableTestMode() (*types.Transaction, error) {
	return _Marketplace.Contract.DisableTestMode(&_Marketplace.TransactOpts)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xe63ea408.
//
// Solidity: function emergencyWithdraw(address token, address to, uint256 amount) returns()
func (_Marketplace *MarketplaceTransactor) EmergencyWithdraw(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "emergencyWithdraw", token, to, amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xe63ea408.
//
// Solidity: function emergencyWithdraw(address token, address to, uint256 amount) returns()
func (_Marketplace *MarketplaceSession) EmergencyWithdraw(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.EmergencyWithdraw(&_Marketplace.TransactOpts, token, to, amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xe63ea408.
//
// Solidity: function emergencyWithdraw(address token, address to, uint256 amount) returns()
func (_Marketplace *MarketplaceTransactorSession) EmergencyWithdraw(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.EmergencyWithdraw(&_Marketplace.TransactOpts, token, to, amount)
}

// EmergencyWithdrawNFT is a paid mutator transaction binding the contract method 0x2abdcf79.
//
// Solidity: function emergencyWithdrawNFT(uint256 tokenId, address to) returns()
func (_Marketplace *MarketplaceTransactor) EmergencyWithdrawNFT(opts *bind.TransactOpts, tokenId *big.Int, to common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "emergencyWithdrawNFT", tokenId, to)
}

// EmergencyWithdrawNFT is a paid mutator transaction binding the contract method 0x2abdcf79.
//
// Solidity: function emergencyWithdrawNFT(uint256 tokenId, address to) returns()
func (_Marketplace *MarketplaceSession) EmergencyWithdrawNFT(tokenId *big.Int, to common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.EmergencyWithdrawNFT(&_Marketplace.TransactOpts, tokenId, to)
}

// EmergencyWithdrawNFT is a paid mutator transaction binding the contract method 0x2abdcf79.
//
// Solidity: function emergencyWithdrawNFT(uint256 tokenId, address to) returns()
func (_Marketplace *MarketplaceTransactorSession) EmergencyWithdrawNFT(tokenId *big.Int, to common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.EmergencyWithdrawNFT(&_Marketplace.TransactOpts, tokenId, to)
}

// EnableTestMode is a paid mutator transaction binding the contract method 0x7dc21e65.
//
// Solidity: function enableTestMode(uint256 initialTimestamp) returns()
func (_Marketplace *MarketplaceTransactor) EnableTestMode(opts *bind.TransactOpts, initialTimestamp *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "enableTestMode", initialTimestamp)
}

// EnableTestMode is a paid mutator transaction binding the contract method 0x7dc21e65.
//
// Solidity: function enableTestMode(uint256 initialTimestamp) returns()
func (_Marketplace *MarketplaceSession) EnableTestMode(initialTimestamp *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.EnableTestMode(&_Marketplace.TransactOpts, initialTimestamp)
}

// EnableTestMode is a paid mutator transaction binding the contract method 0x7dc21e65.
//
// Solidity: function enableTestMode(uint256 initialTimestamp) returns()
func (_Marketplace *MarketplaceTransactorSession) EnableTestMode(initialTimestamp *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.EnableTestMode(&_Marketplace.TransactOpts, initialTimestamp)
}

// Initialize is a paid mutator transaction binding the contract method 0x8754bbc6.
//
// Solidity: function initialize(address _nftContract, address _paymentToken, address _platformFeeRecipient, uint256 _platformFeeRate, uint256 _delistingLimit, uint256 _delistingWindow, address _owner) returns()
func (_Marketplace *MarketplaceTransactor) Initialize(opts *bind.TransactOpts, _nftContract common.Address, _paymentToken common.Address, _platformFeeRecipient common.Address, _platformFeeRate *big.Int, _delistingLimit *big.Int, _delistingWindow *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "initialize", _nftContract, _paymentToken, _platformFeeRecipient, _platformFeeRate, _delistingLimit, _delistingWindow, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x8754bbc6.
//
// Solidity: function initialize(address _nftContract, address _paymentToken, address _platformFeeRecipient, uint256 _platformFeeRate, uint256 _delistingLimit, uint256 _delistingWindow, address _owner) returns()
func (_Marketplace *MarketplaceSession) Initialize(_nftContract common.Address, _paymentToken common.Address, _platformFeeRecipient common.Address, _platformFeeRate *big.Int, _delistingLimit *big.Int, _delistingWindow *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.Initialize(&_Marketplace.TransactOpts, _nftContract, _paymentToken, _platformFeeRecipient, _platformFeeRate, _delistingLimit, _delistingWindow, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x8754bbc6.
//
// Solidity: function initialize(address _nftContract, address _paymentToken, address _platformFeeRecipient, uint256 _platformFeeRate, uint256 _delistingLimit, uint256 _delistingWindow, address _owner) returns()
func (_Marketplace *MarketplaceTransactorSession) Initialize(_nftContract common.Address, _paymentToken common.Address, _platformFeeRecipient common.Address, _platformFeeRate *big.Int, _delistingLimit *big.Int, _delistingWindow *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.Initialize(&_Marketplace.TransactOpts, _nftContract, _paymentToken, _platformFeeRecipient, _platformFeeRate, _delistingLimit, _delistingWindow, _owner)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Marketplace *MarketplaceTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Marketplace *MarketplaceSession) Pause() (*types.Transaction, error) {
	return _Marketplace.Contract.Pause(&_Marketplace.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Marketplace *MarketplaceTransactorSession) Pause() (*types.Transaction, error) {
	return _Marketplace.Contract.Pause(&_Marketplace.TransactOpts)
}

// PlaceBid is a paid mutator transaction binding the contract method 0xf1e2f348.
//
// Solidity: function placeBid(uint256 listingId, address bidder, uint256 bidAmount) returns()
func (_Marketplace *MarketplaceTransactor) PlaceBid(opts *bind.TransactOpts, listingId *big.Int, bidder common.Address, bidAmount *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "placeBid", listingId, bidder, bidAmount)
}

// PlaceBid is a paid mutator transaction binding the contract method 0xf1e2f348.
//
// Solidity: function placeBid(uint256 listingId, address bidder, uint256 bidAmount) returns()
func (_Marketplace *MarketplaceSession) PlaceBid(listingId *big.Int, bidder common.Address, bidAmount *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.PlaceBid(&_Marketplace.TransactOpts, listingId, bidder, bidAmount)
}

// PlaceBid is a paid mutator transaction binding the contract method 0xf1e2f348.
//
// Solidity: function placeBid(uint256 listingId, address bidder, uint256 bidAmount) returns()
func (_Marketplace *MarketplaceTransactorSession) PlaceBid(listingId *big.Int, bidder common.Address, bidAmount *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.PlaceBid(&_Marketplace.TransactOpts, listingId, bidder, bidAmount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Marketplace *MarketplaceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Marketplace *MarketplaceSession) RenounceOwnership() (*types.Transaction, error) {
	return _Marketplace.Contract.RenounceOwnership(&_Marketplace.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Marketplace *MarketplaceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Marketplace.Contract.RenounceOwnership(&_Marketplace.TransactOpts)
}

// SetTestTimestamp is a paid mutator transaction binding the contract method 0x8b0a49a7.
//
// Solidity: function setTestTimestamp(uint256 timestamp) returns()
func (_Marketplace *MarketplaceTransactor) SetTestTimestamp(opts *bind.TransactOpts, timestamp *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setTestTimestamp", timestamp)
}

// SetTestTimestamp is a paid mutator transaction binding the contract method 0x8b0a49a7.
//
// Solidity: function setTestTimestamp(uint256 timestamp) returns()
func (_Marketplace *MarketplaceSession) SetTestTimestamp(timestamp *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.SetTestTimestamp(&_Marketplace.TransactOpts, timestamp)
}

// SetTestTimestamp is a paid mutator transaction binding the contract method 0x8b0a49a7.
//
// Solidity: function setTestTimestamp(uint256 timestamp) returns()
func (_Marketplace *MarketplaceTransactorSession) SetTestTimestamp(timestamp *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.SetTestTimestamp(&_Marketplace.TransactOpts, timestamp)
}

// SettleAuction is a paid mutator transaction binding the contract method 0xe4566f2d.
//
// Solidity: function settleAuction(uint256 listingId, address winner) returns()
func (_Marketplace *MarketplaceTransactor) SettleAuction(opts *bind.TransactOpts, listingId *big.Int, winner common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "settleAuction", listingId, winner)
}

// SettleAuction is a paid mutator transaction binding the contract method 0xe4566f2d.
//
// Solidity: function settleAuction(uint256 listingId, address winner) returns()
func (_Marketplace *MarketplaceSession) SettleAuction(listingId *big.Int, winner common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SettleAuction(&_Marketplace.TransactOpts, listingId, winner)
}

// SettleAuction is a paid mutator transaction binding the contract method 0xe4566f2d.
//
// Solidity: function settleAuction(uint256 listingId, address winner) returns()
func (_Marketplace *MarketplaceTransactorSession) SettleAuction(listingId *big.Int, winner common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SettleAuction(&_Marketplace.TransactOpts, listingId, winner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Marketplace *MarketplaceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Marketplace *MarketplaceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.TransferOwnership(&_Marketplace.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Marketplace *MarketplaceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.TransferOwnership(&_Marketplace.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Marketplace *MarketplaceTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Marketplace *MarketplaceSession) Unpause() (*types.Transaction, error) {
	return _Marketplace.Contract.Unpause(&_Marketplace.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Marketplace *MarketplaceTransactorSession) Unpause() (*types.Transaction, error) {
	return _Marketplace.Contract.Unpause(&_Marketplace.TransactOpts)
}

// UpdateDelistingLimit is a paid mutator transaction binding the contract method 0xff150b23.
//
// Solidity: function updateDelistingLimit(uint256 _delistingLimit) returns()
func (_Marketplace *MarketplaceTransactor) UpdateDelistingLimit(opts *bind.TransactOpts, _delistingLimit *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "updateDelistingLimit", _delistingLimit)
}

// UpdateDelistingLimit is a paid mutator transaction binding the contract method 0xff150b23.
//
// Solidity: function updateDelistingLimit(uint256 _delistingLimit) returns()
func (_Marketplace *MarketplaceSession) UpdateDelistingLimit(_delistingLimit *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.UpdateDelistingLimit(&_Marketplace.TransactOpts, _delistingLimit)
}

// UpdateDelistingLimit is a paid mutator transaction binding the contract method 0xff150b23.
//
// Solidity: function updateDelistingLimit(uint256 _delistingLimit) returns()
func (_Marketplace *MarketplaceTransactorSession) UpdateDelistingLimit(_delistingLimit *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.UpdateDelistingLimit(&_Marketplace.TransactOpts, _delistingLimit)
}

// UpdateDelistingWindow is a paid mutator transaction binding the contract method 0x0994aea2.
//
// Solidity: function updateDelistingWindow(uint256 _delistingWindow) returns()
func (_Marketplace *MarketplaceTransactor) UpdateDelistingWindow(opts *bind.TransactOpts, _delistingWindow *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "updateDelistingWindow", _delistingWindow)
}

// UpdateDelistingWindow is a paid mutator transaction binding the contract method 0x0994aea2.
//
// Solidity: function updateDelistingWindow(uint256 _delistingWindow) returns()
func (_Marketplace *MarketplaceSession) UpdateDelistingWindow(_delistingWindow *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.UpdateDelistingWindow(&_Marketplace.TransactOpts, _delistingWindow)
}

// UpdateDelistingWindow is a paid mutator transaction binding the contract method 0x0994aea2.
//
// Solidity: function updateDelistingWindow(uint256 _delistingWindow) returns()
func (_Marketplace *MarketplaceTransactorSession) UpdateDelistingWindow(_delistingWindow *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.UpdateDelistingWindow(&_Marketplace.TransactOpts, _delistingWindow)
}

// UpdateNFTContract is a paid mutator transaction binding the contract method 0x78627f2b.
//
// Solidity: function updateNFTContract(address _nftContract) returns()
func (_Marketplace *MarketplaceTransactor) UpdateNFTContract(opts *bind.TransactOpts, _nftContract common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "updateNFTContract", _nftContract)
}

// UpdateNFTContract is a paid mutator transaction binding the contract method 0x78627f2b.
//
// Solidity: function updateNFTContract(address _nftContract) returns()
func (_Marketplace *MarketplaceSession) UpdateNFTContract(_nftContract common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.UpdateNFTContract(&_Marketplace.TransactOpts, _nftContract)
}

// UpdateNFTContract is a paid mutator transaction binding the contract method 0x78627f2b.
//
// Solidity: function updateNFTContract(address _nftContract) returns()
func (_Marketplace *MarketplaceTransactorSession) UpdateNFTContract(_nftContract common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.UpdateNFTContract(&_Marketplace.TransactOpts, _nftContract)
}

// UpdatePaymentToken is a paid mutator transaction binding the contract method 0x4ccf1b34.
//
// Solidity: function updatePaymentToken(address _paymentToken) returns()
func (_Marketplace *MarketplaceTransactor) UpdatePaymentToken(opts *bind.TransactOpts, _paymentToken common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "updatePaymentToken", _paymentToken)
}

// UpdatePaymentToken is a paid mutator transaction binding the contract method 0x4ccf1b34.
//
// Solidity: function updatePaymentToken(address _paymentToken) returns()
func (_Marketplace *MarketplaceSession) UpdatePaymentToken(_paymentToken common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.UpdatePaymentToken(&_Marketplace.TransactOpts, _paymentToken)
}

// UpdatePaymentToken is a paid mutator transaction binding the contract method 0x4ccf1b34.
//
// Solidity: function updatePaymentToken(address _paymentToken) returns()
func (_Marketplace *MarketplaceTransactorSession) UpdatePaymentToken(_paymentToken common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.UpdatePaymentToken(&_Marketplace.TransactOpts, _paymentToken)
}

// UpdatePlatformFeeRate is a paid mutator transaction binding the contract method 0xae3b7f06.
//
// Solidity: function updatePlatformFeeRate(uint256 _platformFeeRate) returns()
func (_Marketplace *MarketplaceTransactor) UpdatePlatformFeeRate(opts *bind.TransactOpts, _platformFeeRate *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "updatePlatformFeeRate", _platformFeeRate)
}

// UpdatePlatformFeeRate is a paid mutator transaction binding the contract method 0xae3b7f06.
//
// Solidity: function updatePlatformFeeRate(uint256 _platformFeeRate) returns()
func (_Marketplace *MarketplaceSession) UpdatePlatformFeeRate(_platformFeeRate *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.UpdatePlatformFeeRate(&_Marketplace.TransactOpts, _platformFeeRate)
}

// UpdatePlatformFeeRate is a paid mutator transaction binding the contract method 0xae3b7f06.
//
// Solidity: function updatePlatformFeeRate(uint256 _platformFeeRate) returns()
func (_Marketplace *MarketplaceTransactorSession) UpdatePlatformFeeRate(_platformFeeRate *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.UpdatePlatformFeeRate(&_Marketplace.TransactOpts, _platformFeeRate)
}

// UpdatePlatformFeeRecipient is a paid mutator transaction binding the contract method 0xf5fe7f71.
//
// Solidity: function updatePlatformFeeRecipient(address _platformFeeRecipient) returns()
func (_Marketplace *MarketplaceTransactor) UpdatePlatformFeeRecipient(opts *bind.TransactOpts, _platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "updatePlatformFeeRecipient", _platformFeeRecipient)
}

// UpdatePlatformFeeRecipient is a paid mutator transaction binding the contract method 0xf5fe7f71.
//
// Solidity: function updatePlatformFeeRecipient(address _platformFeeRecipient) returns()
func (_Marketplace *MarketplaceSession) UpdatePlatformFeeRecipient(_platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.UpdatePlatformFeeRecipient(&_Marketplace.TransactOpts, _platformFeeRecipient)
}

// UpdatePlatformFeeRecipient is a paid mutator transaction binding the contract method 0xf5fe7f71.
//
// Solidity: function updatePlatformFeeRecipient(address _platformFeeRecipient) returns()
func (_Marketplace *MarketplaceTransactorSession) UpdatePlatformFeeRecipient(_platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.UpdatePlatformFeeRecipient(&_Marketplace.TransactOpts, _platformFeeRecipient)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Marketplace *MarketplaceTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Marketplace *MarketplaceSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Marketplace.Contract.UpgradeToAndCall(&_Marketplace.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Marketplace *MarketplaceTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Marketplace.Contract.UpgradeToAndCall(&_Marketplace.TransactOpts, newImplementation, data)
}

// MarketplaceAuctionSettledIterator is returned from FilterAuctionSettled and is used to iterate over the raw logs and unpacked data for AuctionSettled events raised by the Marketplace contract.
type MarketplaceAuctionSettledIterator struct {
	Event *MarketplaceAuctionSettled // Event containing the contract specifics and raw log

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
func (it *MarketplaceAuctionSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceAuctionSettled)
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
		it.Event = new(MarketplaceAuctionSettled)
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
func (it *MarketplaceAuctionSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceAuctionSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceAuctionSettled represents a AuctionSettled event raised by the Marketplace contract.
type MarketplaceAuctionSettled struct {
	ListingId   *big.Int
	Seller      common.Address
	Winner      common.Address
	FinalPrice  *big.Int
	PlatformFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAuctionSettled is a free log retrieval operation binding the contract event 0xf2c88560a571e69247330e22df504ac9cd88ebd638c975954fddf870a9039177.
//
// Solidity: event AuctionSettled(uint256 indexed listingId, address indexed seller, address indexed winner, uint256 finalPrice, uint256 platformFee)
func (_Marketplace *MarketplaceFilterer) FilterAuctionSettled(opts *bind.FilterOpts, listingId []*big.Int, seller []common.Address, winner []common.Address) (*MarketplaceAuctionSettledIterator, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "AuctionSettled", listingIdRule, sellerRule, winnerRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceAuctionSettledIterator{contract: _Marketplace.contract, event: "AuctionSettled", logs: logs, sub: sub}, nil
}

// WatchAuctionSettled is a free log subscription operation binding the contract event 0xf2c88560a571e69247330e22df504ac9cd88ebd638c975954fddf870a9039177.
//
// Solidity: event AuctionSettled(uint256 indexed listingId, address indexed seller, address indexed winner, uint256 finalPrice, uint256 platformFee)
func (_Marketplace *MarketplaceFilterer) WatchAuctionSettled(opts *bind.WatchOpts, sink chan<- *MarketplaceAuctionSettled, listingId []*big.Int, seller []common.Address, winner []common.Address) (event.Subscription, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "AuctionSettled", listingIdRule, sellerRule, winnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceAuctionSettled)
				if err := _Marketplace.contract.UnpackLog(event, "AuctionSettled", log); err != nil {
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

// ParseAuctionSettled is a log parse operation binding the contract event 0xf2c88560a571e69247330e22df504ac9cd88ebd638c975954fddf870a9039177.
//
// Solidity: event AuctionSettled(uint256 indexed listingId, address indexed seller, address indexed winner, uint256 finalPrice, uint256 platformFee)
func (_Marketplace *MarketplaceFilterer) ParseAuctionSettled(log types.Log) (*MarketplaceAuctionSettled, error) {
	event := new(MarketplaceAuctionSettled)
	if err := _Marketplace.contract.UnpackLog(event, "AuctionSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceBidPlacedIterator is returned from FilterBidPlaced and is used to iterate over the raw logs and unpacked data for BidPlaced events raised by the Marketplace contract.
type MarketplaceBidPlacedIterator struct {
	Event *MarketplaceBidPlaced // Event containing the contract specifics and raw log

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
func (it *MarketplaceBidPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceBidPlaced)
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
		it.Event = new(MarketplaceBidPlaced)
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
func (it *MarketplaceBidPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceBidPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceBidPlaced represents a BidPlaced event raised by the Marketplace contract.
type MarketplaceBidPlaced struct {
	ListingId      *big.Int
	Bidder         common.Address
	Amount         *big.Int
	PreviousBidder common.Address
	RefundAmount   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBidPlaced is a free log retrieval operation binding the contract event 0x2e296671c28b83e813c76e2acf7481f5a2cc46aaeb9bcf33b3e048f50e9c33e9.
//
// Solidity: event BidPlaced(uint256 indexed listingId, address indexed bidder, uint256 amount, address indexed previousBidder, uint256 refundAmount)
func (_Marketplace *MarketplaceFilterer) FilterBidPlaced(opts *bind.FilterOpts, listingId []*big.Int, bidder []common.Address, previousBidder []common.Address) (*MarketplaceBidPlacedIterator, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	var previousBidderRule []interface{}
	for _, previousBidderItem := range previousBidder {
		previousBidderRule = append(previousBidderRule, previousBidderItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "BidPlaced", listingIdRule, bidderRule, previousBidderRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceBidPlacedIterator{contract: _Marketplace.contract, event: "BidPlaced", logs: logs, sub: sub}, nil
}

// WatchBidPlaced is a free log subscription operation binding the contract event 0x2e296671c28b83e813c76e2acf7481f5a2cc46aaeb9bcf33b3e048f50e9c33e9.
//
// Solidity: event BidPlaced(uint256 indexed listingId, address indexed bidder, uint256 amount, address indexed previousBidder, uint256 refundAmount)
func (_Marketplace *MarketplaceFilterer) WatchBidPlaced(opts *bind.WatchOpts, sink chan<- *MarketplaceBidPlaced, listingId []*big.Int, bidder []common.Address, previousBidder []common.Address) (event.Subscription, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	var previousBidderRule []interface{}
	for _, previousBidderItem := range previousBidder {
		previousBidderRule = append(previousBidderRule, previousBidderItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "BidPlaced", listingIdRule, bidderRule, previousBidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceBidPlaced)
				if err := _Marketplace.contract.UnpackLog(event, "BidPlaced", log); err != nil {
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

// ParseBidPlaced is a log parse operation binding the contract event 0x2e296671c28b83e813c76e2acf7481f5a2cc46aaeb9bcf33b3e048f50e9c33e9.
//
// Solidity: event BidPlaced(uint256 indexed listingId, address indexed bidder, uint256 amount, address indexed previousBidder, uint256 refundAmount)
func (_Marketplace *MarketplaceFilterer) ParseBidPlaced(log types.Log) (*MarketplaceBidPlaced, error) {
	event := new(MarketplaceBidPlaced)
	if err := _Marketplace.contract.UnpackLog(event, "BidPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceBidRefundedIterator is returned from FilterBidRefunded and is used to iterate over the raw logs and unpacked data for BidRefunded events raised by the Marketplace contract.
type MarketplaceBidRefundedIterator struct {
	Event *MarketplaceBidRefunded // Event containing the contract specifics and raw log

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
func (it *MarketplaceBidRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceBidRefunded)
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
		it.Event = new(MarketplaceBidRefunded)
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
func (it *MarketplaceBidRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceBidRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceBidRefunded represents a BidRefunded event raised by the Marketplace contract.
type MarketplaceBidRefunded struct {
	ListingId *big.Int
	Bidder    common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBidRefunded is a free log retrieval operation binding the contract event 0xa59312996dc68b7f0224b341fed75d5afeb3bb20624107d37942d14f59428156.
//
// Solidity: event BidRefunded(uint256 indexed listingId, address indexed bidder, uint256 amount)
func (_Marketplace *MarketplaceFilterer) FilterBidRefunded(opts *bind.FilterOpts, listingId []*big.Int, bidder []common.Address) (*MarketplaceBidRefundedIterator, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "BidRefunded", listingIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceBidRefundedIterator{contract: _Marketplace.contract, event: "BidRefunded", logs: logs, sub: sub}, nil
}

// WatchBidRefunded is a free log subscription operation binding the contract event 0xa59312996dc68b7f0224b341fed75d5afeb3bb20624107d37942d14f59428156.
//
// Solidity: event BidRefunded(uint256 indexed listingId, address indexed bidder, uint256 amount)
func (_Marketplace *MarketplaceFilterer) WatchBidRefunded(opts *bind.WatchOpts, sink chan<- *MarketplaceBidRefunded, listingId []*big.Int, bidder []common.Address) (event.Subscription, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "BidRefunded", listingIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceBidRefunded)
				if err := _Marketplace.contract.UnpackLog(event, "BidRefunded", log); err != nil {
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

// ParseBidRefunded is a log parse operation binding the contract event 0xa59312996dc68b7f0224b341fed75d5afeb3bb20624107d37942d14f59428156.
//
// Solidity: event BidRefunded(uint256 indexed listingId, address indexed bidder, uint256 amount)
func (_Marketplace *MarketplaceFilterer) ParseBidRefunded(log types.Log) (*MarketplaceBidRefunded, error) {
	event := new(MarketplaceBidRefunded)
	if err := _Marketplace.contract.UnpackLog(event, "BidRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Marketplace contract.
type MarketplaceInitializedIterator struct {
	Event *MarketplaceInitialized // Event containing the contract specifics and raw log

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
func (it *MarketplaceInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceInitialized)
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
		it.Event = new(MarketplaceInitialized)
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
func (it *MarketplaceInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceInitialized represents a Initialized event raised by the Marketplace contract.
type MarketplaceInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Marketplace *MarketplaceFilterer) FilterInitialized(opts *bind.FilterOpts) (*MarketplaceInitializedIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MarketplaceInitializedIterator{contract: _Marketplace.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Marketplace *MarketplaceFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MarketplaceInitialized) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceInitialized)
				if err := _Marketplace.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseInitialized(log types.Log) (*MarketplaceInitialized, error) {
	event := new(MarketplaceInitialized)
	if err := _Marketplace.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceItemSoldIterator is returned from FilterItemSold and is used to iterate over the raw logs and unpacked data for ItemSold events raised by the Marketplace contract.
type MarketplaceItemSoldIterator struct {
	Event *MarketplaceItemSold // Event containing the contract specifics and raw log

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
func (it *MarketplaceItemSoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceItemSold)
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
		it.Event = new(MarketplaceItemSold)
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
func (it *MarketplaceItemSoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceItemSoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceItemSold represents a ItemSold event raised by the Marketplace contract.
type MarketplaceItemSold struct {
	ListingId   *big.Int
	Seller      common.Address
	Buyer       common.Address
	Price       *big.Int
	PlatformFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterItemSold is a free log retrieval operation binding the contract event 0xd2b8648ec6ff6bd9ed10162d6ec424ae792c12820a475e8ef52f23fd7da4f1eb.
//
// Solidity: event ItemSold(uint256 indexed listingId, address indexed seller, address indexed buyer, uint256 price, uint256 platformFee)
func (_Marketplace *MarketplaceFilterer) FilterItemSold(opts *bind.FilterOpts, listingId []*big.Int, seller []common.Address, buyer []common.Address) (*MarketplaceItemSoldIterator, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "ItemSold", listingIdRule, sellerRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceItemSoldIterator{contract: _Marketplace.contract, event: "ItemSold", logs: logs, sub: sub}, nil
}

// WatchItemSold is a free log subscription operation binding the contract event 0xd2b8648ec6ff6bd9ed10162d6ec424ae792c12820a475e8ef52f23fd7da4f1eb.
//
// Solidity: event ItemSold(uint256 indexed listingId, address indexed seller, address indexed buyer, uint256 price, uint256 platformFee)
func (_Marketplace *MarketplaceFilterer) WatchItemSold(opts *bind.WatchOpts, sink chan<- *MarketplaceItemSold, listingId []*big.Int, seller []common.Address, buyer []common.Address) (event.Subscription, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "ItemSold", listingIdRule, sellerRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceItemSold)
				if err := _Marketplace.contract.UnpackLog(event, "ItemSold", log); err != nil {
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

// ParseItemSold is a log parse operation binding the contract event 0xd2b8648ec6ff6bd9ed10162d6ec424ae792c12820a475e8ef52f23fd7da4f1eb.
//
// Solidity: event ItemSold(uint256 indexed listingId, address indexed seller, address indexed buyer, uint256 price, uint256 platformFee)
func (_Marketplace *MarketplaceFilterer) ParseItemSold(log types.Log) (*MarketplaceItemSold, error) {
	event := new(MarketplaceItemSold)
	if err := _Marketplace.contract.UnpackLog(event, "ItemSold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceListingCancelledIterator is returned from FilterListingCancelled and is used to iterate over the raw logs and unpacked data for ListingCancelled events raised by the Marketplace contract.
type MarketplaceListingCancelledIterator struct {
	Event *MarketplaceListingCancelled // Event containing the contract specifics and raw log

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
func (it *MarketplaceListingCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceListingCancelled)
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
		it.Event = new(MarketplaceListingCancelled)
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
func (it *MarketplaceListingCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceListingCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceListingCancelled represents a ListingCancelled event raised by the Marketplace contract.
type MarketplaceListingCancelled struct {
	ListingId *big.Int
	Seller    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterListingCancelled is a free log retrieval operation binding the contract event 0x8e25282255ab31897df2b0456bb993ac7f84d376861aefd84901d2d63a7428a2.
//
// Solidity: event ListingCancelled(uint256 indexed listingId, address indexed seller)
func (_Marketplace *MarketplaceFilterer) FilterListingCancelled(opts *bind.FilterOpts, listingId []*big.Int, seller []common.Address) (*MarketplaceListingCancelledIterator, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "ListingCancelled", listingIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceListingCancelledIterator{contract: _Marketplace.contract, event: "ListingCancelled", logs: logs, sub: sub}, nil
}

// WatchListingCancelled is a free log subscription operation binding the contract event 0x8e25282255ab31897df2b0456bb993ac7f84d376861aefd84901d2d63a7428a2.
//
// Solidity: event ListingCancelled(uint256 indexed listingId, address indexed seller)
func (_Marketplace *MarketplaceFilterer) WatchListingCancelled(opts *bind.WatchOpts, sink chan<- *MarketplaceListingCancelled, listingId []*big.Int, seller []common.Address) (event.Subscription, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "ListingCancelled", listingIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceListingCancelled)
				if err := _Marketplace.contract.UnpackLog(event, "ListingCancelled", log); err != nil {
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

// ParseListingCancelled is a log parse operation binding the contract event 0x8e25282255ab31897df2b0456bb993ac7f84d376861aefd84901d2d63a7428a2.
//
// Solidity: event ListingCancelled(uint256 indexed listingId, address indexed seller)
func (_Marketplace *MarketplaceFilterer) ParseListingCancelled(log types.Log) (*MarketplaceListingCancelled, error) {
	event := new(MarketplaceListingCancelled)
	if err := _Marketplace.contract.UnpackLog(event, "ListingCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceListingCreatedIterator is returned from FilterListingCreated and is used to iterate over the raw logs and unpacked data for ListingCreated events raised by the Marketplace contract.
type MarketplaceListingCreatedIterator struct {
	Event *MarketplaceListingCreated // Event containing the contract specifics and raw log

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
func (it *MarketplaceListingCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceListingCreated)
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
		it.Event = new(MarketplaceListingCreated)
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
func (it *MarketplaceListingCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceListingCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceListingCreated represents a ListingCreated event raised by the Marketplace contract.
type MarketplaceListingCreated struct {
	ListingId   *big.Int
	Seller      common.Address
	NftContract common.Address
	TokenId     *big.Int
	ListingType uint8
	Price       *big.Int
	EndTime     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterListingCreated is a free log retrieval operation binding the contract event 0xd12c2f3f0275527fd4502e85db3c970f8329166a46346f8858e097b0b499dee0.
//
// Solidity: event ListingCreated(uint256 indexed listingId, address indexed seller, address indexed nftContract, uint256 tokenId, uint8 listingType, uint256 price, uint256 endTime)
func (_Marketplace *MarketplaceFilterer) FilterListingCreated(opts *bind.FilterOpts, listingId []*big.Int, seller []common.Address, nftContract []common.Address) (*MarketplaceListingCreatedIterator, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "ListingCreated", listingIdRule, sellerRule, nftContractRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceListingCreatedIterator{contract: _Marketplace.contract, event: "ListingCreated", logs: logs, sub: sub}, nil
}

// WatchListingCreated is a free log subscription operation binding the contract event 0xd12c2f3f0275527fd4502e85db3c970f8329166a46346f8858e097b0b499dee0.
//
// Solidity: event ListingCreated(uint256 indexed listingId, address indexed seller, address indexed nftContract, uint256 tokenId, uint8 listingType, uint256 price, uint256 endTime)
func (_Marketplace *MarketplaceFilterer) WatchListingCreated(opts *bind.WatchOpts, sink chan<- *MarketplaceListingCreated, listingId []*big.Int, seller []common.Address, nftContract []common.Address) (event.Subscription, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "ListingCreated", listingIdRule, sellerRule, nftContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceListingCreated)
				if err := _Marketplace.contract.UnpackLog(event, "ListingCreated", log); err != nil {
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

// ParseListingCreated is a log parse operation binding the contract event 0xd12c2f3f0275527fd4502e85db3c970f8329166a46346f8858e097b0b499dee0.
//
// Solidity: event ListingCreated(uint256 indexed listingId, address indexed seller, address indexed nftContract, uint256 tokenId, uint8 listingType, uint256 price, uint256 endTime)
func (_Marketplace *MarketplaceFilterer) ParseListingCreated(log types.Log) (*MarketplaceListingCreated, error) {
	event := new(MarketplaceListingCreated)
	if err := _Marketplace.contract.UnpackLog(event, "ListingCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Marketplace contract.
type MarketplaceOwnershipTransferredIterator struct {
	Event *MarketplaceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MarketplaceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceOwnershipTransferred)
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
		it.Event = new(MarketplaceOwnershipTransferred)
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
func (it *MarketplaceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceOwnershipTransferred represents a OwnershipTransferred event raised by the Marketplace contract.
type MarketplaceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Marketplace *MarketplaceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MarketplaceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceOwnershipTransferredIterator{contract: _Marketplace.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Marketplace *MarketplaceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MarketplaceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceOwnershipTransferred)
				if err := _Marketplace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseOwnershipTransferred(log types.Log) (*MarketplaceOwnershipTransferred, error) {
	event := new(MarketplaceOwnershipTransferred)
	if err := _Marketplace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplacePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Marketplace contract.
type MarketplacePausedIterator struct {
	Event *MarketplacePaused // Event containing the contract specifics and raw log

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
func (it *MarketplacePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplacePaused)
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
		it.Event = new(MarketplacePaused)
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
func (it *MarketplacePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplacePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplacePaused represents a Paused event raised by the Marketplace contract.
type MarketplacePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Marketplace *MarketplaceFilterer) FilterPaused(opts *bind.FilterOpts) (*MarketplacePausedIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MarketplacePausedIterator{contract: _Marketplace.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Marketplace *MarketplaceFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MarketplacePaused) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplacePaused)
				if err := _Marketplace.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParsePaused(log types.Log) (*MarketplacePaused, error) {
	event := new(MarketplacePaused)
	if err := _Marketplace.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Marketplace contract.
type MarketplaceUnpausedIterator struct {
	Event *MarketplaceUnpaused // Event containing the contract specifics and raw log

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
func (it *MarketplaceUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceUnpaused)
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
		it.Event = new(MarketplaceUnpaused)
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
func (it *MarketplaceUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceUnpaused represents a Unpaused event raised by the Marketplace contract.
type MarketplaceUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Marketplace *MarketplaceFilterer) FilterUnpaused(opts *bind.FilterOpts) (*MarketplaceUnpausedIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MarketplaceUnpausedIterator{contract: _Marketplace.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Marketplace *MarketplaceFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MarketplaceUnpaused) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceUnpaused)
				if err := _Marketplace.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseUnpaused(log types.Log) (*MarketplaceUnpaused, error) {
	event := new(MarketplaceUnpaused)
	if err := _Marketplace.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Marketplace contract.
type MarketplaceUpgradedIterator struct {
	Event *MarketplaceUpgraded // Event containing the contract specifics and raw log

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
func (it *MarketplaceUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceUpgraded)
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
		it.Event = new(MarketplaceUpgraded)
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
func (it *MarketplaceUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceUpgraded represents a Upgraded event raised by the Marketplace contract.
type MarketplaceUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Marketplace *MarketplaceFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*MarketplaceUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceUpgradedIterator{contract: _Marketplace.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Marketplace *MarketplaceFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *MarketplaceUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceUpgraded)
				if err := _Marketplace.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseUpgraded(log types.Log) (*MarketplaceUpgraded, error) {
	event := new(MarketplaceUpgraded)
	if err := _Marketplace.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
