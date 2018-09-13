import * as React from 'react'
import {Link} from 'react-router-dom'
import { Icon, Layout, Menu, Row, Col } from 'antd'

import { MenuMode } from '../../styles/types'

// import LayoutContainer from '../../containers/pages/layoutContainer'

interface FooterProps {
  width: number
  mode: MenuMode
}

const Footer: React.SFC<FooterProps> = ({ width, mode }) => (

  const { SubMenu } = Menu
  const { Sider } = Layout

  return (
    <div>
      <Sider width={{width}}>
        <Menu
          mode={mode}
          defaultSelectedKeys={['1']}
          defaultOpenKeys={['Home']}
          style={{ height: '100%', borderRight: 0 }}
        >
          <Menu.Item key={HomeStrings.home}>
            <Icon type={HomeStrings.homeIcon}/><span>{AppStrings.home}</span>
            <Link to={AppPaths.home}/>
          </Menu.Item>

          <SubMenu title={<span><Icon type={HomeStrings.homeIcon}/><span>{AppStrings.create}</span></span>}>
              <Menu.Item key={AppStrings.create}>
                <Icon type={HomeStrings.homeIcon}/><span>{AppStrings.create}</span>
                <Link to={AppPaths.create}/>
            </Menu.Item>
          </SubMenu>

          <SubMenu title={<span><Icon type={HomeStrings.homeIcon} /><span>{AppStrings.read}</span></span>}>
            <Menu.Item key={AppStrings.read}>
                <Icon type={HomeStrings.homeIcon}/><span>{AppStrings.read}</span>
                <Link to={AppPaths.read}/>
              </Menu.Item>
          </SubMenu>

          <Menu.Item key={AppStrings.about}>
            <Icon type={AboutStrings.aboutIcon}/><span>{AppStrings.about}</span>
            <Link to={AppPaths.about}/>
          </Menu.Item>'../helpers/outputStrings'
          <Menu.Item key={AppStrings.overview}>
            <Icon type={OverviewStrings.overviewIcon}/><span>{AppStrings.overview}</span>
            <Link to={AppPaths.overview}/>
          </Menu.Item>
          <Menu.Item key={AppStrings.help}>
            <Icon type={HelpStrings.helpIcon}/><span>{AppStrings.help}</span>
            <Link to={AppPaths.help}/>
          </Menu.Item>
        </Menu>
      </Sider>
    </div>
)

export default Footer


<Layout>



export default Sider
