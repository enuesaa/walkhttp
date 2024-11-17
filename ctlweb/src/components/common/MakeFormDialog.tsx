import { Button, Dialog } from '@radix-ui/themes'
import { MakeForm } from './MakeForm'
import { IoClose } from 'react-icons/io5'

export const MakeFormDialog = () => {
  return (
    <Dialog.Root>
      <Dialog.Trigger>
        <Button variant='outline' className='m-2'>Invoke</Button>
      </Dialog.Trigger>

      <Dialog.Content className='relative'>
        <Dialog.Title>Invoke</Dialog.Title>
        <Dialog.Description></Dialog.Description>

        <MakeForm />

        <Dialog.Close className='absolute top-1 right-1'>
          <Button variant='outline'><IoClose /></Button>
        </Dialog.Close>

      </Dialog.Content>
    </Dialog.Root>
  )
}
