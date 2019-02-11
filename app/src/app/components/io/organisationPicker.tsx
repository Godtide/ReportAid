import * as React from 'react'

import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../store'
import { ActionProps } from '../../store/types'

import { Field, ErrorMessage} from 'formik'
import { Select } from "material-ui-formik-components"

import { getOrganisation } from '../../store/IATI/IATIReader/organisations/organisation/actions'

import { OrganisationProps } from '../../store/IATI/types'

interface OrganisationFormProps {
  name: string
  label: string
}

interface OrganisationDataProps {
  organisations: OrganisationProps
}

interface OrganisationDispatchProps {
  getOrganisation: () => void
}

type OrganisationPickerProps = OrganisationFormProps & OrganisationDataProps & OrganisationDispatchProps

class Organisation extends React.Component<OrganisationPickerProps> {

  constructor (props: OrganisationPickerProps) {
   super(props)
  }

  componentDidMount() {
    this.props.getOrganisation()
  }

  validateOrganisation (value: object) {
    let error
    if (!(value.hasOwnProperty('organisationRef'))) {
      error = 'Required!'
    }
    return error
  }

  render() {

    let organisationRefs: any[] = [{ value: {} as OrganisationProps, label: "" }]
     Object.keys(this.props.organisations).forEach((key) => {
      //console.log(orgKey)
      const values = Object.values(this.props.organisations[key])
      //console.log(values)
      Object.keys(values[1]).forEach((organisationKey) => {
        //console.log('Key: ', reportKey)
        //const ref: OrganisationProps = { organisationRef: organisationKey}
        organisationRefs.push({ value: organisationKey, label: organisationKey })
      })
    })

    //console.log('Refs: ', reportRefs)

    return (
      <React.Fragment>
        <Field
          name={this.props.name}
          label={this.props.label}
          validate={this.validateOrganisation}
          component={Select}
          options={organisationRefs}
        />
        <ErrorMessage name={this.props.name} />
      </React.Fragment>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrganisationDataProps => {
  //console.log(state.orgReader)
  return {
    organisations: state.organisationsReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrganisationDispatchProps => {
  return {
    getOrganisation: () => dispatch(getOrganisation())
  }
}

export const OrganisationPicker = connect<OrganisationDataProps, OrganisationDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(Organisation)
