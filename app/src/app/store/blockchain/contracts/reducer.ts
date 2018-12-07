import { ChainContractActionTypes, ContractProps } from './types'
import { ActionProps } from '../../types'

 const initialOrgContractState: ContractProps = {
    data: {
      contracts: {
        orgContract: {},
        orgReportsContract: {}
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
            orgReportsContract: state.data.contracts.orgReportsContract
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
            orgReportsContract: payload.data.contracts.orgReportsContract
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
