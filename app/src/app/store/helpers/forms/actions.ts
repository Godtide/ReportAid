import { ThunkDispatch } from 'redux-thunk'
import { ApplicationState } from '../../../store'

import { storeAction } from '../../actions'

import { ActionProps, PayloadProps } from '../../types'
import { FormActionTypes, FormData } from './types'

export const write = (payload: PayloadProps): Function => {
  return (actionType: FormActionTypes): PayloadProps => {
    const getProps = storeAction(actionType)(payload) as PayloadProps
    return getProps
  }
}

export const setFormFunctions = (formProps: FormData) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>) => {
    const formFunctions: FormData = {
      submitFunc: formProps.submitFunc,
      resetFunc: formProps.resetFunc
    }
    await dispatch(write({data: formFunctions})(FormActionTypes.FORMFUNCTION_SUCCESS))
  }
}
