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
import Create from '@material-ui/icons/Create'
import List from '@material-ui/icons/List'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../styles/theme'

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

class Sider extends React.Component<WithStyles<typeof styles> & DispatchProps> {

  render() {

    return (
      <div>
        <ExpansionPanel className={this.props.classes.siderMenu}>
          <ExpansionPanelSummary expandIcon={<ExpandMoreIcon />}>
            <h4>{App.headingActivitiesWriter}</h4>
          </ExpansionPanelSummary>
          <ExpansionPanelDetails className={this.props.classes.siderMenu}>
            <MenuList>
              <Link
                to={PathConfig.activitiesWriter}
                onClick={() => {
                  this.props.initialise()
                  this.props.setKey({key: '', keyType: KeyTypes.ACTIVITIES})
                }} >
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.activitiesWriter}>
                    <Create />
                  </IconButton>
                  {Paths.activitiesWriter}
                </MenuItem>
              </Link>
              <Link
                to={PathConfig.activityWriter}
                onClick={() => {
                  this.props.initialise()
                  this.props.setKey({key: '', keyType: KeyTypes.ACTIVITY})
                }} >
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.activityWriter}>
                    <Create />
                  </IconButton>
                  {Paths.activityWriter}
                </MenuItem>
              </Link>
              <Link
                to={PathConfig.activityAdditionalWriter}
                onClick={() => {
                  this.props.initialise()
                  this.props.setKey({key: '', keyType: KeyTypes.ACTIVITYADDITIONAL})
                }} >
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.activityAdditionalWriter}>
                    <Create />
                  </IconButton>
                  {Paths.activityAdditionalWriter}
                </MenuItem>
              </Link>
              <Link
                to={PathConfig.activityDatesWriter}
                onClick={() => {
                  this.props.initialise()
                  this.props.setKey({key: '', keyType: KeyTypes.ACTIVITYDATE})
                }} >
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.activityDatesWriter}>
                    <Create />
                  </IconButton>
                  {Paths.activityDatesWriter}
                </MenuItem>
              </Link>
            </MenuList>
          </ExpansionPanelDetails>
        </ExpansionPanel>

        <ExpansionPanel className={this.props.classes.siderMenu}>
          <ExpansionPanelSummary expandIcon={<ExpandMoreIcon />}>
            <h4>{App.headingActivitiesUpdater}</h4>
          </ExpansionPanelSummary>
          <ExpansionPanelDetails className={this.props.classes.siderMenu}>
            <MenuList>
              <Link to={PathConfig.activitiesUpdater}>
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.activitiesUpdater}>
                    <Create />
                  </IconButton>
                  {Paths.activitiesUpdater}
                </MenuItem>
              </Link>
              <Link to={PathConfig.activityUpdater}>
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.activityUpdater}>
                    <Create />
                  </IconButton>
                  {Paths.activityUpdater}
                </MenuItem>
              </Link>
              <Link to={PathConfig.activityDateUpdater}>
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.activityDateUpdater}>
                    <Create />
                  </IconButton>
                  {Paths.activityDateUpdater}
                </MenuItem>
              </Link>
            </MenuList>
          </ExpansionPanelDetails>
        </ExpansionPanel>

        <ExpansionPanel className={this.props.classes.siderMenu}>
          <ExpansionPanelSummary expandIcon={<ExpandMoreIcon />}>
            <h4>{App.headingActivitiesReader}</h4>
          </ExpansionPanelSummary>
          <ExpansionPanelDetails className={this.props.classes.siderMenu}>
            <MenuList>
              <Link to={PathConfig.activitiesReader}>
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.activitiesReader}>
                    <List />
                  </IconButton>
                  {Paths.activitiesReader}
                </MenuItem>
              </Link>
              <Link to={PathConfig.activityReader}>
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.activityReader}>
                    <List />
                  </IconButton>
                  {Paths.activityReader}
                </MenuItem>
              </Link>
              <Link to={PathConfig.activityAdditionalReader}>
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.activityAdditionalReader}>
                    <List />
                  </IconButton>
                  {Paths.activityAdditionalReader}
                </MenuItem>
              </Link>
              <Link to={PathConfig.activityDatesReader}>
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.activityDatesReader}>
                    <List />
                  </IconButton>
                  {Paths.activityDatesReader}
                </MenuItem>
              </Link>
            </MenuList>
          </ExpansionPanelDetails>
        </ExpansionPanel>

        <ExpansionPanel className={this.props.classes.siderMenu}>
          <ExpansionPanelSummary expandIcon={<ExpandMoreIcon />}>
            <h4>{App.headingOrganisationsWriter}</h4>
          </ExpansionPanelSummary>
          <ExpansionPanelDetails className={this.props.classes.siderMenu}>
            <MenuList>
              <Link
                to={PathConfig.orgWriter}
                onClick={() => {
                  this.props.initialise()
                  this.props.setKey({key: '', keyType: KeyTypes.ORG})
                }} >
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.orgWriter}>
                    <Create />
                  </IconButton>
                  {Paths.orgWriter}
                </MenuItem>
              </Link>
              <Link
                to={PathConfig.organisationsWriter}
                onClick={() => {
                  this.props.initialise()
                  this.props.setKey({key: '', keyType: KeyTypes.ORGANISATIONS})
                }} >
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.organisationsWriter}>
                    <Create />
                  </IconButton>
                  {Paths.organisationsWriter}
                </MenuItem>
              </Link>
              <Link
                to={PathConfig.organisationWriter}
                onClick={() => {
                  this.props.initialise()
                  this.props.setKey({key: '', keyType: KeyTypes.ORGANISATION})
                }} >
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.organisationWriter}>
                    <Create />
                  </IconButton>
                  {Paths.organisationWriter}
                </MenuItem>
              </Link>
              <Link
                to={PathConfig.organisationDocsWriter}
                onClick={() => {
                  this.props.initialise()
                  this.props.setKey({key: '', keyType: KeyTypes.ORGANISATIONDOC})
                }} >
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.organisationDocsWriter}>
                    <Create />
                  </IconButton>
                  {Paths.organisationDocsWriter}
                </MenuItem>
              </Link>
              <Link
                to={PathConfig.organisationBudgetsWriter}
                onClick={() => {
                  this.props.initialise()
                  this.props.setKey({key: '', keyType: KeyTypes.ORGANISATIONBUDGET})
                }} >
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.organisationBudgetsWriter}>
                    <Create />
                  </IconButton>
                  {Paths.organisationBudgetsWriter}
                </MenuItem>
              </Link>
              <Link
                to={PathConfig.organisationExpenditureWriter}
                onClick={() => {
                  this.props.initialise()
                  this.props.setKey({key: '', keyType: KeyTypes.ORGANISATIONEXPENDITURE})
                }} >
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.organisationExpenditureWriter}>
                    <Create />
                  </IconButton>
                  {Paths.organisationExpenditureWriter}
                </MenuItem>
              </Link>
              <Link
                to={PathConfig.organisationRecipientBudgetsWriter}
                onClick={() => {
                  this.props.initialise()
                  this.props.setKey({key: '', keyType: KeyTypes.ORGANISATIONRECIPIENTBUDGET})
                }} >
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.organisationRecipientBudgetsWriter}>
                    <Create />
                  </IconButton>
                  {Paths.organisationRecipientBudgetsWriter}
                </MenuItem>
              </Link>
              <Link
                to={PathConfig.organisationRegionBudgetsWriter}
                onClick={() => {
                  this.props.initialise()
                  this.props.setKey({key: '', keyType: KeyTypes.ORGANISATIONREGIONBUDGET})
                }} >
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.organisationRegionBudgetsWriter}>
                    <Create />
                  </IconButton>
                  {Paths.organisationRegionBudgetsWriter}
                </MenuItem>
              </Link>
              <Link
                to={PathConfig.organisationCountryBudgetsWriter}
                onClick={() => {
                  this.props.initialise()
                  this.props.setKey({key: '', keyType: KeyTypes.ORGANISATIONCOUNTRYBUDGET})
                }} >
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.organisationCountryBudgetsWriter}>
                    <Create />
                  </IconButton>
                  {Paths.organisationCountryBudgetsWriter}
                </MenuItem>
              </Link>
            </MenuList>
          </ExpansionPanelDetails>
        </ExpansionPanel>

        <ExpansionPanel className={this.props.classes.siderMenu}>
          <ExpansionPanelSummary expandIcon={<ExpandMoreIcon />}>
            <h4>{App.headingOrganisationsUpdater}</h4>
          </ExpansionPanelSummary>
          <ExpansionPanelDetails className={this.props.classes.siderMenu}>
            <MenuList>
              <Link to={PathConfig.orgUpdater}>
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.orgUpdater}>
                    <Create />
                  </IconButton>
                  {Paths.orgUpdater}
                </MenuItem>
              </Link>
            </MenuList>
          </ExpansionPanelDetails>
        </ExpansionPanel>

        <ExpansionPanel className={this.props.classes.siderMenu}>
          <ExpansionPanelSummary expandIcon={<ExpandMoreIcon />}>
            <h4>{App.headingOrganisationsReader}</h4>
          </ExpansionPanelSummary>
          <ExpansionPanelDetails className={this.props.classes.siderMenu}>
            <MenuList>
              <Link to={PathConfig.orgsReader}>
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.orgsReader}>
                    <List />
                  </IconButton>
                  {Paths.orgsReader}
                </MenuItem>
              </Link>
              <Link to={PathConfig.organisationsReader}>
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.organisationsReader}>
                    <List />
                  </IconButton>
                  {Paths.organisationsReader}
                </MenuItem>
              </Link>
              <Link to={PathConfig.organisationReader}>
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.organisationReader}>
                    <List />
                  </IconButton>
                  {Paths.organisationReader}
                </MenuItem>
              </Link>
              <Link to={PathConfig.organisationDocsReader}>
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.organisationDocsReader}>
                    <List />
                  </IconButton>
                  {Paths.organisationDocsReader}
                </MenuItem>
              </Link>
              <Link to={PathConfig.organisationBudgetsReader}>
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.organisationBudgetsReader}>
                    <List />
                  </IconButton>
                  {Paths.organisationBudgetsReader}
                </MenuItem>
              </Link>
              <Link to={PathConfig.organisationExpenditureReader}>
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.organisationExpenditureReader}>
                    <List />
                  </IconButton>
                  {Paths.organisationExpenditureReader}
                </MenuItem>
              </Link>
              <Link to={PathConfig.organisationRecipientBudgetsReader}>
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.organisationRecipientBudgetsReader}>
                    <List />
                  </IconButton>
                  {Paths.organisationRecipientBudgetsReader}
                </MenuItem>
              </Link>
              <Link to={PathConfig.organisationRegionBudgetsReader}>
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.organisationRegionBudgetsReader}>
                    <List />
                  </IconButton>
                  {Paths.organisationRegionBudgetsReader}
                </MenuItem>
              </Link>
              <Link to={PathConfig.organisationCountryBudgetsReader}>
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.organisationCountryBudgetsReader}>
                    <List />
                  </IconButton>
                  {Paths.organisationCountryBudgetsReader}
                </MenuItem>
              </Link>
            </MenuList>
          </ExpansionPanelDetails>
        </ExpansionPanel>
      </div>
    )
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): DispatchProps => {
  return {
    initialise: () => dispatch(initialise()),
    setKey: (keyProps: Keys) => dispatch(setKey(keyProps))
  }
}

export const SiderOrganisationMenu = withTheme(withStyles(styles)(connect<null, DispatchProps, {}, ApplicationState>(
  null,
  mapDispatchToProps
)(Sider)))
