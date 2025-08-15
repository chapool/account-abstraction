# CustodialBridge 设计文档

CustodialBridge是基于CPOP账户抽象技术栈的通用钱包Relayer服务，为Web2开发者提供托管式区块链资产管理的桥接服务。

## 文档结构

### 核心设计
- **[CustodialBridge-Design.md](./CustodialBridge-Design.md)** - 完整系统架构和API设计
- **[CustodialBridge-API-Design.md](./CustodialBridge-API-Design.md)** - 基于go-starter的API开发流程
- **[Chain-Sync-Strategy.md](./Chain-Sync-Strategy.md)** - 链上链下数据同步策略

### 实现指南
- **[CustodialBridge-Development-Reference.md](./CustodialBridge-Development-Reference.md)** - 完整开发参考手册
- **[CustodialBridge-Master-Config.md](./CustodialBridge-Master-Config.md)** - 多Master密钥管理配置
- **[CustodialBridge-Contract-API.md](./CustodialBridge-Contract-API.md)** - 通用合约调用API设计（暂未采用）

## 项目概述

### 技术栈
- **后端框架**: Go + go-starter
- **API设计**: Swagger-First开发模式
- **数据库**: PostgreSQL + Redis
- **区块链**: 基于CPOP账户抽象系统
- **容器化**: Docker + Kubernetes

### 核心功能
1. **钱包管理** - AA钱包自动部署和管理
2. **资产转账** - 支持ETH、ERC20、CPOP、NFT转账
3. **批量处理** - 智能批量操作节省75-85% Gas费用
4. **多链支持** - Ethereum、BSC、Polygon等主流网络
5. **实时同步** - 链上链下数据一致性保障

### 开发状态
当前处于设计阶段，等待CPOP合约开发完成后启动实施。

## 快速开始

详细的开发指南请参考 [CustodialBridge-Development-Reference.md](./CustodialBridge-Development-Reference.md)

```bash
# 项目初始化（未来）
git clone https://github.com/allaboutapps/go-starter.git custodial-bridge
cd custodial-bridge

# 修改配置并启动
make build
make dev
```

## 联系方式

如有问题请联系项目团队。