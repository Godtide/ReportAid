import * as React from 'react'
import Grid from '@material-ui/core/Grid'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles';
import { withTheme, styles } from '../styles/theme'

import Header from './layout/header'
import Sider from './layout/sider'
import Content from './layout/content'
import Footer from './layout/footer'
import { AppStrings, PathStrings } from '../utils/strings'

const Func = () => {
}

class Main extends React.Component<WithStyles<typeof styles>> {

  render() {

    return (
      <div className={this.props.classes.root}>
        <Grid container justify='center' spacing={0} wrap='nowrap'>
          <Header
            appTitle={AppStrings.appTitle}
            handleClose={Func}
            home={PathStrings.home}
            homePath={PathStrings.homePath}
            about={PathStrings.about}
            aboutPath={PathStrings.aboutPath}
            overview={PathStrings.overview}
            overviewPath={PathStrings.overviewPath}
            help={PathStrings.help}
            helpPath={PathStrings.helpPath}
          />
        </Grid>
        <Grid container spacing={0}>
          <Sider
            handleClose={Func}
            home={PathStrings.home}
            homePath={PathStrings.homePath}
            about={PathStrings.about}
            aboutPath={PathStrings.aboutPath}
            overview={PathStrings.overview}
            overviewPath={PathStrings.overviewPath}
            help={PathStrings.help}
            helpPath={PathStrings.helpPath}
          />
          <Content
            home={PathStrings.home}
            homePath={PathStrings.homePath}
            about={PathStrings.about}
            aboutPath={PathStrings.aboutPath}
            overview={PathStrings.overview}
            overviewPath={PathStrings.overviewPath}
            help={PathStrings.help}
            helpPath={PathStrings.helpPath}
          />
        </Grid>
        <Grid container spacing={0}>
          <Footer
            copyright={AppStrings.copyright}
          />
        </Grid>
      </div>
    )
  }
}

export default withTheme(withStyles(styles)(Main))
