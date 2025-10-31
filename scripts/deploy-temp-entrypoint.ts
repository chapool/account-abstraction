import { ethers } from 'hardhat'

async function main() {
  console.log('\n🚀 部署临时 EntryPoint 到 BNB Testnet...')
  const [deployer] = await ethers.getSigners()
  console.log('部署者:', deployer.address)
  const EntryPoint = await ethers.getContractFactory('EntryPoint')
  const ep = await EntryPoint.deploy()
  await ep.deployed()
  console.log('✅ EntryPoint 部署完成:', ep.address)
}

main().catch((e) => {
  console.error('❌ 部署临时 EntryPoint 失败:', e)
  process.exit(1)
})
