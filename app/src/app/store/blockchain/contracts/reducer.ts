import { ChainContractActionTypes, ContractProps } from './types'
import { ActionProps } from '../../types'

 const initialOrgContractState: ContractProps = {
    data: {
      contracts: {
        orgContract: {},
        orgReportsContract: {},
        orgReportDocsContract: {},
        orgReportBudgetsContract: {},
        orgReportExpenditureContract: {},
        orgReportRecipientBudgetsContract: {},
        orgReportRegionBudgetsContract: {},
        orgReportCountryBudgetsContract: {}
      }
    }
  }

export const reducer = (state: ContractProps = initialOrgContractState, action: ActionProps): ContractProps => {
  //console.log('Org Contract info: ', action.type, action.payload)
  const payload = action.payload as ContractProps
  if ( typeof payload != 'undefined' ) {
    if ( action.type == ChainContractActionTypes.ADD_ORG ) {
      //console.log('Orgstate: ', state)
      const data: ContractProps = {
        data: {
          contracts: {
            orgContract: payload.data.contracts.orgContract,
            orgReportsContract: state.data.contracts.orgReportsContract,
            orgReportDocsContract: state.data.contracts.orgReportDocsContract,
            orgReportBudgetsContract: state.data.contracts.orgReportBudgetsContract,
            orgReportExpenditureContract: state.data.contracts.orgReportExpenditureContract,
            orgReportRecipientBudgetsContract: state.data.contracts.orgReportRecipientBudgetsContract,
            orgReportRegionBudgetsContract: state.data.contracts.orgReportRegionBudgetsContract,
            orgReportCountryBudgetsContract: state.data.contracts.orgReportCountryBudgetsContract
          }
        }
      }
      return data
    } else if ( action.type == ChainContractActionTypes.ADD_REPORTS ) {
      //console.log('ADD_ORGREPORTS PayloadData ', payload)
      //console.log('ADD_ORGREPORTS State ', state)
      const data: ContractProps = {
        data: {
          contracts: {
            orgContract: state.data.contracts.orgContract,
            orgReportsContract: payload.data.contracts.orgReportsContract,
            orgReportDocsContract: state.data.contracts.orgReportDocsContract,
            orgReportBudgetsContract: state.data.contracts.orgReportBudgetsContract,
            orgReportExpenditureContract: state.data.contracts.orgReportExpenditureContract,
            orgReportRecipientBudgetsContract: state.data.contracts.orgReportRecipientBudgetsContract,
            orgReportRegionBudgetsContract: state.data.contracts.orgReportRegionBudgetsContract,
            orgReportCountryBudgetsContract: state.data.contracts.orgReportCountryBudgetsContract
          }
        }
      }
      //console.log('Blah ', data)
      return data
    } else if ( action.type == ChainContractActionTypes.ADD_REPORTDOCS ) {
      //console.log('ADD_ORGREPORTS PayloadData ', payload)
      //console.log('ADD_ORGREPORTS State ', state)
      const data: ContractProps = {
        data: {
          contracts: {
            orgContract: state.data.contracts.orgContract,
            orgReportsContract: state.data.contracts.orgReportsContract,
            orgReportDocsContract: payload.data.contracts.orgReportDocsContract,
            orgReportBudgetsContract: state.data.contracts.orgReportBudgetsContract,
            orgReportExpenditureContract: state.data.contracts.orgReportExpenditureContract,
            orgReportRecipientBudgetsContract: state.data.contracts.orgReportRecipientBudgetsContract,
            orgReportRegionBudgetsContract: state.data.contracts.orgReportRegionBudgetsContract,
            orgReportCountryBudgetsContract: state.data.contracts.orgReportCountryBudgetsContract
          }
        }
      }
      //console.log('Blah ', data)
      return data
    } else if ( action.type == ChainContractActionTypes.ADD_REPORTBUDGETS ) {
      //console.log('ADD_ORGREPORTS PayloadData ', payload)
      //console.log('ADD_ORGREPORTS State ', state)
      const data: ContractProps = {
        data: {
          contracts: {
            orgContract: state.data.contracts.orgContract,
            orgReportsContract: state.data.contracts.orgReportsContract,
            orgReportDocsContract: state.data.contracts.orgReportDocsContract,
            orgReportBudgetsContract: payload.data.contracts.orgReportBudgetsContract,
            orgReportExpenditureContract: state.data.contracts.orgReportExpenditureContract,
            orgReportRecipientBudgetsContract: state.data.contracts.orgReportRecipientBudgetsContract,
            orgReportRegionBudgetsContract: state.data.contracts.orgReportRegionBudgetsContract,
            orgReportCountryBudgetsContract: state.data.contracts.orgReportCountryBudgetsContract
          }
        }
      }
      //console.log('Blah ', data)
      return data
    } else if ( action.type == ChainContractActionTypes.ADD_REPORTEXPENDITURES ) {
      //console.log('ADD_ORGREPORTS PayloadData ', payload)
      //console.log('ADD_ORGREPORTS State ', state)
      const data: ContractProps = {
        data: {
          contracts: {
            orgContract: state.data.contracts.orgContract,
            orgReportsContract: state.data.contracts.orgReportsContract,
            orgReportDocsContract: state.data.contracts.orgReportDocsContract,
            orgReportBudgetsContract: state.data.contracts.orgReportBudgetsContract,
            orgReportExpenditureContract: payload.data.contracts.orgReportExpenditureContract,
            orgReportRecipientBudgetsContract: state.data.contracts.orgReportRecipientBudgetsContract,
            orgReportRegionBudgetsContract: state.data.contracts.orgReportRegionBudgetsContract,
            orgReportCountryBudgetsContract: state.data.contracts.orgReportCountryBudgetsContract
          }
        }
      }
      //console.log('Blah ', data)
      return data
    } else if ( action.type == ChainContractActionTypes.ADD_REPORTRECIPIENTBUDGETS ) {
      //console.log('ADD_ORGREPORTS PayloadData ', payload)
      //console.log('ADD_ORGREPORTS State ', state)
      const data: ContractProps = {
        data: {
          contracts: {
            orgContract: state.data.contracts.orgContract,
            orgReportsContract: state.data.contracts.orgReportsContract,
            orgReportDocsContract: state.data.contracts.orgReportDocsContract,
            orgReportBudgetsContract: state.data.contracts.orgReportBudgetsContract,
            orgReportExpenditureContract: state.data.contracts.orgReportExpenditureContract,
            orgReportRecipientBudgetsContract: payload.data.contracts.orgReportRecipientBudgetsContract,
            orgReportRegionBudgetsContract: state.data.contracts.orgReportRegionBudgetsContract,
            orgReportCountryBudgetsContract: state.data.contracts.orgReportCountryBudgetsContract
          }
        }
      }
      //console.log('Blah ', data)
      return data
    } else if ( action.type == ChainContractActionTypes.ADD_REPORTREGIONBUDGETS ) {
      //console.log('ADD_ORGREPORTS PayloadData ', payload)
      //console.log('ADD_ORGREPORTS State ', state)
      const data: ContractProps = {
        data: {
          contracts: {
            orgContract: state.data.contracts.orgContract,
            orgReportsContract: state.data.contracts.orgReportsContract,
            orgReportDocsContract: state.data.contracts.orgReportDocsContract,
            orgReportBudgetsContract: state.data.contracts.orgReportBudgetsContract,
            orgReportExpenditureContract: state.data.contracts.orgReportExpenditureContract,
            orgReportRecipientBudgetsContract: state.data.contracts.orgReportRecipientBudgetsContract,
            orgReportRegionBudgetsContract: payload.data.contracts.orgReportRegionBudgetsContract,
            orgReportCountryBudgetsContract: state.data.contracts.orgReportCountryBudgetsContract
          }
        }
      }
      //console.log('Blah ', data)
      return data
    } else if ( action.type == ChainContractActionTypes.ADD_REPORTCOUNTRYBUDGETS ) {
      //console.log('ADD_ORGREPORTS PayloadData ', payload)
      //console.log('ADD_ORGREPORTS State ', state)
      const data: ContractProps = {
        data: {
          contracts: {
            orgContract: state.data.contracts.orgContract,
            orgReportsContract: state.data.contracts.orgReportsContract,
            orgReportDocsContract: state.data.contracts.orgReportDocsContract,
            orgReportBudgetsContract: state.data.contracts.orgReportBudgetsContract,
            orgReportExpenditureContract: state.data.contracts.orgReportExpenditureContract,
            orgReportRecipientBudgetsContract: state.data.contracts.orgReportRecipientBudgetsContract,
            orgReportRegionBudgetsContract: state.data.contracts.orgReportRegionBudgetsContract,
            orgReportCountryBudgetsContract: payload.data.contracts.orgReportCountryBudgetsContract
          }
        }
      }
      //console.log('Blah ', data)
      return data
    } else {
      return state
    }
  } else {
    return state
  }
}
