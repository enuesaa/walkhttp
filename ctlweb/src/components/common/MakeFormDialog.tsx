import { Dialog } from '@radix-ui/themes'
import { MakeForm } from './MakeForm'
import { IoClose } from 'react-icons/io5'

export const MakeFormDialog = () => {
  return (
    <Dialog.Root>
      <Dialog.Trigger>
        <button className='inline-block h-full px-3 font-bold hover:bg-stone-700 rounded'>invoke</button>
      </Dialog.Trigger>

      <Dialog.Content className='relative'>
        <Dialog.Title>Invoke</Dialog.Title>
        <Dialog.Description></Dialog.Description>

        <MakeForm />

        <Dialog.Close className='absolute top-1 right-1'>
          <button className='p-3 text-2xl font-bold hover:bg-stone-700 rounded'><IoClose /></button>
        </Dialog.Close>

      </Dialog.Content>
    </Dialog.Root>
  )
}
