import { OrgGetActionTypes, OrgGetProps } from './types'
import { ActionProps } from '../../../types'

const initialState: OrgGetProps = {
  data: {
    num: 0,
    refs: [],
    names: [],
    types: []
  }
}

export const reducer = (state: OrgGetProps = initialState, action: ActionProps): OrgGetProps => {
  //console.log('Boom!', action.type, action.payload)
  const propsData = action.payload as OrgGetProps
  if ( (action.type == OrgGetActionTypes.NUM_SUCCESS ) ||
       (action.type == OrgGetActionTypes.NUM_SUCCESS ) ) {

    //console.log('Num State: ', state)

    const data = {
      data: {
        num: propsData.data.num,
        refs: state.data.refs,
        names: state.data.names,
        types: state.data.types
      }
    }
    return (<any>Object).assign({}, data)

  } else if ( (action.type == OrgGetActionTypes.REF_SUCCESS ) ||
       (action.type == OrgGetActionTypes.REF_FAILURE ) ) {

     const data = {
       data: {
         num: state.data.num,
         refs: propsData.data.refs,
         names: state.data.names,
         types: state.data.types
       }
     }
     return (<any>Object).assign({}, data)

  } else if ( (action.type == OrgGetActionTypes.NAME_SUCCESS ) ||
       (action.type == OrgGetActionTypes.NAME_FAILURE ) ) {

     const data = {
       data: {
         num: state.data.num,
         refs: state.data.refs,
         names: propsData.data.names,
         types: state.data.types
       }
     }
     return (<any>Object).assign({}, data)

  } else if ( (action.type == OrgGetActionTypes.TYPE_SUCCESS ) ||
       (action.type == OrgGetActionTypes.TYPE_FAILURE ) ) {

     const data = {
       data: {
         num: state.data.num,
         refs: state.data.refs,
         names: state.data.names,
         types: propsData.data.types
       }
     }
     return (<any>Object).assign({}, data)

  } else {
    return state
  }
}
