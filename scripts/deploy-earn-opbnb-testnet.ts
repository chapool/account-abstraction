/**
 * Deploy EARN system (Chapool Earn Vault, VeCPOT Locker, NFT Boost, Reward Distributor, Vault Reader)
 * to opBNB Testnet. Reads core.json for CPPToken, AccountManager, CPNFT, MockUSDT.
 * Writes deployments/opbnbTestnet/earn.json.
 *
 * Usage:
 *   NETWORK=opbnbTestnet npx hardhat run scripts/deploy-earn-opbnb-testnet.ts --network opbnbTestnet
 * Or ensure .env.opBNBTestnet is loaded (HARDHAT_NETWORK=opbnbTestnet).
 */
import { ethers, upgrades } from 'hardhat'
import * as fs from 'fs'
import * as path from 'path'

const MINTER_ROLE = 2

async function main() {
  console.log('\n🚀 部署 EARN 合约到 opBNB Testnet...')

  const [deployer] = await ethers.getSigners()
  console.log('部署者:', deployer.address)
  const { chainId } = await ethers.provider.getNetwork()
  console.log('chainId:', chainId)

  const outDir = path.join(__dirname, '../deployments/opbnbTestnet')
  const corePath = path.join(outDir, 'core.json')
  const earnPath = path.join(outDir, 'earn.json')

  if (!fs.existsSync(corePath)) {
    throw new Error('请先部署 core 合约: deployments/opbnbTestnet/core.json 不存在')
  }

  const core = JSON.parse(fs.readFileSync(corePath, 'utf8'))
  const contracts = core.contracts || {}
  const cppTokenAddr = contracts.CPPToken
  const accountManagerAddr = contracts.AccountManager
  const cpnftAddr = contracts.CPNFT
  const mockUsdtAddr = contracts.MockUSDT

  if (!cppTokenAddr || !accountManagerAddr || !cpnftAddr || !mockUsdtAddr) {
    throw new Error('core.json 缺少 CPPToken / AccountManager / CPNFT / MockUSDT')
  }
  console.log('📄 使用 core:', { CPPToken: cppTokenAddr, AccountManager: accountManagerAddr, CPNFT: cpnftAddr, MockUSDT: mockUsdtAddr })

  let prev: Record<string, string> = {}
  if (fs.existsSync(earnPath)) {
    try {
      const earn = JSON.parse(fs.readFileSync(earnPath, 'utf8'))
      prev = earn.contracts || {}
      console.log('📄 发现已有 earn 部署，将跳过已存在合约')
    } catch {}
  }

  const hasCode = async (address?: string) => {
    if (!address) return false
    try {
      const code = await ethers.provider.getCode(address)
      return !!(code && code !== '0x')
    } catch {
      return false
    }
  }

  const deployed: Record<string, string> = {}

  // 1) MockCPOT (testnet 无主网 CPOT 时使用)
  let mockCpotAddr = prev.MockCPOT
  if (!(await hasCode(mockCpotAddr))) {
    const MockCPOT = await ethers.getContractFactory('MockCPOT')
    const mockCpot = await MockCPOT.deploy()
    await mockCpot.deployed()
    mockCpotAddr = mockCpot.address
    console.log('✅ MockCPOT:', mockCpotAddr)
  } else {
    console.log('⏭️  跳过 MockCPOT，已部署:', mockCpotAddr)
  }
  deployed.MockCPOT = mockCpotAddr!

  // 2) ChapoolEarnVault (UUPS)
  let vaultAddr = prev.ChapoolEarnVault
  if (!(await hasCode(vaultAddr))) {
    const ChapoolEarnVault = await ethers.getContractFactory('ChapoolEarnVault')
    const vault = await upgrades.deployProxy(
      ChapoolEarnVault,
      [mockUsdtAddr, cppTokenAddr, accountManagerAddr, deployer.address],
      { initializer: 'initialize', kind: 'uups', unsafeAllow: ['constructor'] }
    )
    await vault.deployed()
    vaultAddr = vault.address
    console.log('✅ ChapoolEarnVault (proxy):', vaultAddr)
  } else {
    console.log('⏭️  跳过 ChapoolEarnVault，已部署:', vaultAddr)
  }
  deployed.ChapoolEarnVault = vaultAddr!

  const cppToken = await ethers.getContractAt('CPPToken', cppTokenAddr)
  try {
    const hasMinter = await cppToken.hasRole(vaultAddr!, MINTER_ROLE)
    if (!hasMinter) {
      await (await cppToken.grantRole(vaultAddr!, MINTER_ROLE)).wait()
      console.log('✅ CPPToken.grantRole(MINTER) -> ChapoolEarnVault')
    }
  } catch (e) {
    console.log('⚠️  CPP MINTER 授权:', (e as Error).message)
  }

  // 3) VeCPOTLocker (UUPS)
  let lockerAddr = prev.VeCPOTLocker
  if (!(await hasCode(lockerAddr))) {
    const VeCPOTLocker = await ethers.getContractFactory('VeCPOTLocker')
    const locker = await upgrades.deployProxy(
      VeCPOTLocker,
      [mockCpotAddr, deployer.address],
      { initializer: 'initialize', kind: 'uups', unsafeAllow: ['constructor'] }
    )
    await locker.deployed()
    lockerAddr = locker.address
    console.log('✅ VeCPOTLocker (proxy):', lockerAddr)
  } else {
    console.log('⏭️  跳过 VeCPOTLocker，已部署:', lockerAddr)
  }
  deployed.VeCPOTLocker = lockerAddr!

  const vault = await ethers.getContractAt('ChapoolEarnVault', vaultAddr)
  const locker = await ethers.getContractAt('VeCPOTLocker', lockerAddr)
  try {
    await (await locker.setVault(vaultAddr!)).wait()
    console.log('🔧 VeCPOTLocker.setVault:', vaultAddr)
  } catch {}
  try {
    await (await vault.setVecpotLocker(lockerAddr!)).wait()
    console.log('🔧 ChapoolEarnVault.setVecpotLocker:', lockerAddr)
  } catch {}

  // 4) NFTBoostController (UUPS)
  let nftBoostAddr = prev.NFTBoostController
  if (!(await hasCode(nftBoostAddr))) {
    const NFTBoostController = await ethers.getContractFactory('NFTBoostController')
    const nftBoost = await upgrades.deployProxy(
      NFTBoostController,
      [cpnftAddr, deployer.address],
      { initializer: 'initialize', kind: 'uups', unsafeAllow: ['constructor'] }
    )
    await nftBoost.deployed()
    nftBoostAddr = nftBoost.address
    console.log('✅ NFTBoostController (proxy):', nftBoostAddr)
  } else {
    console.log('⏭️  跳过 NFTBoostController，已部署:', nftBoostAddr)
  }
  deployed.NFTBoostController = nftBoostAddr!

  try {
    await (await vault.setNftBoostController(nftBoostAddr!)).wait()
    console.log('🔧 ChapoolEarnVault.setNftBoostController:', nftBoostAddr)
  } catch {}

  // 5) ChapoolRewardDistributor (UUPS)
  let distributorAddr = prev.ChapoolRewardDistributor
  if (!(await hasCode(distributorAddr))) {
    const ChapoolRewardDistributor = await ethers.getContractFactory('ChapoolRewardDistributor')
    const distributor = await upgrades.deployProxy(
      ChapoolRewardDistributor,
      [deployer.address],
      { initializer: 'initialize', kind: 'uups', unsafeAllow: ['constructor'] }
    )
    await distributor.deployed()
    distributorAddr = distributor.address
    console.log('✅ ChapoolRewardDistributor (proxy):', distributorAddr)
  } else {
    console.log('⏭️  跳过 ChapoolRewardDistributor，已部署:', distributorAddr)
  }
  deployed.ChapoolRewardDistributor = distributorAddr!

  const distributor = await ethers.getContractAt('ChapoolRewardDistributor', distributorAddr)
  try {
    await (await distributor.setEarnVault(vaultAddr!)).wait()
    console.log('🔧 ChapoolRewardDistributor.setEarnVault:', vaultAddr)
  } catch {}

  // 6) ChapoolVaultReader (非升级)
  let readerAddr = prev.ChapoolVaultReader
  if (!(await hasCode(readerAddr))) {
    const ChapoolVaultReader = await ethers.getContractFactory('ChapoolVaultReader')
    const reader = await ChapoolVaultReader.deploy(
      vaultAddr!,
      lockerAddr!,
      nftBoostAddr!,
      distributorAddr!,
      deployer.address
    )
    await reader.deployed()
    readerAddr = reader.address
    console.log('✅ ChapoolVaultReader:', readerAddr)
  } else {
    console.log('⏭️  跳过 ChapoolVaultReader，已部署:', readerAddr)
  }
  deployed.ChapoolVaultReader = readerAddr!

  if (!fs.existsSync(outDir)) fs.mkdirSync(outDir, { recursive: true })
  const earnInfo = {
    network: 'opbnbTestnet',
    chainId: chainId,
    deployer: deployer.address,
    timestamp: new Date().toISOString(),
    contracts: deployed,
    coreRef: { CPPToken: cppTokenAddr, AccountManager: accountManagerAddr, CPNFT: cpnftAddr, MockUSDT: mockUsdtAddr }
  }
  fs.writeFileSync(earnPath, JSON.stringify(earnInfo, null, 2))
  console.log('\n💾 已保存:', earnPath)

  console.log('\n🎉 EARN 部署完成:')
  Object.entries(deployed).forEach(([k, v]) => console.log(`  ${k}: ${v}`))
}

main().catch((e) => {
  console.error('❌', e)
  process.exit(1)
})
