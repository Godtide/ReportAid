import { IATIOrganisationExpenditureReportProps } from './types'
import { IATIReportActionTypes } from '../types'
import { ActionProps } from '../../../../types'

const initialState: IATIOrganisationExpenditureReportProps = {
  data: {
    '': {
      data: {
        '': {
          data: {
            '': {
              expenditureLine: '',
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

export const reducer = (state: IATIOrganisationExpenditureReportProps = initialState, action: ActionProps): IATIOrganisationExpenditureReportProps => {

  switch (action.type) {
    case IATIReportActionTypes.TOTALEXPENDITURE_SUCCESS: {
      const data = (action.payload.data as IATIOrganisationExpenditureReportProps)
      return {...state, ...data}
    }
    default:
      return state
  }
}
