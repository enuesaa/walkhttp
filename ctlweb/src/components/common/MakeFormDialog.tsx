import { Dialog } from '@radix-ui/themes'
import { MakeForm } from './MakeForm'
import { IoClose } from 'react-icons/io5'

export const MakeFormDialog = () => {
  return (
    <Dialog.Root>
      <Dialog.Trigger>
        <button className='inline-block h-full px-3 font-bold hover:bg-stone-700'>invoke</button>
      </Dialog.Trigger>

      <Dialog.Content className='relative bg-stone-800'>
        <MakeForm />

        <Dialog.Close className='absolute top-0 right-0'>
          <button className='p-3 text-2xl font-bold hover:bg-stone-700'><IoClose /></button>
        </Dialog.Close>

      </Dialog.Content>
    </Dialog.Root>
  )
}
