import type { AppProps } from 'next/app'
import { Theme } from '@radix-ui/themes'
import '../styles/app.css'
import { Layout } from '@/components/common/Layout'
import { GraphQLProvider } from '@/gql/GraphQLProvider'
import Head from 'next/head'

export default function App({ Component, pageProps }: AppProps) {
  return (
    <>
      <Head>
        <title>walkhttp</title>
      </Head>
      <GraphQLProvider>
        <Theme appearance='dark' className='bg-stone-900 text-stone-300'>
          <Layout>
            <Component {...pageProps} />
          </Layout>
        </Theme>
      </GraphQLProvider>
    </>
  )
}
