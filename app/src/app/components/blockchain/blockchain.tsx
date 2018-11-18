import { Store } from 'redux'

import { Blockchain } from '../../utils/config'

import { setProvider } from './blockchainProvider'
import { setAccount } from './blockchainAccount'

interface BlockchainProviderProps {
  store: Store
}

export const setBlockchain = (props: BlockchainProviderProps) => {

  if (setProvider({store: props.store}) ) {
    (window as Window).setInterval(() => {
      setAccount({store: props.store})
    }, Blockchain.interval)
  }
}
