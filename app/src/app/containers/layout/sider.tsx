import * as React from 'react'
import { Link } from 'react-router-dom'
import Grid from '@material-ui/core/Grid'
import Menu from '@material-ui/core/Menu'
import Paper from '@material-ui/core/Paper'
import MenuItem from '@material-ui/core/MenuItem'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../styles/theme'

import { Paths } from './types'

interface SiderProps {
  handleClose: () => void
}

type  AllProps = SiderProps & Paths

class Sider extends React.Component<WithStyles<typeof styles> & AllProps> {

  public render() {

    return (
      <React.Fragment>
        <Grid item xs={12} sm={2}>
          <div className={this.props.classes.root}>
            <Menu
              id="simple-menu"
              anchorEl={null}
              open={Boolean(true)}
              onClose={this.props.handleClose}
            >
              <MenuItem>
                <Link to={this.props.homePath}/><span>{this.props.home}</span>
              </MenuItem>
              <MenuItem>
                <Link to={this.props.aboutPath}/><span>{this.props.about}</span>
              </MenuItem>
              <MenuItem>
                <Link to={this.props.overviewPath}/><span>{this.props.overview}</span>
              </MenuItem>
              <MenuItem>
                <Link to={this.props.helpPath}/><span>{this.props.help}</span>
              </MenuItem>
            </Menu>
          </div>
        </Grid>
      </React.Fragment>
    )
  }
}

export default withTheme(withStyles(styles)(Sider))
