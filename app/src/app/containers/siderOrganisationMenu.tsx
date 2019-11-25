import * as React from 'react'
import { Link } from 'react-router-dom'
import MenuList from '@material-ui/core/MenuList'
import MenuItem from '@material-ui/core/MenuItem'

import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'
import { ApplicationState } from '../store'
import { ActionProps } from '../store/types'

import { initialise } from '../store/IATI/actions'
import { setKey } from '../store/helpers/keys/actions'
import { Keys, KeyTypes } from '../store/helpers/keys/types'

import IconButton from '@material-ui/core/IconButton'
import Update from '@material-ui/icons/Update'
import Create from '@material-ui/icons/Create'
import List from '@material-ui/icons/List'

import { useStyles } from '../styles/theme';

import { Paths, App } from '../utils/strings'
import { Paths as PathConfig } from '../utils/config'

import ExpansionPanel from '@material-ui/core/ExpansionPanel'
import ExpansionPanelSummary from '@material-ui/core/ExpansionPanelSummary'
import ExpansionPanelDetails from '@material-ui/core/ExpansionPanelDetails'
import ExpandMoreIcon from '@material-ui/icons/ExpandMore'

interface DispatchProps {
  initialise: () => void
  setKey: (keyProps: Keys) => void
}

const defaultProps: DispatchProps = {
  initialise: () => {},
  setKey: (keyProps: Keys) => {}
}

