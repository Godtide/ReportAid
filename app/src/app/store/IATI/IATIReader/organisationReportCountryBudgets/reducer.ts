import { OrgReportCountryBudgetsReaderActionTypes, OrgReportCountryBudgetsReaderProps } from './types'
import { ActionProps, PayloadProps } from '../../../types'

const initialState: OrgReportCountryBudgetsReaderProps = {
  num: 0,
  data: {
    '': {
      num: 0,
      data: {
        '': {
          reportRef: '',
          budgetRef: '',
          countryRef: '',
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

export const reducer = (state: OrgReportCountryBudgetsReaderProps = initialState, action: ActionProps): OrgReportCountryBudgetsReaderProps => {
  //console.log('Boom!', action.type, action.payload)
  const payload = action.payload as PayloadProps
  if ( typeof payload != 'undefined' ) {
    const payloadData = payload.data as OrgReportCountryBudgetsReaderProps
    if ( (action.type == OrgReportCountryBudgetsReaderActionTypes.NUM_SUCCESS ) ||
         (action.type == OrgReportCountryBudgetsReaderActionTypes.NUM_SUCCESS ) ) {

      const data: OrgReportCountryBudgetsReaderProps = {
        num: payloadData.num,
        data: { ...state.data }
      }
      return {...data}

    } else if ( (action.type == OrgReportCountryBudgetsReaderActionTypes.REF_SUCCESS ) ||
         (action.type == OrgReportCountryBudgetsReaderActionTypes.REF_FAILURE ) ) {

       const data: OrgReportCountryBudgetsReaderProps  = {
         num: state.num,
         data: payloadData.data
       }

       return data

    } else if ( (action.type == OrgReportCountryBudgetsReaderActionTypes.NUMBUDGET_SUCCESS ) ||
                (action.type == OrgReportCountryBudgetsReaderActionTypes.NUMBUDGET_FAILURE ) ||
                (action.type == OrgReportCountryBudgetsReaderActionTypes.BUDGET_SUCCESS ) ||
                (action.type == OrgReportCountryBudgetsReaderActionTypes.BUDGET_FAILURE ) ) {

       const data: OrgReportCountryBudgetsReaderProps  = {
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
