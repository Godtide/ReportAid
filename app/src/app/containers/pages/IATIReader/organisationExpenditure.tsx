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

import { initialise, getExpenditure } from '../../../store/IATI/IATIReader/organisations/organisationExpenditure/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { IATIOrganisationExpenditureReport } from '../../../store/IATI/IATIReader/organisations/organisationExpenditure/types'
import { OrganisationsReportProps } from '../../../store/IATI/IATIReader/organisations/types'

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
  organisationsRef: string,
  organisationRef: string,
  expenditure: IATIOrganisationExpenditureReport
}

interface OrganisationExpenditureDispatchProps {
  handleSubmit: (values: any) => void
  initialise: () => void
}

type OrganisationExpenditureReaderProps =  WithStyles<typeof styles> & OrganisationExpenditureProps & OrganisationExpenditureDispatchProps

class Expenditure extends React.Component<OrganisationExpenditureReaderProps> {

  state = {
    submitFunc: (function(submit: boolean) { return submit }),
    resetFunc: (function() { return null })
  }

  constructor (props: OrganisationExpenditureReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.initialise()
  }

  componentDidUpdate(previousProps: OrganisationExpenditureReaderProps) {
    if(previousProps.expenditure != this.props.expenditure) {
      this.state.submitFunc(false)
      this.state.resetFunc()
    }
  }

  handleSubmit = (values: OrganisationsReportProps, setSubmitting: Function, reset: Function) => {
    this.setState({submitFunc: setSubmitting, resetFunc: reset})
    this.props.initialise()
    this.props.handleSubmit(values)
  }

  render() {

    const expenditureData = this.props.expenditure
    let xs = ""
    let num = 0
    Object.keys(expenditureData).forEach((organisationsKey) => {
      //numOrganisations += 1
      xs += `**${OrganisationExpenditureStrings.organisationsReference}**: ${organisationsKey}<br />`
      Object.keys(expenditureData[organisationsKey].data).forEach((organisationKey) => {

        xs += `**${OrganisationExpenditureStrings.organisationReference}**: ${organisationKey}<br />`
        Object.keys(expenditureData[organisationsKey].data[organisationKey].data).forEach((expenditureKey) => {
          if ( expenditureData[organisationsKey].data[organisationKey].data[expenditureKey].hasOwnProperty('expenditureLine') &&
               expenditureData[organisationsKey].data[organisationKey].data[expenditureKey].expenditureLine != "" ) {

            num += 1
            const thisExpenditureData =  expenditureData[organisationsKey].data[organisationKey].data[expenditureKey]

            const expenditureLine = ethers.utils.parseBytes32String(thisExpenditureData.expenditureLine)
            const start = ethers.utils.parseBytes32String(thisExpenditureData.finance.start)
            const end = ethers.utils.parseBytes32String(thisExpenditureData.finance.end)
            xs += `**${OrganisationExpenditureStrings.expenditureReference}**: ${expenditureKey} <br />`
            xs += `**${OrganisationExpenditureStrings.expenditureLine}**: ${expenditureLine} <br />`
            xs += `**${OrganisationExpenditureStrings.value}**: ${thisExpenditureData.finance.value} <br />`
            xs += `**${OrganisationExpenditureStrings.status}**: ${thisExpenditureData.finance.status} <br />`
            xs += `**${OrganisationExpenditureStrings.expenditureStart}**: ${start} <br />`
            xs += `**${OrganisationExpenditureStrings.expenditureEnd}**: ${end} <br /><br />`
          }
        })
      })
    })

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
                    setValue={formProps.setFieldValue}
                    name='organisationsRef'
                    label={OrganisationExpenditureStrings.organisationsReference}
                  />
                  <ErrorMessage name='organisationsRef' />
                  <OrganisationPicker
                    setValue={formProps.setFieldValue}
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
    organisationsRef: state.keys.data.organisations,
    organisationRef: state.keys.data.organisation,
    expenditure: state.organisationExpenditureReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrganisationExpenditureDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(getExpenditure(ownProps)),
    initialise: () => dispatch(initialise())
  }
}

export const OrganisationExpenditure = withTheme(withStyles(styles)(connect<OrganisationExpenditureProps, OrganisationExpenditureDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(Expenditure)))
