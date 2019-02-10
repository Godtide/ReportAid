import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { Formik, Form, Field, FormikProps, ErrorMessage} from 'formik'
import * as Yup from 'yup'
import { LinearProgress } from '@material-ui/core'
import Button from '@material-ui/core/Button'
import FormControl from '@material-ui/core/FormControl'
//import { Select } from 'formik-material-ui'
import { Select } from "material-ui-formik-components"

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { OrganisationsProps } from '../../../store/IATI/types'

import { setOrganisations } from '../../../store/IATI/IATIWriter/organisations/organisations/actions'

import { TransactionHelper, TransactionTypes } from '../../io/transactionHelper'

import { Organisations as OrganisationsStrings } from '../../../utils/strings'
import { Helpers } from '../../../utils/config'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

const organisationsSchema = Yup.object().shape({
  version: Yup
    .string()
    .required('Required')
})

export interface OrganisationsDispatchProps {
  handleSubmit: (values: any) => void
}

type OrganisationsFormProps = WithStyles<typeof styles> & OrganisationsDispatchProps

export class OrganisationsForm extends React.Component<OrganisationsFormProps> {

  state = {
    submitFunc: (function(submit: boolean) { return submit }),
    resetFunc: (function() { return null })
  }

  constructor (props: OrganisationsFormProps) {
   super(props)
  }

  handleSubmit = (values: OrganisationsProps, setSubmitting: Function, reset: Function) => {
    this.setState({submitFunc: setSubmitting, resetFunc: reset})
    this.props.handleSubmit(values)
  }

  render() {

    return (
      <div>
        <h2>{OrganisationsStrings.headingOrganisationsWriter}</h2>
        <div>
          <Formik
            initialValues={ {version: ""} }
            validationSchema={organisationsSchema}
            onSubmit={(values: OrgProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<OrganisationsProps>) => (
              <Form>
                <FormControl fullWidth={true}>
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
        <TransactionHelper
          type={TransactionTypes.ORGREPORT}
          submitFunc={this.state.submitFunc}
          resetFunc={this.state.resetFunc}
        />
      </div>
    )
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrganisationsDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setOrganisation(ownProps))
  }
}

export const Organisations = withTheme(withStyles(styles)(connect<void, OrganisationsDispatchProps, {}, ApplicationState>(
  null,
  mapDispatchToProps
)(OrganisationsForm)))
