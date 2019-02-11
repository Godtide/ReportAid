import { storeAction } from '../../../actions'

import { ActionProps, PayloadProps, TxProps } from '../../../types'
import { IATIWriterActionTypes } from './types'

export const write = (payload: PayloadProps): Function => {
  return (actionType: IATIWriterActionTypes): TxProps => {
    const writerProps = storeAction(actionType)(payload) as TxProps
    return writerProps
  }
}
