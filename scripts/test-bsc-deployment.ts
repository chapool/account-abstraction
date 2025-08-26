import { ethers } from "hardhat";

async function main() {
    console.log("Testing CPOPToken deployment on BSC Testnet...");
    
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
    console.log("Decimals:", await cpopToken.decimals());
    console.log("Total Supply:", ethers.utils.formatEther(await cpopToken.totalSupply()));
    
    // Test batch mint functionality
    console.log("\n🚀 Testing Batch Mint...");
    const recipients = [
        ethers.utils.getAddress("0x742d35Cc6634C0532925a3b8D489D4E1E5aB09dF"),  // Test address 1  
        ethers.utils.getAddress("0x8ba1f109551bD432803012645Hac136c7cb05cAF"),  // Test address 2
        ethers.utils.getAddress("0x4B20993Bc481177ec7E8f571ceCaE8A9e22C02db")   // Test address 3
    ];
    
    const amounts = [
        ethers.utils.parseEther("100"),
        ethers.utils.parseEther("200"), 
        ethers.utils.parseEther("300")
    ];
    
    try {
        console.log("Executing batchMint for 3 addresses...");
        const tx = await cpopToken.batchMint(recipients, amounts);
        console.log("Transaction hash:", tx.hash);
        
        const receipt = await tx.wait();
        console.log("✅ Batch mint successful!");
        console.log("Gas used:", receipt.gasUsed.toString());
        console.log("Block number:", receipt.blockNumber);
        
        // Check balances
        console.log("\n💰 Recipient Balances:");
        for (let i = 0; i < recipients.length; i++) {
            const balance = await cpopToken.balanceOf(recipients[i]);
            console.log(`${recipients[i]}: ${ethers.utils.formatEther(balance)} CPOP`);
        }
        
    } catch (error) {
        console.error("❌ Batch mint failed:", error);
    }
    
    // Test batch transfer functionality
    console.log("\n🔄 Testing Batch Transfer...");
    const transferRecipients = recipients.slice(0, 2); // Use first 2 addresses
    const transferAmounts = [
        ethers.utils.parseEther("10"),
        ethers.utils.parseEther("20")
    ];
    
    try {
        console.log("Executing batchTransfer...");
        const tx = await cpopToken.batchTransfer(transferRecipients, transferAmounts);
        console.log("Transaction hash:", tx.hash);
        
        const receipt = await tx.wait();
        console.log("✅ Batch transfer successful!");
        console.log("Gas used:", receipt.gasUsed.toString());
        
    } catch (error) {
        console.error("❌ Batch transfer failed:", error);
    }
    
    // Final balances
    console.log("\n📊 Final Token Status:");
    const totalSupply = await cpopToken.totalSupply();
    const adminBalance = await cpopToken.balanceOf(await deployer.getAddress());
    
    console.log("Total Supply:", ethers.utils.formatEther(totalSupply));
    console.log("Admin Balance:", ethers.utils.formatEther(adminBalance));
    
    console.log("\n🎉 BSC Testnet deployment test completed!");
    
    return {
        contractAddress,
        network: "BSC Testnet",
        chainId: 97,
        bscScanUrl: `https://testnet.bscscan.com/address/${contractAddress}`
    };
}

main()
    .then((result) => {
        console.log("\n🔗 BSC Testnet Links:");
        console.log("Contract Address:", result.contractAddress);
        console.log("BSCScan URL:", result.bscScanUrl);
        process.exit(0);
    })
    .catch((error) => {
        console.error("❌ Test failed:", error);
        process.exit(1);
    });