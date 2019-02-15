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

import { OrganisationsPicker } from '../../../components/io/organisationsPicker'
import { OrganisationPicker } from '../../../components/io/organisationPicker'

import { initialise, getDocs } from '../../../store/IATI/IATIReader/organisations/organisationDocs/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { IATIOrganisationDocReport } from '../../../store/IATI/IATIReader/organisations/organisationDocs/types'
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
  docs: IATIOrganisationDocReport
}

interface OrganisationDocDispatchProps {
  handleSubmit: (values: any) => void
  initialise: () => void
}

type OrganisationDocsReaderProps =  WithStyles<typeof styles> & OrganisationDocProps & OrganisationDocDispatchProps

class Docs extends React.Component<OrganisationDocsReaderProps> {

  state = {
    organisationsRef: "",
    submitFunc: (function(submit: boolean) { return submit }),
    resetFunc: (function() { return null }),
    submitting: false
  }

  constructor (props: OrganisationDocsReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.initialise()
  }

  componentDidUpdate(previousProps: OrganisationDocsReaderProps) {
    if(this.state.submitting) {
      this.setState({submitting: false})
      this.state.submitFunc(false)
      this.state.resetFunc()
    }
  }

  handleSubmit = (values: OrganisationsReportProps, setSubmitting: Function, reset: Function) => {
    this.setState({submitFunc: setSubmitting, resetFunc: reset, submitting: true})
    this.props.initialise()
    this.props.handleSubmit(values)
  }

  handleOrganisationsChange = (value: string) => {
    this.setState({organisationsRef: value})
  }

  handleOrganisationChange = (value: string) => {
    console.log(value)
  }

  render() {

    const docsData = this.props.docs
    let xs = ""
    let num = 0
    Object.keys(docsData).forEach((organisationsKey) => {
      //numOrganisations += 1
      xs += `**${OrganisationDocStrings.organisationsReference}**: ${organisationsKey}<br />`
      Object.keys(docsData[organisationsKey].data).forEach((organisationKey) => {

        xs += `**${OrganisationDocStrings.organisationReference}**: ${organisationKey}<br />`
        Object.keys(docsData[organisationsKey].data[organisationKey].data).forEach((docKey) => {
          if ( docsData[organisationsKey].data[organisationKey].data[docKey].hasOwnProperty('title') &&
               docsData[organisationsKey].data[organisationKey].data[docKey].title != "" ) {

            num += 1
            const thisDocData =  docsData[organisationsKey].data[organisationKey].data[docKey]
            const category = ethers.utils.parseBytes32String(thisDocData.category)
            const countryRef = ethers.utils.parseBytes32String(thisDocData.countryRef)
            const lang = ethers.utils.parseBytes32String(thisDocData.lang)
            const date = ethers.utils.parseBytes32String(thisDocData.date)
            xs+= `**${OrganisationDocStrings.docReference}**: ${docKey} <br />`
            xs+= `**${OrganisationDocStrings.documentTitle}**: ${thisDocData.title} <br />`
            xs+= `**${OrganisationDocStrings.documentFormat}**: ${thisDocData.format} <br />`
            xs+= `**${OrganisationDocStrings.documentURL}**: ${thisDocData.url} <br />`
            xs+= `**${OrganisationDocStrings.documentCategory}**: ${category} <br />`
            xs+= `**${OrganisationDocStrings.documentCountryRef}**: ${countryRef} <br />`
            xs+= `**${OrganisationDocStrings.documentDesc}**: ${thisDocData.desc} <br />`
            xs+= `**${OrganisationDocStrings.documentLang}**: ${lang} <br />`
            xs+= `**${OrganisationDocStrings.documentDate}**: ${date} <br /><br />`
          }
        })
      })
    })

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
                    changeFunction={this.handleOrganisationsChange}
                    setValue={formProps.setFieldValue}
                    name='organisationsRef'
                    label={OrganisationDocStrings.organisationsReference}
                  />
                  <ErrorMessage name='organisationsRef' />
                  <OrganisationPicker
                    organisationsRef={this.state.organisationsRef}
                    changeFunction={this.handleOrganisationChange}
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
        <p>
          <b>{OrganisationDocStrings.numDocs}</b>: {num}
        </p>
        <h3>{OrganisationDocStrings.organisationDocDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrganisationDocProps => {
  //console.log(state.orgReader)
  return {
    docs: state.organisationDocsReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrganisationDocDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(getDocs(ownProps)),
    initialise: () => dispatch(initialise())
  }
}

export const OrganisationDocs = withTheme(withStyles(styles)(connect<OrganisationDocProps, OrganisationDocDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(Docs)))
