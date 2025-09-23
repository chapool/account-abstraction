import { ethers } from "hardhat";
import { ContractFactory } from "ethers";

async function main() {
  console.log("Deploying CPNFTStaking...");

  // 获取合约工厂
  const CPNFTStakingFactory = await ethers.getContractFactory("CPNFTStaking");
  
  // 部署实现合约
  console.log("Deploying implementation contract...");
  const implementation = await CPNFTStakingFactory.deploy();
  await implementation.deployed();
  console.log("Implementation deployed to:", implementation.address);

  // 部署UUPS代理
  console.log("Deploying UUPS proxy...");
  const UUPSProxyFactory = await ethers.getContractFactory("ERC1967Proxy");
  
  // 准备初始化数据
  const cpnftAddress = "0x..."; // 替换为实际的CPNFT合约地址
  const cpopTokenAddress = "0x..."; // 替换为实际的CPOPToken合约地址
  const accountManagerAddress = "0x..."; // 替换为实际的AccountManager合约地址
  const defaultMasterSignerAddress = "0x..."; // 替换为默认主签名者地址
  const ownerAddress = "0x..."; // 替换为合约拥有者地址
  
  const initData = CPNFTStakingFactory.interface.encodeFunctionData("initialize", [
    cpnftAddress,
    cpopTokenAddress,
    accountManagerAddress,
    defaultMasterSignerAddress,
    ownerAddress
  ]);

  const proxy = await UUPSProxyFactory.deploy(implementation.address, initData);
  await proxy.deployed();
  console.log("UUPS Proxy deployed to:", proxy.address);

  // 获取代理合约实例
  const stakingContract = CPNFTStakingFactory.attach(proxy.address);
  
  // 验证初始化
  const owner = await stakingContract.owner();
  const cpnft = await stakingContract.cpnft();
  const cpopToken = await stakingContract.cpopToken();
  const accountManager = await stakingContract.accountManager();
  const defaultMasterSigner = await stakingContract.defaultMasterSigner();
  const version = await stakingContract.getVersion();
  
  console.log("Contract initialized successfully:");
  console.log("- Owner:", owner);
  console.log("- CPNFT:", cpnft);
  console.log("- CPOPToken:", cpopToken);
  console.log("- AccountManager:", accountManager);
  console.log("- Default Master Signer:", defaultMasterSigner);
  console.log("- Version:", version);

  // 保存部署信息
  const deploymentInfo = {
    implementation: implementation.address,
    proxy: proxy.address,
    cpnft: cpnftAddress,
    cpopToken: cpopTokenAddress,
    accountManager: accountManagerAddress,
    defaultMasterSigner: defaultMasterSignerAddress,
    owner: ownerAddress,
    version: version,
    deploymentTime: new Date().toISOString(),
    network: (await ethers.provider.getNetwork()).name,
    chainId: (await ethers.provider.getNetwork()).chainId
  };

  console.log("\nDeployment completed successfully!");
  console.log("Deployment info:", JSON.stringify(deploymentInfo, null, 2));
  
  // 可以保存到文件
  // const fs = require('fs');
  // fs.writeFileSync('staking-deployment.json', JSON.stringify(deploymentInfo, null, 2));
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
