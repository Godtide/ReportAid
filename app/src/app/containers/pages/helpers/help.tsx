import * as React from 'react'
import { Link } from 'react-router-dom'
import { connect, Dispatch } from 'react-redux'
//import moment from 'moment'

import styled from '../../../styles/styled'
import Page from '../../../components/layout/page'
import Container from '../../../components/layout/container'
/*import LoadingOverlay from '../../components/data/LoadingOverlay'
import LoadingOverlayInner from '../../components/data/LoadingOverlayInner'
import LoadingSpinner from '../../components/data/LoadingSpinner'*/

import { ApplicationState, ConnectedReduxProps } from '../../../store'
import { IATIReader } from '../../../store/IATIReader/types'
import { fetchRequest } from '../../..store/IATIReader/actions'

// Separate state props + dispatch props to their own interfaces.
interface PropsFromState {
  loading: boolean
  data: IATIReader[]
  errors: string
}

// We can use `typeof` here to map our dispatch types to the props, like so.
interface PropsFromDispatch {
  fetchRequest: typeof fetchRequest
}

// Combine both state + dispatch props - as well as any props we want to pass - in a union type.
type AllProps = PropsFromState & PropsFromDispatch & ConnectedReduxProps

class HomeIndexPage extends React.Component<AllProps> {
  public componentDidMount() {
    const { data } = this.props

    if (data.length === 0) {
      this.props.fetchRequest()
    }
  }

  public render() {
    const { loading } = this.props

    return (
      <Page>
        <Container>
            blah
        </Container>
      </Page>
    )
  }
}

// It's usually good practice to only include one context at a time in a connected component.
// Although if necessary, you can always include multiple contexts. Just make sure to
// separate them from each other to prevent prop conflicts.
const mapStateToProps = ({ home }: ApplicationState) => ({
  loading: home.loading,
  errors: home.errors,
  data: home.data
})

// mapDispatchToProps is especially useful for constraining our actions to the connected component.
// You can access these via `this.props`.
const mapDispatchToProps = (dispatch: Dispatch) => ({
  fetchRequest: () => dispatch(fetchRequest())
})

// Now let's connect our component!
// With redux v4's improved typings, we can finally omit generics here.
export default connect(
  mapStateToProps,
  mapDispatchToProps
)(HomeIndexPage)
