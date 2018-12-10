import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { Formik, Form, Field, FormikProps, ErrorMessage} from 'formik'
import * as Yup from 'yup'

import { ApplicationState } from '../../../store'
import { ActionProps, TxData } from '../../../store/types'
import { OrganisationProps } from '../../../store/IATI/types'

import { getOrgs } from '../../../store/IATI/IATIReader/organisationReader/actions'
import { OrgData } from '../../../store/IATI/IATIReader/organisationReader/types'

import { getDictEntries } from './dict'

import { setOrganisation } from '../../../store/IATI/IATIWriter/organisationWriter/actions'

import { LinearProgress } from '@material-ui/core'
import Button from '@material-ui/core/Button'
import { TextField } from 'formik-material-ui'

import { Organisation } from '../../../utils/strings'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

interface OrgReportProps {
  tx: TxData,
  orgs: OrgData
}

export interface OrgDispatchProps {
  handleSubmit: (values: any) => void
  getOrgs: () => void
}

const organisationSchema = Yup.object().shape({
  name: Yup
    .string()
    .required('Required'),
  code: Yup
    .string()
    .required('Required'),
  identifier: Yup
    .string()
    .required('Required')
})

type OrgReportWriterFormProps = WithStyles<typeof styles> & OrgReportProps & OrgDispatchProps

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
      let txSummary = `${Organisation.transactionFail}`
      const txKeys = Object.keys(this.props.tx)
      if (txKeys.length > 0 ) {
        txKey = Object.keys(this.props.tx)[0]
        txSummary = `${Organisation.transactionSuccess}`
      }
      this.setState({txKey: txKey, txSummary: txSummary, toggleSubmitting: submitting})
      this.state.submitFunc(submitting)
      this.state.resetFunc()
    }
  }

  handleSubmit = (values: OrganisationProps, setSubmitting: Function, reset: Function) => {
    const submitting = !this.state.toggleSubmitting
    this.setState({txKey: '', txSummary: '', toggleSubmitting: submitting, submitFunc: setSubmitting, resetFunc: reset})
    setSubmitting(submitting)
    this.props.handleSubmit(values)
  }

  render() {

    const orgs = getDictEntries(this.props.orgs)
    console.log('Blah orgs ', orgs)

    return (
      <div>
        <h2>{Organisation.headingOrgWriter}</h2>
        <div>
          <Formik
            initialValues={ {name: '', code: '', identifier: ''} }
            validationSchema={organisationSchema}
            onSubmit={(values: OrganisationProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<any>) => (
              <Form>
                <Field name='name' label={Organisation.orgName} component={TextField} />
                <ErrorMessage name='name' />
                <br />
                <Field name='code' label={Organisation.code} component={TextField} />
                <ErrorMessage name='code' />
                <br />
                <Field name='identifier' label={Organisation.identifier} component={TextField} />
                <ErrorMessage name='identifier' />
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
        <h3>{Organisation.orgTXHeader}</h3>
        <p>
          <b>{Organisation.transactionSummary}</b>: {this.state.txSummary}<br />
          <b>{Organisation.transactionKey}</b>: {this.state.txKey}
        </p>
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrgReportProps => {
  //console.log(state.orgReader)
  return {
    tx: state.orgReportsForm.data,
    orgs: state.orgReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setOrganisation(ownProps)),
    getOrgs: () => dispatch(getOrgs())
  }
}

export const OrgReportWriter = withTheme(withStyles(styles)(connect<OrgReportProps, OrgDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrgReportForm)))
