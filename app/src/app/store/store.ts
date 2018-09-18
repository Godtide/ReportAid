// import { combineReducers, Dispatch, Reducer, Action, AnyAction } from 'redux'
import { combineReducers, Reducer } from 'redux'
import { Store, createStore, applyMiddleware } from 'redux'
import createSagaMiddleware from 'redux-saga'
//import { all, fork } from 'redux-saga/effects'

/* import IATIWriterSaga from './IATIWriter/sagas'
import { IATIWriterReducer } from './IATIWriter/reducer'
import { IATIWriterState } from './IATIWriter/types'
import IATIReaderSaga from './IATIReader/sagas'
import { IATIReaderReducer } from './IATIReader/reducer'
import { IATIReaderState } from './IATIReader/types'

import { appFormReducer } from './helpers/form'
import { FormState } from './helpers/form' */


import { IATIWriterState } from './IATIWriter/types'
import { IATIWriterReducer } from './IATIWriter/reducer'

// The top-level state object
export interface ApplicationState {
  writer: IATIWriterState
  /* reader: IATIReaderState
  form: appFormReducer */
  // home: HomeState
}

/*
export interface ConnectedReduxProps<A extends Action = AnyAction> {
  dispatch: Dispatch<A>
}*/

export const rootReducer: Reducer<ApplicationState> = combineReducers<ApplicationState>({
  writer: IATIWriterReducer,
  /* reader: IATIReaderReducer,
  form: formReducer */
  // home: HomeReducer
})

// Here we use `redux-saga` to trigger actions asynchronously. `redux-saga` uses something called a
// "generator function", which you can read about here:
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/function*
/* export function* rootSaga() {
  yield all([fork(IATIWriterSaga), fork(IATIReaderSaga)])
} */

export function configureStore(
  initialState: ApplicationState
): Store<ApplicationState> {

  // create the redux-saga middleware
  const sagaMiddleware = createSagaMiddleware()
  const store = createStore(
    rootReducer,
    initialState,
    applyMiddleware(sagaMiddleware)
  )

  // Don't forget to run the root saga, and return the store object.
  // sagaMiddleware.run(rootSaga)
  return store
}
