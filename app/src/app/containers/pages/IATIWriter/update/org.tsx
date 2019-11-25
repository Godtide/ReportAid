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

import { Org as OrgWriter } from '../create/org'
import { OrgPicker } from '../../../../components/io/orgPicker'

import { getOrgRecord } from '../../../../store/IATI/IATIReader/organisations/orgs/actions'

import { Org as OrgStrings } from '../../../../utils/strings'

import { Paths as PathConfig } from '../../../../utils/config'

const activitySchema = Yup.object().shape({
  orgRef: Yup
    .string()
    .required('Required')
})

interface OrgProps {
  orgRef: string
}

interface OrgDispatchProps {
  handleSubmit: (values: any) => void
}

type OrgFormProps =OrgProps & OrgDispatchProps

export class OrgForm extends React.Component<OrgFormProps> {

  constructor (props: OrgFormProps) {
   super(props)
  }

  handleSubmit = (values: OrgProps, setSubmitting: Function, reset: Function) => {
    this.props.handleSubmit({orgRef: values.orgRef})
    history.push(PathConfig.orgWriter)
  }

  render() {

    return (
      <div>
        <h2>{OrgStrings.headingOrgUpdater}</h2>
        <div>
          <Formik
            initialValues={ {orgRef: ""} }
            enableReinitialize={true}
            validationSchema={activitySchema}
            onSubmit={(values: OrgProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<OrgProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <OrgPicker
                    setValue={formProps.setFieldValue}
                    name='orgRef'
                    label={OrgStrings.orgReference}
                  />
                  <ErrorMessage name='orgRef' />
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
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrgProps => {
  //console.log(state.orgReader)
  return {
    orgRef: state.keys.data.org
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgDispatchProps => {
  return {
    handleSubmit: (props: any) => dispatch(getOrgRecord(props))
  }
}

export const Org = connect<OrgProps, OrgDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrgForm)
