import type { AppProps } from 'next/app'
import '@radix-ui/themes/styles.css'
import { Theme } from '@radix-ui/themes'
import '@/styles/app.css'
import { Layout } from '@/components/common/Layout'
import { GraphProvider } from '@/graph/GraphProvider'

export default function App({ Component, pageProps }: AppProps) {
  return (
    <GraphProvider>
      <Theme appearance='dark'>
        <Layout>
          <Component {...pageProps} />
        </Layout>
      </Theme>
    </GraphProvider>
  )
}
