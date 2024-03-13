import { Header } from '@/components/common/Header'
import { Main } from '@/components/common/Main'
import { useMessgaes } from '@/lib/ws'
import { HistoriesTable } from '@/components/invocation/Histories'

export default function TopPage() {
  const messages = useMessgaes()

  return (
    <>
      <Header />
      <Main>
        histories

        <HistoriesTable messages={messages} />
      </Main>
    </>
  )
}
