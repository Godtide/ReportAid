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

import { IATIOrganisationsReport,
         IATIOrganisationsData } from '../../store/IATI/types'

interface OrganisationsFormProps {
  setValue: Function
  name: string
  label: string
}

interface OrganisationsDataProps {
  organisations: IATIOrganisationsReport
}

interface OrganisationsDispatchProps {
  getOrganisations: (isReport: boolean) => void
  setOrganisationsKey: (organisationsRef: string) => void
}

type OrganisationsPickerProps = OrganisationsFormProps & OrganisationsDataProps & OrganisationsDispatchProps

class Organisations extends React.Component<OrganisationsPickerProps> {

  constructor (props: OrganisationsPickerProps) {
   super(props)
  }

  componentDidMount() {
    this.props.getOrganisations(false)
  }

  render() {

    let organisationsRefs: any[] = [{ value: "", label: "" }]
    if(typeof this.props.organisations != "undefined" &&
       this.props.organisations.hasOwnProperty('data')) {

      //console.log ('rendering', this.props.organisations)
      const organisations: Array<IATIOrganisationsData> = this.props.organisations.data
      for (let i = 0; i < organisations.length; i++) {
       const label = organisations[i].organisationsRef
       const value = label
       organisationsRefs.push({ value: value, label: label })
      }
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
    organisations: state.organisationsPicker.data as IATIOrganisationsReport
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrganisationsDispatchProps => {
  return {
    getOrganisations: (isReport: boolean) => dispatch(getOrganisations(isReport)),
    setOrganisationsKey: (organisationsRef: string) => dispatch(setKey({key: organisationsRef, keyType: KeyTypes.ORGANISATIONS})),
  }
}

export const OrganisationsPicker = connect<OrganisationsDataProps, OrganisationsDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(Organisations)
