import * as React from 'react'
//import { Blockchain } from './blockchain'
import { ethers } from 'ethers'

interface OrgsContractProps {
  blockchain: any
}

export class OrgsContract extends React.Component<OrgsContractProps> {

  static organisationsABI = [
    "event OrganisationSet(string _reference, string _name, string _type)",
    "function setOrganisation(string _reference, string _name, string _type)",
    "function getOrganisationExists(string _reference)",
    "function getNumOrganisations()",
    "function getOrganisationReference(uint256 _index)",
    "function getOrganisationName(string _reference)",
    "function getOrganisationType(string _reference)",
  ]

  /*
  organisations: 0x749f861de9e83807e0ebaadaedd88a2f645dc176
  */

  static organisationsAddress = "0x749f861de9e83807e0ebaadaedd88a2f645dc176"

  organisations: any

  constructor (props: OrgsContractProps) {
    super(props)
    //const blockchain = props.blockchain.getProvider()
    //this.organisations = new ethers.Contract(OrgsContract.organisationsAddress, OrgsContract.organisationsABI, blockchain)
  }

  getOrganisations () {
    return this.organisations
  }
}
