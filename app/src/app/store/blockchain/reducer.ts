import { BlockchainActionTypes, BlockchainAction, BlockchainInfoProps } from './types'

const initialState: BlockchainInfoProps = {
  APIProvider: '',
  networkName: '',
  networkChainId: '',
  networkENSAddress: '',
  account: ''
}

export const blockchainReducer = (state: BlockchainInfoProps = initialState, action: BlockchainAction): BlockchainInfoProps => {
  switch (action.type) {
    case BlockchainActionTypes.ADD_DATA:
      return (<any>Object).assign({}, state, action.payload)
    default:
      return state
  }
}
