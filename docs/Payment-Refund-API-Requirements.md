# Payment Refund API Requirements

## 项目概述

基于Payment合约的批量退款API系统，使用Golang开发，集成区块链事件监听和RabbitMQ消息队列。

## 技术架构

### 核心组件
1. **退款API服务** - 提供批量退款接口
2. **事件监听服务** - 监听RefundProcessed事件
3. **消息队列服务** - 将事件数据推送到RabbitMQ
4. **配置管理** - 环境配置和合约配置

### 技术栈
- **语言**: Go 1.19+
- **区块链交互**: go-ethereum
- **消息队列**: RabbitMQ
- **Web框架**: Gin
- **配置管理**: Viper
- **日志**: Logrus
- **数据库**: PostgreSQL (可选，用于记录)

## 功能需求

### 1. 批量退款API

#### 接口设计
```go
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

#### 响应格式
```json
{
    "txHash": "0xabcdef...",
    "blockNumber": 12345678,
    "gasUsed": 150000,
    "refundCount": 3,
    "message": "Batch refund processed successfully"
}
```

### 2. 事件监听服务

#### 监听目标
- **事件**: `RefundProcessed`
- **合约**: Payment Contract (`0x7Da90c3364a3D0B99fCD18c1651E669C4D294D65`)
- **网络**: Sepolia Testnet

与现在的payment event统一集成

#### 事件数据结构
```json
{
    "orderId": "12345",
    "user": "0x1234567890123456789012345678901234567890",
    "token": "0x0000000000000000000000000000000000000000",
    "amount": "1000000000000000000",
    "txHash": "0xabcdef...",
    "blockNumber": 12345678,
    "timestamp": "2025-09-17T07:16:03.603Z"
}
```

### 3. RabbitMQ集成

#### 队列配置
- **Exchange**: `payment.refund`
- **Queue**: `refund.processed`
- **Routing Key**: `refund.processed`

#### 消息格式
```json
{
    "eventType": "RefundProcessed",
    "orderId": "12345",
    "user": "0x1234567890123456789012345678901234567890",
    "token": "0x0000000000000000000000000000000000000000",
    "amount": "1000000000000000000",
    "txHash": "0xabcdef...",
    "blockNumber": 12345678,
    "timestamp": "2025-09-17T07:16:03.603Z"
}
```

## 核心实现要点

### 1. Go绑定使用
- 使用已生成的 `cpop-abis/payment.go` 文件
- 导入方式: `import "github.com/HzBay/account-abstraction/cpop-abis"`
- 主要使用的结构体和函数:
  - `cpop.Payment` - 合约实例
  - `cpop.PaymentBatchRefund()` - 批量退款函数
  - `cpop.PaymentRefundProcessed` - 事件结构体
  - `cpop.PaymentFilterer.WatchRefundProcessed()` - 事件监听