export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
};

export type Header = {
  __typename?: 'Header';
  name: Scalars['String']['output'];
  value: Scalars['String']['output'];
};

export type Invocation = {
  __typename?: 'Invocation';
  id: Scalars['ID']['output'];
  method: Scalars['String']['output'];
  requestBody?: Maybe<Scalars['String']['output']>;
  requestHeaders?: Maybe<Array<Maybe<Header>>>;
  responseBody?: Maybe<Scalars['String']['output']>;
  responseHeaders?: Maybe<Array<Maybe<Header>>>;
  status: Scalars['Int']['output'];
  url: Scalars['String']['output'];
};

export type Query = {
  __typename?: 'Query';
  invocation?: Maybe<Invocation>;
  invocations: Array<Invocation>;
};


export type QueryInvocationArgs = {
  id: Scalars['ID']['input'];
};

export type Subscription = {
  __typename?: 'Subscription';
  invocations: Array<Invocation>;
};
