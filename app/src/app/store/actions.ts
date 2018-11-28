import { ActionProps } from './types'

export const storeAction = (type: string) => (payload: object): ActionProps => {
  return {
    type: type,
    payload: payload
  }
}
