import * as React from 'react'
import { Route, Switch } from 'react-router-dom'

import Home from './pages/helpers/home'
import IATIWriter from './IATIWriter/IATIWriter'
import IATIReaderPage from './IATIReader/IATIReader'

const Routes: React.SFC = () => (
  <div>
    <Switch>
      <Route exact path="/" component={HomePage} />
      <Route component={() => <div>Not Found</div>} />
    </Switch>
  </div>
)

export default Routes

// <Route exact path="/" component={HomePage} />
<Route path="/IATIWriter" component={IATIWriterPage} />
<Route path="/IATIReader" component={IATIReaderPage} />
