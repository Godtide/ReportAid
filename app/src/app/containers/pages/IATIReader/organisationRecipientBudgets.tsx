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

import { getRecipientBudgets } from '../../../store/IATI/IATIReader/organisations/organisationRecipientBudgets/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { IATIOrganisationsData, OrganisationsReportProps } from '../../../store/IATI/IATIReader/organisations/types'

import { OrganisationRecipientBudget as OrganisationRecipientBudgetStrings } from '../../../utils/strings'

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

interface OrganisationRecipientBudgetProps {
  organisations: IATIOrganisationsData
}

interface OrganisationRecipientBudgetDispatchProps {
  handleSubmit: (values: any) => void
}

type OrganisationRecipientBudgetsReaderProps =  WithStyles<typeof styles> & OrganisationRecipientBudgetProps & OrganisationRecipientBudgetDispatchProps

class RecipientBudgets extends React.Component<OrganisationRecipientBudgetsReaderProps> {

  state = {
    organisationsRef: "",
    submitFunc: (function(submit: boolean) { return submit }),
    resetFunc: (function() { return null })
  }

  constructor (props: OrganisationRecipientBudgetsReaderProps) {
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
    if ( budgetsData.length > 0 ) {
      let length = 0
      //console.log ("Orgsdata: ", orgsData, " length ", orgsData.length )
      budgetsData.forEach((reportKey) => {
        xs += `**${OrganisationRecipientBudgetStrings.reportReference}**: ${reportKey}<br />`
        const values = Object.values(this.props.orgRecipientBudgets[reportKey])
        //console.log('Values: ', values)
        xs += `**${OrganisationRecipientBudgetStrings.numBudgets}**: ${values[0]} <br /><br />`
        Object.keys(values[1]).forEach((budgetKey) => {
          //console.log(': ', values[1][budgetKey])
          //const version = ethers.utils.parseBytes32String(values[1][thisKey].version)
          if ( values[1][budgetKey].hasOwnProperty('budgetLine') && values[1][budgetKey].budgetLine != "" ) {
            const budgetLine = ethers.utils.parseBytes32String(values[1][budgetKey].budgetLine)
            const start = ethers.utils.parseBytes32String(values[1][budgetKey].finance.start)
            const end = ethers.utils.parseBytes32String(values[1][budgetKey].finance.end)
            xs+= `**${OrganisationRecipientBudgetStrings.reportingOrgRef}**: ${values[1][budgetKey].report.orgRef} <br />`
            xs+= `**${OrganisationRecipientBudgetStrings.budgetReference}**: ${budgetKey} <br />`
            xs+= `**${OrganisationRecipientBudgetStrings.orgReference}**: ${values[1][budgetKey].orgRef} <br />`
            xs+= `**${OrganisationRecipientBudgetStrings.budgetLine}**: ${budgetLine} <br />`
            xs+= `**${OrganisationRecipientBudgetStrings.value}**: ${values[1][budgetKey].finance.value} <br />`
            xs+= `**${OrganisationRecipientBudgetStrings.status}**: ${values[1][budgetKey].finance.status} <br />`
            xs+= `**${OrganisationRecipientBudgetStrings.budgetStart}**: ${start} <br />`
            xs+= `**${OrganisationRecipientBudgetStrings.budgetEnd}**: ${end} <br /><br />`
          }
        })
        length += 1
        length == budgetsData.length ? xs += "" : xs += "---<br /><br />"
      })
    }*/

    return (
      <div>
        <h2>{OrganisationRecipientBudgetStrings.headingOrganisationRecipientBudgetReader}</h2>
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
                    label={OrganisationRecipientBudgetStrings.organisationsReference}
                  />
                  <ErrorMessage name='organisationsRef' />
                  <OrganisationPicker
                    organisationsRef={this.state.organisationsRef}
                    changeFunction={this.handleOrganisationChange}
                    setValue={formProps.setFieldValue}
                    name='organisationRef'
                    label={OrganisationRecipientBudgetStrings.organisationReference}
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
          <b>{OrganisationRecipientBudgetStrings.numBudgets}</b>: {num}
        </p>
        <h3>{OrganisationRecipientBudgetStrings.organisationRecipientBudgetDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrganisationRecipientBudgetProps => {
  //console.log(state.orgReader)
  return {
    organisations: state.organisationsReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrganisationRecipientBudgetDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(getRecipientBudgets(ownProps))
  }
}

export const OrganisationRecipientBudgets = withTheme(withStyles(styles)(connect<OrganisationRecipientBudgetProps, OrganisationRecipientBudgetDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(RecipientBudgets)))
