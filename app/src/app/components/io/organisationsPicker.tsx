import * as React from 'react'

import * as Yup from 'yup'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../store'
import { ActionProps } from '../../store/types'

import { Field, ErrorMessage } from 'formik'
import { Select } from "material-ui-formik-components"

import { getOrganisations } from '../../store/IATI/IATIReader/organisations/organisations/actions'

import { IATIOrganisationsData } from '../../store/IATI/IATIReader/organisations/types'

interface OrganisationsFormProps {
  changeFunction: Function
  name: string
  label: string
}

interface OrganisationsDataProps {
  organisations: IATIOrganisationsData
}

interface OrganisationsDispatchProps {
  getOrganisations: () => void
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

    let organisationsRefs: any[] = [{ value: "", label: "" }]
    Object.keys(this.props.organisations).forEach((organisationsKey) => {
      organisationsRefs.push({ value: organisationsKey, label: organisationsKey })
    })

    //console.log('Refs: ', reportRefs)

    return (
      <React.Fragment>
        <Field
          name={this.props.name}
          label={this.props.label}
          component={Select}
          onChange={(ev: any) => {
            this.props.changeFunction(ev.target.value)
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
    getOrganisations: () => dispatch(getOrganisations())
  }
}

export const OrganisationsPicker = connect<OrganisationsDataProps, OrganisationsDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(Organisations)
