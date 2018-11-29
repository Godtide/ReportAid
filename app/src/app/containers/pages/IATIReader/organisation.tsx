import * as React from 'react'
import { connect } from 'react-redux'
//import { Dispatch, AnyAction } from 'redux'

import { IATIOrganisations } from '../../../../blockchain/typechain/IATIOrganisations'

import { OrgContractProps } from '../../../store/blockchain/contracts/types'
import { ApplicationState } from '../../../store'
import { Organisation as OrgStrings } from '../../../utils/strings'

import { PlainTextWithTitle } from '../../../components/io/plainText'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'


type OrgReaderProps =  WithStyles<typeof styles> & OrgContractProps

export class OrgReader extends React.Component<OrgReaderProps> {

  contract: IATIOrganisations
  numOrgs: string

  constructor (props: OrgReaderProps) {
    super(props)
    this.contract = props.data.contract as IATIOrganisations
    this.numOrgs = '0'
    this.getOrgs()
  }

  getOrgs = async () => {
    let numOrgs = await this.contract.getNumOrganisations()
    this.numOrgs = numOrgs.toString()
    console.log(this.numOrgs)
  }

  render() {

    return (
      <div>
        <h2>{OrgStrings.headingOrgWriter}</h2>
        <PlainTextWithTitle title='Num Orgs' text={this.numOrgs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrgContractProps => {
  return {
    data: {
      contract: state.chainOrgContract.data.contract
    }
  }
}

export const OrganisationReader = withTheme(withStyles(styles)(connect<OrgContractProps, {}, {}, ApplicationState>(
  mapStateToProps
)(OrgReader)))
