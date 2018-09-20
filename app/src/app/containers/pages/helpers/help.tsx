import * as React from 'react'
import { connect } from 'react-redux'
import PlainText from '../../../components/io/plainText'

import { ApplicationState } from '../../../store'

interface HelpProps {
  title: string
  data: string
}

class Help extends React.Component<HelpProps> {

  public render() {
    return (
      <div>
        <h2>{this.props.title}</h2>
        <PlainText text={this.props.data} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState, ownProps: HelpProps): HelpProps => {
  return {
    title: ownProps.title,
    data: ownProps.data
  }
}

// Now let's connect our component!
// With redux v4's improved typings, we can finally omit generics here.
export default connect(
  mapStateToProps
)(Help)
