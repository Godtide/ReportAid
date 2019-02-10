import { IATIWriterActionTypes } from './types'
import { ActionProps, TxProps, TxData } from '../../types'

const initialState: TxProps = {
  data: {
    '': {}
  }
}

export const reducer = (state: TxProps = initialState, action: ActionProps): TxProps => {

  const payload = action.payload as PayloadProps
  if ( (action.type == IATIWriterActionTypes.ORGANISATIONS_SUCCESS ) ||
       (action.type == IATIWriterActionTypes.ORGANISATION_SUCCESS ) ||
       (action.type == IATIWriterActionTypes.BUDGET_SUCCESS ) ||
       (action.type == IATIWriterActionTypes.RECIPIENTORGBUDGET_SUCCESS ) ||
       (action.type == IATIWriterActionTypes.RECIPIENTREGIONBUDGET_SUCCESS ) ||
       (action.type == IATIWriterActionTypes.RECIPIENTCOUNTRYBUDGET_SUCCESS ) ||
       (action.type == IATIWriterActionTypes.TOTALEXPENDITURE_SUCCESS ) ||
       (action.type == IATIWriterActionTypes.DOCUMENT_SUCCESS ) ) {

   const data = (action.payload.data as TxData)
   return {...state, ...data}

  } else {
    return state
  }
}
