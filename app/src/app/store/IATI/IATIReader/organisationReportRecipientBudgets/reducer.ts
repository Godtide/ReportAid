import { OrgReportRecipientBudgetsReaderActionTypes, OrgReportRecipientBudgetsReaderProps } from './types'
import { ActionProps, PayloadProps } from '../../../types'

const initialState: OrgReportRecipientBudgetsReaderProps = {
  num: 0,
  data: {
    '': {
      num: 0,
      data: {
        '': {
          reportRef: '',
          budgetRef: '',
          orgRef: '',
          budgetLine: '',
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

export const reducer = (state: OrgReportRecipientBudgetsReaderProps = initialState, action: ActionProps): OrgReportRecipientBudgetsReaderProps => {
  //console.log('Boom!', action.type, action.payload)
  const payload = action.payload as PayloadProps
  if ( typeof payload != 'undefined' ) {
    const payloadData = payload.data as OrgReportRecipientBudgetsReaderProps
    if ( (action.type == OrgReportRecipientBudgetsReaderActionTypes.NUM_SUCCESS ) ||
         (action.type == OrgReportRecipientBudgetsReaderActionTypes.NUM_SUCCESS ) ) {

      const data: OrgReportRecipientBudgetsReaderProps = {
        num: payloadData.num,
        data: { ...state.data }
      }
      return {...data}

    } else if ( (action.type == OrgReportRecipientBudgetsReaderActionTypes.REF_SUCCESS ) ||
         (action.type == OrgReportRecipientBudgetsReaderActionTypes.REF_FAILURE ) ) {

       const data: OrgReportRecipientBudgetsReaderProps  = {
         num: state.num,
         data: payloadData.data
       }

       return data

    } else if ( (action.type == OrgReportRecipientBudgetsReaderActionTypes.NUMBUDGET_SUCCESS ) ||
                (action.type == OrgReportRecipientBudgetsReaderActionTypes.NUMBUDGET_FAILURE ) ||
                (action.type == OrgReportRecipientBudgetsReaderActionTypes.BUDGET_SUCCESS ) ||
                (action.type == OrgReportRecipientBudgetsReaderActionTypes.BUDGET_FAILURE ) ) {

       const data: OrgReportRecipientBudgetsReaderProps  = {
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
