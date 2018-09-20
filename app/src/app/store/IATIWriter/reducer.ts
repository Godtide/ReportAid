import { Reducer } from 'redux'
import { IATIWriterProps } from '../../containers/pages/IATIWriter/IATIWriter'
import { IATIWriterActionTypes } from './types'

// Type-safe initialState!
const initialState: IATIWriterProps = {
  title: '',
  data: '',
}

// Thanks to Redux 4's much simpler typings, we can take away a lot of typings on the reducer side,
// everything will remain type-safe.
const reducer: Reducer<IATIWriterProps> = (state = initialState, action) => {
  switch (action.type) {
    case IATIWriterActionTypes.FETCH_REQUEST: {
      return { ...state, loading: true }
    }
    case IATIWriterActionTypes.FETCH_SUCCESS: {
      return { ...state, loading: false, data: action.payload }
    }
    case IATIWriterActionTypes.FETCH_ERROR: {
      return { ...state, loading: false, errors: action.payload }
    }
    default: {
      return state
    }
  }
}

// Instead of using default export, we use named exports. That way we can group these exports
// inside the `index.js` folder.
export { reducer as IATIWriterReducer }
