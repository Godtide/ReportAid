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
import { IATIActivityAdditionalReport, IATIActivityAdditionalData, ActivityAdditionalProps } from '../../../../store/IATI/types'
import { FormData } from '../../../../store/helpers/forms/types'

import { setFormFunctions } from '../../../../store/helpers/forms/actions'
import { setActivityAdditional } from '../../../../store/IATI/IATIWriter/activities/activityAdditionals/actions'

import { ActivitiesPicker } from '../../../../components/io/activitiesPicker'
import { ActivityPicker } from '../../../../components/io/activityPicker'
import { TransactionHelper } from '../../../io/transactionHelper'

import { ActivityAdditional as ActivityAdditionalStrings } from '../../../../utils/strings'
import { Helpers } from '../../../../utils/config'

const activityAdditionalSchema = Yup.object().shape({
  activitiesRef: Yup
    .string()
    .required('Required'),
  activityRef: Yup
    .string()
    .required('Required'),
  additionalRef: Yup
    .string()
    .required('Required'),
  scope: Yup
    .number(),
  capitalSpend: Yup
    .number(),
  collaborationType: Yup
    .number(),
  defaultFlowType: Yup
    .number(),
  defaultFinanceType: Yup
    .number(),
  defaultAidType: Yup
    .string(),
  defaultTiedStatus: Yup
    .number()
})

interface ActivityAdditionalKeyProps {
  activitiesRef: string
  activityRef: string
  additionalRef: string,
  additionals: IATIActivityAdditionalReport
}

export interface ActivityAdditionalDispatchProps {
  handleSubmit: (values: any) => void
  setFormFunctions: (formProps: FormData) => void
}

type ActivityAdditionalFormProps = ActivityAdditionalKeyProps & ActivityAdditionalDispatchProps

export class ActivityAdditionalForm extends React.Component<ActivityAdditionalFormProps> {

  constructor (props: ActivityAdditionalFormProps) {
   super(props)
  }

  handleSubmit = (values: ActivityAdditionalProps, setSubmitting: Function, reset: Function) => {
    this.props.setFormFunctions({submitFunc: setSubmitting, resetFunc: reset})
    this.props.handleSubmit(values)
    //this.props.initialise()
  }

  getAdditionalData = (additional: IATIActivityAdditionalReport): IATIActivityAdditionalData => {

    let newAdditional: IATIActivityAdditionalData = {
       additionalRef: this.props.additionalRef,
       scope: 0,
       capitalSpend: 0,
       collaborationType: 0,
       defaultFlowType: 0,
       defaultFinanceType: 0,
       defaultAidType: "",
       defaultTiedStatus: 0
    }
    if ( typeof additional.data != 'undefined' ) {
      newAdditional.scope = additional.data[0].scope,
      newAdditional.capitalSpend = additional.data[0].capitalSpend,
      newAdditional.collaborationType = additional.data[0].collaborationType,
      newAdditional.defaultFlowType = additional.data[0].defaultFlowType,
      newAdditional.defaultFinanceType = additional.data[0].defaultFinanceType,
      newAdditional.defaultAidType = additional.data[0].defaultAidType,
      newAdditional.defaultTiedStatus = additional.data[0].defaultTiedStatus
    }

    return newAdditional
  }

  render() {

    const additional: IATIActivityAdditionalData = this.getAdditionalData(this.props.additionals)

    return (
      <div>
        <h2>{ActivityAdditionalStrings.headingActivityAdditionalWriter}</h2>
        <div>
          <Formik
            initialValues={ {activitiesRef: this.props.activitiesRef,
                             activityRef: this.props.activityRef,
                             additionalRef: additional.additionalRef,
                             scope: additional.scope,
                             capitalSpend: additional.capitalSpend,
                             collaborationType: additional.collaborationType,
                             defaultFlowType: additional.defaultFlowType,
                             defaultFinanceType: additional.defaultFinanceType,
                             defaultAidType: additional.defaultAidType,
                             defaultTiedStatus: additional.defaultTiedStatus
                            }}
            enableReinitialize={true}
            validationSchema={activityAdditionalSchema}
            onSubmit={(values: ActivityAdditionalProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<ActivityAdditionalProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <Field
                    readOnly
                    name='additionalRef'
                    value={this.props.additionalRef}
                    label={ActivityAdditionalStrings.additionalRef}
                    component={TextField}
                  />
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
                  <Field
                    name="defaultAidType"
                    label={ActivityAdditionalStrings.defaultAidType}
                    component={Select}
                    options={Helpers.defaultAidType}
                  />
                  <ErrorMessage name='defaultAidType' />
                  <Field
                    name="defaultFinanceType"
                    label={ActivityAdditionalStrings.defaultFinanceType}
                    component={Select}
                    options={Helpers.defaultFinanceType}
                  />
                  <ErrorMessage name='defaultFinanceType' />
                  <Field
                    name="scope"
                    label={ActivityAdditionalStrings.scope}
                    component={Select}
                    options={Helpers.scope}
                  />
                  <ErrorMessage name='scope' />
                  <Field
                    name='capitalSpend'
                    label={ActivityAdditionalStrings.capitalSpend}
                    component={TextField}
                  />
                  <ErrorMessage name='capitalSpend' />
                  <Field
                    name="collaborationType"
                    label={ActivityAdditionalStrings.collaborationType}
                    component={Select}
                    options={Helpers.collaborationType}
                  />
                  <ErrorMessage name='collaborationType' />
                  <Field
                    name="defaultFlowType"
                    label={ActivityAdditionalStrings.defaultFlowType}
                    component={Select}
                    options={Helpers.defaultFlowType}
                  />
                  <ErrorMessage name='defaultFlowType' />
                  <Field
                    name="defaultTiedStatus"
                    label={ActivityAdditionalStrings.defaultTiedStatus}
                    component={Select}
                    options={Helpers.defaultTiedStatus}
                  />
                  <ErrorMessage name='defaultTiedStatus' />
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
        <TransactionHelper/>
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): ActivityAdditionalKeyProps => {
  //console.log(state.orgReader)
  return {
    activitiesRef: state.keys.data.activities,
    activityRef: state.keys.data.activity,
    additionalRef: state.keys.data.activityAdditional,
    additionals: state.report.data as IATIActivityAdditionalReport
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): ActivityAdditionalDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setActivityAdditional(ownProps)),
    setFormFunctions: (formProps: FormData) => dispatch(setFormFunctions(formProps))
  }
}

export const ActivityAdditional = connect<ActivityAdditionalKeyProps, ActivityAdditionalDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(ActivityAdditionalForm)
