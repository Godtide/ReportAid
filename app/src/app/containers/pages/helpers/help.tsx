import * as React from 'react'
import { connect } from 'react-redux'
import MarkdownText from '../../../containers/io/markdownText'
import PlainText from '../../../components/io/plainText'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

import { ApplicationState } from '../../../store'

interface HelpProps {
  readonly title: string
  readonly data: string
}

class Help extends React.Component<WithStyles<typeof styles> & HelpProps> {

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
export default withTheme(withStyles(styles)(connect(
  mapStateToProps
)(Help)))
