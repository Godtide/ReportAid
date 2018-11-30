import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { getOverview } from '../../../store/IATI/IATIReader/organisationReader/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'

import { Organisation as OrgStrings } from '../../../utils/strings'

import { PlainTextKeyedWithTitleList } from '../../../components/io/plainText'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

interface OrgProps {
  num: number
  orgs: object
}

interface OrgDispatchProps {
  getOverview: () => void
}

type OrgReaderProps =  WithStyles<typeof styles> & OrgProps & OrgDispatchProps

export class OrgReader extends React.Component<OrgReaderProps> {

  constructor (props: OrgReaderProps) {
    super(props)
    props.getOverview()
  }

  render() {

    const num = this.props.num
    const orgs = this.props.orgs

    console.log('Num: ', num, 'Orgs: ', orgs)

    /*<div>
      <h2>{OrgStrings.headingOrgReader}</h2>
      <p>
        <b>{OrgStrings.numOrgs}</b>: {num}
      </p>
      <hr/>
      <PlainTextKeyedWithTitleList title={OrgStrings.refsHeader} list={refs} />
      <hr />
      <PlainTextKeyedWithTitleList title={OrgStrings.namesHeader} list={names} />
      <hr />
      <PlainTextKeyedWithTitleList title={OrgStrings.typesHeader} list={types} />
    </div>*/

    return (
      <div>
        <h2>{OrgStrings.headingOrgReader}</h2>
        <p>
          <b>{OrgStrings.numOrgs}</b>: {num}
        </p>
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrgProps => {
  //console.log(state.orgReader)
  return {
    num: state.orgReader.num,
    orgs: state.orgReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgDispatchProps => {
  return {
    getOverview: () => dispatch(getOverview())
  }
}

export const OrganisationReader = withTheme(withStyles(styles)(connect<OrgProps, OrgDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrgReader)))
