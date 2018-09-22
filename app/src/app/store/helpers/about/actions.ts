import { Action, ActionCreator } from 'redux'
import { AboutActionTypes, AboutProps } from './types'

export interface AboutRequestDataAction extends Action {
  type: AboutActionTypes.REQ_DATA
  payload: AboutProps
}

export const AboutFetchRequest: ActionCreator<AboutRequestDataAction> = (type, payload) => {
  console.log('blah', type, payload)
  return (
    { type: type, payload: payload }
  )
}
