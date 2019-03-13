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
import { getOrganisations } from '../../store/IATI/IATIReader/organisations/organisations/actions'

import { Keys, KeyTypes } from '../../store/helpers/keys/types'

interface OrganisationsFormProps {
  setValue: Function
  name: string
  label: string
}

interface OrganisationsDataProps {
  organisationsRefs: Array<string>
}

interface OrganisationsDispatchProps {
  getOrganisationsRefs: () => void
  setOrganisationsKey: (organisationsRef: string) => void
}

type OrganisationsPickerProps = OrganisationsFormProps & OrganisationsDataProps & OrganisationsDispatchProps

class Organisations extends React.Component<OrganisationsPickerProps> {

  constructor (props: OrganisationsPickerProps) {
   super(props)
  }

  componentDidMount() {
    this.props.getOrganisationsRefs()
  }

  render() {

    let organisationsRefs: any[] = [{ value: "", label: "" }]
    if ( typeof this.props.organisationsRefs != 'undefined' &&
         typeof this.props.organisationsRefs[0] != 'undefined' &&
         this.props.organisationsRefs[0] != "" ) {
      this.props.organisationsRefs.forEach((key: string) => {
       organisationsRefs.push({ value: key, label: key })
      })
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
    organisationsRefs: state.organisationsPicker.data as Array<string>
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrganisationsDispatchProps => {
  return {
    getOrganisationsRefs: () => dispatch(getOrganisations()),
    setOrganisationsKey: (organisationsRef: string) => dispatch(setKey({key: organisationsRef, keyType: KeyTypes.ORGANISATIONS})),
  }
}

export const OrganisationsPicker = connect<OrganisationsDataProps, OrganisationsDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(Organisations)
