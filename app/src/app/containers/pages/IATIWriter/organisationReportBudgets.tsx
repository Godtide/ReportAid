import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { Formik, Form, Field, FormikProps, ErrorMessage} from 'formik'
import * as Yup from 'yup'

import { ApplicationState } from '../../../store'
import { ActionProps, TxData } from '../../../store/types'
import { IATIOrgReportBudgetProps, OrgReportBudgetProps} from '../../../store/IATI/types'

import { getOrgReports } from '../../../store/IATI/IATIReader/organisationReports/actions'
import { OrgReportData } from '../../../store/IATI/IATIReader/organisationReports/types'

import { setOrganisationReportBudget } from '../../../store/IATI/IATIWriter/organisationReportBudgets/actions'

import { LinearProgress } from '@material-ui/core'
import Button from '@material-ui/core/Button'
import { Select, TextField } from "material-ui-formik-components"
//import { Date } from 'formik-material-ui'

import { OrganisationReportBudget, Transaction } from '../../../utils/strings'
import { Helpers } from '../../../utils/config'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

const reportSchema = Yup.object().shape({
  reportRef: Yup
    .string()
    .required('Required'),
  budgetLine: Yup
    .string()
    .required('Required'),
  value: Yup
    .number()
    .required('Required'),
  status: Yup
    .string()
    .required('Required'),
  start: Yup
    .string()
    .required('Required'),
  end: Yup
    .string()
    .required('Required'),
})

interface BudgetProps {
  tx: TxData,
  orgReports: OrgReportData
}

export interface OrgReportBudgetsDispatchProps {
  handleSubmit: (values: any) => void
  getOrgReports: () => void
}

type OrgReportBudgetsFormProps = WithStyles<typeof styles> & BudgetProps & OrgReportBudgetsDispatchProps

export class OrgReportBudgetsForm extends React.Component<OrgReportBudgetsFormProps> {

  state = {
    txKey: '',
    txSummary: '',
    toggleSubmitting: false,
    submitFunc: (function(submit: boolean) { return submit }),
    resetFunc: (function() { return null })
  }

  constructor (props: OrgReportBudgetsFormProps) {
   super(props)
  }

  componentDidMount() {
    this.props.getOrgReports()
  }

  componentDidUpdate(previousProps: OrgReportBudgetsFormProps) {
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
      this.setState({txKey: txKey, orgReportBudgetsFormtxSummary: txSummary, toggleSubmitting: submitting})
      this.state.submitFunc(submitting)
      this.state.resetFunc()
    }
  }

  handleSubmit = (values: OrgReportBudgetProps, setSubmitting: Function, reset: Function) => {
    //console.log('Values: ', values)
    const submitting = !this.state.toggleSubmitting
    this.setState({txKey: '', txSummary: '', toggleSubmitting: submitting, submitFunc: setSubmitting, resetFunc: reset})
    setSubmitting(submitting)
    this.props.handleSubmit(values)
  }

  render() {

    //console.log('Props orgs: ', this.props.orgs)
    let orgReportsRefs: any = []
    Object.keys(this.props.orgReports).forEach((key) => {
      orgReportsRefs.push({ value: key, label: key })
    })

    return (
      <div>
        <h2>{OrganisationReportBudget.headingOrgReportBudgetWriter}</h2>
        <div>
          <Formik
            initialValues={ {reportRef: orgReportsRefs[0].value,
                             budgetLine: "",
                             value: 0,
                             status: "",
                             start: "",
                             end: ""
                            }}
            validationSchema={reportSchema}
            onSubmit={(values: OrgReportBudgetProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<OrgReportBudgetProps>) => (
              <Form>
                <Field
                  name="reportRef"
                  label={OrganisationReportBudget.reportReference}
                  component={Select}
                  options={orgReportsRefs}
                />
                <ErrorMessage name='reportRef' />
                <br />
                <Field
                  name='budgetLine'
                  label={OrganisationReportBudget.budgetLine}
                  component={TextField}
                />
                <ErrorMessage name='budgetLine' />
                <br />
                <Field
                  name='value'
                  label={OrganisationReportBudget.value}
                  component={TextField}
                />
                <ErrorMessage name='value' />
                <Field
                  name='status'
                  label={OrganisationReportBudget.status}
                  component={TextField}
                />
                <ErrorMessage name='status' />
                <Field
                  name='start'
                  label={OrganisationReportBudget.budgetStart}
                  component={TextField}
                />
                <ErrorMessage name='start' />
                <br />
                <Field
                  name='end'
                  label={OrganisationReportBudget.budgetEnd}
                  component={TextField}
                />
                <ErrorMessage name='end' />
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

const mapStateToProps = (state: ApplicationState): BudgetProps => {
  //console.log(state.orgReader)
  return {
    tx: state.orgReportBudgetsForm.data,
    orgReports: state.orgReportsReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgReportBudgetsDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setOrganisationReportBudget(ownProps)),
    getOrgReports: () => dispatch(getOrgReports())
  }
}

export const OrganisationReportBudgets = withTheme(withStyles(styles)(connect<BudgetProps, OrgReportBudgetsDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrgReportBudgetsForm)))
