import { Reducer } from 'redux'
import { InfoProps } from '../types'
import { AboutStrings } from '../../../utils/strings'

const initialState: InfoProps = {
  title: AboutStrings.heading,
  data: AboutStrings.info
}

export const reducer: Reducer<InfoProps> = (state = initialState): InfoProps => {
  return state
}
