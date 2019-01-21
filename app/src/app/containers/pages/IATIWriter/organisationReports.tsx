import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { Formik, Form, Field, FormikProps, ErrorMessage} from 'formik'
import * as Yup from 'yup'

import { ApplicationState } from '../../../store'
import { ActionProps, TxData } from '../../../store/types'
import { IATIOrgProps, OrgReportProps } from '../../../store/IATI/types'

import { getOrgs } from '../../../store/IATI/IATIReader/organisation/actions'
import { OrgData } from '../../../store/IATI/IATIReader/organisation/types'

//import { getDictEntries } from '../../../utils/dict'

import { setOrganisationReport } from '../../../store/IATI/IATIWriter/organisationReports/actions'

import { LinearProgress } from '@material-ui/core'
import Button from '@material-ui/core/Button'
//import { Select } from 'formik-material-ui'
import { Select } from "material-ui-formik-components"

import { Organisation, OrganisationReport, Transaction } from '../../../utils/strings'
import { Helpers } from '../../../utils/config'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

const reportSchema = Yup.object().shape({
  version: Yup
    .string()
    .required('Required'),
  orgRef: Yup
    .string()
    .required('Required'),
  reportingOrgRef: Yup
    .string()
    .required('Required'),
  reportingOrgType: Yup
    .number()
    .required('Required'),
  reportingOrgIsSecondary: Yup
    .boolean()
    .required('Required'),
  lang: Yup
    .string()
    .required('Required'),
  currency: Yup
    .string()
    .required('Required'),
})

interface OrgProps {
  tx: TxData,
  orgs: OrgData
}

export interface OrgReportsDispatchProps {
  handleSubmit: (values: any) => void
  getOrgs: () => void
}

type OrgReportsFormProps = WithStyles<typeof styles> & OrgProps & OrgReportsDispatchProps

export class OrgReportsForm extends React.Component<OrgReportsFormProps> {

  state = {
    txKey: '',
    txSummary: '',
    toggleSubmitting: false,
    submitFunc: (function(submit: boolean) { return submit }),
    resetFunc: (function() { return null })
  }

  constructor (props: OrgReportsFormProps) {
   super(props)
  }

  componentDidMount() {
    this.props.getOrgs()
  }

  componentDidUpdate(previousProps: OrgReportsFormProps) {
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

    let versions: any = []
    Helpers.reportVersions.forEach( (value: any) => {
      //console.log(value, value.code)
      versions.push({ value: value, label: value })
    })

    //console.log('Props orgs: ', this.props.orgs)
    let orgRefs: any = []
    Object.keys(this.props.orgs).forEach((key) => {
      orgRefs.push({ value: key, label: this.props.orgs[key].name })
    })

    let orgCodes: any = []
    Helpers.organisationCodes.forEach( (value: any) => {
      //console.log(value, value.code)
      orgCodes.push({ value: value.code, label: value.type })
    })

    let isSecondary = [
      { value: false, label: 'No' },
      { value: true, label: 'Yes' }
    ]

    let countries: any = []
    Helpers.countryCodes.forEach( (value: any) => {
      //console.log(value, value.code)
      countries.push({ value: value, label: value })
    })

    let currencies: any = []
    Helpers.currencyCodes.forEach( (value: any) => {
      //console.log(value, value.code)
      currencies.push({ value: value, label: value })
    })

    return (
      <div>
        <h2>{OrganisationReport.headingOrgReportWriter}</h2>
        <div>
          <Formik
            initialValues={ {version: versions[0].value,
                             orgRef: orgRefs[0].value,
                             reportingOrgRef: orgRefs[0].value,
                             reportingOrgType: orgCodes[0].value,
                             reportingOrgIsSecondary: isSecondary[0].value,
                             lang: countries[0].value,
                             currency: currencies[0].value
                            }}
            validationSchema={reportSchema}
            onSubmit={(values: OrgReportProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<OrgReportProps>) => (
              <Form>
                <Field
                  name="version"
                  label={OrganisationReport.version}
                  component={Select}
                  options={versions}
                />
                <ErrorMessage name='version' />
                <br />
                <Field
                  name="orgRef"
                  label={OrganisationReport.orgIdentifier}
                  component={Select}
                  options={orgRefs}
                />
                <ErrorMessage name='orgRef' />
                <br />
                <Field
                  name="reportingOrgRef"
                  label={OrganisationReport.reportingOrgRef}
                  component={Select}
                  options={orgRefs}
                />
                <ErrorMessage name='reportingOrgRef' />
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
                  name="lang"
                  label={OrganisationReport.language}
                  component={Select}
                  options={countries}
                />
                <ErrorMessage name='lang' />
                <br />
                <Field
                  name="currency"
                  label={OrganisationReport.currency}
                  component={Select}
                  options={currencies}
                />
                <ErrorMessage name='currency' />
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

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgReportsDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setOrganisationReport(ownProps)),
    getOrgs: () => dispatch(getOrgs())
  }
}

export const OrganisationReports = withTheme(withStyles(styles)(connect<OrgProps, OrgReportsDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrgReportsForm)))
