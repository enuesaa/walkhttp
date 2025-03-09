import { Invocation } from '@/gql/types'
import { calcDuration, fmtdatetime } from '@/lib/datefmt'
import { judgeStatusColor } from '@/lib/status'
import clsx from 'clsx'
import { HistoryDialogMetaCard } from './HistoryDialogMetaCard'

type Props = {
  invocation: Invocation
}
export const HistoryDialogMeta = ({ invocation }: Props) => {
  const started = fmtdatetime(invocation.started)
  const duration = calcDuration(invocation.started, invocation.received)
  const httpVersion = invocation.httpVersion === '' ? '-' : invocation.httpVersion
  const statusColor = judgeStatusColor(invocation.status)

  return (
    <>
      <section className='flex w-full text-grayer text-center text-normal h-11 leading-[44px]'>
        <div
          className={clsx(`px-5`, {
            'bg-green-800': statusColor === 'green',
            'bg-red-800': statusColor === 'red',
            'bg-blue-800': statusColor === 'blue',
          })}
        >
          {invocation.status}
        </div>
        <div className='bg-stone-900 w-14'>{invocation.method}</div>
        <div className='bg-stone-900 flex-auto px-2 text-left'>{invocation.url}</div>
      </section>

      <div className='py-2' style={{
        backgroundImage: 'radial-gradient(#57534e 1px, transparent 1px)',
        backgroundSize: '20px 20px',
      }}>
        <HistoryDialogMetaCard label='Started' value={started} />
        <HistoryDialogMetaCard label='Duration' value={`${duration.toString()} ms`} />
        <HistoryDialogMetaCard label='HTTP Version' value={httpVersion} />
      </div>
    </>
  )
}
