import * as React from 'react'
import { Link } from 'react-router-dom'
import MenuList from '@material-ui/core/MenuList'
import MenuItem from '@material-ui/core/MenuItem'

import IconButton from '@material-ui/core/IconButton'
import ExpansionPanel from '@material-ui/core/ExpansionPanel'
import ExpansionPanelSummary from '@material-ui/core/ExpansionPanelSummary'
import ExpandMoreIcon from '@material-ui/icons/ExpandMore'
import Home from '@material-ui/icons/Home'
import Info from '@material-ui/icons/Info'
import Help from '@material-ui/icons/Help'
import Panorama from '@material-ui/icons/Panorama'
import Create from '@material-ui/icons/Create'
import Report from '@material-ui/icons/Report'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../styles/theme'

import { Paths } from '../utils/strings'
import { Paths as PathConfig } from '../utils/config'

class ApplicationMenu extends React.Component<WithStyles<typeof styles>> {

  render() {

    return (
      <div>
        <MenuList>
          <Link to={PathConfig.home}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.home}>
                <Home />
              </IconButton>
              {Paths.home}
            </MenuItem>
          </Link>
          <Link to={PathConfig.blockchain}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.blockchain}>
                <Home />
              </IconButton>
              {Paths.blockchain}
            </MenuItem>
          </Link>
          <Link to={PathConfig.about}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.about}>
                <Info />
              </IconButton>
              {Paths.about}
            </MenuItem>
          </Link>
          <Link to={PathConfig.overview}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.overview}>
                <Panorama />
              </IconButton>
              {Paths.overview}
            </MenuItem>
          </Link>
          <Link to={PathConfig.help}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.help}>
                <Help />
              </IconButton>
              {Paths.help}
            </MenuItem>
          </Link>
          <Link to={PathConfig.writer}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.writer}>
                <Create />
              </IconButton>
              {Paths.writer}
            </MenuItem>
          </Link>
          <ExpansionPanel className={this.props.classes.expansion} >
            <ExpansionPanelSummary className={this.props.classes.expansion} expandIcon={<ExpandMoreIcon />}>
              <p>{Paths.writer}</p>
            </ExpansionPanelSummary>
            <MenuList>
              <Link to={PathConfig.orgWriter}>
                <MenuItem>
                  <IconButton className={this.props.classes.button} aria-label={Paths.orgWriter}>
                    <Create />
                  </IconButton>
                  {Paths.orgWriter}
                </MenuItem>
              </Link>
            </MenuList>
          </ExpansionPanel>
          <Link to={PathConfig.reader}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.reader}>
                <Report />
              </IconButton>
              {Paths.reader}
            </MenuItem>
          </Link>
        </MenuList>
      </div>
    )
  }
}

export const AppMenu = withTheme(withStyles(styles)(ApplicationMenu))
