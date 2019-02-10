import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { Formik, Form, Field, FormikProps, ErrorMessage} from 'formik'
import * as Yup from 'yup'
import { LinearProgress } from '@material-ui/core'
import Button from '@material-ui/core/Button'
import FormControl from '@material-ui/core/FormControl'
//import { TextField } from 'formik-material-ui'
import { TextField } from "material-ui-formik-components"

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { OrgProps } from '../../../store/IATI/types'

import { setOrg } from '../../../store/IATI/IATIWriter/organisations/orgs/actions'

import { TransactionHelper, TransactionTypes } from '../../io/transactionHelper'

import { Org as OrgStrings } from '../../../utils/strings'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

const orgSchema = Yup.object().shape({
  name: Yup
    .string()
    .required('Required'),
  code: Yup
    .string()
    .required('Required'),
  identifier: Yup
    .string()
    .required('Required')
})

interface OrgDispatchProps {
  handleSubmit: (values: any) => void
}

type OrgWriterFormProps = WithStyles<typeof styles> & OrgDispatchProps

export class OrgForm extends React.Component<OrgWriterFormProps> {

  state = {
    submitFunc: (function(submit: boolean) { return submit }),
    resetFunc: (function() { return null })
  }

  constructor (props: OrgWriterFormProps) {
   super(props)
  }

  handleSubmit = (values: OrgProps, setSubmitting: Function, reset: Function) => {
    this.setState({submitFunc: setSubmitting, resetFunc: reset})
    this.props.handleSubmit(values)
  }

  render() {

    return (
      <div>
        <h2>{OrgStrings.headingOrgWriter}</h2>
        <div>
          <Formik
            initialValues={ {name: '', code: '', identifier: ''} }
            validationSchema={orgSchema}
            onSubmit={(values: OrgProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<OrgProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <Field name='name' label={OrgStrings.orgName} component={TextField} />
                  <ErrorMessage name='name' />
                  <br />
                  <Field name='code' label={OrgStrings.code} component={TextField} />
                  <ErrorMessage name='code' />
                  <br />
                  <Field name='identifier' label={OrgStrings.identifier} component={TextField} />
                  <ErrorMessage name='identifier' />
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
        <TransactionHelper
          type={TransactionTypes.ORG}
          submitFunc={this.state.submitFunc}
          resetFunc={this.state.resetFunc}
        />
      </div>
    )
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setOrg(ownProps))
  }
}

export const Org = withTheme(withStyles(styles)(connect<void, OrgDispatchProps, {}, ApplicationState>(
  null,
  mapDispatchToProps
)(OrgForm)))
