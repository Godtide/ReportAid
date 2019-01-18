import { Store } from 'redux'
import { Blockchain } from '../../utils/config'

import { AccountProps } from '../../store/blockchain/account/types'
import { InfoProps } from '../../store/blockchain/info/types'
import { ContractProps } from '../../store/blockchain/contracts/types'

import { addAccount } from '../../store/blockchain/account/actions'
import { addInfo } from '../../store/blockchain/info/actions'
import { addOrgContract,
         addOrgReportsContract,
         addOrgReportDocsContract,
         addOrgReportBudgetsContract,
         addOrgReportExpenditureContract,
         addOrgReportRecipientBudgetsContract,
         addOrgReportRegionBudgetsContract,
         addOrgReportCountryBudgetsContract } from '../../store/blockchain/contracts/actions'

import { getProvider } from './blockchainProvider'
import { getAccount } from './blockchainAccount'
import { getOrgContract,
         getOrgReportsContract,
         getOrgReportDocsContract,
         getOrgReportBudgetsContract,
         getOrgReportExpenditureContract,
         getOrgReportRecipientBudgetsContract,
         getOrgReportRegionBudgetsContract,
         getOrgReportCountryBudgetsContract } from './blockchainContracts'

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
        orgContract: state.chainContracts.data.contracts.orgContract,
        orgReportsContract: {},
        orgReportDocsContract: {},
        orgReportBudgetsContract: {},
        orgReportExpenditureContract: {},
        orgReportRecipientBudgetsContract: {},
        orgReportRegionBudgetsContract: {},
        orgReportCountryBudgetsContract: {}
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
        orgReportDocsContract: {},
        orgReportBudgetsContract: {},
        orgReportExpenditureContract: {},
        orgReportRecipientBudgetsContract: {},
        orgReportRegionBudgetsContract: {},
        orgReportCountryBudgetsContract: {}
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

export const setOrgReportDocsContract = (props: SetProps) => {

  const store = props.store
  const state = store.getState()
  let orgReportDocsContractData: ContractProps = {
    data: {
      contracts: {
        orgContract: {},
        orgReportsContract: {},
        orgReportDocsContract: state.chainContracts.data.contracts.orgReportDocsContract,
        orgReportBudgetsContract: {},
        orgReportExpenditureContract: {},
        orgReportRecipientBudgetsContract: {},
        orgReportRegionBudgetsContract: {},
        orgReportCountryBudgetsContract: {}
      }
    }
  }

  if ( !(orgReportDocsContractData.data.contracts.orgReportDocsContract.hasOwnProperty('getReportExists')) ) {
    getOrgReportDocsContract({provider: props.provider}).then((orgReportDocsContract) => {
      if ( typeof orgReportDocsContract != "undefined" ) {
        //console.log('Storing OrgContract', orgReportsContract)
        orgReportDocsContractData.data.contracts.orgReportDocsContract = orgReportDocsContract
        const add = addOrgReportDocsContract as Function
        store.dispatch(add(orgReportDocsContractData))
      }
    })
  }
}

export const setOrgReportBudgetsContract = (props: SetProps) => {

  const store = props.store
  const state = store.getState()
  let orgReportBudgetsContractData: ContractProps = {
    data: {
      contracts: {
        orgContract: {},
        orgReportsContract: {},
        orgReportDocsContract: {},
        orgReportBudgetsContract: state.chainContracts.data.contracts.orgReportBudgetsContract,
        orgReportExpenditureContract: {},
        orgReportRecipientBudgetsContract: {},
        orgReportRegionBudgetsContract: {},
        orgReportCountryBudgetsContract: {}
      }
    }
  }

  if ( !(orgReportBudgetsContractData.data.contracts.orgReportBudgetsContract.hasOwnProperty('getReportExists')) ) {
    getOrgReportBudgetsContract({provider: props.provider}).then((orgReportBudgetsContract) => {
      if ( typeof orgReportBudgetsContract != "undefined" ) {
        //console.log('Storing OrgContract', orgReportsContract)
        orgReportBudgetsContractData.data.contracts.orgReportBudgetsContract = orgReportBudgetsContract
        const add = addOrgReportBudgetsContract as Function
        store.dispatch(add(orgReportBudgetsContractData))
      }
    })
  }
}

export const setOrgReportExpenditureContract = (props: SetProps) => {

  const store = props.store
  const state = store.getState()
  let orgReportExpenditureContractData: ContractProps = {
    data: {
      contracts: {
        orgContract: {},
        orgReportsContract: {},
        orgReportDocsContract: {},
        orgReportBudgetsContract: {},
        orgReportExpenditureContract: state.chainContracts.data.contracts.orgReportExpenditureContract,
        orgReportRecipientBudgetsContract: {},
        orgReportRegionBudgetsContract: {},
        orgReportCountryBudgetsContract: {}
      }
    }
  }

  if ( !(orgReportExpenditureContractData.data.contracts.orgReportExpenditureContract.hasOwnProperty('getReportExists')) ) {
    getOrgReportExpenditureContract({provider: props.provider}).then((orgReportExpenditureContract) => {
      if ( typeof orgReportExpenditureContract != "undefined" ) {
        //console.log('Storing OrgContract', orgReportsContract)
        orgReportExpenditureContractData.data.contracts.orgReportExpenditureContract = orgReportExpenditureContract
        const add = addOrgReportExpenditureContract as Function
        store.dispatch(add(orgReportExpenditureContractData))
      }
    })
  }
}

