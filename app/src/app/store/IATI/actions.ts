import { ThunkDispatch } from 'redux-thunk'
import { ApplicationState } from '../../store'

import { storeAction } from '../actions'

import { initialise as txInitialise } from './IATIWriter/actions'
import { initialise as readDataInitialise } from './IATIReader/actions'

import { ActionProps, PayloadProps } from '../types'
import { IATIReportActionTypes } from './types'

export const initialise = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    await dispatch(txInitialise())
    await dispatch(readDataInitialise())
  }
}
