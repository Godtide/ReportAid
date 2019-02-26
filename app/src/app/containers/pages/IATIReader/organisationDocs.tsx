import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'
import Markdown from 'react-markdown'

import { Formik, Form, Field, FormikProps, ErrorMessage} from 'formik'
import * as Yup from 'yup'
import { LinearProgress } from '@material-ui/core'
import Button from '@material-ui/core/Button'
//import { Date } from 'formik-material-ui'
import FormControl from '@material-ui/core/FormControl'

import { getDictEntries } from '../../../components/io/dict'
import { OrganisationsPicker } from '../../../components/io/organisationsPicker'
import { OrganisationPicker } from '../../../components/io/organisationPicker'
import { FormData } from '../../../store/helpers/forms/types'

import { setFormFunctions } from '../../../store/helpers/forms/actions'
import { initialise, getDocs } from '../../../store/IATI/IATIReader/organisations/organisationDocs/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { IATIOrganisationDocReport } from '../../../store/IATI/IATIReader/organisations/types'
import { OrganisationsReportProps } from '../../../store/IATI/IATIReader/organisations/types'

import { OrganisationDoc as OrganisationDocStrings } from '../../../utils/strings'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

const reportSchema = Yup.object().shape({
  organisationsRef: Yup
    .string()
    .required('Required'),
  organisationRef: Yup
    .string()
    .required('Required')
})

interface OrganisationDocProps {
  submittingFunc: Function,
  resettingFunc: Function
  organisationsRef: string,
  organisationRef: string,
  docs: IATIOrganisationDocReport
}

interface OrganisationDocDispatchProps {
  handleSubmit: (values: any) => void
  initialise: () => void
  setFormFunctions: (formProps: FormData) => void
}

type OrganisationDocsReaderProps =  WithStyles<typeof styles> & OrganisationDocProps & OrganisationDocDispatchProps

class Docs extends React.Component<OrganisationDocsReaderProps> {

  constructor (props: OrganisationDocsReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.initialise()
  }

  componentDidUpdate(previousProps: OrganisationDocsReaderProps) {
    if(previousProps.docs != this.props.docs) {
    this.props.submittingFunc(false)
    this.props.resettingFunc()
    }
  }

  handleSubmit = (values: OrganisationsReportProps, setSubmitting: Function, reset: Function) => {
    this.props.setFormFunctions({submitFunc: setSubmitting, resetFunc: reset})
    this.props.initialise()
    this.props.handleSubmit(values)
  }

  render() {

    const xs = getDictEntries(this.props.docs)

    return (
      <div>
        <h2>{OrganisationDocStrings.headingOrganisationDocReader}</h2>
        <div>
          <Formik
            initialValues={ {organisationsRef: "",
                             organisationRef: ""
                            }}
            validationSchema={reportSchema}
            onSubmit={(values: OrganisationsReportProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<OrganisationsReportProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <OrganisationsPicker
                    setValue={formProps.setFieldValue}
                    name='organisationsRef'
                    label={OrganisationDocStrings.organisationsReference}
                  />
                  <ErrorMessage name='organisationsRef' />
                  <OrganisationPicker
                    setValue={formProps.setFieldValue}
                    name='organisationRef'
                    label={OrganisationDocStrings.organisationReference}
                  />
                  <ErrorMessage name='organisationRef' />
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
        <hr />
        <h3>{OrganisationDocStrings.organisationDocDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrganisationDocProps => {
  //console.log(state.orgReader)
  return {
    submittingFunc: state.forms.data.submitFunc,
    resettingFunc: state.forms.data.resetFunc,
    organisationsRef: state.keys.data.organisations,
    organisationRef: state.keys.data.organisation,
    docs: state.organisationDocsReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrganisationDocDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(getDocs(ownProps)),
    initialise: () => dispatch(initialise()),
    setFormFunctions: (formProps: FormData) => dispatch(setFormFunctions(formProps))
  }
}

export const OrganisationDocs = withTheme(withStyles(styles)(connect<OrganisationDocProps, OrganisationDocDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(Docs)))
