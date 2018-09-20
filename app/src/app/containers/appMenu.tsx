import * as React from 'react'
import { Link } from 'react-router-dom'
import MenuList from '@material-ui/core/MenuList'
import MenuItem from '@material-ui/core/MenuItem'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../styles/theme'

import { PathStrings } from '../utils/strings'

class AppMenu extends React.Component<WithStyles<typeof styles>> {

  public render() {

    return (
      <MenuList>
        <Link to={PathStrings.homePath}>
          <MenuItem>
            {PathStrings.home}
          </MenuItem>
        </Link>
        <Link to={PathStrings.aboutPath}>
          <MenuItem>
            {PathStrings.about}
          </MenuItem>
        </Link>
        <Link to={PathStrings.overviewPath}>
          <MenuItem>
            {PathStrings.overview}
          </MenuItem>
        </Link>
        <Link to={PathStrings.helpPath}>
          <MenuItem>
            {PathStrings.help}
          </MenuItem>
        </Link>
        <Link to={PathStrings.writerPath}>
          <MenuItem>
            {PathStrings.writer}
          </MenuItem>
        </Link>
        <Link to={PathStrings.readerPath}>
          <MenuItem>
            {PathStrings.reader}
          </MenuItem>
        </Link>
      </MenuList>
    )
  }
}

export default withTheme(withStyles(styles)(AppMenu))
