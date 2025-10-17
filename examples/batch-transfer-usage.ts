import { ethers } from "hardhat";

/**
 * NFTHelper Usage Examples
 * 
 * This script demonstrates how to use the NFTHelper contract for batch operations
 */

async function main() {
  const [deployer, user1, user2, user3] = await ethers.getSigners();

  console.log("=".repeat(60));
  console.log("NFTHelper Usage Examples");
  console.log("=".repeat(60));

  // ============================================
  // 1. Deploy NFTHelper
  // ============================================
  console.log("\n1. Deploying NFTHelper...");
  const NFTHelperFactory = await ethers.getContractFactory("NFTHelper");
  const nftHelper = await NFTHelperFactory.deploy();
  await nftHelper.deployed();
  console.log("✅ NFTHelper deployed at:", nftHelper.address);

  // ============================================
  // 2. Example: Batch Transfer NFTs
  // ============================================
  console.log("\n2. Batch Transfer NFTs Example");
  console.log("-".repeat(60));
  
  // Assume CPNFT contract address (replace with actual)
  const CPNFT_ADDRESS = "0x..."; // Replace with your CPNFT contract address
  const cpnft = await ethers.getContractAt("CPNFT", CPNFT_ADDRESS);

  console.log("Step 1: Approve NFTHelper to manage your NFTs");
  console.log("Code:");
  console.log(`  await cpnft.setApprovalForAll("${nftHelper.address}", true);`);
  
  // Check if approved
  console.log("\nStep 2: Check approval status");
  console.log("Code:");
  console.log(`  const isApproved = await nftHelper.isNFTApproved("${CPNFT_ADDRESS}", yourAddress);`);

  console.log("\nStep 3: Batch transfer NFTs to multiple recipients");
  console.log("Code:");
  console.log(`  await nftHelper.batchTransferNFT(
    "${CPNFT_ADDRESS}",
    ["${user1.address}", "${user2.address}"],
    [1, 2] // Token IDs
  );`);

  console.log("\nStep 4: Or transfer multiple NFTs to a single recipient");
  console.log("Code:");
  console.log(`  await nftHelper.batchTransferNFTToSingle(
    "${CPNFT_ADDRESS}",
    "${user1.address}",
    [1, 2, 3] // Token IDs
  );`);

  // ============================================
  // 3. Example: Batch Transfer ERC20 Tokens
  // ============================================
  console.log("\n\n3. Batch Transfer ERC20 Tokens Example");
  console.log("-".repeat(60));

  const TOKEN_ADDRESS = "0x..."; // Replace with your token contract address
  
  console.log("Step 1: Approve NFTHelper to spend your tokens");
  console.log("Code:");
  console.log(`  const token = await ethers.getContractAt("IERC20", "${TOKEN_ADDRESS}");
  const totalAmount = ethers.utils.parseEther("1000");
  await token.approve("${nftHelper.address}", totalAmount);`);

  console.log("\nStep 2: Check allowance");
  console.log("Code:");
  console.log(`  const hasAllowance = await nftHelper.hasTokenAllowance(
    "${TOKEN_ADDRESS}",
    yourAddress,
    totalAmount
  );`);

  console.log("\nStep 3: Batch transfer tokens with different amounts");
  console.log("Code:");
  console.log(`  await nftHelper.batchTransferToken(
    "${TOKEN_ADDRESS}",
    ["${user1.address}", "${user2.address}", "${user3.address}"],
    [
      ethers.utils.parseEther("100"),
      ethers.utils.parseEther("200"),
      ethers.utils.parseEther("300")
    ]
  );`);

  console.log("\nStep 4: Or batch transfer equal amounts to each recipient");
  console.log("Code:");
  console.log(`  await nftHelper.batchTransferTokenEqual(
    "${TOKEN_ADDRESS}",
    ["${user1.address}", "${user2.address}", "${user3.address}"],
    ethers.utils.parseEther("100") // Same amount for each
  );`);

  // ============================================
  // 4. Example: Transfer Multiple Token Types
  // ============================================
  console.log("\n\n4. Transfer Multiple Token Types Example");
  console.log("-".repeat(60));

  const TOKEN1_ADDRESS = "0x..."; // USDT
  const TOKEN2_ADDRESS = "0x..."; // USDC
  
  console.log("Transfer multiple ERC20 tokens to one recipient:");
  console.log("Code:");
  console.log(`  // Approve both tokens first
  await token1.approve("${nftHelper.address}", amount1);
  await token2.approve("${nftHelper.address}", amount2);
  
  // Transfer both tokens
  await nftHelper.batchTransferMultipleTokens(
    ["${TOKEN1_ADDRESS}", "${TOKEN2_ADDRESS}"],
    "${user1.address}",
    [
      ethers.utils.parseEther("100"), // USDT amount
      ethers.utils.parseEther("200")  // USDC amount
    ]
  );`);

  // ============================================
  // 5. Example: Combined NFT + Token Transfer
  // ============================================
  console.log("\n\n5. Combined NFT + Token Transfer Example");
  console.log("-".repeat(60));

  console.log("Transfer both NFTs and tokens in a single transaction:");
  console.log("Code:");
  console.log(`  // Approve both NFTs and tokens first
  await cpnft.setApprovalForAll("${nftHelper.address}", true);
  await token.approve("${nftHelper.address}", totalAmount);
  
  // Combined transfer
  await nftHelper.combinedBatchTransfer(
    "${CPNFT_ADDRESS}",              // NFT contract
    ["${user1.address}", "${user2.address}"],  // NFT recipients
    [1, 2],                           // NFT token IDs
    "${TOKEN_ADDRESS}",              // Token contract
    ["${user1.address}", "${user2.address}"],  // Token recipients
    [
      ethers.utils.parseEther("100"),
      ethers.utils.parseEther("200")
    ]
  );`);

  // ============================================
  // 6. Gas Optimization Tips
  // ============================================
  console.log("\n\n6. Gas Optimization Tips");
  console.log("-".repeat(60));
  console.log("✅ Use batchTransferTokenEqual() when sending same amount to everyone");
  console.log("✅ Use combinedBatchTransfer() to save gas when doing both operations");
  console.log("✅ Set approval once, use multiple times (no need to approve each time)");
  console.log("✅ Group transfers by collection to save on approval operations");

  // ============================================
  // 7. Common Use Cases
  // ============================================
  console.log("\n\n7. Common Use Cases");
  console.log("-".repeat(60));
  console.log("📦 Airdrop: Send NFTs/tokens to multiple winners");
  console.log("🎁 Giveaway: Distribute rewards to community members");
  console.log("💼 Portfolio Management: Move assets between wallets");
  console.log("🏆 Tournament Prizes: Distribute winnings to top players");
  console.log("👥 Team Distribution: Send tokens to team members");
  console.log("🔄 Migration: Move assets to new wallets");

  // ============================================
  // 8. Safety Checks
  // ============================================
  console.log("\n\n8. Important Safety Checks");
  console.log("-".repeat(60));
  console.log("⚠️  Always check approval status before batch operations");
  console.log("⚠️  Verify recipient addresses to avoid sending to wrong address");
  console.log("⚠️  Make sure you own all NFTs before attempting transfer");
  console.log("⚠️  Check token balance and allowance before token transfers");
  console.log("⚠️  Test with small amounts first on testnet");

  console.log("\n" + "=".repeat(60));
  console.log("✅ Examples Complete!");
  console.log("=".repeat(60));
}

