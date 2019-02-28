import { ThunkDispatch } from 'redux-thunk'
import { ApplicationState } from '../../store'

import { storeAction } from '../../actions'

import { ActionProps, PayloadProps } from '../../types'
import { IATIReportActionTypes } from './types'

export const read = (payload: PayloadProps): Function => {
  return (actionType: IATIReportActionTypes): PayloadProps => {
    const getProps = storeAction(actionType)(payload) as PayloadProps
    return getProps
  }
}

export const initialise = (props: IATIReportActionTypes) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    await dispatch(read({data: { data: {} }})(props))
  }
}
