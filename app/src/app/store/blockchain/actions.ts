import { BlockchainInfoProps, BlockchainActionTypes, AddDataAction } from './types'

export const addData = (payload: BlockchainInfoProps): AddDataAction => {
  return {
    type: BlockchainActionTypes.ADD_DATA,
    payload
  }
}
