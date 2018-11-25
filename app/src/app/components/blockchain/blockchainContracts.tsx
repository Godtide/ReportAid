import { ethers } from 'ethers'
import { OrgsContract } from '../../utils/config'

interface OrgWriterProps {
  provider: any
}

export const getOrgContract = async (props: OrgWriterProps) => {

  const provider = props.provider
  let orgContract = undefined
  if ( provider.hasOwnProperty('connection') ) {
    const signer = provider.getSigner()
    orgContract = new ethers.Contract(OrgsContract.organisationsAddress, OrgsContract.organisationsABI, signer)
  }
  //console.log('Contract', orgContract)
  return orgContract
}
