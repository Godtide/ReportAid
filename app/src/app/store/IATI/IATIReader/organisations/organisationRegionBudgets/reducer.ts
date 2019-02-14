import { IATIOrganisationRegionBudgetReportProps } from './types'
import { IATIReportActionTypes } from '../types'
import { ActionProps } from '../../../../types'

const initialState: IATIOrganisationRegionBudgetReportProps = {
  data: {
    '': {
      data: {
        '': {
          data: {
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
          }
        }
      }
    }
  }
}

export const reducer = (state: IATIOrganisationRegionBudgetReportProps = initialState, action: ActionProps): IATIOrganisationRegionBudgetReportProps => {

  switch (action.type) {
    case IATIReportActionTypes.RECIPIENTREGIONBUDGET_SUCCESS: {
      const data = (action.payload.data as IATIOrganisationRegionBudgetReportProps)
      const newState = {...state, ...data}
      //console.log('New state: ', newState)
      return newState
    }
    default:
      return state
  }
}
