import * as React from 'react'
//import { Link } from 'react-router-dom'
import Paper from '@material-ui/core/Paper'
import Grid from '@material-ui/core/Grid'
import Menu from '@material-ui/core/Menu'
import MenuItem from '@material-ui/core/MenuItem'

import { PathStrings } from '../../utils/strings'

interface HeaderInterface {
  headerTitle: string
  home: string
  about: string
  overview: string
  help: string
  handleClose: () => void
}

export type  HeaderProps = PathStrings & HeaderInterface

const Header: React.SFC<HeaderProps> = ( props ) => {

  return (
    <div>
      <Grid container spacing={24}>
        <Grid item xs={4}>
          <Paper className={'blah'}><p>Logo to go here</p></Paper>
        </Grid>
        <Grid item xs={10}>

          <Menu
            id="simple-menu"
            anchorEl={null}
            open={Boolean(null)}
            onClose={props.handleClose}
          >
            <MenuItem onClick={props.handleClose}>{props.home}</MenuItem>
            <MenuItem onClick={props.handleClose}>{props.about}</MenuItem>
            <MenuItem onClick={props.handleClose}>{props.overview}</MenuItem>
            <MenuItem onClick={props.handleClose}>{props.help}</MenuItem>
          </Menu>
        </Grid>
        <Grid item xs={10}>
          <Paper className={'blah'}><h5>{props.headerTitle}</h5></Paper>
        </Grid>
      </Grid>
    </div>
  )
}

export default Header

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
