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
  
  // 查找最新的部署文件
  const deploymentsDir = './deployments/sepoliaCustom'
  if (!fs.existsSync(deploymentsDir)) {
    throw new Error('Deployments directory not found')
  }

  const files = fs.readdirSync(deploymentsDir)
    .filter((file: string) => file.includes('CPNFTUpgradeable_'))
    .sort()
    .reverse()

  if (files.length === 0) {
    throw new Error('No CPNFT upgradeable deployment found')
  }

  const latestDeployment = files[0]
  const deploymentPath = path.join(deploymentsDir, latestDeployment)
  const deploymentInfo = JSON.parse(fs.readFileSync(deploymentPath, 'utf8'))

  console.log(`Found deployment: ${latestDeployment}`)
  console.log(`Proxy address: ${deploymentInfo.contract.proxyAddress}`)
  console.log(`Current implementation: ${deploymentInfo.contract.implementationAddress}`)

  // 部署新的实现合约
  console.log('Deploying new implementation...')
  const CPNFT = await ethers.getContractFactory('CPNFT')
  
  // 使用 OpenZeppelin upgrades 插件进行升级
  const upgraded = await upgrades.upgradeProxy(deploymentInfo.contract.address, CPNFT)
  
  console.log(`Contract upgraded successfully!`)
  
  // 验证升级
  const contract = await ethers.getContractAt('CPNFT', upgraded.address)
  const version = await contract.version()
  console.log(`New contract version: ${version}`)
  
  console.log('✅ Upgrade completed successfully!')

  // 保存升级信息
  const upgradeInfo = {
    timestamp: new Date().toISOString(),
    proxyAddress: deploymentInfo.contract.address,
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
