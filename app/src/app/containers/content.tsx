import * as React from 'react'
import { Switch, Route } from 'react-router-dom'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../styles/theme'

import { Info } from './pages/info/info'
import { InfoTypes } from './pages/info/types'

import { BlockchainInfo } from './pages/blockchain/blockchainInfo'
import { Organisations as OrgWriter } from './pages/IATIWriter/organisation'
import { OrganisationReports as OrgReportsWriter } from './pages/IATIWriter/organisationReports'
import { OrganisationReportBudgets as OrgReportBudgetsWriter } from './pages/IATIWriter/organisationReportBudgets'
import { OrganisationReportExpenditure as OrgReportExpenditureWriter } from './pages/IATIWriter/organisationReportExpenditure'
import { OrganisationReportRecipientBudgets as OrgReportRecipientBudgetsWriter } from './pages/IATIWriter/organisationReportRecipientBudgets'
import { Organisations as OrgReader } from './pages/IATIReader/organisation'
import { OrganisationReports as OrgReportsReader } from './pages/IATIReader/organisationReports'
import { OrganisationReportBudgets as OrgReportBudgetsReader } from './pages/IATIReader/organisationReportBudgets'
import { OrganisationReportExpenditure as OrgReportExpenditureReader } from './pages/IATIReader/organisationReportExpenditure'
import { OrganisationReportRecipientBudgets as OrgReportRecipientBudgetsReader } from './pages/IATIReader/organisationReportRecipientBudgets'

import { Paths } from '../utils/strings'
import { Paths as PathConfig } from '../utils/config'

class AppContent extends React.Component<WithStyles<typeof styles>> {

  render() {

    return (
      <Switch>
        <Route name={Paths.home} exact path={PathConfig.home} render={() => <Info type={InfoTypes.HOME} />} />
        <Route name={Paths.blockchain} exact path={PathConfig.blockchain} render= {() => <BlockchainInfo />} />
        <Route name={Paths.about} exact path={PathConfig.about} render={() => <Info type={InfoTypes.ABOUT} />} />
        <Route name={Paths.overview} path={PathConfig.overview} render={() => <Info type={InfoTypes.OVERVIEW} />} />
        <Route name={Paths.help} path={PathConfig.help} render={() => <Info type={InfoTypes.HELP} />} />
        <Route name={Paths.writer} path={PathConfig.writer} render={() => <Info type={InfoTypes.IATIWriter} />} />
        <Route name={Paths.reader} path={PathConfig.reader} render={() => <Info type={InfoTypes.IATIReader} />} />
        <Route name={Paths.orgWriter} path={PathConfig.orgWriter} render={() => <OrgWriter />} />
        <Route name={Paths.orgReportWriter} path={PathConfig.orgReportWriter} render={() => <OrgReportsWriter />} />
        <Route name={Paths.orgReportBudgetsWriter} path={PathConfig.orgReportBudgetsWriter} render={() => <OrgReportBudgetsWriter />} />
        <Route name={Paths.orgReportExpenditureWriter} path={PathConfig.orgReportExpenditureWriter} render={() => <OrgReportExpenditureWriter />} />
        <Route name={Paths.orgReportRecipientBudgetsWriter} path={PathConfig.orgReportRecipientBudgetsWriter} render={() => <OrgReportRecipientBudgetsWriter />} />
        <Route name={Paths.orgReader} path={PathConfig.orgReader} render={() => <OrgReader />} />
        <Route name={Paths.orgReportReader} path={PathConfig.orgReportReader} render={() => <OrgReportsReader />} />
        <Route name={Paths.orgReportBudgetsReader} path={PathConfig.orgReportBudgetsReader} render={() => <OrgReportBudgetsReader />} />
        <Route name={Paths.orgReportExpenditureReader} path={PathConfig.orgReportExpenditureReader} render={() => <OrgReportExpenditureReader />} />
        <Route name={Paths.orgReportRecipientBudgetsReader} path={PathConfig.orgReportRecipientBudgetsReader} render={() => <OrgReportRecipientBudgetsReader />} />
      </Switch>
    )
  }
}

export const Content = withTheme(withStyles(styles)(AppContent))
