import { gql, useMutation } from 'urql'
import { BrowserInvocationInput, Mutation } from '@/gql/types'

const query = gql`
  mutation MakeBrowserInvocation($invocation: BrowserInvocationInput!) {
    makeBrowserInvocation(input: $invocation)
  }
`

export const useMakeBrowserInvocation = () => useMutation<Mutation, { invocation: BrowserInvocationInput }>(query)
