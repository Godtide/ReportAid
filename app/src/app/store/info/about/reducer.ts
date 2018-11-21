import { Reducer } from 'redux'
import { InfoProps } from '../types'
import { About } from '../../../utils/strings'

const initialState: InfoProps = {
  title: About.heading,
  data: About.info
}

export const reducer: Reducer<InfoProps> = (state = initialState): InfoProps => {
  return state
}
