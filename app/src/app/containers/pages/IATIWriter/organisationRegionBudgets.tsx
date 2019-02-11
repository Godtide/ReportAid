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
import { IATIOrgRegionBudgetProps, OrgRegionBudgetProps, Props } from '../../../store/IATI/types'

import { setRegionBudget } from '../../../store/IATI/IATIWriter/organisations/organisationRegionBudgets/actions'

import { FormikDatePicker } from '../../../components/io/datePicker'
import { OrgPicker } from '../../../components/io/reportPicker'
import { TransactionHelper } from '../../io/transactionHelper'

import { OrganisationRegionBudget } from '../../../utils/strings'
import { Helpers } from '../../../utils/config'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

const reportSchema = Yup.object().shape({
  organisations: Yup
    .object()
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

interface OrgRegionBudgetsDispatchProps {
  handleSubmit: (values: any) => void
}

type OrgRegionBudgetsFormProps = WithStyles<typeof styles> & OrgRegionBudgetsDispatchProps

export class OrgRegionBudgetsForm extends React.Component<OrgRegionBudgetsFormProps> {

  state = {
    submitFunc: (function(submit: boolean) { return submit }),
    resetFunc: (function() { return null })
  }

  constructor (props: OrgRegionBudgetsFormProps) {
   super(props)
  }

  handleSubmit = (values: OrgRegionBudgetProps, setSubmitting: Function, reset: Function) => {
    this.setState({submitFunc: setSubmitting, resetFunc: reset})
    this.props.handleSubmit(values)
  }

  render() {

    return (
      <div>
        <h2>{OrganisationRegionBudget.headingOrgRegionBudgetWriter}</h2>
        <div>
          <Formik
            initialValues={ {organisations: {} as Props,
                             regionRef: 1,
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
            onSubmit={(values: OrgRegionBudgetProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<OrgRegionBudgetProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <OrganisationsPicker name='organisations' label={OrganisationRegionBudget.organisationsReference} />
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

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgRegionBudgetsDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setRegionBudget(ownProps))
  }
}

export const OrganisationRegionBudgets = withTheme(withStyles(styles)(connect<void, OrgRegionBudgetsDispatchProps, {}, ApplicationState>(
  null,
  mapDispatchToProps
)(OrgRegionBudgetsForm)))
