import * as React from 'react'
import { connect } from 'react-redux'
import { Dispatch } from 'redux'
import { HomeState, HomePage, fetchRequest } from '../../../store/helpers/home'
import { ConnectedReduxProps, ApplicationState } from '../../../store'

interface HomeProps extends ConnectedReduxProps<HomeState> {}

// Separate state props + dispatch props to their own interfaces.
interface PropsFromState {
  data: HomePage[]
  errors: string
}

// We can use `typeof` here to map our dispatch types to the props, like so.
interface PropsFromDispatch {
  fetchRequest: typeof fetchRequest
}

type AllProps = PropsFromState & PropsFromDispatch & HomeProps

class Home extends React.Component<AllProps> {
  public render() {
    const { data } = this.props

    return (
      <p>
        {data}
      </p>
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
)(Home)
