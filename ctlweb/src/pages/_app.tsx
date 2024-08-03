import type { AppProps } from 'next/app'
import '@radix-ui/themes/styles.css'
import { Theme } from '@radix-ui/themes'
import '@/styles/app.css'
import { Layout } from '@/components/common/Layout'
import { GraphQLProvider } from '@/gql/GraphQLProvider'

export default function App({ Component, pageProps }: AppProps) {
  return (
    <GraphQLProvider>
      <Theme appearance='dark' accentColor='cyan'>
        <Layout>
          <Component {...pageProps} />
        </Layout>
      </Theme>
    </GraphQLProvider>
  )
}
