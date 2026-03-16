# Chapool TVL 数据指南

> 本文面向运营团队，说明 Chapool 在 DefiLlama 上的 TVL 是如何统计的、数据来自哪里、以及 CPOT 代币上线 CoinGecko 的完整流程。

---

## 一、Chapool TVL 的组成

Chapool 在 DefiLlama 上的 TVL（Total Value Locked，锁仓总价值）由**两部分**构成：

### 1. 用户存入的 USDT（主要部分）

| 项目 | 说明 |
|------|------|
| 来源合约 | `ChapoolEarnVault` |
| 读取方式 | 调用 `ChapoolVaultReader.getTVL()` |
| 计价方式 | USDT 本身就是稳定币，恒定 = $1，无需额外定价 |
| 赎回机制 | 用户存多少 USDT 可取回多少，1:1，不产生 USDT 收益 |

用户将 USDT 存入 EarnVault 后，USDT 就由合约托管。DefiLlama 直接读取合约里的 USDT 总余额，乘以 $1，即为该部分 TVL。

---

### 2. 用户锁定的 CPOT（辅助部分）

| 项目 | 说明 |
|------|------|
| 来源合约 | `VeCPOTLocker` |
| 读取方式 | 读取 CPOT 代币在 VeCPOTLocker 合约里的余额（`balanceOf`）|
| 计价方式 | 依赖 CoinGecko / GeckoTerminal 提供的 CPOT 市场价格 |
| 锁定目的 | 获得 veCPOT Boost，使 CPP 奖励发放速率提升 +1% ~ +5% |

用户锁定 CPOT 的价值会折算成 USD 后加入 TVL。**如果 CPOT 还没有 CoinGecko 价格，这部分 TVL 显示为 $0，但不影响 USDT 部分的统计。**

---

## 二、TVL 计算公式

```
Chapool TVL = 
  (EarnVault 中的 USDT 总量  × $1.00)
+ (VeCPOTLocker 中的 CPOT 总量 × CPOT 市场价格)
```

**举例：**
- EarnVault 中有 500,000 USDT → 贡献 $500,000
- VeCPOTLocker 中有 1,000,000 CPOT，CPOT 价格 $0.05 → 贡献 $50,000
- **总 TVL = $550,000**

---

## 三、DefiLlama 如何读取数据

DefiLlama 每隔一段时间（通常 1 小时）自动调用我们部署在其平台的适配器代码（`projects/chapool/index.js`），流程如下：

```
DefiLlama 调度器
    │
    ▼
调用 ChapoolVaultReader.getTVL()    ──→  获得 USDT 总量
    │
    ▼
调用 CPOT.balanceOf(VeCPOTLocker)   ──→  获得锁定 CPOT 总量
    │
    ▼
查询 CoinGecko/GeckoTerminal 价格   ──→  USDT=$1, CPOT=市场价
    │
    ▼
汇总 → 更新 DefiLlama TVL 页面
```

> DefiLlama 适配器代码已准备好，需要在 opBNB 主网合约部署后填入合约地址，再提交 PR 到 DefiLlama 官方仓库即可上线。

---

## 四、CPOT 如何获得 CoinGecko 价格

CPOT 价格由 **CoinGecko + GeckoTerminal** 联动提供，核心前提是 **在 DEX 上建立真实流动性**。

### 第一步：在 DEX 上创建 CPOT/USDT 交易对

推荐使用 **PancakeSwap（opBNB 主网）**，这是 opBNB 上流动性最好的 DEX，且 GeckoTerminal 已支持其数据索引。

- 在 PancakeSwap 上创建 CPOT/USDT 流动性池
- 注入足够的初始流动性（建议 ≥ $10,000 等值，流动性越高价格越稳定、越容易通过审核）

### 第二步：确认 GeckoTerminal 已自动索引

访问 [geckoterminal.com](https://www.geckoterminal.com)，选择 **opBNB** 链，搜索 CPOT。如果已有流动性，GeckoTerminal 通常会在 **数小时内**自动发现并显示交易对和价格图表。

> GeckoTerminal 是 CoinGecko 旗下的 DEX 数据平台，两者价格数据共享。

### 第三步：向 CoinGecko 提交上币申请

申请入口：[https://www.coingecko.com/en/coins/new](https://www.coingecko.com/en/coins/new)

需要准备的材料：

| 材料 | 说明 |
|------|------|
| 代币名称 | CPOT（Chapool Point Token） |
| 合约地址 | opBNB 主网上的 CPOT 合约地址 |
| 官网 | chapool.io 或官方域名 |
| Logo | 256×256 PNG，透明背景 |
| 社交媒体 | Twitter / Telegram / Discord 链接 |
| 价格来源 | 选 "On-Chain DEX"，填入 GeckoTerminal 上 CPOT 交易对的 URL |
| 白皮书/项目说明 | 简要说明 CPOT 的用途（veCPOT Boost、治理等）|

### 第四步：等待审核

- CoinGecko 审核周期通常为 **1~4 周**
- 审核通过后，CoinGecko 开始实时从 GeckoTerminal（即 DEX）拉取 CPOT 价格
- DefiLlama 随即自动将 CPOT 锁定量换算成 USD 并加入 Chapool TVL

---

## 五、上线 DefiLlama 的完整 Checklist

| 步骤 | 状态 | 负责方 |
|------|------|------|
| EARN 合约部署到 opBNB 主网 | ⬜ 待完成 | 合约团队 |
| 在 `projects/chapool/index.js` 填入主网合约地址 | ⬜ 待完成 | 开发团队 |
| 向 DefiLlama GitHub 提交 PR | ⬜ 待完成 | 开发团队 |
| DefiLlama 合并 PR，TVL 上线 | ⬜ 待完成 | DefiLlama 审核 |
| PancakeSwap(opBNB) 建立 CPOT/USDT 流动性池 | ⬜ 待完成 | 运营团队 |
| GeckoTerminal 自动索引 CPOT 价格 | ⬜ 自动 | GeckoTerminal |
| 向 CoinGecko 提交 CPOT 上币申请 | ⬜ 待完成 | 运营团队 |
| CoinGecko 审核通过，CPOT 价格上线 | ⬜ 待完成 | CoinGecko 审核 |
| CPOT TVL 自动计入 DefiLlama | ⬜ 自动 | DefiLlama |

---

## 六、常见问题

**Q：CPOT 没有价格时，DefiLlama 会报错吗？**  
A：不会。CPOT 部分 TVL 显示为 $0，USDT 部分正常统计，不影响整体数据展示。

**Q：USDT 的价格需要单独提交吗？**  
A：不需要。USDT 是稳定币，DefiLlama 内置价格为 $1.00，无需任何操作。

**Q：TVL 数据多久更新一次？**  
A：DefiLlama 通常每 **1 小时**更新一次链上数据。

**Q：DefiLlama 审核 PR 需要多长时间？**  
A：通常 **3~7 个工作日**，需要在 GitHub 上的 DefiLlama-Adapters 仓库提交 Pull Request。

**Q：流动性池注入多少合适？**  
A：CoinGecko 没有硬性要求，但建议初始流动性 ≥ $10,000，否则价格波动过大，可能影响 TVL 数据质量，也降低审核通过率。
