# ChainBridge - 改进设计方案

基于Todo.md的要求和MasterAggregator项目的设计经验，提出以下改进方案。

## 📋 **核心修改要求**

基于`docs/chain-bridge/Todo.md`的要求：

1. **CPOP统一处理为ERC20**: 不再特殊处理CPOP，统一按ERC20标准处理
2. **移除UserOperation批处理**: 由bundler自动处理，专注于MasterAggregator功能
3. **应用层批处理功能**: 实现`batchMint`和`batchBurn`等高级功能

## 🎯 **改进方案概览**

### 1. **资产管理统一化改进**

**现状问题**: 文档中CPOP被特殊处理，有专门的积分管理逻辑
**改进方案**: 统一按ERC20处理，简化架构

```go
// 改进前：特殊处理CPOP
type AssetManager struct {
    CPOPManager    *CPOPBalanceManager  // ❌ 特殊处理
    AlchemyClient  *AlchemyAPIClient
    LocalERC20Manager *ERC20TokenManager
}

// 改进后：统一ERC20处理
type AssetManager struct {
    AlchemyClient     *AlchemyAPIClient
    ERC20Manager      *UnifiedERC20Manager  // ✅ 统一处理所有ERC20
    TokenRegistry     *TokenRegistryService // ✅ 代币注册表
    MultiChain        *MultiChainManager
}

// 统一的ERC20管理器
type UnifiedERC20Manager struct {
    contractService *ContractService
    batchProcessor  *BatchProcessor
    supportedTokens map[int64]map[string]*TokenConfig
}

type TokenConfig struct {
    Address     string `json:"address"`
    Symbol      string `json:"symbol"`
    Name        string `json:"name"`
    Decimals    int    `json:"decimals"`
    Type        string `json:"type"`        // "standard", "utility", "governance"
    Features    []string `json:"features"`  // ["mintable", "burnable", "pausable"]
    IsNative    bool   `json:"is_native"`   // ETH, BNB, MATIC等
}
```

### 2. **Web2友好的记账系统设计**

设计理念：
1. **Web2开发者视角**: 就像调用传统数据库一样简单
2. **记账系统**: 维护链下状态，自动批量同步到链上
3. **自动批处理**: 根据策略自动决定何时上链
4. **原因追踪**: 每次资产变动都要有明确的业务原因

#### 系统架构

```
Web2后端应用
    ↓ 调用简单的REST API
ChainBridge记账系统
    ↓ 链下记账 + 自动批量策略
    ↓ 自动批量上链
区块链网络 (通过EntryPoint + MasterAggregator)
```

## 🎯 **核心API设计**

### 1. **资产增减接口（核心接口）**

```http
POST /api/v1/balance/adjust
```

**请求体**:
```json
{
  "adjustments": [
    {
      "user_id": "user_123",
      "token_symbol": "CPOP",
      "amount": "+1000.0",
      "reason_type": "reward",
      "reason_detail": "Daily check-in reward",
      "metadata": {
        "activity_id": "daily_checkin_20241215",
        "reward_tier": "premium"
      }
    },
    {
      "user_id": "user_456", 
      "token_symbol": "CPOP",
      "amount": "-50.0",
      "reason_type": "gas_fee",
      "reason_detail": "Transaction gas fee deduction",
      "metadata": {
        "tx_id": "user_transfer_abc123",
        "gas_cost_usd": 2.5
      }
    }
  ],
  "batch_id": "daily_rewards_batch_001", // 可选，用于关联业务批次
  "priority": "normal" // low, normal, high
}
```

**响应**:
```json
{
  "operation_id": "op_550e8400-e29b-41d4-a716-446655440000",
  "status": "recorded",
  "adjustments_count": 2,
  "blockchain_sync": {
    "status": "pending",
    "estimated_sync_time": "5-15 minutes",
    "will_be_batched": true,
    "current_batch_size": 23
  },
  "balances": [
    {
      "user_id": "user_123",
      "token_symbol": "CPOP", 
      "balance_before": "5000.0",
      "balance_after": "6000.0",
      "pending_sync": "+1000.0"
    }
  ]
}
```

