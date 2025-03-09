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
        <div className='bg-stone-800 rounded-lg w-fit min-w-24 h-[50px] mx-3 px-1 text-sm border border-stone-600 text-stone-400'>
          <span className='font-semibold'>Started</span>
          <div className='w-full text-center text-stone-300 mt-[2px]'>
            {started}
          </div>
        </div>
      </div>
    </>
  )
}
