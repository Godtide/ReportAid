import * as React from 'react'
import { Route } from 'react-router-dom'
import { Layout, Breadcrumb } from 'antd'

import Home from '../../containers/pages/helpers/home'
import About from '../../containers/pages/helpers/about'
import Overview from '../../containers/pages/helpers/overview'
import Help from '../../containers/pages/helpers/help'

interface ContentProps {
  homeName: string
  homePath: string
  about: string
  aboutPath: string
  overview: string
  overviewPath: string
  help: string
  helpPath: string
}

const Content: React.SFC<ContentProps> = ({appStrings: {
    home,
    homePath,
    about,
    aboutPath,
    overview,
    overviewPath,
    help,
    helpPath
  }}) => (

  const { Content } = Layout

  return (
    <div>
      <Breadcrumb style={{ margin: '16px 0' }}>
        <Breadcrumb.Item>{home}</Breadcrumb.Item>
        <Breadcrumb.Item>{about}</Breadcrumb.Item>
        <Breadcrumb.Item>{overview}</Breadcrumb.Item>
        <Breadcrumb.Item>{help}</Breadcrumb.Item>
      </Breadcrumb>

      <Content style={{ padding: 24, margin: 0, minHeight: 280 }}>
        <Route name={home} exact path={homePath} component={Home} />
        <Route name={about} path={aboutPath} component={About} />
        <Route name={overview} path={overviewPath} component={Overview} />
        <Route name={help} path={helpPath} component={Help} />
      </Content>
    </div>
)

export default Container

/*
<Route
  name={AppStrings.create}
  path={AppPaths.create}
  render={() => <Writer contracts={this.contractHandler} web3={this.web3Handler} />}
/>
<Route
  name={AppStrings.read}
  path={AppPaths.read}
  render={() => <Reader contracts={this.contractHandler} web3={this.web3Handler} />}
/>
*/
