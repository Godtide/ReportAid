import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../store'
import { ActionProps, TxData } from '../../store/types'

import { Transaction } from '../../utils/strings'
import { Helpers } from '../../utils/config'

export interface TransactionFuncs {
  submitFunc: (submit: boolean) => boolean
  resetFunc: () => void
}

interface TransactionProps {
  tx: TxData
}

type TxProps = TransactionProps & TransactionFuncs

class TX extends React.Component<TxProps> {

  state = {
    txKey: '',
    txSummary: '',
    toggleSubmitting: false
  }

  constructor (props: TxProps) {
   super(props)
  }

  componentDidUpdate(previousProps: TxProps) {
    //console.log(this.props.tx)
    if(previousProps.tx != this.props.tx) {
      const submitting = !this.state.toggleSubmitting
      let txKey = ''
      let txSummary = `${Transaction.fail}`
      const txKeys = Object.keys(this.props.tx)
      if (txKeys.length > 0 ) {
        txKey = Object.keys(this.props.tx)[0]
        if (txKey != "-1" ) {
          txSummary = `${Transaction.success}`
        }
      }
      this.setState({txKey: txKey, txSummary: txSummary, toggleSubmitting: submitting})
      this.props.submitFunc(submitting)
      this.props.resetFunc()
    }
  }

  render() {

    return (
      <React.Fragment>
        <hr />
        <h3>{Transaction.heading}</h3>
        <p>
          <b>{Transaction.summary}</b>: {this.state.txSummary}<br />
          <b>{Transaction.key}</b>: {this.state.txKey}
        </p>
      </React.Fragment>
    )
  }
}

const mapStateToProps = (state: ApplicationState): TransactionProps => {
  //console.log(state.orgReader)
  return {
    tx: state.organisationsWriterForms.data
  }
}

export const TransactionHelper = connect<TransactionProps, {}, {}, ApplicationState>(
  mapStateToProps
)(TX)
