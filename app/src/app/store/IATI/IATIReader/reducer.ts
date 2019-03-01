import { IATIReportActionTypes } from './types'
import { ActionProps, PayloadProps } from '../../types'

const initialState: PayloadProps = {
  data: {}
}

export const reducer = (state: PayloadProps = initialState, action: ActionProps): PayloadProps => {

  switch (action.type) {
    case IATIReportActionTypes.REPORT_INIT: {
      const data = action.payload.data as PayloadProps
      return data
    }
    case IATIReportActionTypes.ORGS_SUCCESS:
    case IATIReportActionTypes.ORGANISATIONS_SUCCESS:
    case IATIReportActionTypes.ORGANISATION_SUCCESS:
    case IATIReportActionTypes.BUDGET_SUCCESS:
    case IATIReportActionTypes.RECIPIENTORGBUDGET_SUCCESS:
    case IATIReportActionTypes.RECIPIENTREGIONBUDGET_SUCCESS:
    case IATIReportActionTypes.RECIPIENTCOUNTRYBUDGET_SUCCESS:
    case IATIReportActionTypes.TOTALEXPENDITURE_SUCCESS:
    case IATIReportActionTypes.DOCUMENT_SUCCESS:
    case IATIReportActionTypes.ACTIVITIES_SUCCESS:
    case IATIReportActionTypes.ORGS_FAILURE:
    case IATIReportActionTypes.ORGANISATIONS_FAILURE:
    case IATIReportActionTypes.ORGANISATION_FAILURE:
    case IATIReportActionTypes.BUDGET_FAILURE:
    case IATIReportActionTypes.RECIPIENTORGBUDGET_FAILURE:
    case IATIReportActionTypes.RECIPIENTREGIONBUDGET_FAILURE:
    case IATIReportActionTypes.RECIPIENTCOUNTRYBUDGET_FAILURE:
    case IATIReportActionTypes.TOTALEXPENDITURE_FAILURE:
    case IATIReportActionTypes.DOCUMENT_FAILURE:
    case IATIReportActionTypes.ACTIVITIES_FAILURE: {
      const data = (action.payload.data as PayloadProps)
      return {...state, ...data}
    }
    default:
      return state
  }
}
