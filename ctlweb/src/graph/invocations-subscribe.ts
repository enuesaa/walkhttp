import { useSubscription } from 'urql'
import { gql } from 'urql'
import { Query } from './types'

export const query = gql`
  subscription {
    invocations {
      id,
      url
    }
  }
`

export const useSubscribeInvocations = () => useSubscription<Query>({ query })[0]
