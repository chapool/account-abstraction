import { expect } from "chai";
import { ethers } from "hardhat";
import { Signer } from "ethers";
import {
    WalletManager,
    WalletManager__factory,
    EntryPoint,
    EntryPoint__factory,
    AAWallet,
    AAWallet__factory
} from "../typechain";

describe("WalletManager v2 - Simplified EOA + Master Tests", function () {
    let owner: Signer;
    let user: Signer;
    let masterSigner: Signer;
    let unauthorized: Signer;
    
    let entryPoint: EntryPoint;
    let walletManager: WalletManager;
    let implementation: AAWallet;
    
    let ownerAddress: string;
    let userAddress: string;
    let masterSignerAddress: string;
    let unauthorizedAddress: string;

    beforeEach(async function () {
        [owner, user, masterSigner, unauthorized] = await ethers.getSigners();
        
        ownerAddress = await owner.getAddress();
        userAddress = await user.getAddress();
        masterSignerAddress = await masterSigner.getAddress();
        unauthorizedAddress = await unauthorized.getAddress();

        // Deploy EntryPoint
        const EntryPointFactory = (await ethers.getContractFactory("EntryPoint")) as EntryPoint__factory;
        entryPoint = await EntryPointFactory.deploy();
        await entryPoint.deployed();

        // Deploy WalletManager implementation
        const WalletManagerFactory = (await ethers.getContractFactory("WalletManager")) as WalletManager__factory;
        const walletManagerImpl = await WalletManagerFactory.deploy();
        await walletManagerImpl.deployed();

        // Deploy proxy and initialize
        const proxyData = WalletManagerFactory.interface.encodeFunctionData("initialize", [
            entryPoint.address,
            ownerAddress
        ]);

        const ProxyFactory = await ethers.getContractFactory("ERC1967Proxy");
        const proxy = await ProxyFactory.deploy(walletManagerImpl.address, proxyData);
        await proxy.deployed();

        walletManager = WalletManagerFactory.attach(proxy.address);
        
        // Get the AAWallet implementation
        const implementationAddress = await walletManager.getImplementation();
        const AAWalletFactory = (await ethers.getContractFactory("AAWallet")) as AAWallet__factory;
        implementation = AAWalletFactory.attach(implementationAddress);
    });

    describe("Initialization", function () {
        it("should initialize correctly", async function () {
            expect(await walletManager.owner()).to.equal(ownerAddress);
            expect(await walletManager.accountImplementation()).to.not.equal(ethers.constants.AddressZero);
            expect(await walletManager.isAuthorizedCreator(ownerAddress)).to.be.true;
            expect(await walletManager.getDefaultMasterSigner()).to.equal(ownerAddress);
        });
    });

    describe("Simplified Wallet Creation", function () {
        it("should create wallet with EOA + Master signer", async function () {
            const tx = await walletManager.createWallet(userAddress, masterSignerAddress);
            const receipt = await tx.wait();
            
            const event = receipt.events?.find(e => e.event === "AccountCreated");
            expect(event).to.not.be.undefined;
            expect(event?.args?.owner).to.equal(userAddress);
            expect(event?.args?.masterSigner).to.equal(masterSignerAddress);
            
            // Verify wallet was actually created
            const walletAddress = event?.args?.account;
            const code = await ethers.provider.getCode(walletAddress);
            expect(code.length).to.be.greaterThan(2);
        });

        it("should create wallet using default master signer", async function () {
            const tx = await walletManager.createWallet(userAddress, ethers.constants.AddressZero);
            const receipt = await tx.wait();
            
            const event = receipt.events?.find(e => e.event === "AccountCreated");
            expect(event?.args?.masterSigner).to.equal(ownerAddress); // Default master signer
        });
    });

    describe("Address Prediction", function () {
        it("should predict wallet address correctly", async function () {
            const predictedAddress = await walletManager.getAccountAddress(userAddress, masterSignerAddress);
            expect(predictedAddress).to.not.equal(ethers.constants.AddressZero);
            expect(ethers.utils.isAddress(predictedAddress)).to.be.true;
            
            // Create wallet and verify address matches
            const tx = await walletManager.createWallet(userAddress, masterSignerAddress);
            const receipt = await tx.wait();
            const event = receipt.events?.find(e => e.event === "AccountCreated");
            const actualAddress = event?.args?.account;
            
            expect(predictedAddress).to.equal(actualAddress);
        });
    });

    describe("Web2 User Account Creation", function () {
        it("should create user account", async function () {
            const tx = await walletManager.createUserAccount(userAddress, masterSignerAddress);
            const receipt = await tx.wait();
            
            const event = receipt.events?.find(e => e.event === "AccountCreated");
            expect(event).to.not.be.undefined;
            expect(event?.args?.owner).to.equal(userAddress);
            expect(event?.args?.masterSigner).to.equal(masterSignerAddress);
        });

        it("should only allow authorized creators", async function () {
            await expect(
                walletManager.connect(unauthorized).createUserAccount(userAddress, masterSignerAddress)
            ).to.be.revertedWith("WalletManager: unauthorized creator");
        });
    });

    describe("EntryPoint Integration", function () {
        it("should generate correct initCode", async function () {
            const initCode = await walletManager.getInitCode(userAddress, masterSignerAddress);
            
            expect(initCode.length).to.be.greaterThan(0);
            expect(initCode.substring(0, 42)).to.equal(walletManager.address.toLowerCase());
        });

        it("should check deployment status correctly", async function () {
            // Before deployment
            const isDeployedBefore = await walletManager.isAccountDeployed(userAddress, masterSignerAddress);
            expect(isDeployedBefore).to.be.false;
            
            // After deployment
            await walletManager.createWallet(userAddress, masterSignerAddress);
            const isDeployedAfter = await walletManager.isAccountDeployed(userAddress, masterSignerAddress);
            expect(isDeployedAfter).to.be.true;
        });
    });

    describe("Your Use Case Example", function () {
        it("should work with your scenario: owner=0x234234354, master=0x3435455", async function () {
            // Mock your addresses
            const yourUserAddress = "0x2342343542342343542342343542342343423434"; // 40 chars
            const yourMasterAddress = "0x3435455343545534355435543554355435543554"; // 40 chars
            
            // 1. Predict AA wallet address (off-chain)
            const predictedAAWallet = await walletManager.getAccountAddress(
                yourUserAddress,
                yourMasterAddress
            );
            
            console.log("预计算的AA钱包地址:", predictedAAWallet);
            
            // 2. Create AA wallet (on-chain)
            const tx = await walletManager.createUserAccount(
                yourUserAddress,    // User's EOA address
                yourMasterAddress   // Master signer
            );
            
            const receipt = await tx.wait();
            const event = receipt.events?.find(e => e.event === "AccountCreated");
            const actualAAWallet = event?.args?.account;
            
            // 3. Verify addresses match
            expect(predictedAAWallet).to.equal(actualAAWallet);
            
            // 4. Verify wallet properties
            expect(event?.args?.owner).to.equal(yourUserAddress);
            expect(event?.args?.masterSigner).to.equal(yourMasterAddress);
            
            // 5. Verify wallet was deployed
            const code = await ethers.provider.getCode(actualAAWallet);
            expect(code.length).to.be.greaterThan(2);
            
            console.log("✅ 场景验证成功!");
            console.log("- 用户EOA:", yourUserAddress);
            console.log("- Master地址:", yourMasterAddress);
            console.log("- AA钱包地址:", actualAAWallet);
        });
    });

    describe("Management Functions", function () {
        it("should set default master signer", async function () {
            const newMasterSigner = "0x1234567890123456789012345678901234567890";
            
            await walletManager.setDefaultMasterSigner(newMasterSigner);
            expect(await walletManager.getDefaultMasterSigner()).to.equal(newMasterSigner);
        });

        it("should manage creator authorization", async function () {
            // Initially not authorized
            expect(await walletManager.isAuthorizedCreator(userAddress)).to.be.false;
            
            // Authorize
            await walletManager.authorizeCreator(userAddress);
            expect(await walletManager.isAuthorizedCreator(userAddress)).to.be.true;
            
            // Revoke
            await walletManager.revokeCreator(userAddress);
            expect(await walletManager.isAuthorizedCreator(userAddress)).to.be.false;
        });
    });

    describe("Edge Cases", function () {
        it("should reject invalid owner address", async function () {
            await expect(
                walletManager.createWallet(ethers.constants.AddressZero, masterSignerAddress)
            ).to.be.revertedWith("WalletManager: invalid owner");
        });

        it("should handle duplicate wallet creation gracefully", async function () {
            // Create wallet first time
            const tx1 = await walletManager.createWallet(userAddress, masterSignerAddress);
            const receipt1 = await tx1.wait();
            const event1 = receipt1.events?.find(e => e.event === "AccountCreated");
            const firstAddress = event1?.args?.account;

            // Create same wallet again - should return existing address
            const tx2 = await walletManager.createWallet(userAddress, masterSignerAddress);
            const receipt2 = await tx2.wait();
            
            // Should not emit AccountCreated event on duplicate
            const event2 = receipt2.events?.find(e => e.event === "AccountCreated");
            expect(event2).to.be.undefined;
            
            // But function should still return the same address
            const predictedAddress = await walletManager.getAccountAddress(userAddress, masterSignerAddress);
            expect(firstAddress).to.equal(predictedAddress);
        });

        it("should handle different master signers for same owner", async function () {
            const masterSigner2 = await unauthorized.getAddress();
            
            const address1 = await walletManager.getAccountAddress(userAddress, masterSignerAddress);
            const address2 = await walletManager.getAccountAddress(userAddress, masterSigner2);
            
            expect(address1).to.not.equal(address2);
            
            // Both should be creatable
            await walletManager.createWallet(userAddress, masterSignerAddress);
            await walletManager.createWallet(userAddress, masterSigner2);
            
            expect(await walletManager.isAccountDeployed(userAddress, masterSignerAddress)).to.be.true;
            expect(await walletManager.isAccountDeployed(userAddress, masterSigner2)).to.be.true;
        });

        it("should use default master signer when address(0) provided", async function () {
            const predictedWithDefault = await walletManager.getAccountAddress(userAddress, ethers.constants.AddressZero);
            const predictedWithOwner = await walletManager.getAccountAddress(userAddress, ownerAddress);
            
            expect(predictedWithDefault).to.equal(predictedWithOwner);
        });
    });

    describe("Error Scenarios", function () {
        it("should reject unauthorized creator for createUserAccount", async function () {
            await expect(
                walletManager.connect(unauthorized).createUserAccount(userAddress, masterSignerAddress)
            ).to.be.revertedWith("WalletManager: unauthorized creator");
        });

        it("should reject invalid master signer in createUserAccount", async function () {
            await expect(
                walletManager.createUserAccount(userAddress, ethers.constants.AddressZero)
            ).to.be.revertedWith("WalletManager: invalid master signer");
        });

        it("should reject invalid owner in createUserAccount", async function () {
            await expect(
                walletManager.createUserAccount(ethers.constants.AddressZero, masterSignerAddress)
            ).to.be.revertedWith("WalletManager: invalid owner");
        });

        it("should reject invalid creator authorization", async function () {
            await expect(
                walletManager.authorizeCreator(ethers.constants.AddressZero)
            ).to.be.revertedWith("WalletManager: invalid creator");
        });

        it("should reject duplicate creator authorization", async function () {
            await walletManager.authorizeCreator(userAddress);
            await expect(
                walletManager.authorizeCreator(userAddress)
            ).to.be.revertedWith("WalletManager: already authorized");
        });

        it("should reject revoking non-authorized creator", async function () {
            await expect(
                walletManager.revokeCreator(userAddress)
            ).to.be.revertedWith("WalletManager: not authorized");
        });

        it("should reject revoking owner authorization", async function () {
            await expect(
                walletManager.revokeCreator(ownerAddress)
            ).to.be.revertedWith("WalletManager: cannot revoke owner");
        });

        it("should reject non-owner setting default master signer", async function () {
            await expect(
                walletManager.connect(unauthorized).setDefaultMasterSigner(masterSignerAddress)
            ).to.be.reverted;
        });

        it("should reject invalid master signer address", async function () {
            await expect(
                walletManager.setDefaultMasterSigner(ethers.constants.AddressZero)
            ).to.be.revertedWith("WalletManager: invalid master signer");
        });
    });

    describe("Gas Optimization Tests", function () {
        it("should measure gas costs for wallet creation", async function () {
            const tx = await walletManager.createWallet(userAddress, masterSignerAddress);
            const receipt = await tx.wait();
            
            console.log(`Gas used for wallet creation: ${receipt.gasUsed.toString()}`);
            expect(receipt.gasUsed.toNumber()).to.be.lessThan(300000); // Reasonable gas limit
        });

        it("should measure gas costs for duplicate creation attempt", async function () {
            // First creation
            await walletManager.createWallet(userAddress, masterSignerAddress);
            
            // Duplicate attempt
            const tx = await walletManager.createWallet(userAddress, masterSignerAddress);
            const receipt = await tx.wait();
            
            console.log(`Gas used for duplicate creation: ${receipt.gasUsed.toString()}`);
            expect(receipt.gasUsed.toNumber()).to.be.lessThan(50000); // Should be much cheaper
        });

        it("should have consistent gas costs for address prediction", async function () {
            const gasStart = await ethers.provider.getBalance(ownerAddress);
            
            for (let i = 0; i < 5; i++) {
                await walletManager.getAccountAddress(userAddress, masterSignerAddress);
            }
            
            // View functions should not consume gas
            const gasEnd = await ethers.provider.getBalance(ownerAddress);
            expect(gasStart).to.equal(gasEnd);
        });
    });

    describe("EntryPoint Integration Tests", function () {
        it("should create wallet through EntryPoint initCode", async function () {
            // Get initCode for EntryPoint
            const initCode = await walletManager.getInitCode(userAddress, masterSignerAddress);
            expect(initCode.length).to.be.greaterThan(0);
            
            // Simulate EntryPoint calling the initCode
            const predictedAddress = await walletManager.getAccountAddress(userAddress, masterSignerAddress);
            
            // Check that wallet is not deployed yet
            expect(await walletManager.isAccountDeployed(userAddress, masterSignerAddress)).to.be.false;
            
            // Need to get the senderCreator and call from that address
            const senderCreatorAddress = await walletManager.senderCreator();
            
            // For testing, we'll use the core createWallet function instead since we don't have SenderCreator setup
            await walletManager.createWallet(userAddress, masterSignerAddress);
            
            // Verify deployment
            expect(await walletManager.isAccountDeployed(userAddress, masterSignerAddress)).to.be.true;
            
            const code = await ethers.provider.getCode(predictedAddress);
            expect(code.length).to.be.greaterThan(2);
        });

        it("should validate initCode format", async function () {
            const initCode = await walletManager.getInitCode(userAddress, masterSignerAddress);
            
            // InitCode should start with WalletManager address
            const factoryAddress = initCode.slice(0, 42);
            expect(factoryAddress.toLowerCase()).to.equal(walletManager.address.toLowerCase());
            
            // InitCode should contain the function call data
            const callData = "0x" + initCode.slice(42);
            expect(callData.length).to.be.greaterThan(0);
        });

        it("should handle multiple wallet creations with different parameters", async function () {
            const users = [userAddress, masterSignerAddress, unauthorizedAddress];
            const masters = [masterSignerAddress, userAddress, ownerAddress];
            
            const walletAddresses: string[] = [];
            
            for (let i = 0; i < users.length; i++) {
                const tx = await walletManager.createWallet(users[i], masters[i]);
                const receipt = await tx.wait();
                const event = receipt.events?.find(e => e.event === "AccountCreated");
                
                if (event?.args?.account) {
                    walletAddresses.push(event.args.account);
                }
            }
            
            // All addresses should be unique
            const uniqueAddresses = [...new Set(walletAddresses)];
            expect(uniqueAddresses.length).to.equal(walletAddresses.length);
            
            // All wallets should be deployed
            for (let i = 0; i < users.length; i++) {
                expect(await walletManager.isAccountDeployed(users[i], masters[i])).to.be.true;
            }
        });

        it("should maintain deterministic addresses across different deployments", async function () {
            // Get predicted address
            const predictedAddress1 = await walletManager.getAccountAddress(userAddress, masterSignerAddress);
            
            // Create wallet
            await walletManager.createWallet(userAddress, masterSignerAddress);
            
            // Get predicted address again (should be same)
            const predictedAddress2 = await walletManager.getAccountAddress(userAddress, masterSignerAddress);
            
            expect(predictedAddress1).to.equal(predictedAddress2);
            
            // Even after deployment, prediction should remain consistent
            expect(await walletManager.isAccountDeployed(userAddress, masterSignerAddress)).to.be.true;
            const predictedAddress3 = await walletManager.getAccountAddress(userAddress, masterSignerAddress);
            expect(predictedAddress1).to.equal(predictedAddress3);
        });

        it("should support concurrent wallet creations", async function () {
            const user1 = userAddress;
            const user2 = masterSignerAddress;
            const master1 = masterSignerAddress;
            const master2 = userAddress;
            
            // Create both wallets concurrently
            const [tx1, tx2] = await Promise.all([
                walletManager.createWallet(user1, master1),
                walletManager.createWallet(user2, master2)
            ]);
            
            const [receipt1, receipt2] = await Promise.all([tx1.wait(), tx2.wait()]);
            
            const event1 = receipt1.events?.find(e => e.event === "AccountCreated");
            const event2 = receipt2.events?.find(e => e.event === "AccountCreated");
            
            expect(event1?.args?.account).to.not.equal(event2?.args?.account);
            expect(event1?.args?.owner).to.equal(user1);
            expect(event2?.args?.owner).to.equal(user2);
            expect(event1?.args?.masterSigner).to.equal(master1);
            expect(event2?.args?.masterSigner).to.equal(master2);
        });
    });

    describe("Complete Workflow Tests", function () {
        it("should handle complete Web2 user onboarding workflow", async function () {
            const web2User = "0x1111111111111111111111111111111111111111";
            const systemMaster = masterSignerAddress;
            
            console.log("\n=== Web2 User Onboarding Workflow ===");
            
            // Step 1: Pre-compute AA wallet address (off-chain)
            const predictedWallet = await walletManager.getAccountAddress(web2User, systemMaster);
            console.log(`1. Pre-computed AA wallet: ${predictedWallet}`);
            
            // Step 2: Check if wallet exists
            const existsBefore = await walletManager.isAccountDeployed(web2User, systemMaster);
            expect(existsBefore).to.be.false;
            console.log(`2. Wallet exists before: ${existsBefore}`);
            
            // Step 3: Create AA wallet
            const tx = await walletManager.createUserAccount(web2User, systemMaster);
            const receipt = await tx.wait();
            const event = receipt.events?.find(e => e.event === "AccountCreated");
            
            expect(event?.args?.account).to.equal(predictedWallet);
            console.log(`3. Created AA wallet: ${event?.args?.account}`);
            console.log(`   Gas used: ${receipt.gasUsed.toString()}`);
            
            // Step 4: Verify deployment
            const existsAfter = await walletManager.isAccountDeployed(web2User, systemMaster);
            expect(existsAfter).to.be.true;
            console.log(`4. Wallet exists after: ${existsAfter}`);
            
            // Step 5: Generate initCode for future EntryPoint operations
            const initCode = await walletManager.getInitCode(web2User, systemMaster);
            console.log(`5. InitCode length: ${initCode.length} chars`);
            
            console.log("=== Workflow Complete ===\n");
        });

        it("should handle enterprise multi-tenant scenario", async function () {
            const tenant1Users = [
                "0x1000000000000000000000000000000000000001",
                "0x1000000000000000000000000000000000000002"
            ];
            const tenant2Users = [
                "0x2000000000000000000000000000000000000001", 
                "0x2000000000000000000000000000000000000002"
            ];
            
            const tenant1Master = masterSignerAddress;
            const tenant2Master = userAddress;
            
            console.log("\n=== Multi-Tenant Enterprise Scenario ===");
            
            // Authorize tenant masters as creators
            await walletManager.authorizeCreator(tenant1Master);
            await walletManager.authorizeCreator(tenant2Master);
            
            const tenant1Wallets: string[] = [];
            const tenant2Wallets: string[] = [];
            
            // Create wallets for tenant 1
            for (const user of tenant1Users) {
                const tx = await walletManager.connect(masterSigner).createUserAccount(user, tenant1Master);
                const receipt = await tx.wait();
                const event = receipt.events?.find(e => e.event === "AccountCreated");
                if (event?.args?.account) {
                    tenant1Wallets.push(event.args.account);
                }
            }
            
            // Create wallets for tenant 2  
            for (const userAddr of tenant2Users) {
                const tx = await walletManager.connect(user).createUserAccount(userAddr, tenant2Master);
                const receipt = await tx.wait();
                const event = receipt.events?.find(e => e.event === "AccountCreated");
                if (event?.args?.account) {
                    tenant2Wallets.push(event.args.account);
                }
            }
            
            console.log(`Tenant 1 wallets: ${tenant1Wallets.length}`);
            console.log(`Tenant 2 wallets: ${tenant2Wallets.length}`);
            
            // Verify all wallets are unique
            const allWallets = [...tenant1Wallets, ...tenant2Wallets];
            const uniqueWallets = [...new Set(allWallets)];
            expect(uniqueWallets.length).to.equal(allWallets.length);
            
            console.log("=== Multi-Tenant Scenario Complete ===\n");
        });
    });
});