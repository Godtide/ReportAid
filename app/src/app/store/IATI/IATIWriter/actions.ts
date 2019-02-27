import { ThunkDispatch } from 'redux-thunk'
import { ApplicationState } from '../../store'

import { storeAction } from '../../actions'

import { ActionProps, PayloadProps, TxProps, TxReport } from '../../types'
import { IATIWriterActionTypes } from './types'

export const write = (payload: PayloadProps): Function => {
  return (actionType: IATIWriterActionTypes): TxProps => {
    const writerProps = storeAction(actionType)(payload) as TxProps
    return writerProps
  }
}

export const initialise = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const initData: TxReport = {}
    await dispatch(write({data: initData})(IATIWriterActionTypes.TX_INIT))
  }
}
