import * as React from 'react'
import {Link} from 'react-router-dom'
import { Layout, Breadcrumb } from 'antd'
import LayoutContainer from '../../containers/pages/layoutContainer'

interface ContentProps {
  title: string
}

const Content: React.SFC<ContentProps> = ({ title }) => (

  const { SubMenu } = Menu
  const { Content } = Layout

  return (
    <div>
      <Breadcrumb style={{ margin: '16px 0' }}>
        <Breadcrumb.Item>{AppStrings.home}</Breadcrumb.Item>
        <Breadcrumb.Item>{AppStrings.create}</Breadcrumb.Item>
        <Breadcrumb.Item>{AppStrings.read}</Breadcrumb.Item>
        <Breadcrumb.Item>{AppStrings.about}</Breadcrumb.Item>
        <Breadcrumb.Item>{AppStrings.overview}</Breadcrumb.Item>
        <Breadcrumb.Item>{AppStrings.help}</Breadcrumb.Item>
      </Breadcrumb>

      <Content style={{ padding: 24, margin: 0, minHeight: 280 }}>

        <Route name={AppStrings.home} exact path={AppPaths.home} component={Home} />
        <Route name={AppStrings.about} path={AppPaths.about} component={About} />
        <Route name={AppStrings.overview} path={AppPaths.overview} component={Overview} />
        <Route name={AppStrings.help} path={AppPaths.help} component={Help} />
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

      </Content>
    </div>
)

export default Container
