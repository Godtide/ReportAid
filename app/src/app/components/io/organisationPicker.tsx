import * as React from 'react'
import * as Yup from 'yup'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../store'
import { ActionProps } from '../../store/types'

import { Field, ErrorMessage} from 'formik'
import { Select } from "material-ui-formik-components"

import { getOrganisation } from '../../store/IATI/IATIReader/organisations/organisation/actions'

import { IATIOrganisationReport, IATIOrganisationReportProps } from '../../store/IATI/IATIReader/organisations/organisation/types'
import { OrganisationProps } from '../../store/IATI/types'

interface OrganisationFormProps {
  organisationsRef: string
  changeFunction: Function
  setValue: Function
  name: string
  label: string
}

interface OrganisationDataProps {
  organisation: IATIOrganisationReport
}

interface OrganisationDispatchProps {
  getOrganisation: (organisationsRef: string) => void
}

type OrganisationPickerProps = OrganisationFormProps & OrganisationDataProps & OrganisationDispatchProps

class Organisation extends React.Component<OrganisationPickerProps> {

  constructor (props: OrganisationPickerProps) {
   super(props)
  }

  componentDidUpdate(previousProps: OrganisationPickerProps) {
    //console.log('Organisations: ', this.props.organisationsRef)
    if(this.props.organisationsRef != "" &&  previousProps.organisationsRef != this.props.organisationsRef) {
      this.props.getOrganisation(this.props.organisationsRef)
      //console.log('Done Updating: ', this.props.organisationsRef)
    }
  }

  render() {

    //console.log ('rendering', this.props.organisation, this.props.organisationsRef)
    let organisationRefs: any[] = [{ value: "", label: "" }]
    //console.log(this.props.organisationsRef)
    //console.log(this.props.organisation[this.props.organisationsRef])
    if ( ( this.props.organisationsRef != "" ) &&
         ( typeof this.props.organisation[this.props.organisationsRef] !== 'undefined' ) &&
         ( this.props.organisation[this.props.organisationsRef].hasOwnProperty('data') ) ) {
      Object.keys(this.props.organisation[this.props.organisationsRef].data).forEach((organisationKey) => {
        //console.log('Org key: ', organisationKey)
        organisationRefs.push({ value: organisationKey, label: organisationKey })
      })
    }

    return (
      <React.Fragment>
        <Field
          name={this.props.name}
          label={this.props.label}
          component={Select}
          onChange={(ev: any) => {
            this.props.changeFunction(ev.target.value)
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
    organisation: state.organisationReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrganisationDispatchProps => {
  return {
    getOrganisation: (organisationsRef: string) => dispatch(getOrganisation({organisationsRef: organisationsRef}))
  }
}

export const OrganisationPicker = connect<OrganisationDataProps, OrganisationDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(Organisation)
