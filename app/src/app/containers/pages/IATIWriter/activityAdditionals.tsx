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

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { ActivityAdditionalProps } from '../../../store/IATI/types'
import { FormData } from '../../../store/helpers/forms/types'

import { initialise } from '../../../store/IATI/IATIWriter/actions'
import { setFormFunctions } from '../../../store/helpers/forms/actions'
import { newKey } from '../../../store/helpers/keys/actions'
import { setActivityAdditional } from '../../../store/IATI/IATIWriter/activities/activityAdditionals/actions'

import { ActivitiesPicker } from '../../../components/io/activitiesPicker'
import { ActivityPicker } from '../../../components/io/activityPicker'
import { OrgPicker } from '../../../components/io/orgPicker'
import { TransactionHelper } from '../../io/transactionHelper'

import { ActivityAdditional as ActivityAdditionalStrings } from '../../../utils/strings'
import { Helpers } from '../../../utils/config'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

const activityAdditionalSchema = Yup.object().shape({
  activitiesRef: Yup
    .string()
    .required('Required'),
  activityRef: Yup
    .string()
    .required('Required'),
  budgetNotProvided: Yup
    .number()
    .required('Required'),
  status: Yup
    .number()
    .required('Required'),
  scope: Yup
    .number()
    .required('Required'),
  capitalSpend: Yup
    .number()
    .required('Required'),
  collaborationType: Yup
    .number()
    .required('Required'),
  defaultFlowType: Yup
    .number()
    .required('Required'),
  defaultFinanceType: Yup
    .number()
    .required('Required'),
  defaultAidType: Yup
    .string()
    .required('Required'),
  defaultTiedStatus: Yup
    .number()
    .required('Required')
})

interface ActivityAdditionalKeyProps {
  activitiesRef: string
  activityRef: string
}

export interface ActivityAdditionalDispatchProps {
  handleSubmit: (values: any) => void
  initialise: () => void
  setFormFunctions: (formProps: FormData) => void
}

type ActivityAdditionalFormProps = WithStyles<typeof styles> & ActivityAdditionalKeyProps & ActivityAdditionalDispatchProps

export class ActivityAdditionalForm extends React.Component<ActivityAdditionalFormProps> {

  constructor (props: ActivityAdditionalFormProps) {
   super(props)
  }

  componentDidMount() {
    this.props.initialise()
  }

  handleSubmit = (values: ActivityAdditionalProps, setSubmitting: Function, reset: Function) => {
    this.props.setFormFunctions({submitFunc: setSubmitting, resetFunc: reset})
    this.props.handleSubmit(values)
    this.props.initialise()
  }

  render() {

    return (
      <div>
        <h2>{ActivityAdditionalStrings.headingActivityAdditionalWriter}</h2>
        <div>
          <Formik
            initialValues={ {activitiesRef: "",
                             activityRef: "",
                             additionalRef: "",
                             budgetNotProvided: 0,
                             status: 0,
                             scope: 0,
                             capitalSpend: 0,
                             collaborationType: 0,
                             defaultFlowType: 0,
                             defaultFinanceType: 0,
                             defaultAidType: "",
                             defaultTiedStatus: 0
                            }}
            enableReinitialize={true}
            validationSchema={activityAdditionalSchema}
            onSubmit={(values: ActivityAdditionalProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<ActivityAdditionalProps>) => (
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
                    name="budgetNotProvided"
                    label={ActivityAdditionalStrings.budgetNotProvided}
                    component={Select}
                    options={Helpers.budgetNotProvided}
                  />
                  <ErrorMessage name='hierarchy' />
                  <Field
                    name="status"
                    label={ActivityAdditionalStrings.status}
                    component={Select}
                    options={Helpers.status}
                  />
                  <ErrorMessage name='status' />
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

const mapStateToProps = (state: ApplicationState): ActivityAdditionalKeyProps => {
  //console.log(state.orgReader)
  return {
    activitiesRef: state.keys.data.activities,
    activityRef: state.keys.data.activity
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): ActivityAdditionalDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setActivityAdditional(ownProps)),
    initialise: () => dispatch(initialise()),
    setFormFunctions: (formProps: FormData) => dispatch(setFormFunctions(formProps))
  }
}

export const ActivityAdditional = withTheme(withStyles(styles)(connect<ActivityAdditionalKeyProps, ActivityAdditionalDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(ActivityAdditionalForm)))
