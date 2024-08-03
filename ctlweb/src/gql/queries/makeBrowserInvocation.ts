import { gql, useMutation } from 'urql'
import { BrowserInvocationInput, Query } from '@/gql/types'

const query = gql`
  mutation MakeBrowserInvocation($invocation: BrowserInvocationInput!) {
    makeBrowserInvocation(input: $invocation)
  }
`

export const useMakeBrowserInvocation = () => useMutation<Query, { invocation: BrowserInvocationInput }>(query)
