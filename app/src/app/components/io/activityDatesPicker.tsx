import * as React from 'react'
import * as Yup from 'yup'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../store'
import { ActionProps } from '../../store/types'

import { Field, ErrorMessage} from 'formik'
import { Select } from "material-ui-formik-components"

import { setKey } from '../../store/helpers/keys/actions'
import { getActivityDates } from '../../store/IATI/IATIReader/activities/activity/actions'

import { Keys, KeyTypes } from '../../store/helpers/keys/types'

import { IATIActivityDatesReport, IATIActivityDatesData, ActivityDatesReportProps, ActivityDatesProps } from '../../store/IATI/types'

interface ActivityDatesFormProps {
  setValue: Function
  name: string
  label: string
}

interface ActivityDatesDataProps {
  activitiesRef: string
  activityRef: string
  activityDates: IATIActivityDatesReport
}

interface ActivityDatesDispatchProps {
  getActivityDates: (activityProps: ActivityDatesReportProps) => void
  setActivityDatesKey: (activityRef: string) => void
}

type ActivityDatesPickerProps = ActivityDatesFormProps & ActivityDatesDataProps & ActivityDatesDispatchProps

class ActivityDates extends React.Component<ActivityDatesPickerProps> {

  constructor (props: ActivityDatesPickerProps) {
   super(props)
  }

  componentDidUpdate(previousProps: ActivityDatesPickerProps) {
    //console.log('ActivityDatess: ', this.props.activitiesRef)
    if(this.props.activitiesRef != "" &&  previousProps.activitiesRef != this.props.activitiesRef) {
      this.props.getActivityDates({isReport: false, activitiesRef: this.props.activitiesRef})
      //console.log('Done Updating: ', this.props.activitiesRef)
    }
  }

  render() {

    //console.log ('rendering', this.props.activity, this.props.activitiesRef)
    let activityRefs: any[] = [{ value: "", label: "" }]

    //console.log(this.props.activitiesRef, this.props.activity.activitiesRef)
    if ( this.props.activitiesRef != "" &&
         this.props.activity.activitiesRef == this.props.activitiesRef ) {

      const activities: Array<IATIActivityDatesData> = this.props.activity.data
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
            this.props.setActivityDatesKey(ev.target.value)
            this.props.setValue(this.props.name, ev.target.value)
          }}
          options={activityRefs}
        />
      </React.Fragment>
    )
  }
}

const mapStateToProps = (state: ApplicationState): ActivityDatesDataProps => {
  //console.log(state.orgReader)
  return {
    activitiesRef: state.keys.data.activities,
    activityRef: state.keys.data.activity,
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): ActivityDatesDispatchProps => {
  return {
    getActivityDates: (activityProps: ActivityDatesReportProps) => dispatch(getActivityDateRecords(activityProps)),
    setActivityDatesKey: (activityDateRef: string) => dispatch(setKey({key: activityDateRef, keyType: KeyTypes.ACTIVITYDATES})),
  }
}

export const ActivityDatesPicker = connect<ActivityDatesDataProps, ActivityDatesDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(ActivityDates)
