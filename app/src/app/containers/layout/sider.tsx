import * as React from 'react'
import { Link } from 'react-router-dom'
import Grid from '@material-ui/core/Grid'
import Menu from '@material-ui/core/Menu'
import Paper from '@material-ui/core/Paper'
import MenuItem from '@material-ui/core/MenuItem'

import { Paths } from './types'

interface SiderProps {
  handleClose: () => void
}

type  AllProps = SiderProps & Paths

const Sider: React.SFC<AllProps> = (props: AllProps) => {

  return (
    <Grid container spacing={24}>
      <Grid item xs={12}>
        <Paper>
          <Menu
            id="simple-menu"
            anchorEl={null}
            open={Boolean(true)}
            onClose={props.handleClose}
          >
            <MenuItem>
              <Link to={props.homePath}/><span>{props.home}</span>
            </MenuItem>
            <MenuItem>
              <Link to={props.aboutPath}/><span>{props.about}</span>
            </MenuItem>
            <MenuItem>
              <Link to={props.overviewPath}/><span>{props.overview}</span>
            </MenuItem>
            <MenuItem>
              <Link to={props.helpPath}/><span>{props.help}</span>
            </MenuItem>
          </Menu>
        </Paper>
      </Grid>
    </Grid>
  )
}

export default Sider
