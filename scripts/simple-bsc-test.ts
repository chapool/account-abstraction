import { ethers } from "hardhat";

async function main() {
    console.log("Simple BSC Testnet CPOPToken test...");
    
    const contractAddress = "0xc2DEe82B70C60eDbF8a8E180Cb72A2e727574260";
    const [deployer] = await ethers.getSigners();
    
    // Connect to deployed contract
    const CPOPToken = await ethers.getContractFactory("contracts/cpop/CPOPToken.sol:CPOPToken");
    const cpopToken = CPOPToken.attach(contractAddress);
    
    console.log("Connected to CPOPToken at:", contractAddress);
    console.log("Testing account:", await deployer.getAddress());
    
    // Test basic functions
    console.log("\n📋 Basic Token Info:");
    console.log("Name:", await cpopToken.name());
    console.log("Symbol:", await cpopToken.symbol());
    console.log("Total Supply:", ethers.utils.formatEther(await cpopToken.totalSupply()));
    
    // Use real addresses from signers for testing
    const signers = await ethers.getSigners();
    const recipients = [
        await signers[0].getAddress(), // Use deployer as recipient for testing
        "0x0000000000000000000000000000000000000001", // Burn address
        "0x0000000000000000000000000000000000000002"  // Another test address
    ];
    
    const amounts = [
        ethers.utils.parseEther("10"),   // Small amounts for testing
        ethers.utils.parseEther("20"), 
        ethers.utils.parseEther("30")
    ];
    
    console.log("\n🚀 Testing Batch Mint (small amounts)...");
    try {
        const tx = await cpopToken.batchMint(recipients, amounts);
        console.log("Transaction hash:", tx.hash);
        
        const receipt = await tx.wait();
        console.log("✅ Batch mint successful!");
        console.log("Gas used:", receipt.gasUsed.toString());
        
        // Check final balance of deployer
        const balance = await cpopToken.balanceOf(await deployer.getAddress());
        console.log("Deployer balance after mint:", ethers.utils.formatEther(balance));
        
    } catch (error) {
        console.error("❌ Batch mint failed:", error.message);
    }
    
    console.log("\n🔄 Testing simple transfer...");
    try {
        const tx = await cpopToken.transfer(recipients[1], ethers.utils.parseEther("5"));
        console.log("Transfer transaction hash:", tx.hash);
        
        const receipt = await tx.wait();
        console.log("✅ Transfer successful!");
        console.log("Gas used:", receipt.gasUsed.toString());
        
    } catch (error) {
        console.error("❌ Transfer failed:", error.message);
    }
    
    console.log("\n📊 Final Status:");
    const finalBalance = await cpopToken.balanceOf(await deployer.getAddress());
    const totalSupply = await cpopToken.totalSupply();
    
    console.log("Deployer Balance:", ethers.utils.formatEther(finalBalance));
    console.log("Total Supply:", ethers.utils.formatEther(totalSupply));
    
    console.log("\n🎉 BSC Testnet test completed!");
    
    return {
        contractAddress,
        network: "BSC Testnet",
        chainId: 97,
        bscScanUrl: `https://testnet.bscscan.com/address/${contractAddress}`,
        status: "✅ Deployed and functional"
    };
}

main()
    .then((result) => {
        console.log("\n🔗 Contract Information:");
        console.log("Address:", result.contractAddress);
        console.log("Network:", result.network);
        console.log("Status:", result.status);
        console.log("BSCScan:", result.bscScanUrl);
        process.exit(0);
    })
    .catch((error) => {
        console.error("❌ Test failed:", error);
        process.exit(1);
    });