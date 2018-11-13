import { Reducer } from 'redux'
import { InfoProps } from '../types'
import { HelpStrings } from '../../../utils/strings'

const initialState: InfoProps = {
  title: HelpStrings.heading,
  data: HelpStrings.info
}

export const reducer: Reducer<InfoProps> = (state = initialState): InfoProps => {
  return state
}
