import { useSubscription } from 'urql'
import { gql } from 'urql'
import { Query } from '@/gql/types'

export const query = gql`
  subscription SubscribeInvocations {
    invocations {
      id
      status
      method
      url
      requestBody
      responseBody
      created
    }
  }
`

export const useSubscribeInvocations = () => useSubscription<Query>({ query })[0]
