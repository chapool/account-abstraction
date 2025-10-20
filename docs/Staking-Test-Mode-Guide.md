# Staking 测试模式使用指南

## 概述

Staking 合约现在支持**测试模式**，允许你控制合约内部的时间流逝，无需等待真实的时间。这对于测试时间相关的功能（如质押奖励、衰减、持续质押奖励等）非常有用。

## 功能说明

### 时间控制机制

- **生产模式**（默认）：合约使用真实的 `block.timestamp`
- **测试模式**：合约使用可控制的测试时间戳
  - 可以快进任意时间（分钟、小时、天）
  - 时间只能向前，不能倒退
  - 只有合约 owner 可以控制

### 核心函数

```solidity
// 启用测试模式
function enableTestMode(uint256 initialTimestamp) external onlyOwner

// 禁用测试模式（恢复使用真实时间）
function disableTestMode() external onlyOwner

// 设置测试时间戳
function setTestTimestamp(uint256 timestamp) external onlyOwner

// 快进指定秒数
function fastForwardTime(uint256 seconds_) external onlyOwner

// 快进指定分钟数
function fastForwardMinutes(uint256 minutes_) external onlyOwner

// 快进指定天数
function fastForwardDays(uint256 days_) external onlyOwner
```

## 使用场景

### 场景 1：测试 SSS 级 NFT 的 180 天衰减

在测试模式下，配合 StakingConfig 的时间单位调整（1天=1分钟），你可以在 3 小时内测试完整的 180 天衰减周期。

```bash
# 1. 启用测试模式
npx hardhat run scripts/enable-test-mode.ts --network sepolia

# 2. 质押 NFT
# ... 进行质押操作 ...

# 3. 快进 180 分钟（相当于 180 天）
npx hardhat run scripts/fast-forward-time.ts --network sepolia -- --minutes 180

# 4. 领取奖励
# ... 进行领取操作 ...

# 5. 禁用测试模式（测试完成后）
npx hardhat run scripts/disable-test-mode.ts --network sepolia
```

### 场景 2：测试持续质押奖励（30天/90天）

```bash
# 启用测试模式
npx hardhat run scripts/enable-test-mode.ts --network sepolia

# 质押 NFT
# ...

# 快进 30 分钟（相当于 30 天）
npx hardhat run scripts/fast-forward-time.ts --network sepolia -- --minutes 30

# 检查奖励（应该获得 10% 持续质押奖励）
# ...

# 继续快进 60 分钟（总共 90 天）
npx hardhat run scripts/fast-forward-time.ts --network sepolia -- --minutes 60

# 检查奖励（应该获得 20% 持续质押奖励）
# ...
```

### 场景 3：测试组合奖励的次日生效机制

```bash
# 启用测试模式
npx hardhat run scripts/enable-test-mode.ts --network sepolia

# 质押第 1 个 NFT
# ...

# 快进 1 天（1 分钟）
npx hardhat run scripts/fast-forward-time.ts --network sepolia -- --minutes 1

# 质押第 2 个 NFT
# ...

# 快进 1 天
npx hardhat run scripts/fast-forward-time.ts --network sepolia -- --minutes 1

# 质押第 3 个 NFT（达到 3 NFT 组合）
# ...

# 快进 1 天（组合奖励次日生效）
npx hardhat run scripts/fast-forward-time.ts --network sepolia -- --minutes 1

# 现在应该可以看到 5% 的组合奖励
# ...
```

## 脚本说明

### 1. enable-test-mode.ts

启用测试模式。

```bash
npx hardhat run scripts/enable-test-mode.ts --network sepolia
```

**输出示例：**
```
⏰ 启用 Staking 合约测试模式...

操作者地址: 0x...
Staking 地址: 0x...

✅ 测试模式已启用！
========================================
测试模式: ✓ 已启用
测试时间戳: 1234567890
对应日期: 2024-01-01T00:00:00.000Z
========================================
```

### 2. fast-forward-time.ts

快进时间。

```bash
# 快进 60 分钟
npx hardhat run scripts/fast-forward-time.ts --network sepolia -- --minutes 60

# 快进 1 天
npx hardhat run scripts/fast-forward-time.ts --network sepolia -- --days 1

# 快进 3600 秒（1 小时）
npx hardhat run scripts/fast-forward-time.ts --network sepolia -- --seconds 3600
```

**参数：**
- `-m, --minutes <数量>`: 快进指定分钟数
- `-d, --days <数量>`: 快进指定天数
- `-s, --seconds <数量>`: 快进指定秒数

**输出示例：**
```
⏰ Staking 合约时间快进工具

当前时间: 2024-01-01T00:00:00.000Z
⏩ 快进 60 分钟...

✅ 时间快进成功！
新时间: 2024-01-01T01:00:00.000Z
实际快进: 60 分钟
快进秒数: 3600 秒
```

### 3. check-time-status.ts

查看当前时间状态。

```bash
npx hardhat run scripts/check-time-status.ts --network sepolia
```

**输出示例：**
```
⏰ Staking 合约时间状态

========================================
时间状态
========================================
测试模式: ✅ 已启用

📍 当前使用时间（测试时间）:
  时间戳: 1234567890
  日期: 2024-01-01T00:00:00.000Z

🕐 真实区块时间:
  时间戳: 1234564290
  日期: 2023-12-31T23:00:00.000Z

⏩ 测试时间领先真实时间: 3600 秒
   相当于: 60.00 分钟
========================================
```

