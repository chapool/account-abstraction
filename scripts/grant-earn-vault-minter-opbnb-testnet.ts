/**
 * Grant CPPToken MINTER_ROLE to ChapoolEarnVault on opBNB Testnet.
 * 必须由 CPPToken 的 admin（即部署 core 的账户）执行。
 *
 * 若当前 .env 不是 core 部署者，可设置环境变量 CORE_DEPLOYER_PRIVATE_KEY 为 core 部署者私钥后再运行。
 *
 * Usage:
 *   HARDHAT_NETWORK=opbnbTestnet npx hardhat run scripts/grant-earn-vault-minter-opbnb-testnet.ts --network opbnbTestnet
 *   # 或使用 core 部署者私钥：
 *   CORE_DEPLOYER_PRIVATE_KEY=0x... npx hardhat run scripts/grant-earn-vault-minter-opbnb-testnet.ts --network opbnbTestnet
 */
import { ethers } from 'hardhat'
import * as fs from 'fs'
import * as path from 'path'

const MINTER_ROLE = 2

async function main() {
  const earnPath = path.join(__dirname, '../deployments/opbnbTestnet/earn.json')
  const corePath = path.join(__dirname, '../deployments/opbnbTestnet/core.json')
  if (!fs.existsSync(earnPath)) throw new Error('deployments/opbnbTestnet/earn.json 不存在')
  const earn = JSON.parse(fs.readFileSync(earnPath, 'utf8'))
  const vaultAddr = earn.contracts?.ChapoolEarnVault
  let cppTokenAddr = earn.coreRef?.CPPToken
  if (!cppTokenAddr && fs.existsSync(corePath)) {
    const core = JSON.parse(fs.readFileSync(corePath, 'utf8'))
    cppTokenAddr = core.contracts?.CPPToken
  }
  if (!vaultAddr || !cppTokenAddr) throw new Error('earn.json 缺少 ChapoolEarnVault 或 CPPToken')

  let signer
  const corePk = process.env.CORE_DEPLOYER_PRIVATE_KEY
  if (corePk) {
    signer = new ethers.Wallet(corePk, ethers.provider)
    console.log('使用 CORE_DEPLOYER_PRIVATE_KEY 对应账户:', signer.address)
  } else {
    ;[signer] = await ethers.getSigners()
    console.log('当前账户:', signer.address)
  }
  const coreDeployer = fs.existsSync(corePath) ? (JSON.parse(fs.readFileSync(corePath, 'utf8')).deployer as string) : ''
  if (coreDeployer && signer.address.toLowerCase() !== coreDeployer.toLowerCase()) {
    console.log('⚠️  当前账户不是 core 部署者 (' + coreDeployer + ')，授权可能 revert。请设置 CORE_DEPLOYER_PRIVATE_KEY 后重试。')
  }
  console.log('CPPToken:', cppTokenAddr)
  console.log('ChapoolEarnVault:', vaultAddr)

  const cpp = await ethers.getContractAt('CPPToken', cppTokenAddr)
  const has = await cpp.hasRole(vaultAddr, MINTER_ROLE)
  if (has) {
    console.log('✅ ChapoolEarnVault 已具备 MINTER_ROLE，无需重复授权')
    return
  }
  const tx = await cpp.connect(signer).grantRole(vaultAddr, MINTER_ROLE)
  await tx.wait()
  console.log('✅ 已授予 ChapoolEarnVault MINTER_ROLE，tx:', tx.hash)
}

main().catch((e) => {
  console.error(e)
  process.exit(1)
})
