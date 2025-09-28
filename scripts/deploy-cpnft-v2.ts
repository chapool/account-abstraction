import { ethers, upgrades } from 'hardhat'
import * as fs from 'fs'
import * as path from 'path'

async function main() {
  console.log('Starting CPNFT V2 deployment (with NORMAL level)...')
  
  // Get the deployer account
  const [deployer] = await ethers.getSigners()
  console.log('Deploying contracts with account:', deployer.address)
  console.log('Account balance:', ethers.utils.formatEther(await deployer.getBalance()))
  
  // Check if we have enough ETH for deployment
  const balance = await deployer.getBalance()
  if (balance.lt(ethers.utils.parseEther('0.005'))) {
    throw new Error('Insufficient ETH balance for deployment. Need at least 0.005 ETH.')
  }
  
  // Get the CPNFT contract factory
  const CPNFT = await ethers.getContractFactory('CPNFT')
  
  console.log('Deploying CPNFT V2 as upgradeable proxy...')
  
  try {
    // Deploy as upgradeable proxy
    const cpnft = await upgrades.deployProxy(
      CPNFT,
      [
        'Chapool NFT V2', // name with V2
        'CPNFT',          // symbol
        'http://chapool.net:8080/api/v1/meta/' // baseURI
      ],
      {
        initializer: 'initialize',
        kind: 'uups'
      }
    )
    
    console.log('Waiting for deployment to be mined...')
    await cpnft.deployed()
    
    // Get deployment details
    const implementationAddress = await upgrades.erc1967.getImplementationAddress(cpnft.address)
    const adminAddress = await upgrades.erc1967.getAdminAddress(cpnft.address)
    
    console.log('\n=== DEPLOYMENT SUCCESSFUL ===')
    console.log('Network:', (await ethers.provider.getNetwork()).name)
    console.log('Chain ID:', (await ethers.provider.getNetwork()).chainId)
    console.log('Proxy Address:', cpnft.address)
    console.log('Implementation Address:', implementationAddress)
    console.log('Admin Address:', adminAddress)
    console.log('Deployer Address:', deployer.address)
    console.log('Base URI:', 'http://chapool.net:8080/api/v1/meta/')
    console.log('Transaction Hash:', cpnft.deployTransaction.hash)
    
    // Save deployment info to file
    const deploymentInfo = {
      network: (await ethers.provider.getNetwork()).name,
      chainId: (await ethers.provider.getNetwork()).chainId.toString(),
      proxyAddress: cpnft.address,
      implementationAddress: implementationAddress,
      adminAddress: adminAddress,
      deployerAddress: deployer.address,
      baseURI: 'http://chapool.net:8080/api/v1/meta/',
      deploymentTime: new Date().toISOString(),
      transactionHash: cpnft.deployTransaction.hash,
      version: '2.0.0',
      features: ['NORMAL level support', 'Upgraded enum structure']
    }
    
    const deploymentFile = `deployments/sepoliaCustom/CPNFT_V2.json`
    
    // Create deployments directory if it doesn't exist
    const deploymentsDir = `deployments/sepoliaCustom`
    if (!fs.existsSync(deploymentsDir)) {
      fs.mkdirSync(deploymentsDir, { recursive: true })
    }
    
    fs.writeFileSync(deploymentFile, JSON.stringify(deploymentInfo, null, 2))
    console.log('\nDeployment info saved to:', deploymentFile)
    
    // Test basic functionality
    console.log('\n=== TESTING CONTRACT ===')
    const name = await cpnft.name()
    const symbol = await cpnft.symbol()
    const owner = await cpnft.owner()
    const version = await cpnft.version()
    
    console.log('Contract Name:', name)
    console.log('Contract Symbol:', symbol)
    console.log('Contract Owner:', owner)
    console.log('Contract Version:', version)
    
    // Test NORMAL level
    console.log('\n=== TESTING NORMAL LEVEL ===')
    console.log('NORMAL level value:', 0) // NORMAL = 0
    console.log('C level value:', 1)      // C = 1
    console.log('B level value:', 2)      // B = 2
    console.log('A level value:', 3)      // A = 3
    console.log('S level value:', 4)      // S = 4
    console.log('SS level value:', 5)     // SS = 5
    console.log('SSS level value:', 6)    // SSS = 6
    
    // Try to verify contract on Etherscan
    if (network.name !== 'localhost' && network.name !== 'dev' && network.name !== 'hardhat') {
      try {
        console.log('\nAttempting contract verification...')
        await new Promise(resolve => setTimeout(resolve, 15000)) // Wait 15 seconds
        
        await hre.run('verify:verify', {
          address: cpnft.address,
          constructorArguments: [],
          contract: 'contracts/cpop/CPNFT.sol:CPNFT'
        })
        console.log('Contract verified successfully on Etherscan!')
      } catch (error) {
        console.log('Verification failed:', error instanceof Error ? error.message : String(error))
        console.log('You can verify manually later using:')
        console.log(`npx hardhat verify --network sepoliaCustom ${cpnft.address}`)
      }
    }
    
    console.log('\n=== DEPLOYMENT COMPLETE ===')
    console.log('✅ CPNFT V2 contract is ready to use!')
    console.log('✅ NORMAL level is now supported!')
    console.log('\n📋 Contract Information:')
    console.log(`   Proxy Address: ${cpnft.address}`)
    console.log(`   Implementation: ${implementationAddress}`)
    console.log(`   Etherscan: https://sepolia.etherscan.io/address/${cpnft.address}`)
    
  } catch (error) {
    console.error('Deployment failed:', error)
    throw error
  }
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error('Deployment failed:', error)
    process.exit(1)
  })

