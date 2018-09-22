import * as React from 'react'
import { bindActionCreators, Dispatch, AnyAction } from 'redux'
import { connect } from 'react-redux'
import { ApplicationState } from '../../../store'
import MarkdownText from '../../../containers/io/markdownText'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

import { AboutFetchRequest, AboutRequestDataAction } from '../../../store/helpers/about/actions'
import { AboutProps, AboutActionTypes } from '../../../store/helpers/about/types'
import { AboutStrings } from '../../../utils/strings'

interface AllProps {
  data: AboutProps
}

class About extends React.Component<WithStyles<typeof styles> & AllProps> {

  componentDidMount() {

    const type = AboutActionTypes.REQ_DATA
    const payload = {
      title: AboutStrings.heading,
      data: AboutStrings.info
    }
    AboutFetchRequest(type, payload)
  }

  render() {

    return (
      <div>
        <h2>{this.props.data.title}</h2>
        <MarkdownText text={this.props.data.data} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): AllProps => {
  //const props = { state.about.title, data: state.about.data }
  return {
    data: { title: state.about.title, data: state.about.data }
  }
}

const mapDispatchToProps = (dispatch: Dispatch<AnyAction>, ownProps: AboutProps) => {
  console.log('bollox', ownProps.title, ownProps.data)
  const type = AboutActionTypes.REQ_DATA
  const payload = {
    title: ownProps.title,
    data: ownProps.data
  }
  return bindActionCreators<AboutRequestDataAction, any>({
    data: AboutFetchRequest(type, payload)
  },dispatch)
}

export default withTheme(withStyles(styles)(connect(
  mapStateToProps,
  mapDispatchToProps
)(About)))
