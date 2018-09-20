import * as React from 'react'
import { connect } from 'react-redux'
import PlainText from '../../../components/io/plainText'

import { ApplicationState } from '../../../store'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

// Separate state props + dispatch props to their own interfaces.
interface OverviewProps {
  readonly title: string
  readonly data: string
}

class Overview extends React.Component<WithStyles<typeof styles> & OverviewProps> {

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
export default withTheme(withStyles(styles)(connect(
  mapStateToProps
)(Overview)))
