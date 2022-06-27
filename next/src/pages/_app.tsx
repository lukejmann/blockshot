import type { AppProps } from 'next/app'
import ThemeProvider, { ThemedGlobalStyle } from 'theme'
import { Provider } from 'react-redux'
import { StrictMode } from 'react'
import store from 'state'

function Updaters() {
  return <></>
}

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <>
      <Provider store={store}>
        <ThemedGlobalStyle />
        <ThemeProvider>
          <Component {...pageProps} />
        </ThemeProvider>
      </Provider>
    </>
  )
}

export default MyApp