### 3. **查询接口**

#### 用户余额查询
```http
GET /api/v1/balance/{user_id}?token_symbol=CPOP
```

```json
{
  "user_id": "user_123",
  "balances": [
    {
      "token_symbol": "CPOP",
      "confirmed_balance": "5000.0",    // 已上链确认的余额
      "pending_balance": "5050.0",      // 包含待上链变动的余额
      "pending_changes": "+50.0",       // 待上链的变动
      "last_sync_time": "2024-12-15T10:30:00Z",
      "next_sync_estimate": "2024-12-15T10:45:00Z"
    }
  ]
}
```

#### 资产变动历史
```http
GET /api/v1/balance/{user_id}/history?token_symbol=CPOP&limit=20
```

```json
{
  "user_id": "user_123",
  "token_symbol": "CPOP",
  "history": [
    {
      "operation_id": "op_123",
      "timestamp": "2024-12-15T10:30:00Z",
      "amount": "+100.0",
      "balance_before": "5000.0",
      "balance_after": "5100.0",
      "reason_type": "reward",
      "reason_detail": "Daily check-in reward",
      "sync_status": "confirmed",
      "blockchain_tx": "0xabc123...",
      "metadata": {
        "activity_id": "daily_checkin_20241215"
      }
    }
  ]
}
```


## 🔧 **核心实现架构**

### 1. **记账系统服务**

```go
// Web2友好的记账系统
type AccountingService struct {
    balanceManager  *BalanceManager
    batchStrategy   *AutoBatchStrategy
    syncManager     *BlockchainSyncManager
    reasonRegistry  *ReasonRegistryService
    logger         *logger.Logger
}

type BalanceAdjustment struct {
    UserID       string                 `json:"user_id"`
    TokenSymbol  string                 `json:"token_symbol"`
    Amount       string                 `json:"amount"`        // "+100.0" 或 "-50.0"
    ReasonType   ReasonType             `json:"reason_type"`
    ReasonDetail string                 `json:"reason_detail"`
    Metadata     map[string]interface{} `json:"metadata"`
}

type ReasonType string
const (
    REASON_REWARD       ReasonType = "reward"        // 奖励发放
    REASON_GAS_FEE      ReasonType = "gas_fee"       // Gas费扣减
    REASON_CONSUMPTION  ReasonType = "consumption"   // 消费扣费
    REASON_REFUND       ReasonType = "refund"        // 退款
    REASON_TRANSFER_IN  ReasonType = "transfer_in"   // 转入
    REASON_TRANSFER_OUT ReasonType = "transfer_out"  // 转出
    REASON_ADJUSTMENT   ReasonType = "adjustment"    // 管理员调整
)

func (as *AccountingService) AdjustBalances(ctx context.Context, req *BalanceAdjustRequest) (*BalanceAdjustResponse, error) {
    operationID := generateOperationID()
    
    // 1. 验证调整请求
    if err := as.validateAdjustments(req.Adjustments); err != nil {
        return nil, fmt.Errorf("validation failed: %w", err)
    }
    
    // 2. 在数据库中记录链下状态变更
    balanceChanges := []BalanceChange{}
    for _, adj := range req.Adjustments {
        change, err := as.balanceManager.RecordBalanceChange(ctx, &BalanceChange{
            OperationID:  operationID,
            UserID:      adj.UserID,
            TokenSymbol: adj.TokenSymbol,
            Amount:      parseAmount(adj.Amount),
            ReasonType:  adj.ReasonType,
            ReasonDetail: adj.ReasonDetail,
            Metadata:    adj.Metadata,
            Status:      "recorded", // recorded -> pending_sync -> synced
        })
        if err != nil {
            return nil, fmt.Errorf("record balance change failed: %w", err)
        }
        balanceChanges = append(balanceChanges, *change)
    }
    
    // 3. 触发批量上链策略评估
    as.batchStrategy.EvaluateBatchTrigger(balanceChanges)
    
    // 4. 构建响应
    return &BalanceAdjustResponse{
        OperationID:      operationID,
        Status:          "recorded",
        AdjustmentsCount: len(req.Adjustments),
        BlockchainSync:   as.getBlockchainSyncInfo(),
        Balances:        as.buildBalanceInfo(balanceChanges),
    }, nil
}
```

