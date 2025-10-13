# 🚀 Sepolia 测试网合约地址 (最新)

**更新时间**: 2025-10-13

## ✅ 核心质押系统

| 合约 | 地址 | 状态 |
|------|------|------|
| **StakingReader** | `0x3eCA9E28583C469CAA654fcb49Bd94Ce25C0262C` | ✅ **最新** (2025-10-13 重新部署) |
| **Staking** | `0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5` | ✅ 当前 |
| **StakingConfig** | `0x37196054B23Be5CB977AA3284A3A844a8fe77861` | ✅ 当前 |
| **CPNFT** | `0xcC63bf57EF4b4fE5635cF0745Ae7E2C75A63c7Ed` | ✅ 当前 |

## ❌ 已废弃的地址

| 合约 | 地址 | 原因 |
|------|------|------|
| StakingReader (旧) | `0x6C9f7Fb0376C961FE79cED8cf09EbBBaDBfF0051` | 引用了旧版 Staking 合约 |
| Staking (旧) | `0x23983f63C7Eb0e920Fa73146293A51B215310Ac2` | 已升级 |

## 🔧 前端配置示例

### JavaScript/TypeScript

```javascript
// ✅ 正确的配置
const CONTRACT_ADDRESSES = {
  STAKING_READER: "0x3eCA9E28583C469CAA654fcb49Bd94Ce25C0262C",
  STAKING: "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5",
  STAKING_CONFIG: "0x37196054B23Be5CB977AA3284A3A844a8fe77861",
  CPNFT: "0xcC63bf57EF4b4fE5635cF0745Ae7E2C75A63c7Ed",
  CPOP_TOKEN: "0xA2d58d11c2752b010C2444fa2f795c6cf4cb76Bc",
  ACCOUNT_MANAGER: "0x2E4f862Ba3ee6D84dd19ae9f002F5D8c0C5675ef",
};

// 使用示例
const stakingReader = new ethers.Contract(
  CONTRACT_ADDRESSES.STAKING_READER,
  StakingReaderABI,
  provider
);

// 查询 pending rewards
const pendingRewards = await stakingReader.getPendingRewards(tokenId);
```

### Vue 3 示例

```javascript
// config/contracts.ts
export const SEPOLIA_CONTRACTS = {
  stakingReader: "0x3eCA9E28583C469CAA654fcb49Bd94Ce25C0262C",
  staking: "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5",
  stakingConfig: "0x37196054B23Be5CB977AA3284A3A844a8fe77861",
  cpnft: "0xcC63bf57EF4b4fE5635cF0745Ae7E2C75A63c7Ed",
}

// 在组件中使用
import { SEPOLIA_CONTRACTS } from '@/config/contracts'

const stakingReaderContract = useContract(
  SEPOLIA_CONTRACTS.stakingReader,
  StakingReaderABI
)

const { data: pendingRewards } = await stakingReaderContract.getPendingRewards(tokenId)
```

## 🧪 测试验证

### 测试 getPendingRewards

```bash
# 使用 Hardhat 脚本测试
npx hardhat run scripts/test-exact-error.ts --network sepoliaCustom

# 预期结果
✅ Pending Rewards: 0.0 CPOP  (如果是新质押)
```

### 常见问题

#### Q: 为什么还是报 "Invalid level" 错误？

**A:** 检查以下几点：

1. **前端是否使用了正确的地址？**
   ```javascript
   // ❌ 错误 - 旧地址
   "0x6C9f7Fb0376C961FE79cED8cf09EbBBaDBfF0051"
   
   // ✅ 正确 - 新地址
   "0x3eCA9E28583C469CAA654fcb49Bd94Ce25C0262C"
   ```

2. **清除浏览器缓存和 localStorage**

3. **重新连接钱包**

4. **检查网络是否为 Sepolia (Chain ID: 11155111)**

#### Q: Token 2645 的状态？

```
✅ Token 2645 存在
✅ Owner: 0xc5cCc3c5e4bbb9519Deaf7a8afA29522DA49E33D
✅ Level: 1 (C)
✅ 已质押: true
✅ Pending Rewards: 0.0 CPOP (刚质押)
```

## 📞 支持

如果仍有问题，请检查：
1. 前端使用的合约地址
2. 网络连接（Sepolia）
3. RPC 端点是否正常
4. Token ID 是否正确

---

**最后更新**: 2025-10-13  
**维护者**: Account Abstraction Team

