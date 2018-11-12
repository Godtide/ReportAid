import { Reducer } from 'redux'
import { BlockchainActionTypes, BlockchainProps } from './types'
import { BlockchainStrings } from '../../utils/strings'

const initialState: BlockchainProps = {
  title: BlockchainStrings.heading,
  APIProvider: '',
  networkName: '',
  networkChainId: '',
  networkENSAddress: '',
  account: '',
}

export const reducer: Reducer<BlockchainProps> = (state = initialState, action): BlockchainProps => {
  switch (action.type) {
    case BlockchainActionTypes.REQ_DATA:
      return (<any>Object).assign({}, state, {
        title: action.payload.title,
        APIProvider: action.payload.APIProvider,
        networkName: action.payload.networkName,
        networkChainId: action.payload.networkChainId,
        networkENSAddress: action.payload.networkENSAddress,
        account: action.payload.account,
      })
    default:
      return state
  }
}
