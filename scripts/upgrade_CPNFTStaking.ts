import { ethers } from "hardhat";

async function main() {
  console.log("Upgrading CPNFTStaking...");

  // 获取当前的代理合约地址
  const proxyAddress = "0x..."; // 替换为实际的代理合约地址
  
  // 获取合约工厂
  const CPNFTStakingFactory = await ethers.getContractFactory("CPNFTStaking");
  
  // 部署新的实现合约
  console.log("Deploying new implementation contract...");
  const newImplementation = await CPNFTStakingFactory.deploy();
  await newImplementation.deployed();
  console.log("New implementation deployed to:", newImplementation.address);

  // 获取代理合约实例
  const proxy = await ethers.getContractAt("ERC1967Proxy", proxyAddress);
  const stakingContract = CPNFTStakingFactory.attach(proxyAddress);
  
  // 获取当前版本
  const currentVersion = await stakingContract.getVersion();
  console.log("Current version:", currentVersion);

  // 执行升级
  console.log("Upgrading proxy to new implementation...");
  const upgradeTx = await stakingContract.upgradeTo(newImplementation.address);
  await upgradeTx.wait();
  console.log("Upgrade transaction hash:", upgradeTx.hash);

  // 验证升级
  const newVersion = await stakingContract.getVersion();
  console.log("New version:", newVersion);
  
  // 验证核心功能
  const owner = await stakingContract.owner();
  const cpnft = await stakingContract.cpnft();
  const cpopToken = await stakingContract.cpopToken();
  const accountManager = await stakingContract.accountManager();
  const defaultMasterSigner = await stakingContract.defaultMasterSigner();
  
  console.log("Contract state after upgrade:");
  console.log("- Owner:", owner);
  console.log("- CPNFT:", cpnft);
  console.log("- CPOPToken:", cpopToken);
  console.log("- AccountManager:", accountManager);
  console.log("- Default Master Signer:", defaultMasterSigner);

  // 测试关键功能
  console.log("\nTesting key functions...");
  
  // 测试收益地址计算
  const testTokenId = 1;
  const testOwner = "0x1234567890123456789012345678901234567890"; // 测试地址
  try {
    const beneficiary = await stakingContract.calculateBeneficiaryAddress(testTokenId, testOwner);
    console.log("✓ calculateBeneficiaryAddress works");
    console.log("  Test beneficiary:", beneficiary);
  } catch (error) {
    console.log("✗ calculateBeneficiaryAddress failed:", error.message);
  }

  // 测试masterSigner计算
  try {
    const masterSigner = await stakingContract.getCalculatedMasterSigner(testTokenId);
    console.log("✓ getCalculatedMasterSigner works");
    console.log("  Test masterSigner:", masterSigner);
  } catch (error) {
    console.log("✗ getCalculatedMasterSigner failed:", error.message);
  }

  console.log("\nUpgrade completed successfully!");
  
  const upgradeInfo = {
    oldImplementation: "0x...", // 可以从代理合约中获取
    newImplementation: newImplementation.address,
    proxy: proxyAddress,
    oldVersion: currentVersion,
    newVersion: newVersion,
    upgradeTime: new Date().toISOString(),
    network: (await ethers.provider.getNetwork()).name,
    chainId: (await ethers.provider.getNetwork()).chainId,
    transactionHash: upgradeTx.hash
  };

  console.log("Upgrade info:", JSON.stringify(upgradeInfo, null, 2));
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
