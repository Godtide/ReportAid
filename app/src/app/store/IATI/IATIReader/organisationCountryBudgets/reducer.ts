import { OrganisationCountryBudgetsReaderActionTypes, OrganisationCountryBudgetsReaderProps } from './types'
import { ActionProps, PayloadProps } from '../../../types'

const initialState: OrganisationCountryBudgetsReaderProps = {
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

export const reducer = (state: OrganisationCountryBudgetsReaderProps = initialState, action: ActionProps): OrganisationCountryBudgetsReaderProps => {
  //console.log('Boom!', action.type, action.payload)
  const payload = action.payload as PayloadProps
  if ( typeof payload != 'undefined' ) {
    const payloadData = payload.data as OrganisationCountryBudgetsReaderProps
    if ( (action.type == OrganisationCountryBudgetsReaderActionTypes.NUM_SUCCESS ) ||
         (action.type == OrganisationCountryBudgetsReaderActionTypes.NUM_SUCCESS ) ) {

      const data: OrganisationCountryBudgetsReaderProps = {
        num: payloadData.num,
        data: { ...state.data }
      }
      return {...data}

    } else if ( (action.type == OrganisationCountryBudgetsReaderActionTypes.REF_SUCCESS ) ||
         (action.type == OrganisationCountryBudgetsReaderActionTypes.REF_FAILURE ) ) {

       const data: OrganisationCountryBudgetsReaderProps  = {
         num: state.num,
         data: payloadData.data
       }

       return data

    } else if ( (action.type == OrganisationCountryBudgetsReaderActionTypes.NUMBUDGET_SUCCESS ) ||
                (action.type == OrganisationCountryBudgetsReaderActionTypes.NUMBUDGET_FAILURE ) ||
                (action.type == OrganisationCountryBudgetsReaderActionTypes.BUDGET_SUCCESS ) ||
                (action.type == OrganisationCountryBudgetsReaderActionTypes.BUDGET_FAILURE ) ) {

       const data: OrganisationCountryBudgetsReaderProps  = {
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
