import { ethers, upgrades } from 'hardhat'

async function main() {
  console.log('\n🚀 开始部署核心合约到 BNB Testnet...')

  const [deployer] = await ethers.getSigners()
  console.log('部署者:', deployer.address)
  const { chainId, name } = await ethers.provider.getNetwork()
  console.log('网络:', name, 'chainId:', chainId)

  const deployed: Record<string, string> = {}

  // helpers
  const fs = require('fs')
  const path = require('path')
  const outDir = path.join(__dirname, '../deployments/bnbTestnet')
  const outfile = path.join(outDir, 'core.json')

  // Load previous deployments if exist
  let prev: any = {}
  if (fs.existsSync(outfile)) {
    try {
      prev = JSON.parse(fs.readFileSync(outfile, 'utf8')) || {}
      console.log('📄 发现历史部署记录，将复用已部署合约')
    } catch {}
  }

  const getPrev = (name: string): string | undefined => {
    const addr = prev?.contracts?.[name]
    return typeof addr === 'string' ? addr : undefined
  }

  const hasCode = async (address?: string) => {
    if (!address) return false
    try { const code = await ethers.provider.getCode(address); return code && code !== '0x' } catch { return false }
  }

  // Use pre-deployed EntryPoint on BNB (required)
  const entryPointAddress = process.env.ENTRYPOINT
  if (!entryPointAddress) {
    throw new Error('请设置环境变量 ENTRYPOINT 为 BNB Testnet 已部署的 EntryPoint 地址')
  }
  console.log('\n🔗 使用已有 EntryPoint:', entryPointAddress)

  // 1) CPPToken (constructor: admin, initialSupply)
  console.log('\n📦 部署 CPPToken...')
  let cppTokenAddress = getPrev('CPPToken')
  if (!(await hasCode(cppTokenAddress))) {
    const CPPToken = await ethers.getContractFactory('CPPToken')
    const initialSupply = ethers.utils.parseEther('1000000000') // 1e9 * 1e18
    const cppToken = await CPPToken.deploy(deployer.address, initialSupply)
    await cppToken.deployed()
    cppTokenAddress = cppToken.address
    console.log('✅ CPPToken:', cppTokenAddress)
  } else {
    console.log('⏭️  跳过 CPPToken，已部署:', cppTokenAddress)
  }
  deployed.CPPToken = cppTokenAddress!
  const CPPTokenF = await ethers.getContractFactory('CPPToken')
  const cppToken = CPPTokenF.attach(cppTokenAddress!)

  // 2) GasPriceOracle (UUPS)
  console.log('\n📦 部署 GasPriceOracle (UUPS)...')
  let gasPriceOracleAddress = getPrev('GasPriceOracle')
  if (!(await hasCode(gasPriceOracleAddress))) {
    const GasPriceOracle = await ethers.getContractFactory('GasPriceOracle')
    const proxy = await upgrades.deployProxy(
      GasPriceOracle,
      [deployer.address, 3600, 500],
      { initializer: 'initialize', kind: 'uups', unsafeAllow: ['constructor'] }
    )
    await proxy.deployed()
    gasPriceOracleAddress = proxy.address
    console.log('✅ GasPriceOracle (proxy):', gasPriceOracleAddress)
  } else {
    console.log('⏭️  跳过 GasPriceOracle，已部署:', gasPriceOracleAddress)
  }
  deployed.GasPriceOracle = gasPriceOracleAddress!

  // 3) MasterAggregator (UUPS)
  console.log('\n📦 部署 MasterAggregator (UUPS)...')
  let masterAggregatorAddress = getPrev('MasterAggregator')
  if (!(await hasCode(masterAggregatorAddress))) {
    const MasterAggregator = await ethers.getContractFactory('MasterAggregator')
    const proxy = await upgrades.deployProxy(
      MasterAggregator,
      [deployer.address, [deployer.address]],
      { initializer: 'initialize', kind: 'uups', unsafeAllow: ['constructor', 'state-variable-assignment'] }
    )
    await proxy.deployed()
    masterAggregatorAddress = proxy.address
    console.log('✅ MasterAggregator (proxy):', masterAggregatorAddress)
  } else {
    console.log('⏭️  跳过 MasterAggregator，已部署:', masterAggregatorAddress)
  }
  deployed.MasterAggregator = masterAggregatorAddress!
  const MasterAggregatorF = await ethers.getContractFactory('MasterAggregator')
  const masterAggregator = MasterAggregatorF.attach(masterAggregatorAddress!)

  // 4) SessionKeyManager (UUPS)
  console.log('\n📦 部署 SessionKeyManager (UUPS)...')
  let sessionKeyManagerAddress = getPrev('SessionKeyManager')
  if (!(await hasCode(sessionKeyManagerAddress))) {
    const SessionKeyManager = await ethers.getContractFactory('SessionKeyManager')
    const proxy = await upgrades.deployProxy(
      SessionKeyManager,
      [deployer.address],
      { initializer: 'initialize', kind: 'uups', unsafeAllow: ['constructor'] }
    )
    await proxy.deployed()
    sessionKeyManagerAddress = proxy.address
    console.log('✅ SessionKeyManager (proxy):', sessionKeyManagerAddress)
  } else {
    console.log('⏭️  跳过 SessionKeyManager，已部署:', sessionKeyManagerAddress)
  }
  deployed.SessionKeyManager = sessionKeyManagerAddress!

  // 5) AccountManager (UUPS)
  console.log('\n📦 部署 AccountManager (UUPS)...')
  let accountManagerAddress = getPrev('AccountManager')
  if (!(await hasCode(accountManagerAddress))) {
    const AccountManager = await ethers.getContractFactory('AccountManager')
    const proxy = await upgrades.deployProxy(
      AccountManager,
      [entryPointAddress, deployer.address],
      { initializer: 'initialize', kind: 'uups', unsafeAllow: ['constructor'] }
    )
    await proxy.deployed()
    accountManagerAddress = proxy.address
    console.log('✅ AccountManager (proxy):', accountManagerAddress)
  } else {
    console.log('⏭️  跳过 AccountManager，已部署:', accountManagerAddress)
  }
  deployed.AccountManager = accountManagerAddress!
  const AccountManagerF = await ethers.getContractFactory('AccountManager')
  const accountManager = AccountManagerF.attach(accountManagerAddress!)

  // Wire: set aggregator on AccountManager (idempotent)
  try {
    await (await accountManager.setMasterAggregator(masterAggregatorAddress!)).wait()
    console.log('🔧 AccountManager.setMasterAggregator:', masterAggregatorAddress)
  } catch {}

  // 6) GasPaymaster (constructor)
  console.log('\n📦 部署 GasPaymaster...')
  let gasPaymasterAddress = getPrev('GasPaymaster')
  if (!(await hasCode(gasPaymasterAddress))) {
    const GasPaymaster = await ethers.getContractFactory('GasPaymaster')
    const fallbackRate = ethers.utils.parseEther('1')
    const burnTokens = false
    const beneficiary = deployer.address
    const gp = await GasPaymaster.deploy(
      entryPointAddress,
      cppTokenAddress!,
      gasPriceOracleAddress!,
      fallbackRate,
      burnTokens,
      beneficiary
    )
    await gp.deployed()
    gasPaymasterAddress = gp.address
    console.log('✅ GasPaymaster:', gasPaymasterAddress)
  } else {
    console.log('⏭️  跳过 GasPaymaster，已部署:', gasPaymasterAddress)
  }
  deployed.GasPaymaster = gasPaymasterAddress!
  const GasPaymasterF = await ethers.getContractFactory('GasPaymaster')
  const gasPaymaster = GasPaymasterF.attach(gasPaymasterAddress!)

  // 7) CPNFT (UUPS)
  console.log('\n📦 部署 CPNFT (UUPS)...')
  let cpnftAddress = getPrev('CPNFT')
  if (!(await hasCode(cpnftAddress))) {
    const CPNFT = await ethers.getContractFactory('CPNFT')
    const proxy = await upgrades.deployProxy(
      CPNFT,
      ['Chapool NFT','CPNFT','http://chapool.net:8080/api/v1/meta/'],
      { initializer: 'initialize', kind: 'uups', unsafeAllow: ['constructor'] }
    )
    await proxy.deployed()
    cpnftAddress = proxy.address
    console.log('✅ CPNFT (proxy):', cpnftAddress)
  } else {
    console.log('⏭️  跳过 CPNFT，已部署:', cpnftAddress)
  }
  deployed.CPNFT = cpnftAddress!
  const CPNFTF = await ethers.getContractFactory('CPNFT')
  const cpnft = CPNFTF.attach(cpnftAddress!)

  // 8) StakingConfig
  console.log('\n📦 部署 StakingConfig...')
  let stakingConfigAddress = getPrev('StakingConfig')
  if (!(await hasCode(stakingConfigAddress))) {
    const StakingConfig = await ethers.getContractFactory('StakingConfig')
    const sc = await StakingConfig.deploy()
    await sc.deployed()
    stakingConfigAddress = sc.address
    console.log('✅ StakingConfig:', stakingConfigAddress)
  } else {
    console.log('⏭️  跳过 StakingConfig，已部署:', stakingConfigAddress)
  }
  deployed.StakingConfig = stakingConfigAddress!
  const StakingConfigF = await ethers.getContractFactory('StakingConfig')
  const stakingConfig = StakingConfigF.attach(stakingConfigAddress!)

  // 9) Staking (UUPS)
  console.log('\n📦 部署 Staking (UUPS)...')
  let stakingAddress = getPrev('Staking')
  if (!(await hasCode(stakingAddress))) {
    const Staking = await ethers.getContractFactory('Staking')
    const proxy = await upgrades.deployProxy(
      Staking,
      [cpnftAddress!, cppTokenAddress!, accountManagerAddress!, stakingConfigAddress!, deployer.address],
      { initializer: 'initialize', kind: 'uups', unsafeAllow: ['constructor'] }
    )
    await proxy.deployed()
    stakingAddress = proxy.address
    console.log('✅ Staking (proxy):', stakingAddress)
  } else {
    console.log('⏭️  跳过 Staking，已部署:', stakingAddress)
  }
  deployed.Staking = stakingAddress!
  const StakingF = await ethers.getContractFactory('Staking')
  const staking = StakingF.attach(stakingAddress!)

  // Wire staking links (idempotent)
  try { await (await stakingConfig.setStakingContract(stakingAddress!)).wait(); console.log('🔧 StakingConfig.setStakingContract:', stakingAddress) } catch {}
  try { await (await cpnft.setStakingContract(stakingAddress!)).wait(); console.log('🔧 CPNFT.setStakingContract:', stakingAddress) } catch {}

  // 10) StakingReader (UUPS)
  console.log('\n📦 部署 StakingReader (UUPS)...')
  let stakingReaderAddress = getPrev('StakingReader')
  if (!(await hasCode(stakingReaderAddress))) {
    const StakingReader = await ethers.getContractFactory('StakingReader')
    const proxy = await upgrades.deployProxy(
      StakingReader,
      [stakingAddress!, stakingConfigAddress!, cpnftAddress!, deployer.address],
      { initializer: 'initialize', kind: 'uups', unsafeAllow: ['constructor'] }
    )
    await proxy.deployed()
    stakingReaderAddress = proxy.address
    console.log('✅ StakingReader (proxy):', stakingReaderAddress)
  } else {
    console.log('⏭️  跳过 StakingReader，已部署:', stakingReaderAddress)
  }
  deployed.StakingReader = stakingReaderAddress!

  // 11) BatchTransfer
  console.log('\n📦 部署 BatchTransfer...')
  let batchTransferAddress = getPrev('BatchTransfer')
  if (!(await hasCode(batchTransferAddress))) {
    const BatchTransfer = await ethers.getContractFactory('BatchTransfer')
    const bt = await BatchTransfer.deploy()
    await bt.deployed()
    batchTransferAddress = bt.address
    console.log('✅ BatchTransfer:', batchTransferAddress)
  } else {
    console.log('⏭️  跳过 BatchTransfer，已部署:', batchTransferAddress)
  }
  deployed.BatchTransfer = batchTransferAddress!

  // 12) MockUSDT
  console.log('\n📦 部署 MockUSDT...')
  let mockUSDTAddress = getPrev('MockUSDT')
  if (!(await hasCode(mockUSDTAddress))) {
    const MockUSDT = await ethers.getContractFactory('MockUSDT')
    const mu = await MockUSDT.deploy()
    await mu.deployed()
    mockUSDTAddress = mu.address
    console.log('✅ MockUSDT:', mockUSDTAddress)
  } else {
    console.log('⏭️  跳过 MockUSDT，已部署:', mockUSDTAddress)
  }
  deployed.MockUSDT = mockUSDTAddress!

  // 13) Payment
  console.log('\n📦 部署 Payment...')
  let paymentAddress = getPrev('Payment')
  if (!(await hasCode(paymentAddress))) {
    const Payment = await ethers.getContractFactory('Payment')
    const p = await Payment.deploy()
    await p.deployed()
    paymentAddress = p.address
    console.log('✅ Payment:', paymentAddress)
  } else {
    console.log('⏭️  跳过 Payment，已部署:', paymentAddress)
  }
  deployed.Payment = paymentAddress!
  const PaymentF = await ethers.getContractFactory('Payment')
  const payment = PaymentF.attach(paymentAddress!)

  // ============================================
  // 后处理权限与白名单配置（幂等）
  // ============================================
  console.log('\n🔐 配置权限与白名单...')
  const MINTER_ROLE = 2
  const BURNER_ROLE = 4

  // 让 Staking 能够 mint 奖励（仅当尚未具备权限）
  try {
    const hasMinter = await cppToken.hasRole(stakingAddress!, MINTER_ROLE)
    if (!hasMinter) {
      await (await cppToken.grantRole(stakingAddress!, MINTER_ROLE)).wait()
      console.log('✅ CPPToken.grantRole(MINTER) -> Staking')
    } else {
      console.log('⏭️  跳过授予 MINTER_ROLE，已具备')
    }
  } catch (e) {
    console.log('⚠️  授权 MINTER_ROLE 时忽略错误:', (e as any)?.message || e)
  }

  // Payment 白名单：MockUSDT 与 CPPToken（若未添加）
  try {
    const isMuAllowed = await payment.isTokenAllowed(mockUSDTAddress!)
    if (!isMuAllowed) { await (await payment.addAllowedToken(mockUSDTAddress!)).wait(); console.log('✅ Payment.addAllowedToken -> MockUSDT') }
  } catch {}
  try {
    const isCppAllowed = await payment.isTokenAllowed(cppTokenAddress!)
    if (!isCppAllowed) { await (await payment.addAllowedToken(cppTokenAddress!)).wait(); console.log('✅ Payment.addAllowedToken -> CPPToken') }
  } catch {}

  // 可选：启用 GasPaymaster 燃烧模式（设置 GP_BURN=true）
  if (process.env.GP_BURN === 'true') {
    try {
      const hasBurner = await cppToken.hasRole(gasPaymasterAddress!, BURNER_ROLE)
      if (!hasBurner) {
        await (await cppToken.grantRole(gasPaymasterAddress!, BURNER_ROLE)).wait()
        console.log('✅ 授予 GasPaymaster BURNER_ROLE')
      }
      await (await gasPaymaster.setTokenHandlingMode(true, ethers.constants.AddressZero)).wait()
      console.log('✅ 启用 GasPaymaster 燃烧模式')
    } catch (e) {
      console.log('⚠️  配置燃烧模式时忽略错误:', (e as any)?.message || e)
    }
  }

  // Save deployment info
  if (!fs.existsSync(outDir)) fs.mkdirSync(outDir, { recursive: true })
  const info = {
    network: 'bnbTestnet',
    chainId,
    deployer: deployer.address,
    entryPoint: entryPointAddress,
    timestamp: new Date().toISOString(),
    contracts: deployed
  }
  fs.writeFileSync(outfile, JSON.stringify(info, null, 2))
  console.log('\n💾 部署信息已保存:', outfile)

  console.log('\n🎉 全部完成:')
  Object.entries(deployed).forEach(([k, v]) => console.log(`- ${k}: ${v}`))
}

main().catch((e) => {
  console.error('❌ 部署脚本出错:', e)
  process.exit(1)
})
