# Payment Contract Deployment Guide

## Overview

The Payment contract is a simple payment system for offline orders that supports ETH and whitelisted ERC20 tokens. This document provides deployment information and usage instructions for the Payment contract deployed on Sepolia testnet.

## Contract Information

### Deployment Details
- **Network**: Sepolia Testnet
- **Payment Contract Address**: `0x7Da90c3364a3D0B99fCD18c1651E669C4D294D65`
- **MockUSDT Contract Address**: `0x0D3E58b48Ef96A1eAF2CD9A553558FFAf490Eb88`
- **Deployer**: `0xa3B605fB633AD0A0DC4B74b10bBfc40fDB050d35`
- **Deployment Date**: 2025-09-17T07:16:03.603Z

### Transaction Hashes
- MockUSDT Deploy: `0x4bb7e1c1104e166113600ffa8e4f4e853ee334b4626c292cb694bd6b96ca6916`
- Payment Deploy: `0x11c71687b5813f31d4074e9920df5c378874f31adcdae81f659a479b4986d1a8`
- Add Token to Whitelist: `0x06c42208ad5ce7f4295446aa4a46eeed020b37557becda794a5ffb9746f1b9a5`

## MockUSDT Token Information
- **Name**: Mock USDT
- **Symbol**: mUSDT
- **Decimals**: 6
- **Total Supply**: 1,000,000 mUSDT
- **Faucet Function**: Available for test users

## Contract Features

### Core Payment Functions
1. **Payment Processing**: `pay(uint256 orderId, address token, uint256 amount)`
2. **Token Management**: Add/remove tokens from whitelist
3. **Withdrawal Functions**: Withdraw ETH and ERC20 tokens
4. **Query Functions**: Check payment status, balances, etc.

### New Batch Refund Feature
The contract includes a new batch refund functionality that allows the owner to process multiple refunds in a single transaction.

#### Function Signature
```solidity
function batchRefund(
    uint256[] calldata orderIds,
    address[] calldata users,
    uint256[] calldata amounts,
    address[] calldata tokens
) external onlyOwner nonReentrant
```

#### Parameters
- `orderIds`: Array of order IDs (for tracking purposes)
- `users`: Array of user addresses to refund
- `amounts`: Array of refund amounts
- `tokens`: Array of token addresses (use `address(0)` for ETH)

#### Security Features
- Only contract owner can execute
- Reentrancy protection
- Array length validation
- Token whitelist validation
- Sufficient balance checks

#### Events
```solidity
event RefundProcessed(
    uint256 indexed orderId,
    address indexed user,
    address indexed token,
    uint256 amount
);
```

## Usage Instructions

### 1. Get Test USDT
```solidity
// Call faucet to get 1000 USDT
await mockUSDT.faucet();
```

### 2. Make a Payment
```solidity
// Pay with ETH
await payment.pay(orderId, ethers.constants.AddressZero, ethers.utils.parseEther("1"), {
    value: ethers.utils.parseEther("1")
});

// Pay with USDT
await payment.pay(orderId, mockUSDTAddress, ethers.utils.parseUnits("100", 6));
```

### 3. Check Payment Status
```solidity
const paymentInfo = await payment.getPayment(orderId);
console.log("Payment Info:", paymentInfo);
```

### 4. Batch Refund Example
```solidity
// Prepare refund data
const orderIds = [1, 2, 3];
const users = [user1, user2, user3];
const amounts = [
    ethers.utils.parseEther("1"),           // 1 ETH
    ethers.utils.parseUnits("100", 6),      // 100 USDT
    ethers.utils.parseUnits("50", 6)        // 50 USDT
];
const tokens = [
    ethers.constants.AddressZero,           // ETH
    mockUSDTAddress,                        // USDT
    mockUSDTAddress                         // USDT
];

// Execute batch refund
await payment.batchRefund(orderIds, users, amounts, tokens);
```

## Contract Management

### Owner Functions
- `addAllowedToken(address token)`: Add token to whitelist
- `removeAllowedToken(address token)`: Remove token from whitelist
- `withdrawETH(uint256 amount)`: Withdraw ETH
- `withdrawAllETH()`: Withdraw all ETH
- `withdrawToken(address token, uint256 amount)`: Withdraw specific tokens
- `withdrawAllTokens(address token)`: Withdraw all tokens of a type
- `transferOwnership(address newOwner)`: Transfer ownership
- `batchRefund(...)`: Process batch refunds

### Query Functions
- `getPayment(uint256 orderId)`: Get payment information
- `isTokenAllowed(address token)`: Check if token is whitelisted
- `isOrderPaid(uint256 orderId)`: Check if order is paid
- `getETHBalance()`: Get contract ETH balance
- `getTokenBalance(address token)`: Get contract token balance

## Security Considerations

1. **Access Control**: Only contract owner can manage tokens and process refunds
2. **Reentrancy Protection**: All state-changing functions use `nonReentrant` modifier
3. **Token Whitelist**: Only whitelisted tokens can be used for payments
4. **Input Validation**: Comprehensive validation of all function parameters
5. **Balance Checks**: Ensures sufficient balance before processing refunds

## Error Codes

- `OnlyOwner()`: Function called by non-owner
- `TokenNotAllowed()`: Token not in whitelist
- `OrderAlreadyPaid()`: Order already has a payment
- `InvalidAmount()`: Invalid amount (zero or exceeds balance)
- `InvalidToken()`: Invalid token address
- `InsufficientBalance()`: Insufficient contract balance
- `TransferFailed()`: Transfer operation failed
- `InvalidOwner()`: Invalid owner address
- `ArrayLengthMismatch()`: Array lengths don't match
- `InvalidUser()`: Invalid user address
- `RefundFailed()`: Refund operation failed

## Testing

The contract has been deployed and tested on Sepolia testnet. You can interact with it using:

1. **Ethers.js**: Use the contract ABI and address
2. **Hardhat**: Use the deployment scripts
3. **Remix**: Import the contract and connect to Sepolia
4. **Web3.js**: Standard web3 interaction

## Support

For questions or issues related to the Payment contract deployment, please refer to the contract source code or contact the development team.

---

**Last Updated**: 2025-09-17T07:16:03.603Z
**Contract Version**: Latest with batch refund functionality
