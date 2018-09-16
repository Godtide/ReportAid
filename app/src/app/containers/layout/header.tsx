import * as React from 'react'
import { Link } from 'react-router-dom'
import { Icon, Layout, Menu, Row, Col } from 'antd'

import { PathStrings } from '../../utils/strings'

interface HeaderInterface {
  headerTitle: string
}

export type  HeaderProps = PathStrings & HeaderInterface

const Header: React.SFC<HeaderProps> = ( props ) => {

  return (
    <div>
      <Layout>
        <Row>
          <Col span={4}>
            <p>Logo to go here</p>
          </Col>
          <Col span={10}>
            <Menu theme="dark" mode="horizontal" defaultSelectedKeys={['home']} style={{ lineHeight: '64px' }} >
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
          </Col>
          <Col span={10} style={{ textAlign: 'right' }}>
            <h5>{props.headerTitle}</h5>
          </Col>
        </Row>
      </Layout>
    </div>
  )
}

export default Header
