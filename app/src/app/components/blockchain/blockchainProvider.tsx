import { Store } from 'redux'

import Web3 from 'web3'
import { ethers } from 'ethers'

import { Provider } from 'ethers/providers/abstract-provider'

import { BlockchainInfoProps, BlockchainObjectProps } from '../../store/blockchain/types'
import { addInfo, addObjects } from '../../store/blockchain/actions'
import { BlockchainStrings } from '../../utils/strings'

interface BlockchainProviderProps {
  store: Store
}

export const setProvider = async (props: BlockchainProviderProps) => {

  const store = props.store
  const state = store.getState()

  let infoData: BlockchainInfoProps = {
    APIName: state.blockchain.APIName,
    networkName: state.blockchain.networkName,
    networkChainId: state.blockchain.networkChainId,
    networkENSAddress: state.blockchain.networkENSAddress
  }

  let objectData: BlockchainObjectProps = {
    web3: state.blockchain.web3 as Web3,
    ethers: state.blockchain.ethers
  }

  if ( objectData.web3.hasOwnProperty('version') ) {

    let thisInfoData: BlockchainInfoProps = {
      APIName: infoData.APIName,
      networkName: '',
      networkChainId: '',
      networkENSAddress: ''
    }

    await (objectData.ethers as Provider).getNetwork().then(function(chainObj: any) {
      thisInfoData.networkName = chainObj.name
      thisInfoData.networkChainId = chainObj.chainId
      thisInfoData.networkENSAddress = chainObj.ensAddress
    })

    if ( thisInfoData.networkName != infoData.networkName ) {
      console.log('New call ',
                  'Name: ', thisInfoData.networkName,
                  ' ChainID: ', thisInfoData.networkChainId,
                  ' ENS Address: ', thisInfoData.networkENSAddress)
      infoData.networkName = thisInfoData.networkName
      infoData.networkChainId = thisInfoData.networkChainId
      infoData.networkENSAddress = thisInfoData.networkENSAddress

      let providers = _getProvider()
      let web3 = providers[0]
      let blockchainProvider = providers[1]

      objectData.web3 = web3
      objectData.ethers = blockchainProvider

      props.store.dispatch(addInfo(infoData))
      props.store.dispatch(addObjects(objectData))
    }

  } else {

    let providers = _getProvider()
    let web3 = providers[0]
    let blockchainProvider = providers[1]

    await blockchainProvider.getNetwork().then(function(chainObj: any) {
      console.log('First call ', 'Name: ', chainObj.name, ' ChainID: ', chainObj.chainId, 'ENS Address: ', chainObj.ensAddress)
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

}

const _getProvider = () => {

  let ethereum = (window as any).ethereum
  let web3 = (window as any).web3
  let blockchainProvider = undefined

  if (ethereum) {
    //console.log('New MetaMask!')
    web3 = new Web3(ethereum)
    blockchainProvider = new ethers.providers.Web3Provider(web3.currentProvider)
    ethereum.enable().then(function(result: any) {
      console.log('Ethereum enabled', result)
    }).catch(function (error: any) {
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

  return [web3, blockchainProvider]
}
