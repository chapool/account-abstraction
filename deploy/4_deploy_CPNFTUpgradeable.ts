import { HardhatRuntimeEnvironment } from 'hardhat/types'
import { DeployFunction } from 'hardhat-deploy/types'
import { ethers } from 'hardhat'

const deployCPNFTUpgradeable: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const provider = ethers.provider
  const from = await provider.getSigner().getAddress()
  const network = await provider.getNetwork()
  
  // Configuration for different networks
  const networkConfig: { [key: number]: { name: string, symbol: string, baseURI: string } } = {
    31337: {  // Hardhat network
      name: "Test CPNFT Collection",
      symbol: "TCPNFT",
      baseURI: "https://test-api.example.com/nft/"
    },
    1337: {   // Ganache
      name: "Test CPNFT Collection",
      symbol: "TCPNFT",
      baseURI: "https://test-api.example.com/nft/"
    },
    11155111: {  // Sepolia
      name: "CPOP NFT Collection",
      symbol: "CPNFT",
      baseURI: "https://api.cpop.io/nft/"
    }
  }

  // Skip deployment if not on a supported network
  if (!networkConfig[network.chainId]) {
    console.log(`Skipping CPNFT Upgradeable deployment on network ${network.chainId}`)
    return
  }

  const config = networkConfig[network.chainId]
  
  console.log(`Deploying CPNFT Upgradeable to network ${network.chainId} (${network.name})`)
  console.log(`Deployer: ${from}`)

  // Force new deployment by adding a unique salt
  const timestamp = Math.floor(Date.now() / 1000)
  const deploymentName = `CPNFTUpgradeable_${timestamp}`

  // Deploy the upgradeable CPNFT contract using OpenZeppelin upgrades
  console.log('Deploying CPNFT upgradeable contract...')
  const CPNFT = await ethers.getContractFactory('CPNFT')
  
  const proxy = await hre.deployments.deploy(deploymentName, {
    contract: 'CPNFT',
    from,
    args: [config.name, config.symbol, config.baseURI],
    gasLimit: 5e6,
    log: true,
  })

  console.log(`✅ CPNFT Proxy deployed to: ${proxy.address}`)
  
  // Verify the deployment
  if (proxy.newlyDeployed) {
    const contract = await ethers.getContractAt('CPNFT', proxy.address)
    
    // Verify contract properties
    const name = await contract.name()
    const symbol = await contract.symbol()
    const owner = await contract.owner()
    const version = await contract.version()
    
    console.log(`Contract Details:`)
    console.log(`  Name: ${name}`)
    console.log(`  Symbol: ${symbol}`)
    console.log(`  Owner: ${owner}`)
    console.log(`  Version: ${version}`)
    console.log(`  Base URI: ${config.baseURI}`)

    // 保存部署信息到文件
    const fs = require('fs')
    const deploymentInfo = {
      timestamp: new Date().toISOString(),
      network: network.name,
      chainId: network.chainId,
      contract: {
        name: "CPNFT",
        address: proxy.address,
        owner: owner,
        deploymentName: deploymentName,
        config: {
          name: config.name,
          symbol: config.symbol,
          baseURI: config.baseURI
        }
      },
      deployer: from,
      transactionHash: proxy.transactionHash,
      upgradeable: true
    }

    // 确保目录存在
    if (!fs.existsSync('./deployments')) {
      fs.mkdirSync('./deployments')
    }
    if (!fs.existsSync('./deployments/sepoliaCustom')) {
      fs.mkdirSync('./deployments/sepoliaCustom')
    }

    fs.writeFileSync(
      `./deployments/sepoliaCustom/${deploymentName}.json`,
      JSON.stringify(deploymentInfo, null, 2)
    )
    console.log(`\n部署信息已保存到: ./deployments/sepoliaCustom/${deploymentName}.json`)
  }
}

// Add tags for selective deployment
deployCPNFTUpgradeable.tags = ['CPNFT', 'NFT', 'Upgradeable']

export default deployCPNFTUpgradeable
