import { ethers } from 'ethers'
import { Contract } from '../../utils/config'

interface ChainProps {
  provider: any
}

export const getOrgContract = async (props: ChainProps) => {

  const provider = props.provider
  let orgContract = undefined
  if ( provider.hasOwnProperty('connection') ) {
    const signer = provider.getSigner()
    orgContract = new ethers.Contract(Contract.organisationsAddress,Contract.organisationsABI, signer)
  }
  //console.log('Contract', orgContract)
  return orgContract
}

export const getOrgReportsContract = async (props: ChainProps) => {

  const provider = props.provider
  let orgReportsContract = undefined
  if ( provider.hasOwnProperty('connection') ) {
    const signer = provider.getSigner()
    orgReportsContract = new ethers.Contract(Contract.organisationReportsAddress,Contract.organisationReportsABI, signer)
  }
  //console.log('Contract', orgReportsContract)
  return orgReportsContract
}
