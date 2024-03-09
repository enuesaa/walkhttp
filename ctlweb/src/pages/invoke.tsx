import { Header } from '@/components/common/Header'
import { Main } from '@/components/common/Main'
import { Invocations } from '@/components/invocation/Invocations'

export default function Page() {
  return (
    <>
      <Header />
      <Main>
        <Invocations />
      </Main>
    </>
  )
}
