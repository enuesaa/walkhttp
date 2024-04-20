import type { AppProps } from 'next/app'
import '@radix-ui/themes/styles.css'
import { Theme } from '@radix-ui/themes'
import { ApolloClient, InMemoryCache, ApolloProvider } from '@apollo/client'
import '@/styles/app.css'

export default function App({ Component, pageProps }: AppProps) {
  const apibase = process.env.NEXT_PUBLIC_API_BASE
  const queryClient = new ApolloClient({
    uri: `${apibase}graph`,
    cache: new InMemoryCache(),
  }) 

  return (
    <ApolloProvider client={queryClient}>
      <Theme appearance='dark'>
        <Component {...pageProps} />
      </Theme>
    </ApolloProvider>
  )
}
