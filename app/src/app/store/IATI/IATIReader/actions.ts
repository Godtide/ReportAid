import { storeAction } from '../actions'

import { ActionProps, PayloadProps } from '../../types'
import { OrganisationsReaderActionTypes } from './types'

export const read = (payload: PayloadProps): Function => {
  return (actionType: IATIReportActionTypes): IATIReportActionTypes => {
    const getProps = storeAction(actionType)(payload) as IATIReportActionTypes
    return getProps
  }
}
