import { storeAction } from '../../actions'
import { InfoProps } from './types'
import { ChainInfoActionTypes } from './types'

export const addInfo = (payload: InfoProps) => storeAction(ChainInfoActionTypes.ADD_INFO)(payload)
