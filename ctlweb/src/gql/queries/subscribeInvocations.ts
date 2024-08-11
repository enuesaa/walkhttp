import { useSubscription } from 'urql'
import { gql } from 'urql'
import { Subscription } from '@/gql/types'

export const query = gql`
  subscription SubscribeInvocations {
    subscribeInvocations {
      id
      status
      method
      url
      requestBody
      requestHeaders {
        name
        value
      }
      responseBody
      responseHeaders {
        name
        value
      }
      started
      received
    }
  }
`

export const useSubscribeInvocations = () => useSubscription<Subscription>({ query })[0]
