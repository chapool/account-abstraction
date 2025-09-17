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

// PaymentPaymentInfo is an auto generated low-level Go binding around an user-defined struct.
type PaymentPaymentInfo struct {
	OrderId   *big.Int
	Payer     common.Address
	Token     common.Address
	Amount    *big.Int
	Timestamp *big.Int
}

// PaymentMetaData contains all meta data concerning the Payment contract.
var PaymentMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ETHWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"PaymentMade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RefundProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenWithdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"addAllowedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"orderIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"batchRefund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getETHBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"getPayment\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structPayment.PaymentInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"isOrderPaid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isTokenAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"pay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"payments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removeAllowedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAllETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"withdrawAllTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PaymentABI is the input ABI used to generate the binding from.
// Deprecated: Use PaymentMetaData.ABI instead.
var PaymentABI = PaymentMetaData.ABI

// Payment is an auto generated Go binding around an Ethereum contract.
type Payment struct {
	PaymentCaller     // Read-only binding to the contract
	PaymentTransactor // Write-only binding to the contract
	PaymentFilterer   // Log filterer for contract events
}

// PaymentCaller is an auto generated read-only Go binding around an Ethereum contract.
type PaymentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PaymentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PaymentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PaymentSession struct {
	Contract     *Payment          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PaymentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PaymentCallerSession struct {
	Contract *PaymentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// PaymentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PaymentTransactorSession struct {
	Contract     *PaymentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PaymentRaw is an auto generated low-level Go binding around an Ethereum contract.
type PaymentRaw struct {
	Contract *Payment // Generic contract binding to access the raw methods on
}

// PaymentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PaymentCallerRaw struct {
	Contract *PaymentCaller // Generic read-only contract binding to access the raw methods on
}

// PaymentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PaymentTransactorRaw struct {
	Contract *PaymentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPayment creates a new instance of Payment, bound to a specific deployed contract.
func NewPayment(address common.Address, backend bind.ContractBackend) (*Payment, error) {
	contract, err := bindPayment(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Payment{PaymentCaller: PaymentCaller{contract: contract}, PaymentTransactor: PaymentTransactor{contract: contract}, PaymentFilterer: PaymentFilterer{contract: contract}}, nil
}

// NewPaymentCaller creates a new read-only instance of Payment, bound to a specific deployed contract.
func NewPaymentCaller(address common.Address, caller bind.ContractCaller) (*PaymentCaller, error) {
	contract, err := bindPayment(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentCaller{contract: contract}, nil
}

// NewPaymentTransactor creates a new write-only instance of Payment, bound to a specific deployed contract.
func NewPaymentTransactor(address common.Address, transactor bind.ContractTransactor) (*PaymentTransactor, error) {
	contract, err := bindPayment(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentTransactor{contract: contract}, nil
}

// NewPaymentFilterer creates a new log filterer instance of Payment, bound to a specific deployed contract.
func NewPaymentFilterer(address common.Address, filterer bind.ContractFilterer) (*PaymentFilterer, error) {
	contract, err := bindPayment(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PaymentFilterer{contract: contract}, nil
}

// bindPayment binds a generic wrapper to an already deployed contract.
func bindPayment(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PaymentMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Payment *PaymentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Payment.Contract.PaymentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Payment *PaymentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Payment.Contract.PaymentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Payment *PaymentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Payment.Contract.PaymentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Payment *PaymentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Payment.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Payment *PaymentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Payment.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Payment *PaymentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Payment.Contract.contract.Transact(opts, method, params...)
}

// AllowedTokens is a free data retrieval call binding the contract method 0xe744092e.
//
// Solidity: function allowedTokens(address ) view returns(bool)
func (_Payment *PaymentCaller) AllowedTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Payment.contract.Call(opts, &out, "allowedTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedTokens is a free data retrieval call binding the contract method 0xe744092e.
//
// Solidity: function allowedTokens(address ) view returns(bool)
func (_Payment *PaymentSession) AllowedTokens(arg0 common.Address) (bool, error) {
	return _Payment.Contract.AllowedTokens(&_Payment.CallOpts, arg0)
}

// AllowedTokens is a free data retrieval call binding the contract method 0xe744092e.
//
// Solidity: function allowedTokens(address ) view returns(bool)
func (_Payment *PaymentCallerSession) AllowedTokens(arg0 common.Address) (bool, error) {
	return _Payment.Contract.AllowedTokens(&_Payment.CallOpts, arg0)
}

// GetETHBalance is a free data retrieval call binding the contract method 0x6e947298.
//
// Solidity: function getETHBalance() view returns(uint256)
func (_Payment *PaymentCaller) GetETHBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Payment.contract.Call(opts, &out, "getETHBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetETHBalance is a free data retrieval call binding the contract method 0x6e947298.
//
// Solidity: function getETHBalance() view returns(uint256)
func (_Payment *PaymentSession) GetETHBalance() (*big.Int, error) {
	return _Payment.Contract.GetETHBalance(&_Payment.CallOpts)
}

// GetETHBalance is a free data retrieval call binding the contract method 0x6e947298.
//
// Solidity: function getETHBalance() view returns(uint256)
func (_Payment *PaymentCallerSession) GetETHBalance() (*big.Int, error) {
	return _Payment.Contract.GetETHBalance(&_Payment.CallOpts)
}

// GetPayment is a free data retrieval call binding the contract method 0x3280a836.
//
// Solidity: function getPayment(uint256 orderId) view returns((uint256,address,address,uint256,uint256))
func (_Payment *PaymentCaller) GetPayment(opts *bind.CallOpts, orderId *big.Int) (PaymentPaymentInfo, error) {
	var out []interface{}
	err := _Payment.contract.Call(opts, &out, "getPayment", orderId)

	if err != nil {
		return *new(PaymentPaymentInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(PaymentPaymentInfo)).(*PaymentPaymentInfo)

	return out0, err

}

// GetPayment is a free data retrieval call binding the contract method 0x3280a836.
//
// Solidity: function getPayment(uint256 orderId) view returns((uint256,address,address,uint256,uint256))
func (_Payment *PaymentSession) GetPayment(orderId *big.Int) (PaymentPaymentInfo, error) {
	return _Payment.Contract.GetPayment(&_Payment.CallOpts, orderId)
}

// GetPayment is a free data retrieval call binding the contract method 0x3280a836.
//
// Solidity: function getPayment(uint256 orderId) view returns((uint256,address,address,uint256,uint256))
func (_Payment *PaymentCallerSession) GetPayment(orderId *big.Int) (PaymentPaymentInfo, error) {
	return _Payment.Contract.GetPayment(&_Payment.CallOpts, orderId)
}

// GetTokenBalance is a free data retrieval call binding the contract method 0x3aecd0e3.
//
// Solidity: function getTokenBalance(address token) view returns(uint256)
func (_Payment *PaymentCaller) GetTokenBalance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Payment.contract.Call(opts, &out, "getTokenBalance", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenBalance is a free data retrieval call binding the contract method 0x3aecd0e3.
//
// Solidity: function getTokenBalance(address token) view returns(uint256)
func (_Payment *PaymentSession) GetTokenBalance(token common.Address) (*big.Int, error) {
	return _Payment.Contract.GetTokenBalance(&_Payment.CallOpts, token)
}

// GetTokenBalance is a free data retrieval call binding the contract method 0x3aecd0e3.
//
// Solidity: function getTokenBalance(address token) view returns(uint256)
func (_Payment *PaymentCallerSession) GetTokenBalance(token common.Address) (*big.Int, error) {
	return _Payment.Contract.GetTokenBalance(&_Payment.CallOpts, token)
}

// IsOrderPaid is a free data retrieval call binding the contract method 0xa84a357c.
//
// Solidity: function isOrderPaid(uint256 orderId) view returns(bool)
func (_Payment *PaymentCaller) IsOrderPaid(opts *bind.CallOpts, orderId *big.Int) (bool, error) {
	var out []interface{}
	err := _Payment.contract.Call(opts, &out, "isOrderPaid", orderId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOrderPaid is a free data retrieval call binding the contract method 0xa84a357c.
//
// Solidity: function isOrderPaid(uint256 orderId) view returns(bool)
func (_Payment *PaymentSession) IsOrderPaid(orderId *big.Int) (bool, error) {
	return _Payment.Contract.IsOrderPaid(&_Payment.CallOpts, orderId)
}

// IsOrderPaid is a free data retrieval call binding the contract method 0xa84a357c.
//
// Solidity: function isOrderPaid(uint256 orderId) view returns(bool)
func (_Payment *PaymentCallerSession) IsOrderPaid(orderId *big.Int) (bool, error) {
	return _Payment.Contract.IsOrderPaid(&_Payment.CallOpts, orderId)
}

// IsTokenAllowed is a free data retrieval call binding the contract method 0xf9eaee0d.
//
// Solidity: function isTokenAllowed(address token) view returns(bool)
func (_Payment *PaymentCaller) IsTokenAllowed(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _Payment.contract.Call(opts, &out, "isTokenAllowed", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenAllowed is a free data retrieval call binding the contract method 0xf9eaee0d.
//
// Solidity: function isTokenAllowed(address token) view returns(bool)
func (_Payment *PaymentSession) IsTokenAllowed(token common.Address) (bool, error) {
	return _Payment.Contract.IsTokenAllowed(&_Payment.CallOpts, token)
}

// IsTokenAllowed is a free data retrieval call binding the contract method 0xf9eaee0d.
//
// Solidity: function isTokenAllowed(address token) view returns(bool)
func (_Payment *PaymentCallerSession) IsTokenAllowed(token common.Address) (bool, error) {
	return _Payment.Contract.IsTokenAllowed(&_Payment.CallOpts, token)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Payment *PaymentCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Payment.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Payment *PaymentSession) Owner() (common.Address, error) {
	return _Payment.Contract.Owner(&_Payment.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Payment *PaymentCallerSession) Owner() (common.Address, error) {
	return _Payment.Contract.Owner(&_Payment.CallOpts)
}

// Payments is a free data retrieval call binding the contract method 0x87d81789.
//
// Solidity: function payments(uint256 ) view returns(uint256 orderId, address payer, address token, uint256 amount, uint256 timestamp)
func (_Payment *PaymentCaller) Payments(opts *bind.CallOpts, arg0 *big.Int) (struct {
	OrderId   *big.Int
	Payer     common.Address
	Token     common.Address
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _Payment.contract.Call(opts, &out, "payments", arg0)

	outstruct := new(struct {
		OrderId   *big.Int
		Payer     common.Address
		Token     common.Address
		Amount    *big.Int
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OrderId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Payer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Token = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Payments is a free data retrieval call binding the contract method 0x87d81789.
//
// Solidity: function payments(uint256 ) view returns(uint256 orderId, address payer, address token, uint256 amount, uint256 timestamp)
func (_Payment *PaymentSession) Payments(arg0 *big.Int) (struct {
	OrderId   *big.Int
	Payer     common.Address
	Token     common.Address
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	return _Payment.Contract.Payments(&_Payment.CallOpts, arg0)
}

// Payments is a free data retrieval call binding the contract method 0x87d81789.
//
// Solidity: function payments(uint256 ) view returns(uint256 orderId, address payer, address token, uint256 amount, uint256 timestamp)
func (_Payment *PaymentCallerSession) Payments(arg0 *big.Int) (struct {
	OrderId   *big.Int
	Payer     common.Address
	Token     common.Address
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	return _Payment.Contract.Payments(&_Payment.CallOpts, arg0)
}

// AddAllowedToken is a paid mutator transaction binding the contract method 0x4178617f.
//
// Solidity: function addAllowedToken(address token) returns()
func (_Payment *PaymentTransactor) AddAllowedToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "addAllowedToken", token)
}

// AddAllowedToken is a paid mutator transaction binding the contract method 0x4178617f.
//
// Solidity: function addAllowedToken(address token) returns()
func (_Payment *PaymentSession) AddAllowedToken(token common.Address) (*types.Transaction, error) {
	return _Payment.Contract.AddAllowedToken(&_Payment.TransactOpts, token)
}

// AddAllowedToken is a paid mutator transaction binding the contract method 0x4178617f.
//
// Solidity: function addAllowedToken(address token) returns()
func (_Payment *PaymentTransactorSession) AddAllowedToken(token common.Address) (*types.Transaction, error) {
	return _Payment.Contract.AddAllowedToken(&_Payment.TransactOpts, token)
}

// BatchRefund is a paid mutator transaction binding the contract method 0xae273057.
//
// Solidity: function batchRefund(uint256[] orderIds, address[] users, uint256[] amounts, address[] tokens) returns()
func (_Payment *PaymentTransactor) BatchRefund(opts *bind.TransactOpts, orderIds []*big.Int, users []common.Address, amounts []*big.Int, tokens []common.Address) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "batchRefund", orderIds, users, amounts, tokens)
}

// BatchRefund is a paid mutator transaction binding the contract method 0xae273057.
//
// Solidity: function batchRefund(uint256[] orderIds, address[] users, uint256[] amounts, address[] tokens) returns()
func (_Payment *PaymentSession) BatchRefund(orderIds []*big.Int, users []common.Address, amounts []*big.Int, tokens []common.Address) (*types.Transaction, error) {
	return _Payment.Contract.BatchRefund(&_Payment.TransactOpts, orderIds, users, amounts, tokens)
}

// BatchRefund is a paid mutator transaction binding the contract method 0xae273057.
//
// Solidity: function batchRefund(uint256[] orderIds, address[] users, uint256[] amounts, address[] tokens) returns()
func (_Payment *PaymentTransactorSession) BatchRefund(orderIds []*big.Int, users []common.Address, amounts []*big.Int, tokens []common.Address) (*types.Transaction, error) {
	return _Payment.Contract.BatchRefund(&_Payment.TransactOpts, orderIds, users, amounts, tokens)
}

// Pay is a paid mutator transaction binding the contract method 0x0313766f.
//
// Solidity: function pay(uint256 orderId, address token, uint256 amount) payable returns()
func (_Payment *PaymentTransactor) Pay(opts *bind.TransactOpts, orderId *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "pay", orderId, token, amount)
}

// Pay is a paid mutator transaction binding the contract method 0x0313766f.
//
// Solidity: function pay(uint256 orderId, address token, uint256 amount) payable returns()
func (_Payment *PaymentSession) Pay(orderId *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Payment.Contract.Pay(&_Payment.TransactOpts, orderId, token, amount)
}

// Pay is a paid mutator transaction binding the contract method 0x0313766f.
//
// Solidity: function pay(uint256 orderId, address token, uint256 amount) payable returns()
func (_Payment *PaymentTransactorSession) Pay(orderId *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Payment.Contract.Pay(&_Payment.TransactOpts, orderId, token, amount)
}

// RemoveAllowedToken is a paid mutator transaction binding the contract method 0x90469a9d.
//
// Solidity: function removeAllowedToken(address token) returns()
func (_Payment *PaymentTransactor) RemoveAllowedToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "removeAllowedToken", token)
}

// RemoveAllowedToken is a paid mutator transaction binding the contract method 0x90469a9d.
//
// Solidity: function removeAllowedToken(address token) returns()
func (_Payment *PaymentSession) RemoveAllowedToken(token common.Address) (*types.Transaction, error) {
	return _Payment.Contract.RemoveAllowedToken(&_Payment.TransactOpts, token)
}

// RemoveAllowedToken is a paid mutator transaction binding the contract method 0x90469a9d.
//
// Solidity: function removeAllowedToken(address token) returns()
func (_Payment *PaymentTransactorSession) RemoveAllowedToken(token common.Address) (*types.Transaction, error) {
	return _Payment.Contract.RemoveAllowedToken(&_Payment.TransactOpts, token)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Payment *PaymentTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Payment *PaymentSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Payment.Contract.TransferOwnership(&_Payment.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Payment *PaymentTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Payment.Contract.TransferOwnership(&_Payment.TransactOpts, newOwner)
}

// WithdrawAllETH is a paid mutator transaction binding the contract method 0x90386bbf.
//
// Solidity: function withdrawAllETH() returns()
func (_Payment *PaymentTransactor) WithdrawAllETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "withdrawAllETH")
}

// WithdrawAllETH is a paid mutator transaction binding the contract method 0x90386bbf.
//
// Solidity: function withdrawAllETH() returns()
func (_Payment *PaymentSession) WithdrawAllETH() (*types.Transaction, error) {
	return _Payment.Contract.WithdrawAllETH(&_Payment.TransactOpts)
}

// WithdrawAllETH is a paid mutator transaction binding the contract method 0x90386bbf.
//
// Solidity: function withdrawAllETH() returns()
func (_Payment *PaymentTransactorSession) WithdrawAllETH() (*types.Transaction, error) {
	return _Payment.Contract.WithdrawAllETH(&_Payment.TransactOpts)
}

// WithdrawAllTokens is a paid mutator transaction binding the contract method 0xa878aee6.
//
// Solidity: function withdrawAllTokens(address token) returns()
func (_Payment *PaymentTransactor) WithdrawAllTokens(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "withdrawAllTokens", token)
}

// WithdrawAllTokens is a paid mutator transaction binding the contract method 0xa878aee6.
//
// Solidity: function withdrawAllTokens(address token) returns()
func (_Payment *PaymentSession) WithdrawAllTokens(token common.Address) (*types.Transaction, error) {
	return _Payment.Contract.WithdrawAllTokens(&_Payment.TransactOpts, token)
}

// WithdrawAllTokens is a paid mutator transaction binding the contract method 0xa878aee6.
//
// Solidity: function withdrawAllTokens(address token) returns()
func (_Payment *PaymentTransactorSession) WithdrawAllTokens(token common.Address) (*types.Transaction, error) {
	return _Payment.Contract.WithdrawAllTokens(&_Payment.TransactOpts, token)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_Payment *PaymentTransactor) WithdrawETH(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "withdrawETH", amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_Payment *PaymentSession) WithdrawETH(amount *big.Int) (*types.Transaction, error) {
	return _Payment.Contract.WithdrawETH(&_Payment.TransactOpts, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_Payment *PaymentTransactorSession) WithdrawETH(amount *big.Int) (*types.Transaction, error) {
	return _Payment.Contract.WithdrawETH(&_Payment.TransactOpts, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x9e281a98.
//
// Solidity: function withdrawToken(address token, uint256 amount) returns()
func (_Payment *PaymentTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "withdrawToken", token, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x9e281a98.
//
// Solidity: function withdrawToken(address token, uint256 amount) returns()
func (_Payment *PaymentSession) WithdrawToken(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Payment.Contract.WithdrawToken(&_Payment.TransactOpts, token, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x9e281a98.
//
// Solidity: function withdrawToken(address token, uint256 amount) returns()
func (_Payment *PaymentTransactorSession) WithdrawToken(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Payment.Contract.WithdrawToken(&_Payment.TransactOpts, token, amount)
}

// PaymentETHWithdrawnIterator is returned from FilterETHWithdrawn and is used to iterate over the raw logs and unpacked data for ETHWithdrawn events raised by the Payment contract.
type PaymentETHWithdrawnIterator struct {
	Event *PaymentETHWithdrawn // Event containing the contract specifics and raw log

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
func (it *PaymentETHWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentETHWithdrawn)
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
		it.Event = new(PaymentETHWithdrawn)
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
func (it *PaymentETHWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentETHWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentETHWithdrawn represents a ETHWithdrawn event raised by the Payment contract.
type PaymentETHWithdrawn struct {
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterETHWithdrawn is a free log retrieval operation binding the contract event 0x94b2de810873337ed265c5f8cf98c9cffefa06b8607f9a2f1fbaebdfbcfbef1c.
//
// Solidity: event ETHWithdrawn(address indexed owner, uint256 amount)
func (_Payment *PaymentFilterer) FilterETHWithdrawn(opts *bind.FilterOpts, owner []common.Address) (*PaymentETHWithdrawnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Payment.contract.FilterLogs(opts, "ETHWithdrawn", ownerRule)
	if err != nil {
		return nil, err
	}
	return &PaymentETHWithdrawnIterator{contract: _Payment.contract, event: "ETHWithdrawn", logs: logs, sub: sub}, nil
}

// WatchETHWithdrawn is a free log subscription operation binding the contract event 0x94b2de810873337ed265c5f8cf98c9cffefa06b8607f9a2f1fbaebdfbcfbef1c.
//
// Solidity: event ETHWithdrawn(address indexed owner, uint256 amount)
func (_Payment *PaymentFilterer) WatchETHWithdrawn(opts *bind.WatchOpts, sink chan<- *PaymentETHWithdrawn, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Payment.contract.WatchLogs(opts, "ETHWithdrawn", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentETHWithdrawn)
				if err := _Payment.contract.UnpackLog(event, "ETHWithdrawn", log); err != nil {
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

// ParseETHWithdrawn is a log parse operation binding the contract event 0x94b2de810873337ed265c5f8cf98c9cffefa06b8607f9a2f1fbaebdfbcfbef1c.
//
// Solidity: event ETHWithdrawn(address indexed owner, uint256 amount)
func (_Payment *PaymentFilterer) ParseETHWithdrawn(log types.Log) (*PaymentETHWithdrawn, error) {
	event := new(PaymentETHWithdrawn)
	if err := _Payment.contract.UnpackLog(event, "ETHWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Payment contract.
type PaymentOwnershipTransferredIterator struct {
	Event *PaymentOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PaymentOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentOwnershipTransferred)
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
		it.Event = new(PaymentOwnershipTransferred)
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
func (it *PaymentOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentOwnershipTransferred represents a OwnershipTransferred event raised by the Payment contract.
type PaymentOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Payment *PaymentFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PaymentOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Payment.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PaymentOwnershipTransferredIterator{contract: _Payment.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Payment *PaymentFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PaymentOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Payment.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentOwnershipTransferred)
				if err := _Payment.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Payment *PaymentFilterer) ParseOwnershipTransferred(log types.Log) (*PaymentOwnershipTransferred, error) {
	event := new(PaymentOwnershipTransferred)
	if err := _Payment.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentPaymentMadeIterator is returned from FilterPaymentMade and is used to iterate over the raw logs and unpacked data for PaymentMade events raised by the Payment contract.
type PaymentPaymentMadeIterator struct {
	Event *PaymentPaymentMade // Event containing the contract specifics and raw log

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
func (it *PaymentPaymentMadeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentPaymentMade)
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
		it.Event = new(PaymentPaymentMade)
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
func (it *PaymentPaymentMadeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentPaymentMadeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentPaymentMade represents a PaymentMade event raised by the Payment contract.
type PaymentPaymentMade struct {
	OrderId   *big.Int
	Payer     common.Address
	Token     common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPaymentMade is a free log retrieval operation binding the contract event 0x32aced27dfd49efcd31ceb0567a1ef533d2ab1481334c3f316047bf16fe1c8e8.
//
// Solidity: event PaymentMade(uint256 indexed orderId, address indexed payer, address indexed token, uint256 amount, uint256 timestamp)
func (_Payment *PaymentFilterer) FilterPaymentMade(opts *bind.FilterOpts, orderId []*big.Int, payer []common.Address, token []common.Address) (*PaymentPaymentMadeIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Payment.contract.FilterLogs(opts, "PaymentMade", orderIdRule, payerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PaymentPaymentMadeIterator{contract: _Payment.contract, event: "PaymentMade", logs: logs, sub: sub}, nil
}

// WatchPaymentMade is a free log subscription operation binding the contract event 0x32aced27dfd49efcd31ceb0567a1ef533d2ab1481334c3f316047bf16fe1c8e8.
//
// Solidity: event PaymentMade(uint256 indexed orderId, address indexed payer, address indexed token, uint256 amount, uint256 timestamp)
func (_Payment *PaymentFilterer) WatchPaymentMade(opts *bind.WatchOpts, sink chan<- *PaymentPaymentMade, orderId []*big.Int, payer []common.Address, token []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Payment.contract.WatchLogs(opts, "PaymentMade", orderIdRule, payerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentPaymentMade)
				if err := _Payment.contract.UnpackLog(event, "PaymentMade", log); err != nil {
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

// ParsePaymentMade is a log parse operation binding the contract event 0x32aced27dfd49efcd31ceb0567a1ef533d2ab1481334c3f316047bf16fe1c8e8.
//
// Solidity: event PaymentMade(uint256 indexed orderId, address indexed payer, address indexed token, uint256 amount, uint256 timestamp)
func (_Payment *PaymentFilterer) ParsePaymentMade(log types.Log) (*PaymentPaymentMade, error) {
	event := new(PaymentPaymentMade)
	if err := _Payment.contract.UnpackLog(event, "PaymentMade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentRefundProcessedIterator is returned from FilterRefundProcessed and is used to iterate over the raw logs and unpacked data for RefundProcessed events raised by the Payment contract.
type PaymentRefundProcessedIterator struct {
	Event *PaymentRefundProcessed // Event containing the contract specifics and raw log

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
func (it *PaymentRefundProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentRefundProcessed)
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
		it.Event = new(PaymentRefundProcessed)
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
func (it *PaymentRefundProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentRefundProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentRefundProcessed represents a RefundProcessed event raised by the Payment contract.
type PaymentRefundProcessed struct {
	OrderId *big.Int
	User    common.Address
	Token   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRefundProcessed is a free log retrieval operation binding the contract event 0x4d60a9438ba7e18c1fed7577dc8932bfe82f683c1e254a5336b6618ab5301641.
//
// Solidity: event RefundProcessed(uint256 indexed orderId, address indexed user, address indexed token, uint256 amount)
func (_Payment *PaymentFilterer) FilterRefundProcessed(opts *bind.FilterOpts, orderId []*big.Int, user []common.Address, token []common.Address) (*PaymentRefundProcessedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Payment.contract.FilterLogs(opts, "RefundProcessed", orderIdRule, userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PaymentRefundProcessedIterator{contract: _Payment.contract, event: "RefundProcessed", logs: logs, sub: sub}, nil
}

// WatchRefundProcessed is a free log subscription operation binding the contract event 0x4d60a9438ba7e18c1fed7577dc8932bfe82f683c1e254a5336b6618ab5301641.
//
// Solidity: event RefundProcessed(uint256 indexed orderId, address indexed user, address indexed token, uint256 amount)
func (_Payment *PaymentFilterer) WatchRefundProcessed(opts *bind.WatchOpts, sink chan<- *PaymentRefundProcessed, orderId []*big.Int, user []common.Address, token []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Payment.contract.WatchLogs(opts, "RefundProcessed", orderIdRule, userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentRefundProcessed)
				if err := _Payment.contract.UnpackLog(event, "RefundProcessed", log); err != nil {
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

// ParseRefundProcessed is a log parse operation binding the contract event 0x4d60a9438ba7e18c1fed7577dc8932bfe82f683c1e254a5336b6618ab5301641.
//
// Solidity: event RefundProcessed(uint256 indexed orderId, address indexed user, address indexed token, uint256 amount)
func (_Payment *PaymentFilterer) ParseRefundProcessed(log types.Log) (*PaymentRefundProcessed, error) {
	event := new(PaymentRefundProcessed)
	if err := _Payment.contract.UnpackLog(event, "RefundProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentTokenAddedIterator is returned from FilterTokenAdded and is used to iterate over the raw logs and unpacked data for TokenAdded events raised by the Payment contract.
type PaymentTokenAddedIterator struct {
	Event *PaymentTokenAdded // Event containing the contract specifics and raw log

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
func (it *PaymentTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentTokenAdded)
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
		it.Event = new(PaymentTokenAdded)
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
func (it *PaymentTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentTokenAdded represents a TokenAdded event raised by the Payment contract.
type PaymentTokenAdded struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenAdded is a free log retrieval operation binding the contract event 0x784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4.
//
// Solidity: event TokenAdded(address indexed token)
func (_Payment *PaymentFilterer) FilterTokenAdded(opts *bind.FilterOpts, token []common.Address) (*PaymentTokenAddedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Payment.contract.FilterLogs(opts, "TokenAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return &PaymentTokenAddedIterator{contract: _Payment.contract, event: "TokenAdded", logs: logs, sub: sub}, nil
}

// WatchTokenAdded is a free log subscription operation binding the contract event 0x784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4.
//
// Solidity: event TokenAdded(address indexed token)
func (_Payment *PaymentFilterer) WatchTokenAdded(opts *bind.WatchOpts, sink chan<- *PaymentTokenAdded, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Payment.contract.WatchLogs(opts, "TokenAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentTokenAdded)
				if err := _Payment.contract.UnpackLog(event, "TokenAdded", log); err != nil {
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

// ParseTokenAdded is a log parse operation binding the contract event 0x784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4.
//
// Solidity: event TokenAdded(address indexed token)
func (_Payment *PaymentFilterer) ParseTokenAdded(log types.Log) (*PaymentTokenAdded, error) {
	event := new(PaymentTokenAdded)
	if err := _Payment.contract.UnpackLog(event, "TokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentTokenRemovedIterator is returned from FilterTokenRemoved and is used to iterate over the raw logs and unpacked data for TokenRemoved events raised by the Payment contract.
type PaymentTokenRemovedIterator struct {
	Event *PaymentTokenRemoved // Event containing the contract specifics and raw log

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
func (it *PaymentTokenRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentTokenRemoved)
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
		it.Event = new(PaymentTokenRemoved)
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
func (it *PaymentTokenRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentTokenRemoved represents a TokenRemoved event raised by the Payment contract.
type PaymentTokenRemoved struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenRemoved is a free log retrieval operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address indexed token)
func (_Payment *PaymentFilterer) FilterTokenRemoved(opts *bind.FilterOpts, token []common.Address) (*PaymentTokenRemovedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Payment.contract.FilterLogs(opts, "TokenRemoved", tokenRule)
	if err != nil {
		return nil, err
	}
	return &PaymentTokenRemovedIterator{contract: _Payment.contract, event: "TokenRemoved", logs: logs, sub: sub}, nil
}

// WatchTokenRemoved is a free log subscription operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address indexed token)
func (_Payment *PaymentFilterer) WatchTokenRemoved(opts *bind.WatchOpts, sink chan<- *PaymentTokenRemoved, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Payment.contract.WatchLogs(opts, "TokenRemoved", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentTokenRemoved)
				if err := _Payment.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
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

// ParseTokenRemoved is a log parse operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address indexed token)
func (_Payment *PaymentFilterer) ParseTokenRemoved(log types.Log) (*PaymentTokenRemoved, error) {
	event := new(PaymentTokenRemoved)
	if err := _Payment.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentTokenWithdrawnIterator is returned from FilterTokenWithdrawn and is used to iterate over the raw logs and unpacked data for TokenWithdrawn events raised by the Payment contract.
type PaymentTokenWithdrawnIterator struct {
	Event *PaymentTokenWithdrawn // Event containing the contract specifics and raw log

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
func (it *PaymentTokenWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentTokenWithdrawn)
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
		it.Event = new(PaymentTokenWithdrawn)
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
func (it *PaymentTokenWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentTokenWithdrawn represents a TokenWithdrawn event raised by the Payment contract.
type PaymentTokenWithdrawn struct {
	Owner  common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdrawn is a free log retrieval operation binding the contract event 0x8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e5620.
//
// Solidity: event TokenWithdrawn(address indexed owner, address indexed token, uint256 amount)
func (_Payment *PaymentFilterer) FilterTokenWithdrawn(opts *bind.FilterOpts, owner []common.Address, token []common.Address) (*PaymentTokenWithdrawnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Payment.contract.FilterLogs(opts, "TokenWithdrawn", ownerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PaymentTokenWithdrawnIterator{contract: _Payment.contract, event: "TokenWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokenWithdrawn is a free log subscription operation binding the contract event 0x8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e5620.
//
// Solidity: event TokenWithdrawn(address indexed owner, address indexed token, uint256 amount)
func (_Payment *PaymentFilterer) WatchTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *PaymentTokenWithdrawn, owner []common.Address, token []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Payment.contract.WatchLogs(opts, "TokenWithdrawn", ownerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentTokenWithdrawn)
				if err := _Payment.contract.UnpackLog(event, "TokenWithdrawn", log); err != nil {
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

// ParseTokenWithdrawn is a log parse operation binding the contract event 0x8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e5620.
//
// Solidity: event TokenWithdrawn(address indexed owner, address indexed token, uint256 amount)
func (_Payment *PaymentFilterer) ParseTokenWithdrawn(log types.Log) (*PaymentTokenWithdrawn, error) {
	event := new(PaymentTokenWithdrawn)
	if err := _Payment.contract.UnpackLog(event, "TokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
