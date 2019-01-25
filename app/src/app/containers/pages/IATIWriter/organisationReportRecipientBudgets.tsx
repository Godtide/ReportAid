import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { Formik, Form, Field, FormikProps, ErrorMessage} from 'formik'
import * as Yup from 'yup'

import { ApplicationState } from '../../../store'
import { ActionProps, TxData } from '../../../store/types'
import { IATIOrgReportRecipientBudgetProps, OrgReportRecipientBudgetProps} from '../../../store/IATI/types'

import { setRecipientBudget } from '../../../store/IATI/IATIWriter/organisationReportRecipientBudgets/actions'

import { LinearProgress } from '@material-ui/core'
import Button from '@material-ui/core/Button'
import FormControl from '@material-ui/core/FormControl'
import { Select, TextField } from "material-ui-formik-components"

import { FormikDatePicker } from '../../../components/io/datePicker'
import { OrganisationPicker } from '../../../components/io/orgPicker'
import { FormikStatusPicker } from '../../../components/io/statusPicker'
import { OrgReportPicker } from '../../../components/io/reportPicker'

import { OrganisationReportRecipientBudget, Transaction } from '../../../utils/strings'
import { Helpers } from '../../../utils/config'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

const reportSchema = Yup.object().shape({
  reportRef: Yup
    .string()
    .required('Required'),
  orgRef: Yup
    .string()
    .required('Required'),
  budgetLine: Yup
    .string()
    .required('Required'),
  value: Yup
    .number()
    .required('Required'),
  status: Yup
    .number()
    .required('Required'),
  startDay: Yup
    .number()
    .required('Required'),
  startMonth: Yup
    .number()
    .required('Required'),
  startYear: Yup
    .number()
    .required('Required'),
  endDay: Yup
    .number()
    .required('Required'),
  endMonth: Yup
    .number()
    .required('Required'),
  endYear: Yup
    .number()
    .required('Required'),
})

const StartDatePickerProps = {
  day: {
    name: 'startDay',
    label: OrganisationReportRecipientBudget.budgetStartDay
  },
  month: {
    name: 'startMonth',
    label: OrganisationReportRecipientBudget.budgetStartMonth
  },
  year: {
    name: 'startYear',
    label: OrganisationReportRecipientBudget.budgetStartYear
  }
}

const EndDatePickerProps = {
  day: {
    name: 'endDay',
    label: OrganisationReportRecipientBudget.budgetEndDay
  },
  month: {
    name: 'endMonth',
    label: OrganisationReportRecipientBudget.budgetEndMonth
  },
  year: {
    name: 'endYear',
    label: OrganisationReportRecipientBudget.budgetEndYear
  }
}

interface RecipientBudgetProps {
  tx: TxData
}

export interface OrgReportRecipientBudgetsDispatchProps {
  handleSubmit: (values: any) => void
}

type OrgReportRecipientBudgetsFormProps = WithStyles<typeof styles> & RecipientBudgetProps & OrgReportRecipientBudgetsDispatchProps

export class OrgReportRecipientBudgetsForm extends React.Component<OrgReportRecipientBudgetsFormProps> {

  state = {
    txKey: '',
    txSummary: '',
    toggleSubmitting: false,
    submitFunc: (function(submit: boolean) { return submit }),
    resetFunc: (function() { return null })
  }

  constructor (props: OrgReportRecipientBudgetsFormProps) {
   super(props)
  }

  componentDidUpdate(previousProps: OrgReportRecipientBudgetsFormProps) {
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

  handleSubmit = (values: OrgReportRecipientBudgetProps, setSubmitting: Function, reset: Function) => {
    //console.log('Values: ', values)
    const submitting = !this.state.toggleSubmitting
    this.setState({txKey: '', txSummary: '', toggleSubmitting: submitting, submitFunc: setSubmitting, resetFunc: reset})
    setSubmitting(submitting)
    this.props.handleSubmit(values)
  }

  render() {

    return (
      <div>
        <h2>{OrganisationReportRecipientBudget.headingOrgReportRecipientBudgetWriter}</h2>
        <div>
          <Formik
            initialValues={ {reportRef: "",
                             orgRef: "",
                             budgetLine: "",
                             value: 0,
                             status: 1,
                             startDay: 1,
                             startMonth: 1,
                             startYear: 1990,
                             endDay: 1,
                             endMonth: 1,
                             endYear: 1990,
                            }}
            validationSchema={reportSchema}
            onSubmit={(values: OrgReportRecipientBudgetProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<OrgReportRecipientBudgetProps>) => (
              <Form>
                <FormControl fullWidth={false}>
                  <OrgReportPicker name='reportRef' label={OrganisationReportRecipientBudget.reportReference} />
                  <OrganisationPicker name='orgRef' label={OrganisationReportRecipientBudget.orgReference} />
                  <Field
                    name='budgetLine'
                    label={OrganisationReportRecipientBudget.budgetLine}
                    component={TextField}
                  />
                  <ErrorMessage name='budgetLine' />
                  <Field
                    name='value'
                    label={OrganisationReportRecipientBudget.value}
                    component={TextField}
                  />
                  <ErrorMessage name='value' />
                  <FormikStatusPicker name='status' label={OrganisationReportRecipientBudget.status} />
                  <FormikDatePicker dates={StartDatePickerProps} />
                  <FormikDatePicker dates={EndDatePickerProps} />
                  <br />
                  {formProps.isSubmitting && <LinearProgress />}
                  <br />
                  <Button type='submit' variant="raised" color="primary" disabled={formProps.isSubmitting}>
                    Submit
                  </Button>
                </FormControl>
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

const mapStateToProps = (state: ApplicationState): RecipientBudgetProps => {
  //console.log(state.orgReader)
  return {
    tx: state.orgReportRecipientBudgetsForm.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgReportRecipientBudgetsDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setRecipientBudget(ownProps))
  }
}

export const OrganisationReportRecipientBudgets = withTheme(withStyles(styles)(connect<RecipientBudgetProps, OrgReportRecipientBudgetsDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrgReportRecipientBudgetsForm)))
