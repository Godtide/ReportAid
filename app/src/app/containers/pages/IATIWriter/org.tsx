import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

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

import { initialise } from '../../../store/IATI/IATIWriter/organisations/actions'
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

interface OrgDispatchProps {
  handleSubmit: (values: any) => void
  initialise: () => void
}

type OrgWriterFormProps = WithStyles<typeof styles> & OrgDispatchProps

export class OrgForm extends React.Component<OrgWriterFormProps> {

  state = {
    orgRef: '',
    submitFunc: (function(submit: boolean) { return submit }),
    resetFunc: (function() { return null })
  }

  constructor (props: OrgWriterFormProps) {
   super(props)
  }

  componentDidMount() {
    this.props.initialise()
    this.initialiseRef()
  }

  initialiseRef = () => {
    const orgRef = ethers.utils.formatBytes32String(shortid.generate())
    this.setState({orgRef: orgRef})
  }

  handleSubmit = (values: OrgProps, setSubmitting: Function, reset: Function) => {
    this.setState({submitFunc: setSubmitting, resetFunc: reset})
    this.props.initialise()
    this.props.handleSubmit(values)
    this.initialiseRef()
  }

  render() {

    return (
      <div>
        <h2>{OrgStrings.headingOrgWriter}</h2>
        <div>
          <Formik
            initialValues={ {orgRef: this.state.orgRef, name: '', code: '', identifier: ''} }
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
                    value={this.state.orgRef}
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
        <TransactionHelper
          submitFunc={this.state.submitFunc}
          resetFunc={this.state.resetFunc}
        />
      </div>
    )
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setOrg(ownProps)),
    initialise: () => dispatch(initialise())
  }
}

export const Org = withTheme(withStyles(styles)(connect<void, OrgDispatchProps, {}, ApplicationState>(
  null,
  mapDispatchToProps
)(OrgForm)))
