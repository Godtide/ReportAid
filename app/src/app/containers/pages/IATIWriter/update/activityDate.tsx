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

import { getActivityDateRecord } from '../../../../store/IATI/IATIReader/activities/activityDates/actions'

import { Activity as ActivityStrings } from '../../../../utils/strings'

import { Paths as PathConfig } from '../../../../utils/config'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../../styles/theme'

const activitySchema = Yup.object().shape({
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
  handleSubmit: (values: any) => void
}

type ActivityDateFormProps = WithStyles<typeof styles> & ActivityDateProps & ActivityDateDispatchProps

export class ActivityDateForm extends React.Component<ActivityDateFormProps> {

  constructor (props: ActivityDateFormProps) {
   super(props)
  }

  handleSubmit = (values: ActivityDateProps, setSubmitting: Function, reset: Function) => {
    this.props.handleSubmit({activitiesRef: values.activitiesRef, activityRef: values.activityRef, dateRef: values.dateRef})
    history.push(PathConfig.activityDatesWriter)
  }

  render() {

    return (
      <div>
        <h2>{ActivityStrings.headingActivityUpdater}</h2>
        <div>
          <Formik
            initialValues={ {activitiesRef: "", activityRef: "", dateRef: ""} }
            enableReinitialize={true}
            validationSchema={activitySchema}
            onSubmit={(values: ActivityDateProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<ActivityDateProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <ActivitiesPicker
                    setValue={formProps.setFieldValue}
                    name='activitiesRef'
                    label={ActivityStrings.activitiesReference}
                  />
                  <ErrorMessage name='activitiesRef' />
                  <ActivityPicker
                    setValue={formProps.setFieldValue}
                    name='activityRef'
                    label={ActivityStrings.activityReference}
                  />
                  <ErrorMessage name='activityReference' />
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
    handleSubmit: (props: any) => dispatch(getActivityDateRecord(props))
  }
}

export const Activity = withTheme(withStyles(styles)(connect<ActivityDateProps, ActivityDateDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(ActivityDateForm)))
