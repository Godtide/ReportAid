import { Reducer } from 'redux'
import { InfoProps } from '../types'
import { IATIWriter } from '../../../utils/strings'

const initialState: InfoProps = {
  title: IATIWriter.heading,
  data: IATIWriter.info
}

export const reducer: Reducer<InfoProps> = (state = initialState): InfoProps => {
  return state
}
