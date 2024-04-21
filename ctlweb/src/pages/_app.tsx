import type { AppProps } from 'next/app'
import '@radix-ui/themes/styles.css'
import { Theme } from '@radix-ui/themes'
import { Client, cacheExchange, fetchExchange, Provider as GraphProvider } from 'urql'
import '@/styles/app.css'

export default function App({ Component, pageProps }: AppProps) {
  const client = new Client({
    url: process.env.NEXT_PUBLIC_GRAPHQL_ENDPOINT as string,
    exchanges: [cacheExchange, fetchExchange],
  })

  return (
    <GraphProvider value={client}>
      <Theme appearance='dark'>
        <Component {...pageProps} />
      </Theme>
    </GraphProvider>
  )
}
