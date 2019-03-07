import * as React from 'react'
import * as Yup from 'yup'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../store'
import { ActionProps } from '../../store/types'

import { Field, ErrorMessage} from 'formik'
import { Select } from "material-ui-formik-components"

import { setKey } from '../../store/helpers/keys/actions'
import { getActivity } from '../../store/IATI/IATIReader/activities/activity/actions'

import { Keys, KeyTypes } from '../../store/helpers/keys/types'

import { IATIActivityReport, IATIActivityData, ActivityReportProps, ActivityProps } from '../../store/IATI/types'

interface ActivityFormProps {
  setValue: Function
  name: string
  label: string
}

interface ActivityDataProps {
  activitiesRef: string
  activity: IATIActivityReport
}

interface ActivityDispatchProps {
  getActivity: (activityProps: ActivityReportProps) => void
  setActivityKey: (activityRef: string) => void
}

type ActivityPickerProps = ActivityFormProps & ActivityDataProps & ActivityDispatchProps

class Activity extends React.Component<ActivityPickerProps> {

  constructor (props: ActivityPickerProps) {
   super(props)
  }

  componentDidUpdate(previousProps: ActivityPickerProps) {
    //console.log('Activitys: ', this.props.activitiesRef)
    if(this.props.activitiesRef != "" &&  previousProps.activitiesRef != this.props.activitiesRef) {
      this.props.getActivity({isReport: false, activitiesRef: this.props.activitiesRef})
      //console.log('Done Updating: ', this.props.activitiesRef)
    }
  }

  render() {

    //console.log ('rendering', this.props.activity, this.props.activitiesRef)
    let activityRefs: any[] = [{ value: "", label: "" }]

    //console.log(this.props.activitiesRef, this.props.activity.activitiesRef)
    if ( this.props.activitiesRef != "" &&
         this.props.activity.activitiesRef == this.props.activitiesRef ) {

      const activities: Array<IATIActivityData> = this.props.activity.data
      //console.log('Data: ', activities)
      for (let i = 0; i < activities.length; i++) {
       const label = activities[i].activityRef
       const value = label
       activityRefs.push({ value: value, label: label })
      }
    }

    return (
      <React.Fragment>
        <Field
          name={this.props.name}
          label={this.props.label}
          component={Select}
          onChange={(ev: any) => {
            this.props.setActivityKey(ev.target.value)
            this.props.setValue(this.props.name, ev.target.value)
          }}
          options={activityRefs}
        />
      </React.Fragment>
    )
  }
}

const mapStateToProps = (state: ApplicationState): ActivityDataProps => {
  //console.log(state.orgReader)
  return {
    activity: state.activityPicker.data as IATIActivityReport,
    activitiesRef: state.keys.data.activities
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): ActivityDispatchProps => {
  return {
    getActivity: (activityProps: ActivityReportProps) => dispatch(getActivity(activityProps)),
    setActivityKey: (activityRef: string) => dispatch(setKey({key: activityRef, keyType: KeyTypes.ACTIVITY})),
  }
}

export const ActivityPicker = connect<ActivityDataProps, ActivityDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(Activity)
