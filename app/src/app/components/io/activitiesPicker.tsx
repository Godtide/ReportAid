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
import { getActivitiesRefs } from '../../store/IATI/IATIReader/activities/activities/actions'

import { Keys, KeyTypes } from '../../store/helpers/keys/types'

interface ActivitiesFormProps {
  setValue: Function
  name: string
  label: string
}

interface ActivitiesDataProps {
  activitiesRefs: Array<string>
}

interface ActivitiesDispatchProps {
  getActivitiesRefs: () => void
  setActivitiesKey: (activitiesRef: string) => void
}

type ActivitiesPickerProps = ActivitiesFormProps & ActivitiesDataProps & ActivitiesDispatchProps

class Activities extends React.Component<ActivitiesPickerProps> {

  constructor (props: ActivitiesPickerProps) {
   super(props)
  }

  componentDidMount() {
    this.props.getActivitiesRefs()
  }

  render() {

    let activitiesRefs: any[] = [{ value: "", label: "" }]
    if ( typeof this.props.activitiesRefs != 'undefined' &&
         typeof this.props.activitiesRefs[0] != 'undefined' &&
         this.props.activitiesRefs[0] != "" ) {
      this.props.activitiesRefs.forEach((key: string) => {
       activitiesRefs.push({ value: key, label: key })
      })
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
    activitiesRefs: state.activitiesPicker.data as Array<string>
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): ActivitiesDispatchProps => {
  return {
    getActivitiesRefs: () => dispatch(getActivitiesRefs()),
    setActivitiesKey: (activitiesRef: string) => dispatch(setKey({key: activitiesRef, keyType: KeyTypes.ACTIVITIES})),
  }
}

export const ActivitiesPicker = connect<ActivitiesDataProps, ActivitiesDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(Activities)
