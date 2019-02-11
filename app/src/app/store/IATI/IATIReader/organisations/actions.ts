import { storeAction } from '../../../actions'

import { ActionProps, PayloadProps } from '../../../types'
import { IATIOrganisationsReportProps, IATIReportActionTypes } from './types'

export const read = (payload: PayloadProps): Function => {
  return (actionType: IATIReportActionTypes): IATIOrganisationsReportProps => {
    const getProps = storeAction(actionType)(payload) as IATIOrganisationsReportProps
    return getProps
  }
}
