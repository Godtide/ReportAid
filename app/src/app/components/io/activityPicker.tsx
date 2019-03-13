import * as React from 'react'
import * as Yup from 'yup'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../store'
import { ActionProps } from '../../store/types'

import { Field, ErrorMessage} from 'formik'
import { Select } from "material-ui-formik-components"

import { setKey } from '../../store/helpers/keys/actions'
import { getActivityRefs } from '../../store/IATI/IATIReader/activities/activity/actions'

import { Keys, KeyTypes } from '../../store/helpers/keys/types'

import { ActivityReportProps, ActivityProps } from '../../store/IATI/types'

interface ActivityFormProps {
  setValue: Function
  name: string
  label: string
}

interface ActivityDataProps {
  activitiesRef: string
  activityRefs: Array<string>
}

interface ActivityDispatchProps {
  getActivityRefs: (activityProps: ActivityReportProps) => void
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
      this.props.getActivityRefs({activitiesRef: this.props.activitiesRef})
      //console.log('Done Updating: ', this.props.activitiesRef)
    }
  }

  render() {

    //console.log ('rendering', this.props.activityRefs, this.props.activitiesRef)
    let activityRefs: any[] = [{ value: "", label: "" }]
    //console.log(this.props.activitiesRef, this.props.activity.activitiesRef)
    if ( typeof this.props.activityRefs != 'undefined' &&
         typeof this.props.activityRefs[0] != 'undefined' &&
         this.props.activityRefs[0] != "" ) {
      this.props.activityRefs.forEach((key: string) => {
       activityRefs.push({ value: key, label: key })
      })
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
  //console.log('New state!: ', state.activityPicker.data)
  return {
    activityRefs: state.activityPicker.data as Array<string>,
    activitiesRef: state.keys.data.activities
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): ActivityDispatchProps => {
  return {
    getActivityRefs: (activityProps: ActivityReportProps) => dispatch(getActivityRefs(activityProps)),
    setActivityKey: (activityRef: string) => dispatch(setKey({key: activityRef, keyType: KeyTypes.ACTIVITY})),
  }
}

export const ActivityPicker = connect<ActivityDataProps, ActivityDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(Activity)
