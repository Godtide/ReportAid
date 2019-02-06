import { OrgReaderActionTypes, OrgReaderProps } from './types'
import { ActionProps, PayloadProps } from '../../../types'

const initialState: OrgReaderProps = {
  num: 0,
  data: {
    '': {
      num: 0,
      data: {
        '': {
          version: '',
          report: {
            reportRef: '',
            orgRef: ''
          },
          reportingOrg: {
            orgRef: '',
            orgType: 0,
            isSecondary: false
          },
          lang: '',
          currency: '',
          lastUpdatedTime: ''
        }
      }
    }
  }
}

export const reducer = (state: OrgReaderProps = initialState, action: ActionProps): OrgReaderProps => {
  //console.log('Boom!', action.type, action.payload)
  const payload = action.payload as PayloadProps
  if ( typeof payload != 'undefined' ) {
    const payloadData = payload.data as OrgReaderProps
    if ( (action.type == OrgReaderActionTypes.NUM_SUCCESS ) ||
         (action.type == OrgReaderActionTypes.NUM_SUCCESS ) ) {

      const data: OrgReaderProps = {
        num: payloadData.num,
        data: { ...state.data }
      }
      return {...data}

    } else if ( (action.type == OrgReaderActionTypes.REF_SUCCESS ) ||
         (action.type == OrgReaderActionTypes.REF_FAILURE ) ) {

       const data: OrgReaderProps  = {
         num: state.num,
         data: payloadData.data
       }

       return data

    } else if ( (action.type == OrgReaderActionTypes.NUMREP_SUCCESS ) ||
                (action.type == OrgReaderActionTypes.NUMREP_FAILURE ) ||
                (action.type == OrgReaderActionTypes.REPORT_SUCCESS ) ||
                (action.type == OrgReaderActionTypes.REPORT_FAILURE ) ) {

       const data: OrgReaderProps  = {
         num: state.num,
         data: {...payloadData.data}
       }

       //console.log (' data: ', data)
       return data

    } else {
      return state
    }
  } else {
    return state
  }
}