export const setOrgReportRecipientBudgetsContract = (props: SetProps) => {

  const store = props.store
  const state = store.getState()
  let orgReportRecipientBudgetsContractData: ContractProps = {
    data: {
      contracts: {
        orgContract: {},
        orgReportsContract: {},
        orgReportDocsContract: {},
        orgReportBudgetsContract: {},
        orgReportExpenditureContract: {},
        orgReportRecipientBudgetsContract: state.chainContracts.data.contracts.orgReportRecipientBudgetsContract,
        orgReportRegionBudgetsContract: {},
        orgReportCountryBudgetsContract: {}
      }
    }
  }

  if ( !(orgReportRecipientBudgetsContractData.data.contracts.orgReportRecipientBudgetsContract.hasOwnProperty('getReportExists')) ) {
    getOrgReportRecipientBudgetsContract({provider: props.provider}).then((orgReportRecipientBudgetsContract) => {
      if ( typeof orgReportRecipientBudgetsContract != "undefined" ) {
        //console.log('Storing OrgContract', orgReportsContract)
        orgReportRecipientBudgetsContractData.data.contracts.orgReportExpenditureContract = orgReportRecipientBudgetsContract
        const add = addOrgReportRecipientBudgetsContract as Function
        store.dispatch(add(orgReportRecipientBudgetsContractData))
      }
    })
  }
}

export const setOrgReportRegionBudgetsContract = (props: SetProps) => {

  const store = props.store
  const state = store.getState()
  let orgReportRegionBudgetsContractData: ContractProps = {
    data: {
      contracts: {
        orgContract: {},
        orgReportsContract: {},
        orgReportDocsContract: {},
        orgReportBudgetsContract: {},
        orgReportExpenditureContract: {},
        orgReportRecipientBudgetsContract: {},
        orgReportRegionBudgetsContract: state.chainContracts.data.contracts.orgReportRegionBudgetsContract,
        orgReportCountryBudgetsContract: {}
      }
    }
  }

  if ( !(orgReportRegionBudgetsContractData.data.contracts.orgReportRegionBudgetsContract.hasOwnProperty('getReportExists')) ) {
    getOrgReportRegionBudgetsContract({provider: props.provider}).then((orgReportRegionBudgetsContract) => {
      if ( typeof orgReportRegionBudgetsContract != "undefined" ) {
        //console.log('Storing OrgContract', orgReportsContract)
        orgReportRegionBudgetsContractData.data.contracts.orgReportRegionBudgetsContract = orgReportRegionBudgetsContract
        const add = addOrgReportRegionBudgetsContract as Function
        store.dispatch(add(orgReportRegionBudgetsContractData))
      }
    })
  }
}

export const setOrgReportCountryBudgetsContract = (props: SetProps) => {

  const store = props.store
  const state = store.getState()
  let orgReportCountryBudgetsContractData: ContractProps = {
    data: {
      contracts: {
        orgContract: {},
        orgReportsContract: {},
        orgReportDocsContract: {},
        orgReportBudgetsContract: {},
        orgReportExpenditureContract: {},
        orgReportRecipientBudgetsContract: {},
        orgReportRegionBudgetsContract: {},
        orgReportCountryBudgetsContract: state.chainContracts.data.contracts.orgReportCountryBudgetsContract
      }
    }
  }

  if ( !(orgReportCountryBudgetsContractData.data.contracts.orgReportCountryBudgetsContract.hasOwnProperty('getReportExists')) ) {
    getOrgReportCountryBudgetsContract({provider: props.provider}).then((orgReportCountryBudgetsContract) => {
      if ( typeof orgReportCountryBudgetsContract != "undefined" ) {
        //console.log('Storing OrgContract', orgReportsContract)
        orgReportCountryBudgetsContractData.data.contracts.orgReportCountryBudgetsContract = orgReportCountryBudgetsContract
        const add = addOrgReportCountryBudgetsContract as Function
        store.dispatch(add(orgReportCountryBudgetsContractData))
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

  setInterval(() => {
    setOrgReportDocsContract({store: store, provider: provider})
  }, Blockchain.checkAccountInterval)

  setInterval(() => {
    setOrgReportBudgetsContract({store: store, provider: provider})
  }, Blockchain.checkAccountInterval)

  setInterval(() => {
    setOrgReportExpenditureContract({store: store, provider: provider})
  }, Blockchain.checkAccountInterval)

  setInterval(() => {
    setOrgReportRecipientBudgetsContract({store: store, provider: provider})
  }, Blockchain.checkAccountInterval)

  setInterval(() => {
    setOrgReportRegionBudgetsContract({store: store, provider: provider})
  }, Blockchain.checkAccountInterval)

  setInterval(() => {
    setOrgReportCountryBudgetsContract({store: store, provider: provider})
  }, Blockchain.checkAccountInterval)

  /* static organisationReportDocsAddress = "0xFa9f7680705968660d36F34D080d5fEeD0614221"
  static organisationReportBudgetsAddress = "0xB155E22D9598cC0e635792070A888127Ae349B0c"
  static organisationReportExpenditureAddress = "0x85f1116DF7FCFE73bc511Ac1a6a69BaB0A3af1dA"
  static organisationReportRecipientBudgetsAddress = "0x7461eB577da59CBEE2618BB82c0d67311AE89960"
  static organisationReportRegionBudgetsAddress = "0x7952136EB509C59bFe8393a0BAeB17D3a5E0a400"
  static organisationReportCountryBudgetsAddress = "0x391Ef15D0640b87c6Fbaa555CaE2ed29dfd9F5c1" */

  const add = addInfo as Function
  store.dispatch(add(infoData))
}
