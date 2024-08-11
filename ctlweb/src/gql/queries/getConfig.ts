import { gql, useQuery } from 'urql'
import { Query } from '@/gql/types'

const query = gql`
  query GetConfig {
    getConfig {
      baseUrl
    }
  }
`

export const useGetConfig = () => useQuery<Query>({ query })[0]
