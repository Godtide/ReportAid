import * as React from 'react'

import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../store'
import { ActionProps } from '../../store/types'

import { Field, ErrorMessage} from 'formik'
import { Select } from "material-ui-formik-components"

//import { setKey } from '../store/helpers/keys/actions'
import { getOrgs } from '../../store/IATI/IATIReader/organisations/orgs/actions'
import { IATIOrgReport, IATIOrgData } from '../../store/IATI/types'

interface OrgProps {
  name: string
  label: string
}

interface OrgDataProps {
  orgs: IATIOrgReport
}

interface OrgDispatchProps {
  getOrgs: (isReport: boolean) => void
}

type OrgPickerProps = OrgProps & OrgDataProps & OrgDispatchProps

class Org extends React.Component<OrgPickerProps> {

  constructor (props: OrgPickerProps) {
   super(props)
  }

  componentDidMount() {
    this.props.getOrgs(false)
  }

  render() {

    let orgRefs : any[] = [{ value: "", label: "" }]
    if(typeof this.props.orgs != "undefined" &&
       this.props.orgs.hasOwnProperty('data')) {

      //console.log('Orgs: ', this.props.orgs)
      const orgs: Array<IATIOrgData> = this.props.orgs.data
      for (let i = 0; i < orgs.length; i++) {
        const label = orgs[i].name
        const value = orgs[i].orgRef
        orgRefs.push({ value: value, label: label })
      }
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
    orgs: state.orgsPicker.data as IATIOrgReport
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgDispatchProps => {
  return {
    getOrgs: (isReport: boolean) => dispatch(getOrgs(isReport))
  }
}

export const OrgPicker = connect<OrgDataProps, OrgDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(Org)
