import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { Formik, Form, Field, FormikProps, ErrorMessage} from 'formik'
import * as Yup from 'yup'

import { ApplicationState } from '../../../store'
import { ActionProps, TxData } from '../../../store/types'
import { IATIOrgProps, OrgReportProps } from '../../../store/IATI/types'

import { getOrgs } from '../../../store/IATI/IATIReader/organisationReader/actions'
import { OrgData } from '../../../store/IATI/IATIReader/organisationReader/types'

//import { getDictEntries } from '../../../utils/dict'

import { setOrganisationReport } from '../../../store/IATI/IATIWriter/organisationReportsWriter/actions'

import { LinearProgress } from '@material-ui/core'
import Button from '@material-ui/core/Button'
//import { Select } from 'formik-material-ui'
import { Select } from "material-ui-formik-components"

import { Organisation, OrganisationReport, Transaction } from '../../../utils/strings'
import { Helpers } from '../../../utils/config'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

interface OrgProps {
  tx: TxData,
  orgs: OrgData
}

export interface OrgDispatchProps {
  handleSubmit: (values: any) => void
  getOrgs: () => void
}

/* reportingOrgIdentifier: '',
                 reportingOrgType: 0,
                 reportingOrgIsSecondary: false,
                 issuingOrgRef: '',
                 version: 'reportingOrgType',
                 lang: '',
                 currency: ''}
                 */

const reportSchema = Yup.object().shape({
  reportingOrgIdentifier: Yup
    .string()
    .required('Required'),
  reportingOrgType: Yup
    .string()
    .required('Required'),
  reportingOrgIsSecondary: Yup
    .string()
    .required('Required'),
  issuingOrgRef: Yup
    .string()
    .required('Required'),
  version: Yup
    .string()
    .required('Required'),
  lang: Yup
    .string()
    .required('Required'),
  currency: Yup
    .string()
    .required('Required'),
})

type OrgReportWriterFormProps = WithStyles<typeof styles> & OrgProps & OrgDispatchProps

export class OrgReportForm extends React.Component<OrgReportWriterFormProps> {

  state = {
    txKey: '',
    txSummary: '',
    toggleSubmitting: false,
    submitFunc: (function(submit: boolean) { return submit }),
    resetFunc: (function() { return null })
  }

  constructor (props: OrgReportWriterFormProps) {
   super(props)
  }

  componentDidMount() {
    this.props.getOrgs()
  }

  componentDidUpdate(previousProps: OrgReportWriterFormProps) {
    //console.log(this.props.tx)
    if(previousProps.tx != this.props.tx) {
      const submitting = !this.state.toggleSubmitting
      let txKey = ''
      let txSummary = `${Transaction.fail}`
      const txKeys = Object.keys(this.props.tx)
      if (txKeys.length > 0 ) {
        txKey = Object.keys(this.props.tx)[0]
        txSummary = `${Transaction.success}`
      }
      this.setState({txKey: txKey, txSummary: txSummary, toggleSubmitting: submitting})
      this.state.submitFunc(submitting)
      this.state.resetFunc()
    }
  }

  handleSubmit = (values: OrgReportProps, setSubmitting: Function, reset: Function) => {
    //console.log('Values: ', values)
    const submitting = !this.state.toggleSubmitting
    this.setState({txKey: '', txSummary: '', toggleSubmitting: submitting, submitFunc: setSubmitting, resetFunc: reset})
    setSubmitting(submitting)
    this.props.handleSubmit(values)
  }

  render() {

    //console.log('Props orgs: ', this.props.orgs)
    let orgRefs: any = []
    Object.keys(this.props.orgs).forEach((key) => {
      orgRefs.push({ value: key, label: this.props.orgs[key].name })
    })

    let isSecondary = [
      { value: 1, label: 'Yes' },
      { value: 0, label: 'No' }
    ]

    let orgCodes: any = []
    Helpers.organisationCodes.forEach( (value: any) => {
      //console.log(value, value.code)
      orgCodes.push({ value: value.code, label: value.type })
    })

    /*
    static reportingOrgIdentifier = 'Reporting Organisation Identifier'
    static reportingOrgType = "Reporting Organisation Type"
    static reportingOrgIsSecondary = "Is Reporting Organisation Secondary?"
    static issuingOrgIdentifier = "Issuing Organisation"
    static version = "Report Version"
    static language = "Report Language"
    static currency = "Report Currency"
    */

    return (
      <div>
        <h2>{OrganisationReport.headingOrgReportWriter}</h2>
        <div>
          <Formik
            initialValues={ {reportingOrgIdentifier: orgRefs[0].label,
                             reportingOrgType: orgCodes[0].label,
                             reportingOrgIsSecondary: isSecondary[0].value,
                             issuingOrgRef: '',
                             version: '',
                             lang: '',
                             currency: ''}
                           }
            validationSchema={reportSchema}
            onSubmit={(values: OrgReportProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<OrgReportProps>) => (
              <Form>
                <Field
                  name="reportingOrgIdentifier"
                  label={OrganisationReport.reportingOrgIdentifier}
                  component={Select}
                  options={orgRefs}
                />
                <ErrorMessage name='reportingOrgIdentifier' />
                <br />
                <Field
                  name="reportingOrgType"
                  label={OrganisationReport.reportingOrgType}
                  component={Select}
                  options={orgCodes}
                />
                <ErrorMessage name='reportingOrgType' />
                <br />
                <Field
                  name="reportingOrgIsSecondary"
                  label={OrganisationReport.reportingOrgIsSecondary}
                  component={Select}
                  options={isSecondary}
                />
                <ErrorMessage name='reportingOrgIsSecondary' />
                <br />
                <Field
                  name="issuingOrgRef"
                  label={OrganisationReport.issuingOrgIdentifier}
                  component={Select}
                  options={orgRefs}
                />
                <ErrorMessage name='issuingOrgRef' />
                <br />
                {formProps.isSubmitting && <LinearProgress />}
                <br />
                <Button type='submit' variant="raised" color="primary" disabled={formProps.isSubmitting}>
                  Submit
                </Button>
              </Form>
            )}
          />
        </div>
        <hr />
        <h3>{Transaction.heading}</h3>
        <p>
          <b>{Transaction.summary}</b>: {this.state.txSummary}<br />
          <b>{Transaction.key}</b>: {this.state.txKey}
        </p>
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrgProps => {
  //console.log(state.orgReader)
  return {
    tx: state.orgReportsForm.data,
    orgs: state.orgReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setOrganisationReport(ownProps)),
    getOrgs: () => dispatch(getOrgs())
  }
}

export const OrgReportWriter = withTheme(withStyles(styles)(connect<OrgProps, OrgDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrgReportForm)))
