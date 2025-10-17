import { HardhatRuntimeEnvironment } from "hardhat/types";
import { DeployFunction } from "hardhat-deploy/types";

const deployBatchTransfer: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const { deployments, getNamedAccounts } = hre;
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  console.log("\n=".repeat(60));
  console.log("Deploying BatchTransfer Contract");
  console.log("=".repeat(60));

  const batchTransfer = await deploy("BatchTransfer", {
    from: deployer,
    args: [],
    log: true,
    waitConfirmations: 1,
  });

  console.log(`\n✅ BatchTransfer deployed at: ${batchTransfer.address}`);
  console.log(`   Deployer: ${deployer}`);
  console.log(`   Transaction Hash: ${batchTransfer.transactionHash}`);

  // Verification info
  if (hre.network.name !== "hardhat" && hre.network.name !== "localhost") {
    console.log("\n📝 To verify the contract, run:");
    console.log(`   npx hardhat verify --network ${hre.network.name} ${batchTransfer.address}`);
  }

  console.log("\n" + "=".repeat(60));
  console.log("✅ Deployment Complete!");
  console.log("=".repeat(60));
  
  return true;
};

export default deployBatchTransfer;
deployBatchTransfer.id = "deploy_batch_transfer";
deployBatchTransfer.tags = ["BatchTransfer"];

