import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { Formik, Form, Field, FormikProps, ErrorMessage} from 'formik'
import * as Yup from 'yup'
import { LinearProgress } from '@material-ui/core'
import Button from '@material-ui/core/Button'
import FormControl from '@material-ui/core/FormControl'
//import { Select } from 'formik-material-ui'
import { TextField, Select } from "material-ui-formik-components"

import { ApplicationState } from '../../../../store'
import { ActionProps } from '../../../../store/types'
import { IATIActivityParticipatingOrgReport, ActivityParticipatingOrgProps } from '../../../../store/IATI/types'
import { FormData } from '../../../../store/helpers/forms/types'

import { setFormFunctions } from '../../../../store/helpers/forms/actions'
import { setActivityParticipatingOrg } from '../../../../store/IATI/IATIWriter/activities/activityParticipatingOrgs/actions'

import { ActivitiesPicker } from '../../../../components/io/activitiesPicker'
import { ActivityPicker } from '../../../../components/io/activityPicker'
import { OrgPicker } from '../../../../components/io/orgPicker'
import { TransactionHelper } from '../../../io/transactionHelper'

import { ActivityParticipatingOrg as ActivityParticipatingOrgStrings } from '../../../../utils/strings'
import { Helpers } from '../../../../utils/config'

const activityDateSchema = Yup.object().shape({
  activitiesRef: Yup
    .string()
    .required('Required'),
  activityRef: Yup
    .string()
    .required('Required'),
  participatingOrgRef: Yup
    .string()
    .required('Required'),
  orgRef: Yup
    .string()
    .required('Required'),
  orgType: Yup
    .number()
    .required('Required'),
  role: Yup
    .number()
    .required('Required'),
  crsChannelCode: Yup
    .number()
    .required('Required'),
  lang: Yup
    .string()
    .required('Required'),
  narrative: Yup
    .string()
    .required('Required')
})

interface ActivityParticipatingOrgKeyProps {
  activitiesRef: string
  activityRef: string
  participatingOrgRef: string
  participatingOrgs: IATIActivityParticipatingOrgReport
}

export interface ActivityParticipatingOrgDispatchProps {
  handleSubmit: (values: any) => void
  setFormFunctions: (formProps: FormData) => void
}

type ActivityParticipatingOrgFormProps = ActivityParticipatingOrgKeyProps & ActivityParticipatingOrgDispatchProps

export class ActivityParticipatingOrgForm extends React.Component<ActivityParticipatingOrgFormProps> {

  constructor (props: ActivityParticipatingOrgFormProps) {
   super(props)
  }

  handleSubmit = (values: ActivityParticipatingOrgProps, setSubmitting: Function, reset: Function) => {
    this.props.setFormFunctions({submitFunc: setSubmitting, resetFunc: reset})
    this.props.handleSubmit(values)
  }

  getParticipatingOrgData = (participatingOrg: IATIActivityParticipatingOrgReport): ActivityParticipatingOrgProps => {

    let newParticipatingOrg: ActivityParticipatingOrgProps = {
      activitiesRef: this.props.activitiesRef,
      activityRef: this.props.activityRef,
      participatingOrgRef: this.props.participatingOrgRef,
      orgRef: "",
      orgType: 0,
      role: 0,
      crsChannelCode: 0,
      lang: "",
      narrative: ""
    }
    if ( typeof participatingOrg.data != 'undefined' ) {
      newParticipatingOrg.orgRef = participatingOrg.data[0].orgRef
      newParticipatingOrg.orgType = participatingOrg.data[0].orgType
      newParticipatingOrg.role = participatingOrg.data[0].role
      newParticipatingOrg.crsChannelCode = participatingOrg.data[0].crsChannelCode
      newParticipatingOrg.lang = participatingOrg.data[0].lang
      newParticipatingOrg.narrative = participatingOrg.data[0].narrative
    }

    return newParticipatingOrg
  }

