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
      <div className={styles.from}>
        <SegmentedControl.Root defaultValue='Server' size='3' radius='full'>
          <SegmentedControl.Item value='Server'>Server</SegmentedControl.Item>
          <SegmentedControl.Item value='Browser'>Browser</SegmentedControl.Item>
        </SegmentedControl.Root>
      </div>

      <div className={styles.method}>
        <Select.Root defaultValue='GET' size='3'>
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
        <TextField.Root placeholder='https://example.com/' size='3' />
      </div>

      <div className={styles.btn}>
        <Button onClick={handleClick}>Call</Button>
      </div>
    </section>
  )
}
