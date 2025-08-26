# CPOP Contracts Go Bindings

这个包提供了CPOP (Custom Paymaster Operations Protocol) 智能合约的Go语言绑定，可以直接在Go项目中使用。

## 快速开始

### 安装

**获取最新版本：**
```bash
go get -u github.com/HzBay/account-abstraction/cpop-abis
```

**首次安装：**
```bash
go get github.com/HzBay/account-abstraction/cpop-abis
```

**重要提示：** 如果您之前使用过AAWallet和WalletManager，请注意这些合约已重构为AAccount和AccountManager。请更新您的代码以使用新的合约接口。

### 基本使用

```go
package main

import (
    "context"
    "log"
    "math/big"
    
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    
    cpop "github.com/HzBay/account-abstraction/cpop-abis"
)

func main() {
    // 连接到以太坊节点
    client, err := ethclient.Dial("https://your-rpc-endpoint")
    if err != nil {
        log.Fatal(err)
    }

    // 使用 CPOP Token 合约
    tokenAddress := common.HexToAddress("0x...") 
    token, err := cpop.NewCPOPToken(tokenAddress, client)
    if err != nil {
        log.Fatal(err)
    }

    // 查询代币余额
    balance, err := token.BalanceOf(&bind.CallOpts{}, common.HexToAddress("0x..."))
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Token balance: %s", balance.String())
}
```

## 可用的合约绑定

### 1. CPOPToken
CPOP 代币合约，支持批量操作和角色管理。

```go
token, err := cpop.NewCPOPToken(tokenAddress, client)

// 查询总供应量
totalSupply, err := token.TotalSupply(&bind.CallOpts{})

// 批量转账
recipients := []common.Address{addr1, addr2, addr3}
amounts := []*big.Int{amount1, amount2, amount3}
tx, err := token.BatchTransfer(auth, recipients, amounts)
```

### 2. AAccount
账户抽象钱包合约。

```go
account, err := cpop.NewAAccount(accountAddress, client)

// 查询钱包拥有者
owner, err := account.GetOwner(&bind.CallOpts{})

// 查询主签名者
masterSigner, err := account.GetMasterSigner(&bind.CallOpts{})

// 执行交易
tx, err := account.Execute(auth, targetAddress, value, data)
```

### 3. AccountManager
钱包管理器合约，负责创建和管理AA钱包。

```go
manager, err := cpop.NewAccountManager(managerAddress, client)

// 预测钱包地址
predictedAddr, err := manager.GetAccountAddress(&bind.CallOpts{}, owner, salt)

// 创建钱包
tx, err := manager.CreateAccount(auth, owner, salt)
```

### 4. MasterAggregator
主聚合器合约，处理签名聚合。

```go
aggregator, err := cpop.NewMasterAggregator(aggregatorAddress, client)

// 聚合签名
tx, err := aggregator.AggregateSignatures(auth, userOps)
```

### 5. GasPaymaster
Gas代付合约。

```go
paymaster, err := cpop.NewGasPaymaster(paymasterAddress, client)

// 查询支持的代币
supportedToken, err := paymaster.SupportedToken(&bind.CallOpts{})

// 获取代币价格
tokenPrice, err := paymaster.GetTokenPrice(&bind.CallOpts{}, tokenAddress)
```

### 6. GasPriceOracle
Gas价格预言机合约。

```go
oracle, err := cpop.NewGasPriceOracle(oracleAddress, client)

// 获取Gas价格
gasPrice, err := oracle.GetGasPrice(&bind.CallOpts{})
```

### 7. SessionKeyManager
会话密钥管理器合约。

```go
skm, err := cpop.NewSessionKeyManager(skmAddress, client)

// 添加会话密钥
tx, err := skm.AddSessionKey(auth, sessionKey, permissions)
```

## 环境变量

推荐使用环境变量管理配置：

