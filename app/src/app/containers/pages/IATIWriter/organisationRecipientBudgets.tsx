import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { Formik, Form, Field, FormikProps, ErrorMessage} from 'formik'
import * as Yup from 'yup'
import { LinearProgress } from '@material-ui/core'
import Button from '@material-ui/core/Button'
import FormControl from '@material-ui/core/FormControl'
import { Select, TextField } from "material-ui-formik-components"

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { IATIOrgRecipientBudgetProps, OrgRecipientBudgetProps, Props } from '../../../store/IATI/types'

import { setRecipientBudget } from '../../../store/IATI/IATIWriter/organisationRecipientBudgets/actions'

import { FormikDatePicker } from '../../../components/io/datePicker'
import { OrganisationPicker } from '../../../components/io/orgPicker'
import { OrgPicker } from '../../../components/io/reportPicker'
import { TransactionHelper, TransactionTypes } from '../../io/transactionHelper'

import { OrganisationRecipientBudget } from '../../../utils/strings'
import { Helpers } from '../../../utils/config'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

const reportSchema = Yup.object().shape({
  report: Yup
    .object()
    .required('Required'),
  recipientOrgRef: Yup
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
    label: OrganisationRecipientBudget.budgetStartDay
  },
  month: {
    name: 'startMonth',
    label: OrganisationRecipientBudget.budgetStartMonth
  },
  year: {
    name: 'startYear',
    label: OrganisationRecipientBudget.budgetStartYear
  }
}

const EndDatePickerProps = {
  day: {
    name: 'endDay',
    label: OrganisationRecipientBudget.budgetEndDay
  },
  month: {
    name: 'endMonth',
    label: OrganisationRecipientBudget.budgetEndMonth
  },
  year: {
    name: 'endYear',
    label: OrganisationRecipientBudget.budgetEndYear
  }
}

interface OrgRecipientBudgetsDispatchProps {
  handleSubmit: (values: any) => void
}

type OrgRecipientBudgetsFormProps = WithStyles<typeof styles> & OrgRecipientBudgetsDispatchProps

export class OrgRecipientBudgetsForm extends React.Component<OrgRecipientBudgetsFormProps> {

  state = {
    submitFunc: (function(submit: boolean) { return submit }),
    resetFunc: (function() { return null })
  }

  constructor (props: OrgRecipientBudgetsFormProps) {
   super(props)
  }

  handleSubmit = (values: OrgRecipientBudgetProps, setSubmitting: Function, reset: Function) => {
    this.setState({submitFunc: setSubmitting, resetFunc: reset})
    this.props.handleSubmit(values)
  }

  render() {

    return (
      <div>
        <h2>{OrganisationRecipientBudget.headingOrgRecipientBudgetWriter}</h2>
        <div>
          <Formik
            initialValues={ {report: {} as Props,
                             recipientOrgRef: "",
                             budgetLine: "",
                             value: 0,
                             status: 1,
                             startDay: 1,
                             startMonth: 1,
                             startYear: 2000,
                             endDay: 1,
                             endMonth: 1,
                             endYear: 2000
                            }}
            validationSchema={reportSchema}
            onSubmit={(values: OrgRecipientBudgetProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<OrgRecipientBudgetProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <OrgPicker name='report' label={OrganisationRecipientBudget.reportReference} />
                  <OrganisationPicker name='recipientOrgRef' label={OrganisationRecipientBudget.orgReference} />
                  <Field
                    name='budgetLine'
                    label={OrganisationRecipientBudget.budgetLine}
                    component={TextField}
                  />
                  <ErrorMessage name='budgetLine' />
                  <Field
                    name='value'
                    label={OrganisationRecipientBudget.value}
                    component={TextField}
                  />
                  <ErrorMessage name='value' />
                  <Field
                    name="status"
                    label={OrganisationRecipientBudget.status}
                    component={Select}
                    options={Helpers.financeStatus}
                  />
                  <ErrorMessage name='status' />
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
        <TransactionHelper
          type={TransactionTypes.ORGREPORTRECIPIENTBUDGET}
          submitFunc={this.state.submitFunc}
          resetFunc={this.state.resetFunc}
        />
      </div>
    )
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgRecipientBudgetsDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setRecipientBudget(ownProps))
  }
}

export const OrganisationRecipientBudgets = withTheme(withStyles(styles)(connect<void, OrgRecipientBudgetsDispatchProps, {}, ApplicationState>(
  null,
  mapDispatchToProps
)(OrgRecipientBudgetsForm)))
