// This file holds our state type, as well as any other types related to this Redux store.

// Response object for GET /heroes
// https://docs.opendota.com/#tag/heroes%2Fpaths%2F~1heroes%2Fget
export interface IATIWriter extends ApiResponse {
  id: number
  name: string
  localized_name: string
  primary_attr: string
  img: string
  icon: string
}

// This type is basically shorthand for `{ [key: string]: any }`. Feel free to replace `any` with
// the expected return type of your API response.
export type ApiResponse = Record<string, any>

// Use `const enum`s for better autocompletion of action type names. These will
// be compiled away leaving only the final value in your compiled code.
//
// Define however naming conventions you'd like for your action types, but
// personally, I use the `@@context/ACTION_TYPE` convention, to follow the convention
// of Redux's `@@INIT` action.
export const enum IATIWriterActionTypes {
  FETCH_REQUEST = '@@IATIWriter/FETCH_REQUEST',
  FETCH_SUCCESS = '@@IATIWriter/FETCH_SUCCESS',
  FETCH_ERROR = '@@IATIWriter/FETCH_ERROR',
  SELECT_HERO = '@@IATIWriter/SELECT_HERO',
  SELECTED = '@@IATIWriter/SELECTED'
}

// Declare state types with `readonly` modifier to get compile time immutability.
// https://github.com/piotrwitek/react-redux-typescript-guide#state-with-type-level-immutability
export interface IATIWriterState {
  readonly loading: boolean
  readonly data: IATIWriter[]
  readonly errors?: string
}
