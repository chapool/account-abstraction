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

## 核心实现要点

### 1. Go绑定使用
- 使用已生成的 `cpop-abis/payment.go` 文件
- 导入方式: `import "github.com/HzBay/account-abstraction/cpop-abis"`
- 主要使用的结构体和函数:
  - `cpop.Payment` - 合约实例
  - `cpop.PaymentBatchRefund()` - 批量退款函数
  - `cpop.PaymentRefundProcessed` - 事件结构体
  - `cpop.PaymentFilterer.WatchRefundProcessed()` - 事件监听

### 2. 关键依赖
```go
// 主要依赖包
github.com/ethereum/go-ethereum v1.13.5
github.com/gin-gonic/gin v1.9.1
github.com/streadway/amqp v1.1.0
github.com/spf13/viper v1.16.0
github.com/sirupsen/logrus v1.9.3
```

## 配置要求

### 环境变量
```yaml
# 区块链配置
ETHEREUM_RPC_URL: "https://sepolia.infura.io/v3/YOUR_PROJECT_ID"
PRIVATE_KEY: "0x..."  # 合约所有者私钥
PAYMENT_CONTRACT_ADDRESS: "0x7Da90c3364a3D0B99fCD18c1651E669C4D294D65"

# RabbitMQ配置
RABBITMQ_URL: "amqp://guest:guest@localhost:5672/"
RABBITMQ_EXCHANGE: "payment.refund"
RABBITMQ_QUEUE: "refund.processed"

# API配置
API_PORT: "8080"
API_HOST: "0.0.0.0"

# 日志配置
LOG_LEVEL: "info"
LOG_FORMAT: "json"
```

## 安全考虑

### 1. 访问控制
- API密钥认证
- IP白名单
- 请求频率限制

### 2. 私钥管理
- 使用环境变量存储私钥
- 考虑使用硬件钱包或密钥管理服务

### 3. 数据验证
- 输入参数验证
- 地址格式验证
- 金额范围检查

## 监控和日志

### 1. 日志记录
- API请求/响应日志
- 区块链交易日志
- 事件监听日志
- 错误日志

### 2. 指标监控
- API响应时间
- 交易成功率
- 事件处理延迟
- 队列消息积压

## 部署要求

### 1. 系统要求
- Go 1.19+
- RabbitMQ 3.8+
- 足够的ETH余额用于Gas费用

### 2. 网络要求
- 稳定的网络连接到以太坊节点
- 访问RabbitMQ服务器的权限

### 3. 扩展性考虑
- 支持水平扩展
- 负载均衡
- 数据库连接池

## 开发计划

### Phase 1: 基础功能 (1-2周)
- [ ] 项目结构搭建
- [ ] 配置管理
- [ ] 基础API框架
- [ ] Payment合约集成

### Phase 2: 核心功能 (2-3周)
- [ ] 批量退款API实现
- [ ] 事件监听服务
- [ ] RabbitMQ集成
- [ ] 错误处理

### Phase 3: 优化和部署 (1-2周)
- [ ] 性能优化
- [ ] 监控和日志
- [ ] 测试用例
- [ ] 部署脚本

## 测试策略

### 1. 单元测试
- 业务逻辑测试
- 合约交互测试
- 消息队列测试

### 2. 集成测试
- API端到端测试
- 事件监听测试
- 区块链交互测试

### 3. 压力测试
- API并发测试
- 事件处理性能测试

---

**创建时间**: 2025-09-17  
**版本**: 1.0  
**状态**: 需求分析完成
