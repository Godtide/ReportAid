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
import { IATIActivityBudgetReport, ActivityBudgetProps } from '../../../../store/IATI/types'
import { FormData } from '../../../../store/helpers/forms/types'

import { setFormFunctions } from '../../../../store/helpers/forms/actions'
import { setActivityBudget } from '../../../../store/IATI/IATIWriter/activities/activityBudgets/actions'

import { FormikDatePicker } from '../../../../components/io/datePicker'
import { ActivitiesPicker } from '../../../../components/io/activitiesPicker'
import { ActivityPicker } from '../../../../components/io/activityPicker'
import { TransactionHelper } from '../../../io/transactionHelper'

import { ActivityBudget } from '../../../../utils/strings'
import { Helpers } from '../../../../utils/config'

const reportSchema = Yup.object().shape({
  activitiesRef: Yup
    .string()
    .required('Required'),
  activityRef: Yup
    .string()
    .required('Required'),
  budgetType: Yup
    .number()
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
    label: ActivityBudget.budgetStartDay
  },
  month: {
    name: 'startMonth',
    label: ActivityBudget.budgetStartMonth
  },
  year: {
    name: 'startYear',
    label: ActivityBudget.budgetStartYear
  }
}

const EndDatePickerProps = {
  day: {
    name: 'endDay',
    label: ActivityBudget.budgetEndDay
  },
  month: {
    name: 'endMonth',
    label: ActivityBudget.budgetEndMonth
  },
  year: {
    name: 'endYear',
    label: ActivityBudget.budgetEndYear
  }
}

interface ActivityBudgetsKeyProps {
  activitiesRef: string
  activityRef: string
  budgetRef: string
  budgets: IATIActivityBudgetReport
}

interface ActivityBudgetsDispatchProps {
  handleSubmit: (values: any) => void
  setFormFunctions: (formProps: FormData) => void
}

type ActivityBudgetsFormProps = ActivityBudgetsKeyProps & ActivityBudgetsDispatchProps

export class ActivityBudgetsForm extends React.Component<ActivityBudgetsFormProps> {

  constructor (props: ActivityBudgetsFormProps) {
   super(props)
  }

  handleSubmit = (values: ActivityBudgetProps, setSubmitting: Function, reset: Function) => {
    this.props.setFormFunctions({submitFunc: setSubmitting, resetFunc: reset})
    this.props.handleSubmit(values)
    //this.props.initialise()
  }

  getData = (budgets: IATIActivityBudgetReport): ActivityBudgetProps => {

    let newBudget: ActivityBudgetProps = {
      activitiesRef: this.props.activitiesRef,
      activityRef: this.props.activityRef,
      budgetRef: this.props.budgetRef,
      budgetType: 0,
      status: 0,
      value: 0,
      startDay: 0,
      startMonth: 0,
      startYear: 0,
      endDay: 0,
      endMonth: 0,
      endYear: 0
    }
    if ( typeof budgets.data != 'undefined' ) {
      const thisStartDate = new Date(budgets.data[0].start)
      const thisEndDate = new Date(budgets.data[0].end)
      //console.log(thisDate.getDay(), thisDate.getMonth(), thisDate.getFullYear())
      newBudget.budgetType = budgets.data[0].budgetType,
      newBudget.status = budgets.data[0].status,
      newBudget.value = budgets.data[0].value,
      newBudget.startDay = thisStartDate.getDay(),
      newBudget.startMonth = thisStartDate.getMonth(),
      newBudget.startYear = thisStartDate.getFullYear(),
      newBudget.endDay = thisStartDate.getDay(),
      newBudget.endMonth = thisStartDate.getMonth(),
      newBudget.endYear = thisStartDate.getFullYear()
    }

    return newBudget
  }

  render() {

    const thisBudget: ActivityBudgetProps = this.getData(this.props.budgets)

    return (
      <div>
        <h2>{ActivityBudget.headingActivityBudgetWriter}</h2>
        <div>
          <Formik
            initialValues={ {activitiesRef: thisBudget.activitiesRef,
                             activityRef: thisBudget.activityRef,
                             budgetRef: thisBudget.budgetRef,
                             budgetType: thisBudget.budgetType,
                             value: thisBudget.value,
                             status: thisBudget.status,
                             startDay: thisBudget.startDay,
                             startMonth: thisBudget.startMonth,
                             startYear: thisBudget.startYear,
                             endDay: thisBudget.endDay,
                             endMonth: thisBudget.endMonth,
                             endYear: thisBudget.endYear
                            }}
            validationSchema={reportSchema}
            onSubmit={(values: ActivityBudgetProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<ActivityBudgetProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <ActivitiesPicker
                    setValue={formProps.setFieldValue}
                    name='activitiesRef'
                    label={ActivityBudget.activitiesReference}
                  />
                  <ErrorMessage name='activitiesRef' />
                  <ActivityPicker
                    setValue={formProps.setFieldValue}
                    name='activityRef'
                    label={ActivityBudget.activityReference}
                  />
                  <ErrorMessage name='activityRef' />
                  <Field
                    name='value'
                    label={ActivityBudget.value}
                    component={TextField}
                  />
                  <ErrorMessage name='value' />
                  <Field
                    name="budgetType"
                    label={ActivityBudget.budgetType}
                    component={Select}
                    options={Helpers.budgetType}
                  />
                  <ErrorMessage name='budgetType' />
                  <Field
                    name="status"
                    label={ActivityBudget.status}
                    component={Select}
                    options={Helpers.financeStatus}
                  />
                  <ErrorMessage name='status' />
                  <FormikDatePicker dates={StartDatePickerProps} />
                  <FormikDatePicker dates={EndDatePickerProps} />
                  <br />
                  {formProps.isSubmitting && formProps.isValidating && <LinearProgress />}
                  <br />
                  <Button type='submit' color="primary" disabled={formProps.isSubmitting}>
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

const mapStateToProps = (state: ApplicationState): ActivityBudgetsKeyProps => {
  //console.log(state.orgReader)
  return {
    activitiesRef: state.keys.data.activities,
    activityRef: state.keys.data.activity,
    budgetRef: state.keys.data.activityBudget,
    budgets: state.report.data as IATIActivityBudgetReport
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): ActivityBudgetsDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setActivityBudget(ownProps)),
    setFormFunctions: (formProps: FormData) => dispatch(setFormFunctions(formProps))
  }
}

export const ActivityBudgets = connect<ActivityBudgetsKeyProps, ActivityBudgetsDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(ActivityBudgetsForm)
