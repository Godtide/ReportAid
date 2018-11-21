import * as React from 'react'
import { Switch, Route } from 'react-router-dom'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../styles/theme'

import { BlockchainInfo } from './pages/blockchain/blockchainInfo'

import { Info } from './pages/info/info'
import { InfoTypes } from './pages/info/types'
import { IATIWriter } from './pages/IATIWriter/IATIWriter'
import { IATIReader } from './pages/IATIReader/IATIReader'

import { Paths, Writer, Reader } from '../utils/strings'
import { Paths as PathConfig } from '../utils/config'

class AppContent extends React.Component<WithStyles<typeof styles>> {

  render() {

    return (
      <Switch>
        <Route
          name={Paths.home}
          exact path={PathConfig.home}
          render={() => <Info type={InfoTypes.HOME} />}
        />
        <Route
          name={Paths.blockchain}
          exact path={PathConfig.blockchain}
          render= {() => <BlockchainInfo />}
        />
        <Route
          name={Paths.about}
          exact path={PathConfig.about}
          render={() => <Info type={InfoTypes.ABOUT} />}
        />
        <Route
          name={Paths.overview}
          path={PathConfig.overview}
          render={() => <Info type={InfoTypes.OVERVIEW} />}
        />
        <Route
          name={Paths.help}
          path={PathConfig.help}
          render={() => <Info type={InfoTypes.HELP} />}
        />
        <Route
          name={Paths.writer}
          path={PathConfig.writer}
          render={() => <IATIWriter title={Writer.heading} data={''} />}
        />
        <Route
          name={Paths.reader}
          path={PathConfig.reader}
          render={() => <IATIReader title={Reader.heading}  data={''} />}
        />
      </Switch>
    )
  }
}

export const Content = withTheme(withStyles(styles)(AppContent))
