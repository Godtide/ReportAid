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
import { OrganisationRecipientBudgetProps } from '../../../store/IATI/types'

import { setRecipientBudget } from '../../../store/IATI/IATIWriter/organisations/organisationRecipientBudgets/actions'

import { FormikDatePicker } from '../../../components/io/datePicker'
import { OrganisationsPicker } from '../../../components/io/organisationsPicker'
import { OrganisationPicker } from '../../../components/io/organisationPicker'
import { OrgPicker } from '../../../components/io/orgPicker'
import { TransactionHelper } from '../../io/transactionHelper'

import { OrganisationRecipientBudget } from '../../../utils/strings'
import { Helpers } from '../../../utils/config'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

const reportSchema = Yup.object().shape({
  organisationsRef: Yup
    .string()
    .required('Required'),
  organisationRef: Yup
    .string()
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

interface OrganisationRecipientBudgetsDispatchProps {
  handleSubmit: (values: any) => void
}

type OrganisationRecipientBudgetsFormProps = WithStyles<typeof styles> & OrganisationRecipientBudgetsDispatchProps

export class OrganisationRecipientBudgetsForm extends React.Component<OrganisationRecipientBudgetsFormProps> {

  state = {
    organisationsRef: "",
    submitFunc: (function(submit: boolean) { return submit }),
    resetFunc: (function() { return null })
  }

  constructor (props: OrganisationRecipientBudgetsFormProps) {
   super(props)
  }

  handleSubmit = (values: OrganisationRecipientBudgetProps, setSubmitting: Function, reset: Function) => {
    this.setState({submitFunc: setSubmitting, resetFunc: reset})
    this.props.handleSubmit(values)
  }

  handleOrganisationsChange = (value: string) => {
    this.setState({organisationsRef: value})
  }

  handleOrganisationChange = (value: string) => {
    console.log(value)
  }

  render() {

    return (
      <div>
        <h2>{OrganisationRecipientBudget.headingOrganisationRecipientBudgetWriter}</h2>
        <div>
          <Formik
            initialValues={ {organisationsRef: "",
                             organisationRef: "",
                             budgetRef: "",
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
            onSubmit={(values: OrganisationRecipientBudgetProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<OrganisationRecipientBudgetProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <OrganisationsPicker
                    changeFunction={this.handleOrganisationsChange}
                    setValue={formProps.setFieldValue}
                    name='organisationsRef'
                    label={OrganisationRecipientBudget.organisationsReference}
                  />
                  <ErrorMessage name='organisationsRef' />
                  <OrganisationPicker
                    organisationsRef={this.state.organisationsRef}
                    changeFunction={this.handleOrganisationChange}
                    name='organisationRef'
                    label={OrganisationRecipientBudget.organisationReference}
                  />
                  <ErrorMessage name='organisationRef' />
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
          submitFunc={this.state.submitFunc}
          resetFunc={this.state.resetFunc}
        />
      </div>
    )
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrganisationRecipientBudgetsDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setRecipientBudget(ownProps))
  }
}

export const OrganisationRecipientBudgets = withTheme(withStyles(styles)(connect<void, OrganisationRecipientBudgetsDispatchProps, {}, ApplicationState>(
  null,
  mapDispatchToProps
)(OrganisationRecipientBudgetsForm)))