### 4. disable-test-mode.ts

禁用测试模式，恢复使用真实时间。

```bash
npx hardhat run scripts/disable-test-mode.ts --network sepolia
```

## 完整测试流程示例

### 测试 SSS 级 NFT 完整生命周期

```bash
# 步骤 1: 启用测试模式
npx hardhat run scripts/enable-test-mode.ts --network sepolia

# 步骤 2: 质押 SSS 级 NFT
# 使用你的质押脚本...

# 步骤 3: 检查初始状态
npx hardhat run scripts/check-time-status.ts --network sepolia

# 步骤 4: 快进 60 分钟（测试第一个衰减周期）
npx hardhat run scripts/fast-forward-time.ts --network sepolia -- --minutes 60

# 步骤 5: 查看奖励
npx hardhat run scripts/test-user-daily-rewards.ts --network sepolia

# 步骤 6: 快进 60 分钟（测试第二个衰减周期）
npx hardhat run scripts/fast-forward-time.ts --network sepolia -- --minutes 60

# 步骤 7: 再次查看奖励（应该看到衰减效果）
npx hardhat run scripts/test-user-daily-rewards.ts --network sepolia

# 步骤 8: 快进 60 分钟（总共 180 分钟，测试最大衰减）
npx hardhat run scripts/fast-forward-time.ts --network sepolia -- --minutes 60

# 步骤 9: 最后查看奖励（应该达到最大衰减）
npx hardhat run scripts/test-user-daily-rewards.ts --network sepolia

# 步骤 10: 领取奖励
# 使用你的领取脚本...

# 步骤 11: 取消质押
# 使用你的取消质押脚本...

# 步骤 12: 测试完成，禁用测试模式
npx hardhat run scripts/disable-test-mode.ts --network sepolia
```

## 注意事项

### ⚠️ 重要提示

1. **测试模式仅供测试使用**
   - 在生产环境中应始终保持测试模式为禁用状态
   - 测试完成后务必禁用测试模式

2. **时间只能向前**
   - 不能将时间设置为过去
   - 如需重置，必须先禁用再重新启用测试模式

3. **配合 StakingConfig 使用**
   - 建议先使用 `update-staking-config-for-testing.ts` 将时间单位调整为分钟
   - 测试完成后使用 `restore-staking-config-production.ts` 恢复生产配置

4. **权限控制**
   - 只有合约 owner 可以控制测试模式
   - 确保私钥安全

5. **已质押的 NFT**
   - 启用测试模式不会影响已质押的 NFT
   - 奖励计算会使用测试时间戳
   - 禁用测试模式后，会恢复使用真实时间

## 配合 StakingConfig 的完整测试流程

### 1. 设置测试环境

```bash
# 步骤 1: 更新 StakingConfig（1天 = 1分钟）
npx hardhat run scripts/update-staking-config-for-testing.ts --network sepolia

# 步骤 2: 启用 Staking 测试模式
npx hardhat run scripts/enable-test-mode.ts --network sepolia
```

### 2. 进行测试

```bash
# 质押 NFT
# ...

# 快进时间测试各种场景
npx hardhat run scripts/fast-forward-time.ts --network sepolia -- --minutes 180

# 检查奖励、领取等操作
# ...
```

### 3. 恢复生产环境

```bash
# 步骤 1: 禁用测试模式
npx hardhat run scripts/disable-test-mode.ts --network sepolia

# 步骤 2: 恢复 StakingConfig 配置
npx hardhat run scripts/restore-staking-config-production.ts --network sepolia
```

## 常见问题

### Q1: 测试模式会影响已质押的 NFT 吗？

A: 会的。启用测试模式后，所有时间相关的计算都会使用测试时间戳，包括已质押的 NFT 的奖励计算。

### Q2: 可以在测试模式下取消质押吗？

A: 可以。所有操作都正常工作，只是时间使用测试时间戳。

### Q3: 如何重置测试时间？

A: 先禁用测试模式，再重新启用即可重置。

```bash
npx hardhat run scripts/disable-test-mode.ts --network sepolia
npx hardhat run scripts/enable-test-mode.ts --network sepolia
```

### Q4: 测试模式下的时间会影响其他合约吗？

A: 不会。只有 Staking 合约内部使用测试时间戳，其他合约仍使用真实的 `block.timestamp`。

### Q5: 可以将时间设置为过去吗？

A: 不可以。时间只能向前，这是为了防止意外情况和保持逻辑一致性。

## 技术细节

### 内部实现

合约内部使用 `_getCurrentTimestamp()` 函数来获取时间：

```solidity
function _getCurrentTimestamp() internal view returns (uint256) {
    return testMode ? testTimestamp : block.timestamp;
}
```

所有原来使用 `block.timestamp` 的地方都被替换为 `_getCurrentTimestamp()`，这样就可以在测试模式和生产模式之间无缝切换。

### 存储变量

```solidity
bool public testMode;           // 是否启用测试模式
uint256 public testTimestamp;   // 测试模式下的时间戳
```

## 总结

测试模式为 Staking 合约提供了强大的时间控制能力，让你可以：

- ✅ 快速测试长期质押场景（180 天衰减）
- ✅ 验证时间相关的奖励机制
- ✅ 测试组合奖励的次日生效
- ✅ 验证持续质押奖励（30天/90天）
- ✅ 加速开发和测试流程

配合 StakingConfig 的时间单位调整，你可以在几小时内完成原本需要数月的测试场景！

