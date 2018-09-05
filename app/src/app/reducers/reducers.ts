// import { TodoModel } from 'app/models'
import { combineReducers } from 'redux-immutable'
import { RouterState } from 'connected-react-router'

export interface RootState {
  //todos: RootState.TodoState
  router: RouterState
}

export namespace RootState {
  // export type TodoState = TodoModel[]
}

// NOTE: current type definition of Reducer in 'react-router-redux' and 'redux-actions' module
// doesn't go well with redux@4
export const rootReducer = combineReducers<RootState>({})
