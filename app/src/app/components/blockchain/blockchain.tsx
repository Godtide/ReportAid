import * as React from 'react'
import { connect } from 'react-redux'
import Web3 from 'web3'

import { BlockchainConfig } from '../../utils/config'

import { Provider } from 'ethers/providers/abstract-provider'

import { BlockchainAccountProps, BlockchainInfoProps, BlockchainObjectProps } from '../../store/blockchain/types'
import { addAccount, addInfo, addObjects } from '../../store/blockchain/actions'

import { getProviders } from './blockchainProvider'
import { getAccount } from './blockchainAccount'

const createAddAccount = (dispatch: any) => {
  let thisAddAccount = (ownProps: any) => {
    dispatch(ownProps)
  }
  return thisAddAccount
}

const createAddInfo = (dispatch: any) => {
  let thisAddInfo = (ownProps: any) => {
    dispatch(ownProps)
  }
  return thisAddInfo
}

const createAddObject = (dispatch: any) => {
  let thisAddObject = (ownProps: any) => {
    dispatch(ownProps)
  }
  return thisAddObject
}

class Blockchain extends React.Component {

  addAccount: any
  addInfo: any
  addObject: any

  constructor(props: any) {
    super(props)
    this.setBlockchain()
    this.addAccount = createAddAccount(props.addAccount)
    this.addInfo = createAddInfo(props.addInfo)
    this.addObject = createAddObject(props.addObject)
  }

  async setBlockchain() {

    const providers = getProviders()
    let objectData: BlockchainObjectProps = {
      provider: providers[0]
    }

    const web3 = providers[1]
    let accountData = {
      account: ''
    }

    let infoData: BlockchainInfoProps = {
      APIName: '',
      networkName: '',
      networkChainId: '',
      networkENSAddress: ''
    }

    const chainObj = await (objectData.provider as Provider).getNetwork()
    console.log(chainObj)
    console.log('First call ', 'Name: ', chainObj.name, ' ChainID: ', chainObj.chainId, 'ENS Address: ', chainObj.ensAddress)
    infoData.networkName = chainObj.name
    infoData.networkChainId = chainObj.chainId as any
    infoData.networkENSAddress = chainObj.ensAddress as string
    infoData.APIName = 'web3 ' + (web3 as Web3).version as string

      //props.store.dispatch(addInfo(infoData))
      //props.store.dispatch(addObjects(objectData))

    (window as Window).setInterval(() => {
      const account = getAccount({provider: objectData.provider})
      console.log('Stored account', accountData.account)
      if ( accountData.account != account ) {
        accountData.account = account
        console.log ('Setting account', accountData.account)
        this.addAccount(accountData.account)
      }
    }, BlockchainConfig.interval)
  }
}

export const SetBlockchain = connect()(Blockchain)
