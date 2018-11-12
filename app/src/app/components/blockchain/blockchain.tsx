import Web3 from 'web3'
import { ethers } from 'ethers'

interface ProviderProps {
  readonly address: string
  readonly port: string
}

const GetProvider = async (props: ProviderProps) => {

  const ethereum = (window as any).ethereum
  let web3 = (window as any).web3
  let blockchainProvider = undefined
  if (ethereum) {
    console.log('New MetaMask!')
    web3 = new Web3(ethereum)
    blockchainProvider = new ethers.providers.Web3Provider(web3.currentProvider)
    try {
        // Request account access if needed
        await ethereum.enable()
    } catch (error) {
      console.log(error)
    }
  } else if (typeof web3 !== 'undefined') {
    console.log('In legacy web3 provider')
    blockchainProvider = new ethers.providers.Web3Provider(web3.currentProvider)
  } else {
    console.log('Running our own blockchain provider')
    const address = 'http://' + props.address + ':' + props.port
    web3 = new Web3(new Web3.providers.HttpProvider(address))
    blockchainProvider = new ethers.providers.Web3Provider(web3)
  }

  blockchainProvider.getNetwork().then(function(chainObj: any) {
    console.log('Name: ', chainObj.name, ' ChainID: ', chainObj.chainId, 'ENS Address: ', chainObj.ensAddress)
  })

  return blockchainProvider
}

// metamask sets its account to web3.eth.accounts[0]
const SetAccount = async () => {
  let account = ''
  const web3 = (window as any).web3
  await web3.eth.getAccounts((error: any, accounts: any) => {
      account = accounts[0]
  })
  web3.defaultAccount = account
  console.log('Account: ', account)
  return account
}

export { GetProvider, SetAccount }
