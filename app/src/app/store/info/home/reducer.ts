import { Reducer } from 'redux'
import { InfoProps } from '../types'
import { Home } from '../../../utils/strings'

const initialState: InfoProps = {
  title: Home.heading,
  data: Home.info
}

export const reducer: Reducer<InfoProps> = (state = initialState): InfoProps => {
  return state
}
