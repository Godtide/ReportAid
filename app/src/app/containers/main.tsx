import * as React from 'react'
import { Switch, Route, Link } from 'react-router-dom'
import MenuList from '@material-ui/core/MenuList'
import MenuItem from '@material-ui/core/MenuItem'
import logo from '../images/owl.png'
import Grid from '@material-ui/core/Grid'
import Paper from '@material-ui/core/Paper'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles';
import { withTheme, styles } from '../styles/theme'

import { AppStrings, PathStrings, HomeStrings, AboutStrings, OverviewStrings, HelpStrings, WriterStrings, ReaderStrings } from '../utils/strings'

import Home from './pages/helpers/home'
import About from './pages/helpers/about'
import Overview from './pages/helpers/overview'
import Help from './pages/helpers/help'
import IATIWriter from './pages/IATIWriter/IATIWriter'
import IATIReader from './pages/IATIReader/IATIReader'

class Main extends React.Component<WithStyles<typeof styles>> {

  render() {

    return (
      <Paper className={this.props.classes.root}>
        <Paper className={this.props.classes.header}>
          <Grid container justify='center' spacing={0}>
            <Grid item xs={12} sm={2}>
               <Paper className={this.props.classes.sider}><img src={logo}/></Paper>
            </Grid>
            <Grid item xs={12} sm={10}>
              <Paper className={this.props.classes.header}><h1>{AppStrings.appTitle}</h1></Paper>
            </Grid>
          </Grid>
        </Paper>
        <Paper className={this.props.classes.content}>
          <Grid container spacing={0}>
            <Grid item xs={12} sm={2}>
              <Paper className={this.props.classes.sider}>
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
                  <Link to={PathStrings.writerPath}>
                    <MenuItem>
                      {PathStrings.writer}
                    </MenuItem>
                  </Link>
                  <Link to={PathStrings.readerPath}>
                    <MenuItem>
                      {PathStrings.reader}
                    </MenuItem>
                  </Link>
                </MenuList>
              </Paper>
            </Grid>
            <Grid item xs={12} sm={10}>
              <Paper className={this.props.classes.content}>
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
                  <Route
                    name={PathStrings.writer}
                    path={PathStrings.writerPath}
                    render={() => <IATIWriter title={WriterStrings.heading} data={''} />}
                  />
                  <Route
                    name={PathStrings.reader}
                    path={PathStrings.readerPath}
                    render={() => <IATIReader title={ReaderStrings.heading}  data={''} />}
                  />
                </Switch>
              </Paper>
            </Grid>
          </Grid>
        </Paper>
        <Paper className={this.props.classes.footer}>
          <Grid container spacing={0}>
            <Grid item xs={2}>
              <Paper className={this.props.classes.footer}>
                <p>{AppStrings.author}</p>
              </Paper>
            </Grid>
            <Grid item xs={10}>
              <Paper className={this.props.classes.footer}>
                <p>{AppStrings.copyright}</p>
              </Paper>
            </Grid>
          </Grid>
        </Paper>
      </Paper>
    )
  }
}

export default withTheme(withStyles(styles)(Main))
