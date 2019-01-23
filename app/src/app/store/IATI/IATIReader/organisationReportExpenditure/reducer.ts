import { OrgReportExpenditureReaderActionTypes, OrgReportExpenditureReaderProps } from './types'
import { ActionProps, PayloadProps } from '../../../types'

const initialState: OrgReportExpenditureReaderProps = {
  num: 0,
  data: {
    '': {
      num: 0,
      data: {
        '': {
          reportRef: '',
          expenditureRef: '',
          expenditureLine: '',
          finance: {
            value: 0,
            status: '',
            start: '',
            end: ''
          }
        }
      }
    }
  }
}

export const reducer = (state: OrgReportExpenditureReaderProps = initialState, action: ActionProps): OrgReportExpenditureReaderProps => {
  //console.log('Boom!', action.type, action.payload)
  const payload = action.payload as PayloadProps
  if ( typeof payload != 'undefined' ) {
    const payloadData = payload.data as OrgReportExpenditureReaderProps
    if ( (action.type == OrgReportExpenditureReaderActionTypes.NUM_SUCCESS ) ||
         (action.type == OrgReportExpenditureReaderActionTypes.NUM_SUCCESS ) ) {

      const data: OrgReportExpenditureReaderProps = {
        num: payloadData.num,
        data: { ...state.data }
      }
      return {...data}

    } else if ( (action.type == OrgReportExpenditureReaderActionTypes.REF_SUCCESS ) ||
         (action.type == OrgReportExpenditureReaderActionTypes.REF_FAILURE ) ) {

       const data: OrgReportExpenditureReaderProps  = {
         num: state.num,
         data: payloadData.data
       }

       return data

    } else if ( (action.type == OrgReportExpenditureReaderActionTypes.NUMEXPENDITURE_SUCCESS ) ||
                (action.type == OrgReportExpenditureReaderActionTypes.NUMEXPENDITURE_FAILURE ) ||
                (action.type == OrgReportExpenditureReaderActionTypes.EXPENDITURE_SUCCESS ) ||
                (action.type == OrgReportExpenditureReaderActionTypes.EXPENDITURE_FAILURE ) ) {

       const data: OrgReportExpenditureReaderProps  = {
         num: state.num,
         data: {...payloadData.data}
       }

       //console.log ('Report data: ', data)
       return data

    } else {
      return state
    }
  } else {
    return state
  }
}
