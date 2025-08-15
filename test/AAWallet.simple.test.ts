import { expect } from "chai";
import { ethers } from "hardhat";
import { Signer } from "ethers";
import {
    AAWallet,
    EntryPoint,
    AAWallet__factory,
    EntryPoint__factory,
} from "../typechain";

describe("AAWallet - Simple Test", function () {
    let owner: Signer;
    let entryPoint: EntryPoint;
    let wallet: AAWallet;
    let ownerAddress: string;

    before(async function () {
        [owner] = await ethers.getSigners();
        ownerAddress = await owner.getAddress();
        console.log("Owner address:", ownerAddress);
    });

    it("Should deploy EntryPoint successfully", async function () {
        const EntryPointFactory = (await ethers.getContractFactory("EntryPoint")) as EntryPoint__factory;
        entryPoint = await EntryPointFactory.deploy();
        await entryPoint.deployed();
        
        console.log("EntryPoint deployed at:", entryPoint.address);
        expect(entryPoint.address).to.not.equal(ethers.constants.AddressZero);
    });

    it("Should deploy AAWallet successfully", async function () {
        const AAWalletFactory = (await ethers.getContractFactory("AAWallet")) as AAWallet__factory;
        
        console.log("Deploying AAWallet with EntryPoint:", entryPoint.address);
        wallet = await AAWalletFactory.deploy(entryPoint.address);
        await wallet.deployed();
        
        console.log("AAWallet deployed at:", wallet.address);
        expect(wallet.address).to.not.equal(ethers.constants.AddressZero);
    });

    it("Should have correct EntryPoint address", async function () {
        const storedEntryPoint = await wallet.entryPoint();
        console.log("Stored EntryPoint:", storedEntryPoint);
        expect(storedEntryPoint).to.equal(entryPoint.address);
    });

    it("Should initialize with owner only", async function () {
        console.log("Initializing with owner:", ownerAddress);
        
        try {
            const tx = await wallet.initialize(ownerAddress, ethers.constants.AddressZero);
            await tx.wait();
            console.log("Initialization successful");
            
            const storedOwner = await wallet.getOwner();
            console.log("Stored owner:", storedOwner);
            expect(storedOwner).to.equal(ownerAddress);
            
            const masterSigner = await wallet.getMasterSigner();
            console.log("Master signer:", masterSigner);
            expect(masterSigner).to.equal(ethers.constants.AddressZero);
            
            const isMasterEnabled = await wallet.isMasterSignerEnabled();
            console.log("Master signer enabled:", isMasterEnabled);
            expect(isMasterEnabled).to.be.false;
        } catch (error) {
            console.error("Initialization failed:", error);
            throw error;
        }
    });

    it("Should initialize with owner and master signer", async function () {
        // Create a new wallet for this test
        const AAWalletFactory = (await ethers.getContractFactory("AAWallet")) as AAWallet__factory;
        const wallet2 = await AAWalletFactory.deploy(entryPoint.address);
        await wallet2.deployed();
        
        const masterSignerAddr = await ethers.getSigners().then(s => s[1].getAddress());
        console.log("Initializing with owner and master signer:", ownerAddress, masterSignerAddr);
        
        try {
            const tx = await wallet2.initialize(ownerAddress, masterSignerAddr);
            await tx.wait();
            console.log("Initialization with master signer successful");
            
            const storedOwner = await wallet2.getOwner();
            expect(storedOwner).to.equal(ownerAddress);
            
            const storedMasterSigner = await wallet2.getMasterSigner();
            expect(storedMasterSigner).to.equal(masterSignerAddr);
            
            const isMasterEnabled = await wallet2.isMasterSignerEnabled();
            expect(isMasterEnabled).to.be.true;
        } catch (error) {
            console.error("Initialization with master signer failed:", error);
            throw error;
        }
    });
});