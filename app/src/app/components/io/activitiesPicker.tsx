import * as React from 'react'

import * as Yup from 'yup'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../store'
import { ActionProps } from '../../store/types'

import { Field, ErrorMessage, FieldProps } from 'formik'
import { Select } from "material-ui-formik-components"
//import MuiSelect from '@material-ui/core/Select'

import { setKey } from '../../store/helpers/keys/actions'
import { getActivities } from '../../store/IATI/IATIReader/activities/activities/actions'

import { Keys, KeyTypes } from '../../store/helpers/keys/types'

import { IATIActivitiesReport,
         IATIActivitiesData } from '../../store/IATI/types'

interface ActivitiesFormProps {
  setValue: Function
  name: string
  label: string
}

interface ActivitiesDataProps {
  activities: IATIActivitiesReport
}

interface ActivitiesDispatchProps {
  getActivities: (isReport: boolean) => void
  setActivitiesKey: (activitiesRef: string) => void
}

type ActivitiesPickerProps = ActivitiesFormProps & ActivitiesDataProps & ActivitiesDispatchProps

class Activities extends React.Component<ActivitiesPickerProps> {

  constructor (props: ActivitiesPickerProps) {
   super(props)
  }

  componentDidMount() {
    this.props.getActivities(false)
  }

  render() {

    let activitiesRefs: any[] = [{ value: "", label: "" }]
    if(typeof this.props.activities != "undefined" &&
       this.props.activities.hasOwnProperty('data')) {

      //console.log ('rendering', this.props.activities)
      const activities: Array<IATIActivitiesData> = this.props.activities.data
      for (let i = 0; i < activities.length; i++) {
       const label = activities[i].activitiesRef
       const value = label
       activitiesRefs.push({ value: value, label: label })
      }
    }

    return (
      <React.Fragment>
        <Field
          name={this.props.name}
          label={this.props.label}
          component={Select}
          onChange={(ev: any) => {
            this.props.setActivitiesKey(ev.target.value)
            this.props.setValue(this.props.name, ev.target.value)
          }}
          options={activitiesRefs}
        />
      </React.Fragment>
    )
  }
}

const mapStateToProps = (state: ApplicationState): ActivitiesDataProps => {
  //console.log(state.orgReader)
  return {
    activities: state.activitiesPicker.data as IATIActivitiesReport
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): ActivitiesDispatchProps => {
  return {
    getActivities: (isReport: boolean) => dispatch(getActivities(isReport)),
    setActivitiesKey: (activitiesRef: string) => dispatch(setKey({key: activitiesRef, keyType: KeyTypes.ACTIVITIES})),
  }
}

export const ActivitiesPicker = connect<ActivitiesDataProps, ActivitiesDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(Activities)
