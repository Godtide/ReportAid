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
import { OrganisationsProps } from '../../../store/IATI/IATIWriter/organisations/types'
import { FormData } from '../../../store/helpers/forms/types'

import { setFormFunctions } from '../../../store/helpers/forms/actions'
import { initialise } from '../../../store/IATI/IATIWriter/organisations/actions'
import { newKey } from '../../../store/helpers/keys/actions'
import { setOrganisations } from '../../../store/IATI/IATIWriter/organisations/organisations/actions'

import { TransactionHelper } from '../../io/transactionHelper'

import { Organisations as OrganisationsStrings } from '../../../utils/strings'
import { Helpers } from '../../../utils/config'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

const organisationsSchema = Yup.object().shape({
  organisationsRef: Yup
    .string()
    .required('Required'),
  version: Yup
    .string()
    .required('Required')
})

interface OrganisationsKeyProps {
  organisationsRef: string
}

export interface OrganisationsDispatchProps {
  handleSubmit: (values: any) => void
  initialise: () => void
  setFormFunctions: (formProps: FormData) => void
  newKey: () => void
}

type OrganisationsFormProps = WithStyles<typeof styles> & OrganisationsKeyProps & OrganisationsDispatchProps

export class OrganisationsForm extends React.Component<OrganisationsFormProps> {

  constructor (props: OrganisationsFormProps) {
   super(props)
  }

  componentDidMount() {
    this.props.initialise()
    this.props.newKey()
  }

  handleSubmit = (values: OrganisationsProps, setSubmitting: Function, reset: Function) => {
    this.props.setFormFunctions({submitFunc: setSubmitting, resetFunc: reset})
    this.props.handleSubmit(values)
    this.props.initialise()
  }

  render() {

    return (
      <div>
        <h2>{OrganisationsStrings.headingOrganisationsWriter}</h2>
        <div>
          <Formik
            initialValues={ {organisationsRef: this.props.organisationsRef, version: ""} }
            enableReinitialize={true}
            validationSchema={organisationsSchema}
            onSubmit={(values: OrganisationsProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<OrganisationsProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <Field
                    readOnly
                    name='organisationsRef'
                    value={this.props.organisationsRef}
                    label={OrganisationsStrings.organisationsReference}
                    component={TextField}
                  />
                  <Field
                    name="version"
                    label={OrganisationsStrings.version}
                    component={Select}
                    options={Helpers.reportVersions}
                  />
                  <ErrorMessage name='version' />
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

const mapStateToProps = (state: ApplicationState): OrganisationsKeyProps => {
  //console.log(state.orgReader)
  return {
    organisationsRef: state.keys.data.newKey
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrganisationsDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setOrganisations(ownProps)),
    initialise: () => dispatch(initialise()),
    setFormFunctions: (formProps: FormData) => dispatch(setFormFunctions(formProps)),
    newKey: () => dispatch(newKey())
  }
}

export const Organisations = withTheme(withStyles(styles)(connect<OrganisationsKeyProps, OrganisationsDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrganisationsForm)))
