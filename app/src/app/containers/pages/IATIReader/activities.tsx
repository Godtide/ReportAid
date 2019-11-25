import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'
import Markdown from 'react-markdown'

import { getDictEntries } from '../../../components/io/dict'

import { initialise } from '../../../store/IATI/IATIReader/actions'
import { getActivities } from '../../../store/IATI/IATIReader/activities/activities/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { IATIActivitiesReport } from '../../../store/IATI/types'

import { Activities as ActivitiesStrings } from '../../../utils/strings'

interface ActivitiesProps {
  activities: IATIActivitiesReport
}

interface ActivitiesDispatchProps {
  initialise: () => void
  getActivities: () => void
}

type ActivitiesReaderProps =  ActivitiesProps & ActivitiesDispatchProps

class ActivitiesReader extends React.Component<ActivitiesReaderProps> {

  constructor (props: ActivitiesReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.initialise()
    this.props.getActivities()
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
    initialise: () => dispatch(initialise()),
    getActivities: () => dispatch(getActivities())
  }
}

export const Activities = connect<ActivitiesProps, ActivitiesDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(ActivitiesReader)
