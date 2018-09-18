import * as React from 'react'
import { connect } from 'react-redux'
import { HomeState, HomePage, fetchRequest } from '../../../store/helpers/home'
import { ApplicationState } from '../../../store'

// Separate state props + dispatch props to their own interfaces.
type AllProps = HomeState & HomePage

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
const mapStateToProps = (ownProps: HomePropsFromState): HomePropsFromState => {
  return {
    data: ownProps.data,
    errors: ownProps.errors
  }
}

// Now let's connect our component!
// With redux v4's improved typings, we can finally omit generics here.
export default connect(
  mapStateToProps
)(Home)
