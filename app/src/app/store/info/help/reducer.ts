import { Reducer } from 'redux'
import { InfoProps } from '../types'
import { Help } from '../../../utils/strings'

const initialState: InfoProps = {
  title: Help.heading,
  data: Help.info
}

export const reducer: Reducer<InfoProps> = (state = initialState): InfoProps => {
  return state
}
