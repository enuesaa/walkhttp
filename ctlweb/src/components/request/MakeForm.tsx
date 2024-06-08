import { Select, TextField, Button, SegmentedControl } from '@radix-ui/themes'
import styles from './MakeForm.css'
import { useMakeServerInvocation } from '@/graph/make-server-invocation'
import { useMakeBrowserInvocation } from '@/graph/make-browser-invocation'
import { FormEventHandler, MouseEventHandler, useState } from 'react'

export const MakeForm = () => {
  const [invoveServerData, invokeServer] = useMakeServerInvocation()
  const [invoveBrowserData, invokeBrowser] = useMakeBrowserInvocation()
  const [fromServer, setFromServer] = useState<boolean>(true)

  const handleSubmit: FormEventHandler<HTMLFormElement> = async (e) => {
    e.preventDefault()

    if (fromServer) {
      const invocation = {
        method: (e.currentTarget.elements.namedItem('method') as HTMLInputElement).value,
        url: (e.currentTarget.elements.namedItem('url') as HTMLInputElement).value,
        requestHeaders: [],
        requestBody: '',
      }
      await invokeServer({ invocation })
    } else {
      const invocation = {
        method: (e.currentTarget.elements.namedItem('method') as HTMLInputElement).value,
        url: (e.currentTarget.elements.namedItem('url') as HTMLInputElement).value,
        requestHeaders: [],
        requestBody: '',
      }
      // call here.
      await invokeBrowser({ invocation })
    }
  }

  const handleChangeFrom = (from: string) => {
    setFromServer(from === 'Server')
  }

  return (
    <form className={styles.main} onSubmit={handleSubmit}>
      <div className={styles.from}>
        <SegmentedControl.Root defaultValue='Server' size='3' radius='full' onValueChange={handleChangeFrom}>
          <SegmentedControl.Item value='Server'>Server</SegmentedControl.Item>
          <SegmentedControl.Item value='Browser'>Browser</SegmentedControl.Item>
        </SegmentedControl.Root>
      </div>

      <div className={styles.method}>
        <Select.Root defaultValue='GET' size='3' name='method'>
          <Select.Trigger />
          <Select.Content>
            <Select.Item value='GET'>GET</Select.Item>
            <Select.Item value='POST'>POST</Select.Item>
            <Select.Item value='PUT'>PUT</Select.Item>
            <Select.Item value='DELETE'>DELETE</Select.Item>
          </Select.Content>
        </Select.Root>
      </div>

      <div className={styles.url}>
        <TextField.Root placeholder='https://example.com/' size='3' name='url' />
      </div>

      <div className={styles.btn}>
        <Button>Call</Button>
      </div>
    </form>
  )
}
