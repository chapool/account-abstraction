# CPOP Go Bindings 使用指南

## 目录结构

```
cpop-abis/                    # ABI 文件目录
├── CPOPToken.abi.json       # CPOP 代币合约 ABI
├── AAWallet.abi.json        # 账户抽象钱包 ABI
├── WalletManager.abi.json   # 钱包管理器 ABI
├── MasterAggregator.abi.json # 主聚合器 ABI
├── GasPaymaster.abi.json    # Gas 代付合约 ABI
├── GasPriceOracle.abi.json  # Gas 价格预言机 ABI
├── SessionKeyManager.abi.json # 会话密钥管理器 ABI
└── generate-bindings.sh     # Go 绑定生成脚本

bindings/                     # Go 绑定文件目录
├── cpop_token.go            # CPOP 代币 Go 绑定
├── aa_wallet.go             # 账户抽象钱包 Go 绑定
├── wallet_manager.go        # 钱包管理器 Go 绑定
├── master_aggregator.go     # 主聚合器 Go 绑定
├── gas_paymaster.go         # Gas 代付合约 Go 绑定
├── gas_price_oracle.go      # Gas 价格预言机 Go 绑定
└── session_key_manager.go   # 会话密钥管理器 Go 绑定
```

## 在后端项目中使用

### 方法1: 通过GitHub直接导入 (推荐)

```bash
# 在你的Go项目中安装
go get github.com/HzBay/account-abstraction/cpop-abis
```

```go
package main

import (
    "context"
    "log"
    "math/big"
    
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    
    // 直接导入 CPOP 合约绑定
    cpop "github.com/HzBay/account-abstraction/cpop-abis"
)

func main() {
    // 连接到以太坊节点
    client, err := ethclient.Dial("https://your-rpc-endpoint")
    if err != nil {
        log.Fatal(err)
    }

    // 使用 CPOP Token 合约
    tokenAddress := common.HexToAddress("0x...") // CPOP Token 合约地址
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

    // 使用 AAWallet 合约
    walletAddress := common.HexToAddress("0x...") // AA Wallet 合约地址
    wallet, err := cpop.NewAAWallet(walletAddress, client)
    if err != nil {
        log.Fatal(err)
    }

    // 查询钱包拥有者
    owner, err := wallet.Owner(&bind.CallOpts{})
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Wallet owner: %s", owner.Hex())
}
```

### 方法2: 复制绑定文件到你的项目

```bash
# 将 cpop-abis 目录复制到你的 Go 项目中
cp -r cpop-abis/ /path/to/your/go/project/pkg/contracts/cpop/
```

然后在代码中导入本地路径：
```go
import "your-module/pkg/contracts/cpop"
```
```

### 3. 具体合约使用示例

#### CPOP Token 合约操作

```go
// 创建 CPOP Token 实例
token, err := cpop.NewCPOPToken(tokenAddress, client)

// 查询总供应量
totalSupply, err := token.TotalSupply(&bind.CallOpts{})

// 查询余额
balance, err := token.BalanceOf(&bind.CallOpts{}, userAddress)

// 转账 (需要私钥签名)
auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
tx, err := token.Transfer(auth, toAddress, amount)
```

#### AA Wallet 合约操作

```go
// 创建 AA Wallet 实例
wallet, err := cpop.NewAAWallet(walletAddress, client)

// 查询钱包拥有者
owner, err := wallet.Owner(&bind.CallOpts{})

// 执行交易
auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
tx, err := wallet.Execute(auth, targetAddress, value, data)
```

#### Gas Paymaster 合约操作

```go
// 创建 Gas Paymaster 实例
paymaster, err := cpop.NewGasPaymaster(paymasterAddress, client)

// 查询支持的代币
supportedToken, err := paymaster.SupportedToken(&bind.CallOpts{})

// 获取代币价格
tokenPrice, err := paymaster.GetTokenPrice(&bind.CallOpts{}, tokenAddress)
```

## 重新生成绑定

如果合约有更新，可以重新生成 Go 绑定：

```bash
# 进入项目目录
cd /path/to/account-abstraction

# 重新编译合约（如果有更新）
npx hardhat compile

# 重新生成 ABI 文件
./scripts/extract-cpop-abis.sh  # 如果有这个脚本

# 或者手动重新生成绑定
cd cpop-abis
./generate-bindings.sh
```

## 依赖包

确保你的 Go 项目包含以下依赖：

```go
go get github.com/ethereum/go-ethereum
```

## 注意事项

1. **合约地址**: 确保使用正确的合约地址，不同网络地址不同
2. **RPC 端点**: 配置正确的以太坊 RPC 端点
3. **私钥管理**: 安全地管理私钥，不要硬编码在代码中
4. **Gas 费用**: 注意交易的 Gas 费用设置
5. **错误处理**: 妥善处理合约调用可能的错误

## 环境变量示例

```bash
# .env 文件
ETH_RPC_URL=https://your-rpc-endpoint
PRIVATE_KEY=your-private-key
CPOP_TOKEN_ADDRESS=0x...
WALLET_MANAGER_ADDRESS=0x...
```