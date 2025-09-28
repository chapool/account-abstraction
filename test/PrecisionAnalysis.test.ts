import { expect } from "chai";
import { ethers, upgrades } from "hardhat";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";

describe("Precision Analysis Test", function () {
    let owner: SignerWithAddress;
    let user1: SignerWithAddress;

    // Contracts
    let cpnft: any;
    let cpopToken: any;
    let accountManager: any;
    let stakingConfig: any;
    let staking: any;
    let stakingReader: any;

    const MINTER_ROLE = 2;

    beforeEach(async function () {
        [owner, user1] = await ethers.getSigners();

        // Deploy contracts
        const CPNFT = await ethers.getContractFactory("CPNFT");
        cpnft = await upgrades.deployProxy(CPNFT, ["TestNFT", "TNFT", "https://api.test.com/"]);
        await cpnft.deployed();

        const MockCPOPToken = await ethers.getContractFactory("CPOPToken");
        cpopToken = await MockCPOPToken.deploy(owner.address, ethers.utils.parseEther("1000000"));
        await cpopToken.deployed();

        const MockAccountManager = await ethers.getContractFactory("AccountManager");
        accountManager = await MockAccountManager.deploy();
        await accountManager.deployed();

        const StakingConfig = await ethers.getContractFactory("StakingConfig");
        stakingConfig = await StakingConfig.deploy();
        await stakingConfig.deployed();

        const Staking = await ethers.getContractFactory("Staking");
        staking = await upgrades.deployProxy(Staking, [
            cpnft.address,
            cpopToken.address,
            accountManager.address,
            stakingConfig.address,
            owner.address
        ]);
        await staking.deployed();

        const StakingReader = await ethers.getContractFactory("StakingReader");
        stakingReader = await StakingReader.deploy(
            staking.address,
            stakingConfig.address,
            cpnft.address
        );
        await stakingReader.deployed();

        // Setup
        await cpnft.connect(owner).setStakingContract(staking.address);
        await cpopToken.connect(owner).grantRole(staking.address, MINTER_ROLE);
        
        // Set large supplies to avoid dynamic multiplier effects
        await cpnft.connect(owner).setLevelSupply(5, 10000); // SS level
    });

    describe("Precision Analysis", function () {
        it("Should analyze precision loss sources", async function () {
            console.log("\n=== 精度损失分析 ===");
            
            // Check quarterly multiplier
            const quarterlyMultiplier = await stakingConfig.getQuarterlyMultiplier();
            console.log(`季度调整乘数: ${quarterlyMultiplier} (${quarterlyMultiplier/10000}x)`);
            
            // Check dynamic multiplier
            const levelStats = await stakingReader.getLevelStats();
            const dynamicMultiplier = levelStats.dynamicMultipliers[5]; // SS level index
            console.log(`动态乘数: ${dynamicMultiplier} (${dynamicMultiplier/10000}x)`);
            
            // Calculate combined multiplier
            const combinedMultiplier = quarterlyMultiplier * dynamicMultiplier / 10000;
            console.log(`组合乘数: ${combinedMultiplier} (${combinedMultiplier/10000}x)`);
            
            // Test a simple calculation
            const baseReward = 50; // SS level base reward
            const adjustedReward = baseReward * quarterlyMultiplier * dynamicMultiplier / 10000 / 10000;
            console.log(`基础收益: ${baseReward} CPP`);
            console.log(`调整后收益: ${adjustedReward} CPP`);
            console.log(`调整比例: ${(adjustedReward/baseReward*100).toFixed(2)}%`);
            
            // Mint and stake SS-level NFT for 90 days
            const tokenId = await cpnft.connect(owner).callStatic.mint(user1.address, 5);
            await cpnft.connect(owner).mint(user1.address, 5);
            await staking.connect(user1).stake(tokenId);
            
            // Wait 90 days
            await ethers.provider.send("evm_increaseTime", [90 * 24 * 60 * 60]);
            await ethers.provider.send("evm_mine", []);
            
            // Claim rewards
            await staking.connect(user1).claimRewards(tokenId);
            
            // Get total rewards
            const stakeInfo = await stakingReader.getStakeDetails(tokenId);
            const totalRewards = stakeInfo.totalRewards.toNumber();
            
            // Calculate expected rewards step by step
            const expectedBase = 90 * 50; // 90 days * 50 CPP
            const expectedWithQuarterly = expectedBase * quarterlyMultiplier / 10000;
            const expectedWithDynamic = expectedWithQuarterly * dynamicMultiplier / 10000;
            const expectedWithContinuous = expectedWithDynamic * 1.30; // 30% continuous bonus
            
            console.log(`\n=== 收益计算分解 ===`);
            console.log(`预期基础收益: ${expectedBase} CPP`);
            console.log(`季度调整后: ${expectedWithQuarterly} CPP`);
            console.log(`动态调整后: ${expectedWithDynamic} CPP`);
            console.log(`连续质押奖励后: ${expectedWithContinuous} CPP`);
            console.log(`实际总收益: ${totalRewards} CPP`);
            
            const difference = Math.abs(totalRewards - expectedWithContinuous);
            const percentageDiff = (difference / expectedWithContinuous * 100).toFixed(2);
            console.log(`差异: ${difference} CPP (${percentageDiff}%)`);
            
            // Clean up
            await staking.connect(user1).unstake(tokenId);
        });
    });
});
