import * as React from 'react'
import * as ReactDOM from 'react-dom'
import { BrowserRouter } from 'react-router-dom'

import App from './app/containers/app'
import configureStore from './app/store/configureStore'
import { Window } from './types'

const initialState = Window.initialReduxState
const store = configureStore(initialState)

// Now you can dispatch navigation actions from anywhere!
// store.dispatch(push('/foo'))

ReactDOM.render((
  <BrowserRouter>
    <App store={store} />
  </BrowserRouter>
), document.getElementById("root"))
