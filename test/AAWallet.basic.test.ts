import { expect } from "chai";
import { ethers } from "hardhat";
import { Signer } from "ethers";
import {
    AAWallet,
    EntryPoint,
    AAWallet__factory,
    EntryPoint__factory,
} from "../typechain";

describe("AAWallet - Basic Functionality Test", function () {
    let owner: Signer;
    let masterSigner: Signer;
    let entryPoint: EntryPoint;
    let walletFactory: AAWallet__factory;
    let ownerAddress: string;
    let masterSignerAddress: string;

    before(async function () {
        [owner, masterSigner] = await ethers.getSigners();
        ownerAddress = await owner.getAddress();
        masterSignerAddress = await masterSigner.getAddress();
        
        console.log("Owner address:", ownerAddress);
        console.log("Master signer address:", masterSignerAddress);
    });

    beforeEach(async function () {
        // Deploy EntryPoint
        const EntryPointFactory = (await ethers.getContractFactory("EntryPoint")) as EntryPoint__factory;
        entryPoint = await EntryPointFactory.deploy();
        await entryPoint.deployed();
        console.log("EntryPoint deployed at:", entryPoint.address);

        // Get wallet factory
        walletFactory = (await ethers.getContractFactory("AAWallet")) as AAWallet__factory;
    });

    it("Should deploy AAWallet with valid EntryPoint", async function () {
        const wallet = await walletFactory.deploy(entryPoint.address);
        await wallet.deployed();
        
        console.log("AAWallet deployed at:", wallet.address);
        expect(wallet.address).to.not.equal(ethers.constants.AddressZero);
        
        // Check EntryPoint is set correctly
        const storedEntryPoint = await wallet.entryPoint();
        expect(storedEntryPoint).to.equal(entryPoint.address);
    });

    it("Should reject deployment with zero EntryPoint", async function () {
        await expect(walletFactory.deploy(ethers.constants.AddressZero))
            .to.be.revertedWith("AAWallet: invalid entryPoint");
    });

    it("Should have correct default state before initialization", async function () {
        const wallet = await walletFactory.deploy(entryPoint.address);
        await wallet.deployed();
        
        // Check default values
        expect(await wallet.getOwner()).to.equal(ethers.constants.AddressZero);
        expect(await wallet.getMasterSigner()).to.equal(ethers.constants.AddressZero);
        expect(await wallet.isMasterSignerEnabled()).to.be.false;
        expect(await wallet.getAggregator()).to.equal(ethers.constants.AddressZero);
    });

    it("Should get nonce from EntryPoint", async function () {
        const wallet = await walletFactory.deploy(entryPoint.address);
        await wallet.deployed();
        
        const nonce = await wallet.getNonce();
        console.log("Initial nonce:", nonce.toString());
        expect(nonce).to.be.a("bigint");
    });

    it("Should support ERC165 interface", async function () {
        const wallet = await walletFactory.deploy(entryPoint.address);
        await wallet.deployed();
        
        // Check ERC165 support
        const erc165InterfaceId = "0x01ffc9a7";
        expect(await wallet.supportsInterface(erc165InterfaceId)).to.be.true;
    });

    it("Should receive ETH", async function () {
        const wallet = await walletFactory.deploy(entryPoint.address);
        await wallet.deployed();
        
        const amount = ethers.utils.parseEther("1");
        
        await owner.sendTransaction({
            to: wallet.address,
            value: amount
        });
        
        const balance = await ethers.provider.getBalance(wallet.address);
        expect(balance).to.equal(amount);
    });

    describe("EntryPoint deposit management", function () {
        let wallet: AAWallet;

        beforeEach(async function () {
            wallet = await walletFactory.deploy(entryPoint.address);
            await wallet.deployed();
        });

        it("Should have zero initial deposit", async function () {
            const deposit = await wallet.getDeposit();
            expect(deposit).to.equal(0);
        });

        it("Should allow adding deposit", async function () {
            const depositAmount = ethers.utils.parseEther("1");
            
            await wallet.addDeposit({ value: depositAmount });
            
            const deposit = await wallet.getDeposit();
            expect(deposit).to.equal(depositAmount);
        });
    });

    describe("Access control before initialization", function () {
        let wallet: AAWallet;

        beforeEach(async function () {
            wallet = await walletFactory.deploy(entryPoint.address);
            await wallet.deployed();
        });

        it("Should reject owner-only functions before initialization", async function () {
            await expect(wallet.connect(owner).setAggregator(masterSignerAddress))
                .to.be.revertedWith("AAWallet: not owner");
        });

        it("Should reject master signer functions before initialization", async function () {
            await expect(wallet.connect(masterSigner).setMasterSigner(ownerAddress))
                .to.be.revertedWith("AAWallet: not master signer");
        });

        it("Should reject execute functions", async function () {
            await expect(wallet.connect(owner).execute(
                ownerAddress,
                ethers.utils.parseEther("0.1"),
                "0x"
            )).to.be.revertedWith("AAWallet: only EntryPoint can execute");
        });
    });

    // Test the issue with initialization specifically
    describe("Initialization Investigation", function () {
        let wallet: AAWallet;

        beforeEach(async function () {
            wallet = await walletFactory.deploy(entryPoint.address);
            await wallet.deployed();
        });

        it("Should investigate initialization error", async function () {
            console.log("About to call initialize...");
            
            try {
                // Try to estimate gas first
                const gasEstimate = await wallet.estimateGas.initialize(ownerAddress);
                console.log("Gas estimate:", gasEstimate.toString());
                
                // Try with specific gas limit
                const tx = await wallet.initialize(ownerAddress, { gasLimit: gasEstimate.mul(2) });
                await tx.wait();
                
                console.log("Initialize succeeded!");
                const storedOwner = await wallet.getOwner();
                expect(storedOwner).to.equal(ownerAddress);
                
            } catch (error: any) {
                console.error("Initialize failed with error:", error.message);
                if (error.data) {
                    console.error("Error data:", error.data);
                }
                throw error;
            }
        });
    });
});