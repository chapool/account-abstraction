# 后端 API 集成指南 - 批量领取/解质押

**版本**: v1.0  
**日期**: 2025-01-27  
**目标**: 为后端开发人员提供智能批量操作合约和接口设计

---

## 📋 目录

1. [背景说明](#背景说明)
2. [合约接口](#合约接口)
3. [API 设计](#api设计)
4. [实现流程](#实现流程)
5. [错误处理](#错误处理)
6. [测试用例](#测试用例)

---

## 背景说明

### 核心需求

用户质押 NFT 时间较长（> 90 天）后，领取奖励会遇到 Gas 超限问题。

### 解决方案

**后端智能分批调用**：
1. 用户在前端点击"领取奖励"
2. 前端调用后端 API
3. 后端自动分批调用合约
4. 返回完整的领取结果

### 优势

- ✅ 用户只需点击一次
- ✅ 后端智能处理分批逻辑
- ✅ 自动重试失败的操作
- ✅ 实时反馈进度
- ✅ 支持批量 NFT 操作

---

## 合约接口

### ⚠️ 重要说明

**不需要在合约中添加新函数！** 所有逻辑都可在链下实现。

### 1. 已存在的合约函数（可直接使用）

#### 查询待领取奖励（无限制，用于显示）
```solidity
function calculatePendingRewards(uint256 tokenId) external view returns (uint256);
```

#### 领取单个 NFT 奖励（限制 90 天）
```solidity
function claimRewards(uint256 tokenId) external nonReentrant whenNotPaused;
```

#### 批量领取（后端管理员专用）

```solidity
/**
 * @dev 批量领取奖励（Backend/Admin only）
 * @param tokenIds 要领取的 NFT 列表
 * @param userAddress 用户地址（奖励将发送到该用户的 AA 账户）
 * @notice 只有合约 owner 可以调用
 */
function batchClaimRewards(
    address userAddress,
    uint256[] calldata tokenIds
) external nonReentrant whenNotPaused onlyOwner;
```

**重要**：
- ⚠️ **只有合约 owner 可以调用**
- 后端需要使用 **owner 私钥** 调用此函数
- 奖励会发送到 `userAddress` 的 AA 账户
- 自动限制单个 NFT 最多领取 90 天的奖励

### 3. 批量质押（后端管理员专用）

```solidity
/**
 * @dev 批量质押（Backend/Admin only）
 * @param userAddress 用户地址（NFT 将质押到该用户账户）
 * @param tokenIds 要质押的 NFT 列表
 * @notice 只有合约 owner 可以调用
 */
function batchStake(
    address userAddress,
    uint256[] calldata tokenIds
) external nonReentrant whenNotPaused onlyOwner;
```

**重要**：
- ⚠️ **只有合约 owner 可以调用**
- 后端需要使用 **owner 私钥** 调用此函数
- NFT 必须属于 `userAddress`
- 质押后 NFT 状态会更新为已质押

### 4. 批量解质押（后端管理员专用）

```solidity
/**
 * @dev 批量解质押（Backend/Admin only）
 * @param tokenIds 要解质押的 NFT 列表
 * @param userAddress 用户地址（奖励将发送到该用户的 AA 账户）
 * @notice 只有合约 owner 可以调用
 */
function batchUnstake(
    address userAddress,
    uint256[] calldata tokenIds
) external nonReentrant whenNotPaused onlyOwner;
```

**重要**：
- ⚠️ **只有合约 owner 可以调用**
- 后端需要使用 **owner 私钥** 调用此函数
- 奖励会发送到 `userAddress` 的 AA 账户
- NFT 会返还给用户

---

## API 设计

### 1. 用户领取奖励 API

#### 接口名称
`POST /api/staking/claim-rewards`

#### 请求参数

```json
{
  "userAddress": "0x...",  // 用户钱包地址
  "tokenIds": [4812, 3416, 3393],  // 要领取的 NFT 列表（可选，空则领取全部）
  "claimAll": true  // 是否领取全部（如果 tokenIds 为空）
}
```

#### 响应结果

```json
{
  "success": true,
  "data": {
    "totalClaimedNFTs": 3,
    "totalRewards": "1723.455",  // CPOP 总奖励
    "claimedBatches": [
      {
        "batchIndex": 1,
        "tokenIds": [4812, 3416, 3393],
        "rewards": "1545.123",
        "status": "success",
        "txHash": "0x..."
      }
    ],
    "pendingNFTs": [
      {
        "tokenId": 4812,
        "remainingDays": 945,  // 剩余待领取天数
        "remainingRewards": "567.890"
      }
    ],
    "estimatedGasCost": {
      "eth": "0.0015",  // 以太坊
      "usd": "3.45"
    }
  },
  "message": "成功领取 3 个 NFT 的奖励，其中 1 个 NFT 仍有待领取"
}
```

#### 业务流程

```
1. 后端接收请求
2. 获取用户的质押 NFT 列表（如果 claimAll=true）
3. 对每个 NFT 检查待领取天数
4. 如果 > 90 天，后端多次调用合约（每次最多 90 天）
5. 记录每次调用结果
6. 如果有剩余的，在 response 中提示用户
7. 返回完整结果
```

### 2. 用户质押 API

#### 接口名称
`POST /api/staking/stake`

#### 请求参数

```json
{
  "userAddress": "0x...",  // 用户钱包地址
  "tokenIds": [4812, 3416, 3393]  // 要质押的 NFT 列表
}
```

#### 响应结果

```json
{
  "success": true,
  "data": {
    "stakedCount": 3,
    "txHash": "0x...",  // 批量质押的交易哈希
    "stakedNFTs": [
      {
        "tokenId": 4812,
        "level": 3  // A级
      },
      {
        "tokenId": 3416,
        "level": 4  // S级
      },
      {
        "tokenId": 3393,
        "level": 2  // B级
      }
    ],
    "failedNFTs": [],  // 质押失败的 NFT（如果有）
    "estimatedGasCost": {
      "eth": "0.0012",
      "usd": "2.76"
    }
  },
  "message": "成功质押 3 个 NFT"
}
```

#### 业务流程

```
1. 后端接收请求
2. 验证用户地址和 NFT 所有权
3. 检查每个 NFT 是否已质押
4. 过滤掉不符合质押条件的 NFT（如 NORMAL 级别）
5. 一次性调用合约 batchStake（最多 50 个）
6. 返回完整结果
```

#### 错误处理

**NFT 已质押**：
```json
{
  "success": false,
  "error": {
    "code": "ALREADY_STAKED",
    "message": "NFT 已质押",
    "tokenId": 4812
  }
}
```

**NFT 不属于用户**：
```json
{
  "success": false,
  "error": {
    "code": "NOT_OWNER",
    "message": "不是该 NFT 的所有者",
    "tokenId": 3416
  }
}
```

**NORMAL 级别 NFT 不能质押**：
```json
{
  "success": false,
  "error": {
    "code": "INVALID_LEVEL",
    "message": "NORMAL 级别 NFT 不能质押",
    "tokenId": 3393
  }
}
```

### 3. 用户解质押 API

#### 接口名称
`POST /api/staking/unstake`

#### 请求参数

```json
{
  "userAddress": "0x...",
  "tokenIds": [4812, 3416],  // 要解质押的 NFT 列表
  "force": false  // 是否强制解质押（忽略提前取回惩罚）
}
```

#### 响应结果

```json
{
  "success": true,
  "data": {
    "unstakedCount": 2,
    "totalRewards": "2156.789",
    "penalty": "0",  // 提前取回惩罚
    "unstakedNFTs": [
      {
        "tokenId": 4812,
        "rewards": "1234.567",
        "penalty": "0",
        "txHash": "0x..."
      },
      {
        "tokenId": 3416,
        "rewards": "922.222",
        "penalty": "0",
        "txHash": "0x..."
      }
    ],
    "nftReturned": true  // NFT 已返还
  }
}
```

---

## 实现流程

### 后端领取奖励实现（伪代码）

```typescript
async function claimRewardsAPI(req, res) {
  const { userAddress, tokenIds, claimAll } = req.body;
  
  try {
    // 1. 获取用户 NFT 列表
    let nftList = [];
    if (claimAll) {
      nftList = await getStakedNFTs(userAddress);
    } else {
      nftList = tokenIds;
    }
    
    // 2. 检查每个 NFT 的待领取天数
    const claimPlan = [];
    for (const tokenId of nftList) {
      const pendingDays = await staking.getPendingDays(tokenId);
      
      if (pendingDays <= 90) {
        // 一次领取完成
        claimPlan.push({
          tokenId,
          needsMultipleClaims: false,
          claimsNeeded: 1
        });
      } else {
        // 需要分批
        const claimsNeeded = Math.ceil(pendingDays / 90);
        claimPlan.push({
          tokenId,
          needsMultipleClaims: true,
          claimsNeeded
        });
      }
    }
    
    // 3. 执行领取
    const results = [];
    for (const plan of claimPlan) {
      if (!plan.needsMultipleClaims) {
        // 正常领取
        const result = await batchClaimRewards([plan.tokenId]);
        results.push(result);
      } else {
        // 多次领取
        let remaining = plan.claimsNeeded;
        while (remaining > 0) {
          const result = await batchClaimRewards([plan.tokenId]);
          results.push(result);
          
          // 检查是否还有剩余
          const newPendingDays = await staking.getPendingDays(plan.tokenId);
          if (newPendingDays === 0) break; // 已全部领取
          
          remaining--;
        }
      }
    }
    
    // 4. 汇总结果
    const response = buildClaimResponse(results, claimPlan);
    res.json({ success: true, data: response });
    
  } catch (error) {
    res.status(500).json({ 
      success: false, 
      error: error.message 
    });
  }
}
```

### 关键逻辑

```typescript
// 检查是否需要分批的逻辑
async function checkIfNeedsBatching(tokenId: number): Promise<{
  needsBatching: boolean;
  batches: number;
  firstBatchRewards: string;
}> {
  const pendingDays = await staking.getPendingDays(tokenId);
  
  if (pendingDays <= 90) {
    return { needsBatching: false, batches: 1, firstBatchRewards: "0" };
  }
  
  // 计算需要分几批
  const batches = Math.ceil(pendingDays / 90);
  
  // 模拟第一批的奖励（90天）
  const firstBatchRewards = await estimateRewards(tokenId, 90);
  
  return {
    needsBatching: true,
    batches,
    firstBatchRewards
  };
}
```

---

## 错误处理

### 常见错误情况

#### 1. 部分 NFT 领取失败

```json
{
  "success": true,
  "partial": true,
  "data": {
    "successful": [
      { "tokenId": 4812, "rewards": "123.45" }
    ],
    "failed": [
      { 
        "tokenId": 3416, 
        "error": "Insufficient gas", 
        "retryable": true 
      }
    ]
  }
}
```

#### 2. Gas 不足

```json
{
  "success": false,
  "error": {
    "code": "INSUFFICIENT_GAS",
    "message": "Gas 不足",
    "required": "0.0015 ETH",
    "available": "0.0005 ETH",
    "suggestion": "请充值后再试"
  }
}
```

#### 3. NFT 未质押

```json
{
  "success": false,
  "error": {
    "code": "NOT_STAKED",
    "message": "NFT 未质押",
    "tokenId": 4812
  }
}
```

### 重试机制

```typescript
async function claimWithRetry(tokenId: number, maxRetries = 3) {
  for (let i = 0; i < maxRetries; i++) {
    try {
      return await batchClaimRewards([tokenId]);
    } catch (error) {
      if (i === maxRetries - 1) throw error;
      
      // 等待后重试
      await sleep(1000 * (i + 1));
    }
  }
}
```

---

## 测试用例

### 测试场景 1: 短期质押领取

**输入**：
```json
{
  "userAddress": "0x...",
  "tokenIds": [4812],
  "claimAll": false
}
```

**预期**：一次调用完成，返回完整奖励

### 测试场景 2: 长期质押领取（需要分批）

**输入**：
```json
{
  "userAddress": "0x...",
  "tokenIds": [3416],  // 已质押 180 天
  "claimAll": false
}
```

**预期**：后端自动调用 2 次，返回 180 天的奖励

### 测试场景 3: 批量领取

**输入**：
```json
{
  "userAddress": "0x...",
  "tokenIds": [4812, 3416, 3393],  // 3 个 NFT
  "claimAll": false
}
```

**预期**：对每个 NFT 分别处理，自动分批

### 测试场景 4: 领取全部

**输入**：
```json
{
  "userAddress": "0x...",
  "claimAll": true
}
```

**预期**：获取用户所有质押 NFT，逐个领取

---

## 前端集成示例

### React Hook 示例

```typescript
// hooks/useClaimRewards.ts
export function useClaimRewards() {
  const [loading, setLoading] = useState(false);
  const [progress, setProgress] = useState(0);
  
  const claimRewards = async (tokenIds?: number[]) => {
    setLoading(true);
    
    try {
      const response = await fetch('/api/staking/claim-rewards', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          userAddress: wallet.address,
          tokenIds,
          claimAll: !tokenIds
        })
      });
      
      const result = await response.json();
      
      if (result.success) {
        // 显示成功消息
        toast.success(`成功领取 ${result.data.totalClaimedNFTs} 个 NFT`);
        
        // 如果有剩余的，提示用户
        if (result.data.pendingNFTs.length > 0) {
          toast.info('部分 NFT 还有待领取奖励，可以稍后继续领取');
        }
      }
      
      return result;
    } finally {
      setLoading(false);
    }
  };
  
  return { claimRewards, loading, progress };
}
```

---

## 性能优化建议

### 1. 批量处理优化

```typescript
// 如果多个 NFT 都可以一次性领取，批量调用
const canBatchTogether = claimPlan.filter(p => !p.needsMultipleClaims);

if (canBatchTogether.length > 5) {
  // 分批调用合约，每批 20 个
  for (let i = 0; i < canBatchTogether.length; i += 20) {
    const batch = canBatchTogether.slice(i, i + 20);
    await batchClaimRewards(batch.map(p => p.tokenId));
  }
}
```

### 2. 并发控制

```typescript
// 限制并发数
const CONCURRENT_LIMIT = 3;

async function claimWithConcurrencyLimit(plans: ClaimPlan[]) {
  const semaphore = new Semaphore(CONCURRENT_LIMIT);
  
  const results = await Promise.all(
    plans.map(plan => 
      semaphore.acquire(() => executeClaim(plan))
    )
  );
  
  return results;
}
```

### 3. Gas 预估

```typescript
async function estimateGasCost(tokenIds: number[]): Promise<{
  gas: string;
  cost: string;
}> {
  const gasEstimate = await staking.estimateGas.batchClaimRewards(tokenIds);
  const gasPrice = await provider.getGasPrice();
  const cost = gasEstimate.mul(gasPrice);
  
  return {
    gas: gasEstimate.toString(),
    cost: ethers.utils.formatEther(cost)
  };
}
```

---

## 安全注意事项

1. **用户认证**：验证用户签名
2. **权限检查**：确保用户是 NFT 所有者
3. **防重放攻击**：使用 nonce
4. **Rate Limiting**：限制 API 调用频率
5. **Gas 限制**：设置合理的 Gas 上限

---

## 总结

### 关键接口

| 功能 | 合约函数 | API 端点 |
|-----|---------|---------|
| 质押 | `batchStake` | `POST /api/staking/stake` |
| 领取奖励 | `batchClaimRewards` | `POST /api/staking/claim-rewards` |
| 解质押 | `batchUnstake` | `POST /api/staking/unstake` |
| 查询待领取天数 | `getPendingDays` | `GET /api/staking/pending-days/:tokenId` |

### 后端职责

1. 接收用户请求
2. 智能分批调用合约
3. 重试失败的调用
4. 返回完整结果

### 优势

- ✅ 用户只需一次操作
- ✅ 后端处理复杂逻辑
- ✅ 自动重试和错误处理
- ✅ 实时反馈进度

---

## 链下实现指南

### ✅ 核心要点

**所有逻辑都可在链下实现，不需要修改合约！**

只需要使用现有的合约函数：
- `staking.stakes(tokenId)` - 获取质押信息
- `staking.getCurrentTimestamp()` - 获取当前时间
- `staking.calculatePendingRewards(tokenId)` - 查询待领取奖励
- `staking.claimRewards(tokenId)` - 领取奖励（自动限制 90 天）
- `staking.batchClaimRewards(tokenIds)` - 批量领取

### 为什么可以链下实现？

1. **所有数据都可以查询**：
   - 质押时间、最后领取时间 → 可以用 `stakes()` 查询
   - 当前时间 → 可以用 `getCurrentTimestamp()` 查询
   - 待领取天数 → 链下计算

2. **合约已有保护**：
   - `claimRewards()` 会自动限制 90 天
   - 后端只需循环调用即可

3. **灵活可控**：
   - 后端可以智能分批
   - 可以重试失败的调用
   - 可以给用户实时反馈

### 链下实现：获取待领取天数

```typescript
/**
 * 获取 NFT 的待领取天数（链下实现）
 */
async function getPendingDays(tokenId: number): Promise<number> {
  // 1. 获取质押信息
  const stakeInfo = await staking.stakes(tokenId);
  
  // 2. 获取当前时间
  const currentTime = await staking.getCurrentTimestamp(); // 或使用 block.timestamp
  
  // 3. 计算天数
  const lastClaimTime = Number(stakeInfo.lastClaimTime);
  const current = Number(currentTime);
  const pendingDays = Math.floor((current - lastClaimTime) / 86400);
  
  return pendingDays;
}
```

### 链下实现：判断是否需要分批

```typescript
/**
 * 判断是否需要分批领取（链下实现）
 */
async function needsBatching(tokenId: number): Promise<{
  needsBatching: boolean;
  batches: number;
  estimatedRewards: string;
}> {
  const pendingDays = await getPendingDays(tokenId);
  
  if (pendingDays <= 90) {
    return {
      needsBatching: false,
      batches: 1,
      estimatedRewards: "0"
    };
  }
  
  // 需要分批
  const batches = Math.ceil(pendingDays / 90);
  
  // 估算第一批奖励（90天）
  const stakeInfo = await staking.stakes(tokenId);
  const estimatedRewards = await staking.calculatePendingRewards(tokenId);
  
  return {
    needsBatching: true,
    batches,
    estimatedRewards: ethers.utils.formatEther(estimatedRewards)
  };
}
```

### 完整的链下实现示例

```typescript
/**
 * 完整的智能分批领取实现（链下逻辑）
 */
class SmartClaimService {
  
  /**
   * 获取用户的领取计划
   */
  async getClaimPlan(userAddress: string): Promise<ClaimPlan[]> {
    // 1. 获取用户所有质押的 NFT
    const nftIds = await this.getUserStakedNFTs(userAddress);
    
    const plan = [];
    
    // 2. 为每个 NFT 制定计划
    for (const tokenId of nftIds) {
      const pendingDays = await this.getPendingDays(tokenId);
      
      if (pendingDays === 0) {
        continue; // 没有待领取奖励
      }
      
      plan.push({
        tokenId,
        pendingDays,
        batchesNeeded: Math.ceil(pendingDays / 90),
        needsMultipleCalls: pendingDays > 90,
        estimatedRewards: await this.estimateRewards(tokenId)
      });
    }
    
    return plan;
  }
  
  /**
   * 执行智能分批领取
   */
  async claimRewards(
    userAddress: string, 
    tokenIds: number[],
    onProgress?: (progress: Progress) => void
  ): Promise<ClaimResult> {
    
    // 1. 检查每个 NFT 的状态
    const plans = await Promise.all(
      tokenIds.map(id => this.getNFTClaimPlan(id))
    );
    
    // 2. 分组：可以直接领取的 vs 需要分批的
    const directClaim = plans.filter(p => !p.needsMultipleCalls);
    const batchedClaim = plans.filter(p => p.needsMultipleCalls);
    
    const results: ClaimResult = {
      successful: [],
      failed: [],
      pending: []
    };
    
    // 3. 先处理可以直接领取的
    if (directClaim.length > 0) {
      const directIds = directClaim.map(p => p.tokenId);
      try {
        const tx = await staking.batchClaimRewards(directIds);
        const receipt = await tx.wait();
        
        results.successful.push({
          tokenIds: directIds,
          txHash: receipt.transactionHash,
          rewards: await this.calculateTotalRewards(directIds)
        });
        
        onProgress?.({ done: directIds.length, total: tokenIds.length });
      } catch (error) {
        results.failed.push({
          tokenIds: directIds,
          error: error.message
        });
      }
    }
    
    // 4. 处理需要分批的
    for (const plan of batchedClaim) {
      const { tokenId, batchesNeeded } = plan;
      
      for (let i = 0; i < batchesNeeded; i++) {
        try {
          // 调用合约（会自动限制为90天）
          const tx = await staking.claimRewards(tokenId);
          const receipt = await tx.wait();
          
          results.successful.push({
            tokenIds: [tokenId],
            txHash: receipt.transactionHash,
            rewards: await this.getRewardsForToken(tokenId)
          });
          
          // 检查是否还有剩余
          const remainingDays = await this.getPendingDays(tokenId);
          if (remainingDays === 0) {
            break; // 全部领取完成
          }
          
          onProgress?.({ 
            done: results.successful.length, 
            total: tokenIds.length 
          });
          
        } catch (error) {
          results.failed.push({
            tokenIds: [tokenId],
            error: error.message
          });
          break; // 失败了就停止这个 NFT 的领取
        }
      }
    }
    
    // 5. 检查还有哪些需要后续领取
    for (const tokenId of tokenIds) {
      const pendingDays = await this.getPendingDays(tokenId);
      if (pendingDays > 0) {
        const estimatedRewards = await this.estimateRewards(tokenId);
        
        results.pending.push({
          tokenId,
          remainingDays: pendingDays,
          estimatedRewards,
          batchesNeeded: Math.ceil(pendingDays / 90)
        });
      }
    }
    
    return results;
  }
  
  // 辅助方法
  private async getPendingDays(tokenId: number): Promise<number> {
    const stakeInfo = await staking.stakes(tokenId);
    const currentTime = await staking.getCurrentTimestamp();
    const pendingDays = Math.floor(
      (Number(currentTime) - Number(stakeInfo.lastClaimTime)) / 86400
    );
    return pendingDays;
  }
  
  private async estimateRewards(tokenId: number): Promise<string> {
    const rewards = await staking.calculatePendingRewards(tokenId);
    return ethers.utils.formatEther(rewards);
  }
  
  private async getUserStakedNFTs(userAddress: string): Promise<number[]> {
    // 从 StakingReader 或合约获取
    const reader = await ethers.getContractAt(
      "StakingReader", 
      STAKING_READER_ADDRESS
    );
    return await reader.getUserStakedTokens(userAddress);
  }
}

// 使用示例
const service = new SmartClaimService();

// 后端 API 调用
app.post('/api/staking/claim-rewards', async (req, res) => {
  const { userAddress, tokenIds, claimAll } = req.body;
  
  try {
    // 获取要领取的 NFT 列表
    const nftIds = claimAll 
      ? await service.getUserStakedNFTs(userAddress)
      : tokenIds;
    
    // 执行智能分批领取
    const result = await service.claimRewards(
      userAddress, 
      nftIds,
      (progress) => {
        // 实时反馈进度给前端
        io.emit('claim-progress', {
          userAddress,
          progress: `${progress.done}/${progress.total}`
        });
      }
    );
    
    res.json({ success: true, data: result });
    
  } catch (error) {
    res.status(500).json({ success: false, error: error.message });
  }
});
```

### 类型定义

```typescript
interface ClaimPlan {
  tokenId: number;
  pendingDays: number;
  batchesNeeded: number;
  needsMultipleCalls: boolean;
  estimatedRewards: string;
}

interface Progress {
  done: number;
  total: number;
}

interface ClaimResult {
  successful: Array<{
    tokenIds: number[];
    txHash: string;
    rewards: string;
  }>;
  failed: Array<{
    tokenIds: number[];
    error: string;
  }>;
  pending: Array<{
    tokenId: number;
    remainingDays: number;
    estimatedRewards: string;
    batchesNeeded: number;
  }>;
}
```

---

## 后端快速开始（简化版）

### 最简单的实现

```typescript
// services/ClaimService.ts - 最简化版本
import { ethers } from "ethers";
import StakingABI from "./abis/Staking.json";

class ClaimService {
  private staking: ethers.Contract;
  
  constructor() {
    this.staking = new ethers.Contract(
      STAKING_ADDRESS,
      StakingABI,
      provider
    );
  }
  
  /**
   * 智能领取 - 自动处理分批
   */
  async smartClaim(tokenIds: number[], userAddress: string) {
    const results = [];
    
    for (const tokenId of tokenIds) {
      let remaining = true;
      
      while (remaining) {
        try {
          // 调用 claimRewards（会自动限制 90 天）
          const tx = await this.staking.connect(userAddress).claimRewards(tokenId);
          await tx.wait();
          
          // 检查是否还有剩余
          const pendingDays = await this.getPendingDays(tokenId);
          
          if (pendingDays === 0) {
            results.push({ tokenId, status: "complete" });
            remaining = false;
          } else {
            results.push({ 
              tokenId, 
              status: "partial", 
              remainingDays: pendingDays 
            });
            // 继续循环，下次再领取
          }
        } catch (error) {
          results.push({ 
            tokenId, 
            status: "failed", 
            error: error.message 
          });
          remaining = false;
        }
      }
    }
    
    return results;
  }
  
  /**
   * 计算待领取天数
   */
  private async getPendingDays(tokenId: number): Promise<number> {
    const stakeInfo = await this.staking.stakes(tokenId);
    const currentTime = await this.staking.getCurrentTimestamp();
    
    return Math.floor(
      (currentTime.toNumber() - stakeInfo.lastClaimTime.toNumber()) / 86400
    );
  }
}
```

### API 实现示例

```typescript
// routes/staking.ts
import express from "express";
const router = express.Router();

router.post("/claim-rewards", async (req, res) => {
  const { userAddress, tokenIds } = req.body;
  
  try {
    const service = new ClaimService();
    const results = await service.smartClaim(tokenIds, userAddress);
    
    res.json({
      success: true,
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

### 使用流程

```typescript
// 前端调用
POST /api/staking/claim-rewards
{
  "userAddress": "0x...",
  "tokenIds": [4812, 3416]
}

// 后端自动处理
1. 对每个 NFT 调用 claimRewards()
2. 检查是否还有待领取天数
3. 如果有，继续调用
4. 直到全部领取完成

// 返回结果
{
  "success": true,
  "results": [
    {
      "tokenId": 4812,
      "status": "complete",
      "txHash": "0x..."
    },
    {
      "tokenId": 3416,
      "status": "complete",
      "txHash": "0x..."
    }
  ]
}
```

---

## ⚠️ 重要变更说明

### 批量函数改为 Owner Only

**v4.1.0 版本变更**：

`batchClaimRewards` 和 `batchUnstake` 现在只有合约 owner 可以调用！

### 新的函数签名

```solidity
// 旧版本 (v4.0.0)
function batchClaimRewards(uint256[] calldata tokenIds) 
    external nonReentrant whenNotPaused;

// 新版本 (v4.1.0)
function batchClaimRewards(
    address userAddress,  // 新增参数
    uint256[] calldata tokenIds
) external nonReentrant whenNotPaused onlyOwner;
```

### 为什么这样修改？

1. **安全性**：防止恶意用户批量操作
2. **用户体验**：后端统一处理，用户只需点击一次
3. **Gas 优化**：后端可以智能分批，自动重试

### 后端调用方式

```typescript
// 使用 owner 私钥
const ownerWallet = new ethers.Wallet(OWNER_PRIVATE_KEY, provider);
const staking = new ethers.Contract(STAKING_ADDRESS, ABI, ownerWallet);

// 调用批量质押
await staking.batchStake(
  "0xDf3715f4693CC308c961AaF0AacD56400E229F43",  // userAddress
  [4812, 3416, 3393]  // tokenIds
);

// 调用批量领取
await staking.batchClaimRewards(
  "0xDf3715f4693CC308c961AaF0AacD56400E229F43",  // userAddress
  [4812, 3416, 3393]  // tokenIds
);
```

### 环境变量配置

```bash
# .env
STAKING_ADDRESS=0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5
STAKING_OWNER_PRIVATE_KEY=0x...  # 合约 owner 私钥
```

### 安全性注意事项

1. **保护 owner 私钥**：
   - 使用环境变量，不要硬编码
   - 使用密钥管理系统（如 AWS KMS）
   - 定期轮换密钥

2. **权限验证**：
   - 验证 userAddress 是否为真实用户
   - 验证 tokenIds 是否属于该用户
   - 添加 rate limiting

3. **监控和日志**：
   - 记录所有批量操作
   - 监控异常调用
   - 设置告警

---

## Go Bindings 使用指南

### 安装

```bash
cd cpop-abis
go mod tidy
```

### 导入

```go
import (
    "github.com/chapool/account-abstraction/cpop-abis"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/ethclient"
)
```

### 连接到 Staking 合约

```go
// 连接到 Sepolia 网络
client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/YOUR_KEY")
if err != nil {
    log.Fatal(err)
}

// Staking 合约地址
stakingAddress := common.HexToAddress("0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5")

// 创建合约实例
staking, err := cpop.NewStaking(stakingAddress, client)
if err != nil {
    log.Fatal(err)
}

// StakingReader 合约地址
readerAddress := common.HexToAddress("0xbbEe6e5FF90f0B6EFF185F73c71b0deE6Fe9D0A6")
reader, err := cpop.NewStakingReader(readerAddress, client)
if err != nil {
    log.Fatal(err)
}
```

### 调用批量质押函数

```go
// 准备参数
userAddress := common.HexToAddress("0xDf3715f4693CC308c961AaF0AacD56400E229F43")
tokenIds := []*big.Int{big.NewInt(4812), big.NewInt(3416), big.NewInt(3393)}

// 获取 owner 私钥
ownerPrivateKey := os.Getenv("STAKING_OWNER_PRIVATE_KEY")
auth, err := bind.NewKeyedTransactorWithChainID(
    privateKey, 
    big.NewInt(11155111), // Sepolia chain ID
)

// 调用批量质押
tx, err := staking.BatchStake(auth, userAddress, tokenIds)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Transaction hash: %s\n", tx.Hash().Hex())

// 等待确认
receipt, err := bind.WaitMined(context.Background(), client, tx)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Transaction confirmed in block: %d\n", receipt.BlockNumber)
```

### 调用批量领取函数

```go
// 准备参数
userAddress := common.HexToAddress("0xDf3715f4693CC308c961AaF0AacD56400E229F43")
tokenIds := []*big.Int{big.NewInt(4812), big.NewInt(3416), big.NewInt(3393)}

// 获取 owner 私钥
ownerPrivateKey := os.Getenv("STAKING_OWNER_PRIVATE_KEY")
auth, err := bind.NewKeyedTransactorWithChainID(
    privateKey, 
    big.NewInt(11155111), // Sepolia chain ID
)

// 调用批量领取
tx, err := staking.BatchClaimRewards(auth, userAddress, tokenIds)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Transaction hash: %s\n", tx.Hash().Hex())

// 等待确认
receipt, err := bind.WaitMined(context.Background(), client, tx)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Transaction confirmed in block: %d\n", receipt.BlockNumber)
```

### 调用批量解质押函数

```go
// 调用批量解质押
tx, err := staking.BatchUnstake(auth, userAddress, tokenIds)
if err != nil {
    log.Fatal(err)
}

// 等待确认...
```

### 查询功能

```go
// 查询组合状态
comboStatus, err := staking.GetComboStatus(nil, userAddress, 3) // Level 3 = A级
if err != nil {
    log.Fatal(err)
}

fmt.Printf("当前数量: %d\n", comboStatus.Count)
fmt.Printf("加成值: %d 基点 (%.2f%%)\n", comboStatus.Bonus, float64(comboStatus.Bonus)/100)

// 查询待领取奖励
pendingRewards, err := staking.CalculatePendingRewards(nil, big.NewInt(4812))
if err != nil {
    log.Fatal(err)
}
fmt.Printf("待领取奖励: %s CPOP\n", pendingRewards.String())
```

### 使用 StakingReader 查询详细信息

```go
// 查询用户奖励统计
stats, err := reader.GetUserRewardStats(nil, userAddress)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("总 NFT 数量: %d\n", stats.TotalNFTs)
fmt.Printf("总待领取奖励: %s CPOP\n", stats.TotalPendingRewards.String())

// 查询组合汇总
comboSummary, err := reader.GetUserComboSummary(nil, userAddress)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("A 级组合加成: %d 基点\n", comboSummary.LevelComboBonuses[2]) // Index 2 = A级
```

---

**文档状态**: ✅ 完成  
**版本**: v4.1.3  
**待执行**: 后端实现 → 测试 → 上线

