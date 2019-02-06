import { OrganisationRecipientBudgetsReaderActionTypes, OrganisationRecipientBudgetsReaderProps } from './types'
import { ActionProps, PayloadProps } from '../../../types'

const initialState: OrganisationRecipientBudgetsReaderProps = {
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

export const reducer = (state: OrganisationRecipientBudgetsReaderProps = initialState, action: ActionProps): OrganisationRecipientBudgetsReaderProps => {
  //console.log('Boom!', action.type, action.payload)
  const payload = action.payload as PayloadProps
  if ( typeof payload != 'undefined' ) {
    const payloadData = payload.data as OrganisationRecipientBudgetsReaderProps
    if ( (action.type == OrganisationRecipientBudgetsReaderActionTypes.NUM_SUCCESS ) ||
         (action.type == OrganisationRecipientBudgetsReaderActionTypes.NUM_SUCCESS ) ) {

      const data: OrganisationRecipientBudgetsReaderProps = {
        num: payloadData.num,
        data: { ...state.data }
      }
      return {...data}

    } else if ( (action.type == OrganisationRecipientBudgetsReaderActionTypes.REF_SUCCESS ) ||
         (action.type == OrganisationRecipientBudgetsReaderActionTypes.REF_FAILURE ) ) {

       const data: OrganisationRecipientBudgetsReaderProps  = {
         num: state.num,
         data: payloadData.data
       }

       return data

    } else if ( (action.type == OrganisationRecipientBudgetsReaderActionTypes.NUMBUDGET_SUCCESS ) ||
                (action.type == OrganisationRecipientBudgetsReaderActionTypes.NUMBUDGET_FAILURE ) ||
                (action.type == OrganisationRecipientBudgetsReaderActionTypes.BUDGET_SUCCESS ) ||
                (action.type == OrganisationRecipientBudgetsReaderActionTypes.BUDGET_FAILURE ) ) {

       const data: OrganisationRecipientBudgetsReaderProps  = {
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
