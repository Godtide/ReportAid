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
import { getActivityAdditionals } from '../../../store/IATI/IATIReader/activities/activityAdditionals/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { IATIActivityAdditionalReport, ActivitiesReportProps } from '../../../store/IATI/types'

import { ActivityAdditional as ActivityAdditionalStrings } from '../../../utils/strings'

const reportSchema = Yup.object().shape({
  activitiesRef: Yup
    .string()
    .required('Required'),
  activityRef: Yup
    .string()
    .required('Required')
})

interface ActivityAdditionalProps {
  submittingFunc: Function,
  resettingFunc: Function
  activitiesRef: string,
  activityRef: string,
  additionals: IATIActivityAdditionalReport
}

interface ActivityAdditionalDispatchProps {
  handleSubmit: (values: any) => void
  initialise: () => void
  setFormFunctions: (formProps: FormData) => void
}

type ActivityAdditionalsReaderProps =  ActivityAdditionalProps & ActivityAdditionalDispatchProps

class Additionals extends React.Component<ActivityAdditionalsReaderProps> {

  constructor (props: ActivityAdditionalsReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.initialise()
  }

  componentDidUpdate(previousProps: ActivityAdditionalsReaderProps) {
    if(previousProps.additionals != this.props.additionals) {
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

    const xs = getDictEntries(this.props.additionals)

    return (
      <div>
        <h2>{ActivityAdditionalStrings.headingActivityAdditionalReader}</h2>
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
                    label={ActivityAdditionalStrings.activitiesReference}
                  />
                  <ErrorMessage name='activitiesRef' />
                  <ActivityPicker
                    setValue={formProps.setFieldValue}
                    name='activityRef'
                    label={ActivityAdditionalStrings.activityReference}
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
        <h3>{ActivityAdditionalStrings.activityAdditionalDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): ActivityAdditionalProps => {
  //console.log(state.orgReader)
  return {
    submittingFunc: state.forms.data.submitFunc,
    resettingFunc: state.forms.data.resetFunc,
    activitiesRef: state.keys.data.activities,
    activityRef: state.keys.data.activity,
    additionals: state.report.data as IATIActivityAdditionalReport
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): ActivityAdditionalDispatchProps => {
  return {
    handleSubmit: (values: ActivitiesReportProps) => dispatch(getActivityAdditionals(values)),
    initialise: () => dispatch(initialise()),
    setFormFunctions: (formProps: FormData) => dispatch(setFormFunctions(formProps))
  }
}

export const ActivityAdditionals = connect<ActivityAdditionalProps, ActivityAdditionalDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(Additionals)