```bash
# .env 文件
ETH_RPC_URL=https://your-rpc-endpoint
PRIVATE_KEY=your-private-key
CHAIN_ID=1

# 合约地址
CPOP_TOKEN_ADDRESS=0x...
ACCOUNT_MANAGER_ADDRESS=0x...
MASTER_AGGREGATOR_ADDRESS=0x...
GAS_PAYMASTER_ADDRESS=0x...
GAS_ORACLE_ADDRESS=0x...
SESSION_KEY_MANAGER_ADDRESS=0x...
```

## 完整示例

```go
package main

import (
    "context"
    "crypto/ecdsa"
    "log"
    "math/big"
    "os"
    
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
    
    cpop "github.com/HzBay/account-abstraction/cpop-abis"
)

func main() {
    // 从环境变量获取配置
    rpcURL := os.Getenv("ETH_RPC_URL")
    privateKeyHex := os.Getenv("PRIVATE_KEY")
    
    // 连接到区块链
    client, err := ethclient.Dial(rpcURL)
    if err != nil {
        log.Fatal("Failed to connect to the Ethereum client:", err)
    }
    
    // 准备私钥
    privateKey, err := crypto.HexToECDSA(privateKeyHex)
    if err != nil {
        log.Fatal("Invalid private key:", err)
    }
    
    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }
    
    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    
    // 获取链ID
    chainID, err := client.NetworkID(context.Background())
    if err != nil {
        log.Fatal("Failed to get network ID:", err)
    }
    
    // 创建交易授权
    auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
    if err != nil {
        log.Fatal("Failed to create transactor:", err)
    }
    
    // 使用CPOP Token
    tokenAddr := common.HexToAddress(os.Getenv("CPOP_TOKEN_ADDRESS"))
    token, err := cpop.NewCPOPToken(tokenAddr, client)
    if err != nil {
        log.Fatal("Failed to instantiate token contract:", err)
    }
    
    // 查询余额
    balance, err := token.BalanceOf(&bind.CallOpts{}, fromAddress)
    if err != nil {
        log.Fatal("Failed to retrieve token balance:", err)
    }
    
    log.Printf("Token balance: %s", balance.String())
    
    // 创建钱包管理器实例
    managerAddr := common.HexToAddress(os.Getenv("ACCOUNT_MANAGER_ADDRESS"))
    manager, err := cpop.NewAccountManager(managerAddr, client)
    if err != nil {
        log.Fatal("Failed to instantiate wallet manager:", err)
    }
    
    // 预测钱包地址
    salt := big.NewInt(12345)
    walletAddr, err := manager.GetAccountAddress(&bind.CallOpts{}, fromAddress, salt)
    if err != nil {
        log.Fatal("Failed to get wallet address:", err)
    }
    
    log.Printf("Predicted wallet address: %s", walletAddr.Hex())
}
```

## 依赖

- Go 1.21+
- github.com/ethereum/go-ethereum v1.16.2+

## 许可证

[MIT License](../LICENSE)

## 迁移指南

如果您从旧版本升级，请注意以下变更：

### 从 AAWallet 迁移到 AAccount

**旧代码：**
```go
wallet, err := cpop.NewAAWallet(walletAddress, client)
owner, err := wallet.Owner(&bind.CallOpts{})
```

**新代码：**
```go
account, err := cpop.NewAAccount(accountAddress, client)
owner, err := account.GetOwner(&bind.CallOpts{})
masterSigner, err := account.GetMasterSigner(&bind.CallOpts{})
```

### 从 WalletManager 迁移到 AccountManager

**旧代码：**
```go
manager, err := cpop.NewWalletManager(managerAddress, client)
```

**新代码：**
```go
manager, err := cpop.NewAccountManager(managerAddress, client)
```

### 更新后的合约功能

- **AAccount** 增加了更完善的主签名者管理
- **AccountManager** 优化了账户创建和管理逻辑
- **MasterAggregator** 更新了账户授权相关的方法名称
- **SessionKeyManager** 增强了会话密钥权限控制

## 贡献

欢迎提交问题和改进建议到项目的GitHub仓库。