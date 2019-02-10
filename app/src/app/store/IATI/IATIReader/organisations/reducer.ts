import { IATIReportActionTypes, IATIOrganisationsReportProps } from './types'
import { ActionProps, PayloadProps } from '../../../types'

/*
ERROR in [at-loader] ./src/app/store/IATI/IATIReader/organisations/reducer.ts:5:3
    TS2322: Type '{ '': { IATIOganisations: {}; '': { IATIOganisation: {}; totalBudget: { '': {}; }; recipientOrgBudget: { '': {}; }; recipientRegionBudget: { '': {}; }; recipientCountryBudget: { '': {}; }; totalExpenditure: { '': {}; }; document: { ...; }; }; }; }' is not assignable to type 'IATIOrganisationsReport'.
  Property 'data' is missing in type '{ '': { IATIOganisations: {}; '': { IATIOganisation: {}; totalBudget: { '': {}; }; recipientOrgBudget: { '': {}; }; recipientRegionBudget: { '': {}; }; recipientCountryBudget: { '': {}; }; totalExpenditure: { '': {}; }; document: { ...; }; }; }; }'.

*/
const initialState: IATIOrganisationsReportProps = {
  data: {
    '': {
      IATIOganisations: {},
      '': {
        IATIOganisation: {},
        totalBudget: {
          '': {

          }
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
  if ( (action.type == IATIReportActionTypes.ORGANISATIONS_SUCCESS ) ||
       (action.type == IATIReportActionTypes.ORGANISATION_SUCCESS ) ||
       (action.type == IATIReportActionTypes.BUDGET_SUCCESS ) ||
       (action.type == IATIReportActionTypes.RECIPIENTORGBUDGET_SUCCESS ) ||
       (action.type == IATIReportActionTypes.RECIPIENTREGIONBUDGET_SUCCESS ) ||
       (action.type == IATIReportActionTypes.RECIPIENTCOUNTRYBUDGET_SUCCESS ) ||
       (action.type == IATIReportActionTypes.TOTALEXPENDITURE_SUCCESS ) ||
       (action.type == IATIReportActionTypes.DOCUMENT_SUCCESS ) ) {
    const data = (action.payload.data as IATIOrganisationsReportProps)
    return {...state, ...data}
  } else {
    return state
  }
}
