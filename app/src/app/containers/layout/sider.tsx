import * as React from 'react'
// import { Link } from 'react-router-dom'
import { withStyles } from '@material-ui/core/styles'
import Grid from '@material-ui/core/Grid'
import Menu from '@material-ui/core/Menu'
import MenuItem from '@material-ui/core/MenuItem'

import { Styles, ClassNames } from '../../styles/theme'

import { Paths } from './types'

interface SiderProps {
  handleClose: () => void
}

type  AllProps = SiderProps & Paths & ClassNames

const Sider: React.SFC<AllProps> = (props: AllProps) => {

  return (
    <Grid container spacing={24}>
      <Grid className={props.classes.main} item xs={12}>
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
      </Grid>
    </Grid>
  )
}

export default withStyles(Styles)(Sider)

/*
<Menu mode={mode} defaultSelectedKeys={['1']} defaultOpenKeys={['Home']} style={{ height: '100%', borderRight: 0 }} >
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
  </Menu.Item>
</Menu>
*/
