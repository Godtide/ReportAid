import * as React from 'react'
import { connect } from 'react-redux'
import PlainText from '../../../components/io/plainText'

import { ApplicationState } from '../../../store'

// Separate state props + dispatch props to their own interfaces.
interface OverviewProps {
  title: string
  data: string
}

class Overview extends React.Component<OverviewProps> {

  public render() {
    return (
      <div>
        <h2>{this.props.title}</h2>
        <PlainText text={this.props.data} />
      </div>
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
