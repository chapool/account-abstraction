import { expect } from "chai";
import { ethers } from "hardhat";
import { Signer } from "ethers";
import {
    AAWallet,
    EntryPoint,
    WalletManager,
    AAWallet__factory,
    EntryPoint__factory,
    WalletManager__factory,
    ERC1967Proxy__factory,
    TestCounter__factory,
} from "../typechain";

describe("AAWallet - Fixed Unit Tests", function () {
    let owner: Signer;
    let masterSigner: Signer;
    let sessionKey: Signer;
    let user: Signer;
    let attacker: Signer;
    
    let entryPoint: EntryPoint;
    let walletManager: WalletManager;
    let walletImplementation: AAWallet;
    let walletFactory: AAWallet__factory;
    
    let ownerAddress: string;
    let masterSignerAddress: string;
    let sessionKeyAddress: string;
    let userAddress: string;
    let attackerAddress: string;

    before(async function () {
        [owner, masterSigner, sessionKey, user, attacker] = await ethers.getSigners();
        
        ownerAddress = await owner.getAddress();
        masterSignerAddress = await masterSigner.getAddress();
        sessionKeyAddress = await sessionKey.getAddress();
        userAddress = await user.getAddress();
        attackerAddress = await attacker.getAddress();

        console.log("🔧 Setting up fixed AAWallet test environment...");
    });

    beforeEach(async function () {
        // Deploy EntryPoint
        const EntryPointFactory = (await ethers.getContractFactory("EntryPoint")) as EntryPoint__factory;
        entryPoint = await EntryPointFactory.deploy();
        await entryPoint.deployed();

        // Deploy AAWallet implementation
        walletFactory = (await ethers.getContractFactory("AAWallet")) as AAWallet__factory;
        walletImplementation = await walletFactory.deploy(entryPoint.address);
        await walletImplementation.deployed();

        // Deploy WalletManager
        const WalletManagerFactory = (await ethers.getContractFactory("WalletManager")) as WalletManager__factory;
        walletManager = await WalletManagerFactory.deploy();
        await walletManager.deployed();

        // Initialize WalletManager
        await walletManager.initialize(
            entryPoint.address,
            ethers.constants.AddressZero, // CPOP token not needed for basic tests
            ownerAddress
        );

        // Set the wallet implementation
        await walletManager.setAccountImplementation(walletImplementation.address);
    });

    describe("1. Implementation Contract Tests", function () {
        it("Should reject initialization on implementation contract", async function () {
            // This should fail because _disableInitializers() was called in constructor
            await expect(walletImplementation.initialize(ownerAddress, ethers.constants.AddressZero))
                .to.be.reverted;
        });

        it("Should have correct default state in implementation", async function () {
            expect(await walletImplementation.getOwner()).to.equal(ethers.constants.AddressZero);
            expect(await walletImplementation.getMasterSigner()).to.equal(ethers.constants.AddressZero);
            expect(await walletImplementation.isMasterSignerEnabled()).to.be.false;
            expect(await walletImplementation.entryPoint()).to.equal(entryPoint.address);
        });
    });

    describe("2. Proxy Deployment via WalletManager", function () {
        it("Should create wallet without master signer", async function () {
            const salt = ethers.utils.keccak256(ethers.utils.toUtf8Bytes("test1"));
            
            await expect(walletManager.createAccount(ownerAddress, salt))
                .to.emit(walletManager, "AccountCreated");

            const walletAddress = await walletManager.getAccountAddress(ownerAddress, salt);
            const wallet = walletFactory.attach(walletAddress);

            expect(await wallet.getOwner()).to.equal(ownerAddress);
            expect(await wallet.getMasterSigner()).to.equal(ethers.constants.AddressZero);
            expect(await wallet.isMasterSignerEnabled()).to.be.false;
        });

        it("Should create wallet with master signer", async function () {
            const salt = ethers.utils.keccak256(ethers.utils.toUtf8Bytes("test2"));
            
            await expect(walletManager.createAccountWithMasterSigner(ownerAddress, salt, masterSignerAddress))
                .to.emit(walletManager, "AccountCreated");

            const walletAddress = await walletManager.getAccountAddress(ownerAddress, salt);
            const wallet = walletFactory.attach(walletAddress);

            expect(await wallet.getOwner()).to.equal(ownerAddress);
            expect(await wallet.getMasterSigner()).to.equal(masterSignerAddress);
            expect(await wallet.isMasterSignerEnabled()).to.be.true;
        });

        it("Should prevent duplicate account creation", async function () {
            const salt = ethers.utils.keccak256(ethers.utils.toUtf8Bytes("test3"));
            
            await walletManager.createAccount(ownerAddress, salt);
            
            // Second creation should revert or return existing address
            await walletManager.createAccount(ownerAddress, salt); // Should not revert, just return existing
            
            const walletAddress = await walletManager.getAccountAddress(ownerAddress, salt);
            expect(walletAddress).to.not.equal(ethers.constants.AddressZero);
        });
    });

    describe("3. Wallet Functionality Tests", function () {
        let wallet: AAWallet;

        beforeEach(async function () {
            const salt = ethers.utils.keccak256(ethers.utils.toUtf8Bytes("functional"));
            await walletManager.createAccount(ownerAddress, salt);
            const walletAddress = await walletManager.getAccountAddress(ownerAddress, salt);
            wallet = walletFactory.attach(walletAddress);
        });

        it("Should have correct initial state", async function () {
            expect(await wallet.getOwner()).to.equal(ownerAddress);
            expect(await wallet.entryPoint()).to.equal(entryPoint.address);
            expect(await wallet.getNonce()).to.be.a("bigint");
            expect(await wallet.getDeposit()).to.equal(0);
        });

        it("Should manage EntryPoint deposits", async function () {
            const depositAmount = ethers.utils.parseEther("1");
            
            await wallet.addDeposit({ value: depositAmount });
            expect(await wallet.getDeposit()).to.equal(depositAmount);
            
            // Only owner can withdraw
            await expect(wallet.connect(attacker).withdrawDepositTo(attackerAddress, depositAmount))
                .to.be.revertedWith("AAWallet: not owner");
                
            await wallet.connect(owner).withdrawDepositTo(userAddress, depositAmount);
            expect(await wallet.getDeposit()).to.equal(0);
        });

        it("Should receive ETH", async function () {
            const amount = ethers.utils.parseEther("0.5");
            
            await owner.sendTransaction({
                to: wallet.address,
                value: amount
            });
            
            const balance = await ethers.provider.getBalance(wallet.address);
            expect(balance).to.equal(amount);
        });

        it("Should manage aggregator", async function () {
            expect(await wallet.getAggregator()).to.equal(ethers.constants.AddressZero);
            
            await expect(wallet.connect(owner).setAggregator(masterSignerAddress))
                .to.emit(wallet, "AggregatorUpdated")
                .withArgs(ethers.constants.AddressZero, masterSignerAddress);
                
            expect(await wallet.getAggregator()).to.equal(masterSignerAddress);
            
            // Non-owner cannot set aggregator
            await expect(wallet.connect(attacker).setAggregator(attackerAddress))
                .to.be.revertedWith("AAWallet: not owner");
        });

        it("Should enforce execute authorization", async function () {
            // Direct execution should fail (only EntryPoint can call)
            await expect(wallet.connect(owner).execute(
                userAddress,
                ethers.utils.parseEther("0.1"),
                "0x"
            )).to.be.revertedWith("AAWallet: only EntryPoint can execute");
        });
    });

    describe("4. Master Signer Management", function () {
        let wallet: AAWallet;

        beforeEach(async function () {
            const salt = ethers.utils.keccak256(ethers.utils.toUtf8Bytes("master"));
            await walletManager.createAccountWithMasterSigner(ownerAddress, salt, masterSignerAddress);
            const walletAddress = await walletManager.getAccountAddress(ownerAddress, salt);
            wallet = walletFactory.attach(walletAddress);
        });

        it("Should allow master signer to update master signer", async function () {
            const newMasterSigner = userAddress;
            
            await expect(wallet.connect(masterSigner).setMasterSigner(newMasterSigner))
                .to.emit(wallet, "MasterSignerUpdated")
                .withArgs(masterSignerAddress, newMasterSigner);
                
            expect(await wallet.getMasterSigner()).to.equal(newMasterSigner);
        });

        it("Should reject non-master signer trying to update", async function () {
            await expect(wallet.connect(attacker).setMasterSigner(attackerAddress))
                .to.be.revertedWith("AAWallet: not master signer");
        });
    });

    describe("5. Session Key Management", function () {
        let wallet: AAWallet;

        beforeEach(async function () {
            const salt = ethers.utils.keccak256(ethers.utils.toUtf8Bytes("session"));
            await walletManager.createAccount(ownerAddress, salt);
            const walletAddress = await walletManager.getAccountAddress(ownerAddress, salt);
            wallet = walletFactory.attach(walletAddress);
        });

        it("Should add session key with proper constraints", async function () {
            const validAfter = Math.floor(Date.now() / 1000) - 3600; // 1 hour ago
            const validUntil = Math.floor(Date.now() / 1000) + 86400; // 24 hours from now
            const permissions = ethers.utils.id("EXECUTE_PERMISSION");

            await expect(wallet.connect(owner).addSessionKey(
                sessionKeyAddress, 
                validAfter, 
                validUntil, 
                permissions
            )).to.emit(wallet, "SessionKeyAdded")
              .withArgs(sessionKeyAddress, validAfter, validUntil, permissions);

            const [isValid, storedValidAfter, storedValidUntil, storedPermissions] = 
                await wallet.getSessionKeyInfo(sessionKeyAddress);
            
            expect(isValid).to.be.true;
            expect(storedValidAfter).to.equal(validAfter);
            expect(storedValidUntil).to.equal(validUntil);
            expect(storedPermissions).to.equal(permissions);
        });

        it("Should revoke session key", async function () {
            const validAfter = Math.floor(Date.now() / 1000) - 3600;
            const validUntil = Math.floor(Date.now() / 1000) + 86400;
            const permissions = ethers.utils.id("EXECUTE_PERMISSION");

            await wallet.connect(owner).addSessionKey(
                sessionKeyAddress, 
                validAfter, 
                validUntil, 
                permissions
            );

            await expect(wallet.connect(owner).revokeSessionKey(sessionKeyAddress))
                .to.emit(wallet, "SessionKeyRevoked")
                .withArgs(sessionKeyAddress);

            const [isValid] = await wallet.getSessionKeyInfo(sessionKeyAddress);
            expect(isValid).to.be.false;
        });

        it("Should validate session key permissions", async function () {
            const validAfter = Math.floor(Date.now() / 1000) - 3600;
            const validUntil = Math.floor(Date.now() / 1000) + 86400;
            const permissions = ethers.constants.HashZero; // Allow all

            await wallet.connect(owner).addSessionKey(
                sessionKeyAddress, 
                validAfter, 
                validUntil, 
                permissions
            );

            const canExecute = await wallet.canSessionKeyExecute(
                sessionKeyAddress,
                userAddress,
                "0x12345678" // Any selector
            );
            
            expect(canExecute).to.be.true;
        });

        it("Should reject invalid session key parameters", async function () {
            const validAfter = Math.floor(Date.now() / 1000);
            const validUntil = validAfter - 3600; // Invalid: until before after
            const permissions = ethers.utils.id("EXECUTE_PERMISSION");

            await expect(wallet.connect(owner).addSessionKey(
                ethers.constants.AddressZero, 
                validAfter, 
                validAfter + 3600, 
                permissions
            )).to.be.revertedWith("AAWallet: invalid session key");

            await expect(wallet.connect(owner).addSessionKey(
                sessionKeyAddress, 
                validAfter, 
                validUntil, 
                permissions
            )).to.be.revertedWith("AAWallet: invalid time range");
        });

        it("Should not allow non-owner to manage session keys", async function () {
            const validAfter = Math.floor(Date.now() / 1000);
            const validUntil = validAfter + 3600;
            const permissions = ethers.utils.id("EXECUTE_PERMISSION");

            await expect(wallet.connect(attacker).addSessionKey(
                sessionKeyAddress, 
                validAfter, 
                validUntil, 
                permissions
            )).to.be.revertedWith("AAWallet: not owner");

            await expect(wallet.connect(attacker).revokeSessionKey(sessionKeyAddress))
                .to.be.revertedWith("AAWallet: not owner");
        });
    });

    describe("6. Interface Support", function () {
        let wallet: AAWallet;

        beforeEach(async function () {
            const salt = ethers.utils.keccak256(ethers.utils.toUtf8Bytes("interface"));
            await walletManager.createAccount(ownerAddress, salt);
            const walletAddress = await walletManager.getAccountAddress(ownerAddress, salt);
            wallet = walletFactory.attach(walletAddress);
        });

        it("Should support ERC165 interface", async function () {
            const erc165InterfaceId = "0x01ffc9a7";
            expect(await wallet.supportsInterface(erc165InterfaceId)).to.be.true;
        });
    });

    after(async function () {
        console.log("🎯 AAWallet fixed unit testing completed!");
        console.log("📊 Test Coverage Summary:");
        console.log("   ✅ Implementation Contract Behavior");
        console.log("   ✅ Proxy Deployment via WalletManager");
        console.log("   ✅ Basic Wallet Functionality");
        console.log("   ✅ Master Signer Management");
        console.log("   ✅ Session Key Management");
        console.log("   ✅ Interface Support");
        console.log("");
        console.log("🔧 Key Fix: AAWallet must be used through proxy pattern");
        console.log("   - Implementation contract has _disableInitializers() in constructor");
        console.log("   - Use WalletManager.createAccount() for proper deployment");
        console.log("   - Direct initialize() calls on implementation will fail with InvalidInitialization()");
    });
});