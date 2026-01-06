import { expect } from "chai";
import { ethers, upgrades } from "hardhat";

describe("Marketplace", function () {
  let deployer: any;
  let seller: any;
  let buyer: any;
  let bidder1: any;
  let bidder2: any;

  let cpnft: any;
  let marketplace: any;
  let usdt: any;

  const platformFeeRate = 500; // 5%
  const delistingLimit = 3;
  const delistingWindow = 24 * 60 * 60;

  beforeEach(async () => {
    [deployer, seller, buyer, bidder1, bidder2] = await ethers.getSigners();

    const CPNFT = await ethers.getContractFactory("CPNFT");
    cpnft = await upgrades.deployProxy(
      CPNFT,
      ["CPNFT", "CPN", "https://metadata/"],
      { initializer: "initialize", kind: "uups" }
    );
    await cpnft.deployed();

    const MockUSDT = await ethers.getContractFactory("MockUSDT");
    usdt = await MockUSDT.connect(deployer).deploy();
    await usdt.deployed();

    const Marketplace = await ethers.getContractFactory("Marketplace");
    marketplace = await upgrades.deployProxy(
      Marketplace,
      [
        cpnft.address,
        usdt.address,
        deployer.address,
        platformFeeRate,
        delistingLimit,
        delistingWindow,
        deployer.address,
      ],
      { initializer: "initialize", kind: "uups" }
    );
    await marketplace.deployed();

    await cpnft.connect(deployer).setMarketplaceContract(marketplace.address);
  });

  async function mintToSeller(level = 0): Promise<number> {
    const tx = await cpnft.connect(deployer).mint(seller.address, level);
    const receipt = await tx.wait();
    const transferEvt = receipt.events?.find((e: any) => e.event === "Transfer");
    const tokenId = transferEvt?.args?.tokenId?.toNumber() ?? 1;
    return tokenId;
  }

  it("初始化参数应正确设置", async () => {
    expect(await marketplace.nftContract()).to.equal(cpnft.address);
    expect(await marketplace.paymentToken()).to.equal(usdt.address);
    expect(await marketplace.platformFeeRecipient()).to.equal(deployer.address);
    expect(await marketplace.platformFeeRate()).to.equal(platformFeeRate);
    expect(await marketplace.delistingLimit()).to.equal(delistingLimit);
    expect(await marketplace.delistingWindow()).to.equal(delistingWindow);
  });

  it("一口价上架与购买流程", async () => {
    const price = ethers.utils.parseEther("100");
    const tokenId = await mintToSeller();

    const listingType = 0; // FixedPrice
    const tx = await marketplace
      .connect(deployer)
      .createListing(seller.address, tokenId, listingType, price, 0, 0);
    const receipt = await tx.wait();
    const createdEvt = receipt.events?.find((e: any) => e.event === "ListingCreated");
    const listingId = createdEvt?.args?.listingId?.toNumber();
    expect(listingId).to.be.a("number");

    expect(await cpnft.ownerOf(tokenId)).to.equal(marketplace.address);

    await usdt.connect(buyer).faucet();
    await usdt.connect(buyer).approve(marketplace.address, price);

    const sellerBalanceBefore = await usdt.balanceOf(seller.address);
    const platformBalanceBefore = await usdt.balanceOf(deployer.address);

    await expect(marketplace.connect(deployer).buyItem(listingId, buyer.address, buyer.address))
      .to.emit(marketplace, "ItemSold")
      .withArgs(
        listingId,
        seller.address,
        buyer.address,
        buyer.address,
        price,
        price.mul(platformFeeRate).div(10000)
      );

    expect(await cpnft.ownerOf(tokenId)).to.equal(buyer.address);

    const platformFee = price.mul(platformFeeRate).div(10000);
    const sellerAmount = price.sub(platformFee);

    const sellerBalanceAfter = await usdt.balanceOf(seller.address);
    const platformBalanceAfter = await usdt.balanceOf(deployer.address);

    expect(sellerBalanceAfter.sub(sellerBalanceBefore)).to.equal(sellerAmount);
    expect(platformBalanceAfter.sub(platformBalanceBefore)).to.equal(platformFee);
  });

  it("一口价购买由第三方支付", async () => {
    const price = ethers.utils.parseEther("100");
    const tokenId = await mintToSeller();
    const listingType = 0;
    const tx = await marketplace
      .connect(deployer)
      .createListing(seller.address, tokenId, listingType, price, 0, 0);
    const receipt = await tx.wait();
    const listingId = receipt.events?.find((e: any) => e.event === "ListingCreated")?.args?.listingId?.toNumber();
    expect(await cpnft.ownerOf(tokenId)).to.equal(marketplace.address);
    await usdt.connect(bidder1).faucet();
    await usdt.connect(bidder1).approve(marketplace.address, price);
    const sellerBefore = await usdt.balanceOf(seller.address);
    const platformBefore = await usdt.balanceOf(deployer.address);
    const payerBefore = await usdt.balanceOf(bidder1.address);
    await expect(marketplace.connect(deployer).buyItem(listingId, buyer.address, bidder1.address))
      .to.emit(marketplace, "ItemSold")
      .withArgs(
        listingId,
        seller.address,
        buyer.address,
        bidder1.address,
        price,
        price.mul(platformFeeRate).div(10000)
      );
    expect(await cpnft.ownerOf(tokenId)).to.equal(buyer.address);
    const fee = price.mul(platformFeeRate).div(10000);
    const sellerAmount = price.sub(fee);
    const sellerAfter = await usdt.balanceOf(seller.address);
    const platformAfter = await usdt.balanceOf(deployer.address);
    const payerAfter = await usdt.balanceOf(bidder1.address);
    expect(sellerAfter.sub(sellerBefore)).to.equal(sellerAmount);
    expect(platformAfter.sub(platformBefore)).to.equal(fee);
    expect(payerBefore.sub(payerAfter)).to.equal(price);
  });

  it("拍卖加价、退款与结算流程", async () => {
    const startPrice = ethers.utils.parseEther("50");
    const tokenId = await mintToSeller();

    const listingType = 1; // Auction
    const duration = 3600;
    const minBidIncBps = 1000; // 10%

    const tx = await marketplace
      .connect(deployer)
      .createListing(seller.address, tokenId, listingType, startPrice, duration, minBidIncBps);
    const receipt = await tx.wait();
    const createdEvt = receipt.events?.find((e: any) => e.event === "ListingCreated");
    const listingId = createdEvt?.args?.listingId?.toNumber();
    expect(listingId).to.be.a("number");

    await usdt.connect(bidder1).faucet();
    await usdt.connect(bidder2).faucet();
    await usdt.connect(bidder1).approve(marketplace.address, ethers.utils.parseEther("1000"));
    await usdt.connect(bidder2).approve(marketplace.address, ethers.utils.parseEther("1000"));

    const bid1 = ethers.utils.parseEther("60");
    await expect(marketplace.connect(deployer).placeBid(listingId, bidder1.address, bidder1.address, bid1))
      .to.emit(marketplace, "BidPlaced")
      .withArgs(listingId, bidder1.address, bidder1.address, bid1, ethers.constants.AddressZero, 0);

    const minNextBid = bid1.add(bid1.mul(minBidIncBps).div(10000)); // 66
    const bid2 = minNextBid;
    await expect(marketplace.connect(deployer).placeBid(listingId, bidder2.address, bidder2.address, bid2))
      .to.emit(marketplace, "BidRefunded")
      .withArgs(listingId, bidder1.address, bidder1.address, bid1)
      .and.to.emit(marketplace, "BidPlaced")
      .withArgs(listingId, bidder2.address, bidder2.address, bid2, bidder1.address, bid1);

    const bidder1BalanceAfterRefund = await usdt.balanceOf(bidder1.address);
    expect(bidder1BalanceAfterRefund).to.be.gte(bid1);

    await marketplace.connect(deployer).enableTestMode((await ethers.provider.getBlock("latest")).timestamp);
    const currentTs = await marketplace.getCurrentTimestamp();
    await marketplace.connect(deployer).setTestTimestamp(currentTs.add(duration));

    const sellerBalanceBefore = await usdt.balanceOf(seller.address);
    const platformBalanceBefore = await usdt.balanceOf(deployer.address);

    await expect(marketplace.connect(deployer).settleAuction(listingId, ethers.constants.AddressZero))
      .to.emit(marketplace, "AuctionSettled")
      .withArgs(
        listingId,
        seller.address,
        bidder2.address,
        bidder2.address,
        bid2,
        bid2.mul(platformFeeRate).div(10000)
      );

    expect(await cpnft.ownerOf(tokenId)).to.equal(bidder2.address);

    const platformFee = bid2.mul(platformFeeRate).div(10000);
    const sellerAmount = bid2.sub(platformFee);
    const sellerBalanceAfter = await usdt.balanceOf(seller.address);
    const platformBalanceAfter = await usdt.balanceOf(deployer.address);
    expect(sellerBalanceAfter.sub(sellerBalanceBefore)).to.equal(sellerAmount);
    expect(platformBalanceAfter.sub(platformBalanceBefore)).to.equal(platformFee);
  });

  it("拍卖出价由第三方支付与退款回第三方", async () => {
    const startPrice = ethers.utils.parseEther("50");
    const tokenId = await mintToSeller();
    const listingType = 1;
    const duration = 3600;
    const minBidIncBps = 1000;
    const tx = await marketplace
      .connect(deployer)
      .createListing(seller.address, tokenId, listingType, startPrice, duration, minBidIncBps);
    const receipt = await tx.wait();
    const listingId = receipt.events?.find((e: any) => e.event === "ListingCreated")?.args?.listingId?.toNumber();
    await usdt.connect(buyer).faucet();
    await usdt.connect(bidder2).faucet();
    await usdt.connect(buyer).approve(marketplace.address, ethers.utils.parseEther("1000"));
    await usdt.connect(bidder2).approve(marketplace.address, ethers.utils.parseEther("1000"));
    const bid1 = ethers.utils.parseEther("60");
    const payer1Before = await usdt.balanceOf(buyer.address);
    await expect(marketplace.connect(deployer).placeBid(listingId, bidder1.address, buyer.address, bid1))
      .to.emit(marketplace, "BidPlaced")
      .withArgs(listingId, bidder1.address, buyer.address, bid1, ethers.constants.AddressZero, 0);
    const minNextBid = bid1.add(bid1.mul(minBidIncBps).div(10000));
    const bid2 = minNextBid;
    await expect(marketplace.connect(deployer).placeBid(listingId, bidder2.address, bidder2.address, bid2))
      .to.emit(marketplace, "BidRefunded")
      .withArgs(listingId, bidder1.address, buyer.address, bid1)
      .and.to.emit(marketplace, "BidPlaced")
      .withArgs(listingId, bidder2.address, bidder2.address, bid2, bidder1.address, bid1);
    const payer1After = await usdt.balanceOf(buyer.address);
    expect(payer1After.sub(payer1Before)).to.equal(0);
    await marketplace.connect(deployer).enableTestMode((await ethers.provider.getBlock("latest")).timestamp);
    const currentTs = await marketplace.getCurrentTimestamp();
    await marketplace.connect(deployer).setTestTimestamp(currentTs.add(duration));
    const sellerBefore = await usdt.balanceOf(seller.address);
    const platformBefore = await usdt.balanceOf(deployer.address);
    await expect(marketplace.connect(deployer).settleAuction(listingId, ethers.constants.AddressZero))
      .to.emit(marketplace, "AuctionSettled")
      .withArgs(
        listingId,
        seller.address,
        bidder2.address,
        bidder2.address,
        bid2,
        bid2.mul(platformFeeRate).div(10000)
      );
    expect(await cpnft.ownerOf(tokenId)).to.equal(bidder2.address);
    const fee = bid2.mul(platformFeeRate).div(10000);
    const sellerAmount = bid2.sub(fee);
    const sellerAfter = await usdt.balanceOf(seller.address);
    const platformAfter = await usdt.balanceOf(deployer.address);
    expect(sellerAfter.sub(sellerBefore)).to.equal(sellerAmount);
    expect(platformAfter.sub(platformBefore)).to.equal(fee);
  });
  it("取消上架频率限制", async () => {
    const price = ethers.utils.parseEther("10");

    const makeListing = async () => {
      const tokenId = await mintToSeller();
      const listingType = 0;
      const tx = await marketplace
        .connect(deployer)
        .createListing(seller.address, tokenId, listingType, price, 0, 0);
      const receipt = await tx.wait();
      const createdEvt = receipt.events?.find((e: any) => e.event === "ListingCreated");
      const listingId = createdEvt?.args?.listingId?.toNumber();
      expect(listingId).to.be.a("number");
      return listingId!;
    };

    const l1 = await makeListing();
    await marketplace.connect(deployer).cancelListing(l1);

    const l2 = await makeListing();
    await marketplace.connect(deployer).cancelListing(l2);

    const l3 = await makeListing();
    await marketplace.connect(deployer).cancelListing(l3);

    const l4 = await makeListing();
    await expect(marketplace.connect(deployer).cancelListing(l4)).to.be.revertedWith(
      "Delisting limit exceeded"
    );
  });

  it("紧急提取 NFT 使用初始化的合约地址", async () => {
    const tokenId = await mintToSeller();
    const price = ethers.utils.parseEther("1");
    const listingType = 0;
    const tx = await marketplace
      .connect(deployer)
      .createListing(seller.address, tokenId, listingType, price, 0, 0);
    await tx.wait();

    expect(await cpnft.ownerOf(tokenId)).to.equal(marketplace.address);
    await marketplace.connect(deployer).emergencyWithdrawNFT(tokenId, buyer.address);
    expect(await cpnft.ownerOf(tokenId)).to.equal(buyer.address);
  });

  it("非 owner 无法调用受限方法", async () => {
    const tokenId = await mintToSeller();
    const price = ethers.utils.parseEther("1");
    await expect(
      marketplace
        .connect(seller)
        .createListing(seller.address, tokenId, 0, price, 0, 0)
    ).to.be.reverted;
  });

  it("createListing 参数校验与拍卖参数校验", async () => {
    const tokenId = await mintToSeller();
    await expect(
      marketplace.connect(deployer).createListing(seller.address, tokenId, 0, 0, 0, 0)
    ).to.be.revertedWith("Price must be greater than 0");

    const tokenId2 = await mintToSeller();
    await expect(
      marketplace
        .connect(deployer)
        .createListing(seller.address, tokenId2, 1, ethers.utils.parseEther("1"), 0, 100)
    ).to.be.revertedWith("Auction duration must be greater than 0");

    const tokenId3 = await mintToSeller();
    await expect(
      marketplace
        .connect(deployer)
        .createListing(seller.address, tokenId3, 1, ethers.utils.parseEther("1"), 3600, 0)
    ).to.be.revertedWith("Min bid increment must be greater than 0");
  });

  it("placeBid 金额校验与最小加价幅度校验", async () => {
    const startPrice = ethers.utils.parseEther("10");
    const tokenId = await mintToSeller();
    const listingType = 1;
    const duration = 3600;
    const minBidIncBps = 1000; // 10%
    const tx = await marketplace
      .connect(deployer)
      .createListing(seller.address, tokenId, listingType, startPrice, duration, minBidIncBps);
    const receipt = await tx.wait();
    const createdEvt = receipt.events?.find((e: any) => e.event === "ListingCreated");
    const listingId = createdEvt?.args?.listingId?.toNumber();

    await usdt.connect(bidder1).faucet();
    await usdt.connect(bidder1).approve(marketplace.address, ethers.utils.parseEther("1000"));

    await expect(
      marketplace.connect(deployer).placeBid(listingId, bidder1.address, bidder1.address, ethers.utils.parseEther("9"))
    ).to.be.revertedWith("Bid must be at least starting price");

    const bid1 = ethers.utils.parseEther("10");
    await marketplace.connect(deployer).placeBid(listingId, bidder1.address, bidder1.address, bid1);

    await usdt.connect(bidder2).faucet();
    await usdt.connect(bidder2).approve(marketplace.address, ethers.utils.parseEther("1000"));

    const tooLow = ethers.utils.parseEther("10.5");
    await expect(
      marketplace.connect(deployer).placeBid(listingId, bidder2.address, bidder2.address, tooLow)
    ).to.be.revertedWith("Bid amount too low");
  });

  it("getHighestBid 无出价时抛错", async () => {
    const tokenId = await mintToSeller();
    const tx = await marketplace
      .connect(deployer)
      .createListing(seller.address, tokenId, 1, ethers.utils.parseEther("1"), 3600, 100);
    const receipt = await tx.wait();
    const createdEvt = receipt.events?.find((e: any) => e.event === "ListingCreated");
    const listingId = createdEvt?.args?.listingId?.toNumber();
    await expect(marketplace.getHighestBid(listingId)).to.be.revertedWith("No bids placed");
  });

  it("pause 阻止上架与购买与出价", async () => {
    await marketplace.connect(deployer).pause();
    const tokenId = await mintToSeller();
    await expect(
      marketplace.connect(deployer).createListing(seller.address, tokenId, 0, ethers.utils.parseEther("1"), 0, 0)
    ).to.be.reverted;
    await marketplace.connect(deployer).unpause();

    const tx = await marketplace
      .connect(deployer)
      .createListing(seller.address, tokenId, 0, ethers.utils.parseEther("1"), 0, 0);
    const receipt = await tx.wait();
    const listingId = receipt.events?.find((e: any) => e.event === "ListingCreated")?.args?.listingId?.toNumber();
    await marketplace.connect(deployer).pause();
    await expect(marketplace.connect(deployer).buyItem(listingId, buyer.address, buyer.address)).to.be.reverted;
    await marketplace.connect(deployer).unpause();
  });

  it("emergencyWithdraw 提取 ERC20 到指定地址", async () => {
    const amount = ethers.utils.parseEther("100");
    await usdt.connect(deployer).mint(marketplace.address, amount);
    const before = await usdt.balanceOf(buyer.address);
    await marketplace.connect(deployer).emergencyWithdraw(usdt.address, buyer.address, amount);
    const after = await usdt.balanceOf(buyer.address);
    expect(after.sub(before)).to.equal(amount);
  });

  it("getDelistingRecord 窗口重置逻辑", async () => {
    const price = ethers.utils.parseEther("1");
    const tokenId = await mintToSeller();
    const tx = await marketplace
      .connect(deployer)
      .createListing(seller.address, tokenId, 0, price, 0, 0);
    const receipt = await tx.wait();
    const listingId = receipt.events?.find((e: any) => e.event === "ListingCreated")?.args?.listingId?.toNumber();
    await marketplace.connect(deployer).cancelListing(listingId);

    const rec1 = await marketplace.getDelistingRecord(seller.address);
    expect(rec1.count).to.equal(1);

    const now = (await marketplace.getCurrentTimestamp()).toNumber();
    await marketplace.connect(deployer).enableTestMode(now);
    await marketplace.connect(deployer).setTestTimestamp(now + delistingWindow + 1);

    const rec2 = await marketplace.getDelistingRecord(seller.address);
    expect(rec2.count).to.equal(0);
    expect(rec2.remainingCount).to.equal(delistingLimit);
  });

  it("updateNFTContract 切换合约后继续上架", async () => {
    const CPNFT = await ethers.getContractFactory("CPNFT");
    const cpnft2 = await upgrades.deployProxy(
      CPNFT,
      ["CPNFT2", "CPN2", "https://metadata2/"],
      { initializer: "initialize", kind: "uups" }
    );
    await cpnft2.deployed();
    await cpnft2.connect(deployer).setMarketplaceContract(marketplace.address);

    await marketplace.connect(deployer).updateNFTContract(cpnft2.address);
    const tokenId = await (async () => {
      const tx = await cpnft2.connect(deployer).mint(seller.address, 0);
      const receipt = await tx.wait();
      const evt = receipt.events?.find((e: any) => e.event === "Transfer");
      return evt?.args?.tokenId?.toNumber();
    })();

    const tx = await marketplace
      .connect(deployer)
      .createListing(seller.address, tokenId, 0, ethers.utils.parseEther("1"), 0, 0);
    const receipt = await tx.wait();
    const listing = await marketplace.getListing(
      receipt.events?.find((e: any) => e.event === "ListingCreated")?.args?.listingId?.toNumber()
    );
    expect(listing.nftContract).to.equal(cpnft2.address);
    expect(await cpnft2.ownerOf(tokenId)).to.equal(marketplace.address);
  });
});
