import * as React from 'react'
import { connect, Dispatch } from 'react-redux'

import { ApplicationState } from '../../store'
import { ThemeColours, MenuMode } from '../../store/layout'
import * as layoutActions from '../../store/layout/actions'

// Separate state props + dispatch props to their own interfaces.

// Props passed from mapStateToProps
interface PropsFromState {
  theme: ThemeColours
  mode: MenuMode
}

// Props passed from mapDispatchToProps
interface PropsFromDispatch {
  setTheme: typeof layoutActions.setTheme
  setMode: typeof layoutActions.setMode
}

// Component-specific props.
interface OtherProps {
  children: (props: LayoutContainerProps) => React.ReactNode
}

// Combine both state + dispatch props - as well as any props we want to pass - in a union type.
type LayoutContainerProps = PropsFromState & PropsFromDispatch

class LayoutContainer extends React.Component<LayoutContainerProps & OtherProps> {
  public render() {
    const { children, ...rest } = this.props
    return children({ ...rest })
  }
}

// It's usually good practice to only include one context at a time in a connected component.
// Although if necessary, you can always include multiple contexts. Just make sure to
// separate them from each other to prevent prop conflicts.
const mapStateToProps = ({ layout }: ApplicationState) => ({
  theme: layout.theme,
  mode: layout.mode
})

// Mapping our action dispatcher to props is especially useful when creating container components.
const mapDispatchToProps = (dispatch: Dispatch) => ({
  setTheme: (theme: ThemeColours) => dispatch(layoutActions.setTheme(theme)),
  setMode: (mode: MenuMode) => dispatch(layoutActions.setMode(mode))
})

// Now let's connect our component!
// With redux v4's improved typings, we can finally omit generics here.
export default connect(
  mapStateToProps,
  mapDispatchToProps
)(LayoutContainer)
