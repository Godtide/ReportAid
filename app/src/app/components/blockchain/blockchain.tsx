import { Store } from 'redux'
import Web3 from 'web3'

import { BlockchainConfig } from '../../utils/config'

import { Web3Provider } from 'ethers/providers/web3-provider'

import { BlockchainAccountProps, BlockchainInfoProps, BlockchainObjectProps } from '../../store/blockchain/types'
import { addAccount, addInfo, addObjects } from '../../store/blockchain/actions'

import { getProviders } from './blockchainProvider'
import { getAccount } from './blockchainAccount'

interface BlockchainProviderProps {
  store: Store
}

export const setBlockchain = async () => {

  const providers = getProviders()
  let objectData = {
    provider: providers[0] as Web3Provider
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

  const chainObj = await objectData.provider.getNetwork()
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
      //props.addAccount(accountData)
    }
  }, BlockchainConfig.interval)
  return true
}

/* const mapDispatchToProps = (dispatch: any): DispatchBlockchainProps => {
  return {
    addAccount: (props: BlockchainAccountProps) => dispatch(addAccount(props))
  }
}

export default connect(
  mapDispatchToProps
)(setBlockchain)*/
