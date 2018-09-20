import * as React from 'react'
import { connect } from 'react-redux'
import PlainText from '../../../components/io/plainText'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

import { ApplicationState } from '../../../store'

interface HomeProps {
  readonly title: string
  readonly data: string
}

class Home extends React.Component<WithStyles<typeof styles> & HomeProps> {

  public render() {
    return (
      <div>
        <h2>{this.props.title}</h2>
        <PlainText text={this.props.data} />
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
export default withTheme(withStyles(styles)(connect(
  mapStateToProps
)(Home)))
