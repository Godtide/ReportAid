import { setProvider } from '../../components/blockchain/blockchainProvider'
import { connect } from 'react-redux'
import { Dispatch } from 'redux'
import { ApplicationState } from '../../store'
import { addInfo, addObjects } from '../../store/blockchain/actions'

import { BlockchainProps } from '../../store/blockchain/types'

const mapStateToProps = (state: ApplicationState): BlockchainProps => {
  return {
    APIName: state.blockchain.APIName,
    networkName: state.blockchain.networkName,
    networkChainId: state.blockchain.networkChainId,
    networkENSAddress: state.blockchain.networkENSAddress,
    account: state.blockchain.account,
    web3: state.blockchain.web3,
    ethers: state.blockchain.ethers
  }
}

const mapDispatchToProps = (dispatch: Dispatch, ownProps: any): any => {
  return({
    addInfo: () => {dispatch(addInfo(ownProps))},
    addObjects: () => {dispatch(addObjects(ownProps))}
  })
}

export const SetProvider = connect<BlockchainProps, any, any, ApplicationState>(
    mapStateToProps, mapDispatchToProps)(
    setProvider
)
