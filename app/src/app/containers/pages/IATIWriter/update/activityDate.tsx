import * as React from 'react'
import { history } from '../../../../utils/history'

import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { Formik, Form, Field, FormikProps, ErrorMessage } from 'formik'
import * as Yup from 'yup'
import { LinearProgress } from '@material-ui/core'
import Button from '@material-ui/core/Button'
import FormControl from '@material-ui/core/FormControl'

import { ApplicationState } from '../../../../store'
import { ActionProps } from '../../../../store/types'
import { IATIActivityReport } from '../../../../store/IATI/types'

import { ActivityDate as ActivityDateWriter } from '../create/activityDates'
import { ActivitiesPicker } from '../../../../components/io/activitiesPicker'
import { ActivityPicker } from '../../../../components/io/activityPicker'
import { ActivityDatePicker } from '../../../../components/io/activityDatePicker'

import { getActivityDateRecord } from '../../../../store/IATI/IATIReader/activities/activityDates/actions'

import { ActivityDates as ActivityDateStrings } from '../../../../utils/strings'

import { Paths as PathConfig } from '../../../../utils/config'

const activityDateSchema = Yup.object().shape({
  activitiesRef: Yup
    .string()
    .required('Required'),
  activityRef: Yup
    .string()
    .required('Required'),
  dateRef: Yup
    .string()
    .required('Required')
})

interface ActivityDateProps {
  activitiesRef: string
  activityRef: string
  dateRef: string
}

interface ActivityDateDispatchProps {
  getActivityDateRecord: (values: any) => void
}

type ActivityDateFormProps = ActivityDateProps & ActivityDateDispatchProps

export class ActivityDateForm extends React.Component<ActivityDateFormProps> {

  constructor (props: ActivityDateFormProps) {
   super(props)
  }

  handleSubmit = (values: ActivityDateProps, setSubmitting: Function, reset: Function) => {
    this.props.getActivityDateRecord({activitiesRef: values.activitiesRef, activityRef: values.activityRef, dateRef: values.dateRef})
    history.push(PathConfig.activityDatesWriter)
  }

  render() {

    return (
      <div>
        <h2>{ActivityDateStrings.headingActivityDatesUpdater}</h2>
        <div>
          <Formik
            initialValues={ {activitiesRef: "", activityRef: "", dateRef: ""} }
            enableReinitialize={true}
            validationSchema={activityDateSchema}
            onSubmit={(values: ActivityDateProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<ActivityDateProps>) => (
              <Form>
                <FormControl fullWidth={true}>
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
                  <ActivityDatePicker
                    setValue={formProps.setFieldValue}
                    name='dateRef'
                    label={ActivityDateStrings.dateRef}
                  />
                  <ErrorMessage name='dateRef' />
                  <br />
                  {formProps.isSubmitting && <LinearProgress />}
                  <br />
                  <Button type='submit' variant="contained" color="primary" disabled={formProps.isSubmitting}>
                    Submit
                  </Button>
                </FormControl>
              </Form>
            )}
          />
        </div>
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): ActivityDateProps => {
  //console.log(state.orgReader)
  return {
    activitiesRef: state.keys.data.activities,
    activityRef: state.keys.data.activity,
    dateRef: state.keys.data.activityDate
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): ActivityDateDispatchProps => {
  return {
    getActivityDateRecord: (props: any) => dispatch(getActivityDateRecord(props))
  }
}

export const ActivityDate = connect<ActivityDateProps, ActivityDateDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(ActivityDateForm)
