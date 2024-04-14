import { gql, useQuery } from '@apollo/client'

const getInvocations = gql`
query {
  invocations(status: 404) {
    id
  }
}
`

export const useGetInvocations = () => useQuery(getInvocations)
