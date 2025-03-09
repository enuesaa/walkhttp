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
      httpVersion
    }
  }
`

export const useListInvocations = () => useQuery<Query>({ query })[0]
