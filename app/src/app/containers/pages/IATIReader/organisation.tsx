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

import { getOrganisation } from '../../../store/IATI/IATIReader/organisations/organisation/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { IATIOrganisationReport } from '../../../store/IATI/IATIReader/organisations/organisation/types'
import { OrganisationsReportProps } from '../../../store/IATI/IATIReader/organisations/types'

import { Organisation as OrganisationStrings } from '../../../utils/strings'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

const reportSchema = Yup.object().shape({
  organisationsRef: Yup
    .string()
    .required('Required')
})

interface OrgProps {
  organisation: IATIOrganisationReport
}

interface OrgDispatchProps {
  handleSubmit: (values: any) => void
}

type OrgReaderProps =  WithStyles<typeof styles> & OrgProps & OrgDispatchProps

class OrganisationReader extends React.Component<OrgReaderProps> {

  state = {
    organisationsRef: "",
    submitFunc: (function(submit: boolean) { return submit }),
    resetFunc: (function() { return null })
  }

  constructor (props: OrgReaderProps) {
    super(props)
  }

  handleSubmit = (values: OrganisationsReportProps, setSubmitting: Function, reset: Function) => {
    this.setState({submitFunc: setSubmitting, resetFunc: reset})
    //console.log('Values ', values)
    this.props.handleSubmit(values)
  }

  handleOrganisationsChange = (value: string) => {
    this.setState({organisationsRef: value})
  }

  handleOrganisationChange = (value: string) => {
    console.log(value)
  }

  render() {

    const organisationData = this.props.organisation
    let xs = ""
    let numOrganisations = 0
    let numOrganisation = 0
    Object.keys(organisationData).forEach((organisationsKey) => {
      numOrganisations += 1
      xs += `**${OrganisationStrings.organisationsReference}**: ${organisationsKey}<br />`
      Object.keys(organisationData[organisationsKey].data).forEach((organisationKey) => {
        if ( organisationData[organisationsKey].data[organisationKey].hasOwnProperty('lang') &&
             organisationData[organisationsKey].data[organisationKey].lang != "" ) {
          numOrganisation += 1
          const language =  ethers.utils.parseBytes32String(organisationData[organisationsKey].data[organisationKey].lang)
          const currency =  ethers.utils.parseBytes32String(organisationData[organisationsKey].data[organisationKey].currency)
          const lastUpdated =  ethers.utils.parseBytes32String(organisationData[organisationsKey].data[organisationKey].lastUpdatedTime)
          xs += `**${OrganisationStrings.organisationReference}**: ${organisationKey}<br />`
          xs += `**${OrganisationStrings.orgRef}**: ${organisationData[organisationsKey].data[organisationKey].orgRef}<br />`
          xs += `**${OrganisationStrings.reportingOrgRef}**: ${organisationData[organisationsKey].data[organisationKey].reportingOrg.orgRef} <br />`
          xs += `**${OrganisationStrings.reportingOrgType}**: ${organisationData[organisationsKey].data[organisationKey].reportingOrg.orgType} <br />`
          xs += `**${OrganisationStrings.reportingOrgIsSecondary}**: ${organisationData[organisationsKey].data[organisationKey].reportingOrg.isSecondary} <br />`
          xs += `**${OrganisationStrings.language}**: ${language} <br />`
          xs += `**${OrganisationStrings.currency}**: ${currency} <br />`
          xs += `**${OrganisationStrings.lastUpdated}**: ${lastUpdated} <br /><br />`
        }
      })
    })

    return (
      <div>
        <h2>{OrganisationStrings.headingOrganisationReader}</h2>
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
                    label={OrganisationStrings.organisationsReference}
                  />
                  <ErrorMessage name='organisationsRef' />
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
          <b>{OrganisationStrings.numOrganisations}</b>: {numOrganisations}
        </p>
        <p>
          <b>{OrganisationStrings.numOrganisation}</b>: {numOrganisation}
        </p>
        <h3>{OrganisationStrings.organisationDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrgProps => {
  //console.log(state.orgReader)
  return {
    organisation: state.organisationReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(getOrganisation(ownProps))
  }
}

export const Organisation = withTheme(withStyles(styles)(connect<OrgProps, OrgDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrganisationReader)))
