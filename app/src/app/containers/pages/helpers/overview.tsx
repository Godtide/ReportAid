import * as React from 'react'
import { connect } from 'react-redux'

import { ApplicationState } from '../../../store'

// Separate state props + dispatch props to their own interfaces.
interface OverviewProps {
  title: string
  data: []
}

class Overview extends React.Component<OverviewProps> {

  public render() {

    const { data } = this.props
    return (
        <p>
          {data}
        </p>
    )
  }
}

const mapStateToProps = (state: ApplicationState, ownProps: OverviewProps): OverviewProps => {
  return {
    title: ownProps.title,
    data: ownProps.data
  }
}

// Now let's connect our component!
// With redux v4's improved typings, we can finally omit generics here.
export default connect(
  mapStateToProps
)(Overview)
