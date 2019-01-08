import Web3 from 'web3'
import { ethers } from 'ethers'

import { Provider } from 'ethers/providers/abstract-provider'
import { Blockchain } from '../../utils/config'

export const getProvider = () => {

  let blockchainProvider: Provider
  let ethereum = (window as any).ethereum
  let web3 = (window as any).web3

  if (ethereum) {
    //console.log('New MetaMask!')
    web3 = new Web3(ethereum)
    blockchainProvider = new ethers.providers.Web3Provider(web3.currentProvider)
    ethereum.enable().then((result: any) => {
      console.log('Ethereum Enable: ', result)
    }).catch((error: any) => {
      console.log(error)
    })
  } else if (typeof web3 !== 'undefined') {
    //console.log('In legacy web3 provider')
    blockchainProvider = new ethers.providers.Web3Provider(web3.currentProvider)
  } else {
    const host = Blockchain.host
    const port = Blockchain.port
    const provider = host + ":" + port
    const network = Blockchain.network
    blockchainProvider = new ethers.providers.JsonRpcProvider(provider, network)
  }

  return blockchainProvider
}
