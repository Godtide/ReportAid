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
import { FormData } from '../../../store/helpers/forms/types'

import { initialise } from '../../../store/IATI/IATIReader/actions'
import { getDictEntries } from '../../../components/io/dict'
import { setFormFunctions } from '../../../store/helpers/forms/actions'
import { getOrganisation } from '../../../store/IATI/IATIReader/organisations/organisation/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { IATIOrganisationReport, OrganisationReportProps } from '../../../store/IATI/types'

import { Organisation as OrganisationStrings } from '../../../utils/strings'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

const reportSchema = Yup.object().shape({
  organisationsRef: Yup
    .string()
    .required('Required')
})

interface OrgProps {
  submittingFunc: Function,
  resettingFunc: Function
  organisationsRef: string
  organisation: IATIOrganisationReport
}

interface OrgDispatchProps {
  handleSubmit: (values: any) => void
  initialise: () => void
  setFormFunctions: (formProps: FormData) => void
}

type OrgReaderProps =  WithStyles<typeof styles> & OrgProps & OrgDispatchProps

class OrganisationReader extends React.Component<OrgReaderProps> {

  constructor (props: OrgReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.initialise()
  }

  componentDidUpdate(previousProps: OrgReaderProps) {
    if(previousProps.organisation != this.props.organisation) {
      this.props.submittingFunc(false)
      this.props.resettingFunc()
    }
  }

  handleSubmit = (values: OrganisationReportProps, setSubmitting: Function, reset: Function) => {
    this.props.setFormFunctions({submitFunc: setSubmitting, resetFunc: reset})
    this.props.initialise()
    this.props.handleSubmit(values)
  }

  render() {

    const xs = getDictEntries(this.props.organisation)

    return (
      <div>
        <h2>{OrganisationStrings.headingOrganisationReader}</h2>
        <div>
          <Formik
            initialValues={ {organisationsRef: "",
                             isReport: true
                            }}
            validationSchema={reportSchema}
            onSubmit={(values: OrganisationReportProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<OrganisationReportProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <OrganisationsPicker
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
        <h3>{OrganisationStrings.organisationDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

/*
<p>
  <b>{OrganisationStrings.numOrganisations}</b>: {numOrganisations}
</p>
*/

const mapStateToProps = (state: ApplicationState): OrgProps => {
  //console.log(state.orgReader)
  return {
    submittingFunc: state.forms.data.submitFunc,
    resettingFunc: state.forms.data.resetFunc,
    organisationsRef: state.keys.data.organisations,
    organisation: state.report.data as IATIOrganisationReport
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgDispatchProps => {
  return {
    handleSubmit: (ownProps: OrganisationReportProps) => dispatch(getOrganisation(ownProps)),
    initialise: () => dispatch(initialise()),
    setFormFunctions: (formProps: FormData) => dispatch(setFormFunctions(formProps))
  }
}

export const Organisation = withTheme(withStyles(styles)(connect<OrgProps, OrgDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrganisationReader)))
