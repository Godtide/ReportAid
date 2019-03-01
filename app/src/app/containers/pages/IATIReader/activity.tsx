import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'
import Markdown from 'react-markdown'

import { Formik, Form, Field, FormikProps, ErrorMessage} from 'formik'
import * as Yup from 'yup'
import { LinearProgress } from '@material-ui/core'
import Button from '@material-ui/core/Button'
//import { Date } from 'formik-material-ui'
import FormControl from '@material-ui/core/FormControl'

import { ActivitiesPicker } from '../../../components/io/activitiesPicker'
import { FormData } from '../../../store/helpers/forms/types'

import { initialise } from '../../../store/IATI/IATIReader/actions'
import { getDictEntries } from '../../../components/io/dict'
import { setFormFunctions } from '../../../store/helpers/forms/actions'
import { getActivity } from '../../../store/IATI/IATIReader/activities/activity/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { IATIActivityReport } from '../../../store/IATI/IATIReader/activities/types'
import { ActivityReportProps } from '../../../store/IATI/IATIReader/activities/types'

import { Activity as ActivityStrings } from '../../../utils/strings'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

const reportSchema = Yup.object().shape({
  activitiesRef: Yup
    .string()
    .required('Required')
})

interface ActivityProps {
  submittingFunc: Function,
  resettingFunc: Function
  activitiesRef: string
  activity: IATIActivityReport
}

interface ActivityDispatchProps {
  handleSubmit: (values: any) => void
  initialise: () => void
  setFormFunctions: (formProps: FormData) => void
}

type ActivityReaderProps =  WithStyles<typeof styles> & ActivityProps & ActivityDispatchProps

class ActivityReader extends React.Component<ActivityReaderProps> {

  constructor (props: ActivityReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.initialise()
  }

  componentDidUpdate(previousProps: ActivityReaderProps) {
    if(previousProps.activity != this.props.activity) {
      this.props.submittingFunc(false)
      this.props.resettingFunc()
    }
  }

  handleSubmit = (values: ActivityReportProps, setSubmitting: Function, reset: Function) => {
    this.props.setFormFunctions({submitFunc: setSubmitting, resetFunc: reset})
    this.props.initialise()
    this.props.handleSubmit(values)
  }

  render() {

    const xs = getDictEntries(this.props.activity)

    return (
      <div>
        <h2>{ActivityStrings.headingActivityReader}</h2>
        <div>
          <Formik
            initialValues={ {activitiesRef: "",
                             isReport: true
                            }}
            validationSchema={reportSchema}
            onSubmit={(values: ActivityReportProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<ActivityReportProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <ActivitiesPicker
                    setValue={formProps.setFieldValue}
                    name='activitiesRef'
                    label={ActivityStrings.activitiesReference}
                  />
                  <ErrorMessage name='activitiesRef' />
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
        <hr />
        <h3>{ActivityStrings.activityDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

/*
<p>
  <b>{ActivityStrings.numActivities}</b>: {numActivities}
</p>
*/

const mapStateToProps = (state: ApplicationState): ActivityProps => {
  //console.log(state.orgReader)
  return {
    submittingFunc: state.forms.data.submitFunc,
    resettingFunc: state.forms.data.resetFunc,
    activitiesRef: state.keys.data.activities,
    activity: state.report.data as IATIActivityReport
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): ActivityDispatchProps => {
  return {
    handleSubmit: (ownProps: ActivityReportProps) => dispatch(getActivity(ownProps)),
    initialise: () => dispatch(initialise()),
    setFormFunctions: (formProps: FormData) => dispatch(setFormFunctions(formProps))
  }
}

export const Activity = withTheme(withStyles(styles)(connect<ActivityProps, ActivityDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(ActivityReader)))
