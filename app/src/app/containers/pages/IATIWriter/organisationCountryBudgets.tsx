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
import { IATIOrgCountryBudgetProps, OrgCountryBudgetProps, Props } from '../../../store/IATI/types'

import { setCountryBudget } from '../../../store/IATI/IATIWriter/organisations/organisationCountryBudgets/actions'

import { FormikDatePicker } from '../../../components/io/datePicker'
import { OrgPicker } from '../../../components/io/reportPicker'
import { TransactionHelper, TransactionTypes } from '../../io/transactionHelper'

import { OrganisationCountryBudget } from '../../../utils/strings'
import { Helpers } from '../../../utils/config'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

const reportSchema = Yup.object().shape({
  organisations: Yup
    .object()
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
    label: OrganisationCountryBudget.budgetStartDay
  },
  month: {
    name: 'startMonth',
    label: OrganisationCountryBudget.budgetStartMonth
  },
  year: {
    name: 'startYear',
    label: OrganisationCountryBudget.budgetStartYear
  }
}

const EndDatePickerProps = {
  day: {
    name: 'endDay',
    label: OrganisationCountryBudget.budgetEndDay
  },
  month: {
    name: 'endMonth',
    label: OrganisationCountryBudget.budgetEndMonth
  },
  year: {
    name: 'endYear',
    label: OrganisationCountryBudget.budgetEndYear
  }
}

interface OrgCountryBudgetsDispatchProps {
  handleSubmit: (values: any) => void
}

type OrgCountryBudgetsFormProps = WithStyles<typeof styles> & OrgCountryBudgetsDispatchProps

export class OrgCountryBudgetsForm extends React.Component<OrgCountryBudgetsFormProps> {

  state = {
    submitFunc: (function(submit: boolean) { return submit }),
    resetFunc: (function() { return null })
  }

  constructor (props: OrgCountryBudgetsFormProps) {
   super(props)
  }

  handleSubmit = (values: OrgCountryBudgetProps, setSubmitting: Function, reset: Function) => {
    this.setState({submitFunc: setSubmitting, resetFunc: reset})
    this.props.handleSubmit(values)
  }

  render() {

    return (
      <div>
        <h2>{OrganisationCountryBudget.headingOrgCountryBudgetWriter}</h2>
        <div>
          <Formik
            initialValues={ {organisations: {} as Props,
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
            onSubmit={(values: OrgCountryBudgetProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<OrgCountryBudgetProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <OrganisationsPicker name='organisations' label={OrganisationCountryBudget.organisationsReference} />
                  <Field
                    name='budgetLine'
                    label={OrganisationCountryBudget.budgetLine}
                    component={TextField}
                  />
                  <ErrorMessage name='budgetLine' />
                  <Field
                    name="countryRef"
                    label={OrganisationCountryBudget.countryReference}
                    component={Select}
                    options={Helpers.countryCodes}
                  />
                  <ErrorMessage name='countryRef' />
                  <Field
                    name='value'
                    label={OrganisationCountryBudget.value}
                    component={TextField}
                  />
                  <ErrorMessage name='value' />
                  <Field
                    name="status"
                    label={OrganisationCountryBudget.status}
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

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgCountryBudgetsDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setCountryBudget(ownProps))
  }
}

export const OrganisationCountryBudgets = withTheme(withStyles(styles)(connect<void, OrgCountryBudgetsDispatchProps, {}, ApplicationState>(
  null,
  mapDispatchToProps
)(OrgCountryBudgetsForm)))
