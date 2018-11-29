import { PayloadProps } from './types'

export const storeAction = (type: string) => (payload: PayloadProps): object => {
  return {
    type: type,
    payload: payload
  }
}
