import { Action, ActionCreator } from 'redux'
import { BlockchainActionTypes, BlockchainProps } from './types'

export interface RequestDataAction extends Action {
  type: BlockchainActionTypes.REQ_DATA
  payload: BlockchainProps
}

export const fetchRequest: ActionCreator<RequestDataAction> = (type, payload) => {
  console.log('blah blah blah', type, payload)
  return (
    { type: type, payload: payload }
  )
}
