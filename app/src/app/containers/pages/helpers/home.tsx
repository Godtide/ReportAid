import * as React from 'react'
import { Home } from './types'
import { fetchRequest } from '../../..store/IATIReader/actions'

// Separate state props + dispatch props to their own interfaces.
interface PropsFromState {
  loading: boolean
  errors: string
  data: Home[]
}

// Combine both state + dispatch props - as well as any props we want to pass - in a union type.
type AllProps = PropsFromState & PropsFromDispatch & ConnectedReduxProps

class Home extends React.Component<AllProps> {
  public componentDidMount() {
    const { data } = this.props

    if (data.length === 0) {
      this.props.fetchRequest()
    }
  }

  public render() {
    const { loading } = this.props

    return (
      <div>
        Home
      </div>
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
