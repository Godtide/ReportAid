import { OrgCountryBudgetsReaderActionTypes, OrgCountryBudgetsReaderProps } from './types'
import { ActionProps, PayloadProps } from '../../../types'

const initialState: OrgCountryBudgetsReaderProps = {
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

export const reducer = (state: OrgCountryBudgetsReaderProps = initialState, action: ActionProps): OrgCountryBudgetsReaderProps => {
  //console.log('Boom!', action.type, action.payload)
  const payload = action.payload as PayloadProps
  if ( typeof payload != 'undefined' ) {
    const payloadData = payload.data as OrgCountryBudgetsReaderProps
    if ( (action.type == OrgCountryBudgetsReaderActionTypes.NUM_SUCCESS ) ||
         (action.type == OrgCountryBudgetsReaderActionTypes.NUM_SUCCESS ) ) {

      const data: OrgCountryBudgetsReaderProps = {
        num: payloadData.num,
        data: { ...state.data }
      }
      return {...data}

    } else if ( (action.type == OrgCountryBudgetsReaderActionTypes.REF_SUCCESS ) ||
         (action.type == OrgCountryBudgetsReaderActionTypes.REF_FAILURE ) ) {

       const data: OrgCountryBudgetsReaderProps  = {
         num: state.num,
         data: payloadData.data
       }

       return data

    } else if ( (action.type == OrgCountryBudgetsReaderActionTypes.NUMBUDGET_SUCCESS ) ||
                (action.type == OrgCountryBudgetsReaderActionTypes.NUMBUDGET_FAILURE ) ||
                (action.type == OrgCountryBudgetsReaderActionTypes.BUDGET_SUCCESS ) ||
                (action.type == OrgCountryBudgetsReaderActionTypes.BUDGET_FAILURE ) ) {

       const data: OrgCountryBudgetsReaderProps  = {
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
