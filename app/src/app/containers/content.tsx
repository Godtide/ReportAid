import * as React from 'react'
import { Switch, Route } from 'react-router-dom'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../styles/theme'

import { Info } from './pages/info/info'
import { InfoTypes } from './pages/info/types'

import { BlockchainInfo } from './pages/blockchain/blockchainInfo'
import { Org as OrgWriter } from './pages/IATIWriter/org'
import { Organisations as OrgsWriter } from './pages/IATIWriter/organisations'
import { OrganisationDocs as OrgDocsWriter } from './pages/IATIWriter/organisationDocs'
import { OrganisationBudgets as OrgBudgetsWriter } from './pages/IATIWriter/organisationBudgets'
import { OrganisationExpenditure as OrgExpenditureWriter } from './pages/IATIWriter/organisationExpenditure'
import { OrganisationRecipientBudgets as OrgRecipientBudgetsWriter } from './pages/IATIWriter/organisationRecipientBudgets'
import { OrganisationRegionBudgets as OrgRegionBudgetsWriter } from './pages/IATIWriter/organisationRegionBudgets'
import { OrganisationCountryBudgets as OrgCountryBudgetsWriter } from './pages/IATIWriter/organisationCountryBudgets'

import { Orgs as OrgReader } from './pages/IATIReader/orgs'
import { Organisations as OrgsReader } from './pages/IATIReader/organisations'
import { OrganisationDocs as OrgDocsReader } from './pages/IATIReader/organisationDocs'
import { OrganisationBudgets as OrgBudgetsReader } from './pages/IATIReader/organisationBudgets'
import { OrganisationExpenditure as OrgExpenditureReader } from './pages/IATIReader/organisationExpenditure'
import { OrganisationRecipientBudgets as OrgRecipientBudgetsReader } from './pages/IATIReader/organisationRecipientBudgets'
import { OrganisationRegionBudgets as OrgRegionBudgetsReader } from './pages/IATIReader/organisationRegionBudgets'
import { OrganisationCountryBudgets as OrgCountryBudgetsReader } from './pages/IATIReader/organisationCountryBudgets'

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
        <Route name={Paths.orgWriter} path={PathConfig.orgWriter} render={() => <OrgsWriter />} />
        <Route name={Paths.orgDocsWriter} path={PathConfig.orgDocsWriter} render={() => <OrgDocsWriter />} />
        <Route name={Paths.orgBudgetsWriter} path={PathConfig.orgBudgetsWriter} render={() => <OrgBudgetsWriter />} />
        <Route name={Paths.orgExpenditureWriter} path={PathConfig.orgExpenditureWriter} render={() => <OrgExpenditureWriter />} />
        <Route name={Paths.orgRecipientBudgetsWriter} path={PathConfig.orgRecipientBudgetsWriter} render={() => <OrgRecipientBudgetsWriter />} />
        <Route name={Paths.orgRegionBudgetsWriter} path={PathConfig.orgRegionBudgetsWriter} render={() => <OrgRegionBudgetsWriter />} />
        <Route name={Paths.orgCountryBudgetsWriter} path={PathConfig.orgCountryBudgetsWriter} render={() => <OrgCountryBudgetsWriter />} />
        <Route name={Paths.orgReader} path={PathConfig.orgReader} render={() => <OrgReader />} />
        <Route name={Paths.orgReader} path={PathConfig.orgReader} render={() => <OrgsReader />} />
        <Route name={Paths.orgDocsReader} path={PathConfig.orgDocsReader} render={() => <OrgDocsReader />} />
        <Route name={Paths.orgBudgetsReader} path={PathConfig.orgBudgetsReader} render={() => <OrgBudgetsReader />} />
        <Route name={Paths.orgExpenditureReader} path={PathConfig.orgExpenditureReader} render={() => <OrgExpenditureReader />} />
        <Route name={Paths.orgRecipientBudgetsReader} path={PathConfig.orgRecipientBudgetsReader} render={() => <OrgRecipientBudgetsReader />} />
        <Route name={Paths.orgRegionBudgetsReader} path={PathConfig.orgRegionBudgetsReader} render={() => <OrgRegionBudgetsReader />} />
        <Route name={Paths.orgCountryBudgetsReader} path={PathConfig.orgCountryBudgetsReader} render={() => <OrgCountryBudgetsReader />} />
      </Switch>
    )
  }
}

export const Content = withTheme(withStyles(styles)(AppContent))
