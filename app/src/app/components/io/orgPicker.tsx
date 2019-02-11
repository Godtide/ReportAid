import * as React from 'react'

import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../store'
import { ActionProps } from '../../store/types'

import { Field, ErrorMessage} from 'formik'
import { Select } from "material-ui-formik-components"

import { getOrgs } from '../../store/IATI/IATIReader/organisations/orgs/actions'
import { IATIOrgProps } from '../../store/IATI/types'

interface OrgProps {
  name: string
  label: string
}

interface OrgDataProps {
  orgs: IATIOrgProps
}

interface OrgDispatchProps {
  getOrgs: () => void
}

type OrgPickerProps = OrgProps & OrgDataProps & OrgDispatchProps

class Org extends React.Component<OrgPickerProps> {

  constructor (props: OrgPickerProps) {
   super(props)
  }

  componentDidMount() {
    this.props.getOrgs()
  }

  render() {

    let orgRefs : any[] = [{ value: "", label: "" }]
    Object.keys(this.props.orgs).forEach((orgKey) => {
      orgRefs.push({ value: orgKey, label: orgRef[orgKey].name })
    })

    return (
      <React.Fragment>
        <Field
          name={this.props.name}
          label={this.props.label}
          component={Select}
          options={orgRefs}
        />
        <ErrorMessage name={this.props.name} />
      </React.Fragment>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrgDataProps => {
  //console.log(state.orgReader)
  return {
    orgs: state.orgsReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgDispatchProps => {
  return {
    getOrgs: () => dispatch(getOrgs())
  }
}

export const OrgPicker = connect<OrgDataProps, OrgDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(Org)
