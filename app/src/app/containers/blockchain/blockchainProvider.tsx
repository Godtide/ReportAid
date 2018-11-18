import { setProvider } from '../../components/blockchain/blockchainProvider'
import { connect } from 'react-redux'
import { Dispatch } from 'redux'
import { addInfo, addObjects } from '../../store/blockchain/actions'

interface BlockchainProviderProps {
  addInfo: () => any,
  addObjects: () => any
}

const mapDispatchToProps = (dispatch: Dispatch, ownProps: any): BlockchainProviderProps => {
  return({
    addInfo: () => {dispatch(addInfo(ownProps))},
    addObjects: () => {dispatch(addObjects(ownProps))}
  })
}

export const SetProvider = connect<void,void,void,void>(
    null, mapDispatchToProps)(
    setProvider
)
