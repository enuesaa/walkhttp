import { RouterProvider } from 'react-router-dom'
import { createRouter } from './routes'
import { ApolloClient, InMemoryCache, ApolloProvider, gql } from '@apollo/client'
import './styles/globals.css'

export const App = () => {
  const client = new ApolloClient({
    uri: 'http://localhost:3000/query/',
    cache: new InMemoryCache(),
  })
  const router = createRouter()

  return (
    <ApolloProvider client={client}>
      <RouterProvider router={router} />
    </ApolloProvider>
  )
}
