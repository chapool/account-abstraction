import { ethers } from "ethers";

// CPNFT ABI - 只包含我们需要的函数
const ABI = [
  "function owner() view returns (address)",
  "function name() view returns (string)",
  "function symbol() view returns (string)",
  "function getCurrentTokenId() view returns (uint256)"
];

async function main() {
  // CPNFT contract address
  const CPNFT_ADDRESS = "0xEa81A317a4Bc82084359028A207e282F8F503d16";
  
  // Connect to Sepolia
  const provider = new ethers.providers.JsonRpcProvider("https://eth-sepolia.g.alchemy.com/v2/_x4NAgu50ejHAhTH2-gJoRNS4PQv7Tjp");
  
  console.log("Checking CPNFT owner...");
  console.log("Contract address:", CPNFT_ADDRESS);
  console.log("Network:", await provider.getNetwork());

  // Get contract instance
  const cpnft = new ethers.Contract(CPNFT_ADDRESS, ABI, provider);

  try {
    // Get owner
    const owner = await cpnft.owner();
    console.log("\nContract Details:");
    console.log("  Owner:", owner);

    // Get additional info
    const name = await cpnft.name();
    const symbol = await cpnft.symbol();
    console.log("  Name:", name);
    console.log("  Symbol:", symbol);
    
    // Get current token ID counter
    const currentTokenId = await cpnft.getCurrentTokenId();
    console.log("  Current Token ID:", currentTokenId.toString());

  } catch (error) {
    console.error("Error fetching contract details:", error);
  }
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });