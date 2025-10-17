# BNB Chain Paymaster 集成指南

## 📋 目录
1. [概述](#概述)
2. [技术架构](#技术架构)
3. [集成步骤](#集成步骤)
4. [代码实现](#代码实现)
5. [测试验证](#测试验证)
6. [最佳实践](#最佳实践)

---

## 概述

本指南详细说明如何将 BNB Chain 的 EOA Paymaster 集成到 Staking 项目中，实现零 Gas 费用户体验。

### 核心优势
- ✅ 用户无需持有 BNB 即可质押
- ✅ 零 Gas 费操作（stake、unstake、claim）
- ✅ 提升用户转化率和活跃度
- ✅ 最小化合约改动

---

## 技术架构

### 整体架构图

```
┌─────────────┐
│   用户钱包   │
│ (EOA Wallet) │
└──────┬──────┘
       │ 1. 发起交易 (Gas Price = 0)
       ↓
┌─────────────────────┐
│   前端 dApp         │
│ - 捕获交易          │
│ - 添加赞助标识      │
└──────┬──────────────┘
       │ 2. 提交到 Paymaster
       ↓
┌──────────────────────┐
│  Paymaster Service   │
│  (NodeReal MegaFuel) │
│ - 验证赞助策略       │
│ - 创建赞助交易       │
│ - 打包 Bundle        │
└──────┬───────────────┘
       │ 3. 提交 Bundle
       ↓
┌──────────────────────┐
│   MEV Builders       │
│ - 构建区块           │
└──────┬───────────────┘
       │ 4. 提交区块
       ↓
┌──────────────────────┐
│   BSC Network        │
│ - Proposer 选择区块  │
│ - 执行交易           │
└──────────────────────┘
```

### 关键组件

1. **智能合约层**（Staking.sol）
   - ⚠️ **无需修改**（重要！）
   - 保持现有逻辑不变

2. **前端 dApp 层**
   - 集成 Paymaster SDK
   - 修改交易发送逻辑
   - 添加赞助策略处理

3. **Paymaster 服务层**
   - NodeReal MegaFuel
   - 或自建 Paymaster 服务

---

## 集成步骤

### 步骤 1: 选择 Paymaster 服务商

#### 选项 A: NodeReal MegaFuel（推荐）

**优点**：
- 官方支持
- 开箱即用
- 有管理后台

**注册步骤**：
```bash
1. 访问 NodeReal 官网
2. 注册账号
3. 创建 Paymaster 项目
4. 获取 API Key
5. 配置赞助策略
```

#### 选项 B: 自建 Paymaster

**优点**：
- 完全控制
- 自定义策略

**缺点**：
- 开发成本高
- 需要维护基础设施

---

### 步骤 2: 安装依赖

```bash
# 安装 ethers.js（如果还没有）
npm install ethers

# 安装 Paymaster SDK（NodeReal 提供）
npm install @nodereal/megafuel-sdk

# 或者使用通用的 Account Abstraction SDK
npm install @account-abstraction/sdk
```

---

### 步骤 3: 配置环境变量

创建 `.env.paymaster` 文件：

```bash
# Paymaster 配置
PAYMASTER_URL=https://open-platform.nodereal.io/paymaster
PAYMASTER_API_KEY=your_api_key_here
PAYMASTER_SPONSOR_ADDRESS=0x...your_sponsor_address

# BSC 配置
BSC_RPC_URL=https://bsc-dataseed.binance.org/
BSC_CHAIN_ID=56

# 赞助策略
MAX_GAS_SPONSORED_PER_TX=500000
DAILY_GAS_LIMIT=10000000
WHITELIST_CONTRACTS=0x...staking_contract_address
```

---

## 代码实现

### 1. 前端集成 - React 示例

#### 创建 Paymaster 配置文件

```typescript
// src/config/paymaster.config.ts

export const PAYMASTER_CONFIG = {
  url: process.env.REACT_APP_PAYMASTER_URL || 'https://open-platform.nodereal.io/paymaster',
  apiKey: process.env.REACT_APP_PAYMASTER_API_KEY || '',
  sponsorAddress: process.env.REACT_APP_PAYMASTER_SPONSOR_ADDRESS || '',
  
  // 赞助策略
  policy: {
    maxGasPerTx: 500000,
    dailyGasLimit: 10000000,
    whitelistedContracts: [
      process.env.REACT_APP_STAKING_CONTRACT_ADDRESS || '',
    ],
    whitelistedMethods: [
      'stake',
      'unstake',
      'claimRewards',
      'batchStake',
      'batchUnstake',
      'batchClaimRewards',
    ],
  },
};
```

#### 创建 Paymaster Hook

```typescript
// src/hooks/usePaymaster.ts

import { ethers } from 'ethers';
import { PAYMASTER_CONFIG } from '../config/paymaster.config';

interface PaymasterTxParams {
  to: string;
  data: string;
  value?: string;
  gasLimit?: number;
}

export function usePaymaster() {
  /**
   * 发送被赞助的交易
   */
  const sendSponsoredTransaction = async (
    params: PaymasterTxParams
  ): Promise<ethers.providers.TransactionResponse> => {
    try {
      const provider = new ethers.providers.Web3Provider(window.ethereum);
      const signer = provider.getSigner();
      const userAddress = await signer.getAddress();

      // 1. 构建交易对象（Gas Price = 0）
      const tx: ethers.providers.TransactionRequest = {
        to: params.to,
        data: params.data,
        value: params.value || '0',
        gasPrice: 0, // 关键：设置为 0
        gasLimit: params.gasLimit || 500000,
        from: userAddress,
        chainId: 56, // BSC mainnet
      };

      // 2. 获取 nonce
      const nonce = await provider.getTransactionCount(userAddress);
      tx.nonce = nonce;

      // 3. 签名交易
      const signedTx = await signer.signTransaction(tx);

      // 4. 提交到 Paymaster 服务
      const paymasterResponse = await fetch(PAYMASTER_CONFIG.url, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${PAYMASTER_CONFIG.apiKey}`,
        },
        body: JSON.stringify({
          signedTransaction: signedTx,
          sponsorPolicy: PAYMASTER_CONFIG.policy,
        }),
      });

      if (!paymasterResponse.ok) {
        throw new Error('Paymaster submission failed');
      }

      const result = await paymasterResponse.json();
      
      // 5. 返回交易哈希
      return {
        hash: result.transactionHash,
        wait: async (confirmations?: number) => {
          return provider.waitForTransaction(result.transactionHash, confirmations);
        },
      } as ethers.providers.TransactionResponse;

    } catch (error) {
      console.error('Sponsored transaction failed:', error);
      throw error;
    }
  };

  /**
   * 检查交易是否符合赞助条件
   */
  const isEligibleForSponsorship = (
    contractAddress: string,
    methodName: string
  ): boolean => {
    const { whitelistedContracts, whitelistedMethods } = PAYMASTER_CONFIG.policy;
    
    return (
      whitelistedContracts.includes(contractAddress.toLowerCase()) &&
      whitelistedMethods.includes(methodName)
    );
  };

  /**
   * 获取用户的赞助配额状态
   */
  const getSponsorshipQuota = async (userAddress: string) => {
    const response = await fetch(`${PAYMASTER_CONFIG.url}/quota/${userAddress}`, {
      headers: {
        'Authorization': `Bearer ${PAYMASTER_CONFIG.apiKey}`,
      },
    });

    return response.json();
  };

  return {
    sendSponsoredTransaction,
    isEligibleForSponsorship,
    getSponsorshipQuota,
  };
}
```

#### 集成到 Staking 操作

```typescript
// src/hooks/useStaking.ts

