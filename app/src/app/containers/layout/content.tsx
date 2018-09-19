import * as React from 'react'
import { Route } from 'react-router-dom'

import Paper from '@material-ui/core/Paper'
import Grid from '@material-ui/core/Grid'
import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../styles/theme'

import Home from '../pages/helpers/home'
import About from '../pages/helpers/about'
import Overview from '../pages/helpers/overview'
import Help from '../pages/helpers/help'

import { Paths } from './types'

export type  AllProps = Paths

class Content extends React.Component<WithStyles<typeof styles> & AllProps> {

    render() {

      return (
        <React.Fragment>
          <Grid item xs={12} sm={10}>
            <Paper className={this.props.classes.paper}>
              <p>blah</p>
              <Route name={this.props.home} exact path={this.props.homePath} component={Home} />
              <Route name={this.props.about} path={this.props.aboutPath} component={About} />
              <Route name={this.props.overview} path={this.props.overviewPath} component={Overview} />
              <Route name={this.props.help} path={this.props.helpPath} component={Help} />
            </Paper>
          </Grid>
        </React.Fragment>
      )
    }
}

export default withTheme(withStyles(styles)(Content))

/*
<Breadcrumb style={{ margin: '16px 0' }}>
  <Breadcrumb.Item>{props.home}</Breadcrumb.Item>
  <Breadcrumb.Item>{props.about}</Breadcrumb.Item>
  <Breadcrumb.Item>{props.overview}</Breadcrumb.Item>
  <Breadcrumb.Item>{props.help}</Breadcrumb.Item>
</Breadcrumb>

<Route
  name={AppStrings.create}
  path={AppPaths.create}
  render={() => <Writer contracts={this.contractHandler} web3={this.web3Handler} />}
/>
<Route
  name={AppStrings.read}
  path={AppPaths.read}
  render={() => <Reader contracts={this.contractHandler} web3={this.web3Handler} />}
/>
*/
