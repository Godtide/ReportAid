import { OrgReportBudgetsReaderActionTypes, OrgReportBudgetsReaderProps } from './types'
import { ActionProps, PayloadProps } from '../../../types'

const initialState: OrgReportBudgetsReaderProps = {
  num: 0,
  data: {
    '': {
      num: 0,
      data: {
        '': {
          report: {
            reportRef: '',
            orgRef: ''
          },
          budgetRef: '',
          budgetLine: '',
          finance: {
            value: 0,
            status: 0,
            start: '',
            end: ''
          }
        }
      }
    }
  }
}

export const reducer = (state: OrgReportBudgetsReaderProps = initialState, action: ActionProps): OrgReportBudgetsReaderProps => {
  //console.log('Boom!', action.type, action.payload)
  const payload = action.payload as PayloadProps
  if ( typeof payload != 'undefined' ) {
    const payloadData = payload.data as OrgReportBudgetsReaderProps
    if ( (action.type == OrgReportBudgetsReaderActionTypes.NUM_SUCCESS ) ||
         (action.type == OrgReportBudgetsReaderActionTypes.NUM_SUCCESS ) ) {

      const data: OrgReportBudgetsReaderProps = {
        num: payloadData.num,
        data: { ...state.data }
      }
      return {...data}

    } else if ( (action.type == OrgReportBudgetsReaderActionTypes.REF_SUCCESS ) ||
         (action.type == OrgReportBudgetsReaderActionTypes.REF_FAILURE ) ) {

       const data: OrgReportBudgetsReaderProps  = {
         num: state.num,
         data: payloadData.data
       }

       return data

    } else if ( (action.type == OrgReportBudgetsReaderActionTypes.NUMBUDGET_SUCCESS ) ||
                (action.type == OrgReportBudgetsReaderActionTypes.NUMBUDGET_FAILURE ) ||
                (action.type == OrgReportBudgetsReaderActionTypes.BUDGET_SUCCESS ) ||
                (action.type == OrgReportBudgetsReaderActionTypes.BUDGET_FAILURE ) ) {

       const data: OrgReportBudgetsReaderProps  = {
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
