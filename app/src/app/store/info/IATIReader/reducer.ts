import { Reducer } from 'redux'
import { InfoProps } from '../types'
import { IATIReader } from '../../../utils/strings'

const initialState: InfoProps = {
  title: IATIReader.heading,
  data: IATIReader.info
}

export const reducer: Reducer<InfoProps> = (state = initialState): InfoProps => {
  return state
}
