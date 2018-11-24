import { BlockchainActionTypes, BlockchainAddAction, BlockchainProps } from './types'

const initialState: BlockchainProps = {
  API: '',
  Name: '',
  ChainId: '',
  ENS: '',
  account: '',
  provider: {},
}

export const reducer = (state: BlockchainProps = initialState, action: BlockchainAddAction): BlockchainProps => {
  if ( (action.type == BlockchainActionTypes.ADD_INFO ) ||
       (action.type == BlockchainActionTypes.ADD_ACCOUNT ) ||
       (action.type == BlockchainActionTypes.ADD_OBJECT )
     ) {
    //console.log(action.payload)
    return (<any>Object).assign({}, state, action.payload)
  } else {
    return state
  }
}
