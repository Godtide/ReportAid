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
import { OrganisationCountryBudgetProps } from '../../../store/IATI/types'

import { setCountryBudget } from '../../../store/IATI/IATIWriter/organisations/organisationCountryBudgets/actions'

import { FormikDatePicker } from '../../../components/io/datePicker'
import { OrganisationsPicker } from '../../../components/io/organisationsPicker'
import { OrganisationPicker } from '../../../components/io/organisationPicker'
import { OrgPicker } from '../../../components/io/orgPicker'
import { TransactionHelper } from '../../io/transactionHelper'

import { OrganisationCountryBudget } from '../../../utils/strings'
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
    .required('Required')
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

interface OrganisationCountryBudgetsDispatchProps {
  handleSubmit: (values: any) => void
}

type OrganisationCountryBudgetsFormProps = WithStyles<typeof styles> & OrganisationCountryBudgetsDispatchProps

export class OrganisationCountryBudgetsForm extends React.Component<OrganisationCountryBudgetsFormProps> {

  state = {
    organisationsRef: "",
    submitFunc: (function(submit: boolean) { return submit }),
    resetFunc: (function() { return null })
  }

  constructor (props: OrganisationCountryBudgetsFormProps) {
   super(props)
  }

  handleSubmit = (values: OrganisationCountryBudgetProps, setSubmitting: Function, reset: Function) => {
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
        <h2>{OrganisationCountryBudget.headingOrganisationCountryBudgetWriter}</h2>
        <div>
          <Formik
            initialValues={ {organisationsRef: "",
                             organisationRef: "",
                             budgetRef: "",
                             countryRef: "",
                             budgetLine: "",
                             value: 0,
                             status: 1,
                             startDay: 0,
                             startMonth: 0,
                             startYear: 0,
                             endDay: 0,
                             endMonth: 0,
                             endYear: 0
                            }}
            validationSchema={reportSchema}
            onSubmit={(values: OrganisationCountryBudgetProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<OrganisationCountryBudgetProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <OrganisationsPicker
                    changeFunction={this.handleOrganisationsChange}
                    setValue={formProps.setFieldValue}
                    name='organisationsRef'
                    label={OrganisationCountryBudget.organisationsReference}
                  />
                  <ErrorMessage name='organisationsRef' />
                  <OrganisationPicker
                    organisationsRef={this.state.organisationsRef}
                    changeFunction={this.handleOrganisationChange}
                    setValue={formProps.setFieldValue}
                    name='organisationRef'
                    label={OrganisationCountryBudget.organisationReference}
                  />
                  <ErrorMessage name='organisationRef' />
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
          submitFunc={this.state.submitFunc}
          resetFunc={this.state.resetFunc}
        />
      </div>
    )
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrganisationCountryBudgetsDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setCountryBudget(ownProps))
  }
}

export const OrganisationCountryBudgets = withTheme(withStyles(styles)(connect<void, OrganisationCountryBudgetsDispatchProps, {}, ApplicationState>(
  null,
  mapDispatchToProps
)(OrganisationCountryBudgetsForm)))
