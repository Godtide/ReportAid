import { OrgRecipientBudgetsReaderActionTypes, OrgRecipientBudgetsReaderProps } from './types'
import { ActionProps, PayloadProps } from '../../../types'

const initialState: OrgRecipientBudgetsReaderProps = {
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
          orgRef: '',
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

export const reducer = (state: OrgRecipientBudgetsReaderProps = initialState, action: ActionProps): OrgRecipientBudgetsReaderProps => {
  //console.log('Boom!', action.type, action.payload)
  const payload = action.payload as PayloadProps
  if ( typeof payload != 'undefined' ) {
    const payloadData = payload.data as OrgRecipientBudgetsReaderProps
    if ( (action.type == OrgRecipientBudgetsReaderActionTypes.NUM_SUCCESS ) ||
         (action.type == OrgRecipientBudgetsReaderActionTypes.NUM_SUCCESS ) ) {

      const data: OrgRecipientBudgetsReaderProps = {
        num: payloadData.num,
        data: { ...state.data }
      }
      return {...data}

    } else if ( (action.type == OrgRecipientBudgetsReaderActionTypes.REF_SUCCESS ) ||
         (action.type == OrgRecipientBudgetsReaderActionTypes.REF_FAILURE ) ) {

       const data: OrgRecipientBudgetsReaderProps  = {
         num: state.num,
         data: payloadData.data
       }

       return data

    } else if ( (action.type == OrgRecipientBudgetsReaderActionTypes.NUMBUDGET_SUCCESS ) ||
                (action.type == OrgRecipientBudgetsReaderActionTypes.NUMBUDGET_FAILURE ) ||
                (action.type == OrgRecipientBudgetsReaderActionTypes.BUDGET_SUCCESS ) ||
                (action.type == OrgRecipientBudgetsReaderActionTypes.BUDGET_FAILURE ) ) {

       const data: OrgRecipientBudgetsReaderProps  = {
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
