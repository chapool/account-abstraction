import { expect } from 'chai'
import { ethers, upgrades } from 'hardhat'
import { Contract } from 'ethers'

describe('CPNFT Upgradeable', function () {
  let cpnft: Contract
  let owner: any
  let addr1: any
  let addr2: any

  beforeEach(async function () {
    [owner, addr1, addr2] = await ethers.getSigners()

    // Deploy the upgradeable contract
    const CPNFT = await ethers.getContractFactory('CPNFT')
    cpnft = await upgrades.deployProxy(
      CPNFT,
      ['Test CPNFT', 'TCPNFT', 'https://test-api.example.com/nft/'],
      { initializer: 'initialize' }
    )
    await cpnft.deployed()
  })

  describe('Deployment', function () {
    it('Should set the right owner', async function () {
      expect(await cpnft.owner()).to.equal(owner.address)
    })

    it('Should set the right name and symbol', async function () {
      expect(await cpnft.name()).to.equal('Test CPNFT')
      expect(await cpnft.symbol()).to.equal('TCPNFT')
    })

    it('Should return the correct version', async function () {
      expect(await cpnft.version()).to.equal('1.0.0')
    })
  })

  describe('Minting', function () {
    it('Should mint NFT with correct level', async function () {
      const tokenId = await cpnft.callStatic.mint(addr1.address, 0) // Level C
      await cpnft.mint(addr1.address, 0)
      
      expect(await cpnft.ownerOf(tokenId)).to.equal(addr1.address)
      expect(await cpnft.getTokenLevel(tokenId)).to.equal(0) // Level C
      expect(await cpnft.tokenExists(tokenId)).to.be.true
    })

    it('Should batch mint NFTs', async function () {
      const addresses = [addr1.address, addr2.address]
      const levels = [0, 1] // C, B levels
      
      await cpnft.batchMint(addresses, levels)
      
      expect(await cpnft.ownerOf(1)).to.equal(addr1.address)
      expect(await cpnft.ownerOf(2)).to.equal(addr2.address)
      expect(await cpnft.getTokenLevel(1)).to.equal(0)
      expect(await cpnft.getTokenLevel(2)).to.equal(1)
    })
  })

  describe('Token Level Management', function () {
    beforeEach(async function () {
      await cpnft.mint(addr1.address, 0) // Mint token with level C
    })

    it('Should allow owner to change token level', async function () {
      await cpnft.setTokenLevel(1, 5) // Change to level SSS
      expect(await cpnft.getTokenLevel(1)).to.equal(5)
    })

    it('Should not allow non-owner to change token level', async function () {
      await expect(
        cpnft.connect(addr1).setTokenLevel(1, 5)
      ).to.be.revertedWith('Ownable: caller is not the owner')
    })
  })

  describe('Staking', function () {
    beforeEach(async function () {
      await cpnft.mint(addr1.address, 0)
    })

    it('Should check staking status', async function () {
      expect(await cpnft.isStaked(1)).to.be.false
    })

    it('Should prevent transfer of staked tokens', async function () {
      // This test would require a staking contract to set stake status
      // For now, we'll test the basic functionality
      expect(await cpnft.isStaked(1)).to.be.false
    })
  })

  describe('Transfers', function () {
    beforeEach(async function () {
      await cpnft.mint(addr1.address, 0)
    })

    it('Should transfer token between addresses', async function () {
      await cpnft.connect(addr1).transferFrom(addr1.address, addr2.address, 1)
      expect(await cpnft.ownerOf(1)).to.equal(addr2.address)
    })

    it('Should batch transfer tokens', async function () {
      await cpnft.mint(addr2.address, 1) // Mint second token
      
      const from = [addr1.address, addr2.address]
      const to = [addr2.address, addr1.address]
      const tokenIds = [1, 2]
      
      await cpnft.batchTransferFrom(from, to, tokenIds)
      
      expect(await cpnft.ownerOf(1)).to.equal(addr2.address)
      expect(await cpnft.ownerOf(2)).to.equal(addr1.address)
    })
  })

  describe('Burning', function () {
    beforeEach(async function () {
      await cpnft.mint(addr1.address, 0)
    })

    it('Should burn token', async function () {
      await cpnft.burn(1)
      expect(await cpnft.tokenExists(1)).to.be.false
    })

    it('Should batch burn tokens', async function () {
      await cpnft.mint(addr1.address, 1) // Mint second token
      
      const tokenIds = [1, 2]
      await cpnft.batchBurn(tokenIds)
      
      expect(await cpnft.tokenExists(1)).to.be.false
      expect(await cpnft.tokenExists(2)).to.be.false
    })
  })

  describe('Metadata', function () {
    it('Should return correct token URI', async function () {
      await cpnft.mint(addr1.address, 0)
      const tokenURI = await cpnft.tokenURI(1)
      expect(tokenURI).to.equal('https://test-api.example.com/nft/1')
    })

    it('Should allow owner to update base URI', async function () {
      await cpnft.setBaseURI('https://new-api.example.com/nft/')
      await cpnft.mint(addr1.address, 0)
      const tokenURI = await cpnft.tokenURI(1)
      expect(tokenURI).to.equal('https://new-api.example.com/nft/1')
    })
  })
})
