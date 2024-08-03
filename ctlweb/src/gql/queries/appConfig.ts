import { gql, useQuery } from 'urql'
import { Query } from '@/gql/types'

const query = gql`
  query GetAppConfig {
    getConfig {
      baseUrl
    }
  }
`

export const useGetAppConfig = () => useQuery<Query>({ query })[0]
