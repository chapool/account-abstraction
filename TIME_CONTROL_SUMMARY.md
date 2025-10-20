# Staking 合约时间控制功能总结

## 🎯 实现目标

为 Staking 合约添加可控制的时间戳功能，方便测试时间相关的质押逻辑，无需等待真实时间流逝。

## ✅ 已完成的修改

### 1. 合约修改（Staking.sol）

#### 新增存储变量
```solidity
bool public testMode;           // 是否启用测试模式
uint256 public testTimestamp;   // 测试模式下的时间戳
```

#### 新增核心函数

1. **内部时间获取函数**
   ```solidity
   function _getCurrentTimestamp() internal view returns (uint256)
   ```
   - 根据 testMode 返回测试时间戳或真实区块时间戳

2. **测试模式管理函数**
   ```solidity
   function enableTestMode(uint256 initialTimestamp) external onlyOwner
   function disableTestMode() external onlyOwner
   function setTestTimestamp(uint256 timestamp) external onlyOwner
   ```

3. **时间快进函数**
   ```solidity
   function fastForwardTime(uint256 seconds_) external onlyOwner
   function fastForwardDays(uint256 days_) external onlyOwner
   function fastForwardMinutes(uint256 minutes_) external onlyOwner
   ```

#### 全局替换
将合约中所有 `block.timestamp` 替换为 `_getCurrentTimestamp()`，涉及以下场景：
- 质押时记录时间
- 领取奖励时更新时间
- 取消质押时计算时间
- 奖励计算中的时间判断
- 组合状态的次日生效机制
- 历史调整记录

### 2. 管理脚本

创建了 4 个管理脚本：

#### enable-test-mode.ts
- 启用测试模式
- 设置初始测试时间戳
- 显示当前状态

#### fast-forward-time.ts  
- 快进指定时间（支持分钟、天、秒）
- 显示快进前后的时间对比
- 参数：`--minutes`, `--days`, `--seconds`

#### check-time-status.ts
- 查看当前时间状态
- 对比测试时间和真实时间
- 显示可用操作提示

#### disable-test-mode.ts
- 禁用测试模式
- 恢复使用真实时间
- 重置测试时间戳

### 3. 文档

#### Staking-Test-Mode-Guide.md
完整的使用指南，包含：
- 功能说明
- 使用场景示例
- 脚本说明
- 完整测试流程
- 常见问题解答
- 技术细节

## 📊 使用示例

### 基本流程

```bash
# 1. 启用测试模式
npx hardhat run scripts/enable-test-mode.ts --network sepolia

# 2. 质押 NFT（使用现有脚本）
# ...

# 3. 快进 180 分钟（模拟 SSS 级完整衰减周期）
npx hardhat run scripts/fast-forward-time.ts --network sepolia -- --minutes 180

# 4. 查看奖励
npx hardhat run scripts/test-user-daily-rewards.ts --network sepolia

# 5. 禁用测试模式
npx hardhat run scripts/disable-test-mode.ts --network sepolia
```

### 配合 StakingConfig 使用

```bash
# 第一步：设置测试环境
npx hardhat run scripts/update-staking-config-for-testing.ts --network sepolia
npx hardhat run scripts/enable-test-mode.ts --network sepolia

# 第二步：进行测试
# ... 质押、快进时间、领取奖励等 ...

# 第三步：恢复生产环境
npx hardhat run scripts/disable-test-mode.ts --network sepolia
npx hardhat run scripts/restore-staking-config-production.ts --network sepolia
```

## 🔑 关键特性

### 1. 安全性
- ✅ 只有 owner 可以控制测试模式
- ✅ 时间只能向前，不能倒退
- ✅ 测试模式不影响其他合约
- ✅ 可以随时禁用恢复正常

### 2. 灵活性
- ✅ 支持秒、分钟、天三种时间单位
- ✅ 可以精确控制时间点
- ✅ 可以查看时间状态
- ✅ 支持多次快进

### 3. 兼容性
- ✅ 不影响现有功能
- ✅ 已质押的 NFT 正常工作
- ✅ 所有奖励计算逻辑不变
- ✅ 可以无缝切换测试/生产模式

## 🧪 测试场景覆盖

现在可以快速测试以下场景：

