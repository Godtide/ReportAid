import * as React from 'react'
import * as ReactDOM from 'react-dom'
import { BrowserRouter } from 'react-router-dom'
import { Provider } from 'react-redux'

import App from './app/containers/app'
import { configureStore } from './app/store'
import { AppStrings } from './app/utils/strings'


const initialState = (window as any).initialReduxState
const store = configureStore(initialState)

// Now you can dispatch navigation actions from anywhere!
// store.dispatch(push('/foo'))

ReactDOM.render((
  <Provider store={store}>
    <BrowserRouter>
      <App appTitle={AppStrings.appTitle}/>
    </BrowserRouter>
  </Provider>
), document.getElementById("root"))
