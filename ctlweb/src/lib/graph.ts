import { gql, useQuery } from '@apollo/client'
import { Query } from './graphtypes'

const getInvocations = gql`
query {
  invocations(status: 404) {
    id
  }
}
`

export const useGetInvocations = () => useQuery<Query>(getInvocations)
