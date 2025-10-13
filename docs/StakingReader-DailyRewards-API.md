# StakingReader 每日收益统计 API 文档

## 📋 概述

StakingReader 合约新增了强大的每日收益统计功能，可以帮助前端展示用户的每日收益详情，包括基础收益、各种加成和最终收益。

## 🎯 核心功能

### 1. 获取用户所有NFT的每日收益汇总

**函数签名：**
```solidity
function getUserDailyRewards(address user) 
    external view 
    returns (UserDailyRewards memory)
```

**返回数据结构：**
```solidity
struct UserDailyRewards {
    uint256 totalBaseReward;           // 总基础收益（CPOP/天）
    uint256 totalDecayedReward;        // 总衰减后收益
    uint256 totalComboBonus;           // 总Combo加成（百分比，10000=100%）
    uint256 totalDynamicMultiplier;    // 总动态乘数（百分比）
    uint256 totalFinalReward;          // 总最终收益
    uint256 totalBonus;                // ⭐ 总加成百分比（从基础到最终）
    uint256[6] baseRewardPerLevel;     // 各等级基础收益
    uint256[6] finalRewardPerLevel;    // 各等级最终收益
    uint256[6] nftCountPerLevel;       // 各等级NFT数量
}
```

**示例调用（JavaScript/TypeScript）：**
```javascript
const dailyRewards = await stakingReaderContract.getUserDailyRewards(userAddress);

console.log("总基础收益:", ethers.utils.formatEther(dailyRewards.totalBaseReward), "CPOP/天");
console.log("总最终收益:", ethers.utils.formatEther(dailyRewards.totalFinalReward), "CPOP/天");

// totalBonus: 10000=100%, 11000=110%, 9000=90%
const bonusValue = Number(dailyRewards.totalBonus);
const bonusPercent = (bonusValue / 100).toFixed(2);
const bonusChange = ((bonusValue - 10000) / 100).toFixed(2);
console.log("⭐ 总乘数:", bonusPercent + "%", `(${bonusChange > 0 ? '+' : ''}${bonusChange}%)`);
```

**前端展示建议：**
```jsx
// React 示例
function DailyRewardsDisplay({ dailyRewards }) {
  const bonusValue = Number(dailyRewards.totalBonus);
  const bonusPercent = (bonusValue / 100).toFixed(2);
  const bonusChange = ((bonusValue - 10000) / 100).toFixed(2);
  const bonusSign = bonusValue > 10000 ? "+" : "";
  
  const baseReward = parseFloat(ethers.utils.formatEther(dailyRewards.totalBaseReward));
  const finalReward = parseFloat(ethers.utils.formatEther(dailyRewards.totalFinalReward));
  
  return (
    <div className="daily-rewards">
      <h3>每日收益</h3>
      <div className="reward-row">
        <span>基础收益:</span>
        <span>{baseReward.toFixed(2)} CPOP/天</span>
      </div>
      <div className="reward-row highlight">
        <span>⭐ 总乘数:</span>
        <span className="bonus">{bonusPercent}% ({bonusSign}{bonusChange}%)</span>
      </div>
      <div className="reward-row total">
        <span>最终收益:</span>
        <span className="final">{finalReward.toFixed(2)} CPOP/天</span>
      </div>
      
      <div className="breakdown">
        <small>├─ Combo加成: +{(Number(dailyRewards.totalComboBonus) / 100).toFixed(2)}%</small>
        <small>└─ 动态乘数: {(Number(dailyRewards.totalDynamicMultiplier) / 100).toFixed(2)}%</small>
      </div>
    </div>
  );
}
```

---

### 2. 按等级查询每日收益

**函数签名：**
```solidity
function getUserDailyRewardsByLevel(address user, uint8 level) 
    external view 
    returns (
        uint256 baseReward,
        uint256 finalReward,
        uint256 nftCount
    )
```

**参数：**
- `user`: 用户地址
- `level`: NFT等级 (1=C, 2=B, 3=A, 4=S, 5=SS, 6=SSS)

**示例调用：**
```javascript
// 查询用户 Level A (3) 的每日收益
const { baseReward, finalReward, nftCount } = 
    await stakingReaderContract.getUserDailyRewardsByLevel(userAddress, 3);

console.log("Level A 质押数量:", nftCount.toString());
console.log("Level A 基础收益:", ethers.utils.formatEther(baseReward), "CPOP/天");
console.log("Level A 最终收益:", ethers.utils.formatEther(finalReward), "CPOP/天");

if (baseReward > 0) {
    const bonus = ((Number(finalReward) / Number(baseReward) - 1) * 100).toFixed(2);
    console.log("Level A 加成:", bonus + "%");
}
```

