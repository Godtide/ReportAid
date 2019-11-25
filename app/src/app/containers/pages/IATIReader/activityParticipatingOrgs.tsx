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

import { getDictEntries } from '../../../components/io/dict'
import { ActivitiesPicker } from '../../../components/io/activitiesPicker'
import { ActivityPicker } from '../../../components/io/activityPicker'
import { FormData } from '../../../store/helpers/forms/types'

import { initialise } from '../../../store/IATI/IATIReader/actions'
import { setFormFunctions } from '../../../store/helpers/forms/actions'
import { getActivityParticipatingOrgs } from '../../../store/IATI/IATIReader/activities/activityParticipatingOrgs/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { IATIActivityParticipatingOrgReport, ActivitiesReportProps } from '../../../store/IATI/types'

import { ActivityParticipatingOrg as ActivityParticipatingOrgStrings } from '../../../utils/strings'

const reportSchema = Yup.object().shape({
  activitiesRef: Yup
    .string()
    .required('Required'),
  activityRef: Yup
    .string()
    .required('Required')
})

interface ActivityParticipatingOrgProps {
  submittingFunc: Function,
  resettingFunc: Function
  activitiesRef: string,
  activityRef: string,
  participatingOrgs: IATIActivityParticipatingOrgReport
}

interface ActivityParticipatingOrgDispatchProps {
  handleSubmit: (values: any) => void
  initialise: () => void
  setFormFunctions: (formProps: FormData) => void
}

type ActivityParticipatingOrgReaderProps = ActivityParticipatingOrgProps & ActivityParticipatingOrgDispatchProps

class ParticipatingOrgs extends React.Component<ActivityParticipatingOrgReaderProps> {

  constructor (props: ActivityParticipatingOrgReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.initialise()
  }

  componentDidUpdate(previousProps: ActivityParticipatingOrgReaderProps) {
    if(previousProps.participatingOrgs != this.props.participatingOrgs) {
    this.props.submittingFunc(false)
    this.props.resettingFunc()
    }
  }

  handleSubmit = (values: ActivitiesReportProps, setSubmitting: Function, reset: Function) => {
    this.props.setFormFunctions({submitFunc: setSubmitting, resetFunc: reset})
    this.props.initialise()
    this.props.handleSubmit(values)
  }

  render() {

    const xs = getDictEntries(this.props.participatingOrgs)

    return (
      <div>
        <h2>{ActivityParticipatingOrgStrings.headingActivityParticipatingOrgReader}</h2>
        <div>
          <Formik
            initialValues={ {activitiesRef: "",
                             activityRef: ""
                            }}
            validationSchema={reportSchema}
            onSubmit={(values: ActivitiesReportProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<ActivitiesReportProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <ActivitiesPicker
                    setValue={formProps.setFieldValue}
                    name='activitiesRef'
                    label={ActivityParticipatingOrgStrings.activitiesReference}
                  />
                  <ErrorMessage name='activitiesRef' />
                  <ActivityPicker
                    setValue={formProps.setFieldValue}
                    name='activityRef'
                    label={ActivityParticipatingOrgStrings.activityReference}
                  />
                  <ErrorMessage name='activityRef' />
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
        <hr />
        <h3>{ActivityParticipatingOrgStrings.participatingOrgDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): ActivityParticipatingOrgProps => {
  //console.log(state.orgReader)
  return {
    submittingFunc: state.forms.data.submitFunc,
    resettingFunc: state.forms.data.resetFunc,
    activitiesRef: state.keys.data.activities,
    activityRef: state.keys.data.activity,
    participatingOrgs: state.report.data as IATIActivityParticipatingOrgReport
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): ActivityParticipatingOrgDispatchProps => {
  return {
    handleSubmit: (values: ActivitiesReportProps) => dispatch(getActivityParticipatingOrgs(values)),
    initialise: () => dispatch(initialise()),
    setFormFunctions: (formProps: FormData) => dispatch(setFormFunctions(formProps))
  }
}

export const ActivityParticipatingOrg = connect<ActivityParticipatingOrgProps, ActivityParticipatingOrgDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(ParticipatingOrgs)
