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
import { IATIOrgReportCountryBudgetProps, OrgReportCountryBudgetProps} from '../../../store/IATI/types'

import { setCountryBudget } from '../../../store/IATI/IATIWriter/organisationReportCountryBudgets/actions'

import { FormikDatePicker } from '../../../components/io/datePicker'
import { OrgReportPicker } from '../../../components/io/reportPicker'
import { TransactionHelper, TransactionTypes } from '../../io/transactionHelper'

import { OrganisationReportCountryBudget } from '../../../utils/strings'
import { Helpers } from '../../../utils/config'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

const reportSchema = Yup.object().shape({
  reportRef: Yup
    .string()
    .required('Required'),
  countryRef: Yup
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
    label: OrganisationReportCountryBudget.budgetStartDay
  },
  month: {
    name: 'startMonth',
    label: OrganisationReportCountryBudget.budgetStartMonth
  },
  year: {
    name: 'startYear',
    label: OrganisationReportCountryBudget.budgetStartYear
  }
}

const EndDatePickerProps = {
  day: {
    name: 'endDay',
    label: OrganisationReportCountryBudget.budgetEndDay
  },
  month: {
    name: 'endMonth',
    label: OrganisationReportCountryBudget.budgetEndMonth
  },
  year: {
    name: 'endYear',
    label: OrganisationReportCountryBudget.budgetEndYear
  }
}

interface OrgReportCountryBudgetsDispatchProps {
  handleSubmit: (values: any) => void
}

type OrgReportCountryBudgetsFormProps = WithStyles<typeof styles> & OrgReportCountryBudgetsDispatchProps

export class OrgReportCountryBudgetsForm extends React.Component<OrgReportCountryBudgetsFormProps> {

  state = {
    submitFunc: (function(submit: boolean) { return submit }),
    resetFunc: (function() { return null })
  }

  constructor (props: OrgReportCountryBudgetsFormProps) {
   super(props)
  }

  handleSubmit = (values: OrgReportCountryBudgetProps, setSubmitting: Function, reset: Function) => {
    this.setState({submitFunc: setSubmitting, resetFunc: reset})
    this.props.handleSubmit(values)
  }

  render() {

    return (
      <div>
        <h2>{OrganisationReportCountryBudget.headingOrgReportCountryBudgetWriter}</h2>
        <div>
          <Formik
            initialValues={ {reportRef: "",
                             countryRef: "",
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
            onSubmit={(values: OrgReportCountryBudgetProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<OrgReportCountryBudgetProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <OrgReportPicker name='reportRef' label={OrganisationReportCountryBudget.reportReference} />
                  <Field
                    name='budgetLine'
                    label={OrganisationReportCountryBudget.budgetLine}
                    component={TextField}
                  />
                  <ErrorMessage name='budgetLine' />
                  <Field
                    name="countryRef"
                    label={OrganisationReportCountryBudget.countryReference}
                    component={Select}
                    options={Helpers.countryCodes}
                  />
                  <ErrorMessage name='countryRef' />
                  <Field
                    name='value'
                    label={OrganisationReportCountryBudget.value}
                    component={TextField}
                  />
                  <ErrorMessage name='value' />
                  <Field
                    name="status"
                    label={OrganisationReportCountryBudget.status}
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
          type={TransactionTypes.ORGREPORTCOUNTRYBUDGET}
          submitFunc={this.state.submitFunc}
          resetFunc={this.state.resetFunc}
        />
      </div>
    )
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgReportCountryBudgetsDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setCountryBudget(ownProps))
  }
}

export const OrganisationReportCountryBudgets = withTheme(withStyles(styles)(connect<void, OrgReportCountryBudgetsDispatchProps, {}, ApplicationState>(
  null,
  mapDispatchToProps
)(OrgReportCountryBudgetsForm)))
