import { OrganisationBudgetsReaderActionTypes, OrganisationBudgetsReaderProps } from './types'
import { ActionProps, PayloadProps } from '../../../types'

const initialState: OrganisationBudgetsReaderProps = {
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

export const reducer = (state: OrganisationBudgetsReaderProps = initialState, action: ActionProps): OrganisationBudgetsReaderProps => {
  //console.log('Boom!', action.type, action.payload)
  const payload = action.payload as PayloadProps
  if ( typeof payload != 'undefined' ) {
    const payloadData = payload.data as OrganisationBudgetsReaderProps
    if ( (action.type == OrganisationBudgetsReaderActionTypes.NUM_SUCCESS ) ||
         (action.type == OrganisationBudgetsReaderActionTypes.NUM_SUCCESS ) ) {

      const data: OrganisationBudgetsReaderProps = {
        num: payloadData.num,
        data: { ...state.data }
      }
      return {...data}

    } else if ( (action.type == OrganisationBudgetsReaderActionTypes.REF_SUCCESS ) ||
         (action.type == OrganisationBudgetsReaderActionTypes.REF_FAILURE ) ) {

       const data: OrganisationBudgetsReaderProps  = {
         num: state.num,
         data: payloadData.data
       }

       return data

    } else if ( (action.type == OrganisationBudgetsReaderActionTypes.NUMBUDGET_SUCCESS ) ||
                (action.type == OrganisationBudgetsReaderActionTypes.NUMBUDGET_FAILURE ) ||
                (action.type == OrganisationBudgetsReaderActionTypes.BUDGET_SUCCESS ) ||
                (action.type == OrganisationBudgetsReaderActionTypes.BUDGET_FAILURE ) ) {

       const data: OrganisationBudgetsReaderProps  = {
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