### 1. 衰减机制（Phase-based Decay）
- SSS 级：180 分钟 → 完整衰减周期
- SS 级：90 分钟 → 完整衰减周期
- S 级：60 分钟 → 完整衰减周期
- A 级：45 分钟 → 完整衰减周期
- B 级：30 分钟 → 完整衰减周期
- C 级：20 分钟 → 完整衰减周期

### 2. 持续质押奖励
- 30 天 → 30 分钟：10% 奖励
- 90 天 → 90 分钟：20% 奖励

### 3. 组合奖励
- 3 NFT：7 分钟 → 5% 奖励（次日生效）
- 5 NFT：15 分钟 → 10% 奖励（次日生效）
- 10 NFT：30 分钟 → 20% 奖励（次日生效）

### 4. 提前取消惩罚
- 7 分钟内取消 → 50% 惩罚

### 5. 季度调整
- 可以快进到下一个季度
- 测试历史调整记录功能

## 📁 文件清单

### 合约修改
- `contracts/CPNFT/Staking.sol` - 添加时间控制功能

### 新增脚本
- `scripts/enable-test-mode.ts` - 启用测试模式
- `scripts/fast-forward-time.ts` - 快进时间
- `scripts/check-time-status.ts` - 查看时间状态
- `scripts/disable-test-mode.ts` - 禁用测试模式

### 文档
- `docs/Staking-Test-Mode-Guide.md` - 完整使用指南
- `TIME_CONTROL_SUMMARY.md` - 本文件

## 🎓 最佳实践

### 1. 测试前准备
```bash
# 检查当前状态
npx hardhat run scripts/check-time-status.ts --network sepolia

# 如果已启用，先禁用
npx hardhat run scripts/disable-test-mode.ts --network sepolia

# 重新启用获得干净的测试环境
npx hardhat run scripts/enable-test-mode.ts --network sepolia
```

### 2. 测试中
```bash
# 随时查看时间状态
npx hardhat run scripts/check-time-status.ts --network sepolia

# 分步快进，观察每个阶段的变化
npx hardhat run scripts/fast-forward-time.ts --network sepolia -- --minutes 30
# 查看奖励...
npx hardhat run scripts/fast-forward-time.ts --network sepolia -- --minutes 30
# 再次查看奖励...
```

### 3. 测试后
```bash
# 务必禁用测试模式
npx hardhat run scripts/disable-test-mode.ts --network sepolia

# 如果修改了 StakingConfig，也要恢复
npx hardhat run scripts/restore-staking-config-production.ts --network sepolia
```

## ⚠️ 注意事项

1. **生产环境**
   - 测试模式仅供测试使用
   - 生产环境必须保持禁用状态
   - 部署时默认为禁用

2. **时间限制**
   - 时间只能向前
   - 不能设置为过去
   - 如需重置，先禁用再启用

3. **权限控制**
   - 只有 owner 可以操作
   - 确保私钥安全
   - 定期检查状态

4. **与其他合约的交互**
   - 测试时间只影响 Staking 合约
   - 其他合约仍使用真实时间
   - 注意跨合约交互的时间一致性

## 🚀 效率提升

### 之前（无时间控制）
- 测试 180 天衰减：需要等待 **6 个月**
- 测试 90 天持续奖励：需要等待 **3 个月**
- 测试组合奖励次日生效：需要等待 **1-3 天**

### 现在（有时间控制）
- 测试 180 天衰减：只需 **3 小时**（180 分钟）
- 测试 90 天持续奖励：只需 **1.5 小时**（90 分钟）
- 测试组合奖励次日生效：只需 **几分钟**

**效率提升：从数月缩短到数小时！** 🎉

## 🔮 未来扩展

可能的增强功能：
- [ ] 添加时间快照功能
- [ ] 支持批量时间操作
- [ ] 添加时间回退功能（仅限测试环境）
- [ ] 集成到自动化测试套件
- [ ] 添加时间事件日志

## 📞 支持

如有问题或建议，请参考：
- 使用指南：`docs/Staking-Test-Mode-Guide.md`
- 合约代码：`contracts/CPNFT/Staking.sol`
- 示例脚本：`scripts/enable-test-mode.ts` 等

---

**创建日期：** 2024-10-20  
**版本：** 1.0.0  
**状态：** ✅ 已完成并可用

