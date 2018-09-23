import { Reducer } from 'redux'
import { HelpActionTypes, InfoProps } from '../types'
import { HelpStrings } from '../../../utils/strings'

const initialState: InfoProps = {
  title: HelpStrings.heading,
  data: HelpStrings.info
}

export const reducer: Reducer<InfoProps> = (state = initialState, action): InfoProps => {
  switch (action.type) {
    case HelpActionTypes.REQ_DATA:
      return (<any>Object).assign({}, state, {
        title: action.payload.title,
        data: action.payload.data
      })
    default:
      return state
  }
}
