import { gql, useQuery } from 'urql'
import { Query } from './types'

const query = gql`
  query {
    appConfig {
      baseUrl
    }
  }
`

export const useGetAppConfig = () => useQuery<Query>({ query })[0]