// Example helper functions
async function exampleBatchNFTTransfer() {
  const [sender] = await ethers.getSigners();
  const nftHelperAddress = "0x..."; // Your deployed NFTHelper address
  const cpnftAddress = "0x...";     // Your CPNFT address
  
  const nftHelper = await ethers.getContractAt("NFTHelper", nftHelperAddress);
  const cpnft = await ethers.getContractAt("CPNFT", cpnftAddress);

  // 1. Approve
  const tx1 = await cpnft.setApprovalForAll(nftHelperAddress, true);
  await tx1.wait();
  console.log("✅ Approved NFTHelper");

  // 2. Batch transfer
  const recipients = [
    "0xRecipient1...",
    "0xRecipient2...",
    "0xRecipient3..."
  ];
  const tokenIds = [1, 2, 3];

  const tx2 = await nftHelper.batchTransferNFT(cpnftAddress, recipients, tokenIds);
  await tx2.wait();
  console.log("✅ Batch transferred NFTs");
}

async function exampleBatchTokenTransfer() {
  const [sender] = await ethers.getSigners();
  const nftHelperAddress = "0x..."; // Your deployed NFTHelper address
  const tokenAddress = "0x...";     // Your token address
  
  const nftHelper = await ethers.getContractAt("NFTHelper", nftHelperAddress);
  const token = await ethers.getContractAt("IERC20", tokenAddress);

  // 1. Approve
  const totalAmount = ethers.utils.parseEther("1000");
  const tx1 = await token.approve(nftHelperAddress, totalAmount);
  await tx1.wait();
  console.log("✅ Approved tokens");

  // 2. Batch transfer
  const recipients = [
    "0xRecipient1...",
    "0xRecipient2...",
    "0xRecipient3..."
  ];
  const amounts = [
    ethers.utils.parseEther("200"),
    ethers.utils.parseEther("300"),
    ethers.utils.parseEther("500")
  ];

  const tx2 = await nftHelper.batchTransferToken(tokenAddress, recipients, amounts);
  await tx2.wait();
  console.log("✅ Batch transferred tokens");
}

// Run the main function
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });

