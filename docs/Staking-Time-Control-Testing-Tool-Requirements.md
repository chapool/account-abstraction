# Staking 时间控制测试工具 - 前端开发需求文档

## 📋 项目概述

### 背景
Staking 合约已升级支持测试模式，可以控制时间流逝以便快速测试。但目前只能通过命令行操作，对非技术测试人员不友好。

### 目标
开发一个 Web 端测试工具，测试人员只需连接钱包、点击按钮即可完成所有测试操作，无需使用命令行。

### 目标用户
- 测试人员（非技术背景）
- 产品经理
- 项目演示

---

## 🎯 核心功能需求

### 1. 钱包连接
**功能描述：**
- 支持 MetaMask 钱包连接
- 自动检测网络（Sepolia 测试网）
- 如果不在 Sepolia 网络，提示切换网络
- 显示连接的钱包地址
- 显示账户余额（Sepolia ETH）

**UI 要求：**
- 顶部导航栏显示"连接钱包"按钮
- 连接后显示地址（缩写格式：0x1234...5678）
- 显示余额（例如：2.31 ETH）
- 提供"断开连接"选项

**技术要点：**
```javascript
// 连接钱包
const accounts = await window.ethereum.request({ 
  method: 'eth_requestAccounts' 
});

// 检查网络
const chainId = await window.ethereum.request({ 
  method: 'eth_chainId' 
});
// Sepolia chainId: 0xaa36a7 (11155111)

// 切换网络
await window.ethereum.request({
  method: 'wallet_switchEthereumChain',
  params: [{ chainId: '0xaa36a7' }],
});
```

**合约地址：**
- Staking Proxy: `0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5`
- 网络：Sepolia (Chain ID: 11155111)

---

### 2. 时间状态显示

**功能描述：**
实时显示当前合约的时间状态

**显示内容：**
```
┌─────────────────────────────────────┐
│  时间状态                            │
├─────────────────────────────────────┤
│  测试模式：  ✅ 已启用 / ❌ 未启用   │
│  测试时间：  2024-10-20 10:30:00   │
│  真实时间：  2024-10-20 10:00:00   │
│  时间差：    +30 分钟              │
└─────────────────────────────────────┘
```

**合约函数调用：**
```javascript
// 读取测试模式状态
const testMode = await stakingContract.testMode();

// 读取测试时间戳
const testTimestamp = await stakingContract.testTimestamp();

// 区块时间
const block = await provider.getBlock('latest');
const blockTimestamp = block.timestamp;
```

---

### 3. 测试模式控制

#### 3.1 启用测试模式

**UI 设计：**
```
┌───────────────────────────────┐
│   启用测试模式                 │
├───────────────────────────────┤
│  □ 使用当前时间               │
│  □ 自定义起始时间             │
│     [日期时间选择器]           │
│                               │
│  [  启用测试模式  ]            │
└───────────────────────────────┘
```

**合约函数：**
```javascript
// 启用测试模式（使用当前时间）
await stakingContract.enableTestMode(0);

// 启用测试模式（自定义时间）
const timestamp = Math.floor(new Date('2024-10-20').getTime() / 1000);
await stakingContract.enableTestMode(timestamp);
```

#### 3.2 禁用测试模式

**UI 设计：**
```
┌───────────────────────────────┐
│  ⚠️  禁用测试模式             │
├───────────────────────────────┤
│  禁用后将恢复使用真实时间      │
│  所有测试时间设置将被清除      │
│                               │
│  [  确认禁用  ]  [  取消  ]   │
└───────────────────────────────┘
```

**合约函数：**
```javascript
await stakingContract.disableTestMode();
```

---

### 4. 时间快进功能

**UI 设计：**
```
┌────────────────────────────────────┐
│   时间快进                          │
├────────────────────────────────────┤
│  快捷操作：                         │
│  [  +1 分钟  ]  [  +30 分钟  ]    │
│  [  +60 分钟 ]  [  +90 分钟  ]    │
│  [  +180 分钟]                     │
│                                    │
│  自定义快进：                       │
│  ┌─────────┬──────────┐           │
│  │  [_60_] │ 分钟 ▼   │           │
│  └─────────┴──────────┘           │
│  单位选项：分钟 / 小时 / 天         │
│                                    │
│  [  确认快进  ]                     │
└────────────────────────────────────┘
```

**合约函数：**
```javascript
// 快进分钟
await stakingContract.fastForwardMinutes(minutes);

// 快进天数
await stakingContract.fastForwardDays(days);

// 快进秒数
await stakingContract.fastForwardTime(seconds);
```

---

### 5. 测试场景预设

**功能描述：**
提供预设的测试场景，一键执行常见测试流程

**UI 设计：**
```
┌────────────────────────────────────────┐
│   测试场景                              │
├────────────────────────────────────────┤
│  场景 1: SSS 级 NFT 衰减测试           │
│  说明：测试 180 天完整衰减周期          │
│  操作：启用测试模式 → 快进180分钟      │
│  [  开始测试  ]                         │
│                                        │
│  场景 2: 持续质押奖励测试 (30天)       │
│  说明：测试 30 天持续质押 10% 奖励     │
│  操作：启用测试模式 → 快进30分钟       │
│  [  开始测试  ]                         │
│                                        │
│  场景 3: 持续质押奖励测试 (90天)       │
│  说明：测试 90 天持续质押 20% 奖励     │
│  操作：启用测试模式 → 快进90分钟       │
│  [  开始测试  ]                         │
│                                        │
│  场景 4: 组合奖励次日生效测试          │
│  说明：测试 3 NFT 组合次日生效机制     │
│  操作：启用测试模式 → 快进1分钟（多次） │
│  [  开始测试  ]                         │
└────────────────────────────────────────┘
```

**实现逻辑：**
```javascript
// 场景执行示例
async function runScenario1() {
  // 1. 检查测试模式
  const testMode = await stakingContract.testMode();
  if (!testMode) {
    await stakingContract.enableTestMode(0);
    await waitForConfirmation();
  }
  
  // 2. 快进时间
  await stakingContract.fastForwardMinutes(180);
  await waitForConfirmation();
  
  // 3. 显示完成提示
  showSuccess('场景测试完成！现在可以查看 NFT 奖励。');
}
```

---

### 6. 操作历史记录

**功能描述：**
显示所有操作记录，方便追踪测试流程

**UI 设计：**
```
┌────────────────────────────────────────┐
│   操作历史                              │
├────────────────────────────────────────┤
│  ✅ 2024-10-20 10:35:20               │
│     启用测试模式                        │
│     交易: 0x1234...5678                │
│                                        │
│  ✅ 2024-10-20 10:36:45               │
│     快进 60 分钟                       │
│     交易: 0xabcd...ef12                │
│                                        │
│  ✅ 2024-10-20 10:38:12               │
│     快进 60 分钟                       │
│     交易: 0x9876...5432                │
│                                        │
│  [  清空历史  ]  [  导出日志  ]        │
└────────────────────────────────────────┘
```

**数据结构：**
```javascript
interface OperationRecord {
  timestamp: number;
  action: string;
  txHash: string;
  status: 'success' | 'pending' | 'failed';
  details?: string;
}
```

---

### 7. 权限检查

**功能描述：**
自动检查当前钱包是否有权限操作测试模式

**检查内容：**
```javascript
// 检查是否为 owner
const owner = await stakingContract.owner();
const currentAccount = accounts[0];
const isOwner = owner.toLowerCase() === currentAccount.toLowerCase();

if (!isOwner) {
  // 显示警告
  alert(`
    ⚠️ 权限不足
    
    当前账户：${currentAccount}
    合约 Owner：${owner}
    
    只有合约 owner 可以控制测试模式。
    请切换到正确的钱包账户。
  `);
}
```

---

## 🎨 UI/UX 设计要求

### 整体布局

```
┌──────────────────────────────────────────────┐
│  🕐 Staking 时间控制测试工具                   │
│                           [连接钱包] [0x12..34]│
├──────────────────────────────────────────────┤
│  ┌─────────────────┐  ┌───────────────────┐ │
│  │   时间状态       │  │   快速操作         │ │
│  │  测试模式: ✅   │  │  [启用测试模式]    │ │
│  │  测试时间: ...  │  │  [禁用测试模式]    │ │
│  │  真实时间: ...  │  │  [快进时间]       │ │
│  │  时间差: ...    │  │                   │ │
│  └─────────────────┘  └───────────────────┘ │
│                                              │
│  ┌──────────────────────────────────────┐   │
│  │   时间快进                            │   │
│  │  [+1分] [+30分] [+60分] [+90分]      │   │
│  │  自定义: [___] [分钟▼] [确认]        │   │
│  └──────────────────────────────────────┘   │
│                                              │
│  ┌──────────────────────────────────────┐   │
│  │   测试场景                            │   │
│  │  □ SSS级衰减测试 (180分钟)           │   │
│  │  □ 持续质押30天测试                  │   │
│  │  □ 持续质押90天测试                  │   │
│  │  □ 组合奖励测试                      │   │
│  └──────────────────────────────────────┘   │
│                                              │
│  ┌──────────────────────────────────────┐   │
│  │   操作历史                            │   │
│  │  ✅ 10:35:20 启用测试模式            │   │
│  │  ✅ 10:36:45 快进 60 分钟            │   │
│  │  ...                                 │   │
│  └──────────────────────────────────────┘   │
└──────────────────────────────────────────────┘
```

