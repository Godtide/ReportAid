import { storeAction } from '../actions'

import { ActionProps, PayloadProps } from '../../types'
import { OrganisationsReaderActionTypes } from './types'

export const get = (payload: PayloadProps): Function => {
  return (actionType: OrganisationsReportActionTypes): OrganisationsReportActionTypes => {
    const getProps = storeAction(actionType)(payload) as OrganisationsReportActionTypes
    return getProps
  }
}
