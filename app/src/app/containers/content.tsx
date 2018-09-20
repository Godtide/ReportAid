import * as React from 'react'
import { Switch, Route } from 'react-router-dom'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../styles/theme'

import Home from './pages/helpers/home'
import About from './pages/helpers/about'
import Overview from './pages/helpers/overview'
import Help from './pages/helpers/help'
import IATIWriter from './pages/IATIWriter/IATIWriter'
import IATIReader from './pages/IATIReader/IATIReader'

import { PathStrings, HomeStrings, AboutStrings, OverviewStrings, HelpStrings, WriterStrings, ReaderStrings } from '../utils/strings'

class Content extends React.Component<WithStyles<typeof styles>> {

    render() {

      return (
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
      )
    }
}

export default withTheme(withStyles(styles)(Content))
