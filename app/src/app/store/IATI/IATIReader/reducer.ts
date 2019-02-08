import { OrganisationsReportActionTypes, IATIOrganisationsReportProps } from './types'
import { ActionProps, PayloadProps } from '../../types'

const initialState: IATIOrganisationsReportProps = {
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

export const reducer = (state: IATIOrganisationsReportProps = initialState, action: ActionProps): IATIOrganisationsReportProps => {

  const payload = action.payload as PayloadProps
  if ( (action.type == OrganisationsReportActionTypes.ORGANISATIONS_SUCCESS ) ||
       (action.type == OrganisationsReportActionTypes.ORGANISATION_SUCCESS ) ||
       (action.type == OrganisationsReportActionTypes.BUDGET_SUCCESS ) ||
       (action.type == OrganisationsReportActionTypes.RECIPIENTORGBUDGET_SUCCESS ) ||
       (action.type == OrganisationsReportActionTypes.RECIPIENTREGIONBUDGET_SUCCESS ) ||
       (action.type == OrganisationsReportActionTypes.RECIPIENTCOUNTRYBUDGET_SUCCESS ) ||
       (action.type == OrganisationsReportActionTypes.TOTALEXPENDITURE_SUCCESS ) ||
       (action.type == OrganisationsReportActionTypes.DOCUMENT_SUCCESS ) ) {
    const data = (action.payload.data as IATIOrganisationsReportProps)
    return {...state, ...data}
  } else {
    return state
  }
}
