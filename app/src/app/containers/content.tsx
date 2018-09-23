import * as React from 'react'
import { Switch, Route } from 'react-router-dom'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../styles/theme'

import Info from './pages/info/info'
import { InfoTypes } from './pages/info/types'
import IATIWriter from './pages/IATIWriter/IATIWriter'
import IATIReader from './pages/IATIReader/IATIReader'

import { PathStrings, WriterStrings, ReaderStrings } from '../utils/strings'

class Content extends React.Component<WithStyles<typeof styles>> {

    render() {

      return (
        <Switch>
          <Route
            name={PathStrings.home}
            exact path={PathStrings.homePath}
            render={() => <Info type={InfoTypes.HOME} />}
          />
          <Route
            name={PathStrings.about}
            exact path={PathStrings.aboutPath}
            render={() => <Info type={InfoTypes.ABOUT} />}
          />
          <Route
            name={PathStrings.overview}
            path={PathStrings.overviewPath}
            render={() => <Info type={InfoTypes.OVERVIEW} />}
          />
          <Route
            name={PathStrings.help}
            path={PathStrings.helpPath}
            render={() => <Info type={InfoTypes.HELP} />}
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
