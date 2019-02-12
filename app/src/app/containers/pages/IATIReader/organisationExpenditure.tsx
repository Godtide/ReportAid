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

import { getExpenditure } from '../../../store/IATI/IATIReader/organisations/organisationExpenditure/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { IATIOrganisationsData, OrganisationsReportProps } from '../../../store/IATI/IATIReader/organisations/types'

import { OrganisationExpenditure as OrganisationExpenditureStrings } from '../../../utils/strings'

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

interface OrganisationExpenditureProps {
  organisations: IATIOrganisationsData
}

interface OrganisationExpenditureDispatchProps {
  handleSubmit: (values: any) => void
}

type OrganisationExpenditureReaderProps =  WithStyles<typeof styles> & OrganisationExpenditureProps & OrganisationExpenditureDispatchProps

class Expenditure extends React.Component<OrganisationExpenditureReaderProps> {

  state = {
    organisationsRef: "",
    submitFunc: (function(submit: boolean) { return submit }),
    resetFunc: (function() { return null })
  }

  constructor (props: OrganisationExpenditureReaderProps) {
    super(props)
  }

  handleSubmit = (values: OrganisationsReportProps, setSubmitting: Function, reset: Function) => {
    this.setState({submitFunc: setSubmitting, resetFunc: reset})
    this.props.handleSubmit(values)
  }

  handleOrganisationsChange = (value: string) => {
    this.setState({organisationsRef: value})
  }

  handleOrganisationChange = (value: string) => {
    console.log(value)
  }

  render() {

    const reportsData = Object.keys(this.props.organisations)
    let xs = ""
    let num = 0
    /*let xs = ""
    if ( expenditureData.length > 0 ) {
      let length = 0
      //console.log ("Orgsdata: ", orgsData, " length ", orgsData.length )
      expenditureData.forEach((reportKey) => {
        xs += `**${OrganisationExpenditureStrings.reportReference}**: ${reportKey}<br />`
        const values = Object.values(this.props.orgExpenditure[reportKey])
        //console.log('Values: ', values)
        xs += `**${OrganisationExpenditureStrings.numExpenditure}**: ${values[0]} <br /><br />`
        Object.keys(values[1]).forEach((expenditureKey) => {
          //console.log(': ', values[1][expenditureKey])
          //const version = ethers.utils.parseBytes32String(values[1][thisKey].version)
          if ( values[1][expenditureKey].hasOwnProperty('expenditureLine') && values[1][expenditureKey].expenditureLine != "" ) {
            const expenditureLine = ethers.utils.parseBytes32String(values[1][expenditureKey].expenditureLine)
            const start = ethers.utils.parseBytes32String(values[1][expenditureKey].finance.start)
            const end = ethers.utils.parseBytes32String(values[1][expenditureKey].finance.end)
            xs+= `**${OrganisationExpenditureStrings.reportingOrgRef}**: ${values[1][expenditureKey].report.orgRef} <br />`
            xs+= `**${OrganisationExpenditureStrings.expenditureReference}**: ${expenditureKey} <br />`
            xs+= `**${OrganisationExpenditureStrings.expenditureLine}**: ${expenditureLine} <br />`
            xs+= `**${OrganisationExpenditureStrings.value}**: ${values[1][expenditureKey].finance.value} <br />`
            xs+= `**${OrganisationExpenditureStrings.status}**: ${values[1][expenditureKey].finance.status} <br />`
            xs+= `**${OrganisationExpenditureStrings.expenditureStart}**: ${start} <br />`
            xs+= `**${OrganisationExpenditureStrings.expenditureEnd}**: ${end} <br /><br />`
          }
        })
        length += 1
        length == expenditureData.length ? xs += "" : xs += "---<br /><br />"
      })
    }*/

    return (
      <div>
        <h2>{OrganisationExpenditureStrings.headingOrganisationExpenditureReader}</h2>
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
                    name='organisationsRef'
                    label={OrganisationExpenditureStrings.organisationsReference}
                  />
                  <ErrorMessage name='organisationsRef' />
                  <OrganisationPicker
                    organisationsRef={this.state.organisationsRef}
                    changeFunction={this.handleOrganisationChange}
                    name='organisationRef'
                    label={OrganisationExpenditureStrings.organisationReference}
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
          <b>{OrganisationExpenditureStrings.numExpenditure}</b>: {num}
        </p>
        <h3>{OrganisationExpenditureStrings.organisationExpenditureDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrganisationExpenditureProps => {
  //console.log(state.orgReader)
  return {
    organisations: state.organisationsReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrganisationExpenditureDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(getExpenditure(ownProps))
  }
}

export const OrganisationExpenditure = withTheme(withStyles(styles)(connect<OrganisationExpenditureProps, OrganisationExpenditureDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(Expenditure)))
