/**
 * 在 opBNB Testnet 上测试已部署的 EARN 合约逻辑（存款、查询、领取、取款）。
 * 使用 .env.opBNBTestnet 的账户，会调用 MockUSDT.mintToSelf 获取测试 USDT。
 *
 * Usage: HARDHAT_NETWORK=opbnbTestnet npx hardhat run scripts/test-earn-opbnb-testnet.ts --network opbnbTestnet
 */
import { ethers } from 'hardhat'
import * as fs from 'fs'
import * as path from 'path'

const parseEther = (n: string) => ethers.utils.parseEther(n)

async function main() {
  console.log('\n🧪 opBNB Testnet EARN 合约逻辑测试\n')

  const earnPath = path.join(__dirname, '../deployments/opbnbTestnet/earn.json')
  const corePath = path.join(__dirname, '../deployments/opbnbTestnet/core.json')
  if (!fs.existsSync(earnPath) || !fs.existsSync(corePath)) throw new Error('缺少 deployments/opbnbTestnet/earn.json 或 core.json')
  const earn = JSON.parse(fs.readFileSync(earnPath, 'utf8'))
  const core = JSON.parse(fs.readFileSync(corePath, 'utf8'))
  const vaultAddr = earn.contracts.ChapoolEarnVault
  const readerAddr = earn.contracts.ChapoolVaultReader
  const lockerAddr = earn.contracts.VeCPOTLocker
  const mockUsdtAddr = earn.coreRef?.MockUSDT || core.contracts.MockUSDT
  const mockCpotAddr = earn.contracts.MockCPOT

  const [signer] = await ethers.getSigners()
  const user = signer.address
  console.log('测试账户:', user)

  const vault = await ethers.getContractAt('ChapoolEarnVault', vaultAddr)
  const reader = await ethers.getContractAt('ChapoolVaultReader', readerAddr)
  const usdt = await ethers.getContractAt('MockUSDT', mockUsdtAddr)
  const cppToken = await ethers.getContractAt('CPPToken', core.contracts.CPPToken)
  const cpot = await ethers.getContractAt('MockCPOT', mockCpotAddr)
  const locker = await ethers.getContractAt('VeCPOTLocker', lockerAddr)
  const nftBoostAddr = earn.contracts.NFTBoostController
  const nftBoost = await ethers.getContractAt('NFTBoostController', nftBoostAddr)
  const cpnftAddr = earn.coreRef?.CPNFT || core.contracts.CPNFT
  const cpnft = await ethers.getContractAt('CPNFT', cpnftAddr)

  const depositAmount = parseEther('100')
  let passed = 0
  let failed = 0

  // 1) MockUSDT.mintToSelf
  try {
    const balBefore = await usdt.balanceOf(user)
    const tx = await usdt.mintToSelf(depositAmount)
    await tx.wait()
    const balAfter = await usdt.balanceOf(user)
    if (balAfter.sub(balBefore).gte(depositAmount)) {
      console.log('✅ MockUSDT.mintToSelf(100) 成功')
      passed++
    } else throw new Error('余额未增加')
  } catch (e) {
    console.log('❌ MockUSDT.mintToSelf 失败:', (e as Error).message)
    failed++
  }

  // 2) approve + deposit
  try {
    await (await usdt.approve(vaultAddr, depositAmount)).wait()
    await (await vault.deposit(depositAmount, user)).wait()
    const balance = await vault.usdtBalance(user)
    if (balance.eq(depositAmount)) {
      console.log('✅ deposit(100 USDT) 成功，usdtBalance=', balance.toString())
      passed++
    } else throw new Error('usdtBalance 不符')
  } catch (e) {
    console.log('❌ deposit 失败:', (e as Error).message)
    failed++
  }

  // 2b) VeCPOT 锁仓与加成：MockCPOT.mint -> lock -> syncBoost（veUnits=amount×days/360，需 >=10 才有 1bps）
  try {
    const cpotAmount = parseEther('10000')
    await (await cpot.mint(user, cpotAmount)).wait()
    await (await cpot.approve(lockerAddr, cpotAmount)).wait()
    await (await locker.lock(parseEther('1000'), 30)).wait()
    const veBoostBps = await locker.getBoostBps(user)
    await (await vault.syncBoost(user)).wait()
    if (veBoostBps.gt(0)) {
      console.log('✅ VeCPOTLocker: lock(1000 CPOT, 30d) 成功，getBoostBps=', veBoostBps.toString())
      passed++
    } else throw new Error('veCPOT boost 应为正')
  } catch (e) {
    console.log('❌ VeCPOTLocker 测试失败:', (e as Error).message)
    failed++
  }

  // 2c) NFT 加成：若当前账户为 CPNFT owner 则 mint -> activateBoost -> syncBoost
  const cpnftOwner = await cpnft.owner()
  if (cpnftOwner.toLowerCase() === user.toLowerCase()) {
    try {
      const levelS = 4 // CPNFT.NFTLevel.S (0=NORMAL,1=C,2=B,3=A,4=S,5=SS,6=SSS)
      const mintTx = await cpnft.mint(user, levelS)
      const receipt = await mintTx.wait()
      const transferLog = receipt.logs?.find((log: any) => log.topics[0] === ethers.utils.id('Transfer(address,address,uint256)'))
      const tokenIdUsed = transferLog && transferLog.topics[3] ? ethers.BigNumber.from(transferLog.topics[3]).toString() : '1'
      await (await nftBoost.activateBoost(tokenIdUsed)).wait()
      await (await vault.syncBoost(user)).wait()
      const nftBps = await nftBoost.getNFTBoostBps(user)
      if (nftBps.gte(50)) {
        console.log('✅ NFTBoostController: mint(S) + activateBoost 成功，getNFTBoostBps=', nftBps.toString())
        passed++
      } else throw new Error('NFT boost 预期 >= 50 bps (S 级)')
    } catch (e) {
      console.log('❌ NFT 加成测试失败:', (e as Error).message)
      failed++
    }
  } else {
    console.log('⏭️  跳过 NFT 加成测试（当前账户非 CPNFT owner）')
  }

  // 3) ChapoolVaultReader.getUserDashboard（含加成后 weightedUSDT >= deposit）
  try {
    const dash = await reader.getUserDashboard(user)
    if (!dash.usdtBalance.eq(depositAmount)) throw new Error('usdtBalance 不符')
    if (dash.weightedUSDT.lt(depositAmount)) throw new Error('加权 USDT 应 >= 存款（有加成）')
    console.log('✅ ChapoolVaultReader.getUserDashboard 正常，weightedUSDT=', dash.weightedUSDT.toString(), 'vecpotBoostBps=', dash.vecpotBoostBps.toString(), 'nftBoostBps=', dash.nftBoostBps.toString())
    passed++
  } catch (e) {
    console.log('❌ getUserDashboard 失败:', (e as Error).message)
    failed++
  }

  // 4) getProtocolOverview
  try {
    const overview = await reader.getProtocolOverview()
    if (overview.tvl.gte(depositAmount) && overview.totalUsers.gte(1)) {
      console.log('✅ getProtocolOverview 正常，TVL=', ethers.utils.formatEther(overview.tvl), 'totalUsers=', overview.totalUsers.toString())
      passed++
    } else throw new Error('overview 数据异常')
  } catch (e) {
    console.log('❌ getProtocolOverview 失败:', (e as Error).message)
    failed++
  }

  // 5) setRewardRate（若当前账户是 vault owner 则设置非零速率，便于后续 claim 测试）
  const vaultOwner = await vault.owner()
  if (vaultOwner.toLowerCase() === user.toLowerCase()) {
    try {
      const smallRate = parseEther('0.001')
      await (await vault.setRewardRate(smallRate)).wait()
      const rate = await vault.rewardRate()
      if (rate.eq(smallRate)) {
        console.log('✅ setRewardRate(0.001 CPP/s) 成功')
        passed++
      } else throw new Error('rewardRate 未更新')
    } catch (e) {
      console.log('❌ setRewardRate 失败:', (e as Error).message)
      failed++
    }
  } else {
    console.log('⏭️  跳过 setRewardRate（当前账户非 vault owner）')
  }

  // 6) getPendingCPP（可能为 0，若刚存款且 rate 刚设）
  try {
    const pending = await vault.getPendingCPP(user)
    console.log('  getPendingCPP =', ethers.utils.formatEther(pending), 'CPP')
    passed++
  } catch (e) {
    console.log('❌ getPendingCPP 失败:', (e as Error).message)
    failed++
  }

  // 7) claimCPP（pending=0 时 revert "No CPP to claim" 为预期行为）
  const pendingBeforeClaim = await vault.getPendingCPP(user)
  try {
    const cppBefore = await cppToken.balanceOf(user)
    await (await vault.claimCPP()).wait()
    const cppAfter = await cppToken.balanceOf(user)
    console.log('✅ claimCPP() 成功，CPP 余额变化:', cppBefore.toString(), '->', cppAfter.toString())
    passed++
  } catch (e: any) {
    if (pendingBeforeClaim.eq(0) && (e?.message?.includes('No CPP to claim') || e?.error?.message?.includes('No CPP to claim'))) {
      console.log('✅ claimCPP 在无待领时正确 revert: No CPP to claim')
      passed++
    } else {
      console.log('❌ claimCPP 失败:', (e as Error).message)
      failed++
    }
  }

  // 8) withdraw
  try {
    await (await vault.withdraw(depositAmount, user)).wait()
    const balanceAfter = await vault.usdtBalance(user)
    const usdtBack = await usdt.balanceOf(user)
    if (balanceAfter.eq(0) && usdtBack.gte(depositAmount)) {
      console.log('✅ withdraw(100 USDT) 成功，USDT 已回到账户')
      passed++
    } else throw new Error('取款后余额异常')
  } catch (e) {
    console.log('❌ withdraw 失败:', (e as Error).message)
    failed++
  }

  console.log('\n--- 结果 ---')
  console.log('通过:', passed, '失败:', failed)
  if (failed > 0) process.exit(1)
  console.log('🎉 全部通过，opBNB Testnet EARN 逻辑正常。\n')
}

main().catch((e) => {
  console.error(e)
  process.exit(1)
})
