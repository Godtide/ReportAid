import { Store } from 'redux'

import Web3 from 'web3'
import { ethers } from 'ethers'

import { BlockchainInfoProps, BlockchainObjectProps } from '../../store/blockchain/types'
import { addInfo, addObjects } from '../../store/blockchain/actions'
import { BlockchainStrings } from '../../utils/strings'

interface BlockchainProviderProps {
  store: Store
}

export const setProvider = async (props: BlockchainProviderProps) => {

  const ethereum = (window as any).ethereum
  let web3: any = (window as any).web3
  let blockchainProvider = undefined

  let infoData: BlockchainInfoProps = {
    APIName: '',
    networkName: '',
    networkChainId: '',
    networkENSAddress: ''
  }

  let objectData: BlockchainObjectProps = {
    web3: {},
    ethers: {}
  }

  if (ethereum) {
    //console.log('New MetaMask!')
    web3 = new Web3(ethereum)
    blockchainProvider = new ethers.providers.Web3Provider(web3.currentProvider)
    try {
        // Request account access if needed
        await ethereum.enable()
    } catch (error) {
      console.log(error)
    }

  } else if (typeof web3 !== 'undefined') {
    //console.log('In legacy web3 provider')
    blockchainProvider = new ethers.providers.Web3Provider(web3.currentProvider)
  } else {
    //console.log('Running our own blockchain provider')
    const address = 'http://' + BlockchainStrings.host + ':' + BlockchainStrings.port
    web3 = new Web3(new Web3.providers.HttpProvider(address))
    blockchainProvider = new ethers.providers.Web3Provider(web3)
  }

  //console.log('Account: ', account)

  await blockchainProvider.getNetwork().then(function(chainObj: any) {
    //console.log('Name: ', chainObj.name, ' ChainID: ', chainObj.chainId, 'ENS Address: ', chainObj.ensAddress)
    infoData.networkName = chainObj.name
    infoData.networkChainId = chainObj.chainId
    infoData.networkENSAddress = chainObj.ensAddress
  })

  infoData.APIName = 'web3 ' + web3.version
  objectData.web3 = web3
  objectData.ethers = blockchainProvider

  props.store.dispatch(addInfo(infoData))
  props.store.dispatch(addObjects(objectData))
}
