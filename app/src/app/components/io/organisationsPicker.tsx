import * as React from 'react'

import * as Yup from 'yup'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../store'
import { ActionProps } from '../../store/types'

import { Field, ErrorMessage, FieldProps } from 'formik'
import { Select } from "material-ui-formik-components"
//import MuiSelect from '@material-ui/core/Select'

import { setOrganisationsKey } from '../../store/helpers/keys/actions'
import { getOrganisations } from '../../store/IATI/IATIReader/organisations/organisations/actions'

import { IATIOrganisationsReport,
         IATIOrganisationsData } from '../../store/IATI/IATIReader/organisations/types'

interface OrganisationsFormProps {
  setValue: Function
  name: string
  label: string
}

interface OrganisationsDataProps {
  organisations: IATIOrganisationsReport
}

interface OrganisationsDispatchProps {
  getOrganisations: () => void
  setOrganisationsKey: (organisationsRef: string) => void
}

type OrganisationsPickerProps = OrganisationsFormProps & OrganisationsDataProps & OrganisationsDispatchProps

class Organisations extends React.Component<OrganisationsPickerProps> {

  constructor (props: OrganisationsPickerProps) {
   super(props)
  }

  componentDidMount() {
    this.props.getOrganisations()
  }

  render() {

    //console.log ('rendering', this.props.organisation, this.props.organisationsRef)
    let organisationsRefs: any[] = [{ value: "", label: "" }]
    const organisations: Array<IATIOrganisationsData> = this.props.organisations.data
    for (let i = 0; i < organisations.length; i++) {
     const label = organisations[i].organisationsRef
     const value = label
     organisationsRefs.push({ value: value, label: label })
    }

    return (
      <React.Fragment>
        <Field
          name={this.props.name}
          label={this.props.label}
          component={Select}
          onChange={(ev: any) => {
            this.props.setOrganisationsKey(ev.target.value)
            this.props.setValue(this.props.name, ev.target.value)
          }}
          options={organisationsRefs}
        />
      </React.Fragment>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrganisationsDataProps => {
  //console.log(state.orgReader)
  return {
    organisations: state.organisationsReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrganisationsDispatchProps => {
  return {
    getOrganisations: () => dispatch(getOrganisations()),
    setOrganisationsKey: (organisationsRef: string) => dispatch(setOrganisationsKey(organisationsRef)),
  }
}

export const OrganisationsPicker = connect<OrganisationsDataProps, OrganisationsDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(Organisations)
