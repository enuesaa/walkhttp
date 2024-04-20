import { gql, useQuery } from '@apollo/client'
import { Query } from './graphtypes'

const listInvocations = gql`
query {
  invocations(status: 200) {
    id,
    url,
    method
  }
}
`

export const useListInvocations = () => useQuery<Query>(listInvocations)
