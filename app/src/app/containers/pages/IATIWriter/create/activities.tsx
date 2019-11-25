import * as React from 'react'
import { RouteComponentProps } from 'react-router-dom';

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
import { IATIActivitiesReport, IATIActivitiesData, ActivitiesProps } from '../../../../store/IATI/types'
import { FormData } from '../../../../store/helpers/forms/types'

import { setFormFunctions } from '../../../../store/helpers/forms/actions'
import { setActivities } from '../../../../store/IATI/IATIWriter/activities/activities/actions'

import { TransactionHelper } from '../../../io/transactionHelper'

import { Activities as ActivitiesStrings } from '../../../../utils/strings'
import { Helpers } from '../../../../utils/config'
import { history } from '../../../../utils/history'

const activitiesSchema = Yup.object().shape({
  activitiesRef: Yup
    .string()
    .required('Required'),
  version: Yup
    .string()
    .required('Required'),
  linkedData: Yup
    .string()
    .required('Required'),
})

interface ActivitiesDataProps {
  activitiesRef: string
  activities: IATIActivitiesReport
}

interface ActivitiesDispatchProps {
  handleSubmit: (values: any) => void
  setFormFunctions: (formProps: FormData) => void
}

type ActivitiesFormProps = ActivitiesDataProps & ActivitiesDispatchProps

class ActivitiesForm extends React.Component<ActivitiesFormProps> {

  constructor (props: ActivitiesFormProps) {
   super(props)
  }

  handleSubmit = (values: ActivitiesProps, setSubmitting: Function, reset: Function) => {
    this.props.setFormFunctions({submitFunc: setSubmitting, resetFunc: reset})
    this.props.handleSubmit(values)
  }

  getActivitiesData = (activities: IATIActivitiesReport): IATIActivitiesData => {
    let newActivities: IATIActivitiesData = {
      activitiesRef: this.props.activitiesRef,
      version: '',
      generatedTime: '',
      linkedData: ''
    }
    if ( typeof activities.data != 'undefined' ) {
      newActivities.version = activities.data[0].version
      newActivities.generatedTime = activities.data[0].generatedTime
      newActivities.linkedData = activities.data[0].linkedData
    }

    return newActivities
  }

  render() {

    const activities: IATIActivitiesData = this.getActivitiesData(this.props.activities)
    //console.log('Activities: ', activities)

    return (
      <div>
        <h2>{ActivitiesStrings.headingActivitiesWriter}</h2>
        <div>
          <Formik
            initialValues={ {activitiesRef: activities.activitiesRef,
                             version: activities.version,
                             linkedData: activities.linkedData} }
            enableReinitialize={true}
            validationSchema={activitiesSchema}
            onSubmit={(values: ActivitiesProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<ActivitiesProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <Field
                    readOnly
                    name='activitiesRef'
                    value={this.props.activitiesRef}
                    label={ActivitiesStrings.activitiesReference}
                    component={TextField}
                  />
                  <Field
                    name="version"
                    label={ActivitiesStrings.version}
                    component={Select}
                    options={Helpers.reportVersions}
                  />
                  <ErrorMessage name='version' />
                  <Field
                    name='linkedData'
                    label={ActivitiesStrings.linkedData}
                    component={TextField}
                  />
                  <ErrorMessage name='linkedData' />
                  <br />
                  {formProps.isSubmitting && <LinearProgress />}
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

const mapStateToProps = (state: ApplicationState): ActivitiesDataProps => {
  //console.log(state.orgReader)
  return {
    activitiesRef: state.keys.data.activities,
    activities: state.report.data as IATIActivitiesReport
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): ActivitiesDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setActivities(ownProps)),
    setFormFunctions: (formProps: FormData) => dispatch(setFormFunctions(formProps))
  }
}

export const Activities = connect<ActivitiesDataProps, ActivitiesDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(ActivitiesForm)
