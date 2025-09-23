import { expect } from 'chai'
import { ethers, upgrades } from 'hardhat'
import { Contract } from 'ethers'

// Mock V2 contract for testing upgrades
const CPNFTUpgradeableV2Source = `
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "./CPNFTUpgradeable.sol";

contract CPNFTUpgradeableV2 is CPNFTUpgradeable {
    // New storage variable (added at the end to maintain compatibility)
    mapping(uint256 => string) private _tokenMetadata;
    
    // New event
    event TokenMetadataSet(uint256 indexed tokenId, string metadata);
    
    // New function to set token metadata
    function setTokenMetadata(uint256 tokenId, string memory metadata) external onlyOwner {
        require(_mintedTokens[tokenId], "Token does not exist");
        _tokenMetadata[tokenId] = metadata;
        emit TokenMetadataSet(tokenId, metadata);
    }
    
    // New function to get token metadata
    function getTokenMetadata(uint256 tokenId) external view returns (string memory) {
        require(_mintedTokens[tokenId], "Token does not exist");
        return _tokenMetadata[tokenId];
    }
    
    // Override version function
    function version() public pure override returns (string memory) {
        return "2.0.0";
    }
}
`

describe('CPNFT Upgrade Tests', function () {
  let cpnft: Contract
  let cpnftV2: Contract
  let owner: any
  let addr1: any

  beforeEach(async function () {
    [owner, addr1] = await ethers.getSigners()

    // Deploy the initial upgradeable contract
    const CPNFT = await ethers.getContractFactory('CPNFT')
    cpnft = await upgrades.deployProxy(
      CPNFT,
      ['Test CPNFT', 'TCPNFT', 'https://test-api.example.com/nft/'],
      { initializer: 'initialize' }
    )
    await cpnft.deployed()

    // Mint some tokens for testing
    await cpnft.mint(addr1.address, 0) // Level C
    await cpnft.mint(addr1.address, 1) // Level B
  })

  describe('Upgrade Process', function () {
    it('Should upgrade to V2 successfully', async function () {
      // Verify initial state
      expect(await cpnft.version()).to.equal('1.0.0')
      expect(await cpnft.ownerOf(1)).to.equal(addr1.address)
      expect(await cpnft.getTokenLevel(1)).to.equal(0)

      // Deploy V2 implementation
      const CPNFTUpgradeableV2 = await ethers.getContractFactory('CPNFTUpgradeableV2')
      const cpnftV2Impl = await CPNFTUpgradeableV2.deploy()
      await cpnftV2Impl.deployed()

      // Upgrade the proxy
      await upgrades.upgradeProxy(cpnft.address, CPNFTUpgradeableV2)

      // Get the upgraded contract instance
      cpnftV2 = await ethers.getContractAt('CPNFTUpgradeableV2', cpnft.address)

      // Verify upgrade
      expect(await cpnftV2.version()).to.equal('2.0.0')
      expect(await cpnftV2.ownerOf(1)).to.equal(addr1.address)
      expect(await cpnftV2.getTokenLevel(1)).to.equal(0)
    })

    it('Should preserve existing data after upgrade', async function () {
      // Upgrade to V2
      const CPNFTUpgradeableV2 = await ethers.getContractFactory('CPNFTUpgradeableV2')
      await upgrades.upgradeProxy(cpnft.address, CPNFTUpgradeableV2)
      
      cpnftV2 = await ethers.getContractAt('CPNFTUpgradeableV2', cpnft.address)

      // Verify all existing data is preserved
      expect(await cpnftV2.name()).to.equal('Test CPNFT')
      expect(await cpnftV2.symbol()).to.equal('TCPNFT')
      expect(await cpnftV2.owner()).to.equal(owner.address)
      expect(await cpnftV2.ownerOf(1)).to.equal(addr1.address)
      expect(await cpnftV2.ownerOf(2)).to.equal(addr1.address)
      expect(await cpnftV2.getTokenLevel(1)).to.equal(0)
      expect(await cpnftV2.getTokenLevel(2)).to.equal(1)
      expect(await cpnftV2.tokenExists(1)).to.be.true
      expect(await cpnftV2.tokenExists(2)).to.be.true
    })

    it('Should allow new functionality after upgrade', async function () {
      // Upgrade to V2
      const CPNFTUpgradeableV2 = await ethers.getContractFactory('CPNFTUpgradeableV2')
      await upgrades.upgradeProxy(cpnft.address, CPNFTUpgradeableV2)
      
      cpnftV2 = await ethers.getContractAt('CPNFTUpgradeableV2', cpnft.address)

      // Test new functionality
      await cpnftV2.setTokenMetadata(1, 'Test metadata')
      expect(await cpnftV2.getTokenMetadata(1)).to.equal('Test metadata')
    })

    it('Should maintain existing functionality after upgrade', async function () {
      // Upgrade to V2
      const CPNFTUpgradeableV2 = await ethers.getContractFactory('CPNFTUpgradeableV2')
      await upgrades.upgradeProxy(cpnft.address, CPNFTUpgradeableV2)
      
      cpnftV2 = await ethers.getContractAt('CPNFTUpgradeableV2', cpnft.address)

      // Test existing functionality still works
      await cpnftV2.mint(owner.address, 2) // Level A
      expect(await cpnftV2.ownerOf(3)).to.equal(owner.address)
      expect(await cpnftV2.getTokenLevel(3)).to.equal(2)

      await cpnftV2.setTokenLevel(3, 5) // Change to SSS
      expect(await cpnftV2.getTokenLevel(3)).to.equal(5)

      await cpnftV2.burn(3)
      expect(await cpnftV2.tokenExists(3)).to.be.false
    })
  })

  describe('Upgrade Authorization', function () {
    it('Should only allow owner to upgrade', async function () {
      const CPNFTUpgradeableV2 = await ethers.getContractFactory('CPNFTUpgradeableV2')
      
      // Try to upgrade with non-owner account
      await expect(
        upgrades.upgradeProxy(cpnft.connect(addr1).address, CPNFTUpgradeableV2)
      ).to.be.reverted
    })
  })

  describe('Storage Layout Compatibility', function () {
    it('Should maintain storage layout compatibility', async function () {
      // This test ensures that the storage layout is compatible
      // by verifying that existing data can be accessed after upgrade
      
      const CPNFTUpgradeableV2 = await ethers.getContractFactory('CPNFTUpgradeableV2')
      await upgrades.upgradeProxy(cpnft.address, CPNFTUpgradeableV2)
      
      cpnftV2 = await ethers.getContractAt('CPNFTUpgradeableV2', cpnft.address)

      // Verify that all storage variables are accessible
      expect(await cpnftV2.getCurrentTokenId()).to.equal(2)
      expect(await cpnftV2.getNextTokenId()).to.equal(3)
    })
  })
})
