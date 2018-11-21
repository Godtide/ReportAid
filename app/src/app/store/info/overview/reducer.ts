import { Reducer } from 'redux'
import { InfoProps } from '../types'
import { Overview } from '../../../utils/strings'

const initialState: InfoProps = {
  title: Overview.heading,
  data: Overview.info
}

export const reducer: Reducer<InfoProps> = (state = initialState): InfoProps => {
  return state
}
