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
import { IATIOrgProps, OrgProps } from '../../../store/IATI/types'

//import { getDictEntries } from '../../../utils/dict'

import { setOrganisation } from '../../../store/IATI/IATIWriter/organisations/actions'

import { OrganisationPicker } from '../../../components/io/orgPicker'
import { TransactionHelper, TransactionTypes } from '../../io/transactionHelper'

import { Org, Organisation } from '../../../utils/strings'
import { Helpers } from '../../../utils/config'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

const reportSchema = Yup.object().shape({
  version: Yup
    .string()
    .required('Required'),
  orgRef: Yup
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
  lang: Yup
    .string()
    .required('Required'),
  currency: Yup
    .string()
    .required('Required'),
})

export interface OrgsDispatchProps {
  handleSubmit: (values: any) => void
}

type OrgsFormProps = WithStyles<typeof styles> & OrgsDispatchProps

export class OrgsForm extends React.Component<OrgsFormProps> {

  state = {
    submitFunc: (function(submit: boolean) { return submit }),
    resetFunc: (function() { return null })
  }

  constructor (props: OrgsFormProps) {
   super(props)
  }

  handleSubmit = (values: OrgProps, setSubmitting: Function, reset: Function) => {
    this.setState({submitFunc: setSubmitting, resetFunc: reset})
    this.props.handleSubmit(values)
  }

  render() {

    return (
      <div>
        <h2>{Organisation.headingOrgWriter}</h2>
        <div>
          <Formik
            initialValues={ {version: "",
                             orgRef: "",
                             reportingOrgRef: "",
                             reportingOrgType: 0,
                             reportingOrgIsSecondary: false,
                             lang: "",
                             currency: ""
                            }}
            validationSchema={reportSchema}
            onSubmit={(values: OrgProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<OrgProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <Field
                    name="version"
                    label={Organisation.version}
                    component={Select}
                    options={Helpers.reportVersions}
                  />
                  <ErrorMessage name='version' />
                  <OrganisationPicker name='orgRef' label={Organisation.orgIdentifier} />
                  <OrganisationPicker name='reportingOrgRef' label={Organisation.reportingOrgRef} />
                  <Field
                    name="reportingOrgType"
                    label={Organisation.reportingOrgType}
                    component={Select}
                    options={Helpers.organisationCodes}
                  />
                  <ErrorMessage name='reportingOrgType' />
                  <Field
                    name="reportingOrgIsSecondary"
                    label={Organisation.reportingOrgIsSecondary}
                    component={Select}
                    options={Helpers.isSecondary}
                  />
                  <ErrorMessage name='reportingOrgIsSecondary' />
                  <Field
                    name="lang"
                    label={Organisation.language}
                    component={Select}
                    options={Helpers.languageCodes}
                  />
                  <ErrorMessage name='lang' />
                  <Field
                    name="currency"
                    label={Organisation.currency}
                    component={Select}
                    options={Helpers.currencyCodes}
                  />
                  <ErrorMessage name='currency' />
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

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgsDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setOrganisation(ownProps))
  }
}

export const Organisations = withTheme(withStyles(styles)(connect<void, OrgsDispatchProps, {}, ApplicationState>(
  null,
  mapDispatchToProps
)(OrgsForm)))
