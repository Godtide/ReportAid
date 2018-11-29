import { ChainInfoActionTypes, InfoProps } from './types'
import { ActionProps } from '../../types'

const initialInfoState: InfoProps = {
  data: {
    API: '',
    Name: '',
    ChainId: '',
    ENS: '',
    provider: {},
  }
}

export const infoReducer = (state: InfoProps = initialInfoState, action: ActionProps): InfoProps => {
  console.log('blockchain info: ', action.type, action.payload)
  if ( action.type == ChainInfoActionTypes.ADD_INFO ) {
    //console.log('Orgstate: ', state)
    return (<any>Object).assign({}, state, action.payload)
  } else {
    return state
  }
}
