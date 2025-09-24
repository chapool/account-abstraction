import { HardhatRuntimeEnvironment } from 'hardhat/types'
import { DeployFunction } from 'hardhat-deploy/types'
import { ethers, upgrades } from 'hardhat'

const deployCPNFT: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  // Use Alchemy RPC URL for Sepolia
  if (hre.network.name === 'sepoliaCustom') {
    hre.network.config.url = 'https://eth-sepolia.g.alchemy.com/v2/_x4NAgu50ejHAhTH2-gJoRNS4PQv7Tjp'
  }
  
  const provider = ethers.provider
  const from = await provider.getSigner().getAddress()
  
  console.log('Deploying CPNFT contract...')
  console.log('Network:', hre.network.name)
  console.log('Deployer address:', from)
  
  // Get the CPNFT contract factory
  const CPNFT = await ethers.getContractFactory('CPNFT')
  
  // Deploy as upgradeable proxy
  const cpnft = await upgrades.deployProxy(
    CPNFT,
    [
      'Chapool NFT', // name
      'CPNFT',       // symbol
      'http://chapool.net:8080/api/v1/meta/' // baseURI
    ],
    {
      initializer: 'initialize',
      kind: 'uups'
    }
  )
  
  await cpnft.deployed()
  
  console.log('CPNFT deployed to:', cpnft.address)
  console.log('Implementation address:', await upgrades.erc1967.getImplementationAddress(cpnft.address))
  console.log('Admin address:', await upgrades.erc1967.getAdminAddress(cpnft.address))
  
  // Save deployment info
  const implementationAddress = await upgrades.erc1967.getImplementationAddress(cpnft.address)
  const adminAddress = await upgrades.erc1967.getAdminAddress(cpnft.address)
  
  const deployment = await hre.deployments.save('CPNFT', {
    address: cpnft.address,
    abi: JSON.parse(CPNFT.interface.format('json')),
    transactionHash: cpnft.deployTransaction.hash,
    receipt: await cpnft.deployTransaction.wait(),
    args: [
      'Chapool NFT',
      'CPNFT',
      'http://chapool.net:8080/api/v1/meta/'
    ],
    libraries: {},
    metadata: JSON.stringify({
      implementation: implementationAddress,
      proxyAdmin: adminAddress
    })
  })
  
  console.log('Deployment saved:', deployment?.address || 'Failed to save deployment info')
  
  // Verify contract on etherscan if not on local network
  if (hre.network.name !== 'dev' && hre.network.name !== 'localhost') {
    try {
      console.log('Waiting for block confirmations...')
      await cpnft.deployTransaction.wait(5)
      
      console.log('Verifying contract on Etherscan...')
      await hre.run('verify:verify', {
        address: cpnft.address,
        constructorArguments: [],
        contract: 'contracts/cpop/CPNFT.sol:CPNFT'
      })
      console.log('Contract verified successfully!')
    } catch (error) {
      console.log('Verification failed:', error instanceof Error ? error.message : String(error))
    }
  }
  
  return true
}

deployCPNFT.tags = ['CPNFT', 'upgradeable']
deployCPNFT.dependencies = []

export default deployCPNFT
