import * as React from 'react'
import { Link } from 'react-router-dom'

import Grid from '@material-ui/core/Grid'
import AppBar from '@material-ui/core/AppBar'
import Toolbar from '@material-ui/core/Toolbar'
import MenuItem from '@material-ui/core/MenuItem'

import IconButton from '@material-ui/core/IconButton'
import ListAlt from '@material-ui/icons/ListAlt'
import List from '@material-ui/icons/List'
import Home from '@material-ui/icons/Home'
import Info from '@material-ui/icons/Info'
import Help from '@material-ui/icons/Help'
import Panorama from '@material-ui/icons/Panorama'
import Create from '@material-ui/icons/Create'

import { useStyles } from '../styles/theme';

import { Paths } from '../utils/strings'
import { Paths as PathConfig } from '../utils/config'

export const ApplicationBar = () => {

  const classes = useStyles()

  return (

    <AppBar position='relative' >
      <Toolbar variant='regular' >
        <Link to={PathConfig.home}>
          <MenuItem>
            <Grid container direction="column" alignItems="center">
              <Grid item>
                <IconButton className={classes.button} aria-label={Paths.home}>
                  <Home />
                </IconButton>
              </Grid>
              <Grid item>
                {Paths.home}
              </Grid>
            </Grid>
          </MenuItem>
        </Link>
        <Link to={PathConfig.blockchain}>
          <MenuItem>
            <Grid container direction="column" alignItems="center">
              <Grid item>
                <IconButton className={classes.button} aria-label={Paths.blockchain}>
                  <ListAlt />
                </IconButton>
              </Grid>
              <Grid item>
                {Paths.blockchain}
              </Grid>
            </Grid>
          </MenuItem>
        </Link>
        <Link to={PathConfig.writer}>
          <MenuItem>
            <Grid container direction="column" alignItems="center">
              <Grid item>
                <IconButton className={classes.button} aria-label={Paths.writer}>
                  <Create />
                </IconButton>
              </Grid>
              <Grid item>
                {Paths.writer}
              </Grid>
            </Grid>
          </MenuItem>
        </Link>
        <Link to={PathConfig.reader}>
          <MenuItem>
            <Grid container direction="column" alignItems="center">
              <Grid item>
                <IconButton className={classes.button} aria-label={Paths.reader}>
                  <List />
                </IconButton>
              </Grid>
              <Grid item>
                {Paths.reader}
              </Grid>
            </Grid>
          </MenuItem>
        </Link>
        <Link to={PathConfig.overview}>
          <MenuItem>
            <Grid container direction="column" alignItems="center">
              <Grid item>
                <IconButton className={classes.button} aria-label={Paths.overview}>
                  <Panorama />
                </IconButton>
              </Grid>
              <Grid item>
                {Paths.overview}
              </Grid>
            </Grid>
          </MenuItem>
        </Link>
        <Link to={PathConfig.help}>
          <MenuItem>
            <Grid container direction="column" alignItems="center">
              <Grid item>
                <IconButton className={classes.button} aria-label={Paths.help}>
                  <Help />
                </IconButton>
              </Grid>
              <Grid item>
                {Paths.help}
              </Grid>
            </Grid>
          </MenuItem>
        </Link>
        <Link to={PathConfig.about}>
          <MenuItem>
            <Grid container direction="column" alignItems="center">
              <Grid item>
                <IconButton className={classes.button} aria-label={Paths.about}>
                  <Info />
                </IconButton>
              </Grid>
              <Grid item>
                {Paths.about}
              </Grid>
            </Grid>
          </MenuItem>
        </Link>
      </Toolbar>
    </AppBar>
  )
}
