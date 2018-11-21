import { ActionTypes, AddAction, BlockchainProps } from './types'

const initialState: BlockchainProps = {
  API: '',
  Name: '',
  ChainId: '',
  ENS: '',
  account: '',
  provider: {},
}

export const reducer = (state: BlockchainProps = initialState, action: AddAction): BlockchainProps => {
  if ( (action.type == ActionTypes.ADD_INFO ) ||
       (action.type == ActionTypes.ADD_ACCOUNT ) ||
       (action.type == ActionTypes.ADD_OBJECT )
     ) {
    //console.log(action.payload)
    return (<any>Object).assign({}, state, action.payload)
  } else {
    return state
  }
}