import { ethers } from 'ethers';
import { usePaymaster } from './usePaymaster';
import StakingABI from '../abis/Staking.json';

const STAKING_CONTRACT_ADDRESS = process.env.REACT_APP_STAKING_CONTRACT_ADDRESS;

export function useStaking() {
  const { sendSponsoredTransaction, isEligibleForSponsorship } = usePaymaster();

  /**
   * 质押 NFT（使用 Paymaster）
   */
  const stakeWithPaymaster = async (tokenId: number) => {
    try {
      // 1. 编码函数调用
      const stakingInterface = new ethers.utils.Interface(StakingABI);
      const data = stakingInterface.encodeFunctionData('stake', [tokenId]);

      // 2. 检查是否符合赞助条件
      if (!isEligibleForSponsorship(STAKING_CONTRACT_ADDRESS, 'stake')) {
        throw new Error('Transaction not eligible for sponsorship');
      }

      // 3. 发送被赞助的交易
      const tx = await sendSponsoredTransaction({
        to: STAKING_CONTRACT_ADDRESS,
        data: data,
        gasLimit: 150000, // 根据 Gas 估算文档
      });

      console.log('Sponsored transaction sent:', tx.hash);

      // 4. 等待确认
      const receipt = await tx.wait();
      console.log('Transaction confirmed:', receipt);

      return receipt;
    } catch (error) {
      console.error('Stake with paymaster failed:', error);
      throw error;
    }
  };

  /**
   * 批量质押（使用 Paymaster）
   */
  const batchStakeWithPaymaster = async (tokenIds: number[]) => {
    const stakingInterface = new ethers.utils.Interface(StakingABI);
    const data = stakingInterface.encodeFunctionData('batchStake', [tokenIds]);

    const tx = await sendSponsoredTransaction({
      to: STAKING_CONTRACT_ADDRESS,
      data: data,
      gasLimit: 100000 * tokenIds.length, // 动态计算
    });

    return tx.wait();
  };

  /**
   * 领取奖励（使用 Paymaster）
   */
  const claimRewardsWithPaymaster = async (tokenId: number) => {
    const stakingInterface = new ethers.utils.Interface(StakingABI);
    const data = stakingInterface.encodeFunctionData('claimRewards', [tokenId]);

    const tx = await sendSponsoredTransaction({
      to: STAKING_CONTRACT_ADDRESS,
      data: data,
      gasLimit: 150000,
    });

    return tx.wait();
  };

  /**
   * 取消质押（使用 Paymaster）
   */
  const unstakeWithPaymaster = async (tokenId: number) => {
    const stakingInterface = new ethers.utils.Interface(StakingABI);
    const data = stakingInterface.encodeFunctionData('unstake', [tokenId]);

    const tx = await sendSponsoredTransaction({
      to: STAKING_CONTRACT_ADDRESS,
      data: data,
      gasLimit: 200000,
    });

    return tx.wait();
  };

  return {
    stakeWithPaymaster,
    batchStakeWithPaymaster,
    claimRewardsWithPaymaster,
    unstakeWithPaymaster,
  };
}
```

#### UI 组件示例

```tsx
// src/components/StakeNFT.tsx

