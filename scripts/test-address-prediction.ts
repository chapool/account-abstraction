import { ethers } from "hardhat";
import { expect } from "chai";

/**
 * Test script to verify address prediction consistency in WalletManager
 * Tests that getAccountAddress returns the same address as createAccount
 */
async function testAddressPrediction() {
    console.log("🔍 Testing WalletManager address prediction consistency...");
    
    // Get signers
    const [deployer, owner1, owner2, masterSigner1, masterSigner2] = await ethers.getSigners();
    
    console.log("Test Addresses:");
    console.log("- Deployer:", deployer.address);
    console.log("- Owner1:", owner1.address);
    console.log("- Owner2:", owner2.address);
    console.log("- Master1:", masterSigner1.address);
    console.log("- Master2:", masterSigner2.address);
    
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
    
    // Deploy MasterAggregator
    console.log("\n📦 Deploying MasterAggregator...");
    const MasterAggregator = await ethers.getContractFactory("MasterAggregator");
    const masterAggregator = await MasterAggregator.deploy();
    await masterAggregator.deployed();
    
    // Initialize MasterAggregator
    await masterAggregator.initialize(
        deployer.address,
        [masterSigner1.address, masterSigner2.address]
    );
    console.log("MasterAggregator deployed to:", masterAggregator.address);
    
    // Deploy WalletManager
    console.log("\n📦 Deploying WalletManager...");
    const WalletManager = await ethers.getContractFactory("WalletManager");
    const walletManager = await WalletManager.deploy();
    await walletManager.deployed();
    
    // Initialize WalletManager
    await walletManager.initialize(
        aaWalletImpl.address,
        entryPoint.address,
        masterSigner1.address, // default master signer
        masterAggregator.address // master aggregator
    );
    console.log("WalletManager deployed to:", walletManager.address);
    
    console.log("\n✅ Setup completed! Starting address prediction tests...\n");
    
    // Test scenarios
    const testCases = [
        {
            name: "Test 1: Owner1 with explicit Master1",
            owner: owner1.address,
            masterSigner: masterSigner1.address
        },
        {
            name: "Test 2: Owner2 with explicit Master2", 
            owner: owner2.address,
            masterSigner: masterSigner2.address
        },
        {
            name: "Test 3: Owner1 with default master (address(0))",
            owner: owner1.address,
            masterSigner: ethers.constants.AddressZero
        },
        {
            name: "Test 4: Owner2 with default master (address(0))",
            owner: owner2.address,
            masterSigner: ethers.constants.AddressZero
        }
    ];
    
    for (const testCase of testCases) {
        console.log(`🧪 ${testCase.name}`);
        
        try {
            // Step 1: Predict address
            const predictedAddress = await walletManager.getAccountAddress(
                testCase.owner,
                testCase.masterSigner
            );
            console.log(`   - Predicted address: ${predictedAddress}`);
            
            // Step 2: Check if wallet already exists
            const existingCode = await ethers.provider.getCode(predictedAddress);
            if (existingCode !== "0x") {
                console.log(`   ⚠️  Wallet already exists, skipping creation`);
                console.log(`   ✅ Prediction matches existing wallet\n`);
                continue;
            }
            
            // Step 3: Create account
            const createTx = await walletManager.createAccount(
                testCase.owner,
                testCase.masterSigner
            );
            const receipt = await createTx.wait();
            
            // Step 4: Extract created address from events
            const accountCreatedEvent = receipt.events?.find(
                event => event.event === "AccountCreated"
            );
            
            if (!accountCreatedEvent) {
                throw new Error("AccountCreated event not found");
            }
            
            const createdAddress = accountCreatedEvent.args?.account;
            console.log(`   - Created address:   ${createdAddress}`);
            
            // Step 5: Compare addresses
            expect(predictedAddress.toLowerCase()).to.equal(createdAddress.toLowerCase());
            console.log(`   ✅ Addresses match perfectly!`);
            
            // Step 6: Verify wallet is properly initialized
            const wallet = AAWallet.attach(createdAddress);
            const walletOwner = await wallet.getOwner();
            const walletMaster = await wallet.getMasterSigner();
            const walletAggregator = await wallet.getAggregator();
            
            console.log(`   - Wallet owner: ${walletOwner}`);
            console.log(`   - Wallet master: ${walletMaster}`);
            console.log(`   - Wallet aggregator: ${walletAggregator}`);
            
            // Verify initialization parameters
            expect(walletOwner).to.equal(testCase.owner);
            
            // Handle default master signer case
            const expectedMaster = testCase.masterSigner === ethers.constants.AddressZero 
                ? masterSigner1.address  // Should use default master
                : testCase.masterSigner;
            expect(walletMaster).to.equal(expectedMaster);
            expect(walletAggregator).to.equal(masterAggregator.address);
            
            console.log(`   ✅ Wallet initialization verified\n`);
            
        } catch (error) {
            console.error(`   ❌ ${testCase.name} failed:`, error);
            throw error;
        }
    }
    
    // Test 5: Multiple predictions for same parameters should be identical
    console.log("🧪 Test 5: Multiple predictions consistency");
    const prediction1 = await walletManager.getAccountAddress(owner1.address, masterSigner2.address);
    const prediction2 = await walletManager.getAccountAddress(owner1.address, masterSigner2.address);
    const prediction3 = await walletManager.getAccountAddress(owner1.address, masterSigner2.address);
    
    expect(prediction1).to.equal(prediction2);
    expect(prediction2).to.equal(prediction3);
    console.log("   ✅ Multiple predictions are identical");
    console.log(`   - Consistent address: ${prediction1}\n`);
    
    // Test 6: Different parameters should produce different addresses
    console.log("🧪 Test 6: Different parameters produce different addresses");
    const addr1 = await walletManager.getAccountAddress(owner1.address, masterSigner1.address);
    const addr2 = await walletManager.getAccountAddress(owner1.address, masterSigner2.address);
    const addr3 = await walletManager.getAccountAddress(owner2.address, masterSigner1.address);
    const addr4 = await walletManager.getAccountAddress(owner2.address, masterSigner2.address);
    
    console.log("   - Owner1 + Master1:", addr1);
    console.log("   - Owner1 + Master2:", addr2);
    console.log("   - Owner2 + Master1:", addr3);
    console.log("   - Owner2 + Master2:", addr4);
    
    // All should be different
    const addresses = [addr1, addr2, addr3, addr4];
    const uniqueAddresses = [...new Set(addresses)];
    expect(uniqueAddresses.length).to.equal(addresses.length);
    console.log("   ✅ All addresses are unique\n");
    
    console.log("🎉 All address prediction tests completed successfully!");
    console.log("✅ Address prediction is working correctly:");
    console.log("   - Predicted addresses match created addresses");
    console.log("   - Wallet initialization parameters are correct");
    console.log("   - Multiple predictions are consistent");
    console.log("   - Different parameters produce different addresses");
}

// Run the test
testAddressPrediction()
    .then(() => {
        console.log("\n✅ Address prediction test completed successfully!");
        process.exit(0);
    })
    .catch((error) => {
        console.error("\n❌ Address prediction test failed:", error);
        process.exit(1);
    });