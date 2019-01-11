import { OrgReportReaderActionTypes, OrgReportReaderProps } from './types'
import { ActionProps, PayloadProps } from '../../../types'

const initialState: OrgReportReaderProps = {
  num: 0,
  data: {
    '': {
      '': {
        version: '',
        orgRef: '',
        reportRef: '',
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

export const reducer = (state: OrgReportReaderProps = initialState, action: ActionProps): OrgReportReaderProps => {
  //console.log('Boom!', action.type, action.payload)
  const payload = action.payload as PayloadProps
  if ( typeof payload != 'undefined' ) {
    const payloadData = payload.data as OrgReportReaderProps
    if ( (action.type == OrgReportReaderActionTypes.NUM_SUCCESS ) ||
         (action.type == OrgReportReaderActionTypes.NUM_SUCCESS ) ) {

      const data: OrgReportReaderProps = {
        num: payloadData.num,
        data: { ...state.data }
      }
      return {...data}

    } else if ( (action.type == OrgReportReaderActionTypes.REF_SUCCESS ) ||
         (action.type == OrgReportReaderActionTypes.REF_FAILURE ) ) {

       const data: OrgReportReaderProps  = {
         num: state.num,
         data: payloadData.data
       }

       return data

    } else if ( (action.type == OrgReportReaderActionTypes.REPORT_SUCCESS ) ||
         (action.type == OrgReportReaderActionTypes.REPORT_FAILURE )  ) {

       const data: OrgReportReaderProps  = {
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
