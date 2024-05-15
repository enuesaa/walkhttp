import { useSubscription } from 'urql'
import { gql } from 'urql'
import { Query } from './types'

export const query = gql`
  subscription {
    invocations {
      id
      status
      method
      url
      requestBody
      responseBody
    }
  }
`

export const useSubscribeInvocations = () => useSubscription<Query>({ query })[0]
