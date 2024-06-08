import { gql, useMutation } from 'urql'
import { Query } from './types'

const query = gql`
  mutation ($invocation: ServerInvocationInput!) {
    makeServerInvocation(invocation: $invocation)
  }
`

export const useMakeServerInvocation = () => useMutation<Query>(query)
