import { Store } from 'redux'
import Web3 from 'web3'
import { Blockchain } from '../../utils/config'

import { Web3Provider } from 'ethers/providers/web3-provider'
import { AccountProps } from '../../store/blockchain/account/types'
import { InfoProps } from '../../store/blockchain/info/types'
import { ContractProps } from '../../store/blockchain/contracts/types'

import { addAccount } from '../../store/blockchain/account/actions'
import { addInfo } from '../../store/blockchain/info/actions'
import { addOrgContract, addOrgReportsContract } from '../../store/blockchain/contracts/actions'

import { getProvider } from './blockchainProvider'
import { getAccount } from './blockchainAccount'
import { getOrgContract, getOrgReportsContract } from './blockchainContracts'

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
      account: state.chainAccount.data.account
    }
  }

  getAccount({provider: props.provider}).then((account) => {
    if ( accountData.data.account != account ) {
      //console.log('Storing Account', account)
      accountData.data.account = account
      const add = addAccount as Function
      store.dispatch(add(accountData))
    }
  })
}

export const setOrgContract = (props: SetProps) => {

  const store = props.store
  const state = store.getState()
  let orgContractData: ContractProps = {
    data: {
      contracts: {
        orgContract: state.chainContracts.data.contracts.orgContract
      }
    }
  }

  if ( !(orgContractData.data.contracts.orgContract.hasOwnProperty('getOrganisationExists')) ) {
    getOrgContract({provider: props.provider}).then((orgContract) => {
      if ( typeof orgContract != "undefined" ) {
        //console.log, web3]('Storing OrgContract', orgContract)
        orgContractData.data.contracts.orgContract = orgContract
        const add = addOrgContract as Function
        store.dispatch(add(orgContractData))
      }
    })
  }
}

export const setOrgReportsContract = (props: SetProps) => {

  const store = props.store
  const state = store.getState()
  let orgReportsContractData: ContractProps = {
    data: {
      contracts: {
        orgContract: {},
        orgReportsContract:  state.chainContracts.data.contracts.orgReportsContract,
      }
    }
  }

  if ( !(orgReportsContractData.data.contracts.orgReportsContract.hasOwnProperty('getReportExists')) ) {
    getOrgReportsContract({provider: props.provider}).then((orgReportsContract) => {
      if ( typeof orgReportsContract != "undefined" ) {
        //console.log('Storing OrgContract', orgReportsContract)
        orgReportsContractData.data.contracts.orgReportsContract = orgReportsContract
        const add = addOrgReportsContract as Function
        store.dispatch(add(orgReportsContractData))
      }
    })
  }
}

export const setBlockchain = async (props: SetBlockchainProps) => {

  //console.log(props.store)
  const store = props.store
  const provider = getProvider()
  const chainObj: any = await provider.getNetwork()
  //console.log(chainObj)
  //console.log('First call ', 'Name: ', chainObj.name, ' ChainID: ', chainObj.chainId, 'ENS Address: ', chainObj.ensAddress)

  let ENSAddress = ""
  if (typeof chainObj.ensAddress != 'undefined') {
    ENSAddress = chainObj.ensAddress
  }
  let infoData: InfoProps = {
    data: {
      Name: chainObj.name,
      ChainId: chainObj.chainId,
      ENS: ENSAddress,
      provider: provider
    }
  }

  setInterval(() => {
    setAccount({store: store, provider: provider})
  }, Blockchain.checkAccountInterval)

  setInterval(() => {
    setOrgContract({store: store, provider: provider})
  }, Blockchain.checkAccountInterval)

  setInterval(() => {
    setOrgReportsContract({store: store, provider: provider})
  }, Blockchain.checkAccountInterval)

  const add = addInfo as Function
  store.dispatch(add(infoData))
}
