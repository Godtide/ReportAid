import { connect } from 'react-redux'
import { Store } from 'redux'

import { ApplicationState } from '../../store'

import Web3 from 'web3'
import { ethers } from 'ethers'

//import { fetchRequest, RequestDataAction } from '../../../store/helpers/about/actions'
import { BlockchainAccountProps } from '../../store/blockchain/types'
import { addAccount } from '../../store/blockchain/actions'

interface OwnProps {
  store: Store
}

export const setAccount = async (props: OwnProps) => {

  console.log('In setting account')
  const store = props.store
  const state = store.getState()
  const web3 = state.blockchain.web3 as Web3
  if ( web3.hasOwnProperty('version') ) {
    console.log('Setting account!')
    console.log(web3)

    let accountData: BlockchainAccountProps = {
      account: ''
    }

    await web3.eth.getAccounts((error: any, accounts: any) => {
      accountData.account = accounts[0]
    })
    .catch((error: any) => {
       console.log('Error!: ', error)
    })
    web3.eth.defaultAccount = accountData.account
    store.dispatch(addAccount(accountData))
    return true
  } else {
    console.log('No accouint')
  }
}
