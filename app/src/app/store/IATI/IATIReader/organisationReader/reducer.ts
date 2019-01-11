import { OrgGetActionTypes, OrgGetProps } from './types'
import { ActionProps, PayloadProps } from '../../../types'

const initialState: OrgGetProps = {
  num: 0,
  data: {
    '': {
      orgRef: '',
      name: '',
      identifier: ''
    }
  }
}

export const reducer = (state: OrgGetProps = initialState, action: ActionProps): OrgGetProps => {
  //console.log('Boom!', action.type, action.payload)
  const payload = action.payload as PayloadProps
  if ( typeof payload != 'undefined' ) {
    const payloadData = payload.data as OrgGetProps
    if ( (action.type == OrgGetActionTypes.NUM_SUCCESS ) ||
         (action.type == OrgGetActionTypes.NUM_SUCCESS ) ) {

      const data: OrgGetProps = {
        num: payloadData.num,
        data: { ...state.data }
      }
      return {...data}

    } else if ( (action.type == OrgGetActionTypes.REF_SUCCESS ) ||
         (action.type == OrgGetActionTypes.REF_FAILURE ) ) {

       const data: OrgGetProps  = {
         num: state.num,
         data: payloadData.data
       }

       return data

    } else if ( (action.type == OrgGetActionTypes.ORG_SUCCESS ) ||
         (action.type == OrgGetActionTypes.ORG_FAILURE ) ||
         (action.type == OrgGetActionTypes.NAME_SUCCESS ) ||
         (action.type == OrgGetActionTypes.NAME_FAILURE ) ||
         (action.type == OrgGetActionTypes.ID_SUCCESS ) ||
         (action.type == OrgGetActionTypes.ID_FAILURE ) ) {

       const data: OrgGetProps  = {
         num: state.num,
         data: {...payloadData.data}
       }
       return data

    } else {
      return state
    }
  } else {
    return state
  }
}
