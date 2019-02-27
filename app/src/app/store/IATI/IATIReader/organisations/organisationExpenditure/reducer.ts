import { IATIBudgetReportProps, IATIReportActionTypes } from '../../types'
import { ActionProps } from '../../../../types'

const initialState: IATIBudgetReportProps = {
  data: {
    organisationsRef: '',
    organisationRef: '',
    data: []
  }
}

export const reducer = (state: IATIBudgetReportProps = initialState, action: ActionProps): IATIBudgetReportProps => {
  switch (action.type) {
    case IATIReportActionTypes.TOTALEXPENDITURE_INIT: {
      const data = (action.payload.data as IATIBudgetReportProps)
      return data
    }
    case IATIReportActionTypes.TOTALEXPENDITURE_SUCCESS: {
      const data = (action.payload.data as IATIBudgetReportProps)
      return {...state, ...data}
    }
    default:
      return state
  }
}
