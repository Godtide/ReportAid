import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'
import Markdown from 'react-markdown'

import { getDictEntries } from '../../../components/io/dict'
import { getOrganisations } from '../../../store/IATI/IATIReader/organisations/organisations/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { IATIOrganisationsReport } from '../../../store/IATI/types'

import { Organisations as OrganisationsStrings } from '../../../utils/strings'

interface OrgProps {
  organisations: IATIOrganisationsReport
}

interface OrgDispatchProps {
  getOrganisations: () => void
}

type OrgReaderProps = OrgProps & OrgDispatchProps

class OrganisationsReader extends React.Component<OrgReaderProps> {

  constructor (props: OrgReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.getOrganisations()
  }

  render() {

    const xs = getDictEntries(this.props.organisations)
    return (
      <div>
        <h2>{OrganisationsStrings.headingOrganisationsReader}</h2>
        <hr />
        <h3>{OrganisationsStrings.organisationsDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrgProps => {
  //console.log(state.orgReader)
  return {
    organisations: state.report.data as IATIOrganisationsReport
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgDispatchProps => {
  return {
    getOrganisations: () => dispatch(getOrganisations())
  }
}

export const Organisations = connect<OrgProps, OrgDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrganisationsReader)
