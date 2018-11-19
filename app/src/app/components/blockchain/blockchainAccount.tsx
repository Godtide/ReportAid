import { Store } from 'redux'

import { BlockchainAccountProps } from '../../store/blockchain/types'
import { addAccount } from '../../store/blockchain/actions'

interface OwnProps {
  store: Store
}

export const setAccount = (props: OwnProps): boolean => {

  //console.log('In set account')
  let result = true
  const store = props.store
  const state = store.getState()
  const provider = state.blockchain.provider
  let accountData: BlockchainAccountProps = {
    account: state.blockchain.account
  }

  if ( provider.hasOwnProperty('connection') ) {
    const signer = provider.getSigner()
    signer.getAddress().then(function(account: string) {
      if ( accountData.account != account ) {
        accountData.account = account
        console.log ('Setting account', account)
        store.dispatch(addAccount(accountData))
      }
    }).catch(function (error: any) {
      console.log(error)
      result = false
    })
  }

  return result
}
