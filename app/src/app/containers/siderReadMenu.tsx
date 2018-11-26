import * as React from 'react'
import { Link } from 'react-router-dom'
import MenuList from '@material-ui/core/MenuList'
import MenuItem from '@material-ui/core/MenuItem'

import IconButton from '@material-ui/core/IconButton'
import List from '@material-ui/icons/List'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../styles/theme'

import { Paths } from '../utils/strings'
import { Paths as PathConfig } from '../utils/config'

class Sider extends React.Component<WithStyles<typeof styles>> {

  render() {

    return (
      <div>
        <MenuList>
          <Link to={PathConfig.orgReader}>
            <MenuItem>
              <IconButton className={this.props.classes.button} aria-label={Paths.orgReader}>
                <List />
              </IconButton>
              {Paths.orgReader}
            </MenuItem>
          </Link>
        </MenuList>
      </div>
    )
  }
}

export const SiderReadMenu = withTheme(withStyles(styles)(Sider))
