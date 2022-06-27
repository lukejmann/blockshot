import { FlattenSimpleInterpolation, ThemedCssFunction } from 'styled-components/macro'

export type Color = string
export interface Colors {
  darkMode: boolean

  white: Color
  black: Color
  transparent: Color
}

declare module 'styled-components/macro' {
  export interface DefaultTheme extends Colors {
    grids: Grids

    // shadows
    shadow1: string
    boxShadow1: string

    // media queries
    mediaWidth: {
      upToExtraSmall: ThemedCssFunction<DefaultTheme>
      upToSmall: ThemedCssFunction<DefaultTheme>
      upToMedium: ThemedCssFunction<DefaultTheme>
      upToLarge: ThemedCssFunction<DefaultTheme>
    }

    // css snippets
    flexColumnNoWrap: FlattenSimpleInterpolation
    flexRowNoWrap: FlattenSimpleInterpolation
  }
}
