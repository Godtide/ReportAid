import * as React from 'react'

import { Switch, Route } from 'react-router-dom'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../styles/theme'

import { Info } from './pages/info/info'
import { InfoTypes } from './pages/info/types'

import { BlockchainInfo } from './pages/blockchain/blockchainInfo'

import { Activities as ActivitiesWriter } from './pages/IATIWriter/create/activities'
import { Activity as ActivityWriter } from './pages/IATIWriter/create/activity'
import { ActivityAdditional as ActivityAdditionalsWriter } from './pages/IATIWriter/create/activityAdditionals'
import { ActivityDate as ActivityDatesWriter } from './pages/IATIWriter/create/activityDates'
import { ActivityParticipatingOrg as ActivityParticipatingOrgWriter } from './pages/IATIWriter/create/activityParticipatingOrgs'
import { ActivitySector as ActivitySectorsWriter } from './pages/IATIWriter/create/activitySectors'
import { ActivityBudgets as ActivityBudgetsWriter } from './pages/IATIWriter/create/activityBudgets'
import { Org as OrgWriter } from './pages/IATIWriter/create/org'
import { Organisations as OrganisationsWriter } from './pages/IATIWriter/create/organisations'
import { Organisation as OrganisationWriter } from './pages/IATIWriter/create/organisation'
import { OrganisationDocs as OrganisationDocsWriter } from './pages/IATIWriter/create/organisationDocs'
import { OrganisationBudgets as OrganisationBudgetsWriter } from './pages/IATIWriter/create/organisationBudgets'
import { OrganisationExpenditure as OrganisationExpenditureWriter } from './pages/IATIWriter/create/organisationExpenditure'
import { OrganisationRecipientBudgets as OrganisationRecipientBudgetsWriter } from './pages/IATIWriter/create/organisationRecipientBudgets'
import { OrganisationRegionBudgets as OrganisationRegionBudgetsWriter } from './pages/IATIWriter/create/organisationRegionBudgets'
import { OrganisationCountryBudgets as OrganisationCountryBudgetsWriter } from './pages/IATIWriter/create/organisationCountryBudgets'

import { Activities as ActivitiesUpdater } from './pages/IATIWriter/update/activities'
import { Activity as ActivityUpdater } from './pages/IATIWriter/update/activity'
import { ActivityDate as ActivityDateUpdater } from './pages/IATIWriter/update/activityDate'
import { Org as OrgUpdater } from './pages/IATIWriter/update/org'

import { Activities as ActivitiesReader } from './pages/IATIReader/activities'
import { Activity as ActivityReader } from './pages/IATIReader/activity'
import { ActivityAdditionals as ActivityAdditionalsReader } from './pages/IATIReader/activityAdditionals'
import { ActivityDates as ActivityDatesReader } from './pages/IATIReader/activityDates'
import { ActivityParticipatingOrg as ActivityParticipatingOrgsReader } from './pages/IATIReader/activityParticipatingOrgs'
import { ActivitySectors as ActivitySectorsReader } from './pages/IATIReader/activitySectors'
import { ActivityBudgets as ActivityBudgetsReader } from './pages/IATIReader/activityBudgets'
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
          name={Paths.activityWriter}
          path={PathConfig.activityWriter}
          render={() => <ActivityWriter />}
        />
        <Route
          name={Paths.activityAdditionalWriter}
          path={PathConfig.activityAdditionalWriter}
          render={() => <ActivityAdditionalsWriter />}
        />
        <Route
          name={Paths.activityDatesWriter}
          path={PathConfig.activityDatesWriter}
          render={() => <ActivityDatesWriter />}
        />
        <Route
          name={Paths.activityParticipatingOrgWriter}
          path={PathConfig.activityParticipatingOrgWriter}
          render={() => <ActivityParticipatingOrgWriter />}
        />
        <Route
          name={Paths.activitySectorsWriter}
          path={PathConfig.activitySectorsWriter}
          render={() => <ActivitySectorsWriter />}
        />
        <Route
          name={Paths.activityBudgetsWriter}
          path={PathConfig.activityBudgetsWriter}
          render={() => <ActivityBudgetsWriter />}
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
          name={Paths.activitiesUpdater}
          path={PathConfig.activitiesUpdater}
          render={() => <ActivitiesUpdater />}
        />
        <Route
          name={Paths.activityUpdater}
          path={PathConfig.activityUpdater}
          render={() => <ActivityUpdater />}
        />
        <Route
          name={Paths.activityDateUpdater}
          path={PathConfig.activityDateUpdater}
          render={() => <ActivityDateUpdater />}
        />

        <Route
          name={Paths.orgUpdater}
          path={PathConfig.orgUpdater}
          render={() => <OrgUpdater />}
        />

        <Route
          name={Paths.activitiesReader}
          path={PathConfig.activitiesReader}
          render={() => <ActivitiesReader />}
        />
        <Route
          name={Paths.activityReader}
          path={PathConfig.activityReader}
          render={() => <ActivityReader />}
        />
        <Route
          name={Paths.activityAdditionalReader}
          path={PathConfig.activityAdditionalReader}
          render={() => <ActivityAdditionalsReader />}
        />
        <Route
          name={Paths.activityDatesReader}
          path={PathConfig.activityDatesReader}
          render={() => <ActivityDatesReader />}
        />
        <Route
          name={Paths.activityParticipatingOrgReader}
          path={PathConfig.activityParticipatingOrgReader}
          render={() => <ActivityParticipatingOrgsReader />}
        />
        <Route
          name={Paths.activitySectorsReader}
          path={PathConfig.activitySectorsReader}
          render={() => <ActivitySectorsReader />}
        />
        <Route
          name={Paths.activityBudgetsReader}
          path={PathConfig.activityBudgetsReader}
          render={() => <ActivityBudgetsReader />}
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
