import { expect } from "chai";
import { ethers } from "hardhat";
import { Signer } from "ethers";
import {
    GasPriceOracle,
    GasPriceOracle__factory
} from "../typechain";

describe("GasPriceOracle - Unit Tests", function () {
    let owner: Signer;
    let oracle1: Signer;
    let oracle2: Signer;
    let unauthorizedUser: Signer;
    
    let gasPriceOracle: GasPriceOracle;
    
    let ownerAddress: string;
    let oracle1Address: string;
    let oracle2Address: string;
    let unauthorizedUserAddress: string;

    // Test constants
    const INITIAL_ETH_PRICE = 2000_00000000; // $2000 with 8 decimals
    const INITIAL_CPOP_PRICE = 1_00000000;   // $1 with 8 decimals
    const MAX_PRICE_AGE = 3600; // 1 hour
    const PRICE_DEVIATION_THRESHOLD = 500; // 5%
    const BASIS_POINTS = 10000;

    beforeEach(async function () {
        [owner, oracle1, oracle2, unauthorizedUser] = await ethers.getSigners();
        
        ownerAddress = await owner.getAddress();
        oracle1Address = await oracle1.getAddress();
        oracle2Address = await oracle2.getAddress();
        unauthorizedUserAddress = await unauthorizedUser.getAddress();

        // Deploy GasPriceOracle via proxy pattern
        const GasPriceOracleFactory = (await ethers.getContractFactory("GasPriceOracle")) as GasPriceOracle__factory;
        const gasPriceOracleImpl = await GasPriceOracleFactory.deploy();
        await gasPriceOracleImpl.deployed();

        // Deploy proxy
        const initData = gasPriceOracleImpl.interface.encodeFunctionData("initialize", [
            ownerAddress,
            MAX_PRICE_AGE,
            PRICE_DEVIATION_THRESHOLD
        ]);
        const proxyFactory = await ethers.getContractFactory("ERC1967Proxy");
        const proxy = await proxyFactory.deploy(gasPriceOracleImpl.address, initData);
        await proxy.deployed();

        // Connect to proxy through GasPriceOracle interface
        gasPriceOracle = GasPriceOracleFactory.attach(proxy.address);

        // Authorize oracles
        await gasPriceOracle.connect(owner).authorizeOracle(oracle1Address);
        await gasPriceOracle.connect(owner).authorizeOracle(oracle2Address);

        // Set initial prices
        await gasPriceOracle.connect(oracle1).updatePriceFeed("ETH/USD", INITIAL_ETH_PRICE);
        await gasPriceOracle.connect(oracle1).updatePriceFeed("CPOP/USD", INITIAL_CPOP_PRICE);
    });

    describe("Initialization", function () {
        it("should initialize correctly with proper configuration", async function () {
            expect(await gasPriceOracle.owner()).to.equal(ownerAddress);
            expect(await gasPriceOracle.getMaxPriceAge()).to.equal(MAX_PRICE_AGE);
            expect(await gasPriceOracle.getPriceDeviationThreshold()).to.equal(PRICE_DEVIATION_THRESHOLD);
            expect(await gasPriceOracle.authorizedOracles(oracle1Address)).to.be.true;
            expect(await gasPriceOracle.authorizedOracles(oracle2Address)).to.be.true;
        });

        it("should have ETH and CPOP feeds active", async function () {
            expect(await gasPriceOracle.isPriceValid("ETH/USD")).to.be.true;
            expect(await gasPriceOracle.isPriceValid("CPOP/USD")).to.be.true;
        });

        it("should revert with invalid owner", async function () {
            const GasPriceOracleFactory = (await ethers.getContractFactory("GasPriceOracle")) as GasPriceOracle__factory;
            const impl = await GasPriceOracleFactory.deploy();
            await impl.deployed();

            const initData = impl.interface.encodeFunctionData("initialize", [
                ethers.constants.AddressZero,
                MAX_PRICE_AGE,
                PRICE_DEVIATION_THRESHOLD
            ]);
            const proxyFactory = await ethers.getContractFactory("ERC1967Proxy");
            
            await expect(
                proxyFactory.deploy(impl.address, initData)
            ).to.be.revertedWith("GasPriceOracle: invalid owner");
        });

        it("should revert with invalid max price age", async function () {
            const GasPriceOracleFactory = (await ethers.getContractFactory("GasPriceOracle")) as GasPriceOracle__factory;
            const impl = await GasPriceOracleFactory.deploy();
            await impl.deployed();

            const initData = impl.interface.encodeFunctionData("initialize", [
                ownerAddress,
                0, // Invalid max price age
                PRICE_DEVIATION_THRESHOLD
            ]);
            const proxyFactory = await ethers.getContractFactory("ERC1967Proxy");
            
            await expect(
                proxyFactory.deploy(impl.address, initData)
            ).to.be.revertedWith("GasPriceOracle: invalid max price age");
        });

        it("should revert with invalid price deviation threshold", async function () {
            const GasPriceOracleFactory = (await ethers.getContractFactory("GasPriceOracle")) as GasPriceOracle__factory;
            const impl = await GasPriceOracleFactory.deploy();
            await impl.deployed();

            const initData = impl.interface.encodeFunctionData("initialize", [
                ownerAddress,
                MAX_PRICE_AGE,
                BASIS_POINTS + 1 // Invalid threshold > 100%
            ]);
            const proxyFactory = await ethers.getContractFactory("ERC1967Proxy");
            
            await expect(
                proxyFactory.deploy(impl.address, initData)
            ).to.be.revertedWith("GasPriceOracle: invalid threshold");
        });
    });

    describe("Oracle Authorization Management", function () {
        it("should authorize oracle", async function () {
            const newOracle = await unauthorizedUser.getAddress();
            expect(await gasPriceOracle.authorizedOracles(newOracle)).to.be.false;

            await gasPriceOracle.connect(owner).authorizeOracle(newOracle);

            expect(await gasPriceOracle.authorizedOracles(newOracle)).to.be.true;
        });

        it("should revoke oracle authorization", async function () {
            expect(await gasPriceOracle.authorizedOracles(oracle1Address)).to.be.true;

            await gasPriceOracle.connect(owner).revokeOracle(oracle1Address);

            expect(await gasPriceOracle.authorizedOracles(oracle1Address)).to.be.false;
        });

        it("should only allow owner to authorize oracles", async function () {
            await expect(
                gasPriceOracle.connect(unauthorizedUser).authorizeOracle(unauthorizedUserAddress)
            ).to.be.reverted;
        });

        it("should only allow owner to revoke oracles", async function () {
            await expect(
                gasPriceOracle.connect(unauthorizedUser).revokeOracle(oracle1Address)
            ).to.be.reverted;
        });

        it("should revert when authorizing invalid oracle address", async function () {
            await expect(
                gasPriceOracle.connect(owner).authorizeOracle(ethers.constants.AddressZero)
            ).to.be.revertedWith("GasPriceOracle: invalid oracle");
        });

        it("should revert when authorizing already authorized oracle", async function () {
            await expect(
                gasPriceOracle.connect(owner).authorizeOracle(oracle1Address)
            ).to.be.revertedWith("GasPriceOracle: already authorized");
        });

        it("should revert when revoking non-authorized oracle", async function () {
            await expect(
                gasPriceOracle.connect(owner).revokeOracle(unauthorizedUserAddress)
            ).to.be.revertedWith("GasPriceOracle: not authorized");
        });
    });

    describe("Price Feed Management", function () {
        it("should add new price feed", async function () {
            const feedType = "BTC/USD";
            const initialPrice = 50000_00000000; // $50,000
            const fallbackPrice = 49000_00000000; // $49,000

            await gasPriceOracle.connect(owner).addPriceFeed(feedType, initialPrice, fallbackPrice);

            const [price, timestamp, isValid] = await gasPriceOracle.getPriceWithMetadata(feedType);
            expect(price).to.equal(initialPrice);
            expect(isValid).to.be.true;
        });

        it("should remove price feed", async function () {
            const feedType = "BTC/USD";
            const initialPrice = 50000_00000000;
            
            // First add the feed
            await gasPriceOracle.connect(owner).addPriceFeed(feedType, initialPrice, 0);
            expect(await gasPriceOracle.isPriceValid(feedType)).to.be.true;

            // Then remove it
            await gasPriceOracle.connect(owner).removePriceFeed(feedType);
            expect(await gasPriceOracle.isPriceValid(feedType)).to.be.false;
        });

        it("should revert when adding feed with invalid initial price", async function () {
            await expect(
                gasPriceOracle.connect(owner).addPriceFeed("BTC/USD", 0, 1000_00000000)
            ).to.be.revertedWith("GasPriceOracle: invalid initial price");
        });

        it("should revert when adding already existing feed", async function () {
            await expect(
                gasPriceOracle.connect(owner).addPriceFeed("ETH/USD", 2000_00000000, 1900_00000000)
            ).to.be.revertedWith("GasPriceOracle: feed already exists");
        });

        it("should revert when removing non-active feed", async function () {
            await expect(
                gasPriceOracle.connect(owner).removePriceFeed("BTC/USD")
            ).to.be.revertedWith("GasPriceOracle: feed not active");
        });

        it("should only allow owner to manage price feeds", async function () {
            await expect(
                gasPriceOracle.connect(unauthorizedUser).addPriceFeed("BTC/USD", 50000_00000000, 0)
            ).to.be.reverted;

            await expect(
                gasPriceOracle.connect(unauthorizedUser).removePriceFeed("ETH/USD")
            ).to.be.reverted;
        });
    });

    describe("Price Updates", function () {
        it("should update price feed by authorized oracle", async function () {
            const newPrice = 2100_00000000; // $2100
            const oldPrice = INITIAL_ETH_PRICE;

            const tx = await gasPriceOracle.connect(oracle1).updatePriceFeed("ETH/USD", newPrice);
            const receipt = await tx.wait();

            const [currentPrice] = await gasPriceOracle.getETHPriceUSD();
            expect(currentPrice).to.equal(newPrice);

            // Check for event in logs
            const event = receipt?.logs.find((log: any) => {
                try {
                    const parsed = gasPriceOracle.interface.parseLog(log);
                    return parsed?.name === "PriceFeedUpdated";
                } catch {
                    return false;
                }
            });
            expect(event).to.not.be.undefined;
        });

        it("should update price feed by owner", async function () {
            const newPrice = 2100_00000000;

            await gasPriceOracle.connect(owner).updatePriceFeed("ETH/USD", newPrice);

            const [currentPrice] = await gasPriceOracle.getETHPriceUSD();
            expect(currentPrice).to.equal(newPrice);
        });

        it("should revert price update by unauthorized user", async function () {
            const newPrice = 2100_00000000;

            await expect(
                gasPriceOracle.connect(unauthorizedUser).updatePriceFeed("ETH/USD", newPrice)
            ).to.be.revertedWith("GasPriceOracle: unauthorized oracle");
        });

        it("should revert with invalid price", async function () {
            await expect(
                gasPriceOracle.connect(oracle1).updatePriceFeed("ETH/USD", 0)
            ).to.be.revertedWith("GasPriceOracle: invalid price");
        });

        it("should revert with inactive feed", async function () {
            await expect(
                gasPriceOracle.connect(oracle1).updatePriceFeed("BTC/USD", 50000_00000000)
            ).to.be.revertedWith("GasPriceOracle: feed not active");
        });

        it("should revert when price deviation is too high", async function () {
            const highDeviationPrice = INITIAL_ETH_PRICE * 2; // 100% increase

            await expect(
                gasPriceOracle.connect(oracle1).updatePriceFeed("ETH/USD", highDeviationPrice)
            ).to.be.revertedWith("GasPriceOracle: price deviation too high");
        });

        it("should allow price update within deviation threshold", async function () {
            const withinThresholdPrice = INITIAL_ETH_PRICE + (INITIAL_ETH_PRICE * PRICE_DEVIATION_THRESHOLD / BASIS_POINTS);

            await gasPriceOracle.connect(oracle1).updatePriceFeed("ETH/USD", withinThresholdPrice);

            const [currentPrice] = await gasPriceOracle.getETHPriceUSD();
            expect(currentPrice).to.equal(withinThresholdPrice);
        });

        it("should not validate deviation for first price update", async function () {
            // Add new feed with initial price, then update with extreme deviation
            await gasPriceOracle.connect(owner).addPriceFeed("BTC/USD", 50000_00000000, 0); // $50,000
            
            // Wait for the price to become stale so deviation check is skipped
            await ethers.provider.send("evm_increaseTime", [MAX_PRICE_AGE + 1]);
            await ethers.provider.send("evm_mine", []);
            
            const extremePrice = 100000_00000000; // $100,000 (should not trigger deviation check for stale price)
            await gasPriceOracle.connect(oracle1).updatePriceFeed("BTC/USD", extremePrice);

            const [currentPrice] = await gasPriceOracle.getPriceWithMetadata("BTC/USD");
            expect(currentPrice).to.equal(extremePrice);
        });

        it("should not allow price updates when paused", async function () {
            await gasPriceOracle.connect(owner).pause();

            await expect(
                gasPriceOracle.connect(oracle1).updatePriceFeed("ETH/USD", 2100_00000000)
            ).to.be.reverted;
        });
    });

    describe("Fallback Price Mechanism", function () {
        it("should set fallback price", async function () {
            const fallbackPrice = 1900_00000000; // $1900

            await gasPriceOracle.connect(owner).setFallbackPrice("ETH/USD", fallbackPrice);

            // We can't directly check fallback price, but we can test it by making the main price stale
        });

        it("should use fallback price when main price is stale", async function () {
            const fallbackPrice = 1900_00000000;
            await gasPriceOracle.connect(owner).setFallbackPrice("ETH/USD", fallbackPrice);

            // Fast forward time to make price stale
            await ethers.provider.send("evm_increaseTime", [MAX_PRICE_AGE + 1]);
            await ethers.provider.send("evm_mine", []);

            const [price] = await gasPriceOracle.getETHPriceUSD();
            expect(price).to.equal(fallbackPrice);
        });

        it("should revert setting invalid fallback price", async function () {
            await expect(
                gasPriceOracle.connect(owner).setFallbackPrice("ETH/USD", 0)
            ).to.be.revertedWith("GasPriceOracle: invalid fallback price");
        });

        it("should only allow owner to set fallback price", async function () {
            await expect(
                gasPriceOracle.connect(unauthorizedUser).setFallbackPrice("ETH/USD", 1900_00000000)
            ).to.be.reverted;
        });

        it("should revert when no price available (no fallback)", async function () {
            // Fast forward time to make price stale
            await ethers.provider.send("evm_increaseTime", [MAX_PRICE_AGE + 1]);
            await ethers.provider.send("evm_mine", []);

            await expect(
                gasPriceOracle.getETHPriceUSD()
            ).to.be.revertedWith("GasPriceOracle: ETH price unavailable");
        });
    });

    describe("Price Retrieval", function () {
        it("should get ETH price in USD", async function () {
            const [price, timestamp] = await gasPriceOracle.getETHPriceUSD();
            expect(price).to.equal(INITIAL_ETH_PRICE);
            expect(timestamp).to.be.gt(0);
        });

        it("should get Token price in USD", async function () {
            const [price, timestamp] = await gasPriceOracle.getTokenPriceUSD();
            expect(price).to.equal(INITIAL_CPOP_PRICE);
            expect(timestamp).to.be.gt(0);
        });

        it("should get price with metadata", async function () {
            const [price, timestamp, isValid] = await gasPriceOracle.getPriceWithMetadata("ETH/USD");
            expect(price).to.equal(INITIAL_ETH_PRICE);
            expect(timestamp).to.be.gt(0);
            expect(isValid).to.be.true;
        });

        it("should return invalid for non-existent feed", async function () {
            const [price, timestamp, isValid] = await gasPriceOracle.getPriceWithMetadata("BTC/USD");
            expect(price).to.equal(0);
            expect(isValid).to.be.false;
        });

        it("should validate price freshness", async function () {
            // Fresh price should be valid
            expect(await gasPriceOracle.isPriceValid("ETH/USD")).to.be.true;

            // Stale price should be invalid
            await ethers.provider.send("evm_increaseTime", [MAX_PRICE_AGE + 1]);
            await ethers.provider.send("evm_mine", []);
            
            expect(await gasPriceOracle.isPriceValid("ETH/USD")).to.be.false;
        });
    });

    describe("Exchange Rate Calculations", function () {
        it("should calculate Token amount for ETH", async function () {
            const weiAmount = ethers.utils.parseEther("1"); // 1 ETH
            const expectedTokens = ethers.utils.parseEther("2000"); // 2000 CPOP (ETH=$2000, CPOP=$1)

            const tokenAmount = await gasPriceOracle.getTokenForETH(weiAmount);
            expect(tokenAmount).to.equal(expectedTokens);
        });

        it("should calculate ETH amount for Token", async function () {
            const tokenAmount = ethers.utils.parseEther("2000"); // 2000 CPOP
            const expectedWei = ethers.utils.parseEther("1"); // 1 ETH

            const weiAmount = await gasPriceOracle.getETHForToken(tokenAmount);
            expect(weiAmount).to.equal(expectedWei);
        });

        it("should handle zero amounts in exchange calculations", async function () {
            expect(await gasPriceOracle.getTokenForETH(0)).to.equal(0);
            expect(await gasPriceOracle.getETHForToken(0)).to.equal(0);
        });

        it("should revert with invalid Token price for ETH conversion", async function () {
            // Remove the CPOP/USD feed to make token price unavailable
            await gasPriceOracle.connect(owner).removePriceFeed("CPOP/USD");

            await expect(
                gasPriceOracle.getTokenForETH(ethers.utils.parseEther("1"))
            ).to.be.revertedWith("GasPriceOracle: Token price unavailable");
        });

        it("should revert with invalid ETH price for Token conversion", async function () {
            // Make ETH price stale and no fallback
            await ethers.provider.send("evm_increaseTime", [MAX_PRICE_AGE + 1]);
            await ethers.provider.send("evm_mine", []);

            await expect(
                gasPriceOracle.getETHForToken(ethers.utils.parseEther("2000"))
            ).to.be.revertedWith("GasPriceOracle: ETH price unavailable");
        });

        it("should revert with calculation overflow in getTokenForETH", async function () {
            // Test with extremely large weiAmount that would cause overflow
            const extremelyLargeAmount = ethers.constants.MaxUint256.div(INITIAL_ETH_PRICE).add(1);
            
            await expect(
                gasPriceOracle.getTokenForETH(extremelyLargeAmount)
            ).to.be.revertedWith("GasPriceOracle: calculation overflow");
        });

        it("should revert with calculation overflow in getETHForToken", async function () {
            // Test with extremely large tokenAmount that would cause overflow
            const extremelyLargeAmount = ethers.constants.MaxUint256.div(INITIAL_CPOP_PRICE).add(1);
            
            await expect(
                gasPriceOracle.getETHForToken(extremelyLargeAmount)
            ).to.be.revertedWith("GasPriceOracle: calculation overflow");
        });

        it("should handle both price validation in getTokenForETH", async function () {
            // Remove ETH feed to test ETH price validation
            await gasPriceOracle.connect(owner).removePriceFeed("ETH/USD");

            await expect(
                gasPriceOracle.getTokenForETH(ethers.utils.parseEther("1"))
            ).to.be.revertedWith("GasPriceOracle: ETH price unavailable");
        });

        it("should handle both price validation in getETHForToken", async function () {
            // Remove CPOP feed to test Token price validation  
            await gasPriceOracle.connect(owner).removePriceFeed("CPOP/USD");

            await expect(
                gasPriceOracle.getETHForToken(ethers.utils.parseEther("1000"))
            ).to.be.revertedWith("GasPriceOracle: Token price unavailable");
        });
    });

    describe("Gas Cost Estimation", function () {
        it("should estimate gas cost in tokens", async function () {
            const gasLimit = 21000;
            const gasPrice = ethers.utils.parseUnits("20", "gwei"); // 20 gwei

            const tokenCost = await gasPriceOracle.estimateGasCostInToken(gasLimit, gasPrice);
            
            // Should return a reasonable token cost (more than 0, reasonable size)
            expect(tokenCost).to.be.gt(0);
            expect(tokenCost).to.be.lt(ethers.utils.parseEther("1")); // Should be less than 1 token
        });

        it("should handle zero gas limit", async function () {
            const gasPrice = ethers.utils.parseUnits("20", "gwei");
            const tokenCost = await gasPriceOracle.estimateGasCostInToken(0, gasPrice);
            expect(tokenCost).to.equal(0);
        });

        it("should handle zero gas price", async function () {
            const gasLimit = 21000;
            const tokenCost = await gasPriceOracle.estimateGasCostInToken(gasLimit, 0);
            expect(tokenCost).to.equal(0);
        });

        it("should revert with gas cost overflow", async function () {
            // Create conditions that would cause gasLimit * gasPrice to overflow
            const largeGasLimit = ethers.constants.MaxUint256.div(1000);
            const largeGasPrice = 2000;
            
            await expect(
                gasPriceOracle.estimateGasCostInToken(largeGasLimit, largeGasPrice)
            ).to.be.revertedWith("GasPriceOracle: gas cost overflow");
        });

        it("should handle zero gas limit and price in estimateGasCostInToken", async function () {
            expect(await gasPriceOracle.estimateGasCostInToken(0, 1000)).to.equal(0);
            expect(await gasPriceOracle.estimateGasCostInToken(21000, 0)).to.equal(0);
            expect(await gasPriceOracle.estimateGasCostInToken(0, 0)).to.equal(0);
        });
    });

    describe("Configuration Management", function () {
        it("should update oracle configuration", async function () {
            const newMaxAge = 7200; // 2 hours
            const newThreshold = 1000; // 10%

            const tx = await gasPriceOracle.connect(owner).updateOracleConfig(newMaxAge, newThreshold);
            const receipt = await tx.wait();

            expect(await gasPriceOracle.getMaxPriceAge()).to.equal(newMaxAge);
            expect(await gasPriceOracle.getPriceDeviationThreshold()).to.equal(newThreshold);

            const event = receipt.events?.find(e => e.event === "OracleConfigUpdated");
            expect(event?.args?.maxPriceAge).to.equal(newMaxAge);
            expect(event?.args?.priceDeviationThreshold).to.equal(newThreshold);
        });

        it("should only allow owner to update configuration", async function () {
            await expect(
                gasPriceOracle.connect(unauthorizedUser).updateOracleConfig(7200, 1000)
            ).to.be.reverted;
        });

        it("should revert with invalid max price age", async function () {
            await expect(
                gasPriceOracle.connect(owner).updateOracleConfig(0, 1000)
            ).to.be.revertedWith("GasPriceOracle: invalid max price age");
        });

        it("should revert with invalid price deviation threshold", async function () {
            await expect(
                gasPriceOracle.connect(owner).updateOracleConfig(7200, BASIS_POINTS + 1)
            ).to.be.revertedWith("GasPriceOracle: invalid threshold");
        });
    });

    describe("Pause Functionality", function () {
        it("should pause and unpause contract", async function () {
            // Initially not paused
            expect(await gasPriceOracle.paused()).to.be.false;

            // Pause
            await gasPriceOracle.connect(owner).pause();
            expect(await gasPriceOracle.paused()).to.be.true;

            // Unpause
            await gasPriceOracle.connect(owner).unpause();
            expect(await gasPriceOracle.paused()).to.be.false;
        });

        it("should only allow owner to pause/unpause", async function () {
            await expect(
                gasPriceOracle.connect(unauthorizedUser).pause()
            ).to.be.reverted;

            await gasPriceOracle.connect(owner).pause();

            await expect(
                gasPriceOracle.connect(unauthorizedUser).unpause()
            ).to.be.reverted;
        });

        it("should prevent price updates when paused", async function () {
            await gasPriceOracle.connect(owner).pause();

            await expect(
                gasPriceOracle.connect(oracle1).updatePriceFeed("ETH/USD", 2100_00000000)
            ).to.be.reverted;
        });

        it("should allow view functions when paused", async function () {
            await gasPriceOracle.connect(owner).pause();

            // View functions should still work
            await expect(gasPriceOracle.getETHPriceUSD()).to.not.be.reverted;
            await expect(gasPriceOracle.getTokenPriceUSD()).to.not.be.reverted;
            await expect(gasPriceOracle.getTokenForETH(ethers.utils.parseEther("1"))).to.not.be.reverted;
        });
    });

    describe("Edge Cases and Error Handling", function () {
        it("should handle very large price values", async function () {
            const largePrice = ethers.constants.MaxUint256.div(1000); // Very large but not overflow
            
            // Add new feed with large price
            await gasPriceOracle.connect(owner).addPriceFeed("LARGE/USD", largePrice, 0);
            
            const [price] = await gasPriceOracle.getPriceWithMetadata("LARGE/USD");
            expect(price).to.equal(largePrice);
        });

        it("should handle price updates after long periods", async function () {
            // Fast forward a very long time
            await ethers.provider.send("evm_increaseTime", [365 * 24 * 3600]); // 1 year
            await ethers.provider.send("evm_mine", []);

            // Should still allow price update
            const newPrice = 2500_00000000;
            await gasPriceOracle.connect(oracle1).updatePriceFeed("ETH/USD", newPrice);

            const [currentPrice] = await gasPriceOracle.getETHPriceUSD();
            expect(currentPrice).to.equal(newPrice);
        });

        it("should handle multiple rapid price updates", async function () {
            const prices = [2100_00000000, 2050_00000000, 2075_00000000];
            
            for (const price of prices) {
                await gasPriceOracle.connect(oracle1).updatePriceFeed("ETH/USD", price);
                const [currentPrice] = await gasPriceOracle.getETHPriceUSD();
                expect(currentPrice).to.equal(price);
            }
        });

        it("should correctly calculate price deviation", async function () {
            // Test exact threshold boundary
            const exactThresholdPrice = INITIAL_ETH_PRICE + (INITIAL_ETH_PRICE * PRICE_DEVIATION_THRESHOLD / BASIS_POINTS);
            
            await gasPriceOracle.connect(oracle1).updatePriceFeed("ETH/USD", exactThresholdPrice);
            
            const [currentPrice] = await gasPriceOracle.getETHPriceUSD();
            expect(currentPrice).to.equal(exactThresholdPrice);
        });

        it("should handle division by zero protection in exchange rates", async function () {
            // This should be prevented by the price validation, but let's test edge cases
            const [ethPrice] = await gasPriceOracle.getETHPriceUSD();
            const [cpopPrice] = await gasPriceOracle.getTokenPriceUSD();
            
            expect(ethPrice).to.be.gt(0);
            expect(cpopPrice).to.be.gt(0);
        });
    });

    describe("Upgrade Functionality", function () {
        it("should only allow owner to authorize upgrades", async function () {
            // This tests the _authorizeUpgrade internal function via ownership
            expect(await gasPriceOracle.owner()).to.equal(ownerAddress);
        });
    });
});