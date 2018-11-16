import { Store } from 'redux'

import { Blockchain } from '../../../utils/config'

import { setProvider } from './blockchainProvider'
import { setAccount } from './blockchainAccount'

interface BlockchainProviderProps {
  store: Store
}

export const setBlockchain = (props: BlockchainProviderProps) => {

  (window as any).setInterval(() => {
    setProvider({store: props.store})
    setAccount({store: props.store})
  }, Blockchain.interval)

}
