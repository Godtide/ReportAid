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
import { OrganisationDocProps } from '../../../store/IATI/types'
import { FormData } from '../../../store/helpers/forms/types'

import { setFormFunctions } from '../../../store/helpers/forms/actions'
import { initialise } from '../../../store/IATI/IATIWriter/actions'
import { newKey } from '../../../store/helpers/keys/actions'
import { setOrganisationDoc } from '../../../store/IATI/IATIWriter/organisations/organisationDocs/actions'

import { FormikDatePicker } from '../../../components/io/datePicker'
import { OrganisationsPicker } from '../../../components/io/organisationsPicker'
import { OrganisationPicker } from '../../../components/io/organisationPicker'
import { OrgPicker } from '../../../components/io/orgPicker'
import { TransactionHelper } from '../../io/transactionHelper'

import { OrganisationDoc } from '../../../utils/strings'
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
  organisationsRef: Yup
    .string()
    .required('Required'),
  organisationRef: Yup
    .string()
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
    .required('Required')
})

const DatePickerProps = {
  day: {
    name: 'day',
    label: OrganisationDoc.documentDay
  },
  month: {
    name: 'month',
    label: OrganisationDoc.documentMonth
  },
  year: {
    name: 'year',
    label: OrganisationDoc.documentYear
  }
}

interface OrganisationDocsKeyProps {
  organisationsRef: string
  organisationRef: string
}

interface OrgDocsDispatchProps {
  handleSubmit: (values: any) => void
  initialise: () => void
  setFormFunctions: (formProps: FormData) => void
}

type OrgDocsFormProps = WithStyles<typeof styles> & OrganisationDocsKeyProps & OrgDocsDispatchProps

export class OrgDocsForm extends React.Component<OrgDocsFormProps> {

  constructor (props: OrgDocsFormProps) {
   super(props)
  }

  componentDidMount() {
    this.props.initialise()
  }

  handleSubmit = (values: OrganisationDocProps, setSubmitting: Function, reset: Function) => {
    this.props.setFormFunctions({submitFunc: setSubmitting, resetFunc: reset})
    this.props.initialise()
    this.props.handleSubmit(values)
  }

  render() {

    return (
      <div>
        <h2>{OrganisationDoc.headingOrganisationDocWriter}</h2>
        <div>
          <Formik
            initialValues={ {organisationsRef: "",
                             organisationRef: "",
                             docRef: "",
                             title: '',
                             format: '',
                             url: '',
                             category: '',
                             countryRef: '',
                             desc: '',
                             lang: '',
                             day: 0,
                             month: 0,
                             year: 0
                            }}
            validationSchema={docSchema}
            onSubmit={(values: OrganisationDocProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<OrganisationDocProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <OrganisationsPicker
                    setValue={formProps.setFieldValue}
                    name='organisationsRef'
                    label={OrganisationDoc.organisationsReference}
                  />
                  <ErrorMessage name='organisationsRef' />
                  <OrganisationPicker
                    setValue={formProps.setFieldValue}
                    name='organisationRef'
                    label={OrganisationDoc.organisationReference}
                  />
                  <ErrorMessage name='organisationRef' />
                  <Field
                    name='title'
                    label={OrganisationDoc.documentTitle}
                    component={TextField}
                  />
                  <ErrorMessage name='title' />
                  <Field
                    name='format'
                    label={OrganisationDoc.documentFormat}
                    component={Select}
                    options={Helpers.documentFormats}
                  />
                  <ErrorMessage name='format' />
                  <Field
                    name='url'
                    label={OrganisationDoc.documentURL}
                    component={TextField}
                  />
                  <ErrorMessage name='url' />
                  <Field
                    name='category'
                    label={OrganisationDoc.documentCategory}
                    component={Select}
                    options={Helpers.documentCategories}
                  />
                  <ErrorMessage name='category' />
                  <Field
                    name='countryRef'
                    label={OrganisationDoc.documentCountryRef}
                    component={Select}
                    options={Helpers.countryCodes}
                  />
                  <ErrorMessage name='countryRef' />
                  <Field
                    name='desc'
                    label={OrganisationDoc.documentDesc}
                    component={TextField}
                  />
                  <ErrorMessage name='desc' />
                  <Field
                    name="lang"
                    label={OrganisationDoc.documentLang}
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
        <TransactionHelper/>
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrganisationDocsKeyProps => {
  return {
    organisationsRef: state.keys.data.organisations,
    organisationRef: state.keys.data.organisation
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgDocsDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setOrganisationDoc(ownProps)),
    initialise: () => dispatch(initialise()),
    setFormFunctions: (formProps: FormData) => dispatch(setFormFunctions(formProps))
  }
}

export const OrganisationDocs = withTheme(withStyles(styles)(connect<OrganisationDocsKeyProps, OrgDocsDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrgDocsForm)))
