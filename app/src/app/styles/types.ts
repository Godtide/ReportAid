const enum ThemeColours {
  DARK = 'dark',
  LIGHT = 'light'
}

const enum MenuMode {
  HORIZONTAL = 'horizontal',
  INLINE = 'inline'
}

export type ThemeColours = ThemeColours.DARK | ThemeColours.LIGHT
export type MenuMode = MenuMode.HORIZONTAL | MenuMode.INLINE

// Declare state types with `readonly` modifier to get compile time immutability.
// https://github.com/piotrwitek/react-redux-typescript-guide#state-with-type-level-immutability
export interface LayoutState {
  readonly theme: ThemeColours
  readonly mode: MenuMode
}
