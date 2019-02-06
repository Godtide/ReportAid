import { OrganisationDocsReaderActionTypes, OrganisationDocsReaderProps } from './types'
import { ActionProps, PayloadProps } from '../../../types'

const initialState: OrganisationDocsReaderProps = {
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
          docRef: '',
          title: '',
          format: '',
          url: '',
          category: '',
          countryRef: '',
          desc: '',
          lang: '',
          date: ''
        }
      }
    }
  }
}

export const reducer = (state: OrganisationDocsReaderProps = initialState, action: ActionProps): OrganisationDocsReaderProps => {
  //console.log('Boom!', action.type, action.payload)
  const payload = action.payload as PayloadProps
  if ( typeof payload != 'undefined' ) {
    const payloadData = payload.data as OrganisationDocsReaderProps
    if ( (action.type == OrganisationDocsReaderActionTypes.NUM_SUCCESS ) ||
         (action.type == OrganisationDocsReaderActionTypes.NUM_SUCCESS ) ) {

      const data: OrganisationDocsReaderProps = {
        num: payloadData.num,
        data: { ...state.data }
      }
      return {...data}

    } else if ( (action.type == OrganisationDocsReaderActionTypes.REF_SUCCESS ) ||
         (action.type == OrganisationDocsReaderActionTypes.REF_FAILURE ) ) {

       const data: OrganisationDocsReaderProps  = {
         num: state.num,
         data: payloadData.data
       }

       return data

    } else if ( (action.type == OrganisationDocsReaderActionTypes.NUMDOC_SUCCESS ) ||
                (action.type == OrganisationDocsReaderActionTypes.NUMDOC_FAILURE ) ||
                (action.type == OrganisationDocsReaderActionTypes.DOC_SUCCESS ) ||
                (action.type == OrganisationDocsReaderActionTypes.DOC_FAILURE ) ) {

       const data: OrganisationDocsReaderProps  = {
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
