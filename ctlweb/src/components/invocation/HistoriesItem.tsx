import { Invocation } from '@/gql/types'
import { Badge, Button, Dialog } from '@radix-ui/themes'
import { History } from './History'

type Props = {
  invocation: Invocation
}
export const HistroiesItem = ({ invocation }: Props) => {
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
        <Button variant='ghost' className='my-2 relative cursor-pointer block py-1 px-5 w-full border border-white/30 text-left'>
          <div className='text-white/30 inline-block mx-2'>
            {started}
          </div>
          <div className='text-[green] inline-block mx-2'>{invocation.method}</div>
          {invocation.url}
        </Button>
      </Dialog.Trigger>

      <Dialog.Content maxWidth='700px' aria-describedby={undefined} className='relative p-8'>
        <Dialog.Title size='5'>Invocation</Dialog.Title>
        <Dialog.Close>
          <Button variant='soft' color='gray' className='absolute top-5 right-5'>
            Close
          </Button>
        </Dialog.Close>
        <History invocation={invocation} />
      </Dialog.Content>
    </Dialog.Root>
  )
}
