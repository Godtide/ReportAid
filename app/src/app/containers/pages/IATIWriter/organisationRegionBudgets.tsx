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
import { OrganisationRegionBudgetProps } from '../../../store/IATI/types'

import { initialise } from '../../../store/IATI/IATIWriter/organisations/actions'
import { setRegionBudget } from '../../../store/IATI/IATIWriter/organisations/organisationRegionBudgets/actions'

import { FormikDatePicker } from '../../../components/io/datePicker'
import { OrganisationsPicker } from '../../../components/io/organisationsPicker'
import { OrganisationPicker } from '../../../components/io/organisationPicker'
import { OrgPicker } from '../../../components/io/orgPicker'
import { TransactionHelper } from '../../io/transactionHelper'

import { OrganisationRegionBudget } from '../../../utils/strings'
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
  regionRef: Yup
    .number()
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
    label: OrganisationRegionBudget.budgetStartDay
  },
  month: {
    name: 'startMonth',
    label: OrganisationRegionBudget.budgetStartMonth
  },
  year: {
    name: 'startYear',
    label: OrganisationRegionBudget.budgetStartYear
  }
}

const EndDatePickerProps = {
  day: {
    name: 'endDay',
    label: OrganisationRegionBudget.budgetEndDay
  },
  month: {
    name: 'endMonth',
    label: OrganisationRegionBudget.budgetEndMonth
  },
  year: {
    name: 'endYear',
    label: OrganisationRegionBudget.budgetEndYear
  }
}

interface OrganisationRegionBudgetsDispatchProps {
  handleSubmit: (values: any) => void
  initialise: () => void
}

type OrganisationRegionBudgetsFormProps = WithStyles<typeof styles> & OrganisationRegionBudgetsDispatchProps

export class OrganisationRegionBudgetsForm extends React.Component<OrganisationRegionBudgetsFormProps> {

  state = {
    organisationsRef: "",
    submitFunc: (function(submit: boolean) { return submit }),
    resetFunc: (function() { return null })
  }

  constructor (props: OrganisationRegionBudgetsFormProps) {
   super(props)
  }

  componentDidMount() {
    this.props.initialise()
  }

  handleSubmit = (values: OrganisationRegionBudgetProps, setSubmitting: Function, reset: Function) => {
    this.setState({submitFunc: setSubmitting, resetFunc: reset})
    this.props.initialise()
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
        <h2>{OrganisationRegionBudget.headingOrganisationRegionBudgetWriter}</h2>
        <div>
          <Formik
            initialValues={ {organisationsRef: "",
                             organisationRef: "",
                             budgetRef: "",
                             regionRef: 1,
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
            onSubmit={(values: OrganisationRegionBudgetProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<OrganisationRegionBudgetProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <OrganisationsPicker
                    changeFunction={this.handleOrganisationsChange}
                    setValue={formProps.setFieldValue}
                    name='organisationsRef'
                    label={OrganisationRegionBudget.organisationsReference}
                  />
                  <ErrorMessage name='organisationsRef' />
                  <OrganisationPicker
                    organisationsRef={this.state.organisationsRef}
                    changeFunction={this.handleOrganisationChange}
                    setValue={formProps.setFieldValue}
                    name='organisationRef'
                    label={OrganisationRegionBudget.organisationReference}
                  />
                  <ErrorMessage name='organisationRef' />
                  <Field
                    name='budgetLine'
                    label={OrganisationRegionBudget.budgetLine}
                    component={TextField}
                  />
                  <ErrorMessage name='budgetLine' />
                  <Field
                    name="regionRef"
                    label={OrganisationRegionBudget.regionReference}
                    component={Select}
                    options={Helpers.oecdDacCrs}
                  />
                  <ErrorMessage name='regionRef' />
                  <Field
                    name='value'
                    label={OrganisationRegionBudget.value}
                    component={TextField}
                  />
                  <ErrorMessage name='value' />
                  <Field
                    name="status"
                    label={OrganisationRegionBudget.status}
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

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrganisationRegionBudgetsDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setRegionBudget(ownProps)),
    initialise: () => dispatch(initialise())
  }
}

export const OrganisationRegionBudgets = withTheme(withStyles(styles)(connect<void, OrganisationRegionBudgetsDispatchProps, {}, ApplicationState>(
  null,
  mapDispatchToProps
)(OrganisationRegionBudgetsForm)))
