//import * as React from 'react'
import Web3 from 'web3'
import { connect } from 'react-redux'
import { Dispatch } from 'redux'
import { ethers } from 'ethers'
import { BlockchainProps, BlockchainInfoProps, BlockchainObjectProps } from '../../store/blockchain/types'
import { BlockchainStrings } from '../../utils/strings'
import { ApplicationState } from '../../store'
import { addInfo, addObjects } from '../../store/blockchain/actions'

interface ProviderDispatchProps {
    addInfo(props: BlockchainInfoProps): void,
    addObjects(props: BlockchainObjectProps): void
}

type ProviderProps = BlockchainProps & ProviderDispatchProps

const Provider = (props: ProviderProps) => {

  console.log('blah')

  let objectData: BlockchainObjectProps = {
    web3: props.web3,
    ethers: props.ethers
  }

  if ( !(objectData.web3.hasOwnProperty('version')) ) {

    let ethereum = (window as any).ethereum
    let web3 = (window as any).web3
    let blockchainProvider: any = undefined

    let infoData: BlockchainInfoProps = {
      APIName: props.APIName,
      networkName: props.networkName,
      networkChainId: props.networkChainId,
      networkENSAddress: props.networkENSAddress
    }

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

    blockchainProvider.getNetwork().then(function(chainObj: any) {
      console.log('First call ', 'Name: ', chainObj.name, ' ChainID: ', chainObj.chainId, 'ENS Address: ', chainObj.ensAddress)
      infoData.networkName = chainObj.name
      infoData.networkChainId = chainObj.chainId
      infoData.networkENSAddress = chainObj.ensAddress
      infoData.APIName = 'web3 ' + web3.version

      objectData.web3 = web3
      objectData.ethers = blockchainProvider

      this.props.addInfo(infoData)
      this.props.addObjects(objectData)

    }).catch(function (error: any) {
      console.log(error)
    })
  }
}

const mapStateToProps = (state: ApplicationState): BlockchainProps => {
  return {
    APIName: state.blockchain.APIName,
    networkName: state.blockchain.networkName,
    networkChainId: state.blockchain.networkChainId,
    networkENSAddress: state.blockchain.networkENSAddress,
    account: state.blockchain.account,
    web3: state.blockchain.web3,
    ethers: state.blockchain.ethers
  }
}

const mapDispatchToProps = (dispatch: Dispatch, ownProps: any): ProviderDispatchProps => {
  return({
    addInfo: () => {dispatch(addInfo(ownProps))},
    addObjects: () => {dispatch(addObjects(ownProps))}
  })
}

export const SetProvider = connect<BlockchainProps,ProviderDispatchProps, any,ApplicationState>(
    mapStateToProps,
    mapDispatchToProps)
(Provider)
