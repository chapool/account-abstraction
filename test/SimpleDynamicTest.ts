import { expect } from "chai";
import { ethers } from "hardhat";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { 
    CPNFT, 
    Staking, 
    StakingConfig, 
    StakingReader,
    CPOPToken,
    AccountManager
} from "../typechain";

describe("Simple Dynamic Adjustment Test", function () {
    let owner: SignerWithAddress;
    let user1: SignerWithAddress;
    
    let cpnft: CPNFT;
    let staking: Staking;
    let config: StakingConfig;
    let reader: StakingReader;
    let cpopToken: CPOPToken;
    let accountManager: AccountManager;

    beforeEach(async function () {
        [owner, user1] = await ethers.getSigners();
        
        // Deploy contracts
        const CPNFTFactory = await ethers.getContractFactory("CPNFT");
        cpnft = await CPNFTFactory.deploy();
        await cpnft.initialize("Test CPNFT", "TCPNFT", "https://api.test.com/metadata/");
        
        const StakingConfigFactory = await ethers.getContractFactory("StakingConfig");
        config = await StakingConfigFactory.deploy();
        
        const StakingFactory = await ethers.getContractFactory("Staking");
        staking = await StakingFactory.deploy();
        
        const ReaderFactory = await ethers.getContractFactory("StakingReader");
        reader = await ReaderFactory.deploy(staking.address, config.address, cpnft.address);
        
        // Deploy required contracts
        const CPOPTokenFactory = await ethers.getContractFactory("CPOPToken");
        cpopToken = await CPOPTokenFactory.deploy(owner.address, ethers.utils.parseEther("1000000"));
        
        const AccountManagerFactory = await ethers.getContractFactory("AccountManager");
        accountManager = await AccountManagerFactory.deploy();
        
        // Initialize staking contract
        await staking.initialize(
            cpnft.address,
            cpopToken.address,
            accountManager.address,
            config.address,
            owner.address
        );
        
        await config.setStakingContract(staking.address);
        await cpnft.setStakingContract(staking.address);
    });

    it("Should test historical adjustment recording", async function () {
        // Mint and stake an A-level NFT
        await cpnft.connect(owner).mint(user1.address, 3); // A level = 3
        await staking.connect(user1).stake(1);
        
        // Record historical adjustment
        await staking.recordHistoricalAdjustment();
        
        // Check that historical adjustment was recorded
        const adjustmentCount = await staking.getHistoricalAdjustmentCount();
        expect(adjustmentCount).to.equal(1);
        
        // Get the adjustment details
        const [timestamp, quarterlyMultiplier] = await staking.getHistoricalAdjustment(0);
        expect(timestamp.toNumber()).to.be.greaterThan(0);
        expect(quarterlyMultiplier).to.equal(10000); // 1.0x default
    });

    it("Should test quarterly adjustment announcement and execution", async function () {
        // Announce quarterly adjustment (1.1x)
        await config.announceQuarterlyAdjustment(11000);
        
        // Check announcement was recorded
        const [multiplier, timestamp, announcementTime, isActive] = await config.getLatestQuarterlyAdjustment();
        expect(multiplier).to.equal(11000);
        expect(isActive).to.be.false; // Not yet executed
        
        // Fast forward 7 days for announcement period
        await ethers.provider.send("evm_increaseTime", [7 * 86400 + 1]); // 7 days + 1 second
        await ethers.provider.send("evm_mine", []);
        
        // Execute the adjustment
        await config.executeQuarterlyAdjustment();
        
        // Check that adjustment was executed
        const [executedMultiplier, executedTimestamp, executedAnnouncementTime, executedIsActive] = await config.getLatestQuarterlyAdjustment();
        expect(executedMultiplier).to.equal(11000);
        expect(executedIsActive).to.be.true;
        expect(executedTimestamp.toNumber()).to.be.greaterThan(0);
        
        // Check that historical adjustment was automatically recorded in staking contract
        const adjustmentCount = await staking.getHistoricalAdjustmentCount();
        expect(adjustmentCount).to.equal(1);
    });

    it("Should test daily reward breakdown with historical adjustments", async function () {
        // Mint and stake an A-level NFT
        await cpnft.connect(owner).mint(user1.address, 3); // A level = 3
        await staking.connect(user1).stake(1);
        
        // Fast forward 2 days
        await ethers.provider.send("evm_increaseTime", [2 * 86400]);
        await ethers.provider.send("evm_mine", []);
        
        // Get daily reward breakdown for day 0
        const breakdown = await reader.getDailyRewardBreakdown(1, 0);
        
        console.log("Daily reward breakdown:", {
            baseReward: breakdown.baseReward.toString(),
            decayedReward: breakdown.decayedReward.toString(),
            quarterlyMultiplier: breakdown.quarterlyMultiplier.toString(),
            dynamicMultiplier: breakdown.dynamicMultiplier.toString(),
            finalReward: breakdown.finalReward.toString()
        });
        
        // Should have base reward of 15 CPP for A level
        expect(breakdown.baseReward).to.equal(ethers.utils.parseEther("15"));
        
        // Should not be decayed yet (first day)
        expect(breakdown.decayedReward).to.equal(breakdown.baseReward);
        
        // Should have default quarterly multiplier (1.0x)
        expect(breakdown.quarterlyMultiplier).to.equal(10000);
        
        // Should have dynamic multiplier (could be 0.8x due to high staking ratio)
        expect(breakdown.dynamicMultiplier.toNumber()).to.be.oneOf([8000, 10000]); // 0.8x or 1.0x
        
        // Final reward should be adjusted based on dynamic multiplier
        const expectedReward = breakdown.baseReward.mul(breakdown.dynamicMultiplier).div(10000);
        expect(breakdown.finalReward).to.equal(expectedReward);
    });
});
