import { ethers } from 'ethers'
import { Store } from 'redux'

import { addContracts } from '../../store/blockchain/contracts/actions'
import { ContractProps } from '../../store/blockchain/contracts/types'

import { Contract } from '../../utils/config'

import { IATIOrgs } from '../../../blockchain/typechain/IATIOrgs'
import { IATIOrganisations } from '../../../blockchain/typechain/IATIOrganisations'
import { IATIOrganisation } from '../../../blockchain/typechain/IATIOrganisation'
import { IATIOrganisationReportDocs } from '../../../blockchain/typechain/IATIOrganisationReportDocs'
import { IATIOrganisationReportBudgets } from '../../../blockchain/typechain/IATIOrganisationReportBudgets'
import { IATIOrganisationReportExpenditure } from '../../../blockchain/typechain/IATIOrganisationReportExpenditure'
import { IATIOrganisationReportRecipientBudgets } from '../../../blockchain/typechain/IATIOrganisationReportRecipientBudgets'
import { IATIOrganisationReportRegionBudgets } from '../../../blockchain/typechain/IATIOrganisationReportRegionBudgets'
import { IATIOrganisationReportCountryBudgets } from '../../../blockchain/typechain/IATIOrganisationReportCountryBudgets'
import { IATIActivities } from '../../../blockchain/typechain/IATIActivities'
import { IATIActivity } from '../../../blockchain/typechain/IATIActivity'

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
            orgs: new ethers.Contract(Contract.orgsAddress,
                                              Contract.orgsABI,
                                              signer) as IATIOrgs,
            organisations: new ethers.Contract(Contract.organisationsAddress,
                                                     Contract.organisationsABI,
                                                     signer) as IATIOrganisations,
            organisation: new ethers.Contract(Contract.organisationAddress,
                                                     Contract.organisationABI,
                                                     signer) as IATIOrganisation,
            orgReportDocs: new ethers.Contract(Contract.organisationReportDocsAddress,
                                                     Contract.organisationReportDocsABI,
                                                     signer) as IATIOrganisationReportDocs,
            orgReportBudgets: new ethers.Contract(Contract.organisationReportBudgetsAddress,
                                                     Contract.organisationReportBudgetsABI,
                                                     signer) as IATIOrganisationReportBudgets,
            orgReportExpenditure: new ethers.Contract(Contract.organisationReportExpenditureAddress,
                                                     Contract.organisationReportExpenditureABI,
                                                     signer) as IATIOrganisationReportExpenditure,
            orgReportRecipientBudgets: new ethers.Contract(Contract.organisationReportRecipientBudgetsAddress,
                                                     Contract.organisationReportRecipientBudgetsABI,
                                                     signer) as IATIOrganisationReportRecipientBudgets,
            orgReportRegionBudgets: new ethers.Contract(Contract.organisationReportRegionBudgetsAddress,
                                                     Contract.organisationReportRegionBudgetsABI,
                                                     signer) as IATIOrganisationReportRegionBudgets,
            orgReportCountryBudgets: new ethers.Contract(Contract.organisationReportCountryBudgetsAddress,
                                                     Contract.organisationReportCountryBudgetsABI,
                                                     signer) as IATIOrganisationReportCountryBudgets,
            activities: new ethers.Contract(Contract.activitiesAddress,
                                                     Contract.activitiesABI,
                                                     signer) as IATIActivities,
            activity: new ethers.Contract(Contract.activityAddress,
                                                     Contract.activityABI,
                                                     signer) as IATIActivity,
          }
        }
      }
      //console.log('Setting contracts ', contractData)
      const add = addContracts as Function
      store.dispatch(add(contractData))
    }
  }
}
