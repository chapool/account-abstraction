# Payment Contract Deployment Guide

## 📋 Contract Information

### Sepolia Testnet Deployment
- **Payment Contract**: `0xb047a38e85A6C4915B36eaB2f77ebF08488c7257`
- **MockUSDT Contract**: `0xa6Ce43c76dEc68D7DA4DBCeEA6513CcfE3e65AeE`
- **Network**: Sepolia Testnet (Chain ID: 11155111)
- **Deployer**: `0xa3B605fB633AD0A0DC4B74b10bBfc40fDB050d35`

## 🔧 Contract Features

### Payment Contract
- Supports ETH and whitelisted ERC20 tokens
- Simple payment flow: create order offline → pay online → withdraw funds
- Owner-controlled token whitelist
- Secure fund management with withdrawal functions

### MockUSDT (Test Token)
- **Name**: Mock USDT
- **Symbol**: mUSDT  
- **Decimals**: 6
- **Faucet**: Anyone can get 1000 USDT for testing

## 🚀 Quick Start

### 1. Get Test USDT
```javascript
// Call faucet function to get 1000 USDT
const mockUSDT = new ethers.Contract(
  "0xa6Ce43c76dEc68D7DA4DBCeEA6513CcfE3e65AeE",
  mockUSDTABI,
  signer
);
await mockUSDT.faucet();
```

### 2. Approve Payment Contract
```javascript
// Approve Payment contract to spend your USDT
const amount = ethers.utils.parseUnits("100", 6); // 100 USDT
await mockUSDT.approve(
  "0xb047a38e85A6C4915B36eaB2f77ebF08488c7257", 
  amount
);
```

### 3. Make Payment
```javascript
const payment = new ethers.Contract(
  "0xb047a38e85A6C4915B36eaB2f77ebF08488c7257",
  paymentABI,
  signer
);

// Pay with USDT
const orderId = 12345; // Your order ID (generated offline)
const tokenAddress = "0xa6Ce43c76dEc68D7DA4DBCeEA6513CcfE3e65AeE"; // MockUSDT
const amount = ethers.utils.parseUnits("100", 6); // 100 USDT

await payment.pay(orderId, tokenAddress, amount);

// Pay with ETH
await payment.pay(orderId, ethers.constants.AddressZero, amount, {
  value: amount
});
```

### 4. Check Payment Status
```javascript
const paymentInfo = await payment.getPayment(orderId);
console.log("Payer:", paymentInfo.payer);
console.log("Amount:", paymentInfo.amount);
console.log("Token:", paymentInfo.token);
console.log("Timestamp:", paymentInfo.timestamp);

// Check if order is paid
const isPaid = await payment.isOrderPaid(orderId);
```

## 📝 API Reference

### Payment Contract Methods

#### Core Functions
- `pay(uint256 orderId, address token, uint256 amount)` - Pay for an order
- `getPayment(uint256 orderId)` - Get payment information
- `isOrderPaid(uint256 orderId)` - Check if order is paid
- `isTokenAllowed(address token)` - Check if token is whitelisted

#### Query Functions
- `getETHBalance()` - Get contract ETH balance
- `getTokenBalance(address token)` - Get contract token balance
- `owner()` - Get contract owner

#### Owner Functions
- `addAllowedToken(address token)` - Add token to whitelist
- `removeAllowedToken(address token)` - Remove token from whitelist
- `withdrawETH(uint256 amount)` - Withdraw ETH
- `withdrawToken(address token, uint256 amount)` - Withdraw tokens

### MockUSDT Methods
- `faucet()` - Get 1000 USDT for testing
- `balanceOf(address account)` - Check balance
- `approve(address spender, uint256 amount)` - Approve spending
- `transfer(address to, uint256 amount)` - Transfer tokens

## 🧪 Testing Examples

### JavaScript (ethers.js)
```javascript
// Initialize contracts
const provider = new ethers.providers.JsonRpcProvider("https://sepolia.infura.io/v3/YOUR_KEY");
const wallet = new ethers.Wallet("YOUR_PRIVATE_KEY", provider);

const payment = new ethers.Contract(PAYMENT_ADDRESS, paymentABI, wallet);
const mockUSDT = new ethers.Contract(MOCKUSDT_ADDRESS, mockUSDTABI, wallet);

// Test flow
async function testPayment() {
  // 1. Get test USDT
  await mockUSDT.faucet();
  console.log("USDT Balance:", await mockUSDT.balanceOf(wallet.address));
  
  // 2. Approve payment contract
  const amount = ethers.utils.parseUnits("50", 6);
  await mockUSDT.approve(PAYMENT_ADDRESS, amount);
  
  // 3. Make payment
  const orderId = Date.now(); // Use timestamp as order ID
  await payment.pay(orderId, MOCKUSDT_ADDRESS, amount);
  
  // 4. Verify payment
  const paymentInfo = await payment.getPayment(orderId);
  console.log("Payment successful:", paymentInfo);
}
```

### React/Web3 Integration
```jsx
import { ethers } from 'ethers';

const PaymentComponent = () => {
  const payForOrder = async (orderId, amount, tokenType) => {
    const provider = new ethers.providers.Web3Provider(window.ethereum);
    const signer = provider.getSigner();
    
    const payment = new ethers.Contract(PAYMENT_ADDRESS, paymentABI, signer);
    
    if (tokenType === 'ETH') {
      await payment.pay(orderId, ethers.constants.AddressZero, amount, {
        value: amount
      });
    } else if (tokenType === 'USDT') {
      const mockUSDT = new ethers.Contract(MOCKUSDT_ADDRESS, mockUSDTABI, signer);
      await mockUSDT.approve(PAYMENT_ADDRESS, amount);
      await payment.pay(orderId, MOCKUSDT_ADDRESS, amount);
    }
  };
  
  return (
    <button onClick={() => payForOrder(12345, ethers.utils.parseUnits("100", 6), "USDT")}>
      Pay with USDT
    </button>
  );
};
```

## 📄 ABI Files

ABI files are available in the `cpop-abis` directory:
- `Payment.abi.json` - Payment contract ABI
- `MockUSDT.abi.json` - MockUSDT contract ABI

## ⚠️ Important Notes

1. **Test Network Only**: These contracts are deployed on Sepolia testnet for testing
2. **Order IDs**: Generate unique order IDs in your backend/frontend
3. **Token Approval**: Always approve before paying with ERC20 tokens
4. **Gas Fees**: Keep some ETH for transaction gas fees
5. **Error Handling**: Implement proper error handling for failed transactions

## 🔍 Etherscan Links

- [Payment Contract](https://sepolia.etherscan.io/address/0xb047a38e85A6C4915B36eaB2f77ebF08488c7257)
- [MockUSDT Contract](https://sepolia.etherscan.io/address/0xa6Ce43c76dEc68D7DA4DBCeEA6513CcfE3e65AeE)

## 💡 Support

If you encounter any issues or need assistance:
1. Check transaction status on [Sepolia Etherscan](https://sepolia.etherscan.io)
2. Verify contract addresses and function calls
3. Ensure sufficient token balance and allowance
4. Contact the development team for technical support

---

**Last Updated**: August 30, 2025  
**Version**: 1.0.0