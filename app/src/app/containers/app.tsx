import * as React from 'react'
import CssBaseline from '@material-ui/core/CssBaseline';

import Header from './layout/header'
import Sider from './layout/sider'
import Content from './layout/content'
import Footer from './layout/footer'
import { AppStrings } from '../utils/strings'

  /*
  readonly appTitle='ReportAid'
  readonly home='home'
  readonly homePath='/'
  readonly about='about'
  readonly aboutPath='/about'
  readonly overview='overview'
  readonly overviewPath='/overview'
  readonly help='help'
  readonly helpPath='/help'
*/

const Func = () => {
}

const App = () => (
  <div>
    <CssBaseline />
    <Header
      headerTitle={AppStrings.appTitle}
      handleClose={Func}
      home={AppStrings.home}
      homePath={AppStrings.homePath}
      about={AppStrings.about}
      aboutPath={AppStrings.aboutPath}
      overview={AppStrings.overview}
      overviewPath={AppStrings.overviewPath}
      help={AppStrings.help}
      helpPath={AppStrings.helpPath}
    />
    <Sider
      handleClose={Func}
      home={AppStrings.home}
      homePath={AppStrings.homePath}
      about={AppStrings.about}
      aboutPath={AppStrings.aboutPath}
      overview={AppStrings.overview}
      overviewPath={AppStrings.overviewPath}
      help={AppStrings.help}
      helpPath={AppStrings.helpPath}
    />
    <Content
      home={AppStrings.home}
      homePath={AppStrings.homePath}
      about={AppStrings.about}
      aboutPath={AppStrings.aboutPath}
      overview={AppStrings.overview}
      overviewPath={AppStrings.overviewPath}
      help={AppStrings.help}
      helpPath={AppStrings.helpPath}
    />
    <Footer
      copyright={AppStrings.copyright}
    />
  </div>
)

export default App
