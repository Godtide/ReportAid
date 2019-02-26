import { IATIOrganisationDocReportProps } from '../types'
import { IATIReportActionTypes } from '../types'
import { ActionProps } from '../../../../types'

const initialState: IATIOrganisationDocReportProps = {
  data: {
    organisationsRef: '',
    organisationRef: '',
    data: []
  }
}

export const reducer = (state: IATIOrganisationDocReportProps = initialState, action: ActionProps): IATIOrganisationDocReportProps => {

  switch (action.type) {
    case IATIReportActionTypes.DOCUMENT_INIT: {
      const data = (action.payload.data as IATIOrganisationDocReportProps)
      return data
    }
    case IATIReportActionTypes.DOCUMENT_SUCCESS: {
      const data = (action.payload.data as IATIOrganisationDocReportProps)
      return {...state, ...data}
    }
    default:
      return state
  }
}
