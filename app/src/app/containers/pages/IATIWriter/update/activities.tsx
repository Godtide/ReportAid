import * as React from 'react'
import { history } from '../../../../utils/history'

import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { Formik, Form, Field, FormikProps, ErrorMessage } from 'formik'
import * as Yup from 'yup'
import { LinearProgress } from '@material-ui/core'
import Button from '@material-ui/core/Button'
import FormControl from '@material-ui/core/FormControl'

import { ApplicationState } from '../../../../store'
import { ActionProps } from '../../../../store/types'
import { IATIActivitiesReport } from '../../../../store/IATI/types'

import { Activities as ActivitiesWriter } from '../create/activities'
import { ActivitiesPicker } from '../../../../components/io/activitiesPicker'

import { getActivitiesRecord } from '../../../../store/IATI/IATIReader/activities/activities/actions'

import { Activities as ActivitiesStrings } from '../../../../utils/strings'

import { Paths as PathConfig } from '../../../../utils/config'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../../styles/theme'

const activitiesSchema = Yup.object().shape({
  activitiesRef: Yup
    .string()
    .required('Required')
})

interface ActivitiesProps {
  activitiesRef: string
}

interface ActivitiesDispatchProps {
  handleSubmit: (values: any) => void
}

type ActivitiesFormProps = WithStyles<typeof styles> & ActivitiesProps & ActivitiesDispatchProps

export class ActivitiesForm extends React.Component<ActivitiesFormProps> {

  constructor (props: ActivitiesFormProps) {
   super(props)
   //console.logthis.props.location.state.isNewRecord
  }

  handleSubmit = (values: ActivitiesProps, setSubmitting: Function, reset: Function) => {
    this.props.handleSubmit({activitiesRef: values.activitiesRef})
    history.push(PathConfig.activitiesWriter)
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
            onSubmit={(values: ActivitiesProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<ActivitiesProps>) => (
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

const mapStateToProps = (state: ApplicationState): ActivitiesProps => {
  //console.log(state.orgReader)
  return {
    activitiesRef: state.keys.data.activities
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): ActivitiesDispatchProps => {
  return {
    handleSubmit: (props: any) => dispatch(getActivitiesRecord(props))
  }
}

export const Activities = withTheme(withStyles(styles)(connect<ActivitiesProps, ActivitiesDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(ActivitiesForm)))