---

### 3. 获取单个NFT的每日收益详情

**函数签名：**
```solidity
function getNFTDailyRewardBreakdown(uint256 tokenId) 
    external view 
    returns (
        uint256 baseReward,
        uint256 decayedReward,
        uint256 comboBonus,
        uint256 dynamicMultiplier,
        uint256 finalReward
    )
```

**示例调用：**
```javascript
const breakdown = await stakingReaderContract.getNFTDailyRewardBreakdown(2645);

console.log("基础收益:", ethers.utils.formatEther(breakdown.baseReward), "CPOP/天");
console.log("衰减后:", ethers.utils.formatEther(breakdown.decayedReward), "CPOP/天");
console.log("Combo加成:", (Number(breakdown.comboBonus) / 100).toFixed(2) + "%");
console.log("动态乘数:", (Number(breakdown.dynamicMultiplier) / 100).toFixed(2) + "%");
console.log("最终收益:", ethers.utils.formatEther(breakdown.finalReward), "CPOP/天");
```

---

## 📊 计算公式说明

### 总乘数（totalBonus）
```
totalBonus = 最终收益 / 基础收益 × 10000
```
- 格式：10000 为基数（100%）
- 例如：
  - `totalBonus = 11000` → 110%（+10% 加成）
  - `totalBonus = 10000` → 100%（无变化）
  - `totalBonus = 9000` → 90%（-10% 衰减）

### Combo加成（totalComboBonus）
```
totalComboBonus = (最终收益 - 衰减后收益) / 衰减后收益 × 10000
```

### 动态乘数（totalDynamicMultiplier）
```
totalDynamicMultiplier = 衰减后收益 / 基础收益 × 10000
```

### 收益计算流程
```
基础收益 
  → 应用衰减 → 衰减后收益
  → 应用季度乘数
  → 应用动态乘数
  → 应用Combo加成 → 最终收益
```

---

## 🎨 前端 UI 设计建议

### 1. 收益概览卡片

```jsx
<Card>
  <h2>每日收益概览</h2>
  
  <div className="main-metric">
    <span className="value">{finalReward.toFixed(2)}</span>
    <span className="unit">CPOP/天</span>
    <span className="bonus">+{totalBonus}%</span>
  </div>
  
  <ProgressBar 
    from={baseReward} 
    to={finalReward}
    label={`+${totalBonus}% 加成`}
  />
  
  <div className="breakdown">
    <div>基础: {baseReward.toFixed(2)}</div>
    <div>Combo: +{comboBonus}%</div>
    <div>动态: {dynamicMult}%</div>
  </div>
</Card>
```

### 2. 等级分布图表

```jsx
const levelData = [
  { level: 'C', base: levelBase[0], final: levelFinal[0], count: levelCount[0] },
  { level: 'B', base: levelBase[1], final: levelFinal[1], count: levelCount[1] },
  // ...
];

<BarChart data={levelData}>
  <Bar dataKey="base" fill="#gray" />
  <Bar dataKey="final" fill="#gold" />
</BarChart>
```

### 3. 收益预测

```javascript
function calculateProjection(dailyReward) {
  return {
    daily: dailyReward,
    weekly: dailyReward * 7,
    monthly: dailyReward * 30,
    yearly: dailyReward * 365
  };
}
```

---

## ⚠️ 注意事项

1. **收益会动态变化**
   - 衰减机制随时间递减
   - Combo状态变化影响加成
   - 平台质押率影响动态乘数

2. **Gas 优化**
   - 批量查询用户所有NFT的收益，一次调用即可
   - 避免循环调用单个NFT的查询

3. **数据精度**
   - 所有收益值以 wei 为单位
   - 百分比以 10000 为基数（10000 = 100%）
   - 使用 `ethers.utils.formatEther()` 转换显示

4. **错误处理**
   ```javascript
   try {
     const rewards = await stakingReaderContract.getUserDailyRewards(user);
     // 处理数据...
   } catch (error) {
     if (error.message.includes("NFT is not staked")) {
       // 处理未质押情况
     }
   }
   ```

---

## 🚀 完整示例

查看测试脚本了解完整用法：
- `scripts/test-user-daily-rewards.ts`

运行测试：
```bash
npx hardhat run scripts/test-user-daily-rewards.ts --network sepoliaCustom
```

---

## 📞 支持

如有问题，请查看：
- [StakingReader 合约源码](../contracts/CPNFT/StakingReader.sol)
- [测试脚本示例](../scripts/test-user-daily-rewards.ts)

