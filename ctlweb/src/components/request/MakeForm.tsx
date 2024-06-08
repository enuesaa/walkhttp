import { Select, TextField, Button, SegmentedControl } from '@radix-ui/themes'
import styles from './MakeForm.css'
import { useMakeServerInvocation } from '@/graph/make-server-invocation'
import { useMakeBrowserInvocation } from '@/graph/make-browser-invocation'
import { FormEventHandler, MouseEventHandler } from 'react'

export const MakeForm = () => {
  const [invoveServerData, invokeServer] = useMakeServerInvocation()
  const [invoveBrowserData, invokeBrowser] = useMakeBrowserInvocation()

  const handleSubmit: FormEventHandler<HTMLFormElement> = async (e) => {
    e.preventDefault()
    const invocation = {
      method: (e.currentTarget.elements.namedItem('method') as HTMLInputElement).value,
      url: (e.currentTarget.elements.namedItem('url') as HTMLInputElement).value,
      requestHeaders: [],
      requestBody: '',
    }
    await invokeServer({ invocation })
  }

  return (
    <form className={styles.main} onSubmit={handleSubmit}>
      <div className={styles.from}>
        <SegmentedControl.Root defaultValue='Server' size='3' radius='full'>
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
