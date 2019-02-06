import { OrgReportDocsReaderActionTypes, OrgReportDocsReaderProps } from './types'
import { ActionProps, PayloadProps } from '../../../types'

const initialState: OrgReportDocsReaderProps = {
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

export const reducer = (state: OrgReportDocsReaderProps = initialState, action: ActionProps): OrgReportDocsReaderProps => {
  //console.log('Boom!', action.type, action.payload)
  const payload = action.payload as PayloadProps
  if ( typeof payload != 'undefined' ) {
    const payloadData = payload.data as OrgReportDocsReaderProps
    if ( (action.type == OrgReportDocsReaderActionTypes.NUM_SUCCESS ) ||
         (action.type == OrgReportDocsReaderActionTypes.NUM_SUCCESS ) ) {

      const data: OrgReportDocsReaderProps = {
        num: payloadData.num,
        data: { ...state.data }
      }
      return {...data}

    } else if ( (action.type == OrgReportDocsReaderActionTypes.REF_SUCCESS ) ||
         (action.type == OrgReportDocsReaderActionTypes.REF_FAILURE ) ) {

       const data: OrgReportDocsReaderProps  = {
         num: state.num,
         data: payloadData.data
       }

       return data

    } else if ( (action.type == OrgReportDocsReaderActionTypes.NUMDOC_SUCCESS ) ||
                (action.type == OrgReportDocsReaderActionTypes.NUMDOC_FAILURE ) ||
                (action.type == OrgReportDocsReaderActionTypes.DOC_SUCCESS ) ||
                (action.type == OrgReportDocsReaderActionTypes.DOC_FAILURE ) ) {

       const data: OrgReportDocsReaderProps  = {
         num: state.num,
         data: {...payloadData.data}
       }

       //console.log ('Report data: ', data)
       return data

    } else {
      return state
    }
  } else {
    return state
  }
}
