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
import { IATIActivityRelatedActivityReport, ActivityRelatedActivityProps } from '../../../../store/IATI/types'
import { FormData } from '../../../../store/helpers/forms/types'

import { setFormFunctions } from '../../../../store/helpers/forms/actions'
import { setRelatedActivity } from '../../../../store/IATI/IATIWriter/activities/activityRelatedActivities/actions'

import { ActivitiesPicker } from '../../../../components/io/activitiesPicker'
import { ActivityPicker } from '../../../../components/io/activityPicker'
import { TransactionHelper } from '../../../io/transactionHelper'

import { ActivityRelatedActivity as ActivityRelatedActivityStrings } from '../../../../utils/strings'
import { Helpers } from '../../../../utils/config'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../../styles/theme'

const activityRelatedActivitySchema = Yup.object().shape({
  activitiesRef: Yup
    .string()
    .required('Required'),
  activityRef: Yup
    .string()
    .required('Required'),
  activityRelatedActivityRef: Yup
    .string()
    .required('Required'),
  relationType: Yup
    .number()
    .required('Required'),
  relatedActivityRef: Yup
    .string()
    .required('Required')
})



interface ActivityRelatedActivityKeyProps {
  activitiesRef: string
  activityRef: string
  activityRelatedActivityRef: string
  relatedActivities: IATIActivityRelatedActivityReport
}

export interface ActivityRelatedActivityDispatchProps {
  handleSubmit: (values: any) => void
  setFormFunctions: (formProps: FormData) => void
}

type ActivityRelatedActivityFormProps = WithStyles<typeof styles> & ActivityRelatedActivityKeyProps & ActivityRelatedActivityDispatchProps

export class ActivityRelatedActivityForm extends React.Component<ActivityRelatedActivityFormProps> {

  constructor (props: ActivityRelatedActivityFormProps) {
   super(props)
  }

  handleSubmit = (values: ActivityRelatedActivityProps, setSubmitting: Function, reset: Function) => {
    this.props.setFormFunctions({submitFunc: setSubmitting, resetFunc: reset})
    this.props.handleSubmit(values)
  }

  getData = (relatedActivities: IATIActivityRelatedActivityReport): ActivityRelatedActivityProps => {

    let newRelatedActivity: ActivityRelatedActivityProps = {
      activitiesRef: this.props.activitiesRef,
      activityRef: this.props.activityRef,
      activityRelatedActivityRef: this.props.activityRelatedActivityRef,
      relationType: 0,
      relatedActivityRef: ""
    }
    if ( typeof relatedActivities.data != 'undefined' ) {
      newRelatedActivity.relationType = relatedActivities.data[0].relationType
      newRelatedActivity.relatedActivityRef = relatedActivities.data[0].relatedActivityRef
    }

    return newRelatedActivity
  }

  render() {

    const thisRelatedActivity: ActivityRelatedActivityProps = this.getData(this.props.relatedActivities)

    return (
      <div>
        <h2>{ActivityRelatedActivityStrings.headingActivityRelatedActivityWriter}</h2>
        <div>
          <Formik
            initialValues={ {activitiesRef: this.props.activitiesRef,
                             activityRef: this.props.activityRef,
                             activityRelatedActivityRef: this.props.activityRelatedActivityRef,
                             relationType: thisRelatedActivity.relationType,
                             relatedActivityRef: thisRelatedActivity.relatedActivityRef
                            }}
            enableReinitialize={true}
            validationSchema={activityRelatedActivitySchema}
            onSubmit={(values: ActivityRelatedActivityProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<ActivityRelatedActivityProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <Field
                    readOnly
                    name='activityRelatedActivityRef'
                    value={this.props.activityRelatedActivityRef}
                    label={ActivityRelatedActivityStrings.activityRelatedActivityRef}
                    component={TextField}
                  />
                  <ActivitiesPicker
                    setValue={formProps.setFieldValue}
                    name='activitiesRef'
                    label={ActivityRelatedActivityStrings.activitiesReference}
                  />
                  <ErrorMessage name='activitiesRef' />
                  <ActivityPicker
                    setValue={formProps.setFieldValue}
                    name='activityRef'
                    label={ActivityRelatedActivityStrings.activityReference}
                  />
                  <ErrorMessage name='activityRef' />
                  <Field
                    name="relationType"
                    label={ActivityRelatedActivityStrings.relationType}
                    component={Select}
                    options={Helpers.relationType}
                  />
                  <ErrorMessage name='relationType' />
                  <ActivityPicker
                    setValue={formProps.setFieldValue}
                    name='relatedActivityRef'
                    label={ActivityRelatedActivityStrings.relatedActivityReference}
                  />
                  <ErrorMessage name='relatedActivityRef' />
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

const mapStateToProps = (state: ApplicationState): ActivityRelatedActivityKeyProps => {
  //console.log(state.orgReader)
  return {
    activitiesRef: state.keys.data.activities,
    activityRef: state.keys.data.activity,
    activityRelatedActivityRef: state.keys.data.activityRelatedActivity,
    relatedActivities: state.report.data as IATIActivityRelatedActivityReport
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): ActivityRelatedActivityDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setRelatedActivity(ownProps)),
    setFormFunctions: (formProps: FormData) => dispatch(setFormFunctions(formProps))
  }
}

export const ActivityRelatedActivity = withTheme(withStyles(styles)(connect<ActivityRelatedActivityKeyProps, ActivityRelatedActivityDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(ActivityRelatedActivityForm)))
