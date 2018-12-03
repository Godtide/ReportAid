import { Action } from 'redux'

export interface DictData {
  [key: string]: object
}

export interface PayloadProps {
  data: object
}

export interface ActionProps extends Action {
  type: string
  payload: PayloadProps
}
