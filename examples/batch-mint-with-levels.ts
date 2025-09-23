import { ethers } from "hardhat";

async function main() {
  console.log("CPNFT Batch Mint with Levels Example");

  // 获取合约实例
  const cpnftAddress = "0x..."; // 替换为实际的CPNFT合约地址
  const CPNFTFactory = await ethers.getContractFactory("CPNFT");
  const cpnft = CPNFTFactory.attach(cpnftAddress);

  // 获取用户账户
  const [owner, user1, user2, user3, user4, user5] = await ethers.getSigners();
  console.log("Owner address:", owner.address);

  // 示例1: 批量铸造不同等级的NFT
  console.log("\n=== 示例1: 批量铸造不同等级的NFT ===");
  
  const recipients = [
    user1.address,
    user2.address,
    user3.address,
    user4.address,
    user5.address
  ];
  
  const levels = [
    0, // C级
    1, // B级
    2, // A级
    3, // S级
    4  // SS级
  ];

  try {
    console.log("Batch minting NFTs with different levels...");
    const tx = await cpnft.connect(owner).batchMint(recipients, levels);
    await tx.wait();
    console.log("✓ Batch mint completed");
    console.log("Transaction hash:", tx.hash);
    
    // 验证铸造结果
    for (let i = 0; i < recipients.length; i++) {
      const balance = await cpnft.balanceOf(recipients[i]);
      console.log(`${recipients[i]}: balance = ${balance}, expected level = ${levels[i]}`);
      
      // 获取最新铸造的tokenId
      const tokenId = await cpnft.getCurrentTokenId();
      const actualLevel = await cpnft.getTokenLevel(tokenId - recipients.length + i + 1);
      console.log(`  Token ID: ${tokenId - recipients.length + i + 1}, Level: ${actualLevel}`);
    }
    
  } catch (error) {
    console.log("✗ Batch mint failed:", error.message);
  }

  // 示例2: 铸造单个NFT带等级
  console.log("\n=== 示例2: 铸造单个NFT带等级 ===");
  
  try {
    console.log("Minting single NFT with SSS level...");
    const tx = await cpnft.connect(owner).mint(user1.address, 5); // SSS级 = 5
    await tx.wait();
    console.log("✓ Single mint completed");
    console.log("Transaction hash:", tx.hash);
    
    // 验证铸造结果
    const tokenId = await cpnft.getCurrentTokenId();
    const level = await cpnft.getTokenLevel(tokenId);
    console.log(`Minted token ID: ${tokenId}, Level: ${level}`);
    
  } catch (error) {
    console.log("✗ Single mint failed:", error.message);
  }

  // 示例3: 使用传统方法批量铸造（无等级）
  console.log("\n=== 示例3: 使用传统方法批量铸造 ===");
  
  const legacyRecipients = [user2.address, user3.address];
  
  try {
    console.log("Batch minting NFTs using legacy method...");
    const tx = await cpnft.connect(owner).batchMintLegacy(legacyRecipients);
    await tx.wait();
    console.log("✓ Legacy batch mint completed");
    console.log("Transaction hash:", tx.hash);
    
    // 验证铸造结果
    for (const recipient of legacyRecipients) {
      const balance = await cpnft.balanceOf(recipient);
      console.log(`${recipient}: balance = ${balance}`);
    }
    
  } catch (error) {
    console.log("✗ Legacy batch mint failed:", error.message);
  }

  // 示例4: 铸造后设置等级
  console.log("\n=== 示例4: 铸造后设置等级 ===");
  
  try {
    // 先铸造NFT
    const tx1 = await cpnft.connect(owner).mintLegacy(user4.address);
    await tx1.wait();
    const tokenId = await cpnft.getCurrentTokenId();
    console.log(`Minted token ID: ${tokenId}`);
    
    // 然后设置等级
    const tx2 = await cpnft.connect(owner).setTokenLevel(tokenId, 3); // S级
    await tx2.wait();
    console.log("✓ Token level set");
    
    // 验证等级
    const level = await cpnft.getTokenLevel(tokenId);
    console.log(`Token ID ${tokenId} level: ${level}`);
    
  } catch (error) {
    console.log("✗ Set token level failed:", error.message);
  }

  // 示例5: 错误处理 - 数组长度不匹配
  console.log("\n=== 示例5: 错误处理 - 数组长度不匹配 ===");
  
  try {
    const shortRecipients = [user1.address];
    const longLevels = [0, 1, 2];
    
    console.log("Attempting batch mint with mismatched arrays...");
    await cpnft.connect(owner).batchMint(shortRecipients, longLevels);
    console.log("✗ This should have failed!");
    
  } catch (error) {
    console.log("✓ Expected error caught:", error.message);
  }

  // 示例6: 查询所有NFT信息
  console.log("\n=== 示例6: 查询所有NFT信息 ===");
  
  try {
    const totalMinted = await cpnft.getCurrentTokenId();
    console.log(`Total NFTs minted: ${totalMinted}`);
    
    // 查询前5个NFT的信息
    for (let i = 1; i <= Math.min(5, totalMinted); i++) {
      const owner = await cpnft.ownerOf(i);
      const level = await cpnft.getTokenLevel(i);
      const isStaked = await cpnft.isStaked(i);
      const tokenURI = await cpnft.tokenURI(i);
      
      console.log(`NFT #${i}:`);
      console.log(`  Owner: ${owner}`);
      console.log(`  Level: ${level}`);
      console.log(`  Is Staked: ${isStaked}`);
      console.log(`  Token URI: ${tokenURI}`);
      console.log("");
    }
    
  } catch (error) {
    console.log("✗ Query failed:", error.message);
  }

  console.log("\n=== 所有示例执行完成 ===");
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
