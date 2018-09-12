import * as React from 'react'
import { Route, Switch } from 'react-router-dom'

import Root from '../../components/layout/root'
import Header from '../../components/layout/header'
//  import Footer from '../../components/layout/footer'
// import { HomePage } from './home/homePage'
import IATIWriterPage from './IATIWriter/IATIWriterPage'
import IATIReaderPage from './IATIReader/IATIReaderPage'

// If your app is big + you have routes with a lot of components, you should consider
// code-splitting your routes! If you bundle stuff up with Webpack, I recommend `react-loadable`.
//
// $ yarn add react-loadable
// $ yarn add --dev @types/react-loadable
//
// The given `pages/` directory provides an example of a directory structure that's easily
// code-splittable.
const Routes: React.SFC = () => (
  <Root>
    <Header title="Example App" />
    <Switch>
      <Route path="/IATIWriter" component={IATIWriterPage} />
      <Route path="/IATIReader" component={IATIReaderPage} />
      <Route component={() => <div>Not Found</div>} />
    </Switch>
  </Root>
)

export default Routes

// <Route exact path="/" component={HomePage} />
