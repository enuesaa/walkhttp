import { RouterProvider } from 'react-router-dom'
import { createRouter } from './routes'
import { ApolloClient, InMemoryCache, ApolloProvider, gql } from '@apollo/client'
import { ChakraProvider } from '@chakra-ui/react'
import { Global } from '@emotion/react'
import { globalStyle } from './styles/global'
import './styles/globals.css'

export const App = () => {
  const client = new ApolloClient({
    uri: 'http://localhost:3000/query/',
    cache: new InMemoryCache(),
  })
  const router = createRouter()

  return (
    <ApolloProvider client={client}>
      <Global styles={globalStyle} />
      <ChakraProvider>
        <RouterProvider router={router} />
      </ChakraProvider>
    </ApolloProvider>
  )
}
