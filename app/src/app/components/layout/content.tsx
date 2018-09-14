import * as React from 'react'
import { Route } from 'react-router-dom'
import { Layout, Breadcrumb } from 'antd'

import Home from '../../containers/pages/helpers/home'
import About from '../../containers/pages/helpers/about'
import Overview from '../../containers/pages/helpers/overview'
import Help from '../../containers/pages/helpers/help'

import { PathStrings } from '../../utils/strings'

export type  ContentProps = PathStrings

const Content: React.SFC<ContentProps> = (props) => {

  return (
    <div>props.
      <Layout>
        <Breadcrumb style={{ margin: '16px 0' }}>
          <Breadcrumb.Item>{props.home}</Breadcrumb.Item>
          <Breadcrumb.Item>{props.about}</Breadcrumb.Item>
          <Breadcrumb.Item>{props.overview}</Breadcrumb.Item>
          <Breadcrumb.Item>{props.help}</Breadcrumb.Item>
        </Breadcrumb>
        <Layout style={{ padding: 24, margin: 0, minHeight: 280 }}>
          <Route name={props.home} exact path={props.homePath} component={Home} />
          <Route name={props.about} path={props.aboutPath} component={About} />
          <Route name={props.overview} path={props.overviewPath} component={Overview} />
          <Route name={props.help} path={props.helpPath} component={Help} />
        </Layout>
      </Layout>
    </div>
  )
}

export default Content

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
