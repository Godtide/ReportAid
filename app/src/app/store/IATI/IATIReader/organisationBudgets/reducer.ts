import { OrgBudgetsReaderActionTypes, OrgBudgetsReaderProps } from './types'
import { ActionProps, PayloadProps } from '../../../types'

const initialState: OrgBudgetsReaderProps = {
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

export const reducer = (state: OrgBudgetsReaderProps = initialState, action: ActionProps): OrgBudgetsReaderProps => {
  //console.log('Boom!', action.type, action.payload)
  const payload = action.payload as PayloadProps
  if ( typeof payload != 'undefined' ) {
    const payloadData = payload.data as OrgBudgetsReaderProps
    if ( (action.type == OrgBudgetsReaderActionTypes.NUM_SUCCESS ) ||
         (action.type == OrgBudgetsReaderActionTypes.NUM_SUCCESS ) ) {

      const data: OrgBudgetsReaderProps = {
        num: payloadData.num,
        data: { ...state.data }
      }
      return {...data}

    } else if ( (action.type == OrgBudgetsReaderActionTypes.REF_SUCCESS ) ||
         (action.type == OrgBudgetsReaderActionTypes.REF_FAILURE ) ) {

       const data: OrgBudgetsReaderProps  = {
         num: state.num,
         data: payloadData.data
       }

       return data

    } else if ( (action.type == OrgBudgetsReaderActionTypes.NUMBUDGET_SUCCESS ) ||
                (action.type == OrgBudgetsReaderActionTypes.NUMBUDGET_FAILURE ) ||
                (action.type == OrgBudgetsReaderActionTypes.BUDGET_SUCCESS ) ||
                (action.type == OrgBudgetsReaderActionTypes.BUDGET_FAILURE ) ) {

       const data: OrgBudgetsReaderProps  = {
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
