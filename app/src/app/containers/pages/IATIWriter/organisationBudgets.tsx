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

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { OrganisationBudgetProps } from '../../../store/IATI/types'

import { setOrganisationBudget } from '../../../store/IATI/IATIWriter/organisations/organisationBudgets/actions'

import { FormikDatePicker } from '../../../components/io/datePicker'
import { OrganisationsPicker } from '../../../components/io/organisationsPicker'
import { OrganisationPicker } from '../../../components/io/organisationPicker'
import { OrgPicker } from '../../../components/io/orgPicker'
import { TransactionHelper } from '../../io/transactionHelper'

import { OrganisationBudget } from '../../../utils/strings'
import { Helpers } from '../../../utils/config'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

/*const StyledSelect = withStyles({
  root: {
    background: 'linear-gradient(45deg, #FE6B8B 30%, #FF8E53 90%)',
    borderRadius: 3,
    border: 0,
    color: 'white',
    height: 48,
    padding: '0 30px',
    boxShadow: '0 3px 5px 2px rgba(255, 105, 135, .3)',
    width: '10%'
  }
})(Select);*/

const reportSchema = Yup.object().shape({
  organisations: Yup
    .object()
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

interface OrganisationBudgetsDispatchProps {
  handleSubmit: (values: any) => void
}

type OrganisationBudgetsFormProps = WithStyles<typeof styles> & OrganisationBudgetsDispatchProps

export class OrganisationBudgetsForm extends React.Component<OrganisationBudgetsFormProps> {

  state = {
    submitFunc: (function(submit: boolean) { return submit }),
    resetFunc: (function() { return null })
  }

  constructor (props: OrganisationBudgetsFormProps) {
   super(props)
  }

  handleSubmit = (values: OrganisationBudgetProps, setSubmitting: Function, reset: Function) => {
    this.setState({submitFunc: setSubmitting, resetFunc: reset})
    this.props.handleSubmit(values)
  }

  render() {

    return (
      <div>
        <h2>{OrganisationBudget.headingOrganisationBudgetWriter}</h2>
        <div>
          <Formik
            initialValues={ {organisations: {},
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
            onSubmit={(values: OrganisationBudgetProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<OrganisationBudgetProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <OrganisationsPicker name='organisations' label={OrganisationBudget.organisationsReference} />
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

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrganisationBudgetsDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setOrganisationBudget(ownProps))
  }
}

export const OrganisationBudgets = withTheme(withStyles(styles)(connect<void, OrganisationBudgetsDispatchProps, {}, ApplicationState>(
  null,
  mapDispatchToProps
)(OrganisationBudgetsForm)))
