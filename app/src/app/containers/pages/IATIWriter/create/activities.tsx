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
import { ActivitiesProps } from '../../../../store/IATI/types'
import { FormData } from '../../../../store/helpers/forms/types'

import { setFormFunctions } from '../../../../store/helpers/forms/actions'
import { initialise } from '../../../../store/IATI/IATIWriter/actions'
import { newKey } from '../../../../store/helpers/keys/actions'
import { setActivities } from '../../../../store/IATI/IATIWriter/activities/activities/actions'

import { TransactionHelper } from '../../../io/transactionHelper'

import { Activities as ActivitiesStrings } from '../../../../utils/strings'
import { Helpers } from '../../../../utils/config'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../../styles/theme'

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

interface ActivitiesKeyProps {
  activitiesRef: string
}

export interface ActivitiesDispatchProps {
  handleSubmit: (values: any) => void
  initialise: () => void
  setFormFunctions: (formProps: FormData) => void
}

type ActivitiesFormProps = WithStyles<typeof styles> & ActivitiesKeyProps & ActivitiesDispatchProps

export class ActivitiesForm extends React.Component<ActivitiesFormProps> {

  constructor (props: ActivitiesFormProps) {
   super(props)
  }

  componentDidMount() {
    this.props.initialise()
  }

  handleSubmit = (values: ActivitiesProps, setSubmitting: Function, reset: Function) => {
    this.props.setFormFunctions({submitFunc: setSubmitting, resetFunc: reset})
    this.props.handleSubmit(values)
    this.props.initialise()
  }

  render() {

    return (
      <div>
        <h2>{ActivitiesStrings.headingActivitiesWriter}</h2>
        <div>
          <Formik
            initialValues={ {activitiesRef: this.props.activitiesRef, version: "", linkedData: ""} }
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

const mapStateToProps = (state: ApplicationState): ActivitiesKeyProps => {
  //console.log(state.orgReader)
  return {
    activitiesRef: state.keys.data.activities
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): ActivitiesDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setActivities(ownProps)),
    initialise: () => dispatch(initialise()),
    setFormFunctions: (formProps: FormData) => dispatch(setFormFunctions(formProps))
  }
}

export const Activities = withTheme(withStyles(styles)(connect<ActivitiesKeyProps, ActivitiesDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(ActivitiesForm)))
