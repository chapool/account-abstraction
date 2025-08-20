import { expect } from "chai";
import { ethers } from "hardhat";
import { Signer } from "ethers";
import { CPOPToken } from "../typechain";

describe("CPOPToken Mint Gas Comparison - 10 Operations", function () {
    let admin: Signer;
    let users: Signer[];
    let userAddresses: string[];
    
    let token: CPOPToken;
    
    const INITIAL_SUPPLY = ethers.utils.parseEther("1000000");
    
    // Role constants
    const ADMIN_ROLE = 1;
    const MINTER_ROLE = 2;

    beforeEach(async function () {
        const signers = await ethers.getSigners();
        admin = signers[0];
        users = signers.slice(1, 11); // Get 10 users
        userAddresses = await Promise.all(users.map(user => user.getAddress()));

        const TokenFactory = await ethers.getContractFactory("contracts/cpop/CPOPToken.sol:CPOPToken");
        token = await TokenFactory.deploy(await admin.getAddress(), INITIAL_SUPPLY);
        await token.deployed();
    });

    describe("Gas Comparison: 10 Individual Mints vs 1 Batch Mint", function () {
        it("should compare gas costs for 10 mint operations", async function () {
            console.log("\n=== 10 MINT OPERATIONS GAS COMPARISON ===");
            
            const mintAmount = ethers.utils.parseEther("100");
            const amounts = Array(10).fill(mintAmount);
            
            // Test 1: 10 Individual Mints
            console.log("\n--- Individual Mints (10 separate transactions) ---");
            let totalIndividualGas = 0;
            const individualGasUsages: number[] = [];
            
            for (let i = 0; i < 10; i++) {
                const tx = await token.mint(userAddresses[i], mintAmount);
                const receipt = await tx.wait();
                const gasUsed = receipt.gasUsed.toNumber();
                individualGasUsages.push(gasUsed);
                totalIndividualGas += gasUsed;
                console.log(`Mint ${i + 1}: ${gasUsed} gas`);
            }
            
            console.log(`Total Individual Mints Gas: ${totalIndividualGas}`);
            console.log(`Average per mint: ${Math.round(totalIndividualGas / 10)}`);
            
            // Reset balances by burning all tokens
            for (let i = 0; i < 10; i++) {
                await token.adminBurn(userAddresses[i], mintAmount);
            }
            
            // Test 2: 1 Batch Mint (10 recipients)
            console.log("\n--- Batch Mint (1 transaction, 10 recipients) ---");
            const batchTx = await token.batchMint(userAddresses, amounts);
            const batchReceipt = await batchTx.wait();
            const batchGas = batchReceipt.gasUsed.toNumber();
            
            console.log(`Batch Mint Gas: ${batchGas}`);
            console.log(`Average per recipient: ${Math.round(batchGas / 10)}`);
            
            // Calculate savings
            const gasSavings = totalIndividualGas - batchGas;
            const savingsPercentage = ((gasSavings / totalIndividualGas) * 100).toFixed(2);
            
            console.log("\n--- Comparison Results ---");
            console.log(`Individual mints (10x): ${totalIndividualGas} gas`);
            console.log(`Batch mint (10x): ${batchGas} gas`);
            console.log(`Gas savings: ${gasSavings} gas`);
            console.log(`Savings percentage: ${savingsPercentage}%`);
            console.log(`Efficiency ratio: ${(totalIndividualGas / batchGas).toFixed(2)}x`);
            
            // Verify batch mint is more efficient
            expect(batchGas).to.be.lessThan(totalIndividualGas);
            
            // Verify all recipients received tokens
            for (let i = 0; i < 10; i++) {
                expect(await token.balanceOf(userAddresses[i])).to.equal(mintAmount);
            }
            
            console.log("==========================================\n");
        });

        it("should compare different batch sizes", async function () {
            console.log("\n=== BATCH SIZE COMPARISON ===");
            
            const mintAmount = ethers.utils.parseEther("100");
            const batchSizes = [1, 2, 5, 10];
            const results: { size: number; gas: number; gasPerRecipient: number }[] = [];
            
            for (const size of batchSizes) {
                const recipients = userAddresses.slice(0, size);
                const amounts = Array(size).fill(mintAmount);
                
                const tx = await token.batchMint(recipients, amounts);
                const receipt = await tx.wait();
                const gasUsed = receipt.gasUsed.toNumber();
                const gasPerRecipient = Math.round(gasUsed / size);
                
                results.push({ size, gas: gasUsed, gasPerRecipient });
                console.log(`Batch size ${size}: ${gasUsed} gas (${gasPerRecipient} gas/recipient)`);
                
                // Clean up for next test
                for (const recipient of recipients) {
                    await token.adminBurn(recipient, mintAmount);
                }
            }
            
            // Calculate efficiency improvement
            const individualGasPerMint = results[0].gasPerRecipient;
            const batch10GasPerMint = results[3].gasPerRecipient;
            const improvement = ((individualGasPerMint - batch10GasPerMint) / individualGasPerMint * 100).toFixed(2);
            
            console.log(`\nGas per mint improvement (1 vs 10): ${improvement}%`);
            console.log("===============================\n");
        });

        it("should test scalability with larger batch sizes", async function () {
            console.log("\n=== SCALABILITY TEST ===");
            
            // We'll need more signers for larger batches
            const allSigners = await ethers.getSigners();
            const maxTestSize = Math.min(20, allSigners.length - 1); // Reserve 1 for admin
            
            const mintAmount = ethers.utils.parseEther("50");
            const testSizes = [5, 10, 15, maxTestSize];
            
            for (const size of testSizes) {
                const recipients = (await Promise.all(
                    allSigners.slice(1, size + 1).map(s => s.getAddress())
                ));
                const amounts = Array(size).fill(mintAmount);
                
                const tx = await token.batchMint(recipients, amounts);
                const receipt = await tx.wait();
                const gasUsed = receipt.gasUsed.toNumber();
                const gasPerRecipient = Math.round(gasUsed / size);
                
                console.log(`Batch size ${size}: ${gasUsed} gas (${gasPerRecipient} gas/recipient)`);
                
                // Verify all mints succeeded
                for (const recipient of recipients) {
                    expect(await token.balanceOf(recipient)).to.equal(mintAmount);
                }
                
                // Clean up
                await token.batchBurn(recipients, amounts);
            }
            
            console.log("========================\n");
        });

        it("should measure gas cost breakdown", async function () {
            console.log("\n=== GAS COST BREAKDOWN ===");
            
            const mintAmount = ethers.utils.parseEther("100");
            
            // Test different operations for comparison
            const operations = [
                { name: "Single mint", fn: () => token.mint(userAddresses[0], mintAmount) },
                { name: "Batch mint (2)", fn: () => token.batchMint(userAddresses.slice(0, 2), [mintAmount, mintAmount]) },
                { name: "Batch mint (5)", fn: () => token.batchMint(userAddresses.slice(0, 5), Array(5).fill(mintAmount)) },
                { name: "Batch mint (10)", fn: () => token.batchMint(userAddresses.slice(0, 10), Array(10).fill(mintAmount)) }
            ];
            
            for (const op of operations) {
                const tx = await op.fn();
                const receipt = await tx.wait();
                const gasUsed = receipt.gasUsed.toNumber();
                
                const recipientCount = op.name.includes("(") ? 
                    parseInt(op.name.match(/\((\d+)\)/)?.[1] || "1") : 1;
                const gasPerRecipient = Math.round(gasUsed / recipientCount);
                
                console.log(`${op.name}: ${gasUsed} gas (${gasPerRecipient} gas/recipient)`);
            }
            
            console.log("===========================\n");
        });
    });
});