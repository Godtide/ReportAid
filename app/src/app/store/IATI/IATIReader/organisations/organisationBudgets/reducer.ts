import { IATIOrganisationBudgetReportProps } from './types'
import { IATIReportActionTypes } from '../types'
import { ActionProps } from '../../../../types'

const initialState: IATIOrganisationBudgetReportProps = {
  data: {
    '': {
      data: {
        '': {
          data: {
            '': {
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

export const reducer = (state: IATIOrganisationBudgetReportProps = initialState, action: ActionProps): IATIOrganisationBudgetReportProps => {

  switch (action.type) {
    case IATIReportActionTypes.BUDGET_SUCCESS: {
      const data = (action.payload.data as IATIOrganisationBudgetReportProps)
      return {...state, ...data}
    }
    default:
      return state
  }
}
