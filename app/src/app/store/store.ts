import { Store, createStore, compose, applyMiddleware } from 'redux'
import { routerMiddleware, connectRouter } from 'connected-react-router/immutable'
import { composeWithDevTools } from 'redux-devtools-extension'

import { logger } from '../middleware'
import { RootState, rootReducer } from '../reducers'


export function configureStore(history: history, initialState?: RootState): Store<RootState> {
  let middleware = applyMiddleware(logger, routerMiddleware(history))

  if (process.env.NODE_ENV !== 'production') {
    middleware = composeWithDevTools(middleware);
  }

  const store = createStore(
    connectRouter(history)(rootReducer), // new root reducer with router state
    initialState,
    middleware
  )

  return store
}