### 交互反馈

1. **Loading 状态**
   - 显示交易确认等待提示
   - 显示交易哈希链接（Etherscan）
   - 预估完成时间

2. **成功提示**
   ```
   ✅ 操作成功！
   
   快进 60 分钟完成
   交易哈希: 0x1234...5678
   Gas 使用: 45,123
   
   [查看交易] [继续测试]
   ```

3. **错误处理**
   ```
   ❌ 操作失败
   
   错误原因: 权限不足
   
   建议:
   1. 确认你是合约 owner
   2. 检查钱包余额是否足够
   3. 查看详细错误信息
   
   [重试] [查看帮助]
   ```

---

## 🔧 技术实现要求

### 技术栈
- **前端框架**: React / Vue / Next.js（任选）
- **Web3 库**: ethers.js v6 或 viem
- **钱包连接**: wagmi / RainbowKit / ConnectKit
- **UI 组件库**: Ant Design / Material-UI / Tailwind CSS
- **状态管理**: React Context / Zustand / Recoil

### 合约交互

**Staking 合约 ABI（需要的函数）：**
```javascript
const stakingABI = [
  "function testMode() view returns (bool)",
  "function testTimestamp() view returns (uint256)",
  "function owner() view returns (address)",
  "function enableTestMode(uint256 initialTimestamp) external",
  "function disableTestMode() external",
  "function fastForwardTime(uint256 seconds_) external",
  "function fastForwardMinutes(uint256 minutes_) external",
  "function fastForwardDays(uint256 days_) external",
  "function setTestTimestamp(uint256 timestamp) external"
];

const stakingAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
const chainId = 11155111; // Sepolia
```

**初始化合约：**
```javascript
import { ethers } from 'ethers';

const provider = new ethers.BrowserProvider(window.ethereum);
const signer = await provider.getSigner();
const stakingContract = new ethers.Contract(
  stakingAddress,
  stakingABI,
  signer
);
```

### 交易确认流程

```javascript
async function executeTransaction(txPromise, actionName) {
  try {
    // 1. 发送交易
    showLoading(`正在发送交易: ${actionName}...`);
    const tx = await txPromise;
    
    // 2. 等待确认
    showLoading(`等待交易确认... (${tx.hash})`);
    const receipt = await tx.wait();
    
    // 3. 显示成功
    showSuccess({
      action: actionName,
      txHash: receipt.hash,
      gasUsed: receipt.gasUsed.toString(),
    });
    
    // 4. 刷新状态
    await refreshStatus();
    
    // 5. 记录历史
    addToHistory({
      timestamp: Date.now(),
      action: actionName,
      txHash: receipt.hash,
      status: 'success',
    });
    
  } catch (error) {
    // 处理错误
    handleError(error, actionName);
  }
}
```

### 错误处理

```javascript
function handleError(error, actionName) {
  let message = '未知错误';
  let suggestion = '请重试或联系技术支持';
  
  if (error.code === 'ACTION_REJECTED') {
    message = '用户取消了交易';
    suggestion = '如需继续，请重新操作';
  } else if (error.code === 'INSUFFICIENT_FUNDS') {
    message = 'ETH 余额不足';
    suggestion = '请从水龙头获取测试 ETH';
  } else if (error.message.includes('Not in test mode')) {
    message = '测试模式未启用';
    suggestion = '请先启用测试模式';
  } else if (error.message.includes('Not authorized')) {
    message = '权限不足';
    suggestion = '只有合约 owner 可以操作';
  }
  
  showError({
    action: actionName,
    message,
    suggestion,
    error: error.message,
  });
}
```

---

## 📱 响应式设计

### 桌面端（≥1024px）
- 三栏布局：左侧状态、中间操作、右侧历史
- 宽度固定，居中显示
- 最大宽度：1200px

