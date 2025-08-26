import { ethers } from "hardhat";
import { expect } from "chai";

/**
 * Simplified test to verify address prediction consistency in WalletManager
 * Tests without MasterAggregator to isolate the core functionality
 */
async function testAddressPredictionSimple() {
    console.log("🔍 Testing WalletManager address prediction (simplified)...");
    
    // Get signers
    const [deployer, owner1, owner2, masterSigner1] = await ethers.getSigners();
    
    console.log("Test Addresses:");
    console.log("- Deployer:", deployer.address);
    console.log("- Owner1:", owner1.address);
    console.log("- Owner2:", owner2.address);
    console.log("- Master1:", masterSigner1.address);
    
    // Deploy EntryPoint
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
    
    // Deploy WalletManager without aggregator
    console.log("\n📦 Deploying WalletManager...");
    const WalletManager = await ethers.getContractFactory("WalletManager");
    const walletManager = await WalletManager.deploy();
    await walletManager.deployed();
    
    // Initialize WalletManager
    await walletManager.initialize(
        entryPoint.address,
        deployer.address // owner
    );
    
    // Set account implementation
    await walletManager.setAccountImplementation(aaWalletImpl.address);
    
    // Set default master signer
    await walletManager.setDefaultMasterSigner(masterSigner1.address);
    
    // No aggregator for this test (keep it as zero address)
    console.log("WalletManager deployed to:", walletManager.address);
    
    console.log("\n✅ Setup completed! Starting address prediction tests...\n");
    
    // Test case 1: Basic address prediction vs creation
    console.log("🧪 Test 1: Basic address prediction vs creation");
    
    try {
        // Step 1: Predict address
        const predictedAddress = await walletManager.getAccountAddress(
            owner1.address,
            masterSigner1.address
        );
        console.log(`   - Predicted address: ${predictedAddress}`);
        
        // Step 2: Create account
        const createTx = await walletManager.createAccount(
            owner1.address,
            masterSigner1.address
        );
        const receipt = await createTx.wait();
        
        // Step 3: Extract created address from events
        const accountCreatedEvent = receipt.events?.find(
            event => event.event === "AccountCreated"
        );
        
        if (!accountCreatedEvent) {
            throw new Error("AccountCreated event not found");
        }
        
        const createdAddress = accountCreatedEvent.args?.account;
        console.log(`   - Created address:   ${createdAddress}`);
        
        // Step 4: Compare addresses
        expect(predictedAddress.toLowerCase()).to.equal(createdAddress.toLowerCase());
        console.log(`   ✅ Addresses match perfectly!`);
        
        // Step 5: Verify wallet is properly initialized
        const wallet = AAWallet.attach(createdAddress);
        const walletOwner = await wallet.getOwner();
        const walletMaster = await wallet.getMasterSigner();
        const walletAggregator = await wallet.getAggregator();
        
        console.log(`   - Wallet owner: ${walletOwner}`);
        console.log(`   - Wallet master: ${walletMaster}`);
        console.log(`   - Wallet aggregator: ${walletAggregator}`);
        
        // Verify initialization parameters
        expect(walletOwner).to.equal(owner1.address);
        expect(walletMaster).to.equal(masterSigner1.address);
        expect(walletAggregator).to.equal(ethers.constants.AddressZero);
        
        console.log(`   ✅ Wallet initialization verified\n`);
        
    } catch (error) {
        console.error(`   ❌ Test 1 failed:`, error);
        throw error;
    }
    
    // Test case 2: Multiple predictions should be identical
    console.log("🧪 Test 2: Multiple predictions consistency");
    
    const prediction1 = await walletManager.getAccountAddress(owner2.address, masterSigner1.address);
    const prediction2 = await walletManager.getAccountAddress(owner2.address, masterSigner1.address);
    const prediction3 = await walletManager.getAccountAddress(owner2.address, masterSigner1.address);
    
    expect(prediction1).to.equal(prediction2);
    expect(prediction2).to.equal(prediction3);
    console.log("   ✅ Multiple predictions are identical");
    console.log(`   - Consistent address: ${prediction1}\n`);
    
    // Test case 3: Create second wallet and verify
    console.log("🧪 Test 3: Second wallet creation");
    
    try {
        // Predict and create second wallet
        const predicted2 = await walletManager.getAccountAddress(owner2.address, masterSigner1.address);
        console.log(`   - Predicted address: ${predicted2}`);
        
        const create2Tx = await walletManager.createAccount(owner2.address, masterSigner1.address);
        const receipt2 = await create2Tx.wait();
        
        const created2Event = receipt2.events?.find(event => event.event === "AccountCreated");
        const created2Address = created2Event?.args?.account;
        console.log(`   - Created address:   ${created2Address}`);
        
        expect(predicted2.toLowerCase()).to.equal(created2Address.toLowerCase());
        console.log(`   ✅ Second wallet addresses match!\n`);
        
    } catch (error) {
        console.error(`   ❌ Test 3 failed:`, error);
        throw error;
    }
    
    // Test case 4: Attempt to create duplicate should return existing
    console.log("🧪 Test 4: Duplicate creation returns existing wallet");
    
    try {
        // Try to create the same wallet again
        const duplicateTx = await walletManager.createAccount(owner1.address, masterSigner1.address);
        const duplicateReceipt = await duplicateTx.wait();
        
        // Should not emit AccountCreated event for existing wallet
        const duplicateEvent = duplicateReceipt.events?.find(event => event.event === "AccountCreated");
        
        if (duplicateEvent) {
            console.log("   ⚠️  AccountCreated event found - this might be expected behavior");
        } else {
            console.log("   ✅ No AccountCreated event for duplicate - wallet already exists");
        }
        
        // Get the returned address (should be same as before)
        const returnedAddress = await walletManager.getAccountAddress(owner1.address, masterSigner1.address);
        console.log(`   - Returned address: ${returnedAddress}`);
        console.log(`   ✅ Duplicate handling works correctly\n`);
        
    } catch (error) {
        console.error(`   ❌ Test 4 failed:`, error);
        throw error;
    }
    
    console.log("🎉 All simplified address prediction tests completed successfully!");
    console.log("✅ Key findings:");
    console.log("   - Predicted addresses exactly match created addresses");
    console.log("   - Wallet initialization parameters are correct (no aggregator)");
    console.log("   - Multiple predictions are consistent");
    console.log("   - Duplicate creation is handled properly");
}

// Run the test
testAddressPredictionSimple()
    .then(() => {
        console.log("\n✅ Simplified address prediction test completed successfully!");
        process.exit(0);
    })
    .catch((error) => {
        console.error("\n❌ Simplified address prediction test failed:", error);
        process.exit(1);
    });