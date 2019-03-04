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
import { ActivityProps } from '../../../store/IATI/types'
import { FormData } from '../../../store/helpers/forms/types'

import { initialise } from '../../../store/IATI/IATIWriter/actions'
import { setFormFunctions } from '../../../store/helpers/forms/actions'
import { newKey } from '../../../store/helpers/keys/actions'
import { setActivity } from '../../../store/IATI/IATIWriter/activities/activity/actions'

import { ActivitiesPicker } from '../../../components/io/activitiesPicker'
import { ActivityPicker } from '../../../components/io/activityPicker'
import { OrgPicker } from '../../../components/io/orgPicker'
import { TransactionHelper } from '../../io/transactionHelper'

import { Activity as ActivityStrings } from '../../../utils/strings'
import { Helpers } from '../../../utils/config'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

const activitySchema = Yup.object().shape({
  activitiesRef: Yup
    .string()
    .required('Required'),
  activityRef: Yup
    .string()
    .required('Required'),
  identifier: Yup
    .string()
    .required('Required'),
  reportingOrgRef: Yup
    .string()
    .required('Required'),
  reportingOrgType: Yup
    .number()
    .required('Required'),
  reportingOrgIsSecondary: Yup
    .boolean()
    .required('Required'),
  title: Yup
    .string()
    .required('Required'),
  lang: Yup
    .string()
    .required('Required'),
  currency: Yup
    .string()
    .required('Required'),
  budgetNotProvided: Yup
    .number(),
  status: Yup
    .number()
    .required('Required'),
  humanitarian: Yup
    .boolean()
    .required('Required'),
  hierarchy: Yup
    .number()
    .required('Required'),
  linkedData: Yup
    .string(),
  description: Yup
    .string()
    .required('Required')
})

interface ActivityKeyProps {
  activitiesRef: string
  activityRef: string
}

export interface ActivityDispatchProps {
  handleSubmit: (values: any) => void
  initialise: () => void
  setFormFunctions: (formProps: FormData) => void
  newKey: () => void
}

type ActivityFormProps = WithStyles<typeof styles> & ActivityKeyProps & ActivityDispatchProps

export class ActivityForm extends React.Component<ActivityFormProps> {

  constructor (props: ActivityFormProps) {
   super(props)
  }

  componentDidMount() {
    this.props.initialise()
    this.props.newKey()
  }

  handleSubmit = (values: ActivityProps, setSubmitting: Function, reset: Function) => {
    this.props.setFormFunctions({submitFunc: setSubmitting, resetFunc: reset})
    this.props.handleSubmit(values)
    this.props.initialise()
  }

  render() {

    return (
      <div>
        <h2>{ActivityStrings.headingActivityWriter}</h2>
        <div>
          <Formik
            initialValues={ {activitiesRef: "",
                             activityRef: this.props.activityRef,
                             identifier: "",
                             reportingOrgRef: "",
                             reportingOrgType: 0,
                             reportingOrgIsSecondary: false,
                             title: "",
                             lang: "",
                             currency: "",
                             budgetNotProvided: 0,
                             status: 0,
                             humanitarian: true,
                             hierarchy: 0,
                             linkedData: "",
                             description: ""
                            }}
            enableReinitialize={true}
            validationSchema={activitySchema}
            onSubmit={(values: ActivityProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<ActivityProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <Field
                    readOnly
                    name='activityRef'
                    value={this.props.activityRef}
                    label={ActivityStrings.activityReference}
                    component={TextField}
                  />
                  <ActivitiesPicker
                    setValue={formProps.setFieldValue}
                    name='activitiesRef'
                    label={ActivityStrings.activitiesReference}
                  />
                  <ErrorMessage name='activitiesRef' />
                  <Field
                    name='title'
                    label={ActivityStrings.title}
                    component={TextField}
                  />
                  <ErrorMessage name='title' />
                  <Field
                    name="description"
                    label={ActivityStrings.description}
                    component={TextField}
                  />
                  <ErrorMessage name='description' />
                  <Field
                    name='identifier'
                    label={ActivityStrings.identifier}
                    component={TextField}
                  />
                  <ErrorMessage name='identifier' />
                  <Field
                    name='linkedData'
                    label={ActivityStrings.linkedData}
                    component={TextField}
                  />
                  <ErrorMessage name='linkedData' />
                  <Field
                    name="lang"
                    label={ActivityStrings.language}
                    component={Select}
                    options={Helpers.languageCodes}
                  />
                  <ErrorMessage name='lang' />
                  <Field
                    name="currency"
                    label={ActivityStrings.currency}
                    component={Select}
                    options={Helpers.currencyCodes}
                  />
                  <ErrorMessage name='currency' />
                  <OrgPicker
                    name='reportingOrgRef'
                    label={ActivityStrings.reportingOrgRef}
                  />
                  <ErrorMessage name='reportingOrgRef' />
                  <Field
                    name="reportingOrgType"
                    label={ActivityStrings.reportingOrgType}
                    component={Select}
                    options={Helpers.organisationCodes}
                  />
                  <ErrorMessage name='reportingOrgType' />
                  <Field
                    name="reportingOrgIsSecondary"
                    label={ActivityStrings.reportingOrgIsSecondary}
                    component={Select}
                    options={Helpers.isSecondary}
                  />
                  <ErrorMessage name='reportingOrgIsSecondary' />
                  <Field
                    name="status"
                    label={ActivityStrings.status}
                    component={Select}
                    options={Helpers.status}
                  />
                  <ErrorMessage name='status' />
                  <Field
                    name="humanitarian"
                    label={ActivityStrings.humanitarian}
                    component={Select}
                    options={Helpers.isSecondary}
                  />
                  <ErrorMessage name='humanitarian' />
                  <Field
                    name="hierarchy"
                    label={ActivityStrings.hierarchy}
                    component={Select}
                    options={Helpers.hierarchy}
                  />
                  <ErrorMessage name='hierarchy' />
                  <Field
                    name="budgetNotProvided"
                    label={ActivityStrings.budgetNotProvided}
                    component={Select}
                    options={Helpers.budgetNotProvided}
                  />
                  <ErrorMessage name='hierarchy' />
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

const mapStateToProps = (state: ApplicationState): ActivityKeyProps => {
  //console.log(state.orgReader)
  return {
    activitiesRef: state.keys.data.activities,
    activityRef: state.keys.data.newKey
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): ActivityDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setActivity(ownProps)),
    initialise: () => dispatch(initialise()),
    setFormFunctions: (formProps: FormData) => dispatch(setFormFunctions(formProps)),
    newKey: () => dispatch(newKey())
  }
}

export const Activity = withTheme(withStyles(styles)(connect<ActivityKeyProps, ActivityDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(ActivityForm)))