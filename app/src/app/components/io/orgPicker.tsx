import * as React from 'react'

import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../store'
import { ActionProps } from '../../store/types'

import { Field, ErrorMessage} from 'formik'
import { Select } from "material-ui-formik-components"

import { getOrgs } from '../../store/IATI/IATIReader/organisations/orgs/actions'
import { IATIOrgReport, IATIOrgData } from '../../store/IATI/IATIReader/organisations/types'

interface OrgProps {
  name: string
  label: string
}

interface OrgDataProps {
  orgs: IATIOrgReport
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

    const orgs: Array<IATIOrgData> = this.props.orgs.data
    console.log('Orgs: ', orgs)
    let orgRefs : any[] = [{ value: "", label: "" }]
    for (let i = 0; i < orgs.length; i++) {
      const label = orgs[i].name
      const value = orgs[i].orgRef
      orgRefs.push({ value: value, label: label })
    }

    return (
      <React.Fragment>
        <Field
          name={this.props.name}
          label={this.props.label}
          component={Select}
          options={orgRefs}
        />
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
