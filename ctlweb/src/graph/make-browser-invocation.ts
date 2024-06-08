import { gql, useMutation } from 'urql'
import { BrowserInvocationInput, Query } from './types'

const query = gql`
  mutation ($invocation: BrowserInvocationInput!) {
    makeBrowserInvocation(invocation: $invocation)
  }
`

export const useMakeBrowserInvocation = () => useMutation<Query, {invocation: BrowserInvocationInput}>(query)
