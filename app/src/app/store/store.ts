//import { Action } from 'redux'
import { combineReducers, Reducer, Store, createStore, applyMiddleware } from 'redux'
//import thunkMiddleware, { ThunkAction } from 'redux-thunk'
import thunkMiddleware from 'redux-thunk'

import { ActionProps} from './types'
import { BlockchainProps } from './blockchain/types'
import { OrgWriterProps } from './IATI/IATIWriter/organisationWriter/types'
import { InfoProps } from './info/types'

import { reducer as blockchainReducer } from './blockchain/reducer'
import { reducer as aboutReducer } from './info/about/reducer'
import { reducer as homeReducer } from './info/home/reducer'
import { reducer as helpReducer } from './info/help/reducer'
import { reducer as overviewReducer } from './info/overview/reducer'
import { reducer as IATIWriterReducer } from './info/IATIWriter/reducer'
import { reducer as orgReducer } from './IATI/IATIWriter/organisationWriter/reducer'
import { reducer as IATIReaderReducer } from './info/IATIReader/reducer'

//export type StoreAction = BlockchainAction & OrgAddAction
//export type ThunkResult<R> = ThunkAction<R, ApplicationState, null, StoreAction>

// The top-level state object
export interface ApplicationState {
  blockchain: BlockchainProps
  about: InfoProps
  home: InfoProps
  help: InfoProps
  overview: InfoProps
  writer: InfoProps
  reader: InfoProps
  orgForm: OrgWriterProps
}

export const rootReducer: Reducer<ApplicationState, ActionProps> = combineReducers<ApplicationState, ActionProps>({
  blockchain: blockchainReducer,
  about: aboutReducer,
  home: homeReducer,
  help: helpReducer,
  overview: overviewReducer,
  writer: IATIWriterReducer,
  reader: IATIReaderReducer,
  orgForm: orgReducer
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
