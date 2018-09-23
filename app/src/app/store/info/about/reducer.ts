import { Reducer } from 'redux'
import { AboutActionTypes, InfoProps } from '../types'
import { AboutStrings } from '../../../utils/strings'

const initialState: InfoProps = {
  title: AboutStrings.heading,
  data: AboutStrings.info
}

export const reducer: Reducer<InfoProps> = (state = initialState, action): InfoProps => {
  switch (action.type) {
    case AboutActionTypes.REQ_DATA:
      return (<any>Object).assign({}, state, {
        title: action.payload.title,
        data: action.payload.data
      })
    default:
      return state
  }
}
