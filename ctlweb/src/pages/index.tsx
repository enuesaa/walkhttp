import { Header } from '@/components/common/Header'
import { Main } from '@/components/common/Main'
import { useMessgaes } from '@/lib/ws'

export default function TopPage() {
  const messages = useMessgaes()

  return (
    <>
      <Header />
      <Main>
        histories

        {messages.map((v, i) => <div key={i}>{v.method} {v.url}</div>)}
      </Main>
    </>
  )
}
