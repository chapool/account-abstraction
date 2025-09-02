import { ethers } from "hardhat";

async function sendETH() {
  const targetAddress = "0x4367a3091ABb1f79B484b1B7A2D7bcd12316d44e";
  const amount = "50"; // 50 ETH
  
  console.log("💰 Sending ETH on local network...\n");
  
  // Get the first account (has 10000 ETH)
  const [sender] = await ethers.getSigners();
  
  console.log(`📤 Sender: ${sender.address}`);
  console.log(`📥 Recipient: ${targetAddress}`);
  console.log(`💵 Amount: ${amount} ETH\n`);
  
  // Check balances before
  const senderBalanceBefore = await sender.getBalance();
  const recipientBalanceBefore = await ethers.provider.getBalance(targetAddress);
  
  console.log("📊 Balances before transaction:");
  console.log(`   Sender: ${ethers.utils.formatEther(senderBalanceBefore)} ETH`);
  console.log(`   Recipient: ${ethers.utils.formatEther(recipientBalanceBefore)} ETH\n`);
  
  // Send transaction
  console.log("🚀 Sending transaction...");
  const tx = await sender.sendTransaction({
    to: targetAddress,
    value: ethers.utils.parseEther(amount)
  });
  
  console.log(`📝 Transaction hash: ${tx.hash}`);
  
  // Wait for confirmation
  const receipt = await tx.wait();
  console.log(`✅ Transaction confirmed in block ${receipt.blockNumber}\n`);
  
  // Check balances after
  const senderBalanceAfter = await sender.getBalance();
  const recipientBalanceAfter = await ethers.provider.getBalance(targetAddress);
  
  console.log("📊 Balances after transaction:");
  console.log(`   Sender: ${ethers.utils.formatEther(senderBalanceAfter)} ETH`);
  console.log(`   Recipient: ${ethers.utils.formatEther(recipientBalanceAfter)} ETH\n`);
  
  // Calculate gas used
  const gasUsed = receipt.gasUsed;
  const gasPrice = receipt.effectiveGasPrice || tx.gasPrice!;
  const gasCost = gasUsed.mul(gasPrice);
  
  console.log("⛽ Transaction details:");
  console.log(`   Gas used: ${gasUsed.toLocaleString()}`);
  console.log(`   Gas price: ${ethers.utils.formatUnits(gasPrice, 'gwei')} gwei`);
  console.log(`   Gas cost: ${ethers.utils.formatEther(gasCost)} ETH\n`);
  
  console.log("🎉 Transfer completed successfully!");
}

sendETH()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error("❌ Transfer failed:", error);
    process.exit(1);
  });