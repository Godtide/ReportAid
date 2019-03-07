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

import { Activity as ActivityWriter } from '../create/activity'
import { ActivitiesPicker } from '../../../../components/io/activitiesPicker'
import { ActivityPicker } from '../../../../components/io/activityPicker'

import { getActivityRecord } from '../../../../store/IATI/IATIReader/activities/activity/actions'

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
    .required('Required')
})

interface ActivityProps {
  activitiesRef: string
  activityRef: string
}

interface ActivityDispatchProps {
  handleSubmit: (values: any) => void
}

type ActivityFormProps = WithStyles<typeof styles> & ActivityProps & ActivityDispatchProps

export class ActivityForm extends React.Component<ActivityFormProps> {

  constructor (props: ActivityFormProps) {
   super(props)
  }

  handleSubmit = (values: ActivityProps, setSubmitting: Function, reset: Function) => {
    this.props.handleSubmit({activitiesRef: values.activitiesRef, activityRef: values.activityRef})
    history.push(PathConfig.activityWriter)
  }

  render() {

    return (
      <div>
        <h2>{ActivityStrings.headingActivityUpdater}</h2>
        <div>
          <Formik
            initialValues={ {activitiesRef: "", activityRef: ""} }
            enableReinitialize={true}
            validationSchema={activitySchema}
            onSubmit={(values: ActivityProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<ActivityProps>) => (
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

const mapStateToProps = (state: ApplicationState): ActivityProps => {
  //console.log(state.orgReader)
  return {
    activitiesRef: state.keys.data.activities,
    activityRef: state.keys.data.activity
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): ActivityDispatchProps => {
  return {
    handleSubmit: (props: any) => dispatch(getActivityRecord(props))
  }
}

export const Activity = withTheme(withStyles(styles)(connect<ActivityProps, ActivityDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(ActivityForm)))
