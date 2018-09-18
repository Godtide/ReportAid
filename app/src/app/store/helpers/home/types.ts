export interface HomePage {
  title: string
  data: string[]
  icon: string
}

// Use `const enum`s for better autocompletion of action type names. These will
// be compiled away leaving only the final value in your compiled code.
//
// Define however naming conventions you'd like for your action types, but
// personally, I use the `@@context/ACTION_TYPE` convention, to follow the convention
// of Redux's `@@INIT` action.
export const enum HomeActionTypes {
  FETCH_REQUEST = '@@IATIReader/FETCH_REQUEST',
  FETCH_SUCCESS = '@@IATIReader/FETCH_SUCCESS',
  FETCH_ERROR = '@@IATIReader/FETCH_ERROR'
}

// Declare state types with `readonly` modifier to get compile time immutability.
// https://github.com/piotrwitek/react-redux-typescript-guide#state-with-type-level-immutability
export interface HomeState {
  readonly data: HomePage[]
  readonly errors?: string
}