import React, { useState } from 'react';
import { useStaking } from '../hooks/useStaking';

export function StakeNFT({ tokenId }: { tokenId: number }) {
  const [loading, setLoading] = useState(false);
  const { stakeWithPaymaster } = useStaking();

  const handleStake = async () => {
    try {
      setLoading(true);
      
      // 使用 Paymaster 质押（零 Gas 费）
      const receipt = await stakeWithPaymaster(tokenId);
      
      alert(`Successfully staked NFT #${tokenId}! No gas fee paid!`);
    } catch (error) {
      console.error('Stake failed:', error);
      alert('Stake failed: ' + error.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="stake-card">
      <h3>NFT #{tokenId}</h3>
      <button 
        onClick={handleStake} 
        disabled={loading}
        className="stake-button"
      >
        {loading ? 'Staking...' : '🎉 Stake (Free Gas!)'}
      </button>
      <p className="gas-info">
        ⚡ Gas fee sponsored by platform
      </p>
    </div>
  );
}
```

---

### 2. NodeReal MegaFuel SDK 集成（推荐方式）

如果使用 NodeReal 的 MegaFuel 服务，可以使用他们的官方 SDK：

```typescript
// src/services/paymasterService.ts

import { MegaFuelSDK } from '@nodereal/megafuel-sdk';

// 初始化 SDK
const megafuel = new MegaFuelSDK({
  apiKey: process.env.REACT_APP_PAYMASTER_API_KEY,
  network: 'bsc-mainnet',
});

export async function sendGaslessTransaction(
  contractAddress: string,
  functionData: string,
  userAddress: string
) {
  // 1. 构建交易
  const tx = {
    to: contractAddress,
    data: functionData,
    from: userAddress,
    gasPrice: 0, // 零 Gas
  };

  // 2. 请求赞助
  const sponsoredTx = await megafuel.sponsorTransaction(tx);

  // 3. 发送交易
  const receipt = await megafuel.sendTransaction(sponsoredTx);

  return receipt;
}

// 使用示例
export async function stakeWithMegaFuel(tokenId: number, userAddress: string) {
  const stakingInterface = new ethers.utils.Interface(StakingABI);
  const data = stakingInterface.encodeFunctionData('stake', [tokenId]);

  return sendGaslessTransaction(
    STAKING_CONTRACT_ADDRESS,
    data,
    userAddress
  );
}
```

---

### 3. 后端服务（可选）

如果需要更精细的控制，可以搭建后端服务：

```typescript
// backend/src/services/paymasterService.ts

import { ethers } from 'ethers';

export class PaymasterService {
  private provider: ethers.providers.JsonRpcProvider;
  private sponsorWallet: ethers.Wallet;

  constructor() {
    this.provider = new ethers.providers.JsonRpcProvider(
      process.env.BSC_RPC_URL
    );
    this.sponsorWallet = new ethers.Wallet(
      process.env.SPONSOR_PRIVATE_KEY,
      this.provider
    );
  }

  /**
   * 验证赞助策略
   */
  async validateSponsorPolicy(
    userAddress: string,
    contractAddress: string,
    functionSelector: string
  ): Promise<boolean> {
    // 1. 检查合约白名单
    const whitelistedContracts = process.env.WHITELISTED_CONTRACTS?.split(',') || [];
    if (!whitelistedContracts.includes(contractAddress.toLowerCase())) {
      return false;
    }

    // 2. 检查用户每日配额
    const userQuota = await this.getUserDailyQuota(userAddress);
    if (userQuota.used >= userQuota.limit) {
      return false;
    }

    // 3. 检查函数白名单
    const allowedFunctions = ['0x...stake', '0x...unstake', '0x...claimRewards'];
    if (!allowedFunctions.includes(functionSelector)) {
      return false;
    }

    return true;
  }

  /**
   * 创建赞助交易
   */
  async createSponsorTransaction(
    userTx: ethers.providers.TransactionRequest,
    estimatedGas: number
  ) {
    // 计算需要的 Gas 费
    const gasPrice = await this.provider.getGasPrice();
    const gasCost = gasPrice.mul(estimatedGas);

    // 创建赞助交易（从 sponsor wallet 到自己，携带足够的 Gas）
    const sponsorTx = {
      to: this.sponsorWallet.address,
      value: 0,
      gasPrice: gasPrice.mul(2), // 更高的 Gas Price 确保被打包
      gasLimit: 21000,
      data: '0x', // 可以添加元数据
    };

    return sponsorTx;
  }

  /**
   * 提交 Bundle
   */
  async submitBundle(userTx: string, sponsorTx: string) {
    // 使用 Flashbots 或 BSC 的 MEV 网络提交 Bundle
    const bundle = [userTx, sponsorTx];
    
    // 提交到多个 builders
    const builders = [
      'https://bsc-builder-1.example.com',
      'https://bsc-builder-2.example.com',
    ];

    const promises = builders.map(builderUrl =>
      fetch(builderUrl, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ bundle }),
      })
    );

    return Promise.all(promises);
  }

  private async getUserDailyQuota(userAddress: string) {
    // 从数据库查询用户配额
    return {
      used: 5000000, // 已使用的 Gas
      limit: 10000000, // 每日限额
    };
  }
}
```

#### API 端点

```typescript
// backend/src/routes/paymaster.routes.ts

