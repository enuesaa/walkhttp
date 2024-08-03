import { gql, useMutation } from 'urql'
import { Query, ServerInvocationInput } from '@/gql/types'

const query = gql`
  mutation ($invocation: ServerInvocationInput!) {
    makeServerInvocation(invocation: $invocation)
  }
`

export const useMakeServerInvocation = () => useMutation<Query, { invocation: ServerInvocationInput }>(query)
