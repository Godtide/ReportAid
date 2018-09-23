import { Reducer } from 'redux'
import { OverviewActionTypes, InfoProps } from '../types'
import { OverviewStrings } from '../../../utils/strings'

const initialState: InfoProps = {
  title: OverviewStrings.heading,
  data: OverviewStrings.info
}

export const reducer: Reducer<InfoProps> = (state = initialState, action): InfoProps => {
  switch (action.type) {
    case OverviewActionTypes.REQ_DATA:
      return (<any>Object).assign({}, state, {
        title: action.payload.title,
        data: action.payload.data
      })
    default:
      return state
  }
}
