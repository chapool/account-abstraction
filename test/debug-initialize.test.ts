import { expect } from "chai";
import { ethers } from "hardhat";
import { Signer } from "ethers";
import {
    AAWallet,
    EntryPoint,
    AAWallet__factory,
    EntryPoint__factory,
} from "../typechain";

describe("AAWallet Initialize Debug", function () {
    let owner: Signer;
    let entryPoint: EntryPoint;
    let wallet: AAWallet;
    let ownerAddress: string;

    before(async function () {
        [owner] = await ethers.getSigners();
        ownerAddress = await owner.getAddress();
        console.log("Owner address:", ownerAddress);
    });

    beforeEach(async function () {
        // Deploy EntryPoint
        const EntryPointFactory = (await ethers.getContractFactory("EntryPoint")) as EntryPoint__factory;
        entryPoint = await EntryPointFactory.deploy();
        await entryPoint.deployed();
        
        // Deploy AAWallet
        const AAWalletFactory = (await ethers.getContractFactory("AAWallet")) as AAWallet__factory;
        wallet = await AAWalletFactory.deploy();
        await wallet.deployed();
        
        console.log("Contracts deployed");
    });

    it("Should debug the initialization error step by step", async function () {
        console.log("=== Starting initialization debug ===");
        
        // Step 1: Check if contract is already initialized
        try {
            // This should work - just reading storage
            const currentOwner = await wallet.getOwner();
            console.log("✓ Current owner (before init):", currentOwner);
            
            const currentMasterSigner = await wallet.getMasterSigner();
            console.log("✓ Current master signer (before init):", currentMasterSigner);
            
        } catch (error) {
            console.log("✗ Error reading current state:", error);
        }

        // Step 2: Try to check if initialize function exists and is callable
        try {
            // Encode the call data manually to see what's happening
            const initCallData = wallet.interface.encodeFunctionData("initialize", [ownerAddress, ethers.constants.AddressZero]);
            console.log("✓ Initialize call data encoded:", initCallData);
            
        } catch (error) {
            console.log("✗ Error encoding initialize call:", error);
        }

        // Step 3: Try to estimate gas for the call
        try {
            const gasEstimate = await wallet.estimateGas.initialize(ownerAddress, ethers.constants.AddressZero);
            console.log("✓ Gas estimate:", gasEstimate.toString());
            
        } catch (error: any) {
            console.log("✗ Gas estimation failed:");
            console.log("  Error message:", error.message);
            console.log("  Error code:", error.code);
            if (error.data) {
                console.log("  Error data:", error.data);
            }
            if (error.reason) {
                console.log("  Error reason:", error.reason);
            }
        }

        // Step 4: Try to call initialize with different approaches
        try {
            console.log("Attempting initialize call...");
            
            // Method 1: Direct call with default gas
            const tx = await wallet.initialize(ownerAddress, ethers.constants.AddressZero);
            await tx.wait();
            console.log("✓ Initialize succeeded!");
            
        } catch (error: any) {
            console.log("✗ Initialize failed:");
            console.log("  Error message:", error.message);
            console.log("  Error code:", error.code);
            if (error.data) {
                console.log("  Error data:", error.data);
                
                // Try to decode the error data
                if (error.data.startsWith("0xf92ee8a9")) {
                    console.log("  This is the custom error 0xf92ee8a9");
                    // Try to decode more of the error data
                    console.log("  Full error data:", error.data);
                }
            }
            
            // Don't rethrow - we want to continue debugging
        }

        // Step 5: Check contract bytecode and storage
        try {
            const code = await ethers.provider.getCode(wallet.address);
            console.log("✓ Contract bytecode length:", code.length);
            console.log("✓ Contract has code:", code !== "0x");
            
        } catch (error) {
            console.log("✗ Error checking contract code:", error);
        }
    });

    it("Should try manual initialization steps", async function () {
        console.log("=== Trying manual initialization ===");
        
        // Let's try to understand what exactly is failing by calling individual parts
        
        // Check if the contract supports the Initializable interface
        try {
            // Call some view functions to see what's accessible
            const entryPointAddr = await wallet.entryPoint();
            console.log("✓ EntryPoint accessible:", entryPointAddr);
            
            const nonce = await wallet.getNonce();
            console.log("✓ Nonce accessible:", nonce.toString());
            
        } catch (error) {
            console.log("✗ Error accessing basic functions:", error);
        }
        
        // Try to understand what the 0xf92ee8a9 error is
        console.log("Error signature 0xf92ee8a9 analysis:");
        console.log("This might be InvalidInitialization() from OpenZeppelin Initializable");
        
        // Let's manually calculate some common error signatures to compare
        const commonErrors = [
            "InvalidInitialization()",
            "NotInitializing()", 
            "AlreadyInitialized()",
        ];
        
        for (const errorSig of commonErrors) {
            const hash = ethers.utils.keccak256(ethers.utils.toUtf8Bytes(errorSig));
            const selector = hash.slice(0, 10);
            console.log(`  ${errorSig}: ${selector}`);
            if (selector === "0xf92ee8a9") {
                console.log(`  *** MATCH FOUND: ${errorSig} ***`);
            }
        }
    });
});