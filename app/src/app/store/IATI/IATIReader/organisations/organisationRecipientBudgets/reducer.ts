import { IATIOrganisationRecipientBudgetReportProps } from './types'
import { IATIReportActionTypes } from '../types'
import { ActionProps } from '../../../../types'

const initialState: IATIOrganisationRecipientBudgetReportProps = {
  data: {
    '': {
      data: {
        '': {
          data: {
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
          }
        }
      }
    }
  }
}

export const reducer = (state: IATIOrganisationRecipientBudgetReportProps = initialState, action: ActionProps): IATIOrganisationRecipientBudgetReportProps => {

  switch (action.type) {
    case IATIReportActionTypes.RECIPIENTCOUNTRYBUDGET_SUCCESS: {
      const data = (action.payload.data as IATIOrganisationRecipientBudgetReportProps)
      return {...state, ...data}
    }
    default:
      return state
  }
}
