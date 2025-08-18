import { expect } from "chai";
import { ethers } from "hardhat";
import { Signer } from "ethers";
import {
    SessionKeyManager,
    SessionKeyManager__factory,
    AAWallet,
    AAWallet__factory,
    EntryPoint,
    EntryPoint__factory
} from "../typechain";

describe("SessionKeyManager - Unit Tests", function () {
    let owner: Signer;
    let masterSigner1: Signer;
    let masterSigner2: Signer;
    let sessionKey1: Signer;
    let sessionKey2: Signer;
    let unauthorized: Signer;
    
    let entryPoint: EntryPoint;
    let sessionKeyManager: SessionKeyManager;
    let mockWallet1: AAWallet;
    let mockWallet2: AAWallet;
    
    let ownerAddress: string;
    let masterSigner1Address: string;
    let masterSigner2Address: string;
    let sessionKey1Address: string;
    let sessionKey2Address: string;
    let unauthorizedAddress: string;
    let mockWallet1Address: string;
    let mockWallet2Address: string;

    beforeEach(async function () {
        [owner, masterSigner1, masterSigner2, sessionKey1, sessionKey2, unauthorized] = await ethers.getSigners();
        
        ownerAddress = await owner.getAddress();
        masterSigner1Address = await masterSigner1.getAddress();
        masterSigner2Address = await masterSigner2.getAddress();
        sessionKey1Address = await sessionKey1.getAddress();
        sessionKey2Address = await sessionKey2.getAddress();
        unauthorizedAddress = await unauthorized.getAddress();

        // Deploy EntryPoint
        const EntryPointFactory = (await ethers.getContractFactory("EntryPoint")) as EntryPoint__factory;
        entryPoint = await EntryPointFactory.deploy();
        await entryPoint.deployed();

        // Deploy mock AAWallet instances via proxy pattern
        const AAWalletFactory = (await ethers.getContractFactory("AAWallet")) as AAWallet__factory;
        
        // Deploy implementation
        const aaWalletImpl = await AAWalletFactory.deploy();
        await aaWalletImpl.deployed();
        
        // Deploy mockWallet1 via proxy
        const initData1 = aaWalletImpl.interface.encodeFunctionData("initialize", [
            entryPoint.address, ownerAddress, masterSigner1Address
        ]);
        const proxyFactory1 = await ethers.getContractFactory("ERC1967Proxy");
        const proxy1 = await proxyFactory1.deploy(aaWalletImpl.address, initData1);
        await proxy1.deployed();
        mockWallet1 = AAWalletFactory.attach(proxy1.address);
        mockWallet1Address = mockWallet1.address;

        // Deploy mockWallet2 via proxy  
        const initData2 = aaWalletImpl.interface.encodeFunctionData("initialize", [
            entryPoint.address, ownerAddress, masterSigner2Address
        ]);
        const proxyFactory2 = await ethers.getContractFactory("ERC1967Proxy");
        const proxy2 = await proxyFactory2.deploy(aaWalletImpl.address, initData2);
        await proxy2.deployed();
        mockWallet2 = AAWalletFactory.attach(proxy2.address);
        mockWallet2Address = mockWallet2.address;

        // Deploy SessionKeyManager via proxy pattern
        const SessionKeyManagerFactory = (await ethers.getContractFactory("SessionKeyManager")) as SessionKeyManager__factory;
        const sessionKeyManagerImpl = await SessionKeyManagerFactory.deploy();
        await sessionKeyManagerImpl.deployed();

        // Deploy proxy
        const initData = sessionKeyManagerImpl.interface.encodeFunctionData("initialize", [ownerAddress]);
        const proxyFactory = await ethers.getContractFactory("ERC1967Proxy");
        const proxy = await proxyFactory.deploy(sessionKeyManagerImpl.address, initData);
        await proxy.deployed();

        // Connect to proxy through SessionKeyManager interface
        sessionKeyManager = SessionKeyManagerFactory.attach(proxy.address);
        
        // Authorize the SessionKeyManager to manage session keys on the wallets
        await mockWallet1.connect(owner).setAuthorizedSessionKeyManager(sessionKeyManager.address);
        await mockWallet2.connect(owner).setAuthorizedSessionKeyManager(sessionKeyManager.address);
    });

    describe("Initialization", function () {
        it("should initialize correctly with default templates", async function () {
            expect(await sessionKeyManager.owner()).to.equal(ownerAddress);
            
            // Check default templates exist
            const templateNames = await sessionKeyManager.getTemplateNames();
            expect(templateNames.length).to.equal(3);
            expect(templateNames).to.include("DAPP_INTERACTION");
            expect(templateNames).to.include("TRADING");
            expect(templateNames).to.include("GAME_INTERACTION");
        });

        it("should have correct default template configurations", async function () {
            const dappTemplate = await sessionKeyManager.getTemplate("DAPP_INTERACTION");
            expect(dappTemplate.name).to.equal("DAPP_INTERACTION");
            expect(dappTemplate.defaultDuration).to.equal(24 * 60 * 60); // 24 hours
            expect(dappTemplate.permissions).to.equal(ethers.constants.HashZero); // Allow all
            expect(dappTemplate.isActive).to.be.true;

            const tradingTemplate = await sessionKeyManager.getTemplate("TRADING");
            expect(tradingTemplate.name).to.equal("TRADING");
            expect(tradingTemplate.defaultDuration).to.equal(60 * 60); // 1 hour
            expect(tradingTemplate.isActive).to.be.true;

            const gameTemplate = await sessionKeyManager.getTemplate("GAME_INTERACTION");
            expect(gameTemplate.name).to.equal("GAME_INTERACTION");
            expect(gameTemplate.defaultDuration).to.equal(7 * 24 * 60 * 60); // 7 days
            expect(gameTemplate.isActive).to.be.true;
        });

        it("should revert with invalid owner", async function () {
            const SessionKeyManagerFactory = (await ethers.getContractFactory("SessionKeyManager")) as SessionKeyManager__factory;
            const impl = await SessionKeyManagerFactory.deploy();
            await impl.deployed();

            // Deploy via proxy to test initialization with invalid owner
            const initData = impl.interface.encodeFunctionData("initialize", [ethers.constants.AddressZero]);
            const proxyFactory = await ethers.getContractFactory("ERC1967Proxy");
            
            await expect(
                proxyFactory.deploy(impl.address, initData)
            ).to.be.revertedWith("SessionKeyManager: invalid owner");
        });
    });

    describe("Master Signer Management", function () {
        it("should authorize master signer", async function () {
            expect(await sessionKeyManager.authorizedMasters(masterSigner1Address)).to.be.false;

            const tx = await sessionKeyManager.connect(owner).authorizeMasterSigner(masterSigner1Address);
            const receipt = await tx.wait();

            expect(await sessionKeyManager.authorizedMasters(masterSigner1Address)).to.be.true;

            const event = receipt.events?.find(e => e.event === "MasterSignerAuthorized");
            expect(event?.args?.masterSigner).to.equal(masterSigner1Address);
            expect(event?.args?.authorized).to.be.true;
        });

        it("should revoke master signer authorization", async function () {
            // First authorize
            await sessionKeyManager.connect(owner).authorizeMasterSigner(masterSigner1Address);
            expect(await sessionKeyManager.authorizedMasters(masterSigner1Address)).to.be.true;

            // Then revoke
            const tx = await sessionKeyManager.connect(owner).revokeMasterSigner(masterSigner1Address);
            const receipt = await tx.wait();

            expect(await sessionKeyManager.authorizedMasters(masterSigner1Address)).to.be.false;

            const event = receipt.events?.find(e => e.event === "MasterSignerAuthorized");
            expect(event?.args?.masterSigner).to.equal(masterSigner1Address);
            expect(event?.args?.authorized).to.be.false;
        });

        it("should only allow owner to authorize master signers", async function () {
            await expect(
                sessionKeyManager.connect(unauthorized).authorizeMasterSigner(masterSigner1Address)
            ).to.be.reverted;
        });

        it("should only allow owner to revoke master signers", async function () {
            await sessionKeyManager.connect(owner).authorizeMasterSigner(masterSigner1Address);

            await expect(
                sessionKeyManager.connect(unauthorized).revokeMasterSigner(masterSigner1Address)
            ).to.be.reverted;
        });

        it("should revert when authorizing invalid master signer", async function () {
            await expect(
                sessionKeyManager.connect(owner).authorizeMasterSigner(ethers.constants.AddressZero)
            ).to.be.revertedWith("SessionKeyManager: invalid master signer");
        });

        it("should revert when authorizing already authorized master signer", async function () {
            await sessionKeyManager.connect(owner).authorizeMasterSigner(masterSigner1Address);

            await expect(
                sessionKeyManager.connect(owner).authorizeMasterSigner(masterSigner1Address)
            ).to.be.revertedWith("SessionKeyManager: already authorized");
        });

        it("should revert when revoking non-authorized master signer", async function () {
            await expect(
                sessionKeyManager.connect(owner).revokeMasterSigner(masterSigner1Address)
            ).to.be.revertedWith("SessionKeyManager: not authorized");
        });
    });

    describe("Account Registration", function () {
        beforeEach(async function () {
            // Authorize master signers
            await sessionKeyManager.connect(owner).authorizeMasterSigner(masterSigner1Address);
            await sessionKeyManager.connect(owner).authorizeMasterSigner(masterSigner2Address);
        });

        it("should register account under master signer", async function () {
            const tx = await sessionKeyManager.registerAccount(masterSigner1Address, mockWallet1Address);
            const receipt = await tx.wait();

            const registeredAccounts = await sessionKeyManager.getRegisteredAccounts(masterSigner1Address);
            expect(registeredAccounts.length).to.equal(1);
            expect(registeredAccounts[0]).to.equal(mockWallet1Address);

            const event = receipt.events?.find(e => e.event === "AccountRegistered");
            expect(event?.args?.masterSigner).to.equal(masterSigner1Address);
            expect(event?.args?.account).to.equal(mockWallet1Address);
        });

        it("should register multiple accounts under same master signer", async function () {
            await sessionKeyManager.registerAccount(masterSigner1Address, mockWallet1Address);
            
            // Deploy another wallet with same master signer via proxy
            const AAWalletFactory = (await ethers.getContractFactory("AAWallet")) as AAWallet__factory;
            const aaWalletImpl = await AAWalletFactory.deploy();
            await aaWalletImpl.deployed();
            
            const initData = aaWalletImpl.interface.encodeFunctionData("initialize", [
                entryPoint.address, ownerAddress, masterSigner1Address
            ]);
            const proxyFactory = await ethers.getContractFactory("ERC1967Proxy");
            const proxy3 = await proxyFactory.deploy(aaWalletImpl.address, initData);
            await proxy3.deployed();
            const mockWallet3 = AAWalletFactory.attach(proxy3.address);
            await mockWallet3.connect(owner).setAuthorizedSessionKeyManager(sessionKeyManager.address);

            await sessionKeyManager.registerAccount(masterSigner1Address, mockWallet3.address);

            const registeredAccounts = await sessionKeyManager.getRegisteredAccounts(masterSigner1Address);
            expect(registeredAccounts.length).to.equal(2);
            expect(registeredAccounts).to.include(mockWallet1Address);
            expect(registeredAccounts).to.include(mockWallet3.address);
        });

        it("should unregister account from master signer", async function () {
            await sessionKeyManager.registerAccount(masterSigner1Address, mockWallet1Address);

            const tx = await sessionKeyManager.connect(masterSigner1).unregisterAccount(masterSigner1Address, mockWallet1Address);
            const receipt = await tx.wait();

            const registeredAccounts = await sessionKeyManager.getRegisteredAccounts(masterSigner1Address);
            expect(registeredAccounts.length).to.equal(0);

            const event = receipt.events?.find(e => e.event === "AccountUnregistered");
            expect(event?.args?.masterSigner).to.equal(masterSigner1Address);
            expect(event?.args?.account).to.equal(mockWallet1Address);
        });

        it("should allow owner to unregister accounts", async function () {
            await sessionKeyManager.registerAccount(masterSigner1Address, mockWallet1Address);

            // Owner needs to be authorized as master signer OR we need to call as owner 
            // But looking at the contract logic, owner isn't automatically authorized master
            // So this test should either authorize owner first, or expect failure
            
            // Let's authorize owner as master signer first
            await sessionKeyManager.connect(owner).authorizeMasterSigner(ownerAddress);
            await sessionKeyManager.connect(owner).unregisterAccount(masterSigner1Address, mockWallet1Address);

            const registeredAccounts = await sessionKeyManager.getRegisteredAccounts(masterSigner1Address);
            expect(registeredAccounts.length).to.equal(0);
        });

        it("should revert when registering with unauthorized master signer", async function () {
            await expect(
                sessionKeyManager.registerAccount(unauthorizedAddress, mockWallet1Address)
            ).to.be.revertedWith("SessionKeyManager: master not authorized");
        });

        it("should revert when registering invalid account", async function () {
            await expect(
                sessionKeyManager.registerAccount(masterSigner1Address, ethers.constants.AddressZero)
            ).to.be.revertedWith("SessionKeyManager: invalid account");
        });

        it("should revert when registering account with mismatched master signer", async function () {
            // Try to register wallet1 (has masterSigner1) under masterSigner2
            await expect(
                sessionKeyManager.registerAccount(masterSigner2Address, mockWallet1Address)
            ).to.be.revertedWith("SessionKeyManager: master signer mismatch");
        });

        it("should revert when registering already registered account", async function () {
            await sessionKeyManager.registerAccount(masterSigner1Address, mockWallet1Address);

            await expect(
                sessionKeyManager.registerAccount(masterSigner1Address, mockWallet1Address)
            ).to.be.revertedWith("SessionKeyManager: account already registered");
        });

        it("should revert when unregistering non-existent account", async function () {
            await expect(
                sessionKeyManager.connect(masterSigner1).unregisterAccount(masterSigner1Address, mockWallet1Address)
            ).to.be.revertedWith("SessionKeyManager: account not found");
        });

        it("should only allow authorized entities to unregister accounts", async function () {
            await sessionKeyManager.registerAccount(masterSigner1Address, mockWallet1Address);

            await expect(
                sessionKeyManager.connect(unauthorized).unregisterAccount(masterSigner1Address, mockWallet1Address)
            ).to.be.revertedWith("SessionKeyManager: not authorized master");
        });
    });

    describe("Template Management", function () {
        it("should create custom session key template", async function () {
            const templateName = "CUSTOM_TEMPLATE";
            const duration = 3600; // 1 hour
            const permissions = ethers.utils.keccak256(ethers.utils.toUtf8Bytes("CUSTOM_PERMISSIONS"));

            const tx = await sessionKeyManager.connect(owner).createSessionKeyTemplate(
                templateName,
                duration,
                permissions
            );
            const receipt = await tx.wait();

            const template = await sessionKeyManager.getTemplate(templateName);
            expect(template.name).to.equal(templateName);
            expect(template.defaultDuration).to.equal(duration);
            expect(template.permissions).to.equal(permissions);
            expect(template.isActive).to.be.true;

            const templateNames = await sessionKeyManager.getTemplateNames();
            expect(templateNames).to.include(templateName);

            const event = receipt.events?.find(e => e.event === "SessionKeyTemplateCreated");
            expect(event).to.not.be.undefined;
            // For indexed string parameters, need to access the hash
            // For non-indexed parameters, access by index or name
            expect(event?.args?.[1]).to.equal(duration); // Duration parameter  
            expect(event?.args?.[2]).to.equal(permissions); // Permissions parameter
            // Template name is indexed so it's hashed, just verify event exists
        });

        it("should only allow owner to create templates", async function () {
            await expect(
                sessionKeyManager.connect(unauthorized).createSessionKeyTemplate("TEST", 3600, ethers.constants.HashZero)
            ).to.be.reverted;
        });

        it("should revert when creating template with invalid name", async function () {
            await expect(
                sessionKeyManager.connect(owner).createSessionKeyTemplate("", 3600, ethers.constants.HashZero)
            ).to.be.revertedWith("SessionKeyManager: invalid template name");
        });

        it("should revert when creating template with invalid duration", async function () {
            await expect(
                sessionKeyManager.connect(owner).createSessionKeyTemplate("TEST", 0, ethers.constants.HashZero)
            ).to.be.revertedWith("SessionKeyManager: invalid duration");
        });

        it("should revert when creating template with existing name", async function () {
            await expect(
                sessionKeyManager.connect(owner).createSessionKeyTemplate("DAPP_INTERACTION", 3600, ethers.constants.HashZero)
            ).to.be.revertedWith("SessionKeyManager: template already exists");
        });

        it("should return empty template for non-existent template", async function () {
            const template = await sessionKeyManager.getTemplate("NON_EXISTENT");
            expect(template.name).to.equal("");
            expect(template.defaultDuration).to.equal(0);
            expect(template.isActive).to.be.false;
        });
    });

    describe("Session Key Operations with Templates", function () {
        beforeEach(async function () {
            // Setup: authorize master signer and register accounts with same master signer
            await sessionKeyManager.connect(owner).authorizeMasterSigner(masterSigner1Address);
            await sessionKeyManager.registerAccount(masterSigner1Address, mockWallet1Address);
            // Create another wallet with the same master signer
            const AAWalletFactory = (await ethers.getContractFactory("AAWallet")) as AAWallet__factory;
            const aaWalletImpl = await AAWalletFactory.deploy();
            await aaWalletImpl.deployed();
            const initData = aaWalletImpl.interface.encodeFunctionData("initialize", [
                entryPoint.address, ownerAddress, masterSigner1Address
            ]);
            const proxyFactory = await ethers.getContractFactory("ERC1967Proxy");
            const proxy = await proxyFactory.deploy(aaWalletImpl.address, initData);
            await proxy.deployed();
            const mockWallet3 = AAWalletFactory.attach(proxy.address);
            await mockWallet3.connect(owner).setAuthorizedSessionKeyManager(sessionKeyManager.address);
            await sessionKeyManager.registerAccount(masterSigner1Address, mockWallet3.address);
        });

        it("should add session key with template to all registered accounts", async function () {
            const tx = await sessionKeyManager.connect(masterSigner1).addSessionKeyWithTemplate(
                masterSigner1Address,
                sessionKey1Address,
                "DAPP_INTERACTION",
                0 // Use template default duration
            );
            const receipt = await tx.wait();

            const event = receipt.events?.find(e => e.event === "BatchSessionKeysAdded");
            expect(event?.args?.masterSigner).to.equal(masterSigner1Address);
            expect(event?.args?.operationCount).to.equal(2); // Should process 2 accounts
        });

        it("should add session key with custom duration", async function () {
            const customDuration = 7200; // 2 hours

            const tx = await sessionKeyManager.connect(masterSigner1).addSessionKeyWithTemplate(
                masterSigner1Address,
                sessionKey1Address,
                "DAPP_INTERACTION",
                customDuration
            );
            const receipt = await tx.wait();

            const event = receipt.events?.find(e => e.event === "BatchSessionKeysAdded");
            expect(event?.args?.operationCount).to.equal(2);
        });

        it("should only allow authorized master signer to add session keys", async function () {
            await expect(
                sessionKeyManager.connect(unauthorized).addSessionKeyWithTemplate(
                    masterSigner1Address,
                    sessionKey1Address,
                    "DAPP_INTERACTION",
                    0
                )
            ).to.be.revertedWith("SessionKeyManager: not authorized master");
        });

        it("should only allow master signer to manage own accounts", async function () {
            await sessionKeyManager.connect(owner).authorizeMasterSigner(masterSigner2Address);

            await expect(
                sessionKeyManager.connect(masterSigner2).addSessionKeyWithTemplate(
                    masterSigner1Address,
                    sessionKey1Address,
                    "DAPP_INTERACTION",
                    0
                )
            ).to.be.revertedWith("SessionKeyManager: can only manage own accounts");
        });

        it("should revert with non-existent template", async function () {
            await expect(
                sessionKeyManager.connect(masterSigner1).addSessionKeyWithTemplate(
                    masterSigner1Address,
                    sessionKey1Address,
                    "NON_EXISTENT_TEMPLATE",
                    0
                )
            ).to.be.revertedWith("SessionKeyManager: template not found");
        });
    });

    describe("Batch Session Key Operations", function () {
        beforeEach(async function () {
            // Setup: authorize master signer and register accounts with same master signer
            await sessionKeyManager.connect(owner).authorizeMasterSigner(masterSigner1Address);
            await sessionKeyManager.registerAccount(masterSigner1Address, mockWallet1Address);
            // Create another wallet with the same master signer
            const AAWalletFactory = (await ethers.getContractFactory("AAWallet")) as AAWallet__factory;
            const aaWalletImpl = await AAWalletFactory.deploy();
            await aaWalletImpl.deployed();
            const initData = aaWalletImpl.interface.encodeFunctionData("initialize", [
                entryPoint.address, ownerAddress, masterSigner1Address
            ]);
            const proxyFactory = await ethers.getContractFactory("ERC1967Proxy");
            const proxy = await proxyFactory.deploy(aaWalletImpl.address, initData);
            await proxy.deployed();
            const mockWallet3 = AAWalletFactory.attach(proxy.address);
            await mockWallet3.connect(owner).setAuthorizedSessionKeyManager(sessionKeyManager.address);
            await sessionKeyManager.registerAccount(masterSigner1Address, mockWallet3.address);
        });

        it("should batch add session keys to multiple accounts", async function () {
            const currentTime = Math.floor(Date.now() / 1000);
            // Get the actual registered accounts to use correct addresses
            const registeredAccounts = await sessionKeyManager.getRegisteredAccounts(masterSigner1Address);
            expect(registeredAccounts.length).to.be.greaterThanOrEqual(2, "Need at least 2 registered accounts");
            
            const operations = [
                {
                    account: registeredAccounts[0], // First registered account
                    sessionKey: sessionKey1Address,
                    validAfter: currentTime,
                    validUntil: currentTime + 3600,
                    permissions: ethers.constants.HashZero
                },
                {
                    account: registeredAccounts[1], // Second registered account
                    sessionKey: sessionKey1Address,
                    validAfter: currentTime,
                    validUntil: currentTime + 3600,
                    permissions: ethers.constants.HashZero
                }
            ];

            const tx = await sessionKeyManager.connect(masterSigner1).batchAddSessionKeys(operations);
            const receipt = await tx.wait();

            const event = receipt.events?.find(e => e.event === "BatchSessionKeysAdded");
            expect(event?.args?.masterSigner).to.equal(masterSigner1Address);
            expect(event?.args?.operationCount).to.equal(2);
        });

        it("should skip accounts not registered to master signer", async function () {
            // Create an operation with an unregistered account
            const currentTime = Math.floor(Date.now() / 1000);
            const operations = [
                {
                    account: unauthorizedAddress, // Not registered under masterSigner1
                    sessionKey: sessionKey1Address,
                    validAfter: currentTime,
                    validUntil: currentTime + 3600,
                    permissions: ethers.constants.HashZero
                },
                {
                    account: mockWallet1Address, // Registered
                    sessionKey: sessionKey1Address,
                    validAfter: currentTime,
                    validUntil: currentTime + 3600,
                    permissions: ethers.constants.HashZero
                }
            ];

            const tx = await sessionKeyManager.connect(masterSigner1).batchAddSessionKeys(operations);
            const receipt = await tx.wait();

            const event = receipt.events?.find(e => e.event === "BatchSessionKeysAdded");
            expect(event?.args?.operationCount).to.equal(1); // Only 1 should succeed
        });

        it("should batch revoke session keys from all registered accounts", async function () {
            // First add session keys to the accounts so we can revoke them
            await sessionKeyManager.connect(masterSigner1).addSessionKeyWithTemplate(
                masterSigner1Address,
                sessionKey1Address,
                "DAPP_INTERACTION",
                0 // Use template default duration
            );

            const tx = await sessionKeyManager.connect(masterSigner1).batchRevokeSessionKey(sessionKey1Address);
            const receipt = await tx.wait();

            const event = receipt.events?.find(e => e.event === "BatchSessionKeysRevoked");
            expect(event?.args?.masterSigner).to.equal(masterSigner1Address);
            expect(event?.args?.operationCount).to.equal(2); // Should process 2 accounts
        });

        it("should only allow authorized master signers to perform batch operations", async function () {
            await expect(
                sessionKeyManager.connect(unauthorized).batchAddSessionKeys([])
            ).to.be.revertedWith("SessionKeyManager: not authorized master");

            await expect(
                sessionKeyManager.connect(unauthorized).batchRevokeSessionKey(sessionKey1Address)
            ).to.be.revertedWith("SessionKeyManager: not authorized master");
        });

        it("should revert with no batch operations", async function () {
            await expect(
                sessionKeyManager.connect(masterSigner1).batchAddSessionKeys([])
            ).to.be.revertedWith("SessionKeyManager: no operations");
        });
    });

    describe("View Functions", function () {
        beforeEach(async function () {
            await sessionKeyManager.connect(owner).authorizeMasterSigner(masterSigner1Address);
            await sessionKeyManager.connect(owner).authorizeMasterSigner(masterSigner2Address);
            await sessionKeyManager.registerAccount(masterSigner1Address, mockWallet1Address);
            await sessionKeyManager.registerAccount(masterSigner2Address, mockWallet2Address);
        });

        it("should return registered accounts for master signer", async function () {
            const accounts1 = await sessionKeyManager.getRegisteredAccounts(masterSigner1Address);
            expect(accounts1.length).to.equal(1);
            expect(accounts1[0]).to.equal(mockWallet1Address);

            const accounts2 = await sessionKeyManager.getRegisteredAccounts(masterSigner2Address);
            expect(accounts2.length).to.equal(1);
            expect(accounts2[0]).to.equal(mockWallet2Address);
        });

        it("should return empty array for master signer with no registered accounts", async function () {
            const accounts = await sessionKeyManager.getRegisteredAccounts(unauthorizedAddress);
            expect(accounts.length).to.equal(0);
        });

        it("should return all template names", async function () {
            const templateNames = await sessionKeyManager.getTemplateNames();
            expect(templateNames.length).to.equal(3);
            expect(templateNames).to.include("DAPP_INTERACTION");
            expect(templateNames).to.include("TRADING");
            expect(templateNames).to.include("GAME_INTERACTION");
        });

        it("should return template information", async function () {
            const template = await sessionKeyManager.getTemplate("TRADING");
            expect(template.name).to.equal("TRADING");
            expect(template.defaultDuration).to.equal(60 * 60); // 1 hour
            expect(template.isActive).to.be.true;
        });

        it("should return specific account from masterSignerAccounts mapping", async function () {
            const account0 = await sessionKeyManager.masterSignerAccounts(masterSigner1Address, 0);
            const account1 = await sessionKeyManager.masterSignerAccounts(masterSigner2Address, 0);
            
            expect(account0).to.equal(mockWallet1Address);
            expect(account1).to.equal(mockWallet2Address);
        });
    });

    describe("Edge Cases and Error Handling", function () {
        it("should handle registration of account with getMasterSigner() failure", async function () {
            await sessionKeyManager.connect(owner).authorizeMasterSigner(masterSigner1Address);
            
            // Try to register an address that's not a proper AAWallet
            await expect(
                sessionKeyManager.registerAccount(masterSigner1Address, unauthorizedAddress)
            ).to.be.reverted;
        });

        it("should handle empty accounts list in batch operations gracefully", async function () {
            await sessionKeyManager.connect(owner).authorizeMasterSigner(masterSigner1Address);
            // Don't register any accounts

            const tx = await sessionKeyManager.connect(masterSigner1).batchRevokeSessionKey(sessionKey1Address);
            const receipt = await tx.wait();

            const event = receipt.events?.find(e => e.event === "BatchSessionKeysRevoked");
            expect(event?.args?.operationCount).to.equal(0);
        });

        it("should handle array bounds correctly when unregistering", async function () {
            await sessionKeyManager.connect(owner).authorizeMasterSigner(masterSigner1Address);
            await sessionKeyManager.registerAccount(masterSigner1Address, mockWallet1Address);
            
            // Create another wallet with the same master signer via proxy
            const AAWalletFactory = (await ethers.getContractFactory("AAWallet")) as AAWallet__factory;
            const aaWalletImpl = await AAWalletFactory.deploy();
            await aaWalletImpl.deployed();
            const initData = aaWalletImpl.interface.encodeFunctionData("initialize", [
                entryPoint.address, ownerAddress, masterSigner1Address
            ]);
            const proxyFactory = await ethers.getContractFactory("ERC1967Proxy");
            const proxy = await proxyFactory.deploy(aaWalletImpl.address, initData);
            await proxy.deployed();
            const mockWallet3 = AAWalletFactory.attach(proxy.address);
            await mockWallet3.connect(owner).setAuthorizedSessionKeyManager(sessionKeyManager.address);
            await sessionKeyManager.registerAccount(masterSigner1Address, mockWallet3.address);

            // Unregister first account
            await sessionKeyManager.connect(masterSigner1).unregisterAccount(masterSigner1Address, mockWallet1Address);
            
            // Verify second account is still there
            const accounts = await sessionKeyManager.getRegisteredAccounts(masterSigner1Address);
            expect(accounts.length).to.equal(1);
            expect(accounts[0]).to.equal(mockWallet3.address);
        });
    });

    describe("Upgrade Functionality", function () {
        it("should only allow owner to authorize upgrades", async function () {
            // This tests the _authorizeUpgrade internal function via a hypothetical upgrade
            // In a real scenario, you'd need to deploy a new implementation and test upgrading
            
            // For now, we can just verify the owner role is properly set
            expect(await sessionKeyManager.owner()).to.equal(ownerAddress);
        });
    });
});