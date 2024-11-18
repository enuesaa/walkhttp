import { Invocation } from '@/gql/types'
import { fmtdate } from '@/lib/datefmt'
import { judgeStatusColor } from '@/lib/status'
import clsx from 'clsx'

type Props = {
  invocation: Invocation
}
export const HistoryDialogMeta = ({ invocation }: Props) => {
  const started = fmtdate(invocation.started)
  const statusColor = judgeStatusColor(invocation.status)

  return (
    <>
      <section className='flex w-full text-grayer text-center text-normal h-11 leading-[44px]'>
        <div className={clsx(`px-5`, {
            'bg-green-800': statusColor === 'green',
            'bg-red-800': statusColor === 'red',
            'bg-blue-800': statusColor === 'blue',
          })}>{invocation.status}</div>
        <div className='bg-stone-900 w-14'>{invocation.method}</div>
        <div className='bg-stone-900 flex-auto px-2 text-left'>{invocation.url}</div>
      </section>

      <div className='text-sm text-stone-400 absolute'>
        {started}
      </div>
    </>
  )
}
