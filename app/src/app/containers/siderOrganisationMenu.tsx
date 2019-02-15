import * as React from 'react'
import { Link } from 'react-router-dom'
import MenuList from '@material-ui/core/MenuList'
import MenuItem from '@material-ui/core/MenuItem'

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

class Sider extends React.Component<WithStyles<typeof styles>> {

  render() {

    return (
      <div>
        <ExpansionPanel className={this.props.classes.siderMenu}>
          <ExpansionPanelSummary expandIcon={<ExpandMoreIcon />}>
            <h3>{App.headingOrganisationsWriter}</h3>
          </ExpansionPanelSummary>
          <ExpansionPanelDetails>
            <MenuList>
              <Link to={PathConfig.orgWriter}>
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.orgWriter}>
                    <Create />
                  </IconButton>
                  {Paths.orgWriter}
                </MenuItem>
              </Link>
              <Link to={PathConfig.organisationsWriter}>
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.organisationsWriter}>
                    <Create />
                  </IconButton>
                  {Paths.organisationsWriter}
                </MenuItem>
              </Link>
              <Link to={PathConfig.organisationWriter}>
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.organisationWriter}>
                    <Create />
                  </IconButton>
                  {Paths.organisationWriter}
                </MenuItem>
              </Link>
              <Link to={PathConfig.organisationDocsWriter}>
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.organisationDocsWriter}>
                    <Create />
                  </IconButton>
                  {Paths.organisationDocsWriter}
                </MenuItem>
              </Link>
              <Link to={PathConfig.organisationBudgetsWriter}>
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.organisationBudgetsWriter}>
                    <Create />
                  </IconButton>
                  {Paths.organisationBudgetsWriter}
                </MenuItem>
              </Link>
              <Link to={PathConfig.organisationExpenditureWriter}>
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.organisationExpenditureWriter}>
                    <Create />
                  </IconButton>
                  {Paths.organisationExpenditureWriter}
                </MenuItem>
              </Link>
              <Link to={PathConfig.organisationRecipientBudgetsWriter}>
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.organisationRecipientBudgetsWriter}>
                    <Create />
                  </IconButton>
                  {Paths.organisationRecipientBudgetsWriter}
                </MenuItem>
              </Link>
              <Link to={PathConfig.organisationRegionBudgetsWriter}>
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.organisationRegionBudgetsWriter}>
                    <Create />
                  </IconButton>
                  {Paths.organisationRegionBudgetsWriter}
                </MenuItem>
              </Link>
              <Link to={PathConfig.organisationCountryBudgetsWriter}>
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
            <h3>{App.headingOrganisationsReader}</h3>
          </ExpansionPanelSummary>
          <ExpansionPanelDetails>
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

export const SiderOrganisationMenu = withTheme(withStyles(styles)(Sider))
