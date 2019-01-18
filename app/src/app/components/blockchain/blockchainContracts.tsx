import { ethers } from 'ethers'
import { Contract } from '../../utils/config'

import { IATIOrganisations } from '../../../blockchain/typechain/IATIOrganisations'
import { IATIOrganisationReports } from '../../../blockchain/typechain/IATIOrganisationReports'
import { IATIOrganisationReportDocs } from '../../../blockchain/typechain/IATIOrganisationReportDocs'
import { IATIOrganisationReportBudgets } from '../../../blockchain/typechain/IATIOrganisationReportBudgets'
import { IATIOrganisationReportExpenditure } from '../../../blockchain/typechain/IATIOrganisationReportExpenditure'
import { IATIOrganisationReportRecipientBudgets } from '../../../blockchain/typechain/IATIOrganisationReportRecipientBudgets'
import { IATIOrganisationReportRegionBudgets } from '../../../blockchain/typechain/IATIOrganisationReportRegionBudgets'
import { IATIOrganisationReportCountryBudgets } from '../../../blockchain/typechain/IATIOrganisationReportCountryBudgets'

interface ChainProps {
  provider: any
}

export const getOrgContract = async (props: ChainProps) => {

  const provider = props.provider
  let orgContract = undefined
  if ( provider.hasOwnProperty('connection') ) {
    const signer = provider.getSigner()
    orgContract = new ethers.Contract(Contract.organisationsAddress,
                                      Contract.organisationsABI,
                                      signer) as IATIOrganisations
  }
  return orgContract
}

export const getOrgReportsContract = async (props: ChainProps) => {

  const provider = props.provider
  let orgReportsContract = undefined
  if ( provider.hasOwnProperty('connection') ) {
    const signer = provider.getSigner()
    orgReportsContract = new ethers.Contract(Contract.organisationReportsAddress,
                                             Contract.organisationReportsABI,
                                             signer) as IATIOrganisationReports
  }
  return orgReportsContract
}

export const getOrgReportDocsContract = async (props: ChainProps) => {

  const provider = props.provider
  let orgReportDocsContract = undefined
  if ( provider.hasOwnProperty('connection') ) {
    const signer = provider.getSigner()
    orgReportDocsContract = new ethers.Contract(Contract.organisationReportDocsAddress,
                                             Contract.organisationReportDocsABI,
                                             signer) as IATIOrganisationReportDocs
  }
  return orgReportDocsContract
}

export const getOrgReportBudgetsContract = async (props: ChainProps) => {

  const provider = props.provider
  let orgReportBudgetsContract = undefined
  if ( provider.hasOwnProperty('connection') ) {
    const signer = provider.getSigner()
    orgReportBudgetsContract = new ethers.Contract(Contract.organisationReportBudgetsAddress,
                                             Contract.organisationReportBudgetsABI,
                                             signer) as IATIOrganisationReportBudgets
  }
  return orgReportBudgetsContract
}

export const getOrgReportExpenditureContract = async (props: ChainProps) => {

  const provider = props.provider
  let orgReportExpenditureContract = undefined
  if ( provider.hasOwnProperty('connection') ) {
    const signer = provider.getSigner()
    orgReportExpenditureContract = new ethers.Contract(Contract.organisationReportExpenditureAddress,
                                             Contract.organisationReportExpenditureABI,
                                             signer) as IATIOrganisationReportExpenditure
  }
  return orgReportExpenditureContract
}

export const getOrgReportRecipientBudgetsContract = async (props: ChainProps) => {

  const provider = props.provider
  let orgReportRecipientBudgetsContract = undefined
  if ( provider.hasOwnProperty('connection') ) {
    const signer = provider.getSigner()
    orgReportRecipientBudgetsContract = new ethers.Contract(Contract.organisationReportRecipientBudgetsAddress,
                                             Contract.organisationReportRecipientBudgetsABI,
                                             signer) as IATIOrganisationReportRecipientBudgets
  }
  return orgReportRecipientBudgetsContract
}

export const getOrgReportRegionBudgetsContract = async (props: ChainProps) => {

  const provider = props.provider
  let orgReportRegionBudgetsContract = undefined
  if ( provider.hasOwnProperty('connection') ) {
    const signer = provider.getSigner()
    orgReportRegionBudgetsContract = new ethers.Contract(Contract.organisationReportRegionBudgetsAddress,
                                             Contract.organisationReportRegionBudgetsABI,
                                             signer) as IATIOrganisationReportRegionBudgets
  }
  return orgReportRegionBudgetsContract
}

export const getOrgReportCountryBudgetsContract = async (props: ChainProps) => {

  const provider = props.provider
  let orgReportCountryBudgetsContract = undefined
  if ( provider.hasOwnProperty('connection') ) {
    const signer = provider.getSigner()
    orgReportCountryBudgetsContract = new ethers.Contract(Contract.organisationReportCountryBudgetsAddress,
                                             Contract.organisationReportCountryBudgetsABI,
                                             signer) as IATIOrganisationReportCountryBudgets
  }
  return orgReportCountryBudgetsContract
}
