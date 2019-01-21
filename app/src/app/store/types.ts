import { Action } from 'redux'

export interface TxData {
  [tx: string]: object
}

export interface DictData {
  [key: string]: object
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
