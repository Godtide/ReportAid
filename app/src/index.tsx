import * as React from 'react'
import * as ReactDOM from 'react-dom'
import { BrowserRouter } from 'react-router-dom'
import { Provider } from 'react-redux'

import App from './app/components/layout/app'
import configureStore from './app/store/configureStore'
import { Window } from './types'

const initialState = Window.initialReduxState
const store = configureStore(initialState)

// Now you can dispatch navigation actions from anywhere!
// store.dispatch(push('/foo'))

ReactDOM.render((
  <Provider store={store}>
    <BrowserRouter>
      <App />
    </BrowserRouter>
  </Provider>
), document.getElementById("root"))
