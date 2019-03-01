import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'
import Markdown from 'react-markdown'

import { getDictEntries } from '../../../components/io/dict'
import { getActivities } from '../../../store/IATI/IATIReader/activities/activities/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { IATIActivitiesReport } from '../../../store/IATI/IATIReader/activities/types'

import { Activities as ActivitiesStrings } from '../../../utils/strings'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

interface ActivitiesProps {
  activities: IATIActivitiesReport
}

interface ActivitiesDispatchProps {
  getActivities: (isReport: boolean) => void
}

type ActivitiesReaderProps =  WithStyles<typeof styles> & ActivitiesProps & ActivitiesDispatchProps

class ActivitiesReader extends React.Component<ActivitiesReaderProps> {

  constructor (props: ActivitiesReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.getActivities(true)
  }

  render() {

    const xs = getDictEntries(this.props.activities)
    return (
      <div>
        <h2>{ActivitiesStrings.headingActivitiesReader}</h2>
        <hr />
        <h3>{ActivitiesStrings.activitiesDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): ActivitiesProps => {
  //console.log(state.orgReader)
  return {
    activities: state.report.data as IATIActivitiesReport
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): ActivitiesDispatchProps => {
  return {
    getActivities: (isReport: boolean) => dispatch(getActivities(isReport))
  }
}

export const Activities = withTheme(withStyles(styles)(connect<ActivitiesProps, ActivitiesDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(ActivitiesReader)))
