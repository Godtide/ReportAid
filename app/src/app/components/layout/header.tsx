import * as React from 'react'
import { NavLink } from 'react-router-dom'
import { Icon, Layout, Menu, Row, Col } from 'antd'
import LayoutContainer from '../../containers/pages/layoutContainer'

import styled from '../../styles/styled'

interface HeaderProps {
  title: string
}

const Header: React.SFC<HeaderProps> = ({ title }) => (

  const { SubMenu } = Menu
  const { Header, Content, Sider, Footer } = Layout

  return (
    <div>
      <Layout>
        <Header>
          <Row>
            <Col span={4}><img src={logo} /></Col>
            <Col span={10}>
              <Menu theme="dark" mode="horizontal" defaultSelectedKeys={['home']} style={{ lineHeight: '64px' }} >
                <Menu.Item key={AppStrings.home}>
                  <Icon type={HomeStrings.icon}/><span>{AppStrings.home}</span>
                  <Link to={AppPaths.home}/>
                </Menu.Item>
                <Menu.Item key={AppStrings.about}>
                  <Icon type={AboutStrings.icon}/><span>{AppStrings.about}</span>
                  <Link to={AppPaths.about}/>
                </Menu.Item>
                <Menu.Item key={AppStrings.overview}>
                  <Icon type={OverviewStrings.icon}/><span>{AppStrings.overview}</span>
                  <Link to={AppPaths.overview}/>
                </Menu.Item>
                <Menu.Item key={AppStrings.help}>
                  <Icon type={HelpStrings.icon}/><span>{AppStrings.help}</span>
                  <Link to={AppPaths.help}/>
                </Menu.Item>
              </Menu>
            </Col>
            <Col span={10} style={{ textAlign: 'right' }}>
              <h5>{title}</h5>
            </Col>
          </Row>
        </Header>
      </Layout>
      <Layout>


)

export default Header

/*
  <Wrapper>
    <HeaderInner>
      <HeaderLeft>
        <Title>{title}</Title>
      </HeaderLeft>
      <HeaderNav>
        <HeaderNavLink exact to="/" activeClassName={HeaderLinkActive}>
          Home
        </HeaderNavLink>
        <HeaderNavLink to="/IATIWriter" activeClassName={HeaderLinkActive}>
          Heroes
        </HeaderNavLink>
        <HeaderNavLink to="/IATIReader" activeClassName={HeaderLinkActive}>
          Teams
        </HeaderNavLink>
      </HeaderNav>
      <HeaderRight>
        <LayoutContainer>
          {({ theme, setTheme }) => (
            <React.Fragment>
              <CurrentTheme>Current theme: {theme}</CurrentTheme>
              <ThemeSwitcherButton onClick={() => setTheme(theme === 'light' ? 'dark' : 'light')}>
                Switch theme
              </ThemeSwitcherButton>
            </React.Fragment>
          )}
        </LayoutContainer>
      </HeaderRight>
    </HeaderInner>
  </Wrapper>

  */
