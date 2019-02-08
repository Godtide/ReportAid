import { storeAction } from '../actions'

import { ActionProps, PayloadProps } from '../../types'
import { OrganisationsReaderActionTypes } from './types'

const write = (payload: PayloadProps): Function => {
  return (actionType: IATIWriterActionTypes): TxProps => {
    const writerProps = storeAction(actionType)(payload) as TxProps
    return writerProps
  }
}
