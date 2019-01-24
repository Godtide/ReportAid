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

class Sider extends React.Component<WithStyles<typeof styles>> {

  render() {

    return (
      <div>

        <h3>{Organisation.headingOrgWriter}</h3>
        <MenuList>
          <Link to={PathConfig.orgWriter}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.orgWriter}>
                <Create />
              </IconButton>
              {Paths.orgWriter}
            </MenuItem>
          </Link>
          <Link to={PathConfig.orgReportWriter}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.orgReportWriter}>
                <Create />
              </IconButton>
              {Paths.orgReportWriter}
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
        </MenuList>

        <hr />

        <h3>{Organisation.headingOrgReader}</h3>
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
        </MenuList>
      </div>
    )
  }
}

export const SiderOrganisationMenu = withTheme(withStyles(styles)(Sider))
