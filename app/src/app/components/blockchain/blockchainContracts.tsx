import { ethers } from 'ethers'
import { Store } from 'redux'

import { addContracts } from '../../store/blockchain/contracts/actions'
import { ContractProps } from '../../store/blockchain/contracts/types'

import { Contract } from '../../utils/config'

import { IATIOrgs } from '../../../blockchain/typechain/IATIOrgs'
import { IATIOrganisationReports } from '../../../blockchain/typechain/IATIOrganisationReports'
import { IATIOrganisationReportDocs } from '../../../blockchain/typechain/IATIOrganisationReportDocs'
import { IATIOrganisationReportBudgets } from '../../../blockchain/typechain/IATIOrganisationReportBudgets'
import { IATIOrganisationReportExpenditure } from '../../../blockchain/typechain/IATIOrganisationReportExpenditure'
import { IATIOrganisationReportRecipientBudgets } from '../../../blockchain/typechain/IATIOrganisationReportRecipientBudgets'
import { IATIOrganisationReportRegionBudgets } from '../../../blockchain/typechain/IATIOrganisationReportRegionBudgets'
import { IATIOrganisationReportCountryBudgets } from '../../../blockchain/typechain/IATIOrganisationReportCountryBudgets'

interface ChainProps {
  store: Store
}

export const setContracts = (props: ChainProps) => {

  const store = props.store
  const state = store.getState()
  const provider = state.chainInfo.data.provider

  if ( provider.hasOwnProperty('connection') ) {
    const state = store.getState()
    const orgContract = state.chainContracts.data.contracts.orgContract
    if ( !(orgContract.hasOwnProperty('getOrganisationExists')) ) {

      const signer = provider.getSigner()
      const contractData: ContractProps = {
        data: {
          contracts: {
            orgContract: new ethers.Contract(Contract.orgsAddress,
                                              Contract.orgsABI,
                                              signer) as IATIOrgs,
            orgReportsContract: new ethers.Contract(Contract.organisationReportsAddress,
                                                     Contract.organisationReportsABI,
                                                     signer) as IATIOrganisationReports,
            orgReportDocsContract: new ethers.Contract(Contract.organisationReportDocsAddress,
                                                     Contract.organisationReportDocsABI,
                                                     signer) as IATIOrganisationReportDocs,
            orgReportBudgetsContract: new ethers.Contract(Contract.organisationReportBudgetsAddress,
                                                     Contract.organisationReportBudgetsABI,
                                                     signer) as IATIOrganisationReportBudgets,
            orgReportExpenditureContract: new ethers.Contract(Contract.organisationReportExpenditureAddress,
                                                     Contract.organisationReportExpenditureABI,
                                                     signer) as IATIOrganisationReportExpenditure,
            orgReportRecipientBudgetsContract: new ethers.Contract(Contract.organisationReportRecipientBudgetsAddress,
                                                     Contract.organisationReportRecipientBudgetsABI,
                                                     signer) as IATIOrganisationReportRecipientBudgets,
            orgReportRegionBudgetsContract: new ethers.Contract(Contract.organisationReportRegionBudgetsAddress,
                                                     Contract.organisationReportRegionBudgetsABI,
                                                     signer) as IATIOrganisationReportRegionBudgets,
            orgReportCountryBudgetsContract: new ethers.Contract(Contract.organisationReportCountryBudgetsAddress,
                                                     Contract.organisationReportCountryBudgetsABI,
                                                     signer) as IATIOrganisationReportCountryBudgets
          }
        }
      }
      //console.log('Setting contracts ', contractData)
      const add = addContracts as Function
      store.dispatch(add(contractData))
    }
  }
}
