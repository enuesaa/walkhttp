import { Invocation } from '@/gql/types'
import { Dialog } from '@radix-ui/themes'
import { IoClose } from 'react-icons/io5'
import { HistoryDialogMeta } from './HistoryDialogMeta'
import { fmtdate } from '@/lib/datefmt'
import { HistoryDialogHeader } from './HistoryDialogHeader'
import { HistoryDialogBody } from './HistoryDialogBody'
import { judgeStatusColor } from '@/lib/status'
import clsx from 'clsx'

type Props = {
  invocation: Invocation
}
export const HistoryDialog = ({ invocation }: Props) => {
  const started = fmtdate(invocation.started)
  const statusColor = judgeStatusColor(invocation.status)

  return (
    <Dialog.Root>
      <Dialog.Trigger>
        <button className='flex cursor-pointer py-1 w-full text-grayer text-center text-base hover:bg-stone-700'>
          <div className='w-20'>{started}</div>
          <div
            className={clsx(`w-10`, {
              'bg-green-800': statusColor === 'green',
              'bg-red-800': statusColor === 'red',
              'bg-blue-800': statusColor === 'blue',
            })}
          >
            {invocation.status}
          </div>
          <div className='w-14'>{invocation.method}</div>
          <div className='flex-auto px-2 text-left'>{invocation.url}</div>
        </button>
      </Dialog.Trigger>

      <Dialog.Content maxWidth='1400px' aria-describedby={undefined} className='relative p-0 bg-stone-800'>
        <Dialog.Title className='p-0 m-0' />
        <HistoryDialogMeta invocation={invocation} />
        <HistoryDialogHeader invocation={invocation} />
        <HistoryDialogBody invocation={invocation} />

        <Dialog.Close className='absolute top-0 right-0'>
          <button className='px-3 py-2 text-2xl font-bold hover:bg-stone-700'>
            <IoClose />
          </button>
        </Dialog.Close>
      </Dialog.Content>
    </Dialog.Root>
  )
}
