import { Select, TextField, Button, SegmentedControl } from '@radix-ui/themes'
import styles from './MakeForm.css'
import { useMakeServerInvocation } from '@/graph/make-server-invocation'
import { MouseEventHandler } from 'react'

export const MakeForm = () => {
  const [data, invoke] = useMakeServerInvocation()

  const handleClick: MouseEventHandler<HTMLButtonElement> = async (e) => {
    e.preventDefault()
    await invoke({
      invocation: {
        method: 'GET',
        url: 'https://example.com/',
        requestHeaders: [],
        requestBody: '',
      },
    })
  }

  return (
    <section className={styles.main}>

      <div className={styles.requestFrom}>
        Request from
        <SegmentedControl.Root defaultValue='Server'>
          <SegmentedControl.Item value='Server'>Server</SegmentedControl.Item>
          <SegmentedControl.Item value='Browser'>Browser</SegmentedControl.Item>
        </SegmentedControl.Root>
      </div>

      <Select.Root defaultValue='GET'>
        <Select.Trigger />
        <Select.Content>
          <Select.Item value='GET'>GET</Select.Item>
          <Select.Item value='POST'>POST</Select.Item>
          <Select.Item value='PUT'>PUT</Select.Item>
          <Select.Item value='DELETE'>DELETE</Select.Item>
        </Select.Content>
      </Select.Root>

      <TextField.Root placeholder='https://example.com/' />
      <Button onClick={handleClick}>call</Button>
    </section>
  )
}
