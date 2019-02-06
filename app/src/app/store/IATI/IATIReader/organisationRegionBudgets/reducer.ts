import { OrgRegionBudgetsReaderActionTypes, OrgRegionBudgetsReaderProps } from './types'
import { ActionProps, PayloadProps } from '../../../types'

const initialState: OrgRegionBudgetsReaderProps = {
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

export const reducer = (state: OrgRegionBudgetsReaderProps = initialState, action: ActionProps): OrgRegionBudgetsReaderProps => {
  //console.log('Boom!', action.type, action.payload)
  const payload = action.payload as PayloadProps
  if ( typeof payload != 'undefined' ) {
    const payloadData = payload.data as OrgRegionBudgetsReaderProps
    if ( (action.type == OrgRegionBudgetsReaderActionTypes.NUM_SUCCESS ) ||
         (action.type == OrgRegionBudgetsReaderActionTypes.NUM_SUCCESS ) ) {

      const data: OrgRegionBudgetsReaderProps = {
        num: payloadData.num,
        data: { ...state.data }
      }
      return {...data}

    } else if ( (action.type == OrgRegionBudgetsReaderActionTypes.REF_SUCCESS ) ||
         (action.type == OrgRegionBudgetsReaderActionTypes.REF_FAILURE ) ) {

       const data: OrgRegionBudgetsReaderProps  = {
         num: state.num,
         data: payloadData.data
       }

       return data

    } else if ( (action.type == OrgRegionBudgetsReaderActionTypes.NUMBUDGET_SUCCESS ) ||
                (action.type == OrgRegionBudgetsReaderActionTypes.NUMBUDGET_FAILURE ) ||
                (action.type == OrgRegionBudgetsReaderActionTypes.BUDGET_SUCCESS ) ||
                (action.type == OrgRegionBudgetsReaderActionTypes.BUDGET_FAILURE ) ) {

       const data: OrgRegionBudgetsReaderProps  = {
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
