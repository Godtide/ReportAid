import { IATIReportActionTypes, IATIOrganisationsReportProps } from './types'
import { ActionProps, PayloadProps } from '../../../types'

const initialState: IATIOrganisationsReportProps = {
  '': {
    IATIOrganisations: {
      version: '',
      generatedTime: ''
    },
    data: {
      '': {
        IATIOrganisation: {
          orgRef: '',
          reportingOrg: {
            orgRef: '',
            orgType: 0,
            isSecondary: false
          },
          lang: '',
          currency: '',
          lastUpdatedTime: ''
        },
        data: {
          totalBudget: {
            '': {
              budgetLine: '',
              finance: {
                value: 0,
                status: 0,
                start: '',
                end: ''
              }
            }
          },
          recipientOrgBudget: {
            '': {
              recipientOrgRef: '',
              budgetLine: '',
              finance: {
                value: 0,
                status: 0,
                start: '',
                end: ''
              }
            }
          },
          recipientRegionBudget: {
            '': {
              regionRef: 0,
              budgetLine: '',
              finance: {
                value: 0,
                status: 0,
                start: '',
                end: ''
              }
            }
          },
          recipientCountryBudget: {
            '': {
              countryRef: '',
              budgetLine: '',
              finance: {
                value: 0,
                status: 0,
                start: '',
                end: ''
              }
            }
          },
          totalExpenditure: {
            '': {
              expenditureLine: '',
              finance: {
                value: 0,
                status: 0,
                start: '',
                end: ''
              }
            }
          },
          document: {
            '': {
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
