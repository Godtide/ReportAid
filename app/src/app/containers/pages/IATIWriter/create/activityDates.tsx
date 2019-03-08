import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { Formik, Form, Field, FormikProps, ErrorMessage} from 'formik'
import * as Yup from 'yup'
import { LinearProgress } from '@material-ui/core'
import Button from '@material-ui/core/Button'
import FormControl from '@material-ui/core/FormControl'
//import { Select } from 'formik-material-ui'
import { TextField, Select } from "material-ui-formik-components"

import { ApplicationState } from '../../../../store'
import { ActionProps } from '../../../../store/types'
import { ActivityDateProps } from '../../../../store/IATI/types'
import { FormData } from '../../../../store/helpers/forms/types'

import { setFormFunctions } from '../../../../store/helpers/forms/actions'
import { setActivityDate } from '../../../../store/IATI/IATIWriter/activities/activityAdditionals/actions'

import { ActivitiesPicker } from '../../../../components/io/activitiesPicker'
import { ActivityPicker } from '../../../../components/io/activityPicker'
import { OrgPicker } from '../../../../components/io/orgPicker'
import { TransactionHelper } from '../../../io/transactionHelper'

import { ActivityDate as ActivityDateStrings } from '../../../../utils/strings'
import { Helpers } from '../../../../utils/config'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../../styles/theme'

export interface ActivityDateProps {
  activitiesRef: string
  activityRef: string
  dateRef: string
  dateType: number
  startDay: number
  startMonth: number
  startYear: number
  endDay: number
  endMonth: number
  endYear: number
  narrative: string
}


const activityDateSchema = Yup.object().shape({
  activitiesRef: Yup
    .string()
    .required('Required'),
  activityRef: Yup
    .string()
    .required('Required'),
  dateRef: Yup
    .string()
    .required('Required'),
  dateType: Yup
    .number(),
  narrative: Yup
    .string()
    .required('Required')
})

const DatePickerProps = {
  day: {
    name: 'Day',
    label: OrganisationBudget.budgetStartDay
  },
  month: {
    name: 'Month',
    label: OrganisationBudget.budgetStartMonth
  },
  year: {
    name: 'Year',
    label: OrganisationBudget.budgetStartYear
  }
}

interface ActivityDateKeyProps {
  activitiesRef: string
  activityRef: string
  dateRef: string
}

export interface ActivityDateDispatchProps {
  handleSubmit: (values: any) => void
  setFormFunctions: (formProps: FormData) => void
}

type ActivityDateFormProps = WithStyles<typeof styles> & ActivityDateKeyProps & ActivityDateDispatchProps

export class ActivityDateForm extends React.Component<ActivityDateFormProps> {

  constructor (props: ActivityDateFormProps) {
   super(props)
  }

  handleSubmit = (values: ActivityDateProps, setSubmitting: Function, reset: Function) => {
    this.props.setFormFunctions({submitFunc: setSubmitting, resetFunc: reset})
    this.props.handleSubmit(values)
  }

  render() {

    return (
      <div>
        <h2>{ActivityDateStrings.headingActivityDateWriter}</h2>
        <div>
          <Formik
            initialValues={ {activitiesRef: "",
                             activityRef: "",
                             dateRef: this.props.dateRef,
                             dateType: 0,
                             day: 0,
                             month: 0,
                             year: 0,
                             narrative: ""
                            }}
            enableReinitialize={true}
            validationSchema={activityAdditionalSchema}
            onSubmit={(values: ActivityDateProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<ActivityDateProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <Field
                    readOnly
                    name='dateRef'
                    value={this.props.dateRef}
                    label={ActivityDateStrings.dateRef}
                    component={TextField}
                  />
                  <ActivitiesPicker
                    setValue={formProps.setFieldValue}
                    name='activitiesRef'
                    label={ActivityDateStrings.activitiesReference}
                  />
                  <ErrorMessage name='activitiesRef' />
                  <ActivityPicker
                    setValue={formProps.setFieldValue}
                    name='activityRef'
                    label={ActivityDateStrings.activityReference}
                  />
                  <ErrorMessage name='activityRef' />
                  <Field
                    name="dateType"
                    label={ActivityDateStrings.dateType}
                    component={Select}
                    options={Helpers.dateCodes}
                  />
                  <ErrorMessage name='dateType' />
                  <FormikDatePicker dates={DatePickerProps} />
                  <Field
                    name='narrative'
                    label={ActivityDateStrings.narrative}
                    component={TextField}
                  />
                  <ErrorMessage name='narrative' />
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
        <TransactionHelper/>
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): ActivityDateKeyProps => {
  //console.log(state.orgReader)
  return {
    activitiesRef: state.keys.data.activities,
    activityRef: state.keys.data.activity,
    dateRef: state.keys.data.activityAdditional
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): ActivityDateDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setActivityDate(ownProps)),
    setFormFunctions: (formProps: FormData) => dispatch(setFormFunctions(formProps))
  }
}

export const ActivityDate = withTheme(withStyles(styles)(connect<ActivityDateKeyProps, ActivityDateDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(ActivityDateForm)))
