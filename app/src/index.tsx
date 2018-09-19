import * as React from 'react'
import { render } from "react-dom"
import { BrowserRouter } from 'react-router-dom'
import { Provider } from 'react-redux'
import CssBaseline from '@material-ui/core/CssBaseline'

import  Main from './app/containers/main'
import { configureStore } from './app/store'
import { AppStrings } from './app/utils/strings'

const initialState = (window as any).initialReduxState
const store = configureStore(initialState)

const App = () => (
    <Provider store={store}>
      <React.Fragment>
        <CssBaseline />
        <BrowserRouter>
          <Main />
        </BrowserRouter>
      </React.Fragment>
    </Provider>
);

render(<App />, document.querySelector("#root"))
