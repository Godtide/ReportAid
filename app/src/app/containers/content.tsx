import * as React from 'react'
import { Switch, Route } from 'react-router-dom'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../styles/theme'

import { Info } from './pages/info/info'
import { InfoTypes } from './pages/info/types'

import { BlockchainInfo } from './pages/blockchain/blockchainInfo'
import { Activities as ActivitiesWriter } from './pages/IATIWriter/activities'
import { Org as OrgWriter } from './pages/IATIWriter/org'
import { Organisations as OrganisationsWriter } from './pages/IATIWriter/organisations'
import { Organisation as OrganisationWriter } from './pages/IATIWriter/organisation'
import { OrganisationDocs as OrganisationDocsWriter } from './pages/IATIWriter/organisationDocs'
import { OrganisationBudgets as OrganisationBudgetsWriter } from './pages/IATIWriter/organisationBudgets'
import { OrganisationExpenditure as OrganisationExpenditureWriter } from './pages/IATIWriter/organisationExpenditure'
import { OrganisationRecipientBudgets as OrganisationRecipientBudgetsWriter } from './pages/IATIWriter/organisationRecipientBudgets'
import { OrganisationRegionBudgets as OrganisationRegionBudgetsWriter } from './pages/IATIWriter/organisationRegionBudgets'
import { OrganisationCountryBudgets as OrganisationCountryBudgetsWriter } from './pages/IATIWriter/organisationCountryBudgets'

import { Activities as ActivitiesReader } from './pages/IATIReader/activities'
import { Orgs as OrgsReader } from './pages/IATIReader/orgs'
import { Organisations as OrganisationsReader } from './pages/IATIReader/organisations'
import { Organisation as OrganisationReader } from './pages/IATIReader/organisation'
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

        <Route
          name={Paths.activitiesWriter}
          path={PathConfig.activitiesWriter}
          render={() => <ActivitiesWriter />}
        />
        <Route
          name={Paths.orgWriter}
          path={PathConfig.orgWriter}
          render={() => <OrgWriter />}
        />
        <Route
          name={Paths.organisationsWriter}
          path={PathConfig.organisationsWriter}
          render={() => <OrganisationsWriter />}
        />
        <Route
          name={Paths.organisationWriter}
          path={PathConfig.organisationWriter}
          render={() => <OrganisationWriter />}
        />
        <Route
          name={Paths.organisationDocsWriter}
          path={PathConfig.organisationDocsWriter}
          render={() => <OrganisationDocsWriter />}
        />
        <Route
          name={Paths.organisationBudgetsWriter}
          path={PathConfig.organisationBudgetsWriter}
          render={() => <OrganisationBudgetsWriter />}
        />
        <Route
          name={Paths.organisationExpenditureWriter}
          path={PathConfig.organisationExpenditureWriter}
          render={() => <OrganisationExpenditureWriter />}
        />
        <Route
          name={Paths.organisationRecipientBudgetsWriter}
          path={PathConfig.organisationRecipientBudgetsWriter}
          render={() => <OrganisationRecipientBudgetsWriter />}
        />
        <Route
          name={Paths.organisationRegionBudgetsWriter}
          path={PathConfig.organisationRegionBudgetsWriter}
          render={() => <OrganisationRegionBudgetsWriter />}
        />
        <Route
          name={Paths.organisationCountryBudgetsWriter}
          path={PathConfig.organisationCountryBudgetsWriter}
          render={() => <OrganisationCountryBudgetsWriter />}
        />

        <Route
          name={Paths.activitiesReader}
          path={PathConfig.activitiesReader}
          render={() => <ActivitiesReader />}
        />
        <Route
          name={Paths.orgsReader}
          path={PathConfig.orgsReader}
          render={() => <OrgsReader />}
        />
        <Route
          name={Paths.organisationsReader}
          path={PathConfig.organisationsReader}
          render={() => <OrganisationsReader />}
        />
        <Route
          name={Paths.organisationReader}
          path={PathConfig.organisationReader}
          render={() => <OrganisationReader />}
        />
        <Route
          name={Paths.organisationDocsReader}
          path={PathConfig.organisationDocsReader}
          render={() => <OrgDocsReader />}
        />
        <Route
          name={Paths.organisationBudgetsReader}
          path={PathConfig.organisationBudgetsReader}
          render={() => <OrgBudgetsReader />}
        />
        <Route
          name={Paths.organisationExpenditureReader}
          path={PathConfig.organisationExpenditureReader}
          render={() => <OrgExpenditureReader />}
        />
        <Route
          name={Paths.organisationRecipientBudgetsReader}
          path={PathConfig.organisationRecipientBudgetsReader}
          render={() => <OrgRecipientBudgetsReader />}
        />
        <Route
          name={Paths.organisationRegionBudgetsReader}
          path={PathConfig.organisationRegionBudgetsReader}
          render={() => <OrgRegionBudgetsReader />}
        />
        <Route
          name={Paths.organisationCountryBudgetsReader}
          path={PathConfig.organisationCountryBudgetsReader}
          render={() => <OrgCountryBudgetsReader />}
        />

      </Switch>
    )
  }
}

export const Content = withTheme(withStyles(styles)(AppContent))