import express from 'express';
import { PaymasterService } from '../services/paymasterService';

const router = express.Router();
const paymasterService = new PaymasterService();

/**
 * POST /api/paymaster/sponsor
 * 请求赞助交易
 */
router.post('/sponsor', async (req, res) => {
  try {
    const { signedTransaction, userAddress, contractAddress } = req.body;

    // 1. 解析交易
    const tx = ethers.utils.parseTransaction(signedTransaction);

    // 2. 验证赞助策略
    const isValid = await paymasterService.validateSponsorPolicy(
      userAddress,
      contractAddress,
      tx.data.slice(0, 10) // function selector
    );

    if (!isValid) {
      return res.status(403).json({ error: 'Not eligible for sponsorship' });
    }

    // 3. 创建并提交 Bundle
    const sponsorTx = await paymasterService.createSponsorTransaction(
      tx,
      200000 // estimated gas
    );

    const signedSponsorTx = await paymasterService.signTransaction(sponsorTx);

    await paymasterService.submitBundle(signedTransaction, signedSponsorTx);

    res.json({
      success: true,
      transactionHash: tx.hash,
      message: 'Transaction sponsored successfully',
    });

  } catch (error) {
    console.error('Sponsor failed:', error);
    res.status(500).json({ error: error.message });
  }
});

