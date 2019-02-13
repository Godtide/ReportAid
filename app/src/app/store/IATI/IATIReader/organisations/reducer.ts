import { IATIReportActionTypes, IATIOrganisationsReportProps } from './types'
import { ActionProps, PayloadProps } from '../../../types'

const initialState: IATIOrganisationsReportProps = {
  data: {
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
}

export const reducer = (state: IATIOrganisationsReportProps = initialState, action: ActionProps): IATIOrganisationsReportProps => {

  switch (action.type) {
    case IATIReportActionTypes.ORGANISATIONS_SUCCESS: {
      const data = (action.payload as IATIOrganisationsReportProps)
      console.log("Data: ", data)
    }
    case IATIReportActionTypes.ORGANISATION_SUCCESS: {
    }
    case IATIReportActionTypes.BUDGET_SUCCESS: {
    }
    case IATIReportActionTypes.RECIPIENTORGBUDGET_SUCCESS: {
    }
    case IATIReportActionTypes.RECIPIENTREGIONBUDGET_SUCCESS: {
    }
    case IATIReportActionTypes.RECIPIENTCOUNTRYBUDGET_SUCCESS: {
    }
    case IATIReportActionTypes.TOTALEXPENDITURE_SUCCESS: {
    }
    case IATIReportActionTypes.DOCUMENT_SUCCESS: {
    }
    default:
      return state
  }
}
