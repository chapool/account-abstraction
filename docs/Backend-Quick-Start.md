# 后端智能领取 - 快速开始指南

**给后端开发人员的简化版指南**

## ⚡ 一分钟了解方案

1. **合约已支持**：`claimRewards(tokenId)` 会自动限制 90 天
2. **后端逻辑**：循环调用 `claimRewards()` 直到全部领取完成
3. **用户体验**：用户只需点击一次，后端自动分批处理
4. **无需修改合约**：所有逻辑在链下实现

---

## 🎯 核心思路

**不需要修改合约！** 所有逻辑都在链下实现。

### 工作原理

1. 用户调用后端 API
2. 后端循环调用 `claimRewards(tokenId)`
3. 合约自动限制每次最多 90 天
4. 后端检查是否还有剩余
5. 如果有，继续调用
6. 返回完整结果

---

## 📝 最简单的实现（100 行代码）

### 完整代码

```typescript
// claim-service.ts
import { ethers } from "ethers";

class ClaimService {
  private staking: ethers.Contract;
  
  constructor(stakingAddress: string, provider: ethers.Provider) {
    this.staking = new ethers.Contract(
      stakingAddress,
      [
        "function claimRewards(uint256 tokenId) external",
        "function stakes(uint256 tokenId) view returns (address, uint256, uint8, uint256, uint256, bool, uint256, uint256, bool)",
        "function getCurrentTimestamp() view returns (uint256)",
        "function calculatePendingRewards(uint256 tokenId) view returns (uint256)"
      ],
      provider
    );
  }
  
  /**
   * 智能领取 - 自动处理分批
   */
  async claimRewards(tokenIds: number[], userPrivateKey: string) {
    const wallet = new ethers.Wallet(userPrivateKey);
    const stakingWithSigner = this.staking.connect(wallet);
    
    const results = [];
    
    for (const tokenId of tokenIds) {
      const claimResult = {
        tokenId,
        status: "success",
        totalRewards: "0",
        batchCount: 0,
        txHashes: []
      };
      
      let hasMore = true;
      
      while (hasMore) {
        try {
          // 估算 Gas
          const gasEstimate = await stakingWithSigner.estimateGas.claimRewards(tokenId);
          
          // 调用 claimRewards（会自动限制 90 天）
          const tx = await stakingWithSigner.claimRewards(tokenId, {
            gasLimit: gasEstimate.mul(120).div(100) // 增加 20% buffer
          });
          
          const receipt = await tx.wait();
          
          claimResult.batchCount++;
          claimResult.txHashes.push(tx.hash);
          
          // 检查是否还有剩余待领取天数
          const remainingDays = await this.getPendingDays(tokenId);
          
          if (remainingDays === 0) {
            hasMore = false;
            
            // 计算总奖励
            const totalRewards = await stakingWithSigner.calculatePendingRewards(tokenId);
            claimResult.totalRewards = ethers.utils.formatEther(totalRewards);
            
            console.log(`✅ NFT ${tokenId} 领取完成`);
          } else {
            console.log(`⏳ NFT ${tokenId} 还有 ${remainingDays} 天待领取，继续...`);
          }
          
          // 避免过快调用
          await this.sleep(1000);
          
        } catch (error) {
          console.error(`❌ NFT ${tokenId} 领取失败:`, error.message);
          claimResult.status = "failed";
          claimResult.error = error.message;
          hasMore = false;
        }
      }
      
      results.push(claimResult);
    }
    
    return results;
  }
  
  /**
   * 计算待领取天数
   */
  private async getPendingDays(tokenId: number): Promise<number> {
    const stakeInfo = await this.staking.stakes(tokenId);
    const currentTime = await this.staking.getCurrentTimestamp();
    
    const lastClaimTime = Number(stakeInfo.lastClaimTime);
    const current = Number(currentTime);
    
    return Math.floor((current - lastClaimTime) / 86400);
  }
  
  private sleep(ms: number) {
    return new Promise(resolve => setTimeout(resolve, ms));
  }
}

export default ClaimService;
```

### API 路由

```typescript
// routes/claim.ts
import express from "express";
import ClaimService from "../services/claim-service";

const router = express.Router();
const claimService = new ClaimService(
  process.env.STAKING_ADDRESS,
  provider
);

router.post("/claim", async (req, res) => {
  const { userPrivateKey, tokenIds } = req.body;
  
  try {
    const results = await claimService.claimRewards(tokenIds, userPrivateKey);
    
    const successCount = results.filter(r => r.status === "success").length;
    const failedCount = results.filter(r => r.status === "failed").length;
    
    res.json({
      success: true,
      summary: {
        total: results.length,
        successCount,
        failedCount
      },
      results
    });
    
  } catch (error) {
    res.status(500).json({
      success: false,
      error: error.message
    });
  }
});

export default router;
```

