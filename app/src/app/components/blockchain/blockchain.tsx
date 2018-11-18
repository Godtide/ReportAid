import { Store } from 'redux'

import { Blockchain } from '../../utils/config'

import { SetProvider } from './blockchainProvider'
import { setAccount } from './blockchainAccount'

interface BlockchainProviderProps {
  store: Store
}

export const setBlockchain = (props: BlockchainProviderProps) => {

  const provider = new SetProvider()
  (window as Window).setInterval(() => {
    setAccount({store: props.store})
  }, Blockchain.interval)
}
