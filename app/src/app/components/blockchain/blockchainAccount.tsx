import { Store } from 'redux'

import { AccountProps } from '../../store/blockchain/account/types'

import { addAccount } from '../../store/blockchain/account/actions'

interface ChainProps {
  store: Store
}

export const setAccount = async (props: ChainProps) => {

  const store = props.store
  const state = store.getState()
  const provider = state.chainInfo.data.provider
  if ( provider.hasOwnProperty('connection') ) {

    let accountData: AccountProps = {
      data: {
        account: state.chainAccount.data.account
      }
    }
    const signer = provider.getSigner()
    const account = await signer.getAddress()
    if ( accountData.data.account != account ) {
      //console.log('Storing Account', account)
      accountData.data.account = account

      const add = addAccount as Function
      store.dispatch(add(accountData))
    }
  }
}
