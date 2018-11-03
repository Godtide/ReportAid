import * as React from 'react'
import { Switch, Route } from 'react-router-dom'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../styles/theme'

import Info from './pages/info/info'
import { InfoTypes } from './pages/info/types'
import IATIWriter from './pages/IATIWriter/IATIWriter'
import IATIReader from './pages/IATIReader/IATIReader'

import Web3Handler from '../utils/web3Handler'
import ContractHandler from '../utils/contractHandler'

import { PathStrings, WriterStrings, ReaderStrings } from '../utils/strings'

class Content extends React.Component<WithStyles<typeof styles>> {

    render() {

      /* const web3Handler = new Web3Handler()
      const contractHandler = new ContractHandler(web3Handler) */

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
