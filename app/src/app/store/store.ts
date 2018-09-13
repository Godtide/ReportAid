import { combineReducers, Dispatch, Action, AnyAction } from 'redux'
import { all, fork } from 'redux-saga/effects'

import IATIWriterSaga from './IATIWriter/sagas'
import { IATIWriterReducer } from './IATIWriter/reducer'
import { IATIWriterState } from './IATIWriter/types'
import IATIReaderSaga from './IATIReader/sagas'
import { IATIReaderReducer } from './IATIReader/reducer'
import { IATIReaderState } from './IATIReader/types'

// The top-level state object
export interface ApplicationState {
  writer: IATIWriterState
  reader: IATIReaderState
}

// Additional props for connected React components. This prop is passed by default with `connect()`
export interface ConnectedReduxProps<A extends Action = AnyAction> {
  dispatch: Dispatch<A>
}

// Whenever an action is dispatched, Redux will update each top-level application state property
// using the reducer with the matching name. It's important that the names match exactly, and that
// the reducer acts on the corresponding ApplicationState property type.
export const rootReducer = combineReducers<ApplicationState>({
  writer: IATIWriterReducer,
  reader: IATIReaderReducer
})

// Here we use `redux-saga` to trigger actions asynchronously. `redux-saga` uses something called a
// "generator function", which you can read about here:
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/function*
export function* rootSaga() {
  yield all([fork(IATIWriterSaga), fork(IATIReaderSaga)])
}
