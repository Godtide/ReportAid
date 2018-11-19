import { Store } from 'redux'

import Web3 from 'web3'
import { ethers } from 'ethers'

// import { Provider } from 'ethers/providers/abstract-provider'

import { BlockchainInfoProps, BlockchainObjectProps } from '../../store/blockchain/types'
import { addInfo, addObjects } from '../../store/blockchain/actions'
import { BlockchainStrings } from '../../utils/strings'

interface BlockchainProviderProps {
  store: Store
}

export const setProvider = (props: BlockchainProviderProps): boolean => {

  let result = true
  const store = props.store
  const state = store.getState()
  let objectData: BlockchainObjectProps = {
    provider: state.blockchain.provider
  }

  if ( !(objectData.provider.hasOwnProperty('getDefaultProvider')) ) {

    let ethereum = (window as any).ethereum
    let web3 = (window as any).web3
    let blockchainProvider: any = undefined

    let infoData: BlockchainInfoProps = {
      APIName: state.blockchain.APIName,
      networkName: state.blockchain.networkName,
      networkChainId: state.blockchain.networkChainId,
      networkENSAddress: state.blockchain.networkENSAddress
    }

    if (ethereum) {
      //console.log('New MetaMask!')
      web3 = new Web3(ethereum)
      blockchainProvider = new ethers.providers.Web3Provider(web3.currentProvider)
      ethereum.enable().then(function(result: any) {
        console.log('Ethereum enabled', result)
      }).catch(function (error: any) {
        console.log(error)
        result = false
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

    blockchainProvider.getNetwork().then(function(chainObj: any) {
      console.log('First call ', 'Name: ', chainObj.name, ' ChainID: ', chainObj.chainId, 'ENS Address: ', chainObj.ensAddress)
      infoData.networkName = chainObj.name
      infoData.networkChainId = chainObj.chainId
      infoData.networkENSAddress = chainObj.ensAddress
      infoData.APIName = 'web3 ' + web3.version

      objectData.provider = blockchainProvider

      props.store.dispatch(addInfo(infoData))
      props.store.dispatch(addObjects(objectData))
    }).catch(function (error: any) {
      console.log(error)
      result = false
    })
  }
  return result
}
