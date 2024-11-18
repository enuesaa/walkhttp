import { Invocation } from '@/gql/types'
import { HistoryHeaders } from './HistoryHeaders'
import { HistoryBody } from './HistoryBody'
import { HistorySectionTitle } from './HistorySectionTitle'
import { fmtdate } from '@/lib/datefmt'

type Props = {
  invocation: Invocation
}
export const History = ({ invocation }: Props) => {
  const started = fmtdate(invocation.started)

  return (
    <>
      <section className='flex w-full text-grayer text-center text-normal h-11 leading-[44px]'>
        <div className='bg-[green] px-5'>{invocation.status}</div>
        <div className='w-14'>{invocation.method}</div>
        <div className='flex-auto px-2 text-left'>{invocation.url}</div>
      </section>

      <ul className='text-sm text-stone-400'>
        <li>started: {started}</li>
      </ul>

      <HistorySectionTitle title='Request Headers' />
      <HistoryHeaders headers={invocation.requestHeaders} />

      <HistorySectionTitle title='Request Body' />
      <HistoryBody body={invocation.requestBody} />

      <HistorySectionTitle title='Response Headers' />
      <HistoryHeaders headers={invocation.responseHeaders} />

      <HistorySectionTitle title='Response Body' />
      <HistoryBody body={invocation.responseBody} />
    </>
  )
}
