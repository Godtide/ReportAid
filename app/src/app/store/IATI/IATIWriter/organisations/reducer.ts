import { IATIWriterActionTypes } from './types'
import { ActionProps, PayloadProps, TxProps, TxData } from '../../../types'

const initialState: TxProps = {
  data: {
    '': {
      summary: '',
      info: {}
    }
  }
}

export const reducer = (state: TxProps = initialState, action: ActionProps): TxProps => {

  switch (action.type) {
    case IATIWriterActionTypes.TX_INIT: {
      const data = action.payload.data as TxProps
      return data
    }
    case IATIWriterActionTypes.ORGS_SUCCESS:
    case IATIWriterActionTypes.ORGANISATIONS_SUCCESS:
    case IATIWriterActionTypes.ORGANISATION_SUCCESS:
    case IATIWriterActionTypes.BUDGET_SUCCESS:
    case IATIWriterActionTypes.RECIPIENTORGBUDGET_SUCCESS:
    case IATIWriterActionTypes.RECIPIENTREGIONBUDGET_SUCCESS:
    case IATIWriterActionTypes.RECIPIENTCOUNTRYBUDGET_SUCCESS:
    case IATIWriterActionTypes.TOTALEXPENDITURE_SUCCESS:
    case IATIWriterActionTypes.DOCUMENT_SUCCESS: {
      const data = (action.payload.data as TxData)
      return {...state, ...data}
    }
    default:
      return state
  }
}
