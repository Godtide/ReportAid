import { Store } from 'redux'
import Web3 from 'web3'

import { BlockchainAccountProps } from '../../store/blockchain/types'
import { addAccount } from '../../store/blockchain/actions'

interface OwnProps {
  store: Store
}

export const setAccount = (props: OwnProps): boolean => {

  let result = true
  const store = props.store
  const state = store.getState()
  const web3 = state.blockchain.web3 as Web3
  let accountData: BlockchainAccountProps = {
    account: state.blockchain.account
  }

  if ( web3.hasOwnProperty('version') ) {

    web3.eth.getAccounts().then(function(accounts: any) {
      if ( accountData.account != accounts[0] ) {
        //console.log('In setting account')
        accountData.account = accounts[0]
        web3.eth.defaultAccount = accountData.account
        store.dispatch(addAccount(accountData))
      }
    }).catch(function (error: any) {
      console.log(error)
      result = false
    })
  }

  return result
}
