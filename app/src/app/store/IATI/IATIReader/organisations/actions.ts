import { storeAction } from '../../../actions'

import { PayloadProps } from '../../../types'
import { IATIReportActionTypes } from '../types'

export const read = (payload: PayloadProps): Function => {
  return (actionType: IATIReportActionTypes): PayloadProps => {
    const getProps = storeAction(actionType)(payload) as PayloadProps
    return getProps
  }
}
