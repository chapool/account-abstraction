# The Graph 集成指南 - NFT 质押统计可视化

## 概述

本指南将带你完整实现一个基于 The Graph 的数据索引和可视化系统，为 NFT 质押合约 (Staking.sol) 提供实时数据查询和图表展示功能。

### 实现目标

- ✅ 实时索引链上质押数据
- ✅ 提供高性能的 GraphQL 查询接口
- ✅ 展示用户质押数量的时间序列图
- ✅ 展示用户累计收益的时间趋势
- ✅ 支持多时间范围查看（24小时、7天、30天、90天）

### 技术栈

**后端（Subgraph）：**
- The Graph Protocol - 去中心化索引服务
- AssemblyScript - 事件处理逻辑
- GraphQL - 数据查询语言

**前端（Dashboard）：**
- React + TypeScript
- Apollo Client - GraphQL 客户端
- Recharts - 图表库
- date-fns - 时间处理

### 架构图

```
┌─────────────┐         ┌──────────────┐         ┌────────────┐
│   Sepolia   │ events  │   Subgraph   │  query  │  Frontend  │
│  Blockchain │────────>│  (Indexer)   │<────────│ Dashboard  │
└─────────────┘         └──────────────┘         └────────────┘
     │                         │                        │
     │ Staking.sol            │ GraphQL API           │ React UI
     │ Events:                │ Entities:             │ Components:
     │ - NFTStaked            │ - User                │ - StakingChart
     │ - NFTUnstaked          │ - HourlyStats         │ - RewardsChart
     │ - RewardsClaimed       │ - DailyStats          │ - Dashboard
     │ - BatchStaked          │ - Activity            │
     │ - BatchUnstaked        │ - GlobalStats         │
```

---

## Subgraph 开发

### 第一步：初始化项目

#### 1. 环境准备

确保已安装 Node.js (v16+) 和 npm：

```bash
node --version  # 应该 >= v16
npm --version
```

#### 2. 安装 Graph CLI

```bash
# 全局安装 Graph CLI
npm install -g @graphprotocol/graph-cli

# 验证安装
graph --version
```

#### 3. 创建 The Graph Studio 账号

1. 访问 https://thegraph.com/studio/
2. 使用钱包（MetaMask）连接并登录
3. 点击 "Create a Subgraph"
4. 输入名称：`nft-staking-stats`（或你喜欢的名称）
5. 选择网络：`Sepolia`
6. 点击创建

#### 4. 初始化本地项目

```bash
# 创建项目目录
mkdir nft-staking-subgraph
cd nft-staking-subgraph

# 初始化 Subgraph
graph init \
  --studio \
  --from-contract 0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5 \
  --network sepolia \
  --contract-name Staking \
  nft-staking-stats

# ⚠️ 参数说明：
# --studio: 使用 The Graph Studio（推荐）
# --from-contract: 合约地址（你的 Staking 合约）
# --network: 网络名称
# --contract-name: 合约名称
# nft-staking-stats: Subgraph 名称（与 Studio 创建的一致）

# 按提示操作：
✔ Protocol · ethereum
✔ Subgraph slug · nft-staking-stats
✔ Directory to create the subgraph in · nft-staking-stats
✔ Ethereum network · sepolia
✔ Contract address · 0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5
✔ Fetching ABI from Etherscan
✔ Contract Name · Staking
```

**💡 如果 Etherscan 获取 ABI 失败：**

```bash
# 手动初始化
graph init --studio nft-staking-stats

# 然后手动复制合约 ABI 到 abis/ 目录
```

#### 5. 项目结构

初始化后的目录结构：

```
nft-staking-subgraph/
├── abis/
│   └── Staking.json          # 合约 ABI（自动获取或手动复制）
├── src/
│   └── mapping.ts            # 事件处理逻辑（我们将编写）
├── schema.graphql            # 数据模型定义（我们将编写）
├── subgraph.yaml             # 子图配置
├── package.json              # 依赖配置
└── tsconfig.json             # TypeScript 配置
```

**关键文件说明：**

- `schema.graphql`: 定义数据结构（实体、字段、关系）
- `subgraph.yaml`: 配置数据源、事件监听、映射关系
- `src/mapping.ts`: 编写事件处理逻辑（AssemblyScript）
- `abis/Staking.json`: 合约 ABI，用于类型生成

---

### 第二步：定义数据模型

编辑 `schema.graphql` 文件，定义我们需要的数据结构：