const Sider = (props: DispatchProps = defaultProps) => {

  const classes = useStyles()

  return (
    <div>
      <ExpansionPanel className={classes.siderMenu}>
        <ExpansionPanelSummary expandIcon={<ExpandMoreIcon />}>
        <IconButton aria-label={App.headingActivitiesWriter}>
          <Create />
        </IconButton>
          <h4>{App.headingActivitiesWriter}</h4>
        </ExpansionPanelSummary>
        <ExpansionPanelDetails className={classes.siderMenu}>
          <MenuList>
            <Link
              to={PathConfig.activitiesWriter}
              onClick={() => {
                props.initialise()
                props.setKey({key: '', keyType: KeyTypes.ACTIVITIES})
              }} >
              <MenuItem>
                {Paths.activitiesWriter}
              </MenuItem>
            </Link>
            <Link
              to={PathConfig.activityWriter}
              onClick={() => {
                props.initialise()
                props.setKey({key: '', keyType: KeyTypes.ACTIVITY})
              }} >
              <MenuItem>
                {Paths.activityWriter}
              </MenuItem>
            </Link>
            <Link
              to={PathConfig.activityAdditionalWriter}
              onClick={() => {
                props.initialise()
                props.setKey({key: '', keyType: KeyTypes.ACTIVITYADDITIONAL})
              }} >
              <MenuItem>
                {Paths.activityAdditionalWriter}
              </MenuItem>
            </Link>
            <Link
              to={PathConfig.activityDatesWriter}
              onClick={() => {
                props.initialise()
                props.setKey({key: '', keyType: KeyTypes.ACTIVITYDATE})
              }} >
              <MenuItem>
                {Paths.activityDatesWriter}
              </MenuItem>
            </Link>
            <Link
              to={PathConfig.activityParticipatingOrgWriter}
              onClick={() => {
                props.initialise()
                props.setKey({key: '', keyType: KeyTypes.ACTIVITYPARTICIPATINGORG})
              }} >
              <MenuItem>
                {Paths.activityParticipatingOrgWriter}
              </MenuItem>
            </Link>
            <Link
              to={PathConfig.activitySectorsWriter}
              onClick={() => {
                props.initialise()
                props.setKey({key: '', keyType: KeyTypes.ACTIVITYSECTOR})
              }} >
              <MenuItem>
                {Paths.activitySectorsWriter}
              </MenuItem>
            </Link>
            <Link
              to={PathConfig.activityBudgetsWriter}
              onClick={() => {
                props.initialise()
                props.setKey({key: '', keyType: KeyTypes.ACTIVITYBUDGET})
              }} >
              <MenuItem>
                {Paths.activityBudgetsWriter}
              </MenuItem>
            </Link>
            <Link
              to={PathConfig.activityTerritoriesWriter}
              onClick={() => {
                props.initialise()
                props.setKey({key: '', keyType: KeyTypes.ACTIVITYTERRITORY})
              }} >
              <MenuItem>
                {Paths.activityTerritoriesWriter}
              </MenuItem>
            </Link>
            <Link
              to={PathConfig.activityTransactionsWriter}
              onClick={() => {
                props.initialise()
                props.setKey({key: '', keyType: KeyTypes.ACTIVITYTRANSACTION})
              }} >
              <MenuItem>
                {Paths.activityTransactionsWriter}
              </MenuItem>
            </Link>
            <Link
              to={PathConfig.activityRelatedActivityWriter}
              onClick={() => {
                props.initialise()
                props.setKey({key: '', keyType: KeyTypes.ACTIVITYRELATEDACTIVITY})
              }} >
              <MenuItem>
                {Paths.activityRelatedActivitiesWriter}
              </MenuItem>
            </Link>
          </MenuList>
        </ExpansionPanelDetails>
      </ExpansionPanel>

      <ExpansionPanel className={classes.siderMenu}>
        <ExpansionPanelSummary expandIcon={<ExpandMoreIcon />}>
          <IconButton aria-label={App.headingActivitiesUpdater}>
            <Update />
          </IconButton>
          <h4>{App.headingActivitiesUpdater}</h4>
        </ExpansionPanelSummary>
        <ExpansionPanelDetails className={classes.siderMenu}>
          <MenuList>
            <Link to={PathConfig.activitiesUpdater}>
              <MenuItem>
                {Paths.activitiesUpdater}
              </MenuItem>
            </Link>
            <Link to={PathConfig.activityUpdater}>
              <MenuItem>
                {Paths.activityUpdater}
              </MenuItem>
            </Link>
            <Link to={PathConfig.activityDateUpdater}>
              <MenuItem>
                {Paths.activityDateUpdater}
              </MenuItem>
            </Link>
          </MenuList>
        </ExpansionPanelDetails>
      </ExpansionPanel>

      <ExpansionPanel className={classes.siderMenu}>
        <ExpansionPanelSummary expandIcon={<ExpandMoreIcon />}>
          <IconButton aria-label={App.headingActivitiesReader}>
            <List />
          </IconButton>
          <h4>{App.headingActivitiesReader}</h4>
        </ExpansionPanelSummary>
        <ExpansionPanelDetails className={classes.siderMenu}>
          <MenuList>
            <Link to={PathConfig.activitiesReader}>
              <MenuItem>
                {Paths.activitiesReader}
              </MenuItem>
            </Link>
            <Link to={PathConfig.activityReader}>
              <MenuItem>
                {Paths.activityReader}
              </MenuItem>
            </Link>
            <Link to={PathConfig.activityAdditionalReader}>
              <MenuItem>
                {Paths.activityAdditionalReader}
              </MenuItem>
            </Link>
            <Link to={PathConfig.activityDatesReader}>
              <MenuItem>
                {Paths.activityDatesReader}
              </MenuItem>
            </Link>
            <Link to={PathConfig.activityParticipatingOrgReader}>
              <MenuItem>
                {Paths.activityParticipatingOrgReader}
              </MenuItem>
            </Link>
            <Link to={PathConfig.activitySectorsReader}>
              <MenuItem>
                {Paths.activitySectorsReader}
              </MenuItem>
            </Link>
            <Link to={PathConfig.activityBudgetsReader}>
              <MenuItem>
                {Paths.activityBudgetsReader}
              </MenuItem>
            </Link>
            <Link to={PathConfig.activityTerritoriesReader}>
              <MenuItem>
                {Paths.activityTerritoriesReader}
              </MenuItem>
            </Link>
            <Link to={PathConfig.activityTransactionsReader}>
              <MenuItem>
                {Paths.activityTransactionsReader}
              </MenuItem>
            </Link>
            <Link to={PathConfig.activityRelatedActivityReader}>
              <MenuItem>
                {Paths.activityRelatedActivitiesReader}
              </MenuItem>
            </Link>
          </MenuList>
        </ExpansionPanelDetails>
      </ExpansionPanel>

      <ExpansionPanel className={classes.siderMenu}>
        <ExpansionPanelSummary expandIcon={<ExpandMoreIcon />}>
          <IconButton aria-label={App.headingOrganisationsWriter}>
            <Create />
          </IconButton>
          <h4>{App.headingOrganisationsWriter}</h4>
        </ExpansionPanelSummary>
        <ExpansionPanelDetails className={classes.siderMenu}>
          <MenuList>
            <Link
              to={PathConfig.orgWriter}
              onClick={() => {
                props.initialise()
                props.setKey({key: '', keyType: KeyTypes.ORG})
              }} >
              <MenuItem>
                {Paths.orgWriter}
              </MenuItem>
            </Link>
            <Link
              to={PathConfig.organisationsWriter}
              onClick={() => {
                props.initialise()
                props.setKey({key: '', keyType: KeyTypes.ORGANISATIONS})
              }} >
              <MenuItem>
                {Paths.organisationsWriter}
              </MenuItem>
            </Link>
            <Link
              to={PathConfig.organisationWriter}
              onClick={() => {
                props.initialise()
                props.setKey({key: '', keyType: KeyTypes.ORGANISATION})
              }} >
              <MenuItem>
                {Paths.organisationWriter}
              </MenuItem>
            </Link>
            <Link
              to={PathConfig.organisationDocsWriter}
              onClick={() => {
                props.initialise()
                props.setKey({key: '', keyType: KeyTypes.ORGANISATIONDOC})
              }} >
              <MenuItem>
                {Paths.organisationDocsWriter}
              </MenuItem>
            </Link>
            <Link
              to={PathConfig.organisationBudgetsWriter}
              onClick={() => {
                props.initialise()
                props.setKey({key: '', keyType: KeyTypes.ORGANISATIONBUDGET})
              }} >
              <MenuItem>
                {Paths.organisationBudgetsWriter}
              </MenuItem>
            </Link>
            <Link
              to={PathConfig.organisationExpenditureWriter}
              onClick={() => {
                props.initialise()
                props.setKey({key: '', keyType: KeyTypes.ORGANISATIONEXPENDITURE})
              }} >
              <MenuItem>
                {Paths.organisationExpenditureWriter}
              </MenuItem>
            </Link>
            <Link
              to={PathConfig.organisationRecipientBudgetsWriter}
              onClick={() => {
                props.initialise()
                props.setKey({key: '', keyType: KeyTypes.ORGANISATIONRECIPIENTBUDGET})
              }} >
              <MenuItem>
                {Paths.organisationRecipientBudgetsWriter}
              </MenuItem>
            </Link>
            <Link
              to={PathConfig.organisationRegionBudgetsWriter}
              onClick={() => {
                props.initialise()
                props.setKey({key: '', keyType: KeyTypes.ORGANISATIONREGIONBUDGET})
              }} >
              <MenuItem>
                {Paths.organisationRegionBudgetsWriter}
              </MenuItem>
            </Link>
            <Link
              to={PathConfig.organisationCountryBudgetsWriter}
              onClick={() => {
                props.initialise()
                props.setKey({key: '', keyType: KeyTypes.ORGANISATIONCOUNTRYBUDGET})
              }} >
              <MenuItem>
                {Paths.organisationCountryBudgetsWriter}
              </MenuItem>
            </Link>
          </MenuList>
        </ExpansionPanelDetails>
      </ExpansionPanel>

      <ExpansionPanel className={classes.siderMenu}>
        <ExpansionPanelSummary expandIcon={<ExpandMoreIcon />}>
          <IconButton aria-label={App.headingOrganisationsUpdater}>
            <Update />
          </IconButton>
          <h4>{App.headingOrganisationsUpdater}</h4>
        </ExpansionPanelSummary>
        <ExpansionPanelDetails className={classes.siderMenu}>
          <MenuList>
            <Link to={PathConfig.orgUpdater}>
              <MenuItem>
                {Paths.orgUpdater}
              </MenuItem>
            </Link>
          </MenuList>
        </ExpansionPanelDetails>
      </ExpansionPanel>

      <ExpansionPanel className={classes.siderMenu}>
        <ExpansionPanelSummary expandIcon={<ExpandMoreIcon />}>
          <IconButton aria-label={App.headingOrganisationsReader}>
            <List />
          </IconButton>
          <h4>{App.headingOrganisationsReader}</h4>
        </ExpansionPanelSummary>
        <ExpansionPanelDetails className={classes.siderMenu}>
          <MenuList>
            <Link to={PathConfig.orgsReader}>
              <MenuItem>
                {Paths.orgsReader}
              </MenuItem>
            </Link>
            <Link to={PathConfig.organisationsReader}>
              <MenuItem>
                {Paths.organisationsReader}
              </MenuItem>
            </Link>
            <Link to={PathConfig.organisationReader}>
              <MenuItem>
                {Paths.organisationReader}
              </MenuItem>
            </Link>
            <Link to={PathConfig.organisationDocsReader}>
              <MenuItem>
                {Paths.organisationDocsReader}
              </MenuItem>
            </Link>
            <Link to={PathConfig.organisationBudgetsReader}>
              <MenuItem>
                {Paths.organisationBudgetsReader}
              </MenuItem>
            </Link>
            <Link to={PathConfig.organisationExpenditureReader}>
              <MenuItem>
                {Paths.organisationExpenditureReader}
              </MenuItem>
            </Link>
            <Link to={PathConfig.organisationRecipientBudgetsReader}>
              <MenuItem>
                {Paths.organisationRecipientBudgetsReader}
              </MenuItem>
            </Link>
            <Link to={PathConfig.organisationRegionBudgetsReader}>
              <MenuItem>
                {Paths.organisationRegionBudgetsReader}
              </MenuItem>
            </Link>
            <Link to={PathConfig.organisationCountryBudgetsReader}>
              <MenuItem>
                {Paths.organisationCountryBudgetsReader}
              </MenuItem>
            </Link>
          </MenuList>
        </ExpansionPanelDetails>
      </ExpansionPanel>
    </div>
  )
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): DispatchProps => {
  return {
    initialise: () => dispatch(initialise()),
    setKey: (keyProps: Keys) => dispatch(setKey(keyProps))
  }
}

export const SiderOrganisationMenu = connect<null, DispatchProps, {}, ApplicationState>(
  null,
  mapDispatchToProps
)(Sider)
