import * as React from 'react'
import * as ReactDOM from 'react-dom'
import createHistory from "history/createBrowserHistory"

import Main from './app/containers/main'
import * as serviceWorker from './app/utils/serviceWorker'
import configureStore from './app/store/configureStore'

import 'typeface-ibm-plex-sans'
import './styles'

// Create a history of your choosing (we're using a browser history in this case)
const history = createHistory()
const initialState = (window as any).initialReduxState
const store = configureStore(history, initialState)

// Now you can dispatch navigation actions from anywhere!
// store.dispatch(push('/foo'))

ReactDOM.render(
  <Main store={store} history={history} />,
  document.getElementById("root")
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: http://bit.ly/CRA-PWA
serviceWorker.unregister()