/**
 * GET /api/paymaster/quota/:address
 * 查询用户配额
 */
router.get('/quota/:address', async (req, res) => {
  const quota = await paymasterService.getUserQuota(req.params.address);
  res.json(quota);
});

export default router;
```

---

## 测试验证

### 1. 本地测试

```typescript
// test/paymaster.test.ts

import { expect } from 'chai';
import { ethers } from 'hardhat';
import { stakeWithMegaFuel } from '../src/services/paymasterService';

describe('Paymaster Integration', () => {
  it('should stake NFT with zero gas fee', async () => {
    const [user] = await ethers.getSigners();
    
    // 获取初始 BNB 余额
    const initialBalance = await user.getBalance();
    
    // 使用 Paymaster 质押
    const receipt = await stakeWithMegaFuel(1, user.address);
    
    // 验证交易成功
    expect(receipt.status).to.equal(1);
    
    // 验证用户 BNB 余额未减少（零 Gas 费）
    const finalBalance = await user.getBalance();
    expect(finalBalance).to.equal(initialBalance);
    
    console.log('✅ Gasless transaction successful!');
  });
});
```

### 2. 测试网测试

```bash
# 1. 部署到 BSC 测试网
npx hardhat deploy --network bscTestnet

# 2. 配置 Paymaster（使用测试 API Key）
export PAYMASTER_API_KEY=test_key_...

# 3. 运行测试脚本
npx hardhat run scripts/test-paymaster.ts --network bscTestnet
```

---

## 最佳实践

### 1. 赞助策略设计

```typescript
// 推荐的赞助策略配置
const SPONSOR_POLICY = {
  // 方式 1: 白名单用户（新用户福利）
  whitelistedUsers: [
    '0x...', // VIP 用户
  ],
  
  // 方式 2: 限额策略
  quota: {
    perUser: {
      daily: 10, // 每天 10 笔免费交易
      weekly: 50,
    },
    perContract: {
      daily: 10000000, // 每天赞助的总 Gas
    },
  },
  
  // 方式 3: 操作类型限制
  allowedOperations: {
    'stake': true,          // 质押免费
    'claimRewards': true,   // 领取免费
    'unstake': false,       // 取消质押需要自付
  },
  
  // 方式 4: 时间限制（活动期）
  timeRestriction: {
    startTime: '2025-01-01',
    endTime: '2025-12-31',
  },
  
  // 方式 5: 条件赞助（持有特定 NFT）
  conditions: {
    requireNFTOwnership: true,
    nftContract: '0x...CPNFT_address',
    minNFTCount: 1,
  },
};
```

### 2. 错误处理

```typescript
async function handlePaymasterTransaction() {
  try {
    const tx = await stakeWithPaymaster(tokenId);
    return tx;
  } catch (error) {
    // 如果 Paymaster 失败，降级到普通交易
    if (error.code === 'PAYMASTER_REJECTED') {
      console.log('Paymaster rejected, falling back to normal tx');
      return stakeNormal(tokenId); // 用户自付 Gas
    }
    
    // 其他错误
    throw error;
  }
}
```

### 3. 用户体验优化

```tsx
// 显示 Gas 费节省信息
function GasSavingsBanner() {
  return (
    <div className="banner">
      <h3>🎉 Free Gas Promotion!</h3>
      <p>Stake your NFTs with ZERO gas fees</p>
      <p>Sponsored by the platform</p>
      <small>Save ~$0.12 per transaction</small>
    </div>
  );
}

