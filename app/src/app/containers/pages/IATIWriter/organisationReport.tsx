import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { withFormik, Formik, Form, Field, FormikProps, ErrorMessage} from 'formik'
import * as Yup from 'yup'

import { ApplicationState } from '../../../store'
import { ActionProps, TxData } from '../../../store/types'
import { IATIOrgProps, OrgReportProps } from '../../../store/IATI/types'

import { getOrgs } from '../../../store/IATI/IATIReader/organisationReader/actions'
import { OrgData } from '../../../store/IATI/IATIReader/organisationReader/types'

import { getDictEntries } from '../../../utils/dict'

import { setOrganisationReport } from '../../../store/IATI/IATIWriter/organisationReportsWriter/actions'

import { LinearProgress } from '@material-ui/core'
import Button from '@material-ui/core/Button'
//import { Select } from 'formik-material-ui'

import { OrganisationReport, Transaction } from '../../../utils/strings'

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

const orgReportSchema = Yup.object().shape({
  orgIdentifier: Yup
    .string()
    /*.matches(/^.*[^-].*$/, {
        message: 'Please select an organisation identifier',
        excludeEmptyString: true
    })*/
    .required('Required'),
  reportingOrgIdentifier: Yup
    .string()
    /*.matches(/^.*[^-].*$/, {
        message: 'Please select an organisation identifier',
        excludeEmptyString: true
    })*/
    .required('Required')
})

type OrgReportWriterFormProps = WithStyles<typeof styles> & OrgProps & OrgDispatchProps


/*const InnerForm = (
    props: any,
    values: OrgReportProps,
    errors: any,
    touched: any,
    setFieldTouched: any,
    setFieldValue: any,
    isSubmitting: any,
    handleSubmit: any
  ) => {

    const orgs = getDictEntries(this.props.orgs) as IATIOrgProps[]
    //console.log('Blah orgs ', orgs)
    const iDs = orgs.map((value: any) => (
      value.identifier
    ))
    const fields = {orgs: iDs}

    return (
      <Form onSubmit={handleSubmit}>
        <label htmlFor='orgIdentifier'>{OrganisationReport.orgIdentifier}: </label>
        <Field
          name='orgIdentifier'
          render={ (props: any) => {
            console.log('Values ', values)
            const defaultOption = <option key='' value=''>Please Select:</option>
            const options = fields.orgs.map((value: any, index: any) => <option key={index} value={value}>{value}</option> )
            const selectOptions = [defaultOption, ...options]
            return (
              <div>
                <select
                  value={values.orgIdentifier}
                  onChange={(value) => setFieldValue('orgIdentifier', value)}
                  {...props}
                >
                  {
                    selectOptions
                  }
                </select>
              </div>
            )
          }}
        />
        <ErrorMessage name='orgIdentifier' />
        <br />
        <label htmlFor='reportingOrgIdentifier'>{OrganisationReport.reportingOrgIdentifier}: </label>
        <Field
          name='reportingOrgIdentifier'
          render={ (props: any) => {
            //console.log('Props ', props)
            const defaultOption = <option key='' value=''>Please Select:</option>
            const options = fields.orgs.map((value: any, index: any) => <option key={index} value={value}>{value}</option> )
            const selectOptions = [defaultOption, ...options]
            return (
              <div>
                <select
                  value={this.state.reportingOrgIdentifier}
                  onChange={this.handleReportingOrgChange}
                  {...props}
                >
                  {
                    selectOptions
                  }
                </select>
              </div>
            )
          }}
        />
        <ErrorMessage name='reportingOrgIdentifier' />
        <br />
        {isSubmitting && <LinearProgress />}
        <br />
        <Button type='submit' variant="raised" color="primary" disabled={isSubmitting}>
          Submit
        </Button>
      </Form>
    )
}*/

export class OrgReportForm extends React.Component<OrgReportWriterFormProps> {

  state = {
    txKey: '',
    txSummary: '',
    orgIdentifier: '',
    reportingOrgIdentifier: '',
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

  handleOrgChange = (event: any) => {
    console.log('Org hell ', event.target.value)
    const thisState = event.target.value
    this.setState({ orgIdentifier: thisState })
  }

  handleReportingOrgChange = (event: any) => {
    console.log('reporting hell ', event.target.value)
    const thisState = event.target.value
    this.setState({ reportingOrgIdentifier: thisState })
  }

  handleSubmit = (values: any, setSubmitting: Function, reset: Function) => {
    const submitting = !this.state.toggleSubmitting
    this.setState({txKey: '', txSummary: '', toggleSubmitting: submitting, submitFunc: setSubmitting, resetFunc: reset})
    setSubmitting(submitting)
    const orgReportDetails: OrgReportProps = {
      orgIdentifier: values.org,
      reportingOrgIdentifier: values.reportingOrg,
      version: '2.03'
    }
    console.log('Org report ', orgReportDetails)
    this.props.handleSubmit(orgReportDetails)
  }

  render() {

    const orgs = getDictEntries(this.props.orgs) as IATIOrgProps[]
    //console.log('Blah orgs ', orgs)
    const iDs = orgs.map((value: any) => (
      value.identifier
    ))
    const fields = {orgs: iDs}

    return (
      <div>
        <h2>{OrganisationReport.headingOrgReportWriter}</h2>
        <div>
          <Formik
            initialValues={ {orgIdentifier: '', reportingOrgIdentifier: '', version: '' } }
            validationSchema={orgReportSchema}
            onSubmit={(values: OrgReportProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={ ({values, errors, handleSubmit, handleChange, isSubmitting, setFieldValue}: any) => (
              <Form>
                <label htmlFor='orgIdentifier'>{OrganisationReport.orgIdentifier}: </label>
                <Field
                  name='orgIdentifier'
                  render={ (props: any) => {
                    console.log('Values ', values)
                    const defaultOption = <option key='' value=''>Please Select:</option>
                    const options = fields.orgs.map((value: any, index: any) => <option key={index} value={value}>{value}</option> )
                    const selectOptions = [defaultOption, ...options]
                    return (
                      <div>
                        <select
                          value={values.orgIdentifier}
                          onChange={(value) => setFieldValue('orgIdentifier', value)}
                          {...props}
                        >
                          {
                            selectOptions
                          }
                        </select>
                      </div>
                    )
                  }}
                />
                <ErrorMessage name='orgIdentifier' />
                <br />
                <label htmlFor='reportingOrgIdentifier'>{OrganisationReport.reportingOrgIdentifier}: </label>
                <Field
                  name='reportingOrgIdentifier'
                  render={ (props: any) => {
                    //console.log('Props ', props)
                    const defaultOption = <option key='' value=''>Please Select:</option>
                    const options = fields.orgs.map((value: any, index: any) => <option key={index} value={value}>{value}</option> )
                    const selectOptions = [defaultOption, ...options]
                    return (
                      <div>
                        <select
                          value={values.reportingOrgIdentifier}
                          onChange={(value) => setFieldValue('reportingOrgIdentifier', value)}
                          {...props}
                        >
                          {
                            selectOptions
                          }
                        </select>
                      </div>
                    )
                  }}
                />
                <ErrorMessage name='reportingOrgIdentifier' />
                <br />
                {isSubmitting && <LinearProgress />}
                <br />
                <Button type='submit' variant="raised" color="primary" disabled={isSubmitting}>
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