```graphql
# schema.graphql

"""
用户实体 - 存储用户的总体统计信息
"""
type User @entity {
  id: ID!                                    # 用户地址（小写）
  address: Bytes!                            # 用户地址
  totalStaked: BigInt!                       # 当前质押总数
  totalStakedAllTime: BigInt!                # 历史总质押数
  totalUnstaked: BigInt!                     # 历史解质押总数
  totalRewardsClaimed: BigInt!               # 历史总收益（wei）
  totalRewardsClaimedDecimal: BigDecimal!    # 历史总收益（格式化为 ETH）
  firstStakeTimestamp: BigInt!               # 首次质押时间
  lastActivityTimestamp: BigInt!             # 最后活动时间
  
  # 关联关系
  hourlyStats: [UserHourlyStats!]! @derivedFrom(field: "user")
  dailyStats: [UserDailyStats!]! @derivedFrom(field: "user")
  activities: [StakingActivity!]! @derivedFrom(field: "user")
}

"""
用户小时统计 - 每小时的用户活动统计
用于绘制小时级别的时间趋势图
"""
type UserHourlyStats @entity {
  id: ID!                                    # {userAddress}-{hourTimestamp}
  user: User!                                # 关联用户
  hour: BigInt!                              # 小时时间戳（对齐到整点）
  hourStartString: String!                   # 小时字符串，例如 "2025-01-09 10:00"
  
  # 该小时的增量数据
  stakedCount: BigInt!                       # 该小时质押数量
  unstakedCount: BigInt!                     # 该小时解质押数量
  rewardsClaimed: BigInt!                    # 该小时领取的收益（wei）
  rewardsClaimedDecimal: BigDecimal!         # 该小时收益（格式化）
  
  # 该小时结束时的累计数据（用于绘制累计趋势图）
  cumulativeStaked: BigInt!                  # 截至该小时的累计质押数
  cumulativeUnstaked: BigInt!                # 截至该小时的累计解质押数
  cumulativeRewards: BigInt!                 # 截至该小时的累计收益
  netStaked: BigInt!                         # 净质押数（质押 - 解质押）
  
  # 各等级统计
  levelStats: [UserHourlyLevelStat!]! @derivedFrom(field: "hourlyStats")
}

"""
用户小时等级统计 - 按 NFT 等级细分的统计
"""
type UserHourlyLevelStat @entity {
  id: ID!                                    # {userHourlyStatsId}-{level}
  hourlyStats: UserHourlyStats!
  level: Int!                                # NFT等级 (1-6: C, B, A, S, SS, SSS)
  
  stakedCount: BigInt!                       # 该等级质押数
  unstakedCount: BigInt!                     # 该等级解质押数
  currentStaked: BigInt!                     # 该等级当前质押数
}

"""
用户日统计 - 每天的汇总数据
用于绘制日级别的时间趋势图（30天、90天范围）
"""
type UserDailyStats @entity {
  id: ID!                                    # {userAddress}-{dayTimestamp}
  user: User!
  day: BigInt!                               # 天时间戳（对齐到0点）
  dayString: String!                         # 日期字符串，例如 "2025-01-09"
  
  # 当天的增量数据
  stakedCount: BigInt!
  unstakedCount: BigInt!
  rewardsClaimed: BigInt!
  rewardsClaimedDecimal: BigDecimal!
  
  # 当天结束时的累计数据
  cumulativeStaked: BigInt!
  cumulativeUnstaked: BigInt!
  cumulativeRewards: BigInt!
  netStaked: BigInt!
}

"""
质押活动记录 - 每次操作的详细记录
用于审计和详细数据查询
"""
type StakingActivity @entity {
  id: ID!                                    # {txHash}-{logIndex}
  user: User!
  action: StakingAction!                     # 操作类型
  tokenIds: [BigInt!]!                       # 涉及的 NFT IDs
  level: Int                                 # NFT等级（单个质押时）
  amount: BigInt                             # 金额（解质押/领取时的收益）
  
  timestamp: BigInt!
  blockNumber: BigInt!
  transactionHash: Bytes!
}

enum StakingAction {
  STAKE
  UNSTAKE
  CLAIM
  BATCH_STAKE
  BATCH_UNSTAKE
  BATCH_CLAIM
}

"""
全局统计 - 平台总体数据
"""
type GlobalStats @entity {
  id: ID!                                    # 固定为 "global"
  totalUsers: BigInt!                        # 总用户数
  totalStaked: BigInt!                       # 总质押数
  totalRewardsPaid: BigInt!                  # 总发放收益
}
```

