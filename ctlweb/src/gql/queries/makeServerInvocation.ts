import { gql, useMutation } from 'urql'
import { Query, ServerInvocationInput } from '@/gql/types'

const query = gql`
  mutation MakeServerInvocation($invocation: ServerInvocationInput!) {
    makeServerInvocation(input: $invocation)
  }
`

export const useMakeServerInvocation = () => useMutation<Query, { invocation: ServerInvocationInput }>(query)
