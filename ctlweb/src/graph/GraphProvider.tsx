import { Client, cacheExchange, fetchExchange, Provider, subscriptionExchange } from 'urql'
import { createClient as createWSClient } from 'graphql-ws'
import { ReactNode } from 'react'

type Props = {
  children: ReactNode
}
export const GraphProvider = ({ children }: Props) => {
  const endpoint = process.env.NEXT_PUBLIC_GRAPH_ENDPOINT as string
  const wsclient = createWSClient({
    url: `ws://${endpoint}`,
  })

  const client = new Client({
    url: `http://${endpoint}`,
    exchanges: [
      cacheExchange,
      fetchExchange,
      subscriptionExchange({
        forwardSubscription: (request) => {
          const input = { ...request, query: request.query || '' }

          return {
            subscribe(sink) {
              const unsubscribe = wsclient.subscribe(input, sink)
              return { unsubscribe }
            },
          }
        },
      }),
    ],
  })

  return <Provider value={client}>{children}</Provider>
}
