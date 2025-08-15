# CPOP Gas Oracle Integration Guide

## 概览

本文档介绍了基于Oracle机制的CPOP代币Gas费用计算系统，该系统支持动态价格获取、链下计算和后期扩展。

## 系统架构

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Price Feeds   │───▶│ GasPriceOracle  │───▶│ CPOPPaymasterV2 │
│  (ETH/USD,      │    │                 │    │                 │
│   CPOP/USD)     │    │                 │    │                 │
└─────────────────┘    └─────────────────┘    └─────────────────┘
                                │                       │
                                ▼                       ▼
                    ┌─────────────────┐    ┌─────────────────┐
                    │ Off-chain Gas   │    │   User Account  │
                    │   Calculator    │    │   (CPOPAccount) │
                    └─────────────────┘    └─────────────────┘
```

## 核心组件

### 1. IGasPriceOracle Interface
Oracle价格接口，定义了价格获取和验证的标准方法。

**主要功能:**
- 获取ETH/USD和CPOP/USD价格
- 计算ETH与CPOP的兑换比率
- 估算Gas费用的CPOP成本
- 价格有效性验证

### 2. GasPriceOracle Contract
实现IGasPriceOracle接口的Oracle合约。

**核心特性:**
- 多价格源支持
- 价格偏差保护
- 回退机制
- 可升级设计

**配置参数:**
```solidity
uint256 public maxPriceAge = 3600;        // 1小时价格有效期
uint256 public priceDeviationThreshold = 500; // 5%偏差阈值
```

### 3. CPOPPaymasterV2 Contract
集成Oracle的增强版Paymaster。

**新增功能:**
- Oracle价格集成
- 回退汇率机制
- 健康状态监控
- 详细成本分析

### 4. Off-chain Gas Calculator
链下Gas费用计算工具，支持实时价格查询和成本优化。

## 部署指南

### 1. 部署Oracle合约

```javascript
// 部署脚本示例
const GasPriceOracle = await ethers.getContractFactory("GasPriceOracle");
const oracle = await upgrades.deployProxy(GasPriceOracle, [
    owner.address,           // 所有者
    3600,                   // 最大价格年龄 (1小时)
    500                     // 价格偏差阈值 (5%)
]);

// 设置初始价格
await oracle.addPriceFeed("ETH/USD", ethers.parseUnits("2000", 8), ethers.parseUnits("1800", 8));
await oracle.addPriceFeed("CPOP/USD", ethers.parseUnits("0.05", 8), ethers.parseUnits("0.04", 8));
```

### 2. 部署增强版Paymaster

```javascript
const CPOPPaymasterV2 = await ethers.getContractFactory("CPOPPaymasterV2");
const paymaster = await CPOPPaymasterV2.deploy(
    entryPoint.address,      // EntryPoint地址
    cpopToken.address,       // CPOP代币地址
    oracle.address,          // Oracle地址
    ethers.parseUnits("50", 18) // 回退汇率 (50 CPOP per ETH)
);
```

### 3. 配置Oracle授权

```javascript
// 授权价格更新者
await oracle.authorizeOracle(priceUpdater.address);

// 设置回退价格
await oracle.setFallbackPrice("ETH/USD", ethers.parseUnits("1800", 8));
await oracle.setFallbackPrice("CPOP/USD", ethers.parseUnits("0.04", 8));
```

## 链下Gas计算使用

### 1. 初始化计算器

```javascript
const { CPOPGasCalculator } = require('./scripts/gas-calculator');

const calculator = new CPOPGasCalculator(
    provider,
    oracleAddress,
    oracleABI
);
```

### 2. 估算用户操作成本

```javascript
const userOp = {
    callGasLimit: '100000',
    verificationGasLimit: '150000',
    preVerificationGas: '21000',
    sender: userAddress
};

const costEstimate = await calculator.estimateUserOpCost(userOp);
console.log('Total CPOP Cost:', costEstimate.costs.cpop);
```

### 3. 检查用户余额

```javascript
const balanceCheck = await calculator.checkUserBalance(
    userAddress,
    userOp,
    cpopTokenContract
);

if (!balanceCheck.hasSufficientBalance) {
    console.log('Insufficient CPOP balance');
    console.log('Shortfall:', balanceCheck.shortfall);
}
```

### 4. 优化Gas参数

```javascript
const optimizations = await calculator.optimizeGasParameters(transaction);
console.log('Recommended:', optimizations.recommended);
```

## Oracle价格更新

### 1. 手动更新

```javascript
// 更新ETH价格
await oracle.updatePriceFeed("ETH/USD", ethers.parseUnits("2100", 8));

