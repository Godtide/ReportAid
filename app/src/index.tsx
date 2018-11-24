import * as React from 'react'
import { render } from "react-dom"
import { HashRouter } from 'react-router-dom'
import { Provider } from 'react-redux'
import CssBaseline from '@material-ui/core/CssBaseline'

import { OrganisationContract } from './app/containers/blockchain/orgsContract'
import { setBlockchain } from './app/components/blockchain/blockchain'

import { Main } from './app/containers/main'
import { configureStore } from './app/store'

const initialState = (window as any).initialReduxState
const store = configureStore(initialState)
setBlockchain({store: store})

const App = () => (
    <Provider store={store}>
      <React.Fragment>
        <CssBaseline />
        <OrganisationContract />
        <HashRouter>
          <Main />
        </HashRouter>
      </React.Fragment>
    </Provider>
);

render(<App />, document.querySelector("#root"))
