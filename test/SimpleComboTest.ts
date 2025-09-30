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

describe("Simple Combo Bonus Test", function () {
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
        
        // Initialize CPNFT
        await cpnft.initialize("Test CPNFT", "TCPNFT", "https://api.test.com/metadata/");
        
        const StakingConfigFactory = await ethers.getContractFactory("StakingConfig");
        config = await StakingConfigFactory.deploy();
        
        const StakingFactory = await ethers.getContractFactory("Staking");
        staking = await StakingFactory.deploy();
        
        const ReaderFactory = await ethers.getContractFactory("StakingReader");
        reader = await ReaderFactory.deploy(staking.address, config.address, cpnft.address);
        
        // Deploy required contracts
        const CPOPTokenFactory = await ethers.getContractFactory("CPOPToken");
        cpopToken = await CPOPTokenFactory.deploy(owner.address, ethers.utils.parseEther("1000000")); // 1M tokens to owner
        
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
        
        // Set staking contract in CPNFT
        await cpnft.setStakingContract(staking.address);
    });

    it("Should test combo bonus calculation", async function () {
        // Mint 5 A-level NFTs to user1
        for (let i = 0; i < 5; i++) {
            await cpnft.connect(owner).mint(user1.address, 3); // A level = 3
        }
        
        // Check combo thresholds from config
        const thresholds = await config.getComboThresholds();
        const bonuses = await config.getComboBonuses();
        
        console.log("Combo thresholds:", thresholds);
        console.log("Combo bonuses:", bonuses);
        
        // Should be [3, 5, 10] for thresholds
        expect(thresholds[0]).to.equal(3);
        expect(thresholds[1]).to.equal(5);
        expect(thresholds[2]).to.equal(10);
        
        // Should be [500, 1000, 2000] for bonuses (5%, 10%, 20%)
        expect(bonuses[0]).to.equal(500);
        expect(bonuses[1]).to.equal(1000);
        expect(bonuses[2]).to.equal(2000);
    });

    it("Should test combo status structure", async function () {
        // Mint 3 A-level NFTs to user1
        for (let i = 0; i < 3; i++) {
            await cpnft.connect(owner).mint(user1.address, 3); // A level = 3
        }
        
        // Stake all 3 NFTs
        await staking.connect(user1).batchStake([1, 2, 3]);
        
        // Check combo status
        const comboStatus = await reader.getComboStatus(user1.address, 3);
        
        console.log("Combo status:", {
            level: comboStatus.level_,
            count: comboStatus.count,
            effectiveFrom: comboStatus.effectiveFrom,
            bonus: comboStatus.bonus,
            isPending: comboStatus.isPending,
            timeUntilEffective: comboStatus.timeUntilEffective
        });
        
        // Should have 3 NFTs staked
        expect(comboStatus.count).to.equal(3);
        
        // Should get 5% bonus for 3 NFTs
        expect(comboStatus.bonus).to.equal(500);
        
        // Should be pending because we just staked (next-day effect)
        expect(comboStatus.isPending).to.be.true;
        expect(comboStatus.timeUntilEffective.toNumber()).to.be.greaterThan(0);
    });
});