// 更新CPOP价格
await oracle.updatePriceFeed("CPOP/USD", ethers.parseUnits("0.055", 8));
```

### 2. 自动化更新脚本

```javascript
// 价格更新示例
async function updatePrices() {
    try {
        // 从外部API获取价格
        const ethPrice = await fetchETHPrice();
        const cpopPrice = await fetchCPOPPrice();
        
        // 更新Oracle
        await oracle.updatePriceFeed("ETH/USD", ethers.parseUnits(ethPrice.toString(), 8));
        await oracle.updatePriceFeed("CPOP/USD", ethers.parseUnits(cpopPrice.toString(), 8));
        
        console.log('Prices updated successfully');
    } catch (error) {
        console.error('Price update failed:', error);
    }
}

// 每5分钟更新一次
setInterval(updatePrices, 5 * 60 * 1000);
```

## 监控和维护

### 1. Oracle健康状态检查

```javascript
const healthStatus = await paymaster.getOracleHealthStatus();
console.log('Oracle Healthy:', healthStatus.isHealthy);
console.log('Last Update:', new Date(healthStatus.lastUpdate * 1000));
```

### 2. 价格偏差监控

```javascript
const ethPrice = await oracle.getETHPriceUSD();
const cpopPrice = await oracle.getCPOPPriceUSD();

// 检查价格是否在合理范围内
if (ethPrice.price < minETHPrice || ethPrice.price > maxETHPrice) {
    console.warn('ETH price out of range:', ethPrice.price);
}
```

### 3. 回退机制使用统计

```javascript
// 监听回退事件
paymaster.on('FallbackRateUsed', (user, fallbackRate, reason) => {
    console.log(`Fallback used for ${user}: ${reason}`);
    // 发送监控警报
});
```

## 费用计算示例

### 链上计算 (Paymaster)
```solidity
// 获取当前汇率
uint256 exchangeRate = _getExchangeRateWithFallback();

// 计算CPOP成本
uint256 cpopCost = gasAmount * exchangeRate;

// 检查用户余额
require(cpopToken.balanceOf(user) >= cpopCost, "Insufficient CPOP");
```

### 链下计算 (Gas Calculator)
```javascript
// 获取当前Gas价格
const gasPrice = await calculator.getCurrentGasPrice();

// 估算Gas限制
const gasLimit = await calculator.estimateGasLimit(transaction);

// 计算CPOP成本
const cpopCost = await calculator.calculateGasCostInCPOP(gasLimit, gasPrice);
```

## 扩展性考虑

### 1. 支持多种代币
Oracle设计支持添加新的价格源：
```javascript
await oracle.addPriceFeed("USDC/USD", initialPrice, fallbackPrice);
```

### 2. 动态汇率调整
可以根据网络拥堵情况动态调整汇率：
```javascript
function getDynamicExchangeRate() {
    const baseRate = oracle.getCPOPForETH(1 ether);
    const congestionMultiplier = getNetworkCongestion();
    return baseRate * congestionMultiplier / 100;
}
```

### 3. 多Oracle聚合
未来可以集成多个Oracle源：
```javascript
function getAggregatedPrice(feedType) {
    const prices = [
        oracle1.getPrice(feedType),
        oracle2.getPrice(feedType),
        oracle3.getPrice(feedType)
    ];
    return calculateMedian(prices);
}
```

## 安全考虑

### 1. 价格操纵保护
- 价格偏差阈值限制
- 多源价格验证
- 时间加权平均价格

### 2. Oracle故障处理
- 回退价格机制
- 自动故障切换
- 人工干预能力

### 3. 访问控制
- Oracle更新权限管理
- Paymaster配置权限
- 紧急暂停功能

## 总结

本Oracle集成系统提供了：

1. **实时价格获取**: 通过Oracle获取ETH和CPOP的实时价格
2. **链下计算优化**: 减少链上计算成本，提供详细的成本分析
3. **回退机制**: 确保系统在Oracle故障时的可用性
4. **可扩展架构**: 支持未来添加更多代币和价格源
5. **监控能力**: 提供全面的健康状态和使用统计

该系统为CPOP代币的Gas费用支付提供了稳定、可靠且可扩展的解决方案。