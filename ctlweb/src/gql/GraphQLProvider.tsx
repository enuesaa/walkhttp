import { Client, cacheExchange, fetchExchange, Provider, subscriptionExchange } from 'urql'
import { createClient as createWSClient } from 'graphql-ws'
import { ReactNode, useEffect, useState } from 'react'

type Props = {
  children: ReactNode
}
export const GraphQLProvider = ({ children }: Props) => {
  const [host, setHost] = useState<string|undefined>()
 
  useEffect(() => setHost(window.location.host), [])
  if (host === undefined) {
    return <></>
  }

  const endpoint = process.env.NEXT_PUBLIC_GRAPH_ENDPOINT as string
  const wsclient = createWSClient({
    url: `ws://${host}${endpoint}`,
  })

  const client = new Client({
    url: `http://${host}${endpoint}`,
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
