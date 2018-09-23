import * as React from 'react'
//import { bindActionCreators, Dispatch, AnyAction } from 'redux'
import { connect } from 'react-redux'
import { ApplicationState } from '../../../store'
import MarkdownText from '../../../containers/io/markdownText'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

//import { fetchRequest, RequestDataAction } from '../../../store/helpers/about/actions'
import { AboutProps } from '../../../store/helpers/about/types'

class About extends React.Component<WithStyles<typeof styles> & AboutProps> {

  render() {

    return (
      <div>
        <h2>{this.props.title}</h2>
        <MarkdownText text={this.props.data} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): AboutProps => {
  return {
    title: state.about.title, data: state.about.data
  }
}

/* const mapDispatchToProps = (dispatch: Dispatch<AnyAction>, ownProps: AboutProps) => {
  console.log('bollox', ownProps.title, ownProps.data)
  const type = AboutActionTypes.REQ_DATA
  const payload = {
    title: ownProps.title,
    data: ownProps.data
  }
  return bindActionCreators<AboutRequestDataAction, any>({
    onSomeEvent: AboutFetchRequest(type, payload)
  },dispatch)
} */

export default withTheme(withStyles(styles)(connect(
  mapStateToProps
)(About)))
