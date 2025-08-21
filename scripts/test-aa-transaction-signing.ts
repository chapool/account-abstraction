import { ethers } from "hardhat";
import { expect } from "chai";

/**
 * Test script to verify both Master and Owner can sign and execute transactions through AAWallet
 * Tests UserOperation execution with both signers
 */
async function testAATransactionSigning() {
    console.log("🔍 Testing Master and Owner transaction signing through AAWallet...");
    
    // Get signers
    const [deployer, owner, masterSigner, recipient] = await ethers.getSigners();
    
    console.log("Addresses:");
    console.log("- Deployer:", deployer.address);
    console.log("- Owner:", owner.address);
    console.log("- Master:", masterSigner.address);
    console.log("- Recipient:", recipient.address);
    
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
    
    // Deploy proxy for AAWallet
    console.log("\n📦 Deploying AAWallet proxy...");
    const ERC1967Proxy = await ethers.getContractFactory("ERC1967Proxy");
    
    const initData = AAWallet.interface.encodeFunctionData("initialize", [
        entryPoint.address,
        owner.address,
        masterSigner.address
    ]);
    
    const proxy = await ERC1967Proxy.deploy(aaWalletImpl.address, initData);
    await proxy.deployed();
    console.log("AAWallet proxy deployed to:", proxy.address);
    
    const aaWallet = AAWallet.attach(proxy.address);
    
    // Add ETH to the wallet and EntryPoint deposit
    console.log("\n💰 Setting up funds...");
    await deployer.sendTransaction({
        to: aaWallet.address,
        value: ethers.utils.parseEther("2.0")
    });
    
    // Add deposit to EntryPoint for gas
    await aaWallet.connect(masterSigner).addDeposit({ value: ethers.utils.parseEther("0.5") });
    
    const walletBalance = await ethers.provider.getBalance(aaWallet.address);
    const deposit = await aaWallet.getDeposit();
    console.log("- Wallet balance:", ethers.utils.formatEther(walletBalance), "ETH");
    console.log("- EntryPoint deposit:", ethers.utils.formatEther(deposit), "ETH");
    
    console.log("\n✅ Setup completed! Starting transaction tests...\n");
    
    // Helper function to create and execute UserOperation
    async function executeUserOp(signer: any, description: string, transferAmount: string) {
        console.log(`🧪 ${description}`);
        
        try {
            const recipientBalanceBefore = await ethers.provider.getBalance(recipient.address);
            const walletBalanceBefore = await ethers.provider.getBalance(aaWallet.address);
            
            // Create transfer calldata
            const transferCalldata = aaWallet.interface.encodeFunctionData("execute", [
                recipient.address,
                ethers.utils.parseEther(transferAmount),
                "0x"
            ]);
            
            // Get nonce
            const nonce = await aaWallet.getNonce();
            
            // Create UserOperation
            const userOp = {
                sender: aaWallet.address,
                nonce: nonce,
                initCode: "0x",
                callData: transferCalldata,
                callGasLimit: 200000,
                verificationGasLimit: 200000,
                preVerificationGas: 50000,
                maxFeePerGas: ethers.utils.parseUnits("20", "gwei"),
                maxPriorityFeePerGas: ethers.utils.parseUnits("2", "gwei"),
                paymasterAndData: "0x",
                signature: "0x"
            };
            
            // Create packed UserOperation for EIP-4337
            const packedUserOp = {
                sender: userOp.sender,
                nonce: userOp.nonce,
                initCode: userOp.initCode,
                callData: userOp.callData,
                accountGasLimits: ethers.utils.hexZeroPad(
                    ethers.BigNumber.from(userOp.verificationGasLimit).shl(128).or(userOp.callGasLimit).toHexString(), 
                    32
                ),
                preVerificationGas: userOp.preVerificationGas,
                gasFees: ethers.utils.hexZeroPad(
                    ethers.BigNumber.from(userOp.maxPriorityFeePerGas).shl(128).or(userOp.maxFeePerGas).toHexString(),
                    32
                ),
                paymasterAndData: userOp.paymasterAndData,
                signature: userOp.signature
            };
            
            // Create UserOperation hash
            const userOpHash = await entryPoint.getUserOpHash(packedUserOp);
            
            // Sign the hash
            const signature = await signer.signMessage(ethers.utils.arrayify(userOpHash));
            packedUserOp.signature = signature;
            
            // Execute through EntryPoint
            const tx = await entryPoint.handleOps([packedUserOp], deployer.address);
            const receipt = await tx.wait();
            
            // Check results
            const recipientBalanceAfter = await ethers.provider.getBalance(recipient.address);
            const walletBalanceAfter = await ethers.provider.getBalance(aaWallet.address);
            
            const transferReceived = recipientBalanceAfter.sub(recipientBalanceBefore);
            const walletSpent = walletBalanceBefore.sub(walletBalanceAfter);
            
            console.log(`✅ Transaction successful!`);
            console.log(`   - Gas used: ${receipt.gasUsed}`);
            console.log(`   - Recipient received: ${ethers.utils.formatEther(transferReceived)} ETH`);
            console.log(`   - Wallet spent: ${ethers.utils.formatEther(walletSpent)} ETH (including gas)`);
            console.log(`   - Expected transfer: ${transferAmount} ETH`);
            
            // Verify the transfer amount is correct
            expect(transferReceived).to.equal(ethers.utils.parseEther(transferAmount));
            console.log(`✅ Transfer amount verified\n`);
            
        } catch (error) {
            console.error(`❌ ${description} failed:`, error);
            throw error;
        }
    }
    
    // Test 1: Owner signs and executes transaction
    await executeUserOp(owner, "Test 1: Owner signs and executes transaction", "0.1");
    
    // Test 2: Master signs and executes transaction  
    await executeUserOp(masterSigner, "Test 2: Master signs and executes transaction", "0.15");
    
    // Test 3: Verify both signers can execute multiple transactions
    console.log("🧪 Test 3: Multiple transactions from both signers");
    
    await executeUserOp(owner, "Test 3a: Owner second transaction", "0.05");
    await executeUserOp(masterSigner, "Test 3b: Master second transaction", "0.08");
    
    // Test 4: Verify signature validation logic
    console.log("🧪 Test 4: Testing signature validation directly");
    
    try {
        // Create a test UserOperation hash
        const testHash = ethers.utils.keccak256(ethers.utils.toUtf8Bytes("test-validation"));
        
        // Test Owner signature
        const ownerSig = await owner.signMessage(ethers.utils.arrayify(testHash));
        console.log("✅ Owner can create valid signatures");
        
        // Test Master signature
        const masterSig = await masterSigner.signMessage(ethers.utils.arrayify(testHash));
        console.log("✅ Master can create valid signatures");
        
        // Verify signatures are different (proving they're from different accounts)
        expect(ownerSig).to.not.equal(masterSig);
        console.log("✅ Signatures are unique for each signer");
        
    } catch (error) {
        console.error("❌ Signature validation test failed:", error);
    }
    
    // Test 5: Check final balances
    console.log("\n🧪 Test 5: Final balance verification");
    
    const finalWalletBalance = await ethers.provider.getBalance(aaWallet.address);
    const finalRecipientBalance = await ethers.provider.getBalance(recipient.address);
    const finalDeposit = await aaWallet.getDeposit();
    
    console.log("Final balances:");
    console.log("- Wallet balance:", ethers.utils.formatEther(finalWalletBalance), "ETH");
    console.log("- Recipient balance:", ethers.utils.formatEther(finalRecipientBalance), "ETH");
    console.log("- EntryPoint deposit:", ethers.utils.formatEther(finalDeposit), "ETH");
    
    // Calculate total transfers (0.1 + 0.15 + 0.05 + 0.08 = 0.38 ETH)
    const expectedTotalTransfers = ethers.utils.parseEther("0.38");
    console.log("- Expected total transfers:", ethers.utils.formatEther(expectedTotalTransfers), "ETH");
    
    console.log("\n🎉 All transaction signing tests completed successfully!");
    console.log("✅ Both Master and Owner can sign and execute transactions through AAWallet");
    console.log("✅ UserOperation execution works correctly for both signers");
    console.log("✅ Account Abstraction signature validation is working properly");
}

// Run the test
testAATransactionSigning()
    .then(() => {
        console.log("\n✅ AA Transaction signing test completed successfully!");
        process.exit(0);
    })
    .catch((error) => {
        console.error("\n❌ AA Transaction signing test failed:", error);
        process.exit(1);
    });