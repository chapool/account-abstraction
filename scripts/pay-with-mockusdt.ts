import { ethers } from "hardhat";
import { writeFileSync } from "fs";
import { join } from "path";

async function main() {
    const [deployer] = await ethers.getSigners();
    console.log("Making payment with account:", deployer.address);
    console.log("Account balance:", ethers.utils.formatEther(await deployer.provider.getBalance(deployer.address)), "ETH");

    // Contract addresses from deployment
    const PAYMENT_CONTRACT_ADDRESS = "0x7Da90c3364a3D0B99fCD18c1651E669C4D294D65";
    const MOCK_USDT_ADDRESS = "0x0D3E58b48Ef96A1eAF2CD9A553558FFAf490Eb88";

    // Payment parameters
    const ORDER_ID = 1; // Order ID for this payment
    const PAYMENT_AMOUNT = ethers.utils.parseUnits("1000", 6); // 1000 USDT (6 decimals)

    console.log("\n=== Payment Details ===");
    console.log("Order ID:", ORDER_ID);
    console.log("Payment Amount:", ethers.utils.formatUnits(PAYMENT_AMOUNT, 6), "USDT");
    console.log("Token Address:", MOCK_USDT_ADDRESS);

    try {
        // Get contract instances
        console.log("\n=== Getting Contract Instances ===");
        const Payment = await ethers.getContractFactory("Payment");
        const payment = Payment.attach(PAYMENT_CONTRACT_ADDRESS);

        const MockUSDT = await ethers.getContractFactory("MockUSDT");
        const mockUSDT = MockUSDT.attach(MOCK_USDT_ADDRESS);

        // Check if token is allowed
        console.log("\n=== Checking Token Status ===");
        const isTokenAllowed = await payment.isTokenAllowed(MOCK_USDT_ADDRESS);
        console.log("MockUSDT is allowed:", isTokenAllowed);

        if (!isTokenAllowed) {
            console.log("❌ MockUSDT is not in the whitelist. Please add it first.");
            return;
        }

        // Check if order is already paid
        console.log("\n=== Checking Order Status ===");
        const isOrderPaid = await payment.isOrderPaid(ORDER_ID);
        console.log("Order is already paid:", isOrderPaid);

        if (isOrderPaid) {
            console.log("❌ Order is already paid. Please use a different order ID.");
            return;
        }

        // Check user's USDT balance
        console.log("\n=== Checking User Balance ===");
        const userUSDTBalance = await mockUSDT.balanceOf(deployer.address);
        console.log("User USDT balance:", ethers.utils.formatUnits(userUSDTBalance, 6), "USDT");

        if (userUSDTBalance.lt(PAYMENT_AMOUNT)) {
            console.log("❌ Insufficient USDT balance. Getting from faucet...");
            
            // Get USDT from faucet
            const faucetTx = await mockUSDT.faucet();
            await faucetTx.wait();
            console.log("✅ Faucet transaction:", faucetTx.hash);
            
            // Check balance again
            const newBalance = await mockUSDT.balanceOf(deployer.address);
            console.log("New USDT balance:", ethers.utils.formatUnits(newBalance, 6), "USDT");
        }

        // Check USDT allowance
        console.log("\n=== Checking Allowance ===");
        const allowance = await mockUSDT.allowance(deployer.address, PAYMENT_CONTRACT_ADDRESS);
        console.log("Current allowance:", ethers.utils.formatUnits(allowance, 6), "USDT");

        if (allowance.lt(PAYMENT_AMOUNT)) {
            console.log("Approving USDT spending...");
            const approveTx = await mockUSDT.approve(PAYMENT_CONTRACT_ADDRESS, PAYMENT_AMOUNT);
            await approveTx.wait();
            console.log("✅ Approval transaction:", approveTx.hash);
        }

        // Make the payment
        console.log("\n=== Making Payment ===");
        console.log("Calling payment.pay()...");
        
        const paymentTx = await payment.pay(ORDER_ID, MOCK_USDT_ADDRESS, PAYMENT_AMOUNT);
        console.log("Payment transaction submitted:", paymentTx.hash);
        
        // Wait for confirmation
        console.log("Waiting for confirmation...");
        const receipt = await paymentTx.wait();
        console.log("✅ Payment confirmed!");
        console.log("Block number:", receipt.blockNumber);
        console.log("Gas used:", receipt.gasUsed.toString());

        // Verify payment
        console.log("\n=== Verifying Payment ===");
        const paymentInfo = await payment.getPayment(ORDER_ID);
        console.log("Payment Info:");
        console.log("  Order ID:", paymentInfo.orderId.toString());
        console.log("  Payer:", paymentInfo.payer);
        console.log("  Token:", paymentInfo.token);
        console.log("  Amount:", ethers.utils.formatUnits(paymentInfo.amount, 6), "USDT");
        console.log("  Timestamp:", new Date(paymentInfo.timestamp.toNumber() * 1000).toISOString());

        // Check contract USDT balance
        const contractUSDTBalance = await mockUSDT.balanceOf(PAYMENT_CONTRACT_ADDRESS);
        console.log("\nContract USDT balance:", ethers.utils.formatUnits(contractUSDTBalance, 6), "USDT");

        // Save payment results
        const paymentResults = {
            network: "sepolia",
            timestamp: new Date().toISOString(),
            payer: deployer.address,
            orderId: ORDER_ID,
            token: MOCK_USDT_ADDRESS,
            amount: PAYMENT_AMOUNT.toString(),
            amountFormatted: ethers.utils.formatUnits(PAYMENT_AMOUNT, 6),
            transactions: {
                payment: paymentTx.hash,
                blockNumber: receipt.blockNumber,
                gasUsed: receipt.gasUsed.toString()
            },
            paymentInfo: {
                orderId: paymentInfo.orderId.toString(),
                payer: paymentInfo.payer,
                token: paymentInfo.token,
                amount: paymentInfo.amount.toString(),
                timestamp: paymentInfo.timestamp.toString()
            }
        };

        const resultsPath = join(__dirname, "../payment-results.json");
        writeFileSync(resultsPath, JSON.stringify(paymentResults, null, 2));
        console.log("\n=== Payment Results Saved ===");
        console.log("Results saved to:", resultsPath);

        console.log("\n🎉 Payment completed successfully!");
        console.log("Order ID:", ORDER_ID);
        console.log("Amount:", ethers.utils.formatUnits(PAYMENT_AMOUNT, 6), "USDT");
        console.log("Transaction:", paymentTx.hash);

    } catch (error) {
        console.error("❌ Payment failed:", error);
        
        if (error.reason) {
            console.error("Reason:", error.reason);
        }
        
        if (error.code) {
            console.error("Error code:", error.code);
        }
    }
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
