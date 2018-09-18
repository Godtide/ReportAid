import * as React from 'react'
import { Route } from 'react-router-dom'

import Home from '../pages/helpers/home'
import About from '../pages/helpers/about'
import Overview from '../pages/helpers/overview'
import Help from '../pages/helpers/help'

import { PathStrings } from '../../utils/strings'

export type  ContentProps = PathStrings

const Content: React.SFC<ContentProps> = (props) => {

  return (
    <div>
      <Route name={props.home} exact path={props.homePath} component={Home} />
      <Route name={props.about} path={props.aboutPath} component={About} />
      <Route name={props.overview} path={props.overviewPath} component={Overview} />
      <Route name={props.help} path={props.helpPath} component={Help} />
    </div>
  )
}

export default Content

/*
<Breadcrumb style={{ margin: '16px 0' }}>
  <Breadcrumb.Item>{props.home}</Breadcrumb.Item>
  <Breadcrumb.Item>{props.about}</Breadcrumb.Item>
  <Breadcrumb.Item>{props.overview}</Breadcrumb.Item>
  <Breadcrumb.Item>{props.help}</Breadcrumb.Item>
</Breadcrumb>

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
