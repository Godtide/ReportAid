import { Action } from 'redux'

export interface TxInfo {
  summary: string,
  info: object
}

export interface TxData {
  [tx: string]: TxInfo
}

export interface TxProps extends PayloadProps {
  data: TxData
}

export interface PayloadProps {
  data: object
}

export interface ActionProps extends Action {
  type: string
  payload: PayloadProps
}
