import * as React from 'react'
import { Switch, Route, Link } from 'react-router-dom'
import MenuList from '@material-ui/core/MenuList'
import MenuItem from '@material-ui/core/MenuItem'
import logo from '../images/owl.png'
import Grid from '@material-ui/core/Grid'
import Paper from '@material-ui/core/Paper'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles';
import { withTheme, styles } from '../styles/theme'

import { AppStrings, PathStrings, HomeStrings, AboutStrings, OverviewStrings, HelpStrings } from '../utils/strings'

import Home from './pages/helpers/home'
import About from './pages/helpers/about'
import Overview from './pages/helpers/overview'
import Help from './pages/helpers/help'

class Main extends React.Component<WithStyles<typeof styles>> {

  render() {

    return (
      <div className={this.props.classes.root}>
        <Paper className={this.props.classes.paper}>
          <Grid container justify='center' spacing={0} wrap='nowrap'>
            <Grid item xs={12} sm={2}>
               <Paper className={this.props.classes.paper}><img src={logo}/></Paper>
            </Grid>
            <Grid item xs={12} sm={10}>
              <Paper className={this.props.classes.paper}><h1>{AppStrings.appTitle}</h1></Paper>
            </Grid>
          </Grid>
          <Grid container spacing={0}>
            <Grid item xs={12} sm={2}>
              <Paper className={this.props.classes.paper}>
                <MenuList>
                  <Link to={PathStrings.homePath}>
                    <MenuItem>
                      {PathStrings.home}
                    </MenuItem>
                  </Link>
                  <Link to={PathStrings.aboutPath}>
                    <MenuItem>
                      {PathStrings.about}
                    </MenuItem>
                  </Link>
                  <Link to={PathStrings.overviewPath}>
                    <MenuItem>
                      {PathStrings.overview}
                    </MenuItem>
                  </Link>
                  <Link to={PathStrings.helpPath}>
                    <MenuItem>
                      {PathStrings.help}
                    </MenuItem>
                  </Link>
                </MenuList>
              </Paper>
            </Grid>
            <Grid item xs={12} sm={10}>
              <Paper className={this.props.classes.paper}>
                <Switch>
                  <Route
                    name={PathStrings.home}
                    exact path={PathStrings.homePath}
                    render={() => <Home title={HomeStrings.heading} data={HomeStrings.info} />}
                  />
                  <Route
                    name={PathStrings.about}
                    path={PathStrings.aboutPath}
                    render={() => <About title={AboutStrings.heading} data={AboutStrings.info} />}
                  />
                  <Route
                    name={PathStrings.overview}
                    path={PathStrings.overviewPath}
                    render={() => <Overview title={OverviewStrings.heading} data={OverviewStrings.info} />}
                  />
                  <Route
                    name={PathStrings.help}
                    path={PathStrings.helpPath}
                    render={() => <Help title={HelpStrings.heading} data={HelpStrings.info} />}
                  />
                </Switch>
              </Paper>
            </Grid>
          </Grid>
          <Grid container spacing={0}>
            <Grid item xs={2}>
              <Paper className={this.props.classes.paper}>
                <p>&nbsp;</p>
              </Paper>
            </Grid>
            <Grid item xs={10}>
              <Paper className={this.props.classes.paper}>
                <h5>{AppStrings.copyright}</h5>
              </Paper>
            </Grid>
          </Grid>
        </Paper>
      </div>
    )
  }
}

export default withTheme(withStyles(styles)(Main))
