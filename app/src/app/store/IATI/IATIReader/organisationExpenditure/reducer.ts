import { OrgExpenditureReaderActionTypes, OrgExpenditureReaderProps } from './types'
import { ActionProps, PayloadProps } from '../../../types'

const initialState: OrgExpenditureReaderProps = {
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

export const reducer = (state: OrgExpenditureReaderProps = initialState, action: ActionProps): OrgExpenditureReaderProps => {
  //console.log('Boom!', action.type, action.payload)
  const payload = action.payload as PayloadProps
  if ( typeof payload != 'undefined' ) {
    const payloadData = payload.data as OrgExpenditureReaderProps
    if ( (action.type == OrgExpenditureReaderActionTypes.NUM_SUCCESS ) ||
         (action.type == OrgExpenditureReaderActionTypes.NUM_SUCCESS ) ) {

      const data: OrgExpenditureReaderProps = {
        num: payloadData.num,
        data: { ...state.data }
      }
      return {...data}

    } else if ( (action.type == OrgExpenditureReaderActionTypes.REF_SUCCESS ) ||
         (action.type == OrgExpenditureReaderActionTypes.REF_FAILURE ) ) {

       const data: OrgExpenditureReaderProps  = {
         num: state.num,
         data: payloadData.data
       }

       return data

    } else if ( (action.type == OrgExpenditureReaderActionTypes.NUMEXPENDITURE_SUCCESS ) ||
                (action.type == OrgExpenditureReaderActionTypes.NUMEXPENDITURE_FAILURE ) ||
                (action.type == OrgExpenditureReaderActionTypes.EXPENDITURE_SUCCESS ) ||
                (action.type == OrgExpenditureReaderActionTypes.EXPENDITURE_FAILURE ) ) {

       const data: OrgExpenditureReaderProps  = {
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
