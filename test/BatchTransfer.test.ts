import { expect } from "chai";
import { ethers } from "hardhat";
import { BatchTransfer, CPNFT, MockUSDT } from "../typechain";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";

describe("BatchTransfer", function () {
  let batchTransfer: BatchTransfer;
  let cpnft: CPNFT;
  let mockToken: MockUSDT;
  let owner: SignerWithAddress;
  let user1: SignerWithAddress;
  let user2: SignerWithAddress;
  let user3: SignerWithAddress;
  let user4: SignerWithAddress;

  beforeEach(async function () {
    [owner, user1, user2, user3, user4] = await ethers.getSigners();

    // Deploy BatchTransfer
    const BatchTransferFactory = await ethers.getContractFactory("BatchTransfer");
    batchTransfer = await BatchTransferFactory.deploy();
    await batchTransfer.deployed();

    // Deploy CPNFT
    const CPNFTFactory = await ethers.getContractFactory("CPNFT");
    const cpnftImpl = await CPNFTFactory.deploy();
    await cpnftImpl.deployed();

    // Deploy proxy for CPNFT
    const ERC1967ProxyFactory = await ethers.getContractFactory("ERC1967Proxy");
    const initData = cpnftImpl.interface.encodeFunctionData("initialize", [
      "CP NFT",
      "CPNFT",
      "https://api.example.com/metadata/"
    ]);
    const proxy = await ERC1967ProxyFactory.deploy(cpnftImpl.address, initData);
    await proxy.deployed();
    
    cpnft = CPNFTFactory.attach(proxy.address);

    // Deploy MockUSDT
    const MockUSDTFactory = await ethers.getContractFactory("MockUSDT");
    mockToken = await MockUSDTFactory.deploy();
    await mockToken.deployed();

    // Mint some NFTs to user1
    await cpnft.mint(user1.address, 0); // tokenId 1
    await cpnft.mint(user1.address, 0); // tokenId 2
    await cpnft.mint(user1.address, 0); // tokenId 3
    await cpnft.mint(user1.address, 0); // tokenId 4

    // Mint some tokens to user1
    await mockToken.mint(user1.address, ethers.utils.parseEther("1000"));
  });

  describe("NFT Batch Transfer", function () {
    it("Should batch transfer NFTs to multiple recipients", async function () {
      // Approve BatchTransfer
      await cpnft.connect(user1).setApprovalForAll(batchTransfer.address, true);

      // Check approval
      const isApproved = await batchTransfer.isNFTApproved(cpnft.address, user1.address);
      expect(isApproved).to.be.true;

      // Batch transfer
      await batchTransfer.connect(user1).batchTransferNFT(
        cpnft.address,
        [user2.address, user3.address],
        [1, 2]
      );

      // Verify transfers
      expect(await cpnft.ownerOf(1)).to.equal(user2.address);
      expect(await cpnft.ownerOf(2)).to.equal(user3.address);
    });

    it("Should batch transfer multiple NFTs to single recipient", async function () {
      await cpnft.connect(user1).setApprovalForAll(batchTransfer.address, true);

      await batchTransfer.connect(user1).batchTransferNFTToSingle(
        cpnft.address,
        user2.address,
        [1, 2, 3]
      );

      expect(await cpnft.ownerOf(1)).to.equal(user2.address);
      expect(await cpnft.ownerOf(2)).to.equal(user2.address);
      expect(await cpnft.ownerOf(3)).to.equal(user2.address);
    });

    it("Should revert when arrays length mismatch", async function () {
      await cpnft.connect(user1).setApprovalForAll(batchTransfer.address, true);

      await expect(
        batchTransfer.connect(user1).batchTransferNFT(
          cpnft.address,
          [user2.address, user3.address],
          [1]
        )
      ).to.be.revertedWith("Arrays length mismatch");
    });

    it("Should revert when not token owner", async function () {
      await cpnft.connect(user1).setApprovalForAll(batchTransfer.address, true);

      await expect(
        batchTransfer.connect(user2).batchTransferNFT(
          cpnft.address,
          [user3.address],
          [1]
        )
      ).to.be.revertedWith("Not token owner");
    });
  });

  describe("Token Batch Transfer", function () {
    it("Should batch transfer tokens to multiple recipients with different amounts", async function () {
      const amount1 = ethers.utils.parseEther("100");
      const amount2 = ethers.utils.parseEther("200");

      // Approve BatchTransfer
      await mockToken.connect(user1).approve(batchTransfer.address, amount1.add(amount2));

      // Check allowance
      const hasAllowance = await batchTransfer.hasTokenAllowance(
        mockToken.address,
        user1.address,
        amount1.add(amount2)
      );
      expect(hasAllowance).to.be.true;

      // Batch transfer
      await batchTransfer.connect(user1).batchTransferToken(
        mockToken.address,
        [user2.address, user3.address],
        [amount1, amount2]
      );

      // Verify balances
      expect(await mockToken.balanceOf(user2.address)).to.equal(amount1);
      expect(await mockToken.balanceOf(user3.address)).to.equal(amount2);
    });

    it("Should batch transfer equal amounts to multiple recipients", async function () {
      const amountPerRecipient = ethers.utils.parseEther("100");
      const totalAmount = amountPerRecipient.mul(3);

      await mockToken.connect(user1).approve(batchTransfer.address, totalAmount);

      await batchTransfer.connect(user1).batchTransferTokenEqual(
        mockToken.address,
        [user2.address, user3.address, user4.address],
        amountPerRecipient
      );

      expect(await mockToken.balanceOf(user2.address)).to.equal(amountPerRecipient);
      expect(await mockToken.balanceOf(user3.address)).to.equal(amountPerRecipient);
      expect(await mockToken.balanceOf(user4.address)).to.equal(amountPerRecipient);
    });

    it("Should batch transfer multiple tokens to single recipient", async function () {
      // Deploy another mock token
      const MockUSDTFactory = await ethers.getContractFactory("MockUSDT");
      const mockToken2 = await MockUSDTFactory.deploy();
      await mockToken2.deployed();

      const amount1 = ethers.utils.parseEther("100");
      const amount2 = ethers.utils.parseEther("200");

      // Mint tokens
      await mockToken.mint(user1.address, amount1);
      await mockToken2.mint(user1.address, amount2);

      // Approve both tokens
      await mockToken.connect(user1).approve(batchTransfer.address, amount1);
      await mockToken2.connect(user1).approve(batchTransfer.address, amount2);

      // Batch transfer
      await batchTransfer.connect(user1).batchTransferMultipleTokens(
        [mockToken.address, mockToken2.address],
        user2.address,
        [amount1, amount2]
      );

      expect(await mockToken.balanceOf(user2.address)).to.equal(amount1);
      expect(await mockToken2.balanceOf(user2.address)).to.equal(amount2);
    });

    it("Should revert when arrays length mismatch", async function () {
      await mockToken.connect(user1).approve(batchTransfer.address, ethers.utils.parseEther("1000"));

      await expect(
        batchTransfer.connect(user1).batchTransferToken(
          mockToken.address,
          [user2.address, user3.address],
          [ethers.utils.parseEther("100")]
        )
      ).to.be.revertedWith("Arrays length mismatch");
    });
  });

  describe("Combined Batch Transfer", function () {
    it("Should transfer both NFTs and tokens in single transaction", async function () {
      const tokenAmount1 = ethers.utils.parseEther("100");
      const tokenAmount2 = ethers.utils.parseEther("200");

      // Approve both NFTs and tokens
      await cpnft.connect(user1).setApprovalForAll(batchTransfer.address, true);
      await mockToken.connect(user1).approve(batchTransfer.address, tokenAmount1.add(tokenAmount2));

      // Combined batch transfer
      await batchTransfer.connect(user1).combinedBatchTransfer(
        cpnft.address,
        [user2.address, user3.address],
        [1, 2],
        mockToken.address,
        [user2.address, user3.address],
        [tokenAmount1, tokenAmount2]
      );

      // Verify NFT transfers
      expect(await cpnft.ownerOf(1)).to.equal(user2.address);
      expect(await cpnft.ownerOf(2)).to.equal(user3.address);

      // Verify token transfers
      expect(await mockToken.balanceOf(user2.address)).to.equal(tokenAmount1);
      expect(await mockToken.balanceOf(user3.address)).to.equal(tokenAmount2);
    });

    it("Should transfer only NFTs when token contract is zero address", async function () {
      await cpnft.connect(user1).setApprovalForAll(batchTransfer.address, true);

      await batchTransfer.connect(user1).combinedBatchTransfer(
        cpnft.address,
        [user2.address, user3.address],
        [1, 2],
        ethers.constants.AddressZero,
        [],
        []
      );

      expect(await cpnft.ownerOf(1)).to.equal(user2.address);
      expect(await cpnft.ownerOf(2)).to.equal(user3.address);
    });

    it("Should transfer only tokens when NFT contract is zero address", async function () {
      const tokenAmount = ethers.utils.parseEther("100");
      await mockToken.connect(user1).approve(batchTransfer.address, tokenAmount.mul(2));

      await batchTransfer.connect(user1).combinedBatchTransfer(
        ethers.constants.AddressZero,
        [],
        [],
        mockToken.address,
        [user2.address, user3.address],
        [tokenAmount, tokenAmount]
      );

      expect(await mockToken.balanceOf(user2.address)).to.equal(tokenAmount);
      expect(await mockToken.balanceOf(user3.address)).to.equal(tokenAmount);
    });
  });

  describe("View Functions", function () {
    it("Should correctly check NFT approval status", async function () {
      expect(await batchTransfer.isNFTApproved(cpnft.address, user1.address)).to.be.false;

      await cpnft.connect(user1).setApprovalForAll(batchTransfer.address, true);

      expect(await batchTransfer.isNFTApproved(cpnft.address, user1.address)).to.be.true;
    });

    it("Should correctly check token allowance", async function () {
      const amount = ethers.utils.parseEther("100");

      expect(await batchTransfer.hasTokenAllowance(mockToken.address, user1.address, amount)).to.be.false;

      await mockToken.connect(user1).approve(batchTransfer.address, amount);

      expect(await batchTransfer.hasTokenAllowance(mockToken.address, user1.address, amount)).to.be.true;
    });
  });
});

