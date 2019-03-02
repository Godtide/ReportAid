import * as React from 'react'
import * as Yup from 'yup'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../store'
import { ActionProps } from '../../store/types'

import { Field, ErrorMessage} from 'formik'
import { Select } from "material-ui-formik-components"

import { setOrganisationKey } from '../../store/helpers/keys/actions'
import { getOrganisation } from '../../store/IATI/IATIReader/organisations/organisation/actions'

import { IATIOrganisationReport, IATIOrganisationData, OrganisationReportProps,OrganisationProps } from '../../store/IATI/types'

interface OrganisationFormProps {
  setValue: Function
  name: string
  label: string
}

interface OrganisationDataProps {
  organisationsRef: string
  organisation: IATIOrganisationReport
}

interface OrganisationDispatchProps {
  getOrganisation: (organisationProps: OrganisationReportProps) => void
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
      this.props.getOrganisation({isReport: false, organisationsRef: this.props.organisationsRef})
      //console.log('Done Updating: ', this.props.organisationsRef)
    }
  }

  render() {

    //console.log ('rendering', this.props.organisation, this.props.organisationsRef)
    let organisationRefs: any[] = [{ value: "", label: "" }]

    //console.log(this.props.organisationsRef, this.props.organisation.organisationsRef)
    if ( this.props.organisationsRef != "" &&
         this.props.organisation.organisationsRef == this.props.organisationsRef ) {

      const organisations: Array<IATIOrganisationData> = this.props.organisation.data
      //console.log('Data: ', organisations)
      for (let i = 0; i < organisations.length; i++) {
       const label = organisations[i].organisationRef
       const value = label
       organisationRefs.push({ value: value, label: label })
      }
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
    organisation: state.organisationPicker.data as IATIOrganisationReport,
    organisationsRef: state.keys.data.organisations
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrganisationDispatchProps => {
  return {
    getOrganisation: (organisationProps: OrganisationReportProps) => dispatch(getOrganisation(organisationProps)),
    setOrganisationKey: (organisationRef: string) => dispatch(setOrganisationKey(organisationRef)),
  }
}

export const OrganisationPicker = connect<OrganisationDataProps, OrganisationDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(Organisation)
