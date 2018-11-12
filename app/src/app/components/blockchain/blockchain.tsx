import Web3 from 'web3'
import { ethers } from 'ethers'

interface ProviderProps {
  readonly web3: any
}

interface Web3Props {
  readonly address: string
  readonly port: string
}

interface AccountProps {
  readonly web3: any
}

const GetWeb3 = async (props: Web3Props) => {

  const ethereum = (window as any).ethereum
  let web3 = (window as any).web3
  if (ethereum) {
    console.log('New Web3 MetaMask!')
    web3 = new Web3(ethereum)
  } else if (typeof web3 !== 'undefined') {
    console.log('Legacy web3 provider')
  } else {
    console.log('Running our own blockchain provider')
    const address = 'http://' + props.address + ':' + props.port
    web3 = new Web3(new Web3.providers.HttpProvider(address))
  }

  return web3
}

const GetProvider = async (props: ProviderProps) => {

  const ethereum = (window as any).ethereum
  const web3 = props.web3
  console.log(web3)
  let blockchainProvider = undefined
  if (ethereum) {
    console.log('New MetaMask!')
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
    blockchainProvider = new ethers.providers.Web3Provider((HttpProvider).web3)
  }

  blockchainProvider.getNetwork().then(function(chainObj: any) {
    console.log('Name: ', chainObj.name, ' ChainID: ', chainObj.chainId, 'ENS Address: ', chainObj.ensAddress)
  })

  return blockchainProvider
}

// metamask sets its account to web3.eth.accounts[0]
const SetAccount = async (props: AccountProps) => {
  let account = ''
  await props.web3.eth.getAccounts((error: any, accounts: any) => {
      account = accounts[0]
  })
  props.web3.eth.defaultAccount = account
  console.log('Account: ', account)
  return account
}

export { GetWeb3, GetProvider, SetAccount }
