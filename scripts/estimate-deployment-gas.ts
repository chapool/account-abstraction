import { ethers } from "hardhat";
import * as dotenv from "dotenv";

dotenv.config({ path: '.env.sepolia' });

async function estimateDeploymentGas() {
  console.log("⛽ Estimating CPOP deployment gas costs...\n");

  const [deployer] = await ethers.getSigners();
  const provider = ethers.provider;
  
  // Get current gas price
  const gasPrice = await provider.getGasPrice();
  const gasPriceGwei = ethers.utils.formatUnits(gasPrice, "gwei");
  console.log(`📊 Current Gas Price: ${gasPriceGwei} Gwei\n`);

  let totalGasEstimate = 0;
  let totalCostEth = 0;

  try {
    const contracts = [
      { name: "EntryPoint", args: [] },
      { name: "CPOPToken", args: [] },
      { name: "GasPriceOracle", args: [] },
      { name: "AAccount", args: [] },
    ];

    console.log("📋 Gas Estimates:");
    console.log("=" .repeat(50));

    for (const contract of contracts) {
      const factory = await ethers.getContractFactory(contract.name);
      const deployTx = factory.getDeployTransaction(...contract.args);
      
      const gasEstimate = await provider.estimateGas(deployTx);
      const costEth = gasPrice.mul(gasEstimate);
      const costEthFormatted = ethers.utils.formatEther(costEth);
      
      totalGasEstimate += gasEstimate.toNumber();
      totalCostEth += parseFloat(costEthFormatted);
      
      console.log(`${contract.name.padEnd(20)} ${gasEstimate.toLocaleString().padStart(8)} gas  ~${costEthFormatted.substring(0, 8)} ETH`);
    }

    // Estimate additional contracts that need constructor args
    const additionalEstimate = 2000000; // Conservative estimate for remaining contracts
    const additionalCostEth = parseFloat(ethers.utils.formatEther(gasPrice.mul(additionalEstimate)));
    
    totalGasEstimate += additionalEstimate;
    totalCostEth += additionalCostEth;

    console.log("=" .repeat(50));
    console.log(`${'Total (estimated)'.padEnd(20)} ${totalGasEstimate.toLocaleString().padStart(8)} gas  ~${totalCostEth.toFixed(6)} ETH`);
    console.log("=" .repeat(50));

    // Add safety margin
    const safetyMargin = 1.5; // 50% safety margin
    const recommendedBalance = totalCostEth * safetyMargin;
    
    console.log(`\n💡 Recommendations:`);
    console.log(`   Estimated cost: ~${totalCostEth.toFixed(6)} ETH`);
    console.log(`   Recommended balance: ~${recommendedBalance.toFixed(6)} ETH (with 50% margin)`);
    
    const currentBalance = await deployer.getBalance();
    const currentBalanceEth = parseFloat(ethers.utils.formatEther(currentBalance));
    
    console.log(`   Current balance: ${currentBalanceEth.toFixed(6)} ETH`);
    
    if (currentBalanceEth >= recommendedBalance) {
      console.log(`   ✅ Sufficient balance for deployment`);
    } else {
      console.log(`   ⚠️  Need ${(recommendedBalance - currentBalanceEth).toFixed(6)} more ETH`);
    }

  } catch (error) {
    console.error("❌ Gas estimation failed:", error);
    throw error;
  }
}

estimateDeploymentGas()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });