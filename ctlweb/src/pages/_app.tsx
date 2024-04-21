import type { AppProps } from 'next/app'
import '@radix-ui/themes/styles.css'
import { Theme } from '@radix-ui/themes'
import { Client, cacheExchange, fetchExchange, Provider as GraphProvider, subscriptionExchange } from 'urql'
import '@/styles/app.css'
import { Layout } from '@/components/common/Layout'
import { createClient as createWSClient } from 'graphql-ws';

export default function App({ Component, pageProps }: AppProps) {
  const wsClient = createWSClient({
    url: 'ws://localhost:3000/graph',
  })
  const client = new Client({
    url: process.env.NEXT_PUBLIC_GRAPHQL_ENDPOINT as string,
    // fetchSubscriptions: true,
    exchanges: [
      cacheExchange,
      fetchExchange,
      subscriptionExchange({
        forwardSubscription(request) {
          const input = { ...request, query: request.query || '' };
          return {
            subscribe(sink) {
              const unsubscribe = wsClient.subscribe(input, sink);
              return { unsubscribe };
            },
          };
        },
      }),
    ],
  })

  return (
    <GraphProvider value={client}>
      <Theme appearance='dark'>
        <Layout>
          <Component {...pageProps} />
        </Layout>
      </Theme>
    </GraphProvider>
  )
}
