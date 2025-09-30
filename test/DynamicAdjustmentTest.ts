import { expect } from "chai";
import { ethers } from "hardhat";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { 
    CPNFT, 
    Staking, 
    StakingConfig, 
    StakingReader,
    ICPOPToken,
    IAccountManager
} from "../typechain";

describe("Dynamic Adjustment Impact on Manual Withdrawal", function () {
    let owner: SignerWithAddress;
    let user1: SignerWithAddress;
    let user2: SignerWithAddress;
    
    let cpnft: CPNFT;
    let staking: Staking;
    let config: StakingConfig;
    let reader: StakingReader;
    let cpopToken: ICPOPToken;
    let accountManager: IAccountManager;

    beforeEach(async function () {
        [owner, user1, user2] = await ethers.getSigners();
        
        // Deploy contracts (simplified for testing)
        const CPNFTFactory = await ethers.getContractFactory("CPNFT");
        cpnft = await CPNFTFactory.deploy();
        
        const StakingConfigFactory = await ethers.getContractFactory("StakingConfig");
        config = await StakingConfigFactory.deploy();
        
        const StakingFactory = await ethers.getContractFactory("Staking");
        staking = await StakingFactory.deploy();
        
        const ReaderFactory = await ethers.getContractFactory("StakingReader");
        reader = await ReaderFactory.deploy(staking.address, config.address, cpnft.address);
        
        // Initialize staking contract
        await staking.initialize(
            cpnft.address,
            cpopToken.address,
            accountManager.address,
            config.address,
            owner.address
        );
        
        // Set staking contract reference in config
        await config.setStakingContract(staking.address);
    });

    it("Should apply historical quarterly adjustments to reward calculation", async function () {
        // Mint A-level NFT to user1
        await cpnft.connect(owner).mint(user1.address, 3); // A level = 3
        
        // Stake the NFT
        await staking.connect(user1).stake(1);
        
        // Fast forward 5 days
        await ethers.provider.send("evm_increaseTime", [5 * 86400]); // 5 days
        await ethers.provider.send("evm_mine", []);
        
        // Execute quarterly adjustment (reduce to 0.9x)
        await config.announceQuarterlyAdjustment(9000); // 0.9x
        await ethers.provider.send("evm_increaseTime", [7 * 86400]); // Wait for announcement period
        await config.executeQuarterlyAdjustment();
        
        // Fast forward 5 more days (total 10 days)
        await ethers.provider.send("evm_increaseTime", [5 * 86400]); // 5 more days
        await ethers.provider.send("evm_mine", []);
        
        // Check daily reward breakdown for day 0 (before adjustment)
        const day0Breakdown = await reader.getDailyRewardBreakdown(1, 0);
        expect(day0Breakdown.quarterlyMultiplier).to.equal(10000); // 1.0x (no adjustment yet)
        
        // Check daily reward breakdown for day 5 (after adjustment)
        const day5Breakdown = await reader.getDailyRewardBreakdown(1, 5);
        expect(day5Breakdown.quarterlyMultiplier).to.equal(9000); // 0.9x (adjusted)
        
        // Calculate total rewards - should reflect historical adjustments
        const pendingRewards = await staking.calculatePendingRewards(1);
        expect(pendingRewards).to.be.greaterThan(0);
        
        // The first 5 days should use 1.0x multiplier, next 5 days should use 0.9x
        const baseReward = await config.getDailyReward(3); // A level = 3, 15 CPP
        const expectedRewards = (baseReward * 5 * 10000 / 10000) + (baseReward * 5 * 9000 / 10000);
        expect(pendingRewards).to.be.closeTo(expectedRewards, ethers.utils.parseEther("0.1"));
    });

    it("Should handle dynamic balance adjustments correctly", async function () {
        // Mint multiple A-level NFTs to create high staking ratio
        for (let i = 0; i < 100; i++) {
            await cpnft.connect(owner).mint(user1.address, 3); // A level = 3
        }
        
        // Stake 80 NFTs to create high staking ratio (>60%)
        const tokenIds = [];
        for (let i = 1; i <= 80; i++) {
            tokenIds.push(i);
        }
        await staking.connect(user1).batchStake(tokenIds);
        
        // Record historical adjustment (this will capture the high staking ratio)
        await staking.recordHistoricalAdjustment();
        
        // Fast forward 5 days
        await ethers.provider.send("evm_increaseTime", [5 * 86400]); // 5 days
        await ethers.provider.send("evm_mine", []);
        
        // Unstake 50 NFTs to reduce staking ratio
        await staking.connect(user1).batchUnstake(tokenIds.slice(0, 50));
        
        // Record another historical adjustment (this will capture the lower staking ratio)
        await staking.recordHistoricalAdjustment();
        
        // Fast forward 5 more days
        await ethers.provider.send("evm_increaseTime", [5 * 86400]); // 5 more days
        await ethers.provider.send("evm_mine", []);
        
        // Check dynamic multipliers for different days
        const day0Breakdown = await reader.getDailyRewardBreakdown(1, 0);
        const day5Breakdown = await reader.getDailyRewardBreakdown(1, 5);
        const day10Breakdown = await reader.getDailyRewardBreakdown(1, 10);
        
        // Day 0 and 5 should have high staking penalty (0.8x)
        expect(day0Breakdown.dynamicMultiplier).to.equal(8000); // 0.8x
        expect(day5Breakdown.dynamicMultiplier).to.equal(8000); // 0.8x
        
        // Day 10 should have normal multiplier (1.0x) after unstaking
        expect(day10Breakdown.dynamicMultiplier).to.equal(10000); // 1.0x
    });

    it("Should demonstrate the example scenario from requirements", async function () {
        // Scenario: User stakes A-level NFT for 91 days with dynamic adjustments
        
        // Mint A-level NFT to user1
        await cpnft.connect(owner).mint(user1.address, 3); // A level = 3
        
        // Stake the NFT
        await staking.connect(user1).stake(1);
        
        // Day 1-10: Normal conditions
        await ethers.provider.send("evm_increaseTime", [10 * 86400]); // 10 days
        await ethers.provider.send("evm_mine", []);
        
        // Day 11: High staking ratio triggers penalty (simulate by staking more NFTs)
        for (let i = 0; i < 200; i++) {
            await cpnft.connect(owner).mint(user2.address, 3); // A level = 3
        }
        
        const moreTokenIds = [];
        for (let i = 2; i <= 200; i++) {
            moreTokenIds.push(i);
        }
        await staking.connect(user2).batchStake(moreTokenIds);
        
        // Record historical adjustment to capture high staking ratio
        await staking.recordHistoricalAdjustment();
        
        // Day 11-45: High staking ratio penalty applies
        await ethers.provider.send("evm_increaseTime", [35 * 86400]); // 35 more days (total 45)
        await ethers.provider.send("evm_mine", []);
        
        // Day 46: Reduce staking ratio back to normal
        await staking.connect(user2).batchUnstake(moreTokenIds);
        
        // Record historical adjustment to capture normal staking ratio
        await staking.recordHistoricalAdjustment();
        
        // Day 46-91: Normal conditions resume
        await ethers.provider.send("evm_increaseTime", [46 * 86400]); // 46 more days (total 91)
        await ethers.provider.send("evm_mine", []);
        
        // Get comprehensive reward calculation
        const comprehensive = await reader.getComprehensiveRewardCalculation(1);
        
        expect(comprehensive.totalDays).to.equal(91);
        
        // Check specific days match the example
        // Day 1-10: Normal reward (15 CPP)
        const day5Breakdown = await reader.getDailyRewardBreakdown(1, 5);
        expect(day5Breakdown.finalReward).to.equal(ethers.utils.parseEther("15")); // 15 CPP
        
        // Day 11-45: High staking penalty (15 * 0.8 = 12 CPP)
        const day25Breakdown = await reader.getDailyRewardBreakdown(1, 25);
        expect(day25Breakdown.finalReward).to.equal(ethers.utils.parseEther("12")); // 12 CPP
        
        // Day 46-91: Normal reward again (15 CPP)
        const day70Breakdown = await reader.getDailyRewardBreakdown(1, 70);
        expect(day70Breakdown.finalReward).to.equal(ethers.utils.parseEther("15")); // 15 CPP
        
        // Calculate total rewards
        const totalRewards = await staking.calculatePendingRewards(1);
        
        // Expected calculation:
        // Days 1-10: 10 * 15 = 150 CPP
        // Days 11-45: 35 * 12 = 420 CPP  
        // Days 46-91: 46 * 15 = 690 CPP
        // Total: 150 + 420 + 690 = 1260 CPP
        const expectedTotal = ethers.utils.parseEther("1260");
        expect(totalRewards).to.be.closeTo(expectedTotal, ethers.utils.parseEther("1"));
    });

    it("Should handle quarterly adjustment during staking period", async function () {
        // Mint A-level NFT to user1
        await cpnft.connect(owner).mint(user1.address, 3); // A level = 3
        
        // Stake the NFT
        await staking.connect(user1).stake(1);
        
        // Fast forward 20 days
        await ethers.provider.send("evm_increaseTime", [20 * 86400]); // 20 days
        await ethers.provider.send("evm_mine", []);
        
        // Announce and execute quarterly adjustment (increase to 1.1x)
        await config.announceQuarterlyAdjustment(11000); // 1.1x
        await ethers.provider.send("evm_increaseTime", [7 * 86400]); // Wait for announcement period
        await config.executeQuarterlyAdjustment();
        
        // Fast forward 20 more days
        await ethers.provider.send("evm_increaseTime", [20 * 86400]); // 20 more days
        await ethers.provider.send("evm_mine", []);
        
        // Check rewards before and after adjustment
        const day10Breakdown = await reader.getDailyRewardBreakdown(1, 10); // Before adjustment
        const day30Breakdown = await reader.getDailyRewardBreakdown(1, 30); // After adjustment
        
        expect(day10Breakdown.quarterlyMultiplier).to.equal(10000); // 1.0x
        expect(day30Breakdown.quarterlyMultiplier).to.equal(11000); // 1.1x
        
        // Calculate total rewards
        const totalRewards = await staking.calculatePendingRewards(1);
        
        // Expected: 40 days total
        // First 27 days (including announcement period): 1.0x multiplier
        // Last 13 days: 1.1x multiplier
        const baseReward = await config.getDailyReward(3); // 15 CPP
        const expectedRewards = (baseReward * 27 * 10000 / 10000) + (baseReward * 13 * 11000 / 10000);
        expect(totalRewards).to.be.closeTo(expectedRewards, ethers.utils.parseEther("0.1"));
    });
});
