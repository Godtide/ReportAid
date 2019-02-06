import { OrganisationsReaderActionTypes, OrganisationsReaderProps } from './types'
import { ActionProps, PayloadProps } from '../../../types'

const initialState: OrganisationsReaderProps = {
  num: 0,
  data: {
    '': {
      num: 0,
      data: {
        '': {
          version: '',
          report: {
            reportRef: '',
            orgRef: ''
          },
          reportingOrg: {
            orgRef: '',
            orgType: 0,
            isSecondary: false
          },
          lang: '',
          currency: '',
          lastUpdatedTime: ''
        }
      }
    }
  }
}

export const reducer = (state: OrganisationsReaderProps = initialState, action: ActionProps): OrganisationsReaderProps => {
  //console.log('Boom!', action.type, action.payload)
  const payload = action.payload as PayloadProps
  if ( typeof payload != 'undefined' ) {
    const payloadData = payload.data as OrganisationsReaderProps
    if ( (action.type == OrganisationsReaderActionTypes.NUM_SUCCESS ) ||
         (action.type == OrganisationsReaderActionTypes.NUM_SUCCESS ) ) {

      const data: OrganisationsReaderProps = {
        num: payloadData.num,
        data: { ...state.data }
      }
      return {...data}

    } else if ( (action.type == OrganisationsReaderActionTypes.REF_SUCCESS ) ||
         (action.type == OrganisationsReaderActionTypes.REF_FAILURE ) ) {

       const data: OrganisationsReaderProps  = {
         num: state.num,
         data: payloadData.data
       }

       return data

    } else if ( (action.type == OrganisationsReaderActionTypes.NUMREP_SUCCESS ) ||
                (action.type == OrganisationsReaderActionTypes.NUMREP_FAILURE ) ||
                (action.type == OrganisationsReaderActionTypes.REPORT_SUCCESS ) ||
                (action.type == OrganisationsReaderActionTypes.REPORT_FAILURE ) ) {

       const data: OrganisationsReaderProps  = {
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
