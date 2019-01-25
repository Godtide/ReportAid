import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../store'
import { ActionProps, TxData } from '../../store/types'

import { Transaction } from '../../utils/strings'
import { Helpers } from '../../utils/config'

export const enum TransactionTypes {
  ORG = 'org',
  ORGREPORT = 'orgReport',
  ORGREPORTBUDGET = 'orgReportBudget',
  ORGREPORTEXPENDITURE = 'orgReportExpenditure',
  ORGREPORTRECIPIENTBUDGET = 'orgReportRecipientBudget'
}

export interface TransactionType {
  type: TransactionTypes
  submitFunc: (submit: boolean) => boolean
  resetFunc: () => void
}

interface TransactionProps {
  tx: TxData
}

type TxProps = TransactionType & TransactionProps

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
        txSummary = `${Transaction.success}`
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

const mapStateToProps = (state: ApplicationState, ownProps: TransactionType): TransactionProps => {
  //console.log(state.orgReader)
  switch (ownProps.type) {
    case TransactionTypes.ORG:
      return { tx: state.orgForm.data }
    case TransactionTypes.ORGREPORT:
      return { tx: state.orgReportsForm.data }
    case TransactionTypes.ORGREPORTBUDGET:
      return { tx: state.orgReportBudgetsForm.data }
    case TransactionTypes.ORGREPORTEXPENDITURE:
      return { tx: state.orgReportExpenditureForm.data }
    case TransactionTypes.ORGREPORTRECIPIENTBUDGET:
      return { tx: state.orgReportRecipientBudgetsForm.data }
    default:
      return { tx: {} }
  }
}

export const TransactionHelper = connect<TransactionProps, {}, TransactionType, ApplicationState>(
  mapStateToProps
)(TX)
