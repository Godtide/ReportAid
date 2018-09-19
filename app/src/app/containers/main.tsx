import * as React from 'react'
import Grid from '@material-ui/core/Grid'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles';
import { withTheme, styles } from '../styles/theme'

import Header from './layout/header'
import Sider from './layout/sider'
import Content from './layout/content'
import Footer from './layout/footer'
import { AppStrings } from '../utils/strings'

const Func = () => {
}

class Main extends React.Component<WithStyles<typeof styles>> {

  render() {

    return (
      <div className={this.props.classes.root}>
        <Grid container spacing={8}>
          <Grid xs={12}>
            <Header
              appTitle={AppStrings.appTitle}
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
          <Grid container spacing={8}>
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
          </Grid>
          <Grid xs={12}>
            <Footer
              copyright={AppStrings.copyright}
            />
          </Grid>
        </Grid>
      </div>
    )
  }
}

export default withTheme(withStyles(styles)(Main))
