import React, { useMemo } from 'react'
import { Text, TextProps as TextPropsOriginal } from 'rebass'
import styled, {
  createGlobalStyle,
  css,
  DefaultTheme,
  ThemeProvider as StyledComponentsThemeProvider,
} from 'styled-components'

import { useIsDarkMode } from '../state/user/hooks'
import { Colors } from './styled'

export * from './components'

type TextProps = Omit<TextPropsOriginal, 'css'>

export const MEDIA_WIDTHS = {
  upToExtraSmall: 500,
  upToSmall: 720,
  upToMedium: 960,
  upToLarge: 1280,
}

const mediaWidthTemplates: { [width in keyof typeof MEDIA_WIDTHS]: typeof css } = Object.keys(MEDIA_WIDTHS).reduce(
  (accumulator, size) => {
    ;(accumulator as any)[size] = (a: any, b: any, c: any) => css`
      @media (max-width: ${(MEDIA_WIDTHS as any)[size]}px) {
        ${css(a, b, c)}
      }
    `
    return accumulator
  },
  {}
) as any

const white = '#FFFFFF'
const black = '#000000'

function colors(darkMode: boolean): Colors {
  return {
    darkMode,
    // base
    white,
    black,
    transparent: 'transparent',
  }
}

function theme(darkMode: boolean): DefaultTheme {
  return {
    ...colors(darkMode),

    grids: {
      sm: 8,
      md: 12,
      lg: 24,
    },

    shadow1: darkMode ? '#000' : '#2F80ED',

    boxShadow1: darkMode ? '1px 2px 4px hsl(220deg 60% 50%);' : '1px 2px 4px hsl(220deg 60% 50%)',

    mediaWidth: mediaWidthTemplates,

    flexColumnNoWrap: css`
      display: flex;
      flex-flow: column nowrap;
    `,
    flexRowNoWrap: css`
      display: flex;
      flex-flow: row nowrap;
    `,
  }
}

export default function ThemeProvider({ children }: { children: React.ReactNode }) {
  const darkMode = useIsDarkMode()

  const themeObject = useMemo(() => {
    const t = theme(darkMode)
    return t
  }, [darkMode])

  return <StyledComponentsThemeProvider theme={themeObject}>{children}</StyledComponentsThemeProvider>
}

const TextWrapper = styled(Text)<{ color: keyof Colors }>`
  color: ${({ color, theme }) => (theme as any)[color]};
  letter-spacing: -0.05rem;
  font-family: 'Soehne', sans-serif;
`

export const T = {
  main(props: TextProps) {
    return <TextWrapper fontWeight={300} color={'text2'} {...props} />
  },
  link(props: TextProps) {
    return <TextWrapper fontWeight={600} color={'primary1'} {...props} />
  },
  body(props: TextProps) {
    return <TextWrapper fontWeight={500} fontSize={16} color={'text1'} {...props} />
  },
}

export const ThemedGlobalStyle = createGlobalStyle`
html,
body {
  padding: 0;
  margin: 0;
  font-family: -apple-system, BlinkMacSystemFont, Segoe UI, Roboto, Oxygen,
    Ubuntu, Cantarell, Fira Sans, Droid Sans, Helvetica Neue, sans-serif;
}

a {
  color: inherit;
  text-decoration: none;
}

* {
  box-sizing: border-box;
}

@font-face {
  font-family: "Soehne";
  font-weight: 600;
  font-style: normal;
  font-display: block;
  font-named-instance: "Semibold";
  src: url(/fonts/test-soehne-breit-fett.woff2) format("woff");
}
@font-face {
  font-family: "K2D";
  font-style: normal;
  font-display: block;
  font-named-instance: "Semibold";
  src: url(/fonts/K2D-SemiBold.ttf) format("truetype");
}
`
