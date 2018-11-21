import { Store } from 'redux'
import Web3 from 'web3'
import { Blockchain } from '../../utils/config'

import { Web3Provider } from 'ethers/providers/web3-provider'
import { AccountProps, InfoProps, ObjectProps } from '../../store/blockchain/types'

import { addAccount, addInfo, addObject } from '../../store/blockchain/actions'

import { getProviders } from './blockchainProvider'
import { getAccount } from './blockchainAccount'

interface SetBlockchainProps {
  store: Store
}

interface SetAccountProps {
  store: Store
  provider: object
}

export const setAccount = (props: SetAccountProps) => {

  const store = props.store
  const state = store.getState()
  let accountData: AccountProps = {
    account: state.blockchain.account
  }

  getAccount({provider: props.provider}).then((account) => {
    if ( accountData.account != account ) {
      //console.log('Storing Account', account)
      accountData.account = account
      store.dispatch(addAccount(accountData))
    }
  })
}

export const setBlockchain = async (props: SetBlockchainProps) => {

  //console.log(props.store)
  const store = props.store
  const providers = getProviders()
  let objectData: ObjectProps = {
    provider: providers[0] as Web3Provider
  }
  const web3 = providers[1]
  const chainObj: any = await (objectData.provider as Web3Provider).getNetwork()
  //console.log(chainObj)
  //console.log('First call ', 'Name: ', chainObj.name, ' ChainID: ', chainObj.chainId, 'ENS Address: ', chainObj.ensAddress)
  let infoData: InfoProps = {
    API: 'web3 ' + (web3 as Web3).version,
    Name: chainObj.name,
    ChainId: chainObj.chainId,
    ENS: chainObj.ensAddress
  }

  store.dispatch(addInfo(infoData))
  store.dispatch(addObject(objectData))

  setInterval(() => {
    setAccount({store: store, provider: objectData.provider})
  }, Blockchain.checkAccountInterval)
}
