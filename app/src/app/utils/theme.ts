enum Colours {
  DARK = 'dark',
  LIGHT = 'light'
}

enum Mode {
  HORIZONTAL = 'horizontal',
  INLINE = 'inline'
}

export type ThemeColours = Colours.DARK | Colours.LIGHT
export type MenuMode = Mode.HORIZONTAL | Mode.INLINE

// Declare state types with `readonly` modifier to get compile time immutability.
// https://github.com/piotrwitek/react-redux-typescript-guide#state-with-type-level-immutability
export interface LayoutState {
  readonly theme: ThemeColours
  readonly mode: MenuMode
}
