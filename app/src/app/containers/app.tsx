import * as React from 'react'
import Grid from '@material-ui/core/Grid'
import Paper from '@material-ui/core/Paper'
//import CssBaseline from '@material-ui/core/CssBaseline'

import { ClassNames } from '../styles/theme'
import CssBaseline from '@material-ui/core/CssBaseline';
import { MuiThemeProvider, withTheme } from '@material-ui/core/styles'
import { Theme } from '../styles/theme'

import Header from './layout/header'
import Sider from './layout/sider'
import Content from './layout/content'
import Footer from './layout/footer'
import { AppStrings } from '../utils/strings'

interface AppProps {
  appTitle: string
}

//type  AllProps = AppProps & ClassNames

const Func = () => {
}

const App: React.SFC<AppProps> = (props: AppProps) => {

  return (
    <React.Fragment>
      <MuiThemeProvider theme={Theme}>
        <CssBaseline />
        <div className='vblah'>
          <Grid container spacing={0}>
            <Grid item xs={12}>
              <Paper><h1>{props.appTitle}</h1></Paper>
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
        </div>
      </MuiThemeProvider>
    </React.Fragment>
  )
}

export default withTheme()(App)