  render() {

    const participatingOrg: ActivityParticipatingOrgProps = this.getParticipatingOrgData(this.props.participatingOrgs)

    return (
      <div>
        <h2>{ActivityParticipatingOrgStrings.headingActivityParticipatingOrgWriter}</h2>
        <div>
          <Formik
            initialValues={ {activitiesRef: this.props.activitiesRef,
                             activityRef: this.props.activityRef,
                             participatingOrgRef: this.props.participatingOrgRef,
                             orgRef: participatingOrg.orgRef,
                             orgType: participatingOrg.orgType,
                             role: participatingOrg.role,
                             crsChannelCode: participatingOrg.crsChannelCode,
                             lang: participatingOrg.lang,
                             narrative: participatingOrg.narrative,
                            }}
            enableReinitialize={true}
            validationSchema={activityDateSchema}
            onSubmit={(values: ActivityParticipatingOrgProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<ActivityParticipatingOrgProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <Field
                    readOnly
                    name='participatingOrgRef'
                    value={this.props.participatingOrgRef}
                    label={ActivityParticipatingOrgStrings.participatingOrgRef}
                    component={TextField}
                  />
                  <ActivitiesPicker
                    setValue={formProps.setFieldValue}
                    name='activitiesRef'
                    label={ActivityParticipatingOrgStrings.activitiesReference}
                  />
                  <ErrorMessage name='activitiesRef' />
                  <ActivityPicker
                    setValue={formProps.setFieldValue}
                    name='activityRef'
                    label={ActivityParticipatingOrgStrings.activityReference}
                  />
                  <ErrorMessage name='activityRef' />
                  <OrgPicker
                    setValue={formProps.setFieldValue}
                    name='orgRef'
                    label={ActivityParticipatingOrgStrings.orgRef}
                  />
                  <ErrorMessage name='orgRef' />
                  <Field
                    name="orgType"
                    label={ActivityParticipatingOrgStrings.orgType}
                    component={Select}
                    options={Helpers.organisationCodes}
                  />
                  <ErrorMessage name='orgType' />
                  <Field
                    name="role"
                    label={ActivityParticipatingOrgStrings.role}
                    component={Select}
                    options={Helpers.participatingOrgRole}
                  />
                  <ErrorMessage name='role' />
                  <Field
                    name="crsChannelCode"
                    label={ActivityParticipatingOrgStrings.crsChannelCode}
                    component={Select}
                    options={Helpers.crsChannelCodes}
                  />
                  <ErrorMessage name='crsChannelCode' />
                  <Field
                    name='narrative'
                    label={ActivityParticipatingOrgStrings.narrative}
                    component={TextField}
                  />
                  <ErrorMessage name='narrative' />
                  <Field
                    name="lang"
                    label={ActivityParticipatingOrgStrings.lang}
                    component={Select}
                    options={Helpers.languageCodes}
                  />
                  <ErrorMessage name='lang' />
                  <br />
                  {formProps.isSubmitting && <LinearProgress />}
                  <br />
                  <Button type='submit' variant="contained" color="primary" disabled={formProps.isSubmitting}>
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

/*
orgType: number
role: number
crsChannelCode: number*/

const mapStateToProps = (state: ApplicationState): ActivityParticipatingOrgKeyProps => {
  //console.log(state.orgReader)
  return {
    activitiesRef: state.keys.data.activities,
    activityRef: state.keys.data.activity,
    participatingOrgRef: state.keys.data.activityParticipatingOrg,
    participatingOrgs: state.report.data as IATIActivityParticipatingOrgReport
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): ActivityParticipatingOrgDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setActivityParticipatingOrg(ownProps)),
    setFormFunctions: (formProps: FormData) => dispatch(setFormFunctions(formProps))
  }
}

export const ActivityParticipatingOrg = connect<ActivityParticipatingOrgKeyProps, ActivityParticipatingOrgDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(ActivityParticipatingOrgForm)
