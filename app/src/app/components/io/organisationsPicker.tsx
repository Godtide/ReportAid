import * as React from 'react'

import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../store'
import { ActionProps } from '../../store/types'

import { Field, ErrorMessage} from 'formik'
import { Select } from "material-ui-formik-components"

import { getOrganisations } from '../../store/IATI/IATIReader/organisations/organisations/actions'

import { OrganisationsProps } from '../../store/IATI/types'

interface OrganisationsFormProps {
  name: string
  label: string
}

interface OrganisationsDataProps {
  organisations: OrganisationsProps
}

interface OrganisationsDispatchProps {
  getOrganisations: () => void
}

type OrganisationsPickerProps = OrganisationsFormProps & OrganisationsDataProps & OrganisationsDispatchProps

class OrganisationsPicker extends React.Component<OrganisationsPickerProps> {

  constructor (props: OrganisationsPickerProps) {
   super(props)
  }

  componentDidMount() {
    this.props.getOrgOrganisationss()
  }

  validateOrgOrganisations (value: object) {
    let error
    if (!(value.hasOwnProperty('orgRef')) &&
        !(value.hasOwnProperty('reportRef')) ) {
      error = 'Required!'
    }
    return error
  }

  render() {

    let organisationsRefs: any[] = [{ value: {} as OrganisationsProps, label: "" }]
     Object.keys(this.props.orgOrganisationss).forEach((orgKey) => {
      //console.log(orgKey)
      const values = Object.values(this.props.orgOrganisationss[orgKey])
      //console.log(values)
      Object.keys(values[1]).forEach((reportKey) => {
        //console.log('Key: ', reportKey)
        const report: OrganisationsProps = { orgRef: orgKey, reportRef: reportKey}
        reportRefs.push({ value: report, label: reportKey })
      })
    })

    //console.log('Refs: ', reportRefs)

    return (
      <React.Fragment>
        <Field
          name={this.props.name}
          label={this.props.label}
          validate={this.validateOrgOrganisations}
          component={Select}
          options={reportRefs}
        />
        <ErrorMessage name={this.props.name} />
      </React.Fragment>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrganisationsDataProps => {
  //console.log(state.orgReader)
  return {
    orgOrganisationss: state.orgOrganisationssReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrganisationsDispatchProps => {
  return {
    getOrgOrganisationss: () => dispatch(getOrgOrganisationss())
  }
}

export const OrgOrganisationsPicker = connect<OrganisationsDataProps, OrganisationsDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrganisationsPicker)
