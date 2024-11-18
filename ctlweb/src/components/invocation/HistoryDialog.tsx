import { Invocation } from '@/gql/types'
import { Dialog } from '@radix-ui/themes'
import { History } from './History'
import { IoClose } from 'react-icons/io5'

type Props = {
  invocation: Invocation
}
export const HistoryDialog = ({ invocation }: Props) => {
  const datefmt = new Intl.DateTimeFormat('en-US', {
    hour: 'numeric',
    hour12: false,
    minute: 'numeric',
    second: 'numeric',
  })
  const started = datefmt.format(new Date(invocation.started))

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

      <Dialog.Content maxWidth='1000px' aria-describedby={undefined} className='relative px-8 py-2'>
        <Dialog.Close className='absolute top-1 right-1'>
          <button className='p-3 text-2xl font-bold hover:bg-gray rounded'><IoClose /></button>
        </Dialog.Close>

        <History invocation={invocation} />
      </Dialog.Content>
    </Dialog.Root>
  )
}
