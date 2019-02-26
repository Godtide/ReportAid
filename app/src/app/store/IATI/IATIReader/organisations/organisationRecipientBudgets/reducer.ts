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
              budgetType: 1,
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

export const reducer = (state: IATIOrganisationRecipientBudgetReportProps = initialState, action: ActionProps): IATIOrganisationRecipientBudgetReportProps => {

  switch (action.type) {
    case IATIReportActionTypes.RECIPIENTORGBUDGET_INIT: {
      const data = (action.payload.data as IATIOrganisationRecipientBudgetReportProps)
      return data
    }
    case IATIReportActionTypes.RECIPIENTORGBUDGET_SUCCESS: {
      const data = (action.payload.data as IATIOrganisationRecipientBudgetReportProps)
      return {...state, ...data}
    }
    default:
      return state
  }
}
