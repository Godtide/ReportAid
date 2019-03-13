import * as React from 'react'
import * as Yup from 'yup'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../store'
import { ActionProps } from '../../store/types'

import { Field, ErrorMessage} from 'formik'
import { Select } from "material-ui-formik-components"

import { setKey } from '../../store/helpers/keys/actions'
import { getOrganisationRefs } from '../../store/IATI/IATIReader/organisations/organisation/actions'

import { Keys, KeyTypes } from '../../store/helpers/keys/types'

import { IATIOrganisationReport, IATIOrganisationData, OrganisationReportProps,OrganisationProps } from '../../store/IATI/types'

interface OrganisationFormProps {
  setValue: Function
  name: string
  label: string
}

interface OrganisationDataProps {
  organisationsRef: string
  organisationRefs: Array<string>
}

interface OrganisationDispatchProps {
  getOrganisationRefs: (organisationProps: OrganisationReportProps) => void
  setOrganisationKey: (organisationRef: string) => void
}

type OrganisationPickerProps = OrganisationFormProps & OrganisationDataProps & OrganisationDispatchProps

class Organisation extends React.Component<OrganisationPickerProps> {

  constructor (props: OrganisationPickerProps) {
   super(props)
  }

  componentDidUpdate(previousProps: OrganisationPickerProps) {
    //console.log('Organisations: ', this.props.organisationsRef)
    if(this.props.organisationsRef != "" &&  previousProps.organisationsRef != this.props.organisationsRef) {
      this.props.getOrganisationRefs({organisationsRef: this.props.organisationsRef})
      //console.log('Done Updating: ', this.props.organisationsRef)
    }
  }

  render() {

    let organisationRefs: any[] = [{ value: "", label: "" }]
    if ( typeof this.props.organisationRefs != 'undefined' &&
         typeof this.props.organisationRefs[0] != 'undefined' &&
         this.props.organisationRefs[0] != "" ) {
      this.props.organisationRefs.forEach((key: string) => {
       organisationRefs.push({ value: key, label: key })
      })
    }

    return (
      <React.Fragment>
        <Field
          name={this.props.name}
          label={this.props.label}
          component={Select}
          onChange={(ev: any) => {
            this.props.setOrganisationKey(ev.target.value)
            this.props.setValue(this.props.name, ev.target.value)
          }}
          options={organisationRefs}
        />
      </React.Fragment>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrganisationDataProps => {
  //console.log(state.orgReader)
  return {
    organisationRefs: state.organisationPicker.data as Array<string>,
    organisationsRef: state.keys.data.organisations
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrganisationDispatchProps => {
  return {
    getOrganisationRefs: (organisationProps: OrganisationReportProps) => dispatch(getOrganisationRefs(organisationProps)),
    setOrganisationKey: (organisationRef: string) => dispatch(setKey({key: organisationRef, keyType: KeyTypes.ORGANISATION})),
  }
}

export const OrganisationPicker = connect<OrganisationDataProps, OrganisationDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(Organisation)
