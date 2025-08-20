import { expect } from "chai";
import { ethers } from "hardhat";
import { Signer } from "ethers";
import { CPOPToken } from "../typechain";

describe("CPOPToken Batch Operations", function () {
    let admin: Signer;
    let minter: Signer;
    let burner: Signer;
    let user1: Signer;
    let user2: Signer;
    let user3: Signer;
    let user4: Signer;
    let user5: Signer;
    
    let adminAddress: string;
    let minterAddress: string;
    let burnerAddress: string;
    let user1Address: string;
    let user2Address: string;
    let user3Address: string;
    let user4Address: string;
    let user5Address: string;
    
    let token: CPOPToken;
    
    const INITIAL_SUPPLY = ethers.utils.parseEther("1000000");
    
    // Role constants
    const ADMIN_ROLE = 1;  // 0001
    const MINTER_ROLE = 2; // 0010  
    const BURNER_ROLE = 4; // 0100

    beforeEach(async function () {
        [admin, minter, burner, user1, user2, user3, user4, user5] = await ethers.getSigners();
        
        adminAddress = await admin.getAddress();
        minterAddress = await minter.getAddress();
        burnerAddress = await burner.getAddress();
        user1Address = await user1.getAddress();
        user2Address = await user2.getAddress();
        user3Address = await user3.getAddress();
        user4Address = await user4.getAddress();
        user5Address = await user5.getAddress();

        const TokenFactory = await ethers.getContractFactory("contracts/cpop/CPOPToken.sol:CPOPToken");
        token = await TokenFactory.deploy(adminAddress, INITIAL_SUPPLY);
        await token.deployed();
        
        // Setup roles for testing
        await token.grantRole(minterAddress, MINTER_ROLE);
        await token.grantRole(burnerAddress, BURNER_ROLE);
    });

    describe("Batch Mint Operations", function () {
        it("should mint tokens to multiple addresses in one transaction", async function () {
            const recipients = [user1Address, user2Address, user3Address];
            const amounts = [
                ethers.utils.parseEther("100"),
                ethers.utils.parseEther("200"), 
                ethers.utils.parseEther("300")
            ];
            
            const initialTotalSupply = await token.totalSupply();
            
            const tx = await token.connect(minter).batchMint(recipients, amounts);
            const receipt = await tx.wait();
            
            console.log(`Batch mint gas for 3 addresses: ${receipt.gasUsed.toString()}`);
            
            // Check balances
            expect(await token.balanceOf(user1Address)).to.equal(amounts[0]);
            expect(await token.balanceOf(user2Address)).to.equal(amounts[1]);
            expect(await token.balanceOf(user3Address)).to.equal(amounts[2]);
            
            // Check total supply
            const expectedTotalSupply = initialTotalSupply.add(amounts[0]).add(amounts[1]).add(amounts[2]);
            expect(await token.totalSupply()).to.equal(expectedTotalSupply);
            
            // Check Transfer events
            const transferEvents = receipt.events?.filter(e => e.event === "Transfer");
            expect(transferEvents?.length).to.equal(3);
        });

        it("should handle large batch mint operations efficiently", async function () {
            const recipients = [user1Address, user2Address, user3Address, user4Address, user5Address];
            const amounts = [
                ethers.utils.parseEther("50"),
                ethers.utils.parseEther("75"),
                ethers.utils.parseEther("100"),
                ethers.utils.parseEther("125"),
                ethers.utils.parseEther("150")
            ];
            
            const tx = await token.connect(minter).batchMint(recipients, amounts);
            const receipt = await tx.wait();
            
            console.log(`Batch mint gas for 5 addresses: ${receipt.gasUsed.toString()}`);
            
            // Verify all recipients received tokens
            for (let i = 0; i < recipients.length; i++) {
                expect(await token.balanceOf(recipients[i])).to.equal(amounts[i]);
            }
            
            // Should be more efficient than individual mints
            expect(receipt.gasUsed.toNumber()).to.be.lessThan(300000); // Reasonable limit for 5 mints
        });

        it("should reject batch mint from non-minter", async function () {
            const recipients = [user1Address, user2Address];
            const amounts = [ethers.utils.parseEther("100"), ethers.utils.parseEther("200")];
            
            await expect(
                token.connect(user1).batchMint(recipients, amounts)
            ).to.be.reverted;
        });

        it("should reject batch mint with mismatched array lengths", async function () {
            const recipients = [user1Address, user2Address];
            const amounts = [ethers.utils.parseEther("100")]; // Length mismatch
            
            await expect(
                token.connect(minter).batchMint(recipients, amounts)
            ).to.be.reverted;
        });

        it("should reject empty batch mint", async function () {
            const recipients: string[] = [];
            const amounts: any[] = [];
            
            await expect(
                token.connect(minter).batchMint(recipients, amounts)
            ).to.be.reverted;
        });
    });

    describe("Batch Burn Operations", function () {
        beforeEach(async function () {
            // Setup some tokens for burning
            const recipients = [user1Address, user2Address, user3Address, user4Address, user5Address];
            const amounts = [
                ethers.utils.parseEther("1000"),
                ethers.utils.parseEther("1000"),
                ethers.utils.parseEther("1000"),
                ethers.utils.parseEther("1000"),
                ethers.utils.parseEther("1000")
            ];
            
            await token.connect(minter).batchMint(recipients, amounts);
        });

        it("should burn tokens from multiple addresses in one transaction", async function () {
            const accounts = [user1Address, user2Address, user3Address];
            const amounts = [
                ethers.utils.parseEther("200"),
                ethers.utils.parseEther("300"),
                ethers.utils.parseEther("400")
            ];
            
            const initialTotalSupply = await token.totalSupply();
            const initialBalances = await Promise.all(
                accounts.map(addr => token.balanceOf(addr))
            );
            
            const tx = await token.connect(burner).batchBurn(accounts, amounts);
            const receipt = await tx.wait();
            
            console.log(`Batch burn gas for 3 addresses: ${receipt.gasUsed.toString()}`);
            
            // Check balances
            for (let i = 0; i < accounts.length; i++) {
                const expectedBalance = initialBalances[i].sub(amounts[i]);
                expect(await token.balanceOf(accounts[i])).to.equal(expectedBalance);
            }
            
            // Check total supply
            const totalBurned = amounts.reduce((sum, amount) => sum.add(amount), ethers.BigNumber.from(0));
            expect(await token.totalSupply()).to.equal(initialTotalSupply.sub(totalBurned));
            
            // Check Transfer events
            const transferEvents = receipt.events?.filter(e => e.event === "Transfer");
            expect(transferEvents?.length).to.equal(3);
        });

        it("should handle large batch burn operations efficiently", async function () {
            const accounts = [user1Address, user2Address, user3Address, user4Address, user5Address];
            const amounts = [
                ethers.utils.parseEther("100"),
                ethers.utils.parseEther("150"),
                ethers.utils.parseEther("200"),
                ethers.utils.parseEther("250"),
                ethers.utils.parseEther("300")
            ];
            
            const tx = await token.connect(burner).batchBurn(accounts, amounts);
            const receipt = await tx.wait();
            
            console.log(`Batch burn gas for 5 addresses: ${receipt.gasUsed.toString()}`);
            
            // Should be efficient
            expect(receipt.gasUsed.toNumber()).to.be.lessThan(250000);
        });

        it("should reject batch burn from non-burner", async function () {
            const accounts = [user1Address, user2Address];
            const amounts = [ethers.utils.parseEther("100"), ethers.utils.parseEther("200")];
            
            await expect(
                token.connect(user1).batchBurn(accounts, amounts)
            ).to.be.reverted;
        });

        it("should reject batch burn with insufficient balance", async function () {
            const accounts = [user1Address];
            
            // Check actual balance first
            const balance = await token.balanceOf(user1Address);
            console.log(`User1 actual balance: ${ethers.utils.formatEther(balance)}`);
            
            const amounts = [balance.add(ethers.utils.parseEther("1"))]; // More than actual balance
            
            await expect(
                token.connect(burner).batchBurn(accounts, amounts)
            ).to.be.reverted;
        });

        it("should reject batch burn with mismatched array lengths", async function () {
            const accounts = [user1Address, user2Address];
            const amounts = [ethers.utils.parseEther("100")]; // Length mismatch
            
            await expect(
                token.connect(burner).batchBurn(accounts, amounts)
            ).to.be.reverted;
        });

        it("should reject empty batch burn", async function () {
            const accounts: string[] = [];
            const amounts: any[] = [];
            
            await expect(
                token.connect(burner).batchBurn(accounts, amounts)
            ).to.be.reverted;
        });
    });

    describe("Gas Efficiency Comparison", function () {
        it("should compare batch vs individual operations", async function () {
            console.log("\n=== BATCH vs INDIVIDUAL GAS COMPARISON ===");
            
            const recipients = [user1Address, user2Address, user3Address];
            const amounts = [
                ethers.utils.parseEther("100"),
                ethers.utils.parseEther("200"),
                ethers.utils.parseEther("300")
            ];
            
            // Individual mints
            let individualGas = 0;
            for (let i = 0; i < recipients.length; i++) {
                const tx = await token.connect(minter).mint(recipients[i], amounts[i]);
                const receipt = await tx.wait();
                individualGas += receipt.gasUsed.toNumber();
            }
            
            // Reset balances for fair comparison
            await token.connect(burner).batchBurn(recipients, amounts);
            
            // Batch mint
            const batchTx = await token.connect(minter).batchMint(recipients, amounts);
            const batchReceipt = await batchTx.wait();
            const batchGas = batchReceipt.gasUsed.toNumber();
            
            console.log(`Individual mints (3x): ${individualGas} gas`);
            console.log(`Batch mint (3x): ${batchGas} gas`);
            console.log(`Gas savings: ${individualGas - batchGas} gas (${((individualGas - batchGas) / individualGas * 100).toFixed(2)}%)`);
            
            expect(batchGas).to.be.lessThan(individualGas);
            console.log("=======================================\n");
        });
    });

    describe("Real-world Scenarios", function () {
        it("should handle airdrop scenario", async function () {
            console.log("\n=== AIRDROP SCENARIO ===");
            
            // Simulate airdrop to 5 users
            const airdropRecipients = [user1Address, user2Address, user3Address, user4Address, user5Address];
            const airdropAmounts = Array(5).fill(ethers.utils.parseEther("50")); // 50 tokens each
            
            const tx = await token.connect(admin).batchMint(airdropRecipients, airdropAmounts);
            const receipt = await tx.wait();
            
            console.log(`Airdrop to 5 users gas: ${receipt.gasUsed.toString()}`);
            console.log(`Average gas per recipient: ${Math.round(receipt.gasUsed.toNumber() / 5)}`);
            
            // Verify all users received tokens
            for (const recipient of airdropRecipients) {
                expect(await token.balanceOf(recipient)).to.equal(ethers.utils.parseEther("50"));
            }
            
            console.log("=========================\n");
        });

        it("should handle fee collection scenario", async function () {
            console.log("\n=== FEE COLLECTION SCENARIO ===");
            
            // First give users some tokens
            const users = [user1Address, user2Address, user3Address, user4Address];
            const initialAmounts = Array(4).fill(ethers.utils.parseEther("500"));
            await token.connect(minter).batchMint(users, initialAmounts);
            
            // Collect fees from users (10 tokens each)
            const feeAmounts = Array(4).fill(ethers.utils.parseEther("10"));
            
            const tx = await token.connect(burner).batchBurn(users, feeAmounts);
            const receipt = await tx.wait();
            
            console.log(`Fee collection from 4 users gas: ${receipt.gasUsed.toString()}`);
            console.log(`Average gas per user: ${Math.round(receipt.gasUsed.toNumber() / 4)}`);
            
            // Verify fees were collected
            for (const user of users) {
                expect(await token.balanceOf(user)).to.equal(ethers.utils.parseEther("490"));
            }
            
            console.log("==============================\n");
        });

        it("should handle mixed operations scenario", async function () {
            console.log("\n=== MIXED OPERATIONS SCENARIO ===");
            
            // Simulate a complex scenario: mint rewards, burn fees
            const rewardRecipients = [user1Address, user2Address];
            const rewardAmounts = [ethers.utils.parseEther("100"), ethers.utils.parseEther("150")];
            
            const feeAccounts = [user3Address, user4Address];
            const feeAmounts = [ethers.utils.parseEther("20"), ethers.utils.parseEther("30")];
            
            // First give fee accounts some tokens
            await token.connect(minter).batchMint(feeAccounts, [ethers.utils.parseEther("100"), ethers.utils.parseEther("100")]);
            
            // Execute rewards and fees in separate transactions
            const rewardTx = await token.connect(admin).batchMint(rewardRecipients, rewardAmounts);
            const rewardReceipt = await rewardTx.wait();
            
            const feeTx = await token.connect(admin).batchBurn(feeAccounts, feeAmounts);
            const feeReceipt = await feeTx.wait();
            
            console.log(`Batch reward mint gas: ${rewardReceipt.gasUsed.toString()}`);
            console.log(`Batch fee burn gas: ${feeReceipt.gasUsed.toString()}`);
            console.log(`Total gas for mixed operations: ${rewardReceipt.gasUsed.add(feeReceipt.gasUsed).toString()}`);
            
            console.log("================================\n");
        });
    });
});