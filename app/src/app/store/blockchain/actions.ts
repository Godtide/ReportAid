import { action } from 'typesafe-actions'
import { BlockchainActionTypes, BlockchainProps } from './types'

// Here we use the `action` helper function provided by `typesafe-actions`.
// This library provides really useful helpers for writing Redux actions in a type-safe manner.
// For more info: https://github.com/piotrwitek/typesafe-actions
export const addData = (payload: BlockchainProps) => action(BlockchainActionTypes.ADD_DATA, payload)