### 2. **自动批量上链策略**

```go
// 自动批量上链策略
type AutoBatchStrategy struct {
    batchTriggers   []BatchTrigger
    syncManager     *BlockchainSyncManager
    logger         *logger.Logger
}

type BatchTrigger struct {
    Name           string        `yaml:"name"`
    TokenSymbol    string        `yaml:"token_symbol"`
    MaxBatchSize   int           `yaml:"max_batch_size"`   // 50个操作
    MaxWaitTime    time.Duration `yaml:"max_wait_time"`    // 15分钟
    MinBatchSize   int           `yaml:"min_batch_size"`   // 5个操作
    ValueThreshold string        `yaml:"value_threshold"`  // 总价值阈值
    Priority       []ReasonType  `yaml:"priority"`         // 高优先级原因
}

func (abs *AutoBatchStrategy) EvaluateBatchTrigger(changes []BalanceChange) {
    // 按代币分组待同步的变更
    pendingByToken := abs.groupPendingChangesByToken()
    
    for tokenSymbol, pendingChanges := range pendingByToken {
        trigger := abs.getTriggerConfig(tokenSymbol)
        
        shouldBatch := false
        reason := ""
        
        // 评估触发条件
        if len(pendingChanges) >= trigger.MaxBatchSize {
            shouldBatch = true
            reason = "max_batch_size_reached"
        } else if abs.hasHighPriorityChanges(pendingChanges, trigger.Priority) {
            shouldBatch = true
            reason = "high_priority_detected"
        } else if abs.getOldestPendingAge(pendingChanges) > trigger.MaxWaitTime {
            shouldBatch = true
            reason = "max_wait_time_exceeded"
        } else if abs.getTotalValue(pendingChanges) > trigger.ValueThreshold {
            shouldBatch = true
            reason = "value_threshold_exceeded"
        }
        
        if shouldBatch {
            abs.logger.Info("Triggering batch sync",
                "token", tokenSymbol,
                "changes_count", len(pendingChanges),
                "reason", reason)
            
            // 触发批量上链
            go abs.syncManager.SyncToBlockchain(tokenSymbol, pendingChanges)
        }
    }
}
```

### 3. **应用层批处理实现**

**正确的层次关系**:
```
应用层批处理 (直接调用合约函数)
    ↓ 调用 batchMint、batchBurn 等合约函数
    ↓
ERC-4337 UserOperation层 (MasterAggregator处理)
    ↓ 签名聚合和Gas优化
    ↓
EntryPoint执行层
```

```go
// 应用层批处理服务（直接调用合约函数）
type ApplicationBatchService struct {
    contractService   *ContractService
    userWalletService *UserWalletService
    tokenRegistry     *TokenRegistryService
    logger           *logger.Logger
}

// 批量mint实现 - 直接调用合约的batchMint函数
func (abs *ApplicationBatchService) ExecuteBatchMint(ctx context.Context, req *BatchMintRequest) (*BatchResponse, error) {
    // 1. 获取代币配置
    tokenConfig, err := abs.tokenRegistry.GetTokenConfig(req.ChainID, req.TokenAddress)
    if err != nil {
        return nil, fmt.Errorf("token not supported: %w", err)
    }
    
    // 2. 验证代币是否支持batchMint
    if !tokenConfig.SupportsBatchMint() {
        return nil, fmt.Errorf("token %s does not support batch mint", req.TokenAddress)
    }
    
    // 3. 收集用户地址和金额
    var addresses []string
    var amounts []string
    
    for _, recipient := range req.Recipients {
        // 获取用户AA钱包地址
        wallet, err := abs.userWalletService.GetUserWallet(recipient.UserID, req.ChainID)
        if err != nil {
            return nil, fmt.Errorf("failed to get wallet for user %s: %w", recipient.UserID, err)
        }
        
        addresses = append(addresses, wallet.AAAddress)
        
        // 转换金额到wei
        amountWei, err := convertToWei(recipient.Amount, tokenConfig.Decimals)
        if err != nil {
            return nil, fmt.Errorf("invalid amount %s: %w", recipient.Amount, err)
        }
        amounts = append(amounts, amountWei.String())
    }
    
    // 4. 直接调用合约的batchMint函数
    contractReq := &ContractCallRequest{
        ExecutorUserID:  req.ExecutorUserID,
        ChainID:        req.ChainID,
        ContractAddress: req.TokenAddress,
        MethodName:     "batchMint",
        MethodSignature: "batchMint(address[],uint256[])",
        Parameters: []ContractParameter{
            {Type: "address[]", Value: addresses},
            {Type: "uint256[]", Value: amounts},
        },
        GasLimit: uint64(80000 + len(req.Recipients)*35000), // 基础gas + 每个mint操作
        GasMode:  "sponsored",
        Priority: "normal",
        Memo:     fmt.Sprintf("Batch Mint: %s", req.Reason),
    }
    
    // 5. 通过通用合约调用服务执行
    return abs.contractService.ExecuteContractCall(ctx, contractReq)
}
```

