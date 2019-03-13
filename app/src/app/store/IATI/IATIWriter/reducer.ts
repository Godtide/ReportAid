import { IATIWriterActionTypes } from '../types'
import { ActionProps, PayloadProps } from '../../types'

const initialState: PayloadProps = {
  data: {}
}

export const reducer = (state: PayloadProps = initialState, action: ActionProps): PayloadProps => {

  switch (action.type) {
    case IATIWriterActionTypes.TX_INIT: {
      const data = action.payload.data as PayloadProps
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
    case IATIWriterActionTypes.DOCUMENT_SUCCESS:
    case IATIWriterActionTypes.ACTIVITIES_SUCCESS:
    case IATIWriterActionTypes.ACTIVITY_SUCCESS:
    case IATIWriterActionTypes.ACTIVITYADDITIONAL_SUCCESS:
    case IATIWriterActionTypes.ACTIVITYDATE_SUCCESS:
    case IATIWriterActionTypes.ACTIVITYPARTICIPATINGORG_SUCCESS:
    case IATIWriterActionTypes.ORGS_FAILURE:
    case IATIWriterActionTypes.ORGANISATIONS_FAILURE:
    case IATIWriterActionTypes.ORGANISATION_FAILURE:
    case IATIWriterActionTypes.BUDGET_FAILURE:
    case IATIWriterActionTypes.RECIPIENTORGBUDGET_FAILURE:
    case IATIWriterActionTypes.RECIPIENTREGIONBUDGET_FAILURE:
    case IATIWriterActionTypes.RECIPIENTCOUNTRYBUDGET_FAILURE:
    case IATIWriterActionTypes.TOTALEXPENDITURE_FAILURE:
    case IATIWriterActionTypes.DOCUMENT_FAILURE:
    case IATIWriterActionTypes.ACTIVITIES_FAILURE:
    case IATIWriterActionTypes.ACTIVITY_FAILURE:
    case IATIWriterActionTypes.ACTIVITYADDITIONAL_FAILURE:
    case IATIWriterActionTypes.ACTIVITYADDITIONAL_FAILURE:
    case IATIWriterActionTypes.ACTIVITYPARTICIPATINGORG_FAILURE: {
      const data = (action.payload.data as PayloadProps)
      return {...state, ...data}
    }
    default:
      return state
  }
}
