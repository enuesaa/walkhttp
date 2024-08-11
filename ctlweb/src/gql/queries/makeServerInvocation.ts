import { gql, useMutation } from 'urql'
import { Mutation, ServerInvocationInput } from '@/gql/types'

const query = gql`
  mutation MakeServerInvocation($invocation: ServerInvocationInput!) {
    makeServerInvocation(input: $invocation)
  }
`

export const useMakeServerInvocation = () => useMutation<Mutation, { invocation: ServerInvocationInput }>(query)
