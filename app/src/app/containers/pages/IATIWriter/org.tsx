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
import { OrgProps } from '../../../store/IATI/IATIWriter/organisations/types'
import { KeyData } from '../../../store/helpers/keys/types'
import { FormData } from '../../../store/helpers/forms/types'

import { initialise } from '../../../store/IATI/IATIWriter/organisations/actions'
import { setFormFunctions } from '../../../store/helpers/forms/actions'
import { newKey } from '../../../store/helpers/keys/actions'
import { setOrg } from '../../../store/IATI/IATIWriter/organisations/orgs/actions'

import { TransactionHelper } from '../../io/transactionHelper'

import { Org as OrgStrings } from '../../../utils/strings'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

const orgSchema = Yup.object().shape({
  orgRef: Yup
    .string()
    .required('Required'),
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

interface OrgFormProps {
  orgRef: string
}

interface OrgDispatchProps {
  handleSubmit: (values: any) => void
  initialise: () => void
  setFormFunctions: (formProps: FormData) => void
  newKey: () => void
}

type OrgWriterFormProps = WithStyles<typeof styles> & OrgFormProps & OrgDispatchProps

export class OrgForm extends React.Component<OrgWriterFormProps> {

  constructor (props: OrgWriterFormProps) {
   super(props)
  }

  componentDidMount() {
    this.props.initialise()
    this.props.newKey()
  }

  handleSubmit = (values: OrgProps, setSubmitting: Function, reset: Function) => {
    this.props.setFormFunctions({submitFunc: setSubmitting, resetFunc: reset})
    this.props.handleSubmit(values)
    this.props.initialise()
  }

  render() {

    return (
      <div>
        <h2>{OrgStrings.headingOrgWriter}</h2>
        <div>
          <Formik
            initialValues={ {orgRef: this.props.orgRef, name: '', code: '', identifier: ''} }
            enableReinitialize={true}
            validationSchema={orgSchema}
            onSubmit={(values: OrgProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<OrgProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <Field
                    readOnly
                    name='orgRef'
                    value={this.props.orgRef}
                    label={OrgStrings.orgIdentifier}
                    component={TextField}
                  />
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
        <TransactionHelper/>
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrgFormProps => {
  //console.log(state.orgReader)
  return {
    orgRef: state.keys.data.newKey
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setOrg(ownProps)),
    initialise: () => dispatch(initialise()),
    setFormFunctions: (formProps: FormData) => dispatch(setFormFunctions(formProps)),
    newKey: () => dispatch(newKey())
  }
}

export const Org = withTheme(withStyles(styles)(connect<OrgFormProps, OrgDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrgForm)))