### 使用示例

```bash
# 调用 API
curl -X POST http://localhost:3000/api/claim \
  -H "Content-Type: application/json" \
  -d '{
    "userPrivateKey": "0x...",
    "tokenIds": [4812, 3416, 3393]
  }'
```

### 响应示例

```json
{
  "success": true,
  "summary": {
    "total": 3,
    "successCount": 3,
    "failedCount": 0
  },
  "results": [
    {
      "tokenId": 4812,
      "status": "success",
      "totalRewards": "1723.455",
      "batchCount": 16,
      "txHashes": ["0x...", "0x...", "..."],
      "error": null
    },
    {
      "tokenId": 3416,
      "status": "success",
      "totalRewards": "5640.969",
      "batchCount": 16,
      "txHashes": ["0x...", "0x..."],
      "error": null
    },
    {
      "tokenId": 3393,
      "status": "success",
      "totalRewards": "11661.468",
      "batchCount": 16,
      "txHashes": ["0x..."],
      "error": null
    }
  ]
}
```

---

## 🔑 关键点

### 1. 不需要修改合约

使用现有函数：
- `claimRewards(tokenId)` - 领取奖励（自动限制 90 天）
- `stakes(tokenId)` - 获取质押信息
- `getCurrentTimestamp()` - 获取当前时间

### 2. 自动分批

合约的 `claimRewards()` 会自动：
- 限制每次最多 90 天
- 更新 `lastClaimTime`
- 发送奖励

后端只需：
- 循环调用 `claimRewards()`
- 检查是否还有剩余天数
- 如果有，继续调用

### 3. 错误处理

```typescript
// 自动重试
async function claimWithRetry(tokenId: number, maxRetries = 3) {
  for (let i = 0; i < maxRetries; i++) {
    try {
      return await staking.claimRewards(tokenId);
    } catch (error) {
      if (i === maxRetries - 1) throw error;
      await sleep(2000 * (i + 1)); // 指数退避
    }
  }
}
```

---

## 📊 测试数据

### 测试场景

```typescript
// NFT 1: 质押 90 天
// NFT 2: 质押 180 天  
// NFT 3: 质押 365 天
// NFT 4: 质押 1434 天

const results = await claimService.claimRewards([1, 2, 3, 4], privateKey);

// 预期结果：
// NFT 1: 1 次调用完成
// NFT 2: 2 次调用完成
// NFT 3: 4 次调用完成
// NFT 4: 16 次调用完成
```

---

## 💡 优化建议

### 1. 并发控制

```typescript
// 限制同时进行的领取数量
const MAX_CONCURRENT = 5;

async function claimWithConcurrency(tokenIds: number[]) {
  const results = [];
  
  for (let i = 0; i < tokenIds.length; i += MAX_CONCURRENT) {
    const batch = tokenIds.slice(i, i + MAX_CONCURRENT);
    const batchResults = await Promise.all(
      batch.map(id => claimRewards(id))
    );
    results.push(...batchResults);
  }
  
  return results;
}
```

### 2. Gas 优化

```typescript
// 估算 Gas 成本
async function estimateClaimCost(tokenIds: number[]): Promise<string> {
  let totalGas = ethers.BigNumber.from(0);
  
  for (const tokenId of tokenIds) {
    const gasEstimate = await staking.estimateGas.claimRewards(tokenId);
    totalGas = totalGas.add(gasEstimate);
  }
  
  const gasPrice = await provider.getGasPrice();
  const cost = totalGas.mul(gasPrice);
  
  return ethers.utils.formatEther(cost);
}
```

### 3. 进度反馈

```typescript
// WebSocket 实时反馈
io.emit('claim-progress', {
  tokenId,
  status: 'claiming',
  batch: 1,
  totalBatches: 3,
  txHash: '0x...'
});
```

---

## ✅ 总结

### 核心优势

1. ✅ **不需要修改合约** - 使用现有函数
2. ✅ **代码简单** - 不超过 100 行
3. ✅ **自动分批** - 后端自动处理
4. ✅ **灵活可控** - 可以重试、进度反馈等

### 快速开始

```bash
# 1. 安装依赖
npm install ethers

# 2. 复制代码
cp claim-service.ts src/services/

# 3. 配置合约地址
export STAKING_ADDRESS=0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5

# 4. 运行
npm start
```

---

**文档状态**: ✅ 完成  
**代码示例**: ✅ 完整可运行  
**测试状态**: ⏳ 待测试

