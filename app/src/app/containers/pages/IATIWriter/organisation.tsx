import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { Formik, Form, Field, FormikProps, ErrorMessage} from 'formik'
import * as Yup from 'yup'

import { ApplicationState } from '../../../store'
import { ActionProps, TxData } from '../../../store/types'
import { OrganisationProps } from '../../../store/IATI/types'

import { setOrganisation } from '../../../store/IATI/IATIWriter/organisation/actions'

import { LinearProgress } from '@material-ui/core'
import Button from '@material-ui/core/Button'
//import { TextField } from 'formik-material-ui'
import { TextField } from "material-ui-formik-components"

import { Organisation, Transaction } from '../../../utils/strings'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

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

interface OrgTXProps {
  tx: TxData
}

export interface OrgDispatchProps {
  handleSubmit: (values: any) => void
}

type OrgWriterFormProps = WithStyles<typeof styles> & OrgTXProps & OrgDispatchProps

export class OrgForm extends React.Component<OrgWriterFormProps> {

  state = {
    txKey: '',
    txSummary: '',
    toggleSubmitting: false,
    submitFunc: (function(submit: boolean) { return submit }),
    resetFunc: (function() { return null })
  }

  constructor (props: OrgWriterFormProps) {
   super(props)
  }

  componentDidUpdate(previousProps: OrgWriterFormProps) {
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

  handleSubmit = (values: OrganisationProps, setSubmitting: Function, reset: Function) => {
    const submitting = !this.state.toggleSubmitting
    this.setState({txKey: '', txSummary: '', toggleSubmitting: submitting, submitFunc: setSubmitting, resetFunc: reset})
    setSubmitting(submitting)
    this.props.handleSubmit(values)
  }

  render() {

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
            render={(formProps: FormikProps<OrganisationProps>) => (
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
        <h3>{Transaction.heading}</h3>
        <p>
          <b>{Transaction.summary}</b>: {this.state.txSummary}<br />
          <b>{Transaction.key}</b>: {this.state.txKey}
        </p>
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrgTXProps => {
  return {
    tx: state.orgForm.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setOrganisation(ownProps))
  }
}

export const Organisations = withTheme(withStyles(styles)(connect<OrgTXProps, OrgDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrgForm)))
