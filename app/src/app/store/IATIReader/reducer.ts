import { Reducer } from 'redux'

import { IATIReaderProps } from '../../containers/pages/IATIReader/IATIReader'
import { IATIReaderActionTypes } from './types'

// Type-safe initialState!
const initialState: IATIReaderProps = {
  title: '',
  data: '',
}

// Thanks to Redux 4's much simpler typings, we can take away a lot of typings on the reducer side,
// everything will remain type-safe.
const reducer: Reducer<IATIReaderProps> = (state = initialState, action) => {
  switch (action.type) {
    case IATIReaderActionTypes.FETCH_REQUEST: {
      return { ...state, loading: true }
    }
    case IATIReaderActionTypes.FETCH_SUCCESS: {
      return { ...state, loading: false, data: action.payload }
    }
    case IATIReaderActionTypes.FETCH_ERROR: {
      return { ...state, loading: false, errors: action.payload }
    }
    default: {
      return state
    }
  }
}

// Instead of using default export, we use named exports. That way we can group these exports
// inside the `index.js` folder.
export { reducer as IATIReaderReducer }