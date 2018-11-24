import * as React from 'react'
import { connect } from 'react-redux'
import { ethers } from 'ethers'
import { Provider } from 'ethers/providers'

import { OrgsContract } from '../../utils/config'
import { ApplicationState } from '../../store'

interface OrgProps {
  provider: object
}

class Org extends React.Component<OrgProps> {

  orgContract: object

  constructor(props: OrgProps) {
    super(props)
    /* const provider = ( props.provider as Provider )
    console.log(provider)
    this.orgContract = new ethers.Contract(OrgsContract.organisationsAddress, OrgsContract.organisationsABI, provider)*/
  }

  getContract () {
    return this.orgContract
  }

  render () {
    return null
  }
}

const mapStateToProps = (state: ApplicationState): OrgProps => {
  //console.log("Mapping state", state.blockchain.provider)
  return {
    provider: state.blockchain.provider
  }
}

export const OrganisationContract = connect<OrgProps, void, void, ApplicationState>(
  mapStateToProps
)(Org)
