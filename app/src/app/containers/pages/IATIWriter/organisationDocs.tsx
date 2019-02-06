import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { Formik, Form, Field, FormikProps, ErrorMessage} from 'formik'
import * as Yup from 'yup'
import { LinearProgress } from '@material-ui/core'
import Button from '@material-ui/core/Button'
//import { Date } from 'formik-material-ui'
import FormControl from '@material-ui/core/FormControl'
import { Select, TextField } from "material-ui-formik-components"

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { IATIOrgReportDocProps, OrgReportDocProps, ReportProps} from '../../../store/IATI/types'

import { setOrganisationReportDoc } from '../../../store/IATI/IATIWriter/organisationReportDocs/actions'

import { FormikDatePicker } from '../../../components/io/datePicker'
import { OrgReportPicker } from '../../../components/io/reportPicker'
import { TransactionHelper, TransactionTypes } from '../../io/transactionHelper'

import { OrganisationReportDoc } from '../../../utils/strings'
import { Helpers } from '../../../utils/config'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

/*const StyledSelect = withStyles({
  root: {
    background: 'linear-gradient(45deg, #FE6B8B 30%, #FF8E53 90%)',
    borderRadius: 3,
    border: 0,
    color: 'white',
    height: 48,
    padding: '0 30px',
    boxShadow: '0 3px 5px 2px rgba(255, 105, 135, .3)',
    width: '10%'
  }
})(Select);*/

const docSchema = Yup.object().shape({
  report: Yup
    .object()
    .required('Required'),
  title: Yup
    .string()
    .required('Required'),
  format: Yup
    .string()
    .required('Required'),
  url: Yup
    .string()
    .required('Required'),
  category: Yup
    .string()
    .required('Required'),
  countryRef: Yup
    .string()
    .required('Required'),
  desc: Yup
    .string()
    .required('Required'),
  lang: Yup
    .string()
    .required('Required'),
  day: Yup
    .number()
    .required('Required'),
  month: Yup
    .number()
    .required('Required'),
  year: Yup
    .number()
    .required('Required'),
})

const DatePickerProps = {
  day: {
    name: 'day',
    label: OrganisationReportDoc.documentDay
  },
  month: {
    name: 'month',
    label: OrganisationReportDoc.documentMonth
  },
  year: {
    name: 'year',
    label: OrganisationReportDoc.documentYear
  }
}

interface OrgReportDocsDispatchProps {
  handleSubmit: (values: any) => void
}

type OrgReportDocsFormProps = WithStyles<typeof styles> & OrgReportDocsDispatchProps

export class OrgReportDocsForm extends React.Component<OrgReportDocsFormProps> {

  state = {
    submitFunc: (function(submit: boolean) { return submit }),
    resetFunc: (function() { return null })
  }

  constructor (props: OrgReportDocsFormProps) {
   super(props)
  }

  handleSubmit = (values: OrgReportDocProps, setSubmitting: Function, reset: Function) => {
    this.setState({submitFunc: setSubmitting, resetFunc: reset})
    this.props.handleSubmit(values)
  }

  render() {

    return (
      <div>
        <h2>{OrganisationReportDoc.headingOrgReportDocWriter}</h2>
        <div>
          <Formik
            initialValues={ {report: {} as ReportProps,
                             title: '',
                             format: '',
                             url: '',
                             category: '',
                             countryRef: '',
                             desc: '',
                             lang: '',
                             day: 1,
                             month: 1,
                             year: 2000
                            }}
            validationSchema={docSchema}
            onSubmit={(values: OrgReportDocProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<OrgReportDocProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <OrgReportPicker name='report' label={OrganisationReportDoc.reportReference} />
                  <Field
                    name='title'
                    label={OrganisationReportDoc.documentTitle}
                    component={TextField}
                  />
                  <ErrorMessage name='title' />
                  <Field
                    name='format'
                    label={OrganisationReportDoc.documentFormat}
                    component={Select}
                    options={Helpers.documentFormats}
                  />
                  <ErrorMessage name='format' />
                  <Field
                    name='url'
                    label={OrganisationReportDoc.documentURL}
                    component={TextField}
                  />
                  <ErrorMessage name='url' />
                  <Field
                    name='category'
                    label={OrganisationReportDoc.documentCategory}
                    component={Select}
                    options={Helpers.documentCategories}
                  />
                  <ErrorMessage name='category' />
                  <Field
                    name='countryRef'
                    label={OrganisationReportDoc.documentCountryRef}
                    component={Select}
                    options={Helpers.countryCodes}
                  />
                  <ErrorMessage name='countryRef' />
                  <Field
                    name='desc'
                    label={OrganisationReportDoc.documentDesc}
                    component={TextField}
                  />
                  <ErrorMessage name='desc' />
                  <Field
                    name="lang"
                    label={OrganisationReportDoc.documentLang}
                    component={Select}
                    options={Helpers.languageCodes}
                  />
                  <ErrorMessage name='lang' />
                  <FormikDatePicker dates={DatePickerProps} />
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
          type={TransactionTypes.ORGREPORTDOC}
          submitFunc={this.state.submitFunc}
          resetFunc={this.state.resetFunc}
        />
      </div>
    )
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgReportDocsDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setOrganisationReportDoc(ownProps))
  }
}

export const OrganisationReportDocs = withTheme(withStyles(styles)(connect<void, OrgReportDocsDispatchProps, {}, ApplicationState>(
  null,
  mapDispatchToProps
)(OrgReportDocsForm)))
