// import { combineReducers, Dispatch, Reducer, Action, AnyAction } from 'redux'
import { combineReducers, Reducer } from 'redux'
import { Store, createStore, applyMiddleware } from 'redux'
import thunkMiddleware from 'redux-thunk'

import { AboutProps } from './helpers/about/types'
import { AboutReducer } from './helpers/about/reducer'

// The top-level state object
export interface ApplicationState {
  about: AboutProps
  // home: HomeState
}

export const rootReducer: Reducer<ApplicationState> = combineReducers<ApplicationState>({
  about: AboutReducer,
})

export function configureStore(
  initialState: ApplicationState
): Store<ApplicationState> {

  // create the redux-saga middleware
  const store = createStore(
    rootReducer,
    initialState,
    applyMiddleware(thunkMiddleware)
  )

  // Don't forget to run the root saga, and return the store object.
  // sagaMiddleware.run(rootSaga)
  return store
}
