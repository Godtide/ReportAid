import { IATIOrganisationCountryBudgetReportProps } from './types'
import { IATIReportActionTypes } from '../types'
import { ActionProps } from '../../../../types'

const initialState: IATIOrganisationCountryBudgetReportProps = {
  data: {
    '': {
      data: {
        '': {
          data: {
            '': {
              budgetLine: '',
              otherRef: '',
              finance: {
                value: 0,
                status: 0,
                start: '',
                end: ''
              }
            }
          }
        }
      }
    }
  }
}

export const reducer = (state: IATIOrganisationCountryBudgetReportProps = initialState, action: ActionProps): IATIOrganisationCountryBudgetReportProps => {

  switch (action.type) {
    case IATIReportActionTypes.RECIPIENTCOUNTRYBUDGET_INIT: {
      const data = (action.payload.data as IATIOrganisationCountryBudgetReportProps)
      return data
    }
    case IATIReportActionTypes.RECIPIENTCOUNTRYBUDGET_SUCCESS: {
      const data = (action.payload.data as IATIOrganisationCountryBudgetReportProps)
      return {...state, ...data}
    }
    default:
      return state
  }
}
