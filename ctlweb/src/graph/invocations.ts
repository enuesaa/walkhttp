import { gql, useQuery } from 'urql'
import { Query } from './types'

const query = gql`
  query {
    invocations {
      id,
      url,
      method
    }
  }
`

export const useListInvocations = () => useQuery<Query>({ query })[0]
