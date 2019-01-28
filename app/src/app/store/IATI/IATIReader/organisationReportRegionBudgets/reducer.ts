import { OrgReportRegionBudgetsReaderActionTypes, OrgReportRegionBudgetsReaderProps } from './types'
import { ActionProps, PayloadProps } from '../../../types'

const initialState: OrgReportRegionBudgetsReaderProps = {
  num: 0,
  data: {
    '': {
      num: 0,
      data: {
        '': {
          reportRef: '',
          budgetRef: '',
          regionRef: 0,
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

export const reducer = (state: OrgReportRegionBudgetsReaderProps = initialState, action: ActionProps): OrgReportRegionBudgetsReaderProps => {
  //console.log('Boom!', action.type, action.payload)
  const payload = action.payload as PayloadProps
  if ( typeof payload != 'undefined' ) {
    const payloadData = payload.data as OrgReportRegionBudgetsReaderProps
    if ( (action.type == OrgReportRegionBudgetsReaderActionTypes.NUM_SUCCESS ) ||
         (action.type == OrgReportRegionBudgetsReaderActionTypes.NUM_SUCCESS ) ) {

      const data: OrgReportRegionBudgetsReaderProps = {
        num: payloadData.num,
        data: { ...state.data }
      }
      return {...data}

    } else if ( (action.type == OrgReportRegionBudgetsReaderActionTypes.REF_SUCCESS ) ||
         (action.type == OrgReportRegionBudgetsReaderActionTypes.REF_FAILURE ) ) {

       const data: OrgReportRegionBudgetsReaderProps  = {
         num: state.num,
         data: payloadData.data
       }

       return data

    } else if ( (action.type == OrgReportRegionBudgetsReaderActionTypes.NUMBUDGET_SUCCESS ) ||
                (action.type == OrgReportRegionBudgetsReaderActionTypes.NUMBUDGET_FAILURE ) ||
                (action.type == OrgReportRegionBudgetsReaderActionTypes.BUDGET_SUCCESS ) ||
                (action.type == OrgReportRegionBudgetsReaderActionTypes.BUDGET_FAILURE ) ) {

       const data: OrgReportRegionBudgetsReaderProps  = {
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
