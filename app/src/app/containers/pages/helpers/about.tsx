import * as React from 'react'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import { ApplicationState } from '../../../store'
import { AboutProps } from '../../../store/helpers/about/types'
import MarkdownText from '../../../containers/io/markdownText'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

import getData from '../../../store/helpers/about/actions'


interface AboutPropsFromState {
  about: ApplicationState
}

interface AboutPropsFromDispatch {
  getData: typeof getData
}

// Combine both state + dispatch props - as well as any props we want to pass - in a union type.
type AllProps = AboutProps & AboutPropsFromState & AboutPropsFromDispatch

class About extends React.Component<WithStyles<typeof styles> & AllProps> {

  componentDidMount() {
    // only fetch the data if there is no data
    if (!this.props.data) this.props.getData()
  }

  render() {
    const data  = this.props.getData()
    console.log(data)

    return (
      <div>
        &nbsp;
      </div>
    )
  }
}

/*
<div>
  <h2>{data.title}</h2>
  <MarkdownText text={data.text} />
</div>
*/

const mapStateToProps = (state: ApplicationState, ownProps: AboutProps): ApplicationState => {
  return {
    about: state.about
  }
}

const mapDispatchToProps = (dispatch: any, ownProps: AboutProps) => {
  return bindActionCreators({
    getData: getData
  }, dispatch)
}

export default withTheme(withStyles(styles)(connect(
  mapStateToProps,
  mapDispatchToProps
)(About)))
