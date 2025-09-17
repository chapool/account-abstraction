# Payment Contract Usage Examples

## 概述

本文档提供了Payment合约的使用示例，包括如何支付ETH和MockUSDT。

## 合约信息

- **Payment合约地址**: `0x7Da90c3364a3D0B99fCD18c1651E669C4D294D65`
- **MockUSDT合约地址**: `0x0D3E58b48Ef96A1eAF2CD9A553558FFAf490Eb88`
- **网络**: Sepolia Testnet

## 使用脚本

### 1. 完整支付脚本 (pay-with-mockusdt.ts)

包含完整验证和错误处理的支付脚本：

```bash
npx hardhat run scripts/pay-with-mockusdt.ts --network sepoliaCustom
```

**功能特性**:
- 自动检查代币白名单状态
- 检查订单是否已支付
- 自动从水龙头获取USDT
- 自动批准代币消费
- 验证支付结果
- 保存支付记录

### 2. 简单支付脚本 (simple-payment.ts)

简化版的支付脚本，适合快速测试：

```bash
npx hardhat run scripts/simple-payment.ts --network sepoliaCustom
```

**功能特性**:
- 基本的支付流程
- 自动处理余额和授权
- 简化的错误处理

## 支付示例

### 支付1000 MockUSDT

```typescript
// 参数设置
const ORDER_ID = 1;
const AMOUNT_USDT = 1000;
const MOCK_USDT_ADDRESS = "0x0D3E58b48Ef96A1eAF2CD9A553558FFAf490Eb88";

// 转换金额 (USDT有6位小数)
const amount = ethers.utils.parseUnits(AMOUNT_USDT.toString(), 6);

// 执行支付
const tx = await payment.pay(ORDER_ID, MOCK_USDT_ADDRESS, amount);
await tx.wait();
```

### 支付ETH

```typescript
// 参数设置
const ORDER_ID = 2;
const AMOUNT_ETH = "1.0";

// 转换金额
const amount = ethers.utils.parseEther(AMOUNT_ETH);

// 执行支付 (需要发送ETH)
const tx = await payment.pay(ORDER_ID, ethers.constants.AddressZero, amount, {
    value: amount
});
await tx.wait();
```

## 成功支付记录

### 最近一次支付 (2025-09-17)

```json
{
  "orderId": 1,
  "amount": "1000.0 USDT",
  "payer": "0xa3B605fB633AD0A0DC4B74b10bBfc40fDB050d35",
  "transaction": "0x9efbd3104957c6edde9c33ecd192525be72ab72234af04d02d05a9bd7feeb0d9",
  "blockNumber": 9221048,
  "gasUsed": "171063"
}
```

## 查询功能

### 检查订单支付状态

```typescript
const isPaid = await payment.isOrderPaid(orderId);
console.log("Order paid:", isPaid);
```

### 获取支付信息

```typescript
const paymentInfo = await payment.getPayment(orderId);
console.log("Payment Info:", {
    orderId: paymentInfo.orderId.toString(),
    payer: paymentInfo.payer,
    token: paymentInfo.token,
    amount: ethers.utils.formatUnits(paymentInfo.amount, 6),
    timestamp: new Date(paymentInfo.timestamp.toNumber() * 1000)
});
```

### 检查代币余额

```typescript
// 检查合约ETH余额
const ethBalance = await payment.getETHBalance();
console.log("Contract ETH balance:", ethers.utils.formatEther(ethBalance));

// 检查合约USDT余额
const usdtBalance = await payment.getTokenBalance(MOCK_USDT_ADDRESS);
console.log("Contract USDT balance:", ethers.utils.formatUnits(usdtBalance, 6));
```

## 错误处理

### 常见错误

1. **OrderAlreadyPaid**: 订单已支付
   ```typescript
   if (await payment.isOrderPaid(orderId)) {
       console.log("Order already paid!");
       return;
   }
   ```

2. **TokenNotAllowed**: 代币不在白名单
   ```typescript
   const isAllowed = await payment.isTokenAllowed(tokenAddress);
   if (!isAllowed) {
       console.log("Token not allowed!");
       return;
   }
   ```

3. **InsufficientBalance**: 余额不足
   ```typescript
   const balance = await mockUSDT.balanceOf(userAddress);
   if (balance.lt(amount)) {
       console.log("Insufficient balance!");
       return;
   }
   ```

## 批量操作

### 批量退款 (仅合约所有者)

```typescript
const orderIds = [1, 2, 3];
const users = [user1, user2, user3];
const amounts = [
    ethers.utils.parseUnits("100", 6),  // 100 USDT
    ethers.utils.parseUnits("200", 6),  // 200 USDT
    ethers.utils.parseEther("1")        // 1 ETH
];
const tokens = [MOCK_USDT_ADDRESS, MOCK_USDT_ADDRESS, ethers.constants.AddressZero];

const tx = await payment.batchRefund(orderIds, users, amounts, tokens);
await tx.wait();
```

## 网络配置

### Sepolia网络配置

```typescript
// hardhat.config.ts
sepoliaCustom: {
    url: process.env.SEPOLIA_RPC_URL,
    accounts: [process.env.PRIVATE_KEY],
    chainId: 11155111
}
```

### 环境变量

```bash
SEPOLIA_RPC_URL=https://sepolia.infura.io/v3/YOUR_PROJECT_ID
PRIVATE_KEY=0x...
```

## 最佳实践

1. **订单ID管理**: 使用唯一的订单ID，避免重复
2. **金额验证**: 确保金额格式正确 (ETH: 18位小数, USDT: 6位小数)
3. **授权检查**: 支付ERC20代币前先检查授权
4. **错误处理**: 添加适当的错误处理和重试机制
5. **Gas估算**: 预估Gas费用，避免交易失败

## 监控和日志

### 事件监听

```typescript
// 监听支付事件
payment.on("PaymentMade", (orderId, payer, token, amount, timestamp) => {
    console.log("Payment received:", {
        orderId: orderId.toString(),
        payer,
        token,
        amount: ethers.utils.formatUnits(amount, 6),
        timestamp: new Date(timestamp.toNumber() * 1000)
    });
});

// 监听退款事件
payment.on("RefundProcessed", (orderId, user, token, amount) => {
    console.log("Refund processed:", {
        orderId: orderId.toString(),
        user,
        token,
        amount: ethers.utils.formatUnits(amount, 6)
    });
});
```

---

**最后更新**: 2025-09-17  
**版本**: 1.0
