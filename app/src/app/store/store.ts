// import { combineReducers, Dispatch, Reducer, Action, AnyAction } from 'redux'
import { combineReducers, Reducer } from 'redux'
import { Store, createStore, applyMiddleware } from 'redux'
import thunkMiddleware from 'redux-thunk'

import { BlockchainInfoProps } from './blockchain/types'
import { InfoProps } from './info/types'

import { blockchainReducer } from './blockchain/reducer'
import { reducer as aboutReducer } from './info/about/reducer'
import { reducer as homeReducer } from './info/home/reducer'
import { reducer as helpReducer } from './info/help/reducer'
import { reducer as overviewReducer } from './info/overview/reducer'

// The top-level state object
export interface ApplicationState {
  blockchain: BlockchainInfoProps
  about: InfoProps
  home: InfoProps
  help: InfoProps
  overview: InfoProps
}

export const rootReducer: Reducer<ApplicationState> = combineReducers<ApplicationState>({
  blockchain: blockchainReducer,
  about: aboutReducer,
  home: homeReducer,
  help: helpReducer,
  overview: overviewReducer
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

  return store
}
