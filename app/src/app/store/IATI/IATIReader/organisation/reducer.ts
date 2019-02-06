import { OrganisationReaderActionTypes, OrganisationReaderProps } from './types'
import { ActionProps, PayloadProps } from '../../../types'

const initialState: OrganisationReaderProps = {
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

export const reducer = (state: OrganisationReaderProps = initialState, action: ActionProps): OrganisationReaderProps => {
  //console.log('Boom!', action.type, action.payload)
  const payload = action.payload as PayloadProps
  if ( typeof payload != 'undefined' ) {
    const payloadData = payload.data as OrganisationReaderProps
    if ( (action.type == OrganisationReaderActionTypes.NUM_SUCCESS ) ||
         (action.type == OrganisationReaderActionTypes.NUM_SUCCESS ) ) {

      const data: OrganisationReaderProps = {
        num: payloadData.num,
        data: { ...state.data }
      }
      return {...data}

    } else if ( (action.type == OrganisationReaderActionTypes.REF_SUCCESS ) ||
         (action.type == OrganisationReaderActionTypes.REF_FAILURE ) ) {

       const data: OrganisationReaderProps  = {
         num: state.num,
         data: payloadData.data
       }

       return data

    } else if ( (action.type == OrganisationReaderActionTypes.NUMREP_SUCCESS ) ||
                (action.type == OrganisationReaderActionTypes.NUMREP_FAILURE ) ||
                (action.type == OrganisationReaderActionTypes.REPORT_SUCCESS ) ||
                (action.type == OrganisationReaderActionTypes.REPORT_FAILURE ) ) {

       const data: OrganisationReaderProps  = {
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
