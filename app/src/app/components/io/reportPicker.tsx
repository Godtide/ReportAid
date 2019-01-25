import * as React from 'react'

import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../store'
import { ActionProps } from '../../store/types'

import { Field, ErrorMessage} from 'formik'
import { Select } from "material-ui-formik-components"

import { getOrgReports } from '../../store/IATI/IATIReader/organisationReports/actions'
import { OrgReportData } from '../../store/IATI/IATIReader/organisationReports/types'

interface ReportProps {
  name: string
  label: string
}

interface ReportDataProps {
  orgReports: OrgReportData
}

interface ReportDispatchProps {
  getOrgReports: () => void
}

type ReportPickerProps = ReportProps & ReportDataProps & ReportDispatchProps

class ReportPicker extends React.Component<ReportPickerProps> {

  constructor (props: ReportPickerProps) {
   super(props)
  }

  componentDidMount() {
    this.props.getOrgReports()
  }

  render() {

    let reportRefs: any[] = [{ value: "", label: "" }]
     Object.keys(this.props.orgReports).forEach((orgKey) => {
      //console.log(orgKey)
      const values = Object.values(this.props.orgReports[orgKey])
      //console.log(values)
      Object.keys(values[1]).forEach((reportKey) => {
        //console.log('Key: ', reportKey)
        reportRefs.push({ value: reportKey, label: reportKey })
      })
    })

    return (
      <React.Fragment>
        <Field
          name={this.props.name}
          label={this.props.label}
          component={Select}
          options={reportRefs}
        />
        <ErrorMessage name={this.props.name} />
      </React.Fragment>
    )
  }
}

const mapStateToProps = (state: ApplicationState): ReportDataProps => {
  //console.log(state.orgReader)
  return {
    orgReports: state.orgReportsReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): ReportDispatchProps => {
  return {
    getOrgReports: () => dispatch(getOrgReports())
  }
}

export const OrgReportPicker = connect<ReportDataProps, ReportDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(ReportPicker)
