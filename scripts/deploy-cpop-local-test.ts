import { ethers } from "hardhat";
import { upgrades } from "hardhat";

interface DeploymentResult {
  contractName: string;
  proxyAddress: string;
  implementationAddress?: string;
  txHash: string;
}

async function testCPOPDeployment() {
  console.log("🧪 Testing CPOP Upgradeable Contracts on Local Network\n");

  const [deployer] = await ethers.getSigners();
  const balance = await deployer.getBalance();
  
  console.log(`📋 Deployer: ${deployer.address}`);
  console.log(`💰 Balance: ${ethers.utils.formatEther(balance)} ETH\n`);

  const deployments: DeploymentResult[] = [];

  try {
    // 1. Deploy EntryPoint first (for local testing)
    console.log("1️⃣ Deploying EntryPoint...");
    const EntryPoint = await ethers.getContractFactory("EntryPoint");
    const entryPoint = await EntryPoint.deploy();
    await entryPoint.deployed();
    console.log(`✅ EntryPoint: ${entryPoint.address}\n`);

    // 2. Deploy CPOP Token (Non-upgradeable)
    console.log("2️⃣ Deploying CPOP Token (non-upgradeable)...");
    const CPOPToken = await ethers.getContractFactory("CPOPToken");
    const cpopToken = await CPOPToken.deploy("CPOP Token", "CPOP", 18);
    await cpopToken.deployed();
    console.log(`✅ CPOP Token: ${cpopToken.address}\n`);

    // 3. Deploy Gas Price Oracle (Upgradeable)
    console.log("3️⃣ Deploying Gas Price Oracle (upgradeable)...");
    const GasPriceOracle = await ethers.getContractFactory("GasPriceOracle");
    const gasPriceOracle = await upgrades.deployProxy(GasPriceOracle, [], {
      initializer: 'initialize',
      kind: 'uups'
    });
    await gasPriceOracle.deployed();
    
    const oracleImpl = await upgrades.erc1967.getImplementationAddress(gasPriceOracle.address);
    deployments.push({
      contractName: "GasPriceOracle",
      proxyAddress: gasPriceOracle.address,
      implementationAddress: oracleImpl,
      txHash: gasPriceOracle.deployTransaction.hash
    });
    console.log(`✅ Gas Price Oracle Proxy: ${gasPriceOracle.address}`);
    console.log(`   Implementation: ${oracleImpl}\n`);

    // 4. Deploy Gas Paymaster (Non-upgradeable)
    console.log("4️⃣ Deploying Gas Paymaster (non-upgradeable)...");
    const GasPaymaster = await ethers.getContractFactory("GasPaymaster");
    const gasPaymaster = await GasPaymaster.deploy(
      entryPoint.address,
      cpopToken.address,
      gasPriceOracle.address,
      ethers.utils.parseUnits("1000000", 18), // fallbackExchangeRate (1M tokens per ETH)
      true, // burnTokens
      deployer.address // beneficiary
    );
    await gasPaymaster.deployed();
    console.log(`✅ Gas Paymaster: ${gasPaymaster.address}\n`);

    // 5. Deploy Master Aggregator (Upgradeable)
    console.log("5️⃣ Deploying Master Aggregator (upgradeable)...");
    const MasterAggregator = await ethers.getContractFactory("MasterAggregator");
    const masterAggregator = await upgrades.deployProxy(MasterAggregator, 
      [entryPoint.address, [deployer.address]], // entryPoint, initial masters
      {
        initializer: 'initialize',
        kind: 'uups'
      }
    );
    await masterAggregator.deployed();
    
    const aggregatorImpl = await upgrades.erc1967.getImplementationAddress(masterAggregator.address);
    deployments.push({
      contractName: "MasterAggregator",
      proxyAddress: masterAggregator.address,
      implementationAddress: aggregatorImpl,
      txHash: masterAggregator.deployTransaction.hash
    });
    console.log(`✅ Master Aggregator Proxy: ${masterAggregator.address}`);
    console.log(`   Implementation: ${aggregatorImpl}\n`);

    // 6. Deploy Session Key Manager (Upgradeable)
    console.log("6️⃣ Deploying Session Key Manager (upgradeable)...");
    const SessionKeyManager = await ethers.getContractFactory("SessionKeyManager");
    const sessionKeyManager = await upgrades.deployProxy(SessionKeyManager,
      [entryPoint.address],
      {
        initializer: 'initialize',
        kind: 'uups'
      }
    );
    await sessionKeyManager.deployed();
    
    const sessionImpl = await upgrades.erc1967.getImplementationAddress(sessionKeyManager.address);
    deployments.push({
      contractName: "SessionKeyManager",
      proxyAddress: sessionKeyManager.address,
      implementationAddress: sessionImpl,
      txHash: sessionKeyManager.deployTransaction.hash
    });
    console.log(`✅ Session Key Manager Proxy: ${sessionKeyManager.address}`);
    console.log(`   Implementation: ${sessionImpl}\n`);

    // 7. Deploy AAccount Implementation
    console.log("7️⃣ Deploying AAccount Implementation...");
    const AAccount = await ethers.getContractFactory("AAccount");
    const accountImpl = await AAccount.deploy();
    await accountImpl.deployed();
    console.log(`✅ AAccount Implementation: ${accountImpl.address}\n`);

    // 8. Deploy Account Manager (Upgradeable)
    console.log("8️⃣ Deploying Account Manager (upgradeable)...");
    const AccountManager = await ethers.getContractFactory("AccountManager");
    const accountManager = await upgrades.deployProxy(AccountManager,
      [
        entryPoint.address,
        accountImpl.address,
        deployer.address, // default master signer
        masterAggregator.address
      ],
      {
        initializer: 'initialize',
        kind: 'uups'
      }
    );
    await accountManager.deployed();
    
    const managerImpl = await upgrades.erc1967.getImplementationAddress(accountManager.address);
    deployments.push({
      contractName: "AccountManager",
      proxyAddress: accountManager.address,
      implementationAddress: managerImpl,
      txHash: accountManager.deployTransaction.hash
    });
    console.log(`✅ Account Manager Proxy: ${accountManager.address}`);
    console.log(`   Implementation: ${managerImpl}\n`);

    // 9. Test contract interactions
    console.log("9️⃣ Testing contract interactions...");
    
    // Test EntryPoint deposit
    const depositAmount = ethers.utils.parseEther("1.0");
    const depositTx = await entryPoint.depositTo(gasPaymaster.address, { value: depositAmount });
    await depositTx.wait();
    console.log(`✅ Added ${ethers.utils.formatEther(depositAmount)} ETH deposit to Gas Paymaster`);

    // Test balance check
    const paymasterBalance = await entryPoint.balanceOf(gasPaymaster.address);
    console.log(`✅ Paymaster balance in EntryPoint: ${ethers.utils.formatEther(paymasterBalance)} ETH`);

    // Test master signer authorization check
    const isAuthorized = await masterAggregator.authorizedMasters(deployer.address);
    console.log(`✅ Deployer authorized as master: ${isAuthorized}`);

    console.log("\n🎉 Local Test Deployment Successful!\n");

    // Summary
    console.log("📊 Deployment Summary:");
    console.log("=" .repeat(80));
    console.log(`${'Contract'.padEnd(20)} ${'Proxy Address'.padEnd(45)} Implementation`);
    console.log("-" .repeat(80));
    console.log(`${'EntryPoint'.padEnd(20)} ${entryPoint.address.padEnd(45)} (non-upgradeable)`);
    console.log(`${'CPOPToken'.padEnd(20)} ${cpopToken.address.padEnd(45)} (non-upgradeable)`);
    console.log(`${'GasPaymaster'.padEnd(20)} ${gasPaymaster.address.padEnd(45)} (non-upgradeable)`);
    console.log(`${'AAccount Impl'.padEnd(20)} ${accountImpl.address.padEnd(45)} (implementation)`);
    
    deployments.forEach(d => {
      console.log(`${d.contractName.padEnd(20)} ${d.proxyAddress.padEnd(45)} ${d.implementationAddress || ''}`);
    });
    
    console.log("=" .repeat(80));
    console.log("\n✨ All contracts deployed successfully and working properly!");
    console.log("🔄 Ready for Sepolia deployment!");

  } catch (error) {
    console.error("❌ Local test deployment failed:", error);
    
    if (deployments.length > 0) {
      console.log("\n📋 Successfully deployed contracts before failure:");
      deployments.forEach(d => {
        console.log(`${d.contractName}: ${d.proxyAddress} (proxy)`);
      });
    }
    
    throw error;
  }
}

// Execute test deployment
testCPOPDeployment()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error("Test deployment failed:", error);
    process.exit(1);
  });