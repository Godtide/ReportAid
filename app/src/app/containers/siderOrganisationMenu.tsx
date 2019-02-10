import * as React from 'react'
import { Link } from 'react-router-dom'
import MenuList from '@material-ui/core/MenuList'
import MenuItem from '@material-ui/core/MenuItem'

import IconButton from '@material-ui/core/IconButton'
import Create from '@material-ui/icons/Create'
import List from '@material-ui/icons/List'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../styles/theme'

import { Paths, Organisation } from '../utils/strings'
import { Paths as PathConfig } from '../utils/config'

/*
static readonly orgWriter='/create-organisation'
static readonly organisationsWriter='/create-organisations-report'
static readonly organisationWriter='/create-organisationsreport'
static readonly organisationDocsWriter = '/create-organisation-report-docs'
static readonly organisationBudgetsWriter = '/create-organisation-report-budget'
static readonly organisationExpenditureWriter = '/create-organisation-report-expenditure'
static readonly organisationRecipientBudgetsWriter = '/create-organisation-report-recipient-budget'
static readonly organisationRegionBudgetsWriter = '/create-organisation-report-region-budget'
static readonly organisationCountryBudgetsWriter = '/create-organisation-report-country-budget'
*/

class Sider extends React.Component<WithStyles<typeof styles>> {

  render() {

    return (
      <div>

        <h3>{Organisation.headingOrganisationWriter}</h3>
        <MenuList>
          <Link to={PathConfig.orgWriter}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.orgWriter}>
                <Create />
              </IconButton>
              {Paths.orgWriter}
            </MenuItem>
          </Link>

          /*<Link to={PathConfig.orgReportWriter}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.orgReportWriter}>
                <Create />
              </IconButton>
              {Paths.orgReportWriter}
            </MenuItem>
          </Link>
          <Link to={PathConfig.orgReportDocsWriter}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.orgReportDocsWriter}>
                <Create />
              </IconButton>
              {Paths.orgReportDocsWriter}
            </MenuItem>
          </Link>
          <Link to={PathConfig.orgReportBudgetsWriter}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.orgReportBudgetsWriter}>
                <Create />
              </IconButton>
              {Paths.orgReportBudgetsWriter}
            </MenuItem>
          </Link>
          <Link to={PathConfig.orgReportExpenditureWriter}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.orgReportExpenditureWriter}>
                <Create />
              </IconButton>
              {Paths.orgReportExpenditureWriter}
            </MenuItem>
          </Link>
          <Link to={PathConfig.orgReportRecipientBudgetsWriter}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.orgReportRecipientBudgetsWriter}>
                <Create />
              </IconButton>
              {Paths.orgReportRecipientBudgetsWriter}
            </MenuItem>
          </Link>
          <Link to={PathConfig.orgReportRegionBudgetsWriter}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.orgReportRegionBudgetsWriter}>
                <Create />
              </IconButton>
              {Paths.orgReportRegionBudgetsWriter}
            </MenuItem>
          </Link>
          <Link to={PathConfig.orgReportCountryBudgetsWriter}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.orgReportCountryBudgetsWriter}>
                <Create />
              </IconButton>
              {Paths.orgReportCountryBudgetsWriter}
            </MenuItem>
          </Link>
          */
        </MenuList>

        /*
        <hr />

        <h3>{Organisation.headingOrganisationReader}</h3>
        <MenuList>
          <Link to={PathConfig.orgReader}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.orgReader}>
                <List />
              </IconButton>
              {Paths.orgReader}
            </MenuItem>
          </Link>
          <Link to={PathConfig.orgReportReader}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.orgReportReader}>
                <List />
              </IconButton>
              {Paths.orgReportReader}
            </MenuItem>
          </Link>
          <Link to={PathConfig.orgReportDocsReader}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.orgReportDocsReader}>
                <List />
              </IconButton>
              {Paths.orgReportDocsReader}
            </MenuItem>
          </Link>
          <Link to={PathConfig.orgReportBudgetsReader}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.orgReportBudgetsReader}>
                <List />
              </IconButton>
              {Paths.orgReportBudgetsReader}
            </MenuItem>
          </Link>
          <Link to={PathConfig.orgReportExpenditureReader}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.orgReportExpenditureReader}>
                <List />
              </IconButton>
              {Paths.orgReportExpenditureReader}
            </MenuItem>
          </Link>
          <Link to={PathConfig.orgReportRecipientBudgetsReader}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.orgReportRecipientBudgetsReader}>
                <List />
              </IconButton>
              {Paths.orgReportRecipientBudgetsReader}
            </MenuItem>
          </Link>
          <Link to={PathConfig.orgReportRegionBudgetsReader}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.orgReportRegionBudgetsReader}>
                <List />
              </IconButton>
              {Paths.orgReportRegionBudgetsReader}
            </MenuItem>
          </Link>
          <Link to={PathConfig.orgReportCountryBudgetsReader}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.orgReportCountryBudgetsReader}>
                <List />
              </IconButton>
              {Paths.orgReportCountryBudgetsReader}
            </MenuItem>
          </Link>
        </MenuList>
        */
      </div>
    )
  }
}

export const SiderOrganisationMenu = withTheme(withStyles(styles)(Sider))
