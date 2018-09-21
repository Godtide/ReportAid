import * as React from 'react'
import { connect } from 'react-redux'
import { ApplicationState } from '../../../store'
import { AboutProps } from '../../../store/helpers/about/types'
import MarkdownText from '../../../containers/io/markdownText'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

// It's usually good practice to only include one context at a time in a connected component.
// Although if necessary, you can always include multiple contexts. Just make sure to
// separate them from each other to prevent prop conflicts.
const mapStateToProps = (state: ApplicationState, ownProps: AboutProps): AboutProps => {
  return {
    title: ownProps.title,
    data: ownProps.data
  }
}

type AllProps = AboutProps & mapStateToProps


class About extends React.Component<WithStyles<typeof styles> & AboutProps> {

  public render() {

    return (
      <div>
        <h2>{this.props.title}</h2>
        <MarkdownText text={this.props.data} />
      </div>
    )
  }
}

// Now let's connect our component!
// With redux v4's improved typings, we can finally omit generics here.
export default withTheme(withStyles(styles)(connect(
  mapStateToProps
)(About)))
