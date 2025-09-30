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

describe("Combo Bonus Next-Day Effect", function () {
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
        
        // For testing, we'll use zero addresses for dependencies we don't need
        // In a real deployment, these would be properly deployed contracts
        const zeroAddress = "0x0000000000000000000000000000000000000000";
        
        // Initialize staking contract
        await staking.initialize(
            cpnft.address,
            zeroAddress, // cpopToken address (not needed for combo bonus tests)
            zeroAddress, // accountManager address (not needed for combo bonus tests)
            config.address,
            owner.address
        );
        
        // Set staking contract reference in config
        await config.setStakingContract(staking.address);
    });

    it("Should implement next-day effect for combo bonus changes", async function () {
        // Mint 6 A-level NFTs to user1
        for (let i = 0; i < 6; i++) {
            await cpnft.connect(owner).mint(user1.address, 3); // A level = 3
        }
        
        // Stake 5 NFTs to get 10% combo bonus
        const tokenIds = [];
        for (let i = 1; i <= 5; i++) {
            tokenIds.push(i);
        }
        
        await staking.connect(user1).batchStake(tokenIds);
        
        // Check initial combo status
        const comboStatus1 = await reader.getComboStatus(user1.address, 3);
        expect(comboStatus1.count).to.equal(5);
        expect(comboStatus1.bonus).to.equal(1000); // 10% bonus
        expect(comboStatus1.isPending).to.be.false;
        
        // Unstake 1 NFT to reduce count to 4
        await staking.connect(user1).unstake(1);
        
        // Check that combo status is now pending for next day
        const comboStatus2 = await reader.getComboStatus(user1.address, 3);
        expect(comboStatus2.count).to.equal(4);
        expect(comboStatus2.bonus).to.equal(0); // No bonus for 4 NFTs
        expect(comboStatus2.isPending).to.be.true;
        expect(comboStatus2.effectiveFrom).to.be.greaterThan(0);
        
        // Check current effective bonus (should still be 10% until next day)
        const effectiveBonus = await staking.getEffectiveComboBonus(user1.address, 3);
        expect(effectiveBonus).to.equal(1000); // Still 10% until next day
        
        // Fast forward time by 1 day + 1 second
        await ethers.provider.send("evm_increaseTime", [86401]); // 1 day + 1 second
        await ethers.provider.send("evm_mine", []);
        
        // Check that bonus is now 0 (no combo bonus for 4 NFTs)
        const effectiveBonusAfter = await staking.getEffectiveComboBonus(user1.address, 3);
        expect(effectiveBonusAfter).to.equal(0);
        
        // Check that combo status is no longer pending
        const comboStatus3 = await reader.getComboStatus(user1.address, 3);
        expect(comboStatus3.isPending).to.be.false;
    });

    it("Should handle different decay states correctly", async function () {
        // Mint 3 B-level NFTs to user2
        for (let i = 0; i < 3; i++) {
            await cpnft.connect(owner).mint(user2.address, 2); // B level = 2
        }
        
        // Stake all 3 NFTs
        await staking.connect(user2).batchStake([6, 7, 8]);
        
        // Check combo bonus (should be 5% for 3 NFTs)
        const effectiveBonus = await staking.getEffectiveComboBonus(user2.address, 2);
        expect(effectiveBonus).to.equal(500); // 5% bonus
        
        // Fast forward 35 days to trigger decay for B level (decay interval is 30 days)
        await ethers.provider.send("evm_increaseTime", [35 * 86400]); // 35 days
        await ethers.provider.send("evm_mine", []);
        
        // Check that combo bonus is still applied to each NFT individually
        // Each NFT should have its own decayed reward with combo bonus applied
        const pendingRewards = await staking.calculatePendingRewards(6);
        expect(pendingRewards).to.be.greaterThan(0);
        
        // The combo bonus should still be 5% even though NFTs are in decay state
        const effectiveBonusAfterDecay = await staking.getEffectiveComboBonus(user2.address, 2);
        expect(effectiveBonusAfterDecay).to.equal(500); // Still 5% bonus
    });

    it("Should handle combo bonus downgrade correctly", async function () {
        // Mint 10 C-level NFTs to user1
        for (let i = 0; i < 10; i++) {
            await cpnft.connect(owner).mint(user1.address, 1); // C level = 1
        }
        
        // Stake all 10 NFTs to get 20% combo bonus
        const tokenIds = [];
        for (let i = 9; i <= 18; i++) {
            tokenIds.push(i);
        }
        await staking.connect(user1).batchStake(tokenIds);
        
        // Check 20% bonus is active
        let effectiveBonus = await staking.getEffectiveComboBonus(user1.address, 1);
        expect(effectiveBonus).to.equal(2000); // 20% bonus
        
        // Unstake 5 NFTs to reduce to 5 NFTs (should get 10% bonus)
        await staking.connect(user1).batchUnstake([9, 10, 11, 12, 13]);
        
        // Check that downgrade is pending for next day
        const comboStatus = await reader.getComboStatus(user1.address, 1);
        expect(comboStatus.count).to.equal(5);
        expect(comboStatus.bonus).to.equal(1000); // 10% bonus
        expect(comboStatus.isPending).to.be.true;
        
        // Current bonus should still be 20%
        effectiveBonus = await staking.getEffectiveComboBonus(user1.address, 1);
        expect(effectiveBonus).to.equal(2000); // Still 20% until next day
        
        // Fast forward 1 day
        await ethers.provider.send("evm_increaseTime", [86401]);
        await ethers.provider.send("evm_mine", []);
        
        // Now should be 10% bonus
        effectiveBonus = await staking.getEffectiveComboBonus(user1.address, 1);
        expect(effectiveBonus).to.equal(1000); // Now 10% bonus
    });
});
