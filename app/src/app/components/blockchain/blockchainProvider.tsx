import Web3 from 'web3'
import { ethers } from 'ethers'

import { Provider } from 'ethers/providers/abstract-provider'
import { BlockchainStrings } from '../../utils/strings'

export const getProviders = (): any[] => {

  let blockchainProvider: Provider
  let ethereum = (window as any).ethereum
  let web3 = (window as any).web3

  if (ethereum) {
    //console.log('New MetaMask!')
    web3 = new Web3(ethereum)
    blockchainProvider = new ethers.providers.Web3Provider(web3.currentProvider)
    ethereum.enable().then((result: any) => {
      console.log('Ethereum enabled', result)
    }).catch((error: any) => {
      console.log(error)
    })
  } else if (typeof web3 !== 'undefined') {
    //console.log('In legacy web3 provider')
    blockchainProvider = new ethers.providers.Web3Provider(web3.currentProvider)
  } else {
    //console.log('Running our own blockchain provider')
    const address = 'http://' + BlockchainStrings.host + ':' + BlockchainStrings.port
    web3 = new Web3(new Web3.providers.HttpProvider(address))
    blockchainProvider = new ethers.providers.Web3Provider(web3)
  }
  return [blockchainProvider, web3]
}
