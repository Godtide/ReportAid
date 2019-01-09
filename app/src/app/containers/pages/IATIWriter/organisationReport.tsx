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

import { getDictEntries } from '../../../utils/dict'

import { setOrganisationReport } from '../../../store/IATI/IATIWriter/organisationReportsWriter/actions'

import { LinearProgress } from '@material-ui/core'
import Button from '@material-ui/core/Button'
import { Select } from 'formik-material-ui'

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

const organisationSchema = Yup.object().shape({
  orgIdentifier: Yup
    .string()
    .required('Required')
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
    const submitting = !this.state.toggleSubmitting
    this.setState({txKey: '', txSummary: '', toggleSubmitting: submitting, submitFunc: setSubmitting, resetFunc: reset})
    setSubmitting(submitting)
    this.props.handleSubmit(values)
  }

  render() {

    return (
      <div>
        <h2>{OrganisationReport.headingOrgReportWriter}</h2>
        <div>
          <Formik
            initialValues={ {orgIdentifier: '', reportingOrgIdentifier: '', version: ''} }
            validationSchema={organisationSchema}
            onSubmit={(values: OrgReportProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<OrgReportProps>) => (
              <Form>
                <Field
                  name='orgIdentifier'
                  label={OrganisationReport.orgIdentifier}
                  component={Select}
                />
                <ErrorMessage name='orgIdentifier' />
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


/* import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { withFormik, ErrorMessage} from 'formik'
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
import { Select } from 'formik-material-ui'

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

type OrgReportWriterFormProps = WithStyles<typeof styles> & OrgProps & OrgDispatchProps

const MyForm = (props: any) => {
  const {
    values,
    touched,
    dirty,
    errors,
    handleChange,
    handleBlur,
    handleSubmit,
    handleReset,
    setFieldValue,
    setFieldTouched,
    isSubmitting
  } = props
  return (
    <form onSubmit={handleSubmit} onReset={handleReset}>
      <MySelect
        label={`${OrganisationReport.orgIdentifier}:`}
        name='orgIdentifier'
        value={values.orgIdentifier}
        onChange={setFieldValue}
        onBlur={setFieldTouched}
        options={{blah: 'blah'}}
      />
      <MySelect
        label={`${OrganisationReport.reportingOrgIdentifier}:`}
        name='orgIdentifier'
        value={values.reportingOrgIdentifier}
        onChange={setFieldValue}
        onBlur={setFieldTouched}
        options={{blah: 'blah'}}
      />
      <br />
      {isSubmitting && <LinearProgress />}
      <br />
      <Button type='submit' variant="raised" color="primary" disabled={isSubmitting}>
        Submit
      </Button>
    </form>
  )
}

interface OrgFormProps {
  orgIdentifier: string,
  reportingOrgIdentifier: string,
  version: string,
  options: object
  handleSubmit: (values: any, setSubmitting: Function, setReset: Function) => void
}


const orgForm = withFormik({
  validationSchema: Yup.object().shape({
    orgIdentifier: Yup
      .string()
      .matches(/^.*[^-].*$/, {
          message: 'Please select an organisation identifier',
          excludeEmptyString: true
      })
      .required('Required'),
    reportingOrgIdentifier: Yup
      .string()
      .matches(/^.*[^-].*$/, {
          message: 'Please select an organisation identifier',
          excludeEmptyString: true
      })
      .required('Required')
  }),
  mapPropsToValues: (props: OrgFormProps) => ({
    orgIdentifier: props.orgIdentifier,
    reportingOrgIdentifier: props.reportingOrgIdentifier,
    version: props.version,
    options: props.options
  }),
  handleSubmit: (values: any, {setSubmitting, setReset}: any) => {
    this.props.handleSubmit(values, setSubmitting, setReset)
  },
  displayName: "MyForm"
})

interface SelectProps {
  label: string,
  name: string,
  options: object,
  value: any,
  onChange: (name: string, value: any) => void,
  onBlur: (name: string, value: boolean) => void
}

class MySelect extends React.Component<SelectProps> {

  constructor (props: SelectProps) {
   super(props)
  }

  handleChange = (value: any) => {
    // this is going to call setFieldValue and manually update values.topcis
    this.props.onChange(this.props.name, value)
  };

  handleBlur = () => {
    // this is going to call setFieldTouched and manually update touched.topcis
    this.props.onBlur(this.props.name, true);
  };

  render() {

    const { fields }: any = {...this.props}
    return (
      <div>
        <label htmlFor={this.props.name}>{this.props.label}</label>
        <Select
          id={this.props.name}
          options={this.props.options}
          onChange={this.handleChange}
          onBlur={this.handleBlur}
          value={this.props.value}
          {...fields}
        />
        <ErrorMessage name={this.props.name} />
      </div>
    );
  }
}

const MyEnhancedForm = orgForm(MyForm)

export class OrgReport extends React.Component<OrgReportWriterFormProps> {

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

  handleSubmit = (values: OrgFormProps, setSubmitting: Function, reset: Function) => {
    const submitting = !this.state.toggleSubmitting
    this.setState({txKey: '', txSummary: '', toggleSubmitting: submitting, submitFunc: setSubmitting, resetFunc: reset})
    setSubmitting(submitting)
    const orgReportDetails: OrgReportProps = {
      orgIdentifier: values.orgIdentifier,
      reportingOrgIdentifier: values.reportingOrgIdentifier,
      version: values.version
    }
    console.log('Submit Org report ', orgReportDetails)
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
        <MyEnhancedForm
          orgIdentifier=''
          reportingOrgIdentifier=''
          version='2.03'
          options={ fields }
          handleSubmit={this.handleSubmit.bind(this)}/>
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
)(OrgReport))) */
