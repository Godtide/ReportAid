import { Reducer } from 'redux'
import { AboutActionTypes, AboutProps } from './types'

const initialState: AboutProps = {
  title: '',
  data: '',
}

export const AboutReducer: Reducer<AboutProps> = (state = initialState, action): AboutProps => {
  console.log('fek!', action)
  switch (action.type) {
    case AboutActionTypes.REQ_DATA:
      console.log('Gee!', action.payload.title, action.payload.data)
      return Object.assign({}, state, {
        title: action.payload.title,
        data: action.payload.data
      })
    default:
      return state
  }
}
