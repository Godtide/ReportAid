import * as React from 'react'
import { connect } from 'react-redux'
import { RouteComponentProps, Route, Switch } from 'react-router-dom'

import { HomeIndexPage } from './homeIndex'

import { ApplicationState, ConnectedReduxProps } from '../../../store'
// import { IATIReader } from '../../../store/IATIReader/types'

// Separate state props + dispatch props to their own interfaces.
interface PropsFromState {
  loading: boolean
  data: {}
  errors: string
}

// Combine both state + dispatch props - as well as any props we want to pass - in a union type.
type AllProps = PropsFromState & RouteComponentProps<{}> & ConnectedReduxProps

class HomePage extends React.Component<AllProps> {
  public render() {
    const { match } = this.props

    return (
      <Switch>
        <Route exact path={match.path + '/'} component={HomeIndex} />
      </Switch>
    )
  }
}

// It's usually good practice to only include one context at a time in a connected component.
// Although if necessary, you can always include multiple contexts. Just make sure to
// separate them from each other to prevent prop conflicts.
const mapStateToProps = ({ reader }: ApplicationState) => ({
  loading: reader.loading,
  errors: reader.errors,
  data: reader.data
})

// Now let's connect our component!
// With redux v4's improved typings, we can finally omit generics here.
export default connect(mapStateToProps)(HomePage)