**💡 关键设计说明：**

- `UserHourlyStats.cumulativeRewards`: 累计收益，用于绘制递增的收益曲线
- `UserHourlyStats.netStaked`: 净质押数，用于绘制质押数量曲线
- `BigDecimal` 类型：用于前端直接显示，无需再转换 wei

### 第三步：配置子图

编辑 `subgraph.yaml`：

```yaml
# subgraph.yaml
specVersion: 0.0.5
schema:
  file: ./schema.graphql
dataSources:
  - kind: ethereum
    name: Staking
    network: sepolia
    source:
      address: "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5"
      abi: Staking
      startBlock: 5000000  # ⚠️ 替换为实际部署的区块号
    mapping:
      kind: ethereum/events
      apiVersion: 0.0.7
      language: wasm/assemblyscript
      entities:
        - User
        - UserHourlyStats
        - UserDailyStats
        - StakingActivity
        - GlobalStats
      abis:
        - name: Staking
          file: ./abis/Staking.json
      eventHandlers:
        - event: NFTStaked(indexed address,indexed uint256,uint8,uint256)
          handler: handleNFTStaked
        - event: NFTUnstaked(indexed address,indexed uint256,uint256,uint256)
          handler: handleNFTUnstaked
        - event: RewardsClaimed(indexed address,indexed uint256,uint256,uint256)
          handler: handleRewardsClaimed
        - event: BatchStaked(indexed address,uint256[],uint256)
          handler: handleBatchStaked
        - event: BatchUnstaked(indexed address,uint256[],uint256,uint256)
          handler: handleBatchUnstaked
      file: ./src/mapping.ts
```

**⚠️ 重要提示：**

- `startBlock`: 必须设置为合约部署的区块号，可以在 Etherscan 查询
- 如果设置太早，索引会很慢；如果设置太晚，会丢失历史数据

**如何查找部署区块号：**

1. 访问 https://sepolia.etherscan.io/
2. 搜索合约地址：`0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5`
3. 在合约页面找到 "Contract Creation" 信息
4. 记录 "Block" 数字

### 第四步：准备合约 ABI

如果初始化时没有自动获取 ABI，手动复制：

```bash
# 从你的主项目复制编译后的 ABI
cp ../artifacts/contracts/CPNFT/Staking.sol/Staking.json ./abis/

# 或者从 Etherscan 下载
# 访问合约页面 → Contract → Code → Export ABI
```

### 第五步：编写事件处理逻辑

编辑 `src/mapping.ts`，这是子图的核心逻辑：

> 由于代码较长，完整代码请参考附录 A，这里展示关键部分：

```typescript
// src/mapping.ts - 关键函数示例

import { BigInt, BigDecimal } from "@graphprotocol/graph-ts"
import { NFTStaked, NFTUnstaked } from "../generated/Staking/Staking"
import { User, UserHourlyStats } from "../generated/schema"

// 将时间戳对齐到小时
function getHourTimestamp(timestamp: BigInt): BigInt {
  return timestamp.div(BigInt.fromI32(3600)).times(BigInt.fromI32(3600))
}

// Wei 转换为 Decimal
function weiToDecimal(wei: BigInt): BigDecimal {
  let divisor = BigInt.fromI32(10).pow(18)
  return wei.toBigDecimal().div(divisor.toBigDecimal())
}

// 处理质押事件
export function handleNFTStaked(event: NFTStaked): void {
  // 1. 更新用户总体数据
  let user = getOrCreateUser(event.params.user, event.block.timestamp)
  user.totalStaked = user.totalStaked.plus(BigInt.fromI32(1))
  user.totalStakedAllTime = user.totalStakedAllTime.plus(BigInt.fromI32(1))
  user.save()
  
  // 2. 更新小时统计
  let hourlyStats = getOrCreateHourlyStats(
    event.params.user, 
    event.block.timestamp,
    user
  )
  hourlyStats.stakedCount = hourlyStats.stakedCount.plus(BigInt.fromI32(1))
  hourlyStats.cumulativeStaked = user.totalStakedAllTime  // 累计值
  hourlyStats.netStaked = user.totalStaked                // 当前值
  hourlyStats.save()
  
  // 3. 创建活动记录
  // ... 详见完整代码
}
```

