# HzBay 跨链充值桥 — 完整技术方案

> **场景**：用户 USDT/USDC 在波场 (Tron)、Solana、Arbitrum 等链，DApp 部署在 opBNB，平台提供 opBNB USDT 流动性，用户转账到平台托管钱包后自动到账。

---

## 目录

1. [方案总览](#1-方案总览)
2. [整体架构](#2-整体架构)
3. [用户充值完整流程](#3-用户充值完整流程)
4. [地址生成策略](#4-地址生成策略)
5. [链上监听服务](#5-链上监听服务)
6. [风控与结算引擎](#6-风控与结算引擎)
7. [流动性管理](#7-流动性管理)
8. [数据库模型](#8-数据库模型)
9. [API 接口设计](#9-api-接口设计)
10. [前端交互设计](#10-前端交互设计)
11. [安全与风险控制](#11-安全与风险控制)
12. [部署架构](#12-部署架构)
13. [费率与成本模型](#13-费率与成本模型)
14. [开发里程碑](#14-开发里程碑)

---

## 1. 方案总览

### 核心思路

采用**中心化托管桥**（Custodial Deposit Bridge）模式：

- 平台在各源链持有**热钱包**作为用户专属收款地址
- 后端**监听服务**检测到账事件
- 风控通过后，**结算引擎**在 opBNB 向用户 EOA 钱包释放等值 USDT
- 平台预存 opBNB USDT 作为流动性池

### 支持的链与代币

| 源链 | 支持代币 | 确认块数 | 预计到账时间 |
|------|---------|---------|------------|
| Tron (TRC-20) | USDT、USDC | 20 块 | ~1 分钟 |
| Solana | USDT、USDC | 32 块 | ~20 秒 |
| Arbitrum | USDT、USDC | 12 块 | ~3 分钟 |
| BNB Chain | USDT、USDC | 15 块 | ~45 秒 |
| Ethereum | USDT、USDC | 12 块 | ~3 分钟 |

### 目标链

| 目标链 | 释放代币 | 来源 |
|--------|---------|------|
| opBNB | USDT (opBNB) | 平台流动性池 |

---

## 2. 整体架构

```mermaid
graph TB
    subgraph 用户侧
        U[用户 EOA 钱包]
        DAPP[HzBay DApp]
    end

    subgraph 充值桥前端 Bridge Frontend
        WEB[充值页面]
        QR[二维码 / 地址展示]
        STATUS[订单状态页]
    end

    subgraph 后端服务 Bridge Backend
        API[REST API 服务]
        ADDR[地址生成服务<br/>Address Manager]
        MON[多链监听服务<br/>Chain Monitor]
        SETTLE[结算引擎<br/>Settlement Engine]
        RISK[风控服务<br/>Risk Engine]
        NOTIFY[通知服务<br/>WebSocket / Webhook]
        DB[(PostgreSQL<br/>订单数据库)]
        MQ[消息队列<br/>Redis Streams]
        CACHE[Redis 缓存]
    end

    subgraph 平台热钱包 Hot Wallets
        TRON_W[Tron 热钱包池]
        SOL_W[Solana 热钱包池]
        ARB_W[ARB 热钱包池]
        BNB_W[opBNB 出款钱包]
    end

    subgraph 链上
        TRON[Tron USDT/USDC]
        SOL[Solana USDT/USDC]
        ARB[ARB USDT/USDC]
        OPBNB[opBNB USDT]
        POOL[平台流动性池<br/>opBNB USDT]
    end

    U -->|发起充值| DAPP
    DAPP --> WEB
    WEB -->|申请报价| API
    API --> ADDR
    ADDR --> DB
    API --> QR
    QR -->|展示| U

    U -->|转账到收款地址| TRON
    U -->|转账到收款地址| SOL
    U -->|转账到收款地址| ARB

    TRON --> TRON_W
    SOL --> SOL_W
    ARB --> ARB_W

    TRON_W -->|监听事件| MON
    SOL_W -->|监听事件| MON
    ARB_W -->|监听事件| MON

    MON -->|推送到账事件| MQ
    MQ --> RISK
    RISK -->|通过| SETTLE
    RISK -->|拒绝| DB
    SETTLE --> DB
    SETTLE -->|发送 USDT| BNB_W
    BNB_W --> POOL
    POOL -->|释放给用户| OPBNB
    OPBNB --> U

    SETTLE --> NOTIFY
    NOTIFY -->|实时推送| STATUS

    style POOL fill:#f0a500,color:#000
    style SETTLE fill:#1a73e8,color:#fff
    style RISK fill:#e53935,color:#fff
    style MON fill:#2e7d32,color:#fff
```

---

## 3. 用户充值完整流程

```mermaid
sequenceDiagram
    actor 用户
    participant DApp
    participant API as Bridge API
    participant DB as 订单数据库
    participant Monitor as 链上监听
    participant Risk as 风控引擎
    participant Settler as 结算引擎
    participant opBNB as opBNB链

    用户->>DApp: 1. 选择源链 + 代币 + 金额<br/>输入 opBNB 钱包地址
    DApp->>API: POST /v1/bridge/quote<br/>{fromChain, token, amount, toAddress}
    API->>API: 计算汇率 + 手续费 + 分配收款地址
    API->>DB: 创建 PENDING 订单
    API-->>DApp: {orderId, depositAddress,<br/>estimatedOut, fee, expireAt}
    DApp-->>用户: 展示收款地址 + 二维码<br/>倒计时15分钟

    用户->>用户: 2. 打开原链钱包
    用户->>Monitor: 向 depositAddress 转账 USDT

    Monitor->>Monitor: 3. 监听到入账交易
    Monitor->>Monitor: 等待确认块数达标
    Monitor->>DB: 更新订单状态 → DETECTED
    Monitor->>Risk: 推送到账事件

    Risk->>Risk: 4. 风控检查<br/>- AML/黑名单<br/>- 金额范围<br/>- 频率限制<br/>- 地址白名单
    alt 风控通过
        Risk->>Settler: 派发结算任务
        Risk->>DB: 更新状态 → RISK_PASSED
    else 风控拒绝
        Risk->>DB: 更新状态 → REJECTED
        Risk-->>用户: 通知人工审核
    end

    Settler->>Settler: 5. 查询流动性池余额
    alt 余额充足
        Settler->>opBNB: 发送 USDT 到用户 toAddress
        opBNB-->>Settler: 交易 Hash
        Settler->>DB: 更新状态 → COMPLETED<br/>记录 txHash
        Settler-->>用户: WebSocket 推送到账通知
    else 余额不足
        Settler->>DB: 更新状态 → PENDING_LIQUIDITY
        Settler-->>用户: 通知延迟到账（人工补充流动性）
    end

    用户->>DApp: 6. 查看 opBNB 余额，开始使用 DApp
```

---

## 4. 地址生成策略

### 4.1 地址分配模型

```mermaid
graph LR
    subgraph 地址池管理
        POOL1[Tron 地址池<br/>100 个地址]
        POOL2[Solana 地址池<br/>50 个地址]
        POOL3[ARB 地址池<br/>50 个地址]
    end

    subgraph 分配策略
        CHECK{地址是否空闲?}
        ALLOC[分配专属地址<br/>绑定 15 分钟]
        REUSE[到期自动回收]
        NEW[从 HD 钱包派生新地址]
    end

    REQ[用户申请] --> CHECK
    CHECK -->|有空闲地址| ALLOC
    CHECK -->|无空闲地址| NEW
    NEW --> ALLOC
    ALLOC --> REUSE
```

### 4.2 HD 钱包派生路径

```
主助记词 (离线保存)
└── m/44'/195'/0'/0/  (Tron)
    ├── index 0  → 地址 T1xxx (长期热钱包)
    ├── index 1  → 地址 T2xxx
    └── ...
└── m/44'/501'/0'/0/  (Solana)
    ├── index 0  → 地址 Sol1xxx
    └── ...
└── m/44'/60'/0'/0/   (ARB/EVM)
    ├── index 0  → 地址 0x1xxx
    └── ...
```

### 4.3 地址生命周期

```mermaid
stateDiagram-v2
    [*] --> IDLE: 派生/初始化
    IDLE --> ALLOCATED: 用户申请，绑定15分钟
    ALLOCATED --> MONITORING: 用户转账，检测到入账
    ALLOCATED --> IDLE: 15分钟无入账，自动释放
    MONITORING --> IDLE: 结算完成，地址归还地址池
    MONITORING --> LOCKED: 风控标记异常
    LOCKED --> IDLE: 人工审核解锁
```

---

## 5. 链上监听服务

### 5.1 多链监听架构

```mermaid
graph TB
    subgraph 监听器集群
        TM[Tron Monitor<br/>TronGrid API / 自建节点]
        SM[Solana Monitor<br/>Helius WebSocket]
        AM[ARB Monitor<br/>Alchemy WebSocket]
        BM[BSC Monitor<br/>Alchemy WebSocket]
    end

    subgraph 事件处理
        DEDUP[去重过滤器<br/>Redis Set]
        CONFIRM[确认数检查器]
        PARSE[交易解析器<br/>提取 token/amount/from/to]
        MATCH[地址匹配器<br/>查询分配地址表]
    end

    subgraph 输出
        MQ[Redis Streams<br/>deposit-events]
    end

    TM -->|原始交易| DEDUP
    SM -->|原始交易| DEDUP
    AM -->|原始交易| DEDUP
    BM -->|原始交易| DEDUP

    DEDUP --> CONFIRM
    CONFIRM -->|已满足确认数| PARSE
    PARSE --> MATCH
    MATCH -->|匹配到活跃地址| MQ
```

### 5.2 各链监听方式

| 链 | 监听方式 | 备选 | 延迟 |
|----|---------|------|------|
| Tron | TronGrid HTTP 轮询 (3s) | 自建全节点 | ~3s |
| Solana | Helius WebSocket logs | QuickNode | <1s |
| Arbitrum | Alchemy `eth_subscribe` | Infura | <1s |
| BSC | Alchemy WebSocket | 自建节点 | <1s |

---

## 6. 风控与结算引擎

### 6.1 风控检查流程

```mermaid
flowchart TD
    IN[收到到账事件] --> A{金额范围检查<br/>最小$1 最大$50,000}
    A -->|不通过| REJ1[拒绝: 金额异常]
    A -->|通过| B{黑名单检查<br/>发款地址/目标地址}
    B -->|命中| REJ2[拒绝: 黑名单]
    B -->|未命中| C{AML 检查<br/>Chainalysis / Elliptic}
    C -->|高风险| MANUAL[人工审核队列]
    C -->|低风险| D{频率限制<br/>同地址24h内≤5笔}
    D -->|超限| REJ3[拒绝: 频率超限]
    D -->|未超限| E{重复交易检查<br/>txHash 唯一性}
    E -->|重复| REJ4[拒绝: 重复交易]
    E -->|唯一| PASS[风控通过 → 结算]

    style PASS fill:#2e7d32,color:#fff
    style MANUAL fill:#f57f17,color:#000
    style REJ1 fill:#c62828,color:#fff
    style REJ2 fill:#c62828,color:#fff
    style REJ3 fill:#c62828,color:#fff
    style REJ4 fill:#c62828,color:#fff
```

### 6.2 结算引擎状态机

```mermaid
stateDiagram-v2
    [*] --> PENDING: 用户申请，创建订单
    PENDING --> DETECTED: 监听到链上入账
    DETECTED --> CONFIRMING: 等待确认块
    CONFIRMING --> RISK_PASSED: 风控通过
    CONFIRMING --> REJECTED: 风控拒绝
    RISK_PASSED --> SETTLING: 开始发送 opBNB USDT
    SETTLING --> COMPLETED: 链上确认成功
    SETTLING --> SETTLE_FAILED: 发送失败（Gas不足等）
    SETTLE_FAILED --> SETTLING: 自动重试（最多3次）
    SETTLE_FAILED --> MANUAL_SETTLE: 超过重试次数，人工处理
    PENDING --> EXPIRED: 15分钟无入账
    COMPLETED --> [*]
    REJECTED --> [*]
    EXPIRED --> [*]
```

### 6.3 汇率计算

```
实际到账金额 = 用户转账金额 × 汇率 - 手续费

手续费 = max(最低手续费, 用户金额 × 费率)

例：
  转入 100 USDT (Tron)
  汇率：1:1（USDT → USDT）
  费率：0.3%，最低 $0.5
  手续费：max(0.5, 100 × 0.003) = max(0.5, 0.3) = $0.5
  实际到账：99.5 USDT (opBNB)
```

---

## 7. 流动性管理

### 7.1 流动性池架构

```mermaid
graph LR
    subgraph 入金热钱包 各链收款
        T[Tron 热钱包<br/>归集账户]
        S[Solana 热钱包<br/>归集账户]
        A[ARB 热钱包<br/>归集账户]
    end

    subgraph 流动性池 opBNB
        HOT[出款热钱包<br/>日常结算<br/>余额: $10k-$50k]
        WARM[温钱包<br/>自动补充<br/>余额: $50k-$200k]
        COLD[冷钱包 多签<br/>主储备<br/>离线保存]
    end

    subgraph 监控告警
        MON[余额监控<br/>每5分钟检查]
        ALERT[告警通知<br/>Telegram / PagerDuty]
    end

    T -->|归集结算收益| COLD
    S -->|归集结算收益| COLD
    A -->|归集结算收益| COLD

    COLD -->|手动补充| WARM
    WARM -->|自动补充<br/>余额低于$20k时| HOT
    HOT -->|结算出款| U[用户钱包]

    MON --> HOT
    MON --> WARM
    MON --> ALERT
```

### 7.2 流动性预警阈值

| 钱包层级 | 目标余额 | 补充触发阈值 | 操作方式 |
|---------|---------|------------|---------|
| 出款热钱包 | $30,000 | < $10,000 | 自动从温钱包补充 |
| 温钱包 | $100,000 | < $30,000 | 告警通知人工从冷钱包转入 |
| 冷钱包 | $500,000+ | < $200,000 | 人工审批充值 |

---

## 8. 数据库模型

```sql
-- 订单表
CREATE TABLE deposit_orders (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_no        VARCHAR(32) UNIQUE NOT NULL,     -- BRG20241201XXXXXX
    
    -- 用户信息
    user_address    VARCHAR(100) NOT NULL,            -- opBNB 目标地址
    
    -- 源链信息
    from_chain      VARCHAR(20) NOT NULL,             -- tron/solana/arb/bsc
    from_token      VARCHAR(20) NOT NULL,             -- USDT/USDC
    from_amount     DECIMAL(36,18) NOT NULL,          -- 用户转入金额
    deposit_address VARCHAR(100) NOT NULL,            -- 平台收款地址
    source_tx_hash  VARCHAR(100),                     -- 用户转账哈希
    source_from     VARCHAR(100),                     -- 用户来源地址
    
    -- 目标链信息
    to_chain        VARCHAR(20) DEFAULT 'opbnb',
    to_token        VARCHAR(20) DEFAULT 'USDT',
    to_amount       DECIMAL(36,18),                  -- 实际到账金额
    settle_tx_hash  VARCHAR(100),                    -- 结算交易哈希
    
    -- 费率
    exchange_rate   DECIMAL(18,8) DEFAULT 1.0,
    fee_rate        DECIMAL(8,6)  DEFAULT 0.003,
    fee_amount      DECIMAL(36,18),
    
    -- 状态
    status          VARCHAR(30) NOT NULL DEFAULT 'PENDING',
    risk_level      VARCHAR(10),                     -- LOW/MEDIUM/HIGH
    risk_reason     TEXT,
    
    -- 时间
    expires_at      TIMESTAMPTZ NOT NULL,
    detected_at     TIMESTAMPTZ,
    settled_at      TIMESTAMPTZ,
    created_at      TIMESTAMPTZ DEFAULT NOW(),
    updated_at      TIMESTAMPTZ DEFAULT NOW()
);

-- 地址池表
CREATE TABLE deposit_addresses (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    chain           VARCHAR(20) NOT NULL,
    address         VARCHAR(100) NOT NULL,
    hd_index        INTEGER NOT NULL,
    status          VARCHAR(20) DEFAULT 'IDLE',      -- IDLE/ALLOCATED/MONITORING/LOCKED
    allocated_to    UUID REFERENCES deposit_orders(id),
    allocated_at    TIMESTAMPTZ,
    expires_at      TIMESTAMPTZ,
    created_at      TIMESTAMPTZ DEFAULT NOW(),
    UNIQUE(chain, address)
);

-- 流动性池记录
CREATE TABLE liquidity_logs (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    wallet_layer    VARCHAR(20) NOT NULL,             -- HOT/WARM/COLD
    balance_before  DECIMAL(36,18),
    balance_after   DECIMAL(36,18),
    change_amount   DECIMAL(36,18),
    change_reason   VARCHAR(100),
    related_order   UUID REFERENCES deposit_orders(id),
    created_at      TIMESTAMPTZ DEFAULT NOW()
);

-- 风控记录
CREATE TABLE risk_checks (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id        UUID REFERENCES deposit_orders(id),
    check_type      VARCHAR(50),                     -- BLACKLIST/AML/FREQUENCY/AMOUNT
    result          VARCHAR(20),                     -- PASS/REJECT/MANUAL
    detail          JSONB,
    created_at      TIMESTAMPTZ DEFAULT NOW()
);
```

---

## 9. API 接口设计

### 9.1 获取报价并创建订单

```
POST /v1/bridge/quote
```

**请求**：
```json
{
  "from_chain": "tron",
  "from_token": "USDT",
  "amount": "100",
  "to_address": "0xUserOpBNBAddress"
}
```

**响应**：
```json
{
  "order_id": "BRG20241201ABCD1234",
  "deposit_address": "TXxxxxxxxxxxxxxx",
  "from_chain": "tron",
  "from_token": "USDT",
  "input_amount": "100",
  "fee_amount": "0.5",
  "estimated_output": "99.5",
  "exchange_rate": "1.0",
  "to_address": "0xUserOpBNBAddress",
  "expires_at": "2024-12-01T12:15:00Z",
  "min_confirmations": 20,
  "qr_code": "data:image/png;base64,..."
}
```

### 9.2 查询订单状态

```
GET /v1/bridge/order/:order_id
```

**响应**：
```json
{
  "order_id": "BRG20241201ABCD1234",
  "status": "COMPLETED",
  "source_tx_hash": "用户转账哈希",
  "settle_tx_hash": "opBNB结算哈希",
  "from_amount": "100",
  "to_amount": "99.5",
  "settled_at": "2024-12-01T12:03:45Z"
}
```

### 9.3 WebSocket 实时状态推送

```
WS /v1/bridge/ws/:order_id
```

**服务端推送事件**：
```json
{ "event": "DETECTED",   "message": "检测到入账，等待确认..." }
{ "event": "CONFIRMING", "message": "已确认 12/20 块..." }
{ "event": "SETTLING",   "message": "正在发送 opBNB USDT..." }
{ "event": "COMPLETED",  "message": "到账成功！", "tx_hash": "0x..." }
```

### 9.4 获取支持的链和汇率

```
GET /v1/bridge/chains
```

**响应**：
```json
{
  "chains": [
    {
      "chain_id": "tron",
      "name": "Tron",
      "tokens": ["USDT", "USDC"],
      "fee_rate": 0.003,
      "min_amount": 1,
      "max_amount": 50000,
      "estimated_time_minutes": 1,
      "min_confirmations": 20
    }
  ]
}
```

---

## 10. 前端交互设计

```mermaid
graph TD
    A[充值入口] --> B[选择源链<br/>Tron / SOL / ARB / BSC]
    B --> C[输入金额<br/>实时显示手续费和到账金额]
    C --> D[确认 opBNB 地址<br/>可自动读取已连接钱包]
    D --> E[获取收款地址<br/>显示 QR 码 + 复制按钮]
    E --> F{倒计时 15 分钟}
    F -->|用户完成转账| G[状态监控页<br/>实时进度条]
    F -->|超时| H[显示过期提示<br/>重新申请]
    G --> I{订单状态}
    I -->|CONFIRMING| J[已到账，确认中...<br/>显示进度 X/20]
    I -->|COMPLETED| K[到账成功 🎉<br/>显示 opBNB txHash]
    I -->|REJECTED| L[需人工审核<br/>显示客服联系方式]
```

---

## 11. 安全与风险控制

### 11.1 私钥安全

```
推荐方案（按安全级别）：

1. 最高安全（生产推荐）
   ├── AWS KMS / GCP Cloud HSM 托管主私钥
   ├── 应用层只持有派生路径，无法导出原始私钥
   └── 签名操作在 KMS 内执行，私钥永不离开 HSM

2. 高安全（中等规模）
   ├── 出款热钱包私钥加密存储于 HashiCorp Vault
   ├── 每次结算从 Vault 动态获取，TTL 30分钟
   └── 审计日志记录每次私钥访问

3. 多签（大额出款）
   ├── 单笔超过 $5,000 需要 2/3 多签确认
   └── 冷钱包补充使用 Gnosis Safe
```

### 11.2 运营风险矩阵

| 风险 | 概率 | 影响 | 缓解措施 |
|------|------|------|---------|
| 流动性不足 | 中 | 高 | 多层钱包架构 + 余额监控告警 |
| 私钥泄露 | 低 | 极高 | HSM + 多签 + 最小权限原则 |
| 链上拥堵结算延迟 | 中 | 中 | 动态 Gas 策略 + EIP-1559 |
| 监听服务宕机 | 低 | 高 | 多节点冗余 + 断点续扫 |
| 用户重复提交 | 高 | 中 | txHash 唯一性检查 |
| AML 合规风险 | 中 | 极高 | 接入 Chainalysis + 限额 + KYC |
| 汇率波动（USDC≠USDT） | 中 | 中 | 实时喂价 + 价格偏差保护 |

### 11.3 合规要求

- **KYC**：单笔超过 $10,000 需要身份验证
- **AML**：接入 Chainalysis KYT 或 Elliptic 进行地址风险评分
- **日志留存**：所有交易记录保存 5 年
- **限额**：
  - 单次：最小 $1，最大 $50,000
  - 日限额：单用户 $100,000
  - 月限额：单用户 $500,000

---

## 12. 部署架构

```mermaid
graph TB
    subgraph CDN / 负载均衡
        CF[Cloudflare CDN]
        LB[负载均衡器]
    end

    subgraph 应用集群 Kubernetes
        API1[API Server Pod 1]
        API2[API Server Pod 2]
        MON1[Monitor Service Pod 1<br/>Tron + SOL]
        MON2[Monitor Service Pod 2<br/>ARB + BSC]
        SETTLE1[Settlement Pod 1]
        SETTLE2[Settlement Pod 2]
        RISK1[Risk Service Pod]
    end

    subgraph 数据层
        PG_M[(PostgreSQL Master)]
        PG_R[(PostgreSQL Replica)]
        REDIS[(Redis Cluster)]
    end

    subgraph 外部服务
        ALCHEMY[Alchemy<br/>ARB + BSC 节点]
        HELIUS[Helius<br/>Solana 节点]
        TRONGRID[TronGrid<br/>Tron 节点]
        CHAINALYSIS[Chainalysis KYT<br/>AML 检查]
        KMS[AWS KMS<br/>私钥管理]
    end

    CF --> LB
    LB --> API1
    LB --> API2
    API1 --> PG_M
    API2 --> PG_M
    PG_M --> PG_R
    API1 --> REDIS
    MON1 --> REDIS
    MON2 --> REDIS
    SETTLE1 --> PG_M
    SETTLE1 --> KMS
    SETTLE2 --> KMS
    RISK1 --> CHAINALYSIS

    MON1 --> TRONGRID
    MON1 --> HELIUS
    MON2 --> ALCHEMY
```

### 基础设施配置

| 组件 | 规格 | 数量 | 月成本估算 |
|------|------|------|-----------|
| API Server | 2vCPU / 4GB | 2 | $80 |
| Monitor Service | 2vCPU / 4GB | 2 | $80 |
| Settlement Service | 2vCPU / 4GB | 2 | $80 |
| PostgreSQL | 4vCPU / 16GB | 1主1从 | $200 |
| Redis | 2vCPU / 8GB | 集群 | $100 |
| Alchemy / Helius | 按量付费 | - | $200 |
| Chainalysis KYT | 按量付费 | - | $500+ |
| **合计** | | | **~$1,240/月** |

---

## 13. 费率与成本模型

### 13.1 平台收费结构

```
手续费 = max(链路成本 + 利润, 最低费用)

各链成本估算（单笔）：
  Tron   → opBNB：$0.10（Tron 几乎免费 + opBNB Gas $0.001）
  Solana → opBNB：$0.01（SOL Gas）+ opBNB Gas $0.001
  ARB    → opBNB：$0.20（ARB Gas）+ opBNB Gas $0.001
  ETH    → opBNB：$3~10（ETH Gas 波动大）

建议费率：
  Tron：0.3% 或最低 $0.5
  Solana：0.3% 或最低 $0.5
  ARB：0.5% 或最低 $1.0
  ETH：0.8% 或最低 $5.0
```

### 13.2 收益模型示例

```
假设：日充值 $100,000（均值每笔 $500，200 笔/天）

收入：$100,000 × 0.3% = $300/天 = $9,000/月

成本：
  基础设施：$1,240/月
  Gas 费：$200/月（估算）
  AML 服务：$500/月
  运营：$500/月
  合计：$2,440/月

净利润：约 $6,560/月（以上为估算，实际以业务量为准）
```

---

## 14. 开发里程碑

```mermaid
gantt
    title 充值桥开发计划
    dateFormat  YYYY-MM-DD
    section 第一阶段 MVP
    需求确认与架构设计        :done,    des1, 2024-12-01, 7d
    数据库设计与 API 定义      :done,    des2, after des1, 5d
    地址生成服务开发          :active,  dev1, after des2, 7d
    Tron 监听服务开发         :         dev2, after des2, 7d
    基础结算引擎开发          :         dev3, after dev1, 7d
    前端充值页面开发          :         fe1,  after des2, 10d
    section 第二阶段 完善
    ARB / SOL 监听开发        :         dev4, after dev2, 7d
    风控服务接入              :         risk1, after dev3, 5d
    流动性监控告警            :         mon1, after dev3, 5d
    WebSocket 实时推送        :         ws1,  after dev3, 5d
    section 第三阶段 上线
    测试网联调测试            :         test1, after dev4, 7d
    安全审计                  :         audit, after test1, 7d
    主网灰度上线（Tron 先行）  :         prod1, after audit, 3d
    全链开放                  :         prod2, after prod1, 7d
```

### 各阶段交付物

| 阶段 | 时间 | 交付物 |
|------|------|--------|
| MVP | 第 1-3 周 | Tron → opBNB 充值通道，基础风控 |
| 完善 | 第 4-5 周 | 全链支持，流动性监控，AML 接入 |
| 上线 | 第 6-8 周 | 安全审计，灰度上线，全链开放 |

---

## 附录：技术选型建议

| 模块 | 推荐方案 | 备选方案 |
|------|---------|---------|
| 后端框架 | Go (go-starter) | Node.js + TypeScript |
| 消息队列 | Redis Streams | RabbitMQ |
| 私钥管理 | AWS KMS | HashiCorp Vault |
| Tron 节点 | TronGrid API | 自建 FullNode |
| Solana 节点 | Helius | QuickNode |
| ARB/EVM 节点 | Alchemy | Infura |
| AML 服务 | Chainalysis KYT | Elliptic |
| 监控告警 | Grafana + PagerDuty | Datadog |
| 部署 | Kubernetes (AWS EKS) | Docker Compose (初期) |

---

*最后更新：2024-12-01*  
*文档版本：v1.0*  
*维护团队：HzBay 基础设施组*
