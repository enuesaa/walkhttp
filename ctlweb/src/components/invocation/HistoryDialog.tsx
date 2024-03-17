import { Invocation } from '@/lib/api'
import { Button, Code, Dialog, Heading } from '@radix-ui/themes'

type Props = {
  invocation: Invocation
}
export const HistoryDialog = ({ invocation }: Props) => {
  return (
    <Dialog.Root>
      <Dialog.Trigger>
        <Button>view</Button>
      </Dialog.Trigger>

      <Dialog.Content style={{ maxWidth: 450 }}>
        <Dialog.Title>{invocation.id}</Dialog.Title>
        
        <Heading>Request Headers</Heading>
        <Code style={{display: 'block'}}>
          {JSON.stringify(invocation.requestHeaders, null, '  ')}
        </Code>

        <Heading>Response Headers</Heading>
        <Code style={{display: 'block'}}>
          {JSON.stringify(invocation.responseHeaders, null, '  ')}
        </Code>

        <Heading>Response Body</Heading>
        <Code style={{display: 'block'}}>
          {invocation.responseBody}
        </Code>

        <Dialog.Close>
          <Button variant='soft'>Close</Button>
        </Dialog.Close>
      </Dialog.Content>
    </Dialog.Root>
  )
}