**完整的 `mapping.ts` 代码请参考文档末尾的 [附录 A](#附录-a-完整-mappingts-代码)**

### 第六步：编译和部署

```bash
# 1. 安装依赖
npm install

# 2. 生成类型定义
graph codegen

# 如果遇到错误，检查：
# - schema.graphql 语法是否正确
# - subgraph.yaml 配置是否正确

# 3. 构建子图
graph build

# 如果构建失败，常见原因：
# - mapping.ts 语法错误
# - 引用了不存在的实体
# - 类型不匹配

# 4. 认证（使用 The Graph Studio 提供的 Deploy Key）
graph auth --studio <YOUR_DEPLOY_KEY>

# Deploy Key 在 Studio 的 Subgraph 页面可以找到

# 5. 部署到 Studio
graph deploy --studio nft-staking-stats

# 按提示选择版本标签
✔ Version Label (e.g. v0.0.1) · v0.0.1

# 成功输出示例：
✔ Deploy to Studio
✔ Version Label: v0.0.1
✔ Upload to IPFS

Deployed to https://thegraph.com/studio/subgraph/nft-staking-stats

Subgraph endpoints:
Queries (HTTP):     https://api.studio.thegraph.com/query/<id>/nft-staking-stats/v0.0.1
```

**部署后等待同步：**

1. 访问 Studio 页面查看同步进度
2. 初次部署需要索引所有历史区块，可能需要几分钟到几小时
3. 同步完成后，可以在 Playground 测试查询

### 第七步：测试子图

在 The Graph Studio 的 Playground 中测试：

```graphql
# 测试查询 - 获取用户基本信息
{
  user(id: "0x你的地址小写") {
    totalStaked
    totalRewardsClaimedDecimal
    firstStakeTimestamp
  }
}

# 测试查询 - 获取最近24小时数据
{
  user(id: "0x你的地址小写") {
    hourlyStats(
      first: 24
      orderBy: hour
      orderDirection: desc
    ) {
      hourStartString
      netStaked
      cumulativeRewards
      rewardsClaimedDecimal
    }
  }
}
```

---

## 前端应用开发

### 第一步：创建 React 项目

```bash
# 创建项目
npx create-react-app nft-staking-dashboard --template typescript
cd nft-staking-dashboard

# 安装依赖
npm install @apollo/client graphql
npm install recharts
npm install date-fns
```

### 第二步：配置 Apollo Client

创建 `src/apollo/client.ts`：

```typescript
// src/apollo/client.ts
import { ApolloClient, InMemoryCache } from '@apollo/client';

// ⚠️ 替换为你的 Subgraph URL（从 The Graph Studio 获取）
const SUBGRAPH_URL = 'https://api.studio.thegraph.com/query/<YOUR_ID>/nft-staking-stats/v0.0.1';

export const client = new ApolloClient({
  uri: SUBGRAPH_URL,
  cache: new InMemoryCache(),
});
```

### 第三步：定义 GraphQL 查询

创建 `src/apollo/queries.ts`：

```typescript
// src/apollo/queries.ts
import { gql } from '@apollo/client';

// 查询用户的时间序列数据（小时级别）
export const GET_USER_HOURLY_STATS = gql`
  query GetUserHourlyStats(
    $userAddress: ID!
    $startTime: BigInt!
    $endTime: BigInt!
  ) {
    user(id: $userAddress) {
      id
      totalStaked
      totalRewardsClaimedDecimal
      
      hourlyStats(
        where: { hour_gte: $startTime, hour_lt: $endTime }
        orderBy: hour
        orderDirection: asc
        first: 1000
      ) {
        hour
        hourStartString
        netStaked
        cumulativeRewards
        rewardsClaimedDecimal
      }
    }
  }
`;

// 查询用户的时间序列数据（日级别）
export const GET_USER_DAILY_STATS = gql`
  query GetUserDailyStats(
    $userAddress: ID!
    $startTime: BigInt!
    $endTime: BigInt!
  ) {
    user(id: $userAddress) {
      id
      dailyStats(
        where: { day_gte: $startTime, day_lt: $endTime }
        orderBy: day
        orderDirection: asc
        first: 365
      ) {
        day
        dayString
        netStaked
        cumulativeRewards
        rewardsClaimedDecimal
      }
    }
  }
`;
```

### 第四步到第九步：创建组件和页面

**完整的前端代码（包括 Hook、Chart 组件、Dashboard 页面、样式文件）请参考 [附录 B](#附录-b-完整前端代码)**

### 运行应用

```bash
npm start

# 应用将在 http://localhost:3000 启动
```

---

## 测试验证

### Subgraph 测试清单

- [ ] 子图成功部署到 Studio
- [ ] 同步进度达到 100%
- [ ] 能在 Playground 查询到数据
- [ ] 查询结果正确反映链上数据
- [ ] 累计值计算正确
- [ ] 时间戳对齐正确

### 前端测试清单

- [ ] 输入有效地址能加载数据
- [ ] 切换时间范围正确刷新
- [ ] 图表数值和标签正确
- [ ] 鼠标悬停显示详情
- [ ] 自动刷新功能正常
- [ ] 空数据友好提示
- [ ] 加载状态显示
- [ ] 错误处理正确

---

## 部署上线

### 前端部署（Vercel）

```bash
# 安装 Vercel CLI
npm install -g vercel

# 登录
vercel login

# 部署到预览环境
vercel

# 部署到生产环境
vercel --prod

# 访问部署的 URL
```

### 子图发布（可选）

```bash
# 从 Studio 发布到去中心化网络
# 需要支付一些 GRT 代币作为信号

# 在 Studio UI 操作：
# 1. 点击 "Publish"
# 2. 添加 GRT 信号
# 3. 确认交易
```

---

## 常见问题

### Q1: 子图同步很慢怎么办？

**A:** 
- 检查 `startBlock` 是否设置正确，不要从创世块开始
- Studio 的同步速度通常较快，耐心等待
- 可以在 Studio 查看同步进度和日志

### Q2: 查询不到数据？

**A:**
- 确认地址使用小写（The Graph 自动转换为小写）
- 检查时间范围是否正确
- 在 Playground 先测试查询
- 查看浏览器控制台的网络请求

### Q3: 数据延迟多久？

**A:**
- The Graph 通常有 1-5 分钟的延迟
- 取决于区块确认时间和同步速度
- 可以在前端显示"最后更新时间"

### Q4: 如何调试 Subgraph？

**A:**
```bash
# 使用本地 Graph Node（高级）
docker-compose up

# 查看构建日志
graph build --debug

# 检查映射逻辑
# 在 mapping.ts 中使用 log.info() 输出调试信息
```

### Q5: 累计值计算不对？

**A:**
- 检查 `getOrCreateHourlyStats` 是否正确设置初始累计值
- 确保每次更新都正确累加
- 在 Playground 查询验证数据

### Q6: 图表显示异常？

**A:**
- 检查数据格式是否正确（BigInt 需要转换）
- 查看浏览器控制台错误
- 验证 Recharts 配置

---

## 参考资源

### 官方文档

- [The Graph 官方文档](https://thegraph.com/docs/)
- [AssemblyScript 文档](https://www.assemblyscript.org/)
- [Apollo Client 文档](https://www.apollographql.com/docs/react/)
- [Recharts 文档](https://recharts.org/)

### 示例项目

- [Uniswap V2 Subgraph](https://github.com/Uniswap/v2-subgraph)
- [Aave Protocol Subgraph](https://github.com/aave/protocol-subgraphs)

### 社区支持

- [The Graph Discord](https://discord.gg/graphprotocol)
- [The Graph Forum](https://forum.thegraph.com/)
- [Stack Exchange - The Graph](https://ethereum.stackexchange.com/questions/tagged/the-graph)

---

## 附录 A: 完整 mapping.ts 代码

> 由于篇幅限制，完整代码请参考前文提供的 `src/mapping.ts` 实现

---

## 附录 B: 完整前端代码

> 完整前端代码请参考前文提供的组件和页面实现

---

## 附录 C: 项目检查清单

### 开发阶段

- [ ] 环境准备完成
- [ ] The Graph Studio 账号创建
- [ ] Subgraph 创建成功
- [ ] Schema 定义完成
- [ ] Mapping 逻辑实现
- [ ] 本地构建成功
- [ ] 部署到 Studio
- [ ] 同步完成
- [ ] Playground 测试通过
- [ ] 前端项目创建
- [ ] Apollo Client 配置
- [ ] GraphQL 查询定义
- [ ] 图表组件实现
- [ ] 本地测试通过

### 上线阶段

- [ ] 前端部署到 Vercel
- [ ] 域名配置（可选）
- [ ] 性能测试
- [ ] 用户验收测试
- [ ] 文档更新
- [ ] 监控配置

---

## 版本历史

| 版本 | 日期 | 说明 |
|------|------|------|
| v1.0 | 2025-01-09 | 初始版本 |

---

**文档维护**: 技术团队  
**最后更新**: 2025-01-09  
**下次审核**: 2025-02-09