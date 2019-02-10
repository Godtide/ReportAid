import { IATIReportActionTypes, IATIOrganisationsReportProps } from './types'
import { ActionProps, PayloadProps } from '../../types'

const initialState: IATIReportProps = {
  data: {
    '': {
      IATIOganisations: {},
      '': {
        IATIOganisation: {},
        totalBudget: {
          '': {}
        },
        recipientOrgBudget: {
          '': {}
        },
        recipientRegionBudget: {
          '': {}
        },
        recipientCountryBudget: {
          '': {}
        },
        totalExpenditure: {
          '': {}
        },
        document: {
          '': {}
        }
      }
    }
  }
}

export const reducer = (state: IATIReportProps = initialState, action: ActionProps): IATIReportProps => {

  const payload = action.payload as PayloadProps
  if ( (action.type == IATIReportActionTypes.ORGANISATIONS_SUCCESS ) ||
       (action.type == IATIReportActionTypes.ORGANISATION_SUCCESS ) ||
       (action.type == IATIReportActionTypes.BUDGET_SUCCESS ) ||
       (action.type == IATIReportActionTypes.RECIPIENTORGBUDGET_SUCCESS ) ||
       (action.type == IATIReportActionTypes.RECIPIENTREGIONBUDGET_SUCCESS ) ||
       (action.type == IATIReportActionTypes.RECIPIENTCOUNTRYBUDGET_SUCCESS ) ||
       (action.type == IATIReportActionTypes.TOTALEXPENDITURE_SUCCESS ) ||
       (action.type == IATIReportActionTypes.DOCUMENT_SUCCESS ) ) {
    const data = (action.payload.data as IATIReportProps)
    return {...state, ...data}
  } else {
    return state
  }
}
