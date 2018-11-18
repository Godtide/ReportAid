import { setProvider } from '../../components/blockchain/blockchainProvider'
import { connect } from 'react-redux'
import { Dispatch } from 'redux'
import { ApplicationState } from '../../store'
import { addInfo, addObjects } from '../../store/blockchain/actions'

const mapDispatchToProps = (dispatch: Dispatch, ownProps: any): any => {
  return({
    addInfo: () => {dispatch(addInfo(ownProps))},
    addObjects: () => {dispatch(addObjects(ownProps))}
  })
}

export const SetProvider = connect<void, any, void, ApplicationState>(
    null, mapDispatchToProps)(
    setProvider
)
