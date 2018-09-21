import { Reducer } from 'redux'
import { AboutActionTypes, AboutProps } from './types'

const initialState: AboutProps = {
  title: '',
  data: '',
}

function resData(state: AboutProps, action: any) {
  const { data } = action
  return data
}

export const AboutReducer: Reducer<AboutProps> = (state = initialState, action) => {
  switch (action.type) {
    case AboutActionTypes.RES_DATA: return resData(state, action)
    default: return state
  }
}
