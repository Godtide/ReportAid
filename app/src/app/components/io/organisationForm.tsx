import * as React from 'react'
import { connect } from 'react-redux'
import { Dispatch, AnyAction } from 'redux'
import { Field } from 'redux-form'
import { ApplicationState } from '../../store'

import { Organisation } from '../../utils/strings'
import { addOrganisation } from '../../store/IATIWriter/organisationWriter/actions'

interface OrganisationFormProps {
  pristine: boolean
  submitting: boolean
}

interface DispatchProps {
  handleSubmit: (ownProps: any) => void
}

type AllProps = OrganisationFormProps & DispatchProps

const Form: React.SFC<AllProps> = (props: AllProps) => {

  return (
    <div>
      <form onSubmit={props.handleSubmit}>
        <div>
          <label>{Organisation.orgName}</label>
          <Field
            name={Organisation.orgName}
            component="input"
            type="text"
            placeholder={Organisation.orgName}
          />
        </div>
        <div>
          <label>{Organisation.identifier}</label>
          <Field
            name={Organisation.identifier}
            component="input"
            type="text"
            placeholder={Organisation.identifier}
          />
        </div>
        <div>
          <label>{Organisation.type}</label>
          <Field
            name={Organisation.type}
            component="input"
            type="text"
            placeholder={Organisation.type}
          />
        </div>
        <div>
          <button type="submit" disabled={props.pristine || props.submitting}>
            Submit
          </button>
        </div>
      </form>
    </div>
  )
}

const mapDispatchToProps = (dispatch: Dispatch<AnyAction>): DispatchProps => {
  return ({
    handleSubmit: (ownProps: any) => dispatch(addOrganisation(ownProps))
  })
}

export const OrganisationForm = connect<{}, DispatchProps, {}, ApplicationState>(
  null,
  mapDispatchToProps
)(Form)
