import { ethers, upgrades, network } from "hardhat";
import * as dotenv from "dotenv";
import * as fs from "fs";

// Load environment variables
dotenv.config({ path: '.env.sepolia' });

interface DeploymentResult {
  contractName: string;
  proxyAddress: string;
  implementationAddress?: string;
  txHash: string;
}

async function deployCPOPContracts() {
  const networkName = network.name;
  const isLocal = networkName === "localhost" || networkName === "hardhat";
  
  console.log(`🚀 Deploying CPOP Contracts to ${networkName}\n`);

  const [deployer] = await ethers.getSigners();
  const balance = await deployer.getBalance();
  
  console.log(`📋 Deployer: ${deployer.address}`);
  console.log(`💰 Balance: ${ethers.utils.formatEther(balance)} ETH\n`);

  if (balance.lt(ethers.utils.parseEther("0.1")) && !isLocal) {
    console.warn("⚠️  Warning: Low balance, deployment may fail\n");
  }

  const deployments: DeploymentResult[] = [];

  try {
    // 1. EntryPoint - Deploy locally, use existing on Sepolia
    let entryPointAddress: string;
    
    if (isLocal) {
      console.log("1️⃣ Deploying EntryPoint (local network)...");
      const EntryPoint = await ethers.getContractFactory("EntryPoint");
      const entryPoint = await EntryPoint.deploy();
      await entryPoint.deployed();
      entryPointAddress = entryPoint.address;
      console.log(`✅ EntryPoint: ${entryPointAddress}\n`);
    } else {
      // Use existing Sepolia EntryPoint
      entryPointAddress = "0x5ff137d4b0fdcd49dca30c7cf57e578a026d2789";
      console.log("1️⃣ Using existing EntryPoint...");
      console.log(`✅ EntryPoint: ${entryPointAddress} (existing)\n`);
    }

    // 2. Deploy CPOP Token (Non-upgradeable)
    console.log("2️⃣ Deploying CPOP Token (non-upgradeable)...");
    const CPOPToken = await ethers.getContractFactory("CPOPToken");
    const initialSupply = ethers.utils.parseUnits("1000000000", 18); // 1 billion tokens
    console.log(`   Constructor: admin=${deployer.address}, initialSupply=${ethers.utils.formatUnits(initialSupply, 18)}`);
    const cpopToken = await CPOPToken.deploy(deployer.address, initialSupply);
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
    const fallbackExchangeRate = ethers.utils.parseUnits("1000000", 18); // 1M tokens per ETH
    console.log(`   Constructor: entryPoint=${entryPointAddress}, token=${cpopToken.address}`);
    console.log(`   Oracle=${gasPriceOracle.address}, fallbackRate=${ethers.utils.formatUnits(fallbackExchangeRate, 18)}`);
    console.log(`   BurnTokens=true, beneficiary=${deployer.address}`);
    
    const gasPaymaster = await GasPaymaster.deploy(
      entryPointAddress,
      cpopToken.address,
      gasPriceOracle.address,
      fallbackExchangeRate,
      true, // burnTokens
      deployer.address // beneficiary
    );
    await gasPaymaster.deployed();
    console.log(`✅ Gas Paymaster: ${gasPaymaster.address}\n`);

    // 5. Deploy Master Aggregator (Upgradeable)
    console.log("5️⃣ Deploying Master Aggregator (upgradeable)...");
    const MasterAggregator = await ethers.getContractFactory("MasterAggregator");
    const masterAggregator = await upgrades.deployProxy(MasterAggregator, 
      [entryPointAddress, [deployer.address]], // entryPoint, initial masters
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
      [entryPointAddress],
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
        entryPointAddress,
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

    // 9. Post-deployment configuration (optional for local)
    if (!isLocal) {
      console.log("9️⃣ Configuring contracts...");
      
      // Add initial deposit to paymaster on Sepolia
      console.log("💰 Adding deposit to Gas Paymaster...");
      const depositAmount = ethers.utils.parseEther("0.05");
      const EntryPoint = await ethers.getContractFactory("EntryPoint");
      const entryPoint = EntryPoint.attach(entryPointAddress);
      const depositTx = await entryPoint.depositTo(gasPaymaster.address, { value: depositAmount });
      await depositTx.wait();
      console.log(`✅ Added ${ethers.utils.formatEther(depositAmount)} ETH deposit\n`);
    }

    // 10. Test contract interactions
    console.log("🧪 Testing contract interactions...");
    
    // Test token details
    const tokenName = await cpopToken.NAME();
    const tokenSymbol = await cpopToken.SYMBOL();
    const tokenDecimals = await cpopToken.DECIMALS();
    console.log(`✅ Token: ${tokenName} (${tokenSymbol}) - ${tokenDecimals} decimals`);

    // Test master authorization
    const isAuthorized = await masterAggregator.authorizedMasters(deployer.address);
    console.log(`✅ Deployer authorized as master: ${isAuthorized}`);

    // Summary
    console.log("\n🎉 CPOP Deployment Complete!\n");
    console.log("📊 Contract Summary:");
    console.log("=".repeat(80));
    console.log(`${'Contract'.padEnd(20)} ${'Address'.padEnd(45)} Type`);
    console.log("-".repeat(80));
    console.log(`${'EntryPoint'.padEnd(20)} ${entryPointAddress.padEnd(45)} ${isLocal ? '(deployed)' : '(existing)'}`);
    console.log(`${'CPOPToken'.padEnd(20)} ${cpopToken.address.padEnd(45)} (non-upgradeable)`);
    console.log(`${'GasPaymaster'.padEnd(20)} ${gasPaymaster.address.padEnd(45)} (non-upgradeable)`);
    console.log(`${'AAccount Impl'.padEnd(20)} ${accountImpl.address.padEnd(45)} (implementation)`);
    
    deployments.forEach(d => {
      console.log(`${d.contractName.padEnd(20)} ${d.proxyAddress.padEnd(45)} (upgradeable)`);
    });
    console.log("=".repeat(80));

    // Save addresses to environment file (Sepolia only)
    if (!isLocal) {
      console.log("\n📝 Updating .env.sepolia with contract addresses...");
      const envUpdates = [
        `ENTRYPOINT_ADDRESS=${entryPointAddress}`,
        `CPOP_TOKEN_ADDRESS=${cpopToken.address}`,
        `GAS_ORACLE_ADDRESS=${gasPriceOracle.address}`,
        `GAS_PAYMASTER_ADDRESS=${gasPaymaster.address}`,
        `MASTER_AGGREGATOR_ADDRESS=${masterAggregator.address}`,
        `SESSION_KEY_MANAGER_ADDRESS=${sessionKeyManager.address}`,
        `ACCOUNT_MANAGER_ADDRESS=${accountManager.address}`,
        `AACCOUNT_IMPLEMENTATION_ADDRESS=${accountImpl.address}`
      ];

      let envContent = "";
      if (fs.existsSync('.env.sepolia')) {
        envContent = fs.readFileSync('.env.sepolia', 'utf8');
      }
      
      envUpdates.forEach(update => {
        const [key] = update.split('=');
        const regex = new RegExp(`^${key}=.*$`, 'm');
        if (regex.test(envContent)) {
          envContent = envContent.replace(regex, update);
        } else {
          envContent += `\n${update}`;
        }
      });
      
      fs.writeFileSync('.env.sepolia', envContent);
      console.log("✅ Contract addresses saved to .env.sepolia");
    }

    console.log("\n✨ All contracts deployed and configured successfully!");
    console.log(`🌐 Network: ${networkName}`);
    console.log(`👛 Deployer: ${deployer.address}`);
    
    if (!isLocal) {
      console.log("🔍 View on Etherscan: https://sepolia.etherscan.io/");
    } else {
      console.log("🔄 Ready for Sepolia deployment!");
    }

  } catch (error) {
    console.error("\n❌ Deployment failed:", error);
    
    if (deployments.length > 0) {
      console.log("\n📋 Successfully deployed contracts before failure:");
      deployments.forEach(d => {
        console.log(`${d.contractName}: ${d.proxyAddress} (proxy)`);
      });
    }
    
    throw error;
  }
}

// Execute deployment
deployCPOPContracts()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error("Deployment failed:", error);
    process.exit(1);
  });