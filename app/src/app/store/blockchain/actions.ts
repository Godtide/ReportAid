import { InfoProps,
         AccountProps,
         ObjectProps,
         ActionTypes,
         AddInfoAction,
         AddAccountAction,
         AddObjectAction } from './types'

export const addInfo = (payload: InfoProps): AddInfoAction => {
  return {
    type: ActionTypes.ADD_INFO,
    payload
  }
}

export const addAccount = (payload: AccountProps): AddAccountAction => {
  //console.log('In payload', payload)
  return {
    type: ActionTypes.ADD_ACCOUNT,
    payload
  }
}

export const addObject = (payload: ObjectProps): AddObjectAction => {
  return {
    type: ActionTypes.ADD_OBJECT,
    payload
  }
}
