import { OrganisationExpenditureReaderActionTypes, OrganisationExpenditureReaderProps } from './types'
import { ActionProps, PayloadProps } from '../../../types'

const initialState: OrganisationExpenditureReaderProps = {
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
          expenditureRef: '',
          expenditureLine: '',
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

export const reducer = (state: OrganisationExpenditureReaderProps = initialState, action: ActionProps): OrganisationExpenditureReaderProps => {
  //console.log('Boom!', action.type, action.payload)
  const payload = action.payload as PayloadProps
  if ( typeof payload != 'undefined' ) {
    const payloadData = payload.data as OrganisationExpenditureReaderProps
    if ( (action.type == OrganisationExpenditureReaderActionTypes.NUM_SUCCESS ) ||
         (action.type == OrganisationExpenditureReaderActionTypes.NUM_SUCCESS ) ) {

      const data: OrganisationExpenditureReaderProps = {
        num: payloadData.num,
        data: { ...state.data }
      }
      return {...data}

    } else if ( (action.type == OrganisationExpenditureReaderActionTypes.REF_SUCCESS ) ||
         (action.type == OrganisationExpenditureReaderActionTypes.REF_FAILURE ) ) {

       const data: OrganisationExpenditureReaderProps  = {
         num: state.num,
         data: payloadData.data
       }

       return data

    } else if ( (action.type == OrganisationExpenditureReaderActionTypes.NUMEXPENDITURE_SUCCESS ) ||
                (action.type == OrganisationExpenditureReaderActionTypes.NUMEXPENDITURE_FAILURE ) ||
                (action.type == OrganisationExpenditureReaderActionTypes.EXPENDITURE_SUCCESS ) ||
                (action.type == OrganisationExpenditureReaderActionTypes.EXPENDITURE_FAILURE ) ) {

       const data: OrganisationExpenditureReaderProps  = {
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
