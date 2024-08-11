import { gql, useQuery } from 'urql'
import { Query } from '@/gql/types'

const query = gql`
  query ListInvocations {
    listInvocations {
      id
      status
      method
      url
      requestBody
      responseBody
      started
      received
    }
  }
`

export const useListInvocations = () => useQuery<Query>({ query })[0]
