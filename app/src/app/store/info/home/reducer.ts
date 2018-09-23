import { Reducer } from 'redux'
import { HomeActionTypes, InfoProps } from '../types'
import { HomeStrings } from '../../../utils/strings'

const initialState: InfoProps = {
  title: HomeStrings.heading,
  data: HomeStrings.info
}

export const reducer: Reducer<InfoProps> = (state = initialState, action): InfoProps => {
  switch (action.type) {
    case HomeActionTypes.REQ_DATA:
      return (<any>Object).assign({}, state, {
        title: action.payload.title,
        data: action.payload.data
      })
    default:
      return state
  }
}
