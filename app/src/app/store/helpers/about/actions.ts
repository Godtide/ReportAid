import { Action, ActionCreator } from 'redux'
import { ActionTypes, AboutProps } from './types'

export interface RequestDataAction extends Action {
  type: ActionTypes.REQ_DATA
  payload: AboutProps
}

export const fetchRequest: ActionCreator<RequestDataAction> = (type, payload) => {
  console.log('blah blah blah', type, payload)
  return (
    { type: type, payload: payload }
  )
}
