import * as React from 'react'
import { Link } from 'react-router-dom'
import MenuList from '@material-ui/core/MenuList'
import MenuItem from '@material-ui/core/MenuItem'

import IconButton from '@material-ui/core/IconButton'
import Home from '@material-ui/icons/Home'
import Info from '@material-ui/icons/Info'
import Help from '@material-ui/icons/Help'
import Panorama from '@material-ui/icons/Panorama'
import Create from '@material-ui/icons/Create'
import Report from '@material-ui/icons/Report'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../styles/theme'

import { PathStrings } from '../utils/strings'

class AppMenu extends React.Component<WithStyles<typeof styles>> {

  public render() {

    return (
      <div>
        <MenuList>
          <Link to={PathStrings.homePath}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={PathStrings.home}>
                <Home />
              </IconButton>
              {PathStrings.home}
            </MenuItem>
          </Link>
          <Link to={PathStrings.blockchainPath}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={PathStrings.blockchain}>
                <Home />
              </IconButton>
              {PathStrings.blockchain}
            </MenuItem>
          </Link>
          <Link to={PathStrings.aboutPath}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={PathStrings.about}>
                <Info />
              </IconButton>
              {PathStrings.about}
            </MenuItem>
          </Link>
          <Link to={PathStrings.overviewPath}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={PathStrings.overview}>
                <Panorama />
              </IconButton>
              {PathStrings.overview}
            </MenuItem>
          </Link>
          <Link to={PathStrings.helpPath}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={PathStrings.help}>
                <Help />
              </IconButton>
              {PathStrings.help}
            </MenuItem>
          </Link>
          <Link to={PathStrings.writerPath}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={PathStrings.writer}>
                <Create />
              </IconButton>
              {PathStrings.writer}
            </MenuItem>
          </Link>
          <Link to={PathStrings.readerPath}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={PathStrings.reader}>
                <Report />
              </IconButton>
              {PathStrings.reader}
            </MenuItem>
          </Link>
        </MenuList>
      </div>
    )
  }
}

export default withTheme(withStyles(styles)(AppMenu))
