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
              countryRef: '',
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

export const reducer = (state: IATIOrganisationCountryBudgetReportProps = initialState, action: ActionProps): IATIOrganisationCountryBudgetReportProps => {

  switch (action.type) {
    case IATIReportActionTypes.RECIPIENTCOUNTRYBUDGET_SUCCESS: {
      const data = (action.payload.data as IATIOrganisationCountryBudgetReportProps)
      const newState = {...state, ...data}
      //console.log('New state: ', newState)
      return newState
    }
    default:
      return state
  }
}
