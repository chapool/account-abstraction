import { ethers } from "hardhat";

async function checkEntryPoint() {
  const ENTRY_POINT_ADDRESS = "0x0000000071727de22e5e9d8baf0edac6f37da032"; // v0.7.0
  
  console.log("🔍 Checking EntryPoint interface on Sepolia...\n");
  
  const [deployer] = await ethers.getSigners();
  console.log(`📋 Checking from: ${deployer.address}\n`);
  
  // Try to interact with EntryPoint
  try {
    const entryPointCode = await ethers.provider.getCode(ENTRY_POINT_ADDRESS);
    console.log(`✅ EntryPoint exists at ${ENTRY_POINT_ADDRESS}`);
    console.log(`   Bytecode length: ${entryPointCode.length} bytes\n`);
    
    // Try to get the EntryPoint factory
    const EntryPoint = await ethers.getContractFactory("EntryPoint");
    const entryPoint = EntryPoint.attach(ENTRY_POINT_ADDRESS);
    
    // Test some basic calls
    console.log("📞 Testing EntryPoint interface calls:");
    
    // Check if it's a contract
    const balance = await ethers.provider.getBalance(ENTRY_POINT_ADDRESS);
    console.log(`   Balance: ${ethers.utils.formatEther(balance)} ETH`);
    
    // Try to call supportsInterface (ERC165)
    const erc165Interface = new ethers.utils.Interface([
      "function supportsInterface(bytes4 interfaceId) view returns (bool)"
    ]);
    
    const erc165Contract = new ethers.Contract(ENTRY_POINT_ADDRESS, erc165Interface, ethers.provider);
    
    // Test IEntryPoint interface ID
    const IEntryPointInterfaceId = "0x1313998b"; // This might need to be calculated
    try {
      const supportsIEntryPoint = await erc165Contract.supportsInterface(IEntryPointInterfaceId);
      console.log(`   Supports IEntryPoint (${IEntryPointInterfaceId}): ${supportsIEntryPoint}`);
    } catch (error) {
      console.log(`   ❌ Error checking IEntryPoint interface: ${error.message}`);
    }
    
    // Try to call a simple EntryPoint method
    try {
      const deposits = await entryPoint.balanceOf(deployer.address);
      console.log(`   Deposit balance: ${ethers.utils.formatEther(deposits)} ETH`);
    } catch (error) {
      console.log(`   ❌ Error calling balanceOf: ${error.message}`);
    }
    
  } catch (error) {
    console.error("❌ Error checking EntryPoint:", error);
  }
}

checkEntryPoint()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error("Check failed:", error);
    process.exit(1);
  });