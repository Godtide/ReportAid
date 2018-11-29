import { ActionProps, PayloadProps } from './types'

export const storeAction = (type: string) => (payload: PayloadProps): ActionProps => {
  return {
    type: type,
    payload: payload
  }
}
