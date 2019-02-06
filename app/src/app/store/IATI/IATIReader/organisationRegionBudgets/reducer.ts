import { OrganisationRegionBudgetsReaderActionTypes, OrganisationRegionBudgetsReaderProps } from './types'
import { ActionProps, PayloadProps } from '../../../types'

const initialState: OrganisationRegionBudgetsReaderProps = {
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

export const reducer = (state: OrganisationRegionBudgetsReaderProps = initialState, action: ActionProps): OrganisationRegionBudgetsReaderProps => {
  //console.log('Boom!', action.type, action.payload)
  const payload = action.payload as PayloadProps
  if ( typeof payload != 'undefined' ) {
    const payloadData = payload.data as OrganisationRegionBudgetsReaderProps
    if ( (action.type == OrganisationRegionBudgetsReaderActionTypes.NUM_SUCCESS ) ||
         (action.type == OrganisationRegionBudgetsReaderActionTypes.NUM_SUCCESS ) ) {

      const data: OrganisationRegionBudgetsReaderProps = {
        num: payloadData.num,
        data: { ...state.data }
      }
      return {...data}

    } else if ( (action.type == OrganisationRegionBudgetsReaderActionTypes.REF_SUCCESS ) ||
         (action.type == OrganisationRegionBudgetsReaderActionTypes.REF_FAILURE ) ) {

       const data: OrganisationRegionBudgetsReaderProps  = {
         num: state.num,
         data: payloadData.data
       }

       return data

    } else if ( (action.type == OrganisationRegionBudgetsReaderActionTypes.NUMBUDGET_SUCCESS ) ||
                (action.type == OrganisationRegionBudgetsReaderActionTypes.NUMBUDGET_FAILURE ) ||
                (action.type == OrganisationRegionBudgetsReaderActionTypes.BUDGET_SUCCESS ) ||
                (action.type == OrganisationRegionBudgetsReaderActionTypes.BUDGET_FAILURE ) ) {

       const data: OrganisationRegionBudgetsReaderProps  = {
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
