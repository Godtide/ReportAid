import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { Formik, Form, Field, FormikProps, ErrorMessage} from 'formik'
import * as Yup from 'yup'
import { LinearProgress } from '@material-ui/core'
import Button from '@material-ui/core/Button'
//import { Select } from 'formik-material-ui'
import { Select } from "material-ui-formik-components"

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { IATIOrgProps, OrgReportProps } from '../../../store/IATI/types'

//import { getDictEntries } from '../../../utils/dict'

import { setOrganisationReport } from '../../../store/IATI/IATIWriter/organisationReports/actions'

import { OrganisationPicker } from '../../../components/io/orgPicker'
import { TransactionHelper, TransactionTypes } from '../../io/transactionHelper'

import { Organisation, OrganisationReport } from '../../../utils/strings'
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

export interface OrgReportsDispatchProps {
  handleSubmit: (values: any) => void
}

type OrgReportsFormProps = WithStyles<typeof styles> & OrgReportsDispatchProps

export class OrgReportsForm extends React.Component<OrgReportsFormProps> {

  state = {
    submitFunc: (function(submit: boolean) { return submit }),
    resetFunc: (function() { return null })
  }

  constructor (props: OrgReportsFormProps) {
   super(props)
  }

  handleSubmit = (values: OrgReportProps, setSubmitting: Function, reset: Function) => {
    this.setState({submitFunc: setSubmitting, resetFunc: reset})
    this.props.handleSubmit(values)
  }

  render() {

    let versions: any[] = []
    Helpers.reportVersions.forEach( (value: any) => {
      //console.log(value, value.code)
      versions.push({ value: value, label: value })
    })

    let orgCodes: any[] = []
    Helpers.organisationCodes.forEach( (value: any) => {
      //console.log(value, value.code)
      orgCodes.push({ value: value.code, label: value.type })
    })

    const isSecondary = [
      { value: false, label: 'No' },
      { value: true, label: 'Yes' }
    ]

    let countries: any[] = [{ value: "", label: "" }]
    Helpers.countryCodes.forEach( (value: any) => {
      //console.log(value, value.code)
      countries.push({ value: value, label: value })
    })

    let currencies: any[] = [{ value: "", label: "" }]
    Helpers.currencyCodes.forEach( (value: any) => {
      //console.log(value, value.code)
      currencies.push({ value: value, label: value })
    })

    return (
      <div>
        <h2>{OrganisationReport.headingOrgReportWriter}</h2>
        <div>
          <Formik
            initialValues={ {version: versions[0].value,
                             orgRef: "",
                             reportingOrgRef: "",
                             reportingOrgType: orgCodes[0].value,
                             reportingOrgIsSecondary: isSecondary[0].value,
                             lang: countries[0].value,
                             currency: currencies[0].value
                            }}
            validationSchema={reportSchema}
            onSubmit={(values: OrgReportProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<OrgReportProps>) => (
              <Form>
                <Field
                  name="version"
                  label={OrganisationReport.version}
                  component={Select}
                  options={versions}
                />
                <ErrorMessage name='version' />
                <OrganisationPicker name='orgRef' label={OrganisationReport.orgIdentifier} />
                <OrganisationPicker name='reportingOrgRef' label={OrganisationReport.reportingOrgRef} />
                <Field
                  name="reportingOrgType"
                  label={OrganisationReport.reportingOrgType}
                  component={Select}
                  options={orgCodes}
                />
                <ErrorMessage name='reportingOrgType' />
                <Field
                  name="reportingOrgIsSecondary"
                  label={OrganisationReport.reportingOrgIsSecondary}
                  component={Select}
                  options={isSecondary}
                />
                <ErrorMessage name='reportingOrgIsSecondary' />
                <Field
                  name="lang"
                  label={OrganisationReport.language}
                  component={Select}
                  options={countries}
                />
                <ErrorMessage name='lang' />
                <Field
                  name="currency"
                  label={OrganisationReport.currency}
                  component={Select}
                  options={currencies}
                />
                <ErrorMessage name='currency' />
                <br />
                {formProps.isSubmitting && <LinearProgress />}
                <br />
                <Button type='submit' variant="raised" color="primary" disabled={formProps.isSubmitting}>
                  Submit
                </Button>
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

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgReportsDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setOrganisationReport(ownProps))
  }
}

export const OrganisationReports = withTheme(withStyles(styles)(connect<void, OrgReportsDispatchProps, {}, ApplicationState>(
  null,
  mapDispatchToProps
)(OrgReportsForm)))