### 平板端（768px - 1023px）
- 两栏布局：上方状态和操作、下方历史
- 自适应宽度

### 移动端（<768px）
- 单栏布局，纵向排列
- 操作按钮全宽
- 历史记录可折叠

---

## 🔐 安全考虑

### 1. 网络检查
```javascript
// 始终确认在 Sepolia 网络
const network = await provider.getNetwork();
if (network.chainId !== 11155111n) {
  throw new Error('请切换到 Sepolia 测试网');
}
```

### 2. 权限验证
```javascript
// 每次操作前验证权限
const owner = await stakingContract.owner();
const signer = await provider.getSigner();
const signerAddress = await signer.getAddress();

if (owner.toLowerCase() !== signerAddress.toLowerCase()) {
  throw new Error('只有合约 owner 可以执行此操作');
}
```

### 3. 交易确认
```javascript
// 重要操作二次确认
const confirmed = await confirm(`
  确认要执行以下操作吗？
  
  操作: ${actionName}
  预估 Gas: ${estimatedGas}
  
  点击确定继续
`);

if (!confirmed) {
  return;
}
```

---

## 📊 数据持久化

### LocalStorage 结构
```javascript
{
  "operationHistory": [
    {
      "timestamp": 1697800000000,
      "action": "启用测试模式",
      "txHash": "0x1234...",
      "status": "success"
    }
  ],
  "lastConnectedWallet": "0x1234...5678",
  "userPreferences": {
    "autoRefresh": true,
    "refreshInterval": 30000
  }
}
```

---

## 🧪 测试要求

### 功能测试
- [ ] 钱包连接/断开
- [ ] 网络切换检测
- [ ] 启用/禁用测试模式
- [ ] 时间快进功能
- [ ] 预设场景执行
- [ ] 权限检查
- [ ] 错误处理

### 兼容性测试
- [ ] Chrome (最新版)
- [ ] Firefox (最新版)
- [ ] Safari (最新版)
- [ ] Edge (最新版)
- [ ] MetaMask (最新版)

---

## 📦 交付物

### 1. 代码仓库
- 完整源代码
- README 说明文档
- 环境配置示例

### 2. 部署版本
- Vercel / Netlify 部署链接
- 支持直接访问使用

### 3. 使用文档
- 用户操作手册
- 常见问题解答
- 视频演示（可选）

---

## 🎯 开发优先级

### P0 (必须)
1. 钱包连接
2. 时间状态显示
3. 启用/禁用测试模式
4. 基本时间快进功能

### P1 (重要)
5. 快捷时间按钮
6. 权限检查
7. 交易状态显示
8. 基础错误处理

### P2 (可选)
9. 测试场景预设
10. 操作历史记录
11. 数据导出功能
12. 自定义主题

---

## 💡 参考资源

### 合约信息
- Staking Proxy: `0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5`
- Network: Sepolia (Chain ID: 11155111)
- RPC URL: 配置在 `.env.sepolia`

### 相关文档
- `docs/Staking-Test-Mode-Guide.md` - 功能说明
- `TIME_CONTROL_SUMMARY.md` - 功能总结
- `contracts/CPNFT/Staking.sol` - 合约源码

### Web3 资源
- ethers.js: https://docs.ethers.org/
- wagmi: https://wagmi.sh/
- RainbowKit: https://www.rainbowkit.com/
- Sepolia Faucet: https://sepoliafaucet.com/

---

## ❓ FAQ

### Q1: 为什么选择 Sepolia 而不是其他测试网？
A: Sepolia 是目前推荐的长期以太坊测试网，稳定且有充足的水龙头资源。

### Q2: 需要支持其他钱包吗？
A: 初期只需支持 MetaMask，后续可扩展支持 WalletConnect。

### Q3: 测试工具需要后端吗？
A: 不需要。所有操作直接与区块链交互，纯前端即可。

### Q4: 如何处理 Gas 费用？
A: 这是测试网，Gas 费用使用测试 ETH，从水龙头免费获取。

### Q5: 工具需要认证吗？
A: 不需要。只要是合约 owner 的钱包地址即可操作。

---

## 📞 联系方式

如有技术问题，请联系：
- 技术负责人：[填写联系方式]
- GitHub Issues：[仓库地址]
- 文档反馈：[反馈渠道]

---

**文档版本：** v1.0
**最后更新：** 2024-10-20
**状态：** ✅ Ready for Development

