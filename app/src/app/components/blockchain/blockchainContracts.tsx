import { ethers } from 'ethers'
import { Store } from 'redux'

import { addContracts } from '../../store/blockchain/contracts/actions'
import { ContractProps } from '../../store/blockchain/contracts/types'

import { Contract } from '../../utils/config'

import { IATIOrgs } from '../../../blockchain/typechain/IATIOrgs'
import { IATIOrganisations } from '../../../blockchain/typechain/IATIOrganisations'
import { IATIOrganisation } from '../../../blockchain/typechain/IATIOrganisation'
import { IATIOrganisationDocs } from '../../../blockchain/typechain/IATIOrganisationDocs'
import { IATIOrganisationBudgets } from '../../../blockchain/typechain/IATIOrganisationBudgets'
import { IATIOrganisationExpenditure } from '../../../blockchain/typechain/IATIOrganisationExpenditure'
import { IATIOrganisationRecipientBudgets } from '../../../blockchain/typechain/IATIOrganisationRecipientBudgets'
import { IATIOrganisationRegionBudgets } from '../../../blockchain/typechain/IATIOrganisationRegionBudgets'
import { IATIOrganisationCountryBudgets } from '../../../blockchain/typechain/IATIOrganisationCountryBudgets'
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
    const orgsContract = state.chainContracts.data.contracts.orgs
    if ( !(orgsContract.hasOwnProperty('getNumOrgs')) ) {

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
            organisationDocs: new ethers.Contract(Contract.organisationDocsAddress,
                                                     Contract.organisationDocsABI,
                                                     signer) as IATIOrganisationDocs,
            organisationBudgets: new ethers.Contract(Contract.organisationBudgetsAddress,
                                                     Contract.organisationBudgetsABI,
                                                     signer) as IATIOrganisationBudgets,
            organisationExpenditure: new ethers.Contract(Contract.organisationExpenditureAddress,
                                                     Contract.organisationExpenditureABI,
                                                     signer) as IATIOrganisationExpenditure,
            organisationRecipientBudgets: new ethers.Contract(Contract.organisationRecipientBudgetsAddress,
                                                     Contract.organisationRecipientBudgetsABI,
                                                     signer) as IATIOrganisationRecipientBudgets,
            organisationRegionBudgets: new ethers.Contract(Contract.organisationRegionBudgetsAddress,
                                                     Contract.organisationRegionBudgetsABI,
                                                     signer) as IATIOrganisationRegionBudgets,
            organisationCountryBudgets: new ethers.Contract(Contract.organisationCountryBudgetsAddress,
                                                     Contract.organisationCountryBudgetsABI,
                                                     signer) as IATIOrganisationCountryBudgets,
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
