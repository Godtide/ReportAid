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
import { IATIOrganisationsData, OrganisationsReportProps } from '../../../store/IATI/IATIReader/organisations/types'

import { Organisation as OrganisationStrings } from '../../../utils/strings'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

const reportSchema = Yup.object().shape({
  organisationsRef: Yup
    .string()
    .required('Required')
})

interface OrgProps {
  organisations: IATIOrganisationsData
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
    console.log('Values ', values)
    this.props.handleSubmit(values)
  }

  handleOrganisationsChange = (value: string) => {
    this.setState({organisationsRef: value})
  }

  handleOrganisationChange = (value: string) => {
    console.log(value)
  }

  render() {

    //console.log(JSON.stringify(this.props.organisations))
    //const reportsData = Object.keys(this.props.organisations)
    const reportsData = this.props.organisations
    let xs = ""
    let num = 0
        /*let xs = ""
    if ( orgsData.length > 0 ) {
      let length = 0
      //console.log ("Orgsdata: ", orgsData, " length ", orgsData.length )
      orgsData.forEach((key) => {
        xs += `**${OrgStrings.orgIdentifier}**: ${key}<br />`
        const values = Object.values(this.props.orgs[key])
        //console.log('Values: ', values)
        xs += `**${OrgStrings.numOrgs}**: ${values[0]} <br /><br />`
        Object.keys(values[1]).forEach((thisKey) => {
          //console.log(': ', values[1][thisKey])
          //const version = ethers.utils.parseBytes32String(values[1][thisKey].version)
          if ( values[1][thisKey].hasOwnProperty('version') && values[1][thisKey].version != "" ) {
            const version = ethers.utils.parseBytes32String(values[1][thisKey].version)
            const language =  ethers.utils.parseBytes32String(values[1][thisKey].lang)
            const currency =  ethers.utils.parseBytes32String(values[1][thisKey].currency)
            const lastUpdated =  ethers.utils.parseBytes32String(values[1][thisKey].lastUpdatedTime)
            xs+= `**${OrgStrings.reportingOrgRef}**: ${values[1][thisKey].reportingOrg.orgRef} <br />`
            xs+= `**${OrgStrings.reportKey}**: ${thisKey} <br />`
            xs+= `**${OrgStrings.reportingOrgType}**: ${values[1][thisKey].reportingOrg.orgType} <br />`
            xs+= `**${OrgStrings.reportingOrgIsSecondary}**: ${values[1][thisKey].reportingOrg.isSecondary} <br />`
            xs+= `**${OrgStrings.version}**: ${version} <br />`
            xs+= `**${OrgStrings.language}**: ${language} <br />`
            xs+= `**${OrgStrings.currency}**: ${currency} <br />`
            xs+= `**${OrgStrings.lastUpdated}**: ${lastUpdated} <br /><br />`
          }
        })
        length += 1
        length == orgsData.length ? xs += "" : xs += "---<br /><br />"
      })
    }*/


    //<ErrorMessage name='organisationsRef' />

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
          <b>{OrganisationStrings.numOrganisation}</b>: {num}
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
    organisations: state.organisationsReader.data
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
