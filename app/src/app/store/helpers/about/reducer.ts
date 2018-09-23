import { Reducer } from 'redux'
import { ActionTypes, AboutProps } from './types'
import { AboutStrings } from '../../../utils/strings'

const initialState: AboutProps = {
  title: AboutStrings.heading,
  data: AboutStrings.info
}

export const reducer: Reducer<AboutProps> = (state = initialState, action): AboutProps => {
  switch (action.type) {
    case ActionTypes.REQ_DATA:
      return Object.assign({}, state, {
        title: action.payload.title,
        data: action.payload.data
      })
    default:
      return state
  }
}
