import { Action, ActionCreator } from 'redux'
import { AboutActionTypes, InfoProps } from '../types'

export interface RequestDataAction extends Action {
  type: AboutActionTypes.REQ_DATA
  payload: InfoProps
}

export const fetchRequest: ActionCreator<RequestDataAction> = (type, payload) => {
  console.log('blah blah blah', type, payload)
  return (
    { type: type, payload: payload }
  )
}
