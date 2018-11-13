import { Reducer } from 'redux'
import { InfoProps } from '../types'
import { OverviewStrings } from '../../../utils/strings'

const initialState: InfoProps = {
  title: OverviewStrings.heading,
  data: OverviewStrings.info
}

export const reducer: Reducer<InfoProps> = (state = initialState): InfoProps => {
  return state
}
