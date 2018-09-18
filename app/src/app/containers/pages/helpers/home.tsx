import * as React from 'react'
import { connect } from 'react-redux'

import { ApplicationState } from '../../../store'

interface HomeProps {
  title: string
  data: []
}

class Home extends React.Component<HomeProps> {

  public render() {
    return (
      <div>
        <h2>{this.props.title}</h2>
        <p>
          {this.props.data}
        </p>
      </div>
    )
  }
}

// It's usually good practice to only include one context at a time in a connected component.
// Although if necessary, you can always include multiple contexts. Just make sure to
// separate them from each other to prevent prop conflicts.
const mapStateToProps = (state: ApplicationState, ownProps: HomeProps): HomeProps => {
  return {
    title: ownProps.title,
    data: ownProps.data
  }
}

// Now let's connect our component!
// With redux v4's improved typings, we can finally omit generics here.
export default connect(
  mapStateToProps
)(Home)
