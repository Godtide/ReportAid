import { Store } from 'redux'
import { Blockchain } from '../../utils/config'

import { setProvider } from './blockchainProvider'
import { setAccount } from './blockchainAccount'
import { setContracts } from './blockchainContracts'

interface SetBlockchainProps {
  store: Store
}

export const setBlockchain = async (props: SetBlockchainProps) => {

  setProvider({store: props.store})
  setAccount({store: props.store})
  setContracts({store: props.store,})
}
