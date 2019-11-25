import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { Formik, Form, Field, FormikProps, ErrorMessage} from 'formik'
import * as Yup from 'yup'
import { LinearProgress } from '@material-ui/core'
import Button from '@material-ui/core/Button'
//import { Date } from 'formik-material-ui'
import FormControl from '@material-ui/core/FormControl'
import { Select, TextField } from "material-ui-formik-components"

import { ApplicationState } from '../../../../store'
import { ActionProps } from '../../../../store/types'
import { OrganisationBudgetProps } from '../../../../store/IATI/types'
import { FormData } from '../../../../store/helpers/forms/types'

import { setFormFunctions } from '../../../../store/helpers/forms/actions'
import { setOrganisationBudget } from '../../../../store/IATI/IATIWriter/organisations/organisationBudgets/actions'

import { FormikDatePicker } from '../../../../components/io/datePicker'
import { OrganisationsPicker } from '../../../../components/io/organisationsPicker'
import { OrganisationPicker } from '../../../../components/io/organisationPicker'
import { OrgPicker } from '../../../../components/io/orgPicker'
import { TransactionHelper } from '../../../io/transactionHelper'

import { OrganisationBudget } from '../../../../utils/strings'
import { Helpers } from '../../../../utils/config'

const reportSchema = Yup.object().shape({
  organisationsRef: Yup
    .string()
    .required('Required'),
  organisationRef: Yup
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
    label: OrganisationBudget.budgetStartDay
  },
  month: {
    name: 'startMonth',
    label: OrganisationBudget.budgetStartMonth
  },
  year: {
    name: 'startYear',
    label: OrganisationBudget.budgetStartYear
  }
}

const EndDatePickerProps = {
  day: {
    name: 'endDay',
    label: OrganisationBudget.budgetEndDay
  },
  month: {
    name: 'endMonth',
    label: OrganisationBudget.budgetEndMonth
  },
  year: {
    name: 'endYear',
    label: OrganisationBudget.budgetEndYear
  }
}

interface OrganisationBudgetsKeyProps {
  organisationsRef: string
  organisationRef: string
}

interface OrganisationBudgetsDispatchProps {
  handleSubmit: (values: any) => void
  setFormFunctions: (formProps: FormData) => void
}

type OrganisationBudgetsFormProps = OrganisationBudgetsKeyProps & OrganisationBudgetsDispatchProps

export class OrganisationBudgetsForm extends React.Component<OrganisationBudgetsFormProps> {

  constructor (props: OrganisationBudgetsFormProps) {
   super(props)
  }

  handleSubmit = (values: OrganisationBudgetProps, setSubmitting: Function, reset: Function) => {
    this.props.setFormFunctions({submitFunc: setSubmitting, resetFunc: reset})
    this.props.handleSubmit(values)
    //this.props.initialise()
  }

  render() {

    return (
      <div>
        <h2>{OrganisationBudget.headingOrganisationBudgetWriter}</h2>
        <div>
          <Formik
            initialValues={ {organisationsRef: "",
                             organisationRef: "",
                             budgetRef: "",
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
            onSubmit={(values: OrganisationBudgetProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<OrganisationBudgetProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <OrganisationsPicker
                    setValue={formProps.setFieldValue}
                    name='organisationsRef'
                    label={OrganisationBudget.organisationsReference}
                  />
                  <ErrorMessage name='organisationsRef' />
                  <OrganisationPicker
                    setValue={formProps.setFieldValue}
                    name='organisationRef'
                    label={OrganisationBudget.organisationReference}
                  />
                  <ErrorMessage name='organisationRef' />
                  <Field
                    name='budgetLine'
                    label={OrganisationBudget.budgetLine}
                    component={TextField}
                  />
                  <ErrorMessage name='budgetLine' />
                  <Field
                    name='value'
                    label={OrganisationBudget.value}
                    component={TextField}
                  />
                  <ErrorMessage name='value' />
                  <Field
                    name="status"
                    label={OrganisationBudget.status}
                    component={Select}
                    options={Helpers.financeStatus}
                  />
                  <ErrorMessage name='status' />
                  <FormikDatePicker dates={StartDatePickerProps} />
                  <FormikDatePicker dates={EndDatePickerProps} />
                  <br />
                  {formProps.isSubmitting && formProps.isValidating && <LinearProgress />}
                  <br />
                  <Button type='submit' variant="contained" color="primary" disabled={formProps.isSubmitting}>
                    Submit
                  </Button>
                </FormControl>
              </Form>
            )}
          />
        </div>
        <TransactionHelper/>
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrganisationBudgetsKeyProps => {
  //console.log(state.orgReader)
  return {
    organisationsRef: state.keys.data.organisations,
    organisationRef: state.keys.data.organisation
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrganisationBudgetsDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setOrganisationBudget(ownProps)),
    setFormFunctions: (formProps: FormData) => dispatch(setFormFunctions(formProps))
  }
}

export const OrganisationBudgets = connect<OrganisationBudgetsKeyProps, OrganisationBudgetsDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrganisationBudgetsForm)
