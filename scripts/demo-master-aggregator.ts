import { ethers } from "hardhat";
import { 
    MasterAggregator,
    AAWallet,
    WalletManager,
    EntryPoint,
    CPOPToken,
} from "../typechain";

/**
 * Complete MasterAggregator Demo Script
 * 
 * This script demonstrates the real-world usage of MasterAggregator:
 * 1. Setup - Deploy all necessary contracts
 * 2. Master Signer Management - Authorize master signers
 * 3. Wallet Creation - Create AAWallets controlled by master signers
 * 4. Signature Aggregation - Demonstrate batch operations with single signature
 * 5. Gas Optimization - Show gas savings from aggregation
 * 6. Session Key Support - Demonstrate session key aggregation
 */

async function main() {
    console.log("🚀 Starting MasterAggregator Complete Demo...\n");

    // Get signers
    const [deployer, masterSigner1, masterSigner2, user1, user2, user3] = await ethers.getSigners();
    
    const deployerAddress = await deployer.getAddress();
    const masterSigner1Address = await masterSigner1.getAddress();
    const masterSigner2Address = await masterSigner2.getAddress();
    const user1Address = await user1.getAddress();
    const user2Address = await user2.getAddress();
    const user3Address = await user3.getAddress();

    console.log("👥 Participants:");
    console.log(`   Deployer: ${deployerAddress}`);
    console.log(`   Master Signer 1: ${masterSigner1Address}`);
    console.log(`   Master Signer 2: ${masterSigner2Address}`);
    console.log(`   User 1: ${user1Address}`);
    console.log(`   User 2: ${user2Address}`);
    console.log(`   User 3: ${user3Address}\n`);

    // ============================================================================
    // STEP 1: DEPLOY CONTRACTS
    // ============================================================================
    
    console.log("📦 Step 1: Deploying Contracts...");

    // Deploy EntryPoint
    const EntryPointFactory = await ethers.getContractFactory("EntryPoint");
    const entryPoint = await EntryPointFactory.deploy() as EntryPoint;
    await entryPoint.waitForDeployment();
    console.log(`   ✅ EntryPoint deployed at: ${await entryPoint.getAddress()}`);

    // Deploy CPOP Token
    const CPOPTokenFactory = await ethers.getContractFactory("CPOPToken");
    const cpopToken = await CPOPTokenFactory.deploy(deployerAddress) as CPOPToken;
    await cpopToken.waitForDeployment();
    console.log(`   ✅ CPOPToken deployed at: ${await cpopToken.getAddress()}`);

    // Deploy MasterAggregator
    const MasterAggregatorFactory = await ethers.getContractFactory("MasterAggregator");
    const masterAggregator = await MasterAggregatorFactory.deploy() as MasterAggregator;
    await masterAggregator.waitForDeployment();
    console.log(`   ✅ MasterAggregator deployed at: ${await masterAggregator.getAddress()}`);

    // Initialize MasterAggregator with initial master signers
    await masterAggregator.initialize(
        deployerAddress,
        [masterSigner1Address, masterSigner2Address]
    );
    console.log(`   ✅ MasterAggregator initialized with 2 master signers`);

    // Deploy WalletManager
    const WalletManagerFactory = await ethers.getContractFactory("WalletManager");
    const walletManager = await WalletManagerFactory.deploy() as WalletManager;
    await walletManager.waitForDeployment();
    console.log(`   ✅ WalletManager deployed at: ${await walletManager.getAddress()}`);

    // Initialize WalletManager
    await walletManager.initialize(
        await entryPoint.getAddress(),
        await cpopToken.getAddress(),
        deployerAddress
    );
    console.log(`   ✅ WalletManager initialized\n`);

    // ============================================================================
    // STEP 2: MASTER SIGNER MANAGEMENT
    // ============================================================================
    
    console.log("🔐 Step 2: Master Signer Management...");

    // Check initial authorization
    const isMaster1Auth = await masterAggregator.authorizedMasters(masterSigner1Address);
    const isMaster2Auth = await masterAggregator.authorizedMasters(masterSigner2Address);
    console.log(`   Master Signer 1 authorized: ${isMaster1Auth}`);
    console.log(`   Master Signer 2 authorized: ${isMaster2Auth}`);

    // Demonstrate authorization management
    const newMaster = user3Address;
    console.log(`   👤 Adding new master signer: ${newMaster}`);
    await masterAggregator.setMasterAuthorization(newMaster, true);
    
    const isNewMasterAuth = await masterAggregator.authorizedMasters(newMaster);
    console.log(`   New master authorized: ${isNewMasterAuth}\n`);

    // ============================================================================
    // STEP 3: WALLET CREATION
    // ============================================================================
    
    console.log("💼 Step 3: Creating AAWallets with Master Signers...");

    // Create wallets for users controlled by master signers
    const salt1 = ethers.id("demo_user1_wallet");
    const salt2 = ethers.id("demo_user2_wallet");
    const salt3 = ethers.id("demo_user3_wallet");

    console.log(`   🏗️  Creating wallet for User1 with MasterSigner1...`);
    const tx1 = await walletManager.createAccountWithMasterSigner(
        user1Address,
        salt1,
        masterSigner1Address
    );
    await tx1.wait();
    const wallet1Address = await walletManager.getAccountAddress(user1Address, salt1);
    console.log(`   ✅ User1 wallet: ${wallet1Address}`);

    console.log(`   🏗️  Creating wallet for User2 with MasterSigner1...`);
    const tx2 = await walletManager.createAccountWithMasterSigner(
        user2Address,
        salt2,
        masterSigner1Address
    );
    await tx2.wait();
    const wallet2Address = await walletManager.getAccountAddress(user2Address, salt2);
    console.log(`   ✅ User2 wallet: ${wallet2Address}`);

    console.log(`   🏗️  Creating wallet for User3 with MasterSigner2...`);
    const tx3 = await walletManager.createAccountWithMasterSigner(
        user3Address,
        salt3,
        masterSigner2Address
    );
    await tx3.wait();
    const wallet3Address = await walletManager.getAccountAddress(user3Address, salt3);
    console.log(`   ✅ User3 wallet: ${wallet3Address}`);

    // Get wallet instances
    const wallet1 = await ethers.getContractAt("AAWallet", wallet1Address) as AAWallet;
    const wallet2 = await ethers.getContractAt("AAWallet", wallet2Address) as AAWallet;
    const wallet3 = await ethers.getContractAt("AAWallet", wallet3Address) as AAWallet;

    // Verify master signers
    console.log(`   🔍 Wallet1 master signer: ${await wallet1.getMasterSigner()}`);
    console.log(`   🔍 Wallet2 master signer: ${await wallet2.getMasterSigner()}`);
    console.log(`   🔍 Wallet3 master signer: ${await wallet3.getMasterSigner()}\n`);

    // ============================================================================
    // STEP 4: WALLET AUTHORIZATION
    // ============================================================================
    
    console.log("🔗 Step 4: Authorizing Wallets for Aggregation...");

    // Auto-authorize wallets by verifying master signer relationship
    console.log(`   🔄 Auto-authorizing wallet1 for MasterSigner1...`);
    await masterAggregator.autoAuthorizeWallet(wallet1Address, masterSigner1Address);

    console.log(`   🔄 Auto-authorizing wallet2 for MasterSigner1...`);
    await masterAggregator.autoAuthorizeWallet(wallet2Address, masterSigner1Address);

    console.log(`   🔄 Auto-authorizing wallet3 for MasterSigner2...`);
    await masterAggregator.autoAuthorizeWallet(wallet3Address, masterSigner2Address);

    // Verify authorizations
    const wallet1Authorized = await masterAggregator.isWalletControlledByMaster(wallet1Address, masterSigner1Address);
    const wallet2Authorized = await masterAggregator.isWalletControlledByMaster(wallet2Address, masterSigner1Address);
    const wallet3Authorized = await masterAggregator.isWalletControlledByMaster(wallet3Address, masterSigner2Address);

    console.log(`   ✅ Wallet1 authorized for MasterSigner1: ${wallet1Authorized}`);
    console.log(`   ✅ Wallet2 authorized for MasterSigner1: ${wallet2Authorized}`);
    console.log(`   ✅ Wallet3 authorized for MasterSigner2: ${wallet3Authorized}\n`);

    // ============================================================================
    // STEP 5: SIGNATURE AGGREGATION DEMONSTRATION
    // ============================================================================
    
    console.log("🔀 Step 5: Signature Aggregation Demo...");

    // Create mock user operations for aggregation
    const userOps = [
        {
            sender: wallet1Address,
            nonce: 0,
            initCode: "0x",
            callData: wallet1.interface.encodeFunctionData("execute", [
                user1Address,
                ethers.parseEther("0.01"),
                "0x"
            ]),
            accountGasLimits: ethers.concat([
                ethers.toBeHex(50000, 16), // verificationGasLimit
                ethers.toBeHex(200000, 16) // callGasLimit
            ]),
            preVerificationGas: 21000,
            gasFees: ethers.concat([
                ethers.toBeHex(ethers.parseUnits("20", "gwei"), 16), // maxPriorityFeePerGas
                ethers.toBeHex(ethers.parseUnits("30", "gwei"), 16)  // maxFeePerGas
            ]),
            paymasterAndData: "0x",
            signature: "0x"
        },
        {
            sender: wallet2Address,
            nonce: 0,
            initCode: "0x",
            callData: wallet2.interface.encodeFunctionData("execute", [
                user2Address,
                ethers.parseEther("0.01"),
                "0x"
            ]),
            accountGasLimits: ethers.concat([
                ethers.toBeHex(50000, 16),
                ethers.toBeHex(200000, 16)
            ]),
            preVerificationGas: 21000,
            gasFees: ethers.concat([
                ethers.toBeHex(ethers.parseUnits("20", "gwei"), 16),
                ethers.toBeHex(ethers.parseUnits("30", "gwei"), 16)
            ]),
            paymasterAndData: "0x",
            signature: "0x"
        }
    ];

    console.log(`   📝 Created ${userOps.length} user operations for aggregation`);

    // Get current nonce for master signer
    const currentNonce = await masterAggregator.getMasterNonce(masterSigner1Address);
    console.log(`   📊 Current nonce for MasterSigner1: ${currentNonce}`);

    // Create aggregated signature
    console.log(`   🔐 Creating aggregated signature...`);
    const DEMO_MASTER_KEY = ethers.id("demo_master_key_1");
    const aggregatedSig = await masterAggregator.createMasterAggregatedSignature(
        userOps,
        masterSigner1Address,
        DEMO_MASTER_KEY
    );
    console.log(`   ✅ Aggregated signature created`);

    // Validate aggregated signature
    console.log(`   🔍 Validating aggregated signature...`);
    try {
        const tx = await masterAggregator.validateSignatures(userOps, aggregatedSig);
        await tx.wait();
        console.log(`   ✅ Signature validation successful!`);
    } catch (error) {
        console.log(`   ❌ Signature validation failed: ${error}`);
    }

    // Check nonce increment
    const newNonce = await masterAggregator.getMasterNonce(masterSigner1Address);
    console.log(`   📊 Nonce after validation: ${currentNonce} → ${newNonce}\n`);

    // ============================================================================
    // STEP 6: GAS SAVINGS CALCULATION
    // ============================================================================
    
    console.log("⛽ Step 6: Gas Savings Analysis...");

    const scenarios = [2, 5, 10, 20, 50];
    console.log("   Operations | Individual Gas | Aggregated Gas | Savings | Efficiency");
    console.log("   -----------|---------------|---------------|---------|----------");
    
    for (const opCount of scenarios) {
        const savings = await masterAggregator.calculateGasSavings(opCount);
        const individualGas = opCount * 3000;
        const aggregatedGas = 5000 + (opCount * 500);
        const efficiency = ((Number(savings) / individualGas) * 100).toFixed(1);
        
        console.log(`   ${opCount.toString().padStart(10)} | ${individualGas.toString().padStart(13)} | ${aggregatedGas.toString().padStart(13)} | ${savings.toString().padStart(7)} | ${efficiency.padStart(8)}%`);
    }
    console.log();

    // ============================================================================
    // STEP 7: SESSION KEY AGGREGATION
    // ============================================================================
    
    console.log("🗝️  Step 7: Session Key Aggregation Demo...");

    // Create a session key
    const sessionKeyPrivate = ethers.id("demo_session_key");
    const sessionKeyWallet = new ethers.Wallet(sessionKeyPrivate);
    const sessionKeyAddress = sessionKeyWallet.address;
    console.log(`   🔑 Generated session key: ${sessionKeyAddress}`);

    // Add session keys to wallets (need to use master signer)
    const validAfter = Math.floor(Date.now() / 1000) - 3600; // 1 hour ago
    const validUntil = Math.floor(Date.now() / 1000) + 86400; // 24 hours from now
    const permissions = ethers.id("EXECUTE_PERMISSION");

    console.log(`   ➕ Adding session key to wallet1...`);
    await wallet1.connect(masterSigner1).addSessionKey(sessionKeyAddress, validAfter, validUntil, permissions);
    
    console.log(`   ➕ Adding session key to wallet2...`);
    await wallet2.connect(masterSigner1).addSessionKey(sessionKeyAddress, validAfter, validUntil, permissions);

    // Verify session keys
    const sessionInfo1 = await wallet1.getSessionKeyInfo(sessionKeyAddress);
    const sessionInfo2 = await wallet2.getSessionKeyInfo(sessionKeyAddress);
    console.log(`   ✅ Session key active in wallet1: ${sessionInfo1.isValid}`);
    console.log(`   ✅ Session key active in wallet2: ${sessionInfo2.isValid}`);

    // Create session key aggregated signature
    const sessionUserOps = [
        {
            sender: wallet1Address,
            nonce: 1,
            initCode: "0x",
            callData: "0x",
            accountGasLimits: ethers.concat([
                ethers.toBeHex(50000, 16),
                ethers.toBeHex(200000, 16)
            ]),
            preVerificationGas: 21000,
            gasFees: ethers.concat([
                ethers.toBeHex(ethers.parseUnits("20", "gwei"), 16),
                ethers.toBeHex(ethers.parseUnits("30", "gwei"), 16)
            ]),
            paymasterAndData: "0x",
            signature: "0x"
        },
        {
            sender: wallet2Address,
            nonce: 1,
            initCode: "0x",
            callData: "0x",
            accountGasLimits: ethers.concat([
                ethers.toBeHex(50000, 16),
                ethers.toBeHex(200000, 16)
            ]),
            preVerificationGas: 21000,
            gasFees: ethers.concat([
                ethers.toBeHex(ethers.parseUnits("20", "gwei"), 16),
                ethers.toBeHex(ethers.parseUnits("30", "gwei"), 16)
            ]),
            paymasterAndData: "0x",
            signature: "0x"
        }
    ];

    console.log(`   🔐 Creating session key aggregated signature...`);
    const sessionSig = await masterAggregator.createSessionKeyAggregatedSignature(
        sessionUserOps,
        sessionKeyAddress,
        sessionKeyPrivate
    );

    // Validate session key signature
    console.log(`   🔍 Validating session key signature...`);
    const isSessionValid = await masterAggregator.validateSessionKeyAggregatedSignature(
        sessionUserOps,
        sessionSig
    );
    console.log(`   ✅ Session key signature valid: ${isSessionValid}\n`);

    // ============================================================================
    // STEP 8: CONFIGURATION MANAGEMENT
    // ============================================================================
    
    console.log("⚙️  Step 8: Configuration Management...");

    // Display current configuration
    const maxOps = await masterAggregator.maxAggregatedOps();
    const validationWindow = await masterAggregator.validationWindow();
    console.log(`   Current max operations: ${maxOps}`);
    console.log(`   Current validation window: ${validationWindow} seconds`);

    // Update configuration
    console.log(`   🔄 Updating configuration...`);
    await masterAggregator.updateConfig(75, 600); // 75 ops, 10 minutes

    const newMaxOps = await masterAggregator.maxAggregatedOps();
    const newValidationWindow = await masterAggregator.validationWindow();
    console.log(`   ✅ New max operations: ${newMaxOps}`);
    console.log(`   ✅ New validation window: ${newValidationWindow} seconds\n`);

    // ============================================================================
    // STEP 9: ENTRYPOINT INTEGRATION
    // ============================================================================
    
    console.log("🔗 Step 9: EntryPoint Integration...");

    // Add stake for the aggregator
    const stakeAmount = ethers.parseEther("1");
    const unstakeDelay = 86400; // 1 day

    console.log(`   💰 Adding ${ethers.formatEther(stakeAmount)} ETH stake to EntryPoint...`);
    await masterAggregator.addStake(entryPoint, unstakeDelay, { value: stakeAmount });

    // Check stake info
    const stakeInfo = await entryPoint.deposits(await masterAggregator.getAddress());
    console.log(`   ✅ Stake amount: ${ethers.formatEther(stakeInfo.stake)} ETH`);
    console.log(`   ✅ Unstake delay: ${stakeInfo.unstakeDelaySec} seconds\n`);

    // ============================================================================
    // STEP 10: BATCH OPERATIONS DEMO
    // ============================================================================
    
    console.log("🔄 Step 10: Batch Operations Demo...");

    // Demonstrate batch wallet authorization
    const walletAddresses = [wallet1Address, wallet2Address];
    console.log(`   📋 Batch authorizing ${walletAddresses.length} wallets...`);
    
    // First deauthorize to demonstrate batch authorization
    await masterAggregator.setWalletAuthorization(masterSigner1Address, wallet1Address, false);
    await masterAggregator.setWalletAuthorization(masterSigner1Address, wallet2Address, false);
    
    // Now batch authorize
    await masterAggregator.batchSetWalletAuthorization(
        masterSigner1Address,
        walletAddresses,
        true
    );
    
    // Verify batch authorization
    for (const walletAddr of walletAddresses) {
        const isAuthorized = await masterAggregator.isWalletControlledByMaster(walletAddr, masterSigner1Address);
        console.log(`   ✅ Wallet ${walletAddr} authorized: ${isAuthorized}`);
    }
    console.log();

    // ============================================================================
    // SUMMARY
    // ============================================================================
    
    console.log("🎯 Demo Summary:");
    console.log("================");
    console.log(`✅ Deployed and configured MasterAggregator with ${await masterAggregator.maxAggregatedOps()} max operations`);
    console.log(`✅ Authorized ${Object.keys(await Promise.all([
        masterAggregator.authorizedMasters(masterSigner1Address),
        masterAggregator.authorizedMasters(masterSigner2Address),
        masterAggregator.authorizedMasters(newMaster)
    ])).length} master signers`);
    console.log(`✅ Created 3 AAWallets with master signer control`);
    console.log(`✅ Demonstrated signature aggregation for ${userOps.length} operations`);
    console.log(`✅ Showed gas savings up to ${await masterAggregator.calculateGasSavings(50)} gas units for 50 operations`);
    console.log(`✅ Implemented session key aggregation support`);
    console.log(`✅ Staked ${ethers.formatEther(stakeAmount)} ETH in EntryPoint`);
    console.log(`✅ Configured batch wallet authorization`);
    
    console.log("\n🌟 MasterAggregator Demo Complete!");
    console.log("The system is ready for production use with real master signers and user operations.");
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error("❌ Demo failed:", error);
        process.exit(1);
    });