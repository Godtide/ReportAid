import * as React from 'react'
import { Link } from 'react-router-dom'
import { Icon, Layout, Menu, Row, Col } from 'antd'

import { PathStrings } from '../../utils/strings'
import { MenuMode } from '../../utils/theme'


interface SiderInterface {
  siderWidth: string
  siderMode: MenuMode
}

export type  SiderProps = SiderInterface & PathStrings

const Sider: React.SFC<SiderProps> = (props) => {

  const width = props.siderWidth
  const mode = props.siderMode

  return (

    <div>
      <Layout width={width}>
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
      </Layout>
    </div>
  )
}

export default Sider
