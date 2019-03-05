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
import { ActivitiesProps, ActivityReportProps } from '../../../../store/IATI/types'
import { FormData } from '../../../../store/helpers/forms/types'
import { Keys, KeyTypes } from '../../../../store/helpers/keys/types'

import { ActivitiesPicker } from '../../../../components/io/activitiesPicker'

//import { setKey } from '../../../../store/helpers/keys/actions'
import { getActivitiesRecord } from '../../../../store/IATI/IATIReader/activities/activities/actions'

import { Activities as ActivitiesStrings } from '../../../../utils/strings'
import { Helpers } from '../../../../utils/config'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../../styles/theme'

const activitiesSchema = Yup.object().shape({
  activitiesRef: Yup
    .string()
    .required('Required')
})

interface ActivitiesKeyProps {
  activitiesRef: string
}

interface ActivitiesDispatchProps {
  handleSubmit: (values: ActivityReportProps) => void
}

type ActivitiesFormProps = WithStyles<typeof styles> & ActivitiesKeyProps & ActivitiesDispatchProps

export class ActivitiesForm extends React.Component<ActivitiesFormProps> {

  constructor (props: ActivitiesFormProps) {
   super(props)
  }

  componentDidMount() {
  }

  handleSubmit = (values: ActivitiesKeyProps, setSubmitting: Function, reset: Function) => {
    this.props.handleSubmit({isReport: true, activitiesRef: values.activitiesRef})
  }

  render() {

    return (
      <div>
        <h2>{ActivitiesStrings.headingActivitiesUpdater}</h2>
        <div>
          <Formik
            initialValues={ {activitiesRef: ""} }
            enableReinitialize={true}
            validationSchema={activitiesSchema}
            onSubmit={(values: ActivitiesKeyProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<ActivitiesKeyProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <ActivitiesPicker
                    setValue={formProps.setFieldValue}
                    name='activitiesRef'
                    label={ActivitiesStrings.activitiesReference}
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
    handleSubmit: (props: ActivityReportProps) => dispatch(getActivitiesRecord(props))
  }
}

export const Activities = withTheme(withStyles(styles)(connect<ActivitiesKeyProps, ActivitiesDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(ActivitiesForm)))
