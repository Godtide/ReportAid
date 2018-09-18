import * as React from 'react'
import Grid from '@material-ui/core/Grid'
import Paper from '@material-ui/core/Paper'
import CssBaseline from '@material-ui/core/CssBaseline'

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

interface AppProps {
  appTitle: string
}

const Func = () => {
}

class App extends React.Component<AppProps> {

  public render() {

    return (
      <Grid container spacing={8}>
        <Grid item xs={12}>
          <Paper className={'blah'}><h1>{this.props.appTitle}</h1></Paper>
        </Grid>
        <Grid item xs={12}>
          <Header
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
        </Grid>
        <Grid item xs={2}>
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
        </Grid>
        <Grid item xs={10}>
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
        </Grid>
        <Grid item xs={12}>
          <Footer
            copyright={AppStrings.copyright}
          />
        </Grid>
      </Grid>
    )
  }
}

export default App