// 配额提示
function QuotaIndicator({ used, limit }: { used: number; limit: number }) {
  const percentage = (used / limit) * 100;
  
  return (
    <div className="quota">
      <p>Free transactions: {limit - used} / {limit}</p>
      <div className="progress-bar">
        <div style={{ width: `${percentage}%` }} />
      </div>
    </div>
  );
}
```

### 4. 监控和分析

```typescript
// 记录 Paymaster 使用情况
async function trackPaymasterUsage(userAddress: string, operation: string, gasSaved: number) {
  await analytics.track({
    event: 'paymaster_used',
    properties: {
      user: userAddress,
      operation: operation,
      gasSaved: gasSaved,
      timestamp: Date.now(),
    },
  });
}

// 每日统计
async function getDailyPaymasterStats() {
  return {
    totalTransactions: 1234,
    totalGasSponsored: 456000000, // wei
    totalCostUSD: 547.20,
    uniqueUsers: 89,
    topOperations: [
      { operation: 'stake', count: 567 },
      { operation: 'claimRewards', count: 456 },
      { operation: 'batchStake', count: 211 },
    ],
  };
}
```

---

## 成本估算

### Paymaster 运营成本

基于你的 Gas 估算文档（1 Gwei）：

| 操作 | Gas | 成本/笔 | 1000用户/天 | 月成本 |
|------|-----|---------|------------|--------|
| Stake | 98,200 | $0.118 | $118 | $3,540 |
| Unstake | 181,200 | $0.218 | $218 | $6,540 |
| Claim | 118,600 | $0.142 | $142 | $4,260 |
| **总计** | - | - | $478/天 | **$14,340/月** |

**优化建议**：
- 设置每用户每日配额（如 5 笔）
- 限制赞助特定操作（如只赞助 stake 和 claim）
- 分阶段推出（测试期 → 限量 → 全量）

---

## 部署清单

- [ ] 注册 Paymaster 服务（NodeReal）
- [ ] 配置赞助策略
- [ ] 前端集成 SDK
- [ ] 测试网测试
- [ ] 设置监控和告警
- [ ] 准备运营资金
- [ ] 主网部署
- [ ] 用户教育和推广

---

## 参考资源

- [BNB Chain Paymaster 官方文档](https://docs.bnbchain.org/bnb-smart-chain/developers/paymaster/overview/)
- [NodeReal MegaFuel](https://nodereal.io/megafuel)
- [Ethers.js 文档](https://docs.ethers.org/)
- [Staking Gas 估算](./STAKING_GAS_ESTIMATION.md)

---

## 常见问题

### Q1: 合约需要修改吗？
**A**: 不需要！Paymaster 是链下解决方案，合约保持不变。

### Q2: 所有用户都能享受零 Gas 费吗？
**A**: 取决于你的赞助策略。可以设置白名单、配额等限制。

### Q3: Paymaster 失败怎么办？
**A**: 应该有降级方案，让用户可以选择自付 Gas 费完成交易。

### Q4: 成本可控吗？
**A**: 可以！通过配额、白名单、操作类型限制等方式控制成本。

### Q5: 多久能集成完成？
**A**: 
- 使用 NodeReal：1-2 周
- 自建服务：4-6 周

---

**准备好开始了吗？从 NodeReal MegaFuel 开始是最快的方式！** 🚀

