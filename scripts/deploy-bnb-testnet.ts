import { ethers, upgrades } from 'hardhat'
import { DeployFunction } from 'hardhat-deploy/types'

/**
 * 部署脚本 - BNB Testnet
 * 按照依赖顺序部署所有合约
 */
async function main() {
  console.log('🚀 开始在 BNB Testnet 上部署合约...\n')

  const [deployer] = await ethers.getSigners()
  console.log('部署者地址:', deployer.address)
  console.log('网络:', await ethers.provider.getNetwork())
  console.log()

  // 检查账户余额
  const balance = await deployer.getBalance()
  console.log('账户余额:', ethers.utils.formatEther(balance), 'BNB')
  console.log()

  const deploymentAddresses: { [key: string]: string } = {}

  try {
    // ============================================
    // 1. 部署 CPOP Token
    // ============================================
    console.log('📦 部署 CPOP Token...')
    const CPOPToken = await ethers.getContractFactory('CPOPToken')
    const cpopToken = await upgrades.deployProxy(
      CPOPToken,
      [
        'Chapool Protocol', // name
        'CPOP',            // symbol
        ethers.utils.parseEther('1000000000') // 10亿代币初始供应
      ],
      {
        initializer: 'initialize',
        kind: 'uups'
      }
    )
    await cpopToken.deployed()
    deploymentAddresses.CPOPToken = cpopToken.address
    console.log('✅ CPOP Token 已部署:', cpopToken.address)

    // ============================================
    // 2. 部署 Gas Price Oracle
    // ============================================
    console.log('📦 部署 Gas Price Oracle...')
    const GasPriceOracle = await ethers.getContractFactory('GasPriceOracle')
    const gasPriceOracle = await upgrades.deployProxy(
      GasPriceOracle,
      [],
      {
        initializer: 'initialize',
        kind: 'uups'
      }
    )
    await gasPriceOracle.deployed()
    deploymentAddresses.GasPriceOracle = gasPriceOracle.address
    console.log('✅ Gas Price Oracle 已部署:', gasPriceOracle.address)

    // ============================================
    // 3. 部署 Gas Paymaster
    // ============================================
    console.log('📦 部署 Gas Paymaster...')
    const GasPaymaster = await ethers.getContractFactory('GasPaymaster')
    const gasPaymaster = await upgrades.deployProxy(
      GasPaymaster,
      [
        gasPriceOracle.address, // _gasPriceOracle
        cpopToken.address,      // _cpopToken
        deployer.address        // _owner
      ],
      {
        initializer: 'initialize',
        kind: 'uups'
      }
    )
    await gasPaymaster.deployed()
    deploymentAddresses.GasPaymaster = gasPaymaster.address
    console.log('✅ Gas Paymaster 已部署:', gasPaymaster.address)

    // ============================================
    // 4. 部署 Master Aggregator
    // ============================================
    console.log('📦 部署 Master Aggregator...')
    const MasterAggregator = await ethers.getContractFactory('MasterAggregator')
    const masterAggregator = await upgrades.deployProxy(
      MasterAggregator,
      [],
      {
        initializer: 'initialize',
        kind: 'uups'
      }
    )
    await masterAggregator.deployed()
    deploymentAddresses.MasterAggregator = masterAggregator.address
    console.log('✅ Master Aggregator 已部署:', masterAggregator.address)

    // ============================================
    // 5. 部署 Session Key Manager
    // ============================================
    console.log('📦 部署 Session Key Manager...')
    const SessionKeyManager = await ethers.getContractFactory('SessionKeyManager')
    const sessionKeyManager = await upgrades.deployProxy(
      SessionKeyManager,
      [],
      {
        initializer: 'initialize',
        kind: 'uups'
      }
    )
    await sessionKeyManager.deployed()
    deploymentAddresses.SessionKeyManager = sessionKeyManager.address
    console.log('✅ Session Key Manager 已部署:', sessionKeyManager.address)

    // ============================================
    // 6. 部署 AAccount Implementation
    // ============================================
    console.log('📦 部署 AAccount Implementation...')
    const AAccount = await ethers.getContractFactory('AAccount')
    const aAccountImpl = await AAccount.deploy()
    await aAccountImpl.deployed()
    deploymentAddresses.AAccountImplementation = aAccountImpl.address
    console.log('✅ AAccount Implementation 已部署:', aAccountImpl.address)

    // ============================================
    // 7. 部署 Account Manager
    // ============================================
    console.log('📦 部署 Account Manager...')
    const AccountManager = await ethers.getContractFactory('AccountManager')
    const accountManager = await upgrades.deployProxy(
      AccountManager,
      [
        aAccountImpl.address,      // _implementation
        masterAggregator.address,  // _aggregator
        sessionKeyManager.address, // _sessionKeyManager
        gasPaymaster.address       // _gasPaymaster
      ],
      {
        initializer: 'initialize',
        kind: 'uups'
      }
    )
    await accountManager.deployed()
    deploymentAddresses.AccountManager = accountManager.address
    console.log('✅ Account Manager 已部署:', accountManager.address)

    // ============================================
    // 8. 部署 CPNFT
    // ============================================
    console.log('📦 部署 CPNFT...')
    const CPNFT = await ethers.getContractFactory('CPNFT')
    const cpnft = await upgrades.deployProxy(
      CPNFT,
      [
        'Chapool NFT', // name
        'CPNFT',       // symbol
        'http://chapool.net:8080/api/v1/meta/' // baseURI
      ],
      {
        initializer: 'initialize',
        kind: 'uups'
      }
    )
    await cpnft.deployed()
    deploymentAddresses.CPNFT = cpnft.address
    console.log('✅ CPNFT 已部署:', cpnft.address)

    // ============================================
    // 9. 部署 Staking Config
    // ============================================
    console.log('📦 部署 Staking Config...')
    const StakingConfig = await ethers.getContractFactory('StakingConfig')
    const stakingConfig = await StakingConfig.deploy()
    await stakingConfig.deployed()
    deploymentAddresses.StakingConfig = stakingConfig.address
    console.log('✅ Staking Config 已部署:', stakingConfig.address)

    // ============================================
    // 10. 部署 Staking
    // ============================================
    console.log('📦 部署 Staking...')
    const Staking = await ethers.getContractFactory('Staking')
    const staking = await upgrades.deployProxy(
      Staking,
      [
        cpnft.address,           // _cpnftContract
        cpopToken.address,       // _cpopTokenContract
        accountManager.address,  // _accountManagerContract
        stakingConfig.address,   // _configContract
        deployer.address         // _owner
      ],
      {
        initializer: 'initialize',
        kind: 'uups'
      }
    )
    await staking.deployed()
    deploymentAddresses.Staking = staking.address
    console.log('✅ Staking 已部署:', staking.address)

    // ============================================
    // 11. 部署 Staking Reader
    // ============================================
    console.log('📦 部署 Staking Reader...')
    const StakingReader = await ethers.getContractFactory('StakingReader')
    const stakingReader = await upgrades.deployProxy(
      StakingReader,
      [
        staking.address,         // _stakingContract
        stakingConfig.address,   // _configContract
        cpnft.address,           // _cpnftContract
        deployer.address         // _owner
      ],
      {
        initializer: 'initialize',
        kind: 'uups'
      }
    )
    await stakingReader.deployed()
    deploymentAddresses.StakingReader = stakingReader.address
    console.log('✅ Staking Reader 已部署:', stakingReader.address)

    // ============================================
    // 12. 部署 Batch Transfer
    // ============================================
    console.log('📦 部署 Batch Transfer...')
    const BatchTransfer = await ethers.getContractFactory('BatchTransfer')
    const batchTransfer = await BatchTransfer.deploy()
    await batchTransfer.deployed()
    deploymentAddresses.BatchTransfer = batchTransfer.address
    console.log('✅ Batch Transfer 已部署:', batchTransfer.address)

    // ============================================
    // 后处理配置
    // ============================================
    console.log('\n🔧 配置合约关系...')

    // 设置 Staking Config 的 Staking 合约地址
    await stakingConfig.setStakingContract(staking.address)
    console.log('✅ Staking Config 已配置 Staking 合约地址')

    // 设置 CPNFT 的 Staking 合约地址
    await cpnft.setStakingContract(staking.address)
    console.log('✅ CPNFT 已配置 Staking 合约地址')

    // ============================================
    // 保存部署信息
    // ============================================
    console.log('\n💾 保存部署信息...')

    const fs = require('fs')
    const path = require('path')

    const deploymentPath = path.join(__dirname, '../deployments/bnbTestnet')
    if (!fs.existsSync(deploymentPath)) {
      fs.mkdirSync(deploymentPath, { recursive: true })
    }

    const deploymentInfo = {
      network: 'bnbTestnet',
      chainId: 97,
      deployer: deployer.address,
      timestamp: new Date().toISOString(),
      contracts: deploymentAddresses
    }

    fs.writeFileSync(
      path.join(deploymentPath, 'deployment.json'),
      JSON.stringify(deploymentInfo, null, 2)
    )

    console.log('✅ 部署信息已保存到:', path.join(deploymentPath, 'deployment.json'))

    // ============================================
    // 部署总结
    // ============================================
    console.log('\n🎉 部署完成！')
    console.log('========================================')
    console.log('合约地址汇总:')
    console.log('========================================')

    Object.entries(deploymentAddresses).forEach(([name, address]) => {
      console.log(`${name.padEnd(20)}: ${address}`)
    })

    console.log('========================================')
    console.log('🔗 BNB Testnet Explorer: https://testnet.bscscan.com/')
    console.log('========================================')

  } catch (error) {
    console.error('❌ 部署失败:', error)
    process.exit(1)
  }
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error('❌ 脚本执行失败:', error)
    process.exit(1)
  })
