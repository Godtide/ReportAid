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
import { getActivitySectors } from '../../../store/IATI/IATIReader/activities/activitySectors/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { IATIActivitySectorReport, ActivitiesReportProps } from '../../../store/IATI/types'

import { ActivitySectors as ActivitySectorsStrings } from '../../../utils/strings'

const reportSchema = Yup.object().shape({
  activitiesRef: Yup
    .string()
    .required('Required'),
  activityRef: Yup
    .string()
    .required('Required')
})

interface ActivitySectorsProps {
  submittingFunc: Function,
  resettingFunc: Function
  activitiesRef: string,
  activityRef: string,
  sectors: IATIActivitySectorReport
}

interface ActivitySectorsDispatchProps {
  handleSubmit: (values: any) => void
  initialise: () => void
  setFormFunctions: (formProps: FormData) => void
}

type ActivitySectorsReaderProps = ActivitySectorsProps & ActivitySectorsDispatchProps

class Sectors extends React.Component<ActivitySectorsReaderProps> {

  constructor (props: ActivitySectorsReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.initialise()
  }

  componentDidUpdate(previousProps: ActivitySectorsReaderProps) {
    if(previousProps.sectors != this.props.sectors) {
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

    const xs = getDictEntries(this.props.sectors)

    return (
      <div>
        <h2>{ActivitySectorsStrings.headingActivitySectorsReader}</h2>
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
                    label={ActivitySectorsStrings.activitiesReference}
                  />
                  <ErrorMessage name='activitiesRef' />
                  <ActivityPicker
                    setValue={formProps.setFieldValue}
                    name='activityRef'
                    label={ActivitySectorsStrings.activityReference}
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
        <h3>{ActivitySectorsStrings.sectorsDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): ActivitySectorsProps => {
  //console.log(state.orgReader)
  return {
    submittingFunc: state.forms.data.submitFunc,
    resettingFunc: state.forms.data.resetFunc,
    activitiesRef: state.keys.data.activities,
    activityRef: state.keys.data.activity,
    sectors: state.report.data as IATIActivitySectorReport
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): ActivitySectorsDispatchProps => {
  return {
    handleSubmit: (values: ActivitiesReportProps) => dispatch(getActivitySectors(values)),
    initialise: () => dispatch(initialise()),
    setFormFunctions: (formProps: FormData) => dispatch(setFormFunctions(formProps))
  }
}

export const ActivitySectors = connect<ActivitySectorsProps, ActivitySectorsDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(Sectors)
