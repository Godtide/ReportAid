import * as React from 'react'
import { connect } from 'react-redux'

import { ApplicationState } from '../../../store'
import { HomePage } from '../../../store/helpers/home/types'
import { fetchRequest } from '../../../store/helpers/home/actions'

// Separate state props + dispatch props to their own interfaces.
interface AboutPropsFromState {
  data: string
  errors: string
}

// Combine both state + dispatch props - as well as any props we want to pass - in a union type.
type AllProps = AboutPropsFromState

class About extends React.Component<AllProps> {

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
)(About)
