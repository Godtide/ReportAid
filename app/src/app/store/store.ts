import * as React from 'react'
import { createStore, applyMiddleware } from 'redux'
import { RouteComponentProps } from 'react-router'
import { routerMiddleware, connectRouter } from 'connected-react-router'
import { createBrowserHistory } from 'history'
import { composeWithDevTools } from 'redux-devtools-extension'
import Immutable from 'immutable'

import { logger } from '../middleware'

const history = createBrowserHistory()
const initialState = Immutable.Map()

export namespace Store {
  export interface Props extends RouteComponentProps<void> {
    history: history,
    initialState: initialState
  }
}

export class Store extends React.Component<Store.Props> {

  constructor(props: Store.Props, context?: any) {
    super(props, context)
  }

  getHistory () {
    return this.props.history
  }

  configureStore () {

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
}
