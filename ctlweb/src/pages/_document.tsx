import { Html, Head, Main, NextScript } from 'next/document'

export default function Document() {
  return (
    <Html lang='ja'>
      <Head>
        <link rel='icon' href='/_/favicon.ico' sizes='any' />
      </Head>
      <body>
        <Main />
        <NextScript />
      </body>
    </Html>
  )
}
