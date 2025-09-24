import { ethers, upgrades } from 'hardhat'
import { Contract } from 'ethers'

async function main() {
  console.log('Starting CPNFT upgrade process...')

  // 获取部署者
  const [deployer] = await ethers.getSigners()
  console.log('Deploying contracts with the account:', deployer.address)

  // 读取现有部署信息
  const fs = require('fs')
  const path = require('path')
  
  // 查找现有的 CPNFT 部署文件
  const deploymentsDir = './deployments/sepoliaCustom'
  if (!fs.existsSync(deploymentsDir)) {
    throw new Error('Deployments directory not found')
  }

  const cpnftDeploymentPath = path.join(deploymentsDir, 'CPNFT.json')
  if (!fs.existsSync(cpnftDeploymentPath)) {
    throw new Error('CPNFT deployment not found')
  }

  const deploymentInfo = JSON.parse(fs.readFileSync(cpnftDeploymentPath, 'utf8'))
  const metadata = JSON.parse(deploymentInfo.metadata)

  console.log(`Found CPNFT deployment`)
  console.log(`Proxy address: ${deploymentInfo.address}`)
  console.log(`Current implementation: ${metadata.implementation}`)

  // 部署新的实现合约
  console.log('Deploying new implementation...')
  const CPNFT = await ethers.getContractFactory('CPNFT')
  
  // 使用 OpenZeppelin upgrades 插件进行升级
  const upgraded = await upgrades.upgradeProxy(deploymentInfo.address, CPNFT)
  
  console.log(`Contract upgraded successfully!`)
  
  // 验证升级
  const contract = await ethers.getContractAt('CPNFT', upgraded.address)
  const version = await contract.version()
  console.log(`New contract version: ${version}`)
  
  console.log('✅ Upgrade completed successfully!')

  // 获取新的实现地址
  const newImplementationAddress = await upgrades.erc1967.getImplementationAddress(upgraded.address)
  
  // 保存升级信息
  const upgradeInfo = {
    timestamp: new Date().toISOString(),
    proxyAddress: deploymentInfo.address,
    oldImplementation: metadata.implementation,
    newImplementation: newImplementationAddress,
    version: version,
    deployer: deployer.address
  }

  const upgradeFileName = `upgrade_${Date.now()}.json`
  fs.writeFileSync(
    path.join(deploymentsDir, upgradeFileName),
    JSON.stringify(upgradeInfo, null, 2)
  )
  
  console.log(`Upgrade info saved to: ${path.join(deploymentsDir, upgradeFileName)}`)
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error)
    process.exit(1)
  })
