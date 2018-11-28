import { Store } from 'redux'
import Web3 from 'web3'
import { Blockchain } from '../../utils/config'

import { Web3Provider } from 'ethers/providers/web3-provider'
import { AccountProps, InfoProps, ObjectProps, OrgContractProps } from '../../store/blockchain/types'

import { addAccount, addInfo, addObject, addOrgContract } from '../../store/blockchain/actions'

import { getProviders } from './blockchainProvider'
import { getAccount } from './blockchainAccount'
import { getOrgContract } from './blockchainContracts'

interface SetBlockchainProps {
  store: Store
}

interface SetProps {
  store: Store
  provider: object
}

export const setAccount = (props: SetProps) => {

  const store = props.store
  const state = store.getState()
  let accountData: AccountProps = {
    data: {
      account: state.blockchain.data.account
    }
  }

  getAccount({provider: props.provider}).then((account) => {
    if ( accountData.data.account != account ) {
      //console.log('Storing Account', account)
      accountData.data.account = account
      store.dispatch(addAccount(accountData))
    }
  })
}

export const setOrgContract = (props: SetProps) => {

  //console.log('Storing OrgContract')

  const store = props.store
  const state = store.getState()
  let orgContractData: OrgContractProps = {
    data: {
      orgContract: state.blockchain.data.orgContract
    }
  }

  getOrgContract({provider: props.provider}).then((orgContract) => {
    if ( typeof orgContract != "undefined" ) {
      //console.log('Storing OrgContract', orgContract)
      orgContractData.data.orgContract = orgContract
      store.dispatch(addOrgContract(orgContractData))
    }
  })
}

export const setBlockchain = async (props: SetBlockchainProps) => {

  //console.log(props.store)
  const store = props.store
  const providers = getProviders()
  let objectData: ObjectProps = {
    data: {
      provider: providers[0] as Web3Provider
    }
  }
  const web3 = providers[1]
  const chainObj: any = await (objectData.data.provider as Web3Provider).getNetwork()
  //console.log(chainObj)
  //console.log('First call ', 'Name: ', chainObj.name, ' ChainID: ', chainObj.chainId, 'ENS Address: ', chainObj.ensAddress)
  let ENSAddress = ""
  if (typeof chainObj.ensAddress != 'undefined') {
    ENSAddress = chainObj.ensAddress
  }
  let infoData: InfoProps = {
    data: {
      API: 'web3 ' + (web3 as Web3).version,
      Name: chainObj.name,
      ChainId: chainObj.chainId,
      ENS: ENSAddress
    }
  }

  store.dispatch(addInfo(infoData))
  store.dispatch(addObject(objectData))

  setOrgContract({store: store, provider: objectData.data.provider})
  setInterval(() => {
    setAccount({store: store, provider: objectData.data.provider})
  }, Blockchain.checkAccountInterval)
}
