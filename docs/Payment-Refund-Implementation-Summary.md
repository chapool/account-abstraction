# Payment Refund API 实现总结

## 需求概述

基于Payment合约的批量退款API系统，使用Golang开发，集成区块链事件监听和RabbitMQ消息队列。

## 已完成准备工作

### ✅ 1. 合约功能更新
- Payment合约已添加批量退款功能 `batchRefund()`
- 新增 `RefundProcessed` 事件用于跟踪退款
- 合约已部署到Sepolia测试网: `0x7Da90c3364a3D0B99fCD18c1651E669C4D294D65`

### ✅ 2. Go绑定更新
- 重新生成了 `cpop-abis/payment.go` 文件
- 包含新的 `BatchRefund()` 函数和 `RefundProcessed` 事件监听功能
- 支持事件过滤和订阅: `WatchRefundProcessed()`

## 技术实现方案

### 1. 批量退款API

#### 接口规范
```
POST /api/v1/payment/refund
Content-Type: application/json

[
    {
        "orderId": "12345",
        "user": "0x1234567890123456789012345678901234567890",
        "amount": "1000000000000000000",
        "token": "0x0000000000000000000000000000000000000000"
    }
]
```

#### 核心实现逻辑
```go
// 使用已生成的Go绑定
import "github.com/HzBay/account-abstraction/cpop-abis"

// 1. 创建合约实例
payment, err := cpop.NewPayment(contractAddress, client)

// 2. 准备批量退款参数
orderIDs := []*big.Int{...}
users := []common.Address{...}
amounts := []*big.Int{...}
tokens := []common.Address{...}

// 3. 执行批量退款
tx, err := payment.BatchRefund(auth, orderIDs, users, amounts, tokens)
```

### 2. 事件监听服务

#### 监听RefundProcessed事件
```go
// 使用Go绑定的事件监听功能
filterer, err := cpop.NewPaymentFilterer(contractAddress, client)

// 创建事件通道
sink := make(chan *cpop.PaymentRefundProcessed)

// 开始监听
sub, err := filterer.WatchRefundProcessed(nil, sink, nil, nil, nil)

// 处理事件
for event := range sink {
    // 将事件数据发送到RabbitMQ
    publishToRabbitMQ(event)
}
```

### 3. RabbitMQ集成

#### 消息格式
```json
{
    "eventType": "RefundProcessed",
    "contractAddress": "0x7Da90c3364a3D0B99fCD18c1651E669C4D294D65",
    "data": {
        "orderId": "12345",
        "user": "0x1234567890123456789012345678901234567890",
        "token": "0x0000000000000000000000000000000000000000",
        "amount": "1000000000000000000",
        "txHash": "0xabcdef...",
        "blockNumber": 12345678,
        "timestamp": "2025-09-17T07:16:03.603Z"
    }
}
```

#### 队列配置
- **Exchange**: `payment.refund`
- **Queue**: `refund.processed`
- **Routing Key**: `refund.processed`

## 关键配置

### 环境变量
```bash
# 区块链配置
ETHEREUM_RPC_URL="https://sepolia.infura.io/v3/YOUR_PROJECT_ID"
PRIVATE_KEY="0x..."  # 合约所有者私钥
PAYMENT_CONTRACT_ADDRESS="0x7Da90c3364a3D0B99fCD18c1651E669C4D294D65"

# RabbitMQ配置
RABBITMQ_URL="amqp://guest:guest@localhost:5672/"
RABBITMQ_EXCHANGE="payment.refund"
RABBITMQ_QUEUE="refund.processed"

# API配置
API_PORT="8080"
API_HOST="0.0.0.0"
```

## 安全考虑

### 1. 访问控制
- 只有合约所有者可以执行批量退款
- API需要身份验证和授权
- 输入参数验证和过滤

### 2. 错误处理
- 合约调用失败处理
- 网络连接异常处理
- 消息队列连接异常处理

### 3. 监控和日志
- 记录所有API请求和响应
- 记录区块链交易状态
- 记录事件监听和处理状态

## 部署架构

### 服务组件
1. **API服务** - 处理退款请求
2. **事件监听服务** - 监听区块链事件
3. **RabbitMQ** - 消息队列服务
4. **以太坊节点** - 区块链连接

### 扩展性
- 支持多实例部署
- 负载均衡
- 水平扩展能力

## 开发优先级

### Phase 1: 核心功能
1. 批量退款API实现
2. Payment合约集成
3. 基础错误处理

### Phase 2: 事件系统
1. RefundProcessed事件监听
2. RabbitMQ集成
3. 消息发布功能

### Phase 3: 完善功能
1. 监控和日志
2. 性能优化
3. 部署脚本

---

**状态**: 技术方案已确定，Go绑定已更新  
**下一步**: 开始API服务开发  
**预计开发时间**: 2-3周
