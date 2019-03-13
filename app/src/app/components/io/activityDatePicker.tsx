import * as React from 'react'
import * as Yup from 'yup'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../store'
import { ActionProps } from '../../store/types'

import { Field, ErrorMessage} from 'formik'
import { Select } from "material-ui-formik-components"

import { setKey } from '../../store/helpers/keys/actions'
import { getActivityDateRefs } from '../../store/IATI/IATIReader/activities/activityDates/actions'

import { Keys, KeyTypes } from '../../store/helpers/keys/types'

import { ActivitiesReportProps, ActivityProps } from '../../store/IATI/types'

interface ActivityDateFormProps {
  setValue: Function
  name: string
  label: string
}

interface ActivityDateDataProps {
  activitiesRef: string
  activityRef: string
  activityDateRefs: Array<string>
}

interface ActivityDateDispatchProps {
  getActivityDateRefs: (activitiesProps: ActivitiesReportProps) => void
  setActivityDateKey: (activityDateRef: string) => void
}

type ActivityDatePickerProps = ActivityDateFormProps & ActivityDateDataProps & ActivityDateDispatchProps

class ActivityDate extends React.Component<ActivityDatePickerProps> {

  constructor (props: ActivityDatePickerProps) {
   super(props)
  }

  componentDidUpdate(previousProps: ActivityDatePickerProps) {
    //console.log('Activity Dates: ', this.props.activityRef, this.props.activitiesRef)
    if((this.props.activitiesRef != "" &&  previousProps.activitiesRef != this.props.activitiesRef &&
       this.props.activityRef != "") ||
       (this.props.activityRef != "" &&  previousProps.activityRef != this.props.activityRef &&
        this.props.activitiesRef != "") ) {
      this.props.getActivityDateRefs({activitiesRef: this.props.activitiesRef, activityRef: this.props.activityRef})
      //console.log('Done Updating: ', this.props.activitiesRef)
    }
  }

  render() {

    //console.log ('rendering', this.props.activityRefs, this.props.activitiesRef)
    let activityDateRefs: any[] = [{ value: "", label: "" }]
    //console.log(this.props.activitiesRef, this.props.activity.activitiesRef)
    if ( typeof this.props.activityDateRefs != 'undefined' &&
         typeof this.props.activityDateRefs[0] != 'undefined' &&
         this.props.activityDateRefs[0] != "" ) {
      this.props.activityDateRefs.forEach((key: string) => {
       activityDateRefs.push({ value: key, label: key })
      })
    }

    return (
      <React.Fragment>
        <Field
          name={this.props.name}
          label={this.props.label}
          component={Select}
          onChange={(ev: any) => {
            this.props.setActivityDateKey(ev.target.value)
            this.props.setValue(this.props.name, ev.target.value)
          }}
          options={activityDateRefs}
        />
      </React.Fragment>
    )
  }
}

const mapStateToProps = (state: ApplicationState): ActivityDateDataProps => {
  //console.log('New state!: ', state.activityPicker.data)
  return {
    activityDateRefs: state.activityDatesPicker.data as Array<string>,
    activityRef: state.keys.data.activity,
    activitiesRef: state.keys.data.activities
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): ActivityDateDispatchProps => {
  return {
    getActivityDateRefs: (props: ActivitiesReportProps) => dispatch(getActivityDateRefs(props)),
    setActivityDateKey: (activityDateRef: string) => dispatch(setKey({key: activityDateRef, keyType: KeyTypes.ACTIVITYDATE})),
  }
}

export const ActivityDatePicker = connect<ActivityDateDataProps, ActivityDateDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(ActivityDate)
