import { ethers } from "hardhat";
import { writeFileSync } from "fs";
import { join } from "path";

async function main() {
    const [deployer] = await ethers.getSigners();
    console.log("Deploying contracts with the account:", deployer.address);
    console.log("Account balance:", (await deployer.provider.getBalance(deployer.address)).toString());

    // Deploy MockUSDT
    console.log("\n=== Deploying MockUSDT ===");
    const MockUSDT = await ethers.getContractFactory("MockUSDT");
    const mockUSDT = await MockUSDT.deploy();
    await mockUSDT.deployed();
    const mockUSDTAddress = mockUSDT.address;
    console.log("MockUSDT deployed to:", mockUSDTAddress);

    // Deploy Payment contract
    console.log("\n=== Deploying Payment ===");
    const Payment = await ethers.getContractFactory("Payment");
    const payment = await Payment.deploy();
    await payment.deployed();
    const paymentAddress = payment.address;
    console.log("Payment deployed to:", paymentAddress);

    // Add MockUSDT to payment whitelist
    console.log("\n=== Adding MockUSDT to whitelist ===");
    const tx = await payment.addAllowedToken(mockUSDTAddress);
    await tx.wait();
    console.log("MockUSDT added to payment whitelist");

    // Verify MockUSDT is in whitelist
    const isAllowed = await payment.isTokenAllowed(mockUSDTAddress);
    console.log("MockUSDT whitelist status:", isAllowed);

    // Get some MockUSDT info
    const name = await mockUSDT.name();
    const symbol = await mockUSDT.symbol();
    const decimals = await mockUSDT.decimals();
    const totalSupply = await mockUSDT.totalSupply();
    
    console.log("\n=== MockUSDT Info ===");
    console.log("Name:", name);
    console.log("Symbol:", symbol);
    console.log("Decimals:", decimals);
    console.log("Total Supply:", ethers.utils.formatUnits(totalSupply, decimals));

    // Save deployment results
    const deploymentResults = {
        network: "sepolia",
        timestamp: new Date().toISOString(),
        deployer: deployer.address,
        contracts: {
            MockUSDT: {
                address: mockUSDTAddress,
                name,
                symbol,
                decimals: decimals.toString(),
                totalSupply: totalSupply.toString()
            },
            Payment: {
                address: paymentAddress,
                owner: deployer.address
            }
        },
        transactions: {
            mockUSDTDeploy: mockUSDT.deployTransaction?.hash,
            paymentDeploy: payment.deployTransaction?.hash,
            addToWhitelist: tx.hash
        }
    };

    const resultsPath = join(__dirname, "../sepolia-payment-deployment.json");
    writeFileSync(resultsPath, JSON.stringify(deploymentResults, null, 2));
    console.log("\n=== Deployment Results Saved ===");
    console.log("Results saved to:", resultsPath);

    // Instructions for users
    console.log("\n=== Usage Instructions ===");
    console.log("1. Get test USDT:");
    console.log(`   mockUSDT.faucet() - gives 1000 USDT to caller`);
    console.log("\n2. Pay for order:");
    console.log(`   payment.pay(orderId, "${mockUSDTAddress}", amount)`);
    console.log("\n3. Check payment:");
    console.log(`   payment.getPayment(orderId)`);
    console.log("\n4. Contract addresses:");
    console.log(`   MockUSDT: ${mockUSDTAddress}`);
    console.log(`   Payment:  ${paymentAddress}`);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });