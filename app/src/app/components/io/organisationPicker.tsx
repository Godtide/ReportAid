import * as React from 'react'
import * as Yup from 'yup'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../store'
import { ActionProps } from '../../store/types'

import { Field, ErrorMessage} from 'formik'
import { Select } from "material-ui-formik-components"

import { getOrganisation } from '../../store/IATI/IATIReader/organisations/organisation/actions'

import { IATIOrganisationsData, IATIOrganisationReportProps } from '../../store/IATI/IATIReader/organisations/types'
import { OrganisationProps } from '../../store/IATI/types'

const organisationSchema = Yup.object().shape({
  organisation: Yup
    .string()
    .required('Required')
})

interface OrganisationFormProps {
  organisationsRef: string
  label: string
}

interface OrganisationDataProps {
  organisations: IATIOrganisationsData
}

interface OrganisationDispatchProps {
  getOrganisation: (organisationsRef: string) => void
}

type OrganisationPickerProps = OrganisationFormProps & OrganisationDataProps & OrganisationDispatchProps

class Organisation extends React.Component<OrganisationPickerProps> {

  constructor (props: OrganisationPickerProps) {
   super(props)
  }

  componentDidMount() {
    if(this.props.organisationsRef != "") {
      this.props.getOrganisation(this.props.organisationsRef)
    }
  }

  render() {

    let organisationRefs: any[] = [{ value: "", label: "" }]
    const organisationReports: IATIOrganisationReportProps = this.props.organisations[this.props.organisationsRef].data
    Object.keys(organisationReports).forEach((organisationKey) => {
       organisationRefs.push({ value: organisationKey, label: organisationKey })
    })

    return (
      <React.Fragment>
        <Field
          name='organisation'
          label={this.props.label}
          validate={organisationSchema}
          component={Select}
          options={organisationRefs}
        />
        <ErrorMessage name='organisation' />
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
    getOrganisation: (organisationsRef: string) => dispatch(getOrganisation({organisationsRef: organisationsRef}))
  }
}

export const OrganisationPicker = connect<OrganisationDataProps, OrganisationDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(Organisation)
