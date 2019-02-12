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

import { getBudgets } from '../../../store/IATI/IATIReader/organisations/organisationBudgets/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { IATIOrganisationsData, OrganisationsReportProps } from '../../../store/IATI/IATIReader/organisations/types'

import { OrganisationBudget as OrgBudgetStrings } from '../../../utils/strings'

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

interface OrgBudgetProps {
  organisations: IATIOrganisationsData
}

interface OrgBudgetDispatchProps {
  handleSubmit: (values: any) => void
}

type OrganisationBudgetsReaderProps =  WithStyles<typeof styles> & OrgBudgetProps & OrgBudgetDispatchProps

class Budgets extends React.Component<OrganisationBudgetsReaderProps> {

  state = {
    organisationsRef: "",
    submitFunc: (function(submit: boolean) { return submit }),
    resetFunc: (function() { return null })
  }

  constructor (props: OrganisationBudgetsReaderProps) {
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
        xs += `**${OrgBudgetStrings.reportReference}**: ${reportKey}<br />`
        const values = Object.values(this.props.orgBudgets[reportKey])
        //console.log('Values: ', values)
        xs += `**${OrgBudgetStrings.numBudgets}**: ${values[0]} <br /><br />`
        Object.keys(values[1]).forEach((budgetKey) => {
          //console.log('Budget: ', values[1][budgetKey])
          //const version = ethers.utils.parseBytes32String(values[1][thisKey].version)
          if ( values[1][budgetKey].hasOwnProperty('budgetLine') && values[1][budgetKey].budgetLine != "" ) {
            const budgetLine = ethers.utils.parseBytes32String(values[1][budgetKey].budgetLine)
            const start = ethers.utils.parseBytes32String(values[1][budgetKey].finance.start)
            const end = ethers.utils.parseBytes32String(values[1][budgetKey].finance.end)
            xs+= `**${OrgBudgetStrings.reportingOrgRef}**: ${values[1][budgetKey].report.orgRef} <br />`
            xs+= `**${OrgBudgetStrings.budgetReference}**: ${budgetKey} <br />`
            xs+= `**${OrgBudgetStrings.budgetLine}**: ${budgetLine} <br />`
            xs+= `**${OrgBudgetStrings.value}**: ${values[1][budgetKey].finance.value} <br />`
            xs+= `**${OrgBudgetStrings.status}**: ${values[1][budgetKey].finance.status} <br />`
            xs+= `**${OrgBudgetStrings.budgetStart}**: ${start} <br />`
            xs+= `**${OrgBudgetStrings.budgetEnd}**: ${end} <br /><br />`
          }
        })
        length += 1
        length == budgetsData.length ? xs += "" : xs += "---<br /><br />"
      })
    }*/

    return (
      <div>
        <h2>{OrgBudgetStrings.headingOrganisationBudgetReader}</h2>
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
                    label={OrgBudgetStrings.organisationsReference}
                  />
                  <ErrorMessage name='organisationsRef' />
                  <OrganisationPicker
                    organisationsRef={this.state.organisationsRef}
                    changeFunction={this.handleOrganisationChange}
                    name='organisationRef'
                    label={OrgBudgetStrings.organisationReference}
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
          <b>{OrgBudgetStrings.numBudgets}</b>: {num}
        </p>
        <h3>{OrgBudgetStrings.organisationBudgetDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrgBudgetProps => {
  //console.log(state.orgReader)
  return {
    organisations: state.organisationsReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgBudgetDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(getBudgets(ownProps))
  }
}

export const OrganisationBudgets = withTheme(withStyles(styles)(connect<OrgBudgetProps, OrgBudgetDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(Budgets)))
