import { gql, useMutation } from 'urql'
import { Query } from './types'

const query = gql`
  mutation ($invocation: BrowserInvocationInput!) {
    makeBrowserInvocation(invocation: $invocation)
  }
`

export const useMakeBrowserInvocation = () => useMutation<Query>(query)[0]
