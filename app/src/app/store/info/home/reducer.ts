import { Reducer } from 'redux'
import { InfoProps } from '../types'
import { HomeStrings } from '../../../utils/strings'

const initialState: InfoProps = {
  title: HomeStrings.heading,
  data: HomeStrings.info
}

export const reducer: Reducer<InfoProps> = (state = initialState): InfoProps => {
  return state
}
