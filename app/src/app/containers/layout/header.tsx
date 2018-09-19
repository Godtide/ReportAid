import * as React from 'react'
//import { Link } from 'react-router-dom'
import Paper from '@material-ui/core/Paper'
import Grid from '@material-ui/core/Grid'
//import Menu from '@material-ui/core/Menu'
//import MenuItem from '@material-ui/core/MenuItem'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../styles/theme'

import { Paths } from './types'

interface HeaderProps {
  appTitle: string
  handleClose: () => void
}

export type  AllProps = HeaderProps & Paths

class Header extends React.Component<WithStyles<typeof styles> & AllProps> {

    public render() {
      return (
        <div className={this.props.classes.root}>
          <Grid container spacing={8}>
            <Grid item xs={2}>
               <Paper><p>logo to go here</p></Paper>
            </Grid>
            <Grid item xs={10}>
              <Paper><h1>{this.props.appTitle}</h1></Paper>
            </Grid>
          </Grid>
        </div>
      )
    }
}

export default withTheme(withStyles(styles)(Header))

/*
<Menu
    id="simple-menu"
    anchorEl={null}
    open={Boolean(true)}
    onClose={props.handleClose}
  >
    <MenuItem onClick={props.handleClose}>{props.home}</MenuItem>
    <MenuItem onClick={props.handleClose}>{props.about}</MenuItem>
    <MenuItem onClick={props.handleClose}>{props.overview}</MenuItem>
    <MenuItem onClick={props.handleClose}>{props.help}</MenuItem>
  </Menu>
  */

/* <Menu theme="dark" mode="horizontal" defaultSelectedKeys={['home']} style={{ lineHeight: '64px' }} >
  <Menu.Item key={props.home}>
    <Icon type={props.homeIcon}/><span>{props.home}</span>
    <Link to={props.homePath}/>
  </Menu.Item>
  <Menu.Item key={props.about}>
    <Icon type={props.aboutIcon}/><span>{props.about}</span>
    <Link to={props.aboutPath}/>
  </Menu.Item>
  <Menu.Item key={props.overview}>
    <Icon type={props.overviewIcon}/><span>{props.overview}</span>
    <Link to={props.overviewPath}/>
  </Menu.Item>
  <Menu.Item key={props.help}>
    <Icon type={props.helpIcon}/><span>{props.help}</span>
    <Link to={props.helpPath}/>
  </Menu.Item> */
