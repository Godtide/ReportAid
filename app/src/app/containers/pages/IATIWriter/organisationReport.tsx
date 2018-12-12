import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { Formik, Form, Field, FieldProps, FormikProps, FieldArray, ErrorMessage} from 'formik'
import * as Yup from 'yup'

import { ApplicationState } from '../../../store'
import { ActionProps, TxData } from '../../../store/types'
import { IATIOrgProps } from '../../../store/IATI/types'

import { getOrgs } from '../../../store/IATI/IATIReader/organisationReader/actions'
import { OrgData } from '../../../store/IATI/IATIReader/organisationReader/types'

import { getDictEntries } from '../../../utils/dict'

import { setOrganisation } from '../../../store/IATI/IATIWriter/organisationWriter/actions'

import { LinearProgress } from '@material-ui/core'
import Button from '@material-ui/core/Button'
import { Select } from 'formik-material-ui'

import { Organisation } from '../../../utils/strings'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

interface id {
  identifier: string
}

interface FormProps {
  orgs: Array<id>
}

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

  handleSubmit = (values: FormProps, setSubmitting: Function, reset: Function) => {
    const submitting = !this.state.toggleSubmitting
    this.setState({txKey: '', txSummary: '', toggleSubmitting: submitting, submitFunc: setSubmitting, resetFunc: reset})
    setSubmitting(submitting)
    this.props.handleSubmit(values)
  }

  render() {

    const orgs = getDictEntries(this.props.orgs) as IATIOrgProps[]
    //console.log('Blah orgs ', orgs)
    const xs = orgs.map((value: IATIOrgProps) => {
      return (
        {identifier: `${value.identifier}`}
      )
    })

    const fields = {orgs: xs}
    console.log('field ', fields)

    return (
      <div>
        <h2>{Organisation.headingOrgWriter}</h2>
        <div>
          <Formik
            initialValues={ fields }
            validationSchema={ organisationSchema }
            onSubmit={(values: FormProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={ ({ values }: FormikProps<any>) => (
              <Form>
                <FieldArray
                  name="orgs"
                  render={ () => (
                    <React.Fragment>
                      <div>
                        {values.orgs && values.orgs.length > 0 ? (
                          values.orgs.map((value: any, index: any) => (
                            <Field component="select" name={index}>
                              <option value={value.identifier}>{value.identifier}</option>
                            </Field>
                          ))
                        ): (
                          <Field component="select" name=''>
                            <option value=''>''</option>
                          </Field>
                        )}
                      </div>
                    </React.Fragment>
                  )}
                />
                <br />
                {values.isSubmitting && <LinearProgress />}
                <br />
                <Button type='submit' variant="raised" color="primary" disabled={values.isSubmitting}>
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
