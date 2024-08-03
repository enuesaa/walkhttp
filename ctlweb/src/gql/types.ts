export type Maybe<T> = T | null
export type InputMaybe<T> = Maybe<T>
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] }
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> }
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> }
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never }
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never }
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string }
  String: { input: string; output: string }
  Boolean: { input: boolean; output: boolean }
  Int: { input: number; output: number }
  Float: { input: number; output: number }
}

export type BrowserInvocationInput = {
  method: Scalars['String']['input']
  requestBody: Scalars['String']['input']
  requestHeaders?: InputMaybe<Array<InputMaybe<HeaderInput>>>
  responseBody: Scalars['String']['input']
  responseHeaders?: InputMaybe<Array<InputMaybe<HeaderInput>>>
  status: Scalars['Int']['input']
  url: Scalars['String']['input']
}

export type Config = {
  __typename?: 'Config'
  baseUrl: Scalars['String']['output']
}

export type Header = {
  __typename?: 'Header'
  name: Scalars['String']['output']
  value: Scalars['String']['output']
}

export type HeaderInput = {
  name: Scalars['String']['input']
  value: Scalars['String']['input']
}

export type Invocation = {
  __typename?: 'Invocation'
  created: Scalars['String']['output']
  id: Scalars['ID']['output']
  method: Scalars['String']['output']
  requestBody: Scalars['String']['output']
  requestHeaders?: Maybe<Array<Maybe<Header>>>
  responseBody: Scalars['String']['output']
  responseHeaders?: Maybe<Array<Maybe<Header>>>
  status: Scalars['Int']['output']
  url: Scalars['String']['output']
}

export type Mutation = {
  __typename?: 'Mutation'
  makeBrowserInvocation?: Maybe<Scalars['Boolean']['output']>
  makeServerInvocation?: Maybe<Scalars['Boolean']['output']>
}

export type MutationMakeBrowserInvocationArgs = {
  input: BrowserInvocationInput
}

export type MutationMakeServerInvocationArgs = {
  input: ServerInvocationInput
}

export type Query = {
  __typename?: 'Query'
  getConfig: Config
  getInvocation?: Maybe<Invocation>
  listInvocations: Array<Invocation>
}

export type QueryGetInvocationArgs = {
  id: Scalars['ID']['input']
}

export type ServerInvocationInput = {
  method: Scalars['String']['input']
  requestBody: Scalars['String']['input']
  requestHeaders?: InputMaybe<Array<InputMaybe<HeaderInput>>>
  url: Scalars['String']['input']
}

export type Subscription = {
  __typename?: 'Subscription'
  subscribeInvocations: Array<Invocation>
}
