import { Invocation } from '@/gql/types'
import { Dialog } from '@radix-ui/themes'
import { IoClose } from 'react-icons/io5'
import { HistoryDialogMeta } from './HistoryDialogMeta'
import { fmtdate } from '@/lib/datefmt'
import { HistoryDialogHeaders } from './HistoryDialogHeaders'

type Props = {
  invocation: Invocation
}
export const HistoryDialog = ({ invocation }: Props) => {
  const started = fmtdate(invocation.started)

  return (
    <Dialog.Root>
      <Dialog.Trigger>
        <button className='flex cursor-pointer py-1 w-full text-grayer text-center text-base hover:bg-stone-700'>
          <div className='w-20'>{started}</div>
          <div className='bg-[green] w-10'>{invocation.status}</div>
          <div className='w-14'>{invocation.method}</div>
          <div className='flex-auto px-2 text-left'>{invocation.url}</div>
        </button>
      </Dialog.Trigger>

      <Dialog.Content maxWidth='1000px' aria-describedby={undefined} className='relative p-0 bg-stone-900'>
        <HistoryDialogMeta invocation={invocation} />
        <HistoryDialogHeaders invocation={invocation} />

        <Dialog.Close className='absolute top-0 right-0'>
          <button className='p-3 text-2xl font-bold hover:bg-stone-700'><IoClose /></button>
        </Dialog.Close>

      </Dialog.Content>
    </Dialog.Root>
  )
}