### 4. **区块链同步管理器**

```go
// 区块链同步管理器
type BlockchainSyncManager struct {
    contractService    *ContractService
    masterAggregator   *MasterAggregatorService
    balanceManager     *BalanceManager
    logger            *logger.Logger
}

func (bsm *BlockchainSyncManager) SyncToBlockchain(tokenSymbol string, changes []BalanceChange) error {
    // 1. 标记变更为同步中
    if err := bsm.balanceManager.UpdateChangesStatus(changes, "syncing"); err != nil {
        return err
    }
    
    // 2. 按操作类型分组
    increments := []BalanceChange{}  // 增加操作 -> batchMint
    decrements := []BalanceChange{}  // 减少操作 -> batchBurn
    
    for _, change := range changes {
        if change.Amount.Sign() > 0 {
            increments = append(increments, change)
        } else {
            decrements = append(decrements, change)
        }
    }
    
    // 3. 执行批量mint操作
    if len(increments) > 0 {
        if err := bsm.executeBatchMint(tokenSymbol, increments); err != nil {
            bsm.balanceManager.UpdateChangesStatus(increments, "sync_failed")
            return err
        }
    }
    
    // 4. 执行批量burn操作
    if len(decrements) > 0 {
        if err := bsm.executeBatchBurn(tokenSymbol, decrements); err != nil {
            bsm.balanceManager.UpdateChangesStatus(decrements, "sync_failed")
            return err
        }
    }
    
    return nil
}

func (bsm *BlockchainSyncManager) executeBatchMint(tokenSymbol string, increments []BalanceChange) error {
    // 获取代币配置
    tokenConfig := bsm.getTokenConfig(tokenSymbol)
    
    // 准备批量mint参数
    var addresses []string
    var amounts []string
    
    for _, change := range increments {
        userWallet := bsm.getUserWallet(change.UserID, tokenConfig.ChainID)
        addresses = append(addresses, userWallet.AAAddress)
        amounts = append(amounts, change.Amount.String())
    }
    
    // 调用合约batchMint函数
    contractReq := &ContractCallRequest{
        ExecutorUserID:  "system",
        ChainID:        tokenConfig.ChainID,
        ContractAddress: tokenConfig.ContractAddress,
        MethodName:     "batchMint",
        MethodSignature: "batchMint(address[],uint256[])",
        Parameters: []ContractParameter{
            {Type: "address[]", Value: addresses},
            {Type: "uint256[]", Value: amounts},
        },
        GasMode: "sponsored",
        Priority: "normal",
        Memo: fmt.Sprintf("Auto batch mint: %d operations", len(increments)),
    }
    
    response, err := bsm.contractService.ExecuteContractCall(context.Background(), contractReq)
    if err != nil {
        return err
    }
    
    // 更新同步状态
    for _, change := range increments {
        bsm.balanceManager.UpdateChangeBlockchainInfo(change.ID, response.TransactionID, "pending_confirmation")
    }
    
    return nil
}
```
