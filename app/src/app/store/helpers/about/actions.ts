import { AboutActionTypes } from './types'
//import { Promise } from 'bluebird'
import { AboutStrings } from '../../../utils/strings'

export default function getData() {
  return (dispatch: any) => {
    dispatch({
      type: AboutActionTypes.REQ_DATA,
      data: { text: AboutStrings.info, title: AboutStrings.heading }
    })
  }
}
