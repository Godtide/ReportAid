import { OrgDocsReaderActionTypes, OrgDocsReaderProps } from './types'
import { ActionProps, PayloadProps } from '../../../types'

const initialState: OrgDocsReaderProps = {
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

export const reducer = (state: OrgDocsReaderProps = initialState, action: ActionProps): OrgDocsReaderProps => {
  //console.log('Boom!', action.type, action.payload)
  const payload = action.payload as PayloadProps
  if ( typeof payload != 'undefined' ) {
    const payloadData = payload.data as OrgDocsReaderProps
    if ( (action.type == OrgDocsReaderActionTypes.NUM_SUCCESS ) ||
         (action.type == OrgDocsReaderActionTypes.NUM_SUCCESS ) ) {

      const data: OrgDocsReaderProps = {
        num: payloadData.num,
        data: { ...state.data }
      }
      return {...data}

    } else if ( (action.type == OrgDocsReaderActionTypes.REF_SUCCESS ) ||
         (action.type == OrgDocsReaderActionTypes.REF_FAILURE ) ) {

       const data: OrgDocsReaderProps  = {
         num: state.num,
         data: payloadData.data
       }

       return data

    } else if ( (action.type == OrgDocsReaderActionTypes.NUMDOC_SUCCESS ) ||
                (action.type == OrgDocsReaderActionTypes.NUMDOC_FAILURE ) ||
                (action.type == OrgDocsReaderActionTypes.DOC_SUCCESS ) ||
                (action.type == OrgDocsReaderActionTypes.DOC_FAILURE ) ) {

       const data: OrgDocsReaderProps  = {
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
