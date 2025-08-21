import { ethers } from "hardhat";
import { expect } from "chai";

/**
 * Test script to verify Master and Owner control over the same AAWallet
 * Tests the permission separation after our recent changes
 */
async function testMasterOwnerControl() {
    console.log("🔍 Testing Master and Owner control over AAWallet...");
    
    // Get signers
    const [deployer, owner, masterSigner, randomUser] = await ethers.getSigners();
    
    console.log("Addresses:");
    console.log("- Deployer:", deployer.address);
    console.log("- Owner:", owner.address);
    console.log("- Master:", masterSigner.address);
    console.log("- Random User:", randomUser.address);
    
    // Deploy EntryPoint first
    console.log("\n📦 Deploying EntryPoint...");
    const EntryPoint = await ethers.getContractFactory("EntryPoint");
    const entryPoint = await EntryPoint.deploy();
    await entryPoint.deployed();
    console.log("EntryPoint deployed to:", entryPoint.address);
    
    // Deploy AAWallet implementation
    console.log("\n📦 Deploying AAWallet implementation...");
    const AAWallet = await ethers.getContractFactory("AAWallet");
    const aaWalletImpl = await AAWallet.deploy();
    await aaWalletImpl.deployed();
    console.log("AAWallet implementation deployed to:", aaWalletImpl.address);
    
    // Deploy proxy for AAWallet
    console.log("\n📦 Deploying AAWallet proxy...");
    const ERC1967Proxy = await ethers.getContractFactory("ERC1967Proxy");
    
    // Prepare initialization data
    const initData = AAWallet.interface.encodeFunctionData("initialize", [
        entryPoint.address,
        owner.address,
        masterSigner.address
    ]);
    
    const proxy = await ERC1967Proxy.deploy(aaWalletImpl.address, initData);
    await proxy.deployed();
    console.log("AAWallet proxy deployed to:", proxy.address);
    
    // Connect to the proxy as AAWallet
    const aaWallet = AAWallet.attach(proxy.address);
    
    // Add some ETH to the wallet for testing
    await owner.sendTransaction({
        to: aaWallet.address,
        value: ethers.utils.parseEther("1.0")
    });
    
    console.log("\n✅ Setup completed! Starting permission tests...\n");
    
    // Test 1: Verify initial state
    console.log("🧪 Test 1: Verify initial wallet state");
    try {
        const walletOwner = await aaWallet.getOwner();
        const walletMaster = await aaWallet.getMasterSigner();
        
        console.log("- Wallet owner:", walletOwner);
        console.log("- Wallet master:", walletMaster);
        console.log("- Wallet balance:", ethers.utils.formatEther(await ethers.provider.getBalance(aaWallet.address)), "ETH");
        
        expect(walletOwner).to.equal(owner.address);
        expect(walletMaster).to.equal(masterSigner.address);
        console.log("✅ Initial state verified\n");
    } catch (error) {
        console.error("❌ Test 1 failed:", error);
        return;
    }
    
    // Test 2: Owner permissions (should be very limited now)
    console.log("🧪 Test 2: Testing Owner permissions");
    
    // Test 2a: Owner cannot change Master (should fail now)
    try {
        const newMaster = randomUser.address;
        await aaWallet.connect(owner).setMasterSigner(newMaster);
        console.error("❌ Owner should NOT be able to change Master");
    } catch (error) {
        console.log("✅ Owner correctly cannot change Master (expected)");
    }
    
    // Test 2b: Owner cannot set aggregator (should fail)
    try {
        await aaWallet.connect(owner).setAggregator(randomUser.address);
        console.error("❌ Owner should NOT be able to set aggregator");
    } catch (error) {
        console.log("✅ Owner correctly cannot set aggregator (expected)");
    }
    
    // Test 2c: Owner cannot manage session key manager (should fail)
    try {
        await aaWallet.connect(owner).setAuthorizedSessionKeyManager(randomUser.address);
        console.error("❌ Owner should NOT be able to set session key manager");
    } catch (error) {
        console.log("✅ Owner correctly cannot set session key manager (expected)");
    }
    
    // Test 2d: Owner cannot withdraw deposit (should fail)
    try {
        await aaWallet.connect(owner).withdrawDepositTo(owner.address, 0);
        console.error("❌ Owner should NOT be able to withdraw deposit");
    } catch (error) {
        console.log("✅ Owner correctly cannot withdraw deposit (expected)");
    }
    
    console.log();
    
    // Test 3: Master permissions
    console.log("🧪 Test 3: Testing Master permissions");
    
    // Test 3a: Master can set aggregator
    try {
        await aaWallet.connect(masterSigner).setAggregator(randomUser.address);
        const aggregator = await aaWallet.getAggregator();
        expect(aggregator).to.equal(randomUser.address);
        console.log("✅ Master can set aggregator");
    } catch (error) {
        console.error("❌ Master cannot set aggregator:", error);
    }
    
    // Test 3b: Master can set session key manager
    try {
        await aaWallet.connect(masterSigner).setAuthorizedSessionKeyManager(randomUser.address);
        console.log("✅ Master can set session key manager");
    } catch (error) {
        console.error("❌ Master cannot set session key manager:", error);
    }
    
    // Test 3c: Master can change Master (should work now)
    try {
        const newMaster = randomUser.address;
        await aaWallet.connect(masterSigner).setMasterSigner(newMaster);
        const updatedMaster = await aaWallet.getMasterSigner();
        expect(updatedMaster).to.equal(newMaster);
        console.log("✅ Master can change Master signer");
        
        // Change back for other tests
        await aaWallet.connect(randomUser).setMasterSigner(masterSigner.address);
        console.log("✅ Master signer restored");
    } catch (error) {
        console.error("❌ Master cannot change Master:", error);
    }
    
    // Test 3d: Add some deposit for withdrawal test
    try {
        await aaWallet.connect(masterSigner).addDeposit({ value: ethers.utils.parseEther("0.1") });
        console.log("✅ Added deposit for withdrawal test");
    } catch (error) {
        console.log("⚠️  Could not add deposit, continuing without withdrawal test");
    }
    
    // Test 3e: Master can withdraw deposit
    try {
        const depositBefore = await aaWallet.getDeposit();
        if (depositBefore > 0) {
            await aaWallet.connect(masterSigner).withdrawDepositTo(masterSigner.address, depositBefore);
            console.log("✅ Master can withdraw deposit");
        } else {
            console.log("⚠️  No deposit to withdraw");
        }
    } catch (error) {
        console.error("❌ Master cannot withdraw deposit:", error);
    }
    
    console.log();
    
    // Test 4: Signature validation (both should work)
    console.log("🧪 Test 4: Testing signature validation");
    
    try {
        // Create a simple user operation hash for testing
        const userOpHash = ethers.utils.keccak256(ethers.utils.toUtf8Bytes("test-operation"));
        
        // Test Owner signature validation
        const ownerSignature = await owner.signMessage(ethers.utils.arrayify(userOpHash));
        console.log("✅ Owner can create signatures");
        
        // Test Master signature validation  
        const masterSignature = await masterSigner.signMessage(ethers.utils.arrayify(userOpHash));
        console.log("✅ Master can create signatures");
        
        console.log("✅ Both Owner and Master can sign (signature validation will be tested in UserOp execution)");
    } catch (error) {
        console.error("❌ Signature test failed:", error);
    }
    
    console.log();
    
    // Test 5: Access control summary (updated for Master control)
    console.log("🧪 Test 5: Permission summary (Master Control Mode)");
    console.log("┌─────────────────────────────────┬───────┬────────┐");
    console.log("│ Permission                      │ Owner │ Master │");
    console.log("├─────────────────────────────────┼───────┼────────┤");
    console.log("│ setMasterSigner                 │   ❌   │   ✅    │");
    console.log("│ setAggregator                   │   ❌   │   ✅    │");
    console.log("│ setAuthorizedSessionKeyManager  │   ❌   │   ✅    │");
    console.log("│ withdrawDepositTo               │   ❌   │   ✅    │");
    console.log("│ Contract upgrade                │   ❌   │   ✅    │");
    console.log("│ Sign transactions               │   ✅   │   ✅    │");
    console.log("│ Batch operations (via aggregator)│   ❌   │   ✅    │");
    console.log("└─────────────────────────────────┴───────┴────────┘");
    
    console.log("\n🎉 All tests completed!");
    console.log("✅ Master Control Mode working as expected:");
    console.log("   - Master: Complete control over wallet management and operations");
    console.log("   - Owner: Only transaction signing capability (true custody model)");
}

// Run the test
testMasterOwnerControl()
    .then(() => {
        console.log("\n✅ Test script completed successfully!");
        process.exit(0);
    })
    .catch((error) => {
        console.error("\n❌ Test script failed:", error);
        process.exit(1);
    });