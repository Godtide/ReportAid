import { BlockchainActionTypes, BlockchainAction, BlockchainProps } from './types'

const initialState: BlockchainProps = {
  API: '',
  Name: '',
  ChainId: '',
  ENS: '',
  account: '',
  provider: {},
  orgContract: {}
}

export const reducer = (state: BlockchainProps = initialState, action: BlockchainAction): BlockchainProps => {
  if ( (action.type == BlockchainActionTypes.ADD_INFO ) ||
       (action.type == BlockchainActionTypes.ADD_ACCOUNT ) ||
       (action.type == BlockchainActionTypes.ADD_OBJECT ) ||
       (action.type == BlockchainActionTypes.ADD_ORGCONTRACT )
     ) {
    //console.log(action.payload)
    return (<any>Object).assign({}, state, action.payload)
  } else {
    return state
  }
}
