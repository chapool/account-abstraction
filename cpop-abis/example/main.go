package main

import (
	"context"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	cpop "github.com/HzBay/account-abstraction/cpop-abis"
)

func main() {
	// 连接到以太坊节点
	rpcURL := os.Getenv("ETH_RPC_URL")
	if rpcURL == "" {
		rpcURL = "https://eth-mainnet.alchemyapi.io/v2/your-api-key"
	}

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatal("Failed to connect to the Ethereum client:", err)
	}

	// 使用示例合约地址 (这些需要替换为实际的合约地址)
	tokenAddress := common.HexToAddress("0x1234567890123456789012345678901234567890")

	// 创建CPOP Token实例
	token, err := cpop.NewCPOPToken(tokenAddress, client)
	if err != nil {
		log.Fatal("Failed to instantiate CPOP token contract:", err)
	}

	// 查询代币信息
	name, err := token.Name(&bind.CallOpts{})
	if err != nil {
		log.Printf("Could not retrieve token name: %v", err)
	} else {
		log.Printf("Token name: %s", name)
	}

	symbol, err := token.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Printf("Could not retrieve token symbol: %v", err)
	} else {
		log.Printf("Token symbol: %s", symbol)
	}

	totalSupply, err := token.TotalSupply(&bind.CallOpts{})
	if err != nil {
		log.Printf("Could not retrieve total supply: %v", err)
	} else {
		log.Printf("Total supply: %s", totalSupply.String())
	}

	log.Println("CPOP Go bindings are working correctly!")
}