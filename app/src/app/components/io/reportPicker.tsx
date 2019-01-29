import * as React from 'react'

import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../store'
import { ActionProps } from '../../store/types'

import { Field, ErrorMessage} from 'formik'
import { Select } from "material-ui-formik-components"

import { getOrgReports } from '../../store/IATI/IATIReader/organisationReports/actions'

import { OrgReportData } from '../../store/IATI/IATIReader/organisationReports/types'
import { ReportProps } from '../../store/IATI/types'

interface ReportFormProps {
  name: string
  label: string
}

interface ReportDataProps {
  orgReports: OrgReportData
}

interface ReportDispatchProps {
  getOrgReports: () => void
}

type ReportPickerProps = ReportFormProps & ReportDataProps & ReportDispatchProps

class ReportPicker extends React.Component<ReportPickerProps> {

  constructor (props: ReportPickerProps) {
   super(props)
  }

  componentDidMount() {
    this.props.getOrgReports()
  }

  validateOrgReport (value: object) {
    let error
    if (!(value.hasOwnProperty('orgRef')) &&
        !(value.hasOwnProperty('reportRef')) ) {
      error = 'Required!'
    }
    return error
  }

  render() {

    let reportRefs: any[] = [{ value: {} as ReportProps, label: "" }]
     Object.keys(this.props.orgReports).forEach((orgKey) => {
      //console.log(orgKey)
      const values = Object.values(this.props.orgReports[orgKey])
      //console.log(values)
      Object.keys(values[1]).forEach((reportKey) => {
        //console.log('Key: ', reportKey)
        const report: ReportProps = { orgRef: orgKey, reportRef: reportKey}
        reportRefs.push({ value: report, label: reportKey })
      })
    })

    //console.log('Refs: ', reportRefs)

    return (
      <React.Fragment>
        <Field
          name={this.props.name}
          label={this.props.label}
          validate={this.validateOrgReport}
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
